/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import "strconv"

// BlockReward gets block and uncle rewards by block number
func (c *Client) BlockReward(blockNum int) (rewards BlockRewards, err error) {
	param := M{
		"blockno": blockNum,
	}

	err = c.call("block", "getblockreward", param, &rewards)
	return
}

// BlockReward gets closest block number by UNIX timestamp
func (c *Client) BlockNumber(timestamp int64, closest string) (blockNumber BlockNumberFromTimestamp, err error) {
	var result string
	param := M{
		"timestamp": strconv.Itoa(int(timestamp)),
		"closest":   closest,
	}

	err = c.call("block", "getblocknobytime", param, &result)

	if err != nil {
		return
	}

	blockNum, err := strconv.ParseInt(result, 10, 64)
	blockNumber.BlockNumber = int(blockNum)
	return
}
