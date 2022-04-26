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
	ABI: "[{\"inputs\":[],\"name\":\"lpFeePercent\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenLocks\",\"outputs\":[{\"internaltype\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internaltype\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internaltype\":\"uint256\",\"name\":\"unlockTime\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lockNonce\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"internaltype\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"lpLockAt\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internaltype\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"}],\"name\":\"lpLocksLength\",\"outputs\":[{\"internaltype\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_Mudra *MudraCaller) LockNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "lockNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_Mudra *MudraSession) LockNonce() (*big.Int, error) {
	return _Mudra.Contract.LockNonce(&_Mudra.CallOpts)
}

// LockNonce is a free data retrieval call binding the contract method 0xb5a9096e.
//
// Solidity: function lockNonce() view returns(uint256)
func (_Mudra *MudraCallerSession) LockNonce() (*big.Int, error) {
	return _Mudra.Contract.LockNonce(&_Mudra.CallOpts)
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

// LpLockAt is a free data retrieval call binding the contract method 0xe4af2d68.
//
// Solidity: function lpLockAt(address lpToken, uint256 index) view returns(uint256)
func (_Mudra *MudraCaller) LpLockAt(opts *bind.CallOpts, lpToken common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "lpLockAt", lpToken, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpLockAt is a free data retrieval call binding the contract method 0xe4af2d68.
//
// Solidity: function lpLockAt(address lpToken, uint256 index) view returns(uint256)
func (_Mudra *MudraSession) LpLockAt(lpToken common.Address, index *big.Int) (*big.Int, error) {
	return _Mudra.Contract.LpLockAt(&_Mudra.CallOpts, lpToken, index)
}

// LpLockAt is a free data retrieval call binding the contract method 0xe4af2d68.
//
// Solidity: function lpLockAt(address lpToken, uint256 index) view returns(uint256)
func (_Mudra *MudraCallerSession) LpLockAt(lpToken common.Address, index *big.Int) (*big.Int, error) {
	return _Mudra.Contract.LpLockAt(&_Mudra.CallOpts, lpToken, index)
}

// LpLocksLength is a free data retrieval call binding the contract method 0xb8c7813c.
//
// Solidity: function lpLocksLength(address lpToken) view returns(uint256)
func (_Mudra *MudraCaller) LpLocksLength(opts *bind.CallOpts, lpToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "lpLocksLength", lpToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpLocksLength is a free data retrieval call binding the contract method 0xb8c7813c.
//
// Solidity: function lpLocksLength(address lpToken) view returns(uint256)
func (_Mudra *MudraSession) LpLocksLength(lpToken common.Address) (*big.Int, error) {
	return _Mudra.Contract.LpLocksLength(&_Mudra.CallOpts, lpToken)
}

// LpLocksLength is a free data retrieval call binding the contract method 0xb8c7813c.
//
// Solidity: function lpLocksLength(address lpToken) view returns(uint256)
func (_Mudra *MudraCallerSession) LpLocksLength(lpToken common.Address) (*big.Int, error) {
	return _Mudra.Contract.LpLocksLength(&_Mudra.CallOpts, lpToken)
}

// TokenLocks is a free data retrieval call binding the contract method 0x946ca949.
//
// Solidity: function tokenLocks(uint256 ) view returns(address lpToken, address owner, uint256 tokenAmount, uint256 unlockTime)
func (_Mudra *MudraCaller) TokenLocks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	LpToken     common.Address
	Owner       common.Address
	TokenAmount *big.Int
	UnlockTime  *big.Int
}, error) {
	var out []interface{}
	err := _Mudra.contract.Call(opts, &out, "tokenLocks", arg0)

	outstruct := new(struct {
		LpToken     common.Address
		Owner       common.Address
		TokenAmount *big.Int
		UnlockTime  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LpToken = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Owner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TokenAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnlockTime = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// TokenLocks is a free data retrieval call binding the contract method 0x946ca949.
//
// Solidity: function tokenLocks(uint256 ) view returns(address lpToken, address owner, uint256 tokenAmount, uint256 unlockTime)
func (_Mudra *MudraSession) TokenLocks(arg0 *big.Int) (struct {
	LpToken     common.Address
	Owner       common.Address
	TokenAmount *big.Int
	UnlockTime  *big.Int
}, error) {
	return _Mudra.Contract.TokenLocks(&_Mudra.CallOpts, arg0)
}

// TokenLocks is a free data retrieval call binding the contract method 0x946ca949.
//
// Solidity: function tokenLocks(uint256 ) view returns(address lpToken, address owner, uint256 tokenAmount, uint256 unlockTime)
func (_Mudra *MudraCallerSession) TokenLocks(arg0 *big.Int) (struct {
	LpToken     common.Address
	Owner       common.Address
	TokenAmount *big.Int
	UnlockTime  *big.Int
}, error) {
	return _Mudra.Contract.TokenLocks(&_Mudra.CallOpts, arg0)
}
