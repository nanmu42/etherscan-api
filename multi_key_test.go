package etherscan

import "testing"

// TestGetKey may fail is you run test parallel
func TestGetKey(t *testing.T) {
	countApiKey, countBackupApiKey, k := 0, 0, ""
	for i := 0; i < 10; i++ {
		k = api.getKey()
		t.Logf("key: %s", k)
		if apiKey == k {
			countApiKey++
		} else if backupApiKey == k {
			countBackupApiKey++
		}
	}
	equal := countApiKey == 5 && countBackupApiKey == 5
	if !equal {
		t.Errorf("api.getKey not working, expected 5 for each key, got main:%d , backup %d", countApiKey, countBackupApiKey)
	}
}
