// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lido

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

// StakingRouterNodeOperatorDigest is an auto generated low-level Go binding around an user-defined struct.
type StakingRouterNodeOperatorDigest struct {
	Id       *big.Int
	IsActive bool
	Summary  StakingRouterNodeOperatorSummary
}

// StakingRouterNodeOperatorSummary is an auto generated low-level Go binding around an user-defined struct.
type StakingRouterNodeOperatorSummary struct {
	IsTargetLimitActive        bool
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}

// StakingRouterStakingModule is an auto generated low-level Go binding around an user-defined struct.
type StakingRouterStakingModule struct {
	Id                    *big.Int
	StakingModuleAddress  common.Address
	StakingModuleFee      uint16
	TreasuryFee           uint16
	TargetShare           uint16
	Status                uint8
	Name                  string
	LastDepositAt         uint64
	LastDepositBlock      *big.Int
	ExitedValidatorsCount *big.Int
}

// StakingRouterStakingModuleDigest is an auto generated low-level Go binding around an user-defined struct.
type StakingRouterStakingModuleDigest struct {
	NodeOperatorsCount       *big.Int
	ActiveNodeOperatorsCount *big.Int
	State                    StakingRouterStakingModule
	Summary                  StakingRouterStakingModuleSummary
}

// StakingRouterStakingModuleSummary is an auto generated low-level Go binding around an user-defined struct.
type StakingRouterStakingModuleSummary struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}

// StakingRouterValidatorsCountsCorrection is an auto generated low-level Go binding around an user-defined struct.
type StakingRouterValidatorsCountsCorrection struct {
	CurrentModuleExitedValidatorsCount       *big.Int
	CurrentNodeOperatorExitedValidatorsCount *big.Int
	CurrentNodeOperatorStuckValidatorsCount  *big.Int
	NewModuleExitedValidatorsCount           *big.Int
	NewNodeOperatorExitedValidatorsCount     *big.Int
	NewNodeOperatorStuckValidatorsCount      *big.Int
}

