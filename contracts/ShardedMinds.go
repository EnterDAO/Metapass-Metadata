// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// ShardedMindsABI is the input ABI used to generate the binding from.
const ShardedMindsABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"},{\"internalType\":\"addresspayable\",\"name\":\"_daoAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_shardedMindsPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_bulkBuyLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxNFTsPerWallet\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxNFTsPerWalletPresale\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_presaleStart\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_officialSaleStart\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseURI\",\"type\":\"string\"}],\"name\":\"BaseURIChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newBulkBuyLimit\",\"type\":\"uint256\"}],\"name\":\"BulkBuyLimitChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newMaxSupply\",\"type\":\"uint256\"}],\"name\":\"MaxSupplyChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newShardedMindsPrice\",\"type\":\"uint256\"}],\"name\":\"ShardedMindsPriceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGene\",\"type\":\"uint256\"}],\"name\":\"TokenMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldGene\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newGene\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumShardedMinds.ShardedMindsEventType\",\"name\":\"eventType\",\"type\":\"uint8\"}],\"name\":\"TokenMorphed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Unpaused\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MINTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"entries\",\"type\":\"address[]\"}],\"name\":\"addToPresaleList\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"bulkBuy\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bulkBuyLimit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"daoAddress\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"geneOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"gene\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isInPresaleWhitelist\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPresale\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isSale\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"isTokenUnique\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNFTsPerWallet\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxNFTsPerWalletPresale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"officialSaleStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"presaleList\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presaleMint\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"presaleStart\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reserveMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reservedNFTsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"royaltyFeeBps\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"royaltyInfo\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"royaltyAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_baseURI\",\"type\":\"string\"}],\"name\":\"setBaseURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_bulkBuyLimit\",\"type\":\"uint256\"}],\"name\":\"setBulkBuyLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_maxSupply\",\"type\":\"uint256\"}],\"name\":\"setMaxSupply\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newShardedMindsPrice\",\"type\":\"uint256\"}],\"name\":\"setShardedMindsPrice\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"shardedMindsPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uniquesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// ShardedMinds is an auto generated Go binding around an Ethereum contract.
type ShardedMinds struct {
	ShardedMindsCaller     // Read-only binding to the contract
	ShardedMindsTransactor // Write-only binding to the contract
	ShardedMindsFilterer   // Log filterer for contract events
}

// ShardedMindsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ShardedMindsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShardedMindsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ShardedMindsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShardedMindsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ShardedMindsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ShardedMindsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ShardedMindsSession struct {
	Contract     *ShardedMinds     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ShardedMindsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ShardedMindsCallerSession struct {
	Contract *ShardedMindsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ShardedMindsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ShardedMindsTransactorSession struct {
	Contract     *ShardedMindsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ShardedMindsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ShardedMindsRaw struct {
	Contract *ShardedMinds // Generic contract binding to access the raw methods on
}

// ShardedMindsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ShardedMindsCallerRaw struct {
	Contract *ShardedMindsCaller // Generic read-only contract binding to access the raw methods on
}

// ShardedMindsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ShardedMindsTransactorRaw struct {
	Contract *ShardedMindsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewShardedMinds creates a new instance of ShardedMinds, bound to a specific deployed contract.
func NewShardedMinds(address common.Address, backend bind.ContractBackend) (*ShardedMinds, error) {
	contract, err := bindShardedMinds(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ShardedMinds{ShardedMindsCaller: ShardedMindsCaller{contract: contract}, ShardedMindsTransactor: ShardedMindsTransactor{contract: contract}, ShardedMindsFilterer: ShardedMindsFilterer{contract: contract}}, nil
}

// NewShardedMindsCaller creates a new read-only instance of ShardedMinds, bound to a specific deployed contract.
func NewShardedMindsCaller(address common.Address, caller bind.ContractCaller) (*ShardedMindsCaller, error) {
	contract, err := bindShardedMinds(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ShardedMindsCaller{contract: contract}, nil
}

// NewShardedMindsTransactor creates a new write-only instance of ShardedMinds, bound to a specific deployed contract.
func NewShardedMindsTransactor(address common.Address, transactor bind.ContractTransactor) (*ShardedMindsTransactor, error) {
	contract, err := bindShardedMinds(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ShardedMindsTransactor{contract: contract}, nil
}

// NewShardedMindsFilterer creates a new log filterer instance of ShardedMinds, bound to a specific deployed contract.
func NewShardedMindsFilterer(address common.Address, filterer bind.ContractFilterer) (*ShardedMindsFilterer, error) {
	contract, err := bindShardedMinds(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ShardedMindsFilterer{contract: contract}, nil
}

// bindShardedMinds binds a generic wrapper to an already deployed contract.
func bindShardedMinds(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ShardedMindsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShardedMinds *ShardedMindsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShardedMinds.Contract.ShardedMindsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShardedMinds *ShardedMindsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardedMinds.Contract.ShardedMindsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShardedMinds *ShardedMindsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShardedMinds.Contract.ShardedMindsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ShardedMinds *ShardedMindsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ShardedMinds.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ShardedMinds *ShardedMindsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardedMinds.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ShardedMinds *ShardedMindsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ShardedMinds.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ShardedMinds *ShardedMindsCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ShardedMinds *ShardedMindsSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ShardedMinds.Contract.DEFAULTADMINROLE(&_ShardedMinds.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_ShardedMinds *ShardedMindsCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _ShardedMinds.Contract.DEFAULTADMINROLE(&_ShardedMinds.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ShardedMinds *ShardedMindsCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ShardedMinds *ShardedMindsSession) MINTERROLE() ([32]byte, error) {
	return _ShardedMinds.Contract.MINTERROLE(&_ShardedMinds.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_ShardedMinds *ShardedMindsCallerSession) MINTERROLE() ([32]byte, error) {
	return _ShardedMinds.Contract.MINTERROLE(&_ShardedMinds.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ShardedMinds *ShardedMindsCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ShardedMinds *ShardedMindsSession) PAUSERROLE() ([32]byte, error) {
	return _ShardedMinds.Contract.PAUSERROLE(&_ShardedMinds.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_ShardedMinds *ShardedMindsCallerSession) PAUSERROLE() ([32]byte, error) {
	return _ShardedMinds.Contract.PAUSERROLE(&_ShardedMinds.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ShardedMinds.Contract.BalanceOf(&_ShardedMinds.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _ShardedMinds.Contract.BalanceOf(&_ShardedMinds.CallOpts, owner)
}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) BulkBuyLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "bulkBuyLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) BulkBuyLimit() (*big.Int, error) {
	return _ShardedMinds.Contract.BulkBuyLimit(&_ShardedMinds.CallOpts)
}

// BulkBuyLimit is a free data retrieval call binding the contract method 0xa49bccca.
//
// Solidity: function bulkBuyLimit() view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) BulkBuyLimit() (*big.Int, error) {
	return _ShardedMinds.Contract.BulkBuyLimit(&_ShardedMinds.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_ShardedMinds *ShardedMindsCaller) DaoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "daoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_ShardedMinds *ShardedMindsSession) DaoAddress() (common.Address, error) {
	return _ShardedMinds.Contract.DaoAddress(&_ShardedMinds.CallOpts)
}

// DaoAddress is a free data retrieval call binding the contract method 0x2131c68c.
//
// Solidity: function daoAddress() view returns(address)
func (_ShardedMinds *ShardedMindsCallerSession) DaoAddress() (common.Address, error) {
	return _ShardedMinds.Contract.DaoAddress(&_ShardedMinds.CallOpts)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_ShardedMinds *ShardedMindsCaller) GeneOf(opts *bind.CallOpts, tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "geneOf", tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_ShardedMinds *ShardedMindsSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _ShardedMinds.Contract.GeneOf(&_ShardedMinds.CallOpts, tokenId)
}

// GeneOf is a free data retrieval call binding the contract method 0x6a5be686.
//
// Solidity: function geneOf(uint256 tokenId) view returns(uint256 gene)
func (_ShardedMinds *ShardedMindsCallerSession) GeneOf(tokenId *big.Int) (*big.Int, error) {
	return _ShardedMinds.Contract.GeneOf(&_ShardedMinds.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ShardedMinds *ShardedMindsCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ShardedMinds *ShardedMindsSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ShardedMinds.Contract.GetApproved(&_ShardedMinds.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_ShardedMinds *ShardedMindsCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _ShardedMinds.Contract.GetApproved(&_ShardedMinds.CallOpts, tokenId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ShardedMinds *ShardedMindsCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ShardedMinds *ShardedMindsSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ShardedMinds.Contract.GetRoleAdmin(&_ShardedMinds.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_ShardedMinds *ShardedMindsCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _ShardedMinds.Contract.GetRoleAdmin(&_ShardedMinds.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ShardedMinds *ShardedMindsCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ShardedMinds *ShardedMindsSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ShardedMinds.Contract.GetRoleMember(&_ShardedMinds.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_ShardedMinds *ShardedMindsCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _ShardedMinds.Contract.GetRoleMember(&_ShardedMinds.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ShardedMinds.Contract.GetRoleMemberCount(&_ShardedMinds.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _ShardedMinds.Contract.GetRoleMemberCount(&_ShardedMinds.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ShardedMinds *ShardedMindsCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ShardedMinds *ShardedMindsSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ShardedMinds.Contract.HasRole(&_ShardedMinds.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_ShardedMinds *ShardedMindsCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _ShardedMinds.Contract.HasRole(&_ShardedMinds.CallOpts, role, account)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ShardedMinds *ShardedMindsCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ShardedMinds *ShardedMindsSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ShardedMinds.Contract.IsApprovedForAll(&_ShardedMinds.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_ShardedMinds *ShardedMindsCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _ShardedMinds.Contract.IsApprovedForAll(&_ShardedMinds.CallOpts, owner, operator)
}

// IsInPresaleWhitelist is a free data retrieval call binding the contract method 0x3f134d3d.
//
// Solidity: function isInPresaleWhitelist(address _address) view returns(bool)
func (_ShardedMinds *ShardedMindsCaller) IsInPresaleWhitelist(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "isInPresaleWhitelist", _address)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInPresaleWhitelist is a free data retrieval call binding the contract method 0x3f134d3d.
//
// Solidity: function isInPresaleWhitelist(address _address) view returns(bool)
func (_ShardedMinds *ShardedMindsSession) IsInPresaleWhitelist(_address common.Address) (bool, error) {
	return _ShardedMinds.Contract.IsInPresaleWhitelist(&_ShardedMinds.CallOpts, _address)
}

// IsInPresaleWhitelist is a free data retrieval call binding the contract method 0x3f134d3d.
//
// Solidity: function isInPresaleWhitelist(address _address) view returns(bool)
func (_ShardedMinds *ShardedMindsCallerSession) IsInPresaleWhitelist(_address common.Address) (bool, error) {
	return _ShardedMinds.Contract.IsInPresaleWhitelist(&_ShardedMinds.CallOpts, _address)
}

// IsPresale is a free data retrieval call binding the contract method 0x95364a84.
//
// Solidity: function isPresale() view returns(bool)
func (_ShardedMinds *ShardedMindsCaller) IsPresale(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "isPresale")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPresale is a free data retrieval call binding the contract method 0x95364a84.
//
// Solidity: function isPresale() view returns(bool)
func (_ShardedMinds *ShardedMindsSession) IsPresale() (bool, error) {
	return _ShardedMinds.Contract.IsPresale(&_ShardedMinds.CallOpts)
}

// IsPresale is a free data retrieval call binding the contract method 0x95364a84.
//
// Solidity: function isPresale() view returns(bool)
func (_ShardedMinds *ShardedMindsCallerSession) IsPresale() (bool, error) {
	return _ShardedMinds.Contract.IsPresale(&_ShardedMinds.CallOpts)
}

// IsSale is a free data retrieval call binding the contract method 0xf8c1c186.
//
// Solidity: function isSale() view returns(bool)
func (_ShardedMinds *ShardedMindsCaller) IsSale(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "isSale")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSale is a free data retrieval call binding the contract method 0xf8c1c186.
//
// Solidity: function isSale() view returns(bool)
func (_ShardedMinds *ShardedMindsSession) IsSale() (bool, error) {
	return _ShardedMinds.Contract.IsSale(&_ShardedMinds.CallOpts)
}

// IsSale is a free data retrieval call binding the contract method 0xf8c1c186.
//
// Solidity: function isSale() view returns(bool)
func (_ShardedMinds *ShardedMindsCallerSession) IsSale() (bool, error) {
	return _ShardedMinds.Contract.IsSale(&_ShardedMinds.CallOpts)
}

// IsTokenUnique is a free data retrieval call binding the contract method 0x35c6ed60.
//
// Solidity: function isTokenUnique(uint256 tokenId) view returns(bool, uint256)
func (_ShardedMinds *ShardedMindsCaller) IsTokenUnique(opts *bind.CallOpts, tokenId *big.Int) (bool, *big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "isTokenUnique", tokenId)

	if err != nil {
		return *new(bool), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// IsTokenUnique is a free data retrieval call binding the contract method 0x35c6ed60.
//
// Solidity: function isTokenUnique(uint256 tokenId) view returns(bool, uint256)
func (_ShardedMinds *ShardedMindsSession) IsTokenUnique(tokenId *big.Int) (bool, *big.Int, error) {
	return _ShardedMinds.Contract.IsTokenUnique(&_ShardedMinds.CallOpts, tokenId)
}

// IsTokenUnique is a free data retrieval call binding the contract method 0x35c6ed60.
//
// Solidity: function isTokenUnique(uint256 tokenId) view returns(bool, uint256)
func (_ShardedMinds *ShardedMindsCallerSession) IsTokenUnique(tokenId *big.Int) (bool, *big.Int, error) {
	return _ShardedMinds.Contract.IsTokenUnique(&_ShardedMinds.CallOpts, tokenId)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_ShardedMinds *ShardedMindsCaller) LastTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "lastTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_ShardedMinds *ShardedMindsSession) LastTokenId() (*big.Int, error) {
	return _ShardedMinds.Contract.LastTokenId(&_ShardedMinds.CallOpts)
}

// LastTokenId is a free data retrieval call binding the contract method 0xf84ddf0b.
//
// Solidity: function lastTokenId() view returns(uint256 tokenId)
func (_ShardedMinds *ShardedMindsCallerSession) LastTokenId() (*big.Int, error) {
	return _ShardedMinds.Contract.LastTokenId(&_ShardedMinds.CallOpts)
}

// MaxNFTsPerWallet is a free data retrieval call binding the contract method 0x8160bf78.
//
// Solidity: function maxNFTsPerWallet() view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) MaxNFTsPerWallet(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "maxNFTsPerWallet")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNFTsPerWallet is a free data retrieval call binding the contract method 0x8160bf78.
//
// Solidity: function maxNFTsPerWallet() view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) MaxNFTsPerWallet() (*big.Int, error) {
	return _ShardedMinds.Contract.MaxNFTsPerWallet(&_ShardedMinds.CallOpts)
}

// MaxNFTsPerWallet is a free data retrieval call binding the contract method 0x8160bf78.
//
// Solidity: function maxNFTsPerWallet() view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) MaxNFTsPerWallet() (*big.Int, error) {
	return _ShardedMinds.Contract.MaxNFTsPerWallet(&_ShardedMinds.CallOpts)
}

// MaxNFTsPerWalletPresale is a free data retrieval call binding the contract method 0xd3d1a0c7.
//
// Solidity: function maxNFTsPerWalletPresale() view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) MaxNFTsPerWalletPresale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "maxNFTsPerWalletPresale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxNFTsPerWalletPresale is a free data retrieval call binding the contract method 0xd3d1a0c7.
//
// Solidity: function maxNFTsPerWalletPresale() view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) MaxNFTsPerWalletPresale() (*big.Int, error) {
	return _ShardedMinds.Contract.MaxNFTsPerWalletPresale(&_ShardedMinds.CallOpts)
}

// MaxNFTsPerWalletPresale is a free data retrieval call binding the contract method 0xd3d1a0c7.
//
// Solidity: function maxNFTsPerWalletPresale() view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) MaxNFTsPerWalletPresale() (*big.Int, error) {
	return _ShardedMinds.Contract.MaxNFTsPerWalletPresale(&_ShardedMinds.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) MaxSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "maxSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) MaxSupply() (*big.Int, error) {
	return _ShardedMinds.Contract.MaxSupply(&_ShardedMinds.CallOpts)
}

// MaxSupply is a free data retrieval call binding the contract method 0xd5abeb01.
//
// Solidity: function maxSupply() view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) MaxSupply() (*big.Int, error) {
	return _ShardedMinds.Contract.MaxSupply(&_ShardedMinds.CallOpts)
}

// Mint0 is a free data retrieval call binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) pure returns()
func (_ShardedMinds *ShardedMindsCaller) Mint0(opts *bind.CallOpts, to common.Address) error {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "mint0", to)

	if err != nil {
		return err
	}

	return err

}

// Mint0 is a free data retrieval call binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) pure returns()
func (_ShardedMinds *ShardedMindsSession) Mint0(to common.Address) error {
	return _ShardedMinds.Contract.Mint0(&_ShardedMinds.CallOpts, to)
}

// Mint0 is a free data retrieval call binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) pure returns()
func (_ShardedMinds *ShardedMindsCallerSession) Mint0(to common.Address) error {
	return _ShardedMinds.Contract.Mint0(&_ShardedMinds.CallOpts, to)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShardedMinds *ShardedMindsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShardedMinds *ShardedMindsSession) Name() (string, error) {
	return _ShardedMinds.Contract.Name(&_ShardedMinds.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ShardedMinds *ShardedMindsCallerSession) Name() (string, error) {
	return _ShardedMinds.Contract.Name(&_ShardedMinds.CallOpts)
}

// OfficialSaleStart is a free data retrieval call binding the contract method 0x338aece1.
//
// Solidity: function officialSaleStart() view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) OfficialSaleStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "officialSaleStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OfficialSaleStart is a free data retrieval call binding the contract method 0x338aece1.
//
// Solidity: function officialSaleStart() view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) OfficialSaleStart() (*big.Int, error) {
	return _ShardedMinds.Contract.OfficialSaleStart(&_ShardedMinds.CallOpts)
}

// OfficialSaleStart is a free data retrieval call binding the contract method 0x338aece1.
//
// Solidity: function officialSaleStart() view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) OfficialSaleStart() (*big.Int, error) {
	return _ShardedMinds.Contract.OfficialSaleStart(&_ShardedMinds.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShardedMinds *ShardedMindsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShardedMinds *ShardedMindsSession) Owner() (common.Address, error) {
	return _ShardedMinds.Contract.Owner(&_ShardedMinds.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ShardedMinds *ShardedMindsCallerSession) Owner() (common.Address, error) {
	return _ShardedMinds.Contract.Owner(&_ShardedMinds.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ShardedMinds *ShardedMindsCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ShardedMinds *ShardedMindsSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ShardedMinds.Contract.OwnerOf(&_ShardedMinds.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_ShardedMinds *ShardedMindsCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _ShardedMinds.Contract.OwnerOf(&_ShardedMinds.CallOpts, tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ShardedMinds *ShardedMindsCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ShardedMinds *ShardedMindsSession) Paused() (bool, error) {
	return _ShardedMinds.Contract.Paused(&_ShardedMinds.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ShardedMinds *ShardedMindsCallerSession) Paused() (bool, error) {
	return _ShardedMinds.Contract.Paused(&_ShardedMinds.CallOpts)
}

// PresaleList is a free data retrieval call binding the contract method 0x12fb92e0.
//
// Solidity: function presaleList(address ) view returns(bool)
func (_ShardedMinds *ShardedMindsCaller) PresaleList(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "presaleList", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PresaleList is a free data retrieval call binding the contract method 0x12fb92e0.
//
// Solidity: function presaleList(address ) view returns(bool)
func (_ShardedMinds *ShardedMindsSession) PresaleList(arg0 common.Address) (bool, error) {
	return _ShardedMinds.Contract.PresaleList(&_ShardedMinds.CallOpts, arg0)
}

// PresaleList is a free data retrieval call binding the contract method 0x12fb92e0.
//
// Solidity: function presaleList(address ) view returns(bool)
func (_ShardedMinds *ShardedMindsCallerSession) PresaleList(arg0 common.Address) (bool, error) {
	return _ShardedMinds.Contract.PresaleList(&_ShardedMinds.CallOpts, arg0)
}

// PresaleStart is a free data retrieval call binding the contract method 0xde8801e5.
//
// Solidity: function presaleStart() view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) PresaleStart(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "presaleStart")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PresaleStart is a free data retrieval call binding the contract method 0xde8801e5.
//
// Solidity: function presaleStart() view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) PresaleStart() (*big.Int, error) {
	return _ShardedMinds.Contract.PresaleStart(&_ShardedMinds.CallOpts)
}

// PresaleStart is a free data retrieval call binding the contract method 0xde8801e5.
//
// Solidity: function presaleStart() view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) PresaleStart() (*big.Int, error) {
	return _ShardedMinds.Contract.PresaleStart(&_ShardedMinds.CallOpts)
}

// ReservedNFTsCount is a free data retrieval call binding the contract method 0x24fbb4b0.
//
// Solidity: function reservedNFTsCount() view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) ReservedNFTsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "reservedNFTsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReservedNFTsCount is a free data retrieval call binding the contract method 0x24fbb4b0.
//
// Solidity: function reservedNFTsCount() view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) ReservedNFTsCount() (*big.Int, error) {
	return _ShardedMinds.Contract.ReservedNFTsCount(&_ShardedMinds.CallOpts)
}

// ReservedNFTsCount is a free data retrieval call binding the contract method 0x24fbb4b0.
//
// Solidity: function reservedNFTsCount() view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) ReservedNFTsCount() (*big.Int, error) {
	return _ShardedMinds.Contract.ReservedNFTsCount(&_ShardedMinds.CallOpts)
}

// RoyaltyFeeBps is a free data retrieval call binding the contract method 0x91192765.
//
// Solidity: function royaltyFeeBps() view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) RoyaltyFeeBps(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "royaltyFeeBps")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RoyaltyFeeBps is a free data retrieval call binding the contract method 0x91192765.
//
// Solidity: function royaltyFeeBps() view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) RoyaltyFeeBps() (*big.Int, error) {
	return _ShardedMinds.Contract.RoyaltyFeeBps(&_ShardedMinds.CallOpts)
}

// RoyaltyFeeBps is a free data retrieval call binding the contract method 0x91192765.
//
// Solidity: function royaltyFeeBps() view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) RoyaltyFeeBps() (*big.Int, error) {
	return _ShardedMinds.Contract.RoyaltyFeeBps(&_ShardedMinds.CallOpts)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 value) view returns(address receiver, uint256 royaltyAmount)
func (_ShardedMinds *ShardedMindsCaller) RoyaltyInfo(opts *bind.CallOpts, tokenId *big.Int, value *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "royaltyInfo", tokenId, value)

	outstruct := new(struct {
		Receiver      common.Address
		RoyaltyAmount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Receiver = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.RoyaltyAmount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 value) view returns(address receiver, uint256 royaltyAmount)
func (_ShardedMinds *ShardedMindsSession) RoyaltyInfo(tokenId *big.Int, value *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _ShardedMinds.Contract.RoyaltyInfo(&_ShardedMinds.CallOpts, tokenId, value)
}

// RoyaltyInfo is a free data retrieval call binding the contract method 0x2a55205a.
//
// Solidity: function royaltyInfo(uint256 tokenId, uint256 value) view returns(address receiver, uint256 royaltyAmount)
func (_ShardedMinds *ShardedMindsCallerSession) RoyaltyInfo(tokenId *big.Int, value *big.Int) (struct {
	Receiver      common.Address
	RoyaltyAmount *big.Int
}, error) {
	return _ShardedMinds.Contract.RoyaltyInfo(&_ShardedMinds.CallOpts, tokenId, value)
}

// ShardedMindsPrice is a free data retrieval call binding the contract method 0x05111989.
//
// Solidity: function shardedMindsPrice() view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) ShardedMindsPrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "shardedMindsPrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ShardedMindsPrice is a free data retrieval call binding the contract method 0x05111989.
//
// Solidity: function shardedMindsPrice() view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) ShardedMindsPrice() (*big.Int, error) {
	return _ShardedMinds.Contract.ShardedMindsPrice(&_ShardedMinds.CallOpts)
}

// ShardedMindsPrice is a free data retrieval call binding the contract method 0x05111989.
//
// Solidity: function shardedMindsPrice() view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) ShardedMindsPrice() (*big.Int, error) {
	return _ShardedMinds.Contract.ShardedMindsPrice(&_ShardedMinds.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShardedMinds *ShardedMindsCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShardedMinds *ShardedMindsSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ShardedMinds.Contract.SupportsInterface(&_ShardedMinds.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_ShardedMinds *ShardedMindsCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _ShardedMinds.Contract.SupportsInterface(&_ShardedMinds.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShardedMinds *ShardedMindsCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShardedMinds *ShardedMindsSession) Symbol() (string, error) {
	return _ShardedMinds.Contract.Symbol(&_ShardedMinds.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ShardedMinds *ShardedMindsCallerSession) Symbol() (string, error) {
	return _ShardedMinds.Contract.Symbol(&_ShardedMinds.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ShardedMinds.Contract.TokenByIndex(&_ShardedMinds.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _ShardedMinds.Contract.TokenByIndex(&_ShardedMinds.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ShardedMinds.Contract.TokenOfOwnerByIndex(&_ShardedMinds.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _ShardedMinds.Contract.TokenOfOwnerByIndex(&_ShardedMinds.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ShardedMinds *ShardedMindsCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ShardedMinds *ShardedMindsSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ShardedMinds.Contract.TokenURI(&_ShardedMinds.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_ShardedMinds *ShardedMindsCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _ShardedMinds.Contract.TokenURI(&_ShardedMinds.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) TotalSupply() (*big.Int, error) {
	return _ShardedMinds.Contract.TotalSupply(&_ShardedMinds.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) TotalSupply() (*big.Int, error) {
	return _ShardedMinds.Contract.TotalSupply(&_ShardedMinds.CallOpts)
}

// UniquesCount is a free data retrieval call binding the contract method 0x2bf15e3a.
//
// Solidity: function uniquesCount() view returns(uint256)
func (_ShardedMinds *ShardedMindsCaller) UniquesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ShardedMinds.contract.Call(opts, &out, "uniquesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UniquesCount is a free data retrieval call binding the contract method 0x2bf15e3a.
//
// Solidity: function uniquesCount() view returns(uint256)
func (_ShardedMinds *ShardedMindsSession) UniquesCount() (*big.Int, error) {
	return _ShardedMinds.Contract.UniquesCount(&_ShardedMinds.CallOpts)
}

// UniquesCount is a free data retrieval call binding the contract method 0x2bf15e3a.
//
// Solidity: function uniquesCount() view returns(uint256)
func (_ShardedMinds *ShardedMindsCallerSession) UniquesCount() (*big.Int, error) {
	return _ShardedMinds.Contract.UniquesCount(&_ShardedMinds.CallOpts)
}

// AddToPresaleList is a paid mutator transaction binding the contract method 0x7204a3c9.
//
// Solidity: function addToPresaleList(address[] entries) returns()
func (_ShardedMinds *ShardedMindsTransactor) AddToPresaleList(opts *bind.TransactOpts, entries []common.Address) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "addToPresaleList", entries)
}

// AddToPresaleList is a paid mutator transaction binding the contract method 0x7204a3c9.
//
// Solidity: function addToPresaleList(address[] entries) returns()
func (_ShardedMinds *ShardedMindsSession) AddToPresaleList(entries []common.Address) (*types.Transaction, error) {
	return _ShardedMinds.Contract.AddToPresaleList(&_ShardedMinds.TransactOpts, entries)
}

// AddToPresaleList is a paid mutator transaction binding the contract method 0x7204a3c9.
//
// Solidity: function addToPresaleList(address[] entries) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) AddToPresaleList(entries []common.Address) (*types.Transaction, error) {
	return _ShardedMinds.Contract.AddToPresaleList(&_ShardedMinds.TransactOpts, entries)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ShardedMinds *ShardedMindsTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ShardedMinds *ShardedMindsSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.Approve(&_ShardedMinds.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.Approve(&_ShardedMinds.TransactOpts, to, tokenId)
}

// BulkBuy is a paid mutator transaction binding the contract method 0xd5a83d3e.
//
// Solidity: function bulkBuy(uint256 amount) payable returns()
func (_ShardedMinds *ShardedMindsTransactor) BulkBuy(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "bulkBuy", amount)
}

// BulkBuy is a paid mutator transaction binding the contract method 0xd5a83d3e.
//
// Solidity: function bulkBuy(uint256 amount) payable returns()
func (_ShardedMinds *ShardedMindsSession) BulkBuy(amount *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.BulkBuy(&_ShardedMinds.TransactOpts, amount)
}

// BulkBuy is a paid mutator transaction binding the contract method 0xd5a83d3e.
//
// Solidity: function bulkBuy(uint256 amount) payable returns()
func (_ShardedMinds *ShardedMindsTransactorSession) BulkBuy(amount *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.BulkBuy(&_ShardedMinds.TransactOpts, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ShardedMinds *ShardedMindsTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ShardedMinds *ShardedMindsSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.Burn(&_ShardedMinds.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.Burn(&_ShardedMinds.TransactOpts, tokenId)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ShardedMinds *ShardedMindsTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ShardedMinds *ShardedMindsSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShardedMinds.Contract.GrantRole(&_ShardedMinds.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShardedMinds.Contract.GrantRole(&_ShardedMinds.TransactOpts, role, account)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_ShardedMinds *ShardedMindsTransactor) Mint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "mint")
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_ShardedMinds *ShardedMindsSession) Mint() (*types.Transaction, error) {
	return _ShardedMinds.Contract.Mint(&_ShardedMinds.TransactOpts)
}

// Mint is a paid mutator transaction binding the contract method 0x1249c58b.
//
// Solidity: function mint() payable returns()
func (_ShardedMinds *ShardedMindsTransactorSession) Mint() (*types.Transaction, error) {
	return _ShardedMinds.Contract.Mint(&_ShardedMinds.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ShardedMinds *ShardedMindsTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ShardedMinds *ShardedMindsSession) Pause() (*types.Transaction, error) {
	return _ShardedMinds.Contract.Pause(&_ShardedMinds.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_ShardedMinds *ShardedMindsTransactorSession) Pause() (*types.Transaction, error) {
	return _ShardedMinds.Contract.Pause(&_ShardedMinds.TransactOpts)
}

// PresaleMint is a paid mutator transaction binding the contract method 0x59533d6c.
//
// Solidity: function presaleMint() payable returns()
func (_ShardedMinds *ShardedMindsTransactor) PresaleMint(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "presaleMint")
}

// PresaleMint is a paid mutator transaction binding the contract method 0x59533d6c.
//
// Solidity: function presaleMint() payable returns()
func (_ShardedMinds *ShardedMindsSession) PresaleMint() (*types.Transaction, error) {
	return _ShardedMinds.Contract.PresaleMint(&_ShardedMinds.TransactOpts)
}

// PresaleMint is a paid mutator transaction binding the contract method 0x59533d6c.
//
// Solidity: function presaleMint() payable returns()
func (_ShardedMinds *ShardedMindsTransactorSession) PresaleMint() (*types.Transaction, error) {
	return _ShardedMinds.Contract.PresaleMint(&_ShardedMinds.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShardedMinds *ShardedMindsTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShardedMinds *ShardedMindsSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShardedMinds.Contract.RenounceOwnership(&_ShardedMinds.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ShardedMinds *ShardedMindsTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ShardedMinds.Contract.RenounceOwnership(&_ShardedMinds.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ShardedMinds *ShardedMindsTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ShardedMinds *ShardedMindsSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShardedMinds.Contract.RenounceRole(&_ShardedMinds.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShardedMinds.Contract.RenounceRole(&_ShardedMinds.TransactOpts, role, account)
}

// ReserveMint is a paid mutator transaction binding the contract method 0x1342ff4c.
//
// Solidity: function reserveMint(uint256 amount) returns()
func (_ShardedMinds *ShardedMindsTransactor) ReserveMint(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "reserveMint", amount)
}

// ReserveMint is a paid mutator transaction binding the contract method 0x1342ff4c.
//
// Solidity: function reserveMint(uint256 amount) returns()
func (_ShardedMinds *ShardedMindsSession) ReserveMint(amount *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.ReserveMint(&_ShardedMinds.TransactOpts, amount)
}

// ReserveMint is a paid mutator transaction binding the contract method 0x1342ff4c.
//
// Solidity: function reserveMint(uint256 amount) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) ReserveMint(amount *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.ReserveMint(&_ShardedMinds.TransactOpts, amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ShardedMinds *ShardedMindsTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ShardedMinds *ShardedMindsSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShardedMinds.Contract.RevokeRole(&_ShardedMinds.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _ShardedMinds.Contract.RevokeRole(&_ShardedMinds.TransactOpts, role, account)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ShardedMinds *ShardedMindsTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ShardedMinds *ShardedMindsSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SafeTransferFrom(&_ShardedMinds.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SafeTransferFrom(&_ShardedMinds.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ShardedMinds *ShardedMindsTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ShardedMinds *ShardedMindsSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SafeTransferFrom0(&_ShardedMinds.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SafeTransferFrom0(&_ShardedMinds.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ShardedMinds *ShardedMindsTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ShardedMinds *ShardedMindsSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SetApprovalForAll(&_ShardedMinds.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SetApprovalForAll(&_ShardedMinds.TransactOpts, operator, approved)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_ShardedMinds *ShardedMindsTransactor) SetBaseURI(opts *bind.TransactOpts, _baseURI string) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "setBaseURI", _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_ShardedMinds *ShardedMindsSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SetBaseURI(&_ShardedMinds.TransactOpts, _baseURI)
}

// SetBaseURI is a paid mutator transaction binding the contract method 0x55f804b3.
//
// Solidity: function setBaseURI(string _baseURI) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) SetBaseURI(_baseURI string) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SetBaseURI(&_ShardedMinds.TransactOpts, _baseURI)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x5e468dfd.
//
// Solidity: function setBulkBuyLimit(uint256 _bulkBuyLimit) returns()
func (_ShardedMinds *ShardedMindsTransactor) SetBulkBuyLimit(opts *bind.TransactOpts, _bulkBuyLimit *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "setBulkBuyLimit", _bulkBuyLimit)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x5e468dfd.
//
// Solidity: function setBulkBuyLimit(uint256 _bulkBuyLimit) returns()
func (_ShardedMinds *ShardedMindsSession) SetBulkBuyLimit(_bulkBuyLimit *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SetBulkBuyLimit(&_ShardedMinds.TransactOpts, _bulkBuyLimit)
}

// SetBulkBuyLimit is a paid mutator transaction binding the contract method 0x5e468dfd.
//
// Solidity: function setBulkBuyLimit(uint256 _bulkBuyLimit) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) SetBulkBuyLimit(_bulkBuyLimit *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SetBulkBuyLimit(&_ShardedMinds.TransactOpts, _bulkBuyLimit)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_ShardedMinds *ShardedMindsTransactor) SetMaxSupply(opts *bind.TransactOpts, _maxSupply *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "setMaxSupply", _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_ShardedMinds *ShardedMindsSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SetMaxSupply(&_ShardedMinds.TransactOpts, _maxSupply)
}

// SetMaxSupply is a paid mutator transaction binding the contract method 0x6f8b44b0.
//
// Solidity: function setMaxSupply(uint256 _maxSupply) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) SetMaxSupply(_maxSupply *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SetMaxSupply(&_ShardedMinds.TransactOpts, _maxSupply)
}

// SetShardedMindsPrice is a paid mutator transaction binding the contract method 0xbf7ddca7.
//
// Solidity: function setShardedMindsPrice(uint256 newShardedMindsPrice) returns()
func (_ShardedMinds *ShardedMindsTransactor) SetShardedMindsPrice(opts *bind.TransactOpts, newShardedMindsPrice *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "setShardedMindsPrice", newShardedMindsPrice)
}

// SetShardedMindsPrice is a paid mutator transaction binding the contract method 0xbf7ddca7.
//
// Solidity: function setShardedMindsPrice(uint256 newShardedMindsPrice) returns()
func (_ShardedMinds *ShardedMindsSession) SetShardedMindsPrice(newShardedMindsPrice *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SetShardedMindsPrice(&_ShardedMinds.TransactOpts, newShardedMindsPrice)
}

// SetShardedMindsPrice is a paid mutator transaction binding the contract method 0xbf7ddca7.
//
// Solidity: function setShardedMindsPrice(uint256 newShardedMindsPrice) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) SetShardedMindsPrice(newShardedMindsPrice *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.SetShardedMindsPrice(&_ShardedMinds.TransactOpts, newShardedMindsPrice)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ShardedMinds *ShardedMindsTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ShardedMinds *ShardedMindsSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.TransferFrom(&_ShardedMinds.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _ShardedMinds.Contract.TransferFrom(&_ShardedMinds.TransactOpts, from, to, tokenId)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShardedMinds *ShardedMindsTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShardedMinds *ShardedMindsSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShardedMinds.Contract.TransferOwnership(&_ShardedMinds.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ShardedMinds *ShardedMindsTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ShardedMinds.Contract.TransferOwnership(&_ShardedMinds.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ShardedMinds *ShardedMindsTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardedMinds.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ShardedMinds *ShardedMindsSession) Unpause() (*types.Transaction, error) {
	return _ShardedMinds.Contract.Unpause(&_ShardedMinds.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_ShardedMinds *ShardedMindsTransactorSession) Unpause() (*types.Transaction, error) {
	return _ShardedMinds.Contract.Unpause(&_ShardedMinds.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ShardedMinds *ShardedMindsTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ShardedMinds.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ShardedMinds *ShardedMindsSession) Receive() (*types.Transaction, error) {
	return _ShardedMinds.Contract.Receive(&_ShardedMinds.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ShardedMinds *ShardedMindsTransactorSession) Receive() (*types.Transaction, error) {
	return _ShardedMinds.Contract.Receive(&_ShardedMinds.TransactOpts)
}

// ShardedMindsApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ShardedMinds contract.
type ShardedMindsApprovalIterator struct {
	Event *ShardedMindsApproval // Event containing the contract specifics and raw log

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
func (it *ShardedMindsApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsApproval)
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
		it.Event = new(ShardedMindsApproval)
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
func (it *ShardedMindsApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsApproval represents a Approval event raised by the ShardedMinds contract.
type ShardedMindsApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ShardedMinds *ShardedMindsFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*ShardedMindsApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShardedMindsApprovalIterator{contract: _ShardedMinds.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ShardedMinds *ShardedMindsFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ShardedMindsApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsApproval)
				if err := _ShardedMinds.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_ShardedMinds *ShardedMindsFilterer) ParseApproval(log types.Log) (*ShardedMindsApproval, error) {
	event := new(ShardedMindsApproval)
	if err := _ShardedMinds.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ShardedMinds contract.
type ShardedMindsApprovalForAllIterator struct {
	Event *ShardedMindsApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ShardedMindsApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsApprovalForAll)
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
		it.Event = new(ShardedMindsApprovalForAll)
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
func (it *ShardedMindsApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsApprovalForAll represents a ApprovalForAll event raised by the ShardedMinds contract.
type ShardedMindsApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ShardedMinds *ShardedMindsFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*ShardedMindsApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &ShardedMindsApprovalForAllIterator{contract: _ShardedMinds.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ShardedMinds *ShardedMindsFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ShardedMindsApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsApprovalForAll)
				if err := _ShardedMinds.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_ShardedMinds *ShardedMindsFilterer) ParseApprovalForAll(log types.Log) (*ShardedMindsApprovalForAll, error) {
	event := new(ShardedMindsApprovalForAll)
	if err := _ShardedMinds.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsBaseURIChangedIterator is returned from FilterBaseURIChanged and is used to iterate over the raw logs and unpacked data for BaseURIChanged events raised by the ShardedMinds contract.
type ShardedMindsBaseURIChangedIterator struct {
	Event *ShardedMindsBaseURIChanged // Event containing the contract specifics and raw log

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
func (it *ShardedMindsBaseURIChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsBaseURIChanged)
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
		it.Event = new(ShardedMindsBaseURIChanged)
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
func (it *ShardedMindsBaseURIChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsBaseURIChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsBaseURIChanged represents a BaseURIChanged event raised by the ShardedMinds contract.
type ShardedMindsBaseURIChanged struct {
	BaseURI string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBaseURIChanged is a free log retrieval operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_ShardedMinds *ShardedMindsFilterer) FilterBaseURIChanged(opts *bind.FilterOpts) (*ShardedMindsBaseURIChangedIterator, error) {

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return &ShardedMindsBaseURIChangedIterator{contract: _ShardedMinds.contract, event: "BaseURIChanged", logs: logs, sub: sub}, nil
}

// WatchBaseURIChanged is a free log subscription operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_ShardedMinds *ShardedMindsFilterer) WatchBaseURIChanged(opts *bind.WatchOpts, sink chan<- *ShardedMindsBaseURIChanged) (event.Subscription, error) {

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "BaseURIChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsBaseURIChanged)
				if err := _ShardedMinds.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
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

// ParseBaseURIChanged is a log parse operation binding the contract event 0x5411e8ebf1636d9e83d5fc4900bf80cbac82e8790da2a4c94db4895e889eedf6.
//
// Solidity: event BaseURIChanged(string baseURI)
func (_ShardedMinds *ShardedMindsFilterer) ParseBaseURIChanged(log types.Log) (*ShardedMindsBaseURIChanged, error) {
	event := new(ShardedMindsBaseURIChanged)
	if err := _ShardedMinds.contract.UnpackLog(event, "BaseURIChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsBulkBuyLimitChangedIterator is returned from FilterBulkBuyLimitChanged and is used to iterate over the raw logs and unpacked data for BulkBuyLimitChanged events raised by the ShardedMinds contract.
type ShardedMindsBulkBuyLimitChangedIterator struct {
	Event *ShardedMindsBulkBuyLimitChanged // Event containing the contract specifics and raw log

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
func (it *ShardedMindsBulkBuyLimitChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsBulkBuyLimitChanged)
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
		it.Event = new(ShardedMindsBulkBuyLimitChanged)
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
func (it *ShardedMindsBulkBuyLimitChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsBulkBuyLimitChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsBulkBuyLimitChanged represents a BulkBuyLimitChanged event raised by the ShardedMinds contract.
type ShardedMindsBulkBuyLimitChanged struct {
	NewBulkBuyLimit *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBulkBuyLimitChanged is a free log retrieval operation binding the contract event 0xa0e0113404674c6f545b966e8ec54db3066a6c720a0054f0bc4b0c900cfff243.
//
// Solidity: event BulkBuyLimitChanged(uint256 newBulkBuyLimit)
func (_ShardedMinds *ShardedMindsFilterer) FilterBulkBuyLimitChanged(opts *bind.FilterOpts) (*ShardedMindsBulkBuyLimitChangedIterator, error) {

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "BulkBuyLimitChanged")
	if err != nil {
		return nil, err
	}
	return &ShardedMindsBulkBuyLimitChangedIterator{contract: _ShardedMinds.contract, event: "BulkBuyLimitChanged", logs: logs, sub: sub}, nil
}

// WatchBulkBuyLimitChanged is a free log subscription operation binding the contract event 0xa0e0113404674c6f545b966e8ec54db3066a6c720a0054f0bc4b0c900cfff243.
//
// Solidity: event BulkBuyLimitChanged(uint256 newBulkBuyLimit)
func (_ShardedMinds *ShardedMindsFilterer) WatchBulkBuyLimitChanged(opts *bind.WatchOpts, sink chan<- *ShardedMindsBulkBuyLimitChanged) (event.Subscription, error) {

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "BulkBuyLimitChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsBulkBuyLimitChanged)
				if err := _ShardedMinds.contract.UnpackLog(event, "BulkBuyLimitChanged", log); err != nil {
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

// ParseBulkBuyLimitChanged is a log parse operation binding the contract event 0xa0e0113404674c6f545b966e8ec54db3066a6c720a0054f0bc4b0c900cfff243.
//
// Solidity: event BulkBuyLimitChanged(uint256 newBulkBuyLimit)
func (_ShardedMinds *ShardedMindsFilterer) ParseBulkBuyLimitChanged(log types.Log) (*ShardedMindsBulkBuyLimitChanged, error) {
	event := new(ShardedMindsBulkBuyLimitChanged)
	if err := _ShardedMinds.contract.UnpackLog(event, "BulkBuyLimitChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsMaxSupplyChangedIterator is returned from FilterMaxSupplyChanged and is used to iterate over the raw logs and unpacked data for MaxSupplyChanged events raised by the ShardedMinds contract.
type ShardedMindsMaxSupplyChangedIterator struct {
	Event *ShardedMindsMaxSupplyChanged // Event containing the contract specifics and raw log

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
func (it *ShardedMindsMaxSupplyChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsMaxSupplyChanged)
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
		it.Event = new(ShardedMindsMaxSupplyChanged)
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
func (it *ShardedMindsMaxSupplyChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsMaxSupplyChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsMaxSupplyChanged represents a MaxSupplyChanged event raised by the ShardedMinds contract.
type ShardedMindsMaxSupplyChanged struct {
	NewMaxSupply *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMaxSupplyChanged is a free log retrieval operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_ShardedMinds *ShardedMindsFilterer) FilterMaxSupplyChanged(opts *bind.FilterOpts) (*ShardedMindsMaxSupplyChangedIterator, error) {

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "MaxSupplyChanged")
	if err != nil {
		return nil, err
	}
	return &ShardedMindsMaxSupplyChangedIterator{contract: _ShardedMinds.contract, event: "MaxSupplyChanged", logs: logs, sub: sub}, nil
}

// WatchMaxSupplyChanged is a free log subscription operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_ShardedMinds *ShardedMindsFilterer) WatchMaxSupplyChanged(opts *bind.WatchOpts, sink chan<- *ShardedMindsMaxSupplyChanged) (event.Subscription, error) {

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "MaxSupplyChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsMaxSupplyChanged)
				if err := _ShardedMinds.contract.UnpackLog(event, "MaxSupplyChanged", log); err != nil {
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

// ParseMaxSupplyChanged is a log parse operation binding the contract event 0x28a10a2e0b5582da7164754cb994f6214b8af6aa7f7e003305fbc09e7106c513.
//
// Solidity: event MaxSupplyChanged(uint256 newMaxSupply)
func (_ShardedMinds *ShardedMindsFilterer) ParseMaxSupplyChanged(log types.Log) (*ShardedMindsMaxSupplyChanged, error) {
	event := new(ShardedMindsMaxSupplyChanged)
	if err := _ShardedMinds.contract.UnpackLog(event, "MaxSupplyChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ShardedMinds contract.
type ShardedMindsOwnershipTransferredIterator struct {
	Event *ShardedMindsOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ShardedMindsOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsOwnershipTransferred)
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
		it.Event = new(ShardedMindsOwnershipTransferred)
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
func (it *ShardedMindsOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsOwnershipTransferred represents a OwnershipTransferred event raised by the ShardedMinds contract.
type ShardedMindsOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShardedMinds *ShardedMindsFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ShardedMindsOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ShardedMindsOwnershipTransferredIterator{contract: _ShardedMinds.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ShardedMinds *ShardedMindsFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ShardedMindsOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsOwnershipTransferred)
				if err := _ShardedMinds.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_ShardedMinds *ShardedMindsFilterer) ParseOwnershipTransferred(log types.Log) (*ShardedMindsOwnershipTransferred, error) {
	event := new(ShardedMindsOwnershipTransferred)
	if err := _ShardedMinds.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ShardedMinds contract.
type ShardedMindsPausedIterator struct {
	Event *ShardedMindsPaused // Event containing the contract specifics and raw log

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
func (it *ShardedMindsPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsPaused)
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
		it.Event = new(ShardedMindsPaused)
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
func (it *ShardedMindsPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsPaused represents a Paused event raised by the ShardedMinds contract.
type ShardedMindsPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ShardedMinds *ShardedMindsFilterer) FilterPaused(opts *bind.FilterOpts) (*ShardedMindsPausedIterator, error) {

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ShardedMindsPausedIterator{contract: _ShardedMinds.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ShardedMinds *ShardedMindsFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ShardedMindsPaused) (event.Subscription, error) {

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsPaused)
				if err := _ShardedMinds.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_ShardedMinds *ShardedMindsFilterer) ParsePaused(log types.Log) (*ShardedMindsPaused, error) {
	event := new(ShardedMindsPaused)
	if err := _ShardedMinds.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the ShardedMinds contract.
type ShardedMindsRoleAdminChangedIterator struct {
	Event *ShardedMindsRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ShardedMindsRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsRoleAdminChanged)
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
		it.Event = new(ShardedMindsRoleAdminChanged)
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
func (it *ShardedMindsRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsRoleAdminChanged represents a RoleAdminChanged event raised by the ShardedMinds contract.
type ShardedMindsRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ShardedMinds *ShardedMindsFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ShardedMindsRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ShardedMindsRoleAdminChangedIterator{contract: _ShardedMinds.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ShardedMinds *ShardedMindsFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ShardedMindsRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsRoleAdminChanged)
				if err := _ShardedMinds.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_ShardedMinds *ShardedMindsFilterer) ParseRoleAdminChanged(log types.Log) (*ShardedMindsRoleAdminChanged, error) {
	event := new(ShardedMindsRoleAdminChanged)
	if err := _ShardedMinds.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the ShardedMinds contract.
type ShardedMindsRoleGrantedIterator struct {
	Event *ShardedMindsRoleGranted // Event containing the contract specifics and raw log

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
func (it *ShardedMindsRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsRoleGranted)
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
		it.Event = new(ShardedMindsRoleGranted)
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
func (it *ShardedMindsRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsRoleGranted represents a RoleGranted event raised by the ShardedMinds contract.
type ShardedMindsRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShardedMinds *ShardedMindsFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ShardedMindsRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ShardedMindsRoleGrantedIterator{contract: _ShardedMinds.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShardedMinds *ShardedMindsFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ShardedMindsRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsRoleGranted)
				if err := _ShardedMinds.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShardedMinds *ShardedMindsFilterer) ParseRoleGranted(log types.Log) (*ShardedMindsRoleGranted, error) {
	event := new(ShardedMindsRoleGranted)
	if err := _ShardedMinds.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the ShardedMinds contract.
type ShardedMindsRoleRevokedIterator struct {
	Event *ShardedMindsRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ShardedMindsRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsRoleRevoked)
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
		it.Event = new(ShardedMindsRoleRevoked)
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
func (it *ShardedMindsRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsRoleRevoked represents a RoleRevoked event raised by the ShardedMinds contract.
type ShardedMindsRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShardedMinds *ShardedMindsFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ShardedMindsRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ShardedMindsRoleRevokedIterator{contract: _ShardedMinds.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShardedMinds *ShardedMindsFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ShardedMindsRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsRoleRevoked)
				if err := _ShardedMinds.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_ShardedMinds *ShardedMindsFilterer) ParseRoleRevoked(log types.Log) (*ShardedMindsRoleRevoked, error) {
	event := new(ShardedMindsRoleRevoked)
	if err := _ShardedMinds.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsShardedMindsPriceChangedIterator is returned from FilterShardedMindsPriceChanged and is used to iterate over the raw logs and unpacked data for ShardedMindsPriceChanged events raised by the ShardedMinds contract.
type ShardedMindsShardedMindsPriceChangedIterator struct {
	Event *ShardedMindsShardedMindsPriceChanged // Event containing the contract specifics and raw log

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
func (it *ShardedMindsShardedMindsPriceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsShardedMindsPriceChanged)
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
		it.Event = new(ShardedMindsShardedMindsPriceChanged)
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
func (it *ShardedMindsShardedMindsPriceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsShardedMindsPriceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsShardedMindsPriceChanged represents a ShardedMindsPriceChanged event raised by the ShardedMinds contract.
type ShardedMindsShardedMindsPriceChanged struct {
	NewShardedMindsPrice *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterShardedMindsPriceChanged is a free log retrieval operation binding the contract event 0x121e74175537529c2c32dbc05daa24b387fcaae82836e8e1b423bb10360418a5.
//
// Solidity: event ShardedMindsPriceChanged(uint256 newShardedMindsPrice)
func (_ShardedMinds *ShardedMindsFilterer) FilterShardedMindsPriceChanged(opts *bind.FilterOpts) (*ShardedMindsShardedMindsPriceChangedIterator, error) {

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "ShardedMindsPriceChanged")
	if err != nil {
		return nil, err
	}
	return &ShardedMindsShardedMindsPriceChangedIterator{contract: _ShardedMinds.contract, event: "ShardedMindsPriceChanged", logs: logs, sub: sub}, nil
}

// WatchShardedMindsPriceChanged is a free log subscription operation binding the contract event 0x121e74175537529c2c32dbc05daa24b387fcaae82836e8e1b423bb10360418a5.
//
// Solidity: event ShardedMindsPriceChanged(uint256 newShardedMindsPrice)
func (_ShardedMinds *ShardedMindsFilterer) WatchShardedMindsPriceChanged(opts *bind.WatchOpts, sink chan<- *ShardedMindsShardedMindsPriceChanged) (event.Subscription, error) {

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "ShardedMindsPriceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsShardedMindsPriceChanged)
				if err := _ShardedMinds.contract.UnpackLog(event, "ShardedMindsPriceChanged", log); err != nil {
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

// ParseShardedMindsPriceChanged is a log parse operation binding the contract event 0x121e74175537529c2c32dbc05daa24b387fcaae82836e8e1b423bb10360418a5.
//
// Solidity: event ShardedMindsPriceChanged(uint256 newShardedMindsPrice)
func (_ShardedMinds *ShardedMindsFilterer) ParseShardedMindsPriceChanged(log types.Log) (*ShardedMindsShardedMindsPriceChanged, error) {
	event := new(ShardedMindsShardedMindsPriceChanged)
	if err := _ShardedMinds.contract.UnpackLog(event, "ShardedMindsPriceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsTokenMintedIterator is returned from FilterTokenMinted and is used to iterate over the raw logs and unpacked data for TokenMinted events raised by the ShardedMinds contract.
type ShardedMindsTokenMintedIterator struct {
	Event *ShardedMindsTokenMinted // Event containing the contract specifics and raw log

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
func (it *ShardedMindsTokenMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsTokenMinted)
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
		it.Event = new(ShardedMindsTokenMinted)
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
func (it *ShardedMindsTokenMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsTokenMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsTokenMinted represents a TokenMinted event raised by the ShardedMinds contract.
type ShardedMindsTokenMinted struct {
	TokenId *big.Int
	NewGene *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTokenMinted is a free log retrieval operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_ShardedMinds *ShardedMindsFilterer) FilterTokenMinted(opts *bind.FilterOpts, tokenId []*big.Int) (*ShardedMindsTokenMintedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "TokenMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShardedMindsTokenMintedIterator{contract: _ShardedMinds.contract, event: "TokenMinted", logs: logs, sub: sub}, nil
}

// WatchTokenMinted is a free log subscription operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_ShardedMinds *ShardedMindsFilterer) WatchTokenMinted(opts *bind.WatchOpts, sink chan<- *ShardedMindsTokenMinted, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "TokenMinted", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsTokenMinted)
				if err := _ShardedMinds.contract.UnpackLog(event, "TokenMinted", log); err != nil {
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

// ParseTokenMinted is a log parse operation binding the contract event 0x5f7666687319b40936f33c188908d86aea154abd3f4127b4fa0a3f04f303c7da.
//
// Solidity: event TokenMinted(uint256 indexed tokenId, uint256 newGene)
func (_ShardedMinds *ShardedMindsFilterer) ParseTokenMinted(log types.Log) (*ShardedMindsTokenMinted, error) {
	event := new(ShardedMindsTokenMinted)
	if err := _ShardedMinds.contract.UnpackLog(event, "TokenMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsTokenMorphedIterator is returned from FilterTokenMorphed and is used to iterate over the raw logs and unpacked data for TokenMorphed events raised by the ShardedMinds contract.
type ShardedMindsTokenMorphedIterator struct {
	Event *ShardedMindsTokenMorphed // Event containing the contract specifics and raw log

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
func (it *ShardedMindsTokenMorphedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsTokenMorphed)
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
		it.Event = new(ShardedMindsTokenMorphed)
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
func (it *ShardedMindsTokenMorphedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsTokenMorphedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsTokenMorphed represents a TokenMorphed event raised by the ShardedMinds contract.
type ShardedMindsTokenMorphed struct {
	TokenId   *big.Int
	OldGene   *big.Int
	NewGene   *big.Int
	Price     *big.Int
	EventType uint8
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTokenMorphed is a free log retrieval operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_ShardedMinds *ShardedMindsFilterer) FilterTokenMorphed(opts *bind.FilterOpts, tokenId []*big.Int) (*ShardedMindsTokenMorphedIterator, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "TokenMorphed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShardedMindsTokenMorphedIterator{contract: _ShardedMinds.contract, event: "TokenMorphed", logs: logs, sub: sub}, nil
}

// WatchTokenMorphed is a free log subscription operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_ShardedMinds *ShardedMindsFilterer) WatchTokenMorphed(opts *bind.WatchOpts, sink chan<- *ShardedMindsTokenMorphed, tokenId []*big.Int) (event.Subscription, error) {

	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "TokenMorphed", tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsTokenMorphed)
				if err := _ShardedMinds.contract.UnpackLog(event, "TokenMorphed", log); err != nil {
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

// ParseTokenMorphed is a log parse operation binding the contract event 0x8c0bdd7bca83c4e0c810cbecf44bc544a9dc0b9f265664e31ce0ce85f07a052b.
//
// Solidity: event TokenMorphed(uint256 indexed tokenId, uint256 oldGene, uint256 newGene, uint256 price, uint8 eventType)
func (_ShardedMinds *ShardedMindsFilterer) ParseTokenMorphed(log types.Log) (*ShardedMindsTokenMorphed, error) {
	event := new(ShardedMindsTokenMorphed)
	if err := _ShardedMinds.contract.UnpackLog(event, "TokenMorphed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ShardedMinds contract.
type ShardedMindsTransferIterator struct {
	Event *ShardedMindsTransfer // Event containing the contract specifics and raw log

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
func (it *ShardedMindsTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsTransfer)
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
		it.Event = new(ShardedMindsTransfer)
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
func (it *ShardedMindsTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsTransfer represents a Transfer event raised by the ShardedMinds contract.
type ShardedMindsTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ShardedMinds *ShardedMindsFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*ShardedMindsTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &ShardedMindsTransferIterator{contract: _ShardedMinds.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ShardedMinds *ShardedMindsFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ShardedMindsTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsTransfer)
				if err := _ShardedMinds.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_ShardedMinds *ShardedMindsFilterer) ParseTransfer(log types.Log) (*ShardedMindsTransfer, error) {
	event := new(ShardedMindsTransfer)
	if err := _ShardedMinds.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ShardedMindsUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ShardedMinds contract.
type ShardedMindsUnpausedIterator struct {
	Event *ShardedMindsUnpaused // Event containing the contract specifics and raw log

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
func (it *ShardedMindsUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ShardedMindsUnpaused)
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
		it.Event = new(ShardedMindsUnpaused)
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
func (it *ShardedMindsUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ShardedMindsUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ShardedMindsUnpaused represents a Unpaused event raised by the ShardedMinds contract.
type ShardedMindsUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ShardedMinds *ShardedMindsFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ShardedMindsUnpausedIterator, error) {

	logs, sub, err := _ShardedMinds.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ShardedMindsUnpausedIterator{contract: _ShardedMinds.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ShardedMinds *ShardedMindsFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ShardedMindsUnpaused) (event.Subscription, error) {

	logs, sub, err := _ShardedMinds.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ShardedMindsUnpaused)
				if err := _ShardedMinds.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_ShardedMinds *ShardedMindsFilterer) ParseUnpaused(log types.Log) (*ShardedMindsUnpaused, error) {
	event := new(ShardedMindsUnpaused)
	if err := _ShardedMinds.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
