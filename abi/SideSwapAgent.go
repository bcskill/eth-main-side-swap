// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// SideSwapAgentMetaData contains all meta data concerning the SideSwapAgent contract.
var SideSwapAgentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sideChainErc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sendAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RechargeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mainChainTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mainChainErc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sideChainToAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SwapMain2SideFilledEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"mainChainTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mainChainErc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sideChainErc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"SwapPairCreatedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sideChainErc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mainChainErc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"fromAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"SwapSide2MainEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mainChainTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"mainChainErc20Addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sideChainErc20Addr\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"createSwapPair\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"createSwapPairTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"mainChainTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"mainChainErc20Addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sideChainToAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fillMain2SideSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filledMainTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"ownerAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sideChainErc20Addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"rechargeSide\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"sideChainErc20Banlance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapMappingMain2Side\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapMappingSide2Main\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sideChainErc20Addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swapSide2Main\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506124e4806100206000396000f3fe6080604052600436106100e85760003560e01c80637b0bdf111161008a578063cd6dc68711610059578063cd6dc687146105dc578063cd9fa9b814610637578063f017a4781461069c578063f2fde38b146106ef576100e8565b80637b0bdf11146103fe5780638da5cb5b1461048e57806395660c77146104e5578063bff281a314610576576100e8565b806334e19907116100c657806334e199071461032e57806338a29e851461036957806354cf2aeb146103bc578063715018a6146103e7576100e8565b80630b762dac146100ed5780631be17b65146101535780631d662f62146101e4575b600080fd5b6101396004803603604081101561010357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610740565b604051808215151515815260200191505060405180910390f35b34801561015f57600080fd5b506101a26004803603602081101561017657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610bf8565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156101f057600080fd5b50610314600480360360c081101561020757600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019064010000000081111561026e57600080fd5b82018360208201111561028057600080fd5b803590602001918460018302840111640100000000831117156102a257600080fd5b9091929391929390803590602001906401000000008111156102c357600080fd5b8201836020820111156102d557600080fd5b803590602001918460018302840111640100000000831117156102f757600080fd5b9091929391929390803560ff169060200190929190505050610c2b565b604051808215151515815260200191505060405180910390f35b34801561033a57600080fd5b506103676004803603602081101561035157600080fd5b81019080803590602001909291905050506111c1565b005b34801561037557600080fd5b506103a26004803603602081101561038c57600080fd5b8101908080359060200190929190505050611295565b604051808215151515815260200191505060405180910390f35b3480156103c857600080fd5b506103d16112b5565b6040518082815260200191505060405180910390f35b3480156103f357600080fd5b506103fc6112bb565b005b6104746004803603608081101561041457600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611446565b604051808215151515815260200191505060405180910390f35b34801561049a57600080fd5b506104a36119ff565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104f157600080fd5b506105346004803603602081101561050857600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611a25565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6105c26004803603604081101561058c57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050611a58565b604051808215151515815260200191505060405180910390f35b3480156105e857600080fd5b50610635600480360360408110156105ff57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050612000565b005b34801561064357600080fd5b506106866004803603602081101561065a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061204c565b6040518082815260200191505060405180910390f35b3480156106a857600080fd5b506106d5600480360360208110156106bf57600080fd5b8101908080359060200190929190505050612064565b604051808215151515815260200191505060405180910390f35b3480156106fb57600080fd5b5061073e6004803603602081101561071257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612084565b005b600061074b33612294565b156107be576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f636f6e7472616374206973206e6f7420616c6c6f77656420746f20737761700081525060200191505060405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461085f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6e6f2070726f787920636f6e747261637420697320616c6c6f7765640000000081525060200191505060405180910390fd5b6000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610966576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f6e6f2073776170207061697220666f72207468697320746f6b656e000000000081525060200191505060405180910390fd5b60075434146109dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f7377617020666565206e6f7420657175616c000000000000000000000000000081525060200191505060405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015610a9857600080fd5b505af1158015610aac573d6000803e3d6000fd5b505050506040513d6020811015610ac257600080fd5b810190808051906020019092919050505050610b2683600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546122a790919063ffffffff16565b600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff167ff19f19fa7d4ae2e098cbb45838aa7151ba4813912afa00a24e5c0a71599ee9488634604051808381526020018281526020019250505060405180910390a4600191505092915050565b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000610c3561232f565b73ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610cf7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600460008a815260200190815260200160002060009054906101000a900460ff1615610d6e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602281526020018061248d6022913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600160008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610e6f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6475706c696361746564206d61696e20636861696e207377617020706169720081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610f70576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f6475706c696361746564207369646520636861696e202073776170207061697281525060200191505060405180910390fd5b86600160008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555087600260008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600460008b815260200190815260200160002060006101000a81548160ff0219169083151502179055506000600560008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508673ffffffffffffffffffffffffffffffffffffffff168873ffffffffffffffffffffffffffffffffffffffff168a7fb04179acee3e3defc848704fc6453616f655decf86895f1cb479e98e7552960989898989896040518080602001806020018460ff1660ff1681526020018381038352888882818152602001925080828437600081840152601f19601f8201169050808301925050508381038252868682818152602001925080828437600081840152601f19601f82011690508083019250505097505050505050505060405180910390a46001905098975050505050505050565b6111c961232f565b73ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461128b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b8060078190555050565b60046020528060005260406000206000915054906101000a900460ff1681565b60075481565b6112c361232f565b73ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611385576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600061145061232f565b73ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614611512576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6003600086815260200190815260200160002060009054906101000a900460ff16156115a6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f6d61696e2074782066696c6c656420616c72656164790000000000000000000081525060200191505060405180910390fd5b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156116ad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f6e6f2073776170207061697220666f72207468697320746f6b656e000000000081525060200191505060405180910390fd5b828173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561172b57600080fd5b505afa15801561173f573d6000803e3d6000fd5b505050506040513d602081101561175557600080fd5b8101908080519060200190929190505050101580156117b3575082600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410155b611808576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806124686025913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b15801561188f57600080fd5b505af11580156118a3573d6000803e3d6000fd5b505050506040513d60208110156118b957600080fd5b81019080805190602001909291905050505060016003600088815260200190815260200160002060006101000a81548160ff02191690831515021790555061194983600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461233790919063ffffffff16565b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16877fcce98297f95dc1cec392b3533c9cb6329f064d83ea3ae3ad09fd3f2d968ec100866040518082815260200191505060405180910390a46001915050949350505050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000611a6333612294565b15611ad6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f636f6e7472616374206973206e6f7420616c6c6f77656420746f20737761700081525060200191505060405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614611b77576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6e6f2070726f787920636f6e747261637420697320616c6c6f7765640000000081525060200191505060405180910390fd5b6000600260008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611c7e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f6e6f2073776170207061697220666f72207468697320746f6b656e000000000081525060200191505060405180910390fd5b60008311611cf4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f54686520726563686172676520616d6f756e7420697320746f6f20736d616c6c81525060200191505060405180910390fd5b828473ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b158015611d7257600080fd5b505afa158015611d86573d6000803e3d6000fd5b505050506040513d6020811015611d9c57600080fd5b81019080805190602001909291905050501015611e04576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806124686025913960400191505060405180910390fd5b8373ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b158015611ebf57600080fd5b505af1158015611ed3573d6000803e3d6000fd5b505050506040513d6020811015611ee957600080fd5b810190808051906020019092919050505050611f4d83600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546122a790919063ffffffff16565b600560008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055503373ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fba435e9191118b57fc3cd6841f3248e60b065acfea46e76f975249ade106b03b856040518082815260200191505060405180910390a3600191505092915050565b81600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806007819055505050565b60056020528060005260406000206000915090505481565b60036020528060005260406000206000915054906101000a900460ff1681565b61208c61232f565b73ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461214e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156121d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806124426026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080823b905060008111915050919050565b600080828401905083811015612325576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600033905090565b600061237983836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612381565b905092915050565b600083831115829061242e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156123f35780820151818401526020810190506123d8565b50505050905090810190601f1680156124205780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838503905080915050939250505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e73756666696369656e7420636f6e7472616374206163636f756e742062616c616e63656d61696e20636861696e2074782068617368206372656174656420616c7265616479a2646970667358221220abbe86909f104e09f06534eebfb23d8318d1da7ec9c2121b8d7fad37aad4b48264736f6c63430006040033",
}

