// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package album

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

// AlbumAlbumInfo is an auto generated low-level Go binding around an user-defined struct.
type AlbumAlbumInfo struct {
	Price   *big.Int
	Creater string
	Title   string
}

// AlbumMetaData contains all meta data concerning the Album contract.
var AlbumMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"albumInfo\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"creater\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"int256\",\"name\":\"price\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"creater\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"}],\"internalType\":\"structAlbum.AlbumInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"_price\",\"type\":\"int256\"},{\"internalType\":\"string\",\"name\":\"_title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_creater\",\"type\":\"string\"}],\"name\":\"setInfo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610798806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c80631a3cd59a14610046578063419d7c821461006f578063d1b31f3b14610091575b600080fd5b610059610054366004610417565b6100a6565b6040516100669190610476565b60405180910390f35b61008261007d366004610417565b61022b565b604051610066939291906104c2565b6100a461009f36600461059a565b610371565b005b6100ca60405180606001604052806000815260200160608152602001606081525090565b600082815481106100dd576100dd610607565b9060005260206000209060030201604051806060016040529081600082015481526020016001820180546101109061061d565b80601f016020809104026020016040519081016040528092919081815260200182805461013c9061061d565b80156101895780601f1061015e57610100808354040283529160200191610189565b820191906000526020600020905b81548152906001019060200180831161016c57829003601f168201915b505050505081526020016002820180546101a29061061d565b80601f01602080910402602001604051908101604052809291908181526020018280546101ce9061061d565b801561021b5780601f106101f05761010080835404028352916020019161021b565b820191906000526020600020905b8154815290600101906020018083116101fe57829003601f168201915b5050505050815250509050919050565b6000818154811061023b57600080fd5b600091825260209091206003909102018054600182018054919350906102609061061d565b80601f016020809104026020016040519081016040528092919081815260200182805461028c9061061d565b80156102d95780601f106102ae576101008083540402835291602001916102d9565b820191906000526020600020905b8154815290600101906020018083116102bc57829003601f168201915b5050505050908060020180546102ee9061061d565b80601f016020809104026020016040519081016040528092919081815260200182805461031a9061061d565b80156103675780601f1061033c57610100808354040283529160200191610367565b820191906000526020600020905b81548152906001019060200180831161034a57829003601f168201915b5050505050905083565b604080516060810182528481526020810184815291810183905260008054600181018255908052815160039091027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56381019182559251919290917f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e564909101906103fa90826106a2565b506040820151600282019061040f90826106a2565b505050505050565b60006020828403121561042957600080fd5b5035919050565b6000815180845260005b818110156104565760208185018101518683018201520161043a565b506000602082860101526020601f19601f83011685010191505092915050565b6020815281516020820152600060208301516060604084015261049c6080840182610430565b90506040840151601f198483030160608501526104b98282610430565b95945050505050565b8381526060602082015260006104db6060830185610430565b82810360408401526104ed8185610430565b9695505050505050565b634e487b7160e01b600052604160045260246000fd5b600082601f83011261051e57600080fd5b813567ffffffffffffffff80821115610539576105396104f7565b604051601f8301601f19908116603f01168101908282118183101715610561576105616104f7565b8160405283815286602085880101111561057a57600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156105af57600080fd5b83359250602084013567ffffffffffffffff808211156105ce57600080fd5b6105da8783880161050d565b935060408601359150808211156105f057600080fd5b506105fd8682870161050d565b9150509250925092565b634e487b7160e01b600052603260045260246000fd5b600181811c9082168061063157607f821691505b60208210810361065157634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561069d57600081815260208120601f850160051c8101602086101561067e5750805b601f850160051c820191505b8181101561040f5782815560010161068a565b505050565b815167ffffffffffffffff8111156106bc576106bc6104f7565b6106d0816106ca845461061d565b84610657565b602080601f83116001811461070557600084156106ed5750858301515b600019600386901b1c1916600185901b17855561040f565b600085815260208120601f198616915b8281101561073457888601518255948401946001909101908401610715565b50858210156107525787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea26469706673582212200735f5d005d8f0b39431cb17635896f973fd8a9f3c0a6580505a0cada4ed43fb64736f6c63430008110033",
}

// AlbumABI is the input ABI used to generate the binding from.
// Deprecated: Use AlbumMetaData.ABI instead.
var AlbumABI = AlbumMetaData.ABI

// AlbumBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AlbumMetaData.Bin instead.
var AlbumBin = AlbumMetaData.Bin

// DeployAlbum deploys a new Ethereum contract, binding an instance of Album to it.
func DeployAlbum(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Album, error) {
	parsed, err := AlbumMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AlbumBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Album{AlbumCaller: AlbumCaller{contract: contract}, AlbumTransactor: AlbumTransactor{contract: contract}, AlbumFilterer: AlbumFilterer{contract: contract}}, nil
}

// Album is an auto generated Go binding around an Ethereum contract.
type Album struct {
	AlbumCaller     // Read-only binding to the contract
	AlbumTransactor // Write-only binding to the contract
	AlbumFilterer   // Log filterer for contract events
}

