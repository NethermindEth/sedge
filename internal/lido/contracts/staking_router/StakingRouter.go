// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package stakingrouter

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
	TargetLimitMode            *big.Int
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
	Id                         *big.Int
	StakingModuleAddress       common.Address
	StakingModuleFee           uint16
	TreasuryFee                uint16
	StakeShareLimit            uint16
	Status                     uint8
	Name                       string
	LastDepositAt              uint64
	LastDepositBlock           *big.Int
	ExitedValidatorsCount      *big.Int
	PriorityExitShareThreshold uint16
	MaxDepositsPerBlock        uint64
	MinDepositBlockDistance    uint64
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

// StakingRouterValidatorExitData is an auto generated low-level Go binding around an user-defined struct.
type StakingRouterValidatorExitData struct {
	StakingModuleId *big.Int
	NodeOperatorId  *big.Int
	Pubkey          []byte
}

// StakingRouterValidatorsCountsCorrection is an auto generated low-level Go binding around an user-defined struct.
type StakingRouterValidatorsCountsCorrection struct {
	CurrentModuleExitedValidatorsCount       *big.Int
	CurrentNodeOperatorExitedValidatorsCount *big.Int
	NewModuleExitedValidatorsCount           *big.Int
	NewNodeOperatorExitedValidatorsCount     *big.Int
}

