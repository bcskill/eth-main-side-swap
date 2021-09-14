package swap

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcom "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"

	sabi "github.com/bcskill/eth-main-side-swap/abi"
	"github.com/bcskill/eth-main-side-swap/common"
	"github.com/bcskill/eth-main-side-swap/model"
	"github.com/bcskill/eth-main-side-swap/util"
)

// NewSwapEngine returns the swapEngine instance
func NewSwapEngine(db *gorm.DB, cfg *util.Config, sideClient, mainClient *ethclient.Client) (*SwapEngine, error) {
	pairs := make([]model.SwapPair, 0)
	db.Find(&pairs)

	swapPairInstances, err := buildSwapPairInstance(pairs)
	if err != nil {
		return nil, err
	}
	sideContractAddrToMainContractAddr := make(map[ethcom.Address]ethcom.Address)
	mainContractAddrToSideContractAddr := make(map[ethcom.Address]ethcom.Address)
	for _, token := range pairs {
		mainContractAddrToSideContractAddr[ethcom.HexToAddress(token.MainChainErc20Addr)] = ethcom.HexToAddress(token.SideChainErc20Addr)
		sideContractAddrToMainContractAddr[ethcom.HexToAddress(token.SideChainErc20Addr)] = ethcom.HexToAddress(token.MainChainErc20Addr)
	}

	keyConfig, err := GetKeyConfig(cfg)
	if err != nil {
		return nil, err
	}

	sidePrivateKey, _, err := BuildKeys(keyConfig.SidePrivateKey)
	if err != nil {
		return nil, err
	}

	mainPrivateKey, _, err := BuildKeys(keyConfig.MainPrivateKey)
	if err != nil {
		return nil, err
	}

	sideChainID, err := sideClient.ChainID(context.Background())
	if err != nil {
		return nil, err

	}
	mainChainID, err := mainClient.ChainID(context.Background())
	if err != nil {
		return nil, err
	}
	MainSwapAgentABI, err := abi.JSON(strings.NewReader(sabi.MainSwapAgentABI))
	if err != nil {
		return nil, err
	}

	SideSwapAgentABI, err := abi.JSON(strings.NewReader(sabi.SideSwapAgentABI))
	if err != nil {
		return nil, err
	}

	swapEngine := &SwapEngine{
		db:                      db,
		config:                  cfg,
		hmacCKey:                keyConfig.HMACKey,
		mainPrivateKey:          mainPrivateKey,
		sidePrivateKey:          sidePrivateKey,
		sideClient:              sideClient,
		mainClient:              mainClient,
		sideChainID:             sideChainID.Int64(),
		mainChainID:             mainChainID.Int64(),
		swapPairsFromERC20Addr:  swapPairInstances,
		side20ToMain20:          sideContractAddrToMainContractAddr,
		main20ToSide20:          mainContractAddrToSideContractAddr,
		mainSwapAgentABI:        &MainSwapAgentABI,
		sideSwapAgentABI:        &SideSwapAgentABI,
		mainSwapAgent:           ethcom.HexToAddress(cfg.ChainConfig.MainSwapAgentAddr),
		sideSwapAgent:           ethcom.HexToAddress(cfg.ChainConfig.SideSwapAgentAddr),
	}

	return swapEngine, nil
}

func (engine *SwapEngine) Start() {
	go engine.monitorSwapRequestDaemon()
	go engine.confirmSwapRequestDaemon()
	go engine.swapInstanceDaemon(SwapMain2Side)
	go engine.swapInstanceDaemon(SwapSide2Main)
	go engine.trackSwapTxDaemon()
	go engine.retryFailedSwapsDaemon()
	go engine.trackRetrySwapTxDaemon()
}

