// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package unicrypt

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

// UnicryptMetaData contains all meta data concerning the Unicrypt contract.
var UnicryptMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIUniFactory\",\"name\":\"_uniswapFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"lockDate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"}],\"name\":\"onDeposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"lpToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"onWithdraw\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"MIGRATION_IN\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ethFee\",\"type\":\"uint256\"},{\"internalType\":\"contractIERCBurn\",\"name\":\"secondaryFeeToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"secondaryTokenFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondaryTokenDiscount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referralPercent\",\"type\":\"uint256\"},{\"internalType\":\"contractIERCBurn\",\"name\":\"referralToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"referralHold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"referralDiscount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getLockedTokenAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumLockedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"}],\"name\":\"getNumLocksForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUserLockForTokenAtIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getUserLockedTokenAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserNumLockedTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"}],\"name\":\"getUserNumLocksForToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"getUserWhitelistStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"getWhitelistedUserAtIndex\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelistedUsersLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"incrementLock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlock_date\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_referral\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_fee_in_eth\",\"type\":\"bool\"},{\"internalType\":\"addresspayable\",\"name\":\"_withdrawer\",\"type\":\"address\"}],\"name\":\"lockLPToken\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"migrate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_unlock_date\",\"type\":\"uint256\"}],\"name\":\"relock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"_devaddr\",\"type\":\"address\"}],\"name\":\"setDev\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_referralPercent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_referralDiscount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_ethFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondaryTokenFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_secondaryTokenDiscount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidityFee\",\"type\":\"uint256\"}],\"name\":\"setFees\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_migrationIn\",\"type\":\"address\"}],\"name\":\"setMigrationIn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIMigrator\",\"name\":\"_migrator\",\"type\":\"address\"}],\"name\":\"setMigrator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERCBurn\",\"name\":\"_referralToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_hold\",\"type\":\"uint256\"}],\"name\":\"setReferralTokenAndHold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_secondaryFeeToken\",\"type\":\"address\"}],\"name\":\"setSecondaryFeeToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"splitLock\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokenLocks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lockDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"unlockDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lockID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"},{\"internalType\":\"addresspayable\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferLockOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniswapFactory\",\"outputs\":[{\"internalType\":\"contractIUniFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"_add\",\"type\":\"bool\"}],\"name\":\"whitelistFeeAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lpToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_lockID\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// UnicryptABI is the input ABI used to generate the binding from.
// Deprecated: Use UnicryptMetaData.ABI instead.
var UnicryptABI = UnicryptMetaData.ABI

// Unicrypt is an auto generated Go binding around an Ethereum contract.
type Unicrypt struct {
	UnicryptCaller     // Read-only binding to the contract
	UnicryptTransactor // Write-only binding to the contract
	UnicryptFilterer   // Log filterer for contract events
}

// UnicryptCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnicryptCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnicryptTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnicryptTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnicryptFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnicryptFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnicryptSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnicryptSession struct {
	Contract     *Unicrypt         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnicryptCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnicryptCallerSession struct {
	Contract *UnicryptCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// UnicryptTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnicryptTransactorSession struct {
	Contract     *UnicryptTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// UnicryptRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnicryptRaw struct {
	Contract *Unicrypt // Generic contract binding to access the raw methods on
}

// UnicryptCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnicryptCallerRaw struct {
	Contract *UnicryptCaller // Generic read-only contract binding to access the raw methods on
}

// UnicryptTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnicryptTransactorRaw struct {
	Contract *UnicryptTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnicrypt creates a new instance of Unicrypt, bound to a specific deployed contract.
