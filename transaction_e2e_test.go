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

func TestClient_ExecutionStatus(t *testing.T) {
	var err error

	// bad execution
	bad, err := api.ExecutionStatus("0x15f8e5ea1079d9a0bb04a4c58ae5fe7654b5b2b4463375ff7ffb490aa0032f3a")
	noError(t, err, "api.ExecutionStatus")
	if bad.IsError != 1 && bad.ErrDescription != "Bad jump destination" {
		t.Errorf("api.ExecutionStatus not working, got\n%+v", bad)
	}

	// good execution
	good, err := api.ExecutionStatus("0xe8253035f1a1e93be24f43a3592a2c6cdbe3360e6f738ff40d46305252b44f5c")
	noError(t, err, "api.ExecutionStatus")
	if good.IsError != 0 && good.ErrDescription != "" {
		t.Errorf("api.ExecutionStatus not working, got\n%+v", good)
	}
}

func TestClient_ReceiptStatus(t *testing.T) {
	var err error

	// bad execution
	bad, err := api.ReceiptStatus("0xe7bbbeb43cf863e20ec865021d63005149c133d7822e8edc1e6cb746d6728d4e")
	noError(t, err, "api.ReceiptStatus")
	if bad != 0 {
		t.Errorf("api.ExecutionStatus not working, got %v, want 0", bad)
	}

	// good execution
	good, err := api.ReceiptStatus("0xe8253035f1a1e93be24f43a3592a2c6cdbe3360e6f738ff40d46305252b44f5c")
	noError(t, err, "api.ReceiptStatus")
	if good != 1 {
		t.Errorf("api.ExecutionStatus not working, got %v, want 1", good)
	}

	// tx before byzantium fork
	before, err := api.ReceiptStatus("0x836b403cc1516eb1337c151ff3660c3ebd528d850e6ac20a75652c705ea769f4")
	if err != ErrPreByzantiumTx {
		t.Errorf("api.ReceiptStatus does not throw error for tx before byzantium fork")
	}
	if before != -1 {
		t.Errorf("api.ExecutionStatus not working, got %v, want -1", before)
	}
}