func (engine *SwapEngine) monitorSwapRequestDaemon() {
	for {
		swapStartTxLogs := make([]model.SwapStartTxLog, 0)
		engine.db.Where("phase = ?", model.SeenRequest).Order("height asc").Limit(BatchSize).Find(&swapStartTxLogs)

		if len(swapStartTxLogs) == 0 {
			time.Sleep(SleepTime * time.Second)
			continue
		}

		for _, swapEventLog := range swapStartTxLogs {
			swap := engine.createSwap(&swapEventLog)
			writeDBErr := func() error {
				tx := engine.db.Begin()
				if err := tx.Error; err != nil {
					return err
				}
				if err := engine.insertSwap(tx, swap); err != nil {
					tx.Rollback()
					return err
				}
				tx.Model(model.SwapStartTxLog{}).Where("tx_hash = ?", swap.StartTxHash).Updates(
					map[string]interface{}{
						"phase":       model.ConfirmRequest,
						"update_time": time.Now().Unix(),
					})
				return tx.Commit().Error
			}()

			if writeDBErr != nil {
				util.Logger.Errorf("write db error: %s", writeDBErr.Error())
				util.SendTelegramMessage(fmt.Sprintf("write db error: %s", writeDBErr.Error()))
			}
		}
	}
}

func (engine *SwapEngine) getSwapHMAC(swap *model.Swap) string {
	material := fmt.Sprintf("%s#%s#%s#%s#%s#%s#%s#%d#%s#%s#%s",
		swap.Status, swap.Sponsor, swap.SourceChainErc20Addr, swap.TargetChainErc20Addr, swap.TargetChainToAddr, swap.Symbol, swap.Amount, swap.Decimals, swap.Direction, swap.StartTxHash, swap.FillTxHash)
	mac := hmac.New(sha256.New, []byte(engine.hmacCKey))
	mac.Write([]byte(material))

	return hex.EncodeToString(mac.Sum(nil))
}

func (engine *SwapEngine) verifySwap(swap *model.Swap) bool {
	return swap.RecordHash == engine.getSwapHMAC(swap)
}

func (engine *SwapEngine) insertSwap(tx *gorm.DB, swap *model.Swap) error {
	swap.RecordHash = engine.getSwapHMAC(swap)
	return tx.Create(swap).Error
}

func (engine *SwapEngine) updateSwap(tx *gorm.DB, swap *model.Swap) {
	swap.RecordHash = engine.getSwapHMAC(swap)
	tx.Save(swap)
}

func (engine *SwapEngine) createSwap(txEventLog *model.SwapStartTxLog) *model.Swap {
	sponsor := txEventLog.Sponsor
	amount := txEventLog.Amount
	swapStartTxHash := txEventLog.TxHash
	swapDirection := SwapMain2Side
	if txEventLog.Chain == common.ChainSide {
		swapDirection = SwapSide2Main
	}

	var mainSourceErc20Addr ethcom.Address
	var sideSourceErc20Addr ethcom.Address
	var targetChainErc20Addr ethcom.Address

	var ok bool
	decimals := 0
	var symbol string
	swapStatus := SwapQuoteRejected
	err := func() error {
		if txEventLog.Chain == common.ChainMain {
			mainSourceErc20Addr = ethcom.HexToAddress(txEventLog.SourceChainErc20Addr)
			if sideSourceErc20Addr, ok = engine.main20ToSide20[ethcom.HexToAddress(txEventLog.SourceChainErc20Addr)]; !ok {
				return fmt.Errorf("unsupported main token contract address: %s", txEventLog.SourceChainErc20Addr)
			}
			targetChainErc20Addr = sideSourceErc20Addr
		} else {
			sideSourceErc20Addr = ethcom.HexToAddress(txEventLog.SourceChainErc20Addr)
			if mainSourceErc20Addr, ok = engine.side20ToMain20[ethcom.HexToAddress(txEventLog.SourceChainErc20Addr)]; !ok {
				return fmt.Errorf("unsupported side token contract address: %s", txEventLog.SourceChainErc20Addr)
			}
			targetChainErc20Addr = mainSourceErc20Addr
		}
		pairInstance, err := engine.GetSwapPairInstance(mainSourceErc20Addr)
		if err != nil {
			return fmt.Errorf("failed to get swap pair for target erc20 %s, error %s", targetChainErc20Addr.String(), err.Error())
		}
		decimals = pairInstance.Decimals
		symbol = pairInstance.Symbol
		swapAmount := big.NewInt(0)
		_, ok = swapAmount.SetString(txEventLog.Amount, 10)
		if !ok {
			return fmt.Errorf("unrecongnized swap amount: %s", txEventLog.Amount)
		}
		if swapAmount.Cmp(pairInstance.LowBound) < 0 || swapAmount.Cmp(pairInstance.UpperBound) > 0 {
			return fmt.Errorf("swap amount is out of bound, expected bound [%s, %s]", pairInstance.LowBound.String(), pairInstance.UpperBound.String())
		}

		swapStatus = SwapTokenReceived
		return nil
	}()

	log := ""
	if err != nil {
		log = err.Error()
	}

	swap := &model.Swap{
		Status:      swapStatus,
		Sponsor:     sponsor,
		SourceChainErc20Addr:   ethcom.HexToAddress(txEventLog.SourceChainErc20Addr).String(),
		TargetChainErc20Addr:   targetChainErc20Addr.String(),
		TargetChainToAddr:      ethcom.HexToAddress(txEventLog.TargetChainToAddr).String(),
		Symbol:      symbol,
		Amount:      amount,
		Decimals:    decimals,
		Direction:   swapDirection,
		StartTxHash: swapStartTxHash,
		FillTxHash:  "",
		Log:         log,
	}

	return swap
}

