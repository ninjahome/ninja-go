// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NinjaChatLicenseABI is the input ABI used to generate the binding from.
const NinjaChatLicenseABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issueAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"recvAddr\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nDays\",\"type\":\"uint32\"}],\"name\":\"BindLicenseEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"payerAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userAddr\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nDays\",\"type\":\"uint32\"}],\"name\":\"ChargeUserEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"issueAddr\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"nDays\",\"type\":\"uint32\"}],\"name\":\"GenerateLicenseEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"issueAddr\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"recvAddr\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nDays\",\"type\":\"uint32\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"BindLicense\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"userAddr\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nDays\",\"type\":\"uint32\"}],\"name\":\"ChargeUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"uint32\",\"name\":\"nDays\",\"type\":\"uint32\"}],\"name\":\"GenerateLicense\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetSettings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"userAddr\",\"type\":\"bytes32\"}],\"name\":\"GetUserLicense\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"Licenses\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"used\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"nDays\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"nAddr\",\"type\":\"address\"}],\"name\":\"SetNinjaAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"}],\"name\":\"SetTokenAddr\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nAddr\",\"type\":\"address\"}],\"name\":\"Setting\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"UserLicenses\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"EndDays\",\"type\":\"uint64\"},{\"internalType\":\"uint32\",\"name\":\"TotalCoins\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ninjaAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractNinjaToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NinjaChatLicense is an auto generated Go binding around an Ethereum contract.
type NinjaChatLicense struct {
	NinjaChatLicenseCaller     // Read-only binding to the contract
	NinjaChatLicenseTransactor // Write-only binding to the contract
	NinjaChatLicenseFilterer   // Log filterer for contract events
}

// NinjaChatLicenseCaller is an auto generated read-only Go binding around an Ethereum contract.
type NinjaChatLicenseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NinjaChatLicenseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NinjaChatLicenseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NinjaChatLicenseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NinjaChatLicenseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NinjaChatLicenseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NinjaChatLicenseSession struct {
	Contract     *NinjaChatLicense // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NinjaChatLicenseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NinjaChatLicenseCallerSession struct {
	Contract *NinjaChatLicenseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// NinjaChatLicenseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NinjaChatLicenseTransactorSession struct {
	Contract     *NinjaChatLicenseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// NinjaChatLicenseRaw is an auto generated low-level Go binding around an Ethereum contract.
type NinjaChatLicenseRaw struct {
	Contract *NinjaChatLicense // Generic contract binding to access the raw methods on
}

// NinjaChatLicenseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NinjaChatLicenseCallerRaw struct {
	Contract *NinjaChatLicenseCaller // Generic read-only contract binding to access the raw methods on
}

// NinjaChatLicenseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NinjaChatLicenseTransactorRaw struct {
	Contract *NinjaChatLicenseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNinjaChatLicense creates a new instance of NinjaChatLicense, bound to a specific deployed contract.
