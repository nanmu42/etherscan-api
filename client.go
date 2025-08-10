/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"time"

	"github.com/nanmu42/etherscan-api/pkg/chain"
)

// Client etherscan API client
// Clients are safe for concurrent use by multiple goroutines.
type Client struct {
	conn    *http.Client
	key     string
	baseURL string
	chain   chain.Chain

	// Verbose when true, talks a lot
	Verbose bool

	// BeforeRequest runs before every client request, in the same goroutine.
	// May be used in rate limit.
	// Request will be aborted, if BeforeRequest returns non-nil err.
	BeforeRequest func(module, action string, param map[string]interface{}) error

	// AfterRequest runs after every client request, even when there is an error.
	AfterRequest func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error)
}

// NewClient initialize a new etherscan API client
// please use pre-defined network value
func NewClient(chain chain.Chain, APIKey string) *Client {
	return NewCustomized(Customization{
		Timeout: 30 * time.Second,
		Key:     APIKey,
		Chain:   chain,
		BaseURL: `https://api.etherscan.io/v2/api`,
	})
}

// Customization is used in NewCustomized()
type Customization struct {
	// Timeout for API call
	Timeout time.Duration
	// API key applied from Etherscan
	Key string
	// Base URL like `https://api.etherscan.io/api?`
	BaseURL string
	// When true, talks a lot
	Verbose bool
	// ChainID to be used
	Chain chain.Chain
	// HTTP Client to be used. Specifying this value will ignore the Timeout value set
	// Set your own timeout.
	Client *http.Client

	// BeforeRequest runs before every client request, in the same goroutine.
	// May be used in rate limit.
	// Request will be aborted, if BeforeRequest returns non-nil err.
	BeforeRequest func(module, action string, param map[string]interface{}) error

	// AfterRequest runs after every client request, even when there is an error.
	AfterRequest func(module, action string, param map[string]interface{}, outcome interface{}, requestErr error)
}

// NewCustomized initialize a customized API client,
// useful when calling against etherscan-family API like BscScan.
func NewCustomized(config Customization) *Client {
	var httpClient *http.Client
	if config.Client != nil {
		httpClient = config.Client
	} else {
		httpClient = &http.Client{Timeout: config.Timeout}
	}
	return &Client{
		conn:          httpClient,
		key:           config.Key,
		chain:         config.Chain,
		baseURL:       config.BaseURL,
		Verbose:       config.Verbose,
		BeforeRequest: config.BeforeRequest,
		AfterRequest:  config.AfterRequest,
	}
}

// call does almost all the dirty work.
func (c *Client) call(module, action string, values url.Values) (bytes.Buffer, error) {
	var content = bytes.Buffer{}

	req, err := http.NewRequest(http.MethodGet, c.craftURL(module, action, values), http.NoBody)
	if err != nil {
		return content, fmt.Errorf("creating request: %w", err)
	}
	req.Header.Set("User-Agent", "etherscan-api(Go)")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if c.Verbose {
		reqDump, err := httputil.DumpRequestOut(req, false)
		if err != nil {
			return content, fmt.Errorf("verbose mode: dumping request: %w", err)
		}

		fmt.Printf("\n%s\n", reqDump)

		defer func() {
			if err != nil {
				fmt.Printf("[Error] %v\n", err)
			}
		}()
	}

	res, err := c.conn.Do(req)
	if err != nil {
		return content, fmt.Errorf("sending request: %w", err)
	}
	defer res.Body.Close()

	if c.Verbose {
		resDump, err := httputil.DumpResponse(res, true)
		if err != nil {
			return content, fmt.Errorf("verbose mode:dumping response: %w", err)
		}

		fmt.Printf("%s\n", resDump)
	}

	if _, err = io.Copy(&content, res.Body); err != nil {
		return content, fmt.Errorf("reading response: %w", err)
	}

	if res.StatusCode != http.StatusOK {
		return content, fmt.Errorf("got non-200 status code; status: %v, status text: %s, response body: %s", res.StatusCode, res.Status, content.String())
	}

	return content, nil
}

// craftURL returns desired URL via param provided
func (c *Client) craftURL(module, action string, values url.Values) string {
	if values == nil {
		values = url.Values{}
	}

	values.Add("module", module)
	values.Add("action", action)
	values.Add("apikey", c.key)
	values.Add("chainid", strconv.Itoa(c.chain.ID()))

	return fmt.Sprintf("%s?%s", c.baseURL, values.Encode())
}
