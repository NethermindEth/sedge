// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package csmodule

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

// ICSAccountingPermitInput is an auto generated low-level Go binding around an user-defined struct.
type ICSAccountingPermitInput struct {
	Value    *big.Int
	Deadline *big.Int
	V        uint8
	R        [32]byte
	S        [32]byte
}

// NodeOperator is an auto generated low-level Go binding around an user-defined struct.
type NodeOperator struct {
	TotalAddedKeys             uint32
	TotalWithdrawnKeys         uint32
	TotalDepositedKeys         uint32
	TotalVettedKeys            uint32
	StuckValidatorsCount       uint32
	DepositableValidatorsCount uint32
	TargetLimit                uint32
	TargetLimitMode            uint8
	TotalExitedKeys            uint32
	EnqueuedCount              uint32
	ManagerAddress             common.Address
	ProposedManagerAddress     common.Address
	RewardAddress              common.Address
	ProposedRewardAddress      common.Address
	ExtendedManagerPermissions bool
	UsedPriorityQueue          bool
}

// NodeOperatorManagementProperties is an auto generated low-level Go binding around an user-defined struct.
type NodeOperatorManagementProperties struct {
	ManagerAddress             common.Address
	RewardAddress              common.Address
	ExtendedManagerPermissions bool
}

// ValidatorWithdrawalInfo is an auto generated low-level Go binding around an user-defined struct.
type ValidatorWithdrawalInfo struct {
	NodeOperatorId *big.Int
	KeyIndex       *big.Int
	Amount         *big.Int
}

