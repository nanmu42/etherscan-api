/*
 * Copyright (c) 2022 Avi Misra
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

type Option func(M)

func WithPagination(page, offset int) Option {
	return func(m M) {
		m["page"] = page
		m["offset"] = offset
	}
}

// GetLogs gets logs that match "topic" emitted by the specified "address" between the "fromBlock" and "toBlock"
func (c *Client) GetLogs(fromBlock, toBlock int, address, topic string, options ...Option) (logs []Log, err error) {
	param := M{
		"fromBlock": fromBlock,
		"toBlock":   toBlock,
		"topic0":    topic,
		"address":   address,
	}

	for _, applyOpt := range options {
		applyOpt(param)
	}

	err = c.call("logs", "getLogs", param, &logs)
	return
}
