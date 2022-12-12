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
	"math/big"
	"testing"
)

func TestClient_AccountBalance(t *testing.T) {
	balance, err := api.AccountBalance("0x0000000000000000000000000000000000000000")
	noError(t, err, "api.AccountBalance")

	if balance.Int().Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("rich man is no longer rich")
	}
}

func TestClient_MultiAccountBalance(t *testing.T) {
	balances, err := api.MultiAccountBalance(
		"0x0000000000000000000000000000000000000000",
		"0x0000000000000000000000000000000000000001",
		"0x0000000000000000000000000000000000000002",
		"0x0000000000000000000000000000000000000003")
	noError(t, err, "api.MultiAccountBalance")

	for i, item := range balances {
		if item.Account == "" {
			t.Errorf("bound error on index %v", i)
		}
		if item.Balance.Int().Cmp(big.NewInt(0)) != 1 {
			t.Errorf("rich man %s at index %v is no longer rich.", item.Account, i)
		}
	}
}

func TestClient_NormalTxByAddress(t *testing.T) {
	const wantLen = 19

	var a, b = 54092, 79728
	txs, err := api.NormalTxByAddress("0xde0b295669a9fd93d5f28d9ec85e40f4cb697bae", &a, &b, 1, 500, false)
	noError(t, err, "api.NormalTxByAddress")

	//j, _ := json.MarshalIndent(txs, "", "  ")
	//fmt.Printf("%s\n", j)

	if len(txs) != wantLen {
		t.Errorf("got txs length %v, want %v", len(txs), wantLen)
	}
}

func TestClient_InternalTxByAddress(t *testing.T) {
	const wantLen = 66

	var a, b = 0, 2702578
	txs, err := api.InternalTxByAddress("0x2c1ba59d6f58433fb1eaee7d20b26ed83bda51a3", &a, &b, 1, 500, false)
	noError(t, err, "api.InternalTxByAddress")

	//j, _ := json.MarshalIndent(txs, "", "  ")
	//fmt.Printf("%s\n", j)

	if len(txs) != wantLen {
		t.Errorf("got txs length %v, want %v", len(txs), wantLen)
	}
}

func TestClient_ERC20Transfers(t *testing.T) {
	const (
		wantLen1 = 3
		wantLen2 = 458
		wantLen3 = 2
	)

	var a, b = 3273004, 3328071
	var contract, address = "0xe0b7927c4af23765cb51314a0e0521a9645f0e2a", "0x4e83362442b8d1bec281594cea3050c8eb01311c"
	txs, err := api.ERC20Transfers(&contract, &address, &a, &b, 1, 500, false)
	noError(t, err, "api.ERC20Transfers 1")

	//j, _ := json.MarshalIndent(txs, "", "  ")
	//fmt.Printf("%s\n", j)

	if len(txs) != wantLen1 {
		t.Errorf("got txs length %v, want %v", len(txs), wantLen1)
	}

	txs, err = api.ERC20Transfers(nil, &address, nil, &b, 1, 500, false)
	noError(t, err, "api.ERC20Transfers 2 asc")
	if len(txs) != wantLen2 {
		t.Errorf("got txs length %v, want %v", len(txs), wantLen2)
	}

	txs, err = api.ERC20Transfers(nil, &address, nil, &b, 1, 500, true)
	noError(t, err, "api.ERC20Transfers 2 desc")

	if len(txs) != wantLen2 {
		t.Errorf("got txs length %v, want %v", len(txs), wantLen2)
	}

	// some ERC20 contract does not have valid decimals information in Etherscan,
	// which brings errors like `json: invalid use of ,string struct tag, trying to unmarshal "" into uint8`
	var specialContract = "0x5eac95ad5b287cf44e058dcf694419333b796123"
	var specialStartHeight = 6024142
	var specialEndHeight = 6485274
	txs, err = api.ERC20Transfers(&specialContract, nil, &specialStartHeight, &specialEndHeight, 1, 500, false)
	noError(t, err, "api.ERC20Transfers 2")
	if len(txs) != wantLen3 {
		t.Errorf("got txs length %v, want %v", len(txs), wantLen3)
	}
}

func TestClient_BlocksMinedByAddress(t *testing.T) {
	const wantLen = 10

	blocks, err := api.BlocksMinedByAddress("0x9dd134d14d1e65f84b706d6f205cd5b1cd03a46b", 1, wantLen)
	noError(t, err, "api.BlocksMinedByAddress")

	//j, _ := json.MarshalIndent(blocks, "", "  ")
	//fmt.Printf("%s\n", j)

	if len(blocks) != wantLen {
		t.Errorf("got txs length %v, want %v", len(blocks), wantLen)
	}
}

func TestClient_UnclesMinedByAddress(t *testing.T) {
	const wantLen = 10

	blocks, err := api.UnclesMinedByAddress("0x9dd134d14d1e65f84b706d6f205cd5b1cd03a46b", 1, wantLen)
	noError(t, err, "api.UnclesMinedByAddress")

	//j, _ := json.MarshalIndent(blocks, "", "  ")
	//fmt.Printf("%s\n", j)

	if len(blocks) != wantLen {
		t.Errorf("got txs length %v, want %v", len(blocks), wantLen)
	}
}

func TestClient_TokenBalance(t *testing.T) {
	balance, err := api.TokenBalance("0x57d90b64a1a57749b0f932f1a3395792e12e7055", "0xe04f27eb70e025b78871a2ad7eabe85e61212761")
	noError(t, err, "api.TokenBalance")

	if balance.Int().Cmp(big.NewInt(0)) != 1 {
		t.Errorf("api.TokenBalance not working, got balance %s", balance.Int().String())
	}
}

func TestClient_ERC721Transfers(t *testing.T) {
	const (
		wantLen = 351
	)

	var a, b = 4708442, 9231168
	var contract, address = "0x06012c8cf97bead5deae237070f9587f8e7a266d", "0x6975be450864c02b4613023c2152ee0743572325"
	txs, err := api.ERC721Transfers(&contract, &address, &a, &b, 1, 500, true)
	noError(t, err, "api.ERC721Transfers")

	j, _ := json.MarshalIndent(txs, "", "  ")
	fmt.Printf("%s\n", j)

	if len(txs) != wantLen {
		t.Errorf("got txs length %v, want %v", len(txs), wantLen)
	}
}

func TestClient_ERC1155Transfers(t *testing.T) {
	const (
		wantLen = 1
	)

	var a, b = 128135633, 1802672
	var contract, address = "0x3edf71a31b80Ff6a45Fdb0858eC54DE98dF047AA", "0x4b986EF20Bb83532911521FB4F6F5605122a0721"
	txs, err := api.ERC1155Transfers(&contract, &address, &b, &a, 0, 0, true)
	noError(t, err, "api.ERC721Transfers")

	j, _ := json.MarshalIndent(txs, "", "  ")
	fmt.Printf("%s\n", j)

	if len(txs) != wantLen {
		t.Errorf("got txs length %v, want %v", len(txs), wantLen)
	}
}
