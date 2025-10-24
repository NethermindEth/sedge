// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package csfeedistributor

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

// ICSFeeDistributorDistributionData is an auto generated low-level Go binding around an user-defined struct.
type ICSFeeDistributorDistributionData struct {
	RefSlot     *big.Int
	TreeRoot    [32]byte
	TreeCid     string
	LogCid      string
	Distributed *big.Int
	Rebate      *big.Int
}

// CsfeedistributorMetaData contains all meta data concerning the Csfeedistributor contract.
var CsfeedistributorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"stETH\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"accounting\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeeSharesDecrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLogCID\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidProof\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReportData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTreeCid\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTreeRoot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToRecover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotAccounting\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotOracle\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAccountingAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroOracleAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroRebateRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroStEthAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalClaimableShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"treeRoot\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"treeCid\",\"type\":\"string\"}],\"name\":\"DistributionDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"logCid\",\"type\":\"string\"}],\"name\":\"DistributionLogUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"ModuleFeeDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"OperatorFeeDistributed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"RebateRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"RebateTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"StETHSharesRecovered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACCOUNTING\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECOVERER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STETH\",\"outputs\":[{\"internalType\":\"contractIStETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"distributeFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesToDistribute\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"distributedShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"distributed\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"distributionDataHistoryCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rebateRecipient\",\"type\":\"address\"}],\"name\":\"finalizeUpgradeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"proof\",\"type\":\"bytes32[]\"}],\"name\":\"getFeesToDistribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sharesToDistribute\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getHistoricalDistributionData\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"treeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"treeCid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"logCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"distributed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rebate\",\"type\":\"uint256\"}],\"internalType\":\"structICSFeeDistributor.DistributionData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitializedVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"hashLeaf\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_rebateRecipient\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"logCid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingSharesToDistribute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_treeRoot\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_treeCid\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_logCid\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"distributed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rebate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"processOracleReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rebateRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rebateRecipient\",\"type\":\"address\"}],\"name\":\"setRebateRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalClaimableShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeCid\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"treeRoot\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f80fd5b50600436106101fd575f3560e01c80636f962e5c11610114578063b3c65015116100a9578063d5ba2dcf11610079578063d5ba2dcf14610498578063e00bfe50146104ab578063e877f068146104d2578063ea6301ab146104e5578063fe3c9b9b14610504575f80fd5b8063b3c6501514610449578063ca15c8731461046a578063d257cf2a1461047d578063d547741f14610485575f80fd5b80639010d07c116100e45780639010d07c146103f557806391d1485414610408578063a217fddf1461041b578063acf1c94814610422575f80fd5b80636f962e5c146103a75780637e9f27ad146103bc578063819d4cc6146103cf5780638980f11f146103e2575f80fd5b806338013f02116101955780634e5b3a62116101655780634e5b3a621461033257806352d8bfc2146103525780635c654ad91461035a5780635e8e8f6f1461036d5780636dc3f2bd14610380575f80fd5b806338013f02146102e65780633d18b6f31461030d57806347d17d9d14610316578063485cc9551461031f575f80fd5b80632f2ff15d116101d05780632f2ff15d146102805780632ffa14e1146102955780633333e109146102a857806336568abe146102d3575f80fd5b806301ffc9a71461020157806314dc6c141461022957806321893f7b1461023f578063248a9ca314610252575b5f80fd5b61021461020f366004611a81565b61050c565b60405190151581526020015b60405180910390f35b6102315f5481565b604051908152602001610220565b61023161024d366004611aa8565b610536565b610231610260366004611b24565b5f9081525f80516020612106833981519152602052604090206001015490565b61029361028e366004611b56565b6106df565b005b6102936102a3366004611bc5565b610715565b6007546102bb906001600160a01b031681565b6040516001600160a01b039091168152602001610220565b6102936102e1366004611b56565b610c65565b6102bb7f0000000000000000000000004d4074628678bd302921c20573eea1ed38ddf7fb81565b61023160065481565b61023160045481565b61029361032d366004611c4f565b610c9d565b610345610340366004611b24565b610db8565b6040516102209190611cba565b610293610f5a565b610293610368366004611d27565b610fb6565b61023161037b366004611aa8565b611031565b6102bb7f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da81565b6103af6110c4565b6040516102209190611d4f565b6102316103ca366004611d61565b611150565b6102936103dd366004611d27565b6111a1565b6102936103f0366004611d27565b6111f0565b6102bb610403366004611d61565b611291565b610214610416366004611b56565b6112c9565b6102315f81565b6102317fb3e25b5404b87e5a838579cb5d7481d61ad96ee284d38ec1e97c07ba64e7f6fc81565b6104516112ff565b60405167ffffffffffffffff9091168152602001610220565b610231610478366004611b24565b611337565b610231611375565b610293610493366004611b56565b61140d565b6102936104a6366004611d81565b61143d565b6102bb7f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe8481565b6102936104e0366004611d81565b61151d565b6102316104f3366004611b24565b60036020525f908152604090205481565b6103af611534565b5f6001600160e01b03198216635a05180f60e01b1480610530575061053082611541565b92915050565b5f336001600160a01b037f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da16146105805760405163a8d664b560e01b815260040160405180910390fd5b61058c85858585611031565b9050805f0361059c57505f6106d7565b8060045410156105bf57604051633c57b48560e21b815260040160405180910390fd5b6004805482900381555f868152600360205260409081902080548401905551638fcb4e5b60e01b81526001600160a01b037f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe841691638fcb4e5b9161065b917f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da918691016001600160a01b03929092168252602082015260400190565b6020604051808303815f875af1158015610677573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061069b9190611d9a565b50847f4b7ab1c192267e83350d06490a852b8dbbb25bfa00fd065b1862cf7accd2ab90826040516106ce91815260200190565b60405180910390a25b949350505050565b5f8281525f80516020612106833981519152602052604090206001015461070581611575565b61070f8383611582565b50505050565b336001600160a01b037f0000000000000000000000004d4074628678bd302921c20573eea1ed38ddf7fb161461075e576040516312d4786560e01b815260040160405180910390fd5b604051633d7ad0b760e21b81523060048201527f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b03169063f5eb42dc90602401602060405180830381865afa1580156107c0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107e49190611d9a565b82846004546107f39190611dc5565b6107fd9190611dc5565b111561081c57604051636edcc52360e01b815260040160405180910390fd5b8215801561082957505f82115b156108475760405163319c9a2160e21b815260040160405180910390fd5b8215610954575f86900361086e576040516312b7aebf60e01b815260040160405180910390fd5b600160405161087d9190611e10565b60405180910390208787604051610895929190611e82565b6040518091039020036108bb576040516312b7aebf60e01b815260040160405180910390fd5b876108d9576040516357e86a3360e01b815260040160405180910390fd5b5f5488036108fa576040516357e86a3360e01b815260040160405180910390fd5b60048054840190555f8890556001610913878983611ef0565b507f26dec7cc117e9b3907dc1f90d2dc5f6e04dbb9f285f5898be2c82ec524dcd42460045489898960405161094b9493929190611fd2565b60405180910390a15b6040518381527f010f65f5f56ba52d759f7b1dc49a3d277570cc2aa631e9c865b073a0ffc2af419060200160405180910390a18115610a5757600754604051638fcb4e5b60e01b81526001600160a01b039182166004820152602481018490527f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe8490911690638fcb4e5b906044016020604051808303815f875af11580156109fe573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a229190611d9a565b506040518281527f7462935fb42d34d84233f737293310ca24e851021f9cb7f2549470cdf6de56bf9060200160405180910390a15b5f849003610a785760405163526ca52560e01b815260040160405180910390fd5b6002604051610a879190611e10565b60405180910390208585604051610a9f929190611e82565b604051809103902003610ac55760405163526ca52560e01b815260040160405180910390fd5b6002610ad2858783611ef0565b507f1f1a488b71a099a0d9cb71f60e14cf90bd1b5b188ca593111a40f533a3130b3b8585604051610b04929190611ffb565b60405180910390a16040518060c001604052808281526020015f54815260200160018054610b3190611dd8565b80601f0160208091040260200160405190810160405280929190818152602001828054610b5d90611dd8565b8015610ba85780601f10610b7f57610100808354040283529160200191610ba8565b820191905f5260205f20905b815481529060010190602001808311610b8b57829003601f168201915b5050505050815260200186868080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920182905250938552505050602080830187905260409283018690526006548252600581529082902083518155908301516001820155908201516002820190610c26908261200e565b5060608201516003820190610c3b908261200e565b506080820151600482015560a0909101516005909101555050600680546001019055505050505050565b6001600160a01b0381163314610c8e5760405163334bd91960e11b815260040160405180910390fd5b610c9882826115d7565b505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff1680610ce75750805467ffffffffffffffff808416911610155b15610d055760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff191667ffffffffffffffff831617600160401b1781556001600160a01b038416610d4e57604051633ef39b8160e01b815260040160405180910390fd5b610d5783611623565b610d5f6116ab565b610d695f85611582565b50805460ff60401b1916815560405167ffffffffffffffff831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050565b610df06040518060c001604052805f81526020015f801916815260200160608152602001606081526020015f81526020015f81525090565b60055f8381526020019081526020015f206040518060c00160405290815f820154815260200160018201548152602001600282018054610e2f90611dd8565b80601f0160208091040260200160405190810160405280929190818152602001828054610e5b90611dd8565b8015610ea65780601f10610e7d57610100808354040283529160200191610ea6565b820191905f5260205f20905b815481529060010190602001808311610e8957829003601f168201915b50505050508152602001600382018054610ebf90611dd8565b80601f0160208091040260200160405190810160405280929190818152602001828054610eeb90611dd8565b8015610f365780601f10610f0d57610100808354040283529160200191610f36565b820191905f5260205f20905b815481529060010190602001808311610f1957829003601f168201915b50505050508152602001600482015481526020016005820154815250509050919050565b610f626116b5565b73a74528edc289b1a597faf83fcff7eff871cc01d96352d8bfc26040518163ffffffff1660e01b81526004015f6040518083038186803b158015610fa4575f80fd5b505af415801561070f573d5f803e3d5ffd5b610fbe6116b5565b604051635c654ad960e01b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d990635c654ad9906044015b5f6040518083038186803b158015611017575f80fd5b505af4158015611029573d5f803e3d5ffd5b505050505050565b5f818103611052576040516309bde33960e01b815260040160405180910390fd5b5f61106984845f546110648a8a611150565b6116de565b905080611089576040516309bde33960e01b815260040160405180910390fd5b5f86815260036020526040902054858111156110b857604051636096ce8160e11b815260040160405180910390fd5b90940395945050505050565b600280546110d190611dd8565b80601f01602080910402602001604051908101604052809291908181526020018280546110fd90611dd8565b80156111485780601f1061111f57610100808354040283529160200191611148565b820191905f5260205f20905b81548152906001019060200180831161112b57829003601f168201915b505050505081565b60408051602081018490529081018290525f9060600160408051601f198184030181528282528051602091820120908301520160405160208183030381529060405280519060200120905092915050565b6111a96116b5565b6040516340cea66360e11b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d99063819d4cc690604401611001565b6111f86116b5565b7f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b0316826001600160a01b03160361124a576040516319efe5d760e21b815260040160405180910390fd5b604051638980f11f60e01b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d990638980f11f90604401611001565b5f8281527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e823717059320006020819052604082206106d790846116f5565b5f9182525f80516020612106833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f6113327ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005467ffffffffffffffff1690565b905090565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200060208190526040822061136e90611700565b9392505050565b60048054604051633d7ad0b760e21b815230928101929092525f917f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b03169063f5eb42dc90602401602060405180830381865afa1580156113df573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906114039190611d9a565b61133291906120ca565b5f8281525f80516020612106833981519152602052604090206001015461143381611575565b61070f83836115d7565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff16806114875750805467ffffffffffffffff808416911610155b156114a55760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff191667ffffffffffffffff831617600160401b1781556114d083611623565b805460ff60401b1916815560405167ffffffffffffffff831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050565b5f61152781611575565b61153082611623565b5050565b600180546110d190611dd8565b5f6001600160e01b03198216637965db0b60e01b148061053057506301ffc9a760e01b6001600160e01b0319831614610530565b61157f8133611709565b50565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000816115af8585611746565b905080156106d7575f8581526020839052604090206115ce90856117e7565b50949350505050565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e823717059320008161160485856117fb565b905080156106d7575f8581526020839052604090206115ce9085611874565b6001600160a01b03811661164a5760405163669766e160e01b815260040160405180910390fd5b6007805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f9f8636d85f90aba9c7d2c9e076c6102a5459d2e063afb71d81328bbb3608a2349060200160405180910390a150565b6116b3611888565b565b6116b37fb3e25b5404b87e5a838579cb5d7481d61ad96ee284d38ec1e97c07ba64e7f6fc611575565b5f826116eb8686856118d1565b1495945050505050565b5f61136e8383611909565b5f610530825490565b61171382826112c9565b6115305760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440160405180910390fd5b5f5f8051602061210683398151915261175f84846112c9565b6117de575f848152602082815260408083206001600160a01b03871684529091529020805460ff191660011790556117943390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050610530565b5f915050610530565b5f61136e836001600160a01b03841661192f565b5f5f8051602061210683398151915261181484846112c9565b156117de575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050610530565b5f61136e836001600160a01b03841661197b565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166116b357604051631afcd79f60e31b815260040160405180910390fd5b5f81815b848110156115ce576118ff828787848181106118f3576118f36120dd565b90506020020135611a55565b91506001016118d5565b5f825f01828154811061191e5761191e6120dd565b905f5260205f200154905092915050565b5f81815260018301602052604081205461197457508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155610530565b505f610530565b5f81815260018301602052604081205480156117de575f61199d6001836120ca565b85549091505f906119b0906001906120ca565b9050808214611a0f575f865f0182815481106119ce576119ce6120dd565b905f5260205f200154905080875f0184815481106119ee576119ee6120dd565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080611a2057611a206120f1565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050610530565b5f818310611a6f575f82815260208490526040902061136e565b5f83815260208390526040902061136e565b5f60208284031215611a91575f80fd5b81356001600160e01b03198116811461136e575f80fd5b5f805f8060608587031215611abb575f80fd5b8435935060208501359250604085013567ffffffffffffffff80821115611ae0575f80fd5b818701915087601f830112611af3575f80fd5b813581811115611b01575f80fd5b8860208260051b8501011115611b15575f80fd5b95989497505060200194505050565b5f60208284031215611b34575f80fd5b5035919050565b80356001600160a01b0381168114611b51575f80fd5b919050565b5f8060408385031215611b67575f80fd5b82359150611b7760208401611b3b565b90509250929050565b5f8083601f840112611b90575f80fd5b50813567ffffffffffffffff811115611ba7575f80fd5b602083019150836020828501011115611bbe575f80fd5b9250929050565b5f805f805f805f8060c0898b031215611bdc575f80fd5b88359750602089013567ffffffffffffffff80821115611bfa575f80fd5b611c068c838d01611b80565b909950975060408b0135915080821115611c1e575f80fd5b50611c2b8b828c01611b80565b999c989b5096999698976060880135976080810135975060a0013595509350505050565b5f8060408385031215611c60575f80fd5b611c6983611b3b565b9150611b7760208401611b3b565b5f81518084525f5b81811015611c9b57602081850181015186830182015201611c7f565b505f602082860101526020601f19601f83011685010191505092915050565b6020815281516020820152602082015160408201525f604083015160c06060840152611ce960e0840182611c77565b90506060840151601f19848303016080850152611d068282611c77565b915050608084015160a084015260a084015160c08401528091505092915050565b5f8060408385031215611d38575f80fd5b611d4183611b3b565b946020939093013593505050565b602081525f61136e6020830184611c77565b5f8060408385031215611d72575f80fd5b50508035926020909101359150565b5f60208284031215611d91575f80fd5b61136e82611b3b565b5f60208284031215611daa575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561053057610530611db1565b600181811c90821680611dec57607f821691505b602082108103611e0a57634e487b7160e01b5f52602260045260245ffd5b50919050565b5f808354611e1d81611dd8565b60018281168015611e355760018114611e4a57611e76565b60ff1984168752821515830287019450611e76565b875f526020805f205f5b85811015611e6d5781548a820152908401908201611e54565b50505082870194505b50929695505050505050565b818382375f9101908152919050565b634e487b7160e01b5f52604160045260245ffd5b601f821115610c9857805f5260205f20601f840160051c81016020851015611eca5750805b601f840160051c820191505b81811015611ee9575f8155600101611ed6565b5050505050565b67ffffffffffffffff831115611f0857611f08611e91565b611f1c83611f168354611dd8565b83611ea5565b5f601f841160018114611f4d575f8515611f365750838201355b5f19600387901b1c1916600186901b178355611ee9565b5f83815260208120601f198716915b82811015611f7c5786850135825560209485019460019092019101611f5c565b5086821015611f98575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b848152836020820152606060408201525f611ff1606083018486611faa565b9695505050505050565b602081525f6106d7602083018486611faa565b815167ffffffffffffffff81111561202857612028611e91565b61203c816120368454611dd8565b84611ea5565b602080601f83116001811461206f575f84156120585750858301515b5f19600386901b1c1916600185901b178555611029565b5f85815260208120601f198616915b8281101561209d5788860151825594840194600190910190840161207e565b50858210156120ba57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b8181038181111561053057610530611db1565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52603160045260245ffdfe02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a164736f6c6343000818000a",
}