func NewNinjaChatLicense(address common.Address, backend bind.ContractBackend) (*NinjaChatLicense, error) {
	contract, err := bindNinjaChatLicense(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NinjaChatLicense{NinjaChatLicenseCaller: NinjaChatLicenseCaller{contract: contract}, NinjaChatLicenseTransactor: NinjaChatLicenseTransactor{contract: contract}, NinjaChatLicenseFilterer: NinjaChatLicenseFilterer{contract: contract}}, nil
}

// NewNinjaChatLicenseCaller creates a new read-only instance of NinjaChatLicense, bound to a specific deployed contract.
func NewNinjaChatLicenseCaller(address common.Address, caller bind.ContractCaller) (*NinjaChatLicenseCaller, error) {
	contract, err := bindNinjaChatLicense(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NinjaChatLicenseCaller{contract: contract}, nil
}

// NewNinjaChatLicenseTransactor creates a new write-only instance of NinjaChatLicense, bound to a specific deployed contract.
func NewNinjaChatLicenseTransactor(address common.Address, transactor bind.ContractTransactor) (*NinjaChatLicenseTransactor, error) {
	contract, err := bindNinjaChatLicense(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NinjaChatLicenseTransactor{contract: contract}, nil
}

// NewNinjaChatLicenseFilterer creates a new log filterer instance of NinjaChatLicense, bound to a specific deployed contract.
func NewNinjaChatLicenseFilterer(address common.Address, filterer bind.ContractFilterer) (*NinjaChatLicenseFilterer, error) {
	contract, err := bindNinjaChatLicense(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NinjaChatLicenseFilterer{contract: contract}, nil
}

// bindNinjaChatLicense binds a generic wrapper to an already deployed contract.
func bindNinjaChatLicense(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NinjaChatLicenseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NinjaChatLicense *NinjaChatLicenseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NinjaChatLicense.Contract.NinjaChatLicenseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NinjaChatLicense *NinjaChatLicenseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.NinjaChatLicenseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NinjaChatLicense *NinjaChatLicenseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.NinjaChatLicenseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NinjaChatLicense *NinjaChatLicenseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NinjaChatLicense.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NinjaChatLicense *NinjaChatLicenseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NinjaChatLicense *NinjaChatLicenseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.contract.Transact(opts, method, params...)
}

// GetSettings is a free data retrieval call binding the contract method 0x9e998bc8.
//
// Solidity: function GetSettings() view returns(address, address)
func (_NinjaChatLicense *NinjaChatLicenseCaller) GetSettings(opts *bind.CallOpts) (common.Address, common.Address, error) {
	var out []interface{}
	err := _NinjaChatLicense.contract.Call(opts, &out, "GetSettings")

	if err != nil {
		return *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetSettings is a free data retrieval call binding the contract method 0x9e998bc8.
//
// Solidity: function GetSettings() view returns(address, address)
func (_NinjaChatLicense *NinjaChatLicenseSession) GetSettings() (common.Address, common.Address, error) {
	return _NinjaChatLicense.Contract.GetSettings(&_NinjaChatLicense.CallOpts)
}

// GetSettings is a free data retrieval call binding the contract method 0x9e998bc8.
//
// Solidity: function GetSettings() view returns(address, address)
func (_NinjaChatLicense *NinjaChatLicenseCallerSession) GetSettings() (common.Address, common.Address, error) {
	return _NinjaChatLicense.Contract.GetSettings(&_NinjaChatLicense.CallOpts)
}

// GetUserLicense is a free data retrieval call binding the contract method 0xa6405426.
//
// Solidity: function GetUserLicense(bytes32 userAddr) view returns(uint64, uint32)
func (_NinjaChatLicense *NinjaChatLicenseCaller) GetUserLicense(opts *bind.CallOpts, userAddr [32]byte) (uint64, uint32, error) {
	var out []interface{}
	err := _NinjaChatLicense.contract.Call(opts, &out, "GetUserLicense", userAddr)

	if err != nil {
		return *new(uint64), *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)

	return out0, out1, err

}

// GetUserLicense is a free data retrieval call binding the contract method 0xa6405426.
//
// Solidity: function GetUserLicense(bytes32 userAddr) view returns(uint64, uint32)
func (_NinjaChatLicense *NinjaChatLicenseSession) GetUserLicense(userAddr [32]byte) (uint64, uint32, error) {
	return _NinjaChatLicense.Contract.GetUserLicense(&_NinjaChatLicense.CallOpts, userAddr)
}

// GetUserLicense is a free data retrieval call binding the contract method 0xa6405426.
//
// Solidity: function GetUserLicense(bytes32 userAddr) view returns(uint64, uint32)
func (_NinjaChatLicense *NinjaChatLicenseCallerSession) GetUserLicense(userAddr [32]byte) (uint64, uint32, error) {
	return _NinjaChatLicense.Contract.GetUserLicense(&_NinjaChatLicense.CallOpts, userAddr)
}

// Licenses is a free data retrieval call binding the contract method 0x4a627817.
//
// Solidity: function Licenses(address , bytes32 ) view returns(bool used, uint32 nDays)
func (_NinjaChatLicense *NinjaChatLicenseCaller) Licenses(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	Used  bool
	NDays uint32
}, error) {
	var out []interface{}
	err := _NinjaChatLicense.contract.Call(opts, &out, "Licenses", arg0, arg1)

	outstruct := new(struct {
		Used  bool
		NDays uint32
	})

	outstruct.Used = out[0].(bool)
	outstruct.NDays = out[1].(uint32)

	return *outstruct, err

}

// Licenses is a free data retrieval call binding the contract method 0x4a627817.
//
// Solidity: function Licenses(address , bytes32 ) view returns(bool used, uint32 nDays)
func (_NinjaChatLicense *NinjaChatLicenseSession) Licenses(arg0 common.Address, arg1 [32]byte) (struct {
	Used  bool
	NDays uint32
}, error) {
	return _NinjaChatLicense.Contract.Licenses(&_NinjaChatLicense.CallOpts, arg0, arg1)
}

// Licenses is a free data retrieval call binding the contract method 0x4a627817.
//
// Solidity: function Licenses(address , bytes32 ) view returns(bool used, uint32 nDays)
func (_NinjaChatLicense *NinjaChatLicenseCallerSession) Licenses(arg0 common.Address, arg1 [32]byte) (struct {
	Used  bool
	NDays uint32
}, error) {
	return _NinjaChatLicense.Contract.Licenses(&_NinjaChatLicense.CallOpts, arg0, arg1)
}

// UserLicenses is a free data retrieval call binding the contract method 0x2abefe3a.
//
// Solidity: function UserLicenses(bytes32 ) view returns(uint64 EndDays, uint32 TotalCoins)
func (_NinjaChatLicense *NinjaChatLicenseCaller) UserLicenses(opts *bind.CallOpts, arg0 [32]byte) (struct {
	EndDays    uint64
	TotalCoins uint32
}, error) {
	var out []interface{}
	err := _NinjaChatLicense.contract.Call(opts, &out, "UserLicenses", arg0)

	outstruct := new(struct {
		EndDays    uint64
		TotalCoins uint32
	})

	outstruct.EndDays = out[0].(uint64)
	outstruct.TotalCoins = out[1].(uint32)

	return *outstruct, err

}

// UserLicenses is a free data retrieval call binding the contract method 0x2abefe3a.
//
// Solidity: function UserLicenses(bytes32 ) view returns(uint64 EndDays, uint32 TotalCoins)
func (_NinjaChatLicense *NinjaChatLicenseSession) UserLicenses(arg0 [32]byte) (struct {
	EndDays    uint64
	TotalCoins uint32
}, error) {
	return _NinjaChatLicense.Contract.UserLicenses(&_NinjaChatLicense.CallOpts, arg0)
}

// UserLicenses is a free data retrieval call binding the contract method 0x2abefe3a.
//
// Solidity: function UserLicenses(bytes32 ) view returns(uint64 EndDays, uint32 TotalCoins)
func (_NinjaChatLicense *NinjaChatLicenseCallerSession) UserLicenses(arg0 [32]byte) (struct {
	EndDays    uint64
	TotalCoins uint32
}, error) {
	return _NinjaChatLicense.Contract.UserLicenses(&_NinjaChatLicense.CallOpts, arg0)
}

// NinjaAddr is a free data retrieval call binding the contract method 0xbfe0f294.
//
// Solidity: function ninjaAddr() view returns(address)
func (_NinjaChatLicense *NinjaChatLicenseCaller) NinjaAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NinjaChatLicense.contract.Call(opts, &out, "ninjaAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NinjaAddr is a free data retrieval call binding the contract method 0xbfe0f294.
//
// Solidity: function ninjaAddr() view returns(address)
func (_NinjaChatLicense *NinjaChatLicenseSession) NinjaAddr() (common.Address, error) {
	return _NinjaChatLicense.Contract.NinjaAddr(&_NinjaChatLicense.CallOpts)
}

// NinjaAddr is a free data retrieval call binding the contract method 0xbfe0f294.
//
// Solidity: function ninjaAddr() view returns(address)
func (_NinjaChatLicense *NinjaChatLicenseCallerSession) NinjaAddr() (common.Address, error) {
	return _NinjaChatLicense.Contract.NinjaAddr(&_NinjaChatLicense.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NinjaChatLicense *NinjaChatLicenseCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NinjaChatLicense.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NinjaChatLicense *NinjaChatLicenseSession) Owner() (common.Address, error) {
	return _NinjaChatLicense.Contract.Owner(&_NinjaChatLicense.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NinjaChatLicense *NinjaChatLicenseCallerSession) Owner() (common.Address, error) {
	return _NinjaChatLicense.Contract.Owner(&_NinjaChatLicense.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_NinjaChatLicense *NinjaChatLicenseCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NinjaChatLicense.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_NinjaChatLicense *NinjaChatLicenseSession) Token() (common.Address, error) {
	return _NinjaChatLicense.Contract.Token(&_NinjaChatLicense.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_NinjaChatLicense *NinjaChatLicenseCallerSession) Token() (common.Address, error) {
	return _NinjaChatLicense.Contract.Token(&_NinjaChatLicense.CallOpts)
}

// BindLicense is a paid mutator transaction binding the contract method 0xf5cb29ac.
//
// Solidity: function BindLicense(address issueAddr, bytes32 recvAddr, bytes32 id, uint32 nDays, bytes signature) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactor) BindLicense(opts *bind.TransactOpts, issueAddr common.Address, recvAddr [32]byte, id [32]byte, nDays uint32, signature []byte) (*types.Transaction, error) {
	return _NinjaChatLicense.contract.Transact(opts, "BindLicense", issueAddr, recvAddr, id, nDays, signature)
}

// BindLicense is a paid mutator transaction binding the contract method 0xf5cb29ac.
//
// Solidity: function BindLicense(address issueAddr, bytes32 recvAddr, bytes32 id, uint32 nDays, bytes signature) returns()
func (_NinjaChatLicense *NinjaChatLicenseSession) BindLicense(issueAddr common.Address, recvAddr [32]byte, id [32]byte, nDays uint32, signature []byte) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.BindLicense(&_NinjaChatLicense.TransactOpts, issueAddr, recvAddr, id, nDays, signature)
}

// BindLicense is a paid mutator transaction binding the contract method 0xf5cb29ac.
//
// Solidity: function BindLicense(address issueAddr, bytes32 recvAddr, bytes32 id, uint32 nDays, bytes signature) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactorSession) BindLicense(issueAddr common.Address, recvAddr [32]byte, id [32]byte, nDays uint32, signature []byte) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.BindLicense(&_NinjaChatLicense.TransactOpts, issueAddr, recvAddr, id, nDays, signature)
}

// ChargeUser is a paid mutator transaction binding the contract method 0xa7656e03.
//
// Solidity: function ChargeUser(bytes32 userAddr, uint32 nDays) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactor) ChargeUser(opts *bind.TransactOpts, userAddr [32]byte, nDays uint32) (*types.Transaction, error) {
	return _NinjaChatLicense.contract.Transact(opts, "ChargeUser", userAddr, nDays)
}

// ChargeUser is a paid mutator transaction binding the contract method 0xa7656e03.
//
// Solidity: function ChargeUser(bytes32 userAddr, uint32 nDays) returns()
func (_NinjaChatLicense *NinjaChatLicenseSession) ChargeUser(userAddr [32]byte, nDays uint32) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.ChargeUser(&_NinjaChatLicense.TransactOpts, userAddr, nDays)
}

// ChargeUser is a paid mutator transaction binding the contract method 0xa7656e03.
//
// Solidity: function ChargeUser(bytes32 userAddr, uint32 nDays) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactorSession) ChargeUser(userAddr [32]byte, nDays uint32) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.ChargeUser(&_NinjaChatLicense.TransactOpts, userAddr, nDays)
}

// GenerateLicense is a paid mutator transaction binding the contract method 0x47795697.
//
// Solidity: function GenerateLicense(bytes32 id, uint32 nDays) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactor) GenerateLicense(opts *bind.TransactOpts, id [32]byte, nDays uint32) (*types.Transaction, error) {
	return _NinjaChatLicense.contract.Transact(opts, "GenerateLicense", id, nDays)
}

// GenerateLicense is a paid mutator transaction binding the contract method 0x47795697.
//
// Solidity: function GenerateLicense(bytes32 id, uint32 nDays) returns()
func (_NinjaChatLicense *NinjaChatLicenseSession) GenerateLicense(id [32]byte, nDays uint32) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.GenerateLicense(&_NinjaChatLicense.TransactOpts, id, nDays)
}

// GenerateLicense is a paid mutator transaction binding the contract method 0x47795697.
//
// Solidity: function GenerateLicense(bytes32 id, uint32 nDays) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactorSession) GenerateLicense(id [32]byte, nDays uint32) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.GenerateLicense(&_NinjaChatLicense.TransactOpts, id, nDays)
}

// SetNinjaAddr is a paid mutator transaction binding the contract method 0xddd57a33.
//
// Solidity: function SetNinjaAddr(address nAddr) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactor) SetNinjaAddr(opts *bind.TransactOpts, nAddr common.Address) (*types.Transaction, error) {
	return _NinjaChatLicense.contract.Transact(opts, "SetNinjaAddr", nAddr)
}