func (engine *SwapEngine) confirmSwapRequestDaemon() {
	for {
		txEventLogs := make([]model.SwapStartTxLog, 0)
		engine.db.Where("status = ? and phase = ?", model.TxStatusConfirmed, model.ConfirmRequest).
			Order("height asc").Limit(BatchSize).Find(&txEventLogs)

		if len(txEventLogs) == 0 {
			time.Sleep(SleepTime * time.Second)
			continue
		}

		util.Logger.Debugf("found %d confirmed event logs", len(txEventLogs))

		for _, txEventLog := range txEventLogs {
			writeDBErr := func() error {
				tx := engine.db.Begin()
				if err := tx.Error; err != nil {
					return err
				}
				swap, err := engine.getSwapByStartTxHash(tx, txEventLog.TxHash)
				if err != nil {
					util.Logger.Errorf("verify hmac of swap failed: %s", txEventLog.TxHash)
					util.SendTelegramMessage(fmt.Sprintf("Urgent alert: verify hmac of swap failed: %s", txEventLog.TxHash))
					return err
				}

				if swap.Status == SwapTokenReceived {
					swap.Status = SwapConfirmed
					engine.updateSwap(tx, swap)
				}

				tx.Model(model.SwapStartTxLog{}).Where("id = ?", txEventLog.Id).Updates(
					map[string]interface{}{
						"phase":       model.AckRequest,
						"update_time": time.Now().Unix(),
					})
				return tx.Commit().Error
			}()

			if writeDBErr != nil {
				util.Logger.Errorf("write db error: %s", writeDBErr.Error())
				util.SendTelegramMessage(fmt.Sprintf("write db error: %s", writeDBErr.Error()))
			}
		}
	}
}

