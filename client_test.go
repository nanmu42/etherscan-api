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

func TestClient_craftURL(t *testing.T) {
	c := New(Ropsten, "abc123")

	const expected = `https://api-ropsten.etherscan.io/api?action=craftURL&apikey=abc123&four=d&four=e&four=f&module=testing&one=1&three=1&three=2&three=3&two=2`
	output := c.craftURL("testing", "craftURL", M{
		"one":   1,
		"two":   "2",
		"three": []int{1, 2, 3},
		"four":  []string{"d", "e", "f"},
	})

	if output != expected {
		t.Fatalf("output != expected, got %s, want %s", output, expected)
	}
}
