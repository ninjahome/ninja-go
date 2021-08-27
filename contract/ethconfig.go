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

// NinjaConfigABI is the input ABI used to generate the binding from.
const NinjaConfigABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"ipAddr\",\"type\":\"bytes4\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"port1\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"port2\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"port3\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"port4\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"port5\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"port6\",\"type\":\"uint16\"}],\"name\":\"AddBootsTrapEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"ipAddr\",\"type\":\"bytes4\"}],\"name\":\"DelBootsTrapEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"ipAddr\",\"type\":\"bytes4\"},{\"internalType\":\"uint16\",\"name\":\"port1\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"port2\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"port3\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"port4\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"port5\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"port6\",\"type\":\"uint16\"}],\"name\":\"AddBootsTrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"ipAddr\",\"type\":\"bytes4\"}],\"name\":\"DelBootsTrap\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"ipAddr\",\"type\":\"bytes4\"}],\"name\":\"GetIPPort\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"idx\",\"type\":\"uint256\"}],\"name\":\"GetIPPortByIdx\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetIpAddrList\",\"outputs\":[{\"internalType\":\"bytes4[]\",\"name\":\"\",\"type\":\"bytes4[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetLicenseConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"tAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lcAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"aAddr\",\"type\":\"bytes\"}],\"name\":\"LicenseConfigSet\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"boots\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"ipAddr\",\"type\":\"bytes4\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lCfg\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"tokenAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"licenseContractAddr\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"accessAddr\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NinjaConfig is an auto generated Go binding around an Ethereum contract.
type NinjaConfig struct {
	NinjaConfigCaller     // Read-only binding to the contract
	NinjaConfigTransactor // Write-only binding to the contract
	NinjaConfigFilterer   // Log filterer for contract events
}

// NinjaConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type NinjaConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NinjaConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NinjaConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NinjaConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NinjaConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NinjaConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NinjaConfigSession struct {
	Contract     *NinjaConfig      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NinjaConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NinjaConfigCallerSession struct {
	Contract *NinjaConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// NinjaConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NinjaConfigTransactorSession struct {
	Contract     *NinjaConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// NinjaConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type NinjaConfigRaw struct {
	Contract *NinjaConfig // Generic contract binding to access the raw methods on
}

// NinjaConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NinjaConfigCallerRaw struct {
	Contract *NinjaConfigCaller // Generic read-only contract binding to access the raw methods on
}

// NinjaConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NinjaConfigTransactorRaw struct {
	Contract *NinjaConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNinjaConfig creates a new instance of NinjaConfig, bound to a specific deployed contract.
func NewNinjaConfig(address common.Address, backend bind.ContractBackend) (*NinjaConfig, error) {
	contract, err := bindNinjaConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NinjaConfig{NinjaConfigCaller: NinjaConfigCaller{contract: contract}, NinjaConfigTransactor: NinjaConfigTransactor{contract: contract}, NinjaConfigFilterer: NinjaConfigFilterer{contract: contract}}, nil
}

// NewNinjaConfigCaller creates a new read-only instance of NinjaConfig, bound to a specific deployed contract.
func NewNinjaConfigCaller(address common.Address, caller bind.ContractCaller) (*NinjaConfigCaller, error) {
	contract, err := bindNinjaConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NinjaConfigCaller{contract: contract}, nil
}

// NewNinjaConfigTransactor creates a new write-only instance of NinjaConfig, bound to a specific deployed contract.
func NewNinjaConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*NinjaConfigTransactor, error) {
	contract, err := bindNinjaConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NinjaConfigTransactor{contract: contract}, nil
}

// NewNinjaConfigFilterer creates a new log filterer instance of NinjaConfig, bound to a specific deployed contract.
func NewNinjaConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*NinjaConfigFilterer, error) {
	contract, err := bindNinjaConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NinjaConfigFilterer{contract: contract}, nil
}

// bindNinjaConfig binds a generic wrapper to an already deployed contract.
func bindNinjaConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NinjaConfigABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NinjaConfig *NinjaConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NinjaConfig.Contract.NinjaConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NinjaConfig *NinjaConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NinjaConfig.Contract.NinjaConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NinjaConfig *NinjaConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NinjaConfig.Contract.NinjaConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NinjaConfig *NinjaConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NinjaConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NinjaConfig *NinjaConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NinjaConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NinjaConfig *NinjaConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NinjaConfig.Contract.contract.Transact(opts, method, params...)
}