func (engine *SwapEngine) swapInstanceDaemon(direction common.SwapDirection) {
	util.Logger.Infof("start swap daemon, direction %s", direction)
	for {

		swaps := make([]model.Swap, 0)
		engine.db.Where("status in (?) and direction = ?", []common.SwapStatus{SwapConfirmed, SwapSending}, direction).Order("id asc").Limit(BatchSize).Find(&swaps)

		if len(swaps) == 0 {
			time.Sleep(SwapSleepSecond * time.Second)
			continue
		}

		util.Logger.Debugf("found %d confirmed swap requests", len(swaps))

		for _, swap := range swaps {
			var swapPairInstance *SwapPairIns
			var mainSourceErc20Addr ethcom.Address
			var err error
			retryCheckErr := func() error {
				if !engine.verifySwap(&swap) {
					return fmt.Errorf("verify hmac of swap failed: %s", swap.StartTxHash)
				}

				if swap.Direction == SwapMain2Side {
					mainSourceErc20Addr = ethcom.HexToAddress(swap.SourceChainErc20Addr)
				} else {
					mainSourceErc20Addr = ethcom.HexToAddress(swap.TargetChainErc20Addr)
				}

				swapPairInstance, err = engine.GetSwapPairInstance(mainSourceErc20Addr)
				if err != nil {
					return fmt.Errorf("swap instance for target erc20 %s doesn't exist, skip this swap", swap.TargetChainErc20Addr)
				}
				return nil
			}()
			if retryCheckErr != nil {
				writeDBErr := func() error {
					tx := engine.db.Begin()
					if err := tx.Error; err != nil {
						return err
					}
					swap.Status = SwapQuoteRejected
					swap.Log = retryCheckErr.Error()
					engine.updateSwap(tx, &swap)
					return tx.Commit().Error
				}()
				if writeDBErr != nil {
					util.Logger.Errorf("write db error: %s", writeDBErr.Error())
					util.SendTelegramMessage(fmt.Sprintf("write db error: %s", writeDBErr.Error()))
				}
				continue
			}

			skip, writeDBErr := func() (bool, error) {
				isSkip := false
				tx := engine.db.Begin()
				if err := tx.Error; err != nil {
					return false, err
				}
				if swap.Status == SwapSending {
					var swapTx model.SwapFillTx
					engine.db.Where("start_swap_tx_hash = ?", swap.StartTxHash).First(&swapTx)
					if swapTx.FillSwapTxHash == "" {
						util.Logger.Infof("retry swap, start tx hash %s, symbol %s, amount %s, direction %s",
							swap.StartTxHash, swap.Symbol, swap.Amount, swap.Direction)
						swap.Status = SwapConfirmed
						engine.updateSwap(tx, &swap)
					} else {
						util.Logger.Infof("swap tx is built successfully, but the swap tx status is uncertain, just mark the swap and swap tx status as sent, swap ID %d", swap.ID)
						tx.Model(model.SwapFillTx{}).Where("fill_swap_tx_hash = ?", swapTx.FillSwapTxHash).Updates(
							map[string]interface{}{
								"status":     model.FillTxSent,
								"updated_at": time.Now().Unix(),
							})
						swap.Status = SwapSent
						swap.FillTxHash = swapTx.FillSwapTxHash
						engine.updateSwap(tx, &swap)

						isSkip = true
					}
				} else {
					swap.Status = SwapSending
					engine.updateSwap(tx, &swap)
				}
				return isSkip, tx.Commit().Error
			}()
			if writeDBErr != nil {
				util.Logger.Errorf("write db error: %s", writeDBErr.Error())
				util.SendTelegramMessage(fmt.Sprintf("write db error: %s", writeDBErr.Error()))
				continue
			}
			if skip {
				util.Logger.Debugf("skip this swap, start tx hash %s", swap.StartTxHash)
				continue
			}

			util.Logger.Infof("Swap token direction %s, sponsor: %s, SourceChainErc20Addr: %s, TargetChainErc20Addr: %s, TargetChainToAddr: %s, amount %s, decimals %d", direction, swap.Sponsor, swap.SourceChainErc20Addr, swap.TargetChainErc20Addr, swap.TargetChainToAddr, swap.Amount, swap.Decimals)
			swapTx, swapErr := engine.doSwap(&swap, swapPairInstance)

			writeDBErr = func() error {
				tx := engine.db.Begin()
				if err := tx.Error; err != nil {
					return err
				}
				if swapErr != nil {
					util.Logger.Errorf("do swap failed: %s, start hash %s", swapErr.Error(), swap.StartTxHash)
					util.SendTelegramMessage(fmt.Sprintf("do swap failed: %s, start hash %s", swapErr.Error(), swap.StartTxHash))
					if swapErr.Error() == core.ErrReplaceUnderpriced.Error() {
						//delete the fill swap tx
						tx.Where("fill_swap_tx_hash = ?", swapTx.FillSwapTxHash).Delete(model.SwapFillTx{})
						// retry this swap
						swap.Status = SwapConfirmed
						swap.Log = fmt.Sprintf("do swap failure: %s", swapErr.Error())

						engine.updateSwap(tx, &swap)
					} else {
						fillTxHash := ""
						if swapTx != nil {
							tx.Model(model.SwapFillTx{}).Where("fill_swap_tx_hash = ?", swapTx.FillSwapTxHash).Updates(
								map[string]interface{}{
									"status":     model.FillTxFailed,
									"updated_at": time.Now().Unix(),
								})
							fillTxHash = swapTx.FillSwapTxHash
						}

						swap.Status = SwapSendFailed
						swap.FillTxHash = fillTxHash
						swap.Log = fmt.Sprintf("do swap failure: %s", swapErr.Error())
						engine.updateSwap(tx, &swap)
					}
				} else {
					tx.Model(model.SwapFillTx{}).Where("fill_swap_tx_hash = ?", swapTx.FillSwapTxHash).Updates(
						map[string]interface{}{
							"status":     model.FillTxSent,
							"updated_at": time.Now().Unix(),
						})

					swap.Status = SwapSent
					swap.FillTxHash = swapTx.FillSwapTxHash
					engine.updateSwap(tx, &swap)
				}

				return tx.Commit().Error
			}()

			if writeDBErr != nil {
				util.Logger.Errorf("write db error: %s", writeDBErr.Error())
				util.SendTelegramMessage(fmt.Sprintf("write db error: %s", writeDBErr.Error()))
			}

			if swap.Direction == SwapMain2Side {
				time.Sleep(time.Duration(engine.config.ChainConfig.SideWaitMilliSecBetweenSwaps) * time.Millisecond)
			} else {
				time.Sleep(time.Duration(engine.config.ChainConfig.MainWaitMilliSecBetweenSwaps) * time.Millisecond)
			}
		}
	}
}

