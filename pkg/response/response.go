package response

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"

	"github.com/farbodan/etherscan-api/pkg/types"
)

type EtherscanResponse interface {
	AccountBalance | []AccountBalance |
		NormalTx | []NormalTx |
		InternalTx | []InternalTx |
		ERC20Transfer | []ERC20Transfer |
		ERC721Transfer | []ERC721Transfer |
		ERC1155Transfer | []ERC1155Transfer |
		MinedBlock | []MinedBlock |
		ContractSource | []ContractSource |
		ExecutionStatus | []ExecutionStatus |
		BlockRewards | []BlockRewards |
		LatestPrice | []LatestPrice |
		Log | []Log |
		GasPrices | []GasPrices |
		StatusReponse | []StatusReponse |
		types.BigInt | []types.BigInt | types.Time | string
}

// envelope is the carrier of nearly every response
type envelope[T EtherscanResponse] struct {
	// 1 for good, 0 for error
	Status int `json:"status,string"`
	// OK for good, other words when Status equals 0
	Message string `json:"message"`
	// where response lies
	Result T `json:"result"`
}

type statusEnvelope struct {
	Status int `json:"status,string"`
}

func ReadResponse[T EtherscanResponse](content bytes.Buffer) (T, error) {
	var ret T

	var envelope envelope[T]
	if err := json.Unmarshal(content.Bytes(), &envelope); err != nil {
		var statusEnv statusEnvelope
		if err := json.Unmarshal(content.Bytes(), &statusEnv); err == nil {
			if statusEnv.Status != 1 {
				return ret, fmt.Errorf("unmarshaling response into %T; message=%s", ret, envelope.Message)
			}
		}
		return ret, fmt.Errorf("unmarshaling etherscan response; body=%w", err)
	}
	if envelope.Status != 1 {
		return ret, fmt.Errorf("etherscan server: %s", envelope.Message)
	}

	return envelope.Result, nil
}

// AccountBalance account and its balance in pair
type AccountBalance struct {
	Account string        `json:"account"`
	Balance *types.BigInt `json:"balance"`
}

type StatusReponse struct {
	Status string `json:"status"`
}

type (
	// EtherscanTx represents
	EtherscanTx interface {
		GetBlockNumber() int
		GetHash() string
	}

	// NormalTx holds info from normal tx query
	NormalTx struct {
		BlockNumber       int           `json:"blockNumber,string"`
		TimeStamp         types.Time    `json:"timeStamp"`
		Hash              string        `json:"hash"`
		Nonce             int           `json:"nonce,string"`
		BlockHash         string        `json:"blockHash"`
		TransactionIndex  int           `json:"transactionIndex,string"`
		From              string        `json:"from"`
		To                string        `json:"to"`
		Value             *types.BigInt `json:"value"`
		Gas               int           `json:"gas,string"`
		GasPrice          *types.BigInt `json:"gasPrice"`
		IsError           int           `json:"isError,string"`
		TxReceiptStatus   string        `json:"txreceipt_status"`
		Input             string        `json:"input"`
		ContractAddress   string        `json:"contractAddress"`
		CumulativeGasUsed int           `json:"cumulativeGasUsed,string"`
		GasUsed           int           `json:"gasUsed,string"`
		Confirmations     int           `json:"confirmations,string"`
	}

	// InternalTx holds info from internal tx query
	InternalTx struct {
		BlockNumber     int           `json:"blockNumber,string"`
		TimeStamp       types.Time    `json:"timeStamp"`
		Hash            string        `json:"hash"`
		From            string        `json:"from"`
		To              string        `json:"to"`
		Value           *types.BigInt `json:"value"`
		ContractAddress string        `json:"contractAddress"`
		Input           string        `json:"input"`
		Type            string        `json:"type"`
		Gas             int           `json:"gas,string"`
		GasUsed         int           `json:"gasUsed,string"`
		TraceID         string        `json:"traceId"`
		IsError         int           `json:"isError,string"`
		ErrCode         string        `json:"errCode"`
	}

	// ERC20Transfer holds info from ERC20 token transfer event query
	ERC20Transfer struct {
		BlockNumber       int           `json:"blockNumber,string"`
		TimeStamp         types.Time    `json:"timeStamp"`
		Hash              string        `json:"hash"`
		Nonce             int           `json:"nonce,string"`
		BlockHash         string        `json:"blockHash"`
		From              string        `json:"from"`
		ContractAddress   string        `json:"contractAddress"`
		To                string        `json:"to"`
		Value             *types.BigInt `json:"value"`
		TokenName         string        `json:"tokenName"`
		TokenSymbol       string        `json:"tokenSymbol"`
		TokenDecimal      int           `json:"tokenDecimal,string"`
		TransactionIndex  int           `json:"transactionIndex,string"`
		Gas               int           `json:"gas,string"`
		GasPrice          *types.BigInt `json:"gasPrice"`
		GasUsed           int           `json:"gasUsed,string"`
		CumulativeGasUsed int           `json:"cumulativeGasUsed,string"`
		Input             string        `json:"input"`
		Confirmations     int           `json:"confirmations,string"`
	}

	// ERC721Transfer holds info from ERC721 token transfer event query
	ERC721Transfer struct {
		BlockNumber       int           `json:"blockNumber,string"`
		TimeStamp         types.Time    `json:"timeStamp"`
		Hash              string        `json:"hash"`
		Nonce             int           `json:"nonce,string"`
		BlockHash         string        `json:"blockHash"`
		From              string        `json:"from"`
		ContractAddress   string        `json:"contractAddress"`
		To                string        `json:"to"`
		TokenID           *types.BigInt `json:"tokenID"`
		TokenName         string        `json:"tokenName"`
		TokenSymbol       string        `json:"tokenSymbol"`
		TokenDecimal      int           `json:"tokenDecimal,string"`
		TransactionIndex  int           `json:"transactionIndex,string"`
		Gas               int           `json:"gas,string"`
		GasPrice          *types.BigInt `json:"gasPrice"`
		GasUsed           int           `json:"gasUsed,string"`
		CumulativeGasUsed int           `json:"cumulativeGasUsed,string"`
		Input             string        `json:"input"`
		Confirmations     int           `json:"confirmations,string"`
	}

	// ERC1155Transfer holds info from ERC1155 token transfer event query
	ERC1155Transfer struct {
		BlockNumber int        `json:"blockNumber,string"`
		TimeStamp   types.Time `json:"timeStamp"`
		Hash        string     `json:"hash"`
		//Nonce           int     `json:"nonce,string"`
		BlockHash       string `json:"blockHash"`
		From            string `json:"from"`
		ContractAddress string `json:"contractAddress"`
		To              string `json:"to"`
		//TokenID         *BigInt `json:"tokenID"`
		TokenName   string `json:"tokenName"`
		TokenSymbol string `json:"tokenSymbol"`
		//TokenDecimal      int     `json:"tokenDecimal,string"`
		//TokenValue int `json:"tokenValue,string"`
		//TransactionIndex  int     `json:"transactionIndex,string"`
		//Gas               int     `json:"gas,string"`
		//GasPrice *BigInt `json:"gasPrice"`
		//GasUsed           int     `json:"gasUsed,string"`
		//CumulativeGasUsed int     `json:"cumulativeGasUsed,string"`
		Input string `json:"input"`
		//Confirmations     int     `json:"confirmations,string"`
	}
)

