/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"math/big"
	"testing"
)

func TestClient_AccountBalance(t *testing.T) {
	balance, err := api.AccountBalance("0x281055afc982d96fab65b3a49cac8b878184cb16")
	noError(t, err, "api.AccountBalance")

	if balance.Int().Cmp(big.NewInt(0)) != 1 {
		t.Fatalf("rich man is no longer rich")
	}
}

func TestClient_MultiAccountBalance(t *testing.T) {
	balances, err := api.MultiAccountBalance(
		"0x281055afc982d96fab65b3a49cac8b878184cb16",
		"0x6f46cf5569aefa1acc1009290c8e043747172d89",
		"0x90e63c3d53e0ea496845b7a03ec7548b70014a91",
		"0x53d284357ec70ce289d6d64134dfac8e511c8a3d")
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
	)

	var a, b = 3273004, 3328071
	var contract, address = "0xe0b7927c4af23765cb51314a0e0521a9645f0e2a", "0x4e83362442b8d1bec281594cea3050c8eb01311c"
	txs, err := api.ERC20Transfers(&contract, &address, &a, &b, 1, 500)
	noError(t, err, "api.ERC20Transfers 1")

	//j, _ := json.MarshalIndent(txs, "", "  ")
	//fmt.Printf("%s\n", j)

	if len(txs) != wantLen1 {
		t.Errorf("got txs length %v, want %v", len(txs), wantLen1)
	}

	txs, err = api.ERC20Transfers(nil, &address, nil, &b, 1, 500)
	noError(t, err, "api.ERC20Transfers 2")
	if len(txs) != wantLen2 {
		t.Errorf("got txs length %v, want %v", len(txs), wantLen2)
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
