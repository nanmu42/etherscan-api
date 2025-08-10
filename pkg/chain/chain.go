package chain

type Chain int

// ChainIDs supported as of 2025/01/02
const (
	EthereumMainnet        Chain = 1
	OpMainnet              Chain = 10
	SepoliaTestnet         Chain = 11155111
	HoleskyTestnet         Chain = 17000
	CronosMainnet          Chain = 25
	ApeChainCurtisTestnet  Chain = 33111
	ApeChainMainnet        Chain = 33139
	ArbitrumOneMainnet     Chain = 42161
	ArbitrumNovaMainnet    Chain = 42170
	CeloMainnet            Chain = 42220
	AvalancheCChain        Chain = 43114
	AvalancheFujiTestnet   Chain = 43113
	XDCMainnet             Chain = 50
	XDCApothemTestnet      Chain = 51
	BNBSmartChainMainnet   Chain = 56
	BNBSmartChainTestnet   Chain = 97
	Gnosis                 Chain = 100
	PolygonMainnet         Chain = 137
	SonicMainnet           Chain = 146
	BitTorrentChainMainnet Chain = 199
	FantomOperaMainnet     Chain = 250
	FraxtalMainnet         Chain = 252
	KromaMainnet           Chain = 255
	zkSyncMainnet          Chain = 324
	zkSyncSepoliaTestnet   Chain = 300
	MoonbeamMainnet        Chain = 1284
	MoonbaseAlphaTestnet   Chain = 1287
	MoonriverMainnet       Chain = 1285
	BitTorrentChainTestnet Chain = 1028
	PolygonZkEVMMainnet    Chain = 1101
	WEMIX30Mainnet         Chain = 1111
	WEMIX30Testnet         Chain = 1112
	FantomTestnet          Chain = 4002
	WorldMainnet           Chain = 480
	WorldSepoliaTestnet    Chain = 4801
	MantleMainnet          Chain = 5000
	MantleSepoliaTestnet   Chain = 5003
	BaseMainnet            Chain = 8453
	BaseSepoliaTestnet     Chain = 84532
	BlastMainnet           Chain = 81457
	CeloAlfajoresTestnet   Chain = 44787
	PolygonAmoyTestnet     Chain = 80002
	PolygonZkEVMCardona    Chain = 2442
	FraxtalTestnet         Chain = 2522
	KromaSepolia           Chain = 2358
	ScrollMainnet          Chain = 534352
	ScrollSepoliaTestnet   Chain = 534351
	SonicBlazeTestnet      Chain = 57054
	LineaMainnet           Chain = 59144
	LineaSepoliaTestnet    Chain = 59141
	SophonMainnet          Chain = 50104
	SophonSepoliaTestnet   Chain = 531050104
	TaikoMainnet           Chain = 167000
	TaikoHeklaTestnet      Chain = 167009
	XaiMainnet             Chain = 660279
	XaiSepoliaTestnet      Chain = 37714555429
	BlastSepoliaTestnet    Chain = 168587773
	OpSepoliaTestnet       Chain = 11155420
	ArbitrumSepoliaTestnet Chain = 421614
)

var chainNames = map[Chain]string{
	EthereumMainnet:        "Ethereum Mainnet",
	OpMainnet:              "OP Mainnet",
	SepoliaTestnet:         "Sepolia Testnet",
	HoleskyTestnet:         "Holesky Testnet",
	CronosMainnet:          "Cronos Mainnet",
	ApeChainCurtisTestnet:  "ApeChain Curtis Testnet",
	ApeChainMainnet:        "ApeChain Mainnet",
	ArbitrumOneMainnet:     "Arbitrum One Mainnet",
	ArbitrumNovaMainnet:    "Arbitrum Nova Mainnet",
	CeloMainnet:            "Celo Mainnet",
	AvalancheCChain:        "Avalanche C-Chain",
	AvalancheFujiTestnet:   "Avalanche Fuji Testnet",
	XDCMainnet:             "XDC Mainnet",
	XDCApothemTestnet:      "XDC Apothem Testnet",
	BNBSmartChainMainnet:   "BNB Smart Chain Mainnet",
	BNBSmartChainTestnet:   "BNB Smart Chain Testnet",
	Gnosis:                 "Gnosis",
	PolygonMainnet:         "Polygon Mainnet",
	SonicMainnet:           "Sonic Mainnet",
	BitTorrentChainMainnet: "BitTorrent Chain Mainnet",
	FantomOperaMainnet:     "Fantom Opera Mainnet",
	FraxtalMainnet:         "Fraxtal Mainnet",
	KromaMainnet:           "Kroma Mainnet",
	zkSyncMainnet:          "zkSync Mainnet",
	zkSyncSepoliaTestnet:   "zkSync Sepolia Testnet",
	MoonbeamMainnet:        "Moonbeam Mainnet",
	MoonbaseAlphaTestnet:   "Moonbase Alpha Testnet",
	MoonriverMainnet:       "Moonriver Mainnet",
	BitTorrentChainTestnet: "BitTorrent Chain Testnet",
	PolygonZkEVMMainnet:    "Polygon zkEVM Mainnet",
	WEMIX30Mainnet:         "WEMIX3.0 Mainnet",
	WEMIX30Testnet:         "WEMIX3.0 Testnet",
	FantomTestnet:          "Fantom Testnet",
	WorldMainnet:           "World Mainnet",
	WorldSepoliaTestnet:    "World Sepolia Testnet",
	MantleMainnet:          "Mantle Mainnet",
	MantleSepoliaTestnet:   "Mantle Sepolia Testnet",
	BaseMainnet:            "Base Mainnet",
	BaseSepoliaTestnet:     "Base Sepolia Testnet",
	BlastMainnet:           "Blast Mainnet",
	CeloAlfajoresTestnet:   "Celo Alfajores Testnet",
	PolygonAmoyTestnet:     "Polygon Amoy Testnet",
	PolygonZkEVMCardona:    "Polygon zkEVM Cardona Testnet",
	FraxtalTestnet:         "Fraxtal Testnet",
	KromaSepolia:           "Kroma Sepolia Testnet",
	ScrollMainnet:          "Scroll Mainnet",
	ScrollSepoliaTestnet:   "Scroll Sepolia Testnet",
	SonicBlazeTestnet:      "Sonic Blaze Testnet",
	LineaMainnet:           "Linea Mainnet",
	LineaSepoliaTestnet:    "Linea Sepolia Testnet",
	SophonMainnet:          "Sophon Mainnet",
	SophonSepoliaTestnet:   "Sophon Sepolia Testnet",
	TaikoMainnet:           "Taiko Mainnet",
	TaikoHeklaTestnet:      "Taiko Hekla L2 Testnet",
	XaiMainnet:             "Xai Mainnet",
	XaiSepoliaTestnet:      "Xai Sepolia Testnet",
	BlastSepoliaTestnet:    "Blast Sepolia Testnet",
	OpSepoliaTestnet:       "OP Sepolia Testnet",
	ArbitrumSepoliaTestnet: "Arbitrum Sepolia Testnet",
}

// String returns the name of the network
func (n Chain) String() string {
	if name, ok := chainNames[n]; ok {
		return name
	}
	return "Unknown Network"
}

// ID returns the chain ID as an int
func (n Chain) ID() int { return int(n) }

// NewNetwork creates a Network from a name and chain ID
func NewNetwork(name string, chainID int) Chain {
	network := Chain(chainID)

	return network
}

// GetByChainID returns a Network by its chain ID
func GetByChainID(chainID int) (Chain, bool) {
	network := Chain(chainID)
	_, exists := chainNames[network]
	return network, exists
}