// StakingrouterMetaData contains all meta data concerning the Stakingrouter contract.
var StakingrouterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_depositContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AppAuthLidoFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"firstArrayLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"secondArrayLength\",\"type\":\"uint256\"}],\"name\":\"ArraysLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositContractZeroAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DirectETHTransfer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyWithdrawalsCredentials\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitedValidatorsCountCannotDecrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractVersionIncrement\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"etherValue\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositsCount\",\"type\":\"uint256\"}],\"name\":\"InvalidDepositsValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidFeeSum\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMaxDepositPerBlockValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidMinDepositBlockDistance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidPriorityExitShareThreshold\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"InvalidPublicKeysBatchLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"code\",\"type\":\"uint256\"}],\"name\":\"InvalidReportData\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"actual\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"}],\"name\":\"InvalidSignaturesBatchLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidStakeShareLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroContractVersionOnInit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reportedExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositedValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"ReportedExitedValidatorsExceedDeposited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModuleAddressExists\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModuleNotActive\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModuleStatusTheSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModuleUnregistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModuleWrongName\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StakingModulesLimitExceeded\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"UnexpectedContractVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"currentModuleExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentNodeOpExitedValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"UnexpectedCurrentValidatorsCount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newModuleTotalExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newModuleTotalExitedValidatorsCountInStakingRouter\",\"type\":\"uint256\"}],\"name\":\"UnexpectedFinalExitedValidatorsCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnrecoverableModuleError\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressAdmin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressLido\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAddressStakingModule\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelRevertData\",\"type\":\"bytes\"}],\"name\":\"ExitedAndStuckValidatorsCountsUpdateFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelRevertData\",\"type\":\"bytes\"}],\"name\":\"RewardsMintedReportFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"stakingModule\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"createdBy\",\"type\":\"address\"}],\"name\":\"StakingModuleAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"}],\"name\":\"StakingModuleExitNotificationFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"unreportedExitedValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"StakingModuleExitedValidatorsIncompleteReporting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakingModuleFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"treasuryFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"StakingModuleFeesSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxDepositsPerBlock\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"StakingModuleMaxDepositsPerBlockSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"minDepositBlockDistance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"StakingModuleMinDepositBlockDistanceSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stakeShareLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"priorityExitShareThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"StakingModuleShareLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"enumStakingRouter.StakingModuleStatus\",\"name\":\"status\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"StakingModuleStatusSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"StakingRouterETHDeposited\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"withdrawalCredentials\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"setBy\",\"type\":\"address\"}],\"name\":\"WithdrawalCredentialsSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"lowLevelRevertData\",\"type\":\"bytes\"}],\"name\":\"WithdrawalsCredentialsChangeFailed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_CONTRACT\",\"outputs\":[{\"internalType\":\"contractIDepositContract\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_PRECISION_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_WITHDRAWAL_CREDENTIALS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STAKING_MODULES_COUNT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_STAKING_MODULE_NAME_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORT_EXITED_VALIDATORS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORT_REWARDS_MINTED_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORT_VALIDATOR_EXITING_STATUS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORT_VALIDATOR_EXIT_TRIGGERED_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MODULE_MANAGE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_MODULE_UNVETTING_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"TOTAL_BASIS_POINTS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNSAFE_SET_EXITED_VALIDATORS_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"_stakingModuleAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_stakeShareLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_priorityExitShareThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingModuleFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_treasuryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDepositsPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minDepositBlockDistance\",\"type\":\"uint256\"}],\"name\":\"addStakingModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_nodeOperatorIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_vettedSigningKeysCounts\",\"type\":\"bytes\"}],\"name\":\"decreaseStakingModuleVettedKeysCountByNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_depositsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_depositCalldata\",\"type\":\"bytes\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeUpgrade_v3\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getAllNodeOperatorDigests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.NodeOperatorSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"internalType\":\"structStakingRouter.NodeOperatorDigest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllStakingModuleDigests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeNodeOperatorsCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"stakingModuleAddress\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"stakingModuleFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"treasuryFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"stakeShareLimit\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"lastDepositAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"lastDepositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"priorityExitShareThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"maxDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minDepositBlockDistance\",\"type\":\"uint64\"}],\"internalType\":\"structStakingRouter.StakingModule\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.StakingModuleSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"internalType\":\"structStakingRouter.StakingModuleDigest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_depositsCount\",\"type\":\"uint256\"}],\"name\":\"getDepositsAllocation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"allocated\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"allocations\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLido\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_nodeOperatorIds\",\"type\":\"uint256[]\"}],\"name\":\"getNodeOperatorDigests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.NodeOperatorSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"internalType\":\"structStakingRouter.NodeOperatorDigest[]\",\"name\":\"digests\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_limit\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorDigests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.NodeOperatorSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"internalType\":\"structStakingRouter.NodeOperatorDigest[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.NodeOperatorSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingFeeAggregateDistribution\",\"outputs\":[{\"internalType\":\"uint96\",\"name\":\"modulesFee\",\"type\":\"uint96\"},{\"internalType\":\"uint96\",\"name\":\"treasuryFee\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"basePrecision\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingFeeAggregateDistributionE4Precision\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"modulesFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"treasuryFee\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModule\",\"outputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"stakingModuleAddress\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"stakingModuleFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"treasuryFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"stakeShareLimit\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"lastDepositAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"lastDepositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"priorityExitShareThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"maxDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minDepositBlockDistance\",\"type\":\"uint64\"}],\"internalType\":\"structStakingRouter.StakingModule\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleActiveValidatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"activeValidatorsCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_stakingModuleIds\",\"type\":\"uint256[]\"}],\"name\":\"getStakingModuleDigests\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"activeNodeOperatorsCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"stakingModuleAddress\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"stakingModuleFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"treasuryFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"stakeShareLimit\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"lastDepositAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"lastDepositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"priorityExitShareThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"maxDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minDepositBlockDistance\",\"type\":\"uint64\"}],\"internalType\":\"structStakingRouter.StakingModule\",\"name\":\"state\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.StakingModuleSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"internalType\":\"structStakingRouter.StakingModuleDigest[]\",\"name\":\"digests\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingModuleIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"stakingModuleIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleIsActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleIsDepositsPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleIsStopped\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleLastDepositBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDepositsValue\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleMaxDepositsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleMaxDepositsPerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleMinDepositBlockDistance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleStatus\",\"outputs\":[{\"internalType\":\"enumStakingRouter.StakingModuleStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"getStakingModuleSummary\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.StakingModuleSummary\",\"name\":\"summary\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingModules\",\"outputs\":[{\"components\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"internalType\":\"address\",\"name\":\"stakingModuleAddress\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"stakingModuleFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"treasuryFee\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"stakeShareLimit\",\"type\":\"uint16\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint64\",\"name\":\"lastDepositAt\",\"type\":\"uint64\"},{\"internalType\":\"uint256\",\"name\":\"lastDepositBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"priorityExitShareThreshold\",\"type\":\"uint16\"},{\"internalType\":\"uint64\",\"name\":\"maxDepositsPerBlock\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"minDepositBlockDistance\",\"type\":\"uint64\"}],\"internalType\":\"structStakingRouter.StakingModule[]\",\"name\":\"res\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingModulesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingRewardsDistribution\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"recipients\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"stakingModuleIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint96[]\",\"name\":\"stakingModuleFees\",\"type\":\"uint96[]\"},{\"internalType\":\"uint96\",\"name\":\"totalFee\",\"type\":\"uint96\"},{\"internalType\":\"uint256\",\"name\":\"precisionPoints\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalFeeE4Precision\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"totalFee\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWithdrawalCredentials\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"}],\"name\":\"hasStakingModule\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lido\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"_withdrawalCredentials\",\"type\":\"bytes32\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"internalType\":\"structStakingRouter.ValidatorExitData[]\",\"name\":\"validatorExitData\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_withdrawalRequestPaidFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_exitType\",\"type\":\"uint256\"}],\"name\":\"onValidatorExitTriggered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onValidatorsCountsByNodeOperatorReportingFinished\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_stakingModuleIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_totalShares\",\"type\":\"uint256[]\"}],\"name\":\"reportRewardsMinted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_nodeOperatorIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"_exitedValidatorsCounts\",\"type\":\"bytes\"}],\"name\":\"reportStakingModuleExitedValidatorsCountByNodeOperator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_proofSlotTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"_eligibleToExitInSec\",\"type\":\"uint256\"}],\"name\":\"reportValidatorExitDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"enumStakingRouter.StakingModuleStatus\",\"name\":\"_status\",\"type\":\"uint8\"}],\"name\":\"setStakingModuleStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_withdrawalCredentials\",\"type\":\"bytes32\"}],\"name\":\"setWithdrawalCredentials\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_triggerUpdateFinish\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentModuleExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentNodeOperatorExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newModuleExitedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"newNodeOperatorExitedValidatorsCount\",\"type\":\"uint256\"}],\"internalType\":\"structStakingRouter.ValidatorsCountsCorrection\",\"name\":\"_correction\",\"type\":\"tuple\"}],\"name\":\"unsafeSetExitedValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"_stakingModuleIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_exitedValidatorsCounts\",\"type\":\"uint256[]\"}],\"name\":\"updateExitedValidatorsCountByStakingModule\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakeShareLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_priorityExitShareThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_stakingModuleFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_treasuryFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_maxDepositsPerBlock\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minDepositBlockDistance\",\"type\":\"uint256\"}],\"name\":\"updateStakingModule\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_stakingModuleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetLimitMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_targetLimit\",\"type\":\"uint256\"}],\"name\":\"updateTargetValidatorsLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052600436106103f35760003560e01c80639010d07c11610208578063c445ea7511610118578063e016e6f7116100ab578063efcdcc0e1161007a578063efcdcc0e14610cd5578063f07ff28a14610d05578063f2aebb6514610d25578063f8bb6d4214610d47578063fa5093eb14610d6757600080fd5b8063e016e6f714610c3f578063e24ce9f114610c61578063e39fdbe914610c81578063e97ee8cc14610cb557600080fd5b8063cb8fd4da116100e7578063cb8fd4da14610bca578063d0a2b1b814610bea578063d547741f14610c0a578063db3c7ba714610c2a57600080fd5b8063c445ea7514610b3a578063c82b1bb114610b5c578063c8ac498014610b8a578063ca15c87314610baa57600080fd5b8063a217fddf1161019b578063aa5a1b9d1161016a578063aa5a1b9d14610a7a578063abd44a2414610aa7578063af12409714610ac7578063ba21ccae14610ae7578063bc1bb19014610b0d57600080fd5b8063a217fddf146109e4578063a4258a8d146109f9578063a734329c14610a19578063aa0b7db714610a6757600080fd5b80639b75b4ef116101d75780639b75b4ef1461095a5780639dd068481461096f5780639fbb7bae1461098f5780639fc5a6ed146109b757600080fd5b80639010d07c146108c6578063909c01de146108e657806391d148541461091a57806396b5d81c1461093a57600080fd5b80634b3a1cb7116103035780636b96736b116102965780637a74884d116102655780637a74884d146108205780637b274031146108545780637c8da51c146108745780638525e3a1146108915780638aa10435146108b157600080fd5b80636b96736b146107975780636d395b7e146107cb57806371416583146107e0578063771895831461080057600080fd5b80636183214d116102d25780636183214d146107085780636608b11b1461072a5780636a516b471461074a5780636ada55b91461077757600080fd5b80634b3a1cb71461069c57806356396715146106b157806357993b85146106c65780636133f985146106e857600080fd5b8063248a9ca3116103865780633240a322116103555780633240a322146105fa57806332c4962c1461062757806336568abe14610647578063473e0433146106675780634a7583b61461068757600080fd5b8063248a9ca314610582578063271662ec146105a25780632c201d31146105b85780632f2ff15d146105da57600080fd5b80631565d2f2116103c25780631565d2f2146104da57806319c64b791461050e5780631d1b9d3c1461052e57806320e948c81461056257600080fd5b806301ffc9a7146104165780630519fbbf1461044b57806307e203ac146104795780630fb31c84146104a657600080fd5b36610411576040516309fb455960e41b815260040160405180910390fd5b600080fd5b34801561042257600080fd5b50610436610431366004614baf565b610da2565b60405190151581526020015b60405180910390f35b34801561045757600080fd5b5061046b610466366004614bd9565b610dcd565b604051908152602001610442565b34801561048557600080fd5b50610499610494366004614bd9565b610e48565b6040516104429190614bf2565b3480156104b257600080fd5b5061046b7fbe1bd143a0dde8a867d58aab054bfdb25250951665c4570e39abc3b3de3c2d6c81565b3480156104e657600080fd5b5061046b7f55180e25fcacf9af017d35d497765476319b23896daa1f9bc2b38fa80b36a16381565b34801561051a57600080fd5b5061046b610529366004614c13565b610e98565b34801561053a57600080fd5b5061046b7f779e5c23cb7a5bcb9bfe1e9a5165a00057f12bcdfd13e374540fdf1a1cd9113781565b34801561056e57600080fd5b5061046b61057d366004614bd9565b610f15565b34801561058e57600080fd5b5061046b61059d366004614bd9565b610f41565b3480156105ae57600080fd5b5061046b61271081565b3480156105c457600080fd5b506105d86105d3366004614c7d565b610f63565b005b3480156105e657600080fd5b506105d86105f5366004614d12565b61100e565b34801561060657600080fd5b5061061a610615366004614bd9565b611030565b6040516104429190614d8c565b34801561063357600080fd5b506105d8610642366004614df2565b6110b1565b34801561065357600080fd5b506105d8610662366004614d12565b611153565b34801561067357600080fd5b5061046b610682366004614bd9565b6111d6565b34801561069357600080fd5b5061046b6111ee565b3480156106a857600080fd5b5061046b602081565b3480156106bd57600080fd5b5061046b61121d565b3480156106d257600080fd5b506106db611247565b6040516104429190614fb9565b3480156106f457600080fd5b506105d861070336600461505e565b611254565b34801561071457600080fd5b5061071d611347565b604051610442919061509a565b34801561073657600080fd5b50610436610745366004614bd9565b611543565b34801561075657600080fd5b5061075f611568565b6040516001600160a01b039091168152602001610442565b34801561078357600080fd5b50610436610792366004614bd9565b611592565b3480156107a357600080fd5b5061075f7f00000000000000000000000000000000219ab540356cbb839cbe05303d7705fa81565b3480156107d757600080fd5b506105d861159b565b3480156107ec57600080fd5b506105d86107fb366004615140565b6115b1565b34801561080c57600080fd5b506105d861081b366004615190565b61173e565b34801561082c57600080fd5b5061046b7fe7c742a54cd11fc9749a47ab34bdcd7327820908e8d0d48b4a5c7f17b029409881565b34801561086057600080fd5b506105d861086f366004615230565b611782565b34801561088057600080fd5b5061046b68056bc75e2d6310000081565b34801561089d57600080fd5b506106db6108ac366004615351565b6119ed565b3480156108bd57600080fd5b5061046b611bce565b3480156108d257600080fd5b5061075f6108e1366004614c13565b611bf8565b3480156108f257600080fd5b5061046b7f240525496a9dc32284b17ce03b43e539e4bd81414634ee54395030d793463b5781565b34801561092657600080fd5b50610436610935366004614d12565b611c24565b34801561094657600080fd5b5061046b610955366004614bd9565b611c5c565b34801561096657600080fd5b5061046b601f81565b34801561097b57600080fd5b506105d861098a366004615385565b611cb4565b34801561099b57600080fd5b506109a4611d3c565b60405161ffff9091168152602001610442565b3480156109c357600080fd5b506109d76109d2366004614bd9565b611d6a565b60405161044291906153ef565b3480156109f057600080fd5b5061046b600081565b348015610a0557600080fd5b506105d8610a143660046153fd565b611d94565b348015610a2557600080fd5b50610436610a34366004614bd9565b60009081527f9b48f5b32acb95b982effe269feac267eead113c4b5af14ffeb9aadac18a6e9c6020526040902054151590565b6105d8610a75366004615487565b612000565b348015610a8657600080fd5b50610a9a610a95366004614c13565b61220c565b60405161044291906154d9565b348015610ab357600080fd5b5061046b610ac23660046154e8565b6122d6565b348015610ad357600080fd5b506105d8610ae23660046154e8565b6124a8565b348015610af357600080fd5b50610afc61264f565b604051610442959493929190615582565b348015610b1957600080fd5b50610b2d610b28366004614bd9565b6129ca565b6040516104429190615641565b348015610b4657600080fd5b5061046b600080516020615eb283398151915281565b348015610b6857600080fd5b50610b7c610b77366004614bd9565b612b3e565b604051610442929190615654565b348015610b9657600080fd5b506105d8610ba5366004614c7d565b612b56565b348015610bb657600080fd5b5061046b610bc5366004614bd9565b612bb5565b348015610bd657600080fd5b5061046b610be5366004614bd9565b612bd9565b348015610bf657600080fd5b506105d8610c0536600461566d565b612c01565b348015610c1657600080fd5b506105d8610c25366004614d12565b612c99565b348015610c3657600080fd5b506105d8612cb6565b348015610c4b57600080fd5b5061046b600080516020615ed283398151915281565b348015610c6d57600080fd5b50610436610c7c366004614bd9565b612e17565b348015610c8d57600080fd5b5061046b7f0766e72e5c008b3df8129fb356d9176eef8544f6241e078b7d61aff604f8812b81565b348015610cc157600080fd5b506105d8610cd0366004614bd9565b612e20565b348015610ce157600080fd5b50610cea612fdf565b6040805161ffff938416815292909116602083015201610442565b348015610d1157600080fd5b5061061a610d203660046156a1565b613026565b348015610d3157600080fd5b50610d3a6131b1565b60405161044291906156e7565b348015610d5357600080fd5b5061061a610d623660046156fa565b613247565b348015610d7357600080fd5b50610d7c6132e3565b604080516001600160601b03948516815293909216602084015290820152606001610442565b60006001600160e01b03198216635a05180f60e01b1480610dc75750610dc782613351565b92915050565b6000610dd882613386565b6001600160a01b031663d087d2886040518163ffffffff1660e01b815260040160206040518083038186803b158015610e1057600080fd5b505afa158015610e24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dc79190615726565b610e6c60405180606001604052806000815260200160008152602001600081525090565b6000610e77836129ca565b602001519050610e8681613391565b60408501526020840152825250919050565b60008080610eb7610eb26801bc16d674ec8000008661576b565b613414565b92509250506000610ec786613660565b9050818181518110610edb57610edb61577f565b602002602001015160c00151838281518110610ef957610ef961577f565b6020026020010151610f0b9190615795565b9695505050505050565b6000610f28610f2383613660565b6136b9565b600501546201000090046001600160401b031692915050565b6000908152600080516020615e92833981519152602052604090206001015490565b7f240525496a9dc32284b17ce03b43e539e4bd81414634ee54395030d793463b57610f8e81336136e9565b610f9a8585858561374d565b610fa386613386565b6001600160a01b031663b643189b868686866040518563ffffffff1660e01b8152600401610fd494939291906157d5565b600060405180830381600087803b158015610fee57600080fd5b505af1158015611002573d6000803e3d6000fd5b50505050505050505050565b61101782610f41565b61102181336136e9565b61102b83836137ec565b505050565b6060610dc782600061104185613386565b6001600160a01b031663a70c70e46040518163ffffffff1660e01b815260040160206040518083038186803b15801561107957600080fd5b505afa15801561108d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d629190615726565b7fbe1bd143a0dde8a867d58aab054bfdb25250951665c4570e39abc3b3de3c2d6c6110dc81336136e9565b6110e587613386565b6001600160a01b03166357f9c34187878787876040518663ffffffff1660e01b8152600401611118959493929190615807565b600060405180830381600087803b15801561113257600080fd5b505af1158015611146573d6000803e3d6000fd5b5050505050505050505050565b6001600160a01b03811633146111c85760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b60648201526084015b60405180910390fd5b6111d2828261381b565b5050565b60006111e4610f2383613660565b6003015492915050565b60006112187f1b3ef9db2d6f0727a31622833b45264c21051726d23ddb6f73b3b65628cafcc35490565b905090565b60006112187fabeb05279af36da5d476d7f950157cd2ea98a4166fa68a6bc97ce3a22fbb93c05490565b60606112186108ac6131b1565b6001600160a01b03831661127b5760405163371262eb60e11b815260040160405180910390fd5b6001600160a01b0382166112a257604051630c75384960e01b815260040160405180910390fd5b6112ac600361384a565b6112b760008461387c565b6112e07f706b9ed9846c161ad535be9b6345c3a7b2cb929e8d4a7254dee9ba6e6f8e5531839055565b6113097fabeb05279af36da5d476d7f950157cd2ea98a4166fa68a6bc97ce3a22fbb93c0829055565b604080518281523360208201527f82e72df77173eab89b00556d791a407a78f4605c5c2f0694967c8c429dd43c7c91015b60405180910390a1505050565b606060006113536111ee565b9050806001600160401b0381111561136d5761136d6151ea565b6040519080825280602002602001820160405280156113a657816020015b6113936149b3565b81526020019060019003908161138b5790505b50915060005b8181101561153e576113bd816136b9565b604080516101a081018252825462ffffff81168252630100000081046001600160a01b03166020830152600160b81b810461ffff90811693830193909352600160c81b810483166060830152600160d81b81049092166080820152600160e81b90910460ff1660a082015260018201805491929160c08401919061144090615839565b80601f016020809104026020016040519081016040528092919081815260200182805461146c90615839565b80156114b95780601f1061148e576101008083540402835291602001916114b9565b820191906000526020600020905b81548152906001019060200180831161149c57829003601f168201915b505050918352505060028201546001600160401b039081166020830152600383015460408301526004830154606083015260059092015461ffff81166080830152620100008104831660a0830152600160501b900490911660c090910152835184908390811061152b5761152b61577f565b60209081029190910101526001016113ac565b505090565b6000805b61155083611d6a565b6002811115611561576115616153b7565b1492915050565b60006112187f706b9ed9846c161ad535be9b6345c3a7b2cb929e8d4a7254dee9ba6e6f8e55315490565b60006002611547565b6115a56002613886565b6115af60036138bc565b565b7f0766e72e5c008b3df8129fb356d9176eef8544f6241e078b7d61aff604f8812b6115dc81336136e9565b3660005b85811015611735578686828181106115fa576115fa61577f565b905060200281019061160c919061586e565b91506116188235613386565b6001600160a01b031663693cc6006020840135611638604086018661588e565b89896040518663ffffffff1660e01b815260040161165a9594939291906158d4565b600060405180830381600087803b15801561167457600080fd5b505af1925050508015611685575060015b611725573d8080156116b3576040519150601f19603f3d011682016040523d82523d6000602084013e6116b8565b606091505b5080516116d857604051638fd297d960e01b815260040160405180910390fd5b602083013583357fb639213d4cc5d7a615491fb0505dd448dee5074f322660125b7171993bf9bb1d61170d604087018761588e565b60405161171b929190615902565b60405180910390a3505b61172e81615916565b90506115e0565b50505050505050565b600080516020615ed283398151915261175781336136e9565b6000611765610f238a613660565b9050611777818a8a8a8a8a8a8a6138ee565b505050505050505050565b7f55180e25fcacf9af017d35d497765476319b23896daa1f9bc2b38fa80b36a1636117ad81336136e9565b60006117bb610f2387613660565b8054604051632cc1db0f60e21b815260048101889052919250630100000090046001600160a01b031690600090829063b3076c3c906024016101006040518083038186803b15801561180c57600080fd5b505afa158015611820573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118449190615931565b50509550505050505082600401548560000151141580611868575080856020015114155b156118935760048381015460405163c7c450d560e01b815291820152602481018290526044016111bf565b60408086015160048086019190915560608701519151631282406d60e31b81526001600160a01b038516926394120368926118da928c929101918252602082015260400190565b600060405180830381600087803b1580156118f457600080fd5b505af1158015611908573d6000803e3d6000fd5b5050505060008061191884613391565b5091509150808760400151111561195357866040015181604051630b72c59d60e21b81526004016111bf929190918252602082015260400190565b8715611002578660400151821461198e5781876040015160405163dcab2a8960e01b81526004016111bf929190918252602082015260400190565b836001600160a01b031663e864299e6040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156119c957600080fd5b505af11580156119dd573d6000803e3d6000fd5b5050505050505050505050505050565b606081516001600160401b03811115611a0857611a086151ea565b604051908082528060200260200182016040528015611a4157816020015b611a2e614a1e565b815260200190600190039081611a265790505b50905060005b8251811015611bc8576000611a74848381518110611a6757611a6761577f565b60200260200101516129ca565b90506000816020015190506040518060800160405280826001600160a01b031663a70c70e46040518163ffffffff1660e01b815260040160206040518083038186803b158015611ac357600080fd5b505afa158015611ad7573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611afb9190615726565b8152602001826001600160a01b0316638469cbd36040518163ffffffff1660e01b815260040160206040518083038186803b158015611b3957600080fd5b505afa158015611b4d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b719190615726565b8152602001838152602001611b9e878681518110611b9157611b9161577f565b6020026020010151610e48565b815250848481518110611bb357611bb361577f565b60209081029190910101525050600101611a47565b50919050565b60006112187f4dd0f6662ba1d6b081f08b350f5e9a6a7b15cf586926ba66f753594928fa64a65490565b6000828152600080516020615e7283398151915260205260408120611c1d9083613b75565b9392505050565b6000918252600080516020615e92833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600080611c6b610f2384613660565b80549091506000908190611c8e90630100000090046001600160a01b0316613391565b5091509150611ca1836004015483613b81565b611cab9082615795565b95945050505050565b600080516020615ed2833981519152611ccd81336136e9565b611cd685613386565b6040516308a679ad60e01b81526004810186905260248101859052604481018490526001600160a01b0391909116906308a679ad90606401600060405180830381600087803b158015611d2857600080fd5b505af1158015611777573d6000803e3d6000fd5b6000806000611d4961264f565b94509450505050611d63826001600160601b031682613b97565b9250505090565b6000611d78610f2383613660565b54600160e81b900460ff166002811115610dc757610dc76153b7565b600080516020615ed2833981519152611dad81336136e9565b6001600160a01b038816611dd457604051632ec8c66160e01b815260040160405180910390fd5b881580611de15750601f89115b15611dff5760405163ac18716960e01b815260040160405180910390fd5b6000611e096111ee565b905060208110611e2c5760405163309eed9960e01b815260040160405180910390fd5b60005b81811015611e7c57611e40816136b9565b546001600160a01b038b81166301000000909204161415611e745760405163050f969d60e41b815260040160405180910390fd5b600101611e2f565b506000611e88826136b9565b90506000611eb47ff9a85ae945d8134f58bd2ee028636634dcb9e812798acb5c806bf1951232a2255490565b611ebf90600161598e565b825462ffffff191662ffffff82161783559050611ee0600183018e8e614a6d565b5081547fffff00ffffffffffff0000000000000000000000000000000000000000ffffff1663010000006001600160a01b038d160260ff60e81b1916178255611f308262ffffff83166000613bb0565b611f3f8162ffffff1684613c0f565b62ffffff81167ff9a85ae945d8134f58bd2ee028636634dcb9e812798acb5c806bf1951232a22555611f99611f758460016159b5565b7f1b3ef9db2d6f0727a31622833b45264c21051726d23ddb6f73b3b65628cafcc355565b8062ffffff167f43b5213f0e1666cd0b8692a73686164c94deb955a59c65e10dee8bb958e7ce3e8c8f8f33604051611fd494939291906159cd565b60405180910390a2611ff1828262ffffff168c8c8c8c8c8c6138ee565b50505050505050505050505050565b7f706b9ed9846c161ad535be9b6345c3a7b2cb929e8d4a7254dee9ba6e6f8e5531546001600160a01b0316336001600160a01b03161461205357604051637e71782360e01b815260040160405180910390fd5b600061205d61121d565b90508061207d5760405163180a97cd60e21b815260040160405180910390fd5b600061208b610f2386613660565b905060008154600160e81b900460ff1660028111156120ac576120ac6153b7565b60028111156120bd576120bd6153b7565b146120db5760405163322e64fb60e11b815260040160405180910390fd5b346120ef6801bc16d674ec80000088615a04565b81146121185760405163023db95b60e21b815260048101829052602481018890526044016111bf565b612123828783613bb0565b86156117355781546040516317dc836b60e31b8152600091829163010000009091046001600160a01b03169063bee41b5890612167908c908b908b90600401615a23565b600060405180830381600087803b15801561218157600080fd5b505af1158015612195573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526121bd9190810190615aa0565b9150915060004790506121f38a876040516020016121dd91815260200190565b6040516020818303038152906040528585613c52565b47846121ff8284615795565b1461114657611146615af9565b612214614af1565b600061221f846129ca565b6020015190506000806000806000856001600160a01b031663b3076c3c896040518263ffffffff1660e01b815260040161225b91815260200190565b6101006040518083038186803b15801561227457600080fd5b505afa158015612288573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122ac9190615931565b968e5260208e019590955260a08d015250505060c089015260e08801525094979650505050505050565b6000600080516020615eb28339815191526122f181336136e9565b6122fb8584613dd0565b6000805b8681101561249d57600088888381811061231b5761231b61577f565b9050602002013590506000612332610f2383613660565b60048101549091508089898681811061234d5761234d61577f565b90506020020135101561237357604051632f789f4960e21b815260040160405180910390fd5b8154600090819061239390630100000090046001600160a01b0316613391565b5091509150808b8b888181106123ab576123ab61577f565b9050602002013511156123f6578a8a878181106123ca576123ca61577f565b9050602002013581604051630b72c59d60e21b81526004016111bf929190918252602082015260400190565b828b8b888181106124095761240961577f565b9050602002013561241a9190615795565b61242490886159b5565b96508282101561246c57847fdd2523ca96a639ba7e17420698937f71eddd8af012ccb36ff5c8fe96141acae961245a8486615795565b60405190815260200160405180910390a25b8a8a8781811061247e5761247e61577f565b90506020020135846004018190555085600101955050505050506122ff565b509695505050505050565b7f779e5c23cb7a5bcb9bfe1e9a5165a00057f12bcdfd13e374540fdf1a1cd911376124d381336136e9565b6124dd8483613dd0565b60005b848110156126475760008484838181106124fc576124fc61577f565b90506020020135111561263f5761252a86868381811061251e5761251e61577f565b90506020020135613386565b6001600160a01b0316638d7e401785858481811061254a5761254a61577f565b905060200201356040518263ffffffff1660e01b815260040161256f91815260200190565b600060405180830381600087803b15801561258957600080fd5b505af192505050801561259a575060015b61263f573d8080156125c8576040519150601f19603f3d011682016040523d82523d6000602084013e6125cd565b606091505b5080516125ed57604051638fd297d960e01b815260040160405180910390fd5b8686838181106125ff576125ff61577f565b905060200201357ff74208fedac7280fd11f8de0be14e00423dc5076da8e8ec8ca90e09257fff1b3826040516126359190615b0f565b60405180910390a2505b6001016124e0565b505050505050565b6060806060600080600080612662613dfa565b80519193509150801580612674575082155b156126b65750506040805160008082526020820181815282840182815260608401909452919850909650909450925068056bc75e2d6310000091506129c39050565b68056bc75e2d631000009350806001600160401b038111156126da576126da6151ea565b604051908082528060200260200182016040528015612703578160200160208202803683370190505b509650806001600160401b0381111561271e5761271e6151ea565b604051908082528060200260200182016040528015612747578160200160208202803683370190505b509750806001600160401b03811115612762576127626151ea565b60405190808252806020026020018201604052801561278b578160200160208202803683370190505b5095506000808060005b848110156129905760008682815181106127b1576127b161577f565b602002602001015160c001511115612988578581815181106127d5576127d561577f565b60200260200101516020015162ffffff168b85815181106127f8576127f861577f565b60200260200101818152505086888783815181106128185761281861577f565b602002602001015160c0015161282e9190615a04565b612838919061576b565b925085818151811061284c5761284c61577f565b6020026020010151600001518c858151811061286a5761286a61577f565b60200260200101906001600160a01b031690816001600160a01b03168152505061271086828151811061289f5761289f61577f565b60200260200101516040015161ffff16846128ba9190615a04565b6128c4919061576b565b915060028682815181106128da576128da61577f565b602002602001015160a0015160028111156128f7576128f76153b7565b1461293057818a858151811061290f5761290f61577f565b60200260200101906001600160601b031690816001600160601b0316815250505b816127108783815181106129465761294661577f565b60200260200101516060015161ffff16856129619190615a04565b61296b919061576b565b6129759190615b22565b61297f908a615b22565b98506001909301925b600101612795565b5086886001600160601b031611156129aa576129aa615af9565b838310156129bc57828a52828b528289525b5050505050505b9091929394565b6129d26149b3565b6129de610f2383613660565b604080516101a081018252825462ffffff81168252630100000081046001600160a01b03166020830152600160b81b810461ffff90811693830193909352600160c81b810483166060830152600160d81b81049092166080820152600160e81b90910460ff1660a082015260018201805491929160c084019190612a6190615839565b80601f0160208091040260200160405190810160405280929190818152602001828054612a8d90615839565b8015612ada5780601f10612aaf57610100808354040283529160200191612ada565b820191906000526020600020905b815481529060010190602001808311612abd57829003601f168201915b505050918352505060028201546001600160401b039081166020830152600383015460408301526004830154606083015260059092015461ffff81166080830152620100008104831660a0830152600160501b900490911660c09091015292915050565b60006060612b4b83613414565b509094909350915050565b600080516020615eb2833981519152612b6f81336136e9565b612b7b8585858561374d565b612b8486613386565b6001600160a01b0316639b00c146868686866040518563ffffffff1660e01b8152600401610fd494939291906157d5565b6000818152600080516020615e7283398151915260205260408120610dc790613ec8565b6000612be7610f2383613660565b60050154600160501b90046001600160401b031692915050565b600080516020615ed2833981519152612c1a81336136e9565b6000612c28610f2385613660565b9050826002811115612c3c57612c3c6153b7565b8154600160e81b900460ff166002811115612c5957612c596153b7565b6002811115612c6a57612c6a6153b7565b1415612c8957604051635ca16fa760e11b815260040160405180910390fd5b612c938184613ed2565b50505050565b612ca282610f41565b612cac81336136e9565b61102b838361381b565b600080516020615eb2833981519152612ccf81336136e9565b6000612cd96111ee565b905060008060005b83811015612e1057612cf2816136b9565b8054909350630100000090046001600160a01b031691506000612d1483613391565b505090508360040154811415612e0757826001600160a01b031663e864299e6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015612d5f57600080fd5b505af1925050508015612d70575060015b612e07573d808015612d9e576040519150601f19603f3d011682016040523d82523d6000602084013e612da3565b606091505b508051612dc357604051638fd297d960e01b815260040160405180910390fd5b845460405162ffffff909116907fe74bf895f0c3a2d6c74c40cbb362fdd9640035fc4226c72e3843809ad2a9d2b590612dfd908490615b0f565b60405180910390a2505b50600101612ce1565b5050505050565b60006001611547565b7fe7c742a54cd11fc9749a47ab34bdcd7327820908e8d0d48b4a5c7f17b0294098612e4b81336136e9565b612e747fabeb05279af36da5d476d7f950157cd2ea98a4166fa68a6bc97ce3a22fbb93c0839055565b6000612e7e6111ee565b905060005b81811015612fa9576000612e96826136b9565b90508160010191508060000160039054906101000a90046001600160a01b03166001600160a01b03166390c09bdb6040518163ffffffff1660e01b8152600401600060405180830381600087803b158015612ef057600080fd5b505af1925050508015612f01575060015b612fa3573d808015612f2f576040519150601f19603f3d011682016040523d82523d6000602084013e612f34565b606091505b508051612f5457604051638fd297d960e01b815260040160405180910390fd5b612f5f826001613ed2565b815460405162ffffff909116907f0d64b11929aa111ca874dd00b5b0cc2d82b741be924ec9e3691e67c71552f62390612f99908490615b0f565b60405180910390a2505b50612e83565b50604080518481523360208201527f82e72df77173eab89b00556d791a407a78f4605c5c2f0694967c8c429dd43c7c910161133a565b6000806000806000612fef6132e3565b92506001600160601b031692506001600160601b031692506130118382613b97565b945061301d8282613b97565b93505050509091565b6060600061303384613386565b905082516001600160401b0381111561304e5761304e6151ea565b60405190808252806020026020018201604052801561308757816020015b613074614b36565b81526020019060019003908161306c5790505b50915060005b83518110156131a95760405180606001604052808583815181106130b3576130b361577f565b60200260200101518152602001836001600160a01b0316635e2fb9088785815181106130e1576130e161577f565b60200260200101516040518263ffffffff1660e01b815260040161310791815260200190565b60206040518083038186803b15801561311f57600080fd5b505afa158015613133573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906131579190615b44565b15158152602001613181878785815181106131745761317461577f565b602002602001015161220c565b8152508382815181106131965761319661577f565b602090810291909101015260010161308d565b505092915050565b606060006131bd6111ee565b9050806001600160401b038111156131d7576131d76151ea565b604051908082528060200260200182016040528015613200578160200160208202803683370190505b50915060005b8181101561153e57613217816136b9565b54835162ffffff909116908490839081106132345761323461577f565b6020908102919091010152600101613206565b60606132db8461325686613386565b604051634febc81b60e01b815260048101879052602481018690526001600160a01b039190911690634febc81b9060440160006040518083038186803b15801561329f57600080fd5b505afa1580156132b3573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610d209190810190615bbc565b949350505050565b6000806000606060006132f461264f565b9650909450925060009150505b82518110156133395782818151811061331c5761331c61577f565b60200260200101518661332f9190615b22565b9550600101613301565b506133448582615bf0565b93505050909192565b9055565b60006001600160e01b03198216637965db0b60e01b1480610dc757506301ffc9a760e01b6001600160e01b0319831614610dc7565b6000610dc782613f88565b6000806000836001600160a01b0316639abddf096040518163ffffffff1660e01b815260040160606040518083038186803b1580156133cf57600080fd5b505afa1580156133e3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906134079190615c18565b9250925092509193909250565b60006060806000613423613dfa565b8051909350909150806001600160401b03811115613443576134436151ea565b60405190808252806020026020018201604052801561346c578160200160208202803683370190505b50935080156136575761347f86836159b5565b91506000816001600160401b0381111561349b5761349b6151ea565b6040519080825280602002602001820160405280156134c4578160200160208202803683370190505b5090506000805b838110156135bf578581815181106134e5576134e561577f565b602002602001015160c001518782815181106135035761350361577f565b602002602001018181525050612710858783815181106135255761352561577f565b60200260200101516080015161ffff1661353f9190615a04565b613549919061576b565b915061359a828783815181106135615761356161577f565b602002602001015160e0015188848151811061357f5761357f61577f565b602002602001015160c0015161359591906159b5565b613fad565b8382815181106135ac576135ac61577f565b60209081029190910101526001016134cb565b50604051632529fbc960e01b8152737e70de6d1877b3711b2beda7ba00013c7142d99390632529fbc9906135fb90899086908d90600401615c46565b60006040518083038186803b15801561361357600080fd5b505af4158015613627573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f1916820160405261364f9190810190615c7c565b909750955050505b50509193909250565b60008181527f9b48f5b32acb95b982effe269feac267eead113c4b5af14ffeb9aadac18a6e9c60208190526040822054806136ae57604051636a0eb14160e11b815260040160405180910390fd5b6132db600182615795565b60009081527f1d2f69fc9b5fe89d7414bf039e8d897c4c487c7603d80de6bcdd2868466f94766020526040902090565b6136f38282611c24565b6111d25761370b816001600160a01b03166014613fbc565b613716836020613fbc565b604051602001613727929190615cb8565b60408051601f198184030181529082905262461bcd60e51b82526111bf91600401615b0f565b613758600884615d2d565b15158061376e575061376b601082615d2d565b15155b1561378f576040516363209a7d60e11b8152600360048201526024016111bf565b600061379c60088561576b565b9050806137aa60108461576b565b146137cb576040516363209a7d60e11b8152600260048201526024016111bf565b80612e10576040516363209a7d60e11b8152600160048201526024016111bf565b6137f68282614157565b6000828152600080516020615e728339815191526020526040902061102b90826141cd565b61382582826141e2565b6000828152600080516020615e728339815191526020526040902061102b9082614256565b613852611bce565b156138705760405163184e52a160e21b815260040160405180910390fd5b6138798161426b565b50565b6111d282826137ec565b6000613890611bce565b90508082146111d2576040516303abe78360e21b815260048101829052602481018390526044016111bf565b6138c4611bce565b6138cf9060016159b5565b81146138705760405163167679d560e01b815260040160405180910390fd5b61271086111561391157604051636f004ebd60e11b815260040160405180910390fd5b61271085111561393457604051630285aacf60e31b815260040160405180910390fd5b8486111561395557604051630285aacf60e31b815260040160405180910390fd5b61271061396284866159b5565b11156139815760405163b65e4c5960e01b815260040160405180910390fd5b80158061399457506001600160401b0381115b156139b2576040516309e7727560e31b815260040160405180910390fd5b6001600160401b038211156139da5760405163e747a27f60e01b815260040160405180910390fd5b875460058901805463ffffffff60c81b19909216600160d81b61ffff8a81169190910261ffff60c81b191691909117600160c81b878316021761ffff60b81b1916600160b81b88831602178b55871669ffffffffffffffffffff1990921691909117620100006001600160401b03858116919091029190911767ffffffffffffffff60501b1916600160501b918416919091021790556040805187815260208101879052339181019190915287907f1730859048adcce16559e75a58fd609e9dbf7d34f39bcb7a45ad388dfbba0e4e9060600160405180910390a260408051858152602081018590523381830152905188917f303c8ac43d1b1f9b898ddd2915a294efa01e9b07c322d7deeb7db332b66f0410919081900360600190a26040805183815233602082015288917f72766c50f14fe492bd1281ceef0a57ad49a02b7e1042fb58723647bf38040f83910160405180910390a26040805182815233602082015288917f4d106b4a7aff347abccca2dd6855d8d59d6cf792f1fdbb272c9858433d94b328910160405180910390a25050505050505050565b6000611c1d83836142ca565b6000818311613b905781611c1d565b5090919050565b600081613ba661271085615a04565b611c1d919061576b565b60028301805467ffffffffffffffff1916426001600160401b031617905543600384015560405181815282907f9151b7f88aca05d432bb395647ef52b2ffc454e3c6afb69c95345af6b5a778c0906020015b60405180910390a2505050565b7f9b48f5b32acb95b982effe269feac267eead113c4b5af14ffeb9aadac18a6e9c613c3b8260016159b5565b600093845260209190915260409092209190915550565b613c5d846030615a04565b825114613c93578151613c71856030615a04565b6040516346b38e7960e11b8152600481019290925260248201526044016111bf565b613c9e846060615a04565b815114613cd4578051613cb2856060615a04565b604051633c11c1f760e21b8152600481019290925260248201526044016111bf565b6000613ce060306142f4565b90506000613cee60606142f4565b905060005b8681101561173557613d148584613d0b603085615a04565b6000603061430d565b613d2d8483613d24606085615a04565b6000606061430d565b7f00000000000000000000000000000000219ab540356cbb839cbe05303d7705fa6001600160a01b031663228951186801bc16d674ec800000858986613d748c8a8a614394565b6040518663ffffffff1660e01b8152600401613d939493929190615d41565b6000604051808303818588803b158015613dac57600080fd5b505af1158015613dc0573d6000803e3d6000fd5b5050505050806001019050613cf3565b8082146111d25760405163098b37e560e31b815260048101839052602481018290526044016111bf565b600060606000613e086111ee565b9050806001600160401b03811115613e2257613e226151ea565b604051908082528060200260200182016040528015613e5b57816020015b613e48614b55565b815260200190600190039081613e405790505b50915060005b81811015613ec257613e728161470f565b838281518110613e8457613e8461577f565b6020026020010181905250828181518110613ea157613ea161577f565b602002602001015160c0015184613eb891906159b5565b9350600101613e61565b50509091565b6000610dc7825490565b8154600090600160e81b900460ff166002811115613ef257613ef26153b7565b9050816002811115613f0657613f066153b7565b816002811115613f1857613f186153b7565b1461102b57816002811115613f2f57613f2f6153b7565b835460ff91909116600160e81b0260ff60e81b1982168117855560405162ffffff9182169190921617907ffd6f15fb2b48a21a60fe3d44d3c3a0433ca01e121b5124a63ec45c30ad925a1790613c029085903390615d8c565b6000613f96610f2383613660565b54630100000090046001600160a01b031692915050565b6000818310613b905781611c1d565b60606000613fcb836002615a04565b613fd69060026159b5565b6001600160401b03811115613fed57613fed6151ea565b6040519080825280601f01601f191660200182016040528015614017576020820181803683370190505b509050600360fc1b816000815181106140325761403261577f565b60200101906001600160f81b031916908160001a905350600f60fb1b816001815181106140615761406161577f565b60200101906001600160f81b031916908160001a9053506000614085846002615a04565b6140909060016159b5565b90505b6001811115614108576f181899199a1a9b1b9c1cb0b131b232b360811b85600f16601081106140c4576140c461577f565b1a60f81b8282815181106140da576140da61577f565b60200101906001600160f81b031916908160001a90535060049490941c9361410181615db2565b9050614093565b508315611c1d5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e7460448201526064016111bf565b6141618282611c24565b6111d2576000828152600080516020615e92833981519152602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b6000611c1d836001600160a01b038416614826565b6141ec8282611c24565b156111d2576000828152600080516020615e92833981519152602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b6000611c1d836001600160a01b038416614875565b6142947f4dd0f6662ba1d6b081f08b350f5e9a6a7b15cf586926ba66f753594928fa64a6829055565b6040518181527ffddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb9060200160405180910390a150565b60008260000182815481106142e1576142e161577f565b9060005260206000200154905092915050565b60408051828152603f92810192909201601f1916905290565b845161431982856159b5565b111580156143315750835161432e82846159b5565b11155b61437d5760405162461bcd60e51b815260206004820152601960248201527f42595445535f41525241595f4f55545f4f465f424f554e44530000000000000060448201526064016111bf565b602083860181019083860101611735828285614968565b6000806143a160406142f4565b905060006143b96143b460406060615795565b6142f4565b90506143ca8483600080604061430d565b6143e38482604060006143de826060615795565b61430d565b6000600286600060801b6040516020016143fe929190615dc9565b60408051601f198184030181529082905261441891615e01565b602060405180830381855afa158015614435573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906144589190615726565b90506000600280856040516020016144709190615e01565b60408051601f198184030181529082905261448a91615e01565b602060405180830381855afa1580156144a7573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906144ca9190615726565b6040516002906144e1908790600090602001615e13565b60408051601f19818403018152908290526144fb91615e01565b602060405180830381855afa158015614518573d6000803e3d6000fd5b5050506040513d601f19601f8201168201806040525081019061453b9190615726565b60408051602081019390935282015260600160408051601f198184030181529082905261456791615e01565b602060405180830381855afa158015614584573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906145a79190615726565b9050600280838a6040516020016145bf929190615e35565b60408051601f19818403018152908290526145d991615e01565b602060405180830381855afa1580156145f6573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906146199190615726565b60408051634059730760d81b60208201526000602882015290810184905260029060600160408051601f198184030181529082905261465791615e01565b602060405180830381855afa158015614674573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906146979190615726565b60408051602081019390935282015260600160408051601f19818403018152908290526146c391615e01565b602060405180830381855afa1580156146e0573d6000803e3d6000fd5b5050506040513d601f19601f820116820180604052508101906147039190615726565b98975050505050505050565b614717614b55565b6000614722836136b9565b80546001600160a01b036301000000820416845262ffffff8116602085015261ffff600160b81b820481166040860152600160c81b820481166060860152600160d81b820416608085015290915060ff600160e81b90910416600281111561478c5761478c6153b7565b8260a0019060028111156147a2576147a26153b7565b908160028111156147b5576147b56153b7565b8152505060008060006147cb8560000151613391565b9194509250905060008560a0015160028111156147ea576147ea6153b7565b146147f65760006147f8565b805b60e0860152600484015461480d908490613b81565b6148179083615795565b60c08601525092949350505050565b600081815260018301602052604081205461486d57508154600181810184556000848152602080822090930184905584548482528286019093526040902091909155610dc7565b506000610dc7565b6000818152600183016020526040812054801561495e576000614899600183615795565b85549091506000906148ad90600190615795565b90508181146149125760008660000182815481106148cd576148cd61577f565b90600052602060002001549050808760000184815481106148f0576148f061577f565b6000918252602080832090910192909255918252600188019052604090208390555b855486908061492357614923615e5b565b600190038181906000526020600020016000905590558560010160008681526020019081526020016000206000905560019350505050610dc7565b6000915050610dc7565b5b601f811115614989578251825260209283019290910190601f1901614969565b801561102b5782518251600019600160086020869003021b01908116901991909116178252505050565b604080516101a08101825260008082526020820181905291810182905260608082018390526080820183905260a0820183905260c082015260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081019190915290565b60405180608001604052806000815260200160008152602001614a3f6149b3565b8152602001614a6860405180606001604052806000815260200160008152602001600081525090565b905290565b828054614a7990615839565b90600052602060002090601f016020900481019282614a9b5760008555614ae1565b82601f10614ab45782800160ff19823516178555614ae1565b82800160010185558215614ae1579182015b82811115614ae1578235825591602001919060010190614ac6565b50614aed929150614b9a565b5090565b60405180610100016040528060008152602001600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6040805160608101825260008082526020820152908101614a68614af1565b604080516101008101825260008082526020820181905291810182905260608101829052608081018290529060a0820190815260200160008152602001600081525090565b5b80821115614aed5760008155600101614b9b565b600060208284031215614bc157600080fd5b81356001600160e01b031981168114611c1d57600080fd5b600060208284031215614beb57600080fd5b5035919050565b81518152602080830151908201526040808301519082015260608101610dc7565b60008060408385031215614c2657600080fd5b50508035926020909101359150565b60008083601f840112614c4757600080fd5b5081356001600160401b03811115614c5e57600080fd5b602083019150836020828501011115614c7657600080fd5b9250929050565b600080600080600060608688031215614c9557600080fd5b8535945060208601356001600160401b0380821115614cb357600080fd5b614cbf89838a01614c35565b90965094506040880135915080821115614cd857600080fd5b50614ce588828901614c35565b969995985093965092949392505050565b80356001600160a01b0381168114614d0d57600080fd5b919050565b60008060408385031215614d2557600080fd5b82359150614d3560208401614cf6565b90509250929050565b805182526020810151602083015260408101516040830152606081015160608301526080810151608083015260a081015160a083015260c081015160c083015260e081015160e08301525050565b602080825282518282018190526000919060409081850190868401855b82811015614de55781518051855286810151151587860152850151614dd086860182614d3e565b50610140939093019290850190600101614da9565b5091979650505050505050565b60008060008060008060a08789031215614e0b57600080fd5b86359550602087013594506040870135935060608701356001600160401b03811115614e3657600080fd5b614e4289828a01614c35565b979a9699509497949695608090950135949350505050565b60005b83811015614e75578181015183820152602001614e5d565b83811115612c935750506000910152565b60008151808452614e9e816020860160208601614e5a565b601f01601f19169290920160200192915050565b805162ffffff16825260006101a06020830151614eda60208601826001600160a01b03169052565b506040830151614ef0604086018261ffff169052565b506060830151614f06606086018261ffff169052565b506080830151614f1c608086018261ffff169052565b5060a0830151614f3160a086018260ff169052565b5060c08301518160c0860152614f4982860182614e86565b91505060e0830151614f6660e08601826001600160401b03169052565b50610100838101519085015261012080840151908501526101408084015161ffff1690850152610160808401516001600160401b0390811691860191909152610180938401511692909301919091525090565b60006020808301818452808551808352604092508286019150828160051b87010184880160005b8381101561505057603f19898403018552815160c081518552888201518986015287820151818987015261501682870182614eb2565b60609384015180518886015260208101516080890152604081015160a0890152939092509050509588019593505090860190600101614fe0565b509098975050505050505050565b60008060006060848603121561507357600080fd5b61507c84614cf6565b925061508a60208501614cf6565b9150604084013590509250925092565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156150ef57603f198886030184526150dd858351614eb2565b945092850192908501906001016150c1565b5092979650505050505050565b60008083601f84011261510e57600080fd5b5081356001600160401b0381111561512557600080fd5b6020830191508360208260051b8501011115614c7657600080fd5b6000806000806060858703121561515657600080fd5b84356001600160401b0381111561516c57600080fd5b615178878288016150fc565b90989097506020870135966040013595509350505050565b600080600080600080600060e0888a0312156151ab57600080fd5b505085359760208701359750604087013596606081013596506080810135955060a0810135945060c0013592509050565b801515811461387957600080fd5b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f191681016001600160401b0381118282101715615228576152286151ea565b604052919050565b60008060008084860360e081121561524757600080fd5b85359450602086013593506040860135615260816151dc565b92506080605f198201121561527457600080fd5b50604051608081018181106001600160401b0382111715615297576152976151ea565b8060405250606086013581526080860135602082015260a0860135604082015260c086013560608201528091505092959194509250565b60006001600160401b038211156152e7576152e76151ea565b5060051b60200190565b600082601f83011261530257600080fd5b81356020615317615312836152ce565b615200565b82815260059290921b8401810191818101908684111561533657600080fd5b8286015b8481101561249d578035835291830191830161533a565b60006020828403121561536357600080fd5b81356001600160401b0381111561537957600080fd5b6132db848285016152f1565b6000806000806080858703121561539b57600080fd5b5050823594602084013594506040840135936060013592509050565b634e487b7160e01b600052602160045260246000fd5b600381106153eb57634e487b7160e01b600052602160045260246000fd5b9052565b60208101610dc782846153cd565b60008060008060008060008060006101008a8c03121561541c57600080fd5b89356001600160401b0381111561543257600080fd5b61543e8c828d01614c35565b909a509850615451905060208b01614cf6565b989b979a509798604081013598506060810135976080820135975060a0820135965060c0820135955060e0909101359350915050565b6000806000806060858703121561549d57600080fd5b843593506020850135925060408501356001600160401b038111156154c157600080fd5b6154cd87828801614c35565b95989497509550505050565b6101008101610dc78284614d3e565b600080600080604085870312156154fe57600080fd5b84356001600160401b038082111561551557600080fd5b615521888389016150fc565b9096509450602087013591508082111561553a57600080fd5b506154cd878288016150fc565b600081518084526020808501945080840160005b838110156155775781518752958201959082019060010161555b565b509495945050505050565b60a0808252865190820181905260009060209060c0840190828a01845b828110156155c45781516001600160a01b03168452928401929084019060010161559f565b505050838103828501526155d88189615547565b8481036040860152875180825283890192509083019060005b818110156156165783516001600160601b0316835292840192918401916001016155f1565b50506001600160601b03871660608601529250615631915050565b8260808301529695505050505050565b602081526000611c1d6020830184614eb2565b8281526040602082015260006132db6040830184615547565b6000806040838503121561568057600080fd5b8235915060208301356003811061569657600080fd5b809150509250929050565b600080604083850312156156b457600080fd5b8235915060208301356001600160401b038111156156d157600080fd5b6156dd858286016152f1565b9150509250929050565b602081526000611c1d6020830184615547565b60008060006060848603121561570f57600080fd5b505081359360208301359350604090920135919050565b60006020828403121561573857600080fd5b5051919050565b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b60008261577a5761577a61573f565b500490565b634e487b7160e01b600052603260045260246000fd5b6000828210156157a7576157a7615755565b500390565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b6040815260006157e96040830186886157ac565b82810360208401526157fc8185876157ac565b979650505050505050565b8581528460208201526080604082015260006158276080830185876157ac565b90508260608301529695505050505050565b600181811c9082168061584d57607f821691505b60208210811415611bc857634e487b7160e01b600052602260045260246000fd5b60008235605e1983360301811261588457600080fd5b9190910192915050565b6000808335601e198436030181126158a557600080fd5b8301803591506001600160401b038211156158bf57600080fd5b602001915036819003821315614c7657600080fd5b8581526080602082015260006158ee6080830186886157ac565b604083019490945250606001529392505050565b6020815260006132db6020830184866157ac565b600060001982141561592a5761592a615755565b5060010190565b600080600080600080600080610100898b03121561594e57600080fd5b505086516020880151604089015160608a015160808b015160a08c015160c08d015160e0909d0151959e949d50929b919a50985090965094509092509050565b600062ffffff8083168185168083038211156159ac576159ac615755565b01949350505050565b600082198211156159c8576159c8615755565b500190565b600060018060a01b038087168352606060208401526159f06060840186886157ac565b915080841660408401525095945050505050565b6000816000190483118215151615615a1e57615a1e615755565b500290565b838152604060208201526000611cab6040830184866157ac565b600082601f830112615a4e57600080fd5b81516001600160401b03811115615a6757615a676151ea565b615a7a601f8201601f1916602001615200565b818152846020838601011115615a8f57600080fd5b6132db826020830160208701614e5a565b60008060408385031215615ab357600080fd5b82516001600160401b0380821115615aca57600080fd5b615ad686838701615a3d565b93506020850151915080821115615aec57600080fd5b506156dd85828601615a3d565b634e487b7160e01b600052600160045260246000fd5b602081526000611c1d6020830184614e86565b60006001600160601b038083168185168083038211156159ac576159ac615755565b600060208284031215615b5657600080fd5b8151611c1d816151dc565b600082601f830112615b7257600080fd5b81516020615b82615312836152ce565b82815260059290921b84018101918181019086841115615ba157600080fd5b8286015b8481101561249d5780518352918301918301615ba5565b600060208284031215615bce57600080fd5b81516001600160401b03811115615be457600080fd5b6132db84828501615b61565b60006001600160601b0383811690831681811015615c1057615c10615755565b039392505050565b600080600060608486031215615c2d57600080fd5b8351925060208401519150604084015190509250925092565b606081526000615c596060830186615547565b8281036020840152615c6b8186615547565b915050826040830152949350505050565b60008060408385031215615c8f57600080fd5b8251915060208301516001600160401b03811115615cac57600080fd5b6156dd85828601615b61565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351615cf0816017850160208801614e5a565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351615d21816028840160208801614e5a565b01602801949350505050565b600082615d3c57615d3c61573f565b500690565b608081526000615d546080830187614e86565b8281036020840152615d668187614e86565b90508281036040840152615d7a8186614e86565b91505082606083015295945050505050565b60408101615d9a82856153cd565b6001600160a01b039290921660209190910152919050565b600081615dc157615dc1615755565b506000190190565b60008351615ddb818460208801614e5a565b6fffffffffffffffffffffffffffffffff19939093169190920190815260100192915050565b60008251615884818460208701614e5a565b60008351615e25818460208801614e5a565b9190910191825250602001919050565b82815260008251615e4d816020850160208701614e5a565b919091016020019392505050565b634e487b7160e01b600052603160045260246000fdfe8f8c450dae5029cd48cd91dd9db65da48fb742893edfc7941250f6721d93cbbe9a627a5d4aa7c17f87ff26e3fe9a42c2b6c559e8b41a42282d0ecebb17c0e4d3c23292b191d95d2a7dd94fc6436eb44338fda9e1307d9394fd27c28157c1b33c3105bcbf19d4417b73ae0e58d508a65ecf75665e46c2622d8521732de6080c48a26469706673582212201f124ae5d922775405c86cfba25065f612acf97234fdbcd17a36237861d0ff3464736f6c63430008090033",
}

// StakingrouterABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingrouterMetaData.ABI instead.
var StakingrouterABI = StakingrouterMetaData.ABI

// StakingrouterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingrouterMetaData.Bin instead.
var StakingrouterBin = StakingrouterMetaData.Bin

// DeployStakingrouter deploys a new Ethereum contract, binding an instance of Stakingrouter to it.
func DeployStakingrouter(auth *bind.TransactOpts, backend bind.ContractBackend, _depositContract common.Address) (common.Address, *types.Transaction, *Stakingrouter, error) {
	parsed, err := StakingrouterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingrouterBin), backend, _depositContract)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Stakingrouter{StakingrouterCaller: StakingrouterCaller{contract: contract}, StakingrouterTransactor: StakingrouterTransactor{contract: contract}, StakingrouterFilterer: StakingrouterFilterer{contract: contract}}, nil
}

// Stakingrouter is an auto generated Go binding around an Ethereum contract.
type Stakingrouter struct {
	StakingrouterCaller     // Read-only binding to the contract
	StakingrouterTransactor // Write-only binding to the contract
	StakingrouterFilterer   // Log filterer for contract events
}

// StakingrouterCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingrouterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingrouterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingrouterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingrouterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingrouterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingrouterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingrouterSession struct {
	Contract     *Stakingrouter    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingrouterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingrouterCallerSession struct {
	Contract *StakingrouterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StakingrouterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingrouterTransactorSession struct {
	Contract     *StakingrouterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StakingrouterRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingrouterRaw struct {
	Contract *Stakingrouter // Generic contract binding to access the raw methods on
}

// StakingrouterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingrouterCallerRaw struct {
	Contract *StakingrouterCaller // Generic read-only contract binding to access the raw methods on
}

// StakingrouterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingrouterTransactorRaw struct {
	Contract *StakingrouterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakingrouter creates a new instance of Stakingrouter, bound to a specific deployed contract.
func NewStakingrouter(address common.Address, backend bind.ContractBackend) (*Stakingrouter, error) {
	contract, err := bindStakingrouter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Stakingrouter{StakingrouterCaller: StakingrouterCaller{contract: contract}, StakingrouterTransactor: StakingrouterTransactor{contract: contract}, StakingrouterFilterer: StakingrouterFilterer{contract: contract}}, nil
}

// NewStakingrouterCaller creates a new read-only instance of Stakingrouter, bound to a specific deployed contract.
func NewStakingrouterCaller(address common.Address, caller bind.ContractCaller) (*StakingrouterCaller, error) {
	contract, err := bindStakingrouter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingrouterCaller{contract: contract}, nil
}

// NewStakingrouterTransactor creates a new write-only instance of Stakingrouter, bound to a specific deployed contract.
func NewStakingrouterTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingrouterTransactor, error) {
	contract, err := bindStakingrouter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingrouterTransactor{contract: contract}, nil
}

// NewStakingrouterFilterer creates a new log filterer instance of Stakingrouter, bound to a specific deployed contract.
func NewStakingrouterFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingrouterFilterer, error) {
	contract, err := bindStakingrouter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingrouterFilterer{contract: contract}, nil
}

// bindStakingrouter binds a generic wrapper to an already deployed contract.
func bindStakingrouter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingrouterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakingrouter *StakingrouterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakingrouter.Contract.StakingrouterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakingrouter *StakingrouterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrouter.Contract.StakingrouterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakingrouter *StakingrouterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakingrouter.Contract.StakingrouterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Stakingrouter *StakingrouterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Stakingrouter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Stakingrouter *StakingrouterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrouter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Stakingrouter *StakingrouterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Stakingrouter.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.DEFAULTADMINROLE(&_Stakingrouter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.DEFAULTADMINROLE(&_Stakingrouter.CallOpts)
}

// DEPOSITCONTRACT is a free data retrieval call binding the contract method 0x6b96736b.
//
// Solidity: function DEPOSIT_CONTRACT() view returns(address)
func (_Stakingrouter *StakingrouterCaller) DEPOSITCONTRACT(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "DEPOSIT_CONTRACT")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DEPOSITCONTRACT is a free data retrieval call binding the contract method 0x6b96736b.
//
// Solidity: function DEPOSIT_CONTRACT() view returns(address)
func (_Stakingrouter *StakingrouterSession) DEPOSITCONTRACT() (common.Address, error) {
	return _Stakingrouter.Contract.DEPOSITCONTRACT(&_Stakingrouter.CallOpts)
}

// DEPOSITCONTRACT is a free data retrieval call binding the contract method 0x6b96736b.
//
// Solidity: function DEPOSIT_CONTRACT() view returns(address)
func (_Stakingrouter *StakingrouterCallerSession) DEPOSITCONTRACT() (common.Address, error) {
	return _Stakingrouter.Contract.DEPOSITCONTRACT(&_Stakingrouter.CallOpts)
}

// FEEPRECISIONPOINTS is a free data retrieval call binding the contract method 0x7c8da51c.
//
// Solidity: function FEE_PRECISION_POINTS() view returns(uint256)
func (_Stakingrouter *StakingrouterCaller) FEEPRECISIONPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "FEE_PRECISION_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FEEPRECISIONPOINTS is a free data retrieval call binding the contract method 0x7c8da51c.
//
// Solidity: function FEE_PRECISION_POINTS() view returns(uint256)
func (_Stakingrouter *StakingrouterSession) FEEPRECISIONPOINTS() (*big.Int, error) {
	return _Stakingrouter.Contract.FEEPRECISIONPOINTS(&_Stakingrouter.CallOpts)
}

// FEEPRECISIONPOINTS is a free data retrieval call binding the contract method 0x7c8da51c.
//
// Solidity: function FEE_PRECISION_POINTS() view returns(uint256)
func (_Stakingrouter *StakingrouterCallerSession) FEEPRECISIONPOINTS() (*big.Int, error) {
	return _Stakingrouter.Contract.FEEPRECISIONPOINTS(&_Stakingrouter.CallOpts)
}

// MANAGEWITHDRAWALCREDENTIALSROLE is a free data retrieval call binding the contract method 0x7a74884d.
//
// Solidity: function MANAGE_WITHDRAWAL_CREDENTIALS_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCaller) MANAGEWITHDRAWALCREDENTIALSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "MANAGE_WITHDRAWAL_CREDENTIALS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEWITHDRAWALCREDENTIALSROLE is a free data retrieval call binding the contract method 0x7a74884d.
//
// Solidity: function MANAGE_WITHDRAWAL_CREDENTIALS_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterSession) MANAGEWITHDRAWALCREDENTIALSROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.MANAGEWITHDRAWALCREDENTIALSROLE(&_Stakingrouter.CallOpts)
}

// MANAGEWITHDRAWALCREDENTIALSROLE is a free data retrieval call binding the contract method 0x7a74884d.
//
// Solidity: function MANAGE_WITHDRAWAL_CREDENTIALS_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCallerSession) MANAGEWITHDRAWALCREDENTIALSROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.MANAGEWITHDRAWALCREDENTIALSROLE(&_Stakingrouter.CallOpts)
}

// MAXSTAKINGMODULESCOUNT is a free data retrieval call binding the contract method 0x4b3a1cb7.
//
// Solidity: function MAX_STAKING_MODULES_COUNT() view returns(uint256)
func (_Stakingrouter *StakingrouterCaller) MAXSTAKINGMODULESCOUNT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "MAX_STAKING_MODULES_COUNT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTAKINGMODULESCOUNT is a free data retrieval call binding the contract method 0x4b3a1cb7.
//
// Solidity: function MAX_STAKING_MODULES_COUNT() view returns(uint256)
func (_Stakingrouter *StakingrouterSession) MAXSTAKINGMODULESCOUNT() (*big.Int, error) {
	return _Stakingrouter.Contract.MAXSTAKINGMODULESCOUNT(&_Stakingrouter.CallOpts)
}

// MAXSTAKINGMODULESCOUNT is a free data retrieval call binding the contract method 0x4b3a1cb7.
//
// Solidity: function MAX_STAKING_MODULES_COUNT() view returns(uint256)
func (_Stakingrouter *StakingrouterCallerSession) MAXSTAKINGMODULESCOUNT() (*big.Int, error) {
	return _Stakingrouter.Contract.MAXSTAKINGMODULESCOUNT(&_Stakingrouter.CallOpts)
}