// CsmoduleMetaData contains all meta data concerning the Csmodule contract.
var CsmoduleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleType\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"lidoLocator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"parametersRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_accounting\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"exitPenalties\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyProposed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotAddKeys\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitedKeysDecrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitedKeysHigherThanTotalDeposited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeysCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReportData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVetKeysPointer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"KeysLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MethodCallIsNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoQueuedKeysToMigrate\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeOperatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToRecover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEligibleForPriorityQueue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughKeys\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseUntilMustBeInFuture\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedExpected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriorityQueueAlreadyUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PriorityQueueMaxDepositsUsed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueIsEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueLookupNoLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResumedExpected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotEligible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotManagerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotProposedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotRewardAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SigningKeysInvalidOffset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAccountingAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroExitPenaltiesAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroLocatorAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroParametersRegistryAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPauseDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroRewardAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroSenderAddress\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"queuePriority\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"BatchEnqueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositableKeysCount\",\"type\":\"uint256\"}],\"name\":\"DepositableSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositedKeysCount\",\"type\":\"uint256\"}],\"name\":\"DepositedSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsStealingPenaltyCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsStealingPenaltyCompensated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposedBlockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stolenAmount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsStealingPenaltyReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"ELRewardsStealingPenaltySettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedKeysCount\",\"type\":\"uint256\"}],\"name\":\"ExitedSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"KeyRemovalChargeApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"extendedManagerPermissions\",\"type\":\"bool\"}],\"name\":\"NodeOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProposedAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProposedAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorManagerAddressChangeProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorManagerAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProposedAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProposedAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorRewardAddressChangeProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorRewardAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"NonceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"ReferrerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"SigningKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"SigningKeyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"StETHSharesRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"TargetValidatorsCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalKeysCount\",\"type\":\"uint256\"}],\"name\":\"TotalSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vettedKeysCount\",\"type\":\"uint256\"}],\"name\":\"VettedSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"VettedSigningKeysCountDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"WithdrawalSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACCOUNTING\",\"outputs\":[{\"internalType\":\"contractICSAccounting\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CREATE_NODE_OPERATOR_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEPOSIT_SIZE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXIT_PENALTIES\",\"outputs\":[{\"internalType\":\"contractICSExitPenalties\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_DISTRIBUTOR\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIDO_LOCATOR\",\"outputs\":[{\"internalType\":\"contractILidoLocator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PARAMETERS_REGISTRY\",\"outputs\":[{\"internalType\":\"contractICSParametersRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_INFINITELY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUEUE_LEGACY_PRIORITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUEUE_LOWEST_PRIORITY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECOVERER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORT_EL_REWARDS_STEALING_PENALTY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_ROUTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STETH\",\"outputs\":[{\"internalType\":\"contractIStETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accounting\",\"outputs\":[{\"internalType\":\"contractICSAccounting\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"addValidatorKeysETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"addValidatorKeysStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"addValidatorKeysWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"cancelELRewardsStealingPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"changeNodeOperatorRewardAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"}],\"name\":\"cleanDepositQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"removed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastRemovedAtDepth\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"compensateELRewardsStealingPenalty\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"confirmNodeOperatorManagerAddressChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"confirmNodeOperatorRewardAddressChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extendedManagerPermissions\",\"type\":\"bool\"}],\"internalType\":\"structNodeOperatorManagementProperties\",\"name\":\"managementProperties\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"createNodeOperator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeOperatorIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vettedSigningKeysCounts\",\"type\":\"bytes\"}],\"name\":\"decreaseVettedSigningKeysCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"queuePriority\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"index\",\"type\":\"uint128\"}],\"name\":\"depositQueueItem\",\"outputs\":[{\"internalType\":\"Batch\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"queuePriority\",\"type\":\"uint256\"}],\"name\":\"depositQueuePointers\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"head\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tail\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"exitDeadlineThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"finalizeUpgradeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveNodeOperatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitializedVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperator\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"totalAddedKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalWithdrawnKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalDepositedKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalVettedKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"targetLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"targetLimitMode\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"totalExitedKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"enqueuedCount\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedRewardAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extendedManagerPermissions\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"usedPriorityQueue\",\"type\":\"bool\"}],\"internalType\":\"structNodeOperator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nodeOperatorIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorIsActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorManagementProperties\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extendedManagerPermissions\",\"type\":\"bool\"}],\"internalType\":\"structNodeOperatorManagementProperties\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorNonWithdrawnKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorSummary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorTotalDepositedKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalDepositedKeys\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeOperatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResumeSinceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"}],\"name\":\"getSigningKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"}],\"name\":\"getSigningKeysWithSignatures\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"keys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingModuleSummary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eligibleToExitInSec\",\"type\":\"uint256\"}],\"name\":\"isValidatorExitDelayPenaltyApplicable\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"name\":\"isValidatorWithdrawn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"migrateToPriorityQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositsCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"obtainDepositData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onExitedAndStuckValidatorsCountsUpdated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"}],\"name\":\"onRewardsMinted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"withdrawalRequestPaidFee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitType\",\"type\":\"uint256\"}],\"name\":\"onValidatorExitTriggered\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onWithdrawalCredentialsChanged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"pauseFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"proposeNodeOperatorManagerAddressChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"proposeNodeOperatorRewardAddressChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"}],\"name\":\"removeKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reportELRewardsStealingPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"eligibleToExitInSec\",\"type\":\"uint256\"}],\"name\":\"reportValidatorExitDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"resetNodeOperatorManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nodeOperatorIds\",\"type\":\"uint256[]\"}],\"name\":\"settleELRewardsStealingPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structValidatorWithdrawalInfo[]\",\"name\":\"withdrawalsInfo\",\"type\":\"tuple[]\"}],\"name\":\"submitWithdrawals\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitedValidatorsKeysCount\",\"type\":\"uint256\"}],\"name\":\"unsafeUpdateValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"updateDepositableValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeOperatorIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"exitedValidatorsCounts\",\"type\":\"bytes\"}],\"name\":\"updateExitedValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetLimit\",\"type\":\"uint256\"}],\"name\":\"updateTargetValidatorsLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600436106104bb575f3560e01c80638573e3511161026d578063a6ab5b9c1161014a578063ca15c873116100be578063e00bfe5011610083578063e00bfe50146112db578063e7705db61461130e578063e864299e14611341578063f3f449c71461134c578063f696ccb31461136b578063fa367c9e1461138a575f80fd5b8063ca15c87314611223578063d087d28814611242578063d547741f14611256578063d614ae0c14611275578063dbba4b48146112a8575f80fd5b8063b187bd261161010f578063b187bd261461112c578063b3076c3c14611140578063b3c650151461119a578063b643189b146111c6578063bee41b58146111e5578063c4d66de814611204575f80fd5b8063a6ab5b9c14611088578063a6b89b81146110a7578063a70c70e414610d6b578063acf1c948146110da578063b055e15c1461110d575f80fd5b806394120368116101e15780639c963aef116101a65780639c963aef14610fd8578063a0c8c47e14610ff7578063a1913f4b1461102f578063a217fddf14611042578063a302ee3814611055578063a4516c9814611069575f80fd5b80639412036814610f045780639417366f14610f235780639624e83e14610f425780639abddf0914610f745780639b00c14614610fb9575f80fd5b80638d7e4017116102325780638d7e401714610e335780638eab3cd014610e525780638ec6902814610e715780639010d07c14610eb257806390c09bdb14610ed157806391d1485414610ee5575f80fd5b80638573e35114610d8f57806388984a9714610dc25780638980f11f14610dd65780638b3ac71d14610df55780638cabe95914610e14575f80fd5b806352d8bfc21161039b5780636a6304cc1161030f578063743f5105116102d4578063743f510514610cbb57806375a401da14610cee57806380231f1514610d0d578063819d4cc614610d2d57806383b57a4e14610d4c5780638469cbd314610d6b575f80fd5b80636a6304cc14610c035780636bb1bfdf14610c225780636dc3f2bd14610c415780636efe37a214610c74578063735dfa2814610c87575f80fd5b806359e25c121161036057806359e25c12146109815780635c654ad9146109ad5780635e2fb908146109cc57806365c14dc7146109fd5780636910dcce14610bb1578063693cc60014610be4575f80fd5b806352d8bfc2146108c657806353433643146108da57806357f9c3411461090f5780635810f6221461092e578063589ff76c1461096d575f80fd5b806336bf3325116104325780633f04f0c8116103f75780633f04f0c8146107445780634004480114610777578063499b8e9a146107965780634febc81b1461084e57806350388cb61461087a5780635204281c146108a7575f80fd5b806336bf33251461069857806337b12b5f146106b457806337ebdf6f146106d3578063388dd1d1146106f2578063389ed26714610711575f80fd5b8063248a9ca311610483578063248a9ca31461058357806328d6d36b146105bd5780632de03aa1146105dc5780632f2ff15d1461060f5780632fc887411461062e57806336568abe14610679575f80fd5b806301ffc9a7146104bf578063046f7da2146104f357806308a679ad1461050957806315dae03e146105285780631b40b23114610564575b5f80fd5b3480156104ca575f80fd5b506104de6104d93660046153a6565b6113bd565b60405190151581526020015b60405180910390f35b3480156104fe575f80fd5b506105076113e7565b005b348015610514575f80fd5b506105076105233660046153cd565b61141c565b348015610533575f80fd5b507f636f6d6d756e6974792d6f6e636861696e2d76310000000000000000000000005b6040519081526020016104ea565b34801561056f575f80fd5b5061050761057e36600461540a565b611522565b34801561058e575f80fd5b5061055661059d366004615438565b5f9081525f80516020615f34833981519152602052604090206001015490565b3480156105c8575f80fd5b506105566105d7366004615438565b61159b565b3480156105e7575f80fd5b506105567f2fc10cc8ae19568712f7a176fb4978616a610650813c9d05326c34abb62749c781565b34801561061a575f80fd5b5061050761062936600461540a565b61163a565b348015610639575f80fd5b506106617f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e42881565b6040516001600160a01b0390911681526020016104ea565b348015610684575f80fd5b5061050761069336600461540a565b611670565b3480156106a3575f80fd5b506105566801bc16d674ec80000081565b3480156106bf575f80fd5b506105076106ce36600461544f565b6116a8565b3480156106de575f80fd5b506105566106ed3660046154bd565b6117d2565b3480156106fd575f80fd5b5061050761070c3660046153cd565b611806565b34801561071c575f80fd5b506105567f139c2898040ef16910dc9f44dc697df79363da767d8bc92f2e310312b816e46d81565b34801561074f575f80fd5b506105567fe85fdec10fe0f93d0792364051df7c3d73e37c17b3a954bffe593960e3cd301281565b348015610782575f80fd5b506105076107913660046154eb565b6119c2565b3480156107a1575f80fd5b5061081c6107b0366004615438565b60408051606080820183525f80835260208084018290529284018190529384526006825292829020825193840183526001810154600160401b90046001600160a01b0390811685526003820154169184019190915260040154600160a01b900460ff1615159082015290565b6040805182516001600160a01b03908116825260208085015190911690820152918101511515908201526060016104ea565b348015610859575f80fd5b5061086d6108683660046154eb565b611ab9565b6040516104ea919061550b565b348015610885575f80fd5b506108996108943660046153cd565b611ba1565b6040516104ea929190615591565b3480156108b2575f80fd5b506105076108c1366004615438565b611bd3565b3480156108d1575f80fd5b50610507611c36565b3480156108e5575f80fd5b506104de6108f43660046154eb565b60809190911b175f9081526007602052604090205460ff1690565b34801561091a575f80fd5b506105076109293660046155f9565b611c92565b348015610939575f80fd5b5061094d610948366004615438565b611d39565b604080516001600160801b039384168152929091166020830152016104ea565b348015610978575f80fd5b50610556611d63565b34801561098c575f80fd5b506109a061099b3660046153cd565b611d91565b6040516104ea919061564e565b3480156109b8575f80fd5b506105076109c7366004615660565b611db1565b3480156109d7575f80fd5b506104de6109e6366004615438565b600954600160c01b90046001600160401b03161190565b348015610a08575f80fd5b50610ba4610a17366004615438565b60408051610200810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c081018290526101e0810191909152505f90815260066020908152604091829020825161020081018452815463ffffffff8082168352600160201b808304821695840195909552600160401b808304821696840196909652600160601b820481166060840152600160801b820481166080840152600160a01b808304821660a0850152600160c01b8304821660c085015260ff600160e01b909304831660e085015260018501548083166101008601529586049091166101208401526001600160a01b0395909404851661014083015260028301548516610160830152600383015485166101808301526004909201549384166101a0820152918304811615156101c0830152600160a81b90920490911615156101e082015290565b6040516104ea919061568a565b348015610bbc575f80fd5b506106617f000000000000000000000000d99cc66fec647e68294c6477b40fc7e0f6f618d081565b348015610bef575f80fd5b50610507610bfe3660046157c4565b611e00565b348015610c0e575f80fd5b50610507610c1d366004615438565b611e74565b348015610c2d575f80fd5b50610507610c3c366004615438565b611eb3565b348015610c4c575f80fd5b506106617f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da81565b610507610c82366004615438565b611ef2565b348015610c92575f80fd5b50610ca6610ca1366004615438565b611fb9565b604080519283526020830191909152016104ea565b348015610cc6575f80fd5b506105567fc72a21b38830f4d6418a239e17db78b945cc7cfee674bac97fd596eaf043847881565b348015610cf9575f80fd5b50610507610d0836600461540a565b6120e7565b348015610d18575f80fd5b506105565f80516020615f1483398151915281565b348015610d38575f80fd5b50610507610d47366004615660565b612135565b348015610d57575f80fd5b506104de610d663660046155f9565b612184565b348015610d76575f80fd5b50600954600160c01b90046001600160401b0316610556565b348015610d9a575f80fd5b506105567f59911a6aa08a72fe3824aec4500dc42335c6d0702b6d5c5c72ceb265a0de930281565b348015610dcd575f80fd5b50610507612229565b348015610de1575f80fd5b50610507610df0366004615660565b612301565b348015610e00575f80fd5b50610507610e0f3660046153cd565b612350565b348015610e1f575f80fd5b50610507610e2e36600461540a565b6125bc565b348015610e3e575f80fd5b50610507610e4d366004615438565b61260a565b348015610e5d575f80fd5b50610507610e6c366004615438565b6126d1565b348015610e7c575f80fd5b50610556610e8b366004615438565b5f9081526006602052604090205463ffffffff600160201b82048116918116919091031690565b348015610ebd575f80fd5b50610661610ecc3660046154eb565b6126dc565b348015610edc575f80fd5b50610507612714565b348015610ef0575f80fd5b506104de610eff36600461540a565b612733565b348015610f0f575f80fd5b50610507610f1e3660046154eb565b612769565b348015610f2e575f80fd5b50610507610f3d366004615438565b612794565b348015610f4d575f80fd5b507f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da610661565b348015610f7f575f80fd5b50600954604080516001600160401b03600160401b8404811682528084166020830152600160801b909304909216908201526060016104ea565b348015610fc4575f80fd5b50610507610fd336600461581c565b612972565b348015610fe3575f80fd5b50610507610ff2366004615882565b6129d4565b348015611002575f80fd5b50610556611011366004615438565b5f90815260066020526040902054600160401b900463ffffffff1690565b61050761103d3660046158de565b612da0565b34801561104d575f80fd5b506105565f81565b348015611060575f80fd5b506105565f1981565b348015611074575f80fd5b5061055661108336600461596c565b612eda565b348015611093575f80fd5b506105076110a23660046159bd565b61313c565b3480156110b2575f80fd5b506105567f000000000000000000000000000000000000000000000000000000000000000481565b3480156110e5575f80fd5b506105567fb3e25b5404b87e5a838579cb5d7481d61ad96ee284d38ec1e97c07ba64e7f6fc81565b348015611118575f80fd5b50610661611127366004615438565b613259565b348015611137575f80fd5b506104de6132a6565b34801561114b575f80fd5b5061115f61115a366004615438565b6132d6565b604080519889526020890197909752958701949094526060860192909252608085015260a084015260c083015260e0820152610100016104ea565b3480156111a5575f80fd5b506111ae61345b565b6040516001600160401b0390911681526020016104ea565b3480156111d1575f80fd5b506105076111e036600461581c565b61348d565b3480156111f0575f80fd5b506108996111ff366004615a66565b6135d7565b34801561120f575f80fd5b5061050761121e366004615aad565b613944565b34801561122e575f80fd5b5061055661123d366004615438565b613af3565b34801561124d575f80fd5b50600554610556565b348015611261575f80fd5b5061050761127036600461540a565b613b2a565b348015611280575f80fd5b506105567f000000000000000000000000000000000000000000000000000000000000000581565b3480156112b3575f80fd5b506106617f000000000000000000000000c1d0b3de6792bf6b4b37eccdcc24e45978cfd2eb81565b3480156112e6575f80fd5b506106617f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe8481565b348015611319575f80fd5b506105567f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea0981565b348015610507575f80fd5b348015611357575f80fd5b50610507611366366004615438565b613b5c565b348015611376575f80fd5b506105076113853660046159bd565b613b93565b348015611395575f80fd5b506106617f00000000000000000000000006cd61045f958a209a0f8d746e103ecc625f419381565b5f6001600160e01b03198216635a05180f60e01b14806113e157506113e182613c74565b92915050565b7f2fc10cc8ae19568712f7a176fb4978616a610650813c9d05326c34abb62749c761141181613ca8565b611419613cb2565b50565b5f80516020615f1483398151915261143381613ca8565b60028311156114555760405163b4fa3fb360e01b815260040160405180910390fd5b63ffffffff82111561147a5760405163b4fa3fb360e01b815260040160405180910390fd5b61148384613d07565b5f8481526006602052604081209084900361149c575f92505b805464ffffffffff60c01b1916600160e01b60ff86160263ffffffff60c01b191617600160c01b63ffffffff851602178155604080518581526020810185905286917ff92eb109ce5b449e9b121c352c6aeb4319538a90738cb95d84f08e41274e92d2910160405180910390a2611513855f613d3d565b61151b613f92565b5050505050565b60405162d74f0b60e71b815260066004820152602481018390526001600160a01b038216604482015273e4d5a7be8d7c3db15755061053f5a49b6a67fffc90636ba78580906064015b5f6040518083038186803b158015611581575f80fd5b505af4158015611593573d5f803e3d5ffd5b505050505050565b5f6115a582613d07565b7f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e4286001600160a01b0316639a7b05086115dd84613fd2565b6040518263ffffffff1660e01b81526004016115fb91815260200190565b602060405180830381865afa158015611616573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113e19190615ac8565b5f8281525f80516020615f34833981519152602052604090206001015461166081613ca8565b61166a8383614021565b50505050565b6001600160a01b03811633146116995760405163334bd91960e11b815260040160405180910390fd5b6116a38282614076565b505050565b7fe85fdec10fe0f93d0792364051df7c3d73e37c17b3a954bffe593960e3cd30126116d281613ca8565b7f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da5f5b8381101561151b575f85858381811061171057611710615adf565b90506020020135905061172281613d07565b6040516325d9153960e11b8152600481018290525f906001600160a01b03851690634bb22a72906024016020604051808303815f875af1158015611768573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061178c9190615b00565b905080156117c85760405182907ef4fe19c0404d2fbb58da6f646c0a3ee5a6994a034213bbd22b072ed1ca5c27905f90a26117c8826001613d3d565b50506001016116f5565b5f6117ff826117e0856140c2565b906001600160801b03165f908152600191909101602052604090205490565b9392505050565b7f59911a6aa08a72fe3824aec4500dc42335c6d0702b6d5c5c72ceb265a0de930261183081613ca8565b61183984613d07565b815f036118595760405163162908e360e11b815260040160405180910390fd5b5f61186385613fd2565b6040516307a994c760e01b8152600481018290529091505f906001600160a01b037f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e42816906307a994c790602401602060405180830381865afa1580156118cb573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118ef9190615ac8565b90507f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da6001600160a01b031663dcab7f838761192b8488615b2f565b6040516001600160e01b031960e085901b168152600481019290925260248201526044015f604051808303815f87803b158015611966575f80fd5b505af1158015611978573d5f803e3d5ffd5b505060408051888152602081018890528993507feec4d6dbe34149c6728a9638eca869d0e5a7fcd85c7a96178f7e9780b4b7fe4b92500160405180910390a2611593866001613d3d565b7f59911a6aa08a72fe3824aec4500dc42335c6d0702b6d5c5c72ceb265a0de93026119ec81613ca8565b6119f583613d07565b7f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da60405163d963ae5560e01b815260048101859052602481018490526001600160a01b03919091169063d963ae55906044015f604051808303815f87803b158015611a5e575f80fd5b505af1158015611a70573d5f803e3d5ffd5b50505050827f1e7ebd3c5f4de9502000b6f7e6e7cf5d4ecb27d6fe1778e43fb9d1d0ca87d0e783604051611aa691815260200190565b60405180910390a26116a3836001613d3d565b600954606090600160c01b90046001600160401b03168084101580611adc575082155b15611af6575050604080515f8152602081019091526113e1565b5f611b018583615b56565b8410611b1657611b118583615b56565b611b18565b835b9050806001600160401b03811115611b3257611b32615b42565b604051908082528060200260200182016040528015611b5b578160200160208202803683370190505b5092505f5b8351811015611b9857611b738187615b2f565b848281518110611b8557611b85615adf565b6020908102919091010152600101611b60565b50505092915050565b606080611baf858585614106565b611bb883614143565b9092509050611bcb85858585855f6141e9565b935093915050565b604051631f46d51760e01b8152600660048201526024810182905273e4d5a7be8d7c3db15755061053f5a49b6a67fffc90631f46d517906044015b5f6040518083038186803b158015611c24575f80fd5b505af415801561151b573d5f803e3d5ffd5b611c3e614277565b73a74528edc289b1a597faf83fcff7eff871cc01d96352d8bfc26040518163ffffffff1660e01b81526004015f6040518083038186803b158015611c80575f80fd5b505af415801561166a573d5f803e3d5ffd5b5f80516020615f14833981519152611ca981613ca8565b611cb286613d07565b6040516344dab94960e01b81526001600160a01b037f00000000000000000000000006cd61045f958a209a0f8d746e103ecc625f419316906344dab94990611d04908990889088908890600401615b91565b5f604051808303815f87803b158015611d1b575f80fd5b505af1158015611d2d573d5f803e3d5ffd5b50505050505050505050565b5f805f611d45846140c2565b546001600160801b0380821696600160801b90920416945092505050565b5f611d8c7fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a025490565b905090565b6060611d9e848484614106565b611da98484846142a0565b949350505050565b611db9614277565b604051635c654ad960e01b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d990635c654ad99060440161156b565b5f80516020615f14833981519152611e1781613ca8565b611e2086613d07565b604051632122c77b60e21b81526001600160a01b037f00000000000000000000000006cd61045f958a209a0f8d746e103ecc625f4193169063848b1dec90611d049089908990899089908990600401615bbb565b60405163612b8c3b60e11b8152600660048201526024810182905273e4d5a7be8d7c3db15755061053f5a49b6a67fffc9063c257187690604401611c0e565b60405163c990450f60e01b8152600660048201526024810182905273e4d5a7be8d7c3db15755061053f5a49b6a67fffc9063c990450f90604401611c0e565b611efc813361433b565b6040516315b5c47760e01b8152600481018290526001600160a01b037f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da16906315b5c4779034906024015f604051808303818588803b158015611f5d575f80fd5b505af1158015611f6f573d5f803e3d5ffd5b5050505050807fb1858b4c2ab6242521725a8f7350a6cb22ad4ecae009c9b63ef114baffb054be34604051611fa691815260200190565b60405180910390a2611419816001613d3d565b5f80828103611fcc57505f928392509050565b5f611fd56143ac565b90505f80805b7f000000000000000000000000000000000000000000000000000000000000000581116120de5761200b816140c2565b6040516304ada34360e41b8152600481018290526006602482015260448101899052606481018690529093506001909101905f90819081908190736eff460627b6798c2907409ea2fdfb287eaa2e5590634ada343090608401608060405180830381865af415801561207f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120a39190615be8565b93509350935093505f8411156120bd579883019885830198505b806120cb57505050506120de565b8186019550818b039a5050505050611fdb565b50505050915091565b6040516317a9a2c160e11b815260066004820152602481018390526001600160a01b038216604482015273e4d5a7be8d7c3db15755061053f5a49b6a67fffc90632f5345829060640161156b565b61213d614277565b6040516340cea66360e11b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d99063819d4cc69060440161156b565b5f61218e86613d07565b60405163d404037960e01b81526001600160a01b037f00000000000000000000000006cd61045f958a209a0f8d746e103ecc625f4193169063d4040379906121e0908990889088908890600401615b91565b602060405180830381865afa1580156121fb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061221f9190615b00565b9695505050505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff1680612272575080546001600160401b03808416911610155b156122905760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b038316908117600160401b1782555f8080556004819055600355815460ff60401b191682556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15050565b612309614277565b604051638980f11f60e01b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d990638980f11f9060440161156b565b61235a833361433b565b5f8381526006602052604090208054600160401b900463ffffffff1683101561239657604051635caf530f60e11b815260040160405180910390fd5b80545f906123af9086908690869063ffffffff166143e2565b90505f6123bb86613fd2565b90505f847f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e4286001600160a01b031663f42d7db5846040518263ffffffff1660e01b815260040161240d91815260200190565b602060405180830381865afa158015612428573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061244c9190615ac8565b6124569190615c26565b90508015612508577f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da604051632207e80f60e21b815260048101899052602481018390526001600160a01b03919091169063881fa03c906044015f604051808303815f87803b1580156124c7575f80fd5b505af11580156124d9573d5f803e3d5ffd5b50506040518992507f1cbb8dafbedbdf4f813a8ed1f50d871def63e1104f8729b677af57905eda90f691505f90a25b835463ffffffff191663ffffffff841617845560405183815287907fdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f09060200160405180910390a2835463ffffffff60601b1916600160601b63ffffffff85160217845560405183815287907f947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd9060200160405180910390a26125ab875f613d3d565b6125b3613f92565b50505050505050565b604051632a5a705b60e01b815260066004820152602481018390526001600160a01b038216604482015273e4d5a7be8d7c3db15755061053f5a49b6a67fffc90632a5a705b9060640161156b565b5f80516020615f1483398151915261262181613ca8565b604051638fcb4e5b60e01b81526001600160a01b037f000000000000000000000000d99cc66fec647e68294c6477b40fc7e0f6f618d081166004830152602482018490527f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe841690638fcb4e5b906044016020604051808303815f875af11580156126ad573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116a39190615ac8565b611419816001613d3d565b5f8281527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000602081905260408220611da9908461456c565b5f80516020615f1483398151915261272b81613ca8565b611419613f92565b5f9182525f80516020615f34833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f80516020615f1483398151915261278081613ca8565b61278c83836001614577565b6116a3613f92565b5f8181526006602052604090206004810154600160a81b900460ff16156127ce57604051634d5bd9a760e01b815260040160405180910390fd5b5f6127d883613fd2565b90505f807f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e4286001600160a01b031663decfec56846040518263ffffffff1660e01b815260040161282a91815260200190565b6040805180830381865afa158015612844573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128689190615c50565b915091507f00000000000000000000000000000000000000000000000000000000000000058263ffffffff16036128b25760405163bba5f23f60e01b815260040160405180910390fd5b6001840154600160201b900463ffffffff165f8190036128e55760405163ebe0edcd60e01b815260040160405180910390fd5b845463ffffffff600160401b90910481169083168110612918576040516327da251f60e21b815260040160405180910390fd5b5f61293863ffffffff841661292d8487615c81565b63ffffffff1661469a565b905061294b888663ffffffff16836146af565b60048701805460ff60a81b1916600160a81b179055612968613f92565b5050505050505050565b5f80516020615f1483398151915261298981613ca8565b5f61299686868686614763565b90505f5b818110156129cb576008810287013560c01c6010820286013560801c6129c182825f614577565b505060010161299a565b50611593613f92565b7f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea096129fe81613ca8565b5f805b83811015612d91575f858583818110612a1c57612a1c615adf565b905060600201803603810190612a329190615cd2565b9050612a40815f0151613d07565b80515f9081526006602090815260409091208054918301519091600160401b900463ffffffff1611612a8557604051635caf530f60e11b815260040160405180910390fd5b815160208301515f9160801b175f8181526007602052604090205490915060ff1615612ab357505050612d89565b5f8181526007602090815260408220805460ff19166001908117909155845463ffffffff600160201b808304821684019091160267ffffffff0000000019909116178555855191860151612b0792916142a0565b9050835f01517f9bc54857932b6f10bb750fdad91f736b82dd4de202ed5c2f9f076773bb5b3fb78560200151866040015184604051612b4893929190615d0c565b60405180910390a2835160405163e83ba79d60e01b8152600197505f91829182916001600160a01b037f00000000000000000000000006cd61045f958a209a0f8d746e103ecc625f4193169163e83ba79d91612ba8918890600401615d2a565b60c060405180830381865afa158015612bc3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612be79190615dbb565b80516020015190915015612c0b578051516001600160f81b03169290920191600191505b80602001516020015115612c32576020810151516001600160f81b03169290920191600191505b7f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da828015612c6d57506040820151516001600160f81b031615155b15612ce4578751604080840151519051632207e80f60e21b81526001600160a01b0384169263881fa03c92612cb6926004019182526001600160f81b0316602082015260400190565b5f604051808303815f87803b158015612ccd575f80fd5b505af1158015612cdf573d5f803e3d5ffd5b505050505b87604001516801bc16d674ec8000001115612d0e5787604001516801bc16d674ec80000003840193505b8315612d7457875160405163e5220e3f60e01b81526004810191909152602481018590526001600160a01b0382169063e5220e3f906044015f604051808303815f87803b158015612d5d575f80fd5b505af1158015612d6f573d5f803e3d5ffd5b505050505b8751612d80905f613d3d565b50505050505050505b600101612a01565b50801561166a5761166a613f92565b612da86147d6565b612db286886147fc565b5f7f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da6040516358a46db560e11b815260048101899052602481018890529091506001600160a01b0382169063b148db6a90604401602060405180830381865afa158015612e21573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612e459190615ac8565b341015612e655760405163162908e360e11b815260040160405180910390fd5b3415612ecc57604051630b96641560e21b81526001600160a01b03898116600483015260248201899052821690632e5990549034906044015f604051808303818588803b158015612eb4575f80fd5b505af1158015612ec6573d5f803e3d5ffd5b50505050505b612968878787878787614870565b5f7fc72a21b38830f4d6418a239e17db78b945cc7cfee674bac97fd596eaf0438478612f0581613ca8565b612f0d6147d6565b6001600160a01b038516612f375760405160016232750f60e21b0319815260040160405180910390fd5b600954600160c01b90046001600160401b03169150612f5582614a3e565b5f828152600660209081526040822091908190612f7490880188615aad565b6001600160a01b031614612f9457612f8f6020870187615aad565b612f96565b865b90505f80612faa6040890160208a01615aad565b6001600160a01b031614612fcd57612fc86040880160208901615aad565b612fcf565b875b6001840180547fffffffff0000000000000000000000000000000000000000ffffffffffffffff16600160401b6001600160a01b03868116919091029190911790915560038501805473ffffffffffffffffffffffffffffffffffffffff1916918316919091179055905061304a6060880160408901615e0a565b156130655760048301805460ff60a01b1916600160a01b1790555b6009805460016001600160401b03600160c01b80840482169290920116026001600160c01b039091161790556001600160a01b03808216908316867ff17baf73d46b0a80157c3ea3dda1bf081a702732d53ff1720f85e55d9f0997c06130d160608c0160408d01615e0a565b604051901515815260200160405180910390a46001600160a01b03861615613129576040516001600160a01b0387169086907f67334334c388385e5f244703f8a8b28b7f4ffe52909130aca69bc62a8e27f09a905f90a35b613131613f92565b505050509392505050565b6131446147d6565b61314e87896147fc565b5f7f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da604051632884698160e01b8152600481018a9052602481018990529091505f906001600160a01b03831690632884698190604401602060405180830381865afa1580156131bf573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131e39190615ac8565b9050801561324b57604051637bcb377f60e11b81526001600160a01b0383169063f7966efe9061321d908d908d9086908990600401615e25565b5f604051808303815f87803b158015613234575f80fd5b505af1158015613246573d5f803e3d5ffd5b505050505b611d2d898989898989614870565b5f8181526006602052604081206004810154600160a01b900460ff1661328c5760038101546001600160a01b03166117ff565b60010154600160401b90046001600160a01b031692915050565b5f6132cf7fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a025490565b4210905090565b5f805f805f805f806132e789613d07565b5f898152600660205260408120907f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da6001600160a01b0316639c5161028c6040518263ffffffff1660e01b815260040161334391815260200190565b602060405180830381865afa15801561335e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906133829190615ac8565b82549091505f906133a29063ffffffff600160401b820481169116615c81565b63ffffffff1690508082111561340157825460029b5063ffffffff600160201b8204811681831603168390039a50600160e01b900460ff16156133fc5782546133f9908b90600160c01b900463ffffffff1661469a565b99505b61341f565b8254600160e01b810460ff169b50600160c01b900463ffffffff1699505b505060018101549054989a9799505f98899889985063ffffffff9283169750600160401b820483169650600160a01b9091049091169350915050565b5f611d8c7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00546001600160401b031690565b5f80516020615f148339815191526134a481613ca8565b5f6134b186868686614763565b90505f5b818110156129cb576008810287013560c01c6010820286013560801c6134da82613d07565b5f8281526006602052604090208054600160601b900463ffffffff168210613515576040516388e1a28160e01b815260040160405180910390fd5b8054600160401b900463ffffffff16821015613544576040516388e1a28160e01b815260040160405180910390fd5b805463ffffffff60601b1916600160601b63ffffffff84160217815560405182815283907f947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd9060200160405180910390a260405183907fe5725d045d5c47bd1483feba445e395dc8647486963e6d54aad9ed03ff7d6ce6905f90a26135c9835f613d3d565b5050508060010190506134b5565b6060805f80516020615f148339815191526135f181613ca8565b6135fa86614143565b9093509150851561393b57855f80805b7f0000000000000000000000000000000000000000000000000000000000000005811180613636575083155b6138b557613643816140c2565b91506001015f61366c8380546001600160801b03165f9081526001909101602052604090205490565b90505b80156138af575f6136808260c01c90565b6001600160401b039081165f8181526006602052604081208054929450608086901c90931692916136c8906136c290600160a01b900463ffffffff168561469a565b8a61469a565b9050808911806136d757508281145b156137175760018201805463ffffffff600160201b80830482168790039091160267ffffffff000000001990911617905561371187614a6a565b50613776565b60018201805463ffffffff600160201b808304821685900382160267ffffffff00000000199092169190911790915561375690869083860390614ac916565b87546001600160801b03165f908152600189016020526040902081905594505b805f03613786575050505061388e565b81546137a4908590600160401b900463ffffffff16838f8f8d6141e9565b815463ffffffff600160401b808304821684019182169081026bffffffff000000000000000019909316929092178455604051918252988201989085907f24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b789060200160405180910390a2825463ffffffff600160a01b808304821685900391821690810263ffffffff60a01b199093169290921785556040519182529086907ff9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed339060200160405180910390a2828b039a508a5f03613887575050505050506138af565b5050505050505b5081546001600160801b03165f90815260018301602052604090205461366f565b5061360a565b8983146138d557604051630bc9ea5560e21b815260040160405180910390fd5b600980546001600160401b03600160801b80830482168e9003821602808216828416178e0190911667ffffffffffffffff1990911677ffffffffffffffff0000000000000000ffffffffffffffff1990921691909117179055613936613f92565b505050505b50935093915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff168061398d575080546001600160401b03808416911610155b156139ab5760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b1781556001600160a01b0383166139f357604051633ef39b8160e01b815260040160405180910390fd5b6139fb614aee565b613a055f84614021565b50613a9c5f80516020615f148339815191527f000000000000000000000000c1d0b3de6792bf6b4b37eccdcc24e45978cfd2eb6001600160a01b031663ef6c064c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613a73573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a979190615e8d565b614021565b50613aa75f19614af6565b805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e823717059320006020819052604082206117ff90614b45565b5f8281525f80516020615f348339815191526020526040902060010154613b5081613ca8565b61166a8383614076565b565b7f139c2898040ef16910dc9f44dc697df79363da767d8bc92f2e310312b816e46d613b8681613ca8565b613b8f82614af6565b5050565b613b9b6147d6565b613ba587896147fc565b5f7f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da6040516358a46db560e11b8152600481018a9052602481018990529091505f906001600160a01b0383169063b148db6a90604401602060405180830381865afa158015613c16573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c3a9190615ac8565b9050801561324b5760405163263f69e960e11b81526001600160a01b03831690634c7ed3d29061321d908d908d9086908990600401615e25565b5f6001600160e01b03198216637965db0b60e01b14806113e157506301ffc9a760e01b6001600160e01b03198316146113e1565b6114198133614b4e565b613cba614b8b565b427fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a02556040517f62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9905f90a1565b600954600160c01b90046001600160401b0316811015613d245750565b604051633ed893db60e21b815260040160405180910390fd5b5f8281526006602052604081208054909163ffffffff600160401b8304811692613d70918491600160601b900416615c81565b63ffffffff1690505f7f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da6001600160a01b03166301a5e9e3876040518263ffffffff1660e01b8152600401613dc791815260200190565b602060405180830381865afa158015613de2573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613e069190615ac8565b84549091505f90613e1e90859063ffffffff16615c81565b63ffffffff169050808210613e35575f9250613e6a565b8454613e509063ffffffff600160601b820481169116615c81565b63ffffffff16821115613e6a57613e678282615b56565b92505b508354600160e01b900460ff1615801590613e8457505f82115b15613ed457835463ffffffff600160201b820481168503811691613ed091600160c01b909104168210613eb7575f613eca565b8554600160c01b900463ffffffff168290035b8461469a565b9250505b8354600160a01b900463ffffffff1682146115935783546009805467ffffffffffffffff60801b198116600160a01b9384900463ffffffff908116600160801b938490046001600160401b039081169190910388011690920217909155855463ffffffff60a01b191690841690910217845560405182815286907ff9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed339060200160405180910390a28415613f8957613f89613f92565b61159386614bb0565b60058054600101908190556040519081527f7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa9060200160405180910390a1565b604051630569b94760e01b8152600481018290525f906001600160a01b037f0000000000000000000000004d72bff1beac69925f8bd12526a39baab069e5da1690630569b947906024016115fb565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e823717059320008161404e8585614d73565b90508015611da9575f85815260208390526040902061406d9085614e1b565b50949350505050565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000816140a38585614e2f565b90508015611da9575f85815260208390526040902061406d9085614ea8565b5f7f000000000000000000000000000000000000000000000000000000000000000482036140f257506001919050565b505f8181526020819052604090205b919050565b5f8381526006602052604090205463ffffffff166141248284615b2f565b11156116a357604051635caf530f60e11b815260040160405180910390fd5b606080614151603084615c26565b6001600160401b0381111561416857614168615b42565b6040519080825280601f01601f191660200182016040528015614192576020820181803683370190505b5061419e606085615c26565b6001600160401b038111156141b5576141b5615b42565b6040519080825280601f01601f1916602001820160405280156141df576020820181803683370190505b5091509150915091565b5f805b858110156129685761422a88614202838a615b2f565b7f059e9c54cf92ba46cc39c6b4acd51d5116e9d49fabee6193530ea918b54be94a9190614ebc565b60018082015460801c85840160308181028a0190810192909252835460209283015260028401546060918202890192830152600384015460408301526004840154910152909250016141ec565b613b5a7fb3e25b5404b87e5a838579cb5d7481d61ad96ee284d38ec1e97c07ba64e7f6fc613ca8565b60605f6142ae603084615c26565b6001600160401b038111156142c5576142c5615b42565b6040519080825280601f01601f1916602001820160405280156142ef576020820181803683370190505b5091505f5b838110156143325761430a866142028388615b2f565b9150603081026020840101600183015460801c601082015282548152506001810190506142f4565b50509392505050565b5f82815260066020526040902060010154600160401b90046001600160a01b03168061437a57604051633ed893db60e21b815260040160405180910390fd5b816001600160a01b0316816001600160a01b0316146116a35760405163743a3f7960e11b815260040160405180910390fd5b7f6e38e7eaa4307e6ee6c66720337876ca65012869fbef035f57219354c17284005f818152815c602052604090209081815d5090565b5f8215806143f85750816143f68486615b2f565b115b80614406575063ffffffff82115b156144245760405163575697ff60e01b815260040160405180910390fd5b604080516030808252606082019092525f91829182918291906020820181803683370190505090508787015b8881111561455e576144867f059e9c54cf92ba46cc39c6b4acd51d5116e9d49fabee6193530ea918b54be94a8b5f198401614ebc565b9450600185015460801c603083015284546020830152868110156144f8576144d27f059e9c54cf92ba46cc39c6b4acd51d5116e9d49fabee6193530ea918b54be94a8b5f198a01614ebc565b93505f92505b60058310156144f45782840154838601556001830192506144d8565b8394505b5f92505b6005831015614515575f838601556001830192506144fc565b600187039650600181039050897fea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b983604051614551919061564e565b60405180910390a2614450565b509498975050505050505050565b5f6117ff8383614ef3565b61458083613d07565b5f838152600660205260409020600181015463ffffffff168084036145a6575050505050565b8154600160401b900463ffffffff168411156145d55760405163cc11217f60e01b815260040160405180910390fd5b821580156145e857508063ffffffff1684105b15614606576040516371a4bd1560e01b815260040160405180910390fd5b600980546fffffffffffffffff000000000000000019811663ffffffff848116600160401b938490046001600160401b03908116919091038901169092021790915560018301805463ffffffff191691861691909117905560405184815285907f0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c2201669060200160405180910390a25050505050565b5f8183106146a857816117ff565b5090919050565b5f8381526006602052604090206001810180548391906004906146e0908490600160201b900463ffffffff16615ea8565b92506101000a81548163ffffffff021916908363ffffffff1602179055505f614708846140c2565b905061471e818663ffffffff80871690614f1916565b5060405163ffffffff84168152859085907fdc891a44aee443f7f65d1abc5710a05ef241c0c5d7a62f12671522f3c14852bc9060200160405180910390a35050505050565b5f61476f600885615ed9565b61477a601084615ed9565b141580614790575061478d600885615eec565b15155b806147a457506147a1601083615eec565b15155b156147c25760405163319c9a2160e21b815260040160405180910390fd5b6147cd600885615ed9565b95945050505050565b6147de6132a6565b15613b5a57604051630286f07360e31b815260040160405180910390fd5b336001600160a01b0382160361481657613b8f823361433b565b61483f7fc72a21b38830f4d6418a239e17db78b945cc7cfee674bac97fd596eaf0438478613ca8565b3361484983614f89565b6001600160a01b031614613b8f576040516310b922ef60e21b815260040160405180910390fd5b61487986614fbf565b5f8681526006602052604081208054909163ffffffff9091169061489c89613fd2565b60405163014dddeb60e51b8152600481018290529091505f906001600160a01b037f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e42816906329bbbd6090602401602060405180830381865afa158015614904573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906149289190615ac8565b8454909150600160201b900463ffffffff168984010381101561495e57604051630911e76760e41b815260040160405180910390fd5b5f61496e8b858c8c8c8c8c614feb565b8554909150600160601b900463ffffffff168085036149e257855463ffffffff60601b1916600160601b918c0163ffffffff8116928302919091178755604051918252908c907f947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd9060200160405180910390a25b855463ffffffff191663ffffffff83161786556040518281528c907fdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f09060200160405180910390a25050614a368a5f613d3d565b611d2d613f92565b7f1b07bc0838fdc4254cbabb5dd0c94d936f872c6758547168d513d8ad1dc3a500613b8f81833361517f565b80546001600160801b03165f90815260018201602052604090205480614aa3576040516363c3654960e01b815260040160405180910390fd5b81546fffffffffffffffffffffffffffffffff19166001600160801b0382161790915590565b60801b67ffffffffffffffff60801b1667ffffffffffffffff60801b19919091161790565b613b5a615195565b614afe6147d6565b805f03614b1e5760405163ad58bfc760e01b815260040160405180910390fd5b5f5f198203614b2f57505f19614b3c565b614b398242615b2f565b90505b613b8f816151de565b5f6113e1825490565b614b588282612733565b613b8f5760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440160405180910390fd5b614b936132a6565b613b5a5760405163b047186b60e01b815260040160405180910390fd5b5f614bba82613fd2565b90505f807f0000000000000000000000009d28ad303c90df524ba960d7a2dac56dcc31e4286001600160a01b031663decfec56846040518263ffffffff1660e01b8152600401614c0c91815260200190565b6040805180830381865afa158015614c26573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190614c4a9190615c50565b5f868152600660205260409020805460018201549395509193509163ffffffff600160a01b909204821691600160201b90910416808211614c8e5750505050505050565b8082037f000000000000000000000000000000000000000000000000000000000000000563ffffffff87161015614d3c57835463ffffffff600160401b90910481168301908181169087161115614d3a578086035f614cf663ffffffff80861690841661469a565b9050614d098b8a63ffffffff16836146af565b60048701549381900393600160a81b900460ff16614d375760048701805460ff60a81b1916600160a81b1790555b50505b505b63ffffffff81161561296857612968887f0000000000000000000000000000000000000000000000000000000000000005836146af565b5f5f80516020615f34833981519152614d8c8484612733565b614e0b575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055614dc13390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506113e1565b5f9150506113e1565b5092915050565b5f6117ff836001600160a01b038416615280565b5f5f80516020615f34833981519152614e488484612733565b15614e0b575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506113e1565b5f6117ff836001600160a01b0384166152cc565b6040805160208082019590955280820193909352606080840192909252805180840390920182526080909201909152805191012090565b5f825f018281548110614f0857614f08615adf565b905f5260205f200154905092915050565b8254600160801b908190046001600160801b039081165f8181526001808801602052604090912060809590951b67ffffffffffffffff60801b1660c09690961b6001600160c01b031916959095179085011792839055845480821690839004821690940116029190911790915590565b5f7f1b07bc0838fdc4254cbabb5dd0c94d936f872c6758547168d513d8ad1dc3a5006117ff81845f918252602052604090205c90565b7f1b07bc0838fdc4254cbabb5dd0c94d936f872c6758547168d513d8ad1dc3a500613b8f81835f61517f565b5f851580615005575063ffffffff6150038789615b2f565b115b156150235760405163575697ff60e01b815260040160405180910390fd5b60308602841415806150385750606086028214155b156150565760405163251f56a160e21b815260040160405180910390fd5b604080516030808252606082019092525f91829182916020820181803683370190505090505f5b8981101561516f576150b07f059e9c54cf92ba46cc39c6b4acd51d5116e9d49fabee6193530ea918b54be94a8d8d614ebc565b60308281028b01601081013591850182905235602085018190529195501715925082156150f057604051630f35a7eb60e21b815260040160405180910390fd5b60208201518455603082015160801b60018501556060810287018035600286015560208101356003860155604081013560048601555060018101905060018b019a508b7fc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a8983604051615162919061564e565b60405180910390a261507d565b50989a9950505050505050505050565b5f83815260208390526040902081815d50505050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16613b5a57604051631afcd79f60e31b815260040160405180910390fd5b6152077fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a02829055565b5f198103615247576040515f1981527f32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e906020015b60405180910390a150565b7f32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e6152724283615b56565b60405190815260200161523c565b5f8181526001830160205260408120546152c557508154600181810184555f8481526020808220909301849055845484825282860190935260409020919091556113e1565b505f6113e1565b5f8181526001830160205260408120548015614e0b575f6152ee600183615b56565b85549091505f9061530190600190615b56565b9050808214615360575f865f01828154811061531f5761531f615adf565b905f5260205f200154905080875f01848154811061533f5761533f615adf565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061537157615371615eff565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f9055600193505050506113e1565b5f602082840312156153b6575f80fd5b81356001600160e01b0319811681146117ff575f80fd5b5f805f606084860312156153df575f80fd5b505081359360208301359350604090920135919050565b6001600160a01b0381168114611419575f80fd5b5f806040838503121561541b575f80fd5b82359150602083013561542d816153f6565b809150509250929050565b5f60208284031215615448575f80fd5b5035919050565b5f8060208385031215615460575f80fd5b82356001600160401b0380821115615476575f80fd5b818501915085601f830112615489575f80fd5b813581811115615497575f80fd5b8660208260051b85010111156154ab575f80fd5b60209290920196919550909350505050565b5f80604083850312156154ce575f80fd5b8235915060208301356001600160801b038116811461542d575f80fd5b5f80604083850312156154fc575f80fd5b50508035926020909101359150565b602080825282518282018190525f9190848201906040850190845b8181101561554257835183529284019291840191600101615526565b50909695505050505050565b5f81518084525f5b8181101561557257602081850181015186830182015201615556565b505f602082860101526020601f19601f83011685010191505092915050565b604081525f6155a3604083018561554e565b82810360208401526147cd818561554e565b5f8083601f8401126155c5575f80fd5b5081356001600160401b038111156155db575f80fd5b6020830191508360208285010111156155f2575f80fd5b9250929050565b5f805f805f6080868803121561560d575f80fd5b853594506020860135935060408601356001600160401b03811115615630575f80fd5b61563c888289016155b5565b96999598509660600135949350505050565b602081525f6117ff602083018461554e565b5f8060408385031215615671575f80fd5b823561567c816153f6565b946020939093013593505050565b815163ffffffff168152610200810160208301516156b0602084018263ffffffff169052565b5060408301516156c8604084018263ffffffff169052565b5060608301516156e0606084018263ffffffff169052565b5060808301516156f8608084018263ffffffff169052565b5060a083015161571060a084018263ffffffff169052565b5060c083015161572860c084018263ffffffff169052565b5060e083015161573d60e084018260ff169052565b506101008381015163ffffffff908116918401919091526101208085015190911690830152610140808401516001600160a01b039081169184019190915261016080850151821690840152610180808501518216908401526101a080850151909116908301526101c0808401511515908301526101e0928301511515929091019190915290565b5f805f805f608086880312156157d8575f80fd5b8535945060208601356001600160401b038111156157f4575f80fd5b615800888289016155b5565b9699909850959660408101359660609091013595509350505050565b5f805f806040858703121561582f575f80fd5b84356001600160401b0380821115615845575f80fd5b615851888389016155b5565b90965094506020870135915080821115615869575f80fd5b50615876878288016155b5565b95989497509550505050565b5f8060208385031215615893575f80fd5b82356001600160401b03808211156158a9575f80fd5b818501915085601f8301126158bc575f80fd5b8135818111156158ca575f80fd5b8660206060830285010111156154ab575f80fd5b5f805f805f805f60a0888a0312156158f4575f80fd5b87356158ff816153f6565b9650602088013595506040880135945060608801356001600160401b0380821115615928575f80fd5b6159348b838c016155b5565b909650945060808a013591508082111561594c575f80fd5b506159598a828b016155b5565b989b979a50959850939692959293505050565b5f805f83850360a081121561597f575f80fd5b843561598a816153f6565b93506060601f198201121561599d575f80fd5b5060208401915060808401356159b2816153f6565b809150509250925092565b5f805f805f805f80888a036101408112156159d6575f80fd5b89356159e1816153f6565b985060208a0135975060408a0135965060608a01356001600160401b0380821115615a0a575f80fd5b615a168d838e016155b5565b909850965060808c0135915080821115615a2e575f80fd5b50615a3b8c828d016155b5565b90955093505060a0609f1982011215615a52575f80fd5b5060a0890190509295985092959890939650565b5f805f60408486031215615a78575f80fd5b8335925060208401356001600160401b03811115615a94575f80fd5b615aa0868287016155b5565b9497909650939450505050565b5f60208284031215615abd575f80fd5b81356117ff816153f6565b5f60208284031215615ad8575f80fd5b5051919050565b634e487b7160e01b5f52603260045260245ffd5b8015158114611419575f80fd5b5f60208284031215615b10575f80fd5b81516117ff81615af3565b634e487b7160e01b5f52601160045260245ffd5b808201808211156113e1576113e1615b1b565b634e487b7160e01b5f52604160045260245ffd5b818103818111156113e1576113e1615b1b565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b848152606060208201525f615baa606083018587615b69565b905082604083015295945050505050565b858152608060208201525f615bd4608083018688615b69565b604083019490945250606001529392505050565b5f805f8060808587031215615bfb575f80fd5b8451935060208501519250604085015191506060850151615c1b81615af3565b939692955090935050565b80820281158282048414176113e1576113e1615b1b565b805163ffffffff81168114614101575f80fd5b5f8060408385031215615c61575f80fd5b615c6a83615c3d565b9150615c7860208401615c3d565b90509250929050565b63ffffffff828116828216039080821115614e1457614e14615b1b565b604051606081016001600160401b0381118282101715615ccc57634e487b7160e01b5f52604160045260245ffd5b60405290565b5f60608284031215615ce2575f80fd5b615cea615c9e565b8235815260208301356020820152604083013560408201528091505092915050565b838152826020820152606060408201525f6147cd606083018461554e565b828152604060208201525f611da9604083018461554e565b5f60408284031215615d52575f80fd5b604051604081018181106001600160401b0382111715615d8057634e487b7160e01b5f52604160045260245ffd5b604052825190915081906001600160f81b0381168114615d9e575f80fd5b81526020830151615dae81615af3565b6020919091015292915050565b5f60c08284031215615dcb575f80fd5b615dd3615c9e565b615ddd8484615d42565b8152615dec8460408501615d42565b6020820152615dfe8460808501615d42565b60408201529392505050565b5f60208284031215615e1a575f80fd5b81356117ff81615af3565b5f610100820190506001600160a01b03861682528460208301528360408301528235606083015260208301356080830152604083013560ff8116808214615e6a575f80fd5b60a084015250606083013560c083015260809092013560e0909101529392505050565b5f60208284031215615e9d575f80fd5b81516117ff816153f6565b63ffffffff818116838216019080821115614e1457614e14615b1b565b634e487b7160e01b5f52601260045260245ffd5b5f82615ee757615ee7615ec5565b500490565b5f82615efa57615efa615ec5565b500690565b634e487b7160e01b5f52603160045260245ffdfebb75b874360e0bfd87f964eadd8276d8efb7c942134fc329b513032d0803e0c602dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a164736f6c6343000818000a",
}

