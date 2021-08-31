package executor

import (
	"context"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcmm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	agent "github.com/bcskill/eth-main-side-swap/abi"
	contractabi "github.com/bcskill/eth-main-side-swap/abi"
	"github.com/bcskill/eth-main-side-swap/common"
	"github.com/bcskill/eth-main-side-swap/util"
)

type MainExecutor struct {
	Chain  string
	Config *util.Config

	SwapAgentAddr    ethcmm.Address
	mainSwapAgentInst *contractabi.MainSwapAgent
	SwapAgentAbi     abi.ABI
	Client           *ethclient.Client
}

func NewMainExecutor(ethClient *ethclient.Client, swapAddr string, config *util.Config) *MainExecutor {
	agentAbi, err := abi.JSON(strings.NewReader(agent.MainSwapAgentABI))
	if err != nil {
		panic("marshal abi error")
	}
	mainSwapAgentInst, err := contractabi.NewMainSwapAgent(ethcmm.HexToAddress(swapAddr), ethClient)
	if err != nil {
		panic(err.Error())
	}

	return &MainExecutor{
		Chain:            common.ChainMain,
		Config:           config,
		SwapAgentAddr:    ethcmm.HexToAddress(swapAddr),
		mainSwapAgentInst: mainSwapAgentInst,
		SwapAgentAbi:     agentAbi,
		Client:           ethClient,
	}
}

func (e *MainExecutor) GetChainName() string {
	return e.Chain
}

func (e *MainExecutor) GetBlockAndTxEvents(height int64) (*common.BlockAndEventLogs, error) {
	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	header, err := e.Client.HeaderByNumber(ctxWithTimeout, big.NewInt(height))
	if err != nil {
		return nil, err
	}

	packageLogs, err := e.GetLogs(header)
	if err != nil {
		return nil, err
	}

	return &common.BlockAndEventLogs{
		Height:          height,
		Chain:           e.Chain,
		BlockHash:       header.Hash().String(),
		ParentBlockHash: header.ParentHash.String(),
		BlockTime:       int64(header.Time),
		Events:          packageLogs,
	}, nil
}

func (e *MainExecutor) GetLogs(header *types.Header) ([]interface{}, error) {
	startEvs, err := e.GetSwapStartLogs(header)
	if err != nil {
		return nil, err
	}
	regiserEvs, err := e.GetSwapPairRegisterLogs(header)
	if err != nil {
		return nil, err
	}
	var res = make([]interface{}, 0, len(startEvs)+len(regiserEvs))
	res = append(append(res, startEvs...), regiserEvs...)
	return res, nil

}

func (e *MainExecutor) GetSwapPairRegisterLogs(header *types.Header) ([]interface{}, error) {
	topics := [][]ethcmm.Hash{{SwapPairRegisterEventHash}}

	blockHash := header.Hash()

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logs, err := e.Client.FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []ethcmm.Address{e.SwapAgentAddr},
	})
	if err != nil {
		return nil, err
	}

	eventModels := make([]interface{}, 0, len(logs))
	for _, log := range logs {
		event, err := ParseSwapPairRegisterEvent(&e.SwapAgentAbi, &log)
		if err != nil {
			util.Logger.Errorf("parse event log error, er=%s", err.Error())
			continue
		}
		if event == nil {
			continue
		}

		eventModel := event.ToSwapPairRegisterLog(&log)
		eventModel.Chain = e.Chain
		util.Logger.Debugf("Found register event, erc20 address: %d, name: %s, symbol: %s, decimals: %d",
			eventModel.ERC20Addr, eventModel.Name, eventModel.Symbol, eventModel.Decimals)
		eventModels = append(eventModels, eventModel)
	}
	return eventModels, nil
}

func (e *MainExecutor) GetSwapStartLogs(header *types.Header) ([]interface{}, error) {
	topics := [][]ethcmm.Hash{{Main2SideSwapStartedEventHash}}

	blockHash := header.Hash()

	ctxWithTimeout, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	logs, err := e.Client.FilterLogs(ctxWithTimeout, ethereum.FilterQuery{
		BlockHash: &blockHash,
		Topics:    topics,
		Addresses: []ethcmm.Address{e.SwapAgentAddr},
	})
	if err != nil {
		return nil, err
	}

	eventModels := make([]interface{}, 0, len(logs))
	for _, log := range logs {
		event, err := ParseMain2SideSwapStartEvent(&e.SwapAgentAbi, &log)
		if err != nil {
			util.Logger.Errorf("parse event log error, er=%s", err.Error())
			continue
		}

		if event == nil {
			continue
		}

		eventModel := event.ToSwapStartTxLog(&log)
		eventModel.Chain = e.Chain
		util.Logger.Debugf("Found Main2Side swap, txHash: %s, token address: %s, amount: %s, fee amount: %s",
			eventModel.TxHash, eventModel.TokenAddr, eventModel.Amount, eventModel.FeeAmount)
		eventModels = append(eventModels, eventModel)
	}
	return eventModels, nil
}
