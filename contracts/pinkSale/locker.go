// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pink

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

// PinkLockCumulativeLockInfo is an auto generated low-level Go binding around an user-defined struct.
type PinkLockCumulativeLockInfo struct {
	Token   common.Address
	Factory common.Address
	Amount  *big.Int
}

// PinkLockLock is an auto generated low-level Go binding around an user-defined struct.
type PinkLockLock struct {
	Id         *big.Int
	Token      common.Address
	Owner      common.Address
	Amount     *big.Int
	LockDate   *big.Int
	UnlockDate *big.Int
}

// PinkMetaData contains all meta data concerning the Pink contract.
var PinkMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"name\":\"LockAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockedAt\",\"type\":\"uint256\"}],\"name\":\"LockRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newUnlockDate\",\"type\":\"uint256\"}],\"name\":\"LockUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"allLocks\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"internalType\":\"structPinkLock.Lock[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allLpTokenLockedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allNormalTokenLockedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"cumulativeLockInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newUnlockDate\",\"type\":\"uint256\"}],\"name\":\"editLock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getCumulativeLpTokenLockInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structPinkLock.CumulativeLockInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getCumulativeLpTokenLockInfoAt\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structPinkLock.CumulativeLockInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getCumulativeNormalTokenLockInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structPinkLock.CumulativeLockInfo[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getCumulativeNormalTokenLockInfoAt\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structPinkLock.CumulativeLockInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getLock\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"internalType\":\"structPinkLock.Lock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"getLocksForToken\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"internalType\":\"structPinkLock.Lock[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalLockCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolManager\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_fee\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isLpToken\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"name\":\"lock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"lpLockCountForUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"lpLockForUserAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"internalType\":\"structPinkLock.Lock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"lpLocksForUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"internalType\":\"structPinkLock.Lock[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"normalLockCountForUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"normalLockForUserAtIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"internalType\":\"structPinkLock.Lock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"normalLocksForUser\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"internalType\":\"structPinkLock.Lock[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"poolManager\",\"outputs\":[{\"internalType\":\"contractIPoolManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newFee\",\"type\":\"uint256\"}],\"name\":\"setFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_poolManager\",\"type\":\"address\"}],\"name\":\"setPoolManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"totalLockCountForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"totalLockCountForUser\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalTokenLockedCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"lockId\",\"type\":\"uint256\"}],\"name\":\"unlock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"withdrawFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PinkABI is the input ABI used to generate the binding from.
// Deprecated: Use PinkMetaData.ABI instead.
var PinkABI = PinkMetaData.ABI

// Pink is an auto generated Go binding around an Ethereum contract.
type Pink struct {
	PinkCaller     // Read-only binding to the contract
	PinkTransactor // Write-only binding to the contract
	PinkFilterer   // Log filterer for contract events
}

