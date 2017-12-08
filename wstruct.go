package monero

/***************************************/
type Balance struct {
	Balance   int64 `json:"balance"`
	UnBalance int64 `json:"unlocked_balance"`
}

/***************************************/
type Address struct {
	Address string `json:"address"`
}

/***************************************/
type Height struct {
	Height int64 `json:"height"`
}

/***************************************/
type TransferInput struct {
	Destinations []Destination `json:"destinations"`
	Mixin        uint64        `json:"mixin"`
	GetTxKey     bool          `json:"get_tx_key"`
}
type Destination struct {
	Amount  int64  `json:"amount"`
	Address string `json:"address"`
}
type Transfer struct {
	Fee    uint64 `json:"fee"`
	TxHash string `json:"tx_hash"`
	TxKey  string `json:"tx_key"`
}

/***************************************/
type GetTransferInput struct {
	Pool bool `json:"pool"`
}

/***************************************/
type IntegratedAddressInput struct {
	PaymentId         string `json:"payment_id,omitempty"`
	IntegratedAddress string `json:"integrated_address,omitempty"`
}
type IntegratedAddress struct {
	PaymentId         string `json:"payment_id"`
	IntegratedAddress string `json:"integrated_address"`
	StandardAddress   string `json:"standard_address"`
}

/***************************************/
type GetPaymentsInput struct {
	PaymentId string `json:"payment_id,omitempty"`
}
type Payments struct {
	Payments []Payment `json:"payments"`
}
type Payment struct {
	Amount      uint64 `json:"amount"`
	BlockHeight uint64 `json:"block_height"`
	PaymentId   string `json:"payment_id"`
	TxHash      string `json:"tx_hash"`
	UnlockTime  uint64 `json:"unlock_time"`
}
