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
	"net/http/httputil"
	"net/url"
	"time"
)

// Client etherscan API client
// Clients are safe for concurrent use by multiple goroutines.
type Client struct {
	coon    *http.Client
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
	return NewCustomized(Customization{
		Timeout: 30 * time.Second,
		Key:     APIKey,
		BaseURL: fmt.Sprintf(`https://%s.etherscan.io/api?`, network.SubDomain()),
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
		httpClient = http.DefaultClient
		httpClient.Timeout = config.Timeout
	}
	return &Client{
		coon:          httpClient,
		key:           config.Key,
		baseURL:       config.BaseURL,
		Verbose:       config.Verbose,
		BeforeRequest: config.BeforeRequest,
		AfterRequest:  config.AfterRequest,
	}
}

// call does almost all the dirty work.
func (c *Client) call(module, action string, param map[string]interface{}, outcome interface{}) (err error) {
	// fire hooks if in need
	if c.BeforeRequest != nil {
		err = c.BeforeRequest(module, action, param)
		if err != nil {
			err = wrapErr(err, "beforeRequest")
			return
		}
	}
	if c.AfterRequest != nil {
		defer c.AfterRequest(module, action, param, outcome, err)
	}

	// recover if there shall be an panic
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("[ouch! panic recovered] please report this with what you did and what you expected, panic detail: %v", r)
		}
	}()

	req, err := http.NewRequest(http.MethodGet, c.craftURL(module, action, param), http.NoBody)
	if err != nil {
		err = wrapErr(err, "http.NewRequest")
		return
	}
	req.Header.Set("User-Agent", "etherscan-api(Go)")
	req.Header.Set("Content-Type", "application/json; charset=utf-8")

	if c.Verbose {
		var reqDump []byte
		reqDump, err = httputil.DumpRequestOut(req, false)
		if err != nil {
			err = wrapErr(err, "verbose mode req dump failed")
			return
		}

		fmt.Printf("\n%s\n", reqDump)

		defer func() {
			if err != nil {
				fmt.Printf("[Error] %v\n", err)
			}
		}()
	}

	res, err := c.coon.Do(req)
	if err != nil {
		err = wrapErr(err, "sending request")
		return
	}
	defer res.Body.Close()

	if c.Verbose {
		var resDump []byte
		resDump, err = httputil.DumpResponse(res, true)
		if err != nil {
			err = wrapErr(err, "verbose mode res dump failed")
			return
		}

		fmt.Printf("%s\n", resDump)
	}

	var content bytes.Buffer
	if _, err = io.Copy(&content, res.Body); err != nil {
		err = wrapErr(err, "reading response")
		return
	}

	if res.StatusCode != http.StatusOK {
		err = fmt.Errorf("response status %v %s, response body: %s", res.StatusCode, res.Status, content.String())
		return
	}

	var envelope Envelope
	err = json.Unmarshal(content.Bytes(), &envelope)
	if err != nil {
		err = wrapErr(err, "json unmarshal envelope")
		return
	}
	if envelope.Status != 1 {
		err = fmt.Errorf("etherscan server: %s", envelope.Message)
		return
	}

	// workaround for missing tokenDecimal for some tokentx calls
	if action == "tokentx" {
		err = json.Unmarshal(bytes.Replace(envelope.Result, []byte(`"tokenDecimal":""`), []byte(`"tokenDecimal":"0"`), -1), outcome)
	} else {
		err = json.Unmarshal(envelope.Result, outcome)
	}
	if err != nil {
		err = wrapErr(err, "json unmarshal outcome")
		return
	}

	return
}

// craftURL returns desired URL via param provided
func (c *Client) craftURL(module, action string, param map[string]interface{}) (URL string) {
	q := url.Values{
		"module": []string{module},
		"action": []string{action},
		"apikey": []string{c.key},
	}

	for k, v := range param {
		q[k] = extractValue(v)
	}

	URL = c.baseURL + q.Encode()
	return
}