// ApiMetaData contains all meta data concerning the Api contract.
var ApiMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AppAuthLidoFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"firstArrayLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondArrayLength\",\"type\":\"uint256\"}],\"name\":\"ArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositContractZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DirectETHTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyWithdrawalsCredentials\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitedValidatorsCountCannotDecrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractVersionIncrement\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"etherValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositsCount\",\"type\":\"uint256\"}],\"name\":\"InvalidDepositsValue\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"InvalidPublicKeysBatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"}],\"name\":\"InvalidReportData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"InvalidSignaturesBatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroContractVersionOnInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reportedExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositedValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"ReportedExitedValidatorsExceedDeposited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModuleAddressExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModuleNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModuleNotPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModuleStatusTheSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModuleUnregistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModuleWrongName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModulesLimitExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"UnexpectedContractVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentModuleExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentNodeOpExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentNodeOpStuckValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"UnexpectedCurrentValidatorsCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnrecoverableModuleError\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"}],\"name\":\"ValueOver100Percent\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"field\",\"type\":\"string\"}],\"name\":\"ZeroAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelRevertData\",\"type\":\"bytes\"}],\"name\":\"ExitedAndStuckValidatorsCountsUpdateFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelRevertData\",\"type\":\"bytes\"}],\"name\":\"RewardsMintedReportFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"name\":\"StakingModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unreportedExitedValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"StakingModuleExitedValidatorsIncompleteReporting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakingModuleFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"treasuryFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"StakingModuleFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumStakingRouter.StakingModuleStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"StakingModuleStatusSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"StakingModuleTargetShareSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingRouterETHDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalCredentials\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"WithdrawalCredentialsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelRevertData\",\"type\":\"bytes\"}],\"name\":\"WithdrawalsCredentialsChangeFailed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_PRECISION_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_WITHDRAWAL_CREDENTIALS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STAKING_MODULES_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STAKING_MODULE_NAME_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORT_EXITED_VALIDATORS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORT_REWARDS_MINTED_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MODULE_MANAGE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MODULE_PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MODULE_RESUME_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNSAFE_SET_EXITED_VALIDATORS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_stakingModuleAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_targetShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingModuleFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_treasuryFee\",\"type\":\"uint256\"}],\"name\":\"addStakingModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_depositsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_depositCalldata\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getAllNodeOperatorDigests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isTargetLimitActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.NodeOperatorSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"internalType\":\"structStakingRouter.NodeOperatorDigest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllStakingModuleDigests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeNodeOperatorsCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"stakingModuleAddress\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"stakingModuleFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"treasuryFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"targetShare\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"lastDepositAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"lastDepositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.StakingModule\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.StakingModuleSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"internalType\":\"structStakingRouter.StakingModuleDigest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_depositsCount\",\"type\":\"uint256\"}],\"name\":\"getDepositsAllocation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allocated\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"allocations\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLido\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_nodeOperatorIds\",\"type\":\"uint256[]\"}],\"name\":\"getNodeOperatorDigests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isTargetLimitActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.NodeOperatorSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"internalType\":\"structStakingRouter.NodeOperatorDigest[]\",\"name\":\"digests\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorDigests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"isTargetLimitActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.NodeOperatorSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"internalType\":\"structStakingRouter.NodeOperatorDigest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"isTargetLimitActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.NodeOperatorSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingFeeAggregateDistribution\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"modulesFee\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"treasuryFee\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"basePrecision\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingFeeAggregateDistributionE4Precision\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"modulesFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"treasuryFee\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModule\",\"outputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"stakingModuleAddress\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"stakingModuleFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"treasuryFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"targetShare\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"lastDepositAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"lastDepositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.StakingModule\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleActiveValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"activeValidatorsCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_stakingModuleIds\",\"type\":\"uint256[]\"}],\"name\":\"getStakingModuleDigests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeNodeOperatorsCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"stakingModuleAddress\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"stakingModuleFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"treasuryFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"targetShare\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"lastDepositAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"lastDepositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.StakingModule\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.StakingModuleSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"internalType\":\"structStakingRouter.StakingModuleDigest[]\",\"name\":\"digests\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingModuleIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakingModuleIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleIsActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleIsDepositsPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleIsStopped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleLastDepositBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDepositsValue\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleMaxDepositsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleStatus\",\"outputs\":[{\"internalType\":\"enumStakingRouter.StakingModuleStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.StakingModuleSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingModules\",\"outputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"stakingModuleAddress\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"stakingModuleFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"treasuryFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"targetShare\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"lastDepositAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"lastDepositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.StakingModule[]\",\"name\":\"res\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingModulesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingRewardsDistribution\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakingModuleIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint96[]\",\"name\":\"stakingModuleFees\",\"type\":\"uint96[]\"},{\"internalType\":\"uint96\",\"name\":\"totalFee\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"precisionPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalFeeE4Precision\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"totalFee\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawalCredentials\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"hasStakingModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lido\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_withdrawalCredentials\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onValidatorsCountsByNodeOperatorReportingFinished\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"pauseStakingModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_stakingModuleIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_totalShares\",\"type\":\"uint256[]\"}],\"name\":\"reportRewardsMinted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_nodeOperatorIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_exitedValidatorsCounts\",\"type\":\"bytes\"}],\"name\":\"reportStakingModuleExitedValidatorsCountByNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_nodeOperatorIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_stuckValidatorsCounts\",\"type\":\"bytes\"}],\"name\":\"reportStakingModuleStuckValidatorsCountByNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"resumeStakingModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"enumStakingRouter.StakingModuleStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStakingModuleStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_withdrawalCredentials\",\"type\":\"bytes32\"}],\"name\":\"setWithdrawalCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_triggerUpdateFinish\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentModuleExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentNodeOperatorExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentNodeOperatorStuckValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newModuleExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newNodeOperatorExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newNodeOperatorStuckValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.ValidatorsCountsCorrection\",\"name\":\"_correction\",\"type\":\"tuple\"}],\"name\":\"unsafeSetExitedValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_stakingModuleIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_exitedValidatorsCounts\",\"type\":\"uint256[]\"}],\"name\":\"updateExitedValidatorsCountByStakingModule\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_refundedValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"updateRefundedValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingModuleFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_treasuryFee\",\"type\":\"uint256\"}],\"name\":\"updateStakingModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isTargetLimitActive\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"_targetLimit\",\"type\":\"uint256\"}],\"name\":\"updateTargetValidatorsLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x60a06040523480156200001157600080fd5b5060405162006098380380620060988339810160408190526200003491620000ae565b806001600160a01b0381166200005d57604051637c5f8bcf60e11b815260040160405180910390fd5b6001600160a01b0316608052620000a37f4dd0f6662ba1d6b081f08b350f5e9a6a7b15cf586926ba66f753594928fa64a6600019620000aa602090811b6200374a17901c565b50620000e0565b9055565b600060208284031215620000c157600080fd5b81516001600160a01b0381168114620000d957600080fd5b9392505050565b608051615f9562000103600039600081816107540152613d220152615f956000f3fe6080604052600436106103d25760003560e01c80639010d07c116101fd578063c82b1bb111610118578063e016e6f7116100ab578063efcdcc0e1161007a578063efcdcc0e14610c4a578063f07ff28a14610c7a578063f2aebb6514610c9a578063f8bb6d4214610cbc578063fa5093eb14610cdc57600080fd5b8063e016e6f714610bc8578063e1b92a5c14610bea578063e24ce9f114610c0a578063e97ee8cc14610c2a57600080fd5b8063d0a2b1b8116100e7578063d0a2b1b814610b53578063d547741f14610b73578063d861c58414610b93578063db3c7ba714610bb357600080fd5b8063c82b1bb114610ac5578063c8ac498014610af3578063ca15c87314610b13578063cb589b9a14610b3357600080fd5b8063a7357c8c11610190578063af1240971161015f578063af12409714610a30578063ba21ccae14610a50578063bc1bb19014610a76578063c445ea7514610aa357600080fd5b8063a7357c8c1461099d578063aa0b7db7146109d0578063aa5a1b9d146109e3578063abd44a2414610a1057600080fd5b80639fbb7bae116101cc5780639fbb7bae146108e55780639fc5a6ed1461090d578063a217fddf1461093a578063a734329c1461094f57600080fd5b80639010d07c1461087057806391d148541461089057806396b5d81c146108b05780639b75b4ef146108d057600080fd5b806356396715116102ed5780636b96736b116102805780638525e3a11161024f5780638525e3a1146107e75780638801da79146108075780638aa104351461083b5780638dc70c571461085057600080fd5b80636b96736b146107425780637443f523146107765780637a74884d146107965780637c8da51c146107ca57600080fd5b80636183214d116102bc5780636183214d146106b35780636608b11b146106d55780636a516b47146106f55780636ada55b91461072257600080fd5b8063563967151461063c57806357993b85146106515780635bf55e40146106735780636133f9851461069357600080fd5b8063271662ec116103655780633e54ee5b116103345780633e54ee5b146105d2578063473e0433146105f25780634a7583b6146106125780634b3a1cb71461062757600080fd5b8063271662ec1461054f5780632f2ff15d146105655780633240a3221461058557806336568abe146105b257600080fd5b80631565d2f2116103a15780631565d2f2146104a757806319c64b79146104db5780631d1b9d3c146104fb578063248a9ca31461052f57600080fd5b806301ffc9a7146103f55780630519fbbf1461042a578063072859c71461045857806307e203ac1461047a57600080fd5b366103f0576040516309fb455960e41b815260040160405180910390fd5b600080fd5b34801561040157600080fd5b50610415610410366004614e4f565b610d17565b60405190151581526020015b60405180910390f35b34801561043657600080fd5b5061044a610445366004614e79565b610d42565b604051908152602001610421565b34801561046457600080fd5b50610478610473366004614f0e565b610dbd565b005b34801561048657600080fd5b5061049a610495366004614e79565b610fcc565b6040516104219190614fa3565b3480156104b357600080fd5b5061044a7f55180e25fcacf9af017d35d497765476319b23896daa1f9bc2b38fa80b36a16381565b3480156104e757600080fd5b5061044a6104f6366004614fc4565b61108b565b34801561050757600080fd5b5061044a7f779e5c23cb7a5bcb9bfe1e9a5165a00057f12bcdfd13e374540fdf1a1cd9113781565b34801561053b57600080fd5b5061044a61054a366004614e79565b611108565b34801561055b57600080fd5b5061044a61271081565b34801561057157600080fd5b50610478610580366004615002565b61112a565b34801561059157600080fd5b506105a56105a0366004614e79565b61114c565b604051610421919061507e565b3480156105be57600080fd5b506104786105cd366004615002565b6111e4565b3480156105de57600080fd5b506104786105ed36600461512c565b611262565b3480156105fe57600080fd5b5061044a61060d366004614e79565b611685565b34801561061e57600080fd5b5061044a61169c565b34801561063357600080fd5b5061044a602081565b34801561064857600080fd5b5061044a6116cb565b34801561065d57600080fd5b506106666116f5565b60405161042191906152c6565b34801561067f57600080fd5b5061047861068e366004614e79565b611702565b34801561069f57600080fd5b506104786106ae36600461536b565b611792565b3480156106bf57600080fd5b506106c86118b6565b60405161042191906153a7565b3480156106e157600080fd5b506104156106f0366004614e79565b611a85565b34801561070157600080fd5b5061070a611aaa565b6040516001600160a01b039091168152602001610421565b34801561072e57600080fd5b5061041561073d366004614e79565b611ad4565b34801561074e57600080fd5b5061070a7f000000000000000000000000000000000000000000000000000000000000000081565b34801561078257600080fd5b50610478610791366004615409565b611add565b3480156107a257600080fd5b5061044a7fe7c742a54cd11fc9749a47ab34bdcd7327820908e8d0d48b4a5c7f17b029409881565b3480156107d657600080fd5b5061044a68056bc75e2d6310000081565b3480156107f357600080fd5b506106666108023660046154cb565b611b7d565b34801561081357600080fd5b5061044a7f9a2f67efb89489040f2c48c3b2c38f719fba1276678d2ced3bd9049fb5edc6b281565b34801561084757600080fd5b5061044a611d67565b34801561085c57600080fd5b5061047861086b3660046154ff565b611d91565b34801561087c57600080fd5b5061070a61088b366004614fc4565b611f22565b34801561089c57600080fd5b506104156108ab366004615002565b611f4e565b3480156108bc57600080fd5b5061044a6108cb366004614e79565b611f86565b3480156108dc57600080fd5b5061044a601f81565b3480156108f157600080fd5b506108fa612045565b60405161ffff9091168152602001610421565b34801561091957600080fd5b5061092d610928366004614e79565b612073565b6040516104219190615569565b34801561094657600080fd5b5061044a600081565b34801561095b57600080fd5b5061041561096a366004614e79565b60009081527f9b48f5b32acb95b982effe269feac267eead113c4b5af14ffeb9aadac18a6e9c6020526040902054151590565b3480156109a957600080fd5b5061044a7eb1e70095ba5bacc3202c3db9faf1f7873186f0ed7b6c84e80c0018dcc6e38e81565b6104786109de366004615577565b61209a565b3480156109ef57600080fd5b50610a036109fe366004614fc4565b612300565b60405161042191906155c9565b348015610a1c57600080fd5b5061044a610a2b36600461561c565b61241f565b348015610a3c57600080fd5b50610478610a4b36600461561c565b612679565b348015610a5c57600080fd5b50610a6561284f565b6040516104219594939291906156b6565b348015610a8257600080fd5b50610a96610a91366004614e79565b612bca565b6040516104219190615775565b348015610aaf57600080fd5b5061044a600080516020615f2083398151915281565b348015610ad157600080fd5b50610ae5610ae0366004614e79565b612d0e565b604051610421929190615788565b348015610aff57600080fd5b50610478610b0e3660046157a1565b612d26565b348015610b1f57600080fd5b5061044a610b2e366004614e79565b612dc9565b348015610b3f57600080fd5b50610478610b4e3660046157a1565b612ded565b348015610b5f57600080fd5b50610478610b6e36600461581a565b612e61565b348015610b7f57600080fd5b50610478610b8e366004615002565b612ef6565b348015610b9f57600080fd5b50610478610bae366004614e79565b612f13565b348015610bbf57600080fd5b50610478612fa4565b348015610bd457600080fd5b5061044a600080516020615f4083398151915281565b348015610bf657600080fd5b50610478610c0536600461584e565b613171565b348015610c1657600080fd5b50610415610c25366004614e79565b6131fd565b348015610c3657600080fd5b50610478610c45366004614e79565b613206565b348015610c5657600080fd5b50610c5f6133c5565b6040805161ffff938416815292909116602083015201610421565b348015610c8657600080fd5b506105a5610c9536600461587a565b61340c565b348015610ca657600080fd5b50610caf6135a0565b60405161042191906158c0565b348015610cc857600080fd5b506105a5610cd736600461584e565b613636565b348015610ce857600080fd5b50610cf16136d8565b604080516001600160601b03948516815293909216602084015290820152606001610421565b60006001600160e01b03198216635a05180f60e01b1480610d3c5750610d3c8261374e565b92915050565b6000610d4d82613783565b6001600160a01b031663d087d2886040518163ffffffff1660e01b815260040160206040518083038186803b158015610d8557600080fd5b505afa158015610d99573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d3c91906158d3565b7f55180e25fcacf9af017d35d497765476319b23896daa1f9bc2b38fa80b36a163610de881336137a5565b6000610df386613809565b8054604051632cc1db0f60e21b815260048101889052919250630100000090046001600160a01b0316906000908190839063b3076c3c906024016101006040518083038186803b158015610e4657600080fd5b505afa158015610e5a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e7e91906158ec565b5050955050509350505083600401548660000151141580610ea3575080866020015114155b80610eb2575081866040015114155b15610ee95760048481015460405163e882688560e01b81529182015260248101829052604481018390526064015b60405180910390fd5b6060860151600480860191909155608087015160a088015160405163f2e2ca6360e01b81529283018b9052602483019190915260448201526001600160a01b0384169063f2e2ca6390606401600060405180830381600087803b158015610f4f57600080fd5b505af1158015610f63573d6000803e3d6000fd5b505050508615610fc157826001600160a01b031663e864299e6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610fa857600080fd5b505af1158015610fbc573d6000803e3d6000fd5b505050505b505050505050505050565b610ff060405180606001604052806000815260200160008152602001600081525090565b6000610ffb83612bca565b9050600081602001519050806001600160a01b0316639abddf096040518163ffffffff1660e01b815260040160606040518083038186803b15801561103f57600080fd5b505afa158015611053573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906110779190615956565b604086015260208501528352509092915050565b600080806110aa6110a56801bc16d674ec800000866159b0565b61381c565b925092505060006110ba866139e1565b90508181815181106110ce576110ce6159c4565b602002602001015160c001518382815181106110ec576110ec6159c4565b60200260200101516110fe91906159da565b9695505050505050565b6000908152600080516020615f00833981519152602052604090206001015490565b61113382611108565b61113d81336137a5565b6111478383613a3a565b505050565b6060600061115983613783565b90506000816001600160a01b031663a70c70e46040518163ffffffff1660e01b815260040160206040518083038186803b15801561119657600080fd5b505afa1580156111aa573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111ce91906158d3565b90506111dc84600083613636565b949350505050565b6001600160a01b03811633146112545760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608401610ee0565b61125e8282613a69565b5050565b600080516020615f4083398151915261127b81336137a5565b6127108411156112bd57604051630cb4392560e31b815260206004820152600c60248201526b5f746172676574536861726560a01b6044820152606401610ee0565b6127106112ca83856159f1565b111561131957604051630cb4392560e31b815260206004820181905260248201527f5f7374616b696e674d6f64756c65466565202b205f74726561737572794665656044820152606401610ee0565b6001600160a01b0385166113685760405163eac0d38960e01b81526020600482015260156024820152745f7374616b696e674d6f64756c654164647265737360581b6044820152606401610ee0565b8515806113755750601f86115b156113935760405163ac18716960e01b815260040160405180910390fd5b600061139d61169c565b9050602081106113c05760405163309eed9960e01b815260040160405180910390fd5b60005b81811015611410576113d481613a98565b546001600160a01b0388811663010000009092041614156114085760405163050f969d60e41b815260040160405180910390fd5b6001016113c3565b50600061141c82613a98565b905060006114487ff9a85ae945d8134f58bd2ee028636634dcb9e812798acb5c806bf1951232a2255490565b611453906001615a09565b825462ffffff191662ffffff82161783559050611474600183018b8b614c69565b508154630100000065ffff0000000160b81b03191663010000006001600160a01b038a160261ffff60d81b191617600160d81b61ffff898116919091029190911763ffffffff60b81b1916600160b81b8883160261ffff60c81b191617600160c81b918716919091021760ff60e81b1916825560028201805467ffffffffffffffff1916426001600160401b03161790554360038301556040805160008152905162ffffff8316917f9151b7f88aca05d432bb395647ef52b2ffc454e3c6afb69c95345af6b5a778c0919081900360200190a26115568162ffffff1684613ac8565b62ffffff81167ff9a85ae945d8134f58bd2ee028636634dcb9e812798acb5c806bf1951232a225556115b061158c8460016159f1565b7f1b3ef9db2d6f0727a31622833b45264c21051726d23ddb6f73b3b65628cafcc355565b8062ffffff167f43b5213f0e1666cd0b8692a73686164c94deb955a59c65e10dee8bb958e7ce3e898c8c336040516115eb9493929190615a59565b60405180910390a26040805188815233602082015262ffffff8316917f065e5bd8e4145dd99cf69bad5871ad52d094aee07a67fcf2f418c89e49d5f20c910160405180910390a260408051878152602081018790523381830152905162ffffff8316917f303c8ac43d1b1f9b898ddd2915a294efa01e9b07c322d7deeb7db332b66f0410919081900360600190a250505050505050505050565b60008061169183613809565b600301549392505050565b60006116c67f1b3ef9db2d6f0727a31622833b45264c21051726d23ddb6f73b3b65628cafcc35490565b905090565b60006116c67fabeb05279af36da5d476d7f950157cd2ea98a4166fa68a6bc97ce3a22fbb93c05490565b60606116c66108026135a0565b7eb1e70095ba5bacc3202c3db9faf1f7873186f0ed7b6c84e80c0018dcc6e38e61172c81336137a5565b600061173783613809565b905060008154600160e81b900460ff16600281111561175857611758615531565b600281111561176957611769615531565b146117875760405163322e64fb60e11b815260040160405180910390fd5b611147816001613b0b565b6001600160a01b0383166117d25760405163eac0d38960e01b81526020600482015260066024820152652fb0b236b4b760d11b6044820152606401610ee0565b6001600160a01b0382166118115760405163eac0d38960e01b81526020600482015260056024820152645f6c69646f60d81b6044820152606401610ee0565b61181b6001613bce565b611826600084613c00565b61184f7f706b9ed9846c161ad535be9b6345c3a7b2cb929e8d4a7254dee9ba6e6f8e5531839055565b6118787fabeb05279af36da5d476d7f950157cd2ea98a4166fa68a6bc97ce3a22fbb93c0829055565b604080518281523360208201527f82e72df77173eab89b00556d791a407a78f4605c5c2f0694967c8c429dd43c7c91015b60405180910390a1505050565b606060006118c261169c565b9050806001600160401b038111156118dc576118dc614ea0565b60405190808252806020026020018201604052801561191557816020015b611902614ced565b8152602001906001900390816118fa5790505b50915060005b81811015611a805761192c81613a98565b6040805161014081018252825462ffffff81168252630100000081046001600160a01b03166020830152600160b81b810461ffff90811693830193909352600160c81b810483166060830152600160d81b81049092166080820152600160e81b90910460ff1660a082015260018201805491929160c0840191906119af90615a90565b80601f01602080910402602001604051908101604052809291908181526020018280546119db90615a90565b8015611a285780601f106119fd57610100808354040283529160200191611a28565b820191906000526020600020905b815481529060010190602001808311611a0b57829003601f168201915b505050918352505060028201546001600160401b03166020820152600382015460408201526004909101546060909101528351849083908110611a6d57611a6d6159c4565b602090810291909101015260010161191b565b505090565b6000805b611a9283612073565b6002811115611aa357611aa3615531565b1492915050565b60006116c67f706b9ed9846c161ad535be9b6345c3a7b2cb929e8d4a7254dee9ba6e6f8e55315490565b60006002611a89565b600080516020615f40833981519152611af681336137a5565b6000611b0186613809565b546040516354f3d42360e11b81526004810187905285151560248201526044810185905263010000009091046001600160a01b03169150819063a9e7a84690606401600060405180830381600087803b158015611b5d57600080fd5b505af1158015611b71573d6000803e3d6000fd5b50505050505050505050565b606081516001600160401b03811115611b9857611b98614ea0565b604051908082528060200260200182016040528015611bd157816020015b611bbe614d40565b815260200190600190039081611bb65790505b50905060005b8251811015611d61576000611c04848381518110611bf757611bf76159c4565b6020026020010151612bca565b90506000816020015190506040518060800160405280826001600160a01b031663a70c70e46040518163ffffffff1660e01b815260040160206040518083038186803b158015611c5357600080fd5b505afa158015611c67573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611c8b91906158d3565b8152602001826001600160a01b0316638469cbd36040518163ffffffff1660e01b815260040160206040518083038186803b158015611cc957600080fd5b505afa158015611cdd573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d0191906158d3565b8152602001838152602001611d2e878681518110611d2157611d216159c4565b6020026020010151610fcc565b815250848481518110611d4357611d436159c4565b6020026020010181905250505080611d5a90615ac5565b9050611bd7565b50919050565b60006116c67f4dd0f6662ba1d6b081f08b350f5e9a6a7b15cf586926ba66f753594928fa64a65490565b600080516020615f40833981519152611daa81336137a5565b612710841115611dec57604051630cb4392560e31b815260206004820152600c60248201526b5f746172676574536861726560a01b6044820152606401610ee0565b612710611df983856159f1565b1115611e4857604051630cb4392560e31b815260206004820181905260248201527f5f7374616b696e674d6f64756c65466565202b205f74726561737572794665656044820152606401610ee0565b6000611e5386613809565b805463ffffffff60c81b1916600160d81b61ffff8881169190910261ffff60c81b191691909117600160c81b868316021761ffff60b81b1916600160b81b918716919091021781556040805187815233602082015291925087917f065e5bd8e4145dd99cf69bad5871ad52d094aee07a67fcf2f418c89e49d5f20c910160405180910390a260408051858152602081018590523381830152905187917f303c8ac43d1b1f9b898ddd2915a294efa01e9b07c322d7deeb7db332b66f0410919081900360600190a2505050505050565b6000828152600080516020615ee083398151915260205260408120611f479083613c0a565b9392505050565b6000918252600080516020615f00833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600080611f9283613809565b90506000808260000160039054906101000a90046001600160a01b03166001600160a01b0316639abddf096040518163ffffffff1660e01b815260040160606040518083038186803b158015611fe757600080fd5b505afa158015611ffb573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061201f9190615956565b5091509150612032836004015483613c16565b61203c90826159da565b95945050505050565b600080600061205261284f565b9450945050505061206c826001600160601b031682613c2c565b9250505090565b600061207e82613809565b54600160e81b900460ff166002811115610d3c57610d3c615531565b7f706b9ed9846c161ad535be9b6345c3a7b2cb929e8d4a7254dee9ba6e6f8e5531546001600160a01b0316336001600160a01b0316146120ed57604051637e71782360e01b815260040160405180910390fd5b60006120f76116cb565b9050806121175760405163180a97cd60e21b815260040160405180910390fd5b600061212285613809565b905060008154600160e81b900460ff16600281111561214357612143615531565b600281111561215457612154615531565b146121725760405163322e64fb60e11b815260040160405180910390fd5b60028101805467ffffffffffffffff1916426001600160401b0316179055436003820155604051348082529086907f9151b7f88aca05d432bb395647ef52b2ffc454e3c6afb69c95345af6b5a778c09060200160405180910390a26121e06801bc16d674ec80000088615ae0565b81146122095760405163023db95b60e21b81526004810182905260248101889052604401610ee0565b86156122f75781546040516317dc836b60e31b8152600091829163010000009091046001600160a01b03169063bee41b589061224d908c908b908b90600401615aff565b600060405180830381600087803b15801561226757600080fd5b505af115801561227b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526122a39190810190615b7c565b9150915060004790506122d98a876040516020016122c391815260200190565b6040516020818303038152906040528585613c45565b47846122e582846159da565b146122f2576122f2615bd5565b505050505b50505050505050565b612308614d8f565b600061231384612bca565b9050600081602001519050600080600080600080600080886001600160a01b031663b3076c3c8d6040518263ffffffff1660e01b815260040161235891815260200190565b6101006040518083038186803b15801561237157600080fd5b505afa158015612385573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906123a991906158ec565b97509750975097509750975097509750878b6000019015159081151581525050868b6020018181525050858b6040018181525050848b6060018181525050838b6080018181525050828b60a0018181525050818b60c0018181525050808b60e00181815250505050505050505050505092915050565b6000600080516020615f2083398151915261243a81336137a5565b8483146124645760405163098b37e560e31b81526004810186905260248101849052604401610ee0565b6000805b8681101561266e576000888883818110612484576124846159c4565b905060200201359050600061249882613809565b6004810154909150808989868181106124b3576124b36159c4565b9050602002013510156124d957604051632f789f4960e21b815260040160405180910390fd5b6000808360000160039054906101000a90046001600160a01b03166001600160a01b0316639abddf096040518163ffffffff1660e01b815260040160606040518083038186803b15801561252c57600080fd5b505afa158015612540573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125649190615956565b5091509150808b8b8881811061257c5761257c6159c4565b9050602002013511156125c7578a8a8781811061259b5761259b6159c4565b9050602002013581604051630b72c59d60e21b8152600401610ee0929190918252602082015260400190565b828b8b888181106125da576125da6159c4565b905060200201356125eb91906159da565b6125f590886159f1565b96508282101561263d57847fdd2523ca96a639ba7e17420698937f71eddd8af012ccb36ff5c8fe96141acae961262b84866159da565b60405190815260200160405180910390a25b8a8a8781811061264f5761264f6159c4565b9050602002013584600401819055508560010195505050505050612468565b509695505050505050565b7f779e5c23cb7a5bcb9bfe1e9a5165a00057f12bcdfd13e374540fdf1a1cd911376126a481336137a5565b8382146126ce5760405163098b37e560e31b81526004810185905260248101839052604401610ee0565b60005b848110156128475760008484838181106126ed576126ed6159c4565b90506020020135111561283f57600061271d878784818110612711576127116159c4565b90506020020135613809565b54630100000090046001600160a01b0316905080638d7e4017868685818110612748576127486159c4565b905060200201356040518263ffffffff1660e01b815260040161276d91815260200190565b600060405180830381600087803b15801561278757600080fd5b505af1925050508015612798575060015b61283d573d8080156127c6576040519150601f19603f3d011682016040523d82523d6000602084013e6127cb565b606091505b5080516127eb57604051638fd297d960e01b815260040160405180910390fd5b8787848181106127fd576127fd6159c4565b905060200201357ff74208fedac7280fd11f8de0be14e00423dc5076da8e8ec8ca90e09257fff1b3826040516128339190615beb565b60405180910390a2505b505b6001016126d1565b505050505050565b6060806060600080600080612862613dc3565b80519193509150801580612874575082155b156128b65750506040805160008082526020820181815282840182815260608401909452919850909650909450925068056bc75e2d631000009150612bc39050565b68056bc75e2d631000009350806001600160401b038111156128da576128da614ea0565b604051908082528060200260200182016040528015612903578160200160208202803683370190505b509650806001600160401b0381111561291e5761291e614ea0565b604051908082528060200260200182016040528015612947578160200160208202803683370190505b509750806001600160401b0381111561296257612962614ea0565b60405190808252806020026020018201604052801561298b578160200160208202803683370190505b5095506000808060005b84811015612b905760008682815181106129b1576129b16159c4565b602002602001015160c001511115612b88578581815181106129d5576129d56159c4565b60200260200101516020015162ffffff168b85815181106129f8576129f86159c4565b6020026020010181815250508688878381518110612a1857612a186159c4565b602002602001015160c00151612a2e9190615ae0565b612a3891906159b0565b9250858181518110612a4c57612a4c6159c4565b6020026020010151600001518c8581518110612a6a57612a6a6159c4565b60200260200101906001600160a01b031690816001600160a01b031681525050612710868281518110612a9f57612a9f6159c4565b60200260200101516040015161ffff1684612aba9190615ae0565b612ac491906159b0565b91506002868281518110612ada57612ada6159c4565b602002602001015160a001516002811115612af757612af7615531565b14612b3057818a8581518110612b0f57612b0f6159c4565b60200260200101906001600160601b031690816001600160601b0316815250505b81612710878381518110612b4657612b466159c4565b60200260200101516060015161ffff1685612b619190615ae0565b612b6b91906159b0565b612b759190615bfe565b612b7f908a615bfe565b98506001909301925b600101612995565b5086886001600160601b03161115612baa57612baa615bd5565b83831015612bbc57828a52828b528289525b5050505050505b9091929394565b612bd2614ced565b612bdb82613809565b6040805161014081018252825462ffffff81168252630100000081046001600160a01b03166020830152600160b81b810461ffff90811693830193909352600160c81b810483166060830152600160d81b81049092166080820152600160e81b90910460ff1660a082015260018201805491929160c084019190612c5e90615a90565b80601f0160208091040260200160405190810160405280929190818152602001828054612c8a90615a90565b8015612cd75780601f10612cac57610100808354040283529160200191612cd7565b820191906000526020600020905b815481529060010190602001808311612cba57829003601f168201915b505050918352505060028201546001600160401b031660208201526003820154604082015260049091015460609091015292915050565b60006060612d1b8361381c565b509094909350915050565b600080516020615f20833981519152612d3f81336137a5565b6000612d4a87613809565b54630100000090046001600160a01b03169050612d6986868686613e91565b604051634d8060a360e11b81526001600160a01b03821690639b00c14690612d9b908990899089908990600401615c20565b600060405180830381600087803b158015612db557600080fd5b505af11580156122f2573d6000803e3d6000fd5b6000818152600080516020615ee083398151915260205260408120610d3c90613f37565b600080516020615f20833981519152612e0681336137a5565b6000612e1187613809565b54630100000090046001600160a01b03169050612e3086868686613e91565b604051629b3d1960e81b81526001600160a01b03821690639b3d190090612d9b908990899089908990600401615c20565b600080516020615f40833981519152612e7a81336137a5565b6000612e8584613809565b9050826002811115612e9957612e99615531565b8154600160e81b900460ff166002811115612eb657612eb6615531565b6002811115612ec757612ec7615531565b1415612ee657604051635ca16fa760e11b815260040160405180910390fd5b612ef08184613b0b565b50505050565b612eff82611108565b612f0981336137a5565b6111478383613a69565b7f9a2f67efb89489040f2c48c3b2c38f719fba1276678d2ced3bd9049fb5edc6b2612f3e81336137a5565b6000612f4983613809565b905060018154600160e81b900460ff166002811115612f6a57612f6a615531565b6002811115612f7b57612f7b615531565b14612f99576040516316c1da1560e21b815260040160405180910390fd5b611147816000613b0b565b600080516020615f20833981519152612fbd81336137a5565b6000612fc761169c565b905060005b81811015611147576000612fdf82613a98565b905060008160000160039054906101000a90046001600160a01b031690506000816001600160a01b0316639abddf096040518163ffffffff1660e01b815260040160606040518083038186803b15801561303857600080fd5b505afa15801561304c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906130709190615956565b50509050826004015481141561316357816001600160a01b031663e864299e6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156130bb57600080fd5b505af19250505080156130cc575060015b613163573d8080156130fa576040519150601f19603f3d011682016040523d82523d6000602084013e6130ff565b606091505b50805161311f57604051638fd297d960e01b815260040160405180910390fd5b835460405162ffffff909116907fe74bf895f0c3a2d6c74c40cbb362fdd9640035fc4226c72e3843809ad2a9d2b590613159908490615beb565b60405180910390a2505b836001019350505050612fcc565b600080516020615f4083398151915261318a81336137a5565b600061319585613809565b5460405163a2e080f160e01b8152600481018690526024810185905263010000009091046001600160a01b03169150819063a2e080f190604401600060405180830381600087803b1580156131e957600080fd5b505af1158015610fc1573d6000803e3d6000fd5b60006001611a89565b7fe7c742a54cd11fc9749a47ab34bdcd7327820908e8d0d48b4a5c7f17b029409861323181336137a5565b61325a7fabeb05279af36da5d476d7f950157cd2ea98a4166fa68a6bc97ce3a22fbb93c0839055565b600061326461169c565b905060005b8181101561338f57600061327c82613a98565b90508160010191508060000160039054906101000a90046001600160a01b03166001600160a01b03166390c09bdb6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156132d657600080fd5b505af19250505080156132e7575060015b613389573d808015613315576040519150601f19603f3d011682016040523d82523d6000602084013e61331a565b606091505b50805161333a57604051638fd297d960e01b815260040160405180910390fd5b613345826001613b0b565b815460405162ffffff909116907f0d64b11929aa111ca874dd00b5b0cc2d82b741be924ec9e3691e67c71552f6239061337f908490615beb565b60405180910390a2505b50613269565b50604080518481523360208201527f82e72df77173eab89b00556d791a407a78f4605c5c2f0694967c8c429dd43c7c91016118a9565b60008060008060006133d56136d8565b92506001600160601b031692506001600160601b031692506133f78382613c2c565b94506134038282613c2c565b93505050509091565b6060600061341984613783565b905082516001600160401b0381111561343457613434614ea0565b60405190808252806020026020018201604052801561346d57816020015b61345a614dd6565b8152602001906001900390816134525790505b50915060005b8351811015613598576040518060600160405280858381518110613499576134996159c4565b60200260200101518152602001836001600160a01b0316635e2fb9088785815181106134c7576134c76159c4565b60200260200101516040518263ffffffff1660e01b81526004016134ed91815260200190565b60206040518083038186803b15801561350557600080fd5b505afa158015613519573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061353d9190615c52565b151581526020016135678787858151811061355a5761355a6159c4565b6020026020010151612300565b81525083828151811061357c5761357c6159c4565b60200260200101819052508061359190615ac5565b9050613473565b505092915050565b606060006135ac61169c565b9050806001600160401b038111156135c6576135c6614ea0565b6040519080825280602002602001820160405280156135ef578160200160208202803683370190505b50915060005b81811015611a805761360681613a98565b54835162ffffff90911690849083908110613623576136236159c4565b60209081029190910101526001016135f5565b6060600061364385613783565b604051634febc81b60e01b815260048101869052602481018590529091506000906001600160a01b03831690634febc81b9060440160006040518083038186803b15801561369057600080fd5b505afa1580156136a4573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526136cc9190810190615c6f565b90506110fe868261340c565b6000806000606060006136e961284f565b9650909450925060009150505b825181101561373657828181518110613711576137116159c4565b6020026020010151866137249190615bfe565b955061372f81615ac5565b90506136f6565b506137418582615cf4565b93505050909192565b9055565b60006001600160e01b03198216637965db0b60e01b1480610d3c57506301ffc9a760e01b6001600160e01b0319831614610d3c565b600061378e82613809565b54630100000090046001600160a01b031692915050565b6137af8282611f4e565b61125e576137c7816001600160a01b03166014613f41565b6137d2836020613f41565b6040516020016137e3929190615d1c565b60408051601f198184030181529082905262461bcd60e51b8252610ee091600401615beb565b6000610d3c613817836139e1565b613a98565b6000606080600061382b613dc3565b8051909350909150806001600160401b0381111561384b5761384b614ea0565b604051908082528060200260200182016040528015613874578160200160208202803683370190505b50935080156139d85761388786836159f1565b91506000816001600160401b038111156138a3576138a3614ea0565b6040519080825280602002602001820160405280156138cc578160200160208202803683370190505b5090506000805b838110156139c7578581815181106138ed576138ed6159c4565b602002602001015160c0015187828151811061390b5761390b6159c4565b6020026020010181815250506127108587838151811061392d5761392d6159c4565b60200260200101516080015161ffff166139479190615ae0565b61395191906159b0565b91506139a282878381518110613969576139696159c4565b602002602001015160e00151888481518110613987576139876159c4565b602002602001015160c0015161399d91906159f1565b6140dc565b8382815181106139b4576139b46159c4565b60209081029190910101526001016138d3565b506139d386838a6140eb565b965050505b50509193909250565b60008181527f9b48f5b32acb95b982effe269feac267eead113c4b5af14ffeb9aadac18a6e9c6020819052604082205480613a2f57604051636a0eb14160e11b815260040160405180910390fd5b6111dc6001826159da565b613a448282614130565b6000828152600080516020615ee08339815191526020526040902061114790826141a6565b613a7382826141bb565b6000828152600080516020615ee083398151915260205260409020611147908261422f565b60009081527f1d2f69fc9b5fe89d7414bf039e8d897c4c487c7603d80de6bcdd2868466f94766020526040902090565b7f9b48f5b32acb95b982effe269feac267eead113c4b5af14ffeb9aadac18a6e9c613af48260016159f1565b600093845260209190915260409092209190915550565b8154600090600160e81b900460ff166002811115613b2b57613b2b615531565b9050816002811115613b3f57613b3f615531565b816002811115613b5157613b51615531565b1461114757816002811115613b6857613b68615531565b835460ff91909116600160e81b0260ff60e81b1982168117855560405162ffffff9182169190921617907ffd6f15fb2b48a21a60fe3d44d3c3a0433ca01e121b5124a63ec45c30ad925a1790613bc19085903390615d91565b60405180910390a2505050565b613bd6611d67565b15613bf45760405163184e52a160e21b815260040160405180910390fd5b613bfd81614244565b50565b61125e8282613a3a565b6000611f4783836142a3565b6000818311613c255781611f47565b5090919050565b600081613c3b61271085615ae0565b611f4791906159b0565b613c50846030615ae0565b825114613c86578151613c64856030615ae0565b6040516346b38e7960e11b815260048101929092526024820152604401610ee0565b613c91846060615ae0565b815114613cc7578051613ca5856060615ae0565b604051633c11c1f760e21b815260048101929092526024820152604401610ee0565b6000613cd360306142cd565b90506000613ce160606142cd565b905060005b868110156122f757613d078584613cfe603085615ae0565b600060306142e6565b613d208483613d17606085615ae0565b600060606142e6565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663228951186801bc16d674ec800000858986613d678c8a8a61436d565b6040518663ffffffff1660e01b8152600401613d869493929190615db7565b6000604051808303818588803b158015613d9f57600080fd5b505af1158015613db3573d6000803e3d6000fd5b5050505050806001019050613ce6565b600060606000613dd161169c565b9050806001600160401b03811115613deb57613deb614ea0565b604051908082528060200260200182016040528015613e2457816020015b613e11614df5565b815260200190600190039081613e095790505b50915060005b81811015613e8b57613e3b816146e8565b838281518110613e4d57613e4d6159c4565b6020026020010181905250828181518110613e6a57613e6a6159c4565b602002602001015160c0015184613e8191906159f1565b9350600101613e2a565b50509091565b613e9c600884615e02565b151580613eb25750613eaf601082615e02565b15155b15613ed3576040516363209a7d60e11b815260036004820152602401610ee0565b6000613ee06008856159b0565b905080613eee6010846159b0565b14613f0f576040516363209a7d60e11b815260026004820152602401610ee0565b80613f30576040516363209a7d60e11b815260016004820152602401610ee0565b5050505050565b6000610d3c825490565b60606000613f50836002615ae0565b613f5b9060026159f1565b6001600160401b03811115613f7257613f72614ea0565b6040519080825280601f01601f191660200182016040528015613f9c576020820181803683370190505b509050600360fc1b81600081518110613fb757613fb76159c4565b60200101906001600160f81b031916908160001a905350600f60fb1b81600181518110613fe657613fe66159c4565b60200101906001600160f81b031916908160001a905350600061400a846002615ae0565b6140159060016159f1565b90505b600181111561408d576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110614049576140496159c4565b1a60f81b82828151811061405f5761405f6159c4565b60200101906001600160f81b031916908160001a90535060049490941c9361408681615e16565b9050614018565b508315611f475760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610ee0565b6000818310613c255781611f47565b6000805b828210156141285761410b858561410685876159da565b614867565b90508061411757614128565b61412181836159f1565b91506140ef565b509392505050565b61413a8282611f4e565b61125e576000828152600080516020615f00833981519152602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b6000611f47836001600160a01b038416614aa5565b6141c58282611f4e565b1561125e576000828152600080516020615f00833981519152602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6000611f47836001600160a01b038416614af4565b61426d7f4dd0f6662ba1d6b081f08b350f5e9a6a7b15cf586926ba66f753594928fa64a6829055565b6040518181527ffddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb9060200160405180910390a150565b60008260000182815481106142ba576142ba6159c4565b9060005260206000200154905092915050565b60408051828152603f92810192909201601f1916905290565b84516142f282856159f1565b1115801561430a5750835161430782846159f1565b11155b6143565760405162461bcd60e51b815260206004820152601960248201527f42595445535f41525241595f4f55545f4f465f424f554e4453000000000000006044820152606401610ee0565b6020838601810190838601016122f7828285614be7565b60008061437a60406142cd565b9050600061439261438d604060606159da565b6142cd565b90506143a3848360008060406142e6565b6143bc8482604060006143b78260606159da565b6142e6565b6000600286600060801b6040516020016143d7929190615e2d565b60408051601f19818403018152908290526143f191615e65565b602060405180830381855afa15801561440e573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061443191906158d3565b90506000600280856040516020016144499190615e65565b60408051601f198184030181529082905261446391615e65565b602060405180830381855afa158015614480573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906144a391906158d3565b6040516002906144ba908790600090602001615e81565b60408051601f19818403018152908290526144d491615e65565b602060405180830381855afa1580156144f1573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061451491906158d3565b60408051602081019390935282015260600160408051601f198184030181529082905261454091615e65565b602060405180830381855afa15801561455d573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061458091906158d3565b9050600280838a604051602001614598929190615ea3565b60408051601f19818403018152908290526145b291615e65565b602060405180830381855afa1580156145cf573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906145f291906158d3565b60408051634059730760d81b60208201526000602882015290810184905260029060600160408051601f198184030181529082905261463091615e65565b602060405180830381855afa15801561464d573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061467091906158d3565b60408051602081019390935282015260600160408051601f198184030181529082905261469c91615e65565b602060405180830381855afa1580156146b9573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906146dc91906158d3565b98975050505050505050565b6146f0614df5565b60006146fb83613a98565b80546001600160a01b036301000000820416845262ffffff8116602085015261ffff600160b81b820481166040860152600160c81b820481166060860152600160d81b820416608085015290915060ff600160e81b90910416600281111561476557614765615531565b8260a00190600281111561477b5761477b615531565b9081600281111561478e5761478e615531565b81525050600080600084600001516001600160a01b0316639abddf096040518163ffffffff1660e01b815260040160606040518083038186803b1580156147d457600080fd5b505afa1580156147e8573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061480c9190615956565b9194509250905060008560a00151600281111561482b5761482b615531565b14614837576000614839565b805b60e0860152600484015461484e908490613c16565b61485890836159da565b60c08601525092949350505050565b825160009060001982846148815760009350505050611f47565b60005b87518110156149535786818151811061489f5761489f6159c4565b60200260200101518882815181106148b9576148b96159c4565b6020026020010151106148cb57614943565b8781815181106148dd576148dd6159c4565b60200260200101518311156149145780935060019150878181518110614905576149056159c4565b60200260200101519250614943565b878181518110614926576149266159c4565b6020026020010151831415614943576149406001836159f1565b91505b61494c81615ac5565b9050614884565b50806149655760009350505050611f47565b60001960005b8851811015614a2457878181518110614986576149866159c4565b60200260200101518982815181106149a0576149a06159c4565b6020026020010151106149b257614a14565b838982815181106149c5576149c56159c4565b60200260200101511180156149f25750818982815181106149e8576149e86159c4565b6020026020010151105b15614a1457888181518110614a0957614a096159c4565b602002602001015191505b614a1d81615ac5565b905061496b565b50614a6e60018311614a365786614a40565b614a408784614c32565b84614a64848b8981518110614a5757614a576159c4565b60200260200101516140dc565b61399d91906159da565b945084888581518110614a8357614a836159c4565b60200260200101818151614a9791906159f1565b905250505050509392505050565b6000818152600183016020526040812054614aec57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610d3c565b506000610d3c565b60008181526001830160205260408120548015614bdd576000614b186001836159da565b8554909150600090614b2c906001906159da565b9050818114614b91576000866000018281548110614b4c57614b4c6159c4565b9060005260206000200154905080876000018481548110614b6f57614b6f6159c4565b6000918252602080832090910192909255918252600188019052604090208390555b8554869080614ba257614ba2615ec9565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610d3c565b6000915050610d3c565b5b601f811115614c08578251825260209283019290910190601f1901614be8565b80156111475782518251600019600160086020869003021b01908116901991909116178252505050565b60008215614c605781614c466001856159da565b614c5091906159b0565b614c5b9060016159f1565b611f47565b50600092915050565b828054614c7590615a90565b90600052602060002090601f016020900481019282614c975760008555614cdd565b82601f10614cb05782800160ff19823516178555614cdd565b82800160010185558215614cdd579182015b82811115614cdd578235825591602001919060010190614cc2565b50614ce9929150614e3a565b5090565b604080516101408101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c082015260e08101829052610100810182905261012081019190915290565b60405180608001604052806000815260200160008152602001614d61614ced565b8152602001614d8a60405180606001604052806000815260200160008152602001600081525090565b905290565b604051806101000160405280600015158152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6040805160608101825260008082526020820152908101614d8a614d8f565b604080516101008101825260008082526020820181905291810182905260608101829052608081018290529060a0820190815260200160008152602001600081525090565b5b80821115614ce95760008155600101614e3b565b600060208284031215614e6157600080fd5b81356001600160e01b031981168114611f4757600080fd5b600060208284031215614e8b57600080fd5b5035919050565b8015158114613bfd57600080fd5b634e487b7160e01b600052604160045260246000fd5b60405160c081016001600160401b0381118282101715614ed857614ed8614ea0565b60405290565b604051601f8201601f191681016001600160401b0381118282101715614f0657614f06614ea0565b604052919050565b600080600080848603610120811215614f2657600080fd5b85359450602086013593506040860135614f3f81614e92565b925060c0605f1982011215614f5357600080fd5b50614f5c614eb6565b606086013581526080860135602082015260a0860135604082015260c0860135606082015260e0860135608082015261010086013560a08201528091505092959194509250565b81518152602080830151908201526040808301519082015260608101610d3c565b60008060408385031215614fd757600080fd5b50508035926020909101359150565b80356001600160a01b0381168114614ffd57600080fd5b919050565b6000806040838503121561501557600080fd5b8235915061502560208401614fe6565b90509250929050565b8051151582526020810151602083015260408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301525050565b602080825282518282018190526000919060409081850190868401855b828110156150d757815180518552868101511515878601528501516150c28686018261502e565b5061014093909301929085019060010161509b565b5091979650505050505050565b60008083601f8401126150f657600080fd5b5081356001600160401b0381111561510d57600080fd5b60208301915083602082850101111561512557600080fd5b9250929050565b60008060008060008060a0878903121561514557600080fd5b86356001600160401b0381111561515b57600080fd5b61516789828a016150e4565b909750955061517a905060208801614fe6565b93506040870135925060608701359150608087013590509295509295509295565b60005b838110156151b657818101518382015260200161519e565b83811115612ef05750506000910152565b600081518084526151df81602086016020860161519b565b601f01601f19169290920160200192915050565b805162ffffff1682526000610140602083015161521b60208601826001600160a01b03169052565b506040830151615231604086018261ffff169052565b506060830151615247606086018261ffff169052565b50608083015161525d608086018261ffff169052565b5060a083015161527260a086018260ff169052565b5060c08301518160c086015261528a828601826151c7565b91505060e08301516152a760e08601826001600160401b03169052565b5061010083810151908501526101209283015192909301919091525090565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561535d57603f19898403018552815160c0815185528882015189860152878201518189870152615323828701826151f3565b60609384015180518886015260208101516080890152604081015160a08901529390925090505095880195935050908601906001016152ed565b509098975050505050505050565b60008060006060848603121561538057600080fd5b61538984614fe6565b925061539760208501614fe6565b9150604084013590509250925092565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156153fc57603f198886030184526153ea8583516151f3565b945092850192908501906001016153ce565b5092979650505050505050565b6000806000806080858703121561541f57600080fd5b8435935060208501359250604085013561543881614e92565b9396929550929360600135925050565b60006001600160401b0382111561546157615461614ea0565b5060051b60200190565b600082601f83011261547c57600080fd5b8135602061549161548c83615448565b614ede565b82815260059290921b840181019181810190868411156154b057600080fd5b8286015b8481101561266e57803583529183019183016154b4565b6000602082840312156154dd57600080fd5b81356001600160401b038111156154f357600080fd5b6111dc8482850161546b565b6000806000806080858703121561551557600080fd5b5050823594602084013594506040840135936060013592509050565b634e487b7160e01b600052602160045260246000fd5b6003811061556557634e487b7160e01b600052602160045260246000fd5b9052565b60208101610d3c8284615547565b6000806000806060858703121561558d57600080fd5b843593506020850135925060408501356001600160401b038111156155b157600080fd5b6155bd878288016150e4565b95989497509550505050565b6101008101610d3c828461502e565b60008083601f8401126155ea57600080fd5b5081356001600160401b0381111561560157600080fd5b6020830191508360208260051b850101111561512557600080fd5b6000806000806040858703121561563257600080fd5b84356001600160401b038082111561564957600080fd5b615655888389016155d8565b9096509450602087013591508082111561566e57600080fd5b506155bd878288016155d8565b600081518084526020808501945080840160005b838110156156ab5781518752958201959082019060010161568f565b509495945050505050565b60a0808252865190820181905260009060209060c0840190828a01845b828110156156f85781516001600160a01b0316845292840192908401906001016156d3565b5050508381038285015261570c818961567b565b8481036040860152875180825283890192509083019060005b8181101561574a5783516001600160601b031683529284019291840191600101615725565b50506001600160601b03871660608601529250615765915050565b8260808301529695505050505050565b602081526000611f4760208301846151f3565b8281526040602082015260006111dc604083018461567b565b6000806000806000606086880312156157b957600080fd5b8535945060208601356001600160401b03808211156157d757600080fd5b6157e389838a016150e4565b909650945060408801359150808211156157fc57600080fd5b50615809888289016150e4565b969995985093965092949392505050565b6000806040838503121561582d57600080fd5b8235915060208301356003811061584357600080fd5b809150509250929050565b60008060006060848603121561586357600080fd5b505081359360208301359350604090920135919050565b6000806040838503121561588d57600080fd5b8235915060208301356001600160401b038111156158aa57600080fd5b6158b68582860161546b565b9150509250929050565b602081526000611f47602083018461567b565b6000602082840312156158e557600080fd5b5051919050565b600080600080600080600080610100898b03121561590957600080fd5b885161591481614e92565b809850506020890151965060408901519550606089015194506080890151935060a0890151925060c0890151915060e089015190509295985092959890939650565b60008060006060848603121561596b57600080fd5b8351925060208401519150604084015190509250925092565b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b6000826159bf576159bf615984565b500490565b634e487b7160e01b600052603260045260246000fd5b6000828210156159ec576159ec61599a565b500390565b60008219821115615a0457615a0461599a565b500190565b600062ffffff808316818516808303821115615a2757615a2761599a565b01949350505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b600060018060a01b03808716835260606020840152615a7c606084018688615a30565b915080841660408401525095945050505050565b600181811c90821680615aa457607f821691505b60208210811415611d6157634e487b7160e01b600052602260045260246000fd5b6000600019821415615ad957615ad961599a565b5060010190565b6000816000190483118215151615615afa57615afa61599a565b500290565b83815260406020820152600061203c604083018486615a30565b600082601f830112615b2a57600080fd5b81516001600160401b03811115615b4357615b43614ea0565b615b56601f8201601f1916602001614ede565b818152846020838601011115615b6b57600080fd5b6111dc82602083016020870161519b565b60008060408385031215615b8f57600080fd5b82516001600160401b0380821115615ba657600080fd5b615bb286838701615b19565b93506020850151915080821115615bc857600080fd5b506158b685828601615b19565b634e487b7160e01b600052600160045260246000fd5b602081526000611f4760208301846151c7565b60006001600160601b03808316818516808303821115615a2757615a2761599a565b604081526000615c34604083018688615a30565b8281036020840152615c47818587615a30565b979650505050505050565b600060208284031215615c6457600080fd5b8151611f4781614e92565b60006020808385031215615c8257600080fd5b82516001600160401b03811115615c9857600080fd5b8301601f81018513615ca957600080fd5b8051615cb761548c82615448565b81815260059190911b82018301908381019087831115615cd657600080fd5b928401925b82841015615c4757835182529284019290840190615cdb565b60006001600160601b0383811690831681811015615d1457615d1461599a565b039392505050565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351615d5481601785016020880161519b565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351615d8581602884016020880161519b565b01602801949350505050565b60408101615d9f8285615547565b6001600160a01b039290921660209190910152919050565b608081526000615dca60808301876151c7565b8281036020840152615ddc81876151c7565b90508281036040840152615df081866151c7565b91505082606083015295945050505050565b600082615e1157615e11615984565b500690565b600081615e2557615e2561599a565b506000190190565b60008351615e3f81846020880161519b565b6fffffffffffffffffffffffffffffffff19939093169190920190815260100192915050565b60008251615e7781846020870161519b565b9190910192915050565b60008351615e9381846020880161519b565b9190910191825250602001919050565b82815260008251615ebb81602085016020870161519b565b919091016020019392505050565b634e487b7160e01b600052603160045260246000fdfe8f8c450dae5029cd48cd91dd9db65da48fb742893edfc7941250f6721d93cbbe9a627a5d4aa7c17f87ff26e3fe9a42c2b6c559e8b41a42282d0ecebb17c0e4d3c23292b191d95d2a7dd94fc6436eb44338fda9e1307d9394fd27c28157c1b33c3105bcbf19d4417b73ae0e58d508a65ecf75665e46c2622d8521732de6080c48a264697066735822122093f74b570f38204664e512a5339a76eebcc26abe3394e4e21b014b1dd0eedf8964736f6c63430008090033",
}

