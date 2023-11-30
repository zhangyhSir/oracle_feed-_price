package priceacquire

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"oracle/config"
	"oracle/contract"
	"oracle/proxy"
	"strconv"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func BtcAveragePrice() float64 {
	priceBF := getPrice("POST", config.BFBTCurl)
	priceCG := getPrice("GET", config.CGBTCurl)
	//priceCC := getPrice("GET", config.CCBTCurl)
	bfStr := string(priceBF)
	fmt.Println("bf price" + bfStr)
	bfF := bfStr[1:12]
	bf, err := strconv.ParseFloat(bfF, 64)
	if err != nil {
		fmt.Println("Error while parsing Bitcoin price:", err)
		return 0
	}

	var CGprice config.BitCoinPrice
	err = json.Unmarshal(priceCG, &CGprice)
	if err != nil {
		fmt.Println("Error while parsing Bitcoin price:", err)
		return 0
	}

	/*var CCprice config.Price
	err = json.Unmarshal(priceCC, &CCprice)
	if err != nil {
		fmt.Println("Error while parsing Bitcoin price:", err)
		return 0
	}*/

	/*fmt.Println(bf)
	fmt.Println(CGprice.BitCoin.Usd)*/
	//fmt.Println(CCprice.Usd)
	return (bf + CGprice.BitCoin.Usd) / 2
}

func EthAveragePrice() float64 {
	priceBF := getPrice("POST", config.BFETHurl)
	priceCG := getPrice("GET", config.CGETHurl)
	priceCC := getPrice("GET", config.CCETHurl)

	bfStr := string(priceBF)
	bfF := bfStr[1:7]
	bf, err := strconv.ParseFloat(bfF, 64)
	if err != nil {
		fmt.Println("Error while parsing Ethereum price:", err)
		return 0
	}

	var CGprice config.EthereumPrice
	err = json.Unmarshal(priceCG, &CGprice)
	if err != nil {
		fmt.Println("Error while parsing Ethereum price:", err)
		return 0
	}

	var CCprice config.Price
	err = json.Unmarshal(priceCC, &CCprice)
	if err != nil {
		fmt.Println("Error while parsing Ethereum price:", err)
		return 0
	}

	/*fmt.Println(bf)
	fmt.Println(CGprice.Ethereum.Usd)
	fmt.Println(CCprice.Usd)*/
	return (bf + CGprice.Ethereum.Usd + CCprice.Usd) / 3
}

func getPrice(method string, url string) []byte {
	data, err := proxy.UseProxy(method, url)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	return []byte(data)
}

func GetPriceBtcChainLink() (*big.Int, string) {
	client, err := ethclient.Dial(config.RpcChainlink)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(config.ChainlinkPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(config.ChainlinkBtc)
	instance, err := contract.NewContract(address, client)
	//instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	/*key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))*/

	id, price, err := instance.GetLatestData(nil)
	//tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("btc id:" + id.String())
	a := price.String()
	//a = a[:5]
	return id, a
}

func GetPriceEthChainLink() (*big.Int, string) {
	client, err := ethclient.Dial(config.RpcChainlink)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(config.ChainlinkPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth := bind.NewKeyedTransactor(privateKey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address := common.HexToAddress(config.ChainlinkEth)
	instance, err := contract.NewContract(address, client)
	//instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	/*key := [32]byte{}
	value := [32]byte{}
	copy(key[:], []byte("foo"))
	copy(value[:], []byte("bar"))*/

	id, price, err := instance.GetLatestData(nil)
	//tx, err := instance.SetItem(auth, key, value)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("eth id:" + id.String())
	a := price.String()
	//a = a[:4]
	return id, a
}
