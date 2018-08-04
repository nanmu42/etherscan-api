/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"errors"
	"testing"
)

func Test_wrapfErr(t *testing.T) {
	const ans = "status 100: continue test"

	err := errors.New("continue test")
	err = wrapfErr(err, "%s %v", "status", "100")

	if err.Error() != ans {
		t.Fatalf("got %v, want %s", err, ans)
	}
}
