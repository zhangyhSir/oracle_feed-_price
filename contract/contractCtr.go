package contract

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"oracle/config"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/ethereum/go-ethereum/common"
	"github.com/zksync-sdk/zksync2-go/accounts"
	"github.com/zksync-sdk/zksync2-go/clients"
)

func DeployContract() (common.Address, common.Address) {

	// 连接zkSync网络
	zp, err := clients.NewDefaultProvider(config.ZkSyncProvider)
	if err != nil {
		log.Panic(err)
	}
	defer zp.Close()

	// 根据相应链的私钥创建歌手对象
	chainID, err := zp.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	es, err := accounts.NewEthSignerFromRawPrivateKey(common.Hex2Bytes(config.PrivateKey), chainID.Int64())
	if err != nil {
		log.Fatal(err)
	}

	// 创建钱包
	w, err := accounts.NewWallet(es, zp)
	if err != nil {
		log.Panic(err)
	}

	// 读取智能合约字节码BitCoin
	btcBytecode, err := os.ReadFile("contract/CoinPrice.sol_btcPrice.zbin")
	if err != nil {
		log.Panic(err)
	}

	//部署智能合约
	btcHash, err := w.DeployWithCreate(btcBytecode, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction: ", btcHash)

	// 等待单元事务完成
	_, err = zp.WaitMined(context.Background(), btcHash)
	if err != nil {
		log.Panic(err)
	}

	btcReceipt, err := zp.GetTransactionReceipt(btcHash)
	if err != nil {
		log.Panic(err)
	}
	btcContractAddress := btcReceipt.ContractAddress
	if err != nil {
		panic(err)
	}

	//----------------
	// 读取智能合约字节码Eth
	ethBytecode, err := os.ReadFile("contract/CoinPrice.sol_ethPrice.zbin")
	if err != nil {
		log.Panic(err)
	}

	//部署智能合约
	ethHash, err := w.DeployWithCreate(ethBytecode, nil, nil, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction: ", ethHash)

	// 等待单元事务完成
	_, err = zp.WaitMined(context.Background(), ethHash)
	if err != nil {
		log.Panic(err)
	}

	ethReceipt, err := zp.GetTransactionReceipt(ethHash)
	if err != nil {
		log.Panic(err)
	}
	ethontractAddress := ethReceipt.ContractAddress
	if err != nil {
		panic(err)
	}

	return btcContractAddress, ethontractAddress
}

func GetBtcPrice(contractAddress common.Address) string {

	// 连接zkSync网络
	zp, err := clients.NewDefaultProvider(config.ZkSyncProvider)
	if err != nil {
		log.Panic(err)
	}
	defer zp.Close()

	// 与智能合约交互
	// 创建Storage智能合约实例
	coinPriceContract, err := NewBtcPrice(contractAddress, zp)

	if err != nil {
		log.Panic(err)
	}

	btcValue, err := coinPriceContract.LatestRoundData(nil)

	if err != nil {
		log.Panic(err)
	}

	return btcValue.Answer.String()
}

func GetEthPrice(contractAddress common.Address) string {

	// 连接zkSync网络
	zp, err := clients.NewDefaultProvider(config.ZkSyncProvider)
	if err != nil {
		log.Panic(err)
	}
	defer zp.Close()

	// 与智能合约交互

	// 创建Storage智能合约实例
	coinPriceContract, err := NewEthPrice(contractAddress, zp)
	if err != nil {
		log.Panic(err)
	}

	// 从存储智能合约中执行方法
	ethValue, err := coinPriceContract.LatestRoundData(nil)
	if err != nil {
		log.Panic(err)
	}

	return ethValue.Answer.String()
}

func GetFormerEthPrice(formerTime *big.Int, contractAddress common.Address) string {

	// 连接zkSync网络
	zp, err := clients.NewDefaultProvider(config.ZkSyncProvider)
	if err != nil {
		log.Panic(err)
	}
	defer zp.Close()

	// 与智能合约交互

	// 创建Storage智能合约实例
	coinPriceContract, err := NewEthPrice(contractAddress, zp)
	if err != nil {
		log.Panic(err)
	}

	// 从存储智能合约中执行方法
	ethValue, err := coinPriceContract.GetRoundData(nil, formerTime)
	if err != nil {
		log.Panic(err)
	}

	return ethValue.Answer.String()
}

func GetFormerBtcPrice(formerTime *big.Int, contractAddress common.Address) string {

	// 连接zkSync网络
	zp, err := clients.NewDefaultProvider(config.ZkSyncProvider)
	if err != nil {
		log.Panic(err)
	}
	defer zp.Close()

	// 与智能合约交互
	// 创建Storage智能合约实例
	coinPriceContract, err := NewBtcPrice(contractAddress, zp)

	if err != nil {
		log.Panic(err)
	}

	btcValue, err := coinPriceContract.GetRoundData(nil, formerTime)
	//btcValue, err := coinPriceContract.LatestRoundData(nil)

	if err != nil {
		log.Panic(err)
	}

	return btcValue.Answer.String()
}

func SetEthPrice(contractAddress common.Address, id *big.Int, price string) string {

	// 连接zkSync网络
	zp, err := clients.NewDefaultProvider(config.ZkSyncProvider)
	if err != nil {
		log.Panic(err)
	}
	defer zp.Close()

	// 根据相应链的私钥创建歌手对象
	chainID, err := zp.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	es, err := accounts.NewEthSignerFromRawPrivateKey(common.Hex2Bytes(config.PrivateKey), chainID.Int64())
	if err != nil {
		log.Fatal(err)
	}

	// 创建钱包
	w, err := accounts.NewWallet(es, zp)
	if err != nil {
		log.Panic(err)
	}

	// 创建Storage智能合约实例
	coinPriceContract, err := NewEthPrice(contractAddress, zp)
	// 读取私钥
	privateKey, err := crypto.ToECDSA(common.Hex2Bytes(config.PrivateKey))
	if err != nil {
		log.Panic(err)
	}

	// 开始配置事务参数
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, w.GetEthSigner().GetDomain().ChainId)
	if err != nil {
		log.Panic(err)
	}

	// 从具有配置的事务参数的存储智能合约执行集合方法
	res := new(big.Int)
	res.SetString(price, 10)

	//nowTime := getNowTime()

	tx, err := coinPriceContract.SetEthCoinPrice(opts, id, res)
	if err != nil {
		log.Panic(err)
	}

	// 等待事务完成
	_, err = zp.WaitMined(context.Background(), tx.Hash())
	if err != nil {
		log.Panic(err)
	}

	return "success"
}

func SetBtcPrice(contractAddress common.Address, id *big.Int, price string) string {
	var (
		PrivateKey     = "64011cbcb7703cfa86b0aab4a5700d258d324c7398ebb39cd375518f297fd21f"
		ZkSyncProvider = "https://testnet.era.zksync.dev"
	)

	// 连接zkSync网络
	zp, err := clients.NewDefaultProvider(ZkSyncProvider)
	if err != nil {
		log.Panic(err)
	}
	defer zp.Close()

	// 根据相应链的私钥创建歌手对象
	chainID, err := zp.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	es, err := accounts.NewEthSignerFromRawPrivateKey(common.Hex2Bytes(PrivateKey), chainID.Int64())
	if err != nil {
		log.Fatal(err)
	}

	// 创建钱包
	w, err := accounts.NewWallet(es, zp)
	if err != nil {
		log.Panic(err)
	}

	// 创建Storage智能合约实例
	coinPriceContract, err := NewBtcPrice(contractAddress, zp)
	// 读取私钥
	privateKey, err := crypto.ToECDSA(common.Hex2Bytes(PrivateKey))
	if err != nil {
		log.Panic(err)
	}

	// 开始配置事务参数
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, w.GetEthSigner().GetDomain().ChainId)
	if err != nil {
		log.Panic(err)
	}

	// 从具有配置的事务参数的存储智能合约执行集合方法
	res := new(big.Int)
	res.SetString(price, 10)

	//nowTime := getNowTime()

	tx, err := coinPriceContract.SetBtcCoinPrice(opts, id, res)
	if err != nil {
		log.Panic(err)
	}

	// 等待事务完成
	_, err = zp.WaitMined(context.Background(), tx.Hash())
	if err != nil {
		log.Panic(err)
	}

	return "success"
}

func getNowTime() *big.Int {
	// 获取当前的日期和时间
	currentTime := time.Now()

	customFormat := "20060102150405"
	// 使用 Format 方法将时间格式化为指定格式的字符串
	formattedTime := currentTime.Format(customFormat)

	// 创建一个新的 big.Int 变量
	bigInt := new(big.Int)

	// 使用 SetString 将字符串转换为 *big.Int
	_, success := bigInt.SetString(formattedTime, 10) // 第二个参数表示进制，10 表示十进制

	if !success {
		fmt.Println("Failed to convert the string to *big.Int")
	} else {
		fmt.Println("Converted big.Int:", bigInt)
	}
	return bigInt
}