// SetNinjaAddr is a paid mutator transaction binding the contract method 0xddd57a33.
//
// Solidity: function SetNinjaAddr(address nAddr) returns()
func (_NinjaChatLicense *NinjaChatLicenseSession) SetNinjaAddr(nAddr common.Address) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.SetNinjaAddr(&_NinjaChatLicense.TransactOpts, nAddr)
}

// SetNinjaAddr is a paid mutator transaction binding the contract method 0xddd57a33.
//
// Solidity: function SetNinjaAddr(address nAddr) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactorSession) SetNinjaAddr(nAddr common.Address) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.SetNinjaAddr(&_NinjaChatLicense.TransactOpts, nAddr)
}

// SetTokenAddr is a paid mutator transaction binding the contract method 0x72e7a223.
//
// Solidity: function SetTokenAddr(address tAddr) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactor) SetTokenAddr(opts *bind.TransactOpts, tAddr common.Address) (*types.Transaction, error) {
	return _NinjaChatLicense.contract.Transact(opts, "SetTokenAddr", tAddr)
}

// SetTokenAddr is a paid mutator transaction binding the contract method 0x72e7a223.
//
// Solidity: function SetTokenAddr(address tAddr) returns()
func (_NinjaChatLicense *NinjaChatLicenseSession) SetTokenAddr(tAddr common.Address) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.SetTokenAddr(&_NinjaChatLicense.TransactOpts, tAddr)
}

