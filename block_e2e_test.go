/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"encoding/json"
	"testing"
)

func TestClient_BlockReward(t *testing.T) {
	const ans = `{"blockNumber":"2165403","timeStamp":"1472533979","blockMiner":"0x13a06d3dfe21e0db5c016c03ea7d2509f7f8d1e3","blockReward":"5314181600000000000","uncles":[{"miner":"0xbcdfc35b86bedf72f0cda046a3c16829a2ef41d1","unclePosition":"0","blockreward":"3750000000000000000"},{"miner":"0x0d0c9855c722ff0c78f21e43aa275a5b8ea60dce","unclePosition":"1","blockreward":"3750000000000000000"}],"uncleInclusionReward":"312500000000000000"}`

	reward, err := api.BlockReward(2165403)
	noError(t, err, "api.BlockReward")

	j, err := json.Marshal(reward)
	noError(t, err, "json.Marshal")
	if string(j) != ans {
		t.Errorf("api.BlockReward not working, got %s, want %s", j, ans)
	}
}
