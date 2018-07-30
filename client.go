/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

// Client etherscan API client
// Clients are safe for concurrent use by multiple goroutines.
type Client struct {
	coon    *http.Client
	network Network
	key     string
	baseURL string

	// Verbose when true, talks a lot
	Verbose bool

	// BeforeRequest runs before every client request, in the same goroutine.
	// May be used in rate limit.
	// Request will be aborted, if BeforeRequest returns non-nil err.
	BeforeRequest func(module, action string, param map[string]interface{}) error

	// AfterRequest runs after every client request, even when there is an error.
	AfterRequest func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error)
}

// New initialize a new etherscan API client
// please use pre-defined network value
func New(network Network, APIKey string) *Client {
	return &Client{
		coon:    &http.Client{},
		network: network,
		key:     APIKey,
		baseURL: fmt.Sprintf(`https://%s.etherscan.io/api?`, network.SubDomain()),
	}
}

// call does almost all the dirty work.
func (c *Client) call(module, action string, param map[string]interface{}, outcome interface{}) (err error) {
	// todo: fire hooks
	// todo: verbose mode

	req, err := http.NewRequest(http.MethodGet, c.craftURL(param), http.NoBody)
	if err != nil {
		err = wrapErr(err, "http.NewRequest")
		return
	}

	res, err := c.coon.Do(req)
	if err != nil {
		err = wrapErr(err, "sending request")
		return
	}
	defer res.Body.Close()

	var content bytes.Buffer
	if _, err = io.Copy(&content, res.Body); err != nil {
		err = wrapErr(err, "reading response")
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status %v %s, response body: %s", res.StatusCode, res.Status, content.String())
		return
	}

	err = json.Unmarshal(content.Bytes(), outcome)
	if err != nil {
		err = wrapErr(err, "json unmarshal")
		return
	}

	return
}

// craftURL returns desired URL via param provided
func (c *Client) craftURL(param map[string]interface{}) (URL string) {
	q := url.Values{
		"apikey": []string{c.key},
	}

	for k, v := range param {
		q[k] = extractValue(v)
	}

	URL = c.baseURL + q.Encode()
	return
}
