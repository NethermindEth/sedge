// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package csexitpenalties

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

// ExitPenaltyInfo is an auto generated low-level Go binding around an user-defined struct.
type ExitPenaltyInfo struct {
	DelayPenalty         MarkedUint248
	StrikesPenalty       MarkedUint248
	WithdrawalRequestFee MarkedUint248
}

// MarkedUint248 is an auto generated low-level Go binding around an user-defined struct.
type MarkedUint248 struct {
	Value   *big.Int
	IsValue bool
}

// CsexitpenaltiesMetaData contains all meta data concerning the Csexitpenalties contract.
var CsexitpenaltiesMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"parametersRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"strikes\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotStrikes\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ValidatorExitDelayNotApplicable\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroModuleAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroParametersRegistryAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroStrikesAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strikesPenalty\",\"type\":\"uint256\"}],\"name\":\"StrikesPenaltyProcessed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"exitType\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalRequestPaidFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawalRequestRecordedFee\",\"type\":\"uint256\"}],\"name\":\"TriggeredExitFeeRecorded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delayPenalty\",\"type\":\"uint256\"}],\"name\":\"ValidatorExitDelayProcessed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACCOUNTING\",\"outputs\":[{\"internalType\":\"contractICSAccounting\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE\",\"outputs\":[{\"internalType\":\"contractICSModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PARAMETERS_REGISTRY\",\"outputs\":[{\"internalType\":\"contractICSParametersRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STRIKES\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STRIKES_EXIT_TYPE_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VOLUNTARY_EXIT_TYPE_ID\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"getExitPenaltyInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint248\",\"name\":\"value\",\"type\":\"uint248\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structMarkedUint248\",\"name\":\"delayPenalty\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint248\",\"name\":\"value\",\"type\":\"uint248\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structMarkedUint248\",\"name\":\"strikesPenalty\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint248\",\"name\":\"value\",\"type\":\"uint248\"},{\"internalType\":\"bool\",\"name\":\"isValue\",\"type\":\"bool\"}],\"internalType\":\"structMarkedUint248\",\"name\":\"withdrawalRequestFee\",\"type\":\"tuple\"}],\"internalType\":\"structExitPenaltyInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eligibleToExitInSec\",\"type\":\"uint256\"}],\"name\":\"isValidatorExitDelayPenaltyApplicable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eligibleToExitInSec\",\"type\":\"uint256\"}],\"name\":\"processExitDelayReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"}],\"name\":\"processStrikesReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalRequestPaidFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitType\",\"type\":\"uint256\"}],\"name\":\"processTriggeredExit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50600436106100a6575f3560e01c80636dc3f2bd1161006e5780636dc3f2bd1461014b578063848b1dec146101725780638af3014214610185578063d4040379146101ac578063e83ba79d146101cf578063e9f6fdc61461023a575f80fd5b8063094d3a34146100aa5780632fc88741146100ee578063320eacf8146101155780633b596df51461012f57806344dab94914610136575b5f80fd5b6100d17f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f81565b6040516001600160a01b0390911681526020015b60405180910390f35b6100d17f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e42881565b61011d600181565b60405160ff90911681526020016100e5565b61011d5f81565b610149610144366004610ce8565b61024d565b005b6100d17f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da81565b610149610180366004610d37565b610516565b6100d17f000000000000000000000000aa328816027f2d32b9f56d190bc9fa4a5c07637f81565b6101bf6101ba366004610ce8565b610750565b60405190151581526020016100e5565b6101e26101dd366004610d90565b6108f5565b60408051825180516001600160f81b03908116835260209182015115158284015281850151805182168486015282015115156060840152939092015180519093166080820152910151151560a082015260c0016100e5565b610149610248366004610d90565b61099b565b336001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f1614610296576040516303f249a160e51b815260040160405180910390fd5b604051630569b94760e01b8152600481018590525f907f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da6001600160a01b031690630569b94790602401602060405180830381865afa1580156102fb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061031f9190610dd8565b60405163134f60a160e31b8152600481018290529091505f906001600160a01b037f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e4281690639a7b050890602401602060405180830381865afa158015610387573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103ab9190610dd8565b90508083116103cd57604051635d36b8d160e01b815260040160405180910390fd5b5f6103d9878787610bbe565b5f818152602081905260409020805491925090600160f81b900460ff16156104045750505050610510565b60405163f91a79b560e01b8152600481018590525f907f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e4286001600160a01b03169063f91a79b590602401602060405180830381865afa158015610469573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061048d9190610dd8565b905060405180604001604052806104a383610bf3565b6001600160f81b0390811682526001602092830152825192909101511515600160f81b02911617825560405189907ff808c54013437847ff14496ccbb2d51171fc03bb72a783b7905a9dc64875700990610502908b908b908690610e17565b60405180910390a250505050505b50505050565b336001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f161461055f576040516303f249a160e51b815260040160405180910390fd5b8015610749575f610571868686610bbe565b5f818152602081905260409020600281015491925090600160f81b900460ff161561059d575050610749565b604051630569b94760e01b8152600481018890525f907f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da6001600160a01b031690630569b94790602401602060405180830381865afa158015610602573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106269190610dd8565b604051630acd981360e41b8152600481018290529091505f906001600160a01b037f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e428169063acd9813090602401602060405180830381865afa15801561068e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106b29190610dd8565b90505f6106bf8783610c2e565b905060405180604001604052806106d583610bf3565b6001600160f81b0390811682526001602092830152825192909101511515600160f81b02911617600285015560405186908b907fc90027a0742ed6a64a8829261af4ad3d833e668083a49a0634cc61c110c8e8db9061073b908d908d908d908890610e3a565b60405180910390a350505050505b5050505050565b5f336001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f161461079a576040516303f249a160e51b815260040160405180910390fd5b604051630569b94760e01b8152600481018690525f907f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da6001600160a01b031690630569b94790602401602060405180830381865afa1580156107ff573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108239190610dd8565b60405163134f60a160e31b8152600481018290529091505f906001600160a01b037f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e4281690639a7b050890602401602060405180830381865afa15801561088b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108af9190610dd8565b90508084116108c2575f925050506108ed565b5f6108ce888888610bbe565b5f90815260208190526040902054600160f81b900460ff161593505050505b949350505050565b6108fd610c45565b5f610909858585610bbe565b5f9081526020818152604091829020825160a08101845281546001600160f81b038082166060840190815260ff600160f81b93849004811615156080860152908452865180880188526001860154808416825284900482161515818801528487015286518088018852600290950154918216855291900416151592820192909252918101919091529150509392505050565b336001600160a01b037f000000000000000000000000aa328816027f2d32b9f56d190bc9fa4a5c07637f16146109e457604051638fc0069560e01b815260040160405180910390fd5b5f6109f0848484610bbe565b5f818152602081905260409020600181015491925090600160f81b900460ff1615610a1c575050505050565b604051630569b94760e01b8152600481018690525f907f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da6001600160a01b031690630569b94790602401602060405180830381865afa158015610a81573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aa59190610dd8565b60405163533c60d960e01b8152600481018290529091505f906001600160a01b037f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e428169063533c60d990602401602060405180830381865afa158015610b0d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b319190610dd8565b90506040518060400160405280610b4783610bf3565b6001600160f81b03908116825260016020928301819052835193909201511515600160f81b029216919091179084015560405187907f8b2d35b9e1a65aeadb4b39934801c82d77a6526706826087b9e47a7ce005d40590610bad90899089908690610e17565b60405180910390a250505050505050565b5f838383604051602001610bd493929190610e60565b6040516020818303038152906040528051906020012090509392505050565b5f6001600160f81b03821115610c2a576040516306dfcc6560e41b815260f860048201526024810183905260440160405180910390fd5b5090565b5f818310610c3c5781610c3e565b825b9392505050565b6040805160a081019091525f60608201818152608083019190915281908152602001610c80604080518082019091525f808252602082015290565b8152602001610c9e604080518082019091525f808252602082015290565b905290565b5f8083601f840112610cb3575f80fd5b50813567ffffffffffffffff811115610cca575f80fd5b602083019150836020828501011115610ce1575f80fd5b9250929050565b5f805f8060608587031215610cfb575f80fd5b84359350602085013567ffffffffffffffff811115610d18575f80fd5b610d2487828801610ca3565b9598909750949560400135949350505050565b5f805f805f60808688031215610d4b575f80fd5b85359450602086013567ffffffffffffffff811115610d68575f80fd5b610d7488828901610ca3565b9699909850959660408101359660609091013595509350505050565b5f805f60408486031215610da2575f80fd5b83359250602084013567ffffffffffffffff811115610dbf575f80fd5b610dcb86828701610ca3565b9497909650939450505050565b5f60208284031215610de8575f80fd5b5051919050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b604081525f610e2a604083018587610def565b9050826020830152949350505050565b606081525f610e4d606083018688610def565b6020830194909452506040015292915050565b838152604060208201525f610e79604083018486610def565b9594505050505056fea164736f6c6343000818000a",
}

