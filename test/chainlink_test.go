package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"oracle/contract"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestChainLink(t *testing.T) {

	client, err := ethclient.Dial("https://goerli.infura.io/v3/75262ececf174300bb831d197389608b")
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA("dbf221e86b903180a42209fcb824ed0a9fc370b2e9555c703c353a45c9715d33")
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

	address := common.HexToAddress("0x39e0b62f308c53C8102e210FAd8aDAf47d51CAF4")
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

	a := price.String()
	a = a[:5]
	//fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0x8d490e535678e9a24360e955d75b27ad307bdfb97a1dca51d0f3035dcee3e870

	fmt.Println("+++++++++++++++++++++++++")
	fmt.Println(id)
	fmt.Println(a)
	fmt.Println("-------------------------")
}