func (tx NormalTx) GetBlockNumber() int { return tx.BlockNumber }
func (tx NormalTx) GetHash() string     { return tx.Hash }

func (tx InternalTx) GetBlockNumber() int { return tx.BlockNumber }
func (tx InternalTx) GetHash() string     { return tx.Hash }

func (tx ERC20Transfer) GetBlockNumber() int { return tx.BlockNumber }
func (tx ERC20Transfer) GetHash() string     { return tx.Hash }

func (tx ERC721Transfer) GetBlockNumber() int { return tx.BlockNumber }
func (tx ERC721Transfer) GetHash() string     { return tx.Hash }

func (tx ERC1155Transfer) GetBlockNumber() int { return tx.BlockNumber }
func (tx ERC1155Transfer) GetHash() string     { return tx.Hash }

// MinedBlock holds info from query for mined block by address
type MinedBlock struct {
	BlockNumber int           `json:"blockNumber,string"`
	TimeStamp   types.Time    `json:"timeStamp"`
	BlockReward *types.BigInt `json:"blockReward"`
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

// ExecutionStatus holds info from query for transaction execution status
type ExecutionStatus struct {
	// 0 = pass, 1 = error
	IsError        int    `json:"isError,string"`
	ErrDescription string `json:"errDescription"`
}

// BlockRewards holds info from query for block and uncle rewards
type BlockRewards struct {
	BlockNumber int           `json:"blockNumber,string"`
	TimeStamp   types.Time    `json:"timeStamp"`
	BlockMiner  string        `json:"blockMiner"`
	BlockReward *types.BigInt `json:"blockReward"`
	Uncles      []struct {
		Miner         string        `json:"miner"`
		UnclePosition int           `json:"unclePosition,string"`
		BlockReward   *types.BigInt `json:"blockreward"`
	} `json:"uncles"`
	UncleInclusionReward *types.BigInt `json:"uncleInclusionReward"`
}

// LatestPrice holds info from query for latest ether price
type LatestPrice struct {
	ETHBTC          float64    `json:"ethbtc,string"`
	ETHBTCTimestamp types.Time `json:"ethbtc_timestamp"`
	ETHUSD          float64    `json:"ethusd,string"`
	ETHUSDTimestamp types.Time `json:"ethusd_timestamp"`
}

type Log struct {
	Address         string   `json:"address"`
	Topics          []string `json:"topics"`
	Data            string   `json:"data"`
	BlockNumber     string   `json:"blockNumber"`
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