// CsexitpenaltiesABI is the input ABI used to generate the binding from.
// Deprecated: Use CsexitpenaltiesMetaData.ABI instead.
var CsexitpenaltiesABI = CsexitpenaltiesMetaData.ABI

// CsexitpenaltiesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CsexitpenaltiesMetaData.Bin instead.
var CsexitpenaltiesBin = CsexitpenaltiesMetaData.Bin

// DeployCsexitpenalties deploys a new Ethereum contract, binding an instance of Csexitpenalties to it.
func DeployCsexitpenalties(auth *bind.TransactOpts, backend bind.ContractBackend, module common.Address, parametersRegistry common.Address, strikes common.Address) (common.Address, *types.Transaction, *Csexitpenalties, error) {
	parsed, err := CsexitpenaltiesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CsexitpenaltiesBin), backend, module, parametersRegistry, strikes)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Csexitpenalties{CsexitpenaltiesCaller: CsexitpenaltiesCaller{contract: contract}, CsexitpenaltiesTransactor: CsexitpenaltiesTransactor{contract: contract}, CsexitpenaltiesFilterer: CsexitpenaltiesFilterer{contract: contract}}, nil
}

// Csexitpenalties is an auto generated Go binding around an Ethereum contract.
type Csexitpenalties struct {
	CsexitpenaltiesCaller     // Read-only binding to the contract
	CsexitpenaltiesTransactor // Write-only binding to the contract
	CsexitpenaltiesFilterer   // Log filterer for contract events
}

