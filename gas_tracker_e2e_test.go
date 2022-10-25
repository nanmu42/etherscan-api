/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"testing"
)

//GasEstiamte generates dynamic data. Best we can do is ensure all fields are populated
func TestClient_GasEstimate(t *testing.T) {
	_, err := api.GasEstimate(20000000)
	noError(t, err, "api.GasEstimate")
}

//GasOracle generates dynamic data. Best we can do is ensure all fields are populated
func TestClient_GasOracle(t *testing.T) {
	gasPrice, err := api.GasOracle()
	noError(t, err, "api.GasOrcale")

	if 0 == len(gasPrice.GasUsedRatio) {
		t.Errorf("gasPrice.GasUsedRatio empty")
	}

}