// ApiABI is the input ABI used to generate the binding from.
// Deprecated: Use ApiMetaData.ABI instead.
var ApiABI = ApiMetaData.ABI

// ApiBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ApiMetaData.Bin instead.
var ApiBin = ApiMetaData.Bin

// DeployApi deploys a new Ethereum contract, binding an instance of Api to it.
func DeployApi(auth *bind.TransactOpts, backend bind.ContractBackend, _depositContract common.Address) (common.Address, *types.Transaction, *Api, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ApiBin), backend, _depositContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// Api is an auto generated Go binding around an Ethereum contract.
type Api struct {
	ApiCaller     // Read-only binding to the contract
	ApiTransactor // Write-only binding to the contract
	ApiFilterer   // Log filterer for contract events
}

// ApiCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApiCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApiTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApiFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApiSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApiSession struct {
	Contract     *Api              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApiCallerSession struct {
	Contract *ApiCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ApiTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApiTransactorSession struct {
	Contract     *ApiTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApiRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApiRaw struct {
	Contract *Api // Generic contract binding to access the raw methods on
}

// ApiCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApiCallerRaw struct {
	Contract *ApiCaller // Generic read-only contract binding to access the raw methods on
}

// ApiTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApiTransactorRaw struct {
	Contract *ApiTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApi creates a new instance of Api, bound to a specific deployed contract.
func NewApi(address common.Address, backend bind.ContractBackend) (*Api, error) {
	contract, err := bindApi(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Api{ApiCaller: ApiCaller{contract: contract}, ApiTransactor: ApiTransactor{contract: contract}, ApiFilterer: ApiFilterer{contract: contract}}, nil
}

// NewApiCaller creates a new read-only instance of Api, bound to a specific deployed contract.
func NewApiCaller(address common.Address, caller bind.ContractCaller) (*ApiCaller, error) {
	contract, err := bindApi(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApiCaller{contract: contract}, nil
}

// NewApiTransactor creates a new write-only instance of Api, bound to a specific deployed contract.
func NewApiTransactor(address common.Address, transactor bind.ContractTransactor) (*ApiTransactor, error) {
	contract, err := bindApi(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApiTransactor{contract: contract}, nil
}

// NewApiFilterer creates a new log filterer instance of Api, bound to a specific deployed contract.
func NewApiFilterer(address common.Address, filterer bind.ContractFilterer) (*ApiFilterer, error) {
	contract, err := bindApi(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApiFilterer{contract: contract}, nil
}

// bindApi binds a generic wrapper to an already deployed contract.
func bindApi(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApiMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.ApiCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.ApiTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Api *ApiCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Api.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Api *ApiTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Api *ApiTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Api.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Api *ApiCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Api *ApiSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Api.Contract.DEFAULTADMINROLE(&_Api.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Api *ApiCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Api.Contract.DEFAULTADMINROLE(&_Api.CallOpts)
}

// DEPOSITCONTRACT is a free data retrieval call binding the contract method 0x6b96736b.
//
// Solidity: function DEPOSIT_CONTRACT() view returns(address)
func (_Api *ApiCaller) DEPOSITCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "DEPOSIT_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITCONTRACT is a free data retrieval call binding the contract method 0x6b96736b.
//
// Solidity: function DEPOSIT_CONTRACT() view returns(address)
func (_Api *ApiSession) DEPOSITCONTRACT() (common.Address, error) {
	return _Api.Contract.DEPOSITCONTRACT(&_Api.CallOpts)
}

// DEPOSITCONTRACT is a free data retrieval call binding the contract method 0x6b96736b.
//
// Solidity: function DEPOSIT_CONTRACT() view returns(address)
func (_Api *ApiCallerSession) DEPOSITCONTRACT() (common.Address, error) {
	return _Api.Contract.DEPOSITCONTRACT(&_Api.CallOpts)
}

// FEEPRECISIONPOINTS is a free data retrieval call binding the contract method 0x7c8da51c.
//
// Solidity: function FEE_PRECISION_POINTS() view returns(uint256)
func (_Api *ApiCaller) FEEPRECISIONPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "FEE_PRECISION_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEPRECISIONPOINTS is a free data retrieval call binding the contract method 0x7c8da51c.
//
// Solidity: function FEE_PRECISION_POINTS() view returns(uint256)
func (_Api *ApiSession) FEEPRECISIONPOINTS() (*big.Int, error) {
	return _Api.Contract.FEEPRECISIONPOINTS(&_Api.CallOpts)
}

// FEEPRECISIONPOINTS is a free data retrieval call binding the contract method 0x7c8da51c.
//
// Solidity: function FEE_PRECISION_POINTS() view returns(uint256)
func (_Api *ApiCallerSession) FEEPRECISIONPOINTS() (*big.Int, error) {
	return _Api.Contract.FEEPRECISIONPOINTS(&_Api.CallOpts)
}

// MANAGEWITHDRAWALCREDENTIALSROLE is a free data retrieval call binding the contract method 0x7a74884d.
//
// Solidity: function MANAGE_WITHDRAWAL_CREDENTIALS_ROLE() view returns(bytes32)
func (_Api *ApiCaller) MANAGEWITHDRAWALCREDENTIALSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "MANAGE_WITHDRAWAL_CREDENTIALS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEWITHDRAWALCREDENTIALSROLE is a free data retrieval call binding the contract method 0x7a74884d.
//
// Solidity: function MANAGE_WITHDRAWAL_CREDENTIALS_ROLE() view returns(bytes32)
func (_Api *ApiSession) MANAGEWITHDRAWALCREDENTIALSROLE() ([32]byte, error) {
	return _Api.Contract.MANAGEWITHDRAWALCREDENTIALSROLE(&_Api.CallOpts)
}

// MANAGEWITHDRAWALCREDENTIALSROLE is a free data retrieval call binding the contract method 0x7a74884d.
//
// Solidity: function MANAGE_WITHDRAWAL_CREDENTIALS_ROLE() view returns(bytes32)
func (_Api *ApiCallerSession) MANAGEWITHDRAWALCREDENTIALSROLE() ([32]byte, error) {
	return _Api.Contract.MANAGEWITHDRAWALCREDENTIALSROLE(&_Api.CallOpts)
}

// MAXSTAKINGMODULESCOUNT is a free data retrieval call binding the contract method 0x4b3a1cb7.
//
// Solidity: function MAX_STAKING_MODULES_COUNT() view returns(uint256)
func (_Api *ApiCaller) MAXSTAKINGMODULESCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "MAX_STAKING_MODULES_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTAKINGMODULESCOUNT is a free data retrieval call binding the contract method 0x4b3a1cb7.
//
// Solidity: function MAX_STAKING_MODULES_COUNT() view returns(uint256)
func (_Api *ApiSession) MAXSTAKINGMODULESCOUNT() (*big.Int, error) {
	return _Api.Contract.MAXSTAKINGMODULESCOUNT(&_Api.CallOpts)
}

// MAXSTAKINGMODULESCOUNT is a free data retrieval call binding the contract method 0x4b3a1cb7.
//
// Solidity: function MAX_STAKING_MODULES_COUNT() view returns(uint256)
func (_Api *ApiCallerSession) MAXSTAKINGMODULESCOUNT() (*big.Int, error) {
	return _Api.Contract.MAXSTAKINGMODULESCOUNT(&_Api.CallOpts)
}

// MAXSTAKINGMODULENAMELENGTH is a free data retrieval call binding the contract method 0x9b75b4ef.
//
// Solidity: function MAX_STAKING_MODULE_NAME_LENGTH() view returns(uint256)
func (_Api *ApiCaller) MAXSTAKINGMODULENAMELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "MAX_STAKING_MODULE_NAME_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTAKINGMODULENAMELENGTH is a free data retrieval call binding the contract method 0x9b75b4ef.
//
// Solidity: function MAX_STAKING_MODULE_NAME_LENGTH() view returns(uint256)
func (_Api *ApiSession) MAXSTAKINGMODULENAMELENGTH() (*big.Int, error) {
	return _Api.Contract.MAXSTAKINGMODULENAMELENGTH(&_Api.CallOpts)
}

// MAXSTAKINGMODULENAMELENGTH is a free data retrieval call binding the contract method 0x9b75b4ef.
//
// Solidity: function MAX_STAKING_MODULE_NAME_LENGTH() view returns(uint256)
func (_Api *ApiCallerSession) MAXSTAKINGMODULENAMELENGTH() (*big.Int, error) {
	return _Api.Contract.MAXSTAKINGMODULENAMELENGTH(&_Api.CallOpts)
}

// REPORTEXITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xc445ea75.
//
// Solidity: function REPORT_EXITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Api *ApiCaller) REPORTEXITEDVALIDATORSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "REPORT_EXITED_VALIDATORS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTEXITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xc445ea75.
//
// Solidity: function REPORT_EXITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Api *ApiSession) REPORTEXITEDVALIDATORSROLE() ([32]byte, error) {
	return _Api.Contract.REPORTEXITEDVALIDATORSROLE(&_Api.CallOpts)
}

// REPORTEXITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xc445ea75.
//
// Solidity: function REPORT_EXITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Api *ApiCallerSession) REPORTEXITEDVALIDATORSROLE() ([32]byte, error) {
	return _Api.Contract.REPORTEXITEDVALIDATORSROLE(&_Api.CallOpts)
}

// REPORTREWARDSMINTEDROLE is a free data retrieval call binding the contract method 0x1d1b9d3c.
//
// Solidity: function REPORT_REWARDS_MINTED_ROLE() view returns(bytes32)
func (_Api *ApiCaller) REPORTREWARDSMINTEDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "REPORT_REWARDS_MINTED_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTREWARDSMINTEDROLE is a free data retrieval call binding the contract method 0x1d1b9d3c.
//
// Solidity: function REPORT_REWARDS_MINTED_ROLE() view returns(bytes32)
func (_Api *ApiSession) REPORTREWARDSMINTEDROLE() ([32]byte, error) {
	return _Api.Contract.REPORTREWARDSMINTEDROLE(&_Api.CallOpts)
}

// REPORTREWARDSMINTEDROLE is a free data retrieval call binding the contract method 0x1d1b9d3c.
//
// Solidity: function REPORT_REWARDS_MINTED_ROLE() view returns(bytes32)
func (_Api *ApiCallerSession) REPORTREWARDSMINTEDROLE() ([32]byte, error) {
	return _Api.Contract.REPORTREWARDSMINTEDROLE(&_Api.CallOpts)
}

// STAKINGMODULEMANAGEROLE is a free data retrieval call binding the contract method 0xe016e6f7.
//
// Solidity: function STAKING_MODULE_MANAGE_ROLE() view returns(bytes32)
func (_Api *ApiCaller) STAKINGMODULEMANAGEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "STAKING_MODULE_MANAGE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMODULEMANAGEROLE is a free data retrieval call binding the contract method 0xe016e6f7.
//
// Solidity: function STAKING_MODULE_MANAGE_ROLE() view returns(bytes32)
func (_Api *ApiSession) STAKINGMODULEMANAGEROLE() ([32]byte, error) {
	return _Api.Contract.STAKINGMODULEMANAGEROLE(&_Api.CallOpts)
}

// STAKINGMODULEMANAGEROLE is a free data retrieval call binding the contract method 0xe016e6f7.
//
// Solidity: function STAKING_MODULE_MANAGE_ROLE() view returns(bytes32)
func (_Api *ApiCallerSession) STAKINGMODULEMANAGEROLE() ([32]byte, error) {
	return _Api.Contract.STAKINGMODULEMANAGEROLE(&_Api.CallOpts)
}

// STAKINGMODULEPAUSEROLE is a free data retrieval call binding the contract method 0xa7357c8c.
//
// Solidity: function STAKING_MODULE_PAUSE_ROLE() view returns(bytes32)
func (_Api *ApiCaller) STAKINGMODULEPAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "STAKING_MODULE_PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMODULEPAUSEROLE is a free data retrieval call binding the contract method 0xa7357c8c.
//
// Solidity: function STAKING_MODULE_PAUSE_ROLE() view returns(bytes32)
func (_Api *ApiSession) STAKINGMODULEPAUSEROLE() ([32]byte, error) {
	return _Api.Contract.STAKINGMODULEPAUSEROLE(&_Api.CallOpts)
}

// STAKINGMODULEPAUSEROLE is a free data retrieval call binding the contract method 0xa7357c8c.
//
// Solidity: function STAKING_MODULE_PAUSE_ROLE() view returns(bytes32)
func (_Api *ApiCallerSession) STAKINGMODULEPAUSEROLE() ([32]byte, error) {
	return _Api.Contract.STAKINGMODULEPAUSEROLE(&_Api.CallOpts)
}

// STAKINGMODULERESUMEROLE is a free data retrieval call binding the contract method 0x8801da79.
//
// Solidity: function STAKING_MODULE_RESUME_ROLE() view returns(bytes32)
func (_Api *ApiCaller) STAKINGMODULERESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "STAKING_MODULE_RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMODULERESUMEROLE is a free data retrieval call binding the contract method 0x8801da79.
//
// Solidity: function STAKING_MODULE_RESUME_ROLE() view returns(bytes32)
func (_Api *ApiSession) STAKINGMODULERESUMEROLE() ([32]byte, error) {
	return _Api.Contract.STAKINGMODULERESUMEROLE(&_Api.CallOpts)
}

// STAKINGMODULERESUMEROLE is a free data retrieval call binding the contract method 0x8801da79.
//
// Solidity: function STAKING_MODULE_RESUME_ROLE() view returns(bytes32)
func (_Api *ApiCallerSession) STAKINGMODULERESUMEROLE() ([32]byte, error) {
	return _Api.Contract.STAKINGMODULERESUMEROLE(&_Api.CallOpts)
}

// TOTALBASISPOINTS is a free data retrieval call binding the contract method 0x271662ec.
//
// Solidity: function TOTAL_BASIS_POINTS() view returns(uint256)
func (_Api *ApiCaller) TOTALBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "TOTAL_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALBASISPOINTS is a free data retrieval call binding the contract method 0x271662ec.
//
// Solidity: function TOTAL_BASIS_POINTS() view returns(uint256)
func (_Api *ApiSession) TOTALBASISPOINTS() (*big.Int, error) {
	return _Api.Contract.TOTALBASISPOINTS(&_Api.CallOpts)
}

// TOTALBASISPOINTS is a free data retrieval call binding the contract method 0x271662ec.
//
// Solidity: function TOTAL_BASIS_POINTS() view returns(uint256)
func (_Api *ApiCallerSession) TOTALBASISPOINTS() (*big.Int, error) {
	return _Api.Contract.TOTALBASISPOINTS(&_Api.CallOpts)
}

// UNSAFESETEXITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0x1565d2f2.
//
// Solidity: function UNSAFE_SET_EXITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Api *ApiCaller) UNSAFESETEXITEDVALIDATORSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "UNSAFE_SET_EXITED_VALIDATORS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNSAFESETEXITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0x1565d2f2.
//
// Solidity: function UNSAFE_SET_EXITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Api *ApiSession) UNSAFESETEXITEDVALIDATORSROLE() ([32]byte, error) {
	return _Api.Contract.UNSAFESETEXITEDVALIDATORSROLE(&_Api.CallOpts)
}

// UNSAFESETEXITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0x1565d2f2.
//
// Solidity: function UNSAFE_SET_EXITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Api *ApiCallerSession) UNSAFESETEXITEDVALIDATORSROLE() ([32]byte, error) {
	return _Api.Contract.UNSAFESETEXITEDVALIDATORSROLE(&_Api.CallOpts)
}

// GetAllNodeOperatorDigests is a free data retrieval call binding the contract method 0x3240a322.
//
// Solidity: function getAllNodeOperatorDigests(uint256 _stakingModuleId) view returns((uint256,bool,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Api *ApiCaller) GetAllNodeOperatorDigests(opts *bind.CallOpts, _stakingModuleId *big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getAllNodeOperatorDigests", _stakingModuleId)

	if err != nil {
		return *new([]StakingRouterNodeOperatorDigest), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingRouterNodeOperatorDigest)).(*[]StakingRouterNodeOperatorDigest)

	return out0, err

}