// PinkCaller is an auto generated read-only Go binding around an Ethereum contract.
type PinkCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinkTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PinkTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinkFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PinkFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PinkSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PinkSession struct {
	Contract     *Pink             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PinkCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PinkCallerSession struct {
	Contract *PinkCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PinkTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PinkTransactorSession struct {
	Contract     *PinkTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PinkRaw is an auto generated low-level Go binding around an Ethereum contract.
type PinkRaw struct {
	Contract *Pink // Generic contract binding to access the raw methods on
}

// PinkCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PinkCallerRaw struct {
	Contract *PinkCaller // Generic read-only contract binding to access the raw methods on
}

// PinkTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PinkTransactorRaw struct {
	Contract *PinkTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPink creates a new instance of Pink, bound to a specific deployed contract.
func NewPink(address common.Address, backend bind.ContractBackend) (*Pink, error) {
	contract, err := bindPink(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pink{PinkCaller: PinkCaller{contract: contract}, PinkTransactor: PinkTransactor{contract: contract}, PinkFilterer: PinkFilterer{contract: contract}}, nil
}

// NewPinkCaller creates a new read-only instance of Pink, bound to a specific deployed contract.
func NewPinkCaller(address common.Address, caller bind.ContractCaller) (*PinkCaller, error) {
	contract, err := bindPink(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PinkCaller{contract: contract}, nil
}

// NewPinkTransactor creates a new write-only instance of Pink, bound to a specific deployed contract.
func NewPinkTransactor(address common.Address, transactor bind.ContractTransactor) (*PinkTransactor, error) {
	contract, err := bindPink(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PinkTransactor{contract: contract}, nil
}

// NewPinkFilterer creates a new log filterer instance of Pink, bound to a specific deployed contract.
func NewPinkFilterer(address common.Address, filterer bind.ContractFilterer) (*PinkFilterer, error) {
	contract, err := bindPink(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PinkFilterer{contract: contract}, nil
}

// bindPink binds a generic wrapper to an already deployed contract.
func bindPink(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PinkABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pink *PinkRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pink.Contract.PinkCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pink *PinkRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pink.Contract.PinkTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pink *PinkRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pink.Contract.PinkTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pink *PinkCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pink.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pink *PinkTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pink.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pink *PinkTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pink.Contract.contract.Transact(opts, method, params...)
}

// AllLocks is a free data retrieval call binding the contract method 0x5fbdf739.
//
// Solidity: function allLocks() view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_Pink *PinkCaller) AllLocks(opts *bind.CallOpts) ([]PinkLockLock, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "allLocks")

	if err != nil {
		return *new([]PinkLockLock), err
	}

	out0 := *abi.ConvertType(out[0], new([]PinkLockLock)).(*[]PinkLockLock)

	return out0, err

}

// AllLocks is a free data retrieval call binding the contract method 0x5fbdf739.
//
// Solidity: function allLocks() view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_Pink *PinkSession) AllLocks() ([]PinkLockLock, error) {
	return _Pink.Contract.AllLocks(&_Pink.CallOpts)
}

// AllLocks is a free data retrieval call binding the contract method 0x5fbdf739.
//
// Solidity: function allLocks() view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_Pink *PinkCallerSession) AllLocks() ([]PinkLockLock, error) {
	return _Pink.Contract.AllLocks(&_Pink.CallOpts)
}

// AllLpTokenLockedCount is a free data retrieval call binding the contract method 0xb982922e.
//
// Solidity: function allLpTokenLockedCount() view returns(uint256)
func (_Pink *PinkCaller) AllLpTokenLockedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "allLpTokenLockedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllLpTokenLockedCount is a free data retrieval call binding the contract method 0xb982922e.
//
// Solidity: function allLpTokenLockedCount() view returns(uint256)
func (_Pink *PinkSession) AllLpTokenLockedCount() (*big.Int, error) {
	return _Pink.Contract.AllLpTokenLockedCount(&_Pink.CallOpts)
}

// AllLpTokenLockedCount is a free data retrieval call binding the contract method 0xb982922e.
//
// Solidity: function allLpTokenLockedCount() view returns(uint256)
func (_Pink *PinkCallerSession) AllLpTokenLockedCount() (*big.Int, error) {
	return _Pink.Contract.AllLpTokenLockedCount(&_Pink.CallOpts)
}

// AllNormalTokenLockedCount is a free data retrieval call binding the contract method 0x475831c8.
//
// Solidity: function allNormalTokenLockedCount() view returns(uint256)
func (_Pink *PinkCaller) AllNormalTokenLockedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "allNormalTokenLockedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AllNormalTokenLockedCount is a free data retrieval call binding the contract method 0x475831c8.
//
// Solidity: function allNormalTokenLockedCount() view returns(uint256)
func (_Pink *PinkSession) AllNormalTokenLockedCount() (*big.Int, error) {
	return _Pink.Contract.AllNormalTokenLockedCount(&_Pink.CallOpts)
}

// AllNormalTokenLockedCount is a free data retrieval call binding the contract method 0x475831c8.
//
// Solidity: function allNormalTokenLockedCount() view returns(uint256)
func (_Pink *PinkCallerSession) AllNormalTokenLockedCount() (*big.Int, error) {
	return _Pink.Contract.AllNormalTokenLockedCount(&_Pink.CallOpts)
}

// CumulativeLockInfo is a free data retrieval call binding the contract method 0xe1444fd6.
//
// Solidity: function cumulativeLockInfo(address ) view returns(address token, address factory, uint256 amount)
func (_Pink *PinkCaller) CumulativeLockInfo(opts *bind.CallOpts, arg0 common.Address) (struct {
	Token   common.Address
	Factory common.Address
	Amount  *big.Int
}, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "cumulativeLockInfo", arg0)

	outstruct := new(struct {
		Token   common.Address
		Factory common.Address
		Amount  *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Factory = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CumulativeLockInfo is a free data retrieval call binding the contract method 0xe1444fd6.
//
// Solidity: function cumulativeLockInfo(address ) view returns(address token, address factory, uint256 amount)
func (_Pink *PinkSession) CumulativeLockInfo(arg0 common.Address) (struct {
	Token   common.Address
	Factory common.Address
	Amount  *big.Int
}, error) {
	return _Pink.Contract.CumulativeLockInfo(&_Pink.CallOpts, arg0)
}

// CumulativeLockInfo is a free data retrieval call binding the contract method 0xe1444fd6.
//
// Solidity: function cumulativeLockInfo(address ) view returns(address token, address factory, uint256 amount)
func (_Pink *PinkCallerSession) CumulativeLockInfo(arg0 common.Address) (struct {
	Token   common.Address
	Factory common.Address
	Amount  *big.Int
}, error) {
	return _Pink.Contract.CumulativeLockInfo(&_Pink.CallOpts, arg0)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Pink *PinkCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Pink *PinkSession) Fee() (*big.Int, error) {
	return _Pink.Contract.Fee(&_Pink.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Pink *PinkCallerSession) Fee() (*big.Int, error) {
	return _Pink.Contract.Fee(&_Pink.CallOpts)
}

// GetCumulativeLpTokenLockInfo is a free data retrieval call binding the contract method 0xaec640c6.
//
// Solidity: function getCumulativeLpTokenLockInfo(uint256 start, uint256 end) view returns((address,address,uint256)[])
func (_Pink *PinkCaller) GetCumulativeLpTokenLockInfo(opts *bind.CallOpts, start *big.Int, end *big.Int) ([]PinkLockCumulativeLockInfo, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "getCumulativeLpTokenLockInfo", start, end)

	if err != nil {
		return *new([]PinkLockCumulativeLockInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]PinkLockCumulativeLockInfo)).(*[]PinkLockCumulativeLockInfo)

	return out0, err

}

// GetCumulativeLpTokenLockInfo is a free data retrieval call binding the contract method 0xaec640c6.
//
// Solidity: function getCumulativeLpTokenLockInfo(uint256 start, uint256 end) view returns((address,address,uint256)[])
func (_Pink *PinkSession) GetCumulativeLpTokenLockInfo(start *big.Int, end *big.Int) ([]PinkLockCumulativeLockInfo, error) {
	return _Pink.Contract.GetCumulativeLpTokenLockInfo(&_Pink.CallOpts, start, end)
}

// GetCumulativeLpTokenLockInfo is a free data retrieval call binding the contract method 0xaec640c6.
//
// Solidity: function getCumulativeLpTokenLockInfo(uint256 start, uint256 end) view returns((address,address,uint256)[])
func (_Pink *PinkCallerSession) GetCumulativeLpTokenLockInfo(start *big.Int, end *big.Int) ([]PinkLockCumulativeLockInfo, error) {
	return _Pink.Contract.GetCumulativeLpTokenLockInfo(&_Pink.CallOpts, start, end)
}

// GetCumulativeLpTokenLockInfoAt is a free data retrieval call binding the contract method 0xa20b8c18.
//
// Solidity: function getCumulativeLpTokenLockInfoAt(uint256 index) view returns((address,address,uint256))
func (_Pink *PinkCaller) GetCumulativeLpTokenLockInfoAt(opts *bind.CallOpts, index *big.Int) (PinkLockCumulativeLockInfo, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "getCumulativeLpTokenLockInfoAt", index)

	if err != nil {
		return *new(PinkLockCumulativeLockInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PinkLockCumulativeLockInfo)).(*PinkLockCumulativeLockInfo)

	return out0, err

}

// GetCumulativeLpTokenLockInfoAt is a free data retrieval call binding the contract method 0xa20b8c18.
//
// Solidity: function getCumulativeLpTokenLockInfoAt(uint256 index) view returns((address,address,uint256))
func (_Pink *PinkSession) GetCumulativeLpTokenLockInfoAt(index *big.Int) (PinkLockCumulativeLockInfo, error) {
	return _Pink.Contract.GetCumulativeLpTokenLockInfoAt(&_Pink.CallOpts, index)
}

// GetCumulativeLpTokenLockInfoAt is a free data retrieval call binding the contract method 0xa20b8c18.
//
// Solidity: function getCumulativeLpTokenLockInfoAt(uint256 index) view returns((address,address,uint256))
func (_Pink *PinkCallerSession) GetCumulativeLpTokenLockInfoAt(index *big.Int) (PinkLockCumulativeLockInfo, error) {
	return _Pink.Contract.GetCumulativeLpTokenLockInfoAt(&_Pink.CallOpts, index)
}

// GetCumulativeNormalTokenLockInfo is a free data retrieval call binding the contract method 0x76c12822.
//
// Solidity: function getCumulativeNormalTokenLockInfo(uint256 start, uint256 end) view returns((address,address,uint256)[])
func (_Pink *PinkCaller) GetCumulativeNormalTokenLockInfo(opts *bind.CallOpts, start *big.Int, end *big.Int) ([]PinkLockCumulativeLockInfo, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "getCumulativeNormalTokenLockInfo", start, end)

	if err != nil {
		return *new([]PinkLockCumulativeLockInfo), err
	}

	out0 := *abi.ConvertType(out[0], new([]PinkLockCumulativeLockInfo)).(*[]PinkLockCumulativeLockInfo)

	return out0, err

}

// GetCumulativeNormalTokenLockInfo is a free data retrieval call binding the contract method 0x76c12822.
//
// Solidity: function getCumulativeNormalTokenLockInfo(uint256 start, uint256 end) view returns((address,address,uint256)[])
func (_Pink *PinkSession) GetCumulativeNormalTokenLockInfo(start *big.Int, end *big.Int) ([]PinkLockCumulativeLockInfo, error) {
	return _Pink.Contract.GetCumulativeNormalTokenLockInfo(&_Pink.CallOpts, start, end)
}

// GetCumulativeNormalTokenLockInfo is a free data retrieval call binding the contract method 0x76c12822.
//
// Solidity: function getCumulativeNormalTokenLockInfo(uint256 start, uint256 end) view returns((address,address,uint256)[])
func (_Pink *PinkCallerSession) GetCumulativeNormalTokenLockInfo(start *big.Int, end *big.Int) ([]PinkLockCumulativeLockInfo, error) {
	return _Pink.Contract.GetCumulativeNormalTokenLockInfo(&_Pink.CallOpts, start, end)
}

// GetCumulativeNormalTokenLockInfoAt is a free data retrieval call binding the contract method 0x7e6706d3.
//
// Solidity: function getCumulativeNormalTokenLockInfoAt(uint256 index) view returns((address,address,uint256))
func (_Pink *PinkCaller) GetCumulativeNormalTokenLockInfoAt(opts *bind.CallOpts, index *big.Int) (PinkLockCumulativeLockInfo, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "getCumulativeNormalTokenLockInfoAt", index)

	if err != nil {
		return *new(PinkLockCumulativeLockInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(PinkLockCumulativeLockInfo)).(*PinkLockCumulativeLockInfo)

	return out0, err

}

// GetCumulativeNormalTokenLockInfoAt is a free data retrieval call binding the contract method 0x7e6706d3.
//
// Solidity: function getCumulativeNormalTokenLockInfoAt(uint256 index) view returns((address,address,uint256))
func (_Pink *PinkSession) GetCumulativeNormalTokenLockInfoAt(index *big.Int) (PinkLockCumulativeLockInfo, error) {
	return _Pink.Contract.GetCumulativeNormalTokenLockInfoAt(&_Pink.CallOpts, index)
}

// GetCumulativeNormalTokenLockInfoAt is a free data retrieval call binding the contract method 0x7e6706d3.
//
// Solidity: function getCumulativeNormalTokenLockInfoAt(uint256 index) view returns((address,address,uint256))
func (_Pink *PinkCallerSession) GetCumulativeNormalTokenLockInfoAt(index *big.Int) (PinkLockCumulativeLockInfo, error) {
	return _Pink.Contract.GetCumulativeNormalTokenLockInfoAt(&_Pink.CallOpts, index)
}

// GetLock is a free data retrieval call binding the contract method 0xd68f4dd1.
//
// Solidity: function getLock(uint256 index) view returns((uint256,address,address,uint256,uint256,uint256))
func (_Pink *PinkCaller) GetLock(opts *bind.CallOpts, index *big.Int) (PinkLockLock, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "getLock", index)

	if err != nil {
		return *new(PinkLockLock), err
	}

	out0 := *abi.ConvertType(out[0], new(PinkLockLock)).(*PinkLockLock)

	return out0, err

}

// GetLock is a free data retrieval call binding the contract method 0xd68f4dd1.
//
// Solidity: function getLock(uint256 index) view returns((uint256,address,address,uint256,uint256,uint256))
func (_Pink *PinkSession) GetLock(index *big.Int) (PinkLockLock, error) {
	return _Pink.Contract.GetLock(&_Pink.CallOpts, index)
}

// GetLock is a free data retrieval call binding the contract method 0xd68f4dd1.
//
// Solidity: function getLock(uint256 index) view returns((uint256,address,address,uint256,uint256,uint256))
func (_Pink *PinkCallerSession) GetLock(index *big.Int) (PinkLockLock, error) {
	return _Pink.Contract.GetLock(&_Pink.CallOpts, index)
}

// GetLocksForToken is a free data retrieval call binding the contract method 0x332f26d7.
//
// Solidity: function getLocksForToken(address token, uint256 start, uint256 end) view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_Pink *PinkCaller) GetLocksForToken(opts *bind.CallOpts, token common.Address, start *big.Int, end *big.Int) ([]PinkLockLock, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "getLocksForToken", token, start, end)

	if err != nil {
		return *new([]PinkLockLock), err
	}

	out0 := *abi.ConvertType(out[0], new([]PinkLockLock)).(*[]PinkLockLock)

	return out0, err

}