func (engine *SwapEngine) doSwap(swap *model.Swap, swapPairInstance *SwapPairIns) (*model.SwapFillTx, error) {
	amount := big.NewInt(0)
	_, ok := amount.SetString(swap.Amount, 10)
	if !ok {
		return nil, fmt.Errorf("invalid swap amount: %s", swap.Amount)
	}

	if swap.Direction == SwapMain2Side {
		sideClientMutex.Lock()
		defer sideClientMutex.Unlock()
		data, err := abiEncodeFillMain2SideSwap(ethcom.HexToHash(swap.StartTxHash), ethcom.HexToAddress(swap.SourceChainErc20Addr), ethcom.HexToAddress(swap.TargetChainToAddr), amount, engine.sideSwapAgentABI)
		if err != nil {
			return nil, err
		}
		signedTx, err := buildSignedTransaction(engine.sideSwapAgent, engine.sideClient, data, engine.sidePrivateKey)
		if err != nil {
			return nil, err
		}
		swapTx := &model.SwapFillTx{
			Direction:       SwapMain2Side,
			StartSwapTxHash: swap.StartTxHash,
			FillSwapTxHash:  signedTx.Hash().String(),
			GasPrice:        signedTx.GasPrice().String(),
			Status:          model.FillTxCreated,
		}
		err = engine.insertSwapTxToDB(swapTx)
		if err != nil {
			return nil, err
		}
		err = engine.sideClient.SendTransaction(context.Background(), signedTx)
		if err != nil {
			util.Logger.Errorf("broadcast tx to Side error: %s", err.Error())
			return nil, err
		}
		util.Logger.Infof("Send transaction to Side, %s/%s", engine.config.ChainConfig.SideExplorerUrl, signedTx.Hash().String())
		return swapTx, nil
	} else {
		mainClientMutex.Lock()
		defer mainClientMutex.Unlock()
		data, err := abiEncodeFillSide2MainSwap(ethcom.HexToHash(swap.StartTxHash), ethcom.HexToAddress(swap.SourceChainErc20Addr), ethcom.HexToAddress(swap.TargetChainToAddr), amount, engine.mainSwapAgentABI)
		signedTx, err := buildSignedTransaction(engine.mainSwapAgent, engine.mainClient, data, engine.mainPrivateKey)
		if err != nil {
			return nil, err
		}
		swapTx := &model.SwapFillTx{
			Direction:       SwapSide2Main,
			StartSwapTxHash: swap.StartTxHash,
			GasPrice:        signedTx.GasPrice().String(),
			FillSwapTxHash:  signedTx.Hash().String(),
			Status:          model.FillTxCreated,
		}
		err = engine.insertSwapTxToDB(swapTx)
		if err != nil {
			return nil, err
		}
		err = engine.mainClient.SendTransaction(context.Background(), signedTx)
		if err != nil {
			util.Logger.Errorf("broadcast tx to Main error: %s", err.Error())
			return nil, err
		} else {
			util.Logger.Infof("Send transaction to Main, %s/%s", engine.config.ChainConfig.MainExplorerUrl, signedTx.Hash().String())
		}
		return swapTx, nil
	}
}

