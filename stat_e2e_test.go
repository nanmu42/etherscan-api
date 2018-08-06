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

func TestClient_EtherTotalSupply(t *testing.T) {
	totalSupply, err := api.EtherTotalSupply()
	noError(t, err, "api.EtherTotalSupply")

	if totalSupply.Int().Cmp(big.NewInt(100)) != 1 {
		t.Errorf("api.EtherTotalSupply not working, totalSupply is %s", totalSupply.Int().String())
	}
}

func TestClient_EtherLatestPrice(t *testing.T) {
	latest, err := api.EtherLatestPrice()
	noError(t, err, "api.EtherLatestPrice")

	if latest.ETHBTC == 0 {
		t.Errorf("ETHBTC got 0")
	}
	if latest.ETHBTCTimestamp.Time().IsZero() {
		t.Errorf("ETHBTCTimestamp is zero")
	}
	if latest.ETHUSD == 0 {
		t.Errorf("ETHUSD got 0")
	}
	if latest.ETHUSDTimestamp.Time().IsZero() {
		t.Errorf("ETHUSDTimestamp is zero")
	}
}

func TestClient_TokenTotalSupply(t *testing.T) {
	totalSupply, err := api.TokenTotalSupply("0x57d90b64a1a57749b0f932f1a3395792e12e7055")
	noError(t, err, "api.TokenTotalSupply")

	if totalSupply.Int().Cmp(big.NewInt(100)) != 1 {
		t.Errorf("api.TokenTotalSupply not working, totalSupply is %s", totalSupply.Int().String())
	}
}