func NewUnicrypt(address common.Address, backend bind.ContractBackend) (*Unicrypt, error) {
	contract, err := bindUnicrypt(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Unicrypt{UnicryptCaller: UnicryptCaller{contract: contract}, UnicryptTransactor: UnicryptTransactor{contract: contract}, UnicryptFilterer: UnicryptFilterer{contract: contract}}, nil
}

// NewUnicryptCaller creates a new read-only instance of Unicrypt, bound to a specific deployed contract.
func NewUnicryptCaller(address common.Address, caller bind.ContractCaller) (*UnicryptCaller, error) {
	contract, err := bindUnicrypt(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnicryptCaller{contract: contract}, nil
}

// NewUnicryptTransactor creates a new write-only instance of Unicrypt, bound to a specific deployed contract.
func NewUnicryptTransactor(address common.Address, transactor bind.ContractTransactor) (*UnicryptTransactor, error) {
	contract, err := bindUnicrypt(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnicryptTransactor{contract: contract}, nil
}

// NewUnicryptFilterer creates a new log filterer instance of Unicrypt, bound to a specific deployed contract.
func NewUnicryptFilterer(address common.Address, filterer bind.ContractFilterer) (*UnicryptFilterer, error) {
	contract, err := bindUnicrypt(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnicryptFilterer{contract: contract}, nil
}

// bindUnicrypt binds a generic wrapper to an already deployed contract.
func bindUnicrypt(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnicryptABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unicrypt *UnicryptRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unicrypt.Contract.UnicryptCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unicrypt *UnicryptRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unicrypt.Contract.UnicryptTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unicrypt *UnicryptRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unicrypt.Contract.UnicryptTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Unicrypt *UnicryptCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Unicrypt.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Unicrypt *UnicryptTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unicrypt.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Unicrypt *UnicryptTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Unicrypt.Contract.contract.Transact(opts, method, params...)
}

// MIGRATIONIN is a free data retrieval call binding the contract method 0x76323bd8.
//
// Solidity: function MIGRATION_IN() view returns(address)
func (_Unicrypt *UnicryptCaller) MIGRATIONIN(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "MIGRATION_IN")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MIGRATIONIN is a free data retrieval call binding the contract method 0x76323bd8.
//
// Solidity: function MIGRATION_IN() view returns(address)
func (_Unicrypt *UnicryptSession) MIGRATIONIN() (common.Address, error) {
	return _Unicrypt.Contract.MIGRATIONIN(&_Unicrypt.CallOpts)
}

// MIGRATIONIN is a free data retrieval call binding the contract method 0x76323bd8.
//
// Solidity: function MIGRATION_IN() view returns(address)
func (_Unicrypt *UnicryptCallerSession) MIGRATIONIN() (common.Address, error) {
	return _Unicrypt.Contract.MIGRATIONIN(&_Unicrypt.CallOpts)
}

// GFees is a free data retrieval call binding the contract method 0x90e1a003.
//
// Solidity: function gFees() view returns(uint256 ethFee, address secondaryFeeToken, uint256 secondaryTokenFee, uint256 secondaryTokenDiscount, uint256 liquidityFee, uint256 referralPercent, address referralToken, uint256 referralHold, uint256 referralDiscount)
func (_Unicrypt *UnicryptCaller) GFees(opts *bind.CallOpts) (struct {
	EthFee                 *big.Int
	SecondaryFeeToken      common.Address
	SecondaryTokenFee      *big.Int
	SecondaryTokenDiscount *big.Int
	LiquidityFee           *big.Int
	ReferralPercent        *big.Int
	ReferralToken          common.Address
	ReferralHold           *big.Int
	ReferralDiscount       *big.Int
}, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "gFees")

	outstruct := new(struct {
		EthFee                 *big.Int
		SecondaryFeeToken      common.Address
		SecondaryTokenFee      *big.Int
		SecondaryTokenDiscount *big.Int
		LiquidityFee           *big.Int
		ReferralPercent        *big.Int
		ReferralToken          common.Address
		ReferralHold           *big.Int
		ReferralDiscount       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.EthFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.SecondaryFeeToken = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.SecondaryTokenFee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.SecondaryTokenDiscount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LiquidityFee = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ReferralPercent = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.ReferralToken = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.ReferralHold = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.ReferralDiscount = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GFees is a free data retrieval call binding the contract method 0x90e1a003.
//
// Solidity: function gFees() view returns(uint256 ethFee, address secondaryFeeToken, uint256 secondaryTokenFee, uint256 secondaryTokenDiscount, uint256 liquidityFee, uint256 referralPercent, address referralToken, uint256 referralHold, uint256 referralDiscount)
func (_Unicrypt *UnicryptSession) GFees() (struct {
	EthFee                 *big.Int
	SecondaryFeeToken      common.Address
	SecondaryTokenFee      *big.Int
	SecondaryTokenDiscount *big.Int
	LiquidityFee           *big.Int
	ReferralPercent        *big.Int
	ReferralToken          common.Address
	ReferralHold           *big.Int
	ReferralDiscount       *big.Int
}, error) {
	return _Unicrypt.Contract.GFees(&_Unicrypt.CallOpts)
}

// GFees is a free data retrieval call binding the contract method 0x90e1a003.
//
// Solidity: function gFees() view returns(uint256 ethFee, address secondaryFeeToken, uint256 secondaryTokenFee, uint256 secondaryTokenDiscount, uint256 liquidityFee, uint256 referralPercent, address referralToken, uint256 referralHold, uint256 referralDiscount)
func (_Unicrypt *UnicryptCallerSession) GFees() (struct {
	EthFee                 *big.Int
	SecondaryFeeToken      common.Address
	SecondaryTokenFee      *big.Int
	SecondaryTokenDiscount *big.Int
	LiquidityFee           *big.Int
	ReferralPercent        *big.Int
	ReferralToken          common.Address
	ReferralHold           *big.Int
	ReferralDiscount       *big.Int
}, error) {
	return _Unicrypt.Contract.GFees(&_Unicrypt.CallOpts)
}

// GetLockedTokenAtIndex is a free data retrieval call binding the contract method 0x14dd79a3.
//
// Solidity: function getLockedTokenAtIndex(uint256 _index) view returns(address)
func (_Unicrypt *UnicryptCaller) GetLockedTokenAtIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "getLockedTokenAtIndex", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLockedTokenAtIndex is a free data retrieval call binding the contract method 0x14dd79a3.
//
// Solidity: function getLockedTokenAtIndex(uint256 _index) view returns(address)
func (_Unicrypt *UnicryptSession) GetLockedTokenAtIndex(_index *big.Int) (common.Address, error) {
	return _Unicrypt.Contract.GetLockedTokenAtIndex(&_Unicrypt.CallOpts, _index)
}

// GetLockedTokenAtIndex is a free data retrieval call binding the contract method 0x14dd79a3.
//
// Solidity: function getLockedTokenAtIndex(uint256 _index) view returns(address)
func (_Unicrypt *UnicryptCallerSession) GetLockedTokenAtIndex(_index *big.Int) (common.Address, error) {
	return _Unicrypt.Contract.GetLockedTokenAtIndex(&_Unicrypt.CallOpts, _index)
}

// GetNumLockedTokens is a free data retrieval call binding the contract method 0x783451e8.
//
// Solidity: function getNumLockedTokens() view returns(uint256)
func (_Unicrypt *UnicryptCaller) GetNumLockedTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "getNumLockedTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumLockedTokens is a free data retrieval call binding the contract method 0x783451e8.
//
// Solidity: function getNumLockedTokens() view returns(uint256)
func (_Unicrypt *UnicryptSession) GetNumLockedTokens() (*big.Int, error) {
	return _Unicrypt.Contract.GetNumLockedTokens(&_Unicrypt.CallOpts)
}

// GetNumLockedTokens is a free data retrieval call binding the contract method 0x783451e8.
//
// Solidity: function getNumLockedTokens() view returns(uint256)
func (_Unicrypt *UnicryptCallerSession) GetNumLockedTokens() (*big.Int, error) {
	return _Unicrypt.Contract.GetNumLockedTokens(&_Unicrypt.CallOpts)
}

// GetNumLocksForToken is a free data retrieval call binding the contract method 0x1f2a1d2f.
//
// Solidity: function getNumLocksForToken(address _lpToken) view returns(uint256)
func (_Unicrypt *UnicryptCaller) GetNumLocksForToken(opts *bind.CallOpts, _lpToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "getNumLocksForToken", _lpToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumLocksForToken is a free data retrieval call binding the contract method 0x1f2a1d2f.
//
// Solidity: function getNumLocksForToken(address _lpToken) view returns(uint256)
func (_Unicrypt *UnicryptSession) GetNumLocksForToken(_lpToken common.Address) (*big.Int, error) {
	return _Unicrypt.Contract.GetNumLocksForToken(&_Unicrypt.CallOpts, _lpToken)
}

// GetNumLocksForToken is a free data retrieval call binding the contract method 0x1f2a1d2f.
//
// Solidity: function getNumLocksForToken(address _lpToken) view returns(uint256)
func (_Unicrypt *UnicryptCallerSession) GetNumLocksForToken(_lpToken common.Address) (*big.Int, error) {
	return _Unicrypt.Contract.GetNumLocksForToken(&_Unicrypt.CallOpts, _lpToken)
}

// GetUserLockForTokenAtIndex is a free data retrieval call binding the contract method 0xd4ff493f.
//
// Solidity: function getUserLockForTokenAtIndex(address _user, address _lpToken, uint256 _index) view returns(uint256, uint256, uint256, uint256, uint256, address)
func (_Unicrypt *UnicryptCaller) GetUserLockForTokenAtIndex(opts *bind.CallOpts, _user common.Address, _lpToken common.Address, _index *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "getUserLockForTokenAtIndex", _user, _lpToken, _index)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, out4, out5, err

}

// GetUserLockForTokenAtIndex is a free data retrieval call binding the contract method 0xd4ff493f.
//
// Solidity: function getUserLockForTokenAtIndex(address _user, address _lpToken, uint256 _index) view returns(uint256, uint256, uint256, uint256, uint256, address)
func (_Unicrypt *UnicryptSession) GetUserLockForTokenAtIndex(_user common.Address, _lpToken common.Address, _index *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, error) {
	return _Unicrypt.Contract.GetUserLockForTokenAtIndex(&_Unicrypt.CallOpts, _user, _lpToken, _index)
}

// GetUserLockForTokenAtIndex is a free data retrieval call binding the contract method 0xd4ff493f.
//
// Solidity: function getUserLockForTokenAtIndex(address _user, address _lpToken, uint256 _index) view returns(uint256, uint256, uint256, uint256, uint256, address)
func (_Unicrypt *UnicryptCallerSession) GetUserLockForTokenAtIndex(_user common.Address, _lpToken common.Address, _index *big.Int) (*big.Int, *big.Int, *big.Int, *big.Int, *big.Int, common.Address, error) {
	return _Unicrypt.Contract.GetUserLockForTokenAtIndex(&_Unicrypt.CallOpts, _user, _lpToken, _index)
}

// GetUserLockedTokenAtIndex is a free data retrieval call binding the contract method 0x903df806.
//
// Solidity: function getUserLockedTokenAtIndex(address _user, uint256 _index) view returns(address)
func (_Unicrypt *UnicryptCaller) GetUserLockedTokenAtIndex(opts *bind.CallOpts, _user common.Address, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "getUserLockedTokenAtIndex", _user, _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetUserLockedTokenAtIndex is a free data retrieval call binding the contract method 0x903df806.
//
// Solidity: function getUserLockedTokenAtIndex(address _user, uint256 _index) view returns(address)
func (_Unicrypt *UnicryptSession) GetUserLockedTokenAtIndex(_user common.Address, _index *big.Int) (common.Address, error) {
	return _Unicrypt.Contract.GetUserLockedTokenAtIndex(&_Unicrypt.CallOpts, _user, _index)
}

// GetUserLockedTokenAtIndex is a free data retrieval call binding the contract method 0x903df806.
//
// Solidity: function getUserLockedTokenAtIndex(address _user, uint256 _index) view returns(address)
func (_Unicrypt *UnicryptCallerSession) GetUserLockedTokenAtIndex(_user common.Address, _index *big.Int) (common.Address, error) {
	return _Unicrypt.Contract.GetUserLockedTokenAtIndex(&_Unicrypt.CallOpts, _user, _index)
}

// GetUserNumLockedTokens is a free data retrieval call binding the contract method 0xa3969815.
//
// Solidity: function getUserNumLockedTokens(address _user) view returns(uint256)
func (_Unicrypt *UnicryptCaller) GetUserNumLockedTokens(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "getUserNumLockedTokens", _user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserNumLockedTokens is a free data retrieval call binding the contract method 0xa3969815.
//
// Solidity: function getUserNumLockedTokens(address _user) view returns(uint256)
func (_Unicrypt *UnicryptSession) GetUserNumLockedTokens(_user common.Address) (*big.Int, error) {
	return _Unicrypt.Contract.GetUserNumLockedTokens(&_Unicrypt.CallOpts, _user)
}

// GetUserNumLockedTokens is a free data retrieval call binding the contract method 0xa3969815.
//
// Solidity: function getUserNumLockedTokens(address _user) view returns(uint256)
func (_Unicrypt *UnicryptCallerSession) GetUserNumLockedTokens(_user common.Address) (*big.Int, error) {
	return _Unicrypt.Contract.GetUserNumLockedTokens(&_Unicrypt.CallOpts, _user)
}

// GetUserNumLocksForToken is a free data retrieval call binding the contract method 0xa69d9c4f.
//
// Solidity: function getUserNumLocksForToken(address _user, address _lpToken) view returns(uint256)
func (_Unicrypt *UnicryptCaller) GetUserNumLocksForToken(opts *bind.CallOpts, _user common.Address, _lpToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "getUserNumLocksForToken", _user, _lpToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserNumLocksForToken is a free data retrieval call binding the contract method 0xa69d9c4f.
//
// Solidity: function getUserNumLocksForToken(address _user, address _lpToken) view returns(uint256)
func (_Unicrypt *UnicryptSession) GetUserNumLocksForToken(_user common.Address, _lpToken common.Address) (*big.Int, error) {
	return _Unicrypt.Contract.GetUserNumLocksForToken(&_Unicrypt.CallOpts, _user, _lpToken)
}

// GetUserNumLocksForToken is a free data retrieval call binding the contract method 0xa69d9c4f.
//
// Solidity: function getUserNumLocksForToken(address _user, address _lpToken) view returns(uint256)
func (_Unicrypt *UnicryptCallerSession) GetUserNumLocksForToken(_user common.Address, _lpToken common.Address) (*big.Int, error) {
	return _Unicrypt.Contract.GetUserNumLocksForToken(&_Unicrypt.CallOpts, _user, _lpToken)
}

// GetUserWhitelistStatus is a free data retrieval call binding the contract method 0xb9863a44.
//
// Solidity: function getUserWhitelistStatus(address _user) view returns(bool)
func (_Unicrypt *UnicryptCaller) GetUserWhitelistStatus(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "getUserWhitelistStatus", _user)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetUserWhitelistStatus is a free data retrieval call binding the contract method 0xb9863a44.
//
// Solidity: function getUserWhitelistStatus(address _user) view returns(bool)
func (_Unicrypt *UnicryptSession) GetUserWhitelistStatus(_user common.Address) (bool, error) {
	return _Unicrypt.Contract.GetUserWhitelistStatus(&_Unicrypt.CallOpts, _user)
}

// GetUserWhitelistStatus is a free data retrieval call binding the contract method 0xb9863a44.
//
// Solidity: function getUserWhitelistStatus(address _user) view returns(bool)
func (_Unicrypt *UnicryptCallerSession) GetUserWhitelistStatus(_user common.Address) (bool, error) {
	return _Unicrypt.Contract.GetUserWhitelistStatus(&_Unicrypt.CallOpts, _user)
}

// GetWhitelistedUserAtIndex is a free data retrieval call binding the contract method 0x8c301df8.
//
// Solidity: function getWhitelistedUserAtIndex(uint256 _index) view returns(address)
func (_Unicrypt *UnicryptCaller) GetWhitelistedUserAtIndex(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "getWhitelistedUserAtIndex", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetWhitelistedUserAtIndex is a free data retrieval call binding the contract method 0x8c301df8.
//
// Solidity: function getWhitelistedUserAtIndex(uint256 _index) view returns(address)
func (_Unicrypt *UnicryptSession) GetWhitelistedUserAtIndex(_index *big.Int) (common.Address, error) {
	return _Unicrypt.Contract.GetWhitelistedUserAtIndex(&_Unicrypt.CallOpts, _index)
}

// GetWhitelistedUserAtIndex is a free data retrieval call binding the contract method 0x8c301df8.
//
// Solidity: function getWhitelistedUserAtIndex(uint256 _index) view returns(address)
func (_Unicrypt *UnicryptCallerSession) GetWhitelistedUserAtIndex(_index *big.Int) (common.Address, error) {
	return _Unicrypt.Contract.GetWhitelistedUserAtIndex(&_Unicrypt.CallOpts, _index)
}

// GetWhitelistedUsersLength is a free data retrieval call binding the contract method 0x4bb18e3f.
//
// Solidity: function getWhitelistedUsersLength() view returns(uint256)
func (_Unicrypt *UnicryptCaller) GetWhitelistedUsersLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "getWhitelistedUsersLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetWhitelistedUsersLength is a free data retrieval call binding the contract method 0x4bb18e3f.
//
// Solidity: function getWhitelistedUsersLength() view returns(uint256)
func (_Unicrypt *UnicryptSession) GetWhitelistedUsersLength() (*big.Int, error) {
	return _Unicrypt.Contract.GetWhitelistedUsersLength(&_Unicrypt.CallOpts)
}

// GetWhitelistedUsersLength is a free data retrieval call binding the contract method 0x4bb18e3f.
//
// Solidity: function getWhitelistedUsersLength() view returns(uint256)
func (_Unicrypt *UnicryptCallerSession) GetWhitelistedUsersLength() (*big.Int, error) {
	return _Unicrypt.Contract.GetWhitelistedUsersLength(&_Unicrypt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Unicrypt *UnicryptCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Unicrypt *UnicryptSession) Owner() (common.Address, error) {
	return _Unicrypt.Contract.Owner(&_Unicrypt.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Unicrypt *UnicryptCallerSession) Owner() (common.Address, error) {
	return _Unicrypt.Contract.Owner(&_Unicrypt.CallOpts)
}

// TokenLocks is a free data retrieval call binding the contract method 0xccebfa3f.
//
// Solidity: function tokenLocks(address , uint256 ) view returns(uint256 lockDate, uint256 amount, uint256 initialAmount, uint256 unlockDate, uint256 lockID, address owner)
func (_Unicrypt *UnicryptCaller) TokenLocks(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	LockDate      *big.Int
	Amount        *big.Int
	InitialAmount *big.Int
	UnlockDate    *big.Int
	LockID        *big.Int
	Owner         common.Address
}, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "tokenLocks", arg0, arg1)

	outstruct := new(struct {
		LockDate      *big.Int
		Amount        *big.Int
		InitialAmount *big.Int
		UnlockDate    *big.Int
		LockID        *big.Int
		Owner         common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.LockDate = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.InitialAmount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.UnlockDate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.LockID = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Owner = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// TokenLocks is a free data retrieval call binding the contract method 0xccebfa3f.
//
// Solidity: function tokenLocks(address , uint256 ) view returns(uint256 lockDate, uint256 amount, uint256 initialAmount, uint256 unlockDate, uint256 lockID, address owner)
func (_Unicrypt *UnicryptSession) TokenLocks(arg0 common.Address, arg1 *big.Int) (struct {
	LockDate      *big.Int
	Amount        *big.Int
	InitialAmount *big.Int
	UnlockDate    *big.Int
	LockID        *big.Int
	Owner         common.Address
}, error) {
	return _Unicrypt.Contract.TokenLocks(&_Unicrypt.CallOpts, arg0, arg1)
}

// TokenLocks is a free data retrieval call binding the contract method 0xccebfa3f.
//
// Solidity: function tokenLocks(address , uint256 ) view returns(uint256 lockDate, uint256 amount, uint256 initialAmount, uint256 unlockDate, uint256 lockID, address owner)
func (_Unicrypt *UnicryptCallerSession) TokenLocks(arg0 common.Address, arg1 *big.Int) (struct {
	LockDate      *big.Int
	Amount        *big.Int
	InitialAmount *big.Int
	UnlockDate    *big.Int
	LockID        *big.Int
	Owner         common.Address
}, error) {
	return _Unicrypt.Contract.TokenLocks(&_Unicrypt.CallOpts, arg0, arg1)
}

// UniswapFactory is a free data retrieval call binding the contract method 0x8bdb2afa.
//
// Solidity: function uniswapFactory() view returns(address)
func (_Unicrypt *UnicryptCaller) UniswapFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Unicrypt.contract.Call(opts, &out, "uniswapFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UniswapFactory is a free data retrieval call binding the contract method 0x8bdb2afa.
//
// Solidity: function uniswapFactory() view returns(address)
func (_Unicrypt *UnicryptSession) UniswapFactory() (common.Address, error) {
	return _Unicrypt.Contract.UniswapFactory(&_Unicrypt.CallOpts)
}

// UniswapFactory is a free data retrieval call binding the contract method 0x8bdb2afa.
//
// Solidity: function uniswapFactory() view returns(address)
func (_Unicrypt *UnicryptCallerSession) UniswapFactory() (common.Address, error) {
	return _Unicrypt.Contract.UniswapFactory(&_Unicrypt.CallOpts)
}

// IncrementLock is a paid mutator transaction binding the contract method 0xa9b07cea.
//
// Solidity: function incrementLock(address _lpToken, uint256 _index, uint256 _lockID, uint256 _amount) returns()
func (_Unicrypt *UnicryptTransactor) IncrementLock(opts *bind.TransactOpts, _lpToken common.Address, _index *big.Int, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "incrementLock", _lpToken, _index, _lockID, _amount)
}

// IncrementLock is a paid mutator transaction binding the contract method 0xa9b07cea.
//
// Solidity: function incrementLock(address _lpToken, uint256 _index, uint256 _lockID, uint256 _amount) returns()
func (_Unicrypt *UnicryptSession) IncrementLock(_lpToken common.Address, _index *big.Int, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.IncrementLock(&_Unicrypt.TransactOpts, _lpToken, _index, _lockID, _amount)
}

// IncrementLock is a paid mutator transaction binding the contract method 0xa9b07cea.
//
// Solidity: function incrementLock(address _lpToken, uint256 _index, uint256 _lockID, uint256 _amount) returns()
func (_Unicrypt *UnicryptTransactorSession) IncrementLock(_lpToken common.Address, _index *big.Int, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.IncrementLock(&_Unicrypt.TransactOpts, _lpToken, _index, _lockID, _amount)
}

// LockLPToken is a paid mutator transaction binding the contract method 0x8af416f6.
//
// Solidity: function lockLPToken(address _lpToken, uint256 _amount, uint256 _unlock_date, address _referral, bool _fee_in_eth, address _withdrawer) payable returns()
func (_Unicrypt *UnicryptTransactor) LockLPToken(opts *bind.TransactOpts, _lpToken common.Address, _amount *big.Int, _unlock_date *big.Int, _referral common.Address, _fee_in_eth bool, _withdrawer common.Address) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "lockLPToken", _lpToken, _amount, _unlock_date, _referral, _fee_in_eth, _withdrawer)
}

// LockLPToken is a paid mutator transaction binding the contract method 0x8af416f6.
//
// Solidity: function lockLPToken(address _lpToken, uint256 _amount, uint256 _unlock_date, address _referral, bool _fee_in_eth, address _withdrawer) payable returns()
func (_Unicrypt *UnicryptSession) LockLPToken(_lpToken common.Address, _amount *big.Int, _unlock_date *big.Int, _referral common.Address, _fee_in_eth bool, _withdrawer common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.LockLPToken(&_Unicrypt.TransactOpts, _lpToken, _amount, _unlock_date, _referral, _fee_in_eth, _withdrawer)
}

// LockLPToken is a paid mutator transaction binding the contract method 0x8af416f6.
//
// Solidity: function lockLPToken(address _lpToken, uint256 _amount, uint256 _unlock_date, address _referral, bool _fee_in_eth, address _withdrawer) payable returns()
func (_Unicrypt *UnicryptTransactorSession) LockLPToken(_lpToken common.Address, _amount *big.Int, _unlock_date *big.Int, _referral common.Address, _fee_in_eth bool, _withdrawer common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.LockLPToken(&_Unicrypt.TransactOpts, _lpToken, _amount, _unlock_date, _referral, _fee_in_eth, _withdrawer)
}

// Migrate is a paid mutator transaction binding the contract method 0xee424278.
//
// Solidity: function migrate(address _lpToken, uint256 _index, uint256 _lockID, uint256 _amount) returns()
func (_Unicrypt *UnicryptTransactor) Migrate(opts *bind.TransactOpts, _lpToken common.Address, _index *big.Int, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "migrate", _lpToken, _index, _lockID, _amount)
}

// Migrate is a paid mutator transaction binding the contract method 0xee424278.
//
// Solidity: function migrate(address _lpToken, uint256 _index, uint256 _lockID, uint256 _amount) returns()
func (_Unicrypt *UnicryptSession) Migrate(_lpToken common.Address, _index *big.Int, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.Migrate(&_Unicrypt.TransactOpts, _lpToken, _index, _lockID, _amount)
}

// Migrate is a paid mutator transaction binding the contract method 0xee424278.
//
// Solidity: function migrate(address _lpToken, uint256 _index, uint256 _lockID, uint256 _amount) returns()
func (_Unicrypt *UnicryptTransactorSession) Migrate(_lpToken common.Address, _index *big.Int, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.Migrate(&_Unicrypt.TransactOpts, _lpToken, _index, _lockID, _amount)
}

// Relock is a paid mutator transaction binding the contract method 0x60491d24.
//
// Solidity: function relock(address _lpToken, uint256 _index, uint256 _lockID, uint256 _unlock_date) returns()
func (_Unicrypt *UnicryptTransactor) Relock(opts *bind.TransactOpts, _lpToken common.Address, _index *big.Int, _lockID *big.Int, _unlock_date *big.Int) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "relock", _lpToken, _index, _lockID, _unlock_date)
}

// Relock is a paid mutator transaction binding the contract method 0x60491d24.
//
// Solidity: function relock(address _lpToken, uint256 _index, uint256 _lockID, uint256 _unlock_date) returns()
func (_Unicrypt *UnicryptSession) Relock(_lpToken common.Address, _index *big.Int, _lockID *big.Int, _unlock_date *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.Relock(&_Unicrypt.TransactOpts, _lpToken, _index, _lockID, _unlock_date)
}

// Relock is a paid mutator transaction binding the contract method 0x60491d24.
//
// Solidity: function relock(address _lpToken, uint256 _index, uint256 _lockID, uint256 _unlock_date) returns()
func (_Unicrypt *UnicryptTransactorSession) Relock(_lpToken common.Address, _index *big.Int, _lockID *big.Int, _unlock_date *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.Relock(&_Unicrypt.TransactOpts, _lpToken, _index, _lockID, _unlock_date)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Unicrypt *UnicryptTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Unicrypt *UnicryptSession) RenounceOwnership() (*types.Transaction, error) {
	return _Unicrypt.Contract.RenounceOwnership(&_Unicrypt.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Unicrypt *UnicryptTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Unicrypt.Contract.RenounceOwnership(&_Unicrypt.TransactOpts)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address _devaddr) returns()
func (_Unicrypt *UnicryptTransactor) SetDev(opts *bind.TransactOpts, _devaddr common.Address) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "setDev", _devaddr)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address _devaddr) returns()
func (_Unicrypt *UnicryptSession) SetDev(_devaddr common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.SetDev(&_Unicrypt.TransactOpts, _devaddr)
}

// SetDev is a paid mutator transaction binding the contract method 0xd477f05f.
//
// Solidity: function setDev(address _devaddr) returns()
func (_Unicrypt *UnicryptTransactorSession) SetDev(_devaddr common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.SetDev(&_Unicrypt.TransactOpts, _devaddr)
}

// SetFees is a paid mutator transaction binding the contract method 0x86f6c3c1.
//
// Solidity: function setFees(uint256 _referralPercent, uint256 _referralDiscount, uint256 _ethFee, uint256 _secondaryTokenFee, uint256 _secondaryTokenDiscount, uint256 _liquidityFee) returns()
func (_Unicrypt *UnicryptTransactor) SetFees(opts *bind.TransactOpts, _referralPercent *big.Int, _referralDiscount *big.Int, _ethFee *big.Int, _secondaryTokenFee *big.Int, _secondaryTokenDiscount *big.Int, _liquidityFee *big.Int) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "setFees", _referralPercent, _referralDiscount, _ethFee, _secondaryTokenFee, _secondaryTokenDiscount, _liquidityFee)
}

// SetFees is a paid mutator transaction binding the contract method 0x86f6c3c1.
//
// Solidity: function setFees(uint256 _referralPercent, uint256 _referralDiscount, uint256 _ethFee, uint256 _secondaryTokenFee, uint256 _secondaryTokenDiscount, uint256 _liquidityFee) returns()
func (_Unicrypt *UnicryptSession) SetFees(_referralPercent *big.Int, _referralDiscount *big.Int, _ethFee *big.Int, _secondaryTokenFee *big.Int, _secondaryTokenDiscount *big.Int, _liquidityFee *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.SetFees(&_Unicrypt.TransactOpts, _referralPercent, _referralDiscount, _ethFee, _secondaryTokenFee, _secondaryTokenDiscount, _liquidityFee)
}

// SetFees is a paid mutator transaction binding the contract method 0x86f6c3c1.
//
// Solidity: function setFees(uint256 _referralPercent, uint256 _referralDiscount, uint256 _ethFee, uint256 _secondaryTokenFee, uint256 _secondaryTokenDiscount, uint256 _liquidityFee) returns()
func (_Unicrypt *UnicryptTransactorSession) SetFees(_referralPercent *big.Int, _referralDiscount *big.Int, _ethFee *big.Int, _secondaryTokenFee *big.Int, _secondaryTokenDiscount *big.Int, _liquidityFee *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.SetFees(&_Unicrypt.TransactOpts, _referralPercent, _referralDiscount, _ethFee, _secondaryTokenFee, _secondaryTokenDiscount, _liquidityFee)
}

// SetMigrationIn is a paid mutator transaction binding the contract method 0x44c427e2.
//
// Solidity: function setMigrationIn(address _migrationIn) returns()
func (_Unicrypt *UnicryptTransactor) SetMigrationIn(opts *bind.TransactOpts, _migrationIn common.Address) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "setMigrationIn", _migrationIn)
}

// SetMigrationIn is a paid mutator transaction binding the contract method 0x44c427e2.
//
// Solidity: function setMigrationIn(address _migrationIn) returns()
func (_Unicrypt *UnicryptSession) SetMigrationIn(_migrationIn common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.SetMigrationIn(&_Unicrypt.TransactOpts, _migrationIn)
}

// SetMigrationIn is a paid mutator transaction binding the contract method 0x44c427e2.
//
// Solidity: function setMigrationIn(address _migrationIn) returns()
func (_Unicrypt *UnicryptTransactorSession) SetMigrationIn(_migrationIn common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.SetMigrationIn(&_Unicrypt.TransactOpts, _migrationIn)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Unicrypt *UnicryptTransactor) SetMigrator(opts *bind.TransactOpts, _migrator common.Address) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "setMigrator", _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Unicrypt *UnicryptSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.SetMigrator(&_Unicrypt.TransactOpts, _migrator)
}

// SetMigrator is a paid mutator transaction binding the contract method 0x23cf3118.
//
// Solidity: function setMigrator(address _migrator) returns()
func (_Unicrypt *UnicryptTransactorSession) SetMigrator(_migrator common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.SetMigrator(&_Unicrypt.TransactOpts, _migrator)
}

// SetReferralTokenAndHold is a paid mutator transaction binding the contract method 0xf02c2643.
//
// Solidity: function setReferralTokenAndHold(address _referralToken, uint256 _hold) returns()
func (_Unicrypt *UnicryptTransactor) SetReferralTokenAndHold(opts *bind.TransactOpts, _referralToken common.Address, _hold *big.Int) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "setReferralTokenAndHold", _referralToken, _hold)
}

