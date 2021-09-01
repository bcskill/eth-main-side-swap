package executor

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	ethcmm "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"

	agent "github.com/bcskill/eth-main-side-swap/abi"
	contractabi "github.com/bcskill/eth-main-side-swap/abi"
	"github.com/bcskill/eth-main-side-swap/common"
	"github.com/bcskill/eth-main-side-swap/util"
)

type SideExecutor struct {
	Chain  string
	Config *util.Config

	SwapAgentAddr    ethcmm.Address
	SideSwapAgentInst *contractabi.MainSwapAgent
	SwapAgentAbi     abi.ABI
	Client           *ethclient.Client
}

func NewSideExecutor(ethClient *ethclient.Client, swapAddr string, config *util.Config) *SideExecutor {
	agentAbi, err := abi.JSON(strings.NewReader(agent.SideSwapAgentABI))
	if err != nil {
		panic("marshal abi error")
	}

	bscSwapAgentInst, err := contractabi.NewMainSwapAgent(ethcmm.HexToAddress(swapAddr), ethClient)
	if err != nil {
		panic(err.Error())
	}

	return &SideExecutor{
		Chain:            common.ChainSide,
		Config:           config,
		SwapAgentAddr:    ethcmm.HexToAddress(swapAddr),
		SideSwapAgentInst: bscSwapAgentInst,
		SwapAgentAbi:     agentAbi,
		Client:           ethClient,
	}
}

func (e *SideExecutor) GetChainName() string {
	return e.Chain
}

func (e *SideExecutor) GetBlockAndTxEvents(height int64) (*common.BlockAndEventLogs, error) {
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
func (e *SideExecutor) GetLogs(header *types.Header) ([]interface{}, error) {
	return e.GetSwapSide2MainEventLogs(header)
}

func (e *SideExecutor) GetSwapSide2MainEventLogs(header *types.Header) ([]interface{}, error) {
	topics := [][]ethcmm.Hash{{Side2MainSwapStartedEventHash}}

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
		event, err := ParseSide2MainSwapStartEvent(&e.SwapAgentAbi, &log)
		if err != nil {
			util.Logger.Errorf("parse event log error, er=%s", err.Error())
			continue
		}
		if event == nil {
			continue
		}

		eventModel := event.ToSwapStartTxLog(&log)
		eventModel.Chain = e.Chain
		util.Logger.Debugf("Found swap side 2 main event, txHash: %s, sideChainErc20Addr address: %s, mainChainErc20Addr address: %s, mainChainToAddr address: %s, amount: %s, fee amount: %s",
			eventModel.TxHash, eventModel.SourceChainErc20Addr, eventModel.TargetChainErc20Addr, eventModel.TargetChainToAddr, eventModel.Amount, eventModel.FeeAmount)
		eventModels = append(eventModels, eventModel)
	}
	return eventModels, nil
}