// MAXSTAKINGMODULENAMELENGTH is a free data retrieval call binding the contract method 0x9b75b4ef.
//
// Solidity: function MAX_STAKING_MODULE_NAME_LENGTH() view returns(uint256)
func (_Stakingrouter *StakingrouterCaller) MAXSTAKINGMODULENAMELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "MAX_STAKING_MODULE_NAME_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSTAKINGMODULENAMELENGTH is a free data retrieval call binding the contract method 0x9b75b4ef.
//
// Solidity: function MAX_STAKING_MODULE_NAME_LENGTH() view returns(uint256)
func (_Stakingrouter *StakingrouterSession) MAXSTAKINGMODULENAMELENGTH() (*big.Int, error) {
	return _Stakingrouter.Contract.MAXSTAKINGMODULENAMELENGTH(&_Stakingrouter.CallOpts)
}

// MAXSTAKINGMODULENAMELENGTH is a free data retrieval call binding the contract method 0x9b75b4ef.
//
// Solidity: function MAX_STAKING_MODULE_NAME_LENGTH() view returns(uint256)
func (_Stakingrouter *StakingrouterCallerSession) MAXSTAKINGMODULENAMELENGTH() (*big.Int, error) {
	return _Stakingrouter.Contract.MAXSTAKINGMODULENAMELENGTH(&_Stakingrouter.CallOpts)
}

// REPORTEXITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xc445ea75.
//
// Solidity: function REPORT_EXITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCaller) REPORTEXITEDVALIDATORSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "REPORT_EXITED_VALIDATORS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTEXITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xc445ea75.
//
// Solidity: function REPORT_EXITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterSession) REPORTEXITEDVALIDATORSROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.REPORTEXITEDVALIDATORSROLE(&_Stakingrouter.CallOpts)
}

// REPORTEXITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0xc445ea75.
//
// Solidity: function REPORT_EXITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCallerSession) REPORTEXITEDVALIDATORSROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.REPORTEXITEDVALIDATORSROLE(&_Stakingrouter.CallOpts)
}

// REPORTREWARDSMINTEDROLE is a free data retrieval call binding the contract method 0x1d1b9d3c.
//
// Solidity: function REPORT_REWARDS_MINTED_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCaller) REPORTREWARDSMINTEDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "REPORT_REWARDS_MINTED_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTREWARDSMINTEDROLE is a free data retrieval call binding the contract method 0x1d1b9d3c.
//
// Solidity: function REPORT_REWARDS_MINTED_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterSession) REPORTREWARDSMINTEDROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.REPORTREWARDSMINTEDROLE(&_Stakingrouter.CallOpts)
}

// REPORTREWARDSMINTEDROLE is a free data retrieval call binding the contract method 0x1d1b9d3c.
//
// Solidity: function REPORT_REWARDS_MINTED_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCallerSession) REPORTREWARDSMINTEDROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.REPORTREWARDSMINTEDROLE(&_Stakingrouter.CallOpts)
}

// REPORTVALIDATOREXITINGSTATUSROLE is a free data retrieval call binding the contract method 0x0fb31c84.
//
// Solidity: function REPORT_VALIDATOR_EXITING_STATUS_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCaller) REPORTVALIDATOREXITINGSTATUSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "REPORT_VALIDATOR_EXITING_STATUS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTVALIDATOREXITINGSTATUSROLE is a free data retrieval call binding the contract method 0x0fb31c84.
//
// Solidity: function REPORT_VALIDATOR_EXITING_STATUS_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterSession) REPORTVALIDATOREXITINGSTATUSROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.REPORTVALIDATOREXITINGSTATUSROLE(&_Stakingrouter.CallOpts)
}

// REPORTVALIDATOREXITINGSTATUSROLE is a free data retrieval call binding the contract method 0x0fb31c84.
//
// Solidity: function REPORT_VALIDATOR_EXITING_STATUS_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCallerSession) REPORTVALIDATOREXITINGSTATUSROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.REPORTVALIDATOREXITINGSTATUSROLE(&_Stakingrouter.CallOpts)
}

// REPORTVALIDATOREXITTRIGGEREDROLE is a free data retrieval call binding the contract method 0xe39fdbe9.
//
// Solidity: function REPORT_VALIDATOR_EXIT_TRIGGERED_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCaller) REPORTVALIDATOREXITTRIGGEREDROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "REPORT_VALIDATOR_EXIT_TRIGGERED_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTVALIDATOREXITTRIGGEREDROLE is a free data retrieval call binding the contract method 0xe39fdbe9.
//
// Solidity: function REPORT_VALIDATOR_EXIT_TRIGGERED_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterSession) REPORTVALIDATOREXITTRIGGEREDROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.REPORTVALIDATOREXITTRIGGEREDROLE(&_Stakingrouter.CallOpts)
}

// REPORTVALIDATOREXITTRIGGEREDROLE is a free data retrieval call binding the contract method 0xe39fdbe9.
//
// Solidity: function REPORT_VALIDATOR_EXIT_TRIGGERED_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCallerSession) REPORTVALIDATOREXITTRIGGEREDROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.REPORTVALIDATOREXITTRIGGEREDROLE(&_Stakingrouter.CallOpts)
}

// STAKINGMODULEMANAGEROLE is a free data retrieval call binding the contract method 0xe016e6f7.
//
// Solidity: function STAKING_MODULE_MANAGE_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCaller) STAKINGMODULEMANAGEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "STAKING_MODULE_MANAGE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMODULEMANAGEROLE is a free data retrieval call binding the contract method 0xe016e6f7.
//
// Solidity: function STAKING_MODULE_MANAGE_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterSession) STAKINGMODULEMANAGEROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.STAKINGMODULEMANAGEROLE(&_Stakingrouter.CallOpts)
}

// STAKINGMODULEMANAGEROLE is a free data retrieval call binding the contract method 0xe016e6f7.
//
// Solidity: function STAKING_MODULE_MANAGE_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCallerSession) STAKINGMODULEMANAGEROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.STAKINGMODULEMANAGEROLE(&_Stakingrouter.CallOpts)
}

// STAKINGMODULEUNVETTINGROLE is a free data retrieval call binding the contract method 0x909c01de.
//
// Solidity: function STAKING_MODULE_UNVETTING_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCaller) STAKINGMODULEUNVETTINGROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "STAKING_MODULE_UNVETTING_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGMODULEUNVETTINGROLE is a free data retrieval call binding the contract method 0x909c01de.
//
// Solidity: function STAKING_MODULE_UNVETTING_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterSession) STAKINGMODULEUNVETTINGROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.STAKINGMODULEUNVETTINGROLE(&_Stakingrouter.CallOpts)
}

// STAKINGMODULEUNVETTINGROLE is a free data retrieval call binding the contract method 0x909c01de.
//
// Solidity: function STAKING_MODULE_UNVETTING_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCallerSession) STAKINGMODULEUNVETTINGROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.STAKINGMODULEUNVETTINGROLE(&_Stakingrouter.CallOpts)
}

// TOTALBASISPOINTS is a free data retrieval call binding the contract method 0x271662ec.
//
// Solidity: function TOTAL_BASIS_POINTS() view returns(uint256)
func (_Stakingrouter *StakingrouterCaller) TOTALBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "TOTAL_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOTALBASISPOINTS is a free data retrieval call binding the contract method 0x271662ec.
//
// Solidity: function TOTAL_BASIS_POINTS() view returns(uint256)
func (_Stakingrouter *StakingrouterSession) TOTALBASISPOINTS() (*big.Int, error) {
	return _Stakingrouter.Contract.TOTALBASISPOINTS(&_Stakingrouter.CallOpts)
}

// TOTALBASISPOINTS is a free data retrieval call binding the contract method 0x271662ec.
//
// Solidity: function TOTAL_BASIS_POINTS() view returns(uint256)
func (_Stakingrouter *StakingrouterCallerSession) TOTALBASISPOINTS() (*big.Int, error) {
	return _Stakingrouter.Contract.TOTALBASISPOINTS(&_Stakingrouter.CallOpts)
}

// UNSAFESETEXITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0x1565d2f2.
//
// Solidity: function UNSAFE_SET_EXITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCaller) UNSAFESETEXITEDVALIDATORSROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "UNSAFE_SET_EXITED_VALIDATORS_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNSAFESETEXITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0x1565d2f2.
//
// Solidity: function UNSAFE_SET_EXITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterSession) UNSAFESETEXITEDVALIDATORSROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.UNSAFESETEXITEDVALIDATORSROLE(&_Stakingrouter.CallOpts)
}

// UNSAFESETEXITEDVALIDATORSROLE is a free data retrieval call binding the contract method 0x1565d2f2.
//
// Solidity: function UNSAFE_SET_EXITED_VALIDATORS_ROLE() view returns(bytes32)
func (_Stakingrouter *StakingrouterCallerSession) UNSAFESETEXITEDVALIDATORSROLE() ([32]byte, error) {
	return _Stakingrouter.Contract.UNSAFESETEXITEDVALIDATORSROLE(&_Stakingrouter.CallOpts)
}

// GetAllNodeOperatorDigests is a free data retrieval call binding the contract method 0x3240a322.
//
// Solidity: function getAllNodeOperatorDigests(uint256 _stakingModuleId) view returns((uint256,bool,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Stakingrouter *StakingrouterCaller) GetAllNodeOperatorDigests(opts *bind.CallOpts, _stakingModuleId *big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getAllNodeOperatorDigests", _stakingModuleId)

	if err != nil {
		return *new([]StakingRouterNodeOperatorDigest), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingRouterNodeOperatorDigest)).(*[]StakingRouterNodeOperatorDigest)

	return out0, err

}

// GetAllNodeOperatorDigests is a free data retrieval call binding the contract method 0x3240a322.
//
// Solidity: function getAllNodeOperatorDigests(uint256 _stakingModuleId) view returns((uint256,bool,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Stakingrouter *StakingrouterSession) GetAllNodeOperatorDigests(_stakingModuleId *big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	return _Stakingrouter.Contract.GetAllNodeOperatorDigests(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetAllNodeOperatorDigests is a free data retrieval call binding the contract method 0x3240a322.
//
// Solidity: function getAllNodeOperatorDigests(uint256 _stakingModuleId) view returns((uint256,bool,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Stakingrouter *StakingrouterCallerSession) GetAllNodeOperatorDigests(_stakingModuleId *big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	return _Stakingrouter.Contract.GetAllNodeOperatorDigests(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetAllStakingModuleDigests is a free data retrieval call binding the contract method 0x57993b85.
//
// Solidity: function getAllStakingModuleDigests() view returns((uint256,uint256,(uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256,uint16,uint64,uint64),(uint256,uint256,uint256))[])
func (_Stakingrouter *StakingrouterCaller) GetAllStakingModuleDigests(opts *bind.CallOpts) ([]StakingRouterStakingModuleDigest, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getAllStakingModuleDigests")

	if err != nil {
		return *new([]StakingRouterStakingModuleDigest), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingRouterStakingModuleDigest)).(*[]StakingRouterStakingModuleDigest)

	return out0, err

}

// GetAllStakingModuleDigests is a free data retrieval call binding the contract method 0x57993b85.
//
// Solidity: function getAllStakingModuleDigests() view returns((uint256,uint256,(uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256,uint16,uint64,uint64),(uint256,uint256,uint256))[])
func (_Stakingrouter *StakingrouterSession) GetAllStakingModuleDigests() ([]StakingRouterStakingModuleDigest, error) {
	return _Stakingrouter.Contract.GetAllStakingModuleDigests(&_Stakingrouter.CallOpts)
}

// GetAllStakingModuleDigests is a free data retrieval call binding the contract method 0x57993b85.
//
// Solidity: function getAllStakingModuleDigests() view returns((uint256,uint256,(uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256,uint16,uint64,uint64),(uint256,uint256,uint256))[])
func (_Stakingrouter *StakingrouterCallerSession) GetAllStakingModuleDigests() ([]StakingRouterStakingModuleDigest, error) {
	return _Stakingrouter.Contract.GetAllStakingModuleDigests(&_Stakingrouter.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Stakingrouter *StakingrouterCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Stakingrouter *StakingrouterSession) GetContractVersion() (*big.Int, error) {
	return _Stakingrouter.Contract.GetContractVersion(&_Stakingrouter.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Stakingrouter *StakingrouterCallerSession) GetContractVersion() (*big.Int, error) {
	return _Stakingrouter.Contract.GetContractVersion(&_Stakingrouter.CallOpts)
}

// GetDepositsAllocation is a free data retrieval call binding the contract method 0xc82b1bb1.
//
// Solidity: function getDepositsAllocation(uint256 _depositsCount) view returns(uint256 allocated, uint256[] allocations)
func (_Stakingrouter *StakingrouterCaller) GetDepositsAllocation(opts *bind.CallOpts, _depositsCount *big.Int) (struct {
	Allocated   *big.Int
	Allocations []*big.Int
}, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getDepositsAllocation", _depositsCount)

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
func (_Stakingrouter *StakingrouterSession) GetDepositsAllocation(_depositsCount *big.Int) (struct {
	Allocated   *big.Int
	Allocations []*big.Int
}, error) {
	return _Stakingrouter.Contract.GetDepositsAllocation(&_Stakingrouter.CallOpts, _depositsCount)
}

// GetDepositsAllocation is a free data retrieval call binding the contract method 0xc82b1bb1.
//
// Solidity: function getDepositsAllocation(uint256 _depositsCount) view returns(uint256 allocated, uint256[] allocations)
func (_Stakingrouter *StakingrouterCallerSession) GetDepositsAllocation(_depositsCount *big.Int) (struct {
	Allocated   *big.Int
	Allocations []*big.Int
}, error) {
	return _Stakingrouter.Contract.GetDepositsAllocation(&_Stakingrouter.CallOpts, _depositsCount)
}

// GetLido is a free data retrieval call binding the contract method 0x6a516b47.
//
// Solidity: function getLido() view returns(address)
func (_Stakingrouter *StakingrouterCaller) GetLido(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getLido")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLido is a free data retrieval call binding the contract method 0x6a516b47.
//
// Solidity: function getLido() view returns(address)
func (_Stakingrouter *StakingrouterSession) GetLido() (common.Address, error) {
	return _Stakingrouter.Contract.GetLido(&_Stakingrouter.CallOpts)
}

// GetLido is a free data retrieval call binding the contract method 0x6a516b47.
//
// Solidity: function getLido() view returns(address)
func (_Stakingrouter *StakingrouterCallerSession) GetLido() (common.Address, error) {
	return _Stakingrouter.Contract.GetLido(&_Stakingrouter.CallOpts)
}

// GetNodeOperatorDigests is a free data retrieval call binding the contract method 0xf07ff28a.
//
// Solidity: function getNodeOperatorDigests(uint256 _stakingModuleId, uint256[] _nodeOperatorIds) view returns((uint256,bool,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[] digests)
func (_Stakingrouter *StakingrouterCaller) GetNodeOperatorDigests(opts *bind.CallOpts, _stakingModuleId *big.Int, _nodeOperatorIds []*big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getNodeOperatorDigests", _stakingModuleId, _nodeOperatorIds)

	if err != nil {
		return *new([]StakingRouterNodeOperatorDigest), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingRouterNodeOperatorDigest)).(*[]StakingRouterNodeOperatorDigest)

	return out0, err

}

// GetNodeOperatorDigests is a free data retrieval call binding the contract method 0xf07ff28a.
//
// Solidity: function getNodeOperatorDigests(uint256 _stakingModuleId, uint256[] _nodeOperatorIds) view returns((uint256,bool,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[] digests)
func (_Stakingrouter *StakingrouterSession) GetNodeOperatorDigests(_stakingModuleId *big.Int, _nodeOperatorIds []*big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	return _Stakingrouter.Contract.GetNodeOperatorDigests(&_Stakingrouter.CallOpts, _stakingModuleId, _nodeOperatorIds)
}

// GetNodeOperatorDigests is a free data retrieval call binding the contract method 0xf07ff28a.
//
// Solidity: function getNodeOperatorDigests(uint256 _stakingModuleId, uint256[] _nodeOperatorIds) view returns((uint256,bool,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[] digests)
func (_Stakingrouter *StakingrouterCallerSession) GetNodeOperatorDigests(_stakingModuleId *big.Int, _nodeOperatorIds []*big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	return _Stakingrouter.Contract.GetNodeOperatorDigests(&_Stakingrouter.CallOpts, _stakingModuleId, _nodeOperatorIds)
}

// GetNodeOperatorDigests0 is a free data retrieval call binding the contract method 0xf8bb6d42.
//
// Solidity: function getNodeOperatorDigests(uint256 _stakingModuleId, uint256 _offset, uint256 _limit) view returns((uint256,bool,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Stakingrouter *StakingrouterCaller) GetNodeOperatorDigests0(opts *bind.CallOpts, _stakingModuleId *big.Int, _offset *big.Int, _limit *big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getNodeOperatorDigests0", _stakingModuleId, _offset, _limit)

	if err != nil {
		return *new([]StakingRouterNodeOperatorDigest), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingRouterNodeOperatorDigest)).(*[]StakingRouterNodeOperatorDigest)

	return out0, err

}

// GetNodeOperatorDigests0 is a free data retrieval call binding the contract method 0xf8bb6d42.
//
// Solidity: function getNodeOperatorDigests(uint256 _stakingModuleId, uint256 _offset, uint256 _limit) view returns((uint256,bool,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Stakingrouter *StakingrouterSession) GetNodeOperatorDigests0(_stakingModuleId *big.Int, _offset *big.Int, _limit *big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	return _Stakingrouter.Contract.GetNodeOperatorDigests0(&_Stakingrouter.CallOpts, _stakingModuleId, _offset, _limit)
}

