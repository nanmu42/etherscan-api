package etherscan

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestClient_GetLogs(t *testing.T) {
	expectedLogs := []Log{
		Log{
			Address:         "0x33990122638b9132ca29c723bdf037f1a891a70c",
			Topics:          []string{"0xf63780e752c6a54a94fc52715dbc5518a3b4c3c2833d301a204226548a2a8545", "0x72657075746174696f6e00000000000000000000000000000000000000000000", "0x000000000000000000000000d9b2f59f3b5c7b3c67047d2f03c3e8052470be92"},
			Data:            "0x",
			BlockNumber:     "0x5c958",
			TransactionHash: "0x0b03498648ae2da924f961dda00dc6bb0a8df15519262b7e012b7d67f4bb7e83",
			LogIndex:        "0x",
		},
	}

	actualLogs, err := api.GetLogs(379224, 379225, "0x33990122638b9132ca29c723bdf037f1a891a70c", "0xf63780e752c6a54a94fc52715dbc5518a3b4c3c2833d301a204226548a2a8545")

	noError(t, err, "api.GetLogs")

	equal := cmp.Equal(expectedLogs, actualLogs)

	if !equal {
		t.Errorf("api.GetLogs not working\n: %s\n", cmp.Diff(expectedLogs, actualLogs))
	}
}