// GetLocksForToken is a free data retrieval call binding the contract method 0x332f26d7.
//
// Solidity: function getLocksForToken(address token, uint256 start, uint256 end) view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_Pink *PinkSession) GetLocksForToken(token common.Address, start *big.Int, end *big.Int) ([]PinkLockLock, error) {
	return _Pink.Contract.GetLocksForToken(&_Pink.CallOpts, token, start, end)
}

// GetLocksForToken is a free data retrieval call binding the contract method 0x332f26d7.
//
// Solidity: function getLocksForToken(address token, uint256 start, uint256 end) view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_Pink *PinkCallerSession) GetLocksForToken(token common.Address, start *big.Int, end *big.Int) ([]PinkLockLock, error) {
	return _Pink.Contract.GetLocksForToken(&_Pink.CallOpts, token, start, end)
}

// GetTotalLockCount is a free data retrieval call binding the contract method 0xfd981c66.
//
// Solidity: function getTotalLockCount() view returns(uint256)
func (_Pink *PinkCaller) GetTotalLockCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "getTotalLockCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalLockCount is a free data retrieval call binding the contract method 0xfd981c66.
//
// Solidity: function getTotalLockCount() view returns(uint256)
func (_Pink *PinkSession) GetTotalLockCount() (*big.Int, error) {
	return _Pink.Contract.GetTotalLockCount(&_Pink.CallOpts)
}

