/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import (
	"fmt"
	"strconv"
)

// BlockReward gets block and uncle rewards by block number
func (c *Client) BlockReward(blockNum int) (rewards BlockRewards, err error) {
	param := M{
		"blockno": blockNum,
	}

	err = c.call("block", "getblockreward", param, &rewards)
	return
}

// BlockNumber gets the closest block number by UNIX timestamp
//
// valid closest option: before, after
func (c *Client) BlockNumber(timestamp int64, closest string) (blockNumber int, err error) {
	var blockNumberStr string

	param := M{
		"timestamp": strconv.FormatInt(timestamp, 10),
		"closest":   closest,
	}

	err = c.call("block", "getblocknobytime", param, &blockNumberStr)

	if err != nil {
		return
	}

	blockNumber, err = strconv.Atoi(blockNumberStr)
	if err != nil {
		err = fmt.Errorf("parsing block number %q: %w", blockNumberStr, err)
		return
	}

	return
}