// SideSwapAgentABI is the input ABI used to generate the binding from.
// Deprecated: Use SideSwapAgentMetaData.ABI instead.
var SideSwapAgentABI = SideSwapAgentMetaData.ABI

// SideSwapAgentBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SideSwapAgentMetaData.Bin instead.
var SideSwapAgentBin = SideSwapAgentMetaData.Bin

// DeploySideSwapAgent deploys a new Ethereum contract, binding an instance of SideSwapAgent to it.
func DeploySideSwapAgent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SideSwapAgent, error) {
	parsed, err := SideSwapAgentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SideSwapAgentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SideSwapAgent{SideSwapAgentCaller: SideSwapAgentCaller{contract: contract}, SideSwapAgentTransactor: SideSwapAgentTransactor{contract: contract}, SideSwapAgentFilterer: SideSwapAgentFilterer{contract: contract}}, nil
}

// SideSwapAgent is an auto generated Go binding around an Ethereum contract.
type SideSwapAgent struct {
	SideSwapAgentCaller     // Read-only binding to the contract
	SideSwapAgentTransactor // Write-only binding to the contract
	SideSwapAgentFilterer   // Log filterer for contract events
}

// SideSwapAgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type SideSwapAgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SideSwapAgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SideSwapAgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SideSwapAgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SideSwapAgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SideSwapAgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SideSwapAgentSession struct {
	Contract     *SideSwapAgent    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SideSwapAgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SideSwapAgentCallerSession struct {
	Contract *SideSwapAgentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// SideSwapAgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SideSwapAgentTransactorSession struct {
	Contract     *SideSwapAgentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// SideSwapAgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type SideSwapAgentRaw struct {
	Contract *SideSwapAgent // Generic contract binding to access the raw methods on
}

// SideSwapAgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SideSwapAgentCallerRaw struct {
	Contract *SideSwapAgentCaller // Generic read-only contract binding to access the raw methods on
}

// SideSwapAgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SideSwapAgentTransactorRaw struct {
	Contract *SideSwapAgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSideSwapAgent creates a new instance of SideSwapAgent, bound to a specific deployed contract.
func NewSideSwapAgent(address common.Address, backend bind.ContractBackend) (*SideSwapAgent, error) {
	contract, err := bindSideSwapAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SideSwapAgent{SideSwapAgentCaller: SideSwapAgentCaller{contract: contract}, SideSwapAgentTransactor: SideSwapAgentTransactor{contract: contract}, SideSwapAgentFilterer: SideSwapAgentFilterer{contract: contract}}, nil
}

// NewSideSwapAgentCaller creates a new read-only instance of SideSwapAgent, bound to a specific deployed contract.
func NewSideSwapAgentCaller(address common.Address, caller bind.ContractCaller) (*SideSwapAgentCaller, error) {
	contract, err := bindSideSwapAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SideSwapAgentCaller{contract: contract}, nil
}

// NewSideSwapAgentTransactor creates a new write-only instance of SideSwapAgent, bound to a specific deployed contract.
func NewSideSwapAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*SideSwapAgentTransactor, error) {
	contract, err := bindSideSwapAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SideSwapAgentTransactor{contract: contract}, nil
}