// GetIPPort is a free data retrieval call binding the contract method 0x2724a2b5.
//
// Solidity: function GetIPPort(bytes4 ipAddr) view returns(uint16, uint16, uint16, uint16, uint16, uint16)
func (_NinjaConfig *NinjaConfigCaller) GetIPPort(opts *bind.CallOpts, ipAddr [4]byte) (uint16, uint16, uint16, uint16, uint16, uint16, error) {
	var out []interface{}
	err := _NinjaConfig.contract.Call(opts, &out, "GetIPPort", ipAddr)

	if err != nil {
		return *new(uint16), *new(uint16), *new(uint16), *new(uint16), *new(uint16), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)
	out2 := *abi.ConvertType(out[2], new(uint16)).(*uint16)
	out3 := *abi.ConvertType(out[3], new(uint16)).(*uint16)
	out4 := *abi.ConvertType(out[4], new(uint16)).(*uint16)
	out5 := *abi.ConvertType(out[5], new(uint16)).(*uint16)

	return out0, out1, out2, out3, out4, out5, err

}

// GetIPPort is a free data retrieval call binding the contract method 0x2724a2b5.
//
// Solidity: function GetIPPort(bytes4 ipAddr) view returns(uint16, uint16, uint16, uint16, uint16, uint16)
func (_NinjaConfig *NinjaConfigSession) GetIPPort(ipAddr [4]byte) (uint16, uint16, uint16, uint16, uint16, uint16, error) {
	return _NinjaConfig.Contract.GetIPPort(&_NinjaConfig.CallOpts, ipAddr)
}

// GetIPPort is a free data retrieval call binding the contract method 0x2724a2b5.
//
// Solidity: function GetIPPort(bytes4 ipAddr) view returns(uint16, uint16, uint16, uint16, uint16, uint16)
func (_NinjaConfig *NinjaConfigCallerSession) GetIPPort(ipAddr [4]byte) (uint16, uint16, uint16, uint16, uint16, uint16, error) {
	return _NinjaConfig.Contract.GetIPPort(&_NinjaConfig.CallOpts, ipAddr)
}

// GetIPPortByIdx is a free data retrieval call binding the contract method 0x1148ce0d.
//
// Solidity: function GetIPPortByIdx(uint256 idx) view returns(uint16, uint16, uint16, uint16, uint16, uint16)
func (_NinjaConfig *NinjaConfigCaller) GetIPPortByIdx(opts *bind.CallOpts, idx *big.Int) (uint16, uint16, uint16, uint16, uint16, uint16, error) {
	var out []interface{}
	err := _NinjaConfig.contract.Call(opts, &out, "GetIPPortByIdx", idx)

	if err != nil {
		return *new(uint16), *new(uint16), *new(uint16), *new(uint16), *new(uint16), *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)
	out1 := *abi.ConvertType(out[1], new(uint16)).(*uint16)
	out2 := *abi.ConvertType(out[2], new(uint16)).(*uint16)
	out3 := *abi.ConvertType(out[3], new(uint16)).(*uint16)
	out4 := *abi.ConvertType(out[4], new(uint16)).(*uint16)
	out5 := *abi.ConvertType(out[5], new(uint16)).(*uint16)

	return out0, out1, out2, out3, out4, out5, err

}

// GetIPPortByIdx is a free data retrieval call binding the contract method 0x1148ce0d.
//
// Solidity: function GetIPPortByIdx(uint256 idx) view returns(uint16, uint16, uint16, uint16, uint16, uint16)
func (_NinjaConfig *NinjaConfigSession) GetIPPortByIdx(idx *big.Int) (uint16, uint16, uint16, uint16, uint16, uint16, error) {
	return _NinjaConfig.Contract.GetIPPortByIdx(&_NinjaConfig.CallOpts, idx)
}

// GetIPPortByIdx is a free data retrieval call binding the contract method 0x1148ce0d.
//
// Solidity: function GetIPPortByIdx(uint256 idx) view returns(uint16, uint16, uint16, uint16, uint16, uint16)
func (_NinjaConfig *NinjaConfigCallerSession) GetIPPortByIdx(idx *big.Int) (uint16, uint16, uint16, uint16, uint16, uint16, error) {
	return _NinjaConfig.Contract.GetIPPortByIdx(&_NinjaConfig.CallOpts, idx)
}

