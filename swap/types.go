package swap

import (
	"crypto/ecdsa"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum/accounts/abi"

	ethcom "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/jinzhu/gorm"

	"github.com/bcskill/eth-main-side-swap/common"
	"github.com/bcskill/eth-main-side-swap/util"
)

const (
	SwapTokenReceived common.SwapStatus = "received"
	SwapQuoteRejected common.SwapStatus = "rejected"
	SwapConfirmed     common.SwapStatus = "confirmed"
	SwapSending       common.SwapStatus = "sending"
	SwapSent          common.SwapStatus = "sent"
	SwapSendFailed    common.SwapStatus = "sent_fail"
	SwapSuccess       common.SwapStatus = "sent_success"

	SwapPairReceived   common.SwapPairStatus = "received"
	SwapPairConfirmed  common.SwapPairStatus = "confirmed"
	SwapPairSending    common.SwapPairStatus = "sending"
	SwapPairSent       common.SwapPairStatus = "sent"
	SwapPairSendFailed common.SwapPairStatus = "sent_fail"
	SwapPairSuccess    common.SwapPairStatus = "sent_success"
	SwapPairFinalized  common.SwapPairStatus = "finalized"

	RetrySwapConfirmed  common.RetrySwapStatus = "confirmed"
	RetrySwapSending    common.RetrySwapStatus = "sending"
	RetrySwapSent       common.RetrySwapStatus = "sent"
	RetrySwapSendFailed common.RetrySwapStatus = "sent_fail"
	RetrySwapSuccess    common.RetrySwapStatus = "sent_success"

	SwapMain2Side common.SwapDirection = "main_side"
	SwapSide2Main common.SwapDirection = "side_main"

	BatchSize                = 50
	TrackSentTxBatchSize     = 100
	SleepTime                = 5
	SwapSleepSecond          = 2
	TrackSwapPairSMBatchSize = 5

	TxFailedStatus = 0x00

	MaxUpperBound = "999999999999999999999999999999999999"
)

var mainClientMutex sync.RWMutex
var sideClientMutex sync.RWMutex

type SwapEngine struct {
	mutex    sync.RWMutex
	db       *gorm.DB
	hmacCKey string
	config   *util.Config
	// key is the side contract addr
	swapPairsFromERC20Addr map[ethcom.Address]*SwapPairIns
	mainClient              *ethclient.Client
	sideClient              *ethclient.Client
	mainPrivateKey          *ecdsa.PrivateKey
	sidePrivateKey          *ecdsa.PrivateKey
	mainChainID             int64
	sideChainID             int64
	side20ToMain20           map[ethcom.Address]ethcom.Address
	main20ToSide20           map[ethcom.Address]ethcom.Address

	mainSwapAgentABI *abi.ABI
	sideSwapAgentABI *abi.ABI

	mainSwapAgent ethcom.Address
	sideSwapAgent ethcom.Address
}

type SwapPairEngine struct {
	mutex   sync.RWMutex
	db      *gorm.DB
	hmacKey string
	config  *util.Config

	swapEngine *SwapEngine

	sideClient       *ethclient.Client
	sidePrivateKey   *ecdsa.PrivateKey
	sideChainID      int64
	sideTxSender     ethcom.Address
	sideSwapAgent    ethcom.Address
	sideSwapAgentABI *abi.ABI
}

type SwapPairIns struct {
	Symbol     string
	Name       string
	Decimals   int
	LowBound   *big.Int
	UpperBound *big.Int

	MainChainErc20Addr ethcom.Address
	SideChainErc20Addr ethcom.Address
}
