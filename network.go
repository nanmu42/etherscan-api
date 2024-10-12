/*
 * Copyright (c) 2018 LI Zhennan
 *
 * Use of this work is governed by a MIT License.
 * You may find a license copy in project root.
 */

package etherscan

const (
	//// Ethereum public networks

	// Mainnet Ethereum mainnet for production
	Mainnet Network = "ethereum"
	Sepolia Network = "sepolia"

	Polygon     Network = "polygon"
	PolygonAmoy Network = "amoy"
)

var MapNetworkToURL = map[Network]string{
	Mainnet:     "https://api.etherscan.io/api?",
	Sepolia:     "https://api-sepolia.etherscan.io/api?",
	Polygon:     "https://api.polygonscan.com/api?",
	PolygonAmoy: "https://api-amoy.polygonscan.com/api?",
}

// Network is ethereum network type (mainnet, ropsten, etc)
type Network string

// SubDomain returns the subdomain of  etherscan API
// via n provided.
func (n Network) SubDomain() (sub string) {
	return string(n)
}
