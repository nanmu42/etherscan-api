/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// Envelope is the carrier of nearly every response
type Envelope struct {
	// 1 for good, 0 for error
	Status json.Number `json:"status"`
	// OK for good, other words when Status equals 0
	Message string `json:"message"`
	// where response lies
	Result json.RawMessage `json:"result"`
}

// AccountBalance account and its balance in pair
type AccountBalance struct {
	Account string  `json:"account"`
	Balance *BigInt `json:"balance"`
}

// NormalTx holds info from normal tx query
type NormalTx struct {
	BlockNumber       int     `json:"blockNumber,string"`
	TimeStamp         Time    `json:"timeStamp"`
	Hash              string  `json:"hash"`
	Nonce             int     `json:"nonce,string"`
	BlockHash         string  `json:"blockHash"`
	TransactionIndex  int     `json:"transactionIndex,string"`
	From              string  `json:"from"`
	To                string  `json:"to"`
	Value             *BigInt `json:"value"`
	Gas               int     `json:"gas,string"`
	GasPrice          *BigInt `json:"gasPrice"`
	IsError           int     `json:"isError,string"`
	TxReceiptStatus   string  `json:"txreceipt_status"`
	Input             string  `json:"input"`
	ContractAddress   string  `json:"contractAddress"`
	CumulativeGasUsed int     `json:"cumulativeGasUsed,string"`
	GasUsed           int     `json:"gasUsed,string"`
	Confirmations     int     `json:"confirmations,string"`
	FunctionName      string  `json:"functionName"`
	MethodId          string  `json:"methodId"`
}

// InternalTx holds info from internal tx query
type InternalTx struct {
	BlockNumber     int     `json:"blockNumber,string"`
	TimeStamp       Time    `json:"timeStamp"`
	Hash            string  `json:"hash"`
	From            string  `json:"from"`
	To              string  `json:"to"`
	Value           *BigInt `json:"value"`
	ContractAddress string  `json:"contractAddress"`
	Input           string  `json:"input"`
	Type            string  `json:"type"`
	Gas             int     `json:"gas,string"`
	GasUsed         int     `json:"gasUsed,string"`
	TraceID         string  `json:"traceId"`
	IsError         int     `json:"isError,string"`
	ErrCode         string  `json:"errCode"`
}

// ERC20Transfer holds info from ERC20 token transfer event query
type ERC20Transfer struct {
	BlockNumber       int     `json:"blockNumber,string"`
	TimeStamp         Time    `json:"timeStamp"`
	Hash              string  `json:"hash"`
	Nonce             int     `json:"nonce,string"`
	BlockHash         string  `json:"blockHash"`
	From              string  `json:"from"`
	ContractAddress   string  `json:"contractAddress"`
	To                string  `json:"to"`
	Value             *BigInt `json:"value"`
	TokenName         string  `json:"tokenName"`
	TokenSymbol       string  `json:"tokenSymbol"`
	TokenDecimal      uint8   `json:"tokenDecimal,string"`
	TransactionIndex  int     `json:"transactionIndex,string"`
	Gas               int     `json:"gas,string"`
	GasPrice          *BigInt `json:"gasPrice"`
	GasUsed           int     `json:"gasUsed,string"`
	CumulativeGasUsed int     `json:"cumulativeGasUsed,string"`
	Input             string  `json:"input"`
	Confirmations     int     `json:"confirmations,string"`
}

// ERC721Transfer holds info from ERC721 token transfer event query
type ERC721Transfer struct {
	BlockNumber       int     `json:"blockNumber,string"`
	TimeStamp         Time    `json:"timeStamp"`
	Hash              string  `json:"hash"`
	Nonce             int     `json:"nonce,string"`
	BlockHash         string  `json:"blockHash"`
	From              string  `json:"from"`
	ContractAddress   string  `json:"contractAddress"`
	To                string  `json:"to"`
	TokenID           *BigInt `json:"tokenID"`
	TokenName         string  `json:"tokenName"`
	TokenSymbol       string  `json:"tokenSymbol"`
	TokenDecimal      uint8   `json:"tokenDecimal,string"`
	TransactionIndex  int     `json:"transactionIndex,string"`
	Gas               int     `json:"gas,string"`
	GasPrice          *BigInt `json:"gasPrice"`
	GasUsed           int     `json:"gasUsed,string"`
	CumulativeGasUsed int     `json:"cumulativeGasUsed,string"`
	Input             string  `json:"input"`
	Confirmations     int     `json:"confirmations,string"`
}

// ERC1155Transfer holds info from ERC1155 token transfer event query
type ERC1155Transfer struct {
	BlockNumber       int     `json:"blockNumber,string"`
	TimeStamp         Time    `json:"timeStamp"`
	Hash              string  `json:"hash"`
	Nonce             int     `json:"nonce,string"`
	BlockHash         string  `json:"blockHash"`
	From              string  `json:"from"`
	ContractAddress   string  `json:"contractAddress"`
	To                string  `json:"to"`
	TokenID           *BigInt `json:"tokenID"`
	TokenName         string  `json:"tokenName"`
	TokenSymbol       string  `json:"tokenSymbol"`
	TokenDecimal      uint8   `json:"tokenDecimal,string"`
	TokenValue        uint8   `json:"tokenValue,string"`
	TransactionIndex  int     `json:"transactionIndex,string"`
	Gas               int     `json:"gas,string"`
	GasPrice          *BigInt `json:"gasPrice"`
	GasUsed           int     `json:"gasUsed,string"`
	CumulativeGasUsed int     `json:"cumulativeGasUsed,string"`
	Input             string  `json:"input"`
	Confirmations     int     `json:"confirmations,string"`
}