// SetReferralTokenAndHold is a paid mutator transaction binding the contract method 0xf02c2643.
//
// Solidity: function setReferralTokenAndHold(address _referralToken, uint256 _hold) returns()
func (_Unicrypt *UnicryptSession) SetReferralTokenAndHold(_referralToken common.Address, _hold *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.SetReferralTokenAndHold(&_Unicrypt.TransactOpts, _referralToken, _hold)
}

// SetReferralTokenAndHold is a paid mutator transaction binding the contract method 0xf02c2643.
//
// Solidity: function setReferralTokenAndHold(address _referralToken, uint256 _hold) returns()
func (_Unicrypt *UnicryptTransactorSession) SetReferralTokenAndHold(_referralToken common.Address, _hold *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.SetReferralTokenAndHold(&_Unicrypt.TransactOpts, _referralToken, _hold)
}

// SetSecondaryFeeToken is a paid mutator transaction binding the contract method 0x8931a4be.
//
// Solidity: function setSecondaryFeeToken(address _secondaryFeeToken) returns()
func (_Unicrypt *UnicryptTransactor) SetSecondaryFeeToken(opts *bind.TransactOpts, _secondaryFeeToken common.Address) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "setSecondaryFeeToken", _secondaryFeeToken)
}

// SetSecondaryFeeToken is a paid mutator transaction binding the contract method 0x8931a4be.
//
// Solidity: function setSecondaryFeeToken(address _secondaryFeeToken) returns()
func (_Unicrypt *UnicryptSession) SetSecondaryFeeToken(_secondaryFeeToken common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.SetSecondaryFeeToken(&_Unicrypt.TransactOpts, _secondaryFeeToken)
}

