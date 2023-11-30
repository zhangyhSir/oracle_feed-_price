package config

//Get请求

const CCBTCurl = "https://min-api.cryptocompare.com/data/price?fsym=BTC&tsyms=USD"
const CCETHurl = "https://min-api.cryptocompare.com/data/price?fsym=ETH&tsyms=USD"

type Price struct {
	Usd float64 `json:"USD"`
}
