package etherscan

import (
	"strings"
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
			BlockHash:       "0xe32a9cac27f823b18454e8d69437d2af41a1b81179c6af2601f1040a72ad444b",
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

func TestClient_GetLogsWithPagination(t *testing.T) {
	expectedLogs := []Log{
		{
			Address: "0x33990122638b9132ca29c723bdf037f1a891a70c",
			Topics: []string{"0xf63780e752c6a54a94fc52715dbc5518a3b4c3c2833d301a204226548a2a8545",
				"0x72657075746174696f6e00000000000000000000000000000000000000000000",
				"0x000000000000000000000000d9b2f59f3b5c7b3c67047d2f03c3e8052470be92"},
			Data:            "0x",
			BlockNumber:     "0x5c958",
			BlockHash:       "0xe32a9cac27f823b18454e8d69437d2af41a1b81179c6af2601f1040a72ad444b",
			TransactionHash: "0x0b03498648ae2da924f961dda00dc6bb0a8df15519262b7e012b7d67f4bb7e83",
			LogIndex:        "0x",
		},
		{
			Address: "0x33990122638b9132ca29c723bdf037f1a891a70c",
			Topics: []string{"0xf63780e752c6a54a94fc52715dbc5518a3b4c3c2833d301a204226548a2a8545",
				"0x6c6f747465727900000000000000000000000000000000000000000000000000",
				"0x0000000000000000000000001f6cc3f7c927e1196c03ac49c5aff0d39c9d103d"},
			Data:            "0x",
			BlockNumber:     "0x5c965",
			BlockHash:       "0x46e257ddbafca078402d0c49ad31c79514995667132c45eddbb8ca7153b6871e",
			TransactionHash: "0x8c72ea19b48947c4339077bd9c9c09a780dfbdb1cafe68db4d29cdf2754adc11",
			LogIndex:        "0x",
		},
		{
			Address: "0x33990122638b9132ca29c723bdf037f1a891a70c",
			Topics: []string{"0xf63780e752c6a54a94fc52715dbc5518a3b4c3c2833d301a204226548a2a8545",
				"0x657870616e736500000000000000000000000000000000000000000000000000",
				"0x000000000000000000000000d7586825a3177b1c6ef341ccb18361cfbc62dd0c"},
			Data:            "0x",
			BlockNumber:     "0x6664c",
			BlockHash:       "0xe8e5f93f338348b17df97b4a6723ce4ee899eb375003c8c6c74255325915a1ed",
			TransactionHash: "0xf9c4f7843dc1f9bf6d248ebe0033b2c51398255eb8973f4af4bae2c3f9313a78",
			LogIndex:        "0x",
		},
		{
			Address: "0x33990122638b9132ca29c723bdf037f1a891a70c",
			Topics: []string{"0xf63780e752c6a54a94fc52715dbc5518a3b4c3c2833d301a204226548a2a8545",
				"0x726f6f7473746f636b0000000000000000000000000000000000000000000000",
				"0x000000000000000000000000697c0123cd103cf4c3446b271e0970f109eae78c"},
			Data:            "0x",
			BlockNumber:     "0x66650",
			BlockHash:       "0x49dd65b6e2f97b294d18ec97e26b349b8215f26964be1f26e0639f79995b1a11",
			TransactionHash: "0xb190139d14140cf98035c5b78fe3b2629db2787ef234258633278985fa99a13a",
			LogIndex:        "0x",
		},
	}

	tests := []struct {
		name       string
		page, off  int
		wantSlice  []Log
		wantErr    bool
		errContain string
	}{
		{
			name: "no page or offset returns all logs",
			page: 0, off: 0,
			wantSlice: expectedLogs,
		},
		{
			name: "page=1, offset=0 returns all logs",
			page: 1, off: 0,
			wantSlice: expectedLogs,
		},
		{
			name: "page=2, offset=0 returns all logs",
			page: 2, off: 0,
			wantSlice: expectedLogs,
		},
		{
			name: "page=1, offset=2 returns first 2 logs of 4",
			page: 1, off: 2,
			wantSlice: expectedLogs[0:2],
		},
		{
			name: "page=2, offset=2 returns last 2 logs of 4",
			page: 2, off: 2,
			wantSlice: expectedLogs[2:4],
		},
		{
			name: "page=5, offset=1 returns no logs as page is beyond available logs",
			page: 5, off: 1,
			wantSlice:  []Log{},
			wantErr:    true,
			errContain: "No records found",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := api.GetLogs(
				379224,
				430000,
				"0x33990122638b9132ca29c723bdf037f1a891a70c",
				"0xf63780e752c6a54a94fc52715dbc5518a3b4c3c2833d301a204226548a2a8545",
				WithPagination(tt.page, tt.off),
			)

			if tt.wantErr {
				if err == nil {
					t.Fatal("expected error, got nil")
				}
				if !strings.Contains(err.Error(), tt.errContain) {
					t.Fatalf("expected error to contain %q, got %v", tt.errContain, err)
				}
				if len(got) != 0 {
					t.Errorf("expected no logs on error, but got %d entries", len(got))
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if diff := cmp.Diff(tt.wantSlice, got); diff != "" {
				t.Errorf("logs mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
