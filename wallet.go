package monero

// con "github.com/Soontao/goHttpDigestClient"

type WalletClient struct {
	*CallClient
}

// start wallet rpc server:
// .\monero-wallet-rpc.exe --rpc-login user:pass --wallet-file D:\work\text\admin\admin --rpc-bind-ip 127.0.0.1 --rpc-bind-port 18082
func NewWalletClient(endpoint, username, password string) *WalletClient {
	return &WalletClient{NewCallClient(endpoint, username, password)}
}

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
func (c *WalletClient) GetTransfers(req GetTransferInput) (Transfer, error) {
	var rep Transfer
	if err := c.Wallet("get_transfers", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}
func (c *WalletClient) MakeIntegratedAddress(req IntegratedAddressInput) (IntegratedAddress, error) {
	var rep IntegratedAddress
	if err := c.Wallet("make_integrated_address", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}
func (c *WalletClient) SplitIntegratedAddress(req IntegratedAddressInput) (IntegratedAddress, error) {
	var rep IntegratedAddress
	if err := c.Wallet("split_integrated_address", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
}
func (c *WalletClient) GetPayments(req GetPaymentsInput) (Payments, error) {
	var rep Payments
	if err := c.Wallet("get_payments", req, &rep); err != nil {
		return rep, err
	}
	return rep, nil
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
