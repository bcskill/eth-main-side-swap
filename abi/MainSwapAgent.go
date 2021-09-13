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

// MainSwapAgentMetaData contains all meta data concerning the MainSwapAgent contract.
var MainSwapAgentMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mainChainErc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sendAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"RechargeEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mainChainErc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sideChainToAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeAmount\",\"type\":\"uint256\"}],\"name\":\"SwapMain2SideEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sponsor\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mainChainErc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sideChainErc20Addr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"SwapPairRegisterEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mainChainErc20Addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"sideChainTxHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"mainChainToAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"SwapSide2MainFilledEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"sideChainTxHash\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"sideChainErc20Addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"mainChainToAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"fillSide2MainSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"filledSideTx\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"ownerAddr\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"mainChainErc20Banlance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mainChainErc20Addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sideChainErc20Addr\",\"type\":\"address\"}],\"name\":\"registerSwapPairToSide\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"registeredMain\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"swapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"mainChainErc20Addr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sideChainToAddr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"swapMain2Side\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapMappingMain2Side\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"swapMappingSide2Main\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506124a2806100206000396000f3fe6080604052600436106100dd5760003560e01c80638da5cb5b1161007f578063da35a26f11610059578063da35a26f146104b3578063e62f4e671461050e578063f188e3c614610577578063f2fde38b14610600576100dd565b80638da5cb5b1461037857806395660c77146103cf578063cd6c8ef814610460576100dd565b806334e19907116100bb57806334e199071461027557806342e2de4e146102b057806354cf2aeb14610336578063715018a614610361576100dd565b80631be17b65146100e25780631fa0bb7c146101735780632a88193814610210575b600080fd5b3480156100ee57600080fd5b506101316004803603602081101561010557600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610651565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561017f57600080fd5b506101f66004803603608081101561019657600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610684565b604051808215151515815260200191505060405180910390f35b34801561021c57600080fd5b5061025f6004803603602081101561023357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610c3d565b6040518082815260200191505060405180910390f35b34801561028157600080fd5b506102ae6004803603602081101561029857600080fd5b8101908080359060200190929190505050610c55565b005b61031c600480360360608110156102c657600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610d29565b604051808215151515815260200191505060405180910390f35b34801561034257600080fd5b5061034b61147d565b6040518082815260200191505060405180910390f35b34801561036d57600080fd5b50610376611483565b005b34801561038457600080fd5b5061038d61160e565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103db57600080fd5b5061041e600480360360208110156103f257600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611634565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561046c57600080fd5b506104996004803603602081101561048357600080fd5b8101908080359060200190929190505050611667565b604051808215151515815260200191505060405180910390f35b3480156104bf57600080fd5b5061050c600480360360408110156104d657600080fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611687565b005b34801561051a57600080fd5b5061055d6004803603602081101561053157600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116d3565b604051808215151515815260200191505060405180910390f35b34801561058357600080fd5b506105e66004803603604081101561059a57600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506116f3565b604051808215151515815260200191505060405180910390f35b34801561060c57600080fd5b5061064f6004803603602081101561062357600080fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050612064565b005b60026020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600061068e612274565b73ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610750576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6004600086815260200190815260200160002060009054906101000a900460ff16156107e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260168152602001807f736964652074782066696c6c656420616c72656164790000000000000000000081525060200191505060405180910390fd5b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156108eb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f6e6f2073776170207061697220666f72207468697320746f6b656e000000000081525060200191505060405180910390fd5b828173ffffffffffffffffffffffffffffffffffffffff166370a08231306040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b15801561096957600080fd5b505afa15801561097d573d6000803e3d6000fd5b505050506040513d602081101561099357600080fd5b8101908080519060200190929190505050101580156109f1575082600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205410155b610a46576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806124486025913960400191505060405180910390fd5b60016004600088815260200190815260200160002060006101000a81548160ff0219169083151502179055508073ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85856040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050602060405180830381600087803b158015610af957600080fd5b505af1158015610b0d573d6000803e3d6000fd5b505050506040513d6020811015610b2357600080fd5b810190808051906020019092919050505050610b8783600560008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000205461227c90919063ffffffff16565b600560008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508373ffffffffffffffffffffffffffffffffffffffff16868273ffffffffffffffffffffffffffffffffffffffff167f8a6ec552fee56a817c37275213226a383c891ccdc5b251d8dae0da50543e623d866040518082815260200191505060405180910390a46001915050949350505050565b60056020528060005260406000206000915090505481565b610c5d612274565b73ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1614610d1f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b8060078190555050565b6000610d34336122c6565b15610da7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f636f6e7472616374206973206e6f7420616c6c6f77656420746f20737761700081525060200191505060405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610e48576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601c8152602001807f6e6f2070726f787920636f6e747261637420697320616c6c6f7765640000000081525060200191505060405180910390fd5b6000600160008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610f4f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f6e6f2073776170207061697220666f72207468697320746f6b656e000000000081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161415610ff2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260188152602001807f73696465436861696e546f41646472206973206572726f72000000000000000081525060200191505060405180910390fd5b60008311611068576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f54686520726563686172676520616d6f756e7420697320746f6f20736d616c6c81525060200191505060405180910390fd5b828573ffffffffffffffffffffffffffffffffffffffff166370a08231336040518263ffffffff1660e01b8152600401808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060206040518083038186803b1580156110e657600080fd5b505afa1580156110fa573d6000803e3d6000fd5b505050506040513d602081101561111057600080fd5b81019080805190602001909291905050501015611178576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806124486025913960400191505060405180910390fd5b60075434146111ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260128152602001807f7377617020666565206e6f7420657175616c000000000000000000000000000081525060200191505060405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff166323b872dd3330866040518463ffffffff1660e01b8152600401808473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018281526020019350505050602060405180830381600087803b1580156112aa57600080fd5b505af11580156112be573d6000803e3d6000fd5b505050506040513d60208110156112d457600080fd5b81019080805190602001909291905050505061133883600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020546122d990919063ffffffff16565b600560008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002081905550600034146113ed57600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f193505050501580156113eb573d6000803e3d6000fd5b505b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f42733bad7741b9044a0b058e7d8038dc72cdd877c6eec0d1db3b57f02be7eb818634604051808381526020018281526020019250505060405180910390a460019150509392505050565b60075481565b61148b612274565b73ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461154d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60016020528060005260406000206000915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60046020528060005260406000206000915054906101000a900460ff1681565b8160078190555080600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b60036020528060005260406000206000915054906101000a900460ff1681565b60006116fd612274565b73ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146117bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600160008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146118c0576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f6475706c696361746564206d61696e20636861696e207377617020706169720081525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16146119c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f6475706c696361746564207369646520636861696e202073776170207061697281525060200191505060405180910390fd5b60608373ffffffffffffffffffffffffffffffffffffffff166306fdde036040518163ffffffff1660e01b815260040160006040518083038186803b158015611a0957600080fd5b505afa158015611a1d573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015611a4757600080fd5b8101908080516040519392919084640100000000821115611a6757600080fd5b83820191506020820185811115611a7d57600080fd5b8251866001820283011164010000000082111715611a9a57600080fd5b8083526020830192505050908051906020019080838360005b83811015611ace578082015181840152602081019050611ab3565b50505050905090810190601f168015611afb5780820380516001836020036101000a031916815260200191505b50604052505050905060608473ffffffffffffffffffffffffffffffffffffffff166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b158015611b4c57600080fd5b505afa158015611b60573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015611b8a57600080fd5b8101908080516040519392919084640100000000821115611baa57600080fd5b83820191506020820185811115611bc057600080fd5b8251866001820283011164010000000082111715611bdd57600080fd5b8083526020830192505050908051906020019080838360005b83811015611c11578082015181840152602081019050611bf6565b50505050905090810190601f168015611c3e5780820380516001836020036101000a031916815260200191505b50604052505050905060008573ffffffffffffffffffffffffffffffffffffffff1663313ce5676040518163ffffffff1660e01b815260040160206040518083038186803b158015611c8f57600080fd5b505afa158015611ca3573d6000803e3d6000fd5b505050506040513d6020811015611cb957600080fd5b810190808051906020019092919050505090506000835111611d43576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600a8152602001807f656d707479206e616d650000000000000000000000000000000000000000000081525060200191505060405180910390fd5b6000825111611dba576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252600c8152602001807f656d7074792073796d626f6c000000000000000000000000000000000000000081525060200191505060405180910390fd5b84600160008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600260008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600560008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055508473ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff167f84938b4885411ca7a208fde741a0feec069b8173eef6d83fcdecc5daf3a2af3d8686866040518080602001806020018460ff1660ff168152602001838103835286818151815260200191508051906020019080838360005b83811015611fb4578082015181840152602081019050611f99565b50505050905090810190601f168015611fe15780820380516001836020036101000a031916815260200191505b50838103825285818151815260200191508051906020019080838360005b8381101561201a578082015181840152602081019050611fff565b50505050905090810190601f1680156120475780820380516001836020036101000a031916815260200191505b509550505050505060405180910390a46001935050505092915050565b61206c612274565b73ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161461212e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156121b4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001806124226026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600033905090565b60006122be83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250612361565b905092915050565b600080823b905060008111915050919050565b600080828401905083811015612357576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081525060200191505060405180910390fd5b8091505092915050565b600083831115829061240e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b838110156123d35780820151818401526020810190506123b8565b50505050905090810190601f1680156124005780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838503905080915050939250505056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373496e73756666696369656e7420636f6e7472616374206163636f756e742062616c616e6365a264697066735822122009e6dabcc8f3ae30c36c6e115dbff38f51b36fdf7775c6bd57019d52a684b3f664736f6c63430006040033",
}