// SetSecondaryFeeToken is a paid mutator transaction binding the contract method 0x8931a4be.
//
// Solidity: function setSecondaryFeeToken(address _secondaryFeeToken) returns()
func (_Unicrypt *UnicryptTransactorSession) SetSecondaryFeeToken(_secondaryFeeToken common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.SetSecondaryFeeToken(&_Unicrypt.TransactOpts, _secondaryFeeToken)
}

// SplitLock is a paid mutator transaction binding the contract method 0x582d5adc.
//
// Solidity: function splitLock(address _lpToken, uint256 _index, uint256 _lockID, uint256 _amount) payable returns()
func (_Unicrypt *UnicryptTransactor) SplitLock(opts *bind.TransactOpts, _lpToken common.Address, _index *big.Int, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "splitLock", _lpToken, _index, _lockID, _amount)
}

// SplitLock is a paid mutator transaction binding the contract method 0x582d5adc.
//
// Solidity: function splitLock(address _lpToken, uint256 _index, uint256 _lockID, uint256 _amount) payable returns()
func (_Unicrypt *UnicryptSession) SplitLock(_lpToken common.Address, _index *big.Int, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.SplitLock(&_Unicrypt.TransactOpts, _lpToken, _index, _lockID, _amount)
}

// SplitLock is a paid mutator transaction binding the contract method 0x582d5adc.
//
// Solidity: function splitLock(address _lpToken, uint256 _index, uint256 _lockID, uint256 _amount) payable returns()
func (_Unicrypt *UnicryptTransactorSession) SplitLock(_lpToken common.Address, _index *big.Int, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.SplitLock(&_Unicrypt.TransactOpts, _lpToken, _index, _lockID, _amount)
}