// AlbumCaller is an auto generated read-only Go binding around an Ethereum contract.
type AlbumCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlbumTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AlbumTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlbumFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AlbumFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlbumSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AlbumSession struct {
	Contract     *Album            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlbumCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AlbumCallerSession struct {
	Contract *AlbumCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// AlbumTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AlbumTransactorSession struct {
	Contract     *AlbumTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlbumRaw is an auto generated low-level Go binding around an Ethereum contract.
type AlbumRaw struct {
	Contract *Album // Generic contract binding to access the raw methods on
}

// AlbumCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AlbumCallerRaw struct {
	Contract *AlbumCaller // Generic read-only contract binding to access the raw methods on
}

// AlbumTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AlbumTransactorRaw struct {
	Contract *AlbumTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAlbum creates a new instance of Album, bound to a specific deployed contract.
func NewAlbum(address common.Address, backend bind.ContractBackend) (*Album, error) {
	contract, err := bindAlbum(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Album{AlbumCaller: AlbumCaller{contract: contract}, AlbumTransactor: AlbumTransactor{contract: contract}, AlbumFilterer: AlbumFilterer{contract: contract}}, nil
}

// NewAlbumCaller creates a new read-only instance of Album, bound to a specific deployed contract.
func NewAlbumCaller(address common.Address, caller bind.ContractCaller) (*AlbumCaller, error) {
	contract, err := bindAlbum(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AlbumCaller{contract: contract}, nil
}

// NewAlbumTransactor creates a new write-only instance of Album, bound to a specific deployed contract.
func NewAlbumTransactor(address common.Address, transactor bind.ContractTransactor) (*AlbumTransactor, error) {
	contract, err := bindAlbum(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AlbumTransactor{contract: contract}, nil
}

// NewAlbumFilterer creates a new log filterer instance of Album, bound to a specific deployed contract.
func NewAlbumFilterer(address common.Address, filterer bind.ContractFilterer) (*AlbumFilterer, error) {
	contract, err := bindAlbum(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AlbumFilterer{contract: contract}, nil
}

// bindAlbum binds a generic wrapper to an already deployed contract.
func bindAlbum(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AlbumABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Album *AlbumRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Album.Contract.AlbumCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Album *AlbumRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Album.Contract.AlbumTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Album *AlbumRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Album.Contract.AlbumTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Album *AlbumCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Album.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Album *AlbumTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Album.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Album *AlbumTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Album.Contract.contract.Transact(opts, method, params...)
}

// AlbumInfo is a free data retrieval call binding the contract method 0x419d7c82.
//
// Solidity: function albumInfo(uint256 ) view returns(int256 price, string creater, string title)
func (_Album *AlbumCaller) AlbumInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Price   *big.Int
	Creater string
	Title   string
}, error) {
	var out []interface{}
	err := _Album.contract.Call(opts, &out, "albumInfo", arg0)

	outstruct := new(struct {
		Price   *big.Int
		Creater string
		Title   string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Price = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Creater = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Title = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// AlbumInfo is a free data retrieval call binding the contract method 0x419d7c82.
//
// Solidity: function albumInfo(uint256 ) view returns(int256 price, string creater, string title)
func (_Album *AlbumSession) AlbumInfo(arg0 *big.Int) (struct {
	Price   *big.Int
	Creater string
	Title   string
}, error) {
	return _Album.Contract.AlbumInfo(&_Album.CallOpts, arg0)
}

// AlbumInfo is a free data retrieval call binding the contract method 0x419d7c82.
//
// Solidity: function albumInfo(uint256 ) view returns(int256 price, string creater, string title)
func (_Album *AlbumCallerSession) AlbumInfo(arg0 *big.Int) (struct {
	Price   *big.Int
	Creater string
	Title   string
}, error) {
	return _Album.Contract.AlbumInfo(&_Album.CallOpts, arg0)
}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 index) view returns((int256,string,string))
func (_Album *AlbumCaller) GetInfo(opts *bind.CallOpts, index *big.Int) (AlbumAlbumInfo, error) {
	var out []interface{}
	err := _Album.contract.Call(opts, &out, "getInfo", index)

	if err != nil {
		return *new(AlbumAlbumInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(AlbumAlbumInfo)).(*AlbumAlbumInfo)

	return out0, err

}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 index) view returns((int256,string,string))
func (_Album *AlbumSession) GetInfo(index *big.Int) (AlbumAlbumInfo, error) {
	return _Album.Contract.GetInfo(&_Album.CallOpts, index)
}

// GetInfo is a free data retrieval call binding the contract method 0x1a3cd59a.
//
// Solidity: function getInfo(uint256 index) view returns((int256,string,string))
func (_Album *AlbumCallerSession) GetInfo(index *big.Int) (AlbumAlbumInfo, error) {
	return _Album.Contract.GetInfo(&_Album.CallOpts, index)
}

// SetInfo is a paid mutator transaction binding the contract method 0xd1b31f3b.
//
// Solidity: function setInfo(int256 _price, string _title, string _creater) returns()
func (_Album *AlbumTransactor) SetInfo(opts *bind.TransactOpts, _price *big.Int, _title string, _creater string) (*types.Transaction, error) {
	return _Album.contract.Transact(opts, "setInfo", _price, _title, _creater)
}

// SetInfo is a paid mutator transaction binding the contract method 0xd1b31f3b.
//
// Solidity: function setInfo(int256 _price, string _title, string _creater) returns()
func (_Album *AlbumSession) SetInfo(_price *big.Int, _title string, _creater string) (*types.Transaction, error) {
	return _Album.Contract.SetInfo(&_Album.TransactOpts, _price, _title, _creater)
}

// SetInfo is a paid mutator transaction binding the contract method 0xd1b31f3b.
//
// Solidity: function setInfo(int256 _price, string _title, string _creater) returns()
func (_Album *AlbumTransactorSession) SetInfo(_price *big.Int, _title string, _creater string) (*types.Transaction, error) {
	return _Album.Contract.SetInfo(&_Album.TransactOpts, _price, _title, _creater)
}
