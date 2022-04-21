// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package mudra

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
)

// MudraMetaData contains all meta data concerning the Mudra contract.
var MudraMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internaltype\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internaltype\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internaltype\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internaltype\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"depositId\",\"type\":\"uint256\"}],\"name\":\"TokensLocked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internaltype\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"indexed\":true,\"internaltype\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internaltype\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TokensWithdrawn\",\"type\":\"event\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allDepositIds\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bnbFee\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositId\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositsByTokenAddress\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositsByWithdrawalAddress\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllDepositIds\",\"outputs\":[{\"internaltype\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getDepositDetails\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internaltype\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getDepositsByTokenAddress\",\"outputs\":[{\"internaltype\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"_withdrawalAddress\",\"type\":\"address\"}],\"name\":\"getDepositsByWithdrawalAddress\",\"outputs\":[{\"internaltype\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internaltype\":\"address\",\"name\":\"_walletAddress\",\"type\":\"address\"}],\"name\":\"getTokenBalanceByAddress\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"}],\"name\":\"getTotalTokenBalance\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"_unlockTime\",\"type\":\"uint256\"},{\"internaltype\":\"bool\",\"name\":\"_feeInBnb\",\"type\":\"bool\"}],\"name\":\"lockTokens\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lockedToken\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"},{\"internaltype\":\"address\",\"name\":\"withdrawalAddress\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"},{\"internaltype\":\"bool\",\"name\":\"withdrawn\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lpFeePercent\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingBnbFees\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"setBnbFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"percent\",\"type\":\"uint256\"}],\"name\":\"setLpFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokensFees\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBnbFees\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internaltype\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"walletTokenBalance\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"addresspayable\",\"name\":\"withdrawalAddress\",\"type\":\"address\"}],\"name\":\"withdrawFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// MudraABI is the input ABI used to generate the binding from.
// Deprecated: Use MudraMetaData.ABI instead.
var MudraABI = MudraMetaData.ABI

// Mudra is an auto generated Go binding around an Ethereum contract.
type Mudra struct {
	MudraCaller     // Read-only binding to the contract
	MudraTransactor // Write-only binding to the contract
	MudraFilterer   // Log filterer for contract events
}

// MudraCaller is an auto generated read-only Go binding around an Ethereum contract.
type MudraCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MudraTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MudraTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MudraFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MudraFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MudraSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MudraSession struct {
	Contract     *Mudra            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MudraCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MudraCallerSession struct {
	Contract *MudraCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MudraTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MudraTransactorSession struct {
	Contract     *MudraTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MudraRaw is an auto generated low-level Go binding around an Ethereum contract.
type MudraRaw struct {
	Contract *Mudra // Generic contract binding to access the raw methods on
}

// MudraCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MudraCallerRaw struct {
	Contract *MudraCaller // Generic read-only contract binding to access the raw methods on
}

// MudraTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MudraTransactorRaw struct {
	Contract *MudraTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMudra creates a new instance of Mudra, bound to a specific deployed contract.
func NewMudra(address common.Address, backend bind.ContractBackend) (*Mudra, error) {
	contract, err := bindMudra(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Mudra{MudraCaller: MudraCaller{contract: contract}, MudraTransactor: MudraTransactor{contract: contract}, MudraFilterer: MudraFilterer{contract: contract}}, nil
}

// NewMudraCaller creates a new read-only instance of Mudra, bound to a specific deployed contract.
func NewMudraCaller(address common.Address, caller bind.ContractCaller) (*MudraCaller, error) {
	contract, err := bindMudra(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MudraCaller{contract: contract}, nil
}

// NewMudraTransactor creates a new write-only instance of Mudra, bound to a specific deployed contract.
func NewMudraTransactor(address common.Address, transactor bind.ContractTransactor) (*MudraTransactor, error) {
	contract, err := bindMudra(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MudraTransactor{contract: contract}, nil
}

// NewMudraFilterer creates a new log filterer instance of Mudra, bound to a specific deployed contract.
func NewMudraFilterer(address common.Address, filterer bind.ContractFilterer) (*MudraFilterer, error) {
	contract, err := bindMudra(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MudraFilterer{contract: contract}, nil
}

// bindMudra binds a generic wrapper to an already deployed contract.
func bindMudra(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MudraABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mudra *MudraRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mudra.Contract.MudraCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mudra *MudraRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mudra.Contract.MudraTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mudra *MudraRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mudra.Contract.MudraTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Mudra *MudraCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Mudra.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Mudra *MudraTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mudra.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Mudra *MudraTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Mudra.Contract.contract.Transact(opts, method, params...)
}

// AllDepositIds is a free data retrieval call binding the contract method 0xc9028aff.
//
// Solidity: function allDepositIds(uint256 ) view returns(uint256)
func (_Mudra *MudraCaller) AllDepositIds(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "allDepositIds", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllDepositIds is a free data retrieval call binding the contract method 0xc9028aff.
//
// Solidity: function allDepositIds(uint256 ) view returns(uint256)
func (_Mudra *MudraSession) AllDepositIds(arg0 *big.Int) (*big.Int, error) {
	return _Mudra.Contract.AllDepositIds(&_Mudra.CallOpts, arg0)
}

// AllDepositIds is a free data retrieval call binding the contract method 0xc9028aff.
//
// Solidity: function allDepositIds(uint256 ) view returns(uint256)
func (_Mudra *MudraCallerSession) AllDepositIds(arg0 *big.Int) (*big.Int, error) {
	return _Mudra.Contract.AllDepositIds(&_Mudra.CallOpts, arg0)
}

// BnbFee is a free data retrieval call binding the contract method 0xefc17415.
//
// Solidity: function bnbFee() view returns(uint256)
func (_Mudra *MudraCaller) BnbFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "bnbFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BnbFee is a free data retrieval call binding the contract method 0xefc17415.
//
// Solidity: function bnbFee() view returns(uint256)
func (_Mudra *MudraSession) BnbFee() (*big.Int, error) {
	return _Mudra.Contract.BnbFee(&_Mudra.CallOpts)
}

// BnbFee is a free data retrieval call binding the contract method 0xefc17415.
//
// Solidity: function bnbFee() view returns(uint256)
func (_Mudra *MudraCallerSession) BnbFee() (*big.Int, error) {
	return _Mudra.Contract.BnbFee(&_Mudra.CallOpts)
}

// DepositId is a free data retrieval call binding the contract method 0x9852099c.
//
// Solidity: function depositId() view returns(uint256)
func (_Mudra *MudraCaller) DepositId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "depositId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositId is a free data retrieval call binding the contract method 0x9852099c.
//
// Solidity: function depositId() view returns(uint256)
func (_Mudra *MudraSession) DepositId() (*big.Int, error) {
	return _Mudra.Contract.DepositId(&_Mudra.CallOpts)
}

// DepositId is a free data retrieval call binding the contract method 0x9852099c.
//
// Solidity: function depositId() view returns(uint256)
func (_Mudra *MudraCallerSession) DepositId() (*big.Int, error) {
	return _Mudra.Contract.DepositId(&_Mudra.CallOpts)
}

// DepositsByTokenAddress is a free data retrieval call binding the contract method 0x3fd97b9b.
//
// Solidity: function depositsByTokenAddress(address , uint256 ) view returns(uint256)
func (_Mudra *MudraCaller) DepositsByTokenAddress(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "depositsByTokenAddress", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositsByTokenAddress is a free data retrieval call binding the contract method 0x3fd97b9b.
//
// Solidity: function depositsByTokenAddress(address , uint256 ) view returns(uint256)
func (_Mudra *MudraSession) DepositsByTokenAddress(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Mudra.Contract.DepositsByTokenAddress(&_Mudra.CallOpts, arg0, arg1)
}

// DepositsByTokenAddress is a free data retrieval call binding the contract method 0x3fd97b9b.
//
// Solidity: function depositsByTokenAddress(address , uint256 ) view returns(uint256)
func (_Mudra *MudraCallerSession) DepositsByTokenAddress(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Mudra.Contract.DepositsByTokenAddress(&_Mudra.CallOpts, arg0, arg1)
}

// DepositsByWithdrawalAddress is a free data retrieval call binding the contract method 0x530680d8.
//
// Solidity: function depositsByWithdrawalAddress(address , uint256 ) view returns(uint256)
func (_Mudra *MudraCaller) DepositsByWithdrawalAddress(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "depositsByWithdrawalAddress", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositsByWithdrawalAddress is a free data retrieval call binding the contract method 0x530680d8.
//
// Solidity: function depositsByWithdrawalAddress(address , uint256 ) view returns(uint256)
func (_Mudra *MudraSession) DepositsByWithdrawalAddress(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Mudra.Contract.DepositsByWithdrawalAddress(&_Mudra.CallOpts, arg0, arg1)
}

// DepositsByWithdrawalAddress is a free data retrieval call binding the contract method 0x530680d8.
//
// Solidity: function depositsByWithdrawalAddress(address , uint256 ) view returns(uint256)
func (_Mudra *MudraCallerSession) DepositsByWithdrawalAddress(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _Mudra.Contract.DepositsByWithdrawalAddress(&_Mudra.CallOpts, arg0, arg1)
}

// GetAllDepositIds is a free data retrieval call binding the contract method 0x6ba03924.
//
// Solidity: function getAllDepositIds() view returns(uint256[])
func (_Mudra *MudraCaller) GetAllDepositIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "getAllDepositIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllDepositIds is a free data retrieval call binding the contract method 0x6ba03924.
//
// Solidity: function getAllDepositIds() view returns(uint256[])
func (_Mudra *MudraSession) GetAllDepositIds() ([]*big.Int, error) {
	return _Mudra.Contract.GetAllDepositIds(&_Mudra.CallOpts)
}

// GetAllDepositIds is a free data retrieval call binding the contract method 0x6ba03924.
//
// Solidity: function getAllDepositIds() view returns(uint256[])
func (_Mudra *MudraCallerSession) GetAllDepositIds() ([]*big.Int, error) {
	return _Mudra.Contract.GetAllDepositIds(&_Mudra.CallOpts)
}

// GetDepositDetails is a free data retrieval call binding the contract method 0x890db72f.
//
// Solidity: function getDepositDetails(uint256 _id) view returns(address, address, uint256, uint256, bool)
func (_Mudra *MudraCaller) GetDepositDetails(opts *bind.CallOpts, _id *big.Int) (common.Address, common.Address, *big.Int, *big.Int, bool, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "getDepositDetails", _id)

	if err != nil {
		return *new(common.Address), *new(common.Address), *new(*big.Int), *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, err

}

// GetDepositDetails is a free data retrieval call binding the contract method 0x890db72f.
//
// Solidity: function getDepositDetails(uint256 _id) view returns(address, address, uint256, uint256, bool)
func (_Mudra *MudraSession) GetDepositDetails(_id *big.Int) (common.Address, common.Address, *big.Int, *big.Int, bool, error) {
	return _Mudra.Contract.GetDepositDetails(&_Mudra.CallOpts, _id)
}

// GetDepositDetails is a free data retrieval call binding the contract method 0x890db72f.
//
// Solidity: function getDepositDetails(uint256 _id) view returns(address, address, uint256, uint256, bool)
func (_Mudra *MudraCallerSession) GetDepositDetails(_id *big.Int) (common.Address, common.Address, *big.Int, *big.Int, bool, error) {
	return _Mudra.Contract.GetDepositDetails(&_Mudra.CallOpts, _id)
}

// GetDepositsByTokenAddress is a free data retrieval call binding the contract method 0x86f65a22.
//
// Solidity: function getDepositsByTokenAddress(address _tokenAddress) view returns(uint256[])
func (_Mudra *MudraCaller) GetDepositsByTokenAddress(opts *bind.CallOpts, _tokenAddress common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "getDepositsByTokenAddress", _tokenAddress)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDepositsByTokenAddress is a free data retrieval call binding the contract method 0x86f65a22.
//
// Solidity: function getDepositsByTokenAddress(address _tokenAddress) view returns(uint256[])
func (_Mudra *MudraSession) GetDepositsByTokenAddress(_tokenAddress common.Address) ([]*big.Int, error) {
	return _Mudra.Contract.GetDepositsByTokenAddress(&_Mudra.CallOpts, _tokenAddress)
}

// GetDepositsByTokenAddress is a free data retrieval call binding the contract method 0x86f65a22.
//
// Solidity: function getDepositsByTokenAddress(address _tokenAddress) view returns(uint256[])
func (_Mudra *MudraCallerSession) GetDepositsByTokenAddress(_tokenAddress common.Address) ([]*big.Int, error) {
	return _Mudra.Contract.GetDepositsByTokenAddress(&_Mudra.CallOpts, _tokenAddress)
}

// GetDepositsByWithdrawalAddress is a free data retrieval call binding the contract method 0x0bd59ad3.
//
// Solidity: function getDepositsByWithdrawalAddress(address _withdrawalAddress) view returns(uint256[])
func (_Mudra *MudraCaller) GetDepositsByWithdrawalAddress(opts *bind.CallOpts, _withdrawalAddress common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "getDepositsByWithdrawalAddress", _withdrawalAddress)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetDepositsByWithdrawalAddress is a free data retrieval call binding the contract method 0x0bd59ad3.
//
// Solidity: function getDepositsByWithdrawalAddress(address _withdrawalAddress) view returns(uint256[])
func (_Mudra *MudraSession) GetDepositsByWithdrawalAddress(_withdrawalAddress common.Address) ([]*big.Int, error) {
	return _Mudra.Contract.GetDepositsByWithdrawalAddress(&_Mudra.CallOpts, _withdrawalAddress)
}

// GetDepositsByWithdrawalAddress is a free data retrieval call binding the contract method 0x0bd59ad3.
//
// Solidity: function getDepositsByWithdrawalAddress(address _withdrawalAddress) view returns(uint256[])
func (_Mudra *MudraCallerSession) GetDepositsByWithdrawalAddress(_withdrawalAddress common.Address) ([]*big.Int, error) {
	return _Mudra.Contract.GetDepositsByWithdrawalAddress(&_Mudra.CallOpts, _withdrawalAddress)
}

// GetTokenBalanceByAddress is a free data retrieval call binding the contract method 0x347c80ba.
//
// Solidity: function getTokenBalanceByAddress(address _tokenAddress, address _walletAddress) view returns(uint256)
func (_Mudra *MudraCaller) GetTokenBalanceByAddress(opts *bind.CallOpts, _tokenAddress common.Address, _walletAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "getTokenBalanceByAddress", _tokenAddress, _walletAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenBalanceByAddress is a free data retrieval call binding the contract method 0x347c80ba.
//
// Solidity: function getTokenBalanceByAddress(address _tokenAddress, address _walletAddress) view returns(uint256)
func (_Mudra *MudraSession) GetTokenBalanceByAddress(_tokenAddress common.Address, _walletAddress common.Address) (*big.Int, error) {
	return _Mudra.Contract.GetTokenBalanceByAddress(&_Mudra.CallOpts, _tokenAddress, _walletAddress)
}

// GetTokenBalanceByAddress is a free data retrieval call binding the contract method 0x347c80ba.
//
// Solidity: function getTokenBalanceByAddress(address _tokenAddress, address _walletAddress) view returns(uint256)
func (_Mudra *MudraCallerSession) GetTokenBalanceByAddress(_tokenAddress common.Address, _walletAddress common.Address) (*big.Int, error) {
	return _Mudra.Contract.GetTokenBalanceByAddress(&_Mudra.CallOpts, _tokenAddress, _walletAddress)
}

// GetTotalTokenBalance is a free data retrieval call binding the contract method 0xadad19bd.
//
// Solidity: function getTotalTokenBalance(address _tokenAddress) view returns(uint256)
func (_Mudra *MudraCaller) GetTotalTokenBalance(opts *bind.CallOpts, _tokenAddress common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "getTotalTokenBalance", _tokenAddress)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalTokenBalance is a free data retrieval call binding the contract method 0xadad19bd.
//
// Solidity: function getTotalTokenBalance(address _tokenAddress) view returns(uint256)
func (_Mudra *MudraSession) GetTotalTokenBalance(_tokenAddress common.Address) (*big.Int, error) {
	return _Mudra.Contract.GetTotalTokenBalance(&_Mudra.CallOpts, _tokenAddress)
}

// GetTotalTokenBalance is a free data retrieval call binding the contract method 0xadad19bd.
//
// Solidity: function getTotalTokenBalance(address _tokenAddress) view returns(uint256)
func (_Mudra *MudraCallerSession) GetTotalTokenBalance(_tokenAddress common.Address) (*big.Int, error) {
	return _Mudra.Contract.GetTotalTokenBalance(&_Mudra.CallOpts, _tokenAddress)
}

// LockedToken is a free data retrieval call binding the contract method 0xbb941cff.
//
// Solidity: function lockedToken(uint256 ) view returns(address tokenAddress, address withdrawalAddress, uint256 tokenAmount, uint256 unlockTime, bool withdrawn)
func (_Mudra *MudraCaller) LockedToken(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenAddress      common.Address
	WithdrawalAddress common.Address
	TokenAmount       *big.Int
	UnlockTime        *big.Int
	Withdrawn         bool
}, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "lockedToken", arg0)

	outstruct := new(struct {
		TokenAddress      common.Address
		WithdrawalAddress common.Address
		TokenAmount       *big.Int
		UnlockTime        *big.Int
		Withdrawn         bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.WithdrawalAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnlockTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Withdrawn = *abi.ConvertType(out[4], new(bool)).(*bool)

	return *outstruct, err

}

// LockedToken is a free data retrieval call binding the contract method 0xbb941cff.
//
// Solidity: function lockedToken(uint256 ) view returns(address tokenAddress, address withdrawalAddress, uint256 tokenAmount, uint256 unlockTime, bool withdrawn)
func (_Mudra *MudraSession) LockedToken(arg0 *big.Int) (struct {
	TokenAddress      common.Address
	WithdrawalAddress common.Address
	TokenAmount       *big.Int
	UnlockTime        *big.Int
	Withdrawn         bool
}, error) {
	return _Mudra.Contract.LockedToken(&_Mudra.CallOpts, arg0)
}

// LockedToken is a free data retrieval call binding the contract method 0xbb941cff.
//
// Solidity: function lockedToken(uint256 ) view returns(address tokenAddress, address withdrawalAddress, uint256 tokenAmount, uint256 unlockTime, bool withdrawn)
func (_Mudra *MudraCallerSession) LockedToken(arg0 *big.Int) (struct {
	TokenAddress      common.Address
	WithdrawalAddress common.Address
	TokenAmount       *big.Int
	UnlockTime        *big.Int
	Withdrawn         bool
}, error) {
	return _Mudra.Contract.LockedToken(&_Mudra.CallOpts, arg0)
}

// LpFeePercent is a free data retrieval call binding the contract method 0xd4dbe3ef.
//
// Solidity: function lpFeePercent() view returns(uint256)
func (_Mudra *MudraCaller) LpFeePercent(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "lpFeePercent")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpFeePercent is a free data retrieval call binding the contract method 0xd4dbe3ef.
//
// Solidity: function lpFeePercent() view returns(uint256)
func (_Mudra *MudraSession) LpFeePercent() (*big.Int, error) {
	return _Mudra.Contract.LpFeePercent(&_Mudra.CallOpts)
}

// LpFeePercent is a free data retrieval call binding the contract method 0xd4dbe3ef.
//
// Solidity: function lpFeePercent() view returns(uint256)
func (_Mudra *MudraCallerSession) LpFeePercent() (*big.Int, error) {
	return _Mudra.Contract.LpFeePercent(&_Mudra.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mudra *MudraCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mudra *MudraSession) Owner() (common.Address, error) {
	return _Mudra.Contract.Owner(&_Mudra.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Mudra *MudraCallerSession) Owner() (common.Address, error) {
	return _Mudra.Contract.Owner(&_Mudra.CallOpts)
}

// RemainingBnbFees is a free data retrieval call binding the contract method 0x213596b8.
//
// Solidity: function remainingBnbFees() view returns(uint256)
func (_Mudra *MudraCaller) RemainingBnbFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "remainingBnbFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingBnbFees is a free data retrieval call binding the contract method 0x213596b8.
//
// Solidity: function remainingBnbFees() view returns(uint256)
func (_Mudra *MudraSession) RemainingBnbFees() (*big.Int, error) {
	return _Mudra.Contract.RemainingBnbFees(&_Mudra.CallOpts)
}

// RemainingBnbFees is a free data retrieval call binding the contract method 0x213596b8.
//
// Solidity: function remainingBnbFees() view returns(uint256)
func (_Mudra *MudraCallerSession) RemainingBnbFees() (*big.Int, error) {
	return _Mudra.Contract.RemainingBnbFees(&_Mudra.CallOpts)
}

// TokensFees is a free data retrieval call binding the contract method 0x0cff87f1.
//
// Solidity: function tokensFees(address ) view returns(uint256)
func (_Mudra *MudraCaller) TokensFees(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "tokensFees", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensFees is a free data retrieval call binding the contract method 0x0cff87f1.
//
// Solidity: function tokensFees(address ) view returns(uint256)
func (_Mudra *MudraSession) TokensFees(arg0 common.Address) (*big.Int, error) {
	return _Mudra.Contract.TokensFees(&_Mudra.CallOpts, arg0)
}

// TokensFees is a free data retrieval call binding the contract method 0x0cff87f1.
//
// Solidity: function tokensFees(address ) view returns(uint256)
func (_Mudra *MudraCallerSession) TokensFees(arg0 common.Address) (*big.Int, error) {
	return _Mudra.Contract.TokensFees(&_Mudra.CallOpts, arg0)
}

// TotalBnbFees is a free data retrieval call binding the contract method 0xae3ff3dd.
//
// Solidity: function totalBnbFees() view returns(uint256)
func (_Mudra *MudraCaller) TotalBnbFees(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "totalBnbFees")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBnbFees is a free data retrieval call binding the contract method 0xae3ff3dd.
//
// Solidity: function totalBnbFees() view returns(uint256)
func (_Mudra *MudraSession) TotalBnbFees() (*big.Int, error) {
	return _Mudra.Contract.TotalBnbFees(&_Mudra.CallOpts)
}

// TotalBnbFees is a free data retrieval call binding the contract method 0xae3ff3dd.
//
// Solidity: function totalBnbFees() view returns(uint256)
func (_Mudra *MudraCallerSession) TotalBnbFees() (*big.Int, error) {
	return _Mudra.Contract.TotalBnbFees(&_Mudra.CallOpts)
}

// WalletTokenBalance is a free data retrieval call binding the contract method 0xb9e7df1c.
//
// Solidity: function walletTokenBalance(address , address ) view returns(uint256)
func (_Mudra *MudraCaller) WalletTokenBalance(opts *bind.CallOpts, arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "walletTokenBalance", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WalletTokenBalance is a free data retrieval call binding the contract method 0xb9e7df1c.
//
// Solidity: function walletTokenBalance(address , address ) view returns(uint256)
func (_Mudra *MudraSession) WalletTokenBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Mudra.Contract.WalletTokenBalance(&_Mudra.CallOpts, arg0, arg1)
}

// WalletTokenBalance is a free data retrieval call binding the contract method 0xb9e7df1c.
//
// Solidity: function walletTokenBalance(address , address ) view returns(uint256)
func (_Mudra *MudraCallerSession) WalletTokenBalance(arg0 common.Address, arg1 common.Address) (*big.Int, error) {
	return _Mudra.Contract.WalletTokenBalance(&_Mudra.CallOpts, arg0, arg1)
}

// LockTokens is a paid mutator transaction binding the contract method 0xe6a478b3.
//
// Solidity: function lockTokens(address _tokenAddress, uint256 _amount, uint256 _unlockTime, bool _feeInBnb) payable returns(uint256 _id)
func (_Mudra *MudraTransactor) LockTokens(opts *bind.TransactOpts, _tokenAddress common.Address, _amount *big.Int, _unlockTime *big.Int, _feeInBnb bool) (*types.Transaction, error) {
	return _Mudra.contract.Transact(opts, "lockTokens", _tokenAddress, _amount, _unlockTime, _feeInBnb)
}

// LockTokens is a paid mutator transaction binding the contract method 0xe6a478b3.
//
// Solidity: function lockTokens(address _tokenAddress, uint256 _amount, uint256 _unlockTime, bool _feeInBnb) payable returns(uint256 _id)
func (_Mudra *MudraSession) LockTokens(_tokenAddress common.Address, _amount *big.Int, _unlockTime *big.Int, _feeInBnb bool) (*types.Transaction, error) {
	return _Mudra.Contract.LockTokens(&_Mudra.TransactOpts, _tokenAddress, _amount, _unlockTime, _feeInBnb)
}

// LockTokens is a paid mutator transaction binding the contract method 0xe6a478b3.
//
// Solidity: function lockTokens(address _tokenAddress, uint256 _amount, uint256 _unlockTime, bool _feeInBnb) payable returns(uint256 _id)
func (_Mudra *MudraTransactorSession) LockTokens(_tokenAddress common.Address, _amount *big.Int, _unlockTime *big.Int, _feeInBnb bool) (*types.Transaction, error) {
	return _Mudra.Contract.LockTokens(&_Mudra.TransactOpts, _tokenAddress, _amount, _unlockTime, _feeInBnb)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mudra *MudraTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Mudra.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mudra *MudraSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mudra.Contract.RenounceOwnership(&_Mudra.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Mudra *MudraTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Mudra.Contract.RenounceOwnership(&_Mudra.TransactOpts)
}

// SetBnbFee is a paid mutator transaction binding the contract method 0x8a47446c.
//
// Solidity: function setBnbFee(uint256 fee) returns()
func (_Mudra *MudraTransactor) SetBnbFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _Mudra.contract.Transact(opts, "setBnbFee", fee)
}

// SetBnbFee is a paid mutator transaction binding the contract method 0x8a47446c.
//
// Solidity: function setBnbFee(uint256 fee) returns()
func (_Mudra *MudraSession) SetBnbFee(fee *big.Int) (*types.Transaction, error) {
	return _Mudra.Contract.SetBnbFee(&_Mudra.TransactOpts, fee)
}

// SetBnbFee is a paid mutator transaction binding the contract method 0x8a47446c.
//
// Solidity: function setBnbFee(uint256 fee) returns()
func (_Mudra *MudraTransactorSession) SetBnbFee(fee *big.Int) (*types.Transaction, error) {
	return _Mudra.Contract.SetBnbFee(&_Mudra.TransactOpts, fee)
}

// SetLpFee is a paid mutator transaction binding the contract method 0x84dfbfe2.
//
// Solidity: function setLpFee(uint256 percent) returns()
func (_Mudra *MudraTransactor) SetLpFee(opts *bind.TransactOpts, percent *big.Int) (*types.Transaction, error) {
	return _Mudra.contract.Transact(opts, "setLpFee", percent)
}

// SetLpFee is a paid mutator transaction binding the contract method 0x84dfbfe2.
//
// Solidity: function setLpFee(uint256 percent) returns()
func (_Mudra *MudraSession) SetLpFee(percent *big.Int) (*types.Transaction, error) {
	return _Mudra.Contract.SetLpFee(&_Mudra.TransactOpts, percent)
}

// SetLpFee is a paid mutator transaction binding the contract method 0x84dfbfe2.
//
// Solidity: function setLpFee(uint256 percent) returns()
func (_Mudra *MudraTransactorSession) SetLpFee(percent *big.Int) (*types.Transaction, error) {
	return _Mudra.Contract.SetLpFee(&_Mudra.TransactOpts, percent)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mudra *MudraTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Mudra.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mudra *MudraSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mudra.Contract.TransferOwnership(&_Mudra.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Mudra *MudraTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Mudra.Contract.TransferOwnership(&_Mudra.TransactOpts, newOwner)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address withdrawalAddress) returns()
func (_Mudra *MudraTransactor) WithdrawFees(opts *bind.TransactOpts, withdrawalAddress common.Address) (*types.Transaction, error) {
	return _Mudra.contract.Transact(opts, "withdrawFees", withdrawalAddress)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address withdrawalAddress) returns()
func (_Mudra *MudraSession) WithdrawFees(withdrawalAddress common.Address) (*types.Transaction, error) {
	return _Mudra.Contract.WithdrawFees(&_Mudra.TransactOpts, withdrawalAddress)
}

// WithdrawFees is a paid mutator transaction binding the contract method 0x164e68de.
//
// Solidity: function withdrawFees(address withdrawalAddress) returns()
func (_Mudra *MudraTransactorSession) WithdrawFees(withdrawalAddress common.Address) (*types.Transaction, error) {
	return _Mudra.Contract.WithdrawFees(&_Mudra.TransactOpts, withdrawalAddress)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x315a095d.
//
// Solidity: function withdrawTokens(uint256 _id) returns()
func (_Mudra *MudraTransactor) WithdrawTokens(opts *bind.TransactOpts, _id *big.Int) (*types.Transaction, error) {
	return _Mudra.contract.Transact(opts, "withdrawTokens", _id)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x315a095d.
//
// Solidity: function withdrawTokens(uint256 _id) returns()
func (_Mudra *MudraSession) WithdrawTokens(_id *big.Int) (*types.Transaction, error) {
	return _Mudra.Contract.WithdrawTokens(&_Mudra.TransactOpts, _id)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x315a095d.
//
// Solidity: function withdrawTokens(uint256 _id) returns()
func (_Mudra *MudraTransactorSession) WithdrawTokens(_id *big.Int) (*types.Transaction, error) {
	return _Mudra.Contract.WithdrawTokens(&_Mudra.TransactOpts, _id)
}

// MudraOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Mudra contract.
type MudraOwnershipTransferredIterator struct {
	Event *MudraOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MudraOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MudraOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MudraOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MudraOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MudraOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MudraOwnershipTransferred represents a OwnershipTransferred event raised by the Mudra contract.
type MudraOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mudra *MudraFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*MudraOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mudra.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &MudraOwnershipTransferredIterator{contract: _Mudra.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mudra *MudraFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *MudraOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Mudra.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MudraOwnershipTransferred)
				if err := _Mudra.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Mudra *MudraFilterer) ParseOwnershipTransferred(log types.Log) (*MudraOwnershipTransferred, error) {
	event := new(MudraOwnershipTransferred)
	if err := _Mudra.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MudraTokensLockedIterator is returned from FilterTokensLocked and is used to iterate over the raw logs and unpacked data for TokensLocked events raised by the Mudra contract.
type MudraTokensLockedIterator struct {
	Event *MudraTokensLocked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MudraTokensLockedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MudraTokensLocked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MudraTokensLocked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MudraTokensLockedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MudraTokensLockedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MudraTokensLocked represents a TokensLocked event raised by the Mudra contract.
type MudraTokensLocked struct {
	TokenAddress common.Address
	Sender       common.Address
	Amount       *big.Int
	UnlockTime   *big.Int
	DepositId    *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokensLocked is a free log retrieval operation binding the contract event 0xa6782d3322abbfbe850e6d5c5c78e8e1df603ea07608bb9a62dd83f40d4feccc.
//
// Solidity: event TokensLocked(address indexed tokenAddress, address indexed sender, uint256 amount, uint256 unlockTime, uint256 depositId)
func (_Mudra *MudraFilterer) FilterTokensLocked(opts *bind.FilterOpts, tokenAddress []common.Address, sender []common.Address) (*MudraTokensLockedIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Mudra.contract.FilterLogs(opts, "TokensLocked", tokenAddressRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &MudraTokensLockedIterator{contract: _Mudra.contract, event: "TokensLocked", logs: logs, sub: sub}, nil
}

// WatchTokensLocked is a free log subscription operation binding the contract event 0xa6782d3322abbfbe850e6d5c5c78e8e1df603ea07608bb9a62dd83f40d4feccc.
//
// Solidity: event TokensLocked(address indexed tokenAddress, address indexed sender, uint256 amount, uint256 unlockTime, uint256 depositId)
func (_Mudra *MudraFilterer) WatchTokensLocked(opts *bind.WatchOpts, sink chan<- *MudraTokensLocked, tokenAddress []common.Address, sender []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Mudra.contract.WatchLogs(opts, "TokensLocked", tokenAddressRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MudraTokensLocked)
				if err := _Mudra.contract.UnpackLog(event, "TokensLocked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensLocked is a log parse operation binding the contract event 0xa6782d3322abbfbe850e6d5c5c78e8e1df603ea07608bb9a62dd83f40d4feccc.
//
// Solidity: event TokensLocked(address indexed tokenAddress, address indexed sender, uint256 amount, uint256 unlockTime, uint256 depositId)
func (_Mudra *MudraFilterer) ParseTokensLocked(log types.Log) (*MudraTokensLocked, error) {
	event := new(MudraTokensLocked)
	if err := _Mudra.contract.UnpackLog(event, "TokensLocked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MudraTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the Mudra contract.
type MudraTokensWithdrawnIterator struct {
	Event *MudraTokensWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *MudraTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MudraTokensWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(MudraTokensWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *MudraTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MudraTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MudraTokensWithdrawn represents a TokensWithdrawn event raised by the Mudra contract.
type MudraTokensWithdrawn struct {
	TokenAddress common.Address
	Receiver     common.Address
	Amount       *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6337ed398c0e8467698c581374fdce4db14922df487b5a39483079f5f59b60a4.
//
// Solidity: event TokensWithdrawn(address indexed tokenAddress, address indexed receiver, uint256 amount)
func (_Mudra *MudraFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, tokenAddress []common.Address, receiver []common.Address) (*MudraTokensWithdrawnIterator, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Mudra.contract.FilterLogs(opts, "TokensWithdrawn", tokenAddressRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &MudraTokensWithdrawnIterator{contract: _Mudra.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6337ed398c0e8467698c581374fdce4db14922df487b5a39483079f5f59b60a4.
//
// Solidity: event TokensWithdrawn(address indexed tokenAddress, address indexed receiver, uint256 amount)
func (_Mudra *MudraFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *MudraTokensWithdrawn, tokenAddress []common.Address, receiver []common.Address) (event.Subscription, error) {

	var tokenAddressRule []interface{}
	for _, tokenAddressItem := range tokenAddress {
		tokenAddressRule = append(tokenAddressRule, tokenAddressItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _Mudra.contract.WatchLogs(opts, "TokensWithdrawn", tokenAddressRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MudraTokensWithdrawn)
				if err := _Mudra.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6337ed398c0e8467698c581374fdce4db14922df487b5a39483079f5f59b60a4.
//
// Solidity: event TokensWithdrawn(address indexed tokenAddress, address indexed receiver, uint256 amount)
func (_Mudra *MudraFilterer) ParseTokensWithdrawn(log types.Log) (*MudraTokensWithdrawn, error) {
	event := new(MudraTokensWithdrawn)
	if err := _Mudra.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