// SetTokenAddr is a paid mutator transaction binding the contract method 0x72e7a223.
//
// Solidity: function SetTokenAddr(address tAddr) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactorSession) SetTokenAddr(tAddr common.Address) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.SetTokenAddr(&_NinjaChatLicense.TransactOpts, tAddr)
}

// Setting is a paid mutator transaction binding the contract method 0xc6e2a335.
//
// Solidity: function Setting(address tAddr, address nAddr) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactor) Setting(opts *bind.TransactOpts, tAddr common.Address, nAddr common.Address) (*types.Transaction, error) {
	return _NinjaChatLicense.contract.Transact(opts, "Setting", tAddr, nAddr)
}

// Setting is a paid mutator transaction binding the contract method 0xc6e2a335.
//
// Solidity: function Setting(address tAddr, address nAddr) returns()
func (_NinjaChatLicense *NinjaChatLicenseSession) Setting(tAddr common.Address, nAddr common.Address) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.Setting(&_NinjaChatLicense.TransactOpts, tAddr, nAddr)
}

// Setting is a paid mutator transaction binding the contract method 0xc6e2a335.
//
// Solidity: function Setting(address tAddr, address nAddr) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactorSession) Setting(tAddr common.Address, nAddr common.Address) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.Setting(&_NinjaChatLicense.TransactOpts, tAddr, nAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NinjaChatLicense.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NinjaChatLicense *NinjaChatLicenseSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.TransferOwnership(&_NinjaChatLicense.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NinjaChatLicense *NinjaChatLicenseTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NinjaChatLicense.Contract.TransferOwnership(&_NinjaChatLicense.TransactOpts, newOwner)
}

