/*
 * Copyright (c) 2022 Avi Misra
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

// GasEstiamte gets estiamted confirmation time (in seconds) at the given gas price
func (c *Client) GasEstimate(gasPrice int) (confirmationTimeInSec string, err error) {
	params := M{"gasPrice": gasPrice}
	err = c.call("gastracker", "gasestimate", params, &confirmationTimeInSec)
	return
}

// GasOracle gets suggested gas prices (in Gwei)
func (c *Client) GasOracle() (gasPrices GasPrices, err error) {
	err = c.call("gastracker", "gasoracle", M{}, &gasPrices)
	return
}
