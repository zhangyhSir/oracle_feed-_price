package main

import (
	"fmt"
	"math/big"
	incrementer "oracle/contract"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestBtc(t *testing.T) {
	btcAddress := common.HexToAddress("0x631E93cc32BdEbE6078D2B8B05175Fa6620C88dA")

	//btcAddress, ethAddress := incrementer.DeployContract()

	btcAverageStr := "25933.981523333332"
	incrementer.SetBtcPrice(btcAddress, btcAverageStr)

	btc := incrementer.GetBtcPrice(btcAddress)
	b := big.NewInt(20230903191537)
	btc1 := incrementer.GetFormerBtcPrice(b, btcAddress)
	fmt.Println("btc chian price:" + btc)
	fmt.Println("+++++:" + btc1)
}