// TransferLockOwnership is a paid mutator transaction binding the contract method 0xbef497fd.
//
// Solidity: function transferLockOwnership(address _lpToken, uint256 _index, uint256 _lockID, address _newOwner) returns()
func (_Unicrypt *UnicryptTransactor) TransferLockOwnership(opts *bind.TransactOpts, _lpToken common.Address, _index *big.Int, _lockID *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "transferLockOwnership", _lpToken, _index, _lockID, _newOwner)
}

// TransferLockOwnership is a paid mutator transaction binding the contract method 0xbef497fd.
//
// Solidity: function transferLockOwnership(address _lpToken, uint256 _index, uint256 _lockID, address _newOwner) returns()
func (_Unicrypt *UnicryptSession) TransferLockOwnership(_lpToken common.Address, _index *big.Int, _lockID *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.TransferLockOwnership(&_Unicrypt.TransactOpts, _lpToken, _index, _lockID, _newOwner)
}

// TransferLockOwnership is a paid mutator transaction binding the contract method 0xbef497fd.
//
// Solidity: function transferLockOwnership(address _lpToken, uint256 _index, uint256 _lockID, address _newOwner) returns()
func (_Unicrypt *UnicryptTransactorSession) TransferLockOwnership(_lpToken common.Address, _index *big.Int, _lockID *big.Int, _newOwner common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.TransferLockOwnership(&_Unicrypt.TransactOpts, _lpToken, _index, _lockID, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Unicrypt *UnicryptTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Unicrypt *UnicryptSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.TransferOwnership(&_Unicrypt.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Unicrypt *UnicryptTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Unicrypt.Contract.TransferOwnership(&_Unicrypt.TransactOpts, newOwner)
}

// WhitelistFeeAccount is a paid mutator transaction binding the contract method 0x91ff1eb1.
//
// Solidity: function whitelistFeeAccount(address _user, bool _add) returns()
func (_Unicrypt *UnicryptTransactor) WhitelistFeeAccount(opts *bind.TransactOpts, _user common.Address, _add bool) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "whitelistFeeAccount", _user, _add)
}