// NewSideSwapAgentFilterer creates a new log filterer instance of SideSwapAgent, bound to a specific deployed contract.
func NewSideSwapAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*SideSwapAgentFilterer, error) {
	contract, err := bindSideSwapAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SideSwapAgentFilterer{contract: contract}, nil
}

// bindSideSwapAgent binds a generic wrapper to an already deployed contract.
func bindSideSwapAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SideSwapAgentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SideSwapAgent *SideSwapAgentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SideSwapAgent.Contract.SideSwapAgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SideSwapAgent *SideSwapAgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.SideSwapAgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SideSwapAgent *SideSwapAgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.SideSwapAgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SideSwapAgent *SideSwapAgentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SideSwapAgent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SideSwapAgent *SideSwapAgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SideSwapAgent *SideSwapAgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.contract.Transact(opts, method, params...)
}

// CreateSwapPairTx is a free data retrieval call binding the contract method 0x38a29e85.
//
// Solidity: function createSwapPairTx(bytes32 ) view returns(bool)
func (_SideSwapAgent *SideSwapAgentCaller) CreateSwapPairTx(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _SideSwapAgent.contract.Call(opts, &out, "createSwapPairTx", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CreateSwapPairTx is a free data retrieval call binding the contract method 0x38a29e85.
//
// Solidity: function createSwapPairTx(bytes32 ) view returns(bool)
func (_SideSwapAgent *SideSwapAgentSession) CreateSwapPairTx(arg0 [32]byte) (bool, error) {
	return _SideSwapAgent.Contract.CreateSwapPairTx(&_SideSwapAgent.CallOpts, arg0)
}

// CreateSwapPairTx is a free data retrieval call binding the contract method 0x38a29e85.
//
// Solidity: function createSwapPairTx(bytes32 ) view returns(bool)
func (_SideSwapAgent *SideSwapAgentCallerSession) CreateSwapPairTx(arg0 [32]byte) (bool, error) {
	return _SideSwapAgent.Contract.CreateSwapPairTx(&_SideSwapAgent.CallOpts, arg0)
}

// FilledMainTx is a free data retrieval call binding the contract method 0xf017a478.
//
// Solidity: function filledMainTx(bytes32 ) view returns(bool)
func (_SideSwapAgent *SideSwapAgentCaller) FilledMainTx(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _SideSwapAgent.contract.Call(opts, &out, "filledMainTx", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FilledMainTx is a free data retrieval call binding the contract method 0xf017a478.
//
// Solidity: function filledMainTx(bytes32 ) view returns(bool)
func (_SideSwapAgent *SideSwapAgentSession) FilledMainTx(arg0 [32]byte) (bool, error) {
	return _SideSwapAgent.Contract.FilledMainTx(&_SideSwapAgent.CallOpts, arg0)
}

// FilledMainTx is a free data retrieval call binding the contract method 0xf017a478.
//
// Solidity: function filledMainTx(bytes32 ) view returns(bool)
func (_SideSwapAgent *SideSwapAgentCallerSession) FilledMainTx(arg0 [32]byte) (bool, error) {
	return _SideSwapAgent.Contract.FilledMainTx(&_SideSwapAgent.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SideSwapAgent *SideSwapAgentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SideSwapAgent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SideSwapAgent *SideSwapAgentSession) Owner() (common.Address, error) {
	return _SideSwapAgent.Contract.Owner(&_SideSwapAgent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SideSwapAgent *SideSwapAgentCallerSession) Owner() (common.Address, error) {
	return _SideSwapAgent.Contract.Owner(&_SideSwapAgent.CallOpts)
}

// SideChainErc20Banlance is a free data retrieval call binding the contract method 0xcd9fa9b8.
//
// Solidity: function sideChainErc20Banlance(address ) view returns(uint256)
func (_SideSwapAgent *SideSwapAgentCaller) SideChainErc20Banlance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SideSwapAgent.contract.Call(opts, &out, "sideChainErc20Banlance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SideChainErc20Banlance is a free data retrieval call binding the contract method 0xcd9fa9b8.
//
// Solidity: function sideChainErc20Banlance(address ) view returns(uint256)
func (_SideSwapAgent *SideSwapAgentSession) SideChainErc20Banlance(arg0 common.Address) (*big.Int, error) {
	return _SideSwapAgent.Contract.SideChainErc20Banlance(&_SideSwapAgent.CallOpts, arg0)
}

// SideChainErc20Banlance is a free data retrieval call binding the contract method 0xcd9fa9b8.
//
// Solidity: function sideChainErc20Banlance(address ) view returns(uint256)
func (_SideSwapAgent *SideSwapAgentCallerSession) SideChainErc20Banlance(arg0 common.Address) (*big.Int, error) {
	return _SideSwapAgent.Contract.SideChainErc20Banlance(&_SideSwapAgent.CallOpts, arg0)
}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_SideSwapAgent *SideSwapAgentCaller) SwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SideSwapAgent.contract.Call(opts, &out, "swapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_SideSwapAgent *SideSwapAgentSession) SwapFee() (*big.Int, error) {
	return _SideSwapAgent.Contract.SwapFee(&_SideSwapAgent.CallOpts)
}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_SideSwapAgent *SideSwapAgentCallerSession) SwapFee() (*big.Int, error) {
	return _SideSwapAgent.Contract.SwapFee(&_SideSwapAgent.CallOpts)
}

// SwapMappingMain2Side is a free data retrieval call binding the contract method 0x95660c77.
//
// Solidity: function swapMappingMain2Side(address ) view returns(address)
func (_SideSwapAgent *SideSwapAgentCaller) SwapMappingMain2Side(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SideSwapAgent.contract.Call(opts, &out, "swapMappingMain2Side", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingMain2Side is a free data retrieval call binding the contract method 0x95660c77.
//
// Solidity: function swapMappingMain2Side(address ) view returns(address)
func (_SideSwapAgent *SideSwapAgentSession) SwapMappingMain2Side(arg0 common.Address) (common.Address, error) {
	return _SideSwapAgent.Contract.SwapMappingMain2Side(&_SideSwapAgent.CallOpts, arg0)
}

// SwapMappingMain2Side is a free data retrieval call binding the contract method 0x95660c77.
//
// Solidity: function swapMappingMain2Side(address ) view returns(address)
func (_SideSwapAgent *SideSwapAgentCallerSession) SwapMappingMain2Side(arg0 common.Address) (common.Address, error) {
	return _SideSwapAgent.Contract.SwapMappingMain2Side(&_SideSwapAgent.CallOpts, arg0)
}

// SwapMappingSide2Main is a free data retrieval call binding the contract method 0x1be17b65.
//
// Solidity: function swapMappingSide2Main(address ) view returns(address)
func (_SideSwapAgent *SideSwapAgentCaller) SwapMappingSide2Main(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _SideSwapAgent.contract.Call(opts, &out, "swapMappingSide2Main", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingSide2Main is a free data retrieval call binding the contract method 0x1be17b65.
//
// Solidity: function swapMappingSide2Main(address ) view returns(address)
func (_SideSwapAgent *SideSwapAgentSession) SwapMappingSide2Main(arg0 common.Address) (common.Address, error) {
	return _SideSwapAgent.Contract.SwapMappingSide2Main(&_SideSwapAgent.CallOpts, arg0)
}

// SwapMappingSide2Main is a free data retrieval call binding the contract method 0x1be17b65.
//
// Solidity: function swapMappingSide2Main(address ) view returns(address)
func (_SideSwapAgent *SideSwapAgentCallerSession) SwapMappingSide2Main(arg0 common.Address) (common.Address, error) {
	return _SideSwapAgent.Contract.SwapMappingSide2Main(&_SideSwapAgent.CallOpts, arg0)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x1d662f62.
//
// Solidity: function createSwapPair(bytes32 mainChainTxHash, address mainChainErc20Addr, address sideChainErc20Addr, string name, string symbol, uint8 decimals) returns(bool)
func (_SideSwapAgent *SideSwapAgentTransactor) CreateSwapPair(opts *bind.TransactOpts, mainChainTxHash [32]byte, mainChainErc20Addr common.Address, sideChainErc20Addr common.Address, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _SideSwapAgent.contract.Transact(opts, "createSwapPair", mainChainTxHash, mainChainErc20Addr, sideChainErc20Addr, name, symbol, decimals)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x1d662f62.
//
// Solidity: function createSwapPair(bytes32 mainChainTxHash, address mainChainErc20Addr, address sideChainErc20Addr, string name, string symbol, uint8 decimals) returns(bool)
func (_SideSwapAgent *SideSwapAgentSession) CreateSwapPair(mainChainTxHash [32]byte, mainChainErc20Addr common.Address, sideChainErc20Addr common.Address, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.CreateSwapPair(&_SideSwapAgent.TransactOpts, mainChainTxHash, mainChainErc20Addr, sideChainErc20Addr, name, symbol, decimals)
}

// CreateSwapPair is a paid mutator transaction binding the contract method 0x1d662f62.
//
// Solidity: function createSwapPair(bytes32 mainChainTxHash, address mainChainErc20Addr, address sideChainErc20Addr, string name, string symbol, uint8 decimals) returns(bool)
func (_SideSwapAgent *SideSwapAgentTransactorSession) CreateSwapPair(mainChainTxHash [32]byte, mainChainErc20Addr common.Address, sideChainErc20Addr common.Address, name string, symbol string, decimals uint8) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.CreateSwapPair(&_SideSwapAgent.TransactOpts, mainChainTxHash, mainChainErc20Addr, sideChainErc20Addr, name, symbol, decimals)
}

// FillMain2SideSwap is a paid mutator transaction binding the contract method 0x7b0bdf11.
//
// Solidity: function fillMain2SideSwap(bytes32 mainChainTxHash, address mainChainErc20Addr, address sideChainToAddr, uint256 amount) payable returns(bool)
func (_SideSwapAgent *SideSwapAgentTransactor) FillMain2SideSwap(opts *bind.TransactOpts, mainChainTxHash [32]byte, mainChainErc20Addr common.Address, sideChainToAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.contract.Transact(opts, "fillMain2SideSwap", mainChainTxHash, mainChainErc20Addr, sideChainToAddr, amount)
}

// FillMain2SideSwap is a paid mutator transaction binding the contract method 0x7b0bdf11.
//
// Solidity: function fillMain2SideSwap(bytes32 mainChainTxHash, address mainChainErc20Addr, address sideChainToAddr, uint256 amount) payable returns(bool)
func (_SideSwapAgent *SideSwapAgentSession) FillMain2SideSwap(mainChainTxHash [32]byte, mainChainErc20Addr common.Address, sideChainToAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.FillMain2SideSwap(&_SideSwapAgent.TransactOpts, mainChainTxHash, mainChainErc20Addr, sideChainToAddr, amount)
}

// FillMain2SideSwap is a paid mutator transaction binding the contract method 0x7b0bdf11.
//
// Solidity: function fillMain2SideSwap(bytes32 mainChainTxHash, address mainChainErc20Addr, address sideChainToAddr, uint256 amount) payable returns(bool)
func (_SideSwapAgent *SideSwapAgentTransactorSession) FillMain2SideSwap(mainChainTxHash [32]byte, mainChainErc20Addr common.Address, sideChainToAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.FillMain2SideSwap(&_SideSwapAgent.TransactOpts, mainChainTxHash, mainChainErc20Addr, sideChainToAddr, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address ownerAddr, uint256 fee) returns()
func (_SideSwapAgent *SideSwapAgentTransactor) Initialize(opts *bind.TransactOpts, ownerAddr common.Address, fee *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.contract.Transact(opts, "initialize", ownerAddr, fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address ownerAddr, uint256 fee) returns()
func (_SideSwapAgent *SideSwapAgentSession) Initialize(ownerAddr common.Address, fee *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.Initialize(&_SideSwapAgent.TransactOpts, ownerAddr, fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address ownerAddr, uint256 fee) returns()
func (_SideSwapAgent *SideSwapAgentTransactorSession) Initialize(ownerAddr common.Address, fee *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.Initialize(&_SideSwapAgent.TransactOpts, ownerAddr, fee)
}

// RechargeSide is a paid mutator transaction binding the contract method 0xbff281a3.
//
// Solidity: function rechargeSide(address sideChainErc20Addr, uint256 amount) payable returns(bool)
func (_SideSwapAgent *SideSwapAgentTransactor) RechargeSide(opts *bind.TransactOpts, sideChainErc20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.contract.Transact(opts, "rechargeSide", sideChainErc20Addr, amount)
}

// RechargeSide is a paid mutator transaction binding the contract method 0xbff281a3.
//
// Solidity: function rechargeSide(address sideChainErc20Addr, uint256 amount) payable returns(bool)
func (_SideSwapAgent *SideSwapAgentSession) RechargeSide(sideChainErc20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.RechargeSide(&_SideSwapAgent.TransactOpts, sideChainErc20Addr, amount)
}

// RechargeSide is a paid mutator transaction binding the contract method 0xbff281a3.
//
// Solidity: function rechargeSide(address sideChainErc20Addr, uint256 amount) payable returns(bool)
func (_SideSwapAgent *SideSwapAgentTransactorSession) RechargeSide(sideChainErc20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.RechargeSide(&_SideSwapAgent.TransactOpts, sideChainErc20Addr, amount)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SideSwapAgent *SideSwapAgentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SideSwapAgent.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SideSwapAgent *SideSwapAgentSession) RenounceOwnership() (*types.Transaction, error) {
	return _SideSwapAgent.Contract.RenounceOwnership(&_SideSwapAgent.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_SideSwapAgent *SideSwapAgentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SideSwapAgent.Contract.RenounceOwnership(&_SideSwapAgent.TransactOpts)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 fee) returns()
func (_SideSwapAgent *SideSwapAgentTransactor) SetSwapFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.contract.Transact(opts, "setSwapFee", fee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 fee) returns()
func (_SideSwapAgent *SideSwapAgentSession) SetSwapFee(fee *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.SetSwapFee(&_SideSwapAgent.TransactOpts, fee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 fee) returns()
func (_SideSwapAgent *SideSwapAgentTransactorSession) SetSwapFee(fee *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.SetSwapFee(&_SideSwapAgent.TransactOpts, fee)
}

// SwapSide2Main is a paid mutator transaction binding the contract method 0x0b762dac.
//
// Solidity: function swapSide2Main(address sideChainErc20Addr, uint256 amount) payable returns(bool)
func (_SideSwapAgent *SideSwapAgentTransactor) SwapSide2Main(opts *bind.TransactOpts, sideChainErc20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.contract.Transact(opts, "swapSide2Main", sideChainErc20Addr, amount)
}

// SwapSide2Main is a paid mutator transaction binding the contract method 0x0b762dac.
//
// Solidity: function swapSide2Main(address sideChainErc20Addr, uint256 amount) payable returns(bool)
func (_SideSwapAgent *SideSwapAgentSession) SwapSide2Main(sideChainErc20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.SwapSide2Main(&_SideSwapAgent.TransactOpts, sideChainErc20Addr, amount)
}

// SwapSide2Main is a paid mutator transaction binding the contract method 0x0b762dac.
//
// Solidity: function swapSide2Main(address sideChainErc20Addr, uint256 amount) payable returns(bool)
func (_SideSwapAgent *SideSwapAgentTransactorSession) SwapSide2Main(sideChainErc20Addr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.SwapSide2Main(&_SideSwapAgent.TransactOpts, sideChainErc20Addr, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SideSwapAgent *SideSwapAgentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SideSwapAgent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SideSwapAgent *SideSwapAgentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.TransferOwnership(&_SideSwapAgent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SideSwapAgent *SideSwapAgentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SideSwapAgent.Contract.TransferOwnership(&_SideSwapAgent.TransactOpts, newOwner)
}

// SideSwapAgentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SideSwapAgent contract.
type SideSwapAgentOwnershipTransferredIterator struct {
	Event *SideSwapAgentOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SideSwapAgentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SideSwapAgentOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SideSwapAgentOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SideSwapAgentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SideSwapAgentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SideSwapAgentOwnershipTransferred represents a OwnershipTransferred event raised by the SideSwapAgent contract.
type SideSwapAgentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SideSwapAgent *SideSwapAgentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*SideSwapAgentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SideSwapAgent.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SideSwapAgentOwnershipTransferredIterator{contract: _SideSwapAgent.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SideSwapAgent *SideSwapAgentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SideSwapAgentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SideSwapAgent.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SideSwapAgentOwnershipTransferred)
				if err := _SideSwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_SideSwapAgent *SideSwapAgentFilterer) ParseOwnershipTransferred(log types.Log) (*SideSwapAgentOwnershipTransferred, error) {
	event := new(SideSwapAgentOwnershipTransferred)
	if err := _SideSwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SideSwapAgentRechargeEventIterator is returned from FilterRechargeEvent and is used to iterate over the raw logs and unpacked data for RechargeEvent events raised by the SideSwapAgent contract.
type SideSwapAgentRechargeEventIterator struct {
	Event *SideSwapAgentRechargeEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SideSwapAgentRechargeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SideSwapAgentRechargeEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SideSwapAgentRechargeEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SideSwapAgentRechargeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SideSwapAgentRechargeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SideSwapAgentRechargeEvent represents a RechargeEvent event raised by the SideSwapAgent contract.
type SideSwapAgentRechargeEvent struct {
	SideChainErc20Addr common.Address
	SendAddr           common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRechargeEvent is a free log retrieval operation binding the contract event 0xba435e9191118b57fc3cd6841f3248e60b065acfea46e76f975249ade106b03b.
//
// Solidity: event RechargeEvent(address indexed sideChainErc20Addr, address indexed sendAddr, uint256 amount)
func (_SideSwapAgent *SideSwapAgentFilterer) FilterRechargeEvent(opts *bind.FilterOpts, sideChainErc20Addr []common.Address, sendAddr []common.Address) (*SideSwapAgentRechargeEventIterator, error) {

	var sideChainErc20AddrRule []interface{}
	for _, sideChainErc20AddrItem := range sideChainErc20Addr {
		sideChainErc20AddrRule = append(sideChainErc20AddrRule, sideChainErc20AddrItem)
	}
	var sendAddrRule []interface{}
	for _, sendAddrItem := range sendAddr {
		sendAddrRule = append(sendAddrRule, sendAddrItem)
	}

	logs, sub, err := _SideSwapAgent.contract.FilterLogs(opts, "RechargeEvent", sideChainErc20AddrRule, sendAddrRule)
	if err != nil {
		return nil, err
	}
	return &SideSwapAgentRechargeEventIterator{contract: _SideSwapAgent.contract, event: "RechargeEvent", logs: logs, sub: sub}, nil
}

// WatchRechargeEvent is a free log subscription operation binding the contract event 0xba435e9191118b57fc3cd6841f3248e60b065acfea46e76f975249ade106b03b.
//
// Solidity: event RechargeEvent(address indexed sideChainErc20Addr, address indexed sendAddr, uint256 amount)
func (_SideSwapAgent *SideSwapAgentFilterer) WatchRechargeEvent(opts *bind.WatchOpts, sink chan<- *SideSwapAgentRechargeEvent, sideChainErc20Addr []common.Address, sendAddr []common.Address) (event.Subscription, error) {

	var sideChainErc20AddrRule []interface{}
	for _, sideChainErc20AddrItem := range sideChainErc20Addr {
		sideChainErc20AddrRule = append(sideChainErc20AddrRule, sideChainErc20AddrItem)
	}
	var sendAddrRule []interface{}
	for _, sendAddrItem := range sendAddr {
		sendAddrRule = append(sendAddrRule, sendAddrItem)
	}

	logs, sub, err := _SideSwapAgent.contract.WatchLogs(opts, "RechargeEvent", sideChainErc20AddrRule, sendAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SideSwapAgentRechargeEvent)
				if err := _SideSwapAgent.contract.UnpackLog(event, "RechargeEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseRechargeEvent is a log parse operation binding the contract event 0xba435e9191118b57fc3cd6841f3248e60b065acfea46e76f975249ade106b03b.
//
// Solidity: event RechargeEvent(address indexed sideChainErc20Addr, address indexed sendAddr, uint256 amount)
func (_SideSwapAgent *SideSwapAgentFilterer) ParseRechargeEvent(log types.Log) (*SideSwapAgentRechargeEvent, error) {
	event := new(SideSwapAgentRechargeEvent)
	if err := _SideSwapAgent.contract.UnpackLog(event, "RechargeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SideSwapAgentSwapMain2SideFilledEventIterator is returned from FilterSwapMain2SideFilledEvent and is used to iterate over the raw logs and unpacked data for SwapMain2SideFilledEvent events raised by the SideSwapAgent contract.
type SideSwapAgentSwapMain2SideFilledEventIterator struct {
	Event *SideSwapAgentSwapMain2SideFilledEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SideSwapAgentSwapMain2SideFilledEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SideSwapAgentSwapMain2SideFilledEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SideSwapAgentSwapMain2SideFilledEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SideSwapAgentSwapMain2SideFilledEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SideSwapAgentSwapMain2SideFilledEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SideSwapAgentSwapMain2SideFilledEvent represents a SwapMain2SideFilledEvent event raised by the SideSwapAgent contract.
type SideSwapAgentSwapMain2SideFilledEvent struct {
	MainChainTxHash    [32]byte
	MainChainErc20Addr common.Address
	SideChainToAddr    common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwapMain2SideFilledEvent is a free log retrieval operation binding the contract event 0xcce98297f95dc1cec392b3533c9cb6329f064d83ea3ae3ad09fd3f2d968ec100.
//
// Solidity: event SwapMain2SideFilledEvent(bytes32 indexed mainChainTxHash, address indexed mainChainErc20Addr, address indexed sideChainToAddr, uint256 amount)
func (_SideSwapAgent *SideSwapAgentFilterer) FilterSwapMain2SideFilledEvent(opts *bind.FilterOpts, mainChainTxHash [][32]byte, mainChainErc20Addr []common.Address, sideChainToAddr []common.Address) (*SideSwapAgentSwapMain2SideFilledEventIterator, error) {

	var mainChainTxHashRule []interface{}
	for _, mainChainTxHashItem := range mainChainTxHash {
		mainChainTxHashRule = append(mainChainTxHashRule, mainChainTxHashItem)
	}
	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var sideChainToAddrRule []interface{}
	for _, sideChainToAddrItem := range sideChainToAddr {
		sideChainToAddrRule = append(sideChainToAddrRule, sideChainToAddrItem)
	}

	logs, sub, err := _SideSwapAgent.contract.FilterLogs(opts, "SwapMain2SideFilledEvent", mainChainTxHashRule, mainChainErc20AddrRule, sideChainToAddrRule)
	if err != nil {
		return nil, err
	}
	return &SideSwapAgentSwapMain2SideFilledEventIterator{contract: _SideSwapAgent.contract, event: "SwapMain2SideFilledEvent", logs: logs, sub: sub}, nil
}

// WatchSwapMain2SideFilledEvent is a free log subscription operation binding the contract event 0xcce98297f95dc1cec392b3533c9cb6329f064d83ea3ae3ad09fd3f2d968ec100.
//
// Solidity: event SwapMain2SideFilledEvent(bytes32 indexed mainChainTxHash, address indexed mainChainErc20Addr, address indexed sideChainToAddr, uint256 amount)
func (_SideSwapAgent *SideSwapAgentFilterer) WatchSwapMain2SideFilledEvent(opts *bind.WatchOpts, sink chan<- *SideSwapAgentSwapMain2SideFilledEvent, mainChainTxHash [][32]byte, mainChainErc20Addr []common.Address, sideChainToAddr []common.Address) (event.Subscription, error) {

	var mainChainTxHashRule []interface{}
	for _, mainChainTxHashItem := range mainChainTxHash {
		mainChainTxHashRule = append(mainChainTxHashRule, mainChainTxHashItem)
	}
	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var sideChainToAddrRule []interface{}
	for _, sideChainToAddrItem := range sideChainToAddr {
		sideChainToAddrRule = append(sideChainToAddrRule, sideChainToAddrItem)
	}

	logs, sub, err := _SideSwapAgent.contract.WatchLogs(opts, "SwapMain2SideFilledEvent", mainChainTxHashRule, mainChainErc20AddrRule, sideChainToAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SideSwapAgentSwapMain2SideFilledEvent)
				if err := _SideSwapAgent.contract.UnpackLog(event, "SwapMain2SideFilledEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapMain2SideFilledEvent is a log parse operation binding the contract event 0xcce98297f95dc1cec392b3533c9cb6329f064d83ea3ae3ad09fd3f2d968ec100.
//
// Solidity: event SwapMain2SideFilledEvent(bytes32 indexed mainChainTxHash, address indexed mainChainErc20Addr, address indexed sideChainToAddr, uint256 amount)
func (_SideSwapAgent *SideSwapAgentFilterer) ParseSwapMain2SideFilledEvent(log types.Log) (*SideSwapAgentSwapMain2SideFilledEvent, error) {
	event := new(SideSwapAgentSwapMain2SideFilledEvent)
	if err := _SideSwapAgent.contract.UnpackLog(event, "SwapMain2SideFilledEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SideSwapAgentSwapPairCreatedEventIterator is returned from FilterSwapPairCreatedEvent and is used to iterate over the raw logs and unpacked data for SwapPairCreatedEvent events raised by the SideSwapAgent contract.
type SideSwapAgentSwapPairCreatedEventIterator struct {
	Event *SideSwapAgentSwapPairCreatedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SideSwapAgentSwapPairCreatedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SideSwapAgentSwapPairCreatedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SideSwapAgentSwapPairCreatedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SideSwapAgentSwapPairCreatedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SideSwapAgentSwapPairCreatedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SideSwapAgentSwapPairCreatedEvent represents a SwapPairCreatedEvent event raised by the SideSwapAgent contract.
type SideSwapAgentSwapPairCreatedEvent struct {
	MainChainTxHash    [32]byte
	MainChainErc20Addr common.Address
	SideChainErc20Addr common.Address
	Name               string
	Symbol             string
	Decimals           uint8
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwapPairCreatedEvent is a free log retrieval operation binding the contract event 0xb04179acee3e3defc848704fc6453616f655decf86895f1cb479e98e75529609.
//
// Solidity: event SwapPairCreatedEvent(bytes32 indexed mainChainTxHash, address indexed mainChainErc20Addr, address indexed sideChainErc20Addr, string name, string symbol, uint8 decimals)
func (_SideSwapAgent *SideSwapAgentFilterer) FilterSwapPairCreatedEvent(opts *bind.FilterOpts, mainChainTxHash [][32]byte, mainChainErc20Addr []common.Address, sideChainErc20Addr []common.Address) (*SideSwapAgentSwapPairCreatedEventIterator, error) {

	var mainChainTxHashRule []interface{}
	for _, mainChainTxHashItem := range mainChainTxHash {
		mainChainTxHashRule = append(mainChainTxHashRule, mainChainTxHashItem)
	}
	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var sideChainErc20AddrRule []interface{}
	for _, sideChainErc20AddrItem := range sideChainErc20Addr {
		sideChainErc20AddrRule = append(sideChainErc20AddrRule, sideChainErc20AddrItem)
	}

	logs, sub, err := _SideSwapAgent.contract.FilterLogs(opts, "SwapPairCreatedEvent", mainChainTxHashRule, mainChainErc20AddrRule, sideChainErc20AddrRule)
	if err != nil {
		return nil, err
	}
	return &SideSwapAgentSwapPairCreatedEventIterator{contract: _SideSwapAgent.contract, event: "SwapPairCreatedEvent", logs: logs, sub: sub}, nil
}

// WatchSwapPairCreatedEvent is a free log subscription operation binding the contract event 0xb04179acee3e3defc848704fc6453616f655decf86895f1cb479e98e75529609.
//
// Solidity: event SwapPairCreatedEvent(bytes32 indexed mainChainTxHash, address indexed mainChainErc20Addr, address indexed sideChainErc20Addr, string name, string symbol, uint8 decimals)
func (_SideSwapAgent *SideSwapAgentFilterer) WatchSwapPairCreatedEvent(opts *bind.WatchOpts, sink chan<- *SideSwapAgentSwapPairCreatedEvent, mainChainTxHash [][32]byte, mainChainErc20Addr []common.Address, sideChainErc20Addr []common.Address) (event.Subscription, error) {

	var mainChainTxHashRule []interface{}
	for _, mainChainTxHashItem := range mainChainTxHash {
		mainChainTxHashRule = append(mainChainTxHashRule, mainChainTxHashItem)
	}
	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var sideChainErc20AddrRule []interface{}
	for _, sideChainErc20AddrItem := range sideChainErc20Addr {
		sideChainErc20AddrRule = append(sideChainErc20AddrRule, sideChainErc20AddrItem)
	}

	logs, sub, err := _SideSwapAgent.contract.WatchLogs(opts, "SwapPairCreatedEvent", mainChainTxHashRule, mainChainErc20AddrRule, sideChainErc20AddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SideSwapAgentSwapPairCreatedEvent)
				if err := _SideSwapAgent.contract.UnpackLog(event, "SwapPairCreatedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapPairCreatedEvent is a log parse operation binding the contract event 0xb04179acee3e3defc848704fc6453616f655decf86895f1cb479e98e75529609.
//
// Solidity: event SwapPairCreatedEvent(bytes32 indexed mainChainTxHash, address indexed mainChainErc20Addr, address indexed sideChainErc20Addr, string name, string symbol, uint8 decimals)
func (_SideSwapAgent *SideSwapAgentFilterer) ParseSwapPairCreatedEvent(log types.Log) (*SideSwapAgentSwapPairCreatedEvent, error) {
	event := new(SideSwapAgentSwapPairCreatedEvent)
	if err := _SideSwapAgent.contract.UnpackLog(event, "SwapPairCreatedEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SideSwapAgentSwapSide2MainEventIterator is returned from FilterSwapSide2MainEvent and is used to iterate over the raw logs and unpacked data for SwapSide2MainEvent events raised by the SideSwapAgent contract.
type SideSwapAgentSwapSide2MainEventIterator struct {
	Event *SideSwapAgentSwapSide2MainEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *SideSwapAgentSwapSide2MainEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SideSwapAgentSwapSide2MainEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(SideSwapAgentSwapSide2MainEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *SideSwapAgentSwapSide2MainEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SideSwapAgentSwapSide2MainEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SideSwapAgentSwapSide2MainEvent represents a SwapSide2MainEvent event raised by the SideSwapAgent contract.
type SideSwapAgentSwapSide2MainEvent struct {
	SideChainErc20Addr common.Address
	MainChainErc20Addr common.Address
	FromAddr           common.Address
	Amount             *big.Int
	FeeAmount          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwapSide2MainEvent is a free log retrieval operation binding the contract event 0xf19f19fa7d4ae2e098cbb45838aa7151ba4813912afa00a24e5c0a71599ee948.
//
// Solidity: event SwapSide2MainEvent(address indexed sideChainErc20Addr, address indexed mainChainErc20Addr, address indexed fromAddr, uint256 amount, uint256 feeAmount)
func (_SideSwapAgent *SideSwapAgentFilterer) FilterSwapSide2MainEvent(opts *bind.FilterOpts, sideChainErc20Addr []common.Address, mainChainErc20Addr []common.Address, fromAddr []common.Address) (*SideSwapAgentSwapSide2MainEventIterator, error) {

	var sideChainErc20AddrRule []interface{}
	for _, sideChainErc20AddrItem := range sideChainErc20Addr {
		sideChainErc20AddrRule = append(sideChainErc20AddrRule, sideChainErc20AddrItem)
	}
	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var fromAddrRule []interface{}
	for _, fromAddrItem := range fromAddr {
		fromAddrRule = append(fromAddrRule, fromAddrItem)
	}

	logs, sub, err := _SideSwapAgent.contract.FilterLogs(opts, "SwapSide2MainEvent", sideChainErc20AddrRule, mainChainErc20AddrRule, fromAddrRule)
	if err != nil {
		return nil, err
	}
	return &SideSwapAgentSwapSide2MainEventIterator{contract: _SideSwapAgent.contract, event: "SwapSide2MainEvent", logs: logs, sub: sub}, nil
}

// WatchSwapSide2MainEvent is a free log subscription operation binding the contract event 0xf19f19fa7d4ae2e098cbb45838aa7151ba4813912afa00a24e5c0a71599ee948.
//
// Solidity: event SwapSide2MainEvent(address indexed sideChainErc20Addr, address indexed mainChainErc20Addr, address indexed fromAddr, uint256 amount, uint256 feeAmount)
func (_SideSwapAgent *SideSwapAgentFilterer) WatchSwapSide2MainEvent(opts *bind.WatchOpts, sink chan<- *SideSwapAgentSwapSide2MainEvent, sideChainErc20Addr []common.Address, mainChainErc20Addr []common.Address, fromAddr []common.Address) (event.Subscription, error) {

	var sideChainErc20AddrRule []interface{}
	for _, sideChainErc20AddrItem := range sideChainErc20Addr {
		sideChainErc20AddrRule = append(sideChainErc20AddrRule, sideChainErc20AddrItem)
	}
	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var fromAddrRule []interface{}
	for _, fromAddrItem := range fromAddr {
		fromAddrRule = append(fromAddrRule, fromAddrItem)
	}

	logs, sub, err := _SideSwapAgent.contract.WatchLogs(opts, "SwapSide2MainEvent", sideChainErc20AddrRule, mainChainErc20AddrRule, fromAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SideSwapAgentSwapSide2MainEvent)
				if err := _SideSwapAgent.contract.UnpackLog(event, "SwapSide2MainEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseSwapSide2MainEvent is a log parse operation binding the contract event 0xf19f19fa7d4ae2e098cbb45838aa7151ba4813912afa00a24e5c0a71599ee948.
//
// Solidity: event SwapSide2MainEvent(address indexed sideChainErc20Addr, address indexed mainChainErc20Addr, address indexed fromAddr, uint256 amount, uint256 feeAmount)
func (_SideSwapAgent *SideSwapAgentFilterer) ParseSwapSide2MainEvent(log types.Log) (*SideSwapAgentSwapSide2MainEvent, error) {
	event := new(SideSwapAgentSwapSide2MainEvent)
	if err := _SideSwapAgent.contract.UnpackLog(event, "SwapSide2MainEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