// CsexitpenaltiesCaller is an auto generated read-only Go binding around an Ethereum contract.
type CsexitpenaltiesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsexitpenaltiesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CsexitpenaltiesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsexitpenaltiesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CsexitpenaltiesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsexitpenaltiesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CsexitpenaltiesSession struct {
	Contract     *Csexitpenalties  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CsexitpenaltiesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CsexitpenaltiesCallerSession struct {
	Contract *CsexitpenaltiesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CsexitpenaltiesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CsexitpenaltiesTransactorSession struct {
	Contract     *CsexitpenaltiesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CsexitpenaltiesRaw is an auto generated low-level Go binding around an Ethereum contract.
type CsexitpenaltiesRaw struct {
	Contract *Csexitpenalties // Generic contract binding to access the raw methods on
}

// CsexitpenaltiesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CsexitpenaltiesCallerRaw struct {
	Contract *CsexitpenaltiesCaller // Generic read-only contract binding to access the raw methods on
}

// CsexitpenaltiesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CsexitpenaltiesTransactorRaw struct {
	Contract *CsexitpenaltiesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCsexitpenalties creates a new instance of Csexitpenalties, bound to a specific deployed contract.
func NewCsexitpenalties(address common.Address, backend bind.ContractBackend) (*Csexitpenalties, error) {
	contract, err := bindCsexitpenalties(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Csexitpenalties{CsexitpenaltiesCaller: CsexitpenaltiesCaller{contract: contract}, CsexitpenaltiesTransactor: CsexitpenaltiesTransactor{contract: contract}, CsexitpenaltiesFilterer: CsexitpenaltiesFilterer{contract: contract}}, nil
}

// NewCsexitpenaltiesCaller creates a new read-only instance of Csexitpenalties, bound to a specific deployed contract.
func NewCsexitpenaltiesCaller(address common.Address, caller bind.ContractCaller) (*CsexitpenaltiesCaller, error) {
	contract, err := bindCsexitpenalties(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CsexitpenaltiesCaller{contract: contract}, nil
}

// NewCsexitpenaltiesTransactor creates a new write-only instance of Csexitpenalties, bound to a specific deployed contract.
func NewCsexitpenaltiesTransactor(address common.Address, transactor bind.ContractTransactor) (*CsexitpenaltiesTransactor, error) {
	contract, err := bindCsexitpenalties(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CsexitpenaltiesTransactor{contract: contract}, nil
}

// NewCsexitpenaltiesFilterer creates a new log filterer instance of Csexitpenalties, bound to a specific deployed contract.
func NewCsexitpenaltiesFilterer(address common.Address, filterer bind.ContractFilterer) (*CsexitpenaltiesFilterer, error) {
	contract, err := bindCsexitpenalties(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CsexitpenaltiesFilterer{contract: contract}, nil
}

// bindCsexitpenalties binds a generic wrapper to an already deployed contract.
func bindCsexitpenalties(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CsexitpenaltiesMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csexitpenalties *CsexitpenaltiesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csexitpenalties.Contract.CsexitpenaltiesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csexitpenalties *CsexitpenaltiesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csexitpenalties.Contract.CsexitpenaltiesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csexitpenalties *CsexitpenaltiesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csexitpenalties.Contract.CsexitpenaltiesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csexitpenalties *CsexitpenaltiesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csexitpenalties.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csexitpenalties *CsexitpenaltiesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csexitpenalties.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csexitpenalties *CsexitpenaltiesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csexitpenalties.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Csexitpenalties *CsexitpenaltiesCaller) ACCOUNTING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csexitpenalties.contract.Call(opts, &out, "ACCOUNTING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Csexitpenalties *CsexitpenaltiesSession) ACCOUNTING() (common.Address, error) {
	return _Csexitpenalties.Contract.ACCOUNTING(&_Csexitpenalties.CallOpts)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Csexitpenalties *CsexitpenaltiesCallerSession) ACCOUNTING() (common.Address, error) {
	return _Csexitpenalties.Contract.ACCOUNTING(&_Csexitpenalties.CallOpts)
}

// MODULE is a free data retrieval call binding the contract method 0x094d3a34.
//
// Solidity: function MODULE() view returns(address)
func (_Csexitpenalties *CsexitpenaltiesCaller) MODULE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csexitpenalties.contract.Call(opts, &out, "MODULE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULE is a free data retrieval call binding the contract method 0x094d3a34.
//
// Solidity: function MODULE() view returns(address)
func (_Csexitpenalties *CsexitpenaltiesSession) MODULE() (common.Address, error) {
	return _Csexitpenalties.Contract.MODULE(&_Csexitpenalties.CallOpts)
}

// MODULE is a free data retrieval call binding the contract method 0x094d3a34.
//
// Solidity: function MODULE() view returns(address)
func (_Csexitpenalties *CsexitpenaltiesCallerSession) MODULE() (common.Address, error) {
	return _Csexitpenalties.Contract.MODULE(&_Csexitpenalties.CallOpts)
}

// PARAMETERSREGISTRY is a free data retrieval call binding the contract method 0x2fc88741.
//
// Solidity: function PARAMETERS_REGISTRY() view returns(address)
func (_Csexitpenalties *CsexitpenaltiesCaller) PARAMETERSREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csexitpenalties.contract.Call(opts, &out, "PARAMETERS_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PARAMETERSREGISTRY is a free data retrieval call binding the contract method 0x2fc88741.
//
// Solidity: function PARAMETERS_REGISTRY() view returns(address)
func (_Csexitpenalties *CsexitpenaltiesSession) PARAMETERSREGISTRY() (common.Address, error) {
	return _Csexitpenalties.Contract.PARAMETERSREGISTRY(&_Csexitpenalties.CallOpts)
}

// PARAMETERSREGISTRY is a free data retrieval call binding the contract method 0x2fc88741.
//
// Solidity: function PARAMETERS_REGISTRY() view returns(address)
func (_Csexitpenalties *CsexitpenaltiesCallerSession) PARAMETERSREGISTRY() (common.Address, error) {
	return _Csexitpenalties.Contract.PARAMETERSREGISTRY(&_Csexitpenalties.CallOpts)
}

// STRIKES is a free data retrieval call binding the contract method 0x8af30142.
//
// Solidity: function STRIKES() view returns(address)
func (_Csexitpenalties *CsexitpenaltiesCaller) STRIKES(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csexitpenalties.contract.Call(opts, &out, "STRIKES")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STRIKES is a free data retrieval call binding the contract method 0x8af30142.
//
// Solidity: function STRIKES() view returns(address)
func (_Csexitpenalties *CsexitpenaltiesSession) STRIKES() (common.Address, error) {
	return _Csexitpenalties.Contract.STRIKES(&_Csexitpenalties.CallOpts)
}

// STRIKES is a free data retrieval call binding the contract method 0x8af30142.
//
// Solidity: function STRIKES() view returns(address)
func (_Csexitpenalties *CsexitpenaltiesCallerSession) STRIKES() (common.Address, error) {
	return _Csexitpenalties.Contract.STRIKES(&_Csexitpenalties.CallOpts)
}

// STRIKESEXITTYPEID is a free data retrieval call binding the contract method 0x320eacf8.
//
// Solidity: function STRIKES_EXIT_TYPE_ID() view returns(uint8)
func (_Csexitpenalties *CsexitpenaltiesCaller) STRIKESEXITTYPEID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Csexitpenalties.contract.Call(opts, &out, "STRIKES_EXIT_TYPE_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STRIKESEXITTYPEID is a free data retrieval call binding the contract method 0x320eacf8.
//
// Solidity: function STRIKES_EXIT_TYPE_ID() view returns(uint8)
func (_Csexitpenalties *CsexitpenaltiesSession) STRIKESEXITTYPEID() (uint8, error) {
	return _Csexitpenalties.Contract.STRIKESEXITTYPEID(&_Csexitpenalties.CallOpts)
}

// STRIKESEXITTYPEID is a free data retrieval call binding the contract method 0x320eacf8.
//
// Solidity: function STRIKES_EXIT_TYPE_ID() view returns(uint8)
func (_Csexitpenalties *CsexitpenaltiesCallerSession) STRIKESEXITTYPEID() (uint8, error) {
	return _Csexitpenalties.Contract.STRIKESEXITTYPEID(&_Csexitpenalties.CallOpts)
}

// VOLUNTARYEXITTYPEID is a free data retrieval call binding the contract method 0x3b596df5.
//
// Solidity: function VOLUNTARY_EXIT_TYPE_ID() view returns(uint8)
func (_Csexitpenalties *CsexitpenaltiesCaller) VOLUNTARYEXITTYPEID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Csexitpenalties.contract.Call(opts, &out, "VOLUNTARY_EXIT_TYPE_ID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// VOLUNTARYEXITTYPEID is a free data retrieval call binding the contract method 0x3b596df5.
//
// Solidity: function VOLUNTARY_EXIT_TYPE_ID() view returns(uint8)
func (_Csexitpenalties *CsexitpenaltiesSession) VOLUNTARYEXITTYPEID() (uint8, error) {
	return _Csexitpenalties.Contract.VOLUNTARYEXITTYPEID(&_Csexitpenalties.CallOpts)
}

// VOLUNTARYEXITTYPEID is a free data retrieval call binding the contract method 0x3b596df5.
//
// Solidity: function VOLUNTARY_EXIT_TYPE_ID() view returns(uint8)
func (_Csexitpenalties *CsexitpenaltiesCallerSession) VOLUNTARYEXITTYPEID() (uint8, error) {
	return _Csexitpenalties.Contract.VOLUNTARYEXITTYPEID(&_Csexitpenalties.CallOpts)
}

// GetExitPenaltyInfo is a free data retrieval call binding the contract method 0xe83ba79d.
//
// Solidity: function getExitPenaltyInfo(uint256 nodeOperatorId, bytes publicKey) view returns(((uint248,bool),(uint248,bool),(uint248,bool)))
func (_Csexitpenalties *CsexitpenaltiesCaller) GetExitPenaltyInfo(opts *bind.CallOpts, nodeOperatorId *big.Int, publicKey []byte) (ExitPenaltyInfo, error) {
	var out []interface{}
	err := _Csexitpenalties.contract.Call(opts, &out, "getExitPenaltyInfo", nodeOperatorId, publicKey)

	if err != nil {
		return *new(ExitPenaltyInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(ExitPenaltyInfo)).(*ExitPenaltyInfo)

	return out0, err

}

// GetExitPenaltyInfo is a free data retrieval call binding the contract method 0xe83ba79d.
//
// Solidity: function getExitPenaltyInfo(uint256 nodeOperatorId, bytes publicKey) view returns(((uint248,bool),(uint248,bool),(uint248,bool)))
func (_Csexitpenalties *CsexitpenaltiesSession) GetExitPenaltyInfo(nodeOperatorId *big.Int, publicKey []byte) (ExitPenaltyInfo, error) {
	return _Csexitpenalties.Contract.GetExitPenaltyInfo(&_Csexitpenalties.CallOpts, nodeOperatorId, publicKey)
}

// GetExitPenaltyInfo is a free data retrieval call binding the contract method 0xe83ba79d.
//
// Solidity: function getExitPenaltyInfo(uint256 nodeOperatorId, bytes publicKey) view returns(((uint248,bool),(uint248,bool),(uint248,bool)))
func (_Csexitpenalties *CsexitpenaltiesCallerSession) GetExitPenaltyInfo(nodeOperatorId *big.Int, publicKey []byte) (ExitPenaltyInfo, error) {
	return _Csexitpenalties.Contract.GetExitPenaltyInfo(&_Csexitpenalties.CallOpts, nodeOperatorId, publicKey)
}

// IsValidatorExitDelayPenaltyApplicable is a free data retrieval call binding the contract method 0xd4040379.
//
// Solidity: function isValidatorExitDelayPenaltyApplicable(uint256 nodeOperatorId, bytes publicKey, uint256 eligibleToExitInSec) view returns(bool)
func (_Csexitpenalties *CsexitpenaltiesCaller) IsValidatorExitDelayPenaltyApplicable(opts *bind.CallOpts, nodeOperatorId *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (bool, error) {
	var out []interface{}
	err := _Csexitpenalties.contract.Call(opts, &out, "isValidatorExitDelayPenaltyApplicable", nodeOperatorId, publicKey, eligibleToExitInSec)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorExitDelayPenaltyApplicable is a free data retrieval call binding the contract method 0xd4040379.
//
// Solidity: function isValidatorExitDelayPenaltyApplicable(uint256 nodeOperatorId, bytes publicKey, uint256 eligibleToExitInSec) view returns(bool)
func (_Csexitpenalties *CsexitpenaltiesSession) IsValidatorExitDelayPenaltyApplicable(nodeOperatorId *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (bool, error) {
	return _Csexitpenalties.Contract.IsValidatorExitDelayPenaltyApplicable(&_Csexitpenalties.CallOpts, nodeOperatorId, publicKey, eligibleToExitInSec)
}

// IsValidatorExitDelayPenaltyApplicable is a free data retrieval call binding the contract method 0xd4040379.
//
// Solidity: function isValidatorExitDelayPenaltyApplicable(uint256 nodeOperatorId, bytes publicKey, uint256 eligibleToExitInSec) view returns(bool)
func (_Csexitpenalties *CsexitpenaltiesCallerSession) IsValidatorExitDelayPenaltyApplicable(nodeOperatorId *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (bool, error) {
	return _Csexitpenalties.Contract.IsValidatorExitDelayPenaltyApplicable(&_Csexitpenalties.CallOpts, nodeOperatorId, publicKey, eligibleToExitInSec)
}

// ProcessExitDelayReport is a paid mutator transaction binding the contract method 0x44dab949.
//
// Solidity: function processExitDelayReport(uint256 nodeOperatorId, bytes publicKey, uint256 eligibleToExitInSec) returns()
func (_Csexitpenalties *CsexitpenaltiesTransactor) ProcessExitDelayReport(opts *bind.TransactOpts, nodeOperatorId *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (*types.Transaction, error) {
	return _Csexitpenalties.contract.Transact(opts, "processExitDelayReport", nodeOperatorId, publicKey, eligibleToExitInSec)
}

// ProcessExitDelayReport is a paid mutator transaction binding the contract method 0x44dab949.
//
// Solidity: function processExitDelayReport(uint256 nodeOperatorId, bytes publicKey, uint256 eligibleToExitInSec) returns()
func (_Csexitpenalties *CsexitpenaltiesSession) ProcessExitDelayReport(nodeOperatorId *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (*types.Transaction, error) {
	return _Csexitpenalties.Contract.ProcessExitDelayReport(&_Csexitpenalties.TransactOpts, nodeOperatorId, publicKey, eligibleToExitInSec)
}

// ProcessExitDelayReport is a paid mutator transaction binding the contract method 0x44dab949.
//
// Solidity: function processExitDelayReport(uint256 nodeOperatorId, bytes publicKey, uint256 eligibleToExitInSec) returns()
func (_Csexitpenalties *CsexitpenaltiesTransactorSession) ProcessExitDelayReport(nodeOperatorId *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (*types.Transaction, error) {
	return _Csexitpenalties.Contract.ProcessExitDelayReport(&_Csexitpenalties.TransactOpts, nodeOperatorId, publicKey, eligibleToExitInSec)
}

// ProcessStrikesReport is a paid mutator transaction binding the contract method 0xe9f6fdc6.
//
// Solidity: function processStrikesReport(uint256 nodeOperatorId, bytes publicKey) returns()
func (_Csexitpenalties *CsexitpenaltiesTransactor) ProcessStrikesReport(opts *bind.TransactOpts, nodeOperatorId *big.Int, publicKey []byte) (*types.Transaction, error) {
	return _Csexitpenalties.contract.Transact(opts, "processStrikesReport", nodeOperatorId, publicKey)
}

// ProcessStrikesReport is a paid mutator transaction binding the contract method 0xe9f6fdc6.
//
// Solidity: function processStrikesReport(uint256 nodeOperatorId, bytes publicKey) returns()
func (_Csexitpenalties *CsexitpenaltiesSession) ProcessStrikesReport(nodeOperatorId *big.Int, publicKey []byte) (*types.Transaction, error) {
	return _Csexitpenalties.Contract.ProcessStrikesReport(&_Csexitpenalties.TransactOpts, nodeOperatorId, publicKey)
}

// ProcessStrikesReport is a paid mutator transaction binding the contract method 0xe9f6fdc6.
//
// Solidity: function processStrikesReport(uint256 nodeOperatorId, bytes publicKey) returns()
func (_Csexitpenalties *CsexitpenaltiesTransactorSession) ProcessStrikesReport(nodeOperatorId *big.Int, publicKey []byte) (*types.Transaction, error) {
	return _Csexitpenalties.Contract.ProcessStrikesReport(&_Csexitpenalties.TransactOpts, nodeOperatorId, publicKey)
}

// ProcessTriggeredExit is a paid mutator transaction binding the contract method 0x848b1dec.
//
// Solidity: function processTriggeredExit(uint256 nodeOperatorId, bytes publicKey, uint256 withdrawalRequestPaidFee, uint256 exitType) returns()
func (_Csexitpenalties *CsexitpenaltiesTransactor) ProcessTriggeredExit(opts *bind.TransactOpts, nodeOperatorId *big.Int, publicKey []byte, withdrawalRequestPaidFee *big.Int, exitType *big.Int) (*types.Transaction, error) {
	return _Csexitpenalties.contract.Transact(opts, "processTriggeredExit", nodeOperatorId, publicKey, withdrawalRequestPaidFee, exitType)
}

// ProcessTriggeredExit is a paid mutator transaction binding the contract method 0x848b1dec.
//
// Solidity: function processTriggeredExit(uint256 nodeOperatorId, bytes publicKey, uint256 withdrawalRequestPaidFee, uint256 exitType) returns()
func (_Csexitpenalties *CsexitpenaltiesSession) ProcessTriggeredExit(nodeOperatorId *big.Int, publicKey []byte, withdrawalRequestPaidFee *big.Int, exitType *big.Int) (*types.Transaction, error) {
	return _Csexitpenalties.Contract.ProcessTriggeredExit(&_Csexitpenalties.TransactOpts, nodeOperatorId, publicKey, withdrawalRequestPaidFee, exitType)
}

// ProcessTriggeredExit is a paid mutator transaction binding the contract method 0x848b1dec.
//
// Solidity: function processTriggeredExit(uint256 nodeOperatorId, bytes publicKey, uint256 withdrawalRequestPaidFee, uint256 exitType) returns()
func (_Csexitpenalties *CsexitpenaltiesTransactorSession) ProcessTriggeredExit(nodeOperatorId *big.Int, publicKey []byte, withdrawalRequestPaidFee *big.Int, exitType *big.Int) (*types.Transaction, error) {
	return _Csexitpenalties.Contract.ProcessTriggeredExit(&_Csexitpenalties.TransactOpts, nodeOperatorId, publicKey, withdrawalRequestPaidFee, exitType)
}

// CsexitpenaltiesStrikesPenaltyProcessedIterator is returned from FilterStrikesPenaltyProcessed and is used to iterate over the raw logs and unpacked data for StrikesPenaltyProcessed events raised by the Csexitpenalties contract.
type CsexitpenaltiesStrikesPenaltyProcessedIterator struct {
	Event *CsexitpenaltiesStrikesPenaltyProcessed // Event containing the contract specifics and raw log

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
func (it *CsexitpenaltiesStrikesPenaltyProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsexitpenaltiesStrikesPenaltyProcessed)
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
		it.Event = new(CsexitpenaltiesStrikesPenaltyProcessed)
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
func (it *CsexitpenaltiesStrikesPenaltyProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsexitpenaltiesStrikesPenaltyProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsexitpenaltiesStrikesPenaltyProcessed represents a StrikesPenaltyProcessed event raised by the Csexitpenalties contract.
type CsexitpenaltiesStrikesPenaltyProcessed struct {
	NodeOperatorId *big.Int
	Pubkey         []byte
	StrikesPenalty *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStrikesPenaltyProcessed is a free log retrieval operation binding the contract event 0x8b2d35b9e1a65aeadb4b39934801c82d77a6526706826087b9e47a7ce005d405.
//
// Solidity: event StrikesPenaltyProcessed(uint256 indexed nodeOperatorId, bytes pubkey, uint256 strikesPenalty)
func (_Csexitpenalties *CsexitpenaltiesFilterer) FilterStrikesPenaltyProcessed(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsexitpenaltiesStrikesPenaltyProcessedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csexitpenalties.contract.FilterLogs(opts, "StrikesPenaltyProcessed", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsexitpenaltiesStrikesPenaltyProcessedIterator{contract: _Csexitpenalties.contract, event: "StrikesPenaltyProcessed", logs: logs, sub: sub}, nil
}

// WatchStrikesPenaltyProcessed is a free log subscription operation binding the contract event 0x8b2d35b9e1a65aeadb4b39934801c82d77a6526706826087b9e47a7ce005d405.
//
// Solidity: event StrikesPenaltyProcessed(uint256 indexed nodeOperatorId, bytes pubkey, uint256 strikesPenalty)
func (_Csexitpenalties *CsexitpenaltiesFilterer) WatchStrikesPenaltyProcessed(opts *bind.WatchOpts, sink chan<- *CsexitpenaltiesStrikesPenaltyProcessed, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csexitpenalties.contract.WatchLogs(opts, "StrikesPenaltyProcessed", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsexitpenaltiesStrikesPenaltyProcessed)
				if err := _Csexitpenalties.contract.UnpackLog(event, "StrikesPenaltyProcessed", log); err != nil {
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

// ParseStrikesPenaltyProcessed is a log parse operation binding the contract event 0x8b2d35b9e1a65aeadb4b39934801c82d77a6526706826087b9e47a7ce005d405.
//
// Solidity: event StrikesPenaltyProcessed(uint256 indexed nodeOperatorId, bytes pubkey, uint256 strikesPenalty)
func (_Csexitpenalties *CsexitpenaltiesFilterer) ParseStrikesPenaltyProcessed(log types.Log) (*CsexitpenaltiesStrikesPenaltyProcessed, error) {
	event := new(CsexitpenaltiesStrikesPenaltyProcessed)
	if err := _Csexitpenalties.contract.UnpackLog(event, "StrikesPenaltyProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsexitpenaltiesTriggeredExitFeeRecordedIterator is returned from FilterTriggeredExitFeeRecorded and is used to iterate over the raw logs and unpacked data for TriggeredExitFeeRecorded events raised by the Csexitpenalties contract.
type CsexitpenaltiesTriggeredExitFeeRecordedIterator struct {
	Event *CsexitpenaltiesTriggeredExitFeeRecorded // Event containing the contract specifics and raw log

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
func (it *CsexitpenaltiesTriggeredExitFeeRecordedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsexitpenaltiesTriggeredExitFeeRecorded)
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
		it.Event = new(CsexitpenaltiesTriggeredExitFeeRecorded)
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
func (it *CsexitpenaltiesTriggeredExitFeeRecordedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsexitpenaltiesTriggeredExitFeeRecordedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsexitpenaltiesTriggeredExitFeeRecorded represents a TriggeredExitFeeRecorded event raised by the Csexitpenalties contract.
type CsexitpenaltiesTriggeredExitFeeRecorded struct {
	NodeOperatorId               *big.Int
	ExitType                     *big.Int
	Pubkey                       []byte
	WithdrawalRequestPaidFee     *big.Int
	WithdrawalRequestRecordedFee *big.Int
	Raw                          types.Log // Blockchain specific contextual infos
}

// FilterTriggeredExitFeeRecorded is a free log retrieval operation binding the contract event 0xc90027a0742ed6a64a8829261af4ad3d833e668083a49a0634cc61c110c8e8db.
//
// Solidity: event TriggeredExitFeeRecorded(uint256 indexed nodeOperatorId, uint256 indexed exitType, bytes pubkey, uint256 withdrawalRequestPaidFee, uint256 withdrawalRequestRecordedFee)
func (_Csexitpenalties *CsexitpenaltiesFilterer) FilterTriggeredExitFeeRecorded(opts *bind.FilterOpts, nodeOperatorId []*big.Int, exitType []*big.Int) (*CsexitpenaltiesTriggeredExitFeeRecordedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var exitTypeRule []interface{}
	for _, exitTypeItem := range exitType {
		exitTypeRule = append(exitTypeRule, exitTypeItem)
	}

	logs, sub, err := _Csexitpenalties.contract.FilterLogs(opts, "TriggeredExitFeeRecorded", nodeOperatorIdRule, exitTypeRule)
	if err != nil {
		return nil, err
	}
	return &CsexitpenaltiesTriggeredExitFeeRecordedIterator{contract: _Csexitpenalties.contract, event: "TriggeredExitFeeRecorded", logs: logs, sub: sub}, nil
}

// WatchTriggeredExitFeeRecorded is a free log subscription operation binding the contract event 0xc90027a0742ed6a64a8829261af4ad3d833e668083a49a0634cc61c110c8e8db.
//
// Solidity: event TriggeredExitFeeRecorded(uint256 indexed nodeOperatorId, uint256 indexed exitType, bytes pubkey, uint256 withdrawalRequestPaidFee, uint256 withdrawalRequestRecordedFee)
func (_Csexitpenalties *CsexitpenaltiesFilterer) WatchTriggeredExitFeeRecorded(opts *bind.WatchOpts, sink chan<- *CsexitpenaltiesTriggeredExitFeeRecorded, nodeOperatorId []*big.Int, exitType []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var exitTypeRule []interface{}
	for _, exitTypeItem := range exitType {
		exitTypeRule = append(exitTypeRule, exitTypeItem)
	}

	logs, sub, err := _Csexitpenalties.contract.WatchLogs(opts, "TriggeredExitFeeRecorded", nodeOperatorIdRule, exitTypeRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsexitpenaltiesTriggeredExitFeeRecorded)
				if err := _Csexitpenalties.contract.UnpackLog(event, "TriggeredExitFeeRecorded", log); err != nil {
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

// ParseTriggeredExitFeeRecorded is a log parse operation binding the contract event 0xc90027a0742ed6a64a8829261af4ad3d833e668083a49a0634cc61c110c8e8db.
//
// Solidity: event TriggeredExitFeeRecorded(uint256 indexed nodeOperatorId, uint256 indexed exitType, bytes pubkey, uint256 withdrawalRequestPaidFee, uint256 withdrawalRequestRecordedFee)
func (_Csexitpenalties *CsexitpenaltiesFilterer) ParseTriggeredExitFeeRecorded(log types.Log) (*CsexitpenaltiesTriggeredExitFeeRecorded, error) {
	event := new(CsexitpenaltiesTriggeredExitFeeRecorded)
	if err := _Csexitpenalties.contract.UnpackLog(event, "TriggeredExitFeeRecorded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsexitpenaltiesValidatorExitDelayProcessedIterator is returned from FilterValidatorExitDelayProcessed and is used to iterate over the raw logs and unpacked data for ValidatorExitDelayProcessed events raised by the Csexitpenalties contract.
type CsexitpenaltiesValidatorExitDelayProcessedIterator struct {
	Event *CsexitpenaltiesValidatorExitDelayProcessed // Event containing the contract specifics and raw log

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
func (it *CsexitpenaltiesValidatorExitDelayProcessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsexitpenaltiesValidatorExitDelayProcessed)
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
		it.Event = new(CsexitpenaltiesValidatorExitDelayProcessed)
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
func (it *CsexitpenaltiesValidatorExitDelayProcessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsexitpenaltiesValidatorExitDelayProcessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsexitpenaltiesValidatorExitDelayProcessed represents a ValidatorExitDelayProcessed event raised by the Csexitpenalties contract.
type CsexitpenaltiesValidatorExitDelayProcessed struct {
	NodeOperatorId *big.Int
	Pubkey         []byte
	DelayPenalty   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterValidatorExitDelayProcessed is a free log retrieval operation binding the contract event 0xf808c54013437847ff14496ccbb2d51171fc03bb72a783b7905a9dc648757009.
//
// Solidity: event ValidatorExitDelayProcessed(uint256 indexed nodeOperatorId, bytes pubkey, uint256 delayPenalty)
func (_Csexitpenalties *CsexitpenaltiesFilterer) FilterValidatorExitDelayProcessed(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsexitpenaltiesValidatorExitDelayProcessedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csexitpenalties.contract.FilterLogs(opts, "ValidatorExitDelayProcessed", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsexitpenaltiesValidatorExitDelayProcessedIterator{contract: _Csexitpenalties.contract, event: "ValidatorExitDelayProcessed", logs: logs, sub: sub}, nil
}

// WatchValidatorExitDelayProcessed is a free log subscription operation binding the contract event 0xf808c54013437847ff14496ccbb2d51171fc03bb72a783b7905a9dc648757009.
//
// Solidity: event ValidatorExitDelayProcessed(uint256 indexed nodeOperatorId, bytes pubkey, uint256 delayPenalty)
func (_Csexitpenalties *CsexitpenaltiesFilterer) WatchValidatorExitDelayProcessed(opts *bind.WatchOpts, sink chan<- *CsexitpenaltiesValidatorExitDelayProcessed, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csexitpenalties.contract.WatchLogs(opts, "ValidatorExitDelayProcessed", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsexitpenaltiesValidatorExitDelayProcessed)
				if err := _Csexitpenalties.contract.UnpackLog(event, "ValidatorExitDelayProcessed", log); err != nil {
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

// ParseValidatorExitDelayProcessed is a log parse operation binding the contract event 0xf808c54013437847ff14496ccbb2d51171fc03bb72a783b7905a9dc648757009.
//
// Solidity: event ValidatorExitDelayProcessed(uint256 indexed nodeOperatorId, bytes pubkey, uint256 delayPenalty)
func (_Csexitpenalties *CsexitpenaltiesFilterer) ParseValidatorExitDelayProcessed(log types.Log) (*CsexitpenaltiesValidatorExitDelayProcessed, error) {
	event := new(CsexitpenaltiesValidatorExitDelayProcessed)
	if err := _Csexitpenalties.contract.UnpackLog(event, "ValidatorExitDelayProcessed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