// CsfeedistributorABI is the input ABI used to generate the binding from.
// Deprecated: Use CsfeedistributorMetaData.ABI instead.
var CsfeedistributorABI = CsfeedistributorMetaData.ABI

// CsfeedistributorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CsfeedistributorMetaData.Bin instead.
var CsfeedistributorBin = CsfeedistributorMetaData.Bin

// DeployCsfeedistributor deploys a new Ethereum contract, binding an instance of Csfeedistributor to it.
func DeployCsfeedistributor(auth *bind.TransactOpts, backend bind.ContractBackend, stETH common.Address, accounting common.Address, oracle common.Address) (common.Address, *types.Transaction, *Csfeedistributor, error) {
	parsed, err := CsfeedistributorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CsfeedistributorBin), backend, stETH, accounting, oracle)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Csfeedistributor{CsfeedistributorCaller: CsfeedistributorCaller{contract: contract}, CsfeedistributorTransactor: CsfeedistributorTransactor{contract: contract}, CsfeedistributorFilterer: CsfeedistributorFilterer{contract: contract}}, nil
}

// Csfeedistributor is an auto generated Go binding around an Ethereum contract.
type Csfeedistributor struct {
	CsfeedistributorCaller     // Read-only binding to the contract
	CsfeedistributorTransactor // Write-only binding to the contract
	CsfeedistributorFilterer   // Log filterer for contract events
}