// GetTotalLockCount is a free data retrieval call binding the contract method 0xfd981c66.
//
// Solidity: function getTotalLockCount() view returns(uint256)
func (_Pink *PinkCallerSession) GetTotalLockCount() (*big.Int, error) {
	return _Pink.Contract.GetTotalLockCount(&_Pink.CallOpts)
}

// LpLockCountForUser is a free data retrieval call binding the contract method 0x07873ef1.
//
// Solidity: function lpLockCountForUser(address user) view returns(uint256)
func (_Pink *PinkCaller) LpLockCountForUser(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "lpLockCountForUser", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LpLockCountForUser is a free data retrieval call binding the contract method 0x07873ef1.
//
// Solidity: function lpLockCountForUser(address user) view returns(uint256)
func (_Pink *PinkSession) LpLockCountForUser(user common.Address) (*big.Int, error) {
	return _Pink.Contract.LpLockCountForUser(&_Pink.CallOpts, user)
}

// LpLockCountForUser is a free data retrieval call binding the contract method 0x07873ef1.
//
// Solidity: function lpLockCountForUser(address user) view returns(uint256)
func (_Pink *PinkCallerSession) LpLockCountForUser(user common.Address) (*big.Int, error) {
	return _Pink.Contract.LpLockCountForUser(&_Pink.CallOpts, user)
}

// LpLockForUserAtIndex is a free data retrieval call binding the contract method 0xeeacf786.
//
// Solidity: function lpLockForUserAtIndex(address user, uint256 index) view returns((uint256,address,address,uint256,uint256,uint256))
func (_Pink *PinkCaller) LpLockForUserAtIndex(opts *bind.CallOpts, user common.Address, index *big.Int) (PinkLockLock, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "lpLockForUserAtIndex", user, index)

	if err != nil {
		return *new(PinkLockLock), err
	}

	out0 := *abi.ConvertType(out[0], new(PinkLockLock)).(*PinkLockLock)

	return out0, err

}

// LpLockForUserAtIndex is a free data retrieval call binding the contract method 0xeeacf786.
//
// Solidity: function lpLockForUserAtIndex(address user, uint256 index) view returns((uint256,address,address,uint256,uint256,uint256))
func (_Pink *PinkSession) LpLockForUserAtIndex(user common.Address, index *big.Int) (PinkLockLock, error) {
	return _Pink.Contract.LpLockForUserAtIndex(&_Pink.CallOpts, user, index)
}

// LpLockForUserAtIndex is a free data retrieval call binding the contract method 0xeeacf786.
//
// Solidity: function lpLockForUserAtIndex(address user, uint256 index) view returns((uint256,address,address,uint256,uint256,uint256))
func (_Pink *PinkCallerSession) LpLockForUserAtIndex(user common.Address, index *big.Int) (PinkLockLock, error) {
	return _Pink.Contract.LpLockForUserAtIndex(&_Pink.CallOpts, user, index)
}

// LpLocksForUser is a free data retrieval call binding the contract method 0xaef0e540.
//
// Solidity: function lpLocksForUser(address user) view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_Pink *PinkCaller) LpLocksForUser(opts *bind.CallOpts, user common.Address) ([]PinkLockLock, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "lpLocksForUser", user)

	if err != nil {
		return *new([]PinkLockLock), err
	}

	out0 := *abi.ConvertType(out[0], new([]PinkLockLock)).(*[]PinkLockLock)

	return out0, err

}

// LpLocksForUser is a free data retrieval call binding the contract method 0xaef0e540.
//
// Solidity: function lpLocksForUser(address user) view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_Pink *PinkSession) LpLocksForUser(user common.Address) ([]PinkLockLock, error) {
	return _Pink.Contract.LpLocksForUser(&_Pink.CallOpts, user)
}

// LpLocksForUser is a free data retrieval call binding the contract method 0xaef0e540.
//
// Solidity: function lpLocksForUser(address user) view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_Pink *PinkCallerSession) LpLocksForUser(user common.Address) ([]PinkLockLock, error) {
	return _Pink.Contract.LpLocksForUser(&_Pink.CallOpts, user)
}

// NormalLockCountForUser is a free data retrieval call binding the contract method 0xeb80bdae.
//
// Solidity: function normalLockCountForUser(address user) view returns(uint256)
func (_Pink *PinkCaller) NormalLockCountForUser(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "normalLockCountForUser", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NormalLockCountForUser is a free data retrieval call binding the contract method 0xeb80bdae.
//
// Solidity: function normalLockCountForUser(address user) view returns(uint256)
func (_Pink *PinkSession) NormalLockCountForUser(user common.Address) (*big.Int, error) {
	return _Pink.Contract.NormalLockCountForUser(&_Pink.CallOpts, user)
}

// NormalLockCountForUser is a free data retrieval call binding the contract method 0xeb80bdae.
//
// Solidity: function normalLockCountForUser(address user) view returns(uint256)
func (_Pink *PinkCallerSession) NormalLockCountForUser(user common.Address) (*big.Int, error) {
	return _Pink.Contract.NormalLockCountForUser(&_Pink.CallOpts, user)
}

// NormalLockForUserAtIndex is a free data retrieval call binding the contract method 0x618df7a3.
//
// Solidity: function normalLockForUserAtIndex(address user, uint256 index) view returns((uint256,address,address,uint256,uint256,uint256))
func (_Pink *PinkCaller) NormalLockForUserAtIndex(opts *bind.CallOpts, user common.Address, index *big.Int) (PinkLockLock, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "normalLockForUserAtIndex", user, index)

	if err != nil {
		return *new(PinkLockLock), err
	}

	out0 := *abi.ConvertType(out[0], new(PinkLockLock)).(*PinkLockLock)

	return out0, err

}

