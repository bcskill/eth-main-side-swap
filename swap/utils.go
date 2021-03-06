package swap

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcom "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	contractabi "github.com/bcskill/eth-main-side-swap/abi"
	"github.com/bcskill/eth-main-side-swap/common"
	"github.com/bcskill/eth-main-side-swap/model"
	"github.com/bcskill/eth-main-side-swap/util"

)

func buildSwapPairInstance(pairs []model.SwapPair) (map[ethcom.Address]*SwapPairIns, error) {
	swapPairInstances := make(map[ethcom.Address]*SwapPairIns, len(pairs))

	for _, pair := range pairs {

		lowBound := big.NewInt(0)
		_, ok := lowBound.SetString(pair.LowBound, 10)
		if !ok {
			panic(fmt.Sprintf("invalid lowBound amount: %s", pair.LowBound))
		}
		upperBound := big.NewInt(0)
		_, ok = upperBound.SetString(pair.UpperBound, 10)
		if !ok {
			panic(fmt.Sprintf("invalid upperBound amount: %s", pair.LowBound))
		}

		swapPairInstances[ethcom.HexToAddress(pair.MainChainErc20Addr)] = &SwapPairIns{
			Symbol:     pair.Symbol,
			Name:       pair.Name,
			Decimals:   pair.Decimals,
			LowBound:   lowBound,
			UpperBound: upperBound,
			MainChainErc20Addr:  ethcom.HexToAddress(pair.MainChainErc20Addr),
			SideChainErc20Addr:  ethcom.HexToAddress(pair.SideChainErc20Addr),
		}

		util.Logger.Infof("Load swap pair, symbol %s, main chain address %s, side chain address %s", pair.Symbol, pair.MainChainErc20Addr, pair.SideChainErc20Addr)
	}

	return swapPairInstances, nil
}

func GetKeyConfig(cfg *util.Config) (*util.KeyConfig, error) {
	if cfg.KeyManagerConfig.KeyType == common.AWSPrivateKey {
		result, err := util.GetSecret(cfg.KeyManagerConfig.AWSSecretName, cfg.KeyManagerConfig.AWSRegion)
		if err != nil {
			return nil, err
		}

		keyConfig := util.KeyConfig{}
		err = json.Unmarshal([]byte(result), &keyConfig)
		if err != nil {
			return nil, err
		}
		return &keyConfig, nil
	} else {
		return &util.KeyConfig{
			HMACKey:        cfg.KeyManagerConfig.LocalHMACKey,
			AdminApiKey:    cfg.KeyManagerConfig.LocalAdminApiKey,
			AdminSecretKey: cfg.KeyManagerConfig.LocalAdminSecretKey,
			SidePrivateKey:  cfg.KeyManagerConfig.LocalSidePrivateKey,
			MainPrivateKey:  cfg.KeyManagerConfig.LocalMainPrivateKey,
		}, nil
	}
}

func abiEncodeFillMain2SideSwap(mainChainTxHash ethcom.Hash, mainChainErc20Addr ethcom.Address, sideChainToAddr ethcom.Address, amount *big.Int, abi *abi.ABI) ([]byte, error) {
	data, err := abi.Pack("fillMain2SideSwap", mainChainTxHash, mainChainErc20Addr, sideChainToAddr, amount)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func abiEncodeERC20Transfer(recipient ethcom.Address, amount *big.Int, abi *abi.ABI) ([]byte, error) {
	data, err := abi.Pack("transfer", recipient, amount)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func abiEncodeFillSide2MainSwap(mainTxHash ethcom.Hash, erc20Addr ethcom.Address, toAddress ethcom.Address, amount *big.Int, abi *abi.ABI) ([]byte, error) {
	data, err := abi.Pack("fillSide2MainSwap", mainTxHash, erc20Addr, toAddress, amount)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func abiEncodeCreateSwapPair(mainChainTxHash ethcom.Hash, mainChainErc20Addr ethcom.Address, sideChainErc20Addr ethcom.Address, name, symbol string, decimals uint8, abi *abi.ABI) ([]byte, error) {
	data, err := abi.Pack("createSwapPair", mainChainTxHash, mainChainErc20Addr, sideChainErc20Addr, name, symbol, decimals)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func buildSignedTransaction(contract ethcom.Address, mainClient *ethclient.Client, txInput []byte, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	txOpts := bind.NewKeyedTransactor(privateKey)

	nonce, err := mainClient.PendingNonceAt(context.Background(), txOpts.From)
	if err != nil {
		return nil, err
	}
	gasPrice, err := mainClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	value := big.NewInt(0)
	msg := ethereum.CallMsg{From: txOpts.From, To: &contract, GasPrice: gasPrice, Value: value, Data: txInput}
	gasLimit, err := mainClient.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
	}
	gasLimit += gasLimit / 2
	rawTx := types.NewTransaction(nonce, contract, value, gasLimit, gasPrice, txInput)
	signedTx, err := txOpts.Signer(txOpts.From, rawTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func buildNativeCoinTransferTx(contract ethcom.Address, mainClient *ethclient.Client, value *big.Int, privateKey *ecdsa.PrivateKey) (*types.Transaction, error) {
	txOpts := bind.NewKeyedTransactor(privateKey)

	nonce, err := mainClient.PendingNonceAt(context.Background(), txOpts.From)
	if err != nil {
		return nil, err
	}
	gasPrice, err := mainClient.SuggestGasPrice(context.Background())
	if err != nil {
		return nil, err
	}
	msg := ethereum.CallMsg{From: txOpts.From, To: &contract, GasPrice: gasPrice, Value: value}
	gasLimit, err := mainClient.EstimateGas(context.Background(), msg)
	if err != nil {
		return nil, fmt.Errorf("failed to estimate gas needed: %v", err)
	}

	rawTx := types.NewTransaction(nonce, contract, value, gasLimit, gasPrice, nil)
	signedTx, err := txOpts.Signer(txOpts.From, rawTx)
	if err != nil {
		return nil, err
	}

	return signedTx, nil
}

func queryDeployedBEP20ContractAddr(erc20Addr ethcom.Address, sideSwapAgentAddr ethcom.Address, txRecipient *types.Receipt, sideClient *ethclient.Client) (ethcom.Address, error) {
	swapAgentInstance, err := contractabi.NewSideSwapAgent(sideSwapAgentAddr, sideClient)
	if err != nil {
		return ethcom.Address{}, err
	}
	if len(txRecipient.Logs) != 2 {
		return ethcom.Address{}, fmt.Errorf("Expected tx logs length in recipient is 2, actual it is %d", len(txRecipient.Logs))
	}
	createSwapEvent, err := swapAgentInstance.ParseSwapPairCreatedEvent(*txRecipient.Logs[1])
	if err != nil || createSwapEvent == nil {
		return ethcom.Address{}, err
	}

	util.Logger.Debugf("Deployed bep20 contact %s for register erc20 %s", createSwapEvent.SideChainErc20Addr.String(), erc20Addr.String())
	return createSwapEvent.SideChainErc20Addr, nil
}

func BuildKeys(privateKeyStr string) (*ecdsa.PrivateKey, *ecdsa.PublicKey, error) {
	if strings.HasPrefix(privateKeyStr, "0x") {
		privateKeyStr = privateKeyStr[2:]
	}
	priKey, err := crypto.HexToECDSA(privateKeyStr)
	if err != nil {
		return nil, nil, err
	}
	publicKey, ok := priKey.Public().(*ecdsa.PublicKey)
	if !ok {
		return nil, nil, fmt.Errorf("get public key error")
	}
	return priKey, publicKey, nil
}
