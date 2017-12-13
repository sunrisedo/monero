package monero

type WalletClient struct {
	*CallClient
}

// start wallet rpc server:
//win .\monero-wallet-rpc.exe --rpc-login user:pass --wallet-file D:\work\text\admin\admin --rpc-bind-ip 127.0.0.1 --rpc-bind-port 18082
func NewWalletClient(endpoint, username, password string) *WalletClient {
	return &WalletClient{NewCallClient(endpoint, username, password)}
}

// curl -u user:pass --digest http://127.0.0.1:18082/json_rp   -d '{"jsonrpc":"2.0","id":"0","method":"'getbalance'"}'
func (c *WalletClient) GetBalance() (Balance, error) {
	var rep Balance
	if err := c.Wallet("getbalance", nil, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

func (c *WalletClient) GetAddress() (Address, error) {
	var rep Address
	if err := c.Wallet("getaddress", nil, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}
func (c *WalletClient) GetHeight() (Height, error) {
	var rep Height
	if err := c.Wallet("getheight", nil, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

func (c *WalletClient) Transfer(req TransferInput) (Transfer, error) {
	var rep Transfer
	if err := c.Wallet("transfer", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}
func (c *WalletClient) TransferSplit(req TransferInput) (Transfer, error) {
	var rep Transfer
	if err := c.Wallet("transfer_split", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

func (c *WalletClient) GetTransfers(req GetTransferInput) (Transfer, error) {
	var rep Transfer
	if err := c.Wallet("get_transfers", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}
func (c *WalletClient) GetTransferByTxId(txid string) (Transfer, error) {
	req := struct {
		Txid string `json:"txid"`
	}{
		txid,
	}
	rep := struct {
		Trade Transfer `json:"transfer"`
	}{}
	if err := c.Wallet("get_transfer_by_txid", req, &rep); err != nil {
		return rep.Trade, err
	}
	return rep.Trade, nil
}
func (c *WalletClient) IncomingTransfers(transferType string) (Transfer, error) {
	req := struct {
		TransferType string `json:"transfer_type"`
	}{
		transferType,
	}
	var rep Transfer
	if err := c.Wallet("incoming_transfers", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}

func (c *WalletClient) MakeIntegratedAddress(paymentId string) (IsAddress, error) {
	req := struct {
		PaymentId string `json:"payment_id,omitempty"`
	}{
		paymentId,
	}
	var rep IsAddress
	if err := c.Wallet("make_integrated_address", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}
func (c *WalletClient) SplitIntegratedAddress(integratedAddress string) (IsAddress, error) {
	req := struct {
		IntegratedAddress string `json:"integrated_address,omitempty"`
	}{
		integratedAddress,
	}
	var rep IsAddress
	if err := c.Wallet("split_integrated_address", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}
func (c *WalletClient) GetAddressBook(entries []uint) ([]Entries, error) {
	req := struct {
		Entries []uint `json:"entries,omitempty"`
	}{
		entries,
	}
	rep := struct {
		Entries []Entries `json:"entries,omitempty"`
	}{}
	if err := c.Wallet("get_address_book", req, &rep); err != nil {
		return rep.Entries, err
	}
	return rep.Entries, nil
}
func (c *WalletClient) AddAddressBook(address, paymentId, description string) (uint, error) {
	req := struct {
		Address     string `json:"address,omitempty"`
		PaymentId   string `json:"payment_id,omitempty"`
		Description string `json:"description,omitempty"`
	}{
		address,
		paymentId,
		description,
	}
	var rep uint
	if err := c.Wallet("add_address_book", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}
func (c *WalletClient) DeleteAddressBook(index uint) error {
	req := struct {
		Index uint `json:"index"`
	}{
		index,
	}
	if err := c.Wallet("delete_address_book", req, nil); err != nil {
		return err
	}
	return nil
}

func (c *WalletClient) GetPayments(paymentId string) ([]Payment, error) {
	req := struct {
		PaymentId string `json:"payment_id,omitempty"`
	}{
		paymentId,
	}
	rep := struct {
		Payments []Payment `json:"payments"`
	}{}
	if err := c.Wallet("get_payments", req, &rep); err != nil {
		return rep.Payments, err
	}
	return rep.Payments, nil
}
func (c *WalletClient) GetBulkPayments(paymentIds []string, minBlockHeight uint64) ([]Payment, error) {
	req := struct {
		PaymentIds     []string `json:"payment_ids,omitempty"`
		MinBlockHeight uint64   `json:"min_block_height,omitempty"`
	}{
		paymentIds,
		minBlockHeight,
	}
	rep := struct {
		Payments []Payment `json:"payments"`
	}{}
	if err := c.Wallet("get_bulk_payments", req, &rep); err != nil {
		return rep.Payments, err
	}
	return rep.Payments, nil
}

func (c *WalletClient) Store() (bool, error) {
	if err := c.Wallet("store", nil, nil); err != nil {
		return false, err
	}
	return true, nil
}
func (c *WalletClient) SweepDust() (bool, error) {
	if err := c.Wallet("sweep_dust", nil, nil); err != nil {
		return false, err
	}
	return true, nil
}

func (c *WalletClient) StopWallet() (bool, error) {
	if err := c.Wallet("stop_wallet", nil, nil); err != nil {
		return false, err
	}
	return true, nil
}

// func (c *WalletClient) StopWallet() (bool, error) {
// 	if err := c.Wallet("stop_wallet", nil, nil); err != nil {
// 		return false, err
// 	}
// 	return true, nil
// }