// MainSwapAgentABI is the input ABI used to generate the binding from.
// Deprecated: Use MainSwapAgentMetaData.ABI instead.
var MainSwapAgentABI = MainSwapAgentMetaData.ABI

// MainSwapAgentBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use MainSwapAgentMetaData.Bin instead.
var MainSwapAgentBin = MainSwapAgentMetaData.Bin

// DeployMainSwapAgent deploys a new Ethereum contract, binding an instance of MainSwapAgent to it.
func DeployMainSwapAgent(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *MainSwapAgent, error) {
	parsed, err := MainSwapAgentMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(MainSwapAgentBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &MainSwapAgent{MainSwapAgentCaller: MainSwapAgentCaller{contract: contract}, MainSwapAgentTransactor: MainSwapAgentTransactor{contract: contract}, MainSwapAgentFilterer: MainSwapAgentFilterer{contract: contract}}, nil
}

// MainSwapAgent is an auto generated Go binding around an Ethereum contract.
type MainSwapAgent struct {
	MainSwapAgentCaller     // Read-only binding to the contract
	MainSwapAgentTransactor // Write-only binding to the contract
	MainSwapAgentFilterer   // Log filterer for contract events
}

// MainSwapAgentCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainSwapAgentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSwapAgentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainSwapAgentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSwapAgentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainSwapAgentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSwapAgentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSwapAgentSession struct {
	Contract     *MainSwapAgent    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainSwapAgentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainSwapAgentCallerSession struct {
	Contract *MainSwapAgentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// MainSwapAgentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainSwapAgentTransactorSession struct {
	Contract     *MainSwapAgentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// MainSwapAgentRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainSwapAgentRaw struct {
	Contract *MainSwapAgent // Generic contract binding to access the raw methods on
}

// MainSwapAgentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainSwapAgentCallerRaw struct {
	Contract *MainSwapAgentCaller // Generic read-only contract binding to access the raw methods on
}

// MainSwapAgentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainSwapAgentTransactorRaw struct {
	Contract *MainSwapAgentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMainSwapAgent creates a new instance of MainSwapAgent, bound to a specific deployed contract.
func NewMainSwapAgent(address common.Address, backend bind.ContractBackend) (*MainSwapAgent, error) {
	contract, err := bindMainSwapAgent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &MainSwapAgent{MainSwapAgentCaller: MainSwapAgentCaller{contract: contract}, MainSwapAgentTransactor: MainSwapAgentTransactor{contract: contract}, MainSwapAgentFilterer: MainSwapAgentFilterer{contract: contract}}, nil
}

// NewMainSwapAgentCaller creates a new read-only instance of MainSwapAgent, bound to a specific deployed contract.
func NewMainSwapAgentCaller(address common.Address, caller bind.ContractCaller) (*MainSwapAgentCaller, error) {
	contract, err := bindMainSwapAgent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainSwapAgentCaller{contract: contract}, nil
}

// NewMainSwapAgentTransactor creates a new write-only instance of MainSwapAgent, bound to a specific deployed contract.
func NewMainSwapAgentTransactor(address common.Address, transactor bind.ContractTransactor) (*MainSwapAgentTransactor, error) {
	contract, err := bindMainSwapAgent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainSwapAgentTransactor{contract: contract}, nil
}

// NewMainSwapAgentFilterer creates a new log filterer instance of MainSwapAgent, bound to a specific deployed contract.
func NewMainSwapAgentFilterer(address common.Address, filterer bind.ContractFilterer) (*MainSwapAgentFilterer, error) {
	contract, err := bindMainSwapAgent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainSwapAgentFilterer{contract: contract}, nil
}

// bindMainSwapAgent binds a generic wrapper to an already deployed contract.
func bindMainSwapAgent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainSwapAgentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainSwapAgent *MainSwapAgentRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainSwapAgent.Contract.MainSwapAgentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainSwapAgent *MainSwapAgentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.MainSwapAgentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainSwapAgent *MainSwapAgentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.MainSwapAgentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_MainSwapAgent *MainSwapAgentCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _MainSwapAgent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_MainSwapAgent *MainSwapAgentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_MainSwapAgent *MainSwapAgentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.contract.Transact(opts, method, params...)
}