// GetAllNodeOperatorDigests is a free data retrieval call binding the contract method 0x3240a322.
//
// Solidity: function getAllNodeOperatorDigests(uint256 _stakingModuleId) view returns((uint256,bool,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Api *ApiSession) GetAllNodeOperatorDigests(_stakingModuleId *big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	return _Api.Contract.GetAllNodeOperatorDigests(&_Api.CallOpts, _stakingModuleId)
}

// GetAllNodeOperatorDigests is a free data retrieval call binding the contract method 0x3240a322.
//
// Solidity: function getAllNodeOperatorDigests(uint256 _stakingModuleId) view returns((uint256,bool,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Api *ApiCallerSession) GetAllNodeOperatorDigests(_stakingModuleId *big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	return _Api.Contract.GetAllNodeOperatorDigests(&_Api.CallOpts, _stakingModuleId)
}

// GetAllStakingModuleDigests is a free data retrieval call binding the contract method 0x57993b85.
//
// Solidity: function getAllStakingModuleDigests() view returns((uint256,uint256,(uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256),(uint256,uint256,uint256))[])
func (_Api *ApiCaller) GetAllStakingModuleDigests(opts *bind.CallOpts) ([]StakingRouterStakingModuleDigest, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getAllStakingModuleDigests")

	if err != nil {
		return *new([]StakingRouterStakingModuleDigest), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingRouterStakingModuleDigest)).(*[]StakingRouterStakingModuleDigest)

	return out0, err

}

// GetAllStakingModuleDigests is a free data retrieval call binding the contract method 0x57993b85.
//
// Solidity: function getAllStakingModuleDigests() view returns((uint256,uint256,(uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256),(uint256,uint256,uint256))[])
func (_Api *ApiSession) GetAllStakingModuleDigests() ([]StakingRouterStakingModuleDigest, error) {
	return _Api.Contract.GetAllStakingModuleDigests(&_Api.CallOpts)
}