// MinedBlock holds info from query for mined block by address
type MinedBlock struct {
	BlockNumber int     `json:"blockNumber,string"`
	TimeStamp   Time    `json:"timeStamp"`
	BlockReward *BigInt `json:"blockReward"`
}

// ContractSource holds info from query for contract source code
type ContractSource struct {
	SourceCode           string `json:"SourceCode"`
	ABI                  string `json:"ABI"`
	ContractName         string `json:"ContractName"`
	CompilerVersion      string `json:"CompilerVersion"`
	OptimizationUsed     int    `json:"OptimizationUsed,string"`
	Runs                 int    `json:"Runs,string"`
	ConstructorArguments string `json:"ConstructorArguments"`
	EVMVersion           string `json:"EVMVersion"`
	Library              string `json:"Library"`
	LicenseType          string `json:"LicenseType"`
	Proxy                string `json:"Proxy"`
	Implementation       string `json:"Implementation"`
	SwarmSource          string `json:"SwarmSource"`
}

// ContractCreation holds info from query for contract creation
type ContractCreation struct {
	ContractAddress  string `json:"contractAddress"`
	ContractCreator  string `json:"contractCreator"`
	TxHash           string `json:"txHash"`
	BlockNumber      int    `json:"blockNumber,string"`
	Timestamp        Time   `json:"timestamp"`
	ContractFactory  string `json:"contractFactory"`
	CreationBytecode string `json:"creationBytecode"`
}

// ExecutionStatus holds info from query for transaction execution status
type ExecutionStatus struct {
	// 0 = pass, 1 = error
	IsError        int    `json:"isError,string"`
	ErrDescription string `json:"errDescription"`
}

// BlockRewards holds info from query for block and uncle rewards
type BlockRewards struct {
	BlockNumber int     `json:"blockNumber,string"`
	TimeStamp   Time    `json:"timeStamp"`
	BlockMiner  string  `json:"blockMiner"`
	BlockReward *BigInt `json:"blockReward"`
	Uncles      []struct {
		Miner         string  `json:"miner"`
		UnclePosition int     `json:"unclePosition,string"`
		BlockReward   *BigInt `json:"blockreward"`
	} `json:"uncles"`
	UncleInclusionReward *BigInt `json:"uncleInclusionReward"`
}

// LatestPrice holds info from query for latest ether price
type LatestPrice struct {
	ETHBTC          float64 `json:"ethbtc,string"`
	ETHBTCTimestamp Time    `json:"ethbtc_timestamp"`
	ETHUSD          float64 `json:"ethusd,string"`
	ETHUSDTimestamp Time    `json:"ethusd_timestamp"`
}

type Log struct {
	Address         string   `json:"address"`
	Topics          []string `json:"topics"`
	Data            string   `json:"data"`
	BlockNumber     string   `json:"blockNumber"`
	Timestamp       string   `json:"timeStamp"`
	TransactionHash string   `json:"transactionHash"`
	BlockHash       string   `json:"blockHash"`
	LogIndex        string   `json:"logIndex"`
	Removed         bool     `json:"removed"`
}

// GasPrices holds info for Gas Oracle queries
// Gas Prices are returned in Gwei
type GasPrices struct {
	LastBlock            int
	SafeGasPrice         float64
	ProposeGasPrice      float64
	FastGasPrice         float64
	SuggestBaseFeeInGwei float64   `json:"suggestBaseFee"`
	GasUsedRatio         []float64 `json:"gasUsedRatio"`
}

func (gp *GasPrices) UnmarshalJSON(data []byte) error {
	_gp := struct {
		LastBlock            string
		SafeGasPrice         string
		ProposeGasPrice      string
		FastGasPrice         string
		SuggestBaseFeeInGwei string `json:"suggestBaseFee"`
		GasUsedRatio         string `json:"gasUsedRatio"`
	}{}

	err := json.Unmarshal(data, &_gp)
	if err != nil {
		return err
	}

	gp.LastBlock, err = strconv.Atoi(_gp.LastBlock)
	if err != nil {
		return fmt.Errorf("Unable to convert LastBlock %s to int: %w", _gp.LastBlock, err)
	}

	gp.SafeGasPrice, err = strconv.ParseFloat(_gp.SafeGasPrice, 64)
	if err != nil {
		return fmt.Errorf("Unable to convert SafeGasPrice %s to float64: %w", _gp.SafeGasPrice, err)
	}

	gp.ProposeGasPrice, err = strconv.ParseFloat(_gp.ProposeGasPrice, 64)
	if err != nil {
		return fmt.Errorf("Unable to convert ProposeGasPrice %s to float64: %w", _gp.ProposeGasPrice, err)
	}

	gp.FastGasPrice, err = strconv.ParseFloat(_gp.FastGasPrice, 64)
	if err != nil {
		return fmt.Errorf("Unable to convert FastGasPrice %s to float64: %w", _gp.FastGasPrice, err)
	}

	gp.SuggestBaseFeeInGwei, err = strconv.ParseFloat(_gp.SuggestBaseFeeInGwei, 64)
	if err != nil {
		return fmt.Errorf("Unable to convert SuggestBaseFeeInGwei %s to float64: %w", _gp.SuggestBaseFeeInGwei, err)
	}

	gasRatios := strings.Split(_gp.GasUsedRatio, ",")
	gp.GasUsedRatio = make([]float64, len(gasRatios))
	for i, gasRatio := range gasRatios {
		gp.GasUsedRatio[i], err = strconv.ParseFloat(gasRatio, 64)
		if err != nil {
			return fmt.Errorf("Unable to convert gasRatio %s to float64: %w", gasRatio, err)
		}
	}

	return nil
}