// NormalLockForUserAtIndex is a free data retrieval call binding the contract method 0x618df7a3.
//
// Solidity: function normalLockForUserAtIndex(address user, uint256 index) view returns((uint256,address,address,uint256,uint256,uint256))
func (_Pink *PinkSession) NormalLockForUserAtIndex(user common.Address, index *big.Int) (PinkLockLock, error) {
	return _Pink.Contract.NormalLockForUserAtIndex(&_Pink.CallOpts, user, index)
}

// NormalLockForUserAtIndex is a free data retrieval call binding the contract method 0x618df7a3.
//
// Solidity: function normalLockForUserAtIndex(address user, uint256 index) view returns((uint256,address,address,uint256,uint256,uint256))
func (_Pink *PinkCallerSession) NormalLockForUserAtIndex(user common.Address, index *big.Int) (PinkLockLock, error) {
	return _Pink.Contract.NormalLockForUserAtIndex(&_Pink.CallOpts, user, index)
}

// NormalLocksForUser is a free data retrieval call binding the contract method 0xda1d8cff.
//
// Solidity: function normalLocksForUser(address user) view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_Pink *PinkCaller) NormalLocksForUser(opts *bind.CallOpts, user common.Address) ([]PinkLockLock, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "normalLocksForUser", user)

	if err != nil {
		return *new([]PinkLockLock), err
	}

	out0 := *abi.ConvertType(out[0], new([]PinkLockLock)).(*[]PinkLockLock)

	return out0, err

}

// NormalLocksForUser is a free data retrieval call binding the contract method 0xda1d8cff.
//
// Solidity: function normalLocksForUser(address user) view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_Pink *PinkSession) NormalLocksForUser(user common.Address) ([]PinkLockLock, error) {
	return _Pink.Contract.NormalLocksForUser(&_Pink.CallOpts, user)
}

