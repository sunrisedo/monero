package monero

// Copyright 2017 Marin Basic <marin@marin-basic.com>. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

import (
	"encoding/json"
)

// BlockCount
// count - unsigned int; Number of blocks in longest chain seen by the node.
// status - string; General RPC error code. "OK" means everything looks good.
type BlockCount struct {
	Count  int    `json:"count"`
	Status string `json:"status"`
}

// BlockTemplate
// blocktemplate_blob - string; Blob on which to try to mine a new block.
// difficulty - unsigned int; Difficulty of next block.
// height - unsigned int; Height on which to mine.
// prev_hash - string; Hash of the most recent block on which to mine the next block.
// reserved_offset - unsigned int; Reserved offset.
// status - string; General RPC error code. "OK" means everything looks good.
type BlockTemplate struct {
	BlockTemplateBlob string `json:"blocktemplate_blob "`
	Difficulty        uint   `json:"difficulty"`
	Height            uint   `json:"height"`
	PrevHash          string `json:"prev_hash"`
	ReservedOffset    uint   `json:"reserved_offset"`
	Status            string `json:"status"`
}

// BlockHeader
// depth - unsigned int; The number of blocks succeeding this block on the blockchain. A larger number means an older block.
// difficulty - unsigned int; The strength of the Monero network based on mining power.
// hash - string; The hash of this block.
// height - unsigned int; The number of blocks preceding this block on the blockchain.
// major_version - unsigned int; The major version of the monero protocol at this block height.
// minor_version - unsigned int; The minor version of the monero protocol at this block height.
// nonce - unsigned int; a cryptographic random one-time number used in mining a Monero block.
// orphan_status - boolean; Usually false. If true, this block is not part of the longest chain.
// prev_hash - string; The hash of the block immediately preceding this block in the chain.
// reward - unsigned int; The amount of new atomic units generated in this block and rewarded to the miner. Note: 1 XMR = 1e12 atomic units.
// timestamp - unsigned int; The time the block was recorded into the blockchain.
type BlockHeader struct {
	Depth        uint64 `json:"depth"`
	Difficulty   uint   `json:"difficulty"`
	Hash         string `json:"hash"`
	Height       uint   `json:"height"`
	MajorVersion uint   `json:"major_version"`
	MinorVersion uint   `json:"minor_version"`
	Nonce        uint   `json:"nonce"`
	OrphanStatus bool   `json:"orphan_status"`
	PrevHash     string `json:"prev_hash"`
	Reward       uint   `json:"reward "`
	Timestamp    uint   `json:"timestamp"`
}

// BlockHeaderResponse
// block_header - A structure containing block header information.
// status - string; General RPC error code. "OK" means everything looks good.
type BlockHeaderResponse struct {
	BlockHeader BlockHeader `json:"block_header"`
	Status      string      `json:"status"`
}

// Block
// blob - string; Hexadecimal blob of block information.
// block_header - A structure containing block header information. See getlastblockheader.
// json - json string; JSON formatted block details:
// status - string; General RPC error code. "OK" means everything looks good.
type Block struct {
	Blob        string      `json:"blob"`
	BlockHeader BlockHeader `json:"block_header"`
	Json        string      `json:"json"`
	Status      string      `status`
}

// BlockDetails
// major_version - Same as in block header.
// minor_version - Same as in block header.
// timestamp - Same as in block header.
// prev_id - Same as prev_hash in block header.
// nonce - Same as in block header.
// miner_tx - Miner transaction information
// version - Transaction version number.
// tx_hashes - List of hashes of non-coinbase transactions in the block. If there are no other transactions, this will be an empty list.
type BlockDetails struct {
	MajorVersion         uint                 `json:"major_version"`
	MinorVersion         uint                 `json:"minor_version"`
	Timestamp            uint                 `json:"timestamp"`
	Nonce                uint                 `json:"nonce"`
	PrevId               string               `json:"prev_id"`
	TxHashes             []string             `json:"tx_hashes"`
	MinerTransactionInfo MinerTransactionInfo `json:"miner_tx"`
}

// Parse json string to BlockDetails struct
func (b *Block) ParseJSON() (BlockDetails, error) {
	var bd BlockDetails
	if err := json.Unmarshal([]byte(b.Json), &bd); err != nil {
		return bd, err
	}
	return bd, nil
}

// MinerTransactionInfo
// version - Transaction version number.
// unlock_time - The block height when the coinbase transaction becomes spendable.
// vin - List of transaction inputs:
// gen - Miner txs are coinbase txs, or "gen".
// height - This block height, a.k.a. when the coinbase is generated.
// vout - List of transaction outputs. Each output contains:
// extra - Usually called the "transaction ID" but can be used to include any random 32 byte/64 character hex string.
// signatures - Contain signatures of tx signers. Coinbased txs do not have signatures.
type MinerTransactionInfo struct {
	Version            uint                 `json:"version"`
	UnlockTime         int                  `json:"unlock_time"`
	TransactionInputs  []TransactionInputs  `json:"vin"`
	TransactionOutputs []TransactionOutputs `json:"vout"`
	Extra              []string             `json:"extra"`
	Signatures         []string             `json:"signatures"`
}

// TransactionInputs
// gen - Miner txs are coinbase txs, or "gen".
type TransactionInputs struct {
	Gen struct {
		Height uint `json:"height"`
	} `json:"gen"`
}

// TransactionOutputs
// amount - The amount of the output, in atomic units.
type TransactionOutputs struct {
	Amount uint `json:"amount"`
	Target struct {
		Key string `json:"key"`
	} `json:"target"`
}