// GetNodeOperatorDigests0 is a free data retrieval call binding the contract method 0xf8bb6d42.
//
// Solidity: function getNodeOperatorDigests(uint256 _stakingModuleId, uint256 _offset, uint256 _limit) view returns((uint256,bool,(uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256))[])
func (_Stakingrouter *StakingrouterCallerSession) GetNodeOperatorDigests0(_stakingModuleId *big.Int, _offset *big.Int, _limit *big.Int) ([]StakingRouterNodeOperatorDigest, error) {
	return _Stakingrouter.Contract.GetNodeOperatorDigests0(&_Stakingrouter.CallOpts, _stakingModuleId, _offset, _limit)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xaa5a1b9d.
//
// Solidity: function getNodeOperatorSummary(uint256 _stakingModuleId, uint256 _nodeOperatorId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) summary)
func (_Stakingrouter *StakingrouterCaller) GetNodeOperatorSummary(opts *bind.CallOpts, _stakingModuleId *big.Int, _nodeOperatorId *big.Int) (StakingRouterNodeOperatorSummary, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getNodeOperatorSummary", _stakingModuleId, _nodeOperatorId)

	if err != nil {
		return *new(StakingRouterNodeOperatorSummary), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingRouterNodeOperatorSummary)).(*StakingRouterNodeOperatorSummary)

	return out0, err

}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xaa5a1b9d.
//
// Solidity: function getNodeOperatorSummary(uint256 _stakingModuleId, uint256 _nodeOperatorId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) summary)
func (_Stakingrouter *StakingrouterSession) GetNodeOperatorSummary(_stakingModuleId *big.Int, _nodeOperatorId *big.Int) (StakingRouterNodeOperatorSummary, error) {
	return _Stakingrouter.Contract.GetNodeOperatorSummary(&_Stakingrouter.CallOpts, _stakingModuleId, _nodeOperatorId)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xaa5a1b9d.
//
// Solidity: function getNodeOperatorSummary(uint256 _stakingModuleId, uint256 _nodeOperatorId) view returns((uint256,uint256,uint256,uint256,uint256,uint256,uint256,uint256) summary)
func (_Stakingrouter *StakingrouterCallerSession) GetNodeOperatorSummary(_stakingModuleId *big.Int, _nodeOperatorId *big.Int) (StakingRouterNodeOperatorSummary, error) {
	return _Stakingrouter.Contract.GetNodeOperatorSummary(&_Stakingrouter.CallOpts, _stakingModuleId, _nodeOperatorId)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stakingrouter *StakingrouterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stakingrouter *StakingrouterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Stakingrouter.Contract.GetRoleAdmin(&_Stakingrouter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Stakingrouter *StakingrouterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Stakingrouter.Contract.GetRoleAdmin(&_Stakingrouter.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Stakingrouter *StakingrouterCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Stakingrouter *StakingrouterSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Stakingrouter.Contract.GetRoleMember(&_Stakingrouter.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Stakingrouter *StakingrouterCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Stakingrouter.Contract.GetRoleMember(&_Stakingrouter.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Stakingrouter *StakingrouterCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Stakingrouter *StakingrouterSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Stakingrouter.Contract.GetRoleMemberCount(&_Stakingrouter.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Stakingrouter *StakingrouterCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Stakingrouter.Contract.GetRoleMemberCount(&_Stakingrouter.CallOpts, role)
}

// GetStakingFeeAggregateDistribution is a free data retrieval call binding the contract method 0xfa5093eb.
//
// Solidity: function getStakingFeeAggregateDistribution() view returns(uint96 modulesFee, uint96 treasuryFee, uint256 basePrecision)
func (_Stakingrouter *StakingrouterCaller) GetStakingFeeAggregateDistribution(opts *bind.CallOpts) (struct {
	ModulesFee    *big.Int
	TreasuryFee   *big.Int
	BasePrecision *big.Int
}, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingFeeAggregateDistribution")

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
func (_Stakingrouter *StakingrouterSession) GetStakingFeeAggregateDistribution() (struct {
	ModulesFee    *big.Int
	TreasuryFee   *big.Int
	BasePrecision *big.Int
}, error) {
	return _Stakingrouter.Contract.GetStakingFeeAggregateDistribution(&_Stakingrouter.CallOpts)
}

// GetStakingFeeAggregateDistribution is a free data retrieval call binding the contract method 0xfa5093eb.
//
// Solidity: function getStakingFeeAggregateDistribution() view returns(uint96 modulesFee, uint96 treasuryFee, uint256 basePrecision)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingFeeAggregateDistribution() (struct {
	ModulesFee    *big.Int
	TreasuryFee   *big.Int
	BasePrecision *big.Int
}, error) {
	return _Stakingrouter.Contract.GetStakingFeeAggregateDistribution(&_Stakingrouter.CallOpts)
}

// GetStakingFeeAggregateDistributionE4Precision is a free data retrieval call binding the contract method 0xefcdcc0e.
//
// Solidity: function getStakingFeeAggregateDistributionE4Precision() view returns(uint16 modulesFee, uint16 treasuryFee)
func (_Stakingrouter *StakingrouterCaller) GetStakingFeeAggregateDistributionE4Precision(opts *bind.CallOpts) (struct {
	ModulesFee  uint16
	TreasuryFee uint16
}, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingFeeAggregateDistributionE4Precision")

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
func (_Stakingrouter *StakingrouterSession) GetStakingFeeAggregateDistributionE4Precision() (struct {
	ModulesFee  uint16
	TreasuryFee uint16
}, error) {
	return _Stakingrouter.Contract.GetStakingFeeAggregateDistributionE4Precision(&_Stakingrouter.CallOpts)
}

// GetStakingFeeAggregateDistributionE4Precision is a free data retrieval call binding the contract method 0xefcdcc0e.
//
// Solidity: function getStakingFeeAggregateDistributionE4Precision() view returns(uint16 modulesFee, uint16 treasuryFee)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingFeeAggregateDistributionE4Precision() (struct {
	ModulesFee  uint16
	TreasuryFee uint16
}, error) {
	return _Stakingrouter.Contract.GetStakingFeeAggregateDistributionE4Precision(&_Stakingrouter.CallOpts)
}

// GetStakingModule is a free data retrieval call binding the contract method 0xbc1bb190.
//
// Solidity: function getStakingModule(uint256 _stakingModuleId) view returns((uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256,uint16,uint64,uint64))
func (_Stakingrouter *StakingrouterCaller) GetStakingModule(opts *bind.CallOpts, _stakingModuleId *big.Int) (StakingRouterStakingModule, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModule", _stakingModuleId)

	if err != nil {
		return *new(StakingRouterStakingModule), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingRouterStakingModule)).(*StakingRouterStakingModule)

	return out0, err

}

// GetStakingModule is a free data retrieval call binding the contract method 0xbc1bb190.
//
// Solidity: function getStakingModule(uint256 _stakingModuleId) view returns((uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256,uint16,uint64,uint64))
func (_Stakingrouter *StakingrouterSession) GetStakingModule(_stakingModuleId *big.Int) (StakingRouterStakingModule, error) {
	return _Stakingrouter.Contract.GetStakingModule(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModule is a free data retrieval call binding the contract method 0xbc1bb190.
//
// Solidity: function getStakingModule(uint256 _stakingModuleId) view returns((uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256,uint16,uint64,uint64))
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModule(_stakingModuleId *big.Int) (StakingRouterStakingModule, error) {
	return _Stakingrouter.Contract.GetStakingModule(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleActiveValidatorsCount is a free data retrieval call binding the contract method 0x96b5d81c.
//
// Solidity: function getStakingModuleActiveValidatorsCount(uint256 _stakingModuleId) view returns(uint256 activeValidatorsCount)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleActiveValidatorsCount(opts *bind.CallOpts, _stakingModuleId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleActiveValidatorsCount", _stakingModuleId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingModuleActiveValidatorsCount is a free data retrieval call binding the contract method 0x96b5d81c.
//
// Solidity: function getStakingModuleActiveValidatorsCount(uint256 _stakingModuleId) view returns(uint256 activeValidatorsCount)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleActiveValidatorsCount(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleActiveValidatorsCount(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleActiveValidatorsCount is a free data retrieval call binding the contract method 0x96b5d81c.
//
// Solidity: function getStakingModuleActiveValidatorsCount(uint256 _stakingModuleId) view returns(uint256 activeValidatorsCount)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleActiveValidatorsCount(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleActiveValidatorsCount(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleDigests is a free data retrieval call binding the contract method 0x8525e3a1.
//
// Solidity: function getStakingModuleDigests(uint256[] _stakingModuleIds) view returns((uint256,uint256,(uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256,uint16,uint64,uint64),(uint256,uint256,uint256))[] digests)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleDigests(opts *bind.CallOpts, _stakingModuleIds []*big.Int) ([]StakingRouterStakingModuleDigest, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleDigests", _stakingModuleIds)

	if err != nil {
		return *new([]StakingRouterStakingModuleDigest), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingRouterStakingModuleDigest)).(*[]StakingRouterStakingModuleDigest)

	return out0, err

}

// GetStakingModuleDigests is a free data retrieval call binding the contract method 0x8525e3a1.
//
// Solidity: function getStakingModuleDigests(uint256[] _stakingModuleIds) view returns((uint256,uint256,(uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256,uint16,uint64,uint64),(uint256,uint256,uint256))[] digests)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleDigests(_stakingModuleIds []*big.Int) ([]StakingRouterStakingModuleDigest, error) {
	return _Stakingrouter.Contract.GetStakingModuleDigests(&_Stakingrouter.CallOpts, _stakingModuleIds)
}

// GetStakingModuleDigests is a free data retrieval call binding the contract method 0x8525e3a1.
//
// Solidity: function getStakingModuleDigests(uint256[] _stakingModuleIds) view returns((uint256,uint256,(uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256,uint16,uint64,uint64),(uint256,uint256,uint256))[] digests)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleDigests(_stakingModuleIds []*big.Int) ([]StakingRouterStakingModuleDigest, error) {
	return _Stakingrouter.Contract.GetStakingModuleDigests(&_Stakingrouter.CallOpts, _stakingModuleIds)
}

// GetStakingModuleIds is a free data retrieval call binding the contract method 0xf2aebb65.
//
// Solidity: function getStakingModuleIds() view returns(uint256[] stakingModuleIds)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetStakingModuleIds is a free data retrieval call binding the contract method 0xf2aebb65.
//
// Solidity: function getStakingModuleIds() view returns(uint256[] stakingModuleIds)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleIds() ([]*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleIds(&_Stakingrouter.CallOpts)
}

// GetStakingModuleIds is a free data retrieval call binding the contract method 0xf2aebb65.
//
// Solidity: function getStakingModuleIds() view returns(uint256[] stakingModuleIds)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleIds() ([]*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleIds(&_Stakingrouter.CallOpts)
}

// GetStakingModuleIsActive is a free data retrieval call binding the contract method 0x6608b11b.
//
// Solidity: function getStakingModuleIsActive(uint256 _stakingModuleId) view returns(bool)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleIsActive(opts *bind.CallOpts, _stakingModuleId *big.Int) (bool, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleIsActive", _stakingModuleId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetStakingModuleIsActive is a free data retrieval call binding the contract method 0x6608b11b.
//
// Solidity: function getStakingModuleIsActive(uint256 _stakingModuleId) view returns(bool)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleIsActive(_stakingModuleId *big.Int) (bool, error) {
	return _Stakingrouter.Contract.GetStakingModuleIsActive(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleIsActive is a free data retrieval call binding the contract method 0x6608b11b.
//
// Solidity: function getStakingModuleIsActive(uint256 _stakingModuleId) view returns(bool)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleIsActive(_stakingModuleId *big.Int) (bool, error) {
	return _Stakingrouter.Contract.GetStakingModuleIsActive(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleIsDepositsPaused is a free data retrieval call binding the contract method 0xe24ce9f1.
//
// Solidity: function getStakingModuleIsDepositsPaused(uint256 _stakingModuleId) view returns(bool)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleIsDepositsPaused(opts *bind.CallOpts, _stakingModuleId *big.Int) (bool, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleIsDepositsPaused", _stakingModuleId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetStakingModuleIsDepositsPaused is a free data retrieval call binding the contract method 0xe24ce9f1.
//
// Solidity: function getStakingModuleIsDepositsPaused(uint256 _stakingModuleId) view returns(bool)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleIsDepositsPaused(_stakingModuleId *big.Int) (bool, error) {
	return _Stakingrouter.Contract.GetStakingModuleIsDepositsPaused(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleIsDepositsPaused is a free data retrieval call binding the contract method 0xe24ce9f1.
//
// Solidity: function getStakingModuleIsDepositsPaused(uint256 _stakingModuleId) view returns(bool)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleIsDepositsPaused(_stakingModuleId *big.Int) (bool, error) {
	return _Stakingrouter.Contract.GetStakingModuleIsDepositsPaused(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleIsStopped is a free data retrieval call binding the contract method 0x6ada55b9.
//
// Solidity: function getStakingModuleIsStopped(uint256 _stakingModuleId) view returns(bool)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleIsStopped(opts *bind.CallOpts, _stakingModuleId *big.Int) (bool, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleIsStopped", _stakingModuleId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetStakingModuleIsStopped is a free data retrieval call binding the contract method 0x6ada55b9.
//
// Solidity: function getStakingModuleIsStopped(uint256 _stakingModuleId) view returns(bool)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleIsStopped(_stakingModuleId *big.Int) (bool, error) {
	return _Stakingrouter.Contract.GetStakingModuleIsStopped(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleIsStopped is a free data retrieval call binding the contract method 0x6ada55b9.
//
// Solidity: function getStakingModuleIsStopped(uint256 _stakingModuleId) view returns(bool)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleIsStopped(_stakingModuleId *big.Int) (bool, error) {
	return _Stakingrouter.Contract.GetStakingModuleIsStopped(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleLastDepositBlock is a free data retrieval call binding the contract method 0x473e0433.
//
// Solidity: function getStakingModuleLastDepositBlock(uint256 _stakingModuleId) view returns(uint256)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleLastDepositBlock(opts *bind.CallOpts, _stakingModuleId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleLastDepositBlock", _stakingModuleId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingModuleLastDepositBlock is a free data retrieval call binding the contract method 0x473e0433.
//
// Solidity: function getStakingModuleLastDepositBlock(uint256 _stakingModuleId) view returns(uint256)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleLastDepositBlock(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleLastDepositBlock(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleLastDepositBlock is a free data retrieval call binding the contract method 0x473e0433.
//
// Solidity: function getStakingModuleLastDepositBlock(uint256 _stakingModuleId) view returns(uint256)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleLastDepositBlock(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleLastDepositBlock(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleMaxDepositsCount is a free data retrieval call binding the contract method 0x19c64b79.
//
// Solidity: function getStakingModuleMaxDepositsCount(uint256 _stakingModuleId, uint256 _maxDepositsValue) view returns(uint256)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleMaxDepositsCount(opts *bind.CallOpts, _stakingModuleId *big.Int, _maxDepositsValue *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleMaxDepositsCount", _stakingModuleId, _maxDepositsValue)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingModuleMaxDepositsCount is a free data retrieval call binding the contract method 0x19c64b79.
//
// Solidity: function getStakingModuleMaxDepositsCount(uint256 _stakingModuleId, uint256 _maxDepositsValue) view returns(uint256)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleMaxDepositsCount(_stakingModuleId *big.Int, _maxDepositsValue *big.Int) (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleMaxDepositsCount(&_Stakingrouter.CallOpts, _stakingModuleId, _maxDepositsValue)
}

// GetStakingModuleMaxDepositsCount is a free data retrieval call binding the contract method 0x19c64b79.
//
// Solidity: function getStakingModuleMaxDepositsCount(uint256 _stakingModuleId, uint256 _maxDepositsValue) view returns(uint256)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleMaxDepositsCount(_stakingModuleId *big.Int, _maxDepositsValue *big.Int) (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleMaxDepositsCount(&_Stakingrouter.CallOpts, _stakingModuleId, _maxDepositsValue)
}

// GetStakingModuleMaxDepositsPerBlock is a free data retrieval call binding the contract method 0x20e948c8.
//
// Solidity: function getStakingModuleMaxDepositsPerBlock(uint256 _stakingModuleId) view returns(uint256)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleMaxDepositsPerBlock(opts *bind.CallOpts, _stakingModuleId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleMaxDepositsPerBlock", _stakingModuleId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingModuleMaxDepositsPerBlock is a free data retrieval call binding the contract method 0x20e948c8.
//
// Solidity: function getStakingModuleMaxDepositsPerBlock(uint256 _stakingModuleId) view returns(uint256)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleMaxDepositsPerBlock(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleMaxDepositsPerBlock(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleMaxDepositsPerBlock is a free data retrieval call binding the contract method 0x20e948c8.
//
// Solidity: function getStakingModuleMaxDepositsPerBlock(uint256 _stakingModuleId) view returns(uint256)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleMaxDepositsPerBlock(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleMaxDepositsPerBlock(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleMinDepositBlockDistance is a free data retrieval call binding the contract method 0xcb8fd4da.
//
// Solidity: function getStakingModuleMinDepositBlockDistance(uint256 _stakingModuleId) view returns(uint256)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleMinDepositBlockDistance(opts *bind.CallOpts, _stakingModuleId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleMinDepositBlockDistance", _stakingModuleId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingModuleMinDepositBlockDistance is a free data retrieval call binding the contract method 0xcb8fd4da.
//
// Solidity: function getStakingModuleMinDepositBlockDistance(uint256 _stakingModuleId) view returns(uint256)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleMinDepositBlockDistance(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleMinDepositBlockDistance(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleMinDepositBlockDistance is a free data retrieval call binding the contract method 0xcb8fd4da.
//
// Solidity: function getStakingModuleMinDepositBlockDistance(uint256 _stakingModuleId) view returns(uint256)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleMinDepositBlockDistance(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleMinDepositBlockDistance(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleNonce is a free data retrieval call binding the contract method 0x0519fbbf.
//
// Solidity: function getStakingModuleNonce(uint256 _stakingModuleId) view returns(uint256)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleNonce(opts *bind.CallOpts, _stakingModuleId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleNonce", _stakingModuleId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingModuleNonce is a free data retrieval call binding the contract method 0x0519fbbf.
//
// Solidity: function getStakingModuleNonce(uint256 _stakingModuleId) view returns(uint256)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleNonce(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleNonce(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleNonce is a free data retrieval call binding the contract method 0x0519fbbf.
//
// Solidity: function getStakingModuleNonce(uint256 _stakingModuleId) view returns(uint256)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleNonce(_stakingModuleId *big.Int) (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModuleNonce(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleStatus is a free data retrieval call binding the contract method 0x9fc5a6ed.
//
// Solidity: function getStakingModuleStatus(uint256 _stakingModuleId) view returns(uint8)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleStatus(opts *bind.CallOpts, _stakingModuleId *big.Int) (uint8, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleStatus", _stakingModuleId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStakingModuleStatus is a free data retrieval call binding the contract method 0x9fc5a6ed.
//
// Solidity: function getStakingModuleStatus(uint256 _stakingModuleId) view returns(uint8)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleStatus(_stakingModuleId *big.Int) (uint8, error) {
	return _Stakingrouter.Contract.GetStakingModuleStatus(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleStatus is a free data retrieval call binding the contract method 0x9fc5a6ed.
//
// Solidity: function getStakingModuleStatus(uint256 _stakingModuleId) view returns(uint8)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleStatus(_stakingModuleId *big.Int) (uint8, error) {
	return _Stakingrouter.Contract.GetStakingModuleStatus(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x07e203ac.
//
// Solidity: function getStakingModuleSummary(uint256 _stakingModuleId) view returns((uint256,uint256,uint256) summary)
func (_Stakingrouter *StakingrouterCaller) GetStakingModuleSummary(opts *bind.CallOpts, _stakingModuleId *big.Int) (StakingRouterStakingModuleSummary, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModuleSummary", _stakingModuleId)

	if err != nil {
		return *new(StakingRouterStakingModuleSummary), err
	}

	out0 := *abi.ConvertType(out[0], new(StakingRouterStakingModuleSummary)).(*StakingRouterStakingModuleSummary)

	return out0, err

}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x07e203ac.
//
// Solidity: function getStakingModuleSummary(uint256 _stakingModuleId) view returns((uint256,uint256,uint256) summary)
func (_Stakingrouter *StakingrouterSession) GetStakingModuleSummary(_stakingModuleId *big.Int) (StakingRouterStakingModuleSummary, error) {
	return _Stakingrouter.Contract.GetStakingModuleSummary(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x07e203ac.
//
// Solidity: function getStakingModuleSummary(uint256 _stakingModuleId) view returns((uint256,uint256,uint256) summary)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModuleSummary(_stakingModuleId *big.Int) (StakingRouterStakingModuleSummary, error) {
	return _Stakingrouter.Contract.GetStakingModuleSummary(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// GetStakingModules is a free data retrieval call binding the contract method 0x6183214d.
//
// Solidity: function getStakingModules() view returns((uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256,uint16,uint64,uint64)[] res)
func (_Stakingrouter *StakingrouterCaller) GetStakingModules(opts *bind.CallOpts) ([]StakingRouterStakingModule, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModules")

	if err != nil {
		return *new([]StakingRouterStakingModule), err
	}

	out0 := *abi.ConvertType(out[0], new([]StakingRouterStakingModule)).(*[]StakingRouterStakingModule)

	return out0, err

}

// GetStakingModules is a free data retrieval call binding the contract method 0x6183214d.
//
// Solidity: function getStakingModules() view returns((uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256,uint16,uint64,uint64)[] res)
func (_Stakingrouter *StakingrouterSession) GetStakingModules() ([]StakingRouterStakingModule, error) {
	return _Stakingrouter.Contract.GetStakingModules(&_Stakingrouter.CallOpts)
}

// GetStakingModules is a free data retrieval call binding the contract method 0x6183214d.
//
// Solidity: function getStakingModules() view returns((uint24,address,uint16,uint16,uint16,uint8,string,uint64,uint256,uint256,uint16,uint64,uint64)[] res)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModules() ([]StakingRouterStakingModule, error) {
	return _Stakingrouter.Contract.GetStakingModules(&_Stakingrouter.CallOpts)
}

// GetStakingModulesCount is a free data retrieval call binding the contract method 0x4a7583b6.
//
// Solidity: function getStakingModulesCount() view returns(uint256)
func (_Stakingrouter *StakingrouterCaller) GetStakingModulesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingModulesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStakingModulesCount is a free data retrieval call binding the contract method 0x4a7583b6.
//
// Solidity: function getStakingModulesCount() view returns(uint256)
func (_Stakingrouter *StakingrouterSession) GetStakingModulesCount() (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModulesCount(&_Stakingrouter.CallOpts)
}

// GetStakingModulesCount is a free data retrieval call binding the contract method 0x4a7583b6.
//
// Solidity: function getStakingModulesCount() view returns(uint256)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingModulesCount() (*big.Int, error) {
	return _Stakingrouter.Contract.GetStakingModulesCount(&_Stakingrouter.CallOpts)
}

// GetStakingRewardsDistribution is a free data retrieval call binding the contract method 0xba21ccae.
//
// Solidity: function getStakingRewardsDistribution() view returns(address[] recipients, uint256[] stakingModuleIds, uint96[] stakingModuleFees, uint96 totalFee, uint256 precisionPoints)
func (_Stakingrouter *StakingrouterCaller) GetStakingRewardsDistribution(opts *bind.CallOpts) (struct {
	Recipients        []common.Address
	StakingModuleIds  []*big.Int
	StakingModuleFees []*big.Int
	TotalFee          *big.Int
	PrecisionPoints   *big.Int
}, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getStakingRewardsDistribution")

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
func (_Stakingrouter *StakingrouterSession) GetStakingRewardsDistribution() (struct {
	Recipients        []common.Address
	StakingModuleIds  []*big.Int
	StakingModuleFees []*big.Int
	TotalFee          *big.Int
	PrecisionPoints   *big.Int
}, error) {
	return _Stakingrouter.Contract.GetStakingRewardsDistribution(&_Stakingrouter.CallOpts)
}

// GetStakingRewardsDistribution is a free data retrieval call binding the contract method 0xba21ccae.
//
// Solidity: function getStakingRewardsDistribution() view returns(address[] recipients, uint256[] stakingModuleIds, uint96[] stakingModuleFees, uint96 totalFee, uint256 precisionPoints)
func (_Stakingrouter *StakingrouterCallerSession) GetStakingRewardsDistribution() (struct {
	Recipients        []common.Address
	StakingModuleIds  []*big.Int
	StakingModuleFees []*big.Int
	TotalFee          *big.Int
	PrecisionPoints   *big.Int
}, error) {
	return _Stakingrouter.Contract.GetStakingRewardsDistribution(&_Stakingrouter.CallOpts)
}

// GetTotalFeeE4Precision is a free data retrieval call binding the contract method 0x9fbb7bae.
//
// Solidity: function getTotalFeeE4Precision() view returns(uint16 totalFee)
func (_Stakingrouter *StakingrouterCaller) GetTotalFeeE4Precision(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getTotalFeeE4Precision")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetTotalFeeE4Precision is a free data retrieval call binding the contract method 0x9fbb7bae.
//
// Solidity: function getTotalFeeE4Precision() view returns(uint16 totalFee)
func (_Stakingrouter *StakingrouterSession) GetTotalFeeE4Precision() (uint16, error) {
	return _Stakingrouter.Contract.GetTotalFeeE4Precision(&_Stakingrouter.CallOpts)
}

// GetTotalFeeE4Precision is a free data retrieval call binding the contract method 0x9fbb7bae.
//
// Solidity: function getTotalFeeE4Precision() view returns(uint16 totalFee)
func (_Stakingrouter *StakingrouterCallerSession) GetTotalFeeE4Precision() (uint16, error) {
	return _Stakingrouter.Contract.GetTotalFeeE4Precision(&_Stakingrouter.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_Stakingrouter *StakingrouterCaller) GetWithdrawalCredentials(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "getWithdrawalCredentials")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_Stakingrouter *StakingrouterSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _Stakingrouter.Contract.GetWithdrawalCredentials(&_Stakingrouter.CallOpts)
}

// GetWithdrawalCredentials is a free data retrieval call binding the contract method 0x56396715.
//
// Solidity: function getWithdrawalCredentials() view returns(bytes32)
func (_Stakingrouter *StakingrouterCallerSession) GetWithdrawalCredentials() ([32]byte, error) {
	return _Stakingrouter.Contract.GetWithdrawalCredentials(&_Stakingrouter.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stakingrouter *StakingrouterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stakingrouter *StakingrouterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Stakingrouter.Contract.HasRole(&_Stakingrouter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Stakingrouter *StakingrouterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Stakingrouter.Contract.HasRole(&_Stakingrouter.CallOpts, role, account)
}

// HasStakingModule is a free data retrieval call binding the contract method 0xa734329c.
//
// Solidity: function hasStakingModule(uint256 _stakingModuleId) view returns(bool)
func (_Stakingrouter *StakingrouterCaller) HasStakingModule(opts *bind.CallOpts, _stakingModuleId *big.Int) (bool, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "hasStakingModule", _stakingModuleId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasStakingModule is a free data retrieval call binding the contract method 0xa734329c.
//
// Solidity: function hasStakingModule(uint256 _stakingModuleId) view returns(bool)
func (_Stakingrouter *StakingrouterSession) HasStakingModule(_stakingModuleId *big.Int) (bool, error) {
	return _Stakingrouter.Contract.HasStakingModule(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// HasStakingModule is a free data retrieval call binding the contract method 0xa734329c.
//
// Solidity: function hasStakingModule(uint256 _stakingModuleId) view returns(bool)
func (_Stakingrouter *StakingrouterCallerSession) HasStakingModule(_stakingModuleId *big.Int) (bool, error) {
	return _Stakingrouter.Contract.HasStakingModule(&_Stakingrouter.CallOpts, _stakingModuleId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stakingrouter *StakingrouterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Stakingrouter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stakingrouter *StakingrouterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stakingrouter.Contract.SupportsInterface(&_Stakingrouter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Stakingrouter *StakingrouterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Stakingrouter.Contract.SupportsInterface(&_Stakingrouter.CallOpts, interfaceId)
}

// AddStakingModule is a paid mutator transaction binding the contract method 0xa4258a8d.
//
// Solidity: function addStakingModule(string _name, address _stakingModuleAddress, uint256 _stakeShareLimit, uint256 _priorityExitShareThreshold, uint256 _stakingModuleFee, uint256 _treasuryFee, uint256 _maxDepositsPerBlock, uint256 _minDepositBlockDistance) returns()
func (_Stakingrouter *StakingrouterTransactor) AddStakingModule(opts *bind.TransactOpts, _name string, _stakingModuleAddress common.Address, _stakeShareLimit *big.Int, _priorityExitShareThreshold *big.Int, _stakingModuleFee *big.Int, _treasuryFee *big.Int, _maxDepositsPerBlock *big.Int, _minDepositBlockDistance *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "addStakingModule", _name, _stakingModuleAddress, _stakeShareLimit, _priorityExitShareThreshold, _stakingModuleFee, _treasuryFee, _maxDepositsPerBlock, _minDepositBlockDistance)
}

// AddStakingModule is a paid mutator transaction binding the contract method 0xa4258a8d.
//
// Solidity: function addStakingModule(string _name, address _stakingModuleAddress, uint256 _stakeShareLimit, uint256 _priorityExitShareThreshold, uint256 _stakingModuleFee, uint256 _treasuryFee, uint256 _maxDepositsPerBlock, uint256 _minDepositBlockDistance) returns()
func (_Stakingrouter *StakingrouterSession) AddStakingModule(_name string, _stakingModuleAddress common.Address, _stakeShareLimit *big.Int, _priorityExitShareThreshold *big.Int, _stakingModuleFee *big.Int, _treasuryFee *big.Int, _maxDepositsPerBlock *big.Int, _minDepositBlockDistance *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.AddStakingModule(&_Stakingrouter.TransactOpts, _name, _stakingModuleAddress, _stakeShareLimit, _priorityExitShareThreshold, _stakingModuleFee, _treasuryFee, _maxDepositsPerBlock, _minDepositBlockDistance)
}

// AddStakingModule is a paid mutator transaction binding the contract method 0xa4258a8d.
//
// Solidity: function addStakingModule(string _name, address _stakingModuleAddress, uint256 _stakeShareLimit, uint256 _priorityExitShareThreshold, uint256 _stakingModuleFee, uint256 _treasuryFee, uint256 _maxDepositsPerBlock, uint256 _minDepositBlockDistance) returns()
func (_Stakingrouter *StakingrouterTransactorSession) AddStakingModule(_name string, _stakingModuleAddress common.Address, _stakeShareLimit *big.Int, _priorityExitShareThreshold *big.Int, _stakingModuleFee *big.Int, _treasuryFee *big.Int, _maxDepositsPerBlock *big.Int, _minDepositBlockDistance *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.AddStakingModule(&_Stakingrouter.TransactOpts, _name, _stakingModuleAddress, _stakeShareLimit, _priorityExitShareThreshold, _stakingModuleFee, _treasuryFee, _maxDepositsPerBlock, _minDepositBlockDistance)
}

// DecreaseStakingModuleVettedKeysCountByNodeOperator is a paid mutator transaction binding the contract method 0x2c201d31.
//
// Solidity: function decreaseStakingModuleVettedKeysCountByNodeOperator(uint256 _stakingModuleId, bytes _nodeOperatorIds, bytes _vettedSigningKeysCounts) returns()
func (_Stakingrouter *StakingrouterTransactor) DecreaseStakingModuleVettedKeysCountByNodeOperator(opts *bind.TransactOpts, _stakingModuleId *big.Int, _nodeOperatorIds []byte, _vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "decreaseStakingModuleVettedKeysCountByNodeOperator", _stakingModuleId, _nodeOperatorIds, _vettedSigningKeysCounts)
}

// DecreaseStakingModuleVettedKeysCountByNodeOperator is a paid mutator transaction binding the contract method 0x2c201d31.
//
// Solidity: function decreaseStakingModuleVettedKeysCountByNodeOperator(uint256 _stakingModuleId, bytes _nodeOperatorIds, bytes _vettedSigningKeysCounts) returns()
func (_Stakingrouter *StakingrouterSession) DecreaseStakingModuleVettedKeysCountByNodeOperator(_stakingModuleId *big.Int, _nodeOperatorIds []byte, _vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _Stakingrouter.Contract.DecreaseStakingModuleVettedKeysCountByNodeOperator(&_Stakingrouter.TransactOpts, _stakingModuleId, _nodeOperatorIds, _vettedSigningKeysCounts)
}

// DecreaseStakingModuleVettedKeysCountByNodeOperator is a paid mutator transaction binding the contract method 0x2c201d31.
//
// Solidity: function decreaseStakingModuleVettedKeysCountByNodeOperator(uint256 _stakingModuleId, bytes _nodeOperatorIds, bytes _vettedSigningKeysCounts) returns()
func (_Stakingrouter *StakingrouterTransactorSession) DecreaseStakingModuleVettedKeysCountByNodeOperator(_stakingModuleId *big.Int, _nodeOperatorIds []byte, _vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _Stakingrouter.Contract.DecreaseStakingModuleVettedKeysCountByNodeOperator(&_Stakingrouter.TransactOpts, _stakingModuleId, _nodeOperatorIds, _vettedSigningKeysCounts)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _depositsCount, uint256 _stakingModuleId, bytes _depositCalldata) payable returns()
func (_Stakingrouter *StakingrouterTransactor) Deposit(opts *bind.TransactOpts, _depositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "deposit", _depositsCount, _stakingModuleId, _depositCalldata)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _depositsCount, uint256 _stakingModuleId, bytes _depositCalldata) payable returns()
func (_Stakingrouter *StakingrouterSession) Deposit(_depositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _Stakingrouter.Contract.Deposit(&_Stakingrouter.TransactOpts, _depositsCount, _stakingModuleId, _depositCalldata)
}

// Deposit is a paid mutator transaction binding the contract method 0xaa0b7db7.
//
// Solidity: function deposit(uint256 _depositsCount, uint256 _stakingModuleId, bytes _depositCalldata) payable returns()
func (_Stakingrouter *StakingrouterTransactorSession) Deposit(_depositsCount *big.Int, _stakingModuleId *big.Int, _depositCalldata []byte) (*types.Transaction, error) {
	return _Stakingrouter.Contract.Deposit(&_Stakingrouter.TransactOpts, _depositsCount, _stakingModuleId, _depositCalldata)
}

// FinalizeUpgradeV3 is a paid mutator transaction binding the contract method 0x6d395b7e.
//
// Solidity: function finalizeUpgrade_v3() returns()
func (_Stakingrouter *StakingrouterTransactor) FinalizeUpgradeV3(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "finalizeUpgrade_v3")
}

// FinalizeUpgradeV3 is a paid mutator transaction binding the contract method 0x6d395b7e.
//
// Solidity: function finalizeUpgrade_v3() returns()
func (_Stakingrouter *StakingrouterSession) FinalizeUpgradeV3() (*types.Transaction, error) {
	return _Stakingrouter.Contract.FinalizeUpgradeV3(&_Stakingrouter.TransactOpts)
}

// FinalizeUpgradeV3 is a paid mutator transaction binding the contract method 0x6d395b7e.
//
// Solidity: function finalizeUpgrade_v3() returns()
func (_Stakingrouter *StakingrouterTransactorSession) FinalizeUpgradeV3() (*types.Transaction, error) {
	return _Stakingrouter.Contract.FinalizeUpgradeV3(&_Stakingrouter.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stakingrouter *StakingrouterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stakingrouter *StakingrouterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakingrouter.Contract.GrantRole(&_Stakingrouter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Stakingrouter *StakingrouterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakingrouter.Contract.GrantRole(&_Stakingrouter.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x6133f985.
//
// Solidity: function initialize(address _admin, address _lido, bytes32 _withdrawalCredentials) returns()
func (_Stakingrouter *StakingrouterTransactor) Initialize(opts *bind.TransactOpts, _admin common.Address, _lido common.Address, _withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "initialize", _admin, _lido, _withdrawalCredentials)
}

// Initialize is a paid mutator transaction binding the contract method 0x6133f985.
//
// Solidity: function initialize(address _admin, address _lido, bytes32 _withdrawalCredentials) returns()
func (_Stakingrouter *StakingrouterSession) Initialize(_admin common.Address, _lido common.Address, _withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Stakingrouter.Contract.Initialize(&_Stakingrouter.TransactOpts, _admin, _lido, _withdrawalCredentials)
}

// Initialize is a paid mutator transaction binding the contract method 0x6133f985.
//
// Solidity: function initialize(address _admin, address _lido, bytes32 _withdrawalCredentials) returns()
func (_Stakingrouter *StakingrouterTransactorSession) Initialize(_admin common.Address, _lido common.Address, _withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Stakingrouter.Contract.Initialize(&_Stakingrouter.TransactOpts, _admin, _lido, _withdrawalCredentials)
}

// OnValidatorExitTriggered is a paid mutator transaction binding the contract method 0x71416583.
//
// Solidity: function onValidatorExitTriggered((uint256,uint256,bytes)[] validatorExitData, uint256 _withdrawalRequestPaidFee, uint256 _exitType) returns()
func (_Stakingrouter *StakingrouterTransactor) OnValidatorExitTriggered(opts *bind.TransactOpts, validatorExitData []StakingRouterValidatorExitData, _withdrawalRequestPaidFee *big.Int, _exitType *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "onValidatorExitTriggered", validatorExitData, _withdrawalRequestPaidFee, _exitType)
}

// OnValidatorExitTriggered is a paid mutator transaction binding the contract method 0x71416583.
//
// Solidity: function onValidatorExitTriggered((uint256,uint256,bytes)[] validatorExitData, uint256 _withdrawalRequestPaidFee, uint256 _exitType) returns()
func (_Stakingrouter *StakingrouterSession) OnValidatorExitTriggered(validatorExitData []StakingRouterValidatorExitData, _withdrawalRequestPaidFee *big.Int, _exitType *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.OnValidatorExitTriggered(&_Stakingrouter.TransactOpts, validatorExitData, _withdrawalRequestPaidFee, _exitType)
}

// OnValidatorExitTriggered is a paid mutator transaction binding the contract method 0x71416583.
//
// Solidity: function onValidatorExitTriggered((uint256,uint256,bytes)[] validatorExitData, uint256 _withdrawalRequestPaidFee, uint256 _exitType) returns()
func (_Stakingrouter *StakingrouterTransactorSession) OnValidatorExitTriggered(validatorExitData []StakingRouterValidatorExitData, _withdrawalRequestPaidFee *big.Int, _exitType *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.OnValidatorExitTriggered(&_Stakingrouter.TransactOpts, validatorExitData, _withdrawalRequestPaidFee, _exitType)
}

// OnValidatorsCountsByNodeOperatorReportingFinished is a paid mutator transaction binding the contract method 0xdb3c7ba7.
//
// Solidity: function onValidatorsCountsByNodeOperatorReportingFinished() returns()
func (_Stakingrouter *StakingrouterTransactor) OnValidatorsCountsByNodeOperatorReportingFinished(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "onValidatorsCountsByNodeOperatorReportingFinished")
}

// OnValidatorsCountsByNodeOperatorReportingFinished is a paid mutator transaction binding the contract method 0xdb3c7ba7.
//
// Solidity: function onValidatorsCountsByNodeOperatorReportingFinished() returns()
func (_Stakingrouter *StakingrouterSession) OnValidatorsCountsByNodeOperatorReportingFinished() (*types.Transaction, error) {
	return _Stakingrouter.Contract.OnValidatorsCountsByNodeOperatorReportingFinished(&_Stakingrouter.TransactOpts)
}

// OnValidatorsCountsByNodeOperatorReportingFinished is a paid mutator transaction binding the contract method 0xdb3c7ba7.
//
// Solidity: function onValidatorsCountsByNodeOperatorReportingFinished() returns()
func (_Stakingrouter *StakingrouterTransactorSession) OnValidatorsCountsByNodeOperatorReportingFinished() (*types.Transaction, error) {
	return _Stakingrouter.Contract.OnValidatorsCountsByNodeOperatorReportingFinished(&_Stakingrouter.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Stakingrouter *StakingrouterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Stakingrouter *StakingrouterSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakingrouter.Contract.RenounceRole(&_Stakingrouter.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Stakingrouter *StakingrouterTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakingrouter.Contract.RenounceRole(&_Stakingrouter.TransactOpts, role, account)
}

// ReportRewardsMinted is a paid mutator transaction binding the contract method 0xaf124097.
//
// Solidity: function reportRewardsMinted(uint256[] _stakingModuleIds, uint256[] _totalShares) returns()
func (_Stakingrouter *StakingrouterTransactor) ReportRewardsMinted(opts *bind.TransactOpts, _stakingModuleIds []*big.Int, _totalShares []*big.Int) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "reportRewardsMinted", _stakingModuleIds, _totalShares)
}

// ReportRewardsMinted is a paid mutator transaction binding the contract method 0xaf124097.
//
// Solidity: function reportRewardsMinted(uint256[] _stakingModuleIds, uint256[] _totalShares) returns()
func (_Stakingrouter *StakingrouterSession) ReportRewardsMinted(_stakingModuleIds []*big.Int, _totalShares []*big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.ReportRewardsMinted(&_Stakingrouter.TransactOpts, _stakingModuleIds, _totalShares)
}

// ReportRewardsMinted is a paid mutator transaction binding the contract method 0xaf124097.
//
// Solidity: function reportRewardsMinted(uint256[] _stakingModuleIds, uint256[] _totalShares) returns()
func (_Stakingrouter *StakingrouterTransactorSession) ReportRewardsMinted(_stakingModuleIds []*big.Int, _totalShares []*big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.ReportRewardsMinted(&_Stakingrouter.TransactOpts, _stakingModuleIds, _totalShares)
}

// ReportStakingModuleExitedValidatorsCountByNodeOperator is a paid mutator transaction binding the contract method 0xc8ac4980.
//
// Solidity: function reportStakingModuleExitedValidatorsCountByNodeOperator(uint256 _stakingModuleId, bytes _nodeOperatorIds, bytes _exitedValidatorsCounts) returns()
func (_Stakingrouter *StakingrouterTransactor) ReportStakingModuleExitedValidatorsCountByNodeOperator(opts *bind.TransactOpts, _stakingModuleId *big.Int, _nodeOperatorIds []byte, _exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "reportStakingModuleExitedValidatorsCountByNodeOperator", _stakingModuleId, _nodeOperatorIds, _exitedValidatorsCounts)
}

// ReportStakingModuleExitedValidatorsCountByNodeOperator is a paid mutator transaction binding the contract method 0xc8ac4980.
//
// Solidity: function reportStakingModuleExitedValidatorsCountByNodeOperator(uint256 _stakingModuleId, bytes _nodeOperatorIds, bytes _exitedValidatorsCounts) returns()
func (_Stakingrouter *StakingrouterSession) ReportStakingModuleExitedValidatorsCountByNodeOperator(_stakingModuleId *big.Int, _nodeOperatorIds []byte, _exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _Stakingrouter.Contract.ReportStakingModuleExitedValidatorsCountByNodeOperator(&_Stakingrouter.TransactOpts, _stakingModuleId, _nodeOperatorIds, _exitedValidatorsCounts)
}

// ReportStakingModuleExitedValidatorsCountByNodeOperator is a paid mutator transaction binding the contract method 0xc8ac4980.
//
// Solidity: function reportStakingModuleExitedValidatorsCountByNodeOperator(uint256 _stakingModuleId, bytes _nodeOperatorIds, bytes _exitedValidatorsCounts) returns()
func (_Stakingrouter *StakingrouterTransactorSession) ReportStakingModuleExitedValidatorsCountByNodeOperator(_stakingModuleId *big.Int, _nodeOperatorIds []byte, _exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _Stakingrouter.Contract.ReportStakingModuleExitedValidatorsCountByNodeOperator(&_Stakingrouter.TransactOpts, _stakingModuleId, _nodeOperatorIds, _exitedValidatorsCounts)
}

// ReportValidatorExitDelay is a paid mutator transaction binding the contract method 0x32c4962c.
//
// Solidity: function reportValidatorExitDelay(uint256 _stakingModuleId, uint256 _nodeOperatorId, uint256 _proofSlotTimestamp, bytes _publicKey, uint256 _eligibleToExitInSec) returns()
func (_Stakingrouter *StakingrouterTransactor) ReportValidatorExitDelay(opts *bind.TransactOpts, _stakingModuleId *big.Int, _nodeOperatorId *big.Int, _proofSlotTimestamp *big.Int, _publicKey []byte, _eligibleToExitInSec *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "reportValidatorExitDelay", _stakingModuleId, _nodeOperatorId, _proofSlotTimestamp, _publicKey, _eligibleToExitInSec)
}

// ReportValidatorExitDelay is a paid mutator transaction binding the contract method 0x32c4962c.
//
// Solidity: function reportValidatorExitDelay(uint256 _stakingModuleId, uint256 _nodeOperatorId, uint256 _proofSlotTimestamp, bytes _publicKey, uint256 _eligibleToExitInSec) returns()
func (_Stakingrouter *StakingrouterSession) ReportValidatorExitDelay(_stakingModuleId *big.Int, _nodeOperatorId *big.Int, _proofSlotTimestamp *big.Int, _publicKey []byte, _eligibleToExitInSec *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.ReportValidatorExitDelay(&_Stakingrouter.TransactOpts, _stakingModuleId, _nodeOperatorId, _proofSlotTimestamp, _publicKey, _eligibleToExitInSec)
}

// ReportValidatorExitDelay is a paid mutator transaction binding the contract method 0x32c4962c.
//
// Solidity: function reportValidatorExitDelay(uint256 _stakingModuleId, uint256 _nodeOperatorId, uint256 _proofSlotTimestamp, bytes _publicKey, uint256 _eligibleToExitInSec) returns()
func (_Stakingrouter *StakingrouterTransactorSession) ReportValidatorExitDelay(_stakingModuleId *big.Int, _nodeOperatorId *big.Int, _proofSlotTimestamp *big.Int, _publicKey []byte, _eligibleToExitInSec *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.ReportValidatorExitDelay(&_Stakingrouter.TransactOpts, _stakingModuleId, _nodeOperatorId, _proofSlotTimestamp, _publicKey, _eligibleToExitInSec)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stakingrouter *StakingrouterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stakingrouter *StakingrouterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakingrouter.Contract.RevokeRole(&_Stakingrouter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Stakingrouter *StakingrouterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Stakingrouter.Contract.RevokeRole(&_Stakingrouter.TransactOpts, role, account)
}

// SetStakingModuleStatus is a paid mutator transaction binding the contract method 0xd0a2b1b8.
//
// Solidity: function setStakingModuleStatus(uint256 _stakingModuleId, uint8 _status) returns()
func (_Stakingrouter *StakingrouterTransactor) SetStakingModuleStatus(opts *bind.TransactOpts, _stakingModuleId *big.Int, _status uint8) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "setStakingModuleStatus", _stakingModuleId, _status)
}

// SetStakingModuleStatus is a paid mutator transaction binding the contract method 0xd0a2b1b8.
//
// Solidity: function setStakingModuleStatus(uint256 _stakingModuleId, uint8 _status) returns()
func (_Stakingrouter *StakingrouterSession) SetStakingModuleStatus(_stakingModuleId *big.Int, _status uint8) (*types.Transaction, error) {
	return _Stakingrouter.Contract.SetStakingModuleStatus(&_Stakingrouter.TransactOpts, _stakingModuleId, _status)
}

// SetStakingModuleStatus is a paid mutator transaction binding the contract method 0xd0a2b1b8.
//
// Solidity: function setStakingModuleStatus(uint256 _stakingModuleId, uint8 _status) returns()
func (_Stakingrouter *StakingrouterTransactorSession) SetStakingModuleStatus(_stakingModuleId *big.Int, _status uint8) (*types.Transaction, error) {
	return _Stakingrouter.Contract.SetStakingModuleStatus(&_Stakingrouter.TransactOpts, _stakingModuleId, _status)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xe97ee8cc.
//
// Solidity: function setWithdrawalCredentials(bytes32 _withdrawalCredentials) returns()
func (_Stakingrouter *StakingrouterTransactor) SetWithdrawalCredentials(opts *bind.TransactOpts, _withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "setWithdrawalCredentials", _withdrawalCredentials)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xe97ee8cc.
//
// Solidity: function setWithdrawalCredentials(bytes32 _withdrawalCredentials) returns()
func (_Stakingrouter *StakingrouterSession) SetWithdrawalCredentials(_withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Stakingrouter.Contract.SetWithdrawalCredentials(&_Stakingrouter.TransactOpts, _withdrawalCredentials)
}

// SetWithdrawalCredentials is a paid mutator transaction binding the contract method 0xe97ee8cc.
//
// Solidity: function setWithdrawalCredentials(bytes32 _withdrawalCredentials) returns()
func (_Stakingrouter *StakingrouterTransactorSession) SetWithdrawalCredentials(_withdrawalCredentials [32]byte) (*types.Transaction, error) {
	return _Stakingrouter.Contract.SetWithdrawalCredentials(&_Stakingrouter.TransactOpts, _withdrawalCredentials)
}

// UnsafeSetExitedValidatorsCount is a paid mutator transaction binding the contract method 0x7b274031.
//
// Solidity: function unsafeSetExitedValidatorsCount(uint256 _stakingModuleId, uint256 _nodeOperatorId, bool _triggerUpdateFinish, (uint256,uint256,uint256,uint256) _correction) returns()
func (_Stakingrouter *StakingrouterTransactor) UnsafeSetExitedValidatorsCount(opts *bind.TransactOpts, _stakingModuleId *big.Int, _nodeOperatorId *big.Int, _triggerUpdateFinish bool, _correction StakingRouterValidatorsCountsCorrection) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "unsafeSetExitedValidatorsCount", _stakingModuleId, _nodeOperatorId, _triggerUpdateFinish, _correction)
}

// UnsafeSetExitedValidatorsCount is a paid mutator transaction binding the contract method 0x7b274031.
//
// Solidity: function unsafeSetExitedValidatorsCount(uint256 _stakingModuleId, uint256 _nodeOperatorId, bool _triggerUpdateFinish, (uint256,uint256,uint256,uint256) _correction) returns()
func (_Stakingrouter *StakingrouterSession) UnsafeSetExitedValidatorsCount(_stakingModuleId *big.Int, _nodeOperatorId *big.Int, _triggerUpdateFinish bool, _correction StakingRouterValidatorsCountsCorrection) (*types.Transaction, error) {
	return _Stakingrouter.Contract.UnsafeSetExitedValidatorsCount(&_Stakingrouter.TransactOpts, _stakingModuleId, _nodeOperatorId, _triggerUpdateFinish, _correction)
}

// UnsafeSetExitedValidatorsCount is a paid mutator transaction binding the contract method 0x7b274031.
//
// Solidity: function unsafeSetExitedValidatorsCount(uint256 _stakingModuleId, uint256 _nodeOperatorId, bool _triggerUpdateFinish, (uint256,uint256,uint256,uint256) _correction) returns()
func (_Stakingrouter *StakingrouterTransactorSession) UnsafeSetExitedValidatorsCount(_stakingModuleId *big.Int, _nodeOperatorId *big.Int, _triggerUpdateFinish bool, _correction StakingRouterValidatorsCountsCorrection) (*types.Transaction, error) {
	return _Stakingrouter.Contract.UnsafeSetExitedValidatorsCount(&_Stakingrouter.TransactOpts, _stakingModuleId, _nodeOperatorId, _triggerUpdateFinish, _correction)
}

// UpdateExitedValidatorsCountByStakingModule is a paid mutator transaction binding the contract method 0xabd44a24.
//
// Solidity: function updateExitedValidatorsCountByStakingModule(uint256[] _stakingModuleIds, uint256[] _exitedValidatorsCounts) returns(uint256)
func (_Stakingrouter *StakingrouterTransactor) UpdateExitedValidatorsCountByStakingModule(opts *bind.TransactOpts, _stakingModuleIds []*big.Int, _exitedValidatorsCounts []*big.Int) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "updateExitedValidatorsCountByStakingModule", _stakingModuleIds, _exitedValidatorsCounts)
}

// UpdateExitedValidatorsCountByStakingModule is a paid mutator transaction binding the contract method 0xabd44a24.
//
// Solidity: function updateExitedValidatorsCountByStakingModule(uint256[] _stakingModuleIds, uint256[] _exitedValidatorsCounts) returns(uint256)
func (_Stakingrouter *StakingrouterSession) UpdateExitedValidatorsCountByStakingModule(_stakingModuleIds []*big.Int, _exitedValidatorsCounts []*big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.UpdateExitedValidatorsCountByStakingModule(&_Stakingrouter.TransactOpts, _stakingModuleIds, _exitedValidatorsCounts)
}

// UpdateExitedValidatorsCountByStakingModule is a paid mutator transaction binding the contract method 0xabd44a24.
//
// Solidity: function updateExitedValidatorsCountByStakingModule(uint256[] _stakingModuleIds, uint256[] _exitedValidatorsCounts) returns(uint256)
func (_Stakingrouter *StakingrouterTransactorSession) UpdateExitedValidatorsCountByStakingModule(_stakingModuleIds []*big.Int, _exitedValidatorsCounts []*big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.UpdateExitedValidatorsCountByStakingModule(&_Stakingrouter.TransactOpts, _stakingModuleIds, _exitedValidatorsCounts)
}

// UpdateStakingModule is a paid mutator transaction binding the contract method 0x77189583.
//
// Solidity: function updateStakingModule(uint256 _stakingModuleId, uint256 _stakeShareLimit, uint256 _priorityExitShareThreshold, uint256 _stakingModuleFee, uint256 _treasuryFee, uint256 _maxDepositsPerBlock, uint256 _minDepositBlockDistance) returns()
func (_Stakingrouter *StakingrouterTransactor) UpdateStakingModule(opts *bind.TransactOpts, _stakingModuleId *big.Int, _stakeShareLimit *big.Int, _priorityExitShareThreshold *big.Int, _stakingModuleFee *big.Int, _treasuryFee *big.Int, _maxDepositsPerBlock *big.Int, _minDepositBlockDistance *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "updateStakingModule", _stakingModuleId, _stakeShareLimit, _priorityExitShareThreshold, _stakingModuleFee, _treasuryFee, _maxDepositsPerBlock, _minDepositBlockDistance)
}

// UpdateStakingModule is a paid mutator transaction binding the contract method 0x77189583.
//
// Solidity: function updateStakingModule(uint256 _stakingModuleId, uint256 _stakeShareLimit, uint256 _priorityExitShareThreshold, uint256 _stakingModuleFee, uint256 _treasuryFee, uint256 _maxDepositsPerBlock, uint256 _minDepositBlockDistance) returns()
func (_Stakingrouter *StakingrouterSession) UpdateStakingModule(_stakingModuleId *big.Int, _stakeShareLimit *big.Int, _priorityExitShareThreshold *big.Int, _stakingModuleFee *big.Int, _treasuryFee *big.Int, _maxDepositsPerBlock *big.Int, _minDepositBlockDistance *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.UpdateStakingModule(&_Stakingrouter.TransactOpts, _stakingModuleId, _stakeShareLimit, _priorityExitShareThreshold, _stakingModuleFee, _treasuryFee, _maxDepositsPerBlock, _minDepositBlockDistance)
}

// UpdateStakingModule is a paid mutator transaction binding the contract method 0x77189583.
//
// Solidity: function updateStakingModule(uint256 _stakingModuleId, uint256 _stakeShareLimit, uint256 _priorityExitShareThreshold, uint256 _stakingModuleFee, uint256 _treasuryFee, uint256 _maxDepositsPerBlock, uint256 _minDepositBlockDistance) returns()
func (_Stakingrouter *StakingrouterTransactorSession) UpdateStakingModule(_stakingModuleId *big.Int, _stakeShareLimit *big.Int, _priorityExitShareThreshold *big.Int, _stakingModuleFee *big.Int, _treasuryFee *big.Int, _maxDepositsPerBlock *big.Int, _minDepositBlockDistance *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.UpdateStakingModule(&_Stakingrouter.TransactOpts, _stakingModuleId, _stakeShareLimit, _priorityExitShareThreshold, _stakingModuleFee, _treasuryFee, _maxDepositsPerBlock, _minDepositBlockDistance)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x9dd06848.
//
// Solidity: function updateTargetValidatorsLimits(uint256 _stakingModuleId, uint256 _nodeOperatorId, uint256 _targetLimitMode, uint256 _targetLimit) returns()
func (_Stakingrouter *StakingrouterTransactor) UpdateTargetValidatorsLimits(opts *bind.TransactOpts, _stakingModuleId *big.Int, _nodeOperatorId *big.Int, _targetLimitMode *big.Int, _targetLimit *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.contract.Transact(opts, "updateTargetValidatorsLimits", _stakingModuleId, _nodeOperatorId, _targetLimitMode, _targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x9dd06848.
//
// Solidity: function updateTargetValidatorsLimits(uint256 _stakingModuleId, uint256 _nodeOperatorId, uint256 _targetLimitMode, uint256 _targetLimit) returns()
func (_Stakingrouter *StakingrouterSession) UpdateTargetValidatorsLimits(_stakingModuleId *big.Int, _nodeOperatorId *big.Int, _targetLimitMode *big.Int, _targetLimit *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.UpdateTargetValidatorsLimits(&_Stakingrouter.TransactOpts, _stakingModuleId, _nodeOperatorId, _targetLimitMode, _targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x9dd06848.
//
// Solidity: function updateTargetValidatorsLimits(uint256 _stakingModuleId, uint256 _nodeOperatorId, uint256 _targetLimitMode, uint256 _targetLimit) returns()
func (_Stakingrouter *StakingrouterTransactorSession) UpdateTargetValidatorsLimits(_stakingModuleId *big.Int, _nodeOperatorId *big.Int, _targetLimitMode *big.Int, _targetLimit *big.Int) (*types.Transaction, error) {
	return _Stakingrouter.Contract.UpdateTargetValidatorsLimits(&_Stakingrouter.TransactOpts, _stakingModuleId, _nodeOperatorId, _targetLimitMode, _targetLimit)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Stakingrouter *StakingrouterTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Stakingrouter.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Stakingrouter *StakingrouterSession) Receive() (*types.Transaction, error) {
	return _Stakingrouter.Contract.Receive(&_Stakingrouter.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Stakingrouter *StakingrouterTransactorSession) Receive() (*types.Transaction, error) {
	return _Stakingrouter.Contract.Receive(&_Stakingrouter.TransactOpts)
}

// StakingrouterContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the Stakingrouter contract.
type StakingrouterContractVersionSetIterator struct {
	Event *StakingrouterContractVersionSet // Event containing the contract specifics and raw log

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
func (it *StakingrouterContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterContractVersionSet)
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
		it.Event = new(StakingrouterContractVersionSet)
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
func (it *StakingrouterContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterContractVersionSet represents a ContractVersionSet event raised by the Stakingrouter contract.
type StakingrouterContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Stakingrouter *StakingrouterFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*StakingrouterContractVersionSetIterator, error) {

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &StakingrouterContractVersionSetIterator{contract: _Stakingrouter.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Stakingrouter *StakingrouterFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *StakingrouterContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterContractVersionSet)
				if err := _Stakingrouter.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseContractVersionSet(log types.Log) (*StakingrouterContractVersionSet, error) {
	event := new(StakingrouterContractVersionSet)
	if err := _Stakingrouter.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterExitedAndStuckValidatorsCountsUpdateFailedIterator is returned from FilterExitedAndStuckValidatorsCountsUpdateFailed and is used to iterate over the raw logs and unpacked data for ExitedAndStuckValidatorsCountsUpdateFailed events raised by the Stakingrouter contract.
type StakingrouterExitedAndStuckValidatorsCountsUpdateFailedIterator struct {
	Event *StakingrouterExitedAndStuckValidatorsCountsUpdateFailed // Event containing the contract specifics and raw log

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
func (it *StakingrouterExitedAndStuckValidatorsCountsUpdateFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterExitedAndStuckValidatorsCountsUpdateFailed)
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
		it.Event = new(StakingrouterExitedAndStuckValidatorsCountsUpdateFailed)
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
func (it *StakingrouterExitedAndStuckValidatorsCountsUpdateFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterExitedAndStuckValidatorsCountsUpdateFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterExitedAndStuckValidatorsCountsUpdateFailed represents a ExitedAndStuckValidatorsCountsUpdateFailed event raised by the Stakingrouter contract.
type StakingrouterExitedAndStuckValidatorsCountsUpdateFailed struct {
	StakingModuleId    *big.Int
	LowLevelRevertData []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterExitedAndStuckValidatorsCountsUpdateFailed is a free log retrieval operation binding the contract event 0xe74bf895f0c3a2d6c74c40cbb362fdd9640035fc4226c72e3843809ad2a9d2b5.
//
// Solidity: event ExitedAndStuckValidatorsCountsUpdateFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Stakingrouter *StakingrouterFilterer) FilterExitedAndStuckValidatorsCountsUpdateFailed(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*StakingrouterExitedAndStuckValidatorsCountsUpdateFailedIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "ExitedAndStuckValidatorsCountsUpdateFailed", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterExitedAndStuckValidatorsCountsUpdateFailedIterator{contract: _Stakingrouter.contract, event: "ExitedAndStuckValidatorsCountsUpdateFailed", logs: logs, sub: sub}, nil
}

// WatchExitedAndStuckValidatorsCountsUpdateFailed is a free log subscription operation binding the contract event 0xe74bf895f0c3a2d6c74c40cbb362fdd9640035fc4226c72e3843809ad2a9d2b5.
//
// Solidity: event ExitedAndStuckValidatorsCountsUpdateFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Stakingrouter *StakingrouterFilterer) WatchExitedAndStuckValidatorsCountsUpdateFailed(opts *bind.WatchOpts, sink chan<- *StakingrouterExitedAndStuckValidatorsCountsUpdateFailed, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "ExitedAndStuckValidatorsCountsUpdateFailed", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterExitedAndStuckValidatorsCountsUpdateFailed)
				if err := _Stakingrouter.contract.UnpackLog(event, "ExitedAndStuckValidatorsCountsUpdateFailed", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseExitedAndStuckValidatorsCountsUpdateFailed(log types.Log) (*StakingrouterExitedAndStuckValidatorsCountsUpdateFailed, error) {
	event := new(StakingrouterExitedAndStuckValidatorsCountsUpdateFailed)
	if err := _Stakingrouter.contract.UnpackLog(event, "ExitedAndStuckValidatorsCountsUpdateFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterRewardsMintedReportFailedIterator is returned from FilterRewardsMintedReportFailed and is used to iterate over the raw logs and unpacked data for RewardsMintedReportFailed events raised by the Stakingrouter contract.
type StakingrouterRewardsMintedReportFailedIterator struct {
	Event *StakingrouterRewardsMintedReportFailed // Event containing the contract specifics and raw log

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
func (it *StakingrouterRewardsMintedReportFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterRewardsMintedReportFailed)
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
		it.Event = new(StakingrouterRewardsMintedReportFailed)
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
func (it *StakingrouterRewardsMintedReportFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterRewardsMintedReportFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterRewardsMintedReportFailed represents a RewardsMintedReportFailed event raised by the Stakingrouter contract.
type StakingrouterRewardsMintedReportFailed struct {
	StakingModuleId    *big.Int
	LowLevelRevertData []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterRewardsMintedReportFailed is a free log retrieval operation binding the contract event 0xf74208fedac7280fd11f8de0be14e00423dc5076da8e8ec8ca90e09257fff1b3.
//
// Solidity: event RewardsMintedReportFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Stakingrouter *StakingrouterFilterer) FilterRewardsMintedReportFailed(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*StakingrouterRewardsMintedReportFailedIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "RewardsMintedReportFailed", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterRewardsMintedReportFailedIterator{contract: _Stakingrouter.contract, event: "RewardsMintedReportFailed", logs: logs, sub: sub}, nil
}

// WatchRewardsMintedReportFailed is a free log subscription operation binding the contract event 0xf74208fedac7280fd11f8de0be14e00423dc5076da8e8ec8ca90e09257fff1b3.
//
// Solidity: event RewardsMintedReportFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Stakingrouter *StakingrouterFilterer) WatchRewardsMintedReportFailed(opts *bind.WatchOpts, sink chan<- *StakingrouterRewardsMintedReportFailed, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "RewardsMintedReportFailed", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterRewardsMintedReportFailed)
				if err := _Stakingrouter.contract.UnpackLog(event, "RewardsMintedReportFailed", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseRewardsMintedReportFailed(log types.Log) (*StakingrouterRewardsMintedReportFailed, error) {
	event := new(StakingrouterRewardsMintedReportFailed)
	if err := _Stakingrouter.contract.UnpackLog(event, "RewardsMintedReportFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Stakingrouter contract.
type StakingrouterRoleAdminChangedIterator struct {
	Event *StakingrouterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StakingrouterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterRoleAdminChanged)
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
		it.Event = new(StakingrouterRoleAdminChanged)
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
func (it *StakingrouterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterRoleAdminChanged represents a RoleAdminChanged event raised by the Stakingrouter contract.
type StakingrouterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stakingrouter *StakingrouterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StakingrouterRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterRoleAdminChangedIterator{contract: _Stakingrouter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Stakingrouter *StakingrouterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StakingrouterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterRoleAdminChanged)
				if err := _Stakingrouter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseRoleAdminChanged(log types.Log) (*StakingrouterRoleAdminChanged, error) {
	event := new(StakingrouterRoleAdminChanged)
	if err := _Stakingrouter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Stakingrouter contract.
type StakingrouterRoleGrantedIterator struct {
	Event *StakingrouterRoleGranted // Event containing the contract specifics and raw log

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
func (it *StakingrouterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterRoleGranted)
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
		it.Event = new(StakingrouterRoleGranted)
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
func (it *StakingrouterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterRoleGranted represents a RoleGranted event raised by the Stakingrouter contract.
type StakingrouterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stakingrouter *StakingrouterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingrouterRoleGrantedIterator, error) {

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

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterRoleGrantedIterator{contract: _Stakingrouter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stakingrouter *StakingrouterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StakingrouterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterRoleGranted)
				if err := _Stakingrouter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseRoleGranted(log types.Log) (*StakingrouterRoleGranted, error) {
	event := new(StakingrouterRoleGranted)
	if err := _Stakingrouter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Stakingrouter contract.
type StakingrouterRoleRevokedIterator struct {
	Event *StakingrouterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StakingrouterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterRoleRevoked)
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
		it.Event = new(StakingrouterRoleRevoked)
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
func (it *StakingrouterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterRoleRevoked represents a RoleRevoked event raised by the Stakingrouter contract.
type StakingrouterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stakingrouter *StakingrouterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StakingrouterRoleRevokedIterator, error) {

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

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterRoleRevokedIterator{contract: _Stakingrouter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Stakingrouter *StakingrouterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StakingrouterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterRoleRevoked)
				if err := _Stakingrouter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseRoleRevoked(log types.Log) (*StakingrouterRoleRevoked, error) {
	event := new(StakingrouterRoleRevoked)
	if err := _Stakingrouter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterStakingModuleAddedIterator is returned from FilterStakingModuleAdded and is used to iterate over the raw logs and unpacked data for StakingModuleAdded events raised by the Stakingrouter contract.
type StakingrouterStakingModuleAddedIterator struct {
	Event *StakingrouterStakingModuleAdded // Event containing the contract specifics and raw log

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
func (it *StakingrouterStakingModuleAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterStakingModuleAdded)
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
		it.Event = new(StakingrouterStakingModuleAdded)
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
func (it *StakingrouterStakingModuleAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterStakingModuleAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterStakingModuleAdded represents a StakingModuleAdded event raised by the Stakingrouter contract.
type StakingrouterStakingModuleAdded struct {
	StakingModuleId *big.Int
	StakingModule   common.Address
	Name            string
	CreatedBy       common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleAdded is a free log retrieval operation binding the contract event 0x43b5213f0e1666cd0b8692a73686164c94deb955a59c65e10dee8bb958e7ce3e.
//
// Solidity: event StakingModuleAdded(uint256 indexed stakingModuleId, address stakingModule, string name, address createdBy)
func (_Stakingrouter *StakingrouterFilterer) FilterStakingModuleAdded(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*StakingrouterStakingModuleAddedIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "StakingModuleAdded", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterStakingModuleAddedIterator{contract: _Stakingrouter.contract, event: "StakingModuleAdded", logs: logs, sub: sub}, nil
}

// WatchStakingModuleAdded is a free log subscription operation binding the contract event 0x43b5213f0e1666cd0b8692a73686164c94deb955a59c65e10dee8bb958e7ce3e.
//
// Solidity: event StakingModuleAdded(uint256 indexed stakingModuleId, address stakingModule, string name, address createdBy)
func (_Stakingrouter *StakingrouterFilterer) WatchStakingModuleAdded(opts *bind.WatchOpts, sink chan<- *StakingrouterStakingModuleAdded, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "StakingModuleAdded", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterStakingModuleAdded)
				if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleAdded", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseStakingModuleAdded(log types.Log) (*StakingrouterStakingModuleAdded, error) {
	event := new(StakingrouterStakingModuleAdded)
	if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterStakingModuleExitNotificationFailedIterator is returned from FilterStakingModuleExitNotificationFailed and is used to iterate over the raw logs and unpacked data for StakingModuleExitNotificationFailed events raised by the Stakingrouter contract.
type StakingrouterStakingModuleExitNotificationFailedIterator struct {
	Event *StakingrouterStakingModuleExitNotificationFailed // Event containing the contract specifics and raw log

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
func (it *StakingrouterStakingModuleExitNotificationFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterStakingModuleExitNotificationFailed)
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
		it.Event = new(StakingrouterStakingModuleExitNotificationFailed)
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
func (it *StakingrouterStakingModuleExitNotificationFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterStakingModuleExitNotificationFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterStakingModuleExitNotificationFailed represents a StakingModuleExitNotificationFailed event raised by the Stakingrouter contract.
type StakingrouterStakingModuleExitNotificationFailed struct {
	StakingModuleId *big.Int
	NodeOperatorId  *big.Int
	PublicKey       []byte
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleExitNotificationFailed is a free log retrieval operation binding the contract event 0xb639213d4cc5d7a615491fb0505dd448dee5074f322660125b7171993bf9bb1d.
//
// Solidity: event StakingModuleExitNotificationFailed(uint256 indexed stakingModuleId, uint256 indexed nodeOperatorId, bytes _publicKey)
func (_Stakingrouter *StakingrouterFilterer) FilterStakingModuleExitNotificationFailed(opts *bind.FilterOpts, stakingModuleId []*big.Int, nodeOperatorId []*big.Int) (*StakingrouterStakingModuleExitNotificationFailedIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}
	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "StakingModuleExitNotificationFailed", stakingModuleIdRule, nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterStakingModuleExitNotificationFailedIterator{contract: _Stakingrouter.contract, event: "StakingModuleExitNotificationFailed", logs: logs, sub: sub}, nil
}

// WatchStakingModuleExitNotificationFailed is a free log subscription operation binding the contract event 0xb639213d4cc5d7a615491fb0505dd448dee5074f322660125b7171993bf9bb1d.
//
// Solidity: event StakingModuleExitNotificationFailed(uint256 indexed stakingModuleId, uint256 indexed nodeOperatorId, bytes _publicKey)
func (_Stakingrouter *StakingrouterFilterer) WatchStakingModuleExitNotificationFailed(opts *bind.WatchOpts, sink chan<- *StakingrouterStakingModuleExitNotificationFailed, stakingModuleId []*big.Int, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}
	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "StakingModuleExitNotificationFailed", stakingModuleIdRule, nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterStakingModuleExitNotificationFailed)
				if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleExitNotificationFailed", log); err != nil {
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

// ParseStakingModuleExitNotificationFailed is a log parse operation binding the contract event 0xb639213d4cc5d7a615491fb0505dd448dee5074f322660125b7171993bf9bb1d.
//
// Solidity: event StakingModuleExitNotificationFailed(uint256 indexed stakingModuleId, uint256 indexed nodeOperatorId, bytes _publicKey)
func (_Stakingrouter *StakingrouterFilterer) ParseStakingModuleExitNotificationFailed(log types.Log) (*StakingrouterStakingModuleExitNotificationFailed, error) {
	event := new(StakingrouterStakingModuleExitNotificationFailed)
	if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleExitNotificationFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterStakingModuleExitedValidatorsIncompleteReportingIterator is returned from FilterStakingModuleExitedValidatorsIncompleteReporting and is used to iterate over the raw logs and unpacked data for StakingModuleExitedValidatorsIncompleteReporting events raised by the Stakingrouter contract.
type StakingrouterStakingModuleExitedValidatorsIncompleteReportingIterator struct {
	Event *StakingrouterStakingModuleExitedValidatorsIncompleteReporting // Event containing the contract specifics and raw log

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
func (it *StakingrouterStakingModuleExitedValidatorsIncompleteReportingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterStakingModuleExitedValidatorsIncompleteReporting)
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
		it.Event = new(StakingrouterStakingModuleExitedValidatorsIncompleteReporting)
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
func (it *StakingrouterStakingModuleExitedValidatorsIncompleteReportingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterStakingModuleExitedValidatorsIncompleteReportingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterStakingModuleExitedValidatorsIncompleteReporting represents a StakingModuleExitedValidatorsIncompleteReporting event raised by the Stakingrouter contract.
type StakingrouterStakingModuleExitedValidatorsIncompleteReporting struct {
	StakingModuleId                 *big.Int
	UnreportedExitedValidatorsCount *big.Int
	Raw                             types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleExitedValidatorsIncompleteReporting is a free log retrieval operation binding the contract event 0xdd2523ca96a639ba7e17420698937f71eddd8af012ccb36ff5c8fe96141acae9.
//
// Solidity: event StakingModuleExitedValidatorsIncompleteReporting(uint256 indexed stakingModuleId, uint256 unreportedExitedValidatorsCount)
func (_Stakingrouter *StakingrouterFilterer) FilterStakingModuleExitedValidatorsIncompleteReporting(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*StakingrouterStakingModuleExitedValidatorsIncompleteReportingIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "StakingModuleExitedValidatorsIncompleteReporting", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterStakingModuleExitedValidatorsIncompleteReportingIterator{contract: _Stakingrouter.contract, event: "StakingModuleExitedValidatorsIncompleteReporting", logs: logs, sub: sub}, nil
}

// WatchStakingModuleExitedValidatorsIncompleteReporting is a free log subscription operation binding the contract event 0xdd2523ca96a639ba7e17420698937f71eddd8af012ccb36ff5c8fe96141acae9.
//
// Solidity: event StakingModuleExitedValidatorsIncompleteReporting(uint256 indexed stakingModuleId, uint256 unreportedExitedValidatorsCount)
func (_Stakingrouter *StakingrouterFilterer) WatchStakingModuleExitedValidatorsIncompleteReporting(opts *bind.WatchOpts, sink chan<- *StakingrouterStakingModuleExitedValidatorsIncompleteReporting, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "StakingModuleExitedValidatorsIncompleteReporting", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterStakingModuleExitedValidatorsIncompleteReporting)
				if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleExitedValidatorsIncompleteReporting", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseStakingModuleExitedValidatorsIncompleteReporting(log types.Log) (*StakingrouterStakingModuleExitedValidatorsIncompleteReporting, error) {
	event := new(StakingrouterStakingModuleExitedValidatorsIncompleteReporting)
	if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleExitedValidatorsIncompleteReporting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterStakingModuleFeesSetIterator is returned from FilterStakingModuleFeesSet and is used to iterate over the raw logs and unpacked data for StakingModuleFeesSet events raised by the Stakingrouter contract.
type StakingrouterStakingModuleFeesSetIterator struct {
	Event *StakingrouterStakingModuleFeesSet // Event containing the contract specifics and raw log

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
func (it *StakingrouterStakingModuleFeesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterStakingModuleFeesSet)
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
		it.Event = new(StakingrouterStakingModuleFeesSet)
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
func (it *StakingrouterStakingModuleFeesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterStakingModuleFeesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterStakingModuleFeesSet represents a StakingModuleFeesSet event raised by the Stakingrouter contract.
type StakingrouterStakingModuleFeesSet struct {
	StakingModuleId  *big.Int
	StakingModuleFee *big.Int
	TreasuryFee      *big.Int
	SetBy            common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleFeesSet is a free log retrieval operation binding the contract event 0x303c8ac43d1b1f9b898ddd2915a294efa01e9b07c322d7deeb7db332b66f0410.
//
// Solidity: event StakingModuleFeesSet(uint256 indexed stakingModuleId, uint256 stakingModuleFee, uint256 treasuryFee, address setBy)
func (_Stakingrouter *StakingrouterFilterer) FilterStakingModuleFeesSet(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*StakingrouterStakingModuleFeesSetIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "StakingModuleFeesSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterStakingModuleFeesSetIterator{contract: _Stakingrouter.contract, event: "StakingModuleFeesSet", logs: logs, sub: sub}, nil
}

// WatchStakingModuleFeesSet is a free log subscription operation binding the contract event 0x303c8ac43d1b1f9b898ddd2915a294efa01e9b07c322d7deeb7db332b66f0410.
//
// Solidity: event StakingModuleFeesSet(uint256 indexed stakingModuleId, uint256 stakingModuleFee, uint256 treasuryFee, address setBy)
func (_Stakingrouter *StakingrouterFilterer) WatchStakingModuleFeesSet(opts *bind.WatchOpts, sink chan<- *StakingrouterStakingModuleFeesSet, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "StakingModuleFeesSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterStakingModuleFeesSet)
				if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleFeesSet", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseStakingModuleFeesSet(log types.Log) (*StakingrouterStakingModuleFeesSet, error) {
	event := new(StakingrouterStakingModuleFeesSet)
	if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleFeesSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterStakingModuleMaxDepositsPerBlockSetIterator is returned from FilterStakingModuleMaxDepositsPerBlockSet and is used to iterate over the raw logs and unpacked data for StakingModuleMaxDepositsPerBlockSet events raised by the Stakingrouter contract.
type StakingrouterStakingModuleMaxDepositsPerBlockSetIterator struct {
	Event *StakingrouterStakingModuleMaxDepositsPerBlockSet // Event containing the contract specifics and raw log

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
func (it *StakingrouterStakingModuleMaxDepositsPerBlockSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterStakingModuleMaxDepositsPerBlockSet)
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
		it.Event = new(StakingrouterStakingModuleMaxDepositsPerBlockSet)
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
func (it *StakingrouterStakingModuleMaxDepositsPerBlockSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterStakingModuleMaxDepositsPerBlockSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterStakingModuleMaxDepositsPerBlockSet represents a StakingModuleMaxDepositsPerBlockSet event raised by the Stakingrouter contract.
type StakingrouterStakingModuleMaxDepositsPerBlockSet struct {
	StakingModuleId     *big.Int
	MaxDepositsPerBlock *big.Int
	SetBy               common.Address
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleMaxDepositsPerBlockSet is a free log retrieval operation binding the contract event 0x72766c50f14fe492bd1281ceef0a57ad49a02b7e1042fb58723647bf38040f83.
//
// Solidity: event StakingModuleMaxDepositsPerBlockSet(uint256 indexed stakingModuleId, uint256 maxDepositsPerBlock, address setBy)
func (_Stakingrouter *StakingrouterFilterer) FilterStakingModuleMaxDepositsPerBlockSet(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*StakingrouterStakingModuleMaxDepositsPerBlockSetIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "StakingModuleMaxDepositsPerBlockSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterStakingModuleMaxDepositsPerBlockSetIterator{contract: _Stakingrouter.contract, event: "StakingModuleMaxDepositsPerBlockSet", logs: logs, sub: sub}, nil
}

// WatchStakingModuleMaxDepositsPerBlockSet is a free log subscription operation binding the contract event 0x72766c50f14fe492bd1281ceef0a57ad49a02b7e1042fb58723647bf38040f83.
//
// Solidity: event StakingModuleMaxDepositsPerBlockSet(uint256 indexed stakingModuleId, uint256 maxDepositsPerBlock, address setBy)
func (_Stakingrouter *StakingrouterFilterer) WatchStakingModuleMaxDepositsPerBlockSet(opts *bind.WatchOpts, sink chan<- *StakingrouterStakingModuleMaxDepositsPerBlockSet, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "StakingModuleMaxDepositsPerBlockSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterStakingModuleMaxDepositsPerBlockSet)
				if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleMaxDepositsPerBlockSet", log); err != nil {
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

// ParseStakingModuleMaxDepositsPerBlockSet is a log parse operation binding the contract event 0x72766c50f14fe492bd1281ceef0a57ad49a02b7e1042fb58723647bf38040f83.
//
// Solidity: event StakingModuleMaxDepositsPerBlockSet(uint256 indexed stakingModuleId, uint256 maxDepositsPerBlock, address setBy)
func (_Stakingrouter *StakingrouterFilterer) ParseStakingModuleMaxDepositsPerBlockSet(log types.Log) (*StakingrouterStakingModuleMaxDepositsPerBlockSet, error) {
	event := new(StakingrouterStakingModuleMaxDepositsPerBlockSet)
	if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleMaxDepositsPerBlockSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterStakingModuleMinDepositBlockDistanceSetIterator is returned from FilterStakingModuleMinDepositBlockDistanceSet and is used to iterate over the raw logs and unpacked data for StakingModuleMinDepositBlockDistanceSet events raised by the Stakingrouter contract.
type StakingrouterStakingModuleMinDepositBlockDistanceSetIterator struct {
	Event *StakingrouterStakingModuleMinDepositBlockDistanceSet // Event containing the contract specifics and raw log

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
func (it *StakingrouterStakingModuleMinDepositBlockDistanceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterStakingModuleMinDepositBlockDistanceSet)
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
		it.Event = new(StakingrouterStakingModuleMinDepositBlockDistanceSet)
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
func (it *StakingrouterStakingModuleMinDepositBlockDistanceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterStakingModuleMinDepositBlockDistanceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterStakingModuleMinDepositBlockDistanceSet represents a StakingModuleMinDepositBlockDistanceSet event raised by the Stakingrouter contract.
type StakingrouterStakingModuleMinDepositBlockDistanceSet struct {
	StakingModuleId         *big.Int
	MinDepositBlockDistance *big.Int
	SetBy                   common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleMinDepositBlockDistanceSet is a free log retrieval operation binding the contract event 0x4d106b4a7aff347abccca2dd6855d8d59d6cf792f1fdbb272c9858433d94b328.
//
// Solidity: event StakingModuleMinDepositBlockDistanceSet(uint256 indexed stakingModuleId, uint256 minDepositBlockDistance, address setBy)
func (_Stakingrouter *StakingrouterFilterer) FilterStakingModuleMinDepositBlockDistanceSet(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*StakingrouterStakingModuleMinDepositBlockDistanceSetIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "StakingModuleMinDepositBlockDistanceSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterStakingModuleMinDepositBlockDistanceSetIterator{contract: _Stakingrouter.contract, event: "StakingModuleMinDepositBlockDistanceSet", logs: logs, sub: sub}, nil
}

// WatchStakingModuleMinDepositBlockDistanceSet is a free log subscription operation binding the contract event 0x4d106b4a7aff347abccca2dd6855d8d59d6cf792f1fdbb272c9858433d94b328.
//
// Solidity: event StakingModuleMinDepositBlockDistanceSet(uint256 indexed stakingModuleId, uint256 minDepositBlockDistance, address setBy)
func (_Stakingrouter *StakingrouterFilterer) WatchStakingModuleMinDepositBlockDistanceSet(opts *bind.WatchOpts, sink chan<- *StakingrouterStakingModuleMinDepositBlockDistanceSet, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "StakingModuleMinDepositBlockDistanceSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterStakingModuleMinDepositBlockDistanceSet)
				if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleMinDepositBlockDistanceSet", log); err != nil {
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

// ParseStakingModuleMinDepositBlockDistanceSet is a log parse operation binding the contract event 0x4d106b4a7aff347abccca2dd6855d8d59d6cf792f1fdbb272c9858433d94b328.
//
// Solidity: event StakingModuleMinDepositBlockDistanceSet(uint256 indexed stakingModuleId, uint256 minDepositBlockDistance, address setBy)
func (_Stakingrouter *StakingrouterFilterer) ParseStakingModuleMinDepositBlockDistanceSet(log types.Log) (*StakingrouterStakingModuleMinDepositBlockDistanceSet, error) {
	event := new(StakingrouterStakingModuleMinDepositBlockDistanceSet)
	if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleMinDepositBlockDistanceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterStakingModuleShareLimitSetIterator is returned from FilterStakingModuleShareLimitSet and is used to iterate over the raw logs and unpacked data for StakingModuleShareLimitSet events raised by the Stakingrouter contract.
type StakingrouterStakingModuleShareLimitSetIterator struct {
	Event *StakingrouterStakingModuleShareLimitSet // Event containing the contract specifics and raw log

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
func (it *StakingrouterStakingModuleShareLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterStakingModuleShareLimitSet)
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
		it.Event = new(StakingrouterStakingModuleShareLimitSet)
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
func (it *StakingrouterStakingModuleShareLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterStakingModuleShareLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterStakingModuleShareLimitSet represents a StakingModuleShareLimitSet event raised by the Stakingrouter contract.
type StakingrouterStakingModuleShareLimitSet struct {
	StakingModuleId            *big.Int
	StakeShareLimit            *big.Int
	PriorityExitShareThreshold *big.Int
	SetBy                      common.Address
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleShareLimitSet is a free log retrieval operation binding the contract event 0x1730859048adcce16559e75a58fd609e9dbf7d34f39bcb7a45ad388dfbba0e4e.
//
// Solidity: event StakingModuleShareLimitSet(uint256 indexed stakingModuleId, uint256 stakeShareLimit, uint256 priorityExitShareThreshold, address setBy)
func (_Stakingrouter *StakingrouterFilterer) FilterStakingModuleShareLimitSet(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*StakingrouterStakingModuleShareLimitSetIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "StakingModuleShareLimitSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterStakingModuleShareLimitSetIterator{contract: _Stakingrouter.contract, event: "StakingModuleShareLimitSet", logs: logs, sub: sub}, nil
}

// WatchStakingModuleShareLimitSet is a free log subscription operation binding the contract event 0x1730859048adcce16559e75a58fd609e9dbf7d34f39bcb7a45ad388dfbba0e4e.
//
// Solidity: event StakingModuleShareLimitSet(uint256 indexed stakingModuleId, uint256 stakeShareLimit, uint256 priorityExitShareThreshold, address setBy)
func (_Stakingrouter *StakingrouterFilterer) WatchStakingModuleShareLimitSet(opts *bind.WatchOpts, sink chan<- *StakingrouterStakingModuleShareLimitSet, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "StakingModuleShareLimitSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterStakingModuleShareLimitSet)
				if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleShareLimitSet", log); err != nil {
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

// ParseStakingModuleShareLimitSet is a log parse operation binding the contract event 0x1730859048adcce16559e75a58fd609e9dbf7d34f39bcb7a45ad388dfbba0e4e.
//
// Solidity: event StakingModuleShareLimitSet(uint256 indexed stakingModuleId, uint256 stakeShareLimit, uint256 priorityExitShareThreshold, address setBy)
func (_Stakingrouter *StakingrouterFilterer) ParseStakingModuleShareLimitSet(log types.Log) (*StakingrouterStakingModuleShareLimitSet, error) {
	event := new(StakingrouterStakingModuleShareLimitSet)
	if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleShareLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterStakingModuleStatusSetIterator is returned from FilterStakingModuleStatusSet and is used to iterate over the raw logs and unpacked data for StakingModuleStatusSet events raised by the Stakingrouter contract.
type StakingrouterStakingModuleStatusSetIterator struct {
	Event *StakingrouterStakingModuleStatusSet // Event containing the contract specifics and raw log

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
func (it *StakingrouterStakingModuleStatusSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterStakingModuleStatusSet)
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
		it.Event = new(StakingrouterStakingModuleStatusSet)
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
func (it *StakingrouterStakingModuleStatusSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterStakingModuleStatusSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterStakingModuleStatusSet represents a StakingModuleStatusSet event raised by the Stakingrouter contract.
type StakingrouterStakingModuleStatusSet struct {
	StakingModuleId *big.Int
	Status          uint8
	SetBy           common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingModuleStatusSet is a free log retrieval operation binding the contract event 0xfd6f15fb2b48a21a60fe3d44d3c3a0433ca01e121b5124a63ec45c30ad925a17.
//
// Solidity: event StakingModuleStatusSet(uint256 indexed stakingModuleId, uint8 status, address setBy)
func (_Stakingrouter *StakingrouterFilterer) FilterStakingModuleStatusSet(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*StakingrouterStakingModuleStatusSetIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "StakingModuleStatusSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterStakingModuleStatusSetIterator{contract: _Stakingrouter.contract, event: "StakingModuleStatusSet", logs: logs, sub: sub}, nil
}

// WatchStakingModuleStatusSet is a free log subscription operation binding the contract event 0xfd6f15fb2b48a21a60fe3d44d3c3a0433ca01e121b5124a63ec45c30ad925a17.
//
// Solidity: event StakingModuleStatusSet(uint256 indexed stakingModuleId, uint8 status, address setBy)
func (_Stakingrouter *StakingrouterFilterer) WatchStakingModuleStatusSet(opts *bind.WatchOpts, sink chan<- *StakingrouterStakingModuleStatusSet, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "StakingModuleStatusSet", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterStakingModuleStatusSet)
				if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleStatusSet", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseStakingModuleStatusSet(log types.Log) (*StakingrouterStakingModuleStatusSet, error) {
	event := new(StakingrouterStakingModuleStatusSet)
	if err := _Stakingrouter.contract.UnpackLog(event, "StakingModuleStatusSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterStakingRouterETHDepositedIterator is returned from FilterStakingRouterETHDeposited and is used to iterate over the raw logs and unpacked data for StakingRouterETHDeposited events raised by the Stakingrouter contract.
type StakingrouterStakingRouterETHDepositedIterator struct {
	Event *StakingrouterStakingRouterETHDeposited // Event containing the contract specifics and raw log

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
func (it *StakingrouterStakingRouterETHDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterStakingRouterETHDeposited)
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
		it.Event = new(StakingrouterStakingRouterETHDeposited)
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
func (it *StakingrouterStakingRouterETHDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterStakingRouterETHDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterStakingRouterETHDeposited represents a StakingRouterETHDeposited event raised by the Stakingrouter contract.
type StakingrouterStakingRouterETHDeposited struct {
	StakingModuleId *big.Int
	Amount          *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterStakingRouterETHDeposited is a free log retrieval operation binding the contract event 0x9151b7f88aca05d432bb395647ef52b2ffc454e3c6afb69c95345af6b5a778c0.
//
// Solidity: event StakingRouterETHDeposited(uint256 indexed stakingModuleId, uint256 amount)
func (_Stakingrouter *StakingrouterFilterer) FilterStakingRouterETHDeposited(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*StakingrouterStakingRouterETHDepositedIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "StakingRouterETHDeposited", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterStakingRouterETHDepositedIterator{contract: _Stakingrouter.contract, event: "StakingRouterETHDeposited", logs: logs, sub: sub}, nil
}

// WatchStakingRouterETHDeposited is a free log subscription operation binding the contract event 0x9151b7f88aca05d432bb395647ef52b2ffc454e3c6afb69c95345af6b5a778c0.
//
// Solidity: event StakingRouterETHDeposited(uint256 indexed stakingModuleId, uint256 amount)
func (_Stakingrouter *StakingrouterFilterer) WatchStakingRouterETHDeposited(opts *bind.WatchOpts, sink chan<- *StakingrouterStakingRouterETHDeposited, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "StakingRouterETHDeposited", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterStakingRouterETHDeposited)
				if err := _Stakingrouter.contract.UnpackLog(event, "StakingRouterETHDeposited", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseStakingRouterETHDeposited(log types.Log) (*StakingrouterStakingRouterETHDeposited, error) {
	event := new(StakingrouterStakingRouterETHDeposited)
	if err := _Stakingrouter.contract.UnpackLog(event, "StakingRouterETHDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterWithdrawalCredentialsSetIterator is returned from FilterWithdrawalCredentialsSet and is used to iterate over the raw logs and unpacked data for WithdrawalCredentialsSet events raised by the Stakingrouter contract.
type StakingrouterWithdrawalCredentialsSetIterator struct {
	Event *StakingrouterWithdrawalCredentialsSet // Event containing the contract specifics and raw log

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
func (it *StakingrouterWithdrawalCredentialsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterWithdrawalCredentialsSet)
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
		it.Event = new(StakingrouterWithdrawalCredentialsSet)
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
func (it *StakingrouterWithdrawalCredentialsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterWithdrawalCredentialsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterWithdrawalCredentialsSet represents a WithdrawalCredentialsSet event raised by the Stakingrouter contract.
type StakingrouterWithdrawalCredentialsSet struct {
	WithdrawalCredentials [32]byte
	SetBy                 common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalCredentialsSet is a free log retrieval operation binding the contract event 0x82e72df77173eab89b00556d791a407a78f4605c5c2f0694967c8c429dd43c7c.
//
// Solidity: event WithdrawalCredentialsSet(bytes32 withdrawalCredentials, address setBy)
func (_Stakingrouter *StakingrouterFilterer) FilterWithdrawalCredentialsSet(opts *bind.FilterOpts) (*StakingrouterWithdrawalCredentialsSetIterator, error) {

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "WithdrawalCredentialsSet")
	if err != nil {
		return nil, err
	}
	return &StakingrouterWithdrawalCredentialsSetIterator{contract: _Stakingrouter.contract, event: "WithdrawalCredentialsSet", logs: logs, sub: sub}, nil
}

// WatchWithdrawalCredentialsSet is a free log subscription operation binding the contract event 0x82e72df77173eab89b00556d791a407a78f4605c5c2f0694967c8c429dd43c7c.
//
// Solidity: event WithdrawalCredentialsSet(bytes32 withdrawalCredentials, address setBy)
func (_Stakingrouter *StakingrouterFilterer) WatchWithdrawalCredentialsSet(opts *bind.WatchOpts, sink chan<- *StakingrouterWithdrawalCredentialsSet) (event.Subscription, error) {

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "WithdrawalCredentialsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterWithdrawalCredentialsSet)
				if err := _Stakingrouter.contract.UnpackLog(event, "WithdrawalCredentialsSet", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseWithdrawalCredentialsSet(log types.Log) (*StakingrouterWithdrawalCredentialsSet, error) {
	event := new(StakingrouterWithdrawalCredentialsSet)
	if err := _Stakingrouter.contract.UnpackLog(event, "WithdrawalCredentialsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingrouterWithdrawalsCredentialsChangeFailedIterator is returned from FilterWithdrawalsCredentialsChangeFailed and is used to iterate over the raw logs and unpacked data for WithdrawalsCredentialsChangeFailed events raised by the Stakingrouter contract.
type StakingrouterWithdrawalsCredentialsChangeFailedIterator struct {
	Event *StakingrouterWithdrawalsCredentialsChangeFailed // Event containing the contract specifics and raw log

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
func (it *StakingrouterWithdrawalsCredentialsChangeFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingrouterWithdrawalsCredentialsChangeFailed)
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
		it.Event = new(StakingrouterWithdrawalsCredentialsChangeFailed)
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
func (it *StakingrouterWithdrawalsCredentialsChangeFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingrouterWithdrawalsCredentialsChangeFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingrouterWithdrawalsCredentialsChangeFailed represents a WithdrawalsCredentialsChangeFailed event raised by the Stakingrouter contract.
type StakingrouterWithdrawalsCredentialsChangeFailed struct {
	StakingModuleId    *big.Int
	LowLevelRevertData []byte
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalsCredentialsChangeFailed is a free log retrieval operation binding the contract event 0x0d64b11929aa111ca874dd00b5b0cc2d82b741be924ec9e3691e67c71552f623.
//
// Solidity: event WithdrawalsCredentialsChangeFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Stakingrouter *StakingrouterFilterer) FilterWithdrawalsCredentialsChangeFailed(opts *bind.FilterOpts, stakingModuleId []*big.Int) (*StakingrouterWithdrawalsCredentialsChangeFailedIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.FilterLogs(opts, "WithdrawalsCredentialsChangeFailed", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return &StakingrouterWithdrawalsCredentialsChangeFailedIterator{contract: _Stakingrouter.contract, event: "WithdrawalsCredentialsChangeFailed", logs: logs, sub: sub}, nil
}

// WatchWithdrawalsCredentialsChangeFailed is a free log subscription operation binding the contract event 0x0d64b11929aa111ca874dd00b5b0cc2d82b741be924ec9e3691e67c71552f623.
//
// Solidity: event WithdrawalsCredentialsChangeFailed(uint256 indexed stakingModuleId, bytes lowLevelRevertData)
func (_Stakingrouter *StakingrouterFilterer) WatchWithdrawalsCredentialsChangeFailed(opts *bind.WatchOpts, sink chan<- *StakingrouterWithdrawalsCredentialsChangeFailed, stakingModuleId []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}

	logs, sub, err := _Stakingrouter.contract.WatchLogs(opts, "WithdrawalsCredentialsChangeFailed", stakingModuleIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingrouterWithdrawalsCredentialsChangeFailed)
				if err := _Stakingrouter.contract.UnpackLog(event, "WithdrawalsCredentialsChangeFailed", log); err != nil {
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
func (_Stakingrouter *StakingrouterFilterer) ParseWithdrawalsCredentialsChangeFailed(log types.Log) (*StakingrouterWithdrawalsCredentialsChangeFailed, error) {
	event := new(StakingrouterWithdrawalsCredentialsChangeFailed)
	if err := _Stakingrouter.contract.UnpackLog(event, "WithdrawalsCredentialsChangeFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
