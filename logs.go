/*
 * Copyright (c) 2022 Avi Misra
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

// GetLogs gets logs that match "topic" emitted by the specified "address" between the "fromBlock" and "toBlock"
func (c *Client) GetLogs(address string, fromBlock, toBlock *int, topic *string, page, offset *int) (logs []Log, err error) {
	param := M{
		"address": address,
	}

	if fromBlock != nil {
		param["fromBlock"] = *fromBlock
	}

	if toBlock != nil {
		param["toBlock"] = *toBlock
	}

	if topic != nil {
		param["topic0"] = *topic
	}

	if page != nil {
		param["page"] = *page
	}

	if offset != nil {
		param["offset"] = *offset
	}

	err = c.call("logs", "getLogs", param, &logs)
	return
}
