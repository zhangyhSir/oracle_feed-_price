package main

import (
	"fmt"
	"math/big"
	"oracle/config"
	"oracle/contract"
	priceacquire "oracle/price_acquire"
	"os"
	"os/signal"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

func main() {
	/*r := gin.Default()
	r.GET("/test", web.WebApiPrice)
	r.Run(":8081")*/

	btcAddress := common.HexToAddress(config.BtcAddress)
	ethAddress := common.HexToAddress(config.EthAddress)
	//btcAddress, ethAddress := contract.DeployContract()
	fmt.Println("btc address:" + btcAddress.String())
	fmt.Println("eth address:" + ethAddress.String())

	setPrice(btcAddress, ethAddress)

	btc := contract.GetBtcPrice(btcAddress)
	eth := contract.GetEthPrice(ethAddress)
	fmt.Println("btc chian price:" + btc)
	fmt.Println("eth chian price:" + eth)

	b := big.NewInt(20230903203331)
	btc1 := contract.GetFormerBtcPrice(b, btcAddress)
	e := big.NewInt(20230903203335)
	eth1 := contract.GetFormerEthPrice(e, ethAddress)
	fmt.Println("btc former price:" + btc1)
	fmt.Println("eth former price:" + eth1)

	c := make(chan os.Signal)
	signal.Notify(c)
	//定时器，两分钟刷新一次
	go func() {
		fmt.Println(" 开始 ", time.Now().Unix())
		myT := time.NewTimer(config.UpdateTime * time.Second)

		for {
			btcFormer := btc
			ethFormer := eth
			// 重置定时器为 120 s
			myT.Reset(20 * time.Second)
			<-myT.C
			fmt.Println(" Two minutes up. ", time.Now().Unix())

			btcId, btcPrice := priceacquire.GetPriceBtcChainLink()
			ethId, ethPrice := priceacquire.GetPriceEthChainLink()
			if btcFormer != btcPrice {
				contract.SetBtcPrice(btcAddress, btcId, btcPrice)
				btcc := contract.GetBtcPrice(btcAddress)
				fmt.Println("btc chian price:" + btcc)
			} else {
				fmt.Println(" btc price No change")
			}
			if ethFormer != ethPrice {
				contract.SetEthPrice(ethAddress, ethId, ethPrice)
				ethh := contract.GetEthPrice(ethAddress)
				fmt.Println("eth chian price:" + ethh)
			} else {
				fmt.Println(" eth price No change")
			}
			/*setPrice(btcAddress, ethAddress)
			btc := contract.GetBtcPrice(btcAddress)
			eth := contract.GetEthPrice(ethAddress)
			fmt.Println("btc chian price:" + btc)
			fmt.Println("eth chian price:" + eth)*/
		}
	}()

	<-c
}

func setPrice(btcAddress, ethAddress common.Address) {
	/*btcAverage := priceacquire.BtcAveragePrice()
	btcAverageStr := strconv.FormatFloat(btcAverage, 'f', -1, 64)
	ethAverage := priceacquire.EthAveragePrice()
	ethAverageStr := strconv.FormatFloat(ethAverage, 'f', -1, 64)*/
	btcId, btcPrice := priceacquire.GetPriceBtcChainLink()
	ethId, ethPrice := priceacquire.GetPriceEthChainLink()
	contract.SetBtcPrice(btcAddress, btcId, btcPrice)
	contract.SetEthPrice(ethAddress, ethId, ethPrice)
}
