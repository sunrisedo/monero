package monero

/***************************************/
type Balance struct {
	Balance   uint64 `json:"balance"`
	UnBalance uint64 `json:"unlocked_balance"`
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
	Mixin        uint          `json:"mixin"`
	GetTxKey     bool          `json:"get_tx_key"`
	Fee          uint          `json:"fee"`
	UnlockTime   uint64        `json:"unlock_time"`
	PaymentId    string        `json:"payment_id"`
	Priority     uint          `json:"priority"`
	DoNotRelay   bool          `json:"do_not_relay"`
	GetTxHex     bool          `json:"get_tx_hex"`
}
type Destination struct {
	Amount  int64  `json:"amount"`
	Address string `json:"address"`
}

type Transfer struct {
	Fee    uint   `json:"fee,omitempty"`
	TxHash string `json:"tx_hash,omitempty"`
	TxKey  string `json:"tx_key,omitempty"`

	Amount    uint64 `json:"amount,omitempty"`
	Height    uint64 `json:"height,omitempty"`
	Note      string `json:"note,omitempty"`
	PaymentId string `json:"payment_id,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	Txid      string `json:"txid,omitempty"`
	Type      string `json:"type,omitempty"`
}

type TransferSplit struct {
	FeeList    []uint64 `json:"fee_list"`
	TxHashList []string `json:"tx_hash_list"`
	TxBlobList []string `json:"tx_blob_list"`
	AmountList []uint64 `json:"amount_list"`
	TxKeyList  []string `json:"tx_key_list"`
}

// type Transfer struct {
// 	Amount    uint   `json:"amount"`
// 	Fee       uint   `json:"fee"`
// 	Height    uint   `json:"height"`
// 	Note      string `json:"note"`
// 	PaymentId string `json:"payment_id"`
// 	Timestamp uint   `json:"timestamp"`
// 	Txid      string `json:"txid"`
// 	Type      string `json:"type"`
// }

type IncomingTransfers struct {
}

/***************************************/
type GetTransferInput struct {
	In             bool   `json:"in,omitempty"`
	Out            bool   `json:"out,omitempty"`
	Pending        bool   `json:"pending,omitempty"`
	Failed         bool   `json:"failed,omitempty"`
	Pool           bool   `json:"pool,omitempty"`
	FilterByHeight bool   `json:"filter_by_height,omitempty"`
	MinHeight      uint64 `json:"min_height,omitempty"`
	MaxHeight      uint64 `json:"max_height,omitempty"`
}

/***************************************/
type IsAddress struct {
	PaymentId         string `json:"payment_id"`
	IntegratedAddress string `json:"integrated_address,omitempty"`
	StandardAddress   string `json:"standard_address,omitempty"`
}

type Entries struct {
	Address     string `json:"address"`
	Description string `json:"description"`
	Index       string `json:"index"`
	PaymentId   string `json:"payment_id"`
}

/***************************************/
type Payment struct {
	Amount      uint64 `json:"amount"`
	BlockHeight uint64 `json:"block_height"`
	PaymentId   string `json:"payment_id"`
	TxHash      string `json:"tx_hash"`
	UnlockTime  uint64 `json:"unlock_time"`
}