// Connection
// avg_download - unsigned int; Average bytes of data downloaded by node.
// avg_upload - unsigned int; Average bytes of data uploaded by node.
// current_download - unsigned int; Current bytes downloaded by node.
// current_upload - unsigned int; Current bytes uploaded by node.
// incoming - boolean; Is the node getting information from your node?
// ip - string; The node's IP address.
// live_time - unsigned int
// local_ip - boolean
// localhost - boolean
// peer_id - string; The node's ID on the network.
// port - stringl The port that the node is using to connect to the network.
// recv_count - unsigned int
// recv_idle_time - unsigned int
// send_count - unsigned int
// send_idle_time - unsigned int
//state - string
type Connection struct {
	AvgDownload     uint   `json:"avg_download"`
	AvgUpload       uint   `json:"avg_upload"`
	CurrentDownload uint   `json:"current_download"`
	CurrentUpload   uint   `json:"current_upload"`
	Incoming        bool   `json:"incoming"`
	Ip              string `json:"ip"`
	LiveTime        uint   `json:"live_time"`
	LocalIp         bool   `json:"local_ip"`
	Localhost       bool   `json:"localhost"`
	PeerId          string `json:"peer_id"`
	Port            string `json:"port"`
	RecvCount       uint   `json:"recv_count"`
	RecvIdleTime    uint   `json:"recv_idle_time"`
	SendCount       uint   `json:"send_count"`
	SendIdleTime    uint   `json:"send_idle_time"`
	State           string `json:"state"`
}

// ConnectionRespose
// connections - List of all connections and their info:
type ConnectionResponse struct {
	Connections []Connection `json:"connections"`
	Status      string       `json:"status"`
}

// Info
// alt_blocks_count - unsigned int; Number of alternative blocks to main chain.
// difficulty - unsigned int; Network difficulty (analogous to the strength of the network)
// grey_peerlist_size - unsigned int; Grey Peerlist Size
// height - unsigned int; Current length of longest chain known to daemon.
// incoming_connections_count - unsigned int; Number of peers connected to and pulling from your node.
// outgoing_connections_count - unsigned int; Number of peers that you are connected to and getting information from.
// status - string; General RPC error code. "OK" means everything looks good.
// target - unsigned int; Current target for next proof of work.
// target_height - unsigned int; The height of the next block in the chain.
// testnet - boolean; States if the node is on the testnet (true) or mainnet (false).
// top_block_hash - string; Hash of the highest block in the chain.
// tx_count - unsigned int; Total number of non-coinbase transaction in the chain.
// tx_pool_siz - unsigned int; Number of transactions that have been broadcast but not included in a block.
// white_peerlist_size - unsigned int; White Peerlist Size
type Info struct {
	AltBlocksCount           uint   `json:"alt_blocks_count"`
	Difficulty               uint   `json:"difficulty"`
	GreyPeerlistSize         uint   `json:"grey_peerlist_size"`
	Height                   uint   `json:"height"`
	IncomingConnectionsCount uint   `json:"incoming_connections_count"`
	OutgoingConnectionsCount uint   `json:"outgoing_connections_count"`
	Status                   string `json:"status"`
	Target                   uint   `json:"target"`
	TargetHeight             uint   `json:"target_height"`
	Testnet                  bool   `json:"testnet"`
	TopBlockHash             string `json:"top_block_hash"`
	TxCount                  uint   `json:"tx_count"`
	TxPoolSiz                uint   `json:"tx_pool_siz"`
	WhitePeerlistSize        uint   `json:"white_peerlist_size"`
}

// HardForkInfo
// earliest_height - unsigned int; Block height at which hard fork would be enabled if voted in.
// enabled - boolean; Tells if hard fork is enforced.
// state - unsigned int; Current hard fork state: 0 (There is likely a hard fork), 1 (An update is needed to fork properly), or 2 (Everything looks good).
// status - string; General RPC error code. "OK" means everything looks good.
// threshold - unsigned int; Minimum percent of votes to trigger hard fork. Default is 80.
// version - unsigned int; The major block version for the fork.
// votes - unsigned int; Number of votes towards hard fork.
// voting - unsigned int; Hard fork voting status.
// window - unsigned int; Number of blocks over which current votes are cast. Default is 10080 blocks.
type HardForkInfo struct {
	EarliestHeight uint   `json:"earliest_height"`
	Enabled        bool   `json:"enabled"`
	State          uint   `json:"state"`
	Status         string `json:"status"`
	Threshold      uint   `json:"threshold"`
	Version        uint   `json:"version"`
	Votes          uint   `json:"votes"`
	Voting         uint   `json:"voting"`
	Window         uint   `json:"window"`
}

// Ban
// ip - unsigned int; IP address to ban, in Int format.
// ban - boolean; Set true to ban.
// seconds - unsigned int; Number of seconds to ban node.
type Ban struct {
	Ip      uint `json:"ip"`
	Ban     bool `json:"ban, omitempty"`
	Seconds uint `json:"seconds"`
}

// Creates new ban
func NewBanRequest(ip uint, ban bool, seconds uint) Ban {
	return Ban{
		Ip:      ip,
		Ban:     ban,
		Seconds: seconds,
	}
}

// BanResponse
// bans - A list of nodes
// status - string; General RPC error code. "OK" means everything looks good.
type BanResponse struct {
	Bans   []Ban  `json:"bans"`
	Status string `json:"status"`
}
