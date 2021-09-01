package executor

import (
	common "github.com/bcskill/eth-main-side-swap/common"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcmm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"math/big"

	"github.com/bcskill/eth-main-side-swap/model"
)

type Executor interface {
	GetBlockAndTxEvents(height int64) (*common.BlockAndEventLogs, error)
	GetChainName() string
}

// ===================  SwapStarted =============
var (
	SwapMain2SideEventName        = "SwapMain2SideEvent"
	SwapSide2MainEventName        = "SwapSide2MainEvent"
	Main2SideSwapStartedEventHash = ethcmm.HexToHash("0x64bf044f989e78b5431b2fd639882f746b1795ff367b910eed4136d067fa390e")
	Side2MainSwapStartedEventHash = ethcmm.HexToHash("0x7090729cb71d105f50d09825b7425766b047a5d225e26e73ceb1d6fd4d9c1e58")
)

type Main2SideSwapStartedEvent struct {
	MainChainErc20Addr  ethcmm.Address
	SideChainErc20Addr  ethcmm.Address
	SideChainToAddr     ethcmm.Address
	Amount              *big.Int
	FeeAmount           *big.Int
}

func (ev *Main2SideSwapStartedEvent) ToSwapStartTxLog(log *types.Log) *model.SwapStartTxLog {
	pack := &model.SwapStartTxLog{
		SourceChainErc20Addr:   ev.MainChainErc20Addr.String(),
		TargetChainErc20Addr:   ev.SideChainErc20Addr.String(),
		TargetChainToAddr:      ev.SideChainToAddr.String(),
		Amount:               ev.Amount.String(),
		FeeAmount:            ev.FeeAmount.String(),
		BlockHash:            log.BlockHash.Hex(),
		TxHash:               log.TxHash.String(),
		Height:               int64(log.BlockNumber),
	}
	return pack
}

func ParseMain2SideSwapStartEvent(abi *abi.ABI, log *types.Log) (*Main2SideSwapStartedEvent, error) {
	var ev Main2SideSwapStartedEvent

	err := abi.UnpackIntoInterface(&ev, SwapMain2SideEventName, log.Data)
	if err != nil {
		return nil, err
	}

	ev.MainChainErc20Addr = ethcmm.BytesToAddress(log.Topics[1].Bytes())
	ev.SideChainErc20Addr = ethcmm.BytesToAddress(log.Topics[2].Bytes())
	ev.SideChainToAddr =    ethcmm.BytesToAddress(log.Topics[3].Bytes())

	return &ev, nil
}

type Side2MainSwapStartedEvent struct {
	SideChainErc20Addr ethcmm.Address
	MainChainErc20Addr ethcmm.Address
	MainChainToAddr    ethcmm.Address
	Amount             *big.Int
	FeeAmount          *big.Int
}

func (ev *Side2MainSwapStartedEvent) ToSwapStartTxLog(log *types.Log) *model.SwapStartTxLog {
	pack := &model.SwapStartTxLog{
		SourceChainErc20Addr: ev.SideChainErc20Addr.String(),
		TargetChainErc20Addr:   ev.MainChainErc20Addr.String(),
		TargetChainToAddr: ev.MainChainToAddr.String(),
		Amount:      ev.Amount.String(),

		FeeAmount: ev.FeeAmount.String(),
		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}
	return pack
}

func ParseSide2MainSwapStartEvent(abi *abi.ABI, log *types.Log) (*Side2MainSwapStartedEvent, error) {
	var ev Side2MainSwapStartedEvent

	err := abi.UnpackIntoInterface(&ev, SwapSide2MainEventName, log.Data)
	if err != nil {
		return nil, err
	}

	ev.SideChainErc20Addr = ethcmm.BytesToAddress(log.Topics[1].Bytes())
	ev.MainChainErc20Addr = ethcmm.BytesToAddress(log.Topics[2].Bytes())
	ev.MainChainToAddr = ethcmm.BytesToAddress(log.Topics[3].Bytes())

	return &ev, nil
}


// =================  SwapPairRegister ===================
var (
	SwapPairRegisterEventName = "SwapPairRegisterEvent"
	SwapPairRegisterEventHash = ethcmm.HexToHash("0x84938b4885411ca7a208fde741a0feec069b8173eef6d83fcdecc5daf3a2af3d")
)

type SwapPairRegisterEvent struct {
	Sponsor            ethcmm.Address
	MainChainErc20Addr ethcmm.Address
	SideChainFromAddr  ethcmm.Address
	Name         string
	Symbol       string
	Decimals     uint8
}

func (ev *SwapPairRegisterEvent) ToSwapPairRegisterLog(log *types.Log) *model.SwapPairRegisterTxLog {
	pack := &model.SwapPairRegisterTxLog{
		Sponsor:            ev.Sponsor.String(),
		MainChainErc20Addr: ev.MainChainErc20Addr.String(),
		SideChainFromAddr:  ev.SideChainFromAddr.String(),
		Symbol:    ev.Symbol,
		Name:      ev.Name,
		Decimals:  int(ev.Decimals),

		BlockHash: log.BlockHash.Hex(),
		TxHash:    log.TxHash.String(),
		Height:    int64(log.BlockNumber),
	}
	return pack
}

func ParseSwapPairRegisterEvent(abi *abi.ABI, log *types.Log) (*SwapPairRegisterEvent, error) {
	var ev SwapPairRegisterEvent

	err := abi.UnpackIntoInterface(&ev, SwapPairRegisterEventName, log.Data)
	if err != nil {
		return nil, err
	}
	ev.Sponsor = ethcmm.BytesToAddress(log.Topics[1].Bytes())
	ev.MainChainErc20Addr = ethcmm.BytesToAddress(log.Topics[2].Bytes())
	ev.SideChainFromAddr = ethcmm.BytesToAddress(log.Topics[3].Bytes())

	return &ev, nil
}
