// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IAggregatorV3MetaData contains all meta data concerning the IAggregatorV3 contract.
var IAggregatorV3MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"_roundId\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IAggregatorV3ABI is the input ABI used to generate the binding from.
// Deprecated: Use IAggregatorV3MetaData.ABI instead.
var IAggregatorV3ABI = IAggregatorV3MetaData.ABI

// IAggregatorV3 is an auto generated Go binding around an Ethereum contract.
type IAggregatorV3 struct {
	IAggregatorV3Caller     // Read-only binding to the contract
	IAggregatorV3Transactor // Write-only binding to the contract
	IAggregatorV3Filterer   // Log filterer for contract events
}

// IAggregatorV3Caller is an auto generated read-only Go binding around an Ethereum contract.
type IAggregatorV3Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAggregatorV3Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IAggregatorV3Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAggregatorV3Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IAggregatorV3Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IAggregatorV3Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IAggregatorV3Session struct {
	Contract     *IAggregatorV3    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IAggregatorV3CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IAggregatorV3CallerSession struct {
	Contract *IAggregatorV3Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// IAggregatorV3TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IAggregatorV3TransactorSession struct {
	Contract     *IAggregatorV3Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// IAggregatorV3Raw is an auto generated low-level Go binding around an Ethereum contract.
type IAggregatorV3Raw struct {
	Contract *IAggregatorV3 // Generic contract binding to access the raw methods on
}

// IAggregatorV3CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IAggregatorV3CallerRaw struct {
	Contract *IAggregatorV3Caller // Generic read-only contract binding to access the raw methods on
}

// IAggregatorV3TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IAggregatorV3TransactorRaw struct {
	Contract *IAggregatorV3Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIAggregatorV3 creates a new instance of IAggregatorV3, bound to a specific deployed contract.
func NewIAggregatorV3(address common.Address, backend bind.ContractBackend) (*IAggregatorV3, error) {
	contract, err := bindIAggregatorV3(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IAggregatorV3{IAggregatorV3Caller: IAggregatorV3Caller{contract: contract}, IAggregatorV3Transactor: IAggregatorV3Transactor{contract: contract}, IAggregatorV3Filterer: IAggregatorV3Filterer{contract: contract}}, nil
}

// NewIAggregatorV3Caller creates a new read-only instance of IAggregatorV3, bound to a specific deployed contract.
func NewIAggregatorV3Caller(address common.Address, caller bind.ContractCaller) (*IAggregatorV3Caller, error) {
	contract, err := bindIAggregatorV3(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IAggregatorV3Caller{contract: contract}, nil
}

// NewIAggregatorV3Transactor creates a new write-only instance of IAggregatorV3, bound to a specific deployed contract.
func NewIAggregatorV3Transactor(address common.Address, transactor bind.ContractTransactor) (*IAggregatorV3Transactor, error) {
	contract, err := bindIAggregatorV3(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IAggregatorV3Transactor{contract: contract}, nil
}

// NewIAggregatorV3Filterer creates a new log filterer instance of IAggregatorV3, bound to a specific deployed contract.
func NewIAggregatorV3Filterer(address common.Address, filterer bind.ContractFilterer) (*IAggregatorV3Filterer, error) {
	contract, err := bindIAggregatorV3(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IAggregatorV3Filterer{contract: contract}, nil
}

// bindIAggregatorV3 binds a generic wrapper to an already deployed contract.
func bindIAggregatorV3(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := IAggregatorV3MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAggregatorV3 *IAggregatorV3Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAggregatorV3.Contract.IAggregatorV3Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAggregatorV3 *IAggregatorV3Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAggregatorV3.Contract.IAggregatorV3Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAggregatorV3 *IAggregatorV3Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAggregatorV3.Contract.IAggregatorV3Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IAggregatorV3 *IAggregatorV3CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IAggregatorV3.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IAggregatorV3 *IAggregatorV3TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IAggregatorV3.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IAggregatorV3 *IAggregatorV3TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IAggregatorV3.Contract.contract.Transact(opts, method, params...)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IAggregatorV3 *IAggregatorV3Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IAggregatorV3.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IAggregatorV3 *IAggregatorV3Session) Decimals() (uint8, error) {
	return _IAggregatorV3.Contract.Decimals(&_IAggregatorV3.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IAggregatorV3 *IAggregatorV3CallerSession) Decimals() (uint8, error) {
	return _IAggregatorV3.Contract.Decimals(&_IAggregatorV3.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_IAggregatorV3 *IAggregatorV3Caller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IAggregatorV3.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_IAggregatorV3 *IAggregatorV3Session) Description() (string, error) {
	return _IAggregatorV3.Contract.Description(&_IAggregatorV3.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_IAggregatorV3 *IAggregatorV3CallerSession) Description() (string, error) {
	return _IAggregatorV3.Contract.Description(&_IAggregatorV3.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_IAggregatorV3 *IAggregatorV3Caller) GetRoundData(opts *bind.CallOpts, _roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _IAggregatorV3.contract.Call(opts, &out, "getRoundData", _roundId)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_IAggregatorV3 *IAggregatorV3Session) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _IAggregatorV3.Contract.GetRoundData(&_IAggregatorV3.CallOpts, _roundId)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 _roundId) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_IAggregatorV3 *IAggregatorV3CallerSession) GetRoundData(_roundId *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _IAggregatorV3.Contract.GetRoundData(&_IAggregatorV3.CallOpts, _roundId)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_IAggregatorV3 *IAggregatorV3Caller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _IAggregatorV3.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_IAggregatorV3 *IAggregatorV3Session) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _IAggregatorV3.Contract.LatestRoundData(&_IAggregatorV3.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_IAggregatorV3 *IAggregatorV3CallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _IAggregatorV3.Contract.LatestRoundData(&_IAggregatorV3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_IAggregatorV3 *IAggregatorV3Caller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IAggregatorV3.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_IAggregatorV3 *IAggregatorV3Session) Version() (*big.Int, error) {
	return _IAggregatorV3.Contract.Version(&_IAggregatorV3.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_IAggregatorV3 *IAggregatorV3CallerSession) Version() (*big.Int, error) {
	return _IAggregatorV3.Contract.Version(&_IAggregatorV3.CallOpts)
}

// BtcPriceMetaData contains all meta data concerning the BtcPrice contract.
var BtcPriceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"btcCoinPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"id\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"id\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"setBtcCoinPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x0002000000000002000100000000000200010000000103550000006001100270000000600010019d0000008001000039000000400010043f0000000101200190000000300000c13d0000000001000031000000040110008c000001100000413d0000000101000367000000000101043b000000e001100270000000630210009c0000003d0000213d000000690210009c0000005e0000213d0000006c0210009c000000b10000613d0000006d0110009c000001100000c13d0000000001000416000000000101004b000001100000c13d000000040100008a00000000011000310000006e02000041000000000301004b000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001100000c13d000000400100043d00000000000104350000006002000041000000600310009c0000000001028019000000400110021000000074011001c70000017b0001042e0000000001000416000000000101004b000001100000c13d000000000100041a00000061011001970000000002000411000000000121019f000000000010041b00000020010000390000010000100443000001200000044300000062010000410000017b0001042e000000640210009c0000008c0000213d000000670210009c000000c50000613d000000680110009c000001100000c13d0000000001000416000000000101004b000001100000c13d000000040100008a00000000011000310000006e02000041000000000301004b000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001100000c13d000000000100041a0000007001100197000000400200043d00000000001204350000006001000041000000600320009c0000000001024019000000400110021000000074011001c70000017b0001042e0000006a0210009c000000e00000613d0000006b0110009c000001100000c13d0000000001000416000000000101004b000001100000c13d000000040100008a00000000011000310000006e02000041000000200310008c000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001100000c13d00000004010000390000000101100367000000000101043b0000006f0210009c000001100000213d00000000001004350000000201000039000000200010043f0000000001000019017a01580000040f000000000201041a000000400100043d000100000001001d017a016a0000040f000000010400002900000000014100490000006002000041000000600310009c0000000001028019000000600340009c000000000204401900000040022002100000006001100210000000000121019f0000017b0001042e000000650210009c000000fa0000613d000000660110009c000001100000c13d0000000001000416000000000101004b000001100000c13d000000040100008a00000000011000310000006e02000041000000000301004b000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001100000c13d0000000101000039000000000201041a000000400100043d000100000001001d017a016a0000040f000000010400002900000000014100490000006002000041000000600310009c0000000001028019000000600340009c000000000204401900000040022002100000006001100210000000000121019f0000017b0001042e0000000001000416000000000101004b000001100000c13d000000040100008a00000000011000310000006e02000041000000000301004b000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001100000c13d000000800000043f00000078010000410000017b0001042e0000000001000416000000000101004b000001100000c13d000000040100008a00000000011000310000006e02000041000000000301004b000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001100000c13d0000000101000039000000000101041a000000400200043d00000000001204350000006001000041000000600320009c0000000001024019000000400110021000000074011001c70000017b0001042e0000000001000416000000000101004b000001100000c13d000000040100008a00000000011000310000006e02000041000000000301004b000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001100000c13d000000400300043d000000750130009c000001120000413d000000760100004100000000001004350000004101000039000000040010043f00000077010000410000017c000104300000000001000416000000000101004b000001100000c13d000000040100008a00000000011000310000006e02000041000000400310008c000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001100000c13d00000001010003670000000402100370000000000202043b0000006f0320009c000001330000a13d00000000010000190000017c000104300000002002300039000000400020043f00000000000304350000002004000039000000400100043d0000000004410436000000000303043300000000003404350000004004100039000000000503004b000001250000613d000000000500001900000000064500190000000007250019000000000707043300000000007604350000002005500039000000000635004b0000011e0000413d000000000243001900000000000204350000005f02300039000000200300008a000000000232016f0000006003000041000000600420009c0000000002038019000000600410009c000000000103801900000040011002100000006002200210000000000112019f0000017b0001042e000000000300041a00000070033001970000000004000411000000000334004b000001460000c13d0000002401100370000000000101043b000100000001001d00000000002004350000000201000039000000200010043f0000000001000019017a01580000040f0000000102000029000000000021041b0000000101000039000000000021041b00000000010000190000017b0001042e000000400100043d00000044021000390000007103000041000000000032043500000024021000390000001d030000390000000000320435000000720200004100000000002104350000000402100039000000200300003900000000003204350000006002000041000000600310009c0000000001028019000000400110021000000073011001c70000017c0001043000000060020000410000000003000414000000600430009c0000000003028019000000600410009c00000000010280190000004001100210000000c002300210000000000112019f00000079011001c70000801002000039017a01750000040f0000000102200190000001680000613d000000000101043b000000000001042d00000000010000190000017c00010430000000200310003900000000002304350000008002100039000000000002043500000060021000390000000000020435000000400210003900000000000204350000000000010435000000a001100039000000000001042d00000178002104230000000102000039000000000001042d0000000002000019000000000001042d0000017a000004320000017b0001042e0000017c0001043000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffff0000000000000000000000000000000000000000000000020000000000000000000000000000004000000100000000000000000000000000000000000000000000000000000000000000000000000000ae12765500000000000000000000000000000000000000000000000000000000f15fcd1d00000000000000000000000000000000000000000000000000000000f15fcd1e00000000000000000000000000000000000000000000000000000000feaf968c00000000000000000000000000000000000000000000000000000000ae12765600000000000000000000000000000000000000000000000000000000ce606ee0000000000000000000000000000000000000000000000000000000007284e415000000000000000000000000000000000000000000000000000000007284e416000000000000000000000000000000000000000000000000000000009a6fc8f500000000000000000000000000000000000000000000000000000000313ce5670000000000000000000000000000000000000000000000000000000054fd4d50800000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffff000000000000000000000000ffffffffffffffffffffffffffffffffffffffff796f752068617665206e6f206163717569726520746f206368616e676500000008c379a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000020000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffe04e487b71000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024000000000000000000000000000000000000000000000000000000000000002000000080000000000000000002000000000000000000000000000000000000400000000000000000000000007db96ff8d57727fa755d43462d5188163e49a384951ace78c52716e6f785fc54",
}

// BtcPriceABI is the input ABI used to generate the binding from.
// Deprecated: Use BtcPriceMetaData.ABI instead.
var BtcPriceABI = BtcPriceMetaData.ABI

// BtcPriceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BtcPriceMetaData.Bin instead.
var BtcPriceBin = BtcPriceMetaData.Bin

// DeployBtcPrice deploys a new Ethereum contract, binding an instance of BtcPrice to it.
func DeployBtcPrice(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BtcPrice, error) {
	parsed, err := BtcPriceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BtcPriceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BtcPrice{BtcPriceCaller: BtcPriceCaller{contract: contract}, BtcPriceTransactor: BtcPriceTransactor{contract: contract}, BtcPriceFilterer: BtcPriceFilterer{contract: contract}}, nil
}

// BtcPrice is an auto generated Go binding around an Ethereum contract.
type BtcPrice struct {
	BtcPriceCaller     // Read-only binding to the contract
	BtcPriceTransactor // Write-only binding to the contract
	BtcPriceFilterer   // Log filterer for contract events
}

// BtcPriceCaller is an auto generated read-only Go binding around an Ethereum contract.
type BtcPriceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcPriceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BtcPriceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcPriceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BtcPriceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BtcPriceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BtcPriceSession struct {
	Contract     *BtcPrice         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BtcPriceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BtcPriceCallerSession struct {
	Contract *BtcPriceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// BtcPriceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BtcPriceTransactorSession struct {
	Contract     *BtcPriceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// BtcPriceRaw is an auto generated low-level Go binding around an Ethereum contract.
type BtcPriceRaw struct {
	Contract *BtcPrice // Generic contract binding to access the raw methods on
}

// BtcPriceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BtcPriceCallerRaw struct {
	Contract *BtcPriceCaller // Generic read-only contract binding to access the raw methods on
}

// BtcPriceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BtcPriceTransactorRaw struct {
	Contract *BtcPriceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBtcPrice creates a new instance of BtcPrice, bound to a specific deployed contract.
func NewBtcPrice(address common.Address, backend bind.ContractBackend) (*BtcPrice, error) {
	contract, err := bindBtcPrice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BtcPrice{BtcPriceCaller: BtcPriceCaller{contract: contract}, BtcPriceTransactor: BtcPriceTransactor{contract: contract}, BtcPriceFilterer: BtcPriceFilterer{contract: contract}}, nil
}

// NewBtcPriceCaller creates a new read-only instance of BtcPrice, bound to a specific deployed contract.
func NewBtcPriceCaller(address common.Address, caller bind.ContractCaller) (*BtcPriceCaller, error) {
	contract, err := bindBtcPrice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BtcPriceCaller{contract: contract}, nil
}

// NewBtcPriceTransactor creates a new write-only instance of BtcPrice, bound to a specific deployed contract.
func NewBtcPriceTransactor(address common.Address, transactor bind.ContractTransactor) (*BtcPriceTransactor, error) {
	contract, err := bindBtcPrice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BtcPriceTransactor{contract: contract}, nil
}

// NewBtcPriceFilterer creates a new log filterer instance of BtcPrice, bound to a specific deployed contract.
func NewBtcPriceFilterer(address common.Address, filterer bind.ContractFilterer) (*BtcPriceFilterer, error) {
	contract, err := bindBtcPrice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BtcPriceFilterer{contract: contract}, nil
}

// bindBtcPrice binds a generic wrapper to an already deployed contract.
func bindBtcPrice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BtcPriceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BtcPrice *BtcPriceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BtcPrice.Contract.BtcPriceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BtcPrice *BtcPriceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BtcPrice.Contract.BtcPriceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BtcPrice *BtcPriceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BtcPrice.Contract.BtcPriceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BtcPrice *BtcPriceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BtcPrice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BtcPrice *BtcPriceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BtcPrice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BtcPrice *BtcPriceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BtcPrice.Contract.contract.Transact(opts, method, params...)
}

// BtcCoinPrice is a free data retrieval call binding the contract method 0xae127656.
//
// Solidity: function btcCoinPrice() view returns(int256)
func (_BtcPrice *BtcPriceCaller) BtcCoinPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BtcPrice.contract.Call(opts, &out, "btcCoinPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BtcCoinPrice is a free data retrieval call binding the contract method 0xae127656.
//
// Solidity: function btcCoinPrice() view returns(int256)
func (_BtcPrice *BtcPriceSession) BtcCoinPrice() (*big.Int, error) {
	return _BtcPrice.Contract.BtcCoinPrice(&_BtcPrice.CallOpts)
}

// BtcCoinPrice is a free data retrieval call binding the contract method 0xae127656.
//
// Solidity: function btcCoinPrice() view returns(int256)
func (_BtcPrice *BtcPriceCallerSession) BtcCoinPrice() (*big.Int, error) {
	return _BtcPrice.Contract.BtcCoinPrice(&_BtcPrice.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_BtcPrice *BtcPriceCaller) ContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BtcPrice.contract.Call(opts, &out, "contractOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_BtcPrice *BtcPriceSession) ContractOwner() (common.Address, error) {
	return _BtcPrice.Contract.ContractOwner(&_BtcPrice.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_BtcPrice *BtcPriceCallerSession) ContractOwner() (common.Address, error) {
	return _BtcPrice.Contract.ContractOwner(&_BtcPrice.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BtcPrice *BtcPriceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _BtcPrice.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BtcPrice *BtcPriceSession) Decimals() (uint8, error) {
	return _BtcPrice.Contract.Decimals(&_BtcPrice.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BtcPrice *BtcPriceCallerSession) Decimals() (uint8, error) {
	return _BtcPrice.Contract.Decimals(&_BtcPrice.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_BtcPrice *BtcPriceCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BtcPrice.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_BtcPrice *BtcPriceSession) Description() (string, error) {
	return _BtcPrice.Contract.Description(&_BtcPrice.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_BtcPrice *BtcPriceCallerSession) Description() (string, error) {
	return _BtcPrice.Contract.Description(&_BtcPrice.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 id) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_BtcPrice *BtcPriceCaller) GetRoundData(opts *bind.CallOpts, id *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _BtcPrice.contract.Call(opts, &out, "getRoundData", id)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 id) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_BtcPrice *BtcPriceSession) GetRoundData(id *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _BtcPrice.Contract.GetRoundData(&_BtcPrice.CallOpts, id)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 id) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_BtcPrice *BtcPriceCallerSession) GetRoundData(id *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _BtcPrice.Contract.GetRoundData(&_BtcPrice.CallOpts, id)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_BtcPrice *BtcPriceCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _BtcPrice.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_BtcPrice *BtcPriceSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _BtcPrice.Contract.LatestRoundData(&_BtcPrice.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_BtcPrice *BtcPriceCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _BtcPrice.Contract.LatestRoundData(&_BtcPrice.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_BtcPrice *BtcPriceCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BtcPrice.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_BtcPrice *BtcPriceSession) Version() (*big.Int, error) {
	return _BtcPrice.Contract.Version(&_BtcPrice.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_BtcPrice *BtcPriceCallerSession) Version() (*big.Int, error) {
	return _BtcPrice.Contract.Version(&_BtcPrice.CallOpts)
}

// SetBtcCoinPrice is a paid mutator transaction binding the contract method 0xf15fcd1e.
//
// Solidity: function setBtcCoinPrice(uint80 id, int256 price) returns()
func (_BtcPrice *BtcPriceTransactor) SetBtcCoinPrice(opts *bind.TransactOpts, id *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BtcPrice.contract.Transact(opts, "setBtcCoinPrice", id, price)
}

// SetBtcCoinPrice is a paid mutator transaction binding the contract method 0xf15fcd1e.
//
// Solidity: function setBtcCoinPrice(uint80 id, int256 price) returns()
func (_BtcPrice *BtcPriceSession) SetBtcCoinPrice(id *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BtcPrice.Contract.SetBtcCoinPrice(&_BtcPrice.TransactOpts, id, price)
}

// SetBtcCoinPrice is a paid mutator transaction binding the contract method 0xf15fcd1e.
//
// Solidity: function setBtcCoinPrice(uint80 id, int256 price) returns()
func (_BtcPrice *BtcPriceTransactorSession) SetBtcCoinPrice(id *big.Int, price *big.Int) (*types.Transaction, error) {
	return _BtcPrice.Contract.SetBtcCoinPrice(&_BtcPrice.TransactOpts, id, price)
}

// EthPriceMetaData contains all meta data concerning the EthPrice contract.
var EthPriceMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"contractOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"description\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ethCoinPrice\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"id\",\"type\":\"uint80\"}],\"name\":\"getRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"latestRoundData\",\"outputs\":[{\"internalType\":\"uint80\",\"name\":\"roundId\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"answer\",\"type\":\"int256\"},{\"internalType\":\"uint256\",\"name\":\"startedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"updatedAt\",\"type\":\"uint256\"},{\"internalType\":\"uint80\",\"name\":\"answeredInRound\",\"type\":\"uint80\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint80\",\"name\":\"id\",\"type\":\"uint80\"},{\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"}],\"name\":\"setEthCoinPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x0002000000000002000100000000000200010000000103550000006001100270000000600010019d0000008001000039000000400010043f0000000101200190000000320000c13d0000000001000031000000040110008c000001190000413d0000000101000367000000000101043b000000e001100270000000630210009c0000003f0000213d000000690210009c0000006f0000213d0000006c0210009c000000c10000613d0000006d0110009c000001190000c13d0000000001000416000000000101004b000001190000c13d000000040100008a00000000011000310000006e02000041000000000301004b000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001190000c13d0000000201000039000000000101041a000000400200043d00000000001204350000006001000041000000600320009c0000000001024019000000400110021000000070011001c70000017b0001042e0000000001000416000000000101004b000001190000c13d000000000100041a00000061011001970000000002000411000000000121019f000000000010041b00000020010000390000010000100443000001200000044300000062010000410000017b0001042e000000640210009c0000009c0000213d000000670210009c000000d50000613d000000680110009c000001190000c13d0000000001000416000000000101004b000001190000c13d000000040100008a00000000011000310000006e02000041000000200310008c000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001190000c13d00000004010000390000000101100367000000000101043b000000710210009c000001190000213d00000000001004350000000101000039000000200010043f0000000001000019017a01580000040f000000000201041a000000400100043d000100000001001d017a016a0000040f000000010400002900000000014100490000006002000041000000600310009c0000000001028019000000600340009c000000000204401900000040022002100000006001100210000000000121019f0000017b0001042e0000006a0210009c000000ef0000613d0000006b0110009c000001190000c13d0000000001000416000000000101004b000001190000c13d000000040100008a00000000011000310000006e02000041000000400310008c000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001190000c13d00000001010003670000000402100370000000000202043b000000710320009c000001190000213d000000000300041a0000006f033001970000000004000411000000000334004b000001460000c13d0000002401100370000000000101043b000100000001001d00000000002004350000000101000039000000200010043f0000000001000019017a01580000040f0000000102000029000000000021041b0000000201000039000000000021041b00000000010000190000017b0001042e000000650210009c000001080000613d000000660110009c000001190000c13d0000000001000416000000000101004b000001190000c13d000000040100008a00000000011000310000006e02000041000000000301004b000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001190000c13d0000000201000039000000000201041a000000400100043d000100000001001d017a016a0000040f000000010400002900000000014100490000006002000041000000600310009c0000000001028019000000600340009c000000000204401900000040022002100000006001100210000000000121019f0000017b0001042e0000000001000416000000000101004b000001190000c13d000000040100008a00000000011000310000006e02000041000000000301004b000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001190000c13d000000800000043f00000078010000410000017b0001042e0000000001000416000000000101004b000001190000c13d000000040100008a00000000011000310000006e02000041000000000301004b000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001190000c13d000000400300043d000000720130009c000001250000413d000000730100004100000000001004350000004101000039000000040010043f00000074010000410000017c000104300000000001000416000000000101004b000001190000c13d000000040100008a00000000011000310000006e02000041000000000301004b000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b000001190000c13d000000400100043d00000000000104350000006002000041000000600310009c0000000001028019000000400110021000000070011001c70000017b0001042e0000000001000416000000000101004b000001190000c13d000000040100008a00000000011000310000006e02000041000000000301004b000000000300001900000000030240190000006e01100197000000000401004b000000000200a0190000006e0110009c00000000010300190000000001026019000000000101004b0000011b0000613d00000000010000190000017c00010430000000000100041a0000006f01100197000000400200043d00000000001204350000006001000041000000600320009c0000000001024019000000400110021000000070011001c70000017b0001042e0000002002300039000000400020043f00000000000304350000002004000039000000400100043d0000000004410436000000000303043300000000003404350000004004100039000000000503004b000001380000613d000000000500001900000000064500190000000007250019000000000707043300000000007604350000002005500039000000000635004b000001310000413d000000000243001900000000000204350000005f02300039000000200300008a000000000232016f0000006003000041000000600420009c0000000002038019000000600410009c000000000103801900000040011002100000006002200210000000000112019f0000017b0001042e000000400100043d00000044021000390000007503000041000000000032043500000024021000390000001d030000390000000000320435000000760200004100000000002104350000000402100039000000200300003900000000003204350000006002000041000000600310009c0000000001028019000000400110021000000077011001c70000017c0001043000000060020000410000000003000414000000600430009c0000000003028019000000600410009c00000000010280190000004001100210000000c002300210000000000112019f00000079011001c70000801002000039017a01750000040f0000000102200190000001680000613d000000000101043b000000000001042d00000000010000190000017c00010430000000200310003900000000002304350000008002100039000000000002043500000060021000390000000000020435000000400210003900000000000204350000000000010435000000a001100039000000000001042d00000178002104230000000102000039000000000001042d0000000002000019000000000001042d0000017a000004320000017b0001042e0000017c0001043000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffff00000000000000000000000000000000000000000000000200000000000000000000000000000040000001000000000000000000000000000000000000000000000000000000000000000000000000007284e41500000000000000000000000000000000000000000000000000000000ce606edf00000000000000000000000000000000000000000000000000000000ce606ee000000000000000000000000000000000000000000000000000000000feaf968c000000000000000000000000000000000000000000000000000000007284e416000000000000000000000000000000000000000000000000000000009a6fc8f50000000000000000000000000000000000000000000000000000000054fd4d4f0000000000000000000000000000000000000000000000000000000054fd4d5000000000000000000000000000000000000000000000000000000000565514b700000000000000000000000000000000000000000000000000000000313ce567000000000000000000000000000000000000000000000000000000005387f92d8000000000000000000000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffffffffffffffffffffffff000000000000000000000000000000000000002000000000000000000000000000000000000000000000000000000000000000000000ffffffffffffffffffff000000000000000000000000000000000000000000000000ffffffffffffffe04e487b71000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000024000000000000000000000000796f752068617665206e6f206163717569726520746f206368616e676500000008c379a00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000640000000000000000000000000000000000000000000000000000000000000020000000800000000000000000020000000000000000000000000000000000004000000000000000000000000013eca16de2214c4a2aeeb943b9817f0c40d771c5193348969c0799099d40f159",
}

// EthPriceABI is the input ABI used to generate the binding from.
// Deprecated: Use EthPriceMetaData.ABI instead.
var EthPriceABI = EthPriceMetaData.ABI

// EthPriceBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EthPriceMetaData.Bin instead.
var EthPriceBin = EthPriceMetaData.Bin

// DeployEthPrice deploys a new Ethereum contract, binding an instance of EthPrice to it.
func DeployEthPrice(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *EthPrice, error) {
	parsed, err := EthPriceMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EthPriceBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EthPrice{EthPriceCaller: EthPriceCaller{contract: contract}, EthPriceTransactor: EthPriceTransactor{contract: contract}, EthPriceFilterer: EthPriceFilterer{contract: contract}}, nil
}

// EthPrice is an auto generated Go binding around an Ethereum contract.
type EthPrice struct {
	EthPriceCaller     // Read-only binding to the contract
	EthPriceTransactor // Write-only binding to the contract
	EthPriceFilterer   // Log filterer for contract events
}

// EthPriceCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthPriceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthPriceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthPriceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthPriceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthPriceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthPriceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthPriceSession struct {
	Contract     *EthPrice         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthPriceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthPriceCallerSession struct {
	Contract *EthPriceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// EthPriceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthPriceTransactorSession struct {
	Contract     *EthPriceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// EthPriceRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthPriceRaw struct {
	Contract *EthPrice // Generic contract binding to access the raw methods on
}

// EthPriceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthPriceCallerRaw struct {
	Contract *EthPriceCaller // Generic read-only contract binding to access the raw methods on
}

// EthPriceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthPriceTransactorRaw struct {
	Contract *EthPriceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEthPrice creates a new instance of EthPrice, bound to a specific deployed contract.
func NewEthPrice(address common.Address, backend bind.ContractBackend) (*EthPrice, error) {
	contract, err := bindEthPrice(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EthPrice{EthPriceCaller: EthPriceCaller{contract: contract}, EthPriceTransactor: EthPriceTransactor{contract: contract}, EthPriceFilterer: EthPriceFilterer{contract: contract}}, nil
}

// NewEthPriceCaller creates a new read-only instance of EthPrice, bound to a specific deployed contract.
func NewEthPriceCaller(address common.Address, caller bind.ContractCaller) (*EthPriceCaller, error) {
	contract, err := bindEthPrice(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthPriceCaller{contract: contract}, nil
}

// NewEthPriceTransactor creates a new write-only instance of EthPrice, bound to a specific deployed contract.
func NewEthPriceTransactor(address common.Address, transactor bind.ContractTransactor) (*EthPriceTransactor, error) {
	contract, err := bindEthPrice(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthPriceTransactor{contract: contract}, nil
}

// NewEthPriceFilterer creates a new log filterer instance of EthPrice, bound to a specific deployed contract.
func NewEthPriceFilterer(address common.Address, filterer bind.ContractFilterer) (*EthPriceFilterer, error) {
	contract, err := bindEthPrice(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthPriceFilterer{contract: contract}, nil
}

// bindEthPrice binds a generic wrapper to an already deployed contract.
func bindEthPrice(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := EthPriceMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthPrice *EthPriceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthPrice.Contract.EthPriceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthPrice *EthPriceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthPrice.Contract.EthPriceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthPrice *EthPriceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthPrice.Contract.EthPriceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EthPrice *EthPriceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EthPrice.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EthPrice *EthPriceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EthPrice.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EthPrice *EthPriceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EthPrice.Contract.contract.Transact(opts, method, params...)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_EthPrice *EthPriceCaller) ContractOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EthPrice.contract.Call(opts, &out, "contractOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_EthPrice *EthPriceSession) ContractOwner() (common.Address, error) {
	return _EthPrice.Contract.ContractOwner(&_EthPrice.CallOpts)
}

// ContractOwner is a free data retrieval call binding the contract method 0xce606ee0.
//
// Solidity: function contractOwner() view returns(address)
func (_EthPrice *EthPriceCallerSession) ContractOwner() (common.Address, error) {
	return _EthPrice.Contract.ContractOwner(&_EthPrice.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EthPrice *EthPriceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _EthPrice.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EthPrice *EthPriceSession) Decimals() (uint8, error) {
	return _EthPrice.Contract.Decimals(&_EthPrice.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EthPrice *EthPriceCallerSession) Decimals() (uint8, error) {
	return _EthPrice.Contract.Decimals(&_EthPrice.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_EthPrice *EthPriceCaller) Description(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _EthPrice.contract.Call(opts, &out, "description")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_EthPrice *EthPriceSession) Description() (string, error) {
	return _EthPrice.Contract.Description(&_EthPrice.CallOpts)
}

// Description is a free data retrieval call binding the contract method 0x7284e416.
//
// Solidity: function description() view returns(string)
func (_EthPrice *EthPriceCallerSession) Description() (string, error) {
	return _EthPrice.Contract.Description(&_EthPrice.CallOpts)
}

// EthCoinPrice is a free data retrieval call binding the contract method 0x5387f92d.
//
// Solidity: function ethCoinPrice() view returns(int256)
func (_EthPrice *EthPriceCaller) EthCoinPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthPrice.contract.Call(opts, &out, "ethCoinPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EthCoinPrice is a free data retrieval call binding the contract method 0x5387f92d.
//
// Solidity: function ethCoinPrice() view returns(int256)
func (_EthPrice *EthPriceSession) EthCoinPrice() (*big.Int, error) {
	return _EthPrice.Contract.EthCoinPrice(&_EthPrice.CallOpts)
}

// EthCoinPrice is a free data retrieval call binding the contract method 0x5387f92d.
//
// Solidity: function ethCoinPrice() view returns(int256)
func (_EthPrice *EthPriceCallerSession) EthCoinPrice() (*big.Int, error) {
	return _EthPrice.Contract.EthCoinPrice(&_EthPrice.CallOpts)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 id) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthPrice *EthPriceCaller) GetRoundData(opts *bind.CallOpts, id *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _EthPrice.contract.Call(opts, &out, "getRoundData", id)

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 id) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthPrice *EthPriceSession) GetRoundData(id *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _EthPrice.Contract.GetRoundData(&_EthPrice.CallOpts, id)
}

// GetRoundData is a free data retrieval call binding the contract method 0x9a6fc8f5.
//
// Solidity: function getRoundData(uint80 id) view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthPrice *EthPriceCallerSession) GetRoundData(id *big.Int) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _EthPrice.Contract.GetRoundData(&_EthPrice.CallOpts, id)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthPrice *EthPriceCaller) LatestRoundData(opts *bind.CallOpts) (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	var out []interface{}
	err := _EthPrice.contract.Call(opts, &out, "latestRoundData")

	outstruct := new(struct {
		RoundId         *big.Int
		Answer          *big.Int
		StartedAt       *big.Int
		UpdatedAt       *big.Int
		AnsweredInRound *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.RoundId = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Answer = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StartedAt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UpdatedAt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.AnsweredInRound = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthPrice *EthPriceSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _EthPrice.Contract.LatestRoundData(&_EthPrice.CallOpts)
}

// LatestRoundData is a free data retrieval call binding the contract method 0xfeaf968c.
//
// Solidity: function latestRoundData() view returns(uint80 roundId, int256 answer, uint256 startedAt, uint256 updatedAt, uint80 answeredInRound)
func (_EthPrice *EthPriceCallerSession) LatestRoundData() (struct {
	RoundId         *big.Int
	Answer          *big.Int
	StartedAt       *big.Int
	UpdatedAt       *big.Int
	AnsweredInRound *big.Int
}, error) {
	return _EthPrice.Contract.LatestRoundData(&_EthPrice.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_EthPrice *EthPriceCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EthPrice.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_EthPrice *EthPriceSession) Version() (*big.Int, error) {
	return _EthPrice.Contract.Version(&_EthPrice.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_EthPrice *EthPriceCallerSession) Version() (*big.Int, error) {
	return _EthPrice.Contract.Version(&_EthPrice.CallOpts)
}

// SetEthCoinPrice is a paid mutator transaction binding the contract method 0x565514b7.
//
// Solidity: function setEthCoinPrice(uint80 id, int256 price) returns()
func (_EthPrice *EthPriceTransactor) SetEthCoinPrice(opts *bind.TransactOpts, id *big.Int, price *big.Int) (*types.Transaction, error) {
	return _EthPrice.contract.Transact(opts, "setEthCoinPrice", id, price)
}

// SetEthCoinPrice is a paid mutator transaction binding the contract method 0x565514b7.
//
// Solidity: function setEthCoinPrice(uint80 id, int256 price) returns()
func (_EthPrice *EthPriceSession) SetEthCoinPrice(id *big.Int, price *big.Int) (*types.Transaction, error) {
	return _EthPrice.Contract.SetEthCoinPrice(&_EthPrice.TransactOpts, id, price)
}

// SetEthCoinPrice is a paid mutator transaction binding the contract method 0x565514b7.
//
// Solidity: function setEthCoinPrice(uint80 id, int256 price) returns()
func (_EthPrice *EthPriceTransactorSession) SetEthCoinPrice(id *big.Int, price *big.Int) (*types.Transaction, error) {
	return _EthPrice.Contract.SetEthCoinPrice(&_EthPrice.TransactOpts, id, price)
}