// NinjaChatLicenseBindLicenseEventIterator is returned from FilterBindLicenseEvent and is used to iterate over the raw logs and unpacked data for BindLicenseEvent events raised by the NinjaChatLicense contract.
type NinjaChatLicenseBindLicenseEventIterator struct {
	Event *NinjaChatLicenseBindLicenseEvent // Event containing the contract specifics and raw log

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
func (it *NinjaChatLicenseBindLicenseEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NinjaChatLicenseBindLicenseEvent)
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
		it.Event = new(NinjaChatLicenseBindLicenseEvent)
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
func (it *NinjaChatLicenseBindLicenseEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NinjaChatLicenseBindLicenseEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NinjaChatLicenseBindLicenseEvent represents a BindLicenseEvent event raised by the NinjaChatLicense contract.
type NinjaChatLicenseBindLicenseEvent struct {
	IssueAddr common.Address
	RecvAddr  [32]byte
	Id        [32]byte
	NDays     uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBindLicenseEvent is a free log retrieval operation binding the contract event 0x58e02a7e54f3f27722bcde4ed8ccc229b7700fa54d77f2412cd993ab808d5c6e.
//
// Solidity: event BindLicenseEvent(address indexed issueAddr, bytes32 recvAddr, bytes32 id, uint32 nDays)
func (_NinjaChatLicense *NinjaChatLicenseFilterer) FilterBindLicenseEvent(opts *bind.FilterOpts, issueAddr []common.Address) (*NinjaChatLicenseBindLicenseEventIterator, error) {

	var issueAddrRule []interface{}
	for _, issueAddrItem := range issueAddr {
		issueAddrRule = append(issueAddrRule, issueAddrItem)
	}

	logs, sub, err := _NinjaChatLicense.contract.FilterLogs(opts, "BindLicenseEvent", issueAddrRule)
	if err != nil {
		return nil, err
	}
	return &NinjaChatLicenseBindLicenseEventIterator{contract: _NinjaChatLicense.contract, event: "BindLicenseEvent", logs: logs, sub: sub}, nil
}

// WatchBindLicenseEvent is a free log subscription operation binding the contract event 0x58e02a7e54f3f27722bcde4ed8ccc229b7700fa54d77f2412cd993ab808d5c6e.
//
// Solidity: event BindLicenseEvent(address indexed issueAddr, bytes32 recvAddr, bytes32 id, uint32 nDays)
func (_NinjaChatLicense *NinjaChatLicenseFilterer) WatchBindLicenseEvent(opts *bind.WatchOpts, sink chan<- *NinjaChatLicenseBindLicenseEvent, issueAddr []common.Address) (event.Subscription, error) {

	var issueAddrRule []interface{}
	for _, issueAddrItem := range issueAddr {
		issueAddrRule = append(issueAddrRule, issueAddrItem)
	}

	logs, sub, err := _NinjaChatLicense.contract.WatchLogs(opts, "BindLicenseEvent", issueAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NinjaChatLicenseBindLicenseEvent)
				if err := _NinjaChatLicense.contract.UnpackLog(event, "BindLicenseEvent", log); err != nil {
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

// ParseBindLicenseEvent is a log parse operation binding the contract event 0x58e02a7e54f3f27722bcde4ed8ccc229b7700fa54d77f2412cd993ab808d5c6e.
//
// Solidity: event BindLicenseEvent(address indexed issueAddr, bytes32 recvAddr, bytes32 id, uint32 nDays)
func (_NinjaChatLicense *NinjaChatLicenseFilterer) ParseBindLicenseEvent(log types.Log) (*NinjaChatLicenseBindLicenseEvent, error) {
	event := new(NinjaChatLicenseBindLicenseEvent)
	if err := _NinjaChatLicense.contract.UnpackLog(event, "BindLicenseEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NinjaChatLicenseChargeUserEventIterator is returned from FilterChargeUserEvent and is used to iterate over the raw logs and unpacked data for ChargeUserEvent events raised by the NinjaChatLicense contract.
type NinjaChatLicenseChargeUserEventIterator struct {
	Event *NinjaChatLicenseChargeUserEvent // Event containing the contract specifics and raw log

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
func (it *NinjaChatLicenseChargeUserEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NinjaChatLicenseChargeUserEvent)
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
		it.Event = new(NinjaChatLicenseChargeUserEvent)
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
func (it *NinjaChatLicenseChargeUserEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NinjaChatLicenseChargeUserEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NinjaChatLicenseChargeUserEvent represents a ChargeUserEvent event raised by the NinjaChatLicense contract.
type NinjaChatLicenseChargeUserEvent struct {
	PayerAddr common.Address
	UserAddr  [32]byte
	NDays     uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterChargeUserEvent is a free log retrieval operation binding the contract event 0x7fa14ba5c04db63b34b318d9a4e3f8b8563edf544bd0d1734b9c74af7d3d4562.
//
// Solidity: event ChargeUserEvent(address indexed payerAddr, bytes32 userAddr, uint32 nDays)
func (_NinjaChatLicense *NinjaChatLicenseFilterer) FilterChargeUserEvent(opts *bind.FilterOpts, payerAddr []common.Address) (*NinjaChatLicenseChargeUserEventIterator, error) {

	var payerAddrRule []interface{}
	for _, payerAddrItem := range payerAddr {
		payerAddrRule = append(payerAddrRule, payerAddrItem)
	}

	logs, sub, err := _NinjaChatLicense.contract.FilterLogs(opts, "ChargeUserEvent", payerAddrRule)
	if err != nil {
		return nil, err
	}
	return &NinjaChatLicenseChargeUserEventIterator{contract: _NinjaChatLicense.contract, event: "ChargeUserEvent", logs: logs, sub: sub}, nil
}

// WatchChargeUserEvent is a free log subscription operation binding the contract event 0x7fa14ba5c04db63b34b318d9a4e3f8b8563edf544bd0d1734b9c74af7d3d4562.
//
// Solidity: event ChargeUserEvent(address indexed payerAddr, bytes32 userAddr, uint32 nDays)
func (_NinjaChatLicense *NinjaChatLicenseFilterer) WatchChargeUserEvent(opts *bind.WatchOpts, sink chan<- *NinjaChatLicenseChargeUserEvent, payerAddr []common.Address) (event.Subscription, error) {

	var payerAddrRule []interface{}
	for _, payerAddrItem := range payerAddr {
		payerAddrRule = append(payerAddrRule, payerAddrItem)
	}

	logs, sub, err := _NinjaChatLicense.contract.WatchLogs(opts, "ChargeUserEvent", payerAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NinjaChatLicenseChargeUserEvent)
				if err := _NinjaChatLicense.contract.UnpackLog(event, "ChargeUserEvent", log); err != nil {
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

// ParseChargeUserEvent is a log parse operation binding the contract event 0x7fa14ba5c04db63b34b318d9a4e3f8b8563edf544bd0d1734b9c74af7d3d4562.
//
// Solidity: event ChargeUserEvent(address indexed payerAddr, bytes32 userAddr, uint32 nDays)
func (_NinjaChatLicense *NinjaChatLicenseFilterer) ParseChargeUserEvent(log types.Log) (*NinjaChatLicenseChargeUserEvent, error) {
	event := new(NinjaChatLicenseChargeUserEvent)
	if err := _NinjaChatLicense.contract.UnpackLog(event, "ChargeUserEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NinjaChatLicenseGenerateLicenseEventIterator is returned from FilterGenerateLicenseEvent and is used to iterate over the raw logs and unpacked data for GenerateLicenseEvent events raised by the NinjaChatLicense contract.
type NinjaChatLicenseGenerateLicenseEventIterator struct {
	Event *NinjaChatLicenseGenerateLicenseEvent // Event containing the contract specifics and raw log

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
func (it *NinjaChatLicenseGenerateLicenseEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NinjaChatLicenseGenerateLicenseEvent)
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
		it.Event = new(NinjaChatLicenseGenerateLicenseEvent)
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
func (it *NinjaChatLicenseGenerateLicenseEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NinjaChatLicenseGenerateLicenseEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NinjaChatLicenseGenerateLicenseEvent represents a GenerateLicenseEvent event raised by the NinjaChatLicense contract.
type NinjaChatLicenseGenerateLicenseEvent struct {
	IssueAddr common.Address
	Id        [32]byte
	NDays     uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterGenerateLicenseEvent is a free log retrieval operation binding the contract event 0xbdc37090a1942c317384128913fe39098646601651d2392e9ae8683e3fae7afb.
//
// Solidity: event GenerateLicenseEvent(address indexed issueAddr, bytes32 id, uint32 nDays)
func (_NinjaChatLicense *NinjaChatLicenseFilterer) FilterGenerateLicenseEvent(opts *bind.FilterOpts, issueAddr []common.Address) (*NinjaChatLicenseGenerateLicenseEventIterator, error) {

	var issueAddrRule []interface{}
	for _, issueAddrItem := range issueAddr {
		issueAddrRule = append(issueAddrRule, issueAddrItem)
	}

	logs, sub, err := _NinjaChatLicense.contract.FilterLogs(opts, "GenerateLicenseEvent", issueAddrRule)
	if err != nil {
		return nil, err
	}
	return &NinjaChatLicenseGenerateLicenseEventIterator{contract: _NinjaChatLicense.contract, event: "GenerateLicenseEvent", logs: logs, sub: sub}, nil
}

// WatchGenerateLicenseEvent is a free log subscription operation binding the contract event 0xbdc37090a1942c317384128913fe39098646601651d2392e9ae8683e3fae7afb.
//
// Solidity: event GenerateLicenseEvent(address indexed issueAddr, bytes32 id, uint32 nDays)
func (_NinjaChatLicense *NinjaChatLicenseFilterer) WatchGenerateLicenseEvent(opts *bind.WatchOpts, sink chan<- *NinjaChatLicenseGenerateLicenseEvent, issueAddr []common.Address) (event.Subscription, error) {

	var issueAddrRule []interface{}
	for _, issueAddrItem := range issueAddr {
		issueAddrRule = append(issueAddrRule, issueAddrItem)
	}

	logs, sub, err := _NinjaChatLicense.contract.WatchLogs(opts, "GenerateLicenseEvent", issueAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NinjaChatLicenseGenerateLicenseEvent)
				if err := _NinjaChatLicense.contract.UnpackLog(event, "GenerateLicenseEvent", log); err != nil {
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

// ParseGenerateLicenseEvent is a log parse operation binding the contract event 0xbdc37090a1942c317384128913fe39098646601651d2392e9ae8683e3fae7afb.
//
// Solidity: event GenerateLicenseEvent(address indexed issueAddr, bytes32 id, uint32 nDays)
func (_NinjaChatLicense *NinjaChatLicenseFilterer) ParseGenerateLicenseEvent(log types.Log) (*NinjaChatLicenseGenerateLicenseEvent, error) {
	event := new(NinjaChatLicenseGenerateLicenseEvent)
	if err := _NinjaChatLicense.contract.UnpackLog(event, "GenerateLicenseEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
