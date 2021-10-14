/*
 * Copyright (c) 2021 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

// GetTransactionReceipt gets ETH Transaction Receipt for a given transaction hash
func (c *Client) GetTransactionReceipt(txHash string) (receipt TxReceipt, err error) {
	param := M{
		"txhash": txHash,
	}

	err = c.call("proxy", "eth_getTransactionReceipt", param, &receipt)
	return
}
