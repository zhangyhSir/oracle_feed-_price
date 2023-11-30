package config

//Get请求

const CGBTCurl = "https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=usd"
const CGETHurl = "https://api.coingecko.com/api/v3/simple/price?ids=ethereum&vs_currencies=usd"

// BitCoinPrice 比特币json数据结构
type BitCoinPrice struct {
	BitCoin BitCoin `json:"bitcoin"`
}

type BitCoin struct {
	Usd float64 `json:"usd"`
}

// EthereumPrice 以太坊数据结构
type EthereumPrice struct {
	Ethereum Ethereum `json:"ethereum"`
}

type Ethereum struct {
	Usd float64 `json:"usd"`
}
