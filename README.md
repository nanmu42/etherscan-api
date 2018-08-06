# etherscan-api

[![Build Status](https://travis-ci.org/nanmu42/etherscan-api.svg?branch=master)](https://travis-ci.org/nanmu42/etherscan-api)
[![Go Report Card](https://goreportcard.com/badge/github.com/nanmu42/etherscan-api)](https://goreportcard.com/report/github.com/nanmu42/etherscan-api)
[![codecov](https://codecov.io/gh/nanmu42/etherscan-api/branch/master/graph/badge.svg)](https://codecov.io/gh/nanmu42/etherscan-api)
[![GoDoc](https://godoc.org/github.com/nanmu42/etherscan-api?status.svg)](https://godoc.org/github.com/nanmu42/etherscan-api)
[![中文文档](https://img.shields.io/badge/%E4%B8%AD%E6%96%87%E6%96%87%E6%A1%A3-%E7%82%B9%E6%88%91%E9%98%85%E8%AF%BB-blue.svg)](https://github.com/nanmu42/etherscan-api/blob/master/README_ZH.md)

Go bindings to the Etherscan.io API, with nearly Full implementation(accounts, transactions, tokens, contracts, blocks, stats), full network support(Mainnet, Ropsten, Kovan, Rinkby, Tobalaba), and only depending on standard library. :wink:

# Usage

Create a API instance and off you go. :rocket:

```go
import (
	"github.com/nanmu42/etherscan-api"
	"fmt"
)

func main() {
	// create a API client for specified ethereum net
	// there are many pre-defined network in package
	client := etherscan.New(etherscan.Mainnet, "[your API key]")

	// (optional) add hooks, e.g. for rate limit
	client.BeforeRequest = func(module, action string, param map[string]interface{}) error {
		// ...
	}
	client.AfterRequest = func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error) {
		// ...
	}

	// check account balance
	balance, err := client.AccountBalance("0x281055afc982d96fab65b3a49cac8b878184cb16")
	if err != nil {
		panic(err)
	}
	// balance in wei, in *big.Int type
	fmt.Println(balance.Int())

	// check token balance
	tokenBalance, err := client.TokenBalance("contractAddress", "holderAddress")

	// check ERC20 transactions from/to a specified address
	transfers, err := client.ERC20Transfers("contractAddress", "address", startBlock, endBlock, page, offset)
}
```

You may find full method list at [GoDoc](https://godoc.org/github.com/nanmu42/etherscan-api).

# Etherscan API Key

You may apply for an API key on [etherscan](https://etherscan.io/apis).

> The Etherscan Ethereum Developer APIs are provided as a community service and without warranty, so please just use what you need and no more. They support both GET/POST requests and a rate limit of 5 requests/sec (exceed and you will be blocked). 

# Paperwork Things

I am not from Etherscan and I just find their service really useful, so I implement this. :smile:

# License

Use of this work is governed by a MIT License.

You may find a license copy in project root.
