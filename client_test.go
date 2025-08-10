/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"net/url"
	"testing"

	"github.com/nanmu42/etherscan-api/pkg/chain"
	"github.com/stretchr/testify/assert"
)

func TestClient_craftURL(t *testing.T) {
	c := NewClient(chain.EthereumMainnet, "abc123")

	const expected = `https://api.etherscan.io/v2/api?action=craftURL&apikey=abc123&chainid=1&four=d&four=e&four=f&module=testing&one=1&three=1&three=2&three=3&two=2`

	output := c.craftURL("testing", "craftURL", url.Values{
		"one":   []string{"1"},
		"two":   []string{"2"},
		"three": []string{"1", "2", "3"},
		"four":  []string{"d", "e", "f"},
	})

	assert.Equal(t, expected, output)
}