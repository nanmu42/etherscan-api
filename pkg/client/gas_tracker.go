/*
 * Copyright (c) 2022 Avi Misra
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/farbodan/etherscan-api/pkg/response"
)

type GasEstimateParams struct {
	GasPrice int `json:"gasPrice"`
}

func (p GasEstimateParams) GetUrlValues() url.Values {
	values := url.Values{}
	values.Add("gasPrice", strconv.Itoa(p.GasPrice))
	return values
}

// GasEstimate gets estimated confirmation time (in seconds) at the given gas price
func (c *Client) GasEstimate(gasPrice int) (time.Duration, error) {
	param := GasEstimateParams{
		GasPrice: gasPrice,
	}

	body, err := c.call("gastracker", "gasestimate", param.GetUrlValues())
	if err != nil {
		return 0, fmt.Errorf("executing GasEstimate request: %w", err)
	}

	confTime, err := response.ReadResponse[string](body)
	if err != nil {
		return 0, fmt.Errorf("reading response: %w", err)
	}

	return time.ParseDuration(confTime + "s")
}

// GasOracle gets suggested gas prices (in Gwei)
func (c *Client) GasOracle() (response.GasPrices, error) {
	body, err := c.call("gastracker", "gasoracle", url.Values{})
	if err != nil {
		return response.GasPrices{}, fmt.Errorf("executing GasOracle request: %w", err)
	}
	return response.ReadResponse[response.GasPrices](body)
}
