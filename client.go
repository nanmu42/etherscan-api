/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import "net/http"

// Client etherscan API client
// Clients are safe for concurrent use by multiple goroutines.
type Client struct {
	network Network
	coon    *http.Client
}

// New initialize a new etherscan API client
// please use pre-defined network value
func New(network Network) *Client {
	return &Client{
		network: network,
		coon:    &http.Client{},
	}
}