// GetIpAddrList is a free data retrieval call binding the contract method 0xe18ebf3e.
//
// Solidity: function GetIpAddrList() view returns(bytes4[])
func (_NinjaConfig *NinjaConfigCaller) GetIpAddrList(opts *bind.CallOpts) ([][4]byte, error) {
	var out []interface{}
	err := _NinjaConfig.contract.Call(opts, &out, "GetIpAddrList")

	if err != nil {
		return *new([][4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][4]byte)).(*[][4]byte)

	return out0, err

}

// GetIpAddrList is a free data retrieval call binding the contract method 0xe18ebf3e.
//
// Solidity: function GetIpAddrList() view returns(bytes4[])
func (_NinjaConfig *NinjaConfigSession) GetIpAddrList() ([][4]byte, error) {
	return _NinjaConfig.Contract.GetIpAddrList(&_NinjaConfig.CallOpts)
}

// GetIpAddrList is a free data retrieval call binding the contract method 0xe18ebf3e.
//
// Solidity: function GetIpAddrList() view returns(bytes4[])
func (_NinjaConfig *NinjaConfigCallerSession) GetIpAddrList() ([][4]byte, error) {
	return _NinjaConfig.Contract.GetIpAddrList(&_NinjaConfig.CallOpts)
}

// GetLicenseConfig is a free data retrieval call binding the contract method 0x5424b4f0.
//
// Solidity: function GetLicenseConfig() view returns(address, address, bytes)
func (_NinjaConfig *NinjaConfigCaller) GetLicenseConfig(opts *bind.CallOpts) (common.Address, common.Address, []byte, error) {
	var out []interface{}
	err := _NinjaConfig.contract.Call(opts, &out, "GetLicenseConfig")

	if err != nil {
		return *new(common.Address), *new(common.Address), *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	out2 := *abi.ConvertType(out[2], new([]byte)).(*[]byte)

	return out0, out1, out2, err

}

// GetLicenseConfig is a free data retrieval call binding the contract method 0x5424b4f0.
//
// Solidity: function GetLicenseConfig() view returns(address, address, bytes)
func (_NinjaConfig *NinjaConfigSession) GetLicenseConfig() (common.Address, common.Address, []byte, error) {
	return _NinjaConfig.Contract.GetLicenseConfig(&_NinjaConfig.CallOpts)
}

// GetLicenseConfig is a free data retrieval call binding the contract method 0x5424b4f0.
//
// Solidity: function GetLicenseConfig() view returns(address, address, bytes)
func (_NinjaConfig *NinjaConfigCallerSession) GetLicenseConfig() (common.Address, common.Address, []byte, error) {
	return _NinjaConfig.Contract.GetLicenseConfig(&_NinjaConfig.CallOpts)
}

// Boots is a free data retrieval call binding the contract method 0x05449707.
//
// Solidity: function boots(uint256 ) view returns(bytes4 ipAddr)
func (_NinjaConfig *NinjaConfigCaller) Boots(opts *bind.CallOpts, arg0 *big.Int) ([4]byte, error) {
	var out []interface{}
	err := _NinjaConfig.contract.Call(opts, &out, "boots", arg0)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// Boots is a free data retrieval call binding the contract method 0x05449707.
//
// Solidity: function boots(uint256 ) view returns(bytes4 ipAddr)
func (_NinjaConfig *NinjaConfigSession) Boots(arg0 *big.Int) ([4]byte, error) {
	return _NinjaConfig.Contract.Boots(&_NinjaConfig.CallOpts, arg0)
}

// Boots is a free data retrieval call binding the contract method 0x05449707.
//
// Solidity: function boots(uint256 ) view returns(bytes4 ipAddr)
func (_NinjaConfig *NinjaConfigCallerSession) Boots(arg0 *big.Int) ([4]byte, error) {
	return _NinjaConfig.Contract.Boots(&_NinjaConfig.CallOpts, arg0)
}

// LCfg is a free data retrieval call binding the contract method 0x75e9ded7.
//
// Solidity: function lCfg() view returns(address tokenAddr, address licenseContractAddr, bytes accessAddr)
func (_NinjaConfig *NinjaConfigCaller) LCfg(opts *bind.CallOpts) (struct {
	TokenAddr           common.Address
	LicenseContractAddr common.Address
	AccessAddr          []byte
}, error) {
	var out []interface{}
	err := _NinjaConfig.contract.Call(opts, &out, "lCfg")

	outstruct := new(struct {
		TokenAddr           common.Address
		LicenseContractAddr common.Address
		AccessAddr          []byte
	})

	outstruct.TokenAddr = out[0].(common.Address)
	outstruct.LicenseContractAddr = out[1].(common.Address)
	outstruct.AccessAddr = out[2].([]byte)

	return *outstruct, err

}

// LCfg is a free data retrieval call binding the contract method 0x75e9ded7.
//
// Solidity: function lCfg() view returns(address tokenAddr, address licenseContractAddr, bytes accessAddr)
func (_NinjaConfig *NinjaConfigSession) LCfg() (struct {
	TokenAddr           common.Address
	LicenseContractAddr common.Address
	AccessAddr          []byte
}, error) {
	return _NinjaConfig.Contract.LCfg(&_NinjaConfig.CallOpts)
}

// LCfg is a free data retrieval call binding the contract method 0x75e9ded7.
//
// Solidity: function lCfg() view returns(address tokenAddr, address licenseContractAddr, bytes accessAddr)
func (_NinjaConfig *NinjaConfigCallerSession) LCfg() (struct {
	TokenAddr           common.Address
	LicenseContractAddr common.Address
	AccessAddr          []byte
}, error) {
	return _NinjaConfig.Contract.LCfg(&_NinjaConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NinjaConfig *NinjaConfigCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NinjaConfig.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NinjaConfig *NinjaConfigSession) Owner() (common.Address, error) {
	return _NinjaConfig.Contract.Owner(&_NinjaConfig.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NinjaConfig *NinjaConfigCallerSession) Owner() (common.Address, error) {
	return _NinjaConfig.Contract.Owner(&_NinjaConfig.CallOpts)
}

// AddBootsTrap is a paid mutator transaction binding the contract method 0xcb60456a.
//
// Solidity: function AddBootsTrap(bytes4 ipAddr, uint16 port1, uint16 port2, uint16 port3, uint16 port4, uint16 port5, uint16 port6) returns()
func (_NinjaConfig *NinjaConfigTransactor) AddBootsTrap(opts *bind.TransactOpts, ipAddr [4]byte, port1 uint16, port2 uint16, port3 uint16, port4 uint16, port5 uint16, port6 uint16) (*types.Transaction, error) {
	return _NinjaConfig.contract.Transact(opts, "AddBootsTrap", ipAddr, port1, port2, port3, port4, port5, port6)
}

// AddBootsTrap is a paid mutator transaction binding the contract method 0xcb60456a.
//
// Solidity: function AddBootsTrap(bytes4 ipAddr, uint16 port1, uint16 port2, uint16 port3, uint16 port4, uint16 port5, uint16 port6) returns()
func (_NinjaConfig *NinjaConfigSession) AddBootsTrap(ipAddr [4]byte, port1 uint16, port2 uint16, port3 uint16, port4 uint16, port5 uint16, port6 uint16) (*types.Transaction, error) {
	return _NinjaConfig.Contract.AddBootsTrap(&_NinjaConfig.TransactOpts, ipAddr, port1, port2, port3, port4, port5, port6)
}

// AddBootsTrap is a paid mutator transaction binding the contract method 0xcb60456a.
//
// Solidity: function AddBootsTrap(bytes4 ipAddr, uint16 port1, uint16 port2, uint16 port3, uint16 port4, uint16 port5, uint16 port6) returns()
func (_NinjaConfig *NinjaConfigTransactorSession) AddBootsTrap(ipAddr [4]byte, port1 uint16, port2 uint16, port3 uint16, port4 uint16, port5 uint16, port6 uint16) (*types.Transaction, error) {
	return _NinjaConfig.Contract.AddBootsTrap(&_NinjaConfig.TransactOpts, ipAddr, port1, port2, port3, port4, port5, port6)
}

// DelBootsTrap is a paid mutator transaction binding the contract method 0x1c4d4a66.
//
// Solidity: function DelBootsTrap(bytes4 ipAddr) returns()
func (_NinjaConfig *NinjaConfigTransactor) DelBootsTrap(opts *bind.TransactOpts, ipAddr [4]byte) (*types.Transaction, error) {
	return _NinjaConfig.contract.Transact(opts, "DelBootsTrap", ipAddr)
}

// DelBootsTrap is a paid mutator transaction binding the contract method 0x1c4d4a66.
//
// Solidity: function DelBootsTrap(bytes4 ipAddr) returns()
func (_NinjaConfig *NinjaConfigSession) DelBootsTrap(ipAddr [4]byte) (*types.Transaction, error) {
	return _NinjaConfig.Contract.DelBootsTrap(&_NinjaConfig.TransactOpts, ipAddr)
}

// DelBootsTrap is a paid mutator transaction binding the contract method 0x1c4d4a66.
//
// Solidity: function DelBootsTrap(bytes4 ipAddr) returns()
func (_NinjaConfig *NinjaConfigTransactorSession) DelBootsTrap(ipAddr [4]byte) (*types.Transaction, error) {
	return _NinjaConfig.Contract.DelBootsTrap(&_NinjaConfig.TransactOpts, ipAddr)
}

// LicenseConfigSet is a paid mutator transaction binding the contract method 0x56bfc10f.
//
// Solidity: function LicenseConfigSet(address tAddr, address lcAddr, bytes aAddr) returns()
func (_NinjaConfig *NinjaConfigTransactor) LicenseConfigSet(opts *bind.TransactOpts, tAddr common.Address, lcAddr common.Address, aAddr []byte) (*types.Transaction, error) {
	return _NinjaConfig.contract.Transact(opts, "LicenseConfigSet", tAddr, lcAddr, aAddr)
}

// LicenseConfigSet is a paid mutator transaction binding the contract method 0x56bfc10f.
//
// Solidity: function LicenseConfigSet(address tAddr, address lcAddr, bytes aAddr) returns()
func (_NinjaConfig *NinjaConfigSession) LicenseConfigSet(tAddr common.Address, lcAddr common.Address, aAddr []byte) (*types.Transaction, error) {
	return _NinjaConfig.Contract.LicenseConfigSet(&_NinjaConfig.TransactOpts, tAddr, lcAddr, aAddr)
}

// LicenseConfigSet is a paid mutator transaction binding the contract method 0x56bfc10f.
//
// Solidity: function LicenseConfigSet(address tAddr, address lcAddr, bytes aAddr) returns()
func (_NinjaConfig *NinjaConfigTransactorSession) LicenseConfigSet(tAddr common.Address, lcAddr common.Address, aAddr []byte) (*types.Transaction, error) {
	return _NinjaConfig.Contract.LicenseConfigSet(&_NinjaConfig.TransactOpts, tAddr, lcAddr, aAddr)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NinjaConfig *NinjaConfigTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NinjaConfig.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NinjaConfig *NinjaConfigSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NinjaConfig.Contract.TransferOwnership(&_NinjaConfig.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NinjaConfig *NinjaConfigTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NinjaConfig.Contract.TransferOwnership(&_NinjaConfig.TransactOpts, newOwner)
}

// NinjaConfigAddBootsTrapEventIterator is returned from FilterAddBootsTrapEvent and is used to iterate over the raw logs and unpacked data for AddBootsTrapEvent events raised by the NinjaConfig contract.
type NinjaConfigAddBootsTrapEventIterator struct {
	Event *NinjaConfigAddBootsTrapEvent // Event containing the contract specifics and raw log

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
func (it *NinjaConfigAddBootsTrapEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NinjaConfigAddBootsTrapEvent)
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
		it.Event = new(NinjaConfigAddBootsTrapEvent)
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
func (it *NinjaConfigAddBootsTrapEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NinjaConfigAddBootsTrapEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NinjaConfigAddBootsTrapEvent represents a AddBootsTrapEvent event raised by the NinjaConfig contract.
type NinjaConfigAddBootsTrapEvent struct {
	IpAddr [4]byte
	Port1  uint16
	Port2  uint16
	Port3  uint16
	Port4  uint16
	Port5  uint16
	Port6  uint16
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterAddBootsTrapEvent is a free log retrieval operation binding the contract event 0xef3183821376fad2c6e6a67caa77af06331905c7d97d27759a410a0169fe26da.
//
// Solidity: event AddBootsTrapEvent(bytes4 ipAddr, uint16 port1, uint16 port2, uint16 port3, uint16 port4, uint16 port5, uint16 port6)
func (_NinjaConfig *NinjaConfigFilterer) FilterAddBootsTrapEvent(opts *bind.FilterOpts) (*NinjaConfigAddBootsTrapEventIterator, error) {

	logs, sub, err := _NinjaConfig.contract.FilterLogs(opts, "AddBootsTrapEvent")
	if err != nil {
		return nil, err
	}
	return &NinjaConfigAddBootsTrapEventIterator{contract: _NinjaConfig.contract, event: "AddBootsTrapEvent", logs: logs, sub: sub}, nil
}

// WatchAddBootsTrapEvent is a free log subscription operation binding the contract event 0xef3183821376fad2c6e6a67caa77af06331905c7d97d27759a410a0169fe26da.
//
// Solidity: event AddBootsTrapEvent(bytes4 ipAddr, uint16 port1, uint16 port2, uint16 port3, uint16 port4, uint16 port5, uint16 port6)
func (_NinjaConfig *NinjaConfigFilterer) WatchAddBootsTrapEvent(opts *bind.WatchOpts, sink chan<- *NinjaConfigAddBootsTrapEvent) (event.Subscription, error) {

	logs, sub, err := _NinjaConfig.contract.WatchLogs(opts, "AddBootsTrapEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NinjaConfigAddBootsTrapEvent)
				if err := _NinjaConfig.contract.UnpackLog(event, "AddBootsTrapEvent", log); err != nil {
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

// ParseAddBootsTrapEvent is a log parse operation binding the contract event 0xef3183821376fad2c6e6a67caa77af06331905c7d97d27759a410a0169fe26da.
//
// Solidity: event AddBootsTrapEvent(bytes4 ipAddr, uint16 port1, uint16 port2, uint16 port3, uint16 port4, uint16 port5, uint16 port6)
func (_NinjaConfig *NinjaConfigFilterer) ParseAddBootsTrapEvent(log types.Log) (*NinjaConfigAddBootsTrapEvent, error) {
	event := new(NinjaConfigAddBootsTrapEvent)
	if err := _NinjaConfig.contract.UnpackLog(event, "AddBootsTrapEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NinjaConfigDelBootsTrapEventIterator is returned from FilterDelBootsTrapEvent and is used to iterate over the raw logs and unpacked data for DelBootsTrapEvent events raised by the NinjaConfig contract.
type NinjaConfigDelBootsTrapEventIterator struct {
	Event *NinjaConfigDelBootsTrapEvent // Event containing the contract specifics and raw log

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
func (it *NinjaConfigDelBootsTrapEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NinjaConfigDelBootsTrapEvent)
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
		it.Event = new(NinjaConfigDelBootsTrapEvent)
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
func (it *NinjaConfigDelBootsTrapEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NinjaConfigDelBootsTrapEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NinjaConfigDelBootsTrapEvent represents a DelBootsTrapEvent event raised by the NinjaConfig contract.
type NinjaConfigDelBootsTrapEvent struct {
	IpAddr [4]byte
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDelBootsTrapEvent is a free log retrieval operation binding the contract event 0x38e6ef8d400448b2bde71df2fe14bad3b93fe34e2cdcc12573e2620c5cd90156.
//
// Solidity: event DelBootsTrapEvent(bytes4 ipAddr)
func (_NinjaConfig *NinjaConfigFilterer) FilterDelBootsTrapEvent(opts *bind.FilterOpts) (*NinjaConfigDelBootsTrapEventIterator, error) {

	logs, sub, err := _NinjaConfig.contract.FilterLogs(opts, "DelBootsTrapEvent")
	if err != nil {
		return nil, err
	}
	return &NinjaConfigDelBootsTrapEventIterator{contract: _NinjaConfig.contract, event: "DelBootsTrapEvent", logs: logs, sub: sub}, nil
}

// WatchDelBootsTrapEvent is a free log subscription operation binding the contract event 0x38e6ef8d400448b2bde71df2fe14bad3b93fe34e2cdcc12573e2620c5cd90156.
//
// Solidity: event DelBootsTrapEvent(bytes4 ipAddr)
func (_NinjaConfig *NinjaConfigFilterer) WatchDelBootsTrapEvent(opts *bind.WatchOpts, sink chan<- *NinjaConfigDelBootsTrapEvent) (event.Subscription, error) {

	logs, sub, err := _NinjaConfig.contract.WatchLogs(opts, "DelBootsTrapEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NinjaConfigDelBootsTrapEvent)
				if err := _NinjaConfig.contract.UnpackLog(event, "DelBootsTrapEvent", log); err != nil {
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

// ParseDelBootsTrapEvent is a log parse operation binding the contract event 0x38e6ef8d400448b2bde71df2fe14bad3b93fe34e2cdcc12573e2620c5cd90156.
//
// Solidity: event DelBootsTrapEvent(bytes4 ipAddr)
func (_NinjaConfig *NinjaConfigFilterer) ParseDelBootsTrapEvent(log types.Log) (*NinjaConfigDelBootsTrapEvent, error) {
	event := new(NinjaConfigDelBootsTrapEvent)
	if err := _NinjaConfig.contract.UnpackLog(event, "DelBootsTrapEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