// GetAllStakingModuleDigests is a free data retrieval call binding the contract method 0x57993b85.
//
// Solidity: function getAllStakingModuleDigests() view returns((uint256,uint256,(uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256),(uint256,uint256,uint256))[])
func (_Api *ApiCallerSession) GetAllStakingModuleDigests() ([]StakingRouterStakingModuleDigest, error) {
	return _Api.Contract.GetAllStakingModuleDigests(&_Api.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Api *ApiCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Api *ApiSession) GetContractVersion() (*big.Int, error) {
	return _Api.Contract.GetContractVersion(&_Api.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Api *ApiCallerSession) GetContractVersion() (*big.Int, error) {
	return _Api.Contract.GetContractVersion(&_Api.CallOpts)
}

// GetDepositsAllocation is a free data retrieval call binding the contract method 0xc82b1bb1.
//
// Solidity: function getDepositsAllocation(uint256 _depositsCount) view returns(uint256 allocated, uint256[] allocations)
func (_Api *ApiCaller) GetDepositsAllocation(opts *bind.CallOpts, _depositsCount *big.Int) (struct {
	Allocated   *big.Int
	Allocations []*big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getDepositsAllocation", _depositsCount)

	outstruct := new(struct {
		Allocated   *big.Int
		Allocations []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Allocated = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Allocations = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetDepositsAllocation is a free data retrieval call binding the contract method 0xc82b1bb1.
//
// Solidity: function getDepositsAllocation(uint256 _depositsCount) view returns(uint256 allocated, uint256[] allocations)
func (_Api *ApiSession) GetDepositsAllocation(_depositsCount *big.Int) (struct {
	Allocated   *big.Int
	Allocations []*big.Int
}, error) {
	return _Api.Contract.GetDepositsAllocation(&_Api.CallOpts, _depositsCount)
}

// GetDepositsAllocation is a free data retrieval call binding the contract method 0xc82b1bb1.
//
// Solidity: function getDepositsAllocation(uint256 _depositsCount) view returns(uint256 allocated, uint256[] allocations)
func (_Api *ApiCallerSession) GetDepositsAllocation(_depositsCount *big.Int) (struct {
	Allocated   *big.Int
	Allocations []*big.Int
}, error) {
	return _Api.Contract.GetDepositsAllocation(&_Api.CallOpts, _depositsCount)
}

// GetLido is a free data retrieval call binding the contract method 0x6a516b47.
//
// Solidity: function getLido() view returns(address)
func (_Api *ApiCaller) GetLido(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getLido")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLido is a free data retrieval call binding the contract method 0x6a516b47.
//
// Solidity: function getLido() view returns(address)
func (_Api *ApiSession) GetLido() (common.Address, error) {
	return _Api.Contract.GetLido(&_Api.CallOpts)
}

// GetLido is a free data retrieval call binding the contract method 0x6a516b47.
//
// Solidity: function getLido() view returns(address)
func (_Api *ApiCallerSession) GetLido() (common.Address, error) {
	return _Api.Contract.GetLido(&_Api.CallOpts)
}

// GetNodeOperatorDigests is a free data retrieval call binding the contract method 0xf07ff28a.
//
// Solidity: function getNodeOperatorDigests(uint256 _stakingModuleId, uint256[] _nodeOperatorIds) view returns((uint256,bool,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[] digests)
func (_Api *ApiCaller) GetNodeOperatorDigests(opts *bind.CallOpts, _stakingModuleId *big.Int, _nodeOperatorIds []*big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getNodeOperatorDigests", _stakingModuleId, _nodeOperatorIds)

	if err != nil {
		return *new([]StakingRouterNodeOperatorDigest), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingRouterNodeOperatorDigest)).(*[]StakingRouterNodeOperatorDigest)

	return out0, err

}

// GetNodeOperatorDigests is a free data retrieval call binding the contract method 0xf07ff28a.
//
// Solidity: function getNodeOperatorDigests(uint256 _stakingModuleId, uint256[] _nodeOperatorIds) view returns((uint256,bool,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[] digests)
func (_Api *ApiSession) GetNodeOperatorDigests(_stakingModuleId *big.Int, _nodeOperatorIds []*big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	return _Api.Contract.GetNodeOperatorDigests(&_Api.CallOpts, _stakingModuleId, _nodeOperatorIds)
}

// GetNodeOperatorDigests is a free data retrieval call binding the contract method 0xf07ff28a.
//
// Solidity: function getNodeOperatorDigests(uint256 _stakingModuleId, uint256[] _nodeOperatorIds) view returns((uint256,bool,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[] digests)
func (_Api *ApiCallerSession) GetNodeOperatorDigests(_stakingModuleId *big.Int, _nodeOperatorIds []*big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	return _Api.Contract.GetNodeOperatorDigests(&_Api.CallOpts, _stakingModuleId, _nodeOperatorIds)
}

// GetNodeOperatorDigests0 is a free data retrieval call binding the contract method 0xf8bb6d42.
//
// Solidity: function getNodeOperatorDigests(uint256 _stakingModuleId, uint256 _offset, uint256 _limit) view returns((uint256,bool,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Api *ApiCaller) GetNodeOperatorDigests0(opts *bind.CallOpts, _stakingModuleId *big.Int, _offset *big.Int, _limit *big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getNodeOperatorDigests0", _stakingModuleId, _offset, _limit)

	if err != nil {
		return *new([]StakingRouterNodeOperatorDigest), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingRouterNodeOperatorDigest)).(*[]StakingRouterNodeOperatorDigest)

	return out0, err

}

// GetNodeOperatorDigests0 is a free data retrieval call binding the contract method 0xf8bb6d42.
//
// Solidity: function getNodeOperatorDigests(uint256 _stakingModuleId, uint256 _offset, uint256 _limit) view returns((uint256,bool,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Api *ApiSession) GetNodeOperatorDigests0(_stakingModuleId *big.Int, _offset *big.Int, _limit *big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	return _Api.Contract.GetNodeOperatorDigests0(&_Api.CallOpts, _stakingModuleId, _offset, _limit)
}

// GetNodeOperatorDigests0 is a free data retrieval call binding the contract method 0xf8bb6d42.
//
// Solidity: function getNodeOperatorDigests(uint256 _stakingModuleId, uint256 _offset, uint256 _limit) view returns((uint256,bool,(bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Api *ApiCallerSession) GetNodeOperatorDigests0(_stakingModuleId *big.Int, _offset *big.Int, _limit *big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	return _Api.Contract.GetNodeOperatorDigests0(&_Api.CallOpts, _stakingModuleId, _offset, _limit)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xaa5a1b9d.
//
// Solidity: function getNodeOperatorSummary(uint256 _stakingModuleId, uint256 _nodeOperatorId) view returns((bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256) summary)
func (_Api *ApiCaller) GetNodeOperatorSummary(opts *bind.CallOpts, _stakingModuleId *big.Int, _nodeOperatorId *big.Int) (StakingRouterNodeOperatorSummary, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getNodeOperatorSummary", _stakingModuleId, _nodeOperatorId)

	if err != nil {
		return *new(StakingRouterNodeOperatorSummary), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingRouterNodeOperatorSummary)).(*StakingRouterNodeOperatorSummary)

	return out0, err

}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xaa5a1b9d.
//
// Solidity: function getNodeOperatorSummary(uint256 _stakingModuleId, uint256 _nodeOperatorId) view returns((bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256) summary)
func (_Api *ApiSession) GetNodeOperatorSummary(_stakingModuleId *big.Int, _nodeOperatorId *big.Int) (StakingRouterNodeOperatorSummary, error) {
	return _Api.Contract.GetNodeOperatorSummary(&_Api.CallOpts, _stakingModuleId, _nodeOperatorId)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xaa5a1b9d.
//
// Solidity: function getNodeOperatorSummary(uint256 _stakingModuleId, uint256 _nodeOperatorId) view returns((bool,uint256,uint256,uint256,uint256,uint256,uint256,uint256) summary)
func (_Api *ApiCallerSession) GetNodeOperatorSummary(_stakingModuleId *big.Int, _nodeOperatorId *big.Int) (StakingRouterNodeOperatorSummary, error) {
	return _Api.Contract.GetNodeOperatorSummary(&_Api.CallOpts, _stakingModuleId, _nodeOperatorId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Api *ApiCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Api *ApiSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Api.Contract.GetRoleAdmin(&_Api.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Api *ApiCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Api.Contract.GetRoleAdmin(&_Api.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Api *ApiCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Api *ApiSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Api.Contract.GetRoleMember(&_Api.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Api *ApiCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Api.Contract.GetRoleMember(&_Api.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Api *ApiCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Api *ApiSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Api.Contract.GetRoleMemberCount(&_Api.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Api *ApiCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Api.Contract.GetRoleMemberCount(&_Api.CallOpts, role)
}

// GetStakingFeeAggregateDistribution is a free data retrieval call binding the contract method 0xfa5093eb.
//
// Solidity: function getStakingFeeAggregateDistribution() view returns(uint96 modulesFee, uint96 treasuryFee, uint256 basePrecision)
func (_Api *ApiCaller) GetStakingFeeAggregateDistribution(opts *bind.CallOpts) (struct {
	ModulesFee    *big.Int
	TreasuryFee   *big.Int
	BasePrecision *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingFeeAggregateDistribution")

	outstruct := new(struct {
		ModulesFee    *big.Int
		TreasuryFee   *big.Int
		BasePrecision *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ModulesFee = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TreasuryFee = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.BasePrecision = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakingFeeAggregateDistribution is a free data retrieval call binding the contract method 0xfa5093eb.
//
// Solidity: function getStakingFeeAggregateDistribution() view returns(uint96 modulesFee, uint96 treasuryFee, uint256 basePrecision)
func (_Api *ApiSession) GetStakingFeeAggregateDistribution() (struct {
	ModulesFee    *big.Int
	TreasuryFee   *big.Int
	BasePrecision *big.Int
}, error) {
	return _Api.Contract.GetStakingFeeAggregateDistribution(&_Api.CallOpts)
}

// GetStakingFeeAggregateDistribution is a free data retrieval call binding the contract method 0xfa5093eb.
//
// Solidity: function getStakingFeeAggregateDistribution() view returns(uint96 modulesFee, uint96 treasuryFee, uint256 basePrecision)
func (_Api *ApiCallerSession) GetStakingFeeAggregateDistribution() (struct {
	ModulesFee    *big.Int
	TreasuryFee   *big.Int
	BasePrecision *big.Int
}, error) {
	return _Api.Contract.GetStakingFeeAggregateDistribution(&_Api.CallOpts)
}

// GetStakingFeeAggregateDistributionE4Precision is a free data retrieval call binding the contract method 0xefcdcc0e.
//
// Solidity: function getStakingFeeAggregateDistributionE4Precision() view returns(uint16 modulesFee, uint16 treasuryFee)
func (_Api *ApiCaller) GetStakingFeeAggregateDistributionE4Precision(opts *bind.CallOpts) (struct {
	ModulesFee  uint16
	TreasuryFee uint16
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingFeeAggregateDistributionE4Precision")

	outstruct := new(struct {
		ModulesFee  uint16
		TreasuryFee uint16
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ModulesFee = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.TreasuryFee = *abi.ConvertType(out[1], new(uint16)).(*uint16)

	return *outstruct, err

}

// GetStakingFeeAggregateDistributionE4Precision is a free data retrieval call binding the contract method 0xefcdcc0e.
//
// Solidity: function getStakingFeeAggregateDistributionE4Precision() view returns(uint16 modulesFee, uint16 treasuryFee)
func (_Api *ApiSession) GetStakingFeeAggregateDistributionE4Precision() (struct {
	ModulesFee  uint16
	TreasuryFee uint16
}, error) {
	return _Api.Contract.GetStakingFeeAggregateDistributionE4Precision(&_Api.CallOpts)
}

// GetStakingFeeAggregateDistributionE4Precision is a free data retrieval call binding the contract method 0xefcdcc0e.
//
// Solidity: function getStakingFeeAggregateDistributionE4Precision() view returns(uint16 modulesFee, uint16 treasuryFee)
func (_Api *ApiCallerSession) GetStakingFeeAggregateDistributionE4Precision() (struct {
	ModulesFee  uint16
	TreasuryFee uint16
}, error) {
	return _Api.Contract.GetStakingFeeAggregateDistributionE4Precision(&_Api.CallOpts)
}

// GetStakingModule is a free data retrieval call binding the contract method 0xbc1bb190.
//
// Solidity: function getStakingModule(uint256 _stakingModuleId) view returns((uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256))
func (_Api *ApiCaller) GetStakingModule(opts *bind.CallOpts, _stakingModuleId *big.Int) (StakingRouterStakingModule, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModule", _stakingModuleId)

	if err != nil {
		return *new(StakingRouterStakingModule), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingRouterStakingModule)).(*StakingRouterStakingModule)

	return out0, err

}

// GetStakingModule is a free data retrieval call binding the contract method 0xbc1bb190.
//
// Solidity: function getStakingModule(uint256 _stakingModuleId) view returns((uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256))
func (_Api *ApiSession) GetStakingModule(_stakingModuleId *big.Int) (StakingRouterStakingModule, error) {
	return _Api.Contract.GetStakingModule(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModule is a free data retrieval call binding the contract method 0xbc1bb190.
//
// Solidity: function getStakingModule(uint256 _stakingModuleId) view returns((uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256))
func (_Api *ApiCallerSession) GetStakingModule(_stakingModuleId *big.Int) (StakingRouterStakingModule, error) {
	return _Api.Contract.GetStakingModule(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleActiveValidatorsCount is a free data retrieval call binding the contract method 0x96b5d81c.
//
// Solidity: function getStakingModuleActiveValidatorsCount(uint256 _stakingModuleId) view returns(uint256 activeValidatorsCount)
func (_Api *ApiCaller) GetStakingModuleActiveValidatorsCount(opts *bind.CallOpts, _stakingModuleId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModuleActiveValidatorsCount", _stakingModuleId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingModuleActiveValidatorsCount is a free data retrieval call binding the contract method 0x96b5d81c.
//
// Solidity: function getStakingModuleActiveValidatorsCount(uint256 _stakingModuleId) view returns(uint256 activeValidatorsCount)
func (_Api *ApiSession) GetStakingModuleActiveValidatorsCount(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Api.Contract.GetStakingModuleActiveValidatorsCount(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleActiveValidatorsCount is a free data retrieval call binding the contract method 0x96b5d81c.
//
// Solidity: function getStakingModuleActiveValidatorsCount(uint256 _stakingModuleId) view returns(uint256 activeValidatorsCount)
func (_Api *ApiCallerSession) GetStakingModuleActiveValidatorsCount(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Api.Contract.GetStakingModuleActiveValidatorsCount(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleDigests is a free data retrieval call binding the contract method 0x8525e3a1.
//
// Solidity: function getStakingModuleDigests(uint256[] _stakingModuleIds) view returns((uint256,uint256,(uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256),(uint256,uint256,uint256))[] digests)
func (_Api *ApiCaller) GetStakingModuleDigests(opts *bind.CallOpts, _stakingModuleIds []*big.Int) ([]StakingRouterStakingModuleDigest, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModuleDigests", _stakingModuleIds)

	if err != nil {
		return *new([]StakingRouterStakingModuleDigest), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingRouterStakingModuleDigest)).(*[]StakingRouterStakingModuleDigest)

	return out0, err

}

// GetStakingModuleDigests is a free data retrieval call binding the contract method 0x8525e3a1.
//
// Solidity: function getStakingModuleDigests(uint256[] _stakingModuleIds) view returns((uint256,uint256,(uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256),(uint256,uint256,uint256))[] digests)
func (_Api *ApiSession) GetStakingModuleDigests(_stakingModuleIds []*big.Int) ([]StakingRouterStakingModuleDigest, error) {
	return _Api.Contract.GetStakingModuleDigests(&_Api.CallOpts, _stakingModuleIds)
}

// GetStakingModuleDigests is a free data retrieval call binding the contract method 0x8525e3a1.
//
// Solidity: function getStakingModuleDigests(uint256[] _stakingModuleIds) view returns((uint256,uint256,(uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256),(uint256,uint256,uint256))[] digests)
func (_Api *ApiCallerSession) GetStakingModuleDigests(_stakingModuleIds []*big.Int) ([]StakingRouterStakingModuleDigest, error) {
	return _Api.Contract.GetStakingModuleDigests(&_Api.CallOpts, _stakingModuleIds)
}

// GetStakingModuleIds is a free data retrieval call binding the contract method 0xf2aebb65.
//
// Solidity: function getStakingModuleIds() view returns(uint256[] stakingModuleIds)
func (_Api *ApiCaller) GetStakingModuleIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModuleIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetStakingModuleIds is a free data retrieval call binding the contract method 0xf2aebb65.
//
// Solidity: function getStakingModuleIds() view returns(uint256[] stakingModuleIds)
func (_Api *ApiSession) GetStakingModuleIds() ([]*big.Int, error) {
	return _Api.Contract.GetStakingModuleIds(&_Api.CallOpts)
}

// GetStakingModuleIds is a free data retrieval call binding the contract method 0xf2aebb65.
//
// Solidity: function getStakingModuleIds() view returns(uint256[] stakingModuleIds)
func (_Api *ApiCallerSession) GetStakingModuleIds() ([]*big.Int, error) {
	return _Api.Contract.GetStakingModuleIds(&_Api.CallOpts)
}

// GetStakingModuleIsActive is a free data retrieval call binding the contract method 0x6608b11b.
//
// Solidity: function getStakingModuleIsActive(uint256 _stakingModuleId) view returns(bool)
func (_Api *ApiCaller) GetStakingModuleIsActive(opts *bind.CallOpts, _stakingModuleId *big.Int) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModuleIsActive", _stakingModuleId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetStakingModuleIsActive is a free data retrieval call binding the contract method 0x6608b11b.
//
// Solidity: function getStakingModuleIsActive(uint256 _stakingModuleId) view returns(bool)
func (_Api *ApiSession) GetStakingModuleIsActive(_stakingModuleId *big.Int) (bool, error) {
	return _Api.Contract.GetStakingModuleIsActive(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleIsActive is a free data retrieval call binding the contract method 0x6608b11b.
//
// Solidity: function getStakingModuleIsActive(uint256 _stakingModuleId) view returns(bool)
func (_Api *ApiCallerSession) GetStakingModuleIsActive(_stakingModuleId *big.Int) (bool, error) {
	return _Api.Contract.GetStakingModuleIsActive(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleIsDepositsPaused is a free data retrieval call binding the contract method 0xe24ce9f1.
//
// Solidity: function getStakingModuleIsDepositsPaused(uint256 _stakingModuleId) view returns(bool)
func (_Api *ApiCaller) GetStakingModuleIsDepositsPaused(opts *bind.CallOpts, _stakingModuleId *big.Int) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModuleIsDepositsPaused", _stakingModuleId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetStakingModuleIsDepositsPaused is a free data retrieval call binding the contract method 0xe24ce9f1.
//
// Solidity: function getStakingModuleIsDepositsPaused(uint256 _stakingModuleId) view returns(bool)
func (_Api *ApiSession) GetStakingModuleIsDepositsPaused(_stakingModuleId *big.Int) (bool, error) {
	return _Api.Contract.GetStakingModuleIsDepositsPaused(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleIsDepositsPaused is a free data retrieval call binding the contract method 0xe24ce9f1.
//
// Solidity: function getStakingModuleIsDepositsPaused(uint256 _stakingModuleId) view returns(bool)
func (_Api *ApiCallerSession) GetStakingModuleIsDepositsPaused(_stakingModuleId *big.Int) (bool, error) {
	return _Api.Contract.GetStakingModuleIsDepositsPaused(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleIsStopped is a free data retrieval call binding the contract method 0x6ada55b9.
//
// Solidity: function getStakingModuleIsStopped(uint256 _stakingModuleId) view returns(bool)
func (_Api *ApiCaller) GetStakingModuleIsStopped(opts *bind.CallOpts, _stakingModuleId *big.Int) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModuleIsStopped", _stakingModuleId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetStakingModuleIsStopped is a free data retrieval call binding the contract method 0x6ada55b9.
//
// Solidity: function getStakingModuleIsStopped(uint256 _stakingModuleId) view returns(bool)
func (_Api *ApiSession) GetStakingModuleIsStopped(_stakingModuleId *big.Int) (bool, error) {
	return _Api.Contract.GetStakingModuleIsStopped(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleIsStopped is a free data retrieval call binding the contract method 0x6ada55b9.
//
// Solidity: function getStakingModuleIsStopped(uint256 _stakingModuleId) view returns(bool)
func (_Api *ApiCallerSession) GetStakingModuleIsStopped(_stakingModuleId *big.Int) (bool, error) {
	return _Api.Contract.GetStakingModuleIsStopped(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleLastDepositBlock is a free data retrieval call binding the contract method 0x473e0433.
//
// Solidity: function getStakingModuleLastDepositBlock(uint256 _stakingModuleId) view returns(uint256)
func (_Api *ApiCaller) GetStakingModuleLastDepositBlock(opts *bind.CallOpts, _stakingModuleId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModuleLastDepositBlock", _stakingModuleId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingModuleLastDepositBlock is a free data retrieval call binding the contract method 0x473e0433.
//
// Solidity: function getStakingModuleLastDepositBlock(uint256 _stakingModuleId) view returns(uint256)
func (_Api *ApiSession) GetStakingModuleLastDepositBlock(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Api.Contract.GetStakingModuleLastDepositBlock(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleLastDepositBlock is a free data retrieval call binding the contract method 0x473e0433.
//
// Solidity: function getStakingModuleLastDepositBlock(uint256 _stakingModuleId) view returns(uint256)
func (_Api *ApiCallerSession) GetStakingModuleLastDepositBlock(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Api.Contract.GetStakingModuleLastDepositBlock(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleMaxDepositsCount is a free data retrieval call binding the contract method 0x19c64b79.
//
// Solidity: function getStakingModuleMaxDepositsCount(uint256 _stakingModuleId, uint256 _maxDepositsValue) view returns(uint256)
func (_Api *ApiCaller) GetStakingModuleMaxDepositsCount(opts *bind.CallOpts, _stakingModuleId *big.Int, _maxDepositsValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModuleMaxDepositsCount", _stakingModuleId, _maxDepositsValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingModuleMaxDepositsCount is a free data retrieval call binding the contract method 0x19c64b79.
//
// Solidity: function getStakingModuleMaxDepositsCount(uint256 _stakingModuleId, uint256 _maxDepositsValue) view returns(uint256)
func (_Api *ApiSession) GetStakingModuleMaxDepositsCount(_stakingModuleId *big.Int, _maxDepositsValue *big.Int) (*big.Int, error) {
	return _Api.Contract.GetStakingModuleMaxDepositsCount(&_Api.CallOpts, _stakingModuleId, _maxDepositsValue)
}

// GetStakingModuleMaxDepositsCount is a free data retrieval call binding the contract method 0x19c64b79.
//
// Solidity: function getStakingModuleMaxDepositsCount(uint256 _stakingModuleId, uint256 _maxDepositsValue) view returns(uint256)
func (_Api *ApiCallerSession) GetStakingModuleMaxDepositsCount(_stakingModuleId *big.Int, _maxDepositsValue *big.Int) (*big.Int, error) {
	return _Api.Contract.GetStakingModuleMaxDepositsCount(&_Api.CallOpts, _stakingModuleId, _maxDepositsValue)
}

// GetStakingModuleNonce is a free data retrieval call binding the contract method 0x0519fbbf.
//
// Solidity: function getStakingModuleNonce(uint256 _stakingModuleId) view returns(uint256)
func (_Api *ApiCaller) GetStakingModuleNonce(opts *bind.CallOpts, _stakingModuleId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModuleNonce", _stakingModuleId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingModuleNonce is a free data retrieval call binding the contract method 0x0519fbbf.
//
// Solidity: function getStakingModuleNonce(uint256 _stakingModuleId) view returns(uint256)
func (_Api *ApiSession) GetStakingModuleNonce(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Api.Contract.GetStakingModuleNonce(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleNonce is a free data retrieval call binding the contract method 0x0519fbbf.
//
// Solidity: function getStakingModuleNonce(uint256 _stakingModuleId) view returns(uint256)
func (_Api *ApiCallerSession) GetStakingModuleNonce(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Api.Contract.GetStakingModuleNonce(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleStatus is a free data retrieval call binding the contract method 0x9fc5a6ed.
//
// Solidity: function getStakingModuleStatus(uint256 _stakingModuleId) view returns(uint8)
func (_Api *ApiCaller) GetStakingModuleStatus(opts *bind.CallOpts, _stakingModuleId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModuleStatus", _stakingModuleId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStakingModuleStatus is a free data retrieval call binding the contract method 0x9fc5a6ed.
//
// Solidity: function getStakingModuleStatus(uint256 _stakingModuleId) view returns(uint8)
func (_Api *ApiSession) GetStakingModuleStatus(_stakingModuleId *big.Int) (uint8, error) {
	return _Api.Contract.GetStakingModuleStatus(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleStatus is a free data retrieval call binding the contract method 0x9fc5a6ed.
//
// Solidity: function getStakingModuleStatus(uint256 _stakingModuleId) view returns(uint8)
func (_Api *ApiCallerSession) GetStakingModuleStatus(_stakingModuleId *big.Int) (uint8, error) {
	return _Api.Contract.GetStakingModuleStatus(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x07e203ac.
//
// Solidity: function getStakingModuleSummary(uint256 _stakingModuleId) view returns((uint256,uint256,uint256) summary)
func (_Api *ApiCaller) GetStakingModuleSummary(opts *bind.CallOpts, _stakingModuleId *big.Int) (StakingRouterStakingModuleSummary, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModuleSummary", _stakingModuleId)

	if err != nil {
		return *new(StakingRouterStakingModuleSummary), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingRouterStakingModuleSummary)).(*StakingRouterStakingModuleSummary)

	return out0, err

}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x07e203ac.
//
// Solidity: function getStakingModuleSummary(uint256 _stakingModuleId) view returns((uint256,uint256,uint256) summary)
func (_Api *ApiSession) GetStakingModuleSummary(_stakingModuleId *big.Int) (StakingRouterStakingModuleSummary, error) {
	return _Api.Contract.GetStakingModuleSummary(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x07e203ac.
//
// Solidity: function getStakingModuleSummary(uint256 _stakingModuleId) view returns((uint256,uint256,uint256) summary)
func (_Api *ApiCallerSession) GetStakingModuleSummary(_stakingModuleId *big.Int) (StakingRouterStakingModuleSummary, error) {
	return _Api.Contract.GetStakingModuleSummary(&_Api.CallOpts, _stakingModuleId)
}

// GetStakingModules is a free data retrieval call binding the contract method 0x6183214d.
//
// Solidity: function getStakingModules() view returns((uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256)[] res)
func (_Api *ApiCaller) GetStakingModules(opts *bind.CallOpts) ([]StakingRouterStakingModule, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModules")

	if err != nil {
		return *new([]StakingRouterStakingModule), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingRouterStakingModule)).(*[]StakingRouterStakingModule)

	return out0, err

}

// GetStakingModules is a free data retrieval call binding the contract method 0x6183214d.
//
// Solidity: function getStakingModules() view returns((uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256)[] res)
func (_Api *ApiSession) GetStakingModules() ([]StakingRouterStakingModule, error) {
	return _Api.Contract.GetStakingModules(&_Api.CallOpts)
}

// GetStakingModules is a free data retrieval call binding the contract method 0x6183214d.
//
// Solidity: function getStakingModules() view returns((uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256)[] res)
func (_Api *ApiCallerSession) GetStakingModules() ([]StakingRouterStakingModule, error) {
	return _Api.Contract.GetStakingModules(&_Api.CallOpts)
}

// GetStakingModulesCount is a free data retrieval call binding the contract method 0x4a7583b6.
//
// Solidity: function getStakingModulesCount() view returns(uint256)
func (_Api *ApiCaller) GetStakingModulesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingModulesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingModulesCount is a free data retrieval call binding the contract method 0x4a7583b6.
//
// Solidity: function getStakingModulesCount() view returns(uint256)
func (_Api *ApiSession) GetStakingModulesCount() (*big.Int, error) {
	return _Api.Contract.GetStakingModulesCount(&_Api.CallOpts)
}

// GetStakingModulesCount is a free data retrieval call binding the contract method 0x4a7583b6.
//
// Solidity: function getStakingModulesCount() view returns(uint256)
func (_Api *ApiCallerSession) GetStakingModulesCount() (*big.Int, error) {
	return _Api.Contract.GetStakingModulesCount(&_Api.CallOpts)
}

// GetStakingRewardsDistribution is a free data retrieval call binding the contract method 0xba21ccae.
//
// Solidity: function getStakingRewardsDistribution() view returns(address[] recipients, uint256[] stakingModuleIds, uint96[] stakingModuleFees, uint96 totalFee, uint256 precisionPoints)
func (_Api *ApiCaller) GetStakingRewardsDistribution(opts *bind.CallOpts) (struct {
	Recipients        []common.Address
	StakingModuleIds  []*big.Int
	StakingModuleFees []*big.Int
	TotalFee          *big.Int
	PrecisionPoints   *big.Int
}, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getStakingRewardsDistribution")

	outstruct := new(struct {
		Recipients        []common.Address
		StakingModuleIds  []*big.Int
		StakingModuleFees []*big.Int
		TotalFee          *big.Int
		PrecisionPoints   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Recipients = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.StakingModuleIds = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)
	outstruct.StakingModuleFees = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.TotalFee = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.PrecisionPoints = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakingRewardsDistribution is a free data retrieval call binding the contract method 0xba21ccae.
//
// Solidity: function getStakingRewardsDistribution() view returns(address[] recipients, uint256[] stakingModuleIds, uint96[] stakingModuleFees, uint96 totalFee, uint256 precisionPoints)
func (_Api *ApiSession) GetStakingRewardsDistribution() (struct {
	Recipients        []common.Address
	StakingModuleIds  []*big.Int
	StakingModuleFees []*big.Int
	TotalFee          *big.Int
	PrecisionPoints   *big.Int
}, error) {
	return _Api.Contract.GetStakingRewardsDistribution(&_Api.CallOpts)
}

// GetStakingRewardsDistribution is a free data retrieval call binding the contract method 0xba21ccae.
//
// Solidity: function getStakingRewardsDistribution() view returns(address[] recipients, uint256[] stakingModuleIds, uint96[] stakingModuleFees, uint96 totalFee, uint256 precisionPoints)
func (_Api *ApiCallerSession) GetStakingRewardsDistribution() (struct {
	Recipients        []common.Address
	StakingModuleIds  []*big.Int
	StakingModuleFees []*big.Int
	TotalFee          *big.Int
	PrecisionPoints   *big.Int
}, error) {
	return _Api.Contract.GetStakingRewardsDistribution(&_Api.CallOpts)
}

// GetTotalFeeE4Precision is a free data retrieval call binding the contract method 0x9fbb7bae.
//
// Solidity: function getTotalFeeE4Precision() view returns(uint16 totalFee)
func (_Api *ApiCaller) GetTotalFeeE4Precision(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getTotalFeeE4Precision")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetTotalFeeE4Precision is a free data retrieval call binding the contract method 0x9fbb7bae.
//
// Solidity: function getTotalFeeE4Precision() view returns(uint16 totalFee)
func (_Api *ApiSession) GetTotalFeeE4Precision() (uint16, error) {
	return _Api.Contract.GetTotalFeeE4Precision(&_Api.CallOpts)
}

// GetTotalFeeE4Precision is a free data retrieval call binding the contract method 0x9fbb7bae.
//
// Solidity: function getTotalFeeE4Precision() view returns(uint16 totalFee)
func (_Api *ApiCallerSession) GetTotalFeeE4Precision() (uint16, error) {
	return _Api.Contract.GetTotalFeeE4Precision(&_Api.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_Api *ApiCaller) GetWithdrawalCredentials(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "getWithdrawalCredentials")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_Api *ApiSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _Api.Contract.GetWithdrawalCredentials(&_Api.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_Api *ApiCallerSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _Api.Contract.GetWithdrawalCredentials(&_Api.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Api *ApiCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Api *ApiSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Api.Contract.HasRole(&_Api.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Api *ApiCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Api.Contract.HasRole(&_Api.CallOpts, role, account)
}

// HasStakingModule is a free data retrieval call binding the contract method 0xa734329c.
//
// Solidity: function hasStakingModule(uint256 _stakingModuleId) view returns(bool)
func (_Api *ApiCaller) HasStakingModule(opts *bind.CallOpts, _stakingModuleId *big.Int) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "hasStakingModule", _stakingModuleId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasStakingModule is a free data retrieval call binding the contract method 0xa734329c.
//
// Solidity: function hasStakingModule(uint256 _stakingModuleId) view returns(bool)
func (_Api *ApiSession) HasStakingModule(_stakingModuleId *big.Int) (bool, error) {
	return _Api.Contract.HasStakingModule(&_Api.CallOpts, _stakingModuleId)
}

// HasStakingModule is a free data retrieval call binding the contract method 0xa734329c.
//
// Solidity: function hasStakingModule(uint256 _stakingModuleId) view returns(bool)
func (_Api *ApiCallerSession) HasStakingModule(_stakingModuleId *big.Int) (bool, error) {
	return _Api.Contract.HasStakingModule(&_Api.CallOpts, _stakingModuleId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Api *ApiCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Api.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Api *ApiSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Api.Contract.SupportsInterface(&_Api.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Api *ApiCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Api.Contract.SupportsInterface(&_Api.CallOpts, interfaceId)
}

// AddStakingModule is a paid mutator transaction binding the contract method 0x3e54ee5b.
//
// Solidity: function addStakingModule(string _name, address _stakingModuleAddress, uint256 _targetShare, uint256 _stakingModuleFee, uint256 _treasuryFee) returns()
func (_Api *ApiTransactor) AddStakingModule(opts *bind.TransactOpts, _name string, _stakingModuleAddress common.Address, _targetShare *big.Int, _stakingModuleFee *big.Int, _treasuryFee *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "addStakingModule", _name, _stakingModuleAddress, _targetShare, _stakingModuleFee, _treasuryFee)
}

// AddStakingModule is a paid mutator transaction binding the contract method 0x3e54ee5b.
//
// Solidity: function addStakingModule(string _name, address _stakingModuleAddress, uint256 _targetShare, uint256 _stakingModuleFee, uint256 _treasuryFee) returns()
func (_Api *ApiSession) AddStakingModule(_name string, _stakingModuleAddress common.Address, _targetShare *big.Int, _stakingModuleFee *big.Int, _treasuryFee *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddStakingModule(&_Api.TransactOpts, _name, _stakingModuleAddress, _targetShare, _stakingModuleFee, _treasuryFee)
}

// AddStakingModule is a paid mutator transaction binding the contract method 0x3e54ee5b.
//
// Solidity: function addStakingModule(string _name, address _stakingModuleAddress, uint256 _targetShare, uint256 _stakingModuleFee, uint256 _treasuryFee) returns()
func (_Api *ApiTransactorSession) AddStakingModule(_name string, _stakingModuleAddress common.Address, _targetShare *big.Int, _stakingModuleFee *big.Int, _treasuryFee *big.Int) (*types.Transaction, error) {
	return _Api.Contract.AddStakingModule(&_Api.TransactOpts, _name, _stakingModuleAddress, _targetShare, _stakingModuleFee, _treasuryFee)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _depositsCount, uint256 _stakingModuleId, bytes _depositCalldata) payable returns()
func (_Api *ApiTransactor) Deposit(opts *bind.TransactOpts, _depositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "deposit", _depositsCount, _stakingModuleId, _depositCalldata)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _depositsCount, uint256 _stakingModuleId, bytes _depositCalldata) payable returns()
func (_Api *ApiSession) Deposit(_depositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts, _depositsCount, _stakingModuleId, _depositCalldata)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _depositsCount, uint256 _stakingModuleId, bytes _depositCalldata) payable returns()
func (_Api *ApiTransactorSession) Deposit(_depositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _Api.Contract.Deposit(&_Api.TransactOpts, _depositsCount, _stakingModuleId, _depositCalldata)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Api *ApiTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Api *ApiSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.Contract.GrantRole(&_Api.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Api *ApiTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.Contract.GrantRole(&_Api.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x6133f985.
//
// Solidity: function initialize(address _admin, address _lido, bytes32 _withdrawalCredentials) returns()
func (_Api *ApiTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _lido common.Address, _withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "initialize", _admin, _lido, _withdrawalCredentials)
}

// Initialize is a paid mutator transaction binding the contract method 0x6133f985.
//
// Solidity: function initialize(address _admin, address _lido, bytes32 _withdrawalCredentials) returns()
func (_Api *ApiSession) Initialize(_admin common.Address, _lido common.Address, _withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Api.Contract.Initialize(&_Api.TransactOpts, _admin, _lido, _withdrawalCredentials)
}

// Initialize is a paid mutator transaction binding the contract method 0x6133f985.
//
// Solidity: function initialize(address _admin, address _lido, bytes32 _withdrawalCredentials) returns()
func (_Api *ApiTransactorSession) Initialize(_admin common.Address, _lido common.Address, _withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Api.Contract.Initialize(&_Api.TransactOpts, _admin, _lido, _withdrawalCredentials)
}

// OnValidatorsCountsByNodeOperatorReportingFinished is a paid mutator transaction binding the contract method 0xdb3c7ba7.
//
// Solidity: function onValidatorsCountsByNodeOperatorReportingFinished() returns()
func (_Api *ApiTransactor) OnValidatorsCountsByNodeOperatorReportingFinished(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "onValidatorsCountsByNodeOperatorReportingFinished")
}

// OnValidatorsCountsByNodeOperatorReportingFinished is a paid mutator transaction binding the contract method 0xdb3c7ba7.
//
// Solidity: function onValidatorsCountsByNodeOperatorReportingFinished() returns()
func (_Api *ApiSession) OnValidatorsCountsByNodeOperatorReportingFinished() (*types.Transaction, error) {
	return _Api.Contract.OnValidatorsCountsByNodeOperatorReportingFinished(&_Api.TransactOpts)
}

// OnValidatorsCountsByNodeOperatorReportingFinished is a paid mutator transaction binding the contract method 0xdb3c7ba7.
//
// Solidity: function onValidatorsCountsByNodeOperatorReportingFinished() returns()
func (_Api *ApiTransactorSession) OnValidatorsCountsByNodeOperatorReportingFinished() (*types.Transaction, error) {
	return _Api.Contract.OnValidatorsCountsByNodeOperatorReportingFinished(&_Api.TransactOpts)
}

// PauseStakingModule is a paid mutator transaction binding the contract method 0x5bf55e40.
//
// Solidity: function pauseStakingModule(uint256 _stakingModuleId) returns()
func (_Api *ApiTransactor) PauseStakingModule(opts *bind.TransactOpts, _stakingModuleId *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "pauseStakingModule", _stakingModuleId)
}

// PauseStakingModule is a paid mutator transaction binding the contract method 0x5bf55e40.
//
// Solidity: function pauseStakingModule(uint256 _stakingModuleId) returns()
func (_Api *ApiSession) PauseStakingModule(_stakingModuleId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.PauseStakingModule(&_Api.TransactOpts, _stakingModuleId)
}

// PauseStakingModule is a paid mutator transaction binding the contract method 0x5bf55e40.
//
// Solidity: function pauseStakingModule(uint256 _stakingModuleId) returns()
func (_Api *ApiTransactorSession) PauseStakingModule(_stakingModuleId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.PauseStakingModule(&_Api.TransactOpts, _stakingModuleId)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Api *ApiTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Api *ApiSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.Contract.RenounceRole(&_Api.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Api *ApiTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.Contract.RenounceRole(&_Api.TransactOpts, role, account)
}

// ReportRewardsMinted is a paid mutator transaction binding the contract method 0xaf124097.
//
// Solidity: function reportRewardsMinted(uint256[] _stakingModuleIds, uint256[] _totalShares) returns()
func (_Api *ApiTransactor) ReportRewardsMinted(opts *bind.TransactOpts, _stakingModuleIds []*big.Int, _totalShares []*big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "reportRewardsMinted", _stakingModuleIds, _totalShares)
}

// ReportRewardsMinted is a paid mutator transaction binding the contract method 0xaf124097.
//
// Solidity: function reportRewardsMinted(uint256[] _stakingModuleIds, uint256[] _totalShares) returns()
func (_Api *ApiSession) ReportRewardsMinted(_stakingModuleIds []*big.Int, _totalShares []*big.Int) (*types.Transaction, error) {
	return _Api.Contract.ReportRewardsMinted(&_Api.TransactOpts, _stakingModuleIds, _totalShares)
}

// ReportRewardsMinted is a paid mutator transaction binding the contract method 0xaf124097.
//
// Solidity: function reportRewardsMinted(uint256[] _stakingModuleIds, uint256[] _totalShares) returns()
func (_Api *ApiTransactorSession) ReportRewardsMinted(_stakingModuleIds []*big.Int, _totalShares []*big.Int) (*types.Transaction, error) {
	return _Api.Contract.ReportRewardsMinted(&_Api.TransactOpts, _stakingModuleIds, _totalShares)
}

// ReportStakingModuleExitedValidatorsCountByNodeOperator is a paid mutator transaction binding the contract method 0xc8ac4980.
//
// Solidity: function reportStakingModuleExitedValidatorsCountByNodeOperator(uint256 _stakingModuleId, bytes _nodeOperatorIds, bytes _exitedValidatorsCounts) returns()
func (_Api *ApiTransactor) ReportStakingModuleExitedValidatorsCountByNodeOperator(opts *bind.TransactOpts, _stakingModuleId *big.Int, _nodeOperatorIds []byte, _exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "reportStakingModuleExitedValidatorsCountByNodeOperator", _stakingModuleId, _nodeOperatorIds, _exitedValidatorsCounts)
}

// ReportStakingModuleExitedValidatorsCountByNodeOperator is a paid mutator transaction binding the contract method 0xc8ac4980.
//
// Solidity: function reportStakingModuleExitedValidatorsCountByNodeOperator(uint256 _stakingModuleId, bytes _nodeOperatorIds, bytes _exitedValidatorsCounts) returns()
func (_Api *ApiSession) ReportStakingModuleExitedValidatorsCountByNodeOperator(_stakingModuleId *big.Int, _nodeOperatorIds []byte, _exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _Api.Contract.ReportStakingModuleExitedValidatorsCountByNodeOperator(&_Api.TransactOpts, _stakingModuleId, _nodeOperatorIds, _exitedValidatorsCounts)
}

// ReportStakingModuleExitedValidatorsCountByNodeOperator is a paid mutator transaction binding the contract method 0xc8ac4980.
//
// Solidity: function reportStakingModuleExitedValidatorsCountByNodeOperator(uint256 _stakingModuleId, bytes _nodeOperatorIds, bytes _exitedValidatorsCounts) returns()
func (_Api *ApiTransactorSession) ReportStakingModuleExitedValidatorsCountByNodeOperator(_stakingModuleId *big.Int, _nodeOperatorIds []byte, _exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _Api.Contract.ReportStakingModuleExitedValidatorsCountByNodeOperator(&_Api.TransactOpts, _stakingModuleId, _nodeOperatorIds, _exitedValidatorsCounts)
}

// ReportStakingModuleStuckValidatorsCountByNodeOperator is a paid mutator transaction binding the contract method 0xcb589b9a.
//
// Solidity: function reportStakingModuleStuckValidatorsCountByNodeOperator(uint256 _stakingModuleId, bytes _nodeOperatorIds, bytes _stuckValidatorsCounts) returns()
func (_Api *ApiTransactor) ReportStakingModuleStuckValidatorsCountByNodeOperator(opts *bind.TransactOpts, _stakingModuleId *big.Int, _nodeOperatorIds []byte, _stuckValidatorsCounts []byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "reportStakingModuleStuckValidatorsCountByNodeOperator", _stakingModuleId, _nodeOperatorIds, _stuckValidatorsCounts)
}

// ReportStakingModuleStuckValidatorsCountByNodeOperator is a paid mutator transaction binding the contract method 0xcb589b9a.
//
// Solidity: function reportStakingModuleStuckValidatorsCountByNodeOperator(uint256 _stakingModuleId, bytes _nodeOperatorIds, bytes _stuckValidatorsCounts) returns()
func (_Api *ApiSession) ReportStakingModuleStuckValidatorsCountByNodeOperator(_stakingModuleId *big.Int, _nodeOperatorIds []byte, _stuckValidatorsCounts []byte) (*types.Transaction, error) {
	return _Api.Contract.ReportStakingModuleStuckValidatorsCountByNodeOperator(&_Api.TransactOpts, _stakingModuleId, _nodeOperatorIds, _stuckValidatorsCounts)
}

// ReportStakingModuleStuckValidatorsCountByNodeOperator is a paid mutator transaction binding the contract method 0xcb589b9a.
//
// Solidity: function reportStakingModuleStuckValidatorsCountByNodeOperator(uint256 _stakingModuleId, bytes _nodeOperatorIds, bytes _stuckValidatorsCounts) returns()
func (_Api *ApiTransactorSession) ReportStakingModuleStuckValidatorsCountByNodeOperator(_stakingModuleId *big.Int, _nodeOperatorIds []byte, _stuckValidatorsCounts []byte) (*types.Transaction, error) {
	return _Api.Contract.ReportStakingModuleStuckValidatorsCountByNodeOperator(&_Api.TransactOpts, _stakingModuleId, _nodeOperatorIds, _stuckValidatorsCounts)
}

// ResumeStakingModule is a paid mutator transaction binding the contract method 0xd861c584.
//
// Solidity: function resumeStakingModule(uint256 _stakingModuleId) returns()
func (_Api *ApiTransactor) ResumeStakingModule(opts *bind.TransactOpts, _stakingModuleId *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "resumeStakingModule", _stakingModuleId)
}

// ResumeStakingModule is a paid mutator transaction binding the contract method 0xd861c584.
//
// Solidity: function resumeStakingModule(uint256 _stakingModuleId) returns()
func (_Api *ApiSession) ResumeStakingModule(_stakingModuleId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ResumeStakingModule(&_Api.TransactOpts, _stakingModuleId)
}

// ResumeStakingModule is a paid mutator transaction binding the contract method 0xd861c584.
//
// Solidity: function resumeStakingModule(uint256 _stakingModuleId) returns()
func (_Api *ApiTransactorSession) ResumeStakingModule(_stakingModuleId *big.Int) (*types.Transaction, error) {
	return _Api.Contract.ResumeStakingModule(&_Api.TransactOpts, _stakingModuleId)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Api *ApiTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Api *ApiSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.Contract.RevokeRole(&_Api.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Api *ApiTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Api.Contract.RevokeRole(&_Api.TransactOpts, role, account)
}

// SetStakingModuleStatus is a paid mutator transaction binding the contract method 0xd0a2b1b8.
//
// Solidity: function setStakingModuleStatus(uint256 _stakingModuleId, uint8 _status) returns()
func (_Api *ApiTransactor) SetStakingModuleStatus(opts *bind.TransactOpts, _stakingModuleId *big.Int, _status uint8) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setStakingModuleStatus", _stakingModuleId, _status)
}

// SetStakingModuleStatus is a paid mutator transaction binding the contract method 0xd0a2b1b8.
//
// Solidity: function setStakingModuleStatus(uint256 _stakingModuleId, uint8 _status) returns()
func (_Api *ApiSession) SetStakingModuleStatus(_stakingModuleId *big.Int, _status uint8) (*types.Transaction, error) {
	return _Api.Contract.SetStakingModuleStatus(&_Api.TransactOpts, _stakingModuleId, _status)
}

// SetStakingModuleStatus is a paid mutator transaction binding the contract method 0xd0a2b1b8.
//
// Solidity: function setStakingModuleStatus(uint256 _stakingModuleId, uint8 _status) returns()
func (_Api *ApiTransactorSession) SetStakingModuleStatus(_stakingModuleId *big.Int, _status uint8) (*types.Transaction, error) {
	return _Api.Contract.SetStakingModuleStatus(&_Api.TransactOpts, _stakingModuleId, _status)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xe97ee8cc.
//
// Solidity: function setWithdrawalCredentials(bytes32 _withdrawalCredentials) returns()
func (_Api *ApiTransactor) SetWithdrawalCredentials(opts *bind.TransactOpts, _withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "setWithdrawalCredentials", _withdrawalCredentials)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xe97ee8cc.
//
// Solidity: function setWithdrawalCredentials(bytes32 _withdrawalCredentials) returns()
func (_Api *ApiSession) SetWithdrawalCredentials(_withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Api.Contract.SetWithdrawalCredentials(&_Api.TransactOpts, _withdrawalCredentials)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xe97ee8cc.
//
// Solidity: function setWithdrawalCredentials(bytes32 _withdrawalCredentials) returns()
func (_Api *ApiTransactorSession) SetWithdrawalCredentials(_withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Api.Contract.SetWithdrawalCredentials(&_Api.TransactOpts, _withdrawalCredentials)
}

// UnsafeSetExitedValidatorsCount is a paid mutator transaction binding the contract method 0x072859c7.
//
// Solidity: function unsafeSetExitedValidatorsCount(uint256 _stakingModuleId, uint256 _nodeOperatorId, bool _triggerUpdateFinish, (uint256,uint256,uint256,uint256,uint256,uint256) _correction) returns()
func (_Api *ApiTransactor) UnsafeSetExitedValidatorsCount(opts *bind.TransactOpts, _stakingModuleId *big.Int, _nodeOperatorId *big.Int, _triggerUpdateFinish bool, _correction StakingRouterValidatorsCountsCorrection) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "unsafeSetExitedValidatorsCount", _stakingModuleId, _nodeOperatorId, _triggerUpdateFinish, _correction)
}

// UnsafeSetExitedValidatorsCount is a paid mutator transaction binding the contract method 0x072859c7.
//
// Solidity: function unsafeSetExitedValidatorsCount(uint256 _stakingModuleId, uint256 _nodeOperatorId, bool _triggerUpdateFinish, (uint256,uint256,uint256,uint256,uint256,uint256) _correction) returns()
func (_Api *ApiSession) UnsafeSetExitedValidatorsCount(_stakingModuleId *big.Int, _nodeOperatorId *big.Int, _triggerUpdateFinish bool, _correction StakingRouterValidatorsCountsCorrection) (*types.Transaction, error) {
	return _Api.Contract.UnsafeSetExitedValidatorsCount(&_Api.TransactOpts, _stakingModuleId, _nodeOperatorId, _triggerUpdateFinish, _correction)
}

// UnsafeSetExitedValidatorsCount is a paid mutator transaction binding the contract method 0x072859c7.
//
// Solidity: function unsafeSetExitedValidatorsCount(uint256 _stakingModuleId, uint256 _nodeOperatorId, bool _triggerUpdateFinish, (uint256,uint256,uint256,uint256,uint256,uint256) _correction) returns()
func (_Api *ApiTransactorSession) UnsafeSetExitedValidatorsCount(_stakingModuleId *big.Int, _nodeOperatorId *big.Int, _triggerUpdateFinish bool, _correction StakingRouterValidatorsCountsCorrection) (*types.Transaction, error) {
	return _Api.Contract.UnsafeSetExitedValidatorsCount(&_Api.TransactOpts, _stakingModuleId, _nodeOperatorId, _triggerUpdateFinish, _correction)
}

// UpdateExitedValidatorsCountByStakingModule is a paid mutator transaction binding the contract method 0xabd44a24.
//
// Solidity: function updateExitedValidatorsCountByStakingModule(uint256[] _stakingModuleIds, uint256[] _exitedValidatorsCounts) returns(uint256)
func (_Api *ApiTransactor) UpdateExitedValidatorsCountByStakingModule(opts *bind.TransactOpts, _stakingModuleIds []*big.Int, _exitedValidatorsCounts []*big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "updateExitedValidatorsCountByStakingModule", _stakingModuleIds, _exitedValidatorsCounts)
}

// UpdateExitedValidatorsCountByStakingModule is a paid mutator transaction binding the contract method 0xabd44a24.
//
// Solidity: function updateExitedValidatorsCountByStakingModule(uint256[] _stakingModuleIds, uint256[] _exitedValidatorsCounts) returns(uint256)
func (_Api *ApiSession) UpdateExitedValidatorsCountByStakingModule(_stakingModuleIds []*big.Int, _exitedValidatorsCounts []*big.Int) (*types.Transaction, error) {
	return _Api.Contract.UpdateExitedValidatorsCountByStakingModule(&_Api.TransactOpts, _stakingModuleIds, _exitedValidatorsCounts)
}

// UpdateExitedValidatorsCountByStakingModule is a paid mutator transaction binding the contract method 0xabd44a24.
//
// Solidity: function updateExitedValidatorsCountByStakingModule(uint256[] _stakingModuleIds, uint256[] _exitedValidatorsCounts) returns(uint256)
func (_Api *ApiTransactorSession) UpdateExitedValidatorsCountByStakingModule(_stakingModuleIds []*big.Int, _exitedValidatorsCounts []*big.Int) (*types.Transaction, error) {
	return _Api.Contract.UpdateExitedValidatorsCountByStakingModule(&_Api.TransactOpts, _stakingModuleIds, _exitedValidatorsCounts)
}

// UpdateRefundedValidatorsCount is a paid mutator transaction binding the contract method 0xe1b92a5c.
//
// Solidity: function updateRefundedValidatorsCount(uint256 _stakingModuleId, uint256 _nodeOperatorId, uint256 _refundedValidatorsCount) returns()
func (_Api *ApiTransactor) UpdateRefundedValidatorsCount(opts *bind.TransactOpts, _stakingModuleId *big.Int, _nodeOperatorId *big.Int, _refundedValidatorsCount *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "updateRefundedValidatorsCount", _stakingModuleId, _nodeOperatorId, _refundedValidatorsCount)
}

// UpdateRefundedValidatorsCount is a paid mutator transaction binding the contract method 0xe1b92a5c.
//
// Solidity: function updateRefundedValidatorsCount(uint256 _stakingModuleId, uint256 _nodeOperatorId, uint256 _refundedValidatorsCount) returns()
func (_Api *ApiSession) UpdateRefundedValidatorsCount(_stakingModuleId *big.Int, _nodeOperatorId *big.Int, _refundedValidatorsCount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.UpdateRefundedValidatorsCount(&_Api.TransactOpts, _stakingModuleId, _nodeOperatorId, _refundedValidatorsCount)
}

// UpdateRefundedValidatorsCount is a paid mutator transaction binding the contract method 0xe1b92a5c.
//
// Solidity: function updateRefundedValidatorsCount(uint256 _stakingModuleId, uint256 _nodeOperatorId, uint256 _refundedValidatorsCount) returns()
func (_Api *ApiTransactorSession) UpdateRefundedValidatorsCount(_stakingModuleId *big.Int, _nodeOperatorId *big.Int, _refundedValidatorsCount *big.Int) (*types.Transaction, error) {
	return _Api.Contract.UpdateRefundedValidatorsCount(&_Api.TransactOpts, _stakingModuleId, _nodeOperatorId, _refundedValidatorsCount)
}

// UpdateStakingModule is a paid mutator transaction binding the contract method 0x8dc70c57.
//
// Solidity: function updateStakingModule(uint256 _stakingModuleId, uint256 _targetShare, uint256 _stakingModuleFee, uint256 _treasuryFee) returns()
func (_Api *ApiTransactor) UpdateStakingModule(opts *bind.TransactOpts, _stakingModuleId *big.Int, _targetShare *big.Int, _stakingModuleFee *big.Int, _treasuryFee *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "updateStakingModule", _stakingModuleId, _targetShare, _stakingModuleFee, _treasuryFee)
}

// UpdateStakingModule is a paid mutator transaction binding the contract method 0x8dc70c57.
//
// Solidity: function updateStakingModule(uint256 _stakingModuleId, uint256 _targetShare, uint256 _stakingModuleFee, uint256 _treasuryFee) returns()
func (_Api *ApiSession) UpdateStakingModule(_stakingModuleId *big.Int, _targetShare *big.Int, _stakingModuleFee *big.Int, _treasuryFee *big.Int) (*types.Transaction, error) {
	return _Api.Contract.UpdateStakingModule(&_Api.TransactOpts, _stakingModuleId, _targetShare, _stakingModuleFee, _treasuryFee)
}

// UpdateStakingModule is a paid mutator transaction binding the contract method 0x8dc70c57.
//
// Solidity: function updateStakingModule(uint256 _stakingModuleId, uint256 _targetShare, uint256 _stakingModuleFee, uint256 _treasuryFee) returns()
func (_Api *ApiTransactorSession) UpdateStakingModule(_stakingModuleId *big.Int, _targetShare *big.Int, _stakingModuleFee *big.Int, _treasuryFee *big.Int) (*types.Transaction, error) {
	return _Api.Contract.UpdateStakingModule(&_Api.TransactOpts, _stakingModuleId, _targetShare, _stakingModuleFee, _treasuryFee)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x7443f523.
//
// Solidity: function updateTargetValidatorsLimits(uint256 _stakingModuleId, uint256 _nodeOperatorId, bool _isTargetLimitActive, uint256 _targetLimit) returns()
func (_Api *ApiTransactor) UpdateTargetValidatorsLimits(opts *bind.TransactOpts, _stakingModuleId *big.Int, _nodeOperatorId *big.Int, _isTargetLimitActive bool, _targetLimit *big.Int) (*types.Transaction, error) {
	return _Api.contract.Transact(opts, "updateTargetValidatorsLimits", _stakingModuleId, _nodeOperatorId, _isTargetLimitActive, _targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x7443f523.
//
// Solidity: function updateTargetValidatorsLimits(uint256 _stakingModuleId, uint256 _nodeOperatorId, bool _isTargetLimitActive, uint256 _targetLimit) returns()
func (_Api *ApiSession) UpdateTargetValidatorsLimits(_stakingModuleId *big.Int, _nodeOperatorId *big.Int, _isTargetLimitActive bool, _targetLimit *big.Int) (*types.Transaction, error) {
	return _Api.Contract.UpdateTargetValidatorsLimits(&_Api.TransactOpts, _stakingModuleId, _nodeOperatorId, _isTargetLimitActive, _targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x7443f523.
//
// Solidity: function updateTargetValidatorsLimits(uint256 _stakingModuleId, uint256 _nodeOperatorId, bool _isTargetLimitActive, uint256 _targetLimit) returns()
func (_Api *ApiTransactorSession) UpdateTargetValidatorsLimits(_stakingModuleId *big.Int, _nodeOperatorId *big.Int, _isTargetLimitActive bool, _targetLimit *big.Int) (*types.Transaction, error) {
	return _Api.Contract.UpdateTargetValidatorsLimits(&_Api.TransactOpts, _stakingModuleId, _nodeOperatorId, _isTargetLimitActive, _targetLimit)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Api.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiSession) Receive() (*types.Transaction, error) {
	return _Api.Contract.Receive(&_Api.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Api *ApiTransactorSession) Receive() (*types.Transaction, error) {
	return _Api.Contract.Receive(&_Api.TransactOpts)
}

// ApiContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the Api contract.
type ApiContractVersionSetIterator struct {
	Event *ApiContractVersionSet // Event containing the contract specifics and raw log

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
func (it *ApiContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiContractVersionSet)
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
		it.Event = new(ApiContractVersionSet)
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
func (it *ApiContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiContractVersionSet represents a ContractVersionSet event raised by the Api contract.
type ApiContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Api *ApiFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*ApiContractVersionSetIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &ApiContractVersionSetIterator{contract: _Api.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Api *ApiFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *ApiContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiContractVersionSet)
				if err := _Api.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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

// ParseContractVersionSet is a log parse operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Api *ApiFilterer) ParseContractVersionSet(log types.Log) (*ApiContractVersionSet, error) {
	event := new(ApiContractVersionSet)
	if err := _Api.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiExitedAndStuckValidatorsCountsUpdateFailedIterator is returned from FilterExitedAndStuckValidatorsCountsUpdateFailed and is used to iterate over the raw logs and unpacked data for ExitedAndStuckValidatorsCountsUpdateFailed events raised by the Api contract.
type ApiExitedAndStuckValidatorsCountsUpdateFailedIterator struct {
	Event *ApiExitedAndStuckValidatorsCountsUpdateFailed // Event containing the contract specifics and raw log

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
func (it *ApiExitedAndStuckValidatorsCountsUpdateFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiExitedAndStuckValidatorsCountsUpdateFailed)
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
		it.Event = new(ApiExitedAndStuckValidatorsCountsUpdateFailed)
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
func (it *ApiExitedAndStuckValidatorsCountsUpdateFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiExitedAndStuckValidatorsCountsUpdateFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiExitedAndStuckValidatorsCountsUpdateFailed represents a ExitedAndStuckValidatorsCountsUpdateFailed event raised by the Api contract.
type ApiExitedAndStuckValidatorsCountsUpdateFailed struct {
	StakingModuleId    *big.Int
	LowLevelRevertData []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExitedAndStuckValidatorsCountsUpdateFailed is a free log retrieval operation binding the contract event 0xe74bf895f0c3a2d6c74c40cbb362fdd9640035fc4226c72e3843809ad2a9d2b5.
//
// Solidity: event ExitedAndStuckValidatorsCountsUpdateFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Api *ApiFilterer) FilterExitedAndStuckValidatorsCountsUpdateFailed(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*ApiExitedAndStuckValidatorsCountsUpdateFailedIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "ExitedAndStuckValidatorsCountsUpdateFailed", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiExitedAndStuckValidatorsCountsUpdateFailedIterator{contract: _Api.contract, event: "ExitedAndStuckValidatorsCountsUpdateFailed", logs: logs, sub: sub}, nil
}

// WatchExitedAndStuckValidatorsCountsUpdateFailed is a free log subscription operation binding the contract event 0xe74bf895f0c3a2d6c74c40cbb362fdd9640035fc4226c72e3843809ad2a9d2b5.
//
// Solidity: event ExitedAndStuckValidatorsCountsUpdateFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Api *ApiFilterer) WatchExitedAndStuckValidatorsCountsUpdateFailed(opts *bind.WatchOpts, sink chan<- *ApiExitedAndStuckValidatorsCountsUpdateFailed, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "ExitedAndStuckValidatorsCountsUpdateFailed", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiExitedAndStuckValidatorsCountsUpdateFailed)
				if err := _Api.contract.UnpackLog(event, "ExitedAndStuckValidatorsCountsUpdateFailed", log); err != nil {
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

// ParseExitedAndStuckValidatorsCountsUpdateFailed is a log parse operation binding the contract event 0xe74bf895f0c3a2d6c74c40cbb362fdd9640035fc4226c72e3843809ad2a9d2b5.
//
// Solidity: event ExitedAndStuckValidatorsCountsUpdateFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Api *ApiFilterer) ParseExitedAndStuckValidatorsCountsUpdateFailed(log types.Log) (*ApiExitedAndStuckValidatorsCountsUpdateFailed, error) {
	event := new(ApiExitedAndStuckValidatorsCountsUpdateFailed)
	if err := _Api.contract.UnpackLog(event, "ExitedAndStuckValidatorsCountsUpdateFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiRewardsMintedReportFailedIterator is returned from FilterRewardsMintedReportFailed and is used to iterate over the raw logs and unpacked data for RewardsMintedReportFailed events raised by the Api contract.
type ApiRewardsMintedReportFailedIterator struct {
	Event *ApiRewardsMintedReportFailed // Event containing the contract specifics and raw log

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
func (it *ApiRewardsMintedReportFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiRewardsMintedReportFailed)
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
		it.Event = new(ApiRewardsMintedReportFailed)
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
func (it *ApiRewardsMintedReportFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiRewardsMintedReportFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiRewardsMintedReportFailed represents a RewardsMintedReportFailed event raised by the Api contract.
type ApiRewardsMintedReportFailed struct {
	StakingModuleId    *big.Int
	LowLevelRevertData []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardsMintedReportFailed is a free log retrieval operation binding the contract event 0xf74208fedac7280fd11f8de0be14e00423dc5076da8e8ec8ca90e09257fff1b3.
//
// Solidity: event RewardsMintedReportFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Api *ApiFilterer) FilterRewardsMintedReportFailed(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*ApiRewardsMintedReportFailedIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "RewardsMintedReportFailed", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiRewardsMintedReportFailedIterator{contract: _Api.contract, event: "RewardsMintedReportFailed", logs: logs, sub: sub}, nil
}

// WatchRewardsMintedReportFailed is a free log subscription operation binding the contract event 0xf74208fedac7280fd11f8de0be14e00423dc5076da8e8ec8ca90e09257fff1b3.
//
// Solidity: event RewardsMintedReportFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Api *ApiFilterer) WatchRewardsMintedReportFailed(opts *bind.WatchOpts, sink chan<- *ApiRewardsMintedReportFailed, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "RewardsMintedReportFailed", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiRewardsMintedReportFailed)
				if err := _Api.contract.UnpackLog(event, "RewardsMintedReportFailed", log); err != nil {
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

// ParseRewardsMintedReportFailed is a log parse operation binding the contract event 0xf74208fedac7280fd11f8de0be14e00423dc5076da8e8ec8ca90e09257fff1b3.
//
// Solidity: event RewardsMintedReportFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Api *ApiFilterer) ParseRewardsMintedReportFailed(log types.Log) (*ApiRewardsMintedReportFailed, error) {
	event := new(ApiRewardsMintedReportFailed)
	if err := _Api.contract.UnpackLog(event, "RewardsMintedReportFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Api contract.
type ApiRoleAdminChangedIterator struct {
	Event *ApiRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *ApiRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiRoleAdminChanged)
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
		it.Event = new(ApiRoleAdminChanged)
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
func (it *ApiRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiRoleAdminChanged represents a RoleAdminChanged event raised by the Api contract.
type ApiRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Api *ApiFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*ApiRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Api.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &ApiRoleAdminChangedIterator{contract: _Api.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Api *ApiFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *ApiRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Api.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiRoleAdminChanged)
				if err := _Api.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Api *ApiFilterer) ParseRoleAdminChanged(log types.Log) (*ApiRoleAdminChanged, error) {
	event := new(ApiRoleAdminChanged)
	if err := _Api.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Api contract.
type ApiRoleGrantedIterator struct {
	Event *ApiRoleGranted // Event containing the contract specifics and raw log

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
func (it *ApiRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiRoleGranted)
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
		it.Event = new(ApiRoleGranted)
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
func (it *ApiRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiRoleGranted represents a RoleGranted event raised by the Api contract.
type ApiRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Api *ApiFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ApiRoleGrantedIterator, error) {

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

	logs, sub, err := _Api.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ApiRoleGrantedIterator{contract: _Api.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Api *ApiFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *ApiRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Api.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiRoleGranted)
				if err := _Api.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Api *ApiFilterer) ParseRoleGranted(log types.Log) (*ApiRoleGranted, error) {
	event := new(ApiRoleGranted)
	if err := _Api.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Api contract.
type ApiRoleRevokedIterator struct {
	Event *ApiRoleRevoked // Event containing the contract specifics and raw log

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
func (it *ApiRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiRoleRevoked)
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
		it.Event = new(ApiRoleRevoked)
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
func (it *ApiRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiRoleRevoked represents a RoleRevoked event raised by the Api contract.
type ApiRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Api *ApiFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*ApiRoleRevokedIterator, error) {

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

	logs, sub, err := _Api.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ApiRoleRevokedIterator{contract: _Api.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Api *ApiFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *ApiRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Api.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiRoleRevoked)
				if err := _Api.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Api *ApiFilterer) ParseRoleRevoked(log types.Log) (*ApiRoleRevoked, error) {
	event := new(ApiRoleRevoked)
	if err := _Api.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiStakingModuleAddedIterator is returned from FilterStakingModuleAdded and is used to iterate over the raw logs and unpacked data for StakingModuleAdded events raised by the Api contract.
type ApiStakingModuleAddedIterator struct {
	Event *ApiStakingModuleAdded // Event containing the contract specifics and raw log

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
func (it *ApiStakingModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiStakingModuleAdded)
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
		it.Event = new(ApiStakingModuleAdded)
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
func (it *ApiStakingModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiStakingModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiStakingModuleAdded represents a StakingModuleAdded event raised by the Api contract.
type ApiStakingModuleAdded struct {
	StakingModuleId *big.Int
	StakingModule   common.Address
	Name            string
	CreatedBy       common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleAdded is a free log retrieval operation binding the contract event 0x43b5213f0e1666cd0b8692a73686164c94deb955a59c65e10dee8bb958e7ce3e.
//
// Solidity: event StakingModuleAdded(uint256 indexed stakingModuleId, address stakingModule, string name, address createdBy)
func (_Api *ApiFilterer) FilterStakingModuleAdded(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*ApiStakingModuleAddedIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "StakingModuleAdded", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiStakingModuleAddedIterator{contract: _Api.contract, event: "StakingModuleAdded", logs: logs, sub: sub}, nil
}

// WatchStakingModuleAdded is a free log subscription operation binding the contract event 0x43b5213f0e1666cd0b8692a73686164c94deb955a59c65e10dee8bb958e7ce3e.
//
// Solidity: event StakingModuleAdded(uint256 indexed stakingModuleId, address stakingModule, string name, address createdBy)
func (_Api *ApiFilterer) WatchStakingModuleAdded(opts *bind.WatchOpts, sink chan<- *ApiStakingModuleAdded, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "StakingModuleAdded", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiStakingModuleAdded)
				if err := _Api.contract.UnpackLog(event, "StakingModuleAdded", log); err != nil {
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

// ParseStakingModuleAdded is a log parse operation binding the contract event 0x43b5213f0e1666cd0b8692a73686164c94deb955a59c65e10dee8bb958e7ce3e.
//
// Solidity: event StakingModuleAdded(uint256 indexed stakingModuleId, address stakingModule, string name, address createdBy)
func (_Api *ApiFilterer) ParseStakingModuleAdded(log types.Log) (*ApiStakingModuleAdded, error) {
	event := new(ApiStakingModuleAdded)
	if err := _Api.contract.UnpackLog(event, "StakingModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiStakingModuleExitedValidatorsIncompleteReportingIterator is returned from FilterStakingModuleExitedValidatorsIncompleteReporting and is used to iterate over the raw logs and unpacked data for StakingModuleExitedValidatorsIncompleteReporting events raised by the Api contract.
type ApiStakingModuleExitedValidatorsIncompleteReportingIterator struct {
	Event *ApiStakingModuleExitedValidatorsIncompleteReporting // Event containing the contract specifics and raw log

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
func (it *ApiStakingModuleExitedValidatorsIncompleteReportingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiStakingModuleExitedValidatorsIncompleteReporting)
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
		it.Event = new(ApiStakingModuleExitedValidatorsIncompleteReporting)
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
func (it *ApiStakingModuleExitedValidatorsIncompleteReportingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiStakingModuleExitedValidatorsIncompleteReportingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiStakingModuleExitedValidatorsIncompleteReporting represents a StakingModuleExitedValidatorsIncompleteReporting event raised by the Api contract.
type ApiStakingModuleExitedValidatorsIncompleteReporting struct {
	StakingModuleId                 *big.Int
	UnreportedExitedValidatorsCount *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleExitedValidatorsIncompleteReporting is a free log retrieval operation binding the contract event 0xdd2523ca96a639ba7e17420698937f71eddd8af012ccb36ff5c8fe96141acae9.
//
// Solidity: event StakingModuleExitedValidatorsIncompleteReporting(uint256 indexed stakingModuleId, uint256 unreportedExitedValidatorsCount)
func (_Api *ApiFilterer) FilterStakingModuleExitedValidatorsIncompleteReporting(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*ApiStakingModuleExitedValidatorsIncompleteReportingIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "StakingModuleExitedValidatorsIncompleteReporting", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiStakingModuleExitedValidatorsIncompleteReportingIterator{contract: _Api.contract, event: "StakingModuleExitedValidatorsIncompleteReporting", logs: logs, sub: sub}, nil
}

// WatchStakingModuleExitedValidatorsIncompleteReporting is a free log subscription operation binding the contract event 0xdd2523ca96a639ba7e17420698937f71eddd8af012ccb36ff5c8fe96141acae9.
//
// Solidity: event StakingModuleExitedValidatorsIncompleteReporting(uint256 indexed stakingModuleId, uint256 unreportedExitedValidatorsCount)
func (_Api *ApiFilterer) WatchStakingModuleExitedValidatorsIncompleteReporting(opts *bind.WatchOpts, sink chan<- *ApiStakingModuleExitedValidatorsIncompleteReporting, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "StakingModuleExitedValidatorsIncompleteReporting", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiStakingModuleExitedValidatorsIncompleteReporting)
				if err := _Api.contract.UnpackLog(event, "StakingModuleExitedValidatorsIncompleteReporting", log); err != nil {
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

// ParseStakingModuleExitedValidatorsIncompleteReporting is a log parse operation binding the contract event 0xdd2523ca96a639ba7e17420698937f71eddd8af012ccb36ff5c8fe96141acae9.
//
// Solidity: event StakingModuleExitedValidatorsIncompleteReporting(uint256 indexed stakingModuleId, uint256 unreportedExitedValidatorsCount)
func (_Api *ApiFilterer) ParseStakingModuleExitedValidatorsIncompleteReporting(log types.Log) (*ApiStakingModuleExitedValidatorsIncompleteReporting, error) {
	event := new(ApiStakingModuleExitedValidatorsIncompleteReporting)
	if err := _Api.contract.UnpackLog(event, "StakingModuleExitedValidatorsIncompleteReporting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiStakingModuleFeesSetIterator is returned from FilterStakingModuleFeesSet and is used to iterate over the raw logs and unpacked data for StakingModuleFeesSet events raised by the Api contract.
type ApiStakingModuleFeesSetIterator struct {
	Event *ApiStakingModuleFeesSet // Event containing the contract specifics and raw log

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
func (it *ApiStakingModuleFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiStakingModuleFeesSet)
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
		it.Event = new(ApiStakingModuleFeesSet)
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
func (it *ApiStakingModuleFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiStakingModuleFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiStakingModuleFeesSet represents a StakingModuleFeesSet event raised by the Api contract.
type ApiStakingModuleFeesSet struct {
	StakingModuleId  *big.Int
	StakingModuleFee *big.Int
	TreasuryFee      *big.Int
	SetBy            common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleFeesSet is a free log retrieval operation binding the contract event 0x303c8ac43d1b1f9b898ddd2915a294efa01e9b07c322d7deeb7db332b66f0410.
//
// Solidity: event StakingModuleFeesSet(uint256 indexed stakingModuleId, uint256 stakingModuleFee, uint256 treasuryFee, address setBy)
func (_Api *ApiFilterer) FilterStakingModuleFeesSet(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*ApiStakingModuleFeesSetIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "StakingModuleFeesSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiStakingModuleFeesSetIterator{contract: _Api.contract, event: "StakingModuleFeesSet", logs: logs, sub: sub}, nil
}

// WatchStakingModuleFeesSet is a free log subscription operation binding the contract event 0x303c8ac43d1b1f9b898ddd2915a294efa01e9b07c322d7deeb7db332b66f0410.
//
// Solidity: event StakingModuleFeesSet(uint256 indexed stakingModuleId, uint256 stakingModuleFee, uint256 treasuryFee, address setBy)
func (_Api *ApiFilterer) WatchStakingModuleFeesSet(opts *bind.WatchOpts, sink chan<- *ApiStakingModuleFeesSet, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "StakingModuleFeesSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiStakingModuleFeesSet)
				if err := _Api.contract.UnpackLog(event, "StakingModuleFeesSet", log); err != nil {
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

// ParseStakingModuleFeesSet is a log parse operation binding the contract event 0x303c8ac43d1b1f9b898ddd2915a294efa01e9b07c322d7deeb7db332b66f0410.
//
// Solidity: event StakingModuleFeesSet(uint256 indexed stakingModuleId, uint256 stakingModuleFee, uint256 treasuryFee, address setBy)
func (_Api *ApiFilterer) ParseStakingModuleFeesSet(log types.Log) (*ApiStakingModuleFeesSet, error) {
	event := new(ApiStakingModuleFeesSet)
	if err := _Api.contract.UnpackLog(event, "StakingModuleFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiStakingModuleStatusSetIterator is returned from FilterStakingModuleStatusSet and is used to iterate over the raw logs and unpacked data for StakingModuleStatusSet events raised by the Api contract.
type ApiStakingModuleStatusSetIterator struct {
	Event *ApiStakingModuleStatusSet // Event containing the contract specifics and raw log

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
func (it *ApiStakingModuleStatusSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiStakingModuleStatusSet)
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
		it.Event = new(ApiStakingModuleStatusSet)
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
func (it *ApiStakingModuleStatusSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiStakingModuleStatusSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiStakingModuleStatusSet represents a StakingModuleStatusSet event raised by the Api contract.
type ApiStakingModuleStatusSet struct {
	StakingModuleId *big.Int
	Status          uint8
	SetBy           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleStatusSet is a free log retrieval operation binding the contract event 0xfd6f15fb2b48a21a60fe3d44d3c3a0433ca01e121b5124a63ec45c30ad925a17.
//
// Solidity: event StakingModuleStatusSet(uint256 indexed stakingModuleId, uint8 status, address setBy)
func (_Api *ApiFilterer) FilterStakingModuleStatusSet(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*ApiStakingModuleStatusSetIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "StakingModuleStatusSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiStakingModuleStatusSetIterator{contract: _Api.contract, event: "StakingModuleStatusSet", logs: logs, sub: sub}, nil
}

// WatchStakingModuleStatusSet is a free log subscription operation binding the contract event 0xfd6f15fb2b48a21a60fe3d44d3c3a0433ca01e121b5124a63ec45c30ad925a17.
//
// Solidity: event StakingModuleStatusSet(uint256 indexed stakingModuleId, uint8 status, address setBy)
func (_Api *ApiFilterer) WatchStakingModuleStatusSet(opts *bind.WatchOpts, sink chan<- *ApiStakingModuleStatusSet, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "StakingModuleStatusSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiStakingModuleStatusSet)
				if err := _Api.contract.UnpackLog(event, "StakingModuleStatusSet", log); err != nil {
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

// ParseStakingModuleStatusSet is a log parse operation binding the contract event 0xfd6f15fb2b48a21a60fe3d44d3c3a0433ca01e121b5124a63ec45c30ad925a17.
//
// Solidity: event StakingModuleStatusSet(uint256 indexed stakingModuleId, uint8 status, address setBy)
func (_Api *ApiFilterer) ParseStakingModuleStatusSet(log types.Log) (*ApiStakingModuleStatusSet, error) {
	event := new(ApiStakingModuleStatusSet)
	if err := _Api.contract.UnpackLog(event, "StakingModuleStatusSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiStakingModuleTargetShareSetIterator is returned from FilterStakingModuleTargetShareSet and is used to iterate over the raw logs and unpacked data for StakingModuleTargetShareSet events raised by the Api contract.
type ApiStakingModuleTargetShareSetIterator struct {
	Event *ApiStakingModuleTargetShareSet // Event containing the contract specifics and raw log

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
func (it *ApiStakingModuleTargetShareSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiStakingModuleTargetShareSet)
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
		it.Event = new(ApiStakingModuleTargetShareSet)
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
func (it *ApiStakingModuleTargetShareSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiStakingModuleTargetShareSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiStakingModuleTargetShareSet represents a StakingModuleTargetShareSet event raised by the Api contract.
type ApiStakingModuleTargetShareSet struct {
	StakingModuleId *big.Int
	TargetShare     *big.Int
	SetBy           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleTargetShareSet is a free log retrieval operation binding the contract event 0x065e5bd8e4145dd99cf69bad5871ad52d094aee07a67fcf2f418c89e49d5f20c.
//
// Solidity: event StakingModuleTargetShareSet(uint256 indexed stakingModuleId, uint256 targetShare, address setBy)
func (_Api *ApiFilterer) FilterStakingModuleTargetShareSet(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*ApiStakingModuleTargetShareSetIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "StakingModuleTargetShareSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiStakingModuleTargetShareSetIterator{contract: _Api.contract, event: "StakingModuleTargetShareSet", logs: logs, sub: sub}, nil
}

// WatchStakingModuleTargetShareSet is a free log subscription operation binding the contract event 0x065e5bd8e4145dd99cf69bad5871ad52d094aee07a67fcf2f418c89e49d5f20c.
//
// Solidity: event StakingModuleTargetShareSet(uint256 indexed stakingModuleId, uint256 targetShare, address setBy)
func (_Api *ApiFilterer) WatchStakingModuleTargetShareSet(opts *bind.WatchOpts, sink chan<- *ApiStakingModuleTargetShareSet, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "StakingModuleTargetShareSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiStakingModuleTargetShareSet)
				if err := _Api.contract.UnpackLog(event, "StakingModuleTargetShareSet", log); err != nil {
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

// ParseStakingModuleTargetShareSet is a log parse operation binding the contract event 0x065e5bd8e4145dd99cf69bad5871ad52d094aee07a67fcf2f418c89e49d5f20c.
//
// Solidity: event StakingModuleTargetShareSet(uint256 indexed stakingModuleId, uint256 targetShare, address setBy)
func (_Api *ApiFilterer) ParseStakingModuleTargetShareSet(log types.Log) (*ApiStakingModuleTargetShareSet, error) {
	event := new(ApiStakingModuleTargetShareSet)
	if err := _Api.contract.UnpackLog(event, "StakingModuleTargetShareSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiStakingRouterETHDepositedIterator is returned from FilterStakingRouterETHDeposited and is used to iterate over the raw logs and unpacked data for StakingRouterETHDeposited events raised by the Api contract.
type ApiStakingRouterETHDepositedIterator struct {
	Event *ApiStakingRouterETHDeposited // Event containing the contract specifics and raw log

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
func (it *ApiStakingRouterETHDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiStakingRouterETHDeposited)
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
		it.Event = new(ApiStakingRouterETHDeposited)
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
func (it *ApiStakingRouterETHDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiStakingRouterETHDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiStakingRouterETHDeposited represents a StakingRouterETHDeposited event raised by the Api contract.
type ApiStakingRouterETHDeposited struct {
	StakingModuleId *big.Int
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingRouterETHDeposited is a free log retrieval operation binding the contract event 0x9151b7f88aca05d432bb395647ef52b2ffc454e3c6afb69c95345af6b5a778c0.
//
// Solidity: event StakingRouterETHDeposited(uint256 indexed stakingModuleId, uint256 amount)
func (_Api *ApiFilterer) FilterStakingRouterETHDeposited(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*ApiStakingRouterETHDepositedIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "StakingRouterETHDeposited", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiStakingRouterETHDepositedIterator{contract: _Api.contract, event: "StakingRouterETHDeposited", logs: logs, sub: sub}, nil
}

// WatchStakingRouterETHDeposited is a free log subscription operation binding the contract event 0x9151b7f88aca05d432bb395647ef52b2ffc454e3c6afb69c95345af6b5a778c0.
//
// Solidity: event StakingRouterETHDeposited(uint256 indexed stakingModuleId, uint256 amount)
func (_Api *ApiFilterer) WatchStakingRouterETHDeposited(opts *bind.WatchOpts, sink chan<- *ApiStakingRouterETHDeposited, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "StakingRouterETHDeposited", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiStakingRouterETHDeposited)
				if err := _Api.contract.UnpackLog(event, "StakingRouterETHDeposited", log); err != nil {
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

// ParseStakingRouterETHDeposited is a log parse operation binding the contract event 0x9151b7f88aca05d432bb395647ef52b2ffc454e3c6afb69c95345af6b5a778c0.
//
// Solidity: event StakingRouterETHDeposited(uint256 indexed stakingModuleId, uint256 amount)
func (_Api *ApiFilterer) ParseStakingRouterETHDeposited(log types.Log) (*ApiStakingRouterETHDeposited, error) {
	event := new(ApiStakingRouterETHDeposited)
	if err := _Api.contract.UnpackLog(event, "StakingRouterETHDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiWithdrawalCredentialsSetIterator is returned from FilterWithdrawalCredentialsSet and is used to iterate over the raw logs and unpacked data for WithdrawalCredentialsSet events raised by the Api contract.
type ApiWithdrawalCredentialsSetIterator struct {
	Event *ApiWithdrawalCredentialsSet // Event containing the contract specifics and raw log

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
func (it *ApiWithdrawalCredentialsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiWithdrawalCredentialsSet)
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
		it.Event = new(ApiWithdrawalCredentialsSet)
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
func (it *ApiWithdrawalCredentialsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiWithdrawalCredentialsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiWithdrawalCredentialsSet represents a WithdrawalCredentialsSet event raised by the Api contract.
type ApiWithdrawalCredentialsSet struct {
	WithdrawalCredentials [32]byte
	SetBy                 common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCredentialsSet is a free log retrieval operation binding the contract event 0x82e72df77173eab89b00556d791a407a78f4605c5c2f0694967c8c429dd43c7c.
//
// Solidity: event WithdrawalCredentialsSet(bytes32 withdrawalCredentials, address setBy)
func (_Api *ApiFilterer) FilterWithdrawalCredentialsSet(opts *bind.FilterOpts) (*ApiWithdrawalCredentialsSetIterator, error) {

	logs, sub, err := _Api.contract.FilterLogs(opts, "WithdrawalCredentialsSet")
	if err != nil {
		return nil, err
	}
	return &ApiWithdrawalCredentialsSetIterator{contract: _Api.contract, event: "WithdrawalCredentialsSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCredentialsSet is a free log subscription operation binding the contract event 0x82e72df77173eab89b00556d791a407a78f4605c5c2f0694967c8c429dd43c7c.
//
// Solidity: event WithdrawalCredentialsSet(bytes32 withdrawalCredentials, address setBy)
func (_Api *ApiFilterer) WatchWithdrawalCredentialsSet(opts *bind.WatchOpts, sink chan<- *ApiWithdrawalCredentialsSet) (event.Subscription, error) {

	logs, sub, err := _Api.contract.WatchLogs(opts, "WithdrawalCredentialsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiWithdrawalCredentialsSet)
				if err := _Api.contract.UnpackLog(event, "WithdrawalCredentialsSet", log); err != nil {
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

// ParseWithdrawalCredentialsSet is a log parse operation binding the contract event 0x82e72df77173eab89b00556d791a407a78f4605c5c2f0694967c8c429dd43c7c.
//
// Solidity: event WithdrawalCredentialsSet(bytes32 withdrawalCredentials, address setBy)
func (_Api *ApiFilterer) ParseWithdrawalCredentialsSet(log types.Log) (*ApiWithdrawalCredentialsSet, error) {
	event := new(ApiWithdrawalCredentialsSet)
	if err := _Api.contract.UnpackLog(event, "WithdrawalCredentialsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ApiWithdrawalsCredentialsChangeFailedIterator is returned from FilterWithdrawalsCredentialsChangeFailed and is used to iterate over the raw logs and unpacked data for WithdrawalsCredentialsChangeFailed events raised by the Api contract.
type ApiWithdrawalsCredentialsChangeFailedIterator struct {
	Event *ApiWithdrawalsCredentialsChangeFailed // Event containing the contract specifics and raw log

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
func (it *ApiWithdrawalsCredentialsChangeFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApiWithdrawalsCredentialsChangeFailed)
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
		it.Event = new(ApiWithdrawalsCredentialsChangeFailed)
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
func (it *ApiWithdrawalsCredentialsChangeFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApiWithdrawalsCredentialsChangeFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApiWithdrawalsCredentialsChangeFailed represents a WithdrawalsCredentialsChangeFailed event raised by the Api contract.
type ApiWithdrawalsCredentialsChangeFailed struct {
	StakingModuleId    *big.Int
	LowLevelRevertData []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsCredentialsChangeFailed is a free log retrieval operation binding the contract event 0x0d64b11929aa111ca874dd00b5b0cc2d82b741be924ec9e3691e67c71552f623.
//
// Solidity: event WithdrawalsCredentialsChangeFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Api *ApiFilterer) FilterWithdrawalsCredentialsChangeFailed(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*ApiWithdrawalsCredentialsChangeFailedIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.FilterLogs(opts, "WithdrawalsCredentialsChangeFailed", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &ApiWithdrawalsCredentialsChangeFailedIterator{contract: _Api.contract, event: "WithdrawalsCredentialsChangeFailed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsCredentialsChangeFailed is a free log subscription operation binding the contract event 0x0d64b11929aa111ca874dd00b5b0cc2d82b741be924ec9e3691e67c71552f623.
//
// Solidity: event WithdrawalsCredentialsChangeFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Api *ApiFilterer) WatchWithdrawalsCredentialsChangeFailed(opts *bind.WatchOpts, sink chan<- *ApiWithdrawalsCredentialsChangeFailed, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Api.contract.WatchLogs(opts, "WithdrawalsCredentialsChangeFailed", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApiWithdrawalsCredentialsChangeFailed)
				if err := _Api.contract.UnpackLog(event, "WithdrawalsCredentialsChangeFailed", log); err != nil {
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

// ParseWithdrawalsCredentialsChangeFailed is a log parse operation binding the contract event 0x0d64b11929aa111ca874dd00b5b0cc2d82b741be924ec9e3691e67c71552f623.
//
// Solidity: event WithdrawalsCredentialsChangeFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Api *ApiFilterer) ParseWithdrawalsCredentialsChangeFailed(log types.Log) (*ApiWithdrawalsCredentialsChangeFailed, error) {
	event := new(ApiWithdrawalsCredentialsChangeFailed)
	if err := _Api.contract.UnpackLog(event, "WithdrawalsCredentialsChangeFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