func (engine *SwapEngine) trackSwapTxDaemon() {
	go func() {
		for {
			time.Sleep(SleepTime * time.Second)

			swapTxs := make([]model.SwapFillTx, 0)
			engine.db.Where("status = ? and track_retry_counter >= ?", model.FillTxSent, engine.config.ChainConfig.MainMaxTrackRetry).
				Order("id asc").Limit(TrackSentTxBatchSize).Find(&swapTxs)

			if len(swapTxs) > 0 {
				util.Logger.Infof("%d fill tx are missing, mark these swaps as failed", len(swapTxs))
			}

			for _, swapTx := range swapTxs {
				chainName := "Main"
				maxRetry := engine.config.ChainConfig.MainMaxTrackRetry
				if swapTx.Direction == SwapMain2Side {
					chainName = "Side"
					maxRetry = engine.config.ChainConfig.SideMaxTrackRetry
				}
				util.Logger.Errorf("The fill tx is sent, however, after %d seconds its status is still uncertain. Mark tx as missing and mark swap as failed, chain %s, fill hash %s", SleepTime*maxRetry, chainName, swapTx.StartSwapTxHash)
				util.SendTelegramMessage(fmt.Sprintf("The fill tx is sent, however, after %d seconds its status is still uncertain. Mark tx as missing and mark swap as failed, chain %s, start hash %s", SleepTime*maxRetry, chainName, swapTx.StartSwapTxHash))

				writeDBErr := func() error {
					tx := engine.db.Begin()
					if err := tx.Error; err != nil {
						return err
					}
					tx.Model(model.SwapFillTx{}).Where("id = ?", swapTx.ID).Updates(
						map[string]interface{}{
							"status":     model.FillTxMissing,
							"updated_at": time.Now().Unix(),
						})

					swap, err := engine.getSwapByStartTxHash(tx, swapTx.StartSwapTxHash)
					if err != nil {
						tx.Rollback()
						return err
					}
					swap.Status = SwapSendFailed
					swap.Log = fmt.Sprintf("track fill tx for more than %d times, the fill tx status is still uncertain", maxRetry)
					engine.updateSwap(tx, swap)

					return tx.Commit().Error
				}()
				if writeDBErr != nil {
					util.Logger.Errorf("write db error: %s", writeDBErr.Error())
					util.SendTelegramMessage(fmt.Sprintf("write db error: %s", writeDBErr.Error()))
				}
			}
		}
	}()

	go func() {
		for {
			time.Sleep(SleepTime * time.Second)

			mainSwapTxs := make([]model.SwapFillTx, 0)
			engine.db.Where("status = ? and direction = ? and track_retry_counter < ?", model.FillTxSent, SwapSide2Main, engine.config.ChainConfig.MainMaxTrackRetry).
				Order("id asc").Limit(TrackSentTxBatchSize).Find(&mainSwapTxs)

			sideSwapTxs := make([]model.SwapFillTx, 0)
			engine.db.Where("status = ? and direction = ? and track_retry_counter < ?", model.FillTxSent, SwapMain2Side, engine.config.ChainConfig.SideMaxTrackRetry).
				Order("id asc").Limit(TrackSentTxBatchSize).Find(&sideSwapTxs)

			swapTxs := append(mainSwapTxs, sideSwapTxs...)

			if len(swapTxs) > 0 {
				util.Logger.Debugf("Track %d non-finalized swap txs", len(swapTxs))
			}

			for _, swapTx := range swapTxs {
				gasPrice := big.NewInt(0)
				gasPrice.SetString(swapTx.GasPrice, 10)

				var client *ethclient.Client
				var chainName string
				if swapTx.Direction == SwapSide2Main {
					client = engine.mainClient
					chainName = "Main"
				} else {
					client = engine.sideClient
					chainName = "Side"
				}
				var txRecipient *types.Receipt
				queryTxStatusErr := func() error {
					block, err := client.BlockByNumber(context.Background(), nil)
					if err != nil {
						util.Logger.Debugf("%s, query block failed: %s", chainName, err.Error())
						return err
					}
					txRecipient, err = client.TransactionReceipt(context.Background(), ethcom.HexToHash(swapTx.FillSwapTxHash))
					if err != nil {
						util.Logger.Debugf("%s, query tx failed: %s", chainName, err.Error())
						return err
					}
					if block.Number().Int64() < txRecipient.BlockNumber.Int64()+engine.config.ChainConfig.MainConfirmNum {
						return fmt.Errorf("%s, swap tx is still not finalized", chainName)
					}
					return nil
				}()

				writeDBErr := func() error {
					tx := engine.db.Begin()
					if err := tx.Error; err != nil {
						return err
					}
					if queryTxStatusErr != nil {
						tx.Model(model.SwapFillTx{}).Where("id = ?", swapTx.ID).Updates(
							map[string]interface{}{
								"track_retry_counter": gorm.Expr("track_retry_counter + 1"),
								"updated_at":          time.Now().Unix(),
							})
					} else {
						txFee := big.NewInt(1).Mul(gasPrice, big.NewInt(int64(txRecipient.GasUsed))).String()
						if txRecipient.Status == TxFailedStatus {
							util.Logger.Infof(fmt.Sprintf("fill swap tx is failed, chain %s, txHash: %s", chainName, txRecipient.TxHash.String()))
							util.SendTelegramMessage(fmt.Sprintf("fill swap tx is failed, chain %s, txHash: %s", chainName, txRecipient.TxHash.String()))
							tx.Model(model.SwapFillTx{}).Where("id = ?", swapTx.ID).Updates(
								map[string]interface{}{
									"status":              model.FillTxFailed,
									"height":              txRecipient.BlockNumber.Int64(),
									"consumed_fee_amount": txFee,
									"updated_at":          time.Now().Unix(),
								})

							swap, err := engine.getSwapByStartTxHash(tx, swapTx.StartSwapTxHash)
							if err != nil {
								tx.Rollback()
								return err
							}
							swap.Status = SwapSendFailed
							swap.Log = "fill tx is failed"
							engine.updateSwap(tx, swap)
						} else {
							util.Logger.Infof(fmt.Sprintf("fill swap tx is success, chain %s, txHash: %s", chainName, txRecipient.TxHash.String()))
							tx.Model(model.SwapFillTx{}).Where("id = ?", swapTx.ID).Updates(
								map[string]interface{}{
									"status":              model.FillTxSuccess,
									"height":              txRecipient.BlockNumber.Int64(),
									"consumed_fee_amount": txFee,
									"updated_at":          time.Now().Unix(),
								})

							swap, err := engine.getSwapByStartTxHash(tx, swapTx.StartSwapTxHash)
							if err != nil {
								tx.Rollback()
								return err
							}
							swap.Status = SwapSuccess
							engine.updateSwap(tx, swap)
						}
					}
					return tx.Commit().Error
				}()
				if writeDBErr != nil {
					util.Logger.Errorf("update db failure: %s", writeDBErr.Error())
					util.SendTelegramMessage(fmt.Sprintf("Upgent alert: update db failure: %s", writeDBErr.Error()))
				}

			}
		}
	}()
}

