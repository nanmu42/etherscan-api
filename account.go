/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

import "math/big"

// SingleAddrBalance get latest address balance
func (c *Client) SingleAddrBalance(address string) (balance *big.Int, err error) {
	param := M{
		"tag":     "latest",
		"address": address,
	}

	var b BigInt
	err = c.call("account", "balance", param, &b)
	balance = (*big.Int)(&b)
	return
}
