package model

import (
	"time"

	"github.com/jinzhu/gorm"

	"github.com/bcskill/eth-main-side-swap/common"
)

type SwapStartTxLog struct {
	Id    int64
	Chain string `gorm:"not null;index:swap_main_side_start_txs_chain"`
	Sponsor string `gorm:"not null;index:swap_main_side_start_txs_sponsor"`
	SourceChainErc20Addr string `gorm:"not null"`
	TargetChainToAddr    string `gorm:"not null"`
	Amount      string `gorm:"not null"`
	FeeAmount   string `gorm:"not null"`

	Status       TxStatus `gorm:"not null;index:swap_main_side_start_txs_status"`
	TxHash       string   `gorm:"not null;index:swap_main_side_start_txs_tx_hash"`
	BlockHash    string   `gorm:"not null"`
	Height       int64    `gorm:"not null"`
	ConfirmedNum int64    `gorm:"not null"`

	Phase TxPhase `gorm:"not null;index:swap_main_side_start_txs_phase"`

	UpdateTime int64
	CreateTime int64
}

func (SwapStartTxLog) TableName() string {
	return "swap_main_side_start_txs"
}

func (l *SwapStartTxLog) BeforeCreate() (err error) {
	l.CreateTime = time.Now().Unix()
	l.UpdateTime = time.Now().Unix()
	return nil
}

type SwapFillTx struct {
	gorm.Model

	Direction         common.SwapDirection `gorm:"not null"`
	StartSwapTxHash   string               `gorm:"not null;index:swap_fill_tx_start_swap_tx_hash"`
	FillSwapTxHash    string               `gorm:"not null;index:swap_fill_tx_fill_swap_tx_hash"`
	GasPrice          string               `gorm:"not null"`
	ConsumedFeeAmount string
	Height            int64
	Status            FillTxStatus `gorm:"not null"`
	TrackRetryHeight  int64
	TrackRetryCounter int64
}

func (SwapFillTx) TableName() string {
	return "swap_fill_txs"
}

type RetrySwap struct {
	gorm.Model

	Status      common.RetrySwapStatus `gorm:"not null"`
	SwapID      uint                   `gorm:"not null"`
	Direction   common.SwapDirection   `gorm:"not null"`
	StartTxHash string                 `gorm:"not null;index:retry_swap_start_tx_hash"`
	FillTxHash  string                 `gorm:"not null"`
	Sponsor     string                 `gorm:"not null;index:retry_swap_sponsor"`
	SourceChainErc20Addr string        `gorm:"not null;index:retry_swap_source_addr"`
	TargetChainErc20Addr string        `gorm:"not null;index:retry_swap_target_addr"`
	TargetChainToAddr    string        `gorm:"not null;index:retry_swap_target_to_addr"`
	Symbol      string                 `gorm:"not null"`
	Amount      string                 `gorm:"not null"`
	Decimals    int                    `gorm:"not null"`

	RecordHash string `gorm:"not null"`
	ErrorMsg   string
}

func (RetrySwap) TableName() string {
	return "retry_swaps"
}

type RetrySwapTx struct {
	gorm.Model

	RetrySwapID         uint                 `gorm:"not null;index:retry_swap_tx_retry_swap_id"`
	StartTxHash         string               `gorm:"not null;index:retry_swap_tx_start_tx_hash"`
	Direction           common.SwapDirection `gorm:"not null"`
	TrackRetryCounter   int64
	RetryFillSwapTxHash string            `gorm:"not null"`
	Status              FillRetryTxStatus `gorm:"not null"`
	ErrorMsg            string            `gorm:"not null"`
	GasPrice            string
	ConsumedFeeAmount   string
	Height              int64
}

func (RetrySwapTx) TableName() string {
	return "retry_swap_txs"
}

type Swap struct {
	gorm.Model

	Status common.SwapStatus `gorm:"not null;index:swap_status"`
	// the user addreess who start this swap
	Sponsor string `gorm:"not null"`
	SourceChainErc20Addr string `gorm:"not null;index:swap_source_addr"`
	TargetChainErc20Addr string `gorm:"not null;index:swap_target_addr"`
	TargetChainToAddr string `gorm:"not null;index:swap_target_to_addr"`
	Symbol    string
	Amount    string               `gorm:"not null;index:swap_amount"`
	Decimals  int                  `gorm:"not null"`
	Direction common.SwapDirection `gorm:"not null;index:swap_direction"`

	// The tx hash confirmed deposit
	StartTxHash string `gorm:"not null;index:swap_start_tx_hash"`
	// The tx hash confirmed withdraw
	FillTxHash string `gorm:"not null;index:swap_fill_tx_hash"`

	// used to log more message about how this swap failed or invalid
	Log string

	RecordHash string `gorm:"not null"`
}

func (Swap) TableName() string {
	return "swaps"
}