// WhitelistFeeAccount is a paid mutator transaction binding the contract method 0x91ff1eb1.
//
// Solidity: function whitelistFeeAccount(address _user, bool _add) returns()
func (_Unicrypt *UnicryptSession) WhitelistFeeAccount(_user common.Address, _add bool) (*types.Transaction, error) {
	return _Unicrypt.Contract.WhitelistFeeAccount(&_Unicrypt.TransactOpts, _user, _add)
}

// WhitelistFeeAccount is a paid mutator transaction binding the contract method 0x91ff1eb1.
//
// Solidity: function whitelistFeeAccount(address _user, bool _add) returns()
func (_Unicrypt *UnicryptTransactorSession) WhitelistFeeAccount(_user common.Address, _add bool) (*types.Transaction, error) {
	return _Unicrypt.Contract.WhitelistFeeAccount(&_Unicrypt.TransactOpts, _user, _add)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4532d776.
//
// Solidity: function withdraw(address _lpToken, uint256 _index, uint256 _lockID, uint256 _amount) returns()
func (_Unicrypt *UnicryptTransactor) Withdraw(opts *bind.TransactOpts, _lpToken common.Address, _index *big.Int, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Unicrypt.contract.Transact(opts, "withdraw", _lpToken, _index, _lockID, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4532d776.
//
// Solidity: function withdraw(address _lpToken, uint256 _index, uint256 _lockID, uint256 _amount) returns()
func (_Unicrypt *UnicryptSession) Withdraw(_lpToken common.Address, _index *big.Int, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.Withdraw(&_Unicrypt.TransactOpts, _lpToken, _index, _lockID, _amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0x4532d776.
//
// Solidity: function withdraw(address _lpToken, uint256 _index, uint256 _lockID, uint256 _amount) returns()
func (_Unicrypt *UnicryptTransactorSession) Withdraw(_lpToken common.Address, _index *big.Int, _lockID *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _Unicrypt.Contract.Withdraw(&_Unicrypt.TransactOpts, _lpToken, _index, _lockID, _amount)
}

// UnicryptOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Unicrypt contract.
type UnicryptOwnershipTransferredIterator struct {
	Event *UnicryptOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UnicryptOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicryptOwnershipTransferred)
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
		it.Event = new(UnicryptOwnershipTransferred)
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
func (it *UnicryptOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicryptOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicryptOwnershipTransferred represents a OwnershipTransferred event raised by the Unicrypt contract.
type UnicryptOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Unicrypt *UnicryptFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UnicryptOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Unicrypt.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UnicryptOwnershipTransferredIterator{contract: _Unicrypt.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Unicrypt *UnicryptFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UnicryptOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Unicrypt.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicryptOwnershipTransferred)
				if err := _Unicrypt.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Unicrypt *UnicryptFilterer) ParseOwnershipTransferred(log types.Log) (*UnicryptOwnershipTransferred, error) {
	event := new(UnicryptOwnershipTransferred)
	if err := _Unicrypt.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicryptOnDepositIterator is returned from FilterOnDeposit and is used to iterate over the raw logs and unpacked data for OnDeposit events raised by the Unicrypt contract.
type UnicryptOnDepositIterator struct {
	Event *UnicryptOnDeposit // Event containing the contract specifics and raw log

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
func (it *UnicryptOnDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicryptOnDeposit)
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
		it.Event = new(UnicryptOnDeposit)
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
func (it *UnicryptOnDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicryptOnDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicryptOnDeposit represents a OnDeposit event raised by the Unicrypt contract.
type UnicryptOnDeposit struct {
	LpToken    common.Address
	User       common.Address
	Amount     *big.Int
	LockDate   *big.Int
	UnlockDate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterOnDeposit is a free log retrieval operation binding the contract event 0x830357565da6ecfc26d8d9f69df488ed6f70361af9a07e570544aeb5c5e765e5.
//
// Solidity: event onDeposit(address lpToken, address user, uint256 amount, uint256 lockDate, uint256 unlockDate)
func (_Unicrypt *UnicryptFilterer) FilterOnDeposit(opts *bind.FilterOpts) (*UnicryptOnDepositIterator, error) {

	logs, sub, err := _Unicrypt.contract.FilterLogs(opts, "onDeposit")
	if err != nil {
		return nil, err
	}
	return &UnicryptOnDepositIterator{contract: _Unicrypt.contract, event: "onDeposit", logs: logs, sub: sub}, nil
}

// WatchOnDeposit is a free log subscription operation binding the contract event 0x830357565da6ecfc26d8d9f69df488ed6f70361af9a07e570544aeb5c5e765e5.
//
// Solidity: event onDeposit(address lpToken, address user, uint256 amount, uint256 lockDate, uint256 unlockDate)
func (_Unicrypt *UnicryptFilterer) WatchOnDeposit(opts *bind.WatchOpts, sink chan<- *UnicryptOnDeposit) (event.Subscription, error) {

	logs, sub, err := _Unicrypt.contract.WatchLogs(opts, "onDeposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicryptOnDeposit)
				if err := _Unicrypt.contract.UnpackLog(event, "onDeposit", log); err != nil {
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

// ParseOnDeposit is a log parse operation binding the contract event 0x830357565da6ecfc26d8d9f69df488ed6f70361af9a07e570544aeb5c5e765e5.
//
// Solidity: event onDeposit(address lpToken, address user, uint256 amount, uint256 lockDate, uint256 unlockDate)
func (_Unicrypt *UnicryptFilterer) ParseOnDeposit(log types.Log) (*UnicryptOnDeposit, error) {
	event := new(UnicryptOnDeposit)
	if err := _Unicrypt.contract.UnpackLog(event, "onDeposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnicryptOnWithdrawIterator is returned from FilterOnWithdraw and is used to iterate over the raw logs and unpacked data for OnWithdraw events raised by the Unicrypt contract.
type UnicryptOnWithdrawIterator struct {
	Event *UnicryptOnWithdraw // Event containing the contract specifics and raw log

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
func (it *UnicryptOnWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnicryptOnWithdraw)
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
		it.Event = new(UnicryptOnWithdraw)
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
func (it *UnicryptOnWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnicryptOnWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnicryptOnWithdraw represents a OnWithdraw event raised by the Unicrypt contract.
type UnicryptOnWithdraw struct {
	LpToken common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterOnWithdraw is a free log retrieval operation binding the contract event 0xccad973dcd043c7d680389db4378bd6b9775db7124092e9e0422c9e46d7985dc.
//
// Solidity: event onWithdraw(address lpToken, uint256 amount)
func (_Unicrypt *UnicryptFilterer) FilterOnWithdraw(opts *bind.FilterOpts) (*UnicryptOnWithdrawIterator, error) {

	logs, sub, err := _Unicrypt.contract.FilterLogs(opts, "onWithdraw")
	if err != nil {
		return nil, err
	}
	return &UnicryptOnWithdrawIterator{contract: _Unicrypt.contract, event: "onWithdraw", logs: logs, sub: sub}, nil
}

// WatchOnWithdraw is a free log subscription operation binding the contract event 0xccad973dcd043c7d680389db4378bd6b9775db7124092e9e0422c9e46d7985dc.
//
// Solidity: event onWithdraw(address lpToken, uint256 amount)
func (_Unicrypt *UnicryptFilterer) WatchOnWithdraw(opts *bind.WatchOpts, sink chan<- *UnicryptOnWithdraw) (event.Subscription, error) {

	logs, sub, err := _Unicrypt.contract.WatchLogs(opts, "onWithdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnicryptOnWithdraw)
				if err := _Unicrypt.contract.UnpackLog(event, "onWithdraw", log); err != nil {
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

// ParseOnWithdraw is a log parse operation binding the contract event 0xccad973dcd043c7d680389db4378bd6b9775db7124092e9e0422c9e46d7985dc.
//
// Solidity: event onWithdraw(address lpToken, uint256 amount)
func (_Unicrypt *UnicryptFilterer) ParseOnWithdraw(log types.Log) (*UnicryptOnWithdraw, error) {
	event := new(UnicryptOnWithdraw)
	if err := _Unicrypt.contract.UnpackLog(event, "onWithdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
