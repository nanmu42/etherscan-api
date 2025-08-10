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

	"github.com/nanmu42/etherscan-api/pkg/response"
)

type LogParams struct {
	FromBlock int    `json:"fromBlock"`
	ToBlock   int    `json:"toBlock"`
	Topic0    string `json:"topic0"`
	Address   string `json:"address"`
}

func (p LogParams) GetUrlValues() url.Values {
	values := url.Values{}
	values.Add("fromBlock", strconv.Itoa(p.FromBlock))
	values.Add("toBlock", strconv.Itoa(p.ToBlock))
	values.Add("topic0", p.Topic0)
	values.Add("address", p.Address)
	return values
}

// GetLogs gets logs that match "topic" emitted by the specified "address" between the "fromBlock" and "toBlock"
func (c *Client) GetLogs(fromBlock, toBlock int, address, topic string) ([]response.Log, error) {
	param := LogParams{
		FromBlock: fromBlock,
		ToBlock:   toBlock,
		Topic0:    topic,
		Address:   address,
	}

	body, err := c.call("logs", "getLogs", param.GetUrlValues())
	if err != nil {
		return nil, fmt.Errorf("executing GetLogs request: %w", err)
	}
	return response.ReadResponse[[]response.Log](body)
}
