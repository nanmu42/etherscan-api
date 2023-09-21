package etherscan

import "testing"

// TestgetKey may fail is you run test parallel
func TestgetKey(t *testing.T) {
	countApiKey, countBackupApiKey, k := 0, 0, ""
	for i := 0; i < 10; i++ {
		k = api.getKey()
		if apiKey == k {
			countApiKey++
		} else if backupApiKey == k {
			countBackupApiKey++
		}
	}
	equal := countApiKey == 5 && countBackupApiKey == 5
	if !equal {
		t.Error("api.getKey not working")
	}
}
