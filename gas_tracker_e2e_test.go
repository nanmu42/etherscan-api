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
	confirmationTime, err := api.GasEstimate(20000000)
	noError(t, err, "api.GasEstimate")

	if 0 == len(confirmationTime) {
		t.Errorf("confirmationTime empty string")
	}
}

//GasOracle generates dynamic data. Best we can do is ensure all fields are populated
func TestClient_GasOracle(t *testing.T) {
	gasPrice, err := api.GasOracle()
	noError(t, err, "api.GasOrcale")

	if 0 == len(gasPrice.LastBlock) {
		t.Errorf("gasPrice.LastBlock empty string")
	}

	if 0 == len(gasPrice.SafeGasPrice) {
		t.Errorf("gasPrice.SafeGasPrice empty string")
	}

	if 0 == len(gasPrice.ProposeGasPrice) {
		t.Errorf("gasPrice.ProposeGasPrice empty string")
	}

	if 0 == len(gasPrice.FastGasPrice) {
		t.Errorf("gasPrice.FastGasPrice empty string")
	}

	if 0 == len(gasPrice.SuggestBaseFeeInGwei) {
		t.Errorf("gasPrice.SuggestBaseFeeInGwei empty string")
	}

	if 0 == len(gasPrice.GasUsedRatio) {
		t.Errorf("gasPrice.GasUsedRatio empty string")
	}

}