// FilledSideTx is a free data retrieval call binding the contract method 0xcd6c8ef8.
//
// Solidity: function filledSideTx(bytes32 ) view returns(bool)
func (_MainSwapAgent *MainSwapAgentCaller) FilledSideTx(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _MainSwapAgent.contract.Call(opts, &out, "filledSideTx", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FilledSideTx is a free data retrieval call binding the contract method 0xcd6c8ef8.
//
// Solidity: function filledSideTx(bytes32 ) view returns(bool)
func (_MainSwapAgent *MainSwapAgentSession) FilledSideTx(arg0 [32]byte) (bool, error) {
	return _MainSwapAgent.Contract.FilledSideTx(&_MainSwapAgent.CallOpts, arg0)
}

// FilledSideTx is a free data retrieval call binding the contract method 0xcd6c8ef8.
//
// Solidity: function filledSideTx(bytes32 ) view returns(bool)
func (_MainSwapAgent *MainSwapAgentCallerSession) FilledSideTx(arg0 [32]byte) (bool, error) {
	return _MainSwapAgent.Contract.FilledSideTx(&_MainSwapAgent.CallOpts, arg0)
}

// MainChainErc20Banlance is a free data retrieval call binding the contract method 0x2a881938.
//
// Solidity: function mainChainErc20Banlance(address ) view returns(uint256)
func (_MainSwapAgent *MainSwapAgentCaller) MainChainErc20Banlance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _MainSwapAgent.contract.Call(opts, &out, "mainChainErc20Banlance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MainChainErc20Banlance is a free data retrieval call binding the contract method 0x2a881938.
//
// Solidity: function mainChainErc20Banlance(address ) view returns(uint256)
func (_MainSwapAgent *MainSwapAgentSession) MainChainErc20Banlance(arg0 common.Address) (*big.Int, error) {
	return _MainSwapAgent.Contract.MainChainErc20Banlance(&_MainSwapAgent.CallOpts, arg0)
}

// MainChainErc20Banlance is a free data retrieval call binding the contract method 0x2a881938.
//
// Solidity: function mainChainErc20Banlance(address ) view returns(uint256)
func (_MainSwapAgent *MainSwapAgentCallerSession) MainChainErc20Banlance(arg0 common.Address) (*big.Int, error) {
	return _MainSwapAgent.Contract.MainChainErc20Banlance(&_MainSwapAgent.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MainSwapAgent *MainSwapAgentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _MainSwapAgent.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MainSwapAgent *MainSwapAgentSession) Owner() (common.Address, error) {
	return _MainSwapAgent.Contract.Owner(&_MainSwapAgent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_MainSwapAgent *MainSwapAgentCallerSession) Owner() (common.Address, error) {
	return _MainSwapAgent.Contract.Owner(&_MainSwapAgent.CallOpts)
}

// RegisteredMain is a free data retrieval call binding the contract method 0xe62f4e67.
//
// Solidity: function registeredMain(address ) view returns(bool)
func (_MainSwapAgent *MainSwapAgentCaller) RegisteredMain(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _MainSwapAgent.contract.Call(opts, &out, "registeredMain", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// RegisteredMain is a free data retrieval call binding the contract method 0xe62f4e67.
//
// Solidity: function registeredMain(address ) view returns(bool)
func (_MainSwapAgent *MainSwapAgentSession) RegisteredMain(arg0 common.Address) (bool, error) {
	return _MainSwapAgent.Contract.RegisteredMain(&_MainSwapAgent.CallOpts, arg0)
}

// RegisteredMain is a free data retrieval call binding the contract method 0xe62f4e67.
//
// Solidity: function registeredMain(address ) view returns(bool)
func (_MainSwapAgent *MainSwapAgentCallerSession) RegisteredMain(arg0 common.Address) (bool, error) {
	return _MainSwapAgent.Contract.RegisteredMain(&_MainSwapAgent.CallOpts, arg0)
}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_MainSwapAgent *MainSwapAgentCaller) SwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _MainSwapAgent.contract.Call(opts, &out, "swapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_MainSwapAgent *MainSwapAgentSession) SwapFee() (*big.Int, error) {
	return _MainSwapAgent.Contract.SwapFee(&_MainSwapAgent.CallOpts)
}

// SwapFee is a free data retrieval call binding the contract method 0x54cf2aeb.
//
// Solidity: function swapFee() view returns(uint256)
func (_MainSwapAgent *MainSwapAgentCallerSession) SwapFee() (*big.Int, error) {
	return _MainSwapAgent.Contract.SwapFee(&_MainSwapAgent.CallOpts)
}

// SwapMappingMain2Side is a free data retrieval call binding the contract method 0x95660c77.
//
// Solidity: function swapMappingMain2Side(address ) view returns(address)
func (_MainSwapAgent *MainSwapAgentCaller) SwapMappingMain2Side(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MainSwapAgent.contract.Call(opts, &out, "swapMappingMain2Side", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingMain2Side is a free data retrieval call binding the contract method 0x95660c77.
//
// Solidity: function swapMappingMain2Side(address ) view returns(address)
func (_MainSwapAgent *MainSwapAgentSession) SwapMappingMain2Side(arg0 common.Address) (common.Address, error) {
	return _MainSwapAgent.Contract.SwapMappingMain2Side(&_MainSwapAgent.CallOpts, arg0)
}

// SwapMappingMain2Side is a free data retrieval call binding the contract method 0x95660c77.
//
// Solidity: function swapMappingMain2Side(address ) view returns(address)
func (_MainSwapAgent *MainSwapAgentCallerSession) SwapMappingMain2Side(arg0 common.Address) (common.Address, error) {
	return _MainSwapAgent.Contract.SwapMappingMain2Side(&_MainSwapAgent.CallOpts, arg0)
}

// SwapMappingSide2Main is a free data retrieval call binding the contract method 0x1be17b65.
//
// Solidity: function swapMappingSide2Main(address ) view returns(address)
func (_MainSwapAgent *MainSwapAgentCaller) SwapMappingSide2Main(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _MainSwapAgent.contract.Call(opts, &out, "swapMappingSide2Main", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SwapMappingSide2Main is a free data retrieval call binding the contract method 0x1be17b65.
//
// Solidity: function swapMappingSide2Main(address ) view returns(address)
func (_MainSwapAgent *MainSwapAgentSession) SwapMappingSide2Main(arg0 common.Address) (common.Address, error) {
	return _MainSwapAgent.Contract.SwapMappingSide2Main(&_MainSwapAgent.CallOpts, arg0)
}

// SwapMappingSide2Main is a free data retrieval call binding the contract method 0x1be17b65.
//
// Solidity: function swapMappingSide2Main(address ) view returns(address)
func (_MainSwapAgent *MainSwapAgentCallerSession) SwapMappingSide2Main(arg0 common.Address) (common.Address, error) {
	return _MainSwapAgent.Contract.SwapMappingSide2Main(&_MainSwapAgent.CallOpts, arg0)
}

// FillSide2MainSwap is a paid mutator transaction binding the contract method 0x1fa0bb7c.
//
// Solidity: function fillSide2MainSwap(bytes32 sideChainTxHash, address sideChainErc20Addr, address mainChainToAddr, uint256 amount) returns(bool)
func (_MainSwapAgent *MainSwapAgentTransactor) FillSide2MainSwap(opts *bind.TransactOpts, sideChainTxHash [32]byte, sideChainErc20Addr common.Address, mainChainToAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MainSwapAgent.contract.Transact(opts, "fillSide2MainSwap", sideChainTxHash, sideChainErc20Addr, mainChainToAddr, amount)
}

// FillSide2MainSwap is a paid mutator transaction binding the contract method 0x1fa0bb7c.
//
// Solidity: function fillSide2MainSwap(bytes32 sideChainTxHash, address sideChainErc20Addr, address mainChainToAddr, uint256 amount) returns(bool)
func (_MainSwapAgent *MainSwapAgentSession) FillSide2MainSwap(sideChainTxHash [32]byte, sideChainErc20Addr common.Address, mainChainToAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.FillSide2MainSwap(&_MainSwapAgent.TransactOpts, sideChainTxHash, sideChainErc20Addr, mainChainToAddr, amount)
}

// FillSide2MainSwap is a paid mutator transaction binding the contract method 0x1fa0bb7c.
//
// Solidity: function fillSide2MainSwap(bytes32 sideChainTxHash, address sideChainErc20Addr, address mainChainToAddr, uint256 amount) returns(bool)
func (_MainSwapAgent *MainSwapAgentTransactorSession) FillSide2MainSwap(sideChainTxHash [32]byte, sideChainErc20Addr common.Address, mainChainToAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.FillSide2MainSwap(&_MainSwapAgent.TransactOpts, sideChainTxHash, sideChainErc20Addr, mainChainToAddr, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 fee, address ownerAddr) returns()
func (_MainSwapAgent *MainSwapAgentTransactor) Initialize(opts *bind.TransactOpts, fee *big.Int, ownerAddr common.Address) (*types.Transaction, error) {
	return _MainSwapAgent.contract.Transact(opts, "initialize", fee, ownerAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 fee, address ownerAddr) returns()
func (_MainSwapAgent *MainSwapAgentSession) Initialize(fee *big.Int, ownerAddr common.Address) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.Initialize(&_MainSwapAgent.TransactOpts, fee, ownerAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(uint256 fee, address ownerAddr) returns()
func (_MainSwapAgent *MainSwapAgentTransactorSession) Initialize(fee *big.Int, ownerAddr common.Address) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.Initialize(&_MainSwapAgent.TransactOpts, fee, ownerAddr)
}

// RegisterSwapPairToSide is a paid mutator transaction binding the contract method 0xf188e3c6.
//
// Solidity: function registerSwapPairToSide(address mainChainErc20Addr, address sideChainErc20Addr) returns(bool)
func (_MainSwapAgent *MainSwapAgentTransactor) RegisterSwapPairToSide(opts *bind.TransactOpts, mainChainErc20Addr common.Address, sideChainErc20Addr common.Address) (*types.Transaction, error) {
	return _MainSwapAgent.contract.Transact(opts, "registerSwapPairToSide", mainChainErc20Addr, sideChainErc20Addr)
}

// RegisterSwapPairToSide is a paid mutator transaction binding the contract method 0xf188e3c6.
//
// Solidity: function registerSwapPairToSide(address mainChainErc20Addr, address sideChainErc20Addr) returns(bool)
func (_MainSwapAgent *MainSwapAgentSession) RegisterSwapPairToSide(mainChainErc20Addr common.Address, sideChainErc20Addr common.Address) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.RegisterSwapPairToSide(&_MainSwapAgent.TransactOpts, mainChainErc20Addr, sideChainErc20Addr)
}

// RegisterSwapPairToSide is a paid mutator transaction binding the contract method 0xf188e3c6.
//
// Solidity: function registerSwapPairToSide(address mainChainErc20Addr, address sideChainErc20Addr) returns(bool)
func (_MainSwapAgent *MainSwapAgentTransactorSession) RegisterSwapPairToSide(mainChainErc20Addr common.Address, sideChainErc20Addr common.Address) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.RegisterSwapPairToSide(&_MainSwapAgent.TransactOpts, mainChainErc20Addr, sideChainErc20Addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MainSwapAgent *MainSwapAgentTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _MainSwapAgent.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MainSwapAgent *MainSwapAgentSession) RenounceOwnership() (*types.Transaction, error) {
	return _MainSwapAgent.Contract.RenounceOwnership(&_MainSwapAgent.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_MainSwapAgent *MainSwapAgentTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _MainSwapAgent.Contract.RenounceOwnership(&_MainSwapAgent.TransactOpts)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 fee) returns()
func (_MainSwapAgent *MainSwapAgentTransactor) SetSwapFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _MainSwapAgent.contract.Transact(opts, "setSwapFee", fee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 fee) returns()
func (_MainSwapAgent *MainSwapAgentSession) SetSwapFee(fee *big.Int) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.SetSwapFee(&_MainSwapAgent.TransactOpts, fee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 fee) returns()
func (_MainSwapAgent *MainSwapAgentTransactorSession) SetSwapFee(fee *big.Int) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.SetSwapFee(&_MainSwapAgent.TransactOpts, fee)
}

// SwapMain2Side is a paid mutator transaction binding the contract method 0x42e2de4e.
//
// Solidity: function swapMain2Side(address mainChainErc20Addr, address sideChainToAddr, uint256 amount) payable returns(bool)
func (_MainSwapAgent *MainSwapAgentTransactor) SwapMain2Side(opts *bind.TransactOpts, mainChainErc20Addr common.Address, sideChainToAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MainSwapAgent.contract.Transact(opts, "swapMain2Side", mainChainErc20Addr, sideChainToAddr, amount)
}

// SwapMain2Side is a paid mutator transaction binding the contract method 0x42e2de4e.
//
// Solidity: function swapMain2Side(address mainChainErc20Addr, address sideChainToAddr, uint256 amount) payable returns(bool)
func (_MainSwapAgent *MainSwapAgentSession) SwapMain2Side(mainChainErc20Addr common.Address, sideChainToAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.SwapMain2Side(&_MainSwapAgent.TransactOpts, mainChainErc20Addr, sideChainToAddr, amount)
}

// SwapMain2Side is a paid mutator transaction binding the contract method 0x42e2de4e.
//
// Solidity: function swapMain2Side(address mainChainErc20Addr, address sideChainToAddr, uint256 amount) payable returns(bool)
func (_MainSwapAgent *MainSwapAgentTransactorSession) SwapMain2Side(mainChainErc20Addr common.Address, sideChainToAddr common.Address, amount *big.Int) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.SwapMain2Side(&_MainSwapAgent.TransactOpts, mainChainErc20Addr, sideChainToAddr, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MainSwapAgent *MainSwapAgentTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _MainSwapAgent.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MainSwapAgent *MainSwapAgentSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.TransferOwnership(&_MainSwapAgent.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_MainSwapAgent *MainSwapAgentTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _MainSwapAgent.Contract.TransferOwnership(&_MainSwapAgent.TransactOpts, newOwner)
}

// MainSwapAgentOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the MainSwapAgent contract.
type MainSwapAgentOwnershipTransferredIterator struct {
	Event *MainSwapAgentOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *MainSwapAgentOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSwapAgentOwnershipTransferred)
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
		it.Event = new(MainSwapAgentOwnershipTransferred)
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
func (it *MainSwapAgentOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSwapAgentOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSwapAgentOwnershipTransferred represents a OwnershipTransferred event raised by the MainSwapAgent contract.
type MainSwapAgentOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MainSwapAgent *MainSwapAgentFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MainSwapAgentOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MainSwapAgent.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MainSwapAgentOwnershipTransferredIterator{contract: _MainSwapAgent.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_MainSwapAgent *MainSwapAgentFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MainSwapAgentOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _MainSwapAgent.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSwapAgentOwnershipTransferred)
				if err := _MainSwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_MainSwapAgent *MainSwapAgentFilterer) ParseOwnershipTransferred(log types.Log) (*MainSwapAgentOwnershipTransferred, error) {
	event := new(MainSwapAgentOwnershipTransferred)
	if err := _MainSwapAgent.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainSwapAgentRechargeEventIterator is returned from FilterRechargeEvent and is used to iterate over the raw logs and unpacked data for RechargeEvent events raised by the MainSwapAgent contract.
type MainSwapAgentRechargeEventIterator struct {
	Event *MainSwapAgentRechargeEvent // Event containing the contract specifics and raw log

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
func (it *MainSwapAgentRechargeEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSwapAgentRechargeEvent)
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
		it.Event = new(MainSwapAgentRechargeEvent)
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
func (it *MainSwapAgentRechargeEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSwapAgentRechargeEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSwapAgentRechargeEvent represents a RechargeEvent event raised by the MainSwapAgent contract.
type MainSwapAgentRechargeEvent struct {
	MainChainErc20Addr common.Address
	SendAddr           common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRechargeEvent is a free log retrieval operation binding the contract event 0xba435e9191118b57fc3cd6841f3248e60b065acfea46e76f975249ade106b03b.
//
// Solidity: event RechargeEvent(address indexed mainChainErc20Addr, address indexed sendAddr, uint256 amount)
func (_MainSwapAgent *MainSwapAgentFilterer) FilterRechargeEvent(opts *bind.FilterOpts, mainChainErc20Addr []common.Address, sendAddr []common.Address) (*MainSwapAgentRechargeEventIterator, error) {

	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var sendAddrRule []interface{}
	for _, sendAddrItem := range sendAddr {
		sendAddrRule = append(sendAddrRule, sendAddrItem)
	}

	logs, sub, err := _MainSwapAgent.contract.FilterLogs(opts, "RechargeEvent", mainChainErc20AddrRule, sendAddrRule)
	if err != nil {
		return nil, err
	}
	return &MainSwapAgentRechargeEventIterator{contract: _MainSwapAgent.contract, event: "RechargeEvent", logs: logs, sub: sub}, nil
}

// WatchRechargeEvent is a free log subscription operation binding the contract event 0xba435e9191118b57fc3cd6841f3248e60b065acfea46e76f975249ade106b03b.
//
// Solidity: event RechargeEvent(address indexed mainChainErc20Addr, address indexed sendAddr, uint256 amount)
func (_MainSwapAgent *MainSwapAgentFilterer) WatchRechargeEvent(opts *bind.WatchOpts, sink chan<- *MainSwapAgentRechargeEvent, mainChainErc20Addr []common.Address, sendAddr []common.Address) (event.Subscription, error) {

	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var sendAddrRule []interface{}
	for _, sendAddrItem := range sendAddr {
		sendAddrRule = append(sendAddrRule, sendAddrItem)
	}

	logs, sub, err := _MainSwapAgent.contract.WatchLogs(opts, "RechargeEvent", mainChainErc20AddrRule, sendAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSwapAgentRechargeEvent)
				if err := _MainSwapAgent.contract.UnpackLog(event, "RechargeEvent", log); err != nil {
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
// Solidity: event RechargeEvent(address indexed mainChainErc20Addr, address indexed sendAddr, uint256 amount)
func (_MainSwapAgent *MainSwapAgentFilterer) ParseRechargeEvent(log types.Log) (*MainSwapAgentRechargeEvent, error) {
	event := new(MainSwapAgentRechargeEvent)
	if err := _MainSwapAgent.contract.UnpackLog(event, "RechargeEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainSwapAgentSwapMain2SideEventIterator is returned from FilterSwapMain2SideEvent and is used to iterate over the raw logs and unpacked data for SwapMain2SideEvent events raised by the MainSwapAgent contract.
type MainSwapAgentSwapMain2SideEventIterator struct {
	Event *MainSwapAgentSwapMain2SideEvent // Event containing the contract specifics and raw log

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
func (it *MainSwapAgentSwapMain2SideEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSwapAgentSwapMain2SideEvent)
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
		it.Event = new(MainSwapAgentSwapMain2SideEvent)
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
func (it *MainSwapAgentSwapMain2SideEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSwapAgentSwapMain2SideEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSwapAgentSwapMain2SideEvent represents a SwapMain2SideEvent event raised by the MainSwapAgent contract.
type MainSwapAgentSwapMain2SideEvent struct {
	Sponsor            common.Address
	MainChainErc20Addr common.Address
	SideChainToAddr    common.Address
	Amount             *big.Int
	FeeAmount          *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwapMain2SideEvent is a free log retrieval operation binding the contract event 0x42733bad7741b9044a0b058e7d8038dc72cdd877c6eec0d1db3b57f02be7eb81.
//
// Solidity: event SwapMain2SideEvent(address indexed sponsor, address indexed mainChainErc20Addr, address indexed sideChainToAddr, uint256 amount, uint256 feeAmount)
func (_MainSwapAgent *MainSwapAgentFilterer) FilterSwapMain2SideEvent(opts *bind.FilterOpts, sponsor []common.Address, mainChainErc20Addr []common.Address, sideChainToAddr []common.Address) (*MainSwapAgentSwapMain2SideEventIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var sideChainToAddrRule []interface{}
	for _, sideChainToAddrItem := range sideChainToAddr {
		sideChainToAddrRule = append(sideChainToAddrRule, sideChainToAddrItem)
	}

	logs, sub, err := _MainSwapAgent.contract.FilterLogs(opts, "SwapMain2SideEvent", sponsorRule, mainChainErc20AddrRule, sideChainToAddrRule)
	if err != nil {
		return nil, err
	}
	return &MainSwapAgentSwapMain2SideEventIterator{contract: _MainSwapAgent.contract, event: "SwapMain2SideEvent", logs: logs, sub: sub}, nil
}

// WatchSwapMain2SideEvent is a free log subscription operation binding the contract event 0x42733bad7741b9044a0b058e7d8038dc72cdd877c6eec0d1db3b57f02be7eb81.
//
// Solidity: event SwapMain2SideEvent(address indexed sponsor, address indexed mainChainErc20Addr, address indexed sideChainToAddr, uint256 amount, uint256 feeAmount)
func (_MainSwapAgent *MainSwapAgentFilterer) WatchSwapMain2SideEvent(opts *bind.WatchOpts, sink chan<- *MainSwapAgentSwapMain2SideEvent, sponsor []common.Address, mainChainErc20Addr []common.Address, sideChainToAddr []common.Address) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var sideChainToAddrRule []interface{}
	for _, sideChainToAddrItem := range sideChainToAddr {
		sideChainToAddrRule = append(sideChainToAddrRule, sideChainToAddrItem)
	}

	logs, sub, err := _MainSwapAgent.contract.WatchLogs(opts, "SwapMain2SideEvent", sponsorRule, mainChainErc20AddrRule, sideChainToAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSwapAgentSwapMain2SideEvent)
				if err := _MainSwapAgent.contract.UnpackLog(event, "SwapMain2SideEvent", log); err != nil {
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

// ParseSwapMain2SideEvent is a log parse operation binding the contract event 0x42733bad7741b9044a0b058e7d8038dc72cdd877c6eec0d1db3b57f02be7eb81.
//
// Solidity: event SwapMain2SideEvent(address indexed sponsor, address indexed mainChainErc20Addr, address indexed sideChainToAddr, uint256 amount, uint256 feeAmount)
func (_MainSwapAgent *MainSwapAgentFilterer) ParseSwapMain2SideEvent(log types.Log) (*MainSwapAgentSwapMain2SideEvent, error) {
	event := new(MainSwapAgentSwapMain2SideEvent)
	if err := _MainSwapAgent.contract.UnpackLog(event, "SwapMain2SideEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainSwapAgentSwapPairRegisterEventIterator is returned from FilterSwapPairRegisterEvent and is used to iterate over the raw logs and unpacked data for SwapPairRegisterEvent events raised by the MainSwapAgent contract.
type MainSwapAgentSwapPairRegisterEventIterator struct {
	Event *MainSwapAgentSwapPairRegisterEvent // Event containing the contract specifics and raw log

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
func (it *MainSwapAgentSwapPairRegisterEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSwapAgentSwapPairRegisterEvent)
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
		it.Event = new(MainSwapAgentSwapPairRegisterEvent)
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
func (it *MainSwapAgentSwapPairRegisterEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSwapAgentSwapPairRegisterEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSwapAgentSwapPairRegisterEvent represents a SwapPairRegisterEvent event raised by the MainSwapAgent contract.
type MainSwapAgentSwapPairRegisterEvent struct {
	Sponsor            common.Address
	MainChainErc20Addr common.Address
	SideChainErc20Addr common.Address
	Name               string
	Symbol             string
	Decimals           uint8
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwapPairRegisterEvent is a free log retrieval operation binding the contract event 0x84938b4885411ca7a208fde741a0feec069b8173eef6d83fcdecc5daf3a2af3d.
//
// Solidity: event SwapPairRegisterEvent(address indexed sponsor, address indexed mainChainErc20Addr, address indexed sideChainErc20Addr, string name, string symbol, uint8 decimals)
func (_MainSwapAgent *MainSwapAgentFilterer) FilterSwapPairRegisterEvent(opts *bind.FilterOpts, sponsor []common.Address, mainChainErc20Addr []common.Address, sideChainErc20Addr []common.Address) (*MainSwapAgentSwapPairRegisterEventIterator, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var sideChainErc20AddrRule []interface{}
	for _, sideChainErc20AddrItem := range sideChainErc20Addr {
		sideChainErc20AddrRule = append(sideChainErc20AddrRule, sideChainErc20AddrItem)
	}

	logs, sub, err := _MainSwapAgent.contract.FilterLogs(opts, "SwapPairRegisterEvent", sponsorRule, mainChainErc20AddrRule, sideChainErc20AddrRule)
	if err != nil {
		return nil, err
	}
	return &MainSwapAgentSwapPairRegisterEventIterator{contract: _MainSwapAgent.contract, event: "SwapPairRegisterEvent", logs: logs, sub: sub}, nil
}

// WatchSwapPairRegisterEvent is a free log subscription operation binding the contract event 0x84938b4885411ca7a208fde741a0feec069b8173eef6d83fcdecc5daf3a2af3d.
//
// Solidity: event SwapPairRegisterEvent(address indexed sponsor, address indexed mainChainErc20Addr, address indexed sideChainErc20Addr, string name, string symbol, uint8 decimals)
func (_MainSwapAgent *MainSwapAgentFilterer) WatchSwapPairRegisterEvent(opts *bind.WatchOpts, sink chan<- *MainSwapAgentSwapPairRegisterEvent, sponsor []common.Address, mainChainErc20Addr []common.Address, sideChainErc20Addr []common.Address) (event.Subscription, error) {

	var sponsorRule []interface{}
	for _, sponsorItem := range sponsor {
		sponsorRule = append(sponsorRule, sponsorItem)
	}
	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var sideChainErc20AddrRule []interface{}
	for _, sideChainErc20AddrItem := range sideChainErc20Addr {
		sideChainErc20AddrRule = append(sideChainErc20AddrRule, sideChainErc20AddrItem)
	}

	logs, sub, err := _MainSwapAgent.contract.WatchLogs(opts, "SwapPairRegisterEvent", sponsorRule, mainChainErc20AddrRule, sideChainErc20AddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSwapAgentSwapPairRegisterEvent)
				if err := _MainSwapAgent.contract.UnpackLog(event, "SwapPairRegisterEvent", log); err != nil {
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

// ParseSwapPairRegisterEvent is a log parse operation binding the contract event 0x84938b4885411ca7a208fde741a0feec069b8173eef6d83fcdecc5daf3a2af3d.
//
// Solidity: event SwapPairRegisterEvent(address indexed sponsor, address indexed mainChainErc20Addr, address indexed sideChainErc20Addr, string name, string symbol, uint8 decimals)
func (_MainSwapAgent *MainSwapAgentFilterer) ParseSwapPairRegisterEvent(log types.Log) (*MainSwapAgentSwapPairRegisterEvent, error) {
	event := new(MainSwapAgentSwapPairRegisterEvent)
	if err := _MainSwapAgent.contract.UnpackLog(event, "SwapPairRegisterEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MainSwapAgentSwapSide2MainFilledEventIterator is returned from FilterSwapSide2MainFilledEvent and is used to iterate over the raw logs and unpacked data for SwapSide2MainFilledEvent events raised by the MainSwapAgent contract.
type MainSwapAgentSwapSide2MainFilledEventIterator struct {
	Event *MainSwapAgentSwapSide2MainFilledEvent // Event containing the contract specifics and raw log

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
func (it *MainSwapAgentSwapSide2MainFilledEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainSwapAgentSwapSide2MainFilledEvent)
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
		it.Event = new(MainSwapAgentSwapSide2MainFilledEvent)
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
func (it *MainSwapAgentSwapSide2MainFilledEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainSwapAgentSwapSide2MainFilledEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainSwapAgentSwapSide2MainFilledEvent represents a SwapSide2MainFilledEvent event raised by the MainSwapAgent contract.
type MainSwapAgentSwapSide2MainFilledEvent struct {
	MainChainErc20Addr common.Address
	SideChainTxHash    [32]byte
	MainChainToAddr    common.Address
	Amount             *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterSwapSide2MainFilledEvent is a free log retrieval operation binding the contract event 0x8a6ec552fee56a817c37275213226a383c891ccdc5b251d8dae0da50543e623d.
//
// Solidity: event SwapSide2MainFilledEvent(address indexed mainChainErc20Addr, bytes32 indexed sideChainTxHash, address indexed mainChainToAddr, uint256 amount)
func (_MainSwapAgent *MainSwapAgentFilterer) FilterSwapSide2MainFilledEvent(opts *bind.FilterOpts, mainChainErc20Addr []common.Address, sideChainTxHash [][32]byte, mainChainToAddr []common.Address) (*MainSwapAgentSwapSide2MainFilledEventIterator, error) {

	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var sideChainTxHashRule []interface{}
	for _, sideChainTxHashItem := range sideChainTxHash {
		sideChainTxHashRule = append(sideChainTxHashRule, sideChainTxHashItem)
	}
	var mainChainToAddrRule []interface{}
	for _, mainChainToAddrItem := range mainChainToAddr {
		mainChainToAddrRule = append(mainChainToAddrRule, mainChainToAddrItem)
	}

	logs, sub, err := _MainSwapAgent.contract.FilterLogs(opts, "SwapSide2MainFilledEvent", mainChainErc20AddrRule, sideChainTxHashRule, mainChainToAddrRule)
	if err != nil {
		return nil, err
	}
	return &MainSwapAgentSwapSide2MainFilledEventIterator{contract: _MainSwapAgent.contract, event: "SwapSide2MainFilledEvent", logs: logs, sub: sub}, nil
}

// WatchSwapSide2MainFilledEvent is a free log subscription operation binding the contract event 0x8a6ec552fee56a817c37275213226a383c891ccdc5b251d8dae0da50543e623d.
//
// Solidity: event SwapSide2MainFilledEvent(address indexed mainChainErc20Addr, bytes32 indexed sideChainTxHash, address indexed mainChainToAddr, uint256 amount)
func (_MainSwapAgent *MainSwapAgentFilterer) WatchSwapSide2MainFilledEvent(opts *bind.WatchOpts, sink chan<- *MainSwapAgentSwapSide2MainFilledEvent, mainChainErc20Addr []common.Address, sideChainTxHash [][32]byte, mainChainToAddr []common.Address) (event.Subscription, error) {

	var mainChainErc20AddrRule []interface{}
	for _, mainChainErc20AddrItem := range mainChainErc20Addr {
		mainChainErc20AddrRule = append(mainChainErc20AddrRule, mainChainErc20AddrItem)
	}
	var sideChainTxHashRule []interface{}
	for _, sideChainTxHashItem := range sideChainTxHash {
		sideChainTxHashRule = append(sideChainTxHashRule, sideChainTxHashItem)
	}
	var mainChainToAddrRule []interface{}
	for _, mainChainToAddrItem := range mainChainToAddr {
		mainChainToAddrRule = append(mainChainToAddrRule, mainChainToAddrItem)
	}

	logs, sub, err := _MainSwapAgent.contract.WatchLogs(opts, "SwapSide2MainFilledEvent", mainChainErc20AddrRule, sideChainTxHashRule, mainChainToAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainSwapAgentSwapSide2MainFilledEvent)
				if err := _MainSwapAgent.contract.UnpackLog(event, "SwapSide2MainFilledEvent", log); err != nil {
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

// ParseSwapSide2MainFilledEvent is a log parse operation binding the contract event 0x8a6ec552fee56a817c37275213226a383c891ccdc5b251d8dae0da50543e623d.
//
// Solidity: event SwapSide2MainFilledEvent(address indexed mainChainErc20Addr, bytes32 indexed sideChainTxHash, address indexed mainChainToAddr, uint256 amount)
func (_MainSwapAgent *MainSwapAgentFilterer) ParseSwapSide2MainFilledEvent(log types.Log) (*MainSwapAgentSwapSide2MainFilledEvent, error) {
	event := new(MainSwapAgentSwapSide2MainFilledEvent)
	if err := _MainSwapAgent.contract.UnpackLog(event, "SwapSide2MainFilledEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
