package web

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"oracle/config"
	"oracle/contract"
)

type RequestPrice struct {
	time string `uri:"time" binding:"required"`
}

func WebApiPrice(ctx *gin.Context) {
	/*var req RequestPrice
	if err := ctx.ShouldBindUri(&req); err != nil{
		ctx.JSONP(400, gin.H{"msg" : err})
		return
	}*/

	btcAddress := common.HexToAddress(config.BtcAddress)
	ethAddress := common.HexToAddress(config.EthAddress)

	var btcPrice string
	var ethPrice string

	btcPrice = contract.GetBtcPrice(btcAddress)
	ethPrice = contract.GetEthPrice(ethAddress)

	ctx.JSONP(200, gin.H{
		"btc": btcPrice,
		"eth": ethPrice,
	})
}