// CsmoduleABI is the input ABI used to generate the binding from.
// Deprecated: Use CsmoduleMetaData.ABI instead.
var CsmoduleABI = CsmoduleMetaData.ABI

// CsmoduleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CsmoduleMetaData.Bin instead.
var CsmoduleBin = CsmoduleMetaData.Bin

// DeployCsmodule deploys a new Ethereum contract, binding an instance of Csmodule to it.
func DeployCsmodule(auth *bind.TransactOpts, backend bind.ContractBackend, moduleType [32]byte, lidoLocator common.Address, parametersRegistry common.Address, _accounting common.Address, exitPenalties common.Address) (common.Address, *types.Transaction, *Csmodule, error) {
	parsed, err := CsmoduleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CsmoduleBin), backend, moduleType, lidoLocator, parametersRegistry, _accounting, exitPenalties)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Csmodule{CsmoduleCaller: CsmoduleCaller{contract: contract}, CsmoduleTransactor: CsmoduleTransactor{contract: contract}, CsmoduleFilterer: CsmoduleFilterer{contract: contract}}, nil
}

// Csmodule is an auto generated Go binding around an Ethereum contract.
type Csmodule struct {
	CsmoduleCaller     // Read-only binding to the contract
	CsmoduleTransactor // Write-only binding to the contract
	CsmoduleFilterer   // Log filterer for contract events
}

// CsmoduleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CsmoduleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsmoduleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CsmoduleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsmoduleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CsmoduleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsmoduleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CsmoduleSession struct {
	Contract     *Csmodule         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CsmoduleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CsmoduleCallerSession struct {
	Contract *CsmoduleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CsmoduleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CsmoduleTransactorSession struct {
	Contract     *CsmoduleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CsmoduleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CsmoduleRaw struct {
	Contract *Csmodule // Generic contract binding to access the raw methods on
}

// CsmoduleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CsmoduleCallerRaw struct {
	Contract *CsmoduleCaller // Generic read-only contract binding to access the raw methods on
}

// CsmoduleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CsmoduleTransactorRaw struct {
	Contract *CsmoduleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCsmodule creates a new instance of Csmodule, bound to a specific deployed contract.
func NewCsmodule(address common.Address, backend bind.ContractBackend) (*Csmodule, error) {
	contract, err := bindCsmodule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Csmodule{CsmoduleCaller: CsmoduleCaller{contract: contract}, CsmoduleTransactor: CsmoduleTransactor{contract: contract}, CsmoduleFilterer: CsmoduleFilterer{contract: contract}}, nil
}

// NewCsmoduleCaller creates a new read-only instance of Csmodule, bound to a specific deployed contract.
func NewCsmoduleCaller(address common.Address, caller bind.ContractCaller) (*CsmoduleCaller, error) {
	contract, err := bindCsmodule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CsmoduleCaller{contract: contract}, nil
}

// NewCsmoduleTransactor creates a new write-only instance of Csmodule, bound to a specific deployed contract.
func NewCsmoduleTransactor(address common.Address, transactor bind.ContractTransactor) (*CsmoduleTransactor, error) {
	contract, err := bindCsmodule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CsmoduleTransactor{contract: contract}, nil
}

// NewCsmoduleFilterer creates a new log filterer instance of Csmodule, bound to a specific deployed contract.
func NewCsmoduleFilterer(address common.Address, filterer bind.ContractFilterer) (*CsmoduleFilterer, error) {
	contract, err := bindCsmodule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CsmoduleFilterer{contract: contract}, nil
}

// bindCsmodule binds a generic wrapper to an already deployed contract.
func bindCsmodule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CsmoduleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csmodule *CsmoduleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csmodule.Contract.CsmoduleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csmodule *CsmoduleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csmodule.Contract.CsmoduleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csmodule *CsmoduleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csmodule.Contract.CsmoduleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csmodule *CsmoduleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csmodule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csmodule *CsmoduleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csmodule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csmodule *CsmoduleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csmodule.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Csmodule *CsmoduleCaller) ACCOUNTING(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "ACCOUNTING")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Csmodule *CsmoduleSession) ACCOUNTING() (common.Address, error) {
	return _Csmodule.Contract.ACCOUNTING(&_Csmodule.CallOpts)
}

// ACCOUNTING is a free data retrieval call binding the contract method 0x6dc3f2bd.
//
// Solidity: function ACCOUNTING() view returns(address)
func (_Csmodule *CsmoduleCallerSession) ACCOUNTING() (common.Address, error) {
	return _Csmodule.Contract.ACCOUNTING(&_Csmodule.CallOpts)
}

// CREATENODEOPERATORROLE is a free data retrieval call binding the contract method 0x743f5105.
//
// Solidity: function CREATE_NODE_OPERATOR_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCaller) CREATENODEOPERATORROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "CREATE_NODE_OPERATOR_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CREATENODEOPERATORROLE is a free data retrieval call binding the contract method 0x743f5105.
//
// Solidity: function CREATE_NODE_OPERATOR_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleSession) CREATENODEOPERATORROLE() ([32]byte, error) {
	return _Csmodule.Contract.CREATENODEOPERATORROLE(&_Csmodule.CallOpts)
}

// CREATENODEOPERATORROLE is a free data retrieval call binding the contract method 0x743f5105.
//
// Solidity: function CREATE_NODE_OPERATOR_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCallerSession) CREATENODEOPERATORROLE() ([32]byte, error) {
	return _Csmodule.Contract.CREATENODEOPERATORROLE(&_Csmodule.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Csmodule.Contract.DEFAULTADMINROLE(&_Csmodule.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Csmodule.Contract.DEFAULTADMINROLE(&_Csmodule.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_Csmodule *CsmoduleCaller) DEPOSITSIZE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "DEPOSIT_SIZE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_Csmodule *CsmoduleSession) DEPOSITSIZE() (*big.Int, error) {
	return _Csmodule.Contract.DEPOSITSIZE(&_Csmodule.CallOpts)
}

// DEPOSITSIZE is a free data retrieval call binding the contract method 0x36bf3325.
//
// Solidity: function DEPOSIT_SIZE() view returns(uint256)
func (_Csmodule *CsmoduleCallerSession) DEPOSITSIZE() (*big.Int, error) {
	return _Csmodule.Contract.DEPOSITSIZE(&_Csmodule.CallOpts)
}

// EXITPENALTIES is a free data retrieval call binding the contract method 0xfa367c9e.
//
// Solidity: function EXIT_PENALTIES() view returns(address)
func (_Csmodule *CsmoduleCaller) EXITPENALTIES(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "EXIT_PENALTIES")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EXITPENALTIES is a free data retrieval call binding the contract method 0xfa367c9e.
//
// Solidity: function EXIT_PENALTIES() view returns(address)
func (_Csmodule *CsmoduleSession) EXITPENALTIES() (common.Address, error) {
	return _Csmodule.Contract.EXITPENALTIES(&_Csmodule.CallOpts)
}

// EXITPENALTIES is a free data retrieval call binding the contract method 0xfa367c9e.
//
// Solidity: function EXIT_PENALTIES() view returns(address)
func (_Csmodule *CsmoduleCallerSession) EXITPENALTIES() (common.Address, error) {
	return _Csmodule.Contract.EXITPENALTIES(&_Csmodule.CallOpts)
}

// FEEDISTRIBUTOR is a free data retrieval call binding the contract method 0x6910dcce.
//
// Solidity: function FEE_DISTRIBUTOR() view returns(address)
func (_Csmodule *CsmoduleCaller) FEEDISTRIBUTOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "FEE_DISTRIBUTOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FEEDISTRIBUTOR is a free data retrieval call binding the contract method 0x6910dcce.
//
// Solidity: function FEE_DISTRIBUTOR() view returns(address)
func (_Csmodule *CsmoduleSession) FEEDISTRIBUTOR() (common.Address, error) {
	return _Csmodule.Contract.FEEDISTRIBUTOR(&_Csmodule.CallOpts)
}

// FEEDISTRIBUTOR is a free data retrieval call binding the contract method 0x6910dcce.
//
// Solidity: function FEE_DISTRIBUTOR() view returns(address)
func (_Csmodule *CsmoduleCallerSession) FEEDISTRIBUTOR() (common.Address, error) {
	return _Csmodule.Contract.FEEDISTRIBUTOR(&_Csmodule.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_Csmodule *CsmoduleCaller) LIDOLOCATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "LIDO_LOCATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_Csmodule *CsmoduleSession) LIDOLOCATOR() (common.Address, error) {
	return _Csmodule.Contract.LIDOLOCATOR(&_Csmodule.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_Csmodule *CsmoduleCallerSession) LIDOLOCATOR() (common.Address, error) {
	return _Csmodule.Contract.LIDOLOCATOR(&_Csmodule.CallOpts)
}

// PARAMETERSREGISTRY is a free data retrieval call binding the contract method 0x2fc88741.
//
// Solidity: function PARAMETERS_REGISTRY() view returns(address)
func (_Csmodule *CsmoduleCaller) PARAMETERSREGISTRY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "PARAMETERS_REGISTRY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PARAMETERSREGISTRY is a free data retrieval call binding the contract method 0x2fc88741.
//
// Solidity: function PARAMETERS_REGISTRY() view returns(address)
func (_Csmodule *CsmoduleSession) PARAMETERSREGISTRY() (common.Address, error) {
	return _Csmodule.Contract.PARAMETERSREGISTRY(&_Csmodule.CallOpts)
}

// PARAMETERSREGISTRY is a free data retrieval call binding the contract method 0x2fc88741.
//
// Solidity: function PARAMETERS_REGISTRY() view returns(address)
func (_Csmodule *CsmoduleCallerSession) PARAMETERSREGISTRY() (common.Address, error) {
	return _Csmodule.Contract.PARAMETERSREGISTRY(&_Csmodule.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Csmodule *CsmoduleCaller) PAUSEINFINITELY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "PAUSE_INFINITELY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Csmodule *CsmoduleSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Csmodule.Contract.PAUSEINFINITELY(&_Csmodule.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Csmodule *CsmoduleCallerSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Csmodule.Contract.PAUSEINFINITELY(&_Csmodule.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleSession) PAUSEROLE() ([32]byte, error) {
	return _Csmodule.Contract.PAUSEROLE(&_Csmodule.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCallerSession) PAUSEROLE() ([32]byte, error) {
	return _Csmodule.Contract.PAUSEROLE(&_Csmodule.CallOpts)
}

// QUEUELEGACYPRIORITY is a free data retrieval call binding the contract method 0xa6b89b81.
//
// Solidity: function QUEUE_LEGACY_PRIORITY() view returns(uint256)
func (_Csmodule *CsmoduleCaller) QUEUELEGACYPRIORITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "QUEUE_LEGACY_PRIORITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUEUELEGACYPRIORITY is a free data retrieval call binding the contract method 0xa6b89b81.
//
// Solidity: function QUEUE_LEGACY_PRIORITY() view returns(uint256)
func (_Csmodule *CsmoduleSession) QUEUELEGACYPRIORITY() (*big.Int, error) {
	return _Csmodule.Contract.QUEUELEGACYPRIORITY(&_Csmodule.CallOpts)
}

// QUEUELEGACYPRIORITY is a free data retrieval call binding the contract method 0xa6b89b81.
//
// Solidity: function QUEUE_LEGACY_PRIORITY() view returns(uint256)
func (_Csmodule *CsmoduleCallerSession) QUEUELEGACYPRIORITY() (*big.Int, error) {
	return _Csmodule.Contract.QUEUELEGACYPRIORITY(&_Csmodule.CallOpts)
}

// QUEUELOWESTPRIORITY is a free data retrieval call binding the contract method 0xd614ae0c.
//
// Solidity: function QUEUE_LOWEST_PRIORITY() view returns(uint256)
func (_Csmodule *CsmoduleCaller) QUEUELOWESTPRIORITY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "QUEUE_LOWEST_PRIORITY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QUEUELOWESTPRIORITY is a free data retrieval call binding the contract method 0xd614ae0c.
//
// Solidity: function QUEUE_LOWEST_PRIORITY() view returns(uint256)
func (_Csmodule *CsmoduleSession) QUEUELOWESTPRIORITY() (*big.Int, error) {
	return _Csmodule.Contract.QUEUELOWESTPRIORITY(&_Csmodule.CallOpts)
}

// QUEUELOWESTPRIORITY is a free data retrieval call binding the contract method 0xd614ae0c.
//
// Solidity: function QUEUE_LOWEST_PRIORITY() view returns(uint256)
func (_Csmodule *CsmoduleCallerSession) QUEUELOWESTPRIORITY() (*big.Int, error) {
	return _Csmodule.Contract.QUEUELOWESTPRIORITY(&_Csmodule.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCaller) RECOVERERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "RECOVERER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleSession) RECOVERERROLE() ([32]byte, error) {
	return _Csmodule.Contract.RECOVERERROLE(&_Csmodule.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCallerSession) RECOVERERROLE() ([32]byte, error) {
	return _Csmodule.Contract.RECOVERERROLE(&_Csmodule.CallOpts)
}

// REPORTELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x8573e351.
//
// Solidity: function REPORT_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCaller) REPORTELREWARDSSTEALINGPENALTYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "REPORT_EL_REWARDS_STEALING_PENALTY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x8573e351.
//
// Solidity: function REPORT_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleSession) REPORTELREWARDSSTEALINGPENALTYROLE() ([32]byte, error) {
	return _Csmodule.Contract.REPORTELREWARDSSTEALINGPENALTYROLE(&_Csmodule.CallOpts)
}

// REPORTELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x8573e351.
//
// Solidity: function REPORT_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCallerSession) REPORTELREWARDSSTEALINGPENALTYROLE() ([32]byte, error) {
	return _Csmodule.Contract.REPORTELREWARDSSTEALINGPENALTYROLE(&_Csmodule.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleSession) RESUMEROLE() ([32]byte, error) {
	return _Csmodule.Contract.RESUMEROLE(&_Csmodule.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCallerSession) RESUMEROLE() ([32]byte, error) {
	return _Csmodule.Contract.RESUMEROLE(&_Csmodule.CallOpts)
}

// SETTLEELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x3f04f0c8.
//
// Solidity: function SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCaller) SETTLEELREWARDSSTEALINGPENALTYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETTLEELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x3f04f0c8.
//
// Solidity: function SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleSession) SETTLEELREWARDSSTEALINGPENALTYROLE() ([32]byte, error) {
	return _Csmodule.Contract.SETTLEELREWARDSSTEALINGPENALTYROLE(&_Csmodule.CallOpts)
}

// SETTLEELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x3f04f0c8.
//
// Solidity: function SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCallerSession) SETTLEELREWARDSSTEALINGPENALTYROLE() ([32]byte, error) {
	return _Csmodule.Contract.SETTLEELREWARDSSTEALINGPENALTYROLE(&_Csmodule.CallOpts)
}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCaller) STAKINGROUTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "STAKING_ROUTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleSession) STAKINGROUTERROLE() ([32]byte, error) {
	return _Csmodule.Contract.STAKINGROUTERROLE(&_Csmodule.CallOpts)
}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCallerSession) STAKINGROUTERROLE() ([32]byte, error) {
	return _Csmodule.Contract.STAKINGROUTERROLE(&_Csmodule.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Csmodule *CsmoduleCaller) STETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "STETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Csmodule *CsmoduleSession) STETH() (common.Address, error) {
	return _Csmodule.Contract.STETH(&_Csmodule.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_Csmodule *CsmoduleCallerSession) STETH() (common.Address, error) {
	return _Csmodule.Contract.STETH(&_Csmodule.CallOpts)
}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCaller) VERIFIERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "VERIFIER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleSession) VERIFIERROLE() ([32]byte, error) {
	return _Csmodule.Contract.VERIFIERROLE(&_Csmodule.CallOpts)
}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_Csmodule *CsmoduleCallerSession) VERIFIERROLE() ([32]byte, error) {
	return _Csmodule.Contract.VERIFIERROLE(&_Csmodule.CallOpts)
}

// Accounting is a free data retrieval call binding the contract method 0x9624e83e.
//
// Solidity: function accounting() view returns(address)
func (_Csmodule *CsmoduleCaller) Accounting(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "accounting")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accounting is a free data retrieval call binding the contract method 0x9624e83e.
//
// Solidity: function accounting() view returns(address)
func (_Csmodule *CsmoduleSession) Accounting() (common.Address, error) {
	return _Csmodule.Contract.Accounting(&_Csmodule.CallOpts)
}

// Accounting is a free data retrieval call binding the contract method 0x9624e83e.
//
// Solidity: function accounting() view returns(address)
func (_Csmodule *CsmoduleCallerSession) Accounting() (common.Address, error) {
	return _Csmodule.Contract.Accounting(&_Csmodule.CallOpts)
}

