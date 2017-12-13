package monero

type DaemonClient struct {
	*CallClient
}

// Creates new daemon client
func NewDaemonClient(endpoint string) *DaemonClient {
	return &DaemonClient{NewCallClient(endpoint, "", "")}
}

// Look up how many blocks are in the longest chain known to the node.
func (c *DaemonClient) GetBlockCount() (BlockCount, error) {
	var bc BlockCount
	if err := c.Daemon("getblockcount", nil, &bc); err != nil {
		return bc, err
	}
	return bc, nil
}

// Look up a block's hash by its height.
func (c *DaemonClient) OnGetBlockHash(blockHeight int) (string, error) {
	var blockHash string
	if err := c.Daemon("on_getblockhash", []int{blockHeight}, &blockHash); err != nil {
		return blockHash, err
	}

	return blockHash, nil
}

// Get BlockTemplate
func (c *DaemonClient) GetBlockTemplate(walletAddress string, reserveSize uint) (BlockTemplate, error) {
	var bt BlockTemplate
	req := struct {
		walletAddress string `json:"wallet_address"`
		reserveSize   uint   `json:"reserve_size"`
	}{
		walletAddress,
		reserveSize,
	}
	if err := c.Daemon("getblocktemplate", req, &bt); err != nil {
		return bt, err
	}
	return bt, nil
}

// Submit a mined block to the network.
func (c *DaemonClient) SubmitBlock(blockBlobData string) (string, error) {
	var status string
	if err := c.Daemon("submitblock", blockBlobData, &status); err != nil {
		return status, err
	}
	return status, nil
}

// Block header information for the most recent block is easily retrieved with this method. No inputs are needed.
func (c *DaemonClient) GetLastBlockHeader() (BlockHeaderResponse, error) {
	var bhr BlockHeaderResponse
	if err := c.Daemon("getlastblockheader", nil, &bhr); err != nil {
		return bhr, err
	}
	return bhr, nil
}

// Block header information can be retrieved using either a block's hash or height.
// This method includes a block's hash as an input parameter to retrieve basic information about the block.
func (c *DaemonClient) GetBlockHeaderByHash(hash string) (BlockHeaderResponse, error) {
	var bhr BlockHeaderResponse
	req := struct {
		Hash string `json:"hash"`
	}{
		hash,
	}
	if err := c.Daemon("getblockheaderbyhash", req, &bhr); err != nil {
		return bhr, err
	}
	return bhr, nil
}

// Similar to GetBlockHeaderByHash above, this method includes a block's height as an input parameter to retrieve basic information about the block.
func (c *DaemonClient) GetBlockHeaderByHeight(height uint64) (BlockHeaderResponse, error) {
	req := struct {
		Height uint64 `json:"height"`
	}{
		height,
	}
	var bhr BlockHeaderResponse
	if err := c.Daemon("getblockheaderbyheight", req, &bhr); err != nil {
		return bhr, err
	}
	return bhr, nil
}

// Full block information can be retrieved by either block height or hash, like with the above block header calls.
// For full block information, both lookups use the same method, but with different input parameters.
func (c *DaemonClient) GetBlock(height uint, hash string) (Block, error) {
	var b Block
	req := struct {
		height uint   `json:"height, omitempty"`
		hash   string `json:"hash, omitempty"`
	}{
		height,
		hash,
	}
	if err := c.Daemon("getblock", req, &b); err != nil {
		return b, err
	}
	return b, nil
}

// Retrieve information about incoming and outgoing connections to your node.
func (c *DaemonClient) GetConnections() (ConnectionResponse, error) {
	var cr ConnectionResponse
	if err := c.Daemon("get_connections", nil, &cr); err != nil {
		return cr, err
	}
	return cr, nil
}

// Retrieve general information about the state of your node and the network.
func (c *DaemonClient) GetInfo() (Info, error) {
	var inf Info
	if err := c.Daemon("get_info", nil, &inf); err != nil {
		return inf, err
	}
	return inf, nil
}

// Look up information regarding hard fork voting and readiness.
func (c *DaemonClient) GetHardForkInfo() (HardForkInfo, error) {
	var hi HardForkInfo
	if err := c.Daemon("hard_fork_info", nil, &hi); err != nil {
		return hi, err
	}
	return hi, nil
}

// Ban another node by IP.
func (c *DaemonClient) SetBans(bans []Ban) (string, error) {
	var status string
	if err := c.Daemon("setbans", nil, &status); err != nil {
		return status, err
	}
	return status, nil
}

// Get bans
func (c *DaemonClient) GetBans() (BanResponse, error) {
	var br BanResponse
	if err := c.Daemon("getbans", nil, &br); err != nil {
		return br, err
	}
	return br, nil
}