// NormalLocksForUser is a free data retrieval call binding the contract method 0xda1d8cff.
//
// Solidity: function normalLocksForUser(address user) view returns((uint256,address,address,uint256,uint256,uint256)[])
func (_Pink *PinkCallerSession) NormalLocksForUser(user common.Address) ([]PinkLockLock, error) {
	return _Pink.Contract.NormalLocksForUser(&_Pink.CallOpts, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pink *PinkCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pink *PinkSession) Owner() (common.Address, error) {
	return _Pink.Contract.Owner(&_Pink.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pink *PinkCallerSession) Owner() (common.Address, error) {
	return _Pink.Contract.Owner(&_Pink.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Pink *PinkCaller) PoolManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "poolManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Pink *PinkSession) PoolManager() (common.Address, error) {
	return _Pink.Contract.PoolManager(&_Pink.CallOpts)
}

// PoolManager is a free data retrieval call binding the contract method 0xdc4c90d3.
//
// Solidity: function poolManager() view returns(address)
func (_Pink *PinkCallerSession) PoolManager() (common.Address, error) {
	return _Pink.Contract.PoolManager(&_Pink.CallOpts)
}

// TotalLockCountForToken is a free data retrieval call binding the contract method 0xe3676f88.
//
// Solidity: function totalLockCountForToken(address token) view returns(uint256)
func (_Pink *PinkCaller) TotalLockCountForToken(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "totalLockCountForToken", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLockCountForToken is a free data retrieval call binding the contract method 0xe3676f88.
//
// Solidity: function totalLockCountForToken(address token) view returns(uint256)
func (_Pink *PinkSession) TotalLockCountForToken(token common.Address) (*big.Int, error) {
	return _Pink.Contract.TotalLockCountForToken(&_Pink.CallOpts, token)
}

// TotalLockCountForToken is a free data retrieval call binding the contract method 0xe3676f88.
//
// Solidity: function totalLockCountForToken(address token) view returns(uint256)
func (_Pink *PinkCallerSession) TotalLockCountForToken(token common.Address) (*big.Int, error) {
	return _Pink.Contract.TotalLockCountForToken(&_Pink.CallOpts, token)
}

// TotalLockCountForUser is a free data retrieval call binding the contract method 0xcd83eadc.
//
// Solidity: function totalLockCountForUser(address user) view returns(uint256)
func (_Pink *PinkCaller) TotalLockCountForUser(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "totalLockCountForUser", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalLockCountForUser is a free data retrieval call binding the contract method 0xcd83eadc.
//
// Solidity: function totalLockCountForUser(address user) view returns(uint256)
func (_Pink *PinkSession) TotalLockCountForUser(user common.Address) (*big.Int, error) {
	return _Pink.Contract.TotalLockCountForUser(&_Pink.CallOpts, user)
}

// TotalLockCountForUser is a free data retrieval call binding the contract method 0xcd83eadc.
//
// Solidity: function totalLockCountForUser(address user) view returns(uint256)
func (_Pink *PinkCallerSession) TotalLockCountForUser(user common.Address) (*big.Int, error) {
	return _Pink.Contract.TotalLockCountForUser(&_Pink.CallOpts, user)
}

// TotalTokenLockedCount is a free data retrieval call binding the contract method 0x1982242c.
//
// Solidity: function totalTokenLockedCount() view returns(uint256)
func (_Pink *PinkCaller) TotalTokenLockedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pink.contract.Call(opts, &out, "totalTokenLockedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalTokenLockedCount is a free data retrieval call binding the contract method 0x1982242c.
//
// Solidity: function totalTokenLockedCount() view returns(uint256)
func (_Pink *PinkSession) TotalTokenLockedCount() (*big.Int, error) {
	return _Pink.Contract.TotalTokenLockedCount(&_Pink.CallOpts)
}

// TotalTokenLockedCount is a free data retrieval call binding the contract method 0x1982242c.
//
// Solidity: function totalTokenLockedCount() view returns(uint256)
func (_Pink *PinkCallerSession) TotalTokenLockedCount() (*big.Int, error) {
	return _Pink.Contract.TotalTokenLockedCount(&_Pink.CallOpts)
}

// EditLock is a paid mutator transaction binding the contract method 0xb3b9aa48.
//
// Solidity: function editLock(uint256 lockId, uint256 newAmount, uint256 newUnlockDate) payable returns()
func (_Pink *PinkTransactor) EditLock(opts *bind.TransactOpts, lockId *big.Int, newAmount *big.Int, newUnlockDate *big.Int) (*types.Transaction, error) {
	return _Pink.contract.Transact(opts, "editLock", lockId, newAmount, newUnlockDate)
}

// EditLock is a paid mutator transaction binding the contract method 0xb3b9aa48.
//
// Solidity: function editLock(uint256 lockId, uint256 newAmount, uint256 newUnlockDate) payable returns()
func (_Pink *PinkSession) EditLock(lockId *big.Int, newAmount *big.Int, newUnlockDate *big.Int) (*types.Transaction, error) {
	return _Pink.Contract.EditLock(&_Pink.TransactOpts, lockId, newAmount, newUnlockDate)
}

// EditLock is a paid mutator transaction binding the contract method 0xb3b9aa48.
//
// Solidity: function editLock(uint256 lockId, uint256 newAmount, uint256 newUnlockDate) payable returns()
func (_Pink *PinkTransactorSession) EditLock(lockId *big.Int, newAmount *big.Int, newUnlockDate *big.Int) (*types.Transaction, error) {
	return _Pink.Contract.EditLock(&_Pink.TransactOpts, lockId, newAmount, newUnlockDate)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _poolManager, uint256 _fee) returns()
func (_Pink *PinkTransactor) Initialize(opts *bind.TransactOpts, _poolManager common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _Pink.contract.Transact(opts, "initialize", _poolManager, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _poolManager, uint256 _fee) returns()
func (_Pink *PinkSession) Initialize(_poolManager common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _Pink.Contract.Initialize(&_Pink.TransactOpts, _poolManager, _fee)
}

// Initialize is a paid mutator transaction binding the contract method 0xcd6dc687.
//
// Solidity: function initialize(address _poolManager, uint256 _fee) returns()
func (_Pink *PinkTransactorSession) Initialize(_poolManager common.Address, _fee *big.Int) (*types.Transaction, error) {
	return _Pink.Contract.Initialize(&_Pink.TransactOpts, _poolManager, _fee)
}

// Lock is a paid mutator transaction binding the contract method 0x64be5b39.
//
// Solidity: function lock(address owner, address token, bool isLpToken, uint256 amount, uint256 unlockDate) payable returns(uint256 id)
func (_Pink *PinkTransactor) Lock(opts *bind.TransactOpts, owner common.Address, token common.Address, isLpToken bool, amount *big.Int, unlockDate *big.Int) (*types.Transaction, error) {
	return _Pink.contract.Transact(opts, "lock", owner, token, isLpToken, amount, unlockDate)
}

// Lock is a paid mutator transaction binding the contract method 0x64be5b39.
//
// Solidity: function lock(address owner, address token, bool isLpToken, uint256 amount, uint256 unlockDate) payable returns(uint256 id)
func (_Pink *PinkSession) Lock(owner common.Address, token common.Address, isLpToken bool, amount *big.Int, unlockDate *big.Int) (*types.Transaction, error) {
	return _Pink.Contract.Lock(&_Pink.TransactOpts, owner, token, isLpToken, amount, unlockDate)
}

// Lock is a paid mutator transaction binding the contract method 0x64be5b39.
//
// Solidity: function lock(address owner, address token, bool isLpToken, uint256 amount, uint256 unlockDate) payable returns(uint256 id)
func (_Pink *PinkTransactorSession) Lock(owner common.Address, token common.Address, isLpToken bool, amount *big.Int, unlockDate *big.Int) (*types.Transaction, error) {
	return _Pink.Contract.Lock(&_Pink.TransactOpts, owner, token, isLpToken, amount, unlockDate)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pink *PinkTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pink.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pink *PinkSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pink.Contract.RenounceOwnership(&_Pink.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Pink *PinkTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Pink.Contract.RenounceOwnership(&_Pink.TransactOpts)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_Pink *PinkTransactor) SetFee(opts *bind.TransactOpts, newFee *big.Int) (*types.Transaction, error) {
	return _Pink.contract.Transact(opts, "setFee", newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_Pink *PinkSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _Pink.Contract.SetFee(&_Pink.TransactOpts, newFee)
}

// SetFee is a paid mutator transaction binding the contract method 0x69fe0e2d.
//
// Solidity: function setFee(uint256 newFee) returns()
func (_Pink *PinkTransactorSession) SetFee(newFee *big.Int) (*types.Transaction, error) {
	return _Pink.Contract.SetFee(&_Pink.TransactOpts, newFee)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolManager) returns()
func (_Pink *PinkTransactor) SetPoolManager(opts *bind.TransactOpts, _poolManager common.Address) (*types.Transaction, error) {
	return _Pink.contract.Transact(opts, "setPoolManager", _poolManager)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolManager) returns()
func (_Pink *PinkSession) SetPoolManager(_poolManager common.Address) (*types.Transaction, error) {
	return _Pink.Contract.SetPoolManager(&_Pink.TransactOpts, _poolManager)
}

// SetPoolManager is a paid mutator transaction binding the contract method 0x7aef6715.
//
// Solidity: function setPoolManager(address _poolManager) returns()
func (_Pink *PinkTransactorSession) SetPoolManager(_poolManager common.Address) (*types.Transaction, error) {
	return _Pink.Contract.SetPoolManager(&_Pink.TransactOpts, _poolManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pink *PinkTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pink.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pink *PinkSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pink.Contract.TransferOwnership(&_Pink.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pink *PinkTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pink.Contract.TransferOwnership(&_Pink.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 lockId) returns()
func (_Pink *PinkTransactor) Unlock(opts *bind.TransactOpts, lockId *big.Int) (*types.Transaction, error) {
	return _Pink.contract.Transact(opts, "unlock", lockId)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 lockId) returns()
func (_Pink *PinkSession) Unlock(lockId *big.Int) (*types.Transaction, error) {
	return _Pink.Contract.Unlock(&_Pink.TransactOpts, lockId)
}

// Unlock is a paid mutator transaction binding the contract method 0x6198e339.
//
// Solidity: function unlock(uint256 lockId) returns()
func (_Pink *PinkTransactorSession) Unlock(lockId *big.Int) (*types.Transaction, error) {
	return _Pink.Contract.Unlock(&_Pink.TransactOpts, lockId)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Pink *PinkTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Pink.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Pink *PinkSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Pink.Contract.UpgradeTo(&_Pink.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Pink *PinkTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Pink.Contract.UpgradeTo(&_Pink.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Pink *PinkTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Pink.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Pink *PinkSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Pink.Contract.UpgradeToAndCall(&_Pink.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Pink *PinkTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Pink.Contract.UpgradeToAndCall(&_Pink.TransactOpts, newImplementation, data)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_Pink *PinkTransactor) WithdrawFee(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pink.contract.Transact(opts, "withdrawFee")
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_Pink *PinkSession) WithdrawFee() (*types.Transaction, error) {
	return _Pink.Contract.WithdrawFee(&_Pink.TransactOpts)
}

// WithdrawFee is a paid mutator transaction binding the contract method 0xe941fa78.
//
// Solidity: function withdrawFee() returns()
func (_Pink *PinkTransactorSession) WithdrawFee() (*types.Transaction, error) {
	return _Pink.Contract.WithdrawFee(&_Pink.TransactOpts)
}

// PinkAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Pink contract.
type PinkAdminChangedIterator struct {
	Event *PinkAdminChanged // Event containing the contract specifics and raw log

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
func (it *PinkAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PinkAdminChanged)
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
		it.Event = new(PinkAdminChanged)
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
func (it *PinkAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PinkAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PinkAdminChanged represents a AdminChanged event raised by the Pink contract.
type PinkAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Pink *PinkFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*PinkAdminChangedIterator, error) {

	logs, sub, err := _Pink.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &PinkAdminChangedIterator{contract: _Pink.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Pink *PinkFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *PinkAdminChanged) (event.Subscription, error) {

	logs, sub, err := _Pink.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PinkAdminChanged)
				if err := _Pink.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Pink *PinkFilterer) ParseAdminChanged(log types.Log) (*PinkAdminChanged, error) {
	event := new(PinkAdminChanged)
	if err := _Pink.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PinkBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Pink contract.
type PinkBeaconUpgradedIterator struct {
	Event *PinkBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *PinkBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PinkBeaconUpgraded)
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
		it.Event = new(PinkBeaconUpgraded)
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
func (it *PinkBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PinkBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PinkBeaconUpgraded represents a BeaconUpgraded event raised by the Pink contract.
type PinkBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Pink *PinkFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*PinkBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Pink.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &PinkBeaconUpgradedIterator{contract: _Pink.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Pink *PinkFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *PinkBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Pink.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PinkBeaconUpgraded)
				if err := _Pink.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Pink *PinkFilterer) ParseBeaconUpgraded(log types.Log) (*PinkBeaconUpgraded, error) {
	event := new(PinkBeaconUpgraded)
	if err := _Pink.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PinkLockAddedIterator is returned from FilterLockAdded and is used to iterate over the raw logs and unpacked data for LockAdded events raised by the Pink contract.
type PinkLockAddedIterator struct {
	Event *PinkLockAdded // Event containing the contract specifics and raw log

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
func (it *PinkLockAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PinkLockAdded)
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
		it.Event = new(PinkLockAdded)
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
func (it *PinkLockAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PinkLockAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PinkLockAdded represents a LockAdded event raised by the Pink contract.
type PinkLockAdded struct {
	Id         *big.Int
	Token      common.Address
	Owner      common.Address
	Amount     *big.Int
	UnlockDate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLockAdded is a free log retrieval operation binding the contract event 0x694af1cc8727cdd0afbdd53d9b87b69248bd490224e9dd090e788546506e076f.
//
// Solidity: event LockAdded(uint256 indexed id, address token, address owner, uint256 amount, uint256 unlockDate)
func (_Pink *PinkFilterer) FilterLockAdded(opts *bind.FilterOpts, id []*big.Int) (*PinkLockAddedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Pink.contract.FilterLogs(opts, "LockAdded", idRule)
	if err != nil {
		return nil, err
	}
	return &PinkLockAddedIterator{contract: _Pink.contract, event: "LockAdded", logs: logs, sub: sub}, nil
}

// WatchLockAdded is a free log subscription operation binding the contract event 0x694af1cc8727cdd0afbdd53d9b87b69248bd490224e9dd090e788546506e076f.
//
// Solidity: event LockAdded(uint256 indexed id, address token, address owner, uint256 amount, uint256 unlockDate)
func (_Pink *PinkFilterer) WatchLockAdded(opts *bind.WatchOpts, sink chan<- *PinkLockAdded, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Pink.contract.WatchLogs(opts, "LockAdded", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PinkLockAdded)
				if err := _Pink.contract.UnpackLog(event, "LockAdded", log); err != nil {
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

// ParseLockAdded is a log parse operation binding the contract event 0x694af1cc8727cdd0afbdd53d9b87b69248bd490224e9dd090e788546506e076f.
//
// Solidity: event LockAdded(uint256 indexed id, address token, address owner, uint256 amount, uint256 unlockDate)
func (_Pink *PinkFilterer) ParseLockAdded(log types.Log) (*PinkLockAdded, error) {
	event := new(PinkLockAdded)
	if err := _Pink.contract.UnpackLog(event, "LockAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PinkLockRemovedIterator is returned from FilterLockRemoved and is used to iterate over the raw logs and unpacked data for LockRemoved events raised by the Pink contract.
type PinkLockRemovedIterator struct {
	Event *PinkLockRemoved // Event containing the contract specifics and raw log

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
func (it *PinkLockRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PinkLockRemoved)
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
		it.Event = new(PinkLockRemoved)
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
func (it *PinkLockRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PinkLockRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PinkLockRemoved represents a LockRemoved event raised by the Pink contract.
type PinkLockRemoved struct {
	Id         *big.Int
	Token      common.Address
	Owner      common.Address
	Amount     *big.Int
	UnlockedAt *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLockRemoved is a free log retrieval operation binding the contract event 0xc6532367992b32e42a62dd89fc3574876d97d081a263aa6e030f34b79b7e6e8b.
//
// Solidity: event LockRemoved(uint256 indexed id, address token, address owner, uint256 amount, uint256 unlockedAt)
func (_Pink *PinkFilterer) FilterLockRemoved(opts *bind.FilterOpts, id []*big.Int) (*PinkLockRemovedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Pink.contract.FilterLogs(opts, "LockRemoved", idRule)
	if err != nil {
		return nil, err
	}
	return &PinkLockRemovedIterator{contract: _Pink.contract, event: "LockRemoved", logs: logs, sub: sub}, nil
}

// WatchLockRemoved is a free log subscription operation binding the contract event 0xc6532367992b32e42a62dd89fc3574876d97d081a263aa6e030f34b79b7e6e8b.
//
// Solidity: event LockRemoved(uint256 indexed id, address token, address owner, uint256 amount, uint256 unlockedAt)
func (_Pink *PinkFilterer) WatchLockRemoved(opts *bind.WatchOpts, sink chan<- *PinkLockRemoved, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Pink.contract.WatchLogs(opts, "LockRemoved", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PinkLockRemoved)
				if err := _Pink.contract.UnpackLog(event, "LockRemoved", log); err != nil {
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

// ParseLockRemoved is a log parse operation binding the contract event 0xc6532367992b32e42a62dd89fc3574876d97d081a263aa6e030f34b79b7e6e8b.
//
// Solidity: event LockRemoved(uint256 indexed id, address token, address owner, uint256 amount, uint256 unlockedAt)
func (_Pink *PinkFilterer) ParseLockRemoved(log types.Log) (*PinkLockRemoved, error) {
	event := new(PinkLockRemoved)
	if err := _Pink.contract.UnpackLog(event, "LockRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PinkLockUpdatedIterator is returned from FilterLockUpdated and is used to iterate over the raw logs and unpacked data for LockUpdated events raised by the Pink contract.
type PinkLockUpdatedIterator struct {
	Event *PinkLockUpdated // Event containing the contract specifics and raw log

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
func (it *PinkLockUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PinkLockUpdated)
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
		it.Event = new(PinkLockUpdated)
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
func (it *PinkLockUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PinkLockUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PinkLockUpdated represents a LockUpdated event raised by the Pink contract.
type PinkLockUpdated struct {
	Id            *big.Int
	Token         common.Address
	Owner         common.Address
	NewAmount     *big.Int
	NewUnlockDate *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLockUpdated is a free log retrieval operation binding the contract event 0xa8b26360df8d5e154ffa5a8a7e894e85f781acfbbef0b744fb9551d8fd0fd36c.
//
// Solidity: event LockUpdated(uint256 indexed id, address token, address owner, uint256 newAmount, uint256 newUnlockDate)
func (_Pink *PinkFilterer) FilterLockUpdated(opts *bind.FilterOpts, id []*big.Int) (*PinkLockUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Pink.contract.FilterLogs(opts, "LockUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &PinkLockUpdatedIterator{contract: _Pink.contract, event: "LockUpdated", logs: logs, sub: sub}, nil
}

// WatchLockUpdated is a free log subscription operation binding the contract event 0xa8b26360df8d5e154ffa5a8a7e894e85f781acfbbef0b744fb9551d8fd0fd36c.
//
// Solidity: event LockUpdated(uint256 indexed id, address token, address owner, uint256 newAmount, uint256 newUnlockDate)
func (_Pink *PinkFilterer) WatchLockUpdated(opts *bind.WatchOpts, sink chan<- *PinkLockUpdated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Pink.contract.WatchLogs(opts, "LockUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PinkLockUpdated)
				if err := _Pink.contract.UnpackLog(event, "LockUpdated", log); err != nil {
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

// ParseLockUpdated is a log parse operation binding the contract event 0xa8b26360df8d5e154ffa5a8a7e894e85f781acfbbef0b744fb9551d8fd0fd36c.
//
// Solidity: event LockUpdated(uint256 indexed id, address token, address owner, uint256 newAmount, uint256 newUnlockDate)
func (_Pink *PinkFilterer) ParseLockUpdated(log types.Log) (*PinkLockUpdated, error) {
	event := new(PinkLockUpdated)
	if err := _Pink.contract.UnpackLog(event, "LockUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PinkOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Pink contract.
type PinkOwnershipTransferredIterator struct {
	Event *PinkOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PinkOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PinkOwnershipTransferred)
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
		it.Event = new(PinkOwnershipTransferred)
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
func (it *PinkOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PinkOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PinkOwnershipTransferred represents a OwnershipTransferred event raised by the Pink contract.
type PinkOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pink *PinkFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PinkOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pink.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PinkOwnershipTransferredIterator{contract: _Pink.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Pink *PinkFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PinkOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Pink.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PinkOwnershipTransferred)
				if err := _Pink.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Pink *PinkFilterer) ParseOwnershipTransferred(log types.Log) (*PinkOwnershipTransferred, error) {
	event := new(PinkOwnershipTransferred)
	if err := _Pink.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PinkUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Pink contract.
type PinkUpgradedIterator struct {
	Event *PinkUpgraded // Event containing the contract specifics and raw log

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
func (it *PinkUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PinkUpgraded)
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
		it.Event = new(PinkUpgraded)
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
func (it *PinkUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PinkUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PinkUpgraded represents a Upgraded event raised by the Pink contract.
type PinkUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Pink *PinkFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*PinkUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Pink.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &PinkUpgradedIterator{contract: _Pink.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Pink *PinkFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *PinkUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Pink.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PinkUpgraded)
				if err := _Pink.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Pink *PinkFilterer) ParseUpgraded(log types.Log) (*PinkUpgraded, error) {
	event := new(PinkUpgraded)
	if err := _Pink.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