// DepositQueueItem is a free data retrieval call binding the contract method 0x37ebdf6f.
//
// Solidity: function depositQueueItem(uint256 queuePriority, uint128 index) view returns(uint256)
func (_Csmodule *CsmoduleCaller) DepositQueueItem(opts *bind.CallOpts, queuePriority *big.Int, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "depositQueueItem", queuePriority, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositQueueItem is a free data retrieval call binding the contract method 0x37ebdf6f.
//
// Solidity: function depositQueueItem(uint256 queuePriority, uint128 index) view returns(uint256)
func (_Csmodule *CsmoduleSession) DepositQueueItem(queuePriority *big.Int, index *big.Int) (*big.Int, error) {
	return _Csmodule.Contract.DepositQueueItem(&_Csmodule.CallOpts, queuePriority, index)
}

// DepositQueueItem is a free data retrieval call binding the contract method 0x37ebdf6f.
//
// Solidity: function depositQueueItem(uint256 queuePriority, uint128 index) view returns(uint256)
func (_Csmodule *CsmoduleCallerSession) DepositQueueItem(queuePriority *big.Int, index *big.Int) (*big.Int, error) {
	return _Csmodule.Contract.DepositQueueItem(&_Csmodule.CallOpts, queuePriority, index)
}

// DepositQueuePointers is a free data retrieval call binding the contract method 0x5810f622.
//
// Solidity: function depositQueuePointers(uint256 queuePriority) view returns(uint128 head, uint128 tail)
func (_Csmodule *CsmoduleCaller) DepositQueuePointers(opts *bind.CallOpts, queuePriority *big.Int) (struct {
	Head *big.Int
	Tail *big.Int
}, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "depositQueuePointers", queuePriority)

	outstruct := new(struct {
		Head *big.Int
		Tail *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Head = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Tail = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// DepositQueuePointers is a free data retrieval call binding the contract method 0x5810f622.
//
// Solidity: function depositQueuePointers(uint256 queuePriority) view returns(uint128 head, uint128 tail)
func (_Csmodule *CsmoduleSession) DepositQueuePointers(queuePriority *big.Int) (struct {
	Head *big.Int
	Tail *big.Int
}, error) {
	return _Csmodule.Contract.DepositQueuePointers(&_Csmodule.CallOpts, queuePriority)
}

// DepositQueuePointers is a free data retrieval call binding the contract method 0x5810f622.
//
// Solidity: function depositQueuePointers(uint256 queuePriority) view returns(uint128 head, uint128 tail)
func (_Csmodule *CsmoduleCallerSession) DepositQueuePointers(queuePriority *big.Int) (struct {
	Head *big.Int
	Tail *big.Int
}, error) {
	return _Csmodule.Contract.DepositQueuePointers(&_Csmodule.CallOpts, queuePriority)
}

// ExitDeadlineThreshold is a free data retrieval call binding the contract method 0x28d6d36b.
//
// Solidity: function exitDeadlineThreshold(uint256 nodeOperatorId) view returns(uint256)
func (_Csmodule *CsmoduleCaller) ExitDeadlineThreshold(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "exitDeadlineThreshold", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExitDeadlineThreshold is a free data retrieval call binding the contract method 0x28d6d36b.
//
// Solidity: function exitDeadlineThreshold(uint256 nodeOperatorId) view returns(uint256)
func (_Csmodule *CsmoduleSession) ExitDeadlineThreshold(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csmodule.Contract.ExitDeadlineThreshold(&_Csmodule.CallOpts, nodeOperatorId)
}

// ExitDeadlineThreshold is a free data retrieval call binding the contract method 0x28d6d36b.
//
// Solidity: function exitDeadlineThreshold(uint256 nodeOperatorId) view returns(uint256)
func (_Csmodule *CsmoduleCallerSession) ExitDeadlineThreshold(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csmodule.Contract.ExitDeadlineThreshold(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_Csmodule *CsmoduleCaller) GetActiveNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getActiveNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_Csmodule *CsmoduleSession) GetActiveNodeOperatorsCount() (*big.Int, error) {
	return _Csmodule.Contract.GetActiveNodeOperatorsCount(&_Csmodule.CallOpts)
}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_Csmodule *CsmoduleCallerSession) GetActiveNodeOperatorsCount() (*big.Int, error) {
	return _Csmodule.Contract.GetActiveNodeOperatorsCount(&_Csmodule.CallOpts)
}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_Csmodule *CsmoduleCaller) GetInitializedVersion(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getInitializedVersion")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_Csmodule *CsmoduleSession) GetInitializedVersion() (uint64, error) {
	return _Csmodule.Contract.GetInitializedVersion(&_Csmodule.CallOpts)
}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_Csmodule *CsmoduleCallerSession) GetInitializedVersion() (uint64, error) {
	return _Csmodule.Contract.GetInitializedVersion(&_Csmodule.CallOpts)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x65c14dc7.
//
// Solidity: function getNodeOperator(uint256 nodeOperatorId) view returns((uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint8,uint32,uint32,address,address,address,address,bool,bool))
func (_Csmodule *CsmoduleCaller) GetNodeOperator(opts *bind.CallOpts, nodeOperatorId *big.Int) (NodeOperator, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getNodeOperator", nodeOperatorId)

	if err != nil {
		return *new(NodeOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeOperator)).(*NodeOperator)

	return out0, err

}

// GetNodeOperator is a free data retrieval call binding the contract method 0x65c14dc7.
//
// Solidity: function getNodeOperator(uint256 nodeOperatorId) view returns((uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint8,uint32,uint32,address,address,address,address,bool,bool))
func (_Csmodule *CsmoduleSession) GetNodeOperator(nodeOperatorId *big.Int) (NodeOperator, error) {
	return _Csmodule.Contract.GetNodeOperator(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x65c14dc7.
//
// Solidity: function getNodeOperator(uint256 nodeOperatorId) view returns((uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint8,uint32,uint32,address,address,address,address,bool,bool))
func (_Csmodule *CsmoduleCallerSession) GetNodeOperator(nodeOperatorId *big.Int) (NodeOperator, error) {
	return _Csmodule.Contract.GetNodeOperator(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 offset, uint256 limit) view returns(uint256[] nodeOperatorIds)
func (_Csmodule *CsmoduleCaller) GetNodeOperatorIds(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getNodeOperatorIds", offset, limit)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 offset, uint256 limit) view returns(uint256[] nodeOperatorIds)
func (_Csmodule *CsmoduleSession) GetNodeOperatorIds(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _Csmodule.Contract.GetNodeOperatorIds(&_Csmodule.CallOpts, offset, limit)
}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 offset, uint256 limit) view returns(uint256[] nodeOperatorIds)
func (_Csmodule *CsmoduleCallerSession) GetNodeOperatorIds(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _Csmodule.Contract.GetNodeOperatorIds(&_Csmodule.CallOpts, offset, limit)
}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 nodeOperatorId) view returns(bool)
func (_Csmodule *CsmoduleCaller) GetNodeOperatorIsActive(opts *bind.CallOpts, nodeOperatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getNodeOperatorIsActive", nodeOperatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 nodeOperatorId) view returns(bool)
func (_Csmodule *CsmoduleSession) GetNodeOperatorIsActive(nodeOperatorId *big.Int) (bool, error) {
	return _Csmodule.Contract.GetNodeOperatorIsActive(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 nodeOperatorId) view returns(bool)
func (_Csmodule *CsmoduleCallerSession) GetNodeOperatorIsActive(nodeOperatorId *big.Int) (bool, error) {
	return _Csmodule.Contract.GetNodeOperatorIsActive(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorManagementProperties is a free data retrieval call binding the contract method 0x499b8e9a.
//
// Solidity: function getNodeOperatorManagementProperties(uint256 nodeOperatorId) view returns((address,address,bool))
func (_Csmodule *CsmoduleCaller) GetNodeOperatorManagementProperties(opts *bind.CallOpts, nodeOperatorId *big.Int) (NodeOperatorManagementProperties, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getNodeOperatorManagementProperties", nodeOperatorId)

	if err != nil {
		return *new(NodeOperatorManagementProperties), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeOperatorManagementProperties)).(*NodeOperatorManagementProperties)

	return out0, err

}

// GetNodeOperatorManagementProperties is a free data retrieval call binding the contract method 0x499b8e9a.
//
// Solidity: function getNodeOperatorManagementProperties(uint256 nodeOperatorId) view returns((address,address,bool))
func (_Csmodule *CsmoduleSession) GetNodeOperatorManagementProperties(nodeOperatorId *big.Int) (NodeOperatorManagementProperties, error) {
	return _Csmodule.Contract.GetNodeOperatorManagementProperties(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorManagementProperties is a free data retrieval call binding the contract method 0x499b8e9a.
//
// Solidity: function getNodeOperatorManagementProperties(uint256 nodeOperatorId) view returns((address,address,bool))
func (_Csmodule *CsmoduleCallerSession) GetNodeOperatorManagementProperties(nodeOperatorId *big.Int) (NodeOperatorManagementProperties, error) {
	return _Csmodule.Contract.GetNodeOperatorManagementProperties(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorNonWithdrawnKeys is a free data retrieval call binding the contract method 0x8ec69028.
//
// Solidity: function getNodeOperatorNonWithdrawnKeys(uint256 nodeOperatorId) view returns(uint256)
func (_Csmodule *CsmoduleCaller) GetNodeOperatorNonWithdrawnKeys(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getNodeOperatorNonWithdrawnKeys", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorNonWithdrawnKeys is a free data retrieval call binding the contract method 0x8ec69028.
//
// Solidity: function getNodeOperatorNonWithdrawnKeys(uint256 nodeOperatorId) view returns(uint256)
func (_Csmodule *CsmoduleSession) GetNodeOperatorNonWithdrawnKeys(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csmodule.Contract.GetNodeOperatorNonWithdrawnKeys(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorNonWithdrawnKeys is a free data retrieval call binding the contract method 0x8ec69028.
//
// Solidity: function getNodeOperatorNonWithdrawnKeys(uint256 nodeOperatorId) view returns(uint256)
func (_Csmodule *CsmoduleCallerSession) GetNodeOperatorNonWithdrawnKeys(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csmodule.Contract.GetNodeOperatorNonWithdrawnKeys(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorOwner is a free data retrieval call binding the contract method 0xb055e15c.
//
// Solidity: function getNodeOperatorOwner(uint256 nodeOperatorId) view returns(address)
func (_Csmodule *CsmoduleCaller) GetNodeOperatorOwner(opts *bind.CallOpts, nodeOperatorId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getNodeOperatorOwner", nodeOperatorId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetNodeOperatorOwner is a free data retrieval call binding the contract method 0xb055e15c.
//
// Solidity: function getNodeOperatorOwner(uint256 nodeOperatorId) view returns(address)
func (_Csmodule *CsmoduleSession) GetNodeOperatorOwner(nodeOperatorId *big.Int) (common.Address, error) {
	return _Csmodule.Contract.GetNodeOperatorOwner(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorOwner is a free data retrieval call binding the contract method 0xb055e15c.
//
// Solidity: function getNodeOperatorOwner(uint256 nodeOperatorId) view returns(address)
func (_Csmodule *CsmoduleCallerSession) GetNodeOperatorOwner(nodeOperatorId *big.Int) (common.Address, error) {
	return _Csmodule.Contract.GetNodeOperatorOwner(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xb3076c3c.
//
// Solidity: function getNodeOperatorSummary(uint256 nodeOperatorId) view returns(uint256 targetLimitMode, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_Csmodule *CsmoduleCaller) GetNodeOperatorSummary(opts *bind.CallOpts, nodeOperatorId *big.Int) (struct {
	TargetLimitMode            *big.Int
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getNodeOperatorSummary", nodeOperatorId)

	outstruct := new(struct {
		TargetLimitMode            *big.Int
		TargetValidatorsCount      *big.Int
		StuckValidatorsCount       *big.Int
		RefundedValidatorsCount    *big.Int
		StuckPenaltyEndTimestamp   *big.Int
		TotalExitedValidators      *big.Int
		TotalDepositedValidators   *big.Int
		DepositableValidatorsCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TargetLimitMode = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TargetValidatorsCount = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.StuckValidatorsCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RefundedValidatorsCount = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.StuckPenaltyEndTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.TotalExitedValidators = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.TotalDepositedValidators = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.DepositableValidatorsCount = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xb3076c3c.
//
// Solidity: function getNodeOperatorSummary(uint256 nodeOperatorId) view returns(uint256 targetLimitMode, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_Csmodule *CsmoduleSession) GetNodeOperatorSummary(nodeOperatorId *big.Int) (struct {
	TargetLimitMode            *big.Int
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _Csmodule.Contract.GetNodeOperatorSummary(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xb3076c3c.
//
// Solidity: function getNodeOperatorSummary(uint256 nodeOperatorId) view returns(uint256 targetLimitMode, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_Csmodule *CsmoduleCallerSession) GetNodeOperatorSummary(nodeOperatorId *big.Int) (struct {
	TargetLimitMode            *big.Int
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _Csmodule.Contract.GetNodeOperatorSummary(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorTotalDepositedKeys is a free data retrieval call binding the contract method 0xa0c8c47e.
//
// Solidity: function getNodeOperatorTotalDepositedKeys(uint256 nodeOperatorId) view returns(uint256 totalDepositedKeys)
func (_Csmodule *CsmoduleCaller) GetNodeOperatorTotalDepositedKeys(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getNodeOperatorTotalDepositedKeys", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorTotalDepositedKeys is a free data retrieval call binding the contract method 0xa0c8c47e.
//
// Solidity: function getNodeOperatorTotalDepositedKeys(uint256 nodeOperatorId) view returns(uint256 totalDepositedKeys)
func (_Csmodule *CsmoduleSession) GetNodeOperatorTotalDepositedKeys(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csmodule.Contract.GetNodeOperatorTotalDepositedKeys(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorTotalDepositedKeys is a free data retrieval call binding the contract method 0xa0c8c47e.
//
// Solidity: function getNodeOperatorTotalDepositedKeys(uint256 nodeOperatorId) view returns(uint256 totalDepositedKeys)
func (_Csmodule *CsmoduleCallerSession) GetNodeOperatorTotalDepositedKeys(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csmodule.Contract.GetNodeOperatorTotalDepositedKeys(&_Csmodule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_Csmodule *CsmoduleCaller) GetNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_Csmodule *CsmoduleSession) GetNodeOperatorsCount() (*big.Int, error) {
	return _Csmodule.Contract.GetNodeOperatorsCount(&_Csmodule.CallOpts)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_Csmodule *CsmoduleCallerSession) GetNodeOperatorsCount() (*big.Int, error) {
	return _Csmodule.Contract.GetNodeOperatorsCount(&_Csmodule.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Csmodule *CsmoduleCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Csmodule *CsmoduleSession) GetNonce() (*big.Int, error) {
	return _Csmodule.Contract.GetNonce(&_Csmodule.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_Csmodule *CsmoduleCallerSession) GetNonce() (*big.Int, error) {
	return _Csmodule.Contract.GetNonce(&_Csmodule.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Csmodule *CsmoduleCaller) GetResumeSinceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getResumeSinceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Csmodule *CsmoduleSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Csmodule.Contract.GetResumeSinceTimestamp(&_Csmodule.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Csmodule *CsmoduleCallerSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Csmodule.Contract.GetResumeSinceTimestamp(&_Csmodule.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csmodule *CsmoduleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csmodule *CsmoduleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Csmodule.Contract.GetRoleAdmin(&_Csmodule.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csmodule *CsmoduleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Csmodule.Contract.GetRoleAdmin(&_Csmodule.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csmodule *CsmoduleCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csmodule *CsmoduleSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Csmodule.Contract.GetRoleMember(&_Csmodule.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csmodule *CsmoduleCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Csmodule.Contract.GetRoleMember(&_Csmodule.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csmodule *CsmoduleCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csmodule *CsmoduleSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Csmodule.Contract.GetRoleMemberCount(&_Csmodule.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csmodule *CsmoduleCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Csmodule.Contract.GetRoleMemberCount(&_Csmodule.CallOpts, role)
}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes)
func (_Csmodule *CsmoduleCaller) GetSigningKeys(opts *bind.CallOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]byte, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getSigningKeys", nodeOperatorId, startIndex, keysCount)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes)
func (_Csmodule *CsmoduleSession) GetSigningKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]byte, error) {
	return _Csmodule.Contract.GetSigningKeys(&_Csmodule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes)
func (_Csmodule *CsmoduleCallerSession) GetSigningKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]byte, error) {
	return _Csmodule.Contract.GetSigningKeys(&_Csmodule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetSigningKeysWithSignatures is a free data retrieval call binding the contract method 0x50388cb6.
//
// Solidity: function getSigningKeysWithSignatures(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys, bytes signatures)
func (_Csmodule *CsmoduleCaller) GetSigningKeysWithSignatures(opts *bind.CallOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (struct {
	Keys       []byte
	Signatures []byte
}, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getSigningKeysWithSignatures", nodeOperatorId, startIndex, keysCount)

	outstruct := new(struct {
		Keys       []byte
		Signatures []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Keys = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.Signatures = *abi.ConvertType(out[1], new([]byte)).(*[]byte)

	return *outstruct, err

}

// GetSigningKeysWithSignatures is a free data retrieval call binding the contract method 0x50388cb6.
//
// Solidity: function getSigningKeysWithSignatures(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys, bytes signatures)
func (_Csmodule *CsmoduleSession) GetSigningKeysWithSignatures(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (struct {
	Keys       []byte
	Signatures []byte
}, error) {
	return _Csmodule.Contract.GetSigningKeysWithSignatures(&_Csmodule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetSigningKeysWithSignatures is a free data retrieval call binding the contract method 0x50388cb6.
//
// Solidity: function getSigningKeysWithSignatures(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys, bytes signatures)
func (_Csmodule *CsmoduleCallerSession) GetSigningKeysWithSignatures(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (struct {
	Keys       []byte
	Signatures []byte
}, error) {
	return _Csmodule.Contract.GetSigningKeysWithSignatures(&_Csmodule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x9abddf09.
//
// Solidity: function getStakingModuleSummary() view returns(uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_Csmodule *CsmoduleCaller) GetStakingModuleSummary(opts *bind.CallOpts) (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getStakingModuleSummary")

	outstruct := new(struct {
		TotalExitedValidators      *big.Int
		TotalDepositedValidators   *big.Int
		DepositableValidatorsCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalExitedValidators = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalDepositedValidators = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DepositableValidatorsCount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x9abddf09.
//
// Solidity: function getStakingModuleSummary() view returns(uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_Csmodule *CsmoduleSession) GetStakingModuleSummary() (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _Csmodule.Contract.GetStakingModuleSummary(&_Csmodule.CallOpts)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x9abddf09.
//
// Solidity: function getStakingModuleSummary() view returns(uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_Csmodule *CsmoduleCallerSession) GetStakingModuleSummary() (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _Csmodule.Contract.GetStakingModuleSummary(&_Csmodule.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_Csmodule *CsmoduleCaller) GetType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "getType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_Csmodule *CsmoduleSession) GetType() ([32]byte, error) {
	return _Csmodule.Contract.GetType(&_Csmodule.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_Csmodule *CsmoduleCallerSession) GetType() ([32]byte, error) {
	return _Csmodule.Contract.GetType(&_Csmodule.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csmodule *CsmoduleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csmodule *CsmoduleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Csmodule.Contract.HasRole(&_Csmodule.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csmodule *CsmoduleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Csmodule.Contract.HasRole(&_Csmodule.CallOpts, role, account)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Csmodule *CsmoduleCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Csmodule *CsmoduleSession) IsPaused() (bool, error) {
	return _Csmodule.Contract.IsPaused(&_Csmodule.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Csmodule *CsmoduleCallerSession) IsPaused() (bool, error) {
	return _Csmodule.Contract.IsPaused(&_Csmodule.CallOpts)
}

// IsValidatorExitDelayPenaltyApplicable is a free data retrieval call binding the contract method 0x83b57a4e.
//
// Solidity: function isValidatorExitDelayPenaltyApplicable(uint256 nodeOperatorId, uint256 , bytes publicKey, uint256 eligibleToExitInSec) view returns(bool)
func (_Csmodule *CsmoduleCaller) IsValidatorExitDelayPenaltyApplicable(opts *bind.CallOpts, nodeOperatorId *big.Int, arg1 *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (bool, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "isValidatorExitDelayPenaltyApplicable", nodeOperatorId, arg1, publicKey, eligibleToExitInSec)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorExitDelayPenaltyApplicable is a free data retrieval call binding the contract method 0x83b57a4e.
//
// Solidity: function isValidatorExitDelayPenaltyApplicable(uint256 nodeOperatorId, uint256 , bytes publicKey, uint256 eligibleToExitInSec) view returns(bool)
func (_Csmodule *CsmoduleSession) IsValidatorExitDelayPenaltyApplicable(nodeOperatorId *big.Int, arg1 *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (bool, error) {
	return _Csmodule.Contract.IsValidatorExitDelayPenaltyApplicable(&_Csmodule.CallOpts, nodeOperatorId, arg1, publicKey, eligibleToExitInSec)
}

// IsValidatorExitDelayPenaltyApplicable is a free data retrieval call binding the contract method 0x83b57a4e.
//
// Solidity: function isValidatorExitDelayPenaltyApplicable(uint256 nodeOperatorId, uint256 , bytes publicKey, uint256 eligibleToExitInSec) view returns(bool)
func (_Csmodule *CsmoduleCallerSession) IsValidatorExitDelayPenaltyApplicable(nodeOperatorId *big.Int, arg1 *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (bool, error) {
	return _Csmodule.Contract.IsValidatorExitDelayPenaltyApplicable(&_Csmodule.CallOpts, nodeOperatorId, arg1, publicKey, eligibleToExitInSec)
}

// IsValidatorWithdrawn is a free data retrieval call binding the contract method 0x53433643.
//
// Solidity: function isValidatorWithdrawn(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_Csmodule *CsmoduleCaller) IsValidatorWithdrawn(opts *bind.CallOpts, nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "isValidatorWithdrawn", nodeOperatorId, keyIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorWithdrawn is a free data retrieval call binding the contract method 0x53433643.
//
// Solidity: function isValidatorWithdrawn(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_Csmodule *CsmoduleSession) IsValidatorWithdrawn(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _Csmodule.Contract.IsValidatorWithdrawn(&_Csmodule.CallOpts, nodeOperatorId, keyIndex)
}

// IsValidatorWithdrawn is a free data retrieval call binding the contract method 0x53433643.
//
// Solidity: function isValidatorWithdrawn(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_Csmodule *CsmoduleCallerSession) IsValidatorWithdrawn(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _Csmodule.Contract.IsValidatorWithdrawn(&_Csmodule.CallOpts, nodeOperatorId, keyIndex)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csmodule *CsmoduleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Csmodule.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csmodule *CsmoduleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Csmodule.Contract.SupportsInterface(&_Csmodule.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csmodule *CsmoduleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Csmodule.Contract.SupportsInterface(&_Csmodule.CallOpts, interfaceId)
}

// AddValidatorKeysETH is a paid mutator transaction binding the contract method 0xa1913f4b.
//
// Solidity: function addValidatorKeysETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures) payable returns()
func (_Csmodule *CsmoduleTransactor) AddValidatorKeysETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "addValidatorKeysETH", from, nodeOperatorId, keysCount, publicKeys, signatures)
}

// AddValidatorKeysETH is a paid mutator transaction binding the contract method 0xa1913f4b.
//
// Solidity: function addValidatorKeysETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures) payable returns()
func (_Csmodule *CsmoduleSession) AddValidatorKeysETH(from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte) (*types.Transaction, error) {
	return _Csmodule.Contract.AddValidatorKeysETH(&_Csmodule.TransactOpts, from, nodeOperatorId, keysCount, publicKeys, signatures)
}

// AddValidatorKeysETH is a paid mutator transaction binding the contract method 0xa1913f4b.
//
// Solidity: function addValidatorKeysETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures) payable returns()
func (_Csmodule *CsmoduleTransactorSession) AddValidatorKeysETH(from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte) (*types.Transaction, error) {
	return _Csmodule.Contract.AddValidatorKeysETH(&_Csmodule.TransactOpts, from, nodeOperatorId, keysCount, publicKeys, signatures)
}

// AddValidatorKeysStETH is a paid mutator transaction binding the contract method 0xf696ccb3.
//
// Solidity: function addValidatorKeysStETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csmodule *CsmoduleTransactor) AddValidatorKeysStETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "addValidatorKeysStETH", from, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysStETH is a paid mutator transaction binding the contract method 0xf696ccb3.
//
// Solidity: function addValidatorKeysStETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csmodule *CsmoduleSession) AddValidatorKeysStETH(from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csmodule.Contract.AddValidatorKeysStETH(&_Csmodule.TransactOpts, from, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysStETH is a paid mutator transaction binding the contract method 0xf696ccb3.
//
// Solidity: function addValidatorKeysStETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csmodule *CsmoduleTransactorSession) AddValidatorKeysStETH(from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csmodule.Contract.AddValidatorKeysStETH(&_Csmodule.TransactOpts, from, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysWstETH is a paid mutator transaction binding the contract method 0xa6ab5b9c.
//
// Solidity: function addValidatorKeysWstETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csmodule *CsmoduleTransactor) AddValidatorKeysWstETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "addValidatorKeysWstETH", from, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysWstETH is a paid mutator transaction binding the contract method 0xa6ab5b9c.
//
// Solidity: function addValidatorKeysWstETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csmodule *CsmoduleSession) AddValidatorKeysWstETH(from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csmodule.Contract.AddValidatorKeysWstETH(&_Csmodule.TransactOpts, from, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysWstETH is a paid mutator transaction binding the contract method 0xa6ab5b9c.
//
// Solidity: function addValidatorKeysWstETH(address from, uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csmodule *CsmoduleTransactorSession) AddValidatorKeysWstETH(from common.Address, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csmodule.Contract.AddValidatorKeysWstETH(&_Csmodule.TransactOpts, from, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// CancelELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x40044801.
//
// Solidity: function cancelELRewardsStealingPenalty(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csmodule *CsmoduleTransactor) CancelELRewardsStealingPenalty(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "cancelELRewardsStealingPenalty", nodeOperatorId, amount)
}

// CancelELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x40044801.
//
// Solidity: function cancelELRewardsStealingPenalty(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csmodule *CsmoduleSession) CancelELRewardsStealingPenalty(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.CancelELRewardsStealingPenalty(&_Csmodule.TransactOpts, nodeOperatorId, amount)
}

// CancelELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x40044801.
//
// Solidity: function cancelELRewardsStealingPenalty(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csmodule *CsmoduleTransactorSession) CancelELRewardsStealingPenalty(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.CancelELRewardsStealingPenalty(&_Csmodule.TransactOpts, nodeOperatorId, amount)
}

// ChangeNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x75a401da.
//
// Solidity: function changeNodeOperatorRewardAddress(uint256 nodeOperatorId, address newAddress) returns()
func (_Csmodule *CsmoduleTransactor) ChangeNodeOperatorRewardAddress(opts *bind.TransactOpts, nodeOperatorId *big.Int, newAddress common.Address) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "changeNodeOperatorRewardAddress", nodeOperatorId, newAddress)
}

// ChangeNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x75a401da.
//
// Solidity: function changeNodeOperatorRewardAddress(uint256 nodeOperatorId, address newAddress) returns()
func (_Csmodule *CsmoduleSession) ChangeNodeOperatorRewardAddress(nodeOperatorId *big.Int, newAddress common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.ChangeNodeOperatorRewardAddress(&_Csmodule.TransactOpts, nodeOperatorId, newAddress)
}

// ChangeNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x75a401da.
//
// Solidity: function changeNodeOperatorRewardAddress(uint256 nodeOperatorId, address newAddress) returns()
func (_Csmodule *CsmoduleTransactorSession) ChangeNodeOperatorRewardAddress(nodeOperatorId *big.Int, newAddress common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.ChangeNodeOperatorRewardAddress(&_Csmodule.TransactOpts, nodeOperatorId, newAddress)
}

// CleanDepositQueue is a paid mutator transaction binding the contract method 0x735dfa28.
//
// Solidity: function cleanDepositQueue(uint256 maxItems) returns(uint256 removed, uint256 lastRemovedAtDepth)
func (_Csmodule *CsmoduleTransactor) CleanDepositQueue(opts *bind.TransactOpts, maxItems *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "cleanDepositQueue", maxItems)
}

// CleanDepositQueue is a paid mutator transaction binding the contract method 0x735dfa28.
//
// Solidity: function cleanDepositQueue(uint256 maxItems) returns(uint256 removed, uint256 lastRemovedAtDepth)
func (_Csmodule *CsmoduleSession) CleanDepositQueue(maxItems *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.CleanDepositQueue(&_Csmodule.TransactOpts, maxItems)
}

// CleanDepositQueue is a paid mutator transaction binding the contract method 0x735dfa28.
//
// Solidity: function cleanDepositQueue(uint256 maxItems) returns(uint256 removed, uint256 lastRemovedAtDepth)
func (_Csmodule *CsmoduleTransactorSession) CleanDepositQueue(maxItems *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.CleanDepositQueue(&_Csmodule.TransactOpts, maxItems)
}

// CompensateELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x6efe37a2.
//
// Solidity: function compensateELRewardsStealingPenalty(uint256 nodeOperatorId) payable returns()
func (_Csmodule *CsmoduleTransactor) CompensateELRewardsStealingPenalty(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "compensateELRewardsStealingPenalty", nodeOperatorId)
}

// CompensateELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x6efe37a2.
//
// Solidity: function compensateELRewardsStealingPenalty(uint256 nodeOperatorId) payable returns()
func (_Csmodule *CsmoduleSession) CompensateELRewardsStealingPenalty(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.CompensateELRewardsStealingPenalty(&_Csmodule.TransactOpts, nodeOperatorId)
}

// CompensateELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x6efe37a2.
//
// Solidity: function compensateELRewardsStealingPenalty(uint256 nodeOperatorId) payable returns()
func (_Csmodule *CsmoduleTransactorSession) CompensateELRewardsStealingPenalty(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.CompensateELRewardsStealingPenalty(&_Csmodule.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x6bb1bfdf.
//
// Solidity: function confirmNodeOperatorManagerAddressChange(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleTransactor) ConfirmNodeOperatorManagerAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "confirmNodeOperatorManagerAddressChange", nodeOperatorId)
}

// ConfirmNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x6bb1bfdf.
//
// Solidity: function confirmNodeOperatorManagerAddressChange(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleSession) ConfirmNodeOperatorManagerAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.ConfirmNodeOperatorManagerAddressChange(&_Csmodule.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x6bb1bfdf.
//
// Solidity: function confirmNodeOperatorManagerAddressChange(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleTransactorSession) ConfirmNodeOperatorManagerAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.ConfirmNodeOperatorManagerAddressChange(&_Csmodule.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x5204281c.
//
// Solidity: function confirmNodeOperatorRewardAddressChange(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleTransactor) ConfirmNodeOperatorRewardAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "confirmNodeOperatorRewardAddressChange", nodeOperatorId)
}

// ConfirmNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x5204281c.
//
// Solidity: function confirmNodeOperatorRewardAddressChange(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleSession) ConfirmNodeOperatorRewardAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.ConfirmNodeOperatorRewardAddressChange(&_Csmodule.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x5204281c.
//
// Solidity: function confirmNodeOperatorRewardAddressChange(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleTransactorSession) ConfirmNodeOperatorRewardAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.ConfirmNodeOperatorRewardAddressChange(&_Csmodule.TransactOpts, nodeOperatorId)
}

// CreateNodeOperator is a paid mutator transaction binding the contract method 0xa4516c98.
//
// Solidity: function createNodeOperator(address from, (address,address,bool) managementProperties, address referrer) returns(uint256 nodeOperatorId)
func (_Csmodule *CsmoduleTransactor) CreateNodeOperator(opts *bind.TransactOpts, from common.Address, managementProperties NodeOperatorManagementProperties, referrer common.Address) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "createNodeOperator", from, managementProperties, referrer)
}

// CreateNodeOperator is a paid mutator transaction binding the contract method 0xa4516c98.
//
// Solidity: function createNodeOperator(address from, (address,address,bool) managementProperties, address referrer) returns(uint256 nodeOperatorId)
func (_Csmodule *CsmoduleSession) CreateNodeOperator(from common.Address, managementProperties NodeOperatorManagementProperties, referrer common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.CreateNodeOperator(&_Csmodule.TransactOpts, from, managementProperties, referrer)
}

// CreateNodeOperator is a paid mutator transaction binding the contract method 0xa4516c98.
//
// Solidity: function createNodeOperator(address from, (address,address,bool) managementProperties, address referrer) returns(uint256 nodeOperatorId)
func (_Csmodule *CsmoduleTransactorSession) CreateNodeOperator(from common.Address, managementProperties NodeOperatorManagementProperties, referrer common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.CreateNodeOperator(&_Csmodule.TransactOpts, from, managementProperties, referrer)
}

// DecreaseVettedSigningKeysCount is a paid mutator transaction binding the contract method 0xb643189b.
//
// Solidity: function decreaseVettedSigningKeysCount(bytes nodeOperatorIds, bytes vettedSigningKeysCounts) returns()
func (_Csmodule *CsmoduleTransactor) DecreaseVettedSigningKeysCount(opts *bind.TransactOpts, nodeOperatorIds []byte, vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "decreaseVettedSigningKeysCount", nodeOperatorIds, vettedSigningKeysCounts)
}

// DecreaseVettedSigningKeysCount is a paid mutator transaction binding the contract method 0xb643189b.
//
// Solidity: function decreaseVettedSigningKeysCount(bytes nodeOperatorIds, bytes vettedSigningKeysCounts) returns()
func (_Csmodule *CsmoduleSession) DecreaseVettedSigningKeysCount(nodeOperatorIds []byte, vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _Csmodule.Contract.DecreaseVettedSigningKeysCount(&_Csmodule.TransactOpts, nodeOperatorIds, vettedSigningKeysCounts)
}

// DecreaseVettedSigningKeysCount is a paid mutator transaction binding the contract method 0xb643189b.
//
// Solidity: function decreaseVettedSigningKeysCount(bytes nodeOperatorIds, bytes vettedSigningKeysCounts) returns()
func (_Csmodule *CsmoduleTransactorSession) DecreaseVettedSigningKeysCount(nodeOperatorIds []byte, vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _Csmodule.Contract.DecreaseVettedSigningKeysCount(&_Csmodule.TransactOpts, nodeOperatorIds, vettedSigningKeysCounts)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x88984a97.
//
// Solidity: function finalizeUpgradeV2() returns()
func (_Csmodule *CsmoduleTransactor) FinalizeUpgradeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "finalizeUpgradeV2")
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x88984a97.
//
// Solidity: function finalizeUpgradeV2() returns()
func (_Csmodule *CsmoduleSession) FinalizeUpgradeV2() (*types.Transaction, error) {
	return _Csmodule.Contract.FinalizeUpgradeV2(&_Csmodule.TransactOpts)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x88984a97.
//
// Solidity: function finalizeUpgradeV2() returns()
func (_Csmodule *CsmoduleTransactorSession) FinalizeUpgradeV2() (*types.Transaction, error) {
	return _Csmodule.Contract.FinalizeUpgradeV2(&_Csmodule.TransactOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csmodule *CsmoduleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csmodule *CsmoduleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.GrantRole(&_Csmodule.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csmodule *CsmoduleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.GrantRole(&_Csmodule.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_Csmodule *CsmoduleTransactor) Initialize(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "initialize", admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_Csmodule *CsmoduleSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.Initialize(&_Csmodule.TransactOpts, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address admin) returns()
func (_Csmodule *CsmoduleTransactorSession) Initialize(admin common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.Initialize(&_Csmodule.TransactOpts, admin)
}

// MigrateToPriorityQueue is a paid mutator transaction binding the contract method 0x9417366f.
//
// Solidity: function migrateToPriorityQueue(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleTransactor) MigrateToPriorityQueue(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "migrateToPriorityQueue", nodeOperatorId)
}

// MigrateToPriorityQueue is a paid mutator transaction binding the contract method 0x9417366f.
//
// Solidity: function migrateToPriorityQueue(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleSession) MigrateToPriorityQueue(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.MigrateToPriorityQueue(&_Csmodule.TransactOpts, nodeOperatorId)
}

// MigrateToPriorityQueue is a paid mutator transaction binding the contract method 0x9417366f.
//
// Solidity: function migrateToPriorityQueue(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleTransactorSession) MigrateToPriorityQueue(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.MigrateToPriorityQueue(&_Csmodule.TransactOpts, nodeOperatorId)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_Csmodule *CsmoduleTransactor) ObtainDepositData(opts *bind.TransactOpts, depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "obtainDepositData", depositsCount, arg1)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_Csmodule *CsmoduleSession) ObtainDepositData(depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _Csmodule.Contract.ObtainDepositData(&_Csmodule.TransactOpts, depositsCount, arg1)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_Csmodule *CsmoduleTransactorSession) ObtainDepositData(depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _Csmodule.Contract.ObtainDepositData(&_Csmodule.TransactOpts, depositsCount, arg1)
}

// OnExitedAndStuckValidatorsCountsUpdated is a paid mutator transaction binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() returns()
func (_Csmodule *CsmoduleTransactor) OnExitedAndStuckValidatorsCountsUpdated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "onExitedAndStuckValidatorsCountsUpdated")
}

// OnExitedAndStuckValidatorsCountsUpdated is a paid mutator transaction binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() returns()
func (_Csmodule *CsmoduleSession) OnExitedAndStuckValidatorsCountsUpdated() (*types.Transaction, error) {
	return _Csmodule.Contract.OnExitedAndStuckValidatorsCountsUpdated(&_Csmodule.TransactOpts)
}

// OnExitedAndStuckValidatorsCountsUpdated is a paid mutator transaction binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() returns()
func (_Csmodule *CsmoduleTransactorSession) OnExitedAndStuckValidatorsCountsUpdated() (*types.Transaction, error) {
	return _Csmodule.Contract.OnExitedAndStuckValidatorsCountsUpdated(&_Csmodule.TransactOpts)
}

// OnRewardsMinted is a paid mutator transaction binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 totalShares) returns()
func (_Csmodule *CsmoduleTransactor) OnRewardsMinted(opts *bind.TransactOpts, totalShares *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "onRewardsMinted", totalShares)
}

// OnRewardsMinted is a paid mutator transaction binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 totalShares) returns()
func (_Csmodule *CsmoduleSession) OnRewardsMinted(totalShares *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.OnRewardsMinted(&_Csmodule.TransactOpts, totalShares)
}

// OnRewardsMinted is a paid mutator transaction binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 totalShares) returns()
func (_Csmodule *CsmoduleTransactorSession) OnRewardsMinted(totalShares *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.OnRewardsMinted(&_Csmodule.TransactOpts, totalShares)
}

// OnValidatorExitTriggered is a paid mutator transaction binding the contract method 0x693cc600.
//
// Solidity: function onValidatorExitTriggered(uint256 nodeOperatorId, bytes publicKey, uint256 withdrawalRequestPaidFee, uint256 exitType) returns()
func (_Csmodule *CsmoduleTransactor) OnValidatorExitTriggered(opts *bind.TransactOpts, nodeOperatorId *big.Int, publicKey []byte, withdrawalRequestPaidFee *big.Int, exitType *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "onValidatorExitTriggered", nodeOperatorId, publicKey, withdrawalRequestPaidFee, exitType)
}

// OnValidatorExitTriggered is a paid mutator transaction binding the contract method 0x693cc600.
//
// Solidity: function onValidatorExitTriggered(uint256 nodeOperatorId, bytes publicKey, uint256 withdrawalRequestPaidFee, uint256 exitType) returns()
func (_Csmodule *CsmoduleSession) OnValidatorExitTriggered(nodeOperatorId *big.Int, publicKey []byte, withdrawalRequestPaidFee *big.Int, exitType *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.OnValidatorExitTriggered(&_Csmodule.TransactOpts, nodeOperatorId, publicKey, withdrawalRequestPaidFee, exitType)
}

// OnValidatorExitTriggered is a paid mutator transaction binding the contract method 0x693cc600.
//
// Solidity: function onValidatorExitTriggered(uint256 nodeOperatorId, bytes publicKey, uint256 withdrawalRequestPaidFee, uint256 exitType) returns()
func (_Csmodule *CsmoduleTransactorSession) OnValidatorExitTriggered(nodeOperatorId *big.Int, publicKey []byte, withdrawalRequestPaidFee *big.Int, exitType *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.OnValidatorExitTriggered(&_Csmodule.TransactOpts, nodeOperatorId, publicKey, withdrawalRequestPaidFee, exitType)
}

// OnWithdrawalCredentialsChanged is a paid mutator transaction binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() returns()
func (_Csmodule *CsmoduleTransactor) OnWithdrawalCredentialsChanged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "onWithdrawalCredentialsChanged")
}

// OnWithdrawalCredentialsChanged is a paid mutator transaction binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() returns()
func (_Csmodule *CsmoduleSession) OnWithdrawalCredentialsChanged() (*types.Transaction, error) {
	return _Csmodule.Contract.OnWithdrawalCredentialsChanged(&_Csmodule.TransactOpts)
}

// OnWithdrawalCredentialsChanged is a paid mutator transaction binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() returns()
func (_Csmodule *CsmoduleTransactorSession) OnWithdrawalCredentialsChanged() (*types.Transaction, error) {
	return _Csmodule.Contract.OnWithdrawalCredentialsChanged(&_Csmodule.TransactOpts)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Csmodule *CsmoduleTransactor) PauseFor(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "pauseFor", duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Csmodule *CsmoduleSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.PauseFor(&_Csmodule.TransactOpts, duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Csmodule *CsmoduleTransactorSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.PauseFor(&_Csmodule.TransactOpts, duration)
}

// ProposeNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x8cabe959.
//
// Solidity: function proposeNodeOperatorManagerAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_Csmodule *CsmoduleTransactor) ProposeNodeOperatorManagerAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "proposeNodeOperatorManagerAddressChange", nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x8cabe959.
//
// Solidity: function proposeNodeOperatorManagerAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_Csmodule *CsmoduleSession) ProposeNodeOperatorManagerAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.ProposeNodeOperatorManagerAddressChange(&_Csmodule.TransactOpts, nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x8cabe959.
//
// Solidity: function proposeNodeOperatorManagerAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_Csmodule *CsmoduleTransactorSession) ProposeNodeOperatorManagerAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.ProposeNodeOperatorManagerAddressChange(&_Csmodule.TransactOpts, nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x1b40b231.
//
// Solidity: function proposeNodeOperatorRewardAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_Csmodule *CsmoduleTransactor) ProposeNodeOperatorRewardAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "proposeNodeOperatorRewardAddressChange", nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x1b40b231.
//
// Solidity: function proposeNodeOperatorRewardAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_Csmodule *CsmoduleSession) ProposeNodeOperatorRewardAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.ProposeNodeOperatorRewardAddressChange(&_Csmodule.TransactOpts, nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x1b40b231.
//
// Solidity: function proposeNodeOperatorRewardAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_Csmodule *CsmoduleTransactorSession) ProposeNodeOperatorRewardAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.ProposeNodeOperatorRewardAddressChange(&_Csmodule.TransactOpts, nodeOperatorId, proposedAddress)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csmodule *CsmoduleTransactor) RecoverERC1155(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "recoverERC1155", token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csmodule *CsmoduleSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.RecoverERC1155(&_Csmodule.TransactOpts, token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csmodule *CsmoduleTransactorSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.RecoverERC1155(&_Csmodule.TransactOpts, token, tokenId)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csmodule *CsmoduleTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "recoverERC20", token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csmodule *CsmoduleSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.RecoverERC20(&_Csmodule.TransactOpts, token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csmodule *CsmoduleTransactorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.RecoverERC20(&_Csmodule.TransactOpts, token, amount)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csmodule *CsmoduleTransactor) RecoverERC721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "recoverERC721", token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csmodule *CsmoduleSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.RecoverERC721(&_Csmodule.TransactOpts, token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csmodule *CsmoduleTransactorSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.RecoverERC721(&_Csmodule.TransactOpts, token, tokenId)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csmodule *CsmoduleTransactor) RecoverEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "recoverEther")
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csmodule *CsmoduleSession) RecoverEther() (*types.Transaction, error) {
	return _Csmodule.Contract.RecoverEther(&_Csmodule.TransactOpts)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csmodule *CsmoduleTransactorSession) RecoverEther() (*types.Transaction, error) {
	return _Csmodule.Contract.RecoverEther(&_Csmodule.TransactOpts)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0x8b3ac71d.
//
// Solidity: function removeKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) returns()
func (_Csmodule *CsmoduleTransactor) RemoveKeys(opts *bind.TransactOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "removeKeys", nodeOperatorId, startIndex, keysCount)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0x8b3ac71d.
//
// Solidity: function removeKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) returns()
func (_Csmodule *CsmoduleSession) RemoveKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.RemoveKeys(&_Csmodule.TransactOpts, nodeOperatorId, startIndex, keysCount)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0x8b3ac71d.
//
// Solidity: function removeKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) returns()
func (_Csmodule *CsmoduleTransactorSession) RemoveKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.RemoveKeys(&_Csmodule.TransactOpts, nodeOperatorId, startIndex, keysCount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csmodule *CsmoduleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csmodule *CsmoduleSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.RenounceRole(&_Csmodule.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csmodule *CsmoduleTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.RenounceRole(&_Csmodule.TransactOpts, role, callerConfirmation)
}

// ReportELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x388dd1d1.
//
// Solidity: function reportELRewardsStealingPenalty(uint256 nodeOperatorId, bytes32 blockHash, uint256 amount) returns()
func (_Csmodule *CsmoduleTransactor) ReportELRewardsStealingPenalty(opts *bind.TransactOpts, nodeOperatorId *big.Int, blockHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "reportELRewardsStealingPenalty", nodeOperatorId, blockHash, amount)
}

// ReportELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x388dd1d1.
//
// Solidity: function reportELRewardsStealingPenalty(uint256 nodeOperatorId, bytes32 blockHash, uint256 amount) returns()
func (_Csmodule *CsmoduleSession) ReportELRewardsStealingPenalty(nodeOperatorId *big.Int, blockHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.ReportELRewardsStealingPenalty(&_Csmodule.TransactOpts, nodeOperatorId, blockHash, amount)
}

// ReportELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x388dd1d1.
//
// Solidity: function reportELRewardsStealingPenalty(uint256 nodeOperatorId, bytes32 blockHash, uint256 amount) returns()
func (_Csmodule *CsmoduleTransactorSession) ReportELRewardsStealingPenalty(nodeOperatorId *big.Int, blockHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.ReportELRewardsStealingPenalty(&_Csmodule.TransactOpts, nodeOperatorId, blockHash, amount)
}

// ReportValidatorExitDelay is a paid mutator transaction binding the contract method 0x57f9c341.
//
// Solidity: function reportValidatorExitDelay(uint256 nodeOperatorId, uint256 , bytes publicKey, uint256 eligibleToExitInSec) returns()
func (_Csmodule *CsmoduleTransactor) ReportValidatorExitDelay(opts *bind.TransactOpts, nodeOperatorId *big.Int, arg1 *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "reportValidatorExitDelay", nodeOperatorId, arg1, publicKey, eligibleToExitInSec)
}

// ReportValidatorExitDelay is a paid mutator transaction binding the contract method 0x57f9c341.
//
// Solidity: function reportValidatorExitDelay(uint256 nodeOperatorId, uint256 , bytes publicKey, uint256 eligibleToExitInSec) returns()
func (_Csmodule *CsmoduleSession) ReportValidatorExitDelay(nodeOperatorId *big.Int, arg1 *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.ReportValidatorExitDelay(&_Csmodule.TransactOpts, nodeOperatorId, arg1, publicKey, eligibleToExitInSec)
}

// ReportValidatorExitDelay is a paid mutator transaction binding the contract method 0x57f9c341.
//
// Solidity: function reportValidatorExitDelay(uint256 nodeOperatorId, uint256 , bytes publicKey, uint256 eligibleToExitInSec) returns()
func (_Csmodule *CsmoduleTransactorSession) ReportValidatorExitDelay(nodeOperatorId *big.Int, arg1 *big.Int, publicKey []byte, eligibleToExitInSec *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.ReportValidatorExitDelay(&_Csmodule.TransactOpts, nodeOperatorId, arg1, publicKey, eligibleToExitInSec)
}

// ResetNodeOperatorManagerAddress is a paid mutator transaction binding the contract method 0x6a6304cc.
//
// Solidity: function resetNodeOperatorManagerAddress(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleTransactor) ResetNodeOperatorManagerAddress(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "resetNodeOperatorManagerAddress", nodeOperatorId)
}

// ResetNodeOperatorManagerAddress is a paid mutator transaction binding the contract method 0x6a6304cc.
//
// Solidity: function resetNodeOperatorManagerAddress(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleSession) ResetNodeOperatorManagerAddress(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.ResetNodeOperatorManagerAddress(&_Csmodule.TransactOpts, nodeOperatorId)
}

// ResetNodeOperatorManagerAddress is a paid mutator transaction binding the contract method 0x6a6304cc.
//
// Solidity: function resetNodeOperatorManagerAddress(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleTransactorSession) ResetNodeOperatorManagerAddress(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.ResetNodeOperatorManagerAddress(&_Csmodule.TransactOpts, nodeOperatorId)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Csmodule *CsmoduleTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Csmodule *CsmoduleSession) Resume() (*types.Transaction, error) {
	return _Csmodule.Contract.Resume(&_Csmodule.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Csmodule *CsmoduleTransactorSession) Resume() (*types.Transaction, error) {
	return _Csmodule.Contract.Resume(&_Csmodule.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csmodule *CsmoduleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csmodule *CsmoduleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.RevokeRole(&_Csmodule.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csmodule *CsmoduleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csmodule.Contract.RevokeRole(&_Csmodule.TransactOpts, role, account)
}

// SettleELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x37b12b5f.
//
// Solidity: function settleELRewardsStealingPenalty(uint256[] nodeOperatorIds) returns()
func (_Csmodule *CsmoduleTransactor) SettleELRewardsStealingPenalty(opts *bind.TransactOpts, nodeOperatorIds []*big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "settleELRewardsStealingPenalty", nodeOperatorIds)
}

// SettleELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x37b12b5f.
//
// Solidity: function settleELRewardsStealingPenalty(uint256[] nodeOperatorIds) returns()
func (_Csmodule *CsmoduleSession) SettleELRewardsStealingPenalty(nodeOperatorIds []*big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.SettleELRewardsStealingPenalty(&_Csmodule.TransactOpts, nodeOperatorIds)
}

// SettleELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x37b12b5f.
//
// Solidity: function settleELRewardsStealingPenalty(uint256[] nodeOperatorIds) returns()
func (_Csmodule *CsmoduleTransactorSession) SettleELRewardsStealingPenalty(nodeOperatorIds []*big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.SettleELRewardsStealingPenalty(&_Csmodule.TransactOpts, nodeOperatorIds)
}

// SubmitWithdrawals is a paid mutator transaction binding the contract method 0x9c963aef.
//
// Solidity: function submitWithdrawals((uint256,uint256,uint256)[] withdrawalsInfo) returns()
func (_Csmodule *CsmoduleTransactor) SubmitWithdrawals(opts *bind.TransactOpts, withdrawalsInfo []ValidatorWithdrawalInfo) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "submitWithdrawals", withdrawalsInfo)
}

// SubmitWithdrawals is a paid mutator transaction binding the contract method 0x9c963aef.
//
// Solidity: function submitWithdrawals((uint256,uint256,uint256)[] withdrawalsInfo) returns()
func (_Csmodule *CsmoduleSession) SubmitWithdrawals(withdrawalsInfo []ValidatorWithdrawalInfo) (*types.Transaction, error) {
	return _Csmodule.Contract.SubmitWithdrawals(&_Csmodule.TransactOpts, withdrawalsInfo)
}

// SubmitWithdrawals is a paid mutator transaction binding the contract method 0x9c963aef.
//
// Solidity: function submitWithdrawals((uint256,uint256,uint256)[] withdrawalsInfo) returns()
func (_Csmodule *CsmoduleTransactorSession) SubmitWithdrawals(withdrawalsInfo []ValidatorWithdrawalInfo) (*types.Transaction, error) {
	return _Csmodule.Contract.SubmitWithdrawals(&_Csmodule.TransactOpts, withdrawalsInfo)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0x94120368.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 nodeOperatorId, uint256 exitedValidatorsKeysCount) returns()
func (_Csmodule *CsmoduleTransactor) UnsafeUpdateValidatorsCount(opts *bind.TransactOpts, nodeOperatorId *big.Int, exitedValidatorsKeysCount *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "unsafeUpdateValidatorsCount", nodeOperatorId, exitedValidatorsKeysCount)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0x94120368.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 nodeOperatorId, uint256 exitedValidatorsKeysCount) returns()
func (_Csmodule *CsmoduleSession) UnsafeUpdateValidatorsCount(nodeOperatorId *big.Int, exitedValidatorsKeysCount *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.UnsafeUpdateValidatorsCount(&_Csmodule.TransactOpts, nodeOperatorId, exitedValidatorsKeysCount)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0x94120368.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 nodeOperatorId, uint256 exitedValidatorsKeysCount) returns()
func (_Csmodule *CsmoduleTransactorSession) UnsafeUpdateValidatorsCount(nodeOperatorId *big.Int, exitedValidatorsKeysCount *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.UnsafeUpdateValidatorsCount(&_Csmodule.TransactOpts, nodeOperatorId, exitedValidatorsKeysCount)
}

// UpdateDepositableValidatorsCount is a paid mutator transaction binding the contract method 0x8eab3cd0.
//
// Solidity: function updateDepositableValidatorsCount(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleTransactor) UpdateDepositableValidatorsCount(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "updateDepositableValidatorsCount", nodeOperatorId)
}

// UpdateDepositableValidatorsCount is a paid mutator transaction binding the contract method 0x8eab3cd0.
//
// Solidity: function updateDepositableValidatorsCount(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleSession) UpdateDepositableValidatorsCount(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.UpdateDepositableValidatorsCount(&_Csmodule.TransactOpts, nodeOperatorId)
}

// UpdateDepositableValidatorsCount is a paid mutator transaction binding the contract method 0x8eab3cd0.
//
// Solidity: function updateDepositableValidatorsCount(uint256 nodeOperatorId) returns()
func (_Csmodule *CsmoduleTransactorSession) UpdateDepositableValidatorsCount(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.UpdateDepositableValidatorsCount(&_Csmodule.TransactOpts, nodeOperatorId)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes nodeOperatorIds, bytes exitedValidatorsCounts) returns()
func (_Csmodule *CsmoduleTransactor) UpdateExitedValidatorsCount(opts *bind.TransactOpts, nodeOperatorIds []byte, exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "updateExitedValidatorsCount", nodeOperatorIds, exitedValidatorsCounts)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes nodeOperatorIds, bytes exitedValidatorsCounts) returns()
func (_Csmodule *CsmoduleSession) UpdateExitedValidatorsCount(nodeOperatorIds []byte, exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _Csmodule.Contract.UpdateExitedValidatorsCount(&_Csmodule.TransactOpts, nodeOperatorIds, exitedValidatorsCounts)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes nodeOperatorIds, bytes exitedValidatorsCounts) returns()
func (_Csmodule *CsmoduleTransactorSession) UpdateExitedValidatorsCount(nodeOperatorIds []byte, exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _Csmodule.Contract.UpdateExitedValidatorsCount(&_Csmodule.TransactOpts, nodeOperatorIds, exitedValidatorsCounts)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x08a679ad.
//
// Solidity: function updateTargetValidatorsLimits(uint256 nodeOperatorId, uint256 targetLimitMode, uint256 targetLimit) returns()
func (_Csmodule *CsmoduleTransactor) UpdateTargetValidatorsLimits(opts *bind.TransactOpts, nodeOperatorId *big.Int, targetLimitMode *big.Int, targetLimit *big.Int) (*types.Transaction, error) {
	return _Csmodule.contract.Transact(opts, "updateTargetValidatorsLimits", nodeOperatorId, targetLimitMode, targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x08a679ad.
//
// Solidity: function updateTargetValidatorsLimits(uint256 nodeOperatorId, uint256 targetLimitMode, uint256 targetLimit) returns()
func (_Csmodule *CsmoduleSession) UpdateTargetValidatorsLimits(nodeOperatorId *big.Int, targetLimitMode *big.Int, targetLimit *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.UpdateTargetValidatorsLimits(&_Csmodule.TransactOpts, nodeOperatorId, targetLimitMode, targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x08a679ad.
//
// Solidity: function updateTargetValidatorsLimits(uint256 nodeOperatorId, uint256 targetLimitMode, uint256 targetLimit) returns()
func (_Csmodule *CsmoduleTransactorSession) UpdateTargetValidatorsLimits(nodeOperatorId *big.Int, targetLimitMode *big.Int, targetLimit *big.Int) (*types.Transaction, error) {
	return _Csmodule.Contract.UpdateTargetValidatorsLimits(&_Csmodule.TransactOpts, nodeOperatorId, targetLimitMode, targetLimit)
}

// CsmoduleBatchEnqueuedIterator is returned from FilterBatchEnqueued and is used to iterate over the raw logs and unpacked data for BatchEnqueued events raised by the Csmodule contract.
type CsmoduleBatchEnqueuedIterator struct {
	Event *CsmoduleBatchEnqueued // Event containing the contract specifics and raw log

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
func (it *CsmoduleBatchEnqueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleBatchEnqueued)
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
		it.Event = new(CsmoduleBatchEnqueued)
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
func (it *CsmoduleBatchEnqueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleBatchEnqueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleBatchEnqueued represents a BatchEnqueued event raised by the Csmodule contract.
type CsmoduleBatchEnqueued struct {
	QueuePriority  *big.Int
	NodeOperatorId *big.Int
	Count          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBatchEnqueued is a free log retrieval operation binding the contract event 0xdc891a44aee443f7f65d1abc5710a05ef241c0c5d7a62f12671522f3c14852bc.
//
// Solidity: event BatchEnqueued(uint256 indexed queuePriority, uint256 indexed nodeOperatorId, uint256 count)
func (_Csmodule *CsmoduleFilterer) FilterBatchEnqueued(opts *bind.FilterOpts, queuePriority []*big.Int, nodeOperatorId []*big.Int) (*CsmoduleBatchEnqueuedIterator, error) {

	var queuePriorityRule []interface{}
	for _, queuePriorityItem := range queuePriority {
		queuePriorityRule = append(queuePriorityRule, queuePriorityItem)
	}
	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "BatchEnqueued", queuePriorityRule, nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleBatchEnqueuedIterator{contract: _Csmodule.contract, event: "BatchEnqueued", logs: logs, sub: sub}, nil
}

// WatchBatchEnqueued is a free log subscription operation binding the contract event 0xdc891a44aee443f7f65d1abc5710a05ef241c0c5d7a62f12671522f3c14852bc.
//
// Solidity: event BatchEnqueued(uint256 indexed queuePriority, uint256 indexed nodeOperatorId, uint256 count)
func (_Csmodule *CsmoduleFilterer) WatchBatchEnqueued(opts *bind.WatchOpts, sink chan<- *CsmoduleBatchEnqueued, queuePriority []*big.Int, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var queuePriorityRule []interface{}
	for _, queuePriorityItem := range queuePriority {
		queuePriorityRule = append(queuePriorityRule, queuePriorityItem)
	}
	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "BatchEnqueued", queuePriorityRule, nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleBatchEnqueued)
				if err := _Csmodule.contract.UnpackLog(event, "BatchEnqueued", log); err != nil {
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

// ParseBatchEnqueued is a log parse operation binding the contract event 0xdc891a44aee443f7f65d1abc5710a05ef241c0c5d7a62f12671522f3c14852bc.
//
// Solidity: event BatchEnqueued(uint256 indexed queuePriority, uint256 indexed nodeOperatorId, uint256 count)
func (_Csmodule *CsmoduleFilterer) ParseBatchEnqueued(log types.Log) (*CsmoduleBatchEnqueued, error) {
	event := new(CsmoduleBatchEnqueued)
	if err := _Csmodule.contract.UnpackLog(event, "BatchEnqueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleDepositableSigningKeysCountChangedIterator is returned from FilterDepositableSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for DepositableSigningKeysCountChanged events raised by the Csmodule contract.
type CsmoduleDepositableSigningKeysCountChangedIterator struct {
	Event *CsmoduleDepositableSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmoduleDepositableSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleDepositableSigningKeysCountChanged)
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
		it.Event = new(CsmoduleDepositableSigningKeysCountChanged)
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
func (it *CsmoduleDepositableSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleDepositableSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleDepositableSigningKeysCountChanged represents a DepositableSigningKeysCountChanged event raised by the Csmodule contract.
type CsmoduleDepositableSigningKeysCountChanged struct {
	NodeOperatorId       *big.Int
	DepositableKeysCount *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDepositableSigningKeysCountChanged is a free log retrieval operation binding the contract event 0xf9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed33.
//
// Solidity: event DepositableSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositableKeysCount)
func (_Csmodule *CsmoduleFilterer) FilterDepositableSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleDepositableSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "DepositableSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleDepositableSigningKeysCountChangedIterator{contract: _Csmodule.contract, event: "DepositableSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchDepositableSigningKeysCountChanged is a free log subscription operation binding the contract event 0xf9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed33.
//
// Solidity: event DepositableSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositableKeysCount)
func (_Csmodule *CsmoduleFilterer) WatchDepositableSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CsmoduleDepositableSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "DepositableSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleDepositableSigningKeysCountChanged)
				if err := _Csmodule.contract.UnpackLog(event, "DepositableSigningKeysCountChanged", log); err != nil {
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

// ParseDepositableSigningKeysCountChanged is a log parse operation binding the contract event 0xf9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed33.
//
// Solidity: event DepositableSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositableKeysCount)
func (_Csmodule *CsmoduleFilterer) ParseDepositableSigningKeysCountChanged(log types.Log) (*CsmoduleDepositableSigningKeysCountChanged, error) {
	event := new(CsmoduleDepositableSigningKeysCountChanged)
	if err := _Csmodule.contract.UnpackLog(event, "DepositableSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleDepositedSigningKeysCountChangedIterator is returned from FilterDepositedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for DepositedSigningKeysCountChanged events raised by the Csmodule contract.
type CsmoduleDepositedSigningKeysCountChangedIterator struct {
	Event *CsmoduleDepositedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmoduleDepositedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleDepositedSigningKeysCountChanged)
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
		it.Event = new(CsmoduleDepositedSigningKeysCountChanged)
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
func (it *CsmoduleDepositedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleDepositedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleDepositedSigningKeysCountChanged represents a DepositedSigningKeysCountChanged event raised by the Csmodule contract.
type CsmoduleDepositedSigningKeysCountChanged struct {
	NodeOperatorId     *big.Int
	DepositedKeysCount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b78.
//
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedKeysCount)
func (_Csmodule *CsmoduleFilterer) FilterDepositedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleDepositedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "DepositedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleDepositedSigningKeysCountChangedIterator{contract: _Csmodule.contract, event: "DepositedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchDepositedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b78.
//
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedKeysCount)
func (_Csmodule *CsmoduleFilterer) WatchDepositedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CsmoduleDepositedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "DepositedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleDepositedSigningKeysCountChanged)
				if err := _Csmodule.contract.UnpackLog(event, "DepositedSigningKeysCountChanged", log); err != nil {
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

// ParseDepositedSigningKeysCountChanged is a log parse operation binding the contract event 0x24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b78.
//
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedKeysCount)
func (_Csmodule *CsmoduleFilterer) ParseDepositedSigningKeysCountChanged(log types.Log) (*CsmoduleDepositedSigningKeysCountChanged, error) {
	event := new(CsmoduleDepositedSigningKeysCountChanged)
	if err := _Csmodule.contract.UnpackLog(event, "DepositedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleELRewardsStealingPenaltyCancelledIterator is returned from FilterELRewardsStealingPenaltyCancelled and is used to iterate over the raw logs and unpacked data for ELRewardsStealingPenaltyCancelled events raised by the Csmodule contract.
type CsmoduleELRewardsStealingPenaltyCancelledIterator struct {
	Event *CsmoduleELRewardsStealingPenaltyCancelled // Event containing the contract specifics and raw log

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
func (it *CsmoduleELRewardsStealingPenaltyCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleELRewardsStealingPenaltyCancelled)
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
		it.Event = new(CsmoduleELRewardsStealingPenaltyCancelled)
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
func (it *CsmoduleELRewardsStealingPenaltyCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleELRewardsStealingPenaltyCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleELRewardsStealingPenaltyCancelled represents a ELRewardsStealingPenaltyCancelled event raised by the Csmodule contract.
type CsmoduleELRewardsStealingPenaltyCancelled struct {
	NodeOperatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterELRewardsStealingPenaltyCancelled is a free log retrieval operation binding the contract event 0x1e7ebd3c5f4de9502000b6f7e6e7cf5d4ecb27d6fe1778e43fb9d1d0ca87d0e7.
//
// Solidity: event ELRewardsStealingPenaltyCancelled(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csmodule *CsmoduleFilterer) FilterELRewardsStealingPenaltyCancelled(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleELRewardsStealingPenaltyCancelledIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "ELRewardsStealingPenaltyCancelled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleELRewardsStealingPenaltyCancelledIterator{contract: _Csmodule.contract, event: "ELRewardsStealingPenaltyCancelled", logs: logs, sub: sub}, nil
}

// WatchELRewardsStealingPenaltyCancelled is a free log subscription operation binding the contract event 0x1e7ebd3c5f4de9502000b6f7e6e7cf5d4ecb27d6fe1778e43fb9d1d0ca87d0e7.
//
// Solidity: event ELRewardsStealingPenaltyCancelled(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csmodule *CsmoduleFilterer) WatchELRewardsStealingPenaltyCancelled(opts *bind.WatchOpts, sink chan<- *CsmoduleELRewardsStealingPenaltyCancelled, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "ELRewardsStealingPenaltyCancelled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleELRewardsStealingPenaltyCancelled)
				if err := _Csmodule.contract.UnpackLog(event, "ELRewardsStealingPenaltyCancelled", log); err != nil {
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

// ParseELRewardsStealingPenaltyCancelled is a log parse operation binding the contract event 0x1e7ebd3c5f4de9502000b6f7e6e7cf5d4ecb27d6fe1778e43fb9d1d0ca87d0e7.
//
// Solidity: event ELRewardsStealingPenaltyCancelled(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csmodule *CsmoduleFilterer) ParseELRewardsStealingPenaltyCancelled(log types.Log) (*CsmoduleELRewardsStealingPenaltyCancelled, error) {
	event := new(CsmoduleELRewardsStealingPenaltyCancelled)
	if err := _Csmodule.contract.UnpackLog(event, "ELRewardsStealingPenaltyCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleELRewardsStealingPenaltyCompensatedIterator is returned from FilterELRewardsStealingPenaltyCompensated and is used to iterate over the raw logs and unpacked data for ELRewardsStealingPenaltyCompensated events raised by the Csmodule contract.
type CsmoduleELRewardsStealingPenaltyCompensatedIterator struct {
	Event *CsmoduleELRewardsStealingPenaltyCompensated // Event containing the contract specifics and raw log

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
func (it *CsmoduleELRewardsStealingPenaltyCompensatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleELRewardsStealingPenaltyCompensated)
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
		it.Event = new(CsmoduleELRewardsStealingPenaltyCompensated)
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
func (it *CsmoduleELRewardsStealingPenaltyCompensatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleELRewardsStealingPenaltyCompensatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleELRewardsStealingPenaltyCompensated represents a ELRewardsStealingPenaltyCompensated event raised by the Csmodule contract.
type CsmoduleELRewardsStealingPenaltyCompensated struct {
	NodeOperatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterELRewardsStealingPenaltyCompensated is a free log retrieval operation binding the contract event 0xb1858b4c2ab6242521725a8f7350a6cb22ad4ecae009c9b63ef114baffb054be.
//
// Solidity: event ELRewardsStealingPenaltyCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csmodule *CsmoduleFilterer) FilterELRewardsStealingPenaltyCompensated(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleELRewardsStealingPenaltyCompensatedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "ELRewardsStealingPenaltyCompensated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleELRewardsStealingPenaltyCompensatedIterator{contract: _Csmodule.contract, event: "ELRewardsStealingPenaltyCompensated", logs: logs, sub: sub}, nil
}

// WatchELRewardsStealingPenaltyCompensated is a free log subscription operation binding the contract event 0xb1858b4c2ab6242521725a8f7350a6cb22ad4ecae009c9b63ef114baffb054be.
//
// Solidity: event ELRewardsStealingPenaltyCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csmodule *CsmoduleFilterer) WatchELRewardsStealingPenaltyCompensated(opts *bind.WatchOpts, sink chan<- *CsmoduleELRewardsStealingPenaltyCompensated, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "ELRewardsStealingPenaltyCompensated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleELRewardsStealingPenaltyCompensated)
				if err := _Csmodule.contract.UnpackLog(event, "ELRewardsStealingPenaltyCompensated", log); err != nil {
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

// ParseELRewardsStealingPenaltyCompensated is a log parse operation binding the contract event 0xb1858b4c2ab6242521725a8f7350a6cb22ad4ecae009c9b63ef114baffb054be.
//
// Solidity: event ELRewardsStealingPenaltyCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csmodule *CsmoduleFilterer) ParseELRewardsStealingPenaltyCompensated(log types.Log) (*CsmoduleELRewardsStealingPenaltyCompensated, error) {
	event := new(CsmoduleELRewardsStealingPenaltyCompensated)
	if err := _Csmodule.contract.UnpackLog(event, "ELRewardsStealingPenaltyCompensated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleELRewardsStealingPenaltyReportedIterator is returned from FilterELRewardsStealingPenaltyReported and is used to iterate over the raw logs and unpacked data for ELRewardsStealingPenaltyReported events raised by the Csmodule contract.
type CsmoduleELRewardsStealingPenaltyReportedIterator struct {
	Event *CsmoduleELRewardsStealingPenaltyReported // Event containing the contract specifics and raw log

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
func (it *CsmoduleELRewardsStealingPenaltyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleELRewardsStealingPenaltyReported)
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
		it.Event = new(CsmoduleELRewardsStealingPenaltyReported)
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
func (it *CsmoduleELRewardsStealingPenaltyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleELRewardsStealingPenaltyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleELRewardsStealingPenaltyReported represents a ELRewardsStealingPenaltyReported event raised by the Csmodule contract.
type CsmoduleELRewardsStealingPenaltyReported struct {
	NodeOperatorId    *big.Int
	ProposedBlockHash [32]byte
	StolenAmount      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterELRewardsStealingPenaltyReported is a free log retrieval operation binding the contract event 0xeec4d6dbe34149c6728a9638eca869d0e5a7fcd85c7a96178f7e9780b4b7fe4b.
//
// Solidity: event ELRewardsStealingPenaltyReported(uint256 indexed nodeOperatorId, bytes32 proposedBlockHash, uint256 stolenAmount)
func (_Csmodule *CsmoduleFilterer) FilterELRewardsStealingPenaltyReported(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleELRewardsStealingPenaltyReportedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "ELRewardsStealingPenaltyReported", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleELRewardsStealingPenaltyReportedIterator{contract: _Csmodule.contract, event: "ELRewardsStealingPenaltyReported", logs: logs, sub: sub}, nil
}

// WatchELRewardsStealingPenaltyReported is a free log subscription operation binding the contract event 0xeec4d6dbe34149c6728a9638eca869d0e5a7fcd85c7a96178f7e9780b4b7fe4b.
//
// Solidity: event ELRewardsStealingPenaltyReported(uint256 indexed nodeOperatorId, bytes32 proposedBlockHash, uint256 stolenAmount)
func (_Csmodule *CsmoduleFilterer) WatchELRewardsStealingPenaltyReported(opts *bind.WatchOpts, sink chan<- *CsmoduleELRewardsStealingPenaltyReported, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "ELRewardsStealingPenaltyReported", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleELRewardsStealingPenaltyReported)
				if err := _Csmodule.contract.UnpackLog(event, "ELRewardsStealingPenaltyReported", log); err != nil {
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

// ParseELRewardsStealingPenaltyReported is a log parse operation binding the contract event 0xeec4d6dbe34149c6728a9638eca869d0e5a7fcd85c7a96178f7e9780b4b7fe4b.
//
// Solidity: event ELRewardsStealingPenaltyReported(uint256 indexed nodeOperatorId, bytes32 proposedBlockHash, uint256 stolenAmount)
func (_Csmodule *CsmoduleFilterer) ParseELRewardsStealingPenaltyReported(log types.Log) (*CsmoduleELRewardsStealingPenaltyReported, error) {
	event := new(CsmoduleELRewardsStealingPenaltyReported)
	if err := _Csmodule.contract.UnpackLog(event, "ELRewardsStealingPenaltyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleELRewardsStealingPenaltySettledIterator is returned from FilterELRewardsStealingPenaltySettled and is used to iterate over the raw logs and unpacked data for ELRewardsStealingPenaltySettled events raised by the Csmodule contract.
type CsmoduleELRewardsStealingPenaltySettledIterator struct {
	Event *CsmoduleELRewardsStealingPenaltySettled // Event containing the contract specifics and raw log

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
func (it *CsmoduleELRewardsStealingPenaltySettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleELRewardsStealingPenaltySettled)
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
		it.Event = new(CsmoduleELRewardsStealingPenaltySettled)
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
func (it *CsmoduleELRewardsStealingPenaltySettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleELRewardsStealingPenaltySettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleELRewardsStealingPenaltySettled represents a ELRewardsStealingPenaltySettled event raised by the Csmodule contract.
type CsmoduleELRewardsStealingPenaltySettled struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterELRewardsStealingPenaltySettled is a free log retrieval operation binding the contract event 0x00f4fe19c0404d2fbb58da6f646c0a3ee5a6994a034213bbd22b072ed1ca5c27.
//
// Solidity: event ELRewardsStealingPenaltySettled(uint256 indexed nodeOperatorId)
func (_Csmodule *CsmoduleFilterer) FilterELRewardsStealingPenaltySettled(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleELRewardsStealingPenaltySettledIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "ELRewardsStealingPenaltySettled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleELRewardsStealingPenaltySettledIterator{contract: _Csmodule.contract, event: "ELRewardsStealingPenaltySettled", logs: logs, sub: sub}, nil
}

// WatchELRewardsStealingPenaltySettled is a free log subscription operation binding the contract event 0x00f4fe19c0404d2fbb58da6f646c0a3ee5a6994a034213bbd22b072ed1ca5c27.
//
// Solidity: event ELRewardsStealingPenaltySettled(uint256 indexed nodeOperatorId)
func (_Csmodule *CsmoduleFilterer) WatchELRewardsStealingPenaltySettled(opts *bind.WatchOpts, sink chan<- *CsmoduleELRewardsStealingPenaltySettled, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "ELRewardsStealingPenaltySettled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleELRewardsStealingPenaltySettled)
				if err := _Csmodule.contract.UnpackLog(event, "ELRewardsStealingPenaltySettled", log); err != nil {
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

// ParseELRewardsStealingPenaltySettled is a log parse operation binding the contract event 0x00f4fe19c0404d2fbb58da6f646c0a3ee5a6994a034213bbd22b072ed1ca5c27.
//
// Solidity: event ELRewardsStealingPenaltySettled(uint256 indexed nodeOperatorId)
func (_Csmodule *CsmoduleFilterer) ParseELRewardsStealingPenaltySettled(log types.Log) (*CsmoduleELRewardsStealingPenaltySettled, error) {
	event := new(CsmoduleELRewardsStealingPenaltySettled)
	if err := _Csmodule.contract.UnpackLog(event, "ELRewardsStealingPenaltySettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleERC1155RecoveredIterator is returned from FilterERC1155Recovered and is used to iterate over the raw logs and unpacked data for ERC1155Recovered events raised by the Csmodule contract.
type CsmoduleERC1155RecoveredIterator struct {
	Event *CsmoduleERC1155Recovered // Event containing the contract specifics and raw log

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
func (it *CsmoduleERC1155RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleERC1155Recovered)
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
		it.Event = new(CsmoduleERC1155Recovered)
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
func (it *CsmoduleERC1155RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleERC1155RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleERC1155Recovered represents a ERC1155Recovered event raised by the Csmodule contract.
type CsmoduleERC1155Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC1155Recovered is a free log retrieval operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csmodule *CsmoduleFilterer) FilterERC1155Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsmoduleERC1155RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleERC1155RecoveredIterator{contract: _Csmodule.contract, event: "ERC1155Recovered", logs: logs, sub: sub}, nil
}

// WatchERC1155Recovered is a free log subscription operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csmodule *CsmoduleFilterer) WatchERC1155Recovered(opts *bind.WatchOpts, sink chan<- *CsmoduleERC1155Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleERC1155Recovered)
				if err := _Csmodule.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
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
func (_Csmodule *CsmoduleFilterer) ParseERC1155Recovered(log types.Log) (*CsmoduleERC1155Recovered, error) {
	event := new(CsmoduleERC1155Recovered)
	if err := _Csmodule.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleERC20RecoveredIterator is returned from FilterERC20Recovered and is used to iterate over the raw logs and unpacked data for ERC20Recovered events raised by the Csmodule contract.
type CsmoduleERC20RecoveredIterator struct {
	Event *CsmoduleERC20Recovered // Event containing the contract specifics and raw log

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
func (it *CsmoduleERC20RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleERC20Recovered)
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
		it.Event = new(CsmoduleERC20Recovered)
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
func (it *CsmoduleERC20RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleERC20RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleERC20Recovered represents a ERC20Recovered event raised by the Csmodule contract.
type CsmoduleERC20Recovered struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Recovered is a free log retrieval operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csmodule *CsmoduleFilterer) FilterERC20Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsmoduleERC20RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleERC20RecoveredIterator{contract: _Csmodule.contract, event: "ERC20Recovered", logs: logs, sub: sub}, nil
}

// WatchERC20Recovered is a free log subscription operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csmodule *CsmoduleFilterer) WatchERC20Recovered(opts *bind.WatchOpts, sink chan<- *CsmoduleERC20Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleERC20Recovered)
				if err := _Csmodule.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
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
func (_Csmodule *CsmoduleFilterer) ParseERC20Recovered(log types.Log) (*CsmoduleERC20Recovered, error) {
	event := new(CsmoduleERC20Recovered)
	if err := _Csmodule.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleERC721RecoveredIterator is returned from FilterERC721Recovered and is used to iterate over the raw logs and unpacked data for ERC721Recovered events raised by the Csmodule contract.
type CsmoduleERC721RecoveredIterator struct {
	Event *CsmoduleERC721Recovered // Event containing the contract specifics and raw log

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
func (it *CsmoduleERC721RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleERC721Recovered)
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
		it.Event = new(CsmoduleERC721Recovered)
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
func (it *CsmoduleERC721RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleERC721RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleERC721Recovered represents a ERC721Recovered event raised by the Csmodule contract.
type CsmoduleERC721Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721Recovered is a free log retrieval operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csmodule *CsmoduleFilterer) FilterERC721Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsmoduleERC721RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleERC721RecoveredIterator{contract: _Csmodule.contract, event: "ERC721Recovered", logs: logs, sub: sub}, nil
}

// WatchERC721Recovered is a free log subscription operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csmodule *CsmoduleFilterer) WatchERC721Recovered(opts *bind.WatchOpts, sink chan<- *CsmoduleERC721Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleERC721Recovered)
				if err := _Csmodule.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
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
func (_Csmodule *CsmoduleFilterer) ParseERC721Recovered(log types.Log) (*CsmoduleERC721Recovered, error) {
	event := new(CsmoduleERC721Recovered)
	if err := _Csmodule.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleEtherRecoveredIterator is returned from FilterEtherRecovered and is used to iterate over the raw logs and unpacked data for EtherRecovered events raised by the Csmodule contract.
type CsmoduleEtherRecoveredIterator struct {
	Event *CsmoduleEtherRecovered // Event containing the contract specifics and raw log

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
func (it *CsmoduleEtherRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleEtherRecovered)
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
		it.Event = new(CsmoduleEtherRecovered)
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
func (it *CsmoduleEtherRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleEtherRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleEtherRecovered represents a EtherRecovered event raised by the Csmodule contract.
type CsmoduleEtherRecovered struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEtherRecovered is a free log retrieval operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csmodule *CsmoduleFilterer) FilterEtherRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CsmoduleEtherRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleEtherRecoveredIterator{contract: _Csmodule.contract, event: "EtherRecovered", logs: logs, sub: sub}, nil
}

// WatchEtherRecovered is a free log subscription operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csmodule *CsmoduleFilterer) WatchEtherRecovered(opts *bind.WatchOpts, sink chan<- *CsmoduleEtherRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleEtherRecovered)
				if err := _Csmodule.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
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
func (_Csmodule *CsmoduleFilterer) ParseEtherRecovered(log types.Log) (*CsmoduleEtherRecovered, error) {
	event := new(CsmoduleEtherRecovered)
	if err := _Csmodule.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleExitedSigningKeysCountChangedIterator is returned from FilterExitedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for ExitedSigningKeysCountChanged events raised by the Csmodule contract.
type CsmoduleExitedSigningKeysCountChangedIterator struct {
	Event *CsmoduleExitedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmoduleExitedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleExitedSigningKeysCountChanged)
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
		it.Event = new(CsmoduleExitedSigningKeysCountChanged)
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
func (it *CsmoduleExitedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleExitedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleExitedSigningKeysCountChanged represents a ExitedSigningKeysCountChanged event raised by the Csmodule contract.
type CsmoduleExitedSigningKeysCountChanged struct {
	NodeOperatorId  *big.Int
	ExitedKeysCount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExitedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c220166.
//
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedKeysCount)
func (_Csmodule *CsmoduleFilterer) FilterExitedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleExitedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "ExitedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleExitedSigningKeysCountChangedIterator{contract: _Csmodule.contract, event: "ExitedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchExitedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c220166.
//
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedKeysCount)
func (_Csmodule *CsmoduleFilterer) WatchExitedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CsmoduleExitedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "ExitedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleExitedSigningKeysCountChanged)
				if err := _Csmodule.contract.UnpackLog(event, "ExitedSigningKeysCountChanged", log); err != nil {
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

// ParseExitedSigningKeysCountChanged is a log parse operation binding the contract event 0x0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c220166.
//
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedKeysCount)
func (_Csmodule *CsmoduleFilterer) ParseExitedSigningKeysCountChanged(log types.Log) (*CsmoduleExitedSigningKeysCountChanged, error) {
	event := new(CsmoduleExitedSigningKeysCountChanged)
	if err := _Csmodule.contract.UnpackLog(event, "ExitedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Csmodule contract.
type CsmoduleInitializedIterator struct {
	Event *CsmoduleInitialized // Event containing the contract specifics and raw log

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
func (it *CsmoduleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleInitialized)
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
		it.Event = new(CsmoduleInitialized)
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
func (it *CsmoduleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleInitialized represents a Initialized event raised by the Csmodule contract.
type CsmoduleInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csmodule *CsmoduleFilterer) FilterInitialized(opts *bind.FilterOpts) (*CsmoduleInitializedIterator, error) {

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CsmoduleInitializedIterator{contract: _Csmodule.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csmodule *CsmoduleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CsmoduleInitialized) (event.Subscription, error) {

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleInitialized)
				if err := _Csmodule.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Csmodule *CsmoduleFilterer) ParseInitialized(log types.Log) (*CsmoduleInitialized, error) {
	event := new(CsmoduleInitialized)
	if err := _Csmodule.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleKeyRemovalChargeAppliedIterator is returned from FilterKeyRemovalChargeApplied and is used to iterate over the raw logs and unpacked data for KeyRemovalChargeApplied events raised by the Csmodule contract.
type CsmoduleKeyRemovalChargeAppliedIterator struct {
	Event *CsmoduleKeyRemovalChargeApplied // Event containing the contract specifics and raw log

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
func (it *CsmoduleKeyRemovalChargeAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleKeyRemovalChargeApplied)
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
		it.Event = new(CsmoduleKeyRemovalChargeApplied)
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
func (it *CsmoduleKeyRemovalChargeAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleKeyRemovalChargeAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleKeyRemovalChargeApplied represents a KeyRemovalChargeApplied event raised by the Csmodule contract.
type CsmoduleKeyRemovalChargeApplied struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKeyRemovalChargeApplied is a free log retrieval operation binding the contract event 0x1cbb8dafbedbdf4f813a8ed1f50d871def63e1104f8729b677af57905eda90f6.
//
// Solidity: event KeyRemovalChargeApplied(uint256 indexed nodeOperatorId)
func (_Csmodule *CsmoduleFilterer) FilterKeyRemovalChargeApplied(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleKeyRemovalChargeAppliedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "KeyRemovalChargeApplied", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleKeyRemovalChargeAppliedIterator{contract: _Csmodule.contract, event: "KeyRemovalChargeApplied", logs: logs, sub: sub}, nil
}

// WatchKeyRemovalChargeApplied is a free log subscription operation binding the contract event 0x1cbb8dafbedbdf4f813a8ed1f50d871def63e1104f8729b677af57905eda90f6.
//
// Solidity: event KeyRemovalChargeApplied(uint256 indexed nodeOperatorId)
func (_Csmodule *CsmoduleFilterer) WatchKeyRemovalChargeApplied(opts *bind.WatchOpts, sink chan<- *CsmoduleKeyRemovalChargeApplied, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "KeyRemovalChargeApplied", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleKeyRemovalChargeApplied)
				if err := _Csmodule.contract.UnpackLog(event, "KeyRemovalChargeApplied", log); err != nil {
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

// ParseKeyRemovalChargeApplied is a log parse operation binding the contract event 0x1cbb8dafbedbdf4f813a8ed1f50d871def63e1104f8729b677af57905eda90f6.
//
// Solidity: event KeyRemovalChargeApplied(uint256 indexed nodeOperatorId)
func (_Csmodule *CsmoduleFilterer) ParseKeyRemovalChargeApplied(log types.Log) (*CsmoduleKeyRemovalChargeApplied, error) {
	event := new(CsmoduleKeyRemovalChargeApplied)
	if err := _Csmodule.contract.UnpackLog(event, "KeyRemovalChargeApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleNodeOperatorAddedIterator is returned from FilterNodeOperatorAdded and is used to iterate over the raw logs and unpacked data for NodeOperatorAdded events raised by the Csmodule contract.
type CsmoduleNodeOperatorAddedIterator struct {
	Event *CsmoduleNodeOperatorAdded // Event containing the contract specifics and raw log

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
func (it *CsmoduleNodeOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleNodeOperatorAdded)
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
		it.Event = new(CsmoduleNodeOperatorAdded)
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
func (it *CsmoduleNodeOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleNodeOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleNodeOperatorAdded represents a NodeOperatorAdded event raised by the Csmodule contract.
type CsmoduleNodeOperatorAdded struct {
	NodeOperatorId             *big.Int
	ManagerAddress             common.Address
	RewardAddress              common.Address
	ExtendedManagerPermissions bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorAdded is a free log retrieval operation binding the contract event 0xf17baf73d46b0a80157c3ea3dda1bf081a702732d53ff1720f85e55d9f0997c0.
//
// Solidity: event NodeOperatorAdded(uint256 indexed nodeOperatorId, address indexed managerAddress, address indexed rewardAddress, bool extendedManagerPermissions)
func (_Csmodule *CsmoduleFilterer) FilterNodeOperatorAdded(opts *bind.FilterOpts, nodeOperatorId []*big.Int, managerAddress []common.Address, rewardAddress []common.Address) (*CsmoduleNodeOperatorAddedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var managerAddressRule []interface{}
	for _, managerAddressItem := range managerAddress {
		managerAddressRule = append(managerAddressRule, managerAddressItem)
	}
	var rewardAddressRule []interface{}
	for _, rewardAddressItem := range rewardAddress {
		rewardAddressRule = append(rewardAddressRule, rewardAddressItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "NodeOperatorAdded", nodeOperatorIdRule, managerAddressRule, rewardAddressRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleNodeOperatorAddedIterator{contract: _Csmodule.contract, event: "NodeOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorAdded is a free log subscription operation binding the contract event 0xf17baf73d46b0a80157c3ea3dda1bf081a702732d53ff1720f85e55d9f0997c0.
//
// Solidity: event NodeOperatorAdded(uint256 indexed nodeOperatorId, address indexed managerAddress, address indexed rewardAddress, bool extendedManagerPermissions)
func (_Csmodule *CsmoduleFilterer) WatchNodeOperatorAdded(opts *bind.WatchOpts, sink chan<- *CsmoduleNodeOperatorAdded, nodeOperatorId []*big.Int, managerAddress []common.Address, rewardAddress []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var managerAddressRule []interface{}
	for _, managerAddressItem := range managerAddress {
		managerAddressRule = append(managerAddressRule, managerAddressItem)
	}
	var rewardAddressRule []interface{}
	for _, rewardAddressItem := range rewardAddress {
		rewardAddressRule = append(rewardAddressRule, rewardAddressItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "NodeOperatorAdded", nodeOperatorIdRule, managerAddressRule, rewardAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleNodeOperatorAdded)
				if err := _Csmodule.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
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

// ParseNodeOperatorAdded is a log parse operation binding the contract event 0xf17baf73d46b0a80157c3ea3dda1bf081a702732d53ff1720f85e55d9f0997c0.
//
// Solidity: event NodeOperatorAdded(uint256 indexed nodeOperatorId, address indexed managerAddress, address indexed rewardAddress, bool extendedManagerPermissions)
func (_Csmodule *CsmoduleFilterer) ParseNodeOperatorAdded(log types.Log) (*CsmoduleNodeOperatorAdded, error) {
	event := new(CsmoduleNodeOperatorAdded)
	if err := _Csmodule.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleNodeOperatorManagerAddressChangeProposedIterator is returned from FilterNodeOperatorManagerAddressChangeProposed and is used to iterate over the raw logs and unpacked data for NodeOperatorManagerAddressChangeProposed events raised by the Csmodule contract.
type CsmoduleNodeOperatorManagerAddressChangeProposedIterator struct {
	Event *CsmoduleNodeOperatorManagerAddressChangeProposed // Event containing the contract specifics and raw log

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
func (it *CsmoduleNodeOperatorManagerAddressChangeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleNodeOperatorManagerAddressChangeProposed)
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
		it.Event = new(CsmoduleNodeOperatorManagerAddressChangeProposed)
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
func (it *CsmoduleNodeOperatorManagerAddressChangeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleNodeOperatorManagerAddressChangeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleNodeOperatorManagerAddressChangeProposed represents a NodeOperatorManagerAddressChangeProposed event raised by the Csmodule contract.
type CsmoduleNodeOperatorManagerAddressChangeProposed struct {
	NodeOperatorId     *big.Int
	OldProposedAddress common.Address
	NewProposedAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorManagerAddressChangeProposed is a free log retrieval operation binding the contract event 0x4048f15a706950765ca59f99d0fa6fe8edaaa3f3e3d0337417082e2131df82fb.
//
// Solidity: event NodeOperatorManagerAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_Csmodule *CsmoduleFilterer) FilterNodeOperatorManagerAddressChangeProposed(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (*CsmoduleNodeOperatorManagerAddressChangeProposedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldProposedAddressRule []interface{}
	for _, oldProposedAddressItem := range oldProposedAddress {
		oldProposedAddressRule = append(oldProposedAddressRule, oldProposedAddressItem)
	}
	var newProposedAddressRule []interface{}
	for _, newProposedAddressItem := range newProposedAddress {
		newProposedAddressRule = append(newProposedAddressRule, newProposedAddressItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "NodeOperatorManagerAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleNodeOperatorManagerAddressChangeProposedIterator{contract: _Csmodule.contract, event: "NodeOperatorManagerAddressChangeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorManagerAddressChangeProposed is a free log subscription operation binding the contract event 0x4048f15a706950765ca59f99d0fa6fe8edaaa3f3e3d0337417082e2131df82fb.
//
// Solidity: event NodeOperatorManagerAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_Csmodule *CsmoduleFilterer) WatchNodeOperatorManagerAddressChangeProposed(opts *bind.WatchOpts, sink chan<- *CsmoduleNodeOperatorManagerAddressChangeProposed, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldProposedAddressRule []interface{}
	for _, oldProposedAddressItem := range oldProposedAddress {
		oldProposedAddressRule = append(oldProposedAddressRule, oldProposedAddressItem)
	}
	var newProposedAddressRule []interface{}
	for _, newProposedAddressItem := range newProposedAddress {
		newProposedAddressRule = append(newProposedAddressRule, newProposedAddressItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "NodeOperatorManagerAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleNodeOperatorManagerAddressChangeProposed)
				if err := _Csmodule.contract.UnpackLog(event, "NodeOperatorManagerAddressChangeProposed", log); err != nil {
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

// ParseNodeOperatorManagerAddressChangeProposed is a log parse operation binding the contract event 0x4048f15a706950765ca59f99d0fa6fe8edaaa3f3e3d0337417082e2131df82fb.
//
// Solidity: event NodeOperatorManagerAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_Csmodule *CsmoduleFilterer) ParseNodeOperatorManagerAddressChangeProposed(log types.Log) (*CsmoduleNodeOperatorManagerAddressChangeProposed, error) {
	event := new(CsmoduleNodeOperatorManagerAddressChangeProposed)
	if err := _Csmodule.contract.UnpackLog(event, "NodeOperatorManagerAddressChangeProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleNodeOperatorManagerAddressChangedIterator is returned from FilterNodeOperatorManagerAddressChanged and is used to iterate over the raw logs and unpacked data for NodeOperatorManagerAddressChanged events raised by the Csmodule contract.
type CsmoduleNodeOperatorManagerAddressChangedIterator struct {
	Event *CsmoduleNodeOperatorManagerAddressChanged // Event containing the contract specifics and raw log

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
func (it *CsmoduleNodeOperatorManagerAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleNodeOperatorManagerAddressChanged)
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
		it.Event = new(CsmoduleNodeOperatorManagerAddressChanged)
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
func (it *CsmoduleNodeOperatorManagerAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleNodeOperatorManagerAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleNodeOperatorManagerAddressChanged represents a NodeOperatorManagerAddressChanged event raised by the Csmodule contract.
type CsmoduleNodeOperatorManagerAddressChanged struct {
	NodeOperatorId *big.Int
	OldAddress     common.Address
	NewAddress     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorManagerAddressChanged is a free log retrieval operation binding the contract event 0x862021f23449d6e8516867bd839be15a3d8698a7561c5c2c35069074b7e91e61.
//
// Solidity: event NodeOperatorManagerAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_Csmodule *CsmoduleFilterer) FilterNodeOperatorManagerAddressChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (*CsmoduleNodeOperatorManagerAddressChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "NodeOperatorManagerAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleNodeOperatorManagerAddressChangedIterator{contract: _Csmodule.contract, event: "NodeOperatorManagerAddressChanged", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorManagerAddressChanged is a free log subscription operation binding the contract event 0x862021f23449d6e8516867bd839be15a3d8698a7561c5c2c35069074b7e91e61.
//
// Solidity: event NodeOperatorManagerAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_Csmodule *CsmoduleFilterer) WatchNodeOperatorManagerAddressChanged(opts *bind.WatchOpts, sink chan<- *CsmoduleNodeOperatorManagerAddressChanged, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "NodeOperatorManagerAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleNodeOperatorManagerAddressChanged)
				if err := _Csmodule.contract.UnpackLog(event, "NodeOperatorManagerAddressChanged", log); err != nil {
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

// ParseNodeOperatorManagerAddressChanged is a log parse operation binding the contract event 0x862021f23449d6e8516867bd839be15a3d8698a7561c5c2c35069074b7e91e61.
//
// Solidity: event NodeOperatorManagerAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_Csmodule *CsmoduleFilterer) ParseNodeOperatorManagerAddressChanged(log types.Log) (*CsmoduleNodeOperatorManagerAddressChanged, error) {
	event := new(CsmoduleNodeOperatorManagerAddressChanged)
	if err := _Csmodule.contract.UnpackLog(event, "NodeOperatorManagerAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleNodeOperatorRewardAddressChangeProposedIterator is returned from FilterNodeOperatorRewardAddressChangeProposed and is used to iterate over the raw logs and unpacked data for NodeOperatorRewardAddressChangeProposed events raised by the Csmodule contract.
type CsmoduleNodeOperatorRewardAddressChangeProposedIterator struct {
	Event *CsmoduleNodeOperatorRewardAddressChangeProposed // Event containing the contract specifics and raw log

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
func (it *CsmoduleNodeOperatorRewardAddressChangeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleNodeOperatorRewardAddressChangeProposed)
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
		it.Event = new(CsmoduleNodeOperatorRewardAddressChangeProposed)
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
func (it *CsmoduleNodeOperatorRewardAddressChangeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleNodeOperatorRewardAddressChangeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleNodeOperatorRewardAddressChangeProposed represents a NodeOperatorRewardAddressChangeProposed event raised by the Csmodule contract.
type CsmoduleNodeOperatorRewardAddressChangeProposed struct {
	NodeOperatorId     *big.Int
	OldProposedAddress common.Address
	NewProposedAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRewardAddressChangeProposed is a free log retrieval operation binding the contract event 0xb5878cdb1d66f971efe3b138a71c64bc5bc519314db2533e0e4cde954409ea5a.
//
// Solidity: event NodeOperatorRewardAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_Csmodule *CsmoduleFilterer) FilterNodeOperatorRewardAddressChangeProposed(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (*CsmoduleNodeOperatorRewardAddressChangeProposedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldProposedAddressRule []interface{}
	for _, oldProposedAddressItem := range oldProposedAddress {
		oldProposedAddressRule = append(oldProposedAddressRule, oldProposedAddressItem)
	}
	var newProposedAddressRule []interface{}
	for _, newProposedAddressItem := range newProposedAddress {
		newProposedAddressRule = append(newProposedAddressRule, newProposedAddressItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "NodeOperatorRewardAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleNodeOperatorRewardAddressChangeProposedIterator{contract: _Csmodule.contract, event: "NodeOperatorRewardAddressChangeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRewardAddressChangeProposed is a free log subscription operation binding the contract event 0xb5878cdb1d66f971efe3b138a71c64bc5bc519314db2533e0e4cde954409ea5a.
//
// Solidity: event NodeOperatorRewardAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_Csmodule *CsmoduleFilterer) WatchNodeOperatorRewardAddressChangeProposed(opts *bind.WatchOpts, sink chan<- *CsmoduleNodeOperatorRewardAddressChangeProposed, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldProposedAddressRule []interface{}
	for _, oldProposedAddressItem := range oldProposedAddress {
		oldProposedAddressRule = append(oldProposedAddressRule, oldProposedAddressItem)
	}
	var newProposedAddressRule []interface{}
	for _, newProposedAddressItem := range newProposedAddress {
		newProposedAddressRule = append(newProposedAddressRule, newProposedAddressItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "NodeOperatorRewardAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleNodeOperatorRewardAddressChangeProposed)
				if err := _Csmodule.contract.UnpackLog(event, "NodeOperatorRewardAddressChangeProposed", log); err != nil {
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

// ParseNodeOperatorRewardAddressChangeProposed is a log parse operation binding the contract event 0xb5878cdb1d66f971efe3b138a71c64bc5bc519314db2533e0e4cde954409ea5a.
//
// Solidity: event NodeOperatorRewardAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_Csmodule *CsmoduleFilterer) ParseNodeOperatorRewardAddressChangeProposed(log types.Log) (*CsmoduleNodeOperatorRewardAddressChangeProposed, error) {
	event := new(CsmoduleNodeOperatorRewardAddressChangeProposed)
	if err := _Csmodule.contract.UnpackLog(event, "NodeOperatorRewardAddressChangeProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleNodeOperatorRewardAddressChangedIterator is returned from FilterNodeOperatorRewardAddressChanged and is used to iterate over the raw logs and unpacked data for NodeOperatorRewardAddressChanged events raised by the Csmodule contract.
type CsmoduleNodeOperatorRewardAddressChangedIterator struct {
	Event *CsmoduleNodeOperatorRewardAddressChanged // Event containing the contract specifics and raw log

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
func (it *CsmoduleNodeOperatorRewardAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleNodeOperatorRewardAddressChanged)
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
		it.Event = new(CsmoduleNodeOperatorRewardAddressChanged)
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
func (it *CsmoduleNodeOperatorRewardAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleNodeOperatorRewardAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleNodeOperatorRewardAddressChanged represents a NodeOperatorRewardAddressChanged event raised by the Csmodule contract.
type CsmoduleNodeOperatorRewardAddressChanged struct {
	NodeOperatorId *big.Int
	OldAddress     common.Address
	NewAddress     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRewardAddressChanged is a free log retrieval operation binding the contract event 0x069ac7cd8230db015b7250c8e5425149cf1a3e912d9569f497165e55b3b6b7b2.
//
// Solidity: event NodeOperatorRewardAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_Csmodule *CsmoduleFilterer) FilterNodeOperatorRewardAddressChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (*CsmoduleNodeOperatorRewardAddressChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "NodeOperatorRewardAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleNodeOperatorRewardAddressChangedIterator{contract: _Csmodule.contract, event: "NodeOperatorRewardAddressChanged", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRewardAddressChanged is a free log subscription operation binding the contract event 0x069ac7cd8230db015b7250c8e5425149cf1a3e912d9569f497165e55b3b6b7b2.
//
// Solidity: event NodeOperatorRewardAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_Csmodule *CsmoduleFilterer) WatchNodeOperatorRewardAddressChanged(opts *bind.WatchOpts, sink chan<- *CsmoduleNodeOperatorRewardAddressChanged, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var oldAddressRule []interface{}
	for _, oldAddressItem := range oldAddress {
		oldAddressRule = append(oldAddressRule, oldAddressItem)
	}
	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "NodeOperatorRewardAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleNodeOperatorRewardAddressChanged)
				if err := _Csmodule.contract.UnpackLog(event, "NodeOperatorRewardAddressChanged", log); err != nil {
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

// ParseNodeOperatorRewardAddressChanged is a log parse operation binding the contract event 0x069ac7cd8230db015b7250c8e5425149cf1a3e912d9569f497165e55b3b6b7b2.
//
// Solidity: event NodeOperatorRewardAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_Csmodule *CsmoduleFilterer) ParseNodeOperatorRewardAddressChanged(log types.Log) (*CsmoduleNodeOperatorRewardAddressChanged, error) {
	event := new(CsmoduleNodeOperatorRewardAddressChanged)
	if err := _Csmodule.contract.UnpackLog(event, "NodeOperatorRewardAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleNonceChangedIterator is returned from FilterNonceChanged and is used to iterate over the raw logs and unpacked data for NonceChanged events raised by the Csmodule contract.
type CsmoduleNonceChangedIterator struct {
	Event *CsmoduleNonceChanged // Event containing the contract specifics and raw log

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
func (it *CsmoduleNonceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleNonceChanged)
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
		it.Event = new(CsmoduleNonceChanged)
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
func (it *CsmoduleNonceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleNonceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleNonceChanged represents a NonceChanged event raised by the Csmodule contract.
type CsmoduleNonceChanged struct {
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNonceChanged is a free log retrieval operation binding the contract event 0x7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa.
//
// Solidity: event NonceChanged(uint256 nonce)
func (_Csmodule *CsmoduleFilterer) FilterNonceChanged(opts *bind.FilterOpts) (*CsmoduleNonceChangedIterator, error) {

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "NonceChanged")
	if err != nil {
		return nil, err
	}
	return &CsmoduleNonceChangedIterator{contract: _Csmodule.contract, event: "NonceChanged", logs: logs, sub: sub}, nil
}

// WatchNonceChanged is a free log subscription operation binding the contract event 0x7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa.
//
// Solidity: event NonceChanged(uint256 nonce)
func (_Csmodule *CsmoduleFilterer) WatchNonceChanged(opts *bind.WatchOpts, sink chan<- *CsmoduleNonceChanged) (event.Subscription, error) {

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "NonceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleNonceChanged)
				if err := _Csmodule.contract.UnpackLog(event, "NonceChanged", log); err != nil {
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

// ParseNonceChanged is a log parse operation binding the contract event 0x7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa.
//
// Solidity: event NonceChanged(uint256 nonce)
func (_Csmodule *CsmoduleFilterer) ParseNonceChanged(log types.Log) (*CsmoduleNonceChanged, error) {
	event := new(CsmoduleNonceChanged)
	if err := _Csmodule.contract.UnpackLog(event, "NonceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmodulePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Csmodule contract.
type CsmodulePausedIterator struct {
	Event *CsmodulePaused // Event containing the contract specifics and raw log

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
func (it *CsmodulePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmodulePaused)
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
		it.Event = new(CsmodulePaused)
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
func (it *CsmodulePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmodulePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmodulePaused represents a Paused event raised by the Csmodule contract.
type CsmodulePaused struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Csmodule *CsmoduleFilterer) FilterPaused(opts *bind.FilterOpts) (*CsmodulePausedIterator, error) {

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CsmodulePausedIterator{contract: _Csmodule.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Csmodule *CsmoduleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CsmodulePaused) (event.Subscription, error) {

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmodulePaused)
				if err := _Csmodule.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Csmodule *CsmoduleFilterer) ParsePaused(log types.Log) (*CsmodulePaused, error) {
	event := new(CsmodulePaused)
	if err := _Csmodule.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleReferrerSetIterator is returned from FilterReferrerSet and is used to iterate over the raw logs and unpacked data for ReferrerSet events raised by the Csmodule contract.
type CsmoduleReferrerSetIterator struct {
	Event *CsmoduleReferrerSet // Event containing the contract specifics and raw log

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
func (it *CsmoduleReferrerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleReferrerSet)
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
		it.Event = new(CsmoduleReferrerSet)
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
func (it *CsmoduleReferrerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleReferrerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleReferrerSet represents a ReferrerSet event raised by the Csmodule contract.
type CsmoduleReferrerSet struct {
	NodeOperatorId *big.Int
	Referrer       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReferrerSet is a free log retrieval operation binding the contract event 0x67334334c388385e5f244703f8a8b28b7f4ffe52909130aca69bc62a8e27f09a.
//
// Solidity: event ReferrerSet(uint256 indexed nodeOperatorId, address indexed referrer)
func (_Csmodule *CsmoduleFilterer) FilterReferrerSet(opts *bind.FilterOpts, nodeOperatorId []*big.Int, referrer []common.Address) (*CsmoduleReferrerSetIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "ReferrerSet", nodeOperatorIdRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleReferrerSetIterator{contract: _Csmodule.contract, event: "ReferrerSet", logs: logs, sub: sub}, nil
}

// WatchReferrerSet is a free log subscription operation binding the contract event 0x67334334c388385e5f244703f8a8b28b7f4ffe52909130aca69bc62a8e27f09a.
//
// Solidity: event ReferrerSet(uint256 indexed nodeOperatorId, address indexed referrer)
func (_Csmodule *CsmoduleFilterer) WatchReferrerSet(opts *bind.WatchOpts, sink chan<- *CsmoduleReferrerSet, nodeOperatorId []*big.Int, referrer []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "ReferrerSet", nodeOperatorIdRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleReferrerSet)
				if err := _Csmodule.contract.UnpackLog(event, "ReferrerSet", log); err != nil {
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

// ParseReferrerSet is a log parse operation binding the contract event 0x67334334c388385e5f244703f8a8b28b7f4ffe52909130aca69bc62a8e27f09a.
//
// Solidity: event ReferrerSet(uint256 indexed nodeOperatorId, address indexed referrer)
func (_Csmodule *CsmoduleFilterer) ParseReferrerSet(log types.Log) (*CsmoduleReferrerSet, error) {
	event := new(CsmoduleReferrerSet)
	if err := _Csmodule.contract.UnpackLog(event, "ReferrerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Csmodule contract.
type CsmoduleResumedIterator struct {
	Event *CsmoduleResumed // Event containing the contract specifics and raw log

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
func (it *CsmoduleResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleResumed)
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
		it.Event = new(CsmoduleResumed)
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
func (it *CsmoduleResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleResumed represents a Resumed event raised by the Csmodule contract.
type CsmoduleResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Csmodule *CsmoduleFilterer) FilterResumed(opts *bind.FilterOpts) (*CsmoduleResumedIterator, error) {

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &CsmoduleResumedIterator{contract: _Csmodule.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Csmodule *CsmoduleFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *CsmoduleResumed) (event.Subscription, error) {

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleResumed)
				if err := _Csmodule.contract.UnpackLog(event, "Resumed", log); err != nil {
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

// ParseResumed is a log parse operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Csmodule *CsmoduleFilterer) ParseResumed(log types.Log) (*CsmoduleResumed, error) {
	event := new(CsmoduleResumed)
	if err := _Csmodule.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Csmodule contract.
type CsmoduleRoleAdminChangedIterator struct {
	Event *CsmoduleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CsmoduleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleRoleAdminChanged)
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
		it.Event = new(CsmoduleRoleAdminChanged)
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
func (it *CsmoduleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleRoleAdminChanged represents a RoleAdminChanged event raised by the Csmodule contract.
type CsmoduleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Csmodule *CsmoduleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CsmoduleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleRoleAdminChangedIterator{contract: _Csmodule.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Csmodule *CsmoduleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CsmoduleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleRoleAdminChanged)
				if err := _Csmodule.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Csmodule *CsmoduleFilterer) ParseRoleAdminChanged(log types.Log) (*CsmoduleRoleAdminChanged, error) {
	event := new(CsmoduleRoleAdminChanged)
	if err := _Csmodule.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Csmodule contract.
type CsmoduleRoleGrantedIterator struct {
	Event *CsmoduleRoleGranted // Event containing the contract specifics and raw log

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
func (it *CsmoduleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleRoleGranted)
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
		it.Event = new(CsmoduleRoleGranted)
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
func (it *CsmoduleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleRoleGranted represents a RoleGranted event raised by the Csmodule contract.
type CsmoduleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csmodule *CsmoduleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CsmoduleRoleGrantedIterator, error) {

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

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleRoleGrantedIterator{contract: _Csmodule.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csmodule *CsmoduleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CsmoduleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleRoleGranted)
				if err := _Csmodule.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Csmodule *CsmoduleFilterer) ParseRoleGranted(log types.Log) (*CsmoduleRoleGranted, error) {
	event := new(CsmoduleRoleGranted)
	if err := _Csmodule.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Csmodule contract.
type CsmoduleRoleRevokedIterator struct {
	Event *CsmoduleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CsmoduleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleRoleRevoked)
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
		it.Event = new(CsmoduleRoleRevoked)
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
func (it *CsmoduleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleRoleRevoked represents a RoleRevoked event raised by the Csmodule contract.
type CsmoduleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csmodule *CsmoduleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CsmoduleRoleRevokedIterator, error) {

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

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleRoleRevokedIterator{contract: _Csmodule.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csmodule *CsmoduleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CsmoduleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleRoleRevoked)
				if err := _Csmodule.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Csmodule *CsmoduleFilterer) ParseRoleRevoked(log types.Log) (*CsmoduleRoleRevoked, error) {
	event := new(CsmoduleRoleRevoked)
	if err := _Csmodule.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleSigningKeyAddedIterator is returned from FilterSigningKeyAdded and is used to iterate over the raw logs and unpacked data for SigningKeyAdded events raised by the Csmodule contract.
type CsmoduleSigningKeyAddedIterator struct {
	Event *CsmoduleSigningKeyAdded // Event containing the contract specifics and raw log

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
func (it *CsmoduleSigningKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleSigningKeyAdded)
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
		it.Event = new(CsmoduleSigningKeyAdded)
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
func (it *CsmoduleSigningKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleSigningKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleSigningKeyAdded represents a SigningKeyAdded event raised by the Csmodule contract.
type CsmoduleSigningKeyAdded struct {
	NodeOperatorId *big.Int
	Pubkey         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSigningKeyAdded is a free log retrieval operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed nodeOperatorId, bytes pubkey)
func (_Csmodule *CsmoduleFilterer) FilterSigningKeyAdded(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleSigningKeyAddedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "SigningKeyAdded", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleSigningKeyAddedIterator{contract: _Csmodule.contract, event: "SigningKeyAdded", logs: logs, sub: sub}, nil
}

// WatchSigningKeyAdded is a free log subscription operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed nodeOperatorId, bytes pubkey)
func (_Csmodule *CsmoduleFilterer) WatchSigningKeyAdded(opts *bind.WatchOpts, sink chan<- *CsmoduleSigningKeyAdded, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "SigningKeyAdded", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleSigningKeyAdded)
				if err := _Csmodule.contract.UnpackLog(event, "SigningKeyAdded", log); err != nil {
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

// ParseSigningKeyAdded is a log parse operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed nodeOperatorId, bytes pubkey)
func (_Csmodule *CsmoduleFilterer) ParseSigningKeyAdded(log types.Log) (*CsmoduleSigningKeyAdded, error) {
	event := new(CsmoduleSigningKeyAdded)
	if err := _Csmodule.contract.UnpackLog(event, "SigningKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleSigningKeyRemovedIterator is returned from FilterSigningKeyRemoved and is used to iterate over the raw logs and unpacked data for SigningKeyRemoved events raised by the Csmodule contract.
type CsmoduleSigningKeyRemovedIterator struct {
	Event *CsmoduleSigningKeyRemoved // Event containing the contract specifics and raw log

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
func (it *CsmoduleSigningKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleSigningKeyRemoved)
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
		it.Event = new(CsmoduleSigningKeyRemoved)
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
func (it *CsmoduleSigningKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleSigningKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleSigningKeyRemoved represents a SigningKeyRemoved event raised by the Csmodule contract.
type CsmoduleSigningKeyRemoved struct {
	NodeOperatorId *big.Int
	Pubkey         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSigningKeyRemoved is a free log retrieval operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed nodeOperatorId, bytes pubkey)
func (_Csmodule *CsmoduleFilterer) FilterSigningKeyRemoved(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleSigningKeyRemovedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "SigningKeyRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleSigningKeyRemovedIterator{contract: _Csmodule.contract, event: "SigningKeyRemoved", logs: logs, sub: sub}, nil
}

// WatchSigningKeyRemoved is a free log subscription operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed nodeOperatorId, bytes pubkey)
func (_Csmodule *CsmoduleFilterer) WatchSigningKeyRemoved(opts *bind.WatchOpts, sink chan<- *CsmoduleSigningKeyRemoved, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "SigningKeyRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleSigningKeyRemoved)
				if err := _Csmodule.contract.UnpackLog(event, "SigningKeyRemoved", log); err != nil {
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

// ParseSigningKeyRemoved is a log parse operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed nodeOperatorId, bytes pubkey)
func (_Csmodule *CsmoduleFilterer) ParseSigningKeyRemoved(log types.Log) (*CsmoduleSigningKeyRemoved, error) {
	event := new(CsmoduleSigningKeyRemoved)
	if err := _Csmodule.contract.UnpackLog(event, "SigningKeyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleStETHSharesRecoveredIterator is returned from FilterStETHSharesRecovered and is used to iterate over the raw logs and unpacked data for StETHSharesRecovered events raised by the Csmodule contract.
type CsmoduleStETHSharesRecoveredIterator struct {
	Event *CsmoduleStETHSharesRecovered // Event containing the contract specifics and raw log

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
func (it *CsmoduleStETHSharesRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleStETHSharesRecovered)
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
		it.Event = new(CsmoduleStETHSharesRecovered)
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
func (it *CsmoduleStETHSharesRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleStETHSharesRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleStETHSharesRecovered represents a StETHSharesRecovered event raised by the Csmodule contract.
type CsmoduleStETHSharesRecovered struct {
	Recipient common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStETHSharesRecovered is a free log retrieval operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csmodule *CsmoduleFilterer) FilterStETHSharesRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CsmoduleStETHSharesRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleStETHSharesRecoveredIterator{contract: _Csmodule.contract, event: "StETHSharesRecovered", logs: logs, sub: sub}, nil
}

// WatchStETHSharesRecovered is a free log subscription operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csmodule *CsmoduleFilterer) WatchStETHSharesRecovered(opts *bind.WatchOpts, sink chan<- *CsmoduleStETHSharesRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleStETHSharesRecovered)
				if err := _Csmodule.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
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
func (_Csmodule *CsmoduleFilterer) ParseStETHSharesRecovered(log types.Log) (*CsmoduleStETHSharesRecovered, error) {
	event := new(CsmoduleStETHSharesRecovered)
	if err := _Csmodule.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleTargetValidatorsCountChangedIterator is returned from FilterTargetValidatorsCountChanged and is used to iterate over the raw logs and unpacked data for TargetValidatorsCountChanged events raised by the Csmodule contract.
type CsmoduleTargetValidatorsCountChangedIterator struct {
	Event *CsmoduleTargetValidatorsCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmoduleTargetValidatorsCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleTargetValidatorsCountChanged)
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
		it.Event = new(CsmoduleTargetValidatorsCountChanged)
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
func (it *CsmoduleTargetValidatorsCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleTargetValidatorsCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleTargetValidatorsCountChanged represents a TargetValidatorsCountChanged event raised by the Csmodule contract.
type CsmoduleTargetValidatorsCountChanged struct {
	NodeOperatorId        *big.Int
	TargetLimitMode       *big.Int
	TargetValidatorsCount *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTargetValidatorsCountChanged is a free log retrieval operation binding the contract event 0xf92eb109ce5b449e9b121c352c6aeb4319538a90738cb95d84f08e41274e92d2.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetLimitMode, uint256 targetValidatorsCount)
func (_Csmodule *CsmoduleFilterer) FilterTargetValidatorsCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleTargetValidatorsCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "TargetValidatorsCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleTargetValidatorsCountChangedIterator{contract: _Csmodule.contract, event: "TargetValidatorsCountChanged", logs: logs, sub: sub}, nil
}

// WatchTargetValidatorsCountChanged is a free log subscription operation binding the contract event 0xf92eb109ce5b449e9b121c352c6aeb4319538a90738cb95d84f08e41274e92d2.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetLimitMode, uint256 targetValidatorsCount)
func (_Csmodule *CsmoduleFilterer) WatchTargetValidatorsCountChanged(opts *bind.WatchOpts, sink chan<- *CsmoduleTargetValidatorsCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "TargetValidatorsCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleTargetValidatorsCountChanged)
				if err := _Csmodule.contract.UnpackLog(event, "TargetValidatorsCountChanged", log); err != nil {
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

// ParseTargetValidatorsCountChanged is a log parse operation binding the contract event 0xf92eb109ce5b449e9b121c352c6aeb4319538a90738cb95d84f08e41274e92d2.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetLimitMode, uint256 targetValidatorsCount)
func (_Csmodule *CsmoduleFilterer) ParseTargetValidatorsCountChanged(log types.Log) (*CsmoduleTargetValidatorsCountChanged, error) {
	event := new(CsmoduleTargetValidatorsCountChanged)
	if err := _Csmodule.contract.UnpackLog(event, "TargetValidatorsCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleTotalSigningKeysCountChangedIterator is returned from FilterTotalSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for TotalSigningKeysCountChanged events raised by the Csmodule contract.
type CsmoduleTotalSigningKeysCountChangedIterator struct {
	Event *CsmoduleTotalSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmoduleTotalSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleTotalSigningKeysCountChanged)
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
		it.Event = new(CsmoduleTotalSigningKeysCountChanged)
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
func (it *CsmoduleTotalSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleTotalSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleTotalSigningKeysCountChanged represents a TotalSigningKeysCountChanged event raised by the Csmodule contract.
type CsmoduleTotalSigningKeysCountChanged struct {
	NodeOperatorId *big.Int
	TotalKeysCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalSigningKeysCountChanged is a free log retrieval operation binding the contract event 0xdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f0.
//
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalKeysCount)
func (_Csmodule *CsmoduleFilterer) FilterTotalSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleTotalSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "TotalSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleTotalSigningKeysCountChangedIterator{contract: _Csmodule.contract, event: "TotalSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchTotalSigningKeysCountChanged is a free log subscription operation binding the contract event 0xdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f0.
//
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalKeysCount)
func (_Csmodule *CsmoduleFilterer) WatchTotalSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CsmoduleTotalSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "TotalSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleTotalSigningKeysCountChanged)
				if err := _Csmodule.contract.UnpackLog(event, "TotalSigningKeysCountChanged", log); err != nil {
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

// ParseTotalSigningKeysCountChanged is a log parse operation binding the contract event 0xdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f0.
//
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalKeysCount)
func (_Csmodule *CsmoduleFilterer) ParseTotalSigningKeysCountChanged(log types.Log) (*CsmoduleTotalSigningKeysCountChanged, error) {
	event := new(CsmoduleTotalSigningKeysCountChanged)
	if err := _Csmodule.contract.UnpackLog(event, "TotalSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleVettedSigningKeysCountChangedIterator is returned from FilterVettedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for VettedSigningKeysCountChanged events raised by the Csmodule contract.
type CsmoduleVettedSigningKeysCountChangedIterator struct {
	Event *CsmoduleVettedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CsmoduleVettedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleVettedSigningKeysCountChanged)
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
		it.Event = new(CsmoduleVettedSigningKeysCountChanged)
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
func (it *CsmoduleVettedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleVettedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleVettedSigningKeysCountChanged represents a VettedSigningKeysCountChanged event raised by the Csmodule contract.
type CsmoduleVettedSigningKeysCountChanged struct {
	NodeOperatorId  *big.Int
	VettedKeysCount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVettedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd.
//
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 vettedKeysCount)
func (_Csmodule *CsmoduleFilterer) FilterVettedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleVettedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "VettedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleVettedSigningKeysCountChangedIterator{contract: _Csmodule.contract, event: "VettedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchVettedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd.
//
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 vettedKeysCount)
func (_Csmodule *CsmoduleFilterer) WatchVettedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CsmoduleVettedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "VettedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleVettedSigningKeysCountChanged)
				if err := _Csmodule.contract.UnpackLog(event, "VettedSigningKeysCountChanged", log); err != nil {
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

// ParseVettedSigningKeysCountChanged is a log parse operation binding the contract event 0x947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd.
//
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 vettedKeysCount)
func (_Csmodule *CsmoduleFilterer) ParseVettedSigningKeysCountChanged(log types.Log) (*CsmoduleVettedSigningKeysCountChanged, error) {
	event := new(CsmoduleVettedSigningKeysCountChanged)
	if err := _Csmodule.contract.UnpackLog(event, "VettedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleVettedSigningKeysCountDecreasedIterator is returned from FilterVettedSigningKeysCountDecreased and is used to iterate over the raw logs and unpacked data for VettedSigningKeysCountDecreased events raised by the Csmodule contract.
type CsmoduleVettedSigningKeysCountDecreasedIterator struct {
	Event *CsmoduleVettedSigningKeysCountDecreased // Event containing the contract specifics and raw log

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
func (it *CsmoduleVettedSigningKeysCountDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleVettedSigningKeysCountDecreased)
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
		it.Event = new(CsmoduleVettedSigningKeysCountDecreased)
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
func (it *CsmoduleVettedSigningKeysCountDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleVettedSigningKeysCountDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleVettedSigningKeysCountDecreased represents a VettedSigningKeysCountDecreased event raised by the Csmodule contract.
type CsmoduleVettedSigningKeysCountDecreased struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVettedSigningKeysCountDecreased is a free log retrieval operation binding the contract event 0xe5725d045d5c47bd1483feba445e395dc8647486963e6d54aad9ed03ff7d6ce6.
//
// Solidity: event VettedSigningKeysCountDecreased(uint256 indexed nodeOperatorId)
func (_Csmodule *CsmoduleFilterer) FilterVettedSigningKeysCountDecreased(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleVettedSigningKeysCountDecreasedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "VettedSigningKeysCountDecreased", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleVettedSigningKeysCountDecreasedIterator{contract: _Csmodule.contract, event: "VettedSigningKeysCountDecreased", logs: logs, sub: sub}, nil
}

// WatchVettedSigningKeysCountDecreased is a free log subscription operation binding the contract event 0xe5725d045d5c47bd1483feba445e395dc8647486963e6d54aad9ed03ff7d6ce6.
//
// Solidity: event VettedSigningKeysCountDecreased(uint256 indexed nodeOperatorId)
func (_Csmodule *CsmoduleFilterer) WatchVettedSigningKeysCountDecreased(opts *bind.WatchOpts, sink chan<- *CsmoduleVettedSigningKeysCountDecreased, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "VettedSigningKeysCountDecreased", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleVettedSigningKeysCountDecreased)
				if err := _Csmodule.contract.UnpackLog(event, "VettedSigningKeysCountDecreased", log); err != nil {
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

// ParseVettedSigningKeysCountDecreased is a log parse operation binding the contract event 0xe5725d045d5c47bd1483feba445e395dc8647486963e6d54aad9ed03ff7d6ce6.
//
// Solidity: event VettedSigningKeysCountDecreased(uint256 indexed nodeOperatorId)
func (_Csmodule *CsmoduleFilterer) ParseVettedSigningKeysCountDecreased(log types.Log) (*CsmoduleVettedSigningKeysCountDecreased, error) {
	event := new(CsmoduleVettedSigningKeysCountDecreased)
	if err := _Csmodule.contract.UnpackLog(event, "VettedSigningKeysCountDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsmoduleWithdrawalSubmittedIterator is returned from FilterWithdrawalSubmitted and is used to iterate over the raw logs and unpacked data for WithdrawalSubmitted events raised by the Csmodule contract.
type CsmoduleWithdrawalSubmittedIterator struct {
	Event *CsmoduleWithdrawalSubmitted // Event containing the contract specifics and raw log

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
func (it *CsmoduleWithdrawalSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsmoduleWithdrawalSubmitted)
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
		it.Event = new(CsmoduleWithdrawalSubmitted)
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
func (it *CsmoduleWithdrawalSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsmoduleWithdrawalSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsmoduleWithdrawalSubmitted represents a WithdrawalSubmitted event raised by the Csmodule contract.
type CsmoduleWithdrawalSubmitted struct {
	NodeOperatorId *big.Int
	KeyIndex       *big.Int
	Amount         *big.Int
	Pubkey         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalSubmitted is a free log retrieval operation binding the contract event 0x9bc54857932b6f10bb750fdad91f736b82dd4de202ed5c2f9f076773bb5b3fb7.
//
// Solidity: event WithdrawalSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex, uint256 amount, bytes pubkey)
func (_Csmodule *CsmoduleFilterer) FilterWithdrawalSubmitted(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsmoduleWithdrawalSubmittedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.FilterLogs(opts, "WithdrawalSubmitted", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsmoduleWithdrawalSubmittedIterator{contract: _Csmodule.contract, event: "WithdrawalSubmitted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalSubmitted is a free log subscription operation binding the contract event 0x9bc54857932b6f10bb750fdad91f736b82dd4de202ed5c2f9f076773bb5b3fb7.
//
// Solidity: event WithdrawalSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex, uint256 amount, bytes pubkey)
func (_Csmodule *CsmoduleFilterer) WatchWithdrawalSubmitted(opts *bind.WatchOpts, sink chan<- *CsmoduleWithdrawalSubmitted, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csmodule.contract.WatchLogs(opts, "WithdrawalSubmitted", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsmoduleWithdrawalSubmitted)
				if err := _Csmodule.contract.UnpackLog(event, "WithdrawalSubmitted", log); err != nil {
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

// ParseWithdrawalSubmitted is a log parse operation binding the contract event 0x9bc54857932b6f10bb750fdad91f736b82dd4de202ed5c2f9f076773bb5b3fb7.
//
// Solidity: event WithdrawalSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex, uint256 amount, bytes pubkey)
func (_Csmodule *CsmoduleFilterer) ParseWithdrawalSubmitted(log types.Log) (*CsmoduleWithdrawalSubmitted, error) {
	event := new(CsmoduleWithdrawalSubmitted)
	if err := _Csmodule.contract.UnpackLog(event, "WithdrawalSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