// CsfeedistributorCaller is an auto generated read-only Go binding around an Ethereum contract.
type CsfeedistributorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsfeedistributorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CsfeedistributorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsfeedistributorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CsfeedistributorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsfeedistributorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CsfeedistributorSession struct {
	Contract     *Csfeedistributor // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CsfeedistributorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CsfeedistributorCallerSession struct {
	Contract *CsfeedistributorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CsfeedistributorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CsfeedistributorTransactorSession struct {
	Contract     *CsfeedistributorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CsfeedistributorRaw is an auto generated low-level Go binding around an Ethereum contract.
type CsfeedistributorRaw struct {
	Contract *Csfeedistributor // Generic contract binding to access the raw methods on
}

// CsfeedistributorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CsfeedistributorCallerRaw struct {
	Contract *CsfeedistributorCaller // Generic read-only contract binding to access the raw methods on
}

// CsfeedistributorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CsfeedistributorTransactorRaw struct {
	Contract *CsfeedistributorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCsfeedistributor creates a new instance of Csfeedistributor, bound to a specific deployed contract.
func NewCsfeedistributor(address common.Address, backend bind.ContractBackend) (*Csfeedistributor, error) {
	contract, err := bindCsfeedistributor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Csfeedistributor{CsfeedistributorCaller: CsfeedistributorCaller{contract: contract}, CsfeedistributorTransactor: CsfeedistributorTransactor{contract: contract}, CsfeedistributorFilterer: CsfeedistributorFilterer{contract: contract}}, nil
}

// NewCsfeedistributorCaller creates a new read-only instance of Csfeedistributor, bound to a specific deployed contract.
func NewCsfeedistributorCaller(address common.Address, caller bind.ContractCaller) (*CsfeedistributorCaller, error) {
	contract, err := bindCsfeedistributor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorCaller{contract: contract}, nil
}

// NewCsfeedistributorTransactor creates a new write-only instance of Csfeedistributor, bound to a specific deployed contract.
func NewCsfeedistributorTransactor(address common.Address, transactor bind.ContractTransactor) (*CsfeedistributorTransactor, error) {
	contract, err := bindCsfeedistributor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorTransactor{contract: contract}, nil
}

// NewCsfeedistributorFilterer creates a new log filterer instance of Csfeedistributor, bound to a specific deployed contract.
func NewCsfeedistributorFilterer(address common.Address, filterer bind.ContractFilterer) (*CsfeedistributorFilterer, error) {
	contract, err := bindCsfeedistributor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorFilterer{contract: contract}, nil
}

// bindCsfeedistributor binds a generic wrapper to an already deployed contract.
func bindCsfeedistributor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CsfeedistributorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csfeedistributor *CsfeedistributorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csfeedistributor.Contract.CsfeedistributorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csfeedistributor *CsfeedistributorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.CsfeedistributorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csfeedistributor *CsfeedistributorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.CsfeedistributorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csfeedistributor *CsfeedistributorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csfeedistributor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csfeedistributor *CsfeedistributorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csfeedistributor *CsfeedistributorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Csfeedistributor *CsfeedistributorCaller) ACCOUNTING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "ACCOUNTING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Csfeedistributor *CsfeedistributorSession) ACCOUNTING() (common.Address, error) {
	return _Csfeedistributor.Contract.ACCOUNTING(&_Csfeedistributor.CallOpts)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Csfeedistributor *CsfeedistributorCallerSession) ACCOUNTING() (common.Address, error) {
	return _Csfeedistributor.Contract.ACCOUNTING(&_Csfeedistributor.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Csfeedistributor.Contract.DEFAULTADMINROLE(&_Csfeedistributor.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Csfeedistributor.Contract.DEFAULTADMINROLE(&_Csfeedistributor.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_Csfeedistributor *CsfeedistributorCaller) ORACLE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "ORACLE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_Csfeedistributor *CsfeedistributorSession) ORACLE() (common.Address, error) {
	return _Csfeedistributor.Contract.ORACLE(&_Csfeedistributor.CallOpts)
}

// ORACLE is a free data retrieval call binding the contract method 0x38013f02.
//
// Solidity: function ORACLE() view returns(address)
func (_Csfeedistributor *CsfeedistributorCallerSession) ORACLE() (common.Address, error) {
	return _Csfeedistributor.Contract.ORACLE(&_Csfeedistributor.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCaller) RECOVERERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "RECOVERER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorSession) RECOVERERROLE() ([32]byte, error) {
	return _Csfeedistributor.Contract.RECOVERERROLE(&_Csfeedistributor.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCallerSession) RECOVERERROLE() ([32]byte, error) {
	return _Csfeedistributor.Contract.RECOVERERROLE(&_Csfeedistributor.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Csfeedistributor *CsfeedistributorCaller) STETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "STETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Csfeedistributor *CsfeedistributorSession) STETH() (common.Address, error) {
	return _Csfeedistributor.Contract.STETH(&_Csfeedistributor.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Csfeedistributor *CsfeedistributorCallerSession) STETH() (common.Address, error) {
	return _Csfeedistributor.Contract.STETH(&_Csfeedistributor.CallOpts)
}

// DistributedShares is a free data retrieval call binding the contract method 0xea6301ab.
//
// Solidity: function distributedShares(uint256 nodeOperatorId) view returns(uint256 distributed)
func (_Csfeedistributor *CsfeedistributorCaller) DistributedShares(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "distributedShares", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistributedShares is a free data retrieval call binding the contract method 0xea6301ab.
//
// Solidity: function distributedShares(uint256 nodeOperatorId) view returns(uint256 distributed)
func (_Csfeedistributor *CsfeedistributorSession) DistributedShares(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csfeedistributor.Contract.DistributedShares(&_Csfeedistributor.CallOpts, nodeOperatorId)
}

// DistributedShares is a free data retrieval call binding the contract method 0xea6301ab.
//
// Solidity: function distributedShares(uint256 nodeOperatorId) view returns(uint256 distributed)
func (_Csfeedistributor *CsfeedistributorCallerSession) DistributedShares(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csfeedistributor.Contract.DistributedShares(&_Csfeedistributor.CallOpts, nodeOperatorId)
}

// DistributionDataHistoryCount is a free data retrieval call binding the contract method 0x3d18b6f3.
//
// Solidity: function distributionDataHistoryCount() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCaller) DistributionDataHistoryCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "distributionDataHistoryCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DistributionDataHistoryCount is a free data retrieval call binding the contract method 0x3d18b6f3.
//
// Solidity: function distributionDataHistoryCount() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorSession) DistributionDataHistoryCount() (*big.Int, error) {
	return _Csfeedistributor.Contract.DistributionDataHistoryCount(&_Csfeedistributor.CallOpts)
}

// DistributionDataHistoryCount is a free data retrieval call binding the contract method 0x3d18b6f3.
//
// Solidity: function distributionDataHistoryCount() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCallerSession) DistributionDataHistoryCount() (*big.Int, error) {
	return _Csfeedistributor.Contract.DistributionDataHistoryCount(&_Csfeedistributor.CallOpts)
}

// GetFeesToDistribute is a free data retrieval call binding the contract method 0x5e8e8f6f.
//
// Solidity: function getFeesToDistribute(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] proof) view returns(uint256 sharesToDistribute)
func (_Csfeedistributor *CsfeedistributorCaller) GetFeesToDistribute(opts *bind.CallOpts, nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, proof [][32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "getFeesToDistribute", nodeOperatorId, cumulativeFeeShares, proof)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFeesToDistribute is a free data retrieval call binding the contract method 0x5e8e8f6f.
//
// Solidity: function getFeesToDistribute(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] proof) view returns(uint256 sharesToDistribute)
func (_Csfeedistributor *CsfeedistributorSession) GetFeesToDistribute(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, proof [][32]byte) (*big.Int, error) {
	return _Csfeedistributor.Contract.GetFeesToDistribute(&_Csfeedistributor.CallOpts, nodeOperatorId, cumulativeFeeShares, proof)
}

// GetFeesToDistribute is a free data retrieval call binding the contract method 0x5e8e8f6f.
//
// Solidity: function getFeesToDistribute(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] proof) view returns(uint256 sharesToDistribute)
func (_Csfeedistributor *CsfeedistributorCallerSession) GetFeesToDistribute(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, proof [][32]byte) (*big.Int, error) {
	return _Csfeedistributor.Contract.GetFeesToDistribute(&_Csfeedistributor.CallOpts, nodeOperatorId, cumulativeFeeShares, proof)
}

// GetHistoricalDistributionData is a free data retrieval call binding the contract method 0x4e5b3a62.
//
// Solidity: function getHistoricalDistributionData(uint256 index) view returns((uint256,bytes32,string,string,uint256,uint256))
func (_Csfeedistributor *CsfeedistributorCaller) GetHistoricalDistributionData(opts *bind.CallOpts, index *big.Int) (ICSFeeDistributorDistributionData, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "getHistoricalDistributionData", index)

	if err != nil {
		return *new(ICSFeeDistributorDistributionData), err
	}

	out0 := *abi.ConvertType(out[0], new(ICSFeeDistributorDistributionData)).(*ICSFeeDistributorDistributionData)

	return out0, err

}

// GetHistoricalDistributionData is a free data retrieval call binding the contract method 0x4e5b3a62.
//
// Solidity: function getHistoricalDistributionData(uint256 index) view returns((uint256,bytes32,string,string,uint256,uint256))
func (_Csfeedistributor *CsfeedistributorSession) GetHistoricalDistributionData(index *big.Int) (ICSFeeDistributorDistributionData, error) {
	return _Csfeedistributor.Contract.GetHistoricalDistributionData(&_Csfeedistributor.CallOpts, index)
}

// GetHistoricalDistributionData is a free data retrieval call binding the contract method 0x4e5b3a62.
//
// Solidity: function getHistoricalDistributionData(uint256 index) view returns((uint256,bytes32,string,string,uint256,uint256))
func (_Csfeedistributor *CsfeedistributorCallerSession) GetHistoricalDistributionData(index *big.Int) (ICSFeeDistributorDistributionData, error) {
	return _Csfeedistributor.Contract.GetHistoricalDistributionData(&_Csfeedistributor.CallOpts, index)
}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_Csfeedistributor *CsfeedistributorCaller) GetInitializedVersion(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "getInitializedVersion")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_Csfeedistributor *CsfeedistributorSession) GetInitializedVersion() (uint64, error) {
	return _Csfeedistributor.Contract.GetInitializedVersion(&_Csfeedistributor.CallOpts)
}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_Csfeedistributor *CsfeedistributorCallerSession) GetInitializedVersion() (uint64, error) {
	return _Csfeedistributor.Contract.GetInitializedVersion(&_Csfeedistributor.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Csfeedistributor.Contract.GetRoleAdmin(&_Csfeedistributor.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Csfeedistributor.Contract.GetRoleAdmin(&_Csfeedistributor.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csfeedistributor *CsfeedistributorCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csfeedistributor *CsfeedistributorSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Csfeedistributor.Contract.GetRoleMember(&_Csfeedistributor.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csfeedistributor *CsfeedistributorCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Csfeedistributor.Contract.GetRoleMember(&_Csfeedistributor.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csfeedistributor *CsfeedistributorSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Csfeedistributor.Contract.GetRoleMemberCount(&_Csfeedistributor.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Csfeedistributor.Contract.GetRoleMemberCount(&_Csfeedistributor.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csfeedistributor *CsfeedistributorCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csfeedistributor *CsfeedistributorSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Csfeedistributor.Contract.HasRole(&_Csfeedistributor.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csfeedistributor *CsfeedistributorCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Csfeedistributor.Contract.HasRole(&_Csfeedistributor.CallOpts, role, account)
}

// HashLeaf is a free data retrieval call binding the contract method 0x7e9f27ad.
//
// Solidity: function hashLeaf(uint256 nodeOperatorId, uint256 shares) pure returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCaller) HashLeaf(opts *bind.CallOpts, nodeOperatorId *big.Int, shares *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "hashLeaf", nodeOperatorId, shares)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashLeaf is a free data retrieval call binding the contract method 0x7e9f27ad.
//
// Solidity: function hashLeaf(uint256 nodeOperatorId, uint256 shares) pure returns(bytes32)
func (_Csfeedistributor *CsfeedistributorSession) HashLeaf(nodeOperatorId *big.Int, shares *big.Int) ([32]byte, error) {
	return _Csfeedistributor.Contract.HashLeaf(&_Csfeedistributor.CallOpts, nodeOperatorId, shares)
}

// HashLeaf is a free data retrieval call binding the contract method 0x7e9f27ad.
//
// Solidity: function hashLeaf(uint256 nodeOperatorId, uint256 shares) pure returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCallerSession) HashLeaf(nodeOperatorId *big.Int, shares *big.Int) ([32]byte, error) {
	return _Csfeedistributor.Contract.HashLeaf(&_Csfeedistributor.CallOpts, nodeOperatorId, shares)
}

// LogCid is a free data retrieval call binding the contract method 0x6f962e5c.
//
// Solidity: function logCid() view returns(string)
func (_Csfeedistributor *CsfeedistributorCaller) LogCid(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "logCid")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LogCid is a free data retrieval call binding the contract method 0x6f962e5c.
//
// Solidity: function logCid() view returns(string)
func (_Csfeedistributor *CsfeedistributorSession) LogCid() (string, error) {
	return _Csfeedistributor.Contract.LogCid(&_Csfeedistributor.CallOpts)
}

// LogCid is a free data retrieval call binding the contract method 0x6f962e5c.
//
// Solidity: function logCid() view returns(string)
func (_Csfeedistributor *CsfeedistributorCallerSession) LogCid() (string, error) {
	return _Csfeedistributor.Contract.LogCid(&_Csfeedistributor.CallOpts)
}

// PendingSharesToDistribute is a free data retrieval call binding the contract method 0xd257cf2a.
//
// Solidity: function pendingSharesToDistribute() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCaller) PendingSharesToDistribute(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "pendingSharesToDistribute")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingSharesToDistribute is a free data retrieval call binding the contract method 0xd257cf2a.
//
// Solidity: function pendingSharesToDistribute() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorSession) PendingSharesToDistribute() (*big.Int, error) {
	return _Csfeedistributor.Contract.PendingSharesToDistribute(&_Csfeedistributor.CallOpts)
}

// PendingSharesToDistribute is a free data retrieval call binding the contract method 0xd257cf2a.
//
// Solidity: function pendingSharesToDistribute() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCallerSession) PendingSharesToDistribute() (*big.Int, error) {
	return _Csfeedistributor.Contract.PendingSharesToDistribute(&_Csfeedistributor.CallOpts)
}

// RebateRecipient is a free data retrieval call binding the contract method 0x3333e109.
//
// Solidity: function rebateRecipient() view returns(address)
func (_Csfeedistributor *CsfeedistributorCaller) RebateRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "rebateRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RebateRecipient is a free data retrieval call binding the contract method 0x3333e109.
//
// Solidity: function rebateRecipient() view returns(address)
func (_Csfeedistributor *CsfeedistributorSession) RebateRecipient() (common.Address, error) {
	return _Csfeedistributor.Contract.RebateRecipient(&_Csfeedistributor.CallOpts)
}

// RebateRecipient is a free data retrieval call binding the contract method 0x3333e109.
//
// Solidity: function rebateRecipient() view returns(address)
func (_Csfeedistributor *CsfeedistributorCallerSession) RebateRecipient() (common.Address, error) {
	return _Csfeedistributor.Contract.RebateRecipient(&_Csfeedistributor.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csfeedistributor *CsfeedistributorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csfeedistributor *CsfeedistributorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Csfeedistributor.Contract.SupportsInterface(&_Csfeedistributor.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csfeedistributor *CsfeedistributorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Csfeedistributor.Contract.SupportsInterface(&_Csfeedistributor.CallOpts, interfaceId)
}

// TotalClaimableShares is a free data retrieval call binding the contract method 0x47d17d9d.
//
// Solidity: function totalClaimableShares() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCaller) TotalClaimableShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "totalClaimableShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalClaimableShares is a free data retrieval call binding the contract method 0x47d17d9d.
//
// Solidity: function totalClaimableShares() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorSession) TotalClaimableShares() (*big.Int, error) {
	return _Csfeedistributor.Contract.TotalClaimableShares(&_Csfeedistributor.CallOpts)
}

// TotalClaimableShares is a free data retrieval call binding the contract method 0x47d17d9d.
//
// Solidity: function totalClaimableShares() view returns(uint256)
func (_Csfeedistributor *CsfeedistributorCallerSession) TotalClaimableShares() (*big.Int, error) {
	return _Csfeedistributor.Contract.TotalClaimableShares(&_Csfeedistributor.CallOpts)
}

// TreeCid is a free data retrieval call binding the contract method 0xfe3c9b9b.
//
// Solidity: function treeCid() view returns(string)
func (_Csfeedistributor *CsfeedistributorCaller) TreeCid(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "treeCid")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TreeCid is a free data retrieval call binding the contract method 0xfe3c9b9b.
//
// Solidity: function treeCid() view returns(string)
func (_Csfeedistributor *CsfeedistributorSession) TreeCid() (string, error) {
	return _Csfeedistributor.Contract.TreeCid(&_Csfeedistributor.CallOpts)
}

// TreeCid is a free data retrieval call binding the contract method 0xfe3c9b9b.
//
// Solidity: function treeCid() view returns(string)
func (_Csfeedistributor *CsfeedistributorCallerSession) TreeCid() (string, error) {
	return _Csfeedistributor.Contract.TreeCid(&_Csfeedistributor.CallOpts)
}

// TreeRoot is a free data retrieval call binding the contract method 0x14dc6c14.
//
// Solidity: function treeRoot() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCaller) TreeRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csfeedistributor.contract.Call(opts, &out, "treeRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// TreeRoot is a free data retrieval call binding the contract method 0x14dc6c14.
//
// Solidity: function treeRoot() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorSession) TreeRoot() ([32]byte, error) {
	return _Csfeedistributor.Contract.TreeRoot(&_Csfeedistributor.CallOpts)
}

// TreeRoot is a free data retrieval call binding the contract method 0x14dc6c14.
//
// Solidity: function treeRoot() view returns(bytes32)
func (_Csfeedistributor *CsfeedistributorCallerSession) TreeRoot() ([32]byte, error) {
	return _Csfeedistributor.Contract.TreeRoot(&_Csfeedistributor.CallOpts)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x21893f7b.
//
// Solidity: function distributeFees(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] proof) returns(uint256 sharesToDistribute)
func (_Csfeedistributor *CsfeedistributorTransactor) DistributeFees(opts *bind.TransactOpts, nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "distributeFees", nodeOperatorId, cumulativeFeeShares, proof)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x21893f7b.
//
// Solidity: function distributeFees(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] proof) returns(uint256 sharesToDistribute)
func (_Csfeedistributor *CsfeedistributorSession) DistributeFees(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.DistributeFees(&_Csfeedistributor.TransactOpts, nodeOperatorId, cumulativeFeeShares, proof)
}

// DistributeFees is a paid mutator transaction binding the contract method 0x21893f7b.
//
// Solidity: function distributeFees(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] proof) returns(uint256 sharesToDistribute)
func (_Csfeedistributor *CsfeedistributorTransactorSession) DistributeFees(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, proof [][32]byte) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.DistributeFees(&_Csfeedistributor.TransactOpts, nodeOperatorId, cumulativeFeeShares, proof)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0xd5ba2dcf.
//
// Solidity: function finalizeUpgradeV2(address _rebateRecipient) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) FinalizeUpgradeV2(opts *bind.TransactOpts, _rebateRecipient common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "finalizeUpgradeV2", _rebateRecipient)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0xd5ba2dcf.
//
// Solidity: function finalizeUpgradeV2(address _rebateRecipient) returns()
func (_Csfeedistributor *CsfeedistributorSession) FinalizeUpgradeV2(_rebateRecipient common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.FinalizeUpgradeV2(&_Csfeedistributor.TransactOpts, _rebateRecipient)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0xd5ba2dcf.
//
// Solidity: function finalizeUpgradeV2(address _rebateRecipient) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) FinalizeUpgradeV2(_rebateRecipient common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.FinalizeUpgradeV2(&_Csfeedistributor.TransactOpts, _rebateRecipient)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csfeedistributor *CsfeedistributorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.GrantRole(&_Csfeedistributor.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.GrantRole(&_Csfeedistributor.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address _rebateRecipient) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, _rebateRecipient common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "initialize", admin, _rebateRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address _rebateRecipient) returns()
func (_Csfeedistributor *CsfeedistributorSession) Initialize(admin common.Address, _rebateRecipient common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.Initialize(&_Csfeedistributor.TransactOpts, admin, _rebateRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address admin, address _rebateRecipient) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) Initialize(admin common.Address, _rebateRecipient common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.Initialize(&_Csfeedistributor.TransactOpts, admin, _rebateRecipient)
}

// ProcessOracleReport is a paid mutator transaction binding the contract method 0x2ffa14e1.
//
// Solidity: function processOracleReport(bytes32 _treeRoot, string _treeCid, string _logCid, uint256 distributed, uint256 rebate, uint256 refSlot) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) ProcessOracleReport(opts *bind.TransactOpts, _treeRoot [32]byte, _treeCid string, _logCid string, distributed *big.Int, rebate *big.Int, refSlot *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "processOracleReport", _treeRoot, _treeCid, _logCid, distributed, rebate, refSlot)
}

// ProcessOracleReport is a paid mutator transaction binding the contract method 0x2ffa14e1.
//
// Solidity: function processOracleReport(bytes32 _treeRoot, string _treeCid, string _logCid, uint256 distributed, uint256 rebate, uint256 refSlot) returns()
func (_Csfeedistributor *CsfeedistributorSession) ProcessOracleReport(_treeRoot [32]byte, _treeCid string, _logCid string, distributed *big.Int, rebate *big.Int, refSlot *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.ProcessOracleReport(&_Csfeedistributor.TransactOpts, _treeRoot, _treeCid, _logCid, distributed, rebate, refSlot)
}

// ProcessOracleReport is a paid mutator transaction binding the contract method 0x2ffa14e1.
//
// Solidity: function processOracleReport(bytes32 _treeRoot, string _treeCid, string _logCid, uint256 distributed, uint256 rebate, uint256 refSlot) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) ProcessOracleReport(_treeRoot [32]byte, _treeCid string, _logCid string, distributed *big.Int, rebate *big.Int, refSlot *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.ProcessOracleReport(&_Csfeedistributor.TransactOpts, _treeRoot, _treeCid, _logCid, distributed, rebate, refSlot)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) RecoverERC1155(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "recoverERC1155", token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csfeedistributor *CsfeedistributorSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverERC1155(&_Csfeedistributor.TransactOpts, token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverERC1155(&_Csfeedistributor.TransactOpts, token, tokenId)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "recoverERC20", token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csfeedistributor *CsfeedistributorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverERC20(&_Csfeedistributor.TransactOpts, token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverERC20(&_Csfeedistributor.TransactOpts, token, amount)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) RecoverERC721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "recoverERC721", token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csfeedistributor *CsfeedistributorSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverERC721(&_Csfeedistributor.TransactOpts, token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverERC721(&_Csfeedistributor.TransactOpts, token, tokenId)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csfeedistributor *CsfeedistributorTransactor) RecoverEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "recoverEther")
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csfeedistributor *CsfeedistributorSession) RecoverEther() (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverEther(&_Csfeedistributor.TransactOpts)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) RecoverEther() (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RecoverEther(&_Csfeedistributor.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csfeedistributor *CsfeedistributorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RenounceRole(&_Csfeedistributor.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RenounceRole(&_Csfeedistributor.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csfeedistributor *CsfeedistributorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RevokeRole(&_Csfeedistributor.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.RevokeRole(&_Csfeedistributor.TransactOpts, role, account)
}

// SetRebateRecipient is a paid mutator transaction binding the contract method 0xe877f068.
//
// Solidity: function setRebateRecipient(address _rebateRecipient) returns()
func (_Csfeedistributor *CsfeedistributorTransactor) SetRebateRecipient(opts *bind.TransactOpts, _rebateRecipient common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.contract.Transact(opts, "setRebateRecipient", _rebateRecipient)
}

// SetRebateRecipient is a paid mutator transaction binding the contract method 0xe877f068.
//
// Solidity: function setRebateRecipient(address _rebateRecipient) returns()
func (_Csfeedistributor *CsfeedistributorSession) SetRebateRecipient(_rebateRecipient common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.SetRebateRecipient(&_Csfeedistributor.TransactOpts, _rebateRecipient)
}

// SetRebateRecipient is a paid mutator transaction binding the contract method 0xe877f068.
//
// Solidity: function setRebateRecipient(address _rebateRecipient) returns()
func (_Csfeedistributor *CsfeedistributorTransactorSession) SetRebateRecipient(_rebateRecipient common.Address) (*types.Transaction, error) {
	return _Csfeedistributor.Contract.SetRebateRecipient(&_Csfeedistributor.TransactOpts, _rebateRecipient)
}

// CsfeedistributorDistributionDataUpdatedIterator is returned from FilterDistributionDataUpdated and is used to iterate over the raw logs and unpacked data for DistributionDataUpdated events raised by the Csfeedistributor contract.
type CsfeedistributorDistributionDataUpdatedIterator struct {
	Event *CsfeedistributorDistributionDataUpdated // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorDistributionDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorDistributionDataUpdated)
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
		it.Event = new(CsfeedistributorDistributionDataUpdated)
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
func (it *CsfeedistributorDistributionDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorDistributionDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorDistributionDataUpdated represents a DistributionDataUpdated event raised by the Csfeedistributor contract.
type CsfeedistributorDistributionDataUpdated struct {
	TotalClaimableShares *big.Int
	TreeRoot             [32]byte
	TreeCid              string
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDistributionDataUpdated is a free log retrieval operation binding the contract event 0x26dec7cc117e9b3907dc1f90d2dc5f6e04dbb9f285f5898be2c82ec524dcd424.
//
// Solidity: event DistributionDataUpdated(uint256 totalClaimableShares, bytes32 treeRoot, string treeCid)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterDistributionDataUpdated(opts *bind.FilterOpts) (*CsfeedistributorDistributionDataUpdatedIterator, error) {

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "DistributionDataUpdated")
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorDistributionDataUpdatedIterator{contract: _Csfeedistributor.contract, event: "DistributionDataUpdated", logs: logs, sub: sub}, nil
}

// WatchDistributionDataUpdated is a free log subscription operation binding the contract event 0x26dec7cc117e9b3907dc1f90d2dc5f6e04dbb9f285f5898be2c82ec524dcd424.
//
// Solidity: event DistributionDataUpdated(uint256 totalClaimableShares, bytes32 treeRoot, string treeCid)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchDistributionDataUpdated(opts *bind.WatchOpts, sink chan<- *CsfeedistributorDistributionDataUpdated) (event.Subscription, error) {

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "DistributionDataUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorDistributionDataUpdated)
				if err := _Csfeedistributor.contract.UnpackLog(event, "DistributionDataUpdated", log); err != nil {
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

// ParseDistributionDataUpdated is a log parse operation binding the contract event 0x26dec7cc117e9b3907dc1f90d2dc5f6e04dbb9f285f5898be2c82ec524dcd424.
//
// Solidity: event DistributionDataUpdated(uint256 totalClaimableShares, bytes32 treeRoot, string treeCid)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseDistributionDataUpdated(log types.Log) (*CsfeedistributorDistributionDataUpdated, error) {
	event := new(CsfeedistributorDistributionDataUpdated)
	if err := _Csfeedistributor.contract.UnpackLog(event, "DistributionDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorDistributionLogUpdatedIterator is returned from FilterDistributionLogUpdated and is used to iterate over the raw logs and unpacked data for DistributionLogUpdated events raised by the Csfeedistributor contract.
type CsfeedistributorDistributionLogUpdatedIterator struct {
	Event *CsfeedistributorDistributionLogUpdated // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorDistributionLogUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorDistributionLogUpdated)
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
		it.Event = new(CsfeedistributorDistributionLogUpdated)
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
func (it *CsfeedistributorDistributionLogUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorDistributionLogUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorDistributionLogUpdated represents a DistributionLogUpdated event raised by the Csfeedistributor contract.
type CsfeedistributorDistributionLogUpdated struct {
	LogCid string
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDistributionLogUpdated is a free log retrieval operation binding the contract event 0x1f1a488b71a099a0d9cb71f60e14cf90bd1b5b188ca593111a40f533a3130b3b.
//
// Solidity: event DistributionLogUpdated(string logCid)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterDistributionLogUpdated(opts *bind.FilterOpts) (*CsfeedistributorDistributionLogUpdatedIterator, error) {

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "DistributionLogUpdated")
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorDistributionLogUpdatedIterator{contract: _Csfeedistributor.contract, event: "DistributionLogUpdated", logs: logs, sub: sub}, nil
}

// WatchDistributionLogUpdated is a free log subscription operation binding the contract event 0x1f1a488b71a099a0d9cb71f60e14cf90bd1b5b188ca593111a40f533a3130b3b.
//
// Solidity: event DistributionLogUpdated(string logCid)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchDistributionLogUpdated(opts *bind.WatchOpts, sink chan<- *CsfeedistributorDistributionLogUpdated) (event.Subscription, error) {

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "DistributionLogUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorDistributionLogUpdated)
				if err := _Csfeedistributor.contract.UnpackLog(event, "DistributionLogUpdated", log); err != nil {
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

// ParseDistributionLogUpdated is a log parse operation binding the contract event 0x1f1a488b71a099a0d9cb71f60e14cf90bd1b5b188ca593111a40f533a3130b3b.
//
// Solidity: event DistributionLogUpdated(string logCid)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseDistributionLogUpdated(log types.Log) (*CsfeedistributorDistributionLogUpdated, error) {
	event := new(CsfeedistributorDistributionLogUpdated)
	if err := _Csfeedistributor.contract.UnpackLog(event, "DistributionLogUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorERC1155RecoveredIterator is returned from FilterERC1155Recovered and is used to iterate over the raw logs and unpacked data for ERC1155Recovered events raised by the Csfeedistributor contract.
type CsfeedistributorERC1155RecoveredIterator struct {
	Event *CsfeedistributorERC1155Recovered // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorERC1155RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorERC1155Recovered)
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
		it.Event = new(CsfeedistributorERC1155Recovered)
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
func (it *CsfeedistributorERC1155RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorERC1155RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorERC1155Recovered represents a ERC1155Recovered event raised by the Csfeedistributor contract.
type CsfeedistributorERC1155Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC1155Recovered is a free log retrieval operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterERC1155Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsfeedistributorERC1155RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorERC1155RecoveredIterator{contract: _Csfeedistributor.contract, event: "ERC1155Recovered", logs: logs, sub: sub}, nil
}

// WatchERC1155Recovered is a free log subscription operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchERC1155Recovered(opts *bind.WatchOpts, sink chan<- *CsfeedistributorERC1155Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorERC1155Recovered)
				if err := _Csfeedistributor.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
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

// ParseERC1155Recovered is a log parse operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseERC1155Recovered(log types.Log) (*CsfeedistributorERC1155Recovered, error) {
	event := new(CsfeedistributorERC1155Recovered)
	if err := _Csfeedistributor.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorERC20RecoveredIterator is returned from FilterERC20Recovered and is used to iterate over the raw logs and unpacked data for ERC20Recovered events raised by the Csfeedistributor contract.
type CsfeedistributorERC20RecoveredIterator struct {
	Event *CsfeedistributorERC20Recovered // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorERC20RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorERC20Recovered)
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
		it.Event = new(CsfeedistributorERC20Recovered)
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
func (it *CsfeedistributorERC20RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorERC20RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorERC20Recovered represents a ERC20Recovered event raised by the Csfeedistributor contract.
type CsfeedistributorERC20Recovered struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Recovered is a free log retrieval operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterERC20Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsfeedistributorERC20RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorERC20RecoveredIterator{contract: _Csfeedistributor.contract, event: "ERC20Recovered", logs: logs, sub: sub}, nil
}

// WatchERC20Recovered is a free log subscription operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchERC20Recovered(opts *bind.WatchOpts, sink chan<- *CsfeedistributorERC20Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorERC20Recovered)
				if err := _Csfeedistributor.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
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

// ParseERC20Recovered is a log parse operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseERC20Recovered(log types.Log) (*CsfeedistributorERC20Recovered, error) {
	event := new(CsfeedistributorERC20Recovered)
	if err := _Csfeedistributor.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorERC721RecoveredIterator is returned from FilterERC721Recovered and is used to iterate over the raw logs and unpacked data for ERC721Recovered events raised by the Csfeedistributor contract.
type CsfeedistributorERC721RecoveredIterator struct {
	Event *CsfeedistributorERC721Recovered // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorERC721RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorERC721Recovered)
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
		it.Event = new(CsfeedistributorERC721Recovered)
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
func (it *CsfeedistributorERC721RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorERC721RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorERC721Recovered represents a ERC721Recovered event raised by the Csfeedistributor contract.
type CsfeedistributorERC721Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721Recovered is a free log retrieval operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterERC721Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsfeedistributorERC721RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorERC721RecoveredIterator{contract: _Csfeedistributor.contract, event: "ERC721Recovered", logs: logs, sub: sub}, nil
}

// WatchERC721Recovered is a free log subscription operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchERC721Recovered(opts *bind.WatchOpts, sink chan<- *CsfeedistributorERC721Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorERC721Recovered)
				if err := _Csfeedistributor.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
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

// ParseERC721Recovered is a log parse operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseERC721Recovered(log types.Log) (*CsfeedistributorERC721Recovered, error) {
	event := new(CsfeedistributorERC721Recovered)
	if err := _Csfeedistributor.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorEtherRecoveredIterator is returned from FilterEtherRecovered and is used to iterate over the raw logs and unpacked data for EtherRecovered events raised by the Csfeedistributor contract.
type CsfeedistributorEtherRecoveredIterator struct {
	Event *CsfeedistributorEtherRecovered // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorEtherRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorEtherRecovered)
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
		it.Event = new(CsfeedistributorEtherRecovered)
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
func (it *CsfeedistributorEtherRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorEtherRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorEtherRecovered represents a EtherRecovered event raised by the Csfeedistributor contract.
type CsfeedistributorEtherRecovered struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEtherRecovered is a free log retrieval operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterEtherRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CsfeedistributorEtherRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorEtherRecoveredIterator{contract: _Csfeedistributor.contract, event: "EtherRecovered", logs: logs, sub: sub}, nil
}

// WatchEtherRecovered is a free log subscription operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchEtherRecovered(opts *bind.WatchOpts, sink chan<- *CsfeedistributorEtherRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorEtherRecovered)
				if err := _Csfeedistributor.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
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

// ParseEtherRecovered is a log parse operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseEtherRecovered(log types.Log) (*CsfeedistributorEtherRecovered, error) {
	event := new(CsfeedistributorEtherRecovered)
	if err := _Csfeedistributor.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Csfeedistributor contract.
type CsfeedistributorInitializedIterator struct {
	Event *CsfeedistributorInitialized // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorInitialized)
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
		it.Event = new(CsfeedistributorInitialized)
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
func (it *CsfeedistributorInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorInitialized represents a Initialized event raised by the Csfeedistributor contract.
type CsfeedistributorInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterInitialized(opts *bind.FilterOpts) (*CsfeedistributorInitializedIterator, error) {

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorInitializedIterator{contract: _Csfeedistributor.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CsfeedistributorInitialized) (event.Subscription, error) {

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorInitialized)
				if err := _Csfeedistributor.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseInitialized(log types.Log) (*CsfeedistributorInitialized, error) {
	event := new(CsfeedistributorInitialized)
	if err := _Csfeedistributor.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorModuleFeeDistributedIterator is returned from FilterModuleFeeDistributed and is used to iterate over the raw logs and unpacked data for ModuleFeeDistributed events raised by the Csfeedistributor contract.
type CsfeedistributorModuleFeeDistributedIterator struct {
	Event *CsfeedistributorModuleFeeDistributed // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorModuleFeeDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorModuleFeeDistributed)
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
		it.Event = new(CsfeedistributorModuleFeeDistributed)
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
func (it *CsfeedistributorModuleFeeDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorModuleFeeDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorModuleFeeDistributed represents a ModuleFeeDistributed event raised by the Csfeedistributor contract.
type CsfeedistributorModuleFeeDistributed struct {
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterModuleFeeDistributed is a free log retrieval operation binding the contract event 0x010f65f5f56ba52d759f7b1dc49a3d277570cc2aa631e9c865b073a0ffc2af41.
//
// Solidity: event ModuleFeeDistributed(uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterModuleFeeDistributed(opts *bind.FilterOpts) (*CsfeedistributorModuleFeeDistributedIterator, error) {

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "ModuleFeeDistributed")
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorModuleFeeDistributedIterator{contract: _Csfeedistributor.contract, event: "ModuleFeeDistributed", logs: logs, sub: sub}, nil
}

// WatchModuleFeeDistributed is a free log subscription operation binding the contract event 0x010f65f5f56ba52d759f7b1dc49a3d277570cc2aa631e9c865b073a0ffc2af41.
//
// Solidity: event ModuleFeeDistributed(uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchModuleFeeDistributed(opts *bind.WatchOpts, sink chan<- *CsfeedistributorModuleFeeDistributed) (event.Subscription, error) {

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "ModuleFeeDistributed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorModuleFeeDistributed)
				if err := _Csfeedistributor.contract.UnpackLog(event, "ModuleFeeDistributed", log); err != nil {
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

// ParseModuleFeeDistributed is a log parse operation binding the contract event 0x010f65f5f56ba52d759f7b1dc49a3d277570cc2aa631e9c865b073a0ffc2af41.
//
// Solidity: event ModuleFeeDistributed(uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseModuleFeeDistributed(log types.Log) (*CsfeedistributorModuleFeeDistributed, error) {
	event := new(CsfeedistributorModuleFeeDistributed)
	if err := _Csfeedistributor.contract.UnpackLog(event, "ModuleFeeDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorOperatorFeeDistributedIterator is returned from FilterOperatorFeeDistributed and is used to iterate over the raw logs and unpacked data for OperatorFeeDistributed events raised by the Csfeedistributor contract.
type CsfeedistributorOperatorFeeDistributedIterator struct {
	Event *CsfeedistributorOperatorFeeDistributed // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorOperatorFeeDistributedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorOperatorFeeDistributed)
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
		it.Event = new(CsfeedistributorOperatorFeeDistributed)
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
func (it *CsfeedistributorOperatorFeeDistributedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorOperatorFeeDistributedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorOperatorFeeDistributed represents a OperatorFeeDistributed event raised by the Csfeedistributor contract.
type CsfeedistributorOperatorFeeDistributed struct {
	NodeOperatorId *big.Int
	Shares         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOperatorFeeDistributed is a free log retrieval operation binding the contract event 0x4b7ab1c192267e83350d06490a852b8dbbb25bfa00fd065b1862cf7accd2ab90.
//
// Solidity: event OperatorFeeDistributed(uint256 indexed nodeOperatorId, uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterOperatorFeeDistributed(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsfeedistributorOperatorFeeDistributedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "OperatorFeeDistributed", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorOperatorFeeDistributedIterator{contract: _Csfeedistributor.contract, event: "OperatorFeeDistributed", logs: logs, sub: sub}, nil
}

// WatchOperatorFeeDistributed is a free log subscription operation binding the contract event 0x4b7ab1c192267e83350d06490a852b8dbbb25bfa00fd065b1862cf7accd2ab90.
//
// Solidity: event OperatorFeeDistributed(uint256 indexed nodeOperatorId, uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchOperatorFeeDistributed(opts *bind.WatchOpts, sink chan<- *CsfeedistributorOperatorFeeDistributed, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "OperatorFeeDistributed", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorOperatorFeeDistributed)
				if err := _Csfeedistributor.contract.UnpackLog(event, "OperatorFeeDistributed", log); err != nil {
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

// ParseOperatorFeeDistributed is a log parse operation binding the contract event 0x4b7ab1c192267e83350d06490a852b8dbbb25bfa00fd065b1862cf7accd2ab90.
//
// Solidity: event OperatorFeeDistributed(uint256 indexed nodeOperatorId, uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseOperatorFeeDistributed(log types.Log) (*CsfeedistributorOperatorFeeDistributed, error) {
	event := new(CsfeedistributorOperatorFeeDistributed)
	if err := _Csfeedistributor.contract.UnpackLog(event, "OperatorFeeDistributed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorRebateRecipientSetIterator is returned from FilterRebateRecipientSet and is used to iterate over the raw logs and unpacked data for RebateRecipientSet events raised by the Csfeedistributor contract.
type CsfeedistributorRebateRecipientSetIterator struct {
	Event *CsfeedistributorRebateRecipientSet // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorRebateRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorRebateRecipientSet)
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
		it.Event = new(CsfeedistributorRebateRecipientSet)
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
func (it *CsfeedistributorRebateRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorRebateRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorRebateRecipientSet represents a RebateRecipientSet event raised by the Csfeedistributor contract.
type CsfeedistributorRebateRecipientSet struct {
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRebateRecipientSet is a free log retrieval operation binding the contract event 0x9f8636d85f90aba9c7d2c9e076c6102a5459d2e063afb71d81328bbb3608a234.
//
// Solidity: event RebateRecipientSet(address recipient)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterRebateRecipientSet(opts *bind.FilterOpts) (*CsfeedistributorRebateRecipientSetIterator, error) {

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "RebateRecipientSet")
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorRebateRecipientSetIterator{contract: _Csfeedistributor.contract, event: "RebateRecipientSet", logs: logs, sub: sub}, nil
}

// WatchRebateRecipientSet is a free log subscription operation binding the contract event 0x9f8636d85f90aba9c7d2c9e076c6102a5459d2e063afb71d81328bbb3608a234.
//
// Solidity: event RebateRecipientSet(address recipient)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchRebateRecipientSet(opts *bind.WatchOpts, sink chan<- *CsfeedistributorRebateRecipientSet) (event.Subscription, error) {

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "RebateRecipientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorRebateRecipientSet)
				if err := _Csfeedistributor.contract.UnpackLog(event, "RebateRecipientSet", log); err != nil {
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

// ParseRebateRecipientSet is a log parse operation binding the contract event 0x9f8636d85f90aba9c7d2c9e076c6102a5459d2e063afb71d81328bbb3608a234.
//
// Solidity: event RebateRecipientSet(address recipient)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseRebateRecipientSet(log types.Log) (*CsfeedistributorRebateRecipientSet, error) {
	event := new(CsfeedistributorRebateRecipientSet)
	if err := _Csfeedistributor.contract.UnpackLog(event, "RebateRecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorRebateTransferredIterator is returned from FilterRebateTransferred and is used to iterate over the raw logs and unpacked data for RebateTransferred events raised by the Csfeedistributor contract.
type CsfeedistributorRebateTransferredIterator struct {
	Event *CsfeedistributorRebateTransferred // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorRebateTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorRebateTransferred)
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
		it.Event = new(CsfeedistributorRebateTransferred)
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
func (it *CsfeedistributorRebateTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorRebateTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorRebateTransferred represents a RebateTransferred event raised by the Csfeedistributor contract.
type CsfeedistributorRebateTransferred struct {
	Shares *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRebateTransferred is a free log retrieval operation binding the contract event 0x7462935fb42d34d84233f737293310ca24e851021f9cb7f2549470cdf6de56bf.
//
// Solidity: event RebateTransferred(uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterRebateTransferred(opts *bind.FilterOpts) (*CsfeedistributorRebateTransferredIterator, error) {

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "RebateTransferred")
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorRebateTransferredIterator{contract: _Csfeedistributor.contract, event: "RebateTransferred", logs: logs, sub: sub}, nil
}

// WatchRebateTransferred is a free log subscription operation binding the contract event 0x7462935fb42d34d84233f737293310ca24e851021f9cb7f2549470cdf6de56bf.
//
// Solidity: event RebateTransferred(uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchRebateTransferred(opts *bind.WatchOpts, sink chan<- *CsfeedistributorRebateTransferred) (event.Subscription, error) {

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "RebateTransferred")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorRebateTransferred)
				if err := _Csfeedistributor.contract.UnpackLog(event, "RebateTransferred", log); err != nil {
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

// ParseRebateTransferred is a log parse operation binding the contract event 0x7462935fb42d34d84233f737293310ca24e851021f9cb7f2549470cdf6de56bf.
//
// Solidity: event RebateTransferred(uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseRebateTransferred(log types.Log) (*CsfeedistributorRebateTransferred, error) {
	event := new(CsfeedistributorRebateTransferred)
	if err := _Csfeedistributor.contract.UnpackLog(event, "RebateTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Csfeedistributor contract.
type CsfeedistributorRoleAdminChangedIterator struct {
	Event *CsfeedistributorRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorRoleAdminChanged)
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
		it.Event = new(CsfeedistributorRoleAdminChanged)
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
func (it *CsfeedistributorRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorRoleAdminChanged represents a RoleAdminChanged event raised by the Csfeedistributor contract.
type CsfeedistributorRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CsfeedistributorRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorRoleAdminChangedIterator{contract: _Csfeedistributor.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CsfeedistributorRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorRoleAdminChanged)
				if err := _Csfeedistributor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseRoleAdminChanged(log types.Log) (*CsfeedistributorRoleAdminChanged, error) {
	event := new(CsfeedistributorRoleAdminChanged)
	if err := _Csfeedistributor.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Csfeedistributor contract.
type CsfeedistributorRoleGrantedIterator struct {
	Event *CsfeedistributorRoleGranted // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorRoleGranted)
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
		it.Event = new(CsfeedistributorRoleGranted)
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
func (it *CsfeedistributorRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorRoleGranted represents a RoleGranted event raised by the Csfeedistributor contract.
type CsfeedistributorRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CsfeedistributorRoleGrantedIterator, error) {

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

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorRoleGrantedIterator{contract: _Csfeedistributor.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CsfeedistributorRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorRoleGranted)
				if err := _Csfeedistributor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseRoleGranted(log types.Log) (*CsfeedistributorRoleGranted, error) {
	event := new(CsfeedistributorRoleGranted)
	if err := _Csfeedistributor.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Csfeedistributor contract.
type CsfeedistributorRoleRevokedIterator struct {
	Event *CsfeedistributorRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorRoleRevoked)
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
		it.Event = new(CsfeedistributorRoleRevoked)
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
func (it *CsfeedistributorRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorRoleRevoked represents a RoleRevoked event raised by the Csfeedistributor contract.
type CsfeedistributorRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CsfeedistributorRoleRevokedIterator, error) {

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

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorRoleRevokedIterator{contract: _Csfeedistributor.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CsfeedistributorRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorRoleRevoked)
				if err := _Csfeedistributor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Csfeedistributor *CsfeedistributorFilterer) ParseRoleRevoked(log types.Log) (*CsfeedistributorRoleRevoked, error) {
	event := new(CsfeedistributorRoleRevoked)
	if err := _Csfeedistributor.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsfeedistributorStETHSharesRecoveredIterator is returned from FilterStETHSharesRecovered and is used to iterate over the raw logs and unpacked data for StETHSharesRecovered events raised by the Csfeedistributor contract.
type CsfeedistributorStETHSharesRecoveredIterator struct {
	Event *CsfeedistributorStETHSharesRecovered // Event containing the contract specifics and raw log

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
func (it *CsfeedistributorStETHSharesRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsfeedistributorStETHSharesRecovered)
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
		it.Event = new(CsfeedistributorStETHSharesRecovered)
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
func (it *CsfeedistributorStETHSharesRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsfeedistributorStETHSharesRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsfeedistributorStETHSharesRecovered represents a StETHSharesRecovered event raised by the Csfeedistributor contract.
type CsfeedistributorStETHSharesRecovered struct {
	Recipient common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStETHSharesRecovered is a free log retrieval operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) FilterStETHSharesRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CsfeedistributorStETHSharesRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.FilterLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsfeedistributorStETHSharesRecoveredIterator{contract: _Csfeedistributor.contract, event: "StETHSharesRecovered", logs: logs, sub: sub}, nil
}

// WatchStETHSharesRecovered is a free log subscription operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) WatchStETHSharesRecovered(opts *bind.WatchOpts, sink chan<- *CsfeedistributorStETHSharesRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csfeedistributor.contract.WatchLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsfeedistributorStETHSharesRecovered)
				if err := _Csfeedistributor.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
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

// ParseStETHSharesRecovered is a log parse operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csfeedistributor *CsfeedistributorFilterer) ParseStETHSharesRecovered(log types.Log) (*CsfeedistributorStETHSharesRecovered, error) {
	event := new(CsfeedistributorStETHSharesRecovered)
	if err := _Csfeedistributor.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