func (engine *SwapEngine) getSwapByStartTxHash(tx *gorm.DB, txHash string) (*model.Swap, error) {
	swap := model.Swap{}
	err := tx.Where("start_tx_hash = ?", txHash).First(&swap).Error
	if err != nil {
		return nil, err
	}
	if !engine.verifySwap(&swap) {
		return nil, fmt.Errorf("hmac verification failure")
	}
	return &swap, nil
}

func (engine *SwapEngine) insertSwapTxToDB(data *model.SwapFillTx) error {
	tx := engine.db.Begin()
	if err := tx.Error; err != nil {
		return err
	}

	if err := tx.Create(data).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (engine *SwapEngine) AddSwapPairInstance(swapPair *model.SwapPair) error {
	lowBound := big.NewInt(0)
	_, ok := lowBound.SetString(swapPair.LowBound, 10)
	if !ok {
		return fmt.Errorf("invalid lowBound amount: %s", swapPair.LowBound)
	}
	upperBound := big.NewInt(0)
	_, ok = upperBound.SetString(swapPair.UpperBound, 10)
	if !ok {
		return fmt.Errorf("invalid upperBound amount: %s", swapPair.LowBound)
	}

	engine.mutex.Lock()
	defer engine.mutex.Unlock()
	engine.swapPairsFromERC20Addr[ethcom.HexToAddress(swapPair.MainChainErc20Addr)] = &SwapPairIns{
		Symbol:     swapPair.Symbol,
		Name:       swapPair.Name,
		Decimals:   swapPair.Decimals,
		LowBound:   lowBound,
		UpperBound: upperBound,
		MainChainErc20Addr:  ethcom.HexToAddress(swapPair.MainChainErc20Addr),
		SideChainErc20Addr:  ethcom.HexToAddress(swapPair.SideChainErc20Addr),
	}
	engine.main20ToSide20[ethcom.HexToAddress(swapPair.MainChainErc20Addr)] = ethcom.HexToAddress(swapPair.SideChainErc20Addr)
	engine.side20ToMain20[ethcom.HexToAddress(swapPair.SideChainErc20Addr)] = ethcom.HexToAddress(swapPair.MainChainErc20Addr)

	util.Logger.Infof("Create new swap pair, symbol %s, main chain address %s, side chain address %s", swapPair.Symbol, swapPair.MainChainErc20Addr,swapPair.SideChainErc20Addr)

	return nil
}

func (engine *SwapEngine) GetSwapPairInstance(erc20Addr ethcom.Address) (*SwapPairIns, error) {
	engine.mutex.RLock()
	defer engine.mutex.RUnlock()

	tokenInstance, ok := engine.swapPairsFromERC20Addr[erc20Addr]
	if !ok {
		return nil, fmt.Errorf("swap instance doesn't exist")
	}
	return tokenInstance, nil
}

func (engine *SwapEngine) UpdateSwapInstance(swapPair *model.SwapPair) {
	engine.mutex.Lock()
	defer engine.mutex.Unlock()

	sideTokenAddr := ethcom.HexToAddress(swapPair.SideChainErc20Addr)
	tokenInstance, ok := engine.swapPairsFromERC20Addr[sideTokenAddr]
	if !ok {
		return
	}

	if !swapPair.Available {
		delete(engine.swapPairsFromERC20Addr, sideTokenAddr)
		return
	}

	upperBound := big.NewInt(0)
	_, ok = upperBound.SetString(swapPair.UpperBound, 10)
	tokenInstance.UpperBound = upperBound

	lowBound := big.NewInt(0)
	_, ok = upperBound.SetString(swapPair.LowBound, 10)
	tokenInstance.LowBound = lowBound

	engine.swapPairsFromERC20Addr[sideTokenAddr] = tokenInstance
}
