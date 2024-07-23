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
}

// NodeOperatorManagementProperties is an auto generated low-level Go binding around an user-defined struct.
type NodeOperatorManagementProperties struct {
	ManagerAddress             common.Address
	RewardAddress              common.Address
	ExtendedManagerPermissions bool
}

// CSModuleMetaData contains all meta data concerning the CSModule contract.
var CSModuleMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"moduleType\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"minSlashingPenaltyQuotient\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"elRewardsStealingFine\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxKeysPerOperatorEA\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lidoLocator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyActivated\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyProposed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AlreadyWithdrawn\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"EmptyKey\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitedKeysDecrease\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitedKeysHigherThanTotalDeposited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidKeysCount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidReportData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidVetKeysPointer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxSigningKeysCountExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MethodCallIsNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeOperatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToJoinYet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToRecover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughKeys\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotSupported\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseUntilMustBeInFuture\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedExpected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueIsEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueLookupNoLimit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResumedExpected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SameAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotEligible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotManagerAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotProposedAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotRewardAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SigningKeysInvalidOffset\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"StuckKeysHigherThanNonExited\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAccountingAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroLocatorAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPauseDuration\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"name\":\"BatchEnqueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositableKeysCount\",\"type\":\"uint256\"}],\"name\":\"DepositableSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositedKeysCount\",\"type\":\"uint256\"}],\"name\":\"DepositedSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsStealingPenaltyCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"proposedBlockHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stolenAmount\",\"type\":\"uint256\"}],\"name\":\"ELRewardsStealingPenaltyReported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"ELRewardsStealingPenaltySettled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitedKeysCount\",\"type\":\"uint256\"}],\"name\":\"ExitedSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"name\":\"InitialSlashingSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"KeyRemovalChargeApplied\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"KeyRemovalChargeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProposedAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProposedAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorManagerAddressChangeProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorManagerAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldProposedAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newProposedAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorRewardAddressChangeProposed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"NodeOperatorRewardAddressChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"nonce\",\"type\":\"uint256\"}],\"name\":\"NonceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"PublicRelease\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"ReferrerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"SigningKeyAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"}],\"name\":\"SigningKeyRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"StETHSharesRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stuckKeysCount\",\"type\":\"uint256\"}],\"name\":\"StuckSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"}],\"name\":\"TargetValidatorsCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalKeysCount\",\"type\":\"uint256\"}],\"name\":\"TotalSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"vettedKeysCount\",\"type\":\"uint256\"}],\"name\":\"VettedSigningKeysCountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"VettedSigningKeysCountDecreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"WithdrawalSubmitted\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EL_REWARDS_STEALING_FINE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INITIAL_SLASHING_PENALTY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIDO_LOCATOR\",\"outputs\":[{\"internalType\":\"contractILidoLocator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_SIGNING_KEYS_PER_OPERATOR_BEFORE_PUBLIC_RELEASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_INFINITELY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECOVERER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"REPORT_EL_REWARDS_STEALING_PENALTY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STAKING_ROUTER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"STETH\",\"outputs\":[{\"internalType\":\"contractIStETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERIFIER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"accounting\",\"outputs\":[{\"internalType\":\"contractICSAccounting\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"activatePublicRelease\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extendedManagerPermissions\",\"type\":\"bool\"}],\"internalType\":\"structNodeOperatorManagementProperties\",\"name\":\"managementProperties\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"eaProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"addNodeOperatorETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extendedManagerPermissions\",\"type\":\"bool\"}],\"internalType\":\"structNodeOperatorManagementProperties\",\"name\":\"managementProperties\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"eaProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"addNodeOperatorStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extendedManagerPermissions\",\"type\":\"bool\"}],\"internalType\":\"structNodeOperatorManagementProperties\",\"name\":\"managementProperties\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"},{\"internalType\":\"bytes32[]\",\"name\":\"eaProof\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"}],\"name\":\"addNodeOperatorWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"name\":\"addValidatorKeysETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"addValidatorKeysStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"addValidatorKeysWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"cancelELRewardsStealingPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"changeNodeOperatorRewardAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stETHAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimRewardsStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stEthAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimRewardsUnstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wstETHAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimRewardsWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxItems\",\"type\":\"uint256\"}],\"name\":\"cleanDepositQueue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"compensateELRewardsStealingPenalty\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"confirmNodeOperatorManagerAddressChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"confirmNodeOperatorRewardAddressChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeOperatorIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"vettedSigningKeysCounts\",\"type\":\"bytes\"}],\"name\":\"decreaseVettedSigningKeysCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"depositQueue\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"head\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"tail\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"index\",\"type\":\"uint128\"}],\"name\":\"depositQueueItem\",\"outputs\":[{\"internalType\":\"Batch\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stETHAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"depositStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wstETHAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"depositWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"earlyAdoption\",\"outputs\":[{\"internalType\":\"contractICSEarlyAdoption\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveNodeOperatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperator\",\"outputs\":[{\"components\":[{\"internalType\":\"uint32\",\"name\":\"totalAddedKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalWithdrawnKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalDepositedKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"totalVettedKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"targetLimit\",\"type\":\"uint32\"},{\"internalType\":\"uint8\",\"name\":\"targetLimitMode\",\"type\":\"uint8\"},{\"internalType\":\"uint32\",\"name\":\"totalExitedKeys\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"enqueuedCount\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"managerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedManagerAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"proposedRewardAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"extendedManagerPermissions\",\"type\":\"bool\"}],\"internalType\":\"structNodeOperator\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"offset\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nodeOperatorIds\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorIsActive\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorNonWithdrawnKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getNodeOperatorSummary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refundedValidatorsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckPenaltyEndTimestamp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNodeOperatorsCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonce\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResumeSinceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"}],\"name\":\"getSigningKeys\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"}],\"name\":\"getSigningKeysWithSignatures\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"keys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStakingModuleSummary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalExitedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDepositedValidators\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"depositableValidatorsCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getType\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_accounting\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_earlyAdoption\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_keyRemovalCharge\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"name\":\"isValidatorSlashed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"name\":\"isValidatorWithdrawn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"keyRemovalCharge\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"normalizeQueue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositsCount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"obtainDepositData\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"publicKeys\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signatures\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onExitedAndStuckValidatorsCountsUpdated\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalShares\",\"type\":\"uint256\"}],\"name\":\"onRewardsMinted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"onWithdrawalCredentialsChanged\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"pauseFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"proposeNodeOperatorManagerAddressChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"proposedAddress\",\"type\":\"address\"}],\"name\":\"proposeNodeOperatorRewardAddressChange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"publicRelease\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverStETHShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"startIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"}],\"name\":\"removeKeys\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"blockHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"reportELRewardsStealingPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"resetNodeOperatorManagerAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"setKeyRemovalCharge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"nodeOperatorIds\",\"type\":\"uint256[]\"}],\"name\":\"settleELRewardsStealingPenalty\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"}],\"name\":\"submitInitialSlashing\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"keyIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSlashed\",\"type\":\"bool\"}],\"name\":\"submitWithdrawal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitedValidatorsKeysCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stuckValidatorsKeysCount\",\"type\":\"uint256\"}],\"name\":\"unsafeUpdateValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeOperatorIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"exitedValidatorsCounts\",\"type\":\"bytes\"}],\"name\":\"updateExitedValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"updateRefundedValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"nodeOperatorIds\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"stuckValidatorsCounts\",\"type\":\"bytes\"}],\"name\":\"updateStuckValidatorsCount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetLimitMode\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"targetLimit\",\"type\":\"uint256\"}],\"name\":\"updateTargetValidatorsLimits\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052600436106104e7575f3560e01c80638b3ac71d11610283578063ba1557ae11610155578063dbba4b48116100c9578063f2e2ca6311610083578063f2e2ca6314611224578063f3f449c714611243578063f408b55114611262578063f617eecc14611281578063f96d3952146112cb578063fe7ed3cd146112ea575f80fd5b8063dbba4b4814611138578063e00bfe501461116b578063e1aa105d1461119e578063e21a430b146111bd578063e7705db6146111dd578063e864299e14611210575f80fd5b8063cb17ed3e1161011a578063cb17ed3e14611077578063d087d288146110aa578063d3931457146110be578063d547741f146110d2578063d6477919146110f1578063d9df8c9214611124575f80fd5b8063ba1557ae14610fc8578063bdac46a214610fe7578063be2030941461101a578063bee41b5814611039578063ca15c87314611058575f80fd5b80639b3d1900116101f7578063acc446eb116101b1578063acc446eb14610eca578063acf1c94814610ee9578063b1520dc514610f1c578063b187bd2614610f3b578063b3076c3c14610f4f578063b643189b14610fa9575f80fd5b80639b3d190014610e465780639ec3c24c14610e65578063a217fddf14610e84578063a2e080f114610e97578063a302ee3814610eb6578063a70c70e414610c3e575f80fd5b806390c09bdb1161024857806390c09bdb14610d7157806391d1485414610d85578063946654ad14610da45780639624e83e14610dc35780639abddf0914610de25780639b00c14614610e27575f80fd5b80638b3ac71d14610cb45780638cabe95914610cd35780638d7e401714610cf25780638ec6902814610d115780639010d07c14610d52575f80fd5b80635204281c116103bc5780636a5f2c4a1161033057806380231f15116102ea57806380231f1514610be0578063819d4cc614610c005780638409d4fe14610c1f5780638469cbd314610c3e5780638573e35114610c625780638980f11f14610c95575f80fd5b80636a5f2c4a14610b325780636a6304cc14610b515780636bb1bfdf14610b705780636efe37a214610b8f578063735dfa2814610ba257806375a401da14610bc1575f80fd5b806359e25c121161038157806359e25c12146108e85780635a73bdc8146109145780635c654ad9146109285780635e169fb8146109475780635e2fb9081461096657806365c14dc714610997575f80fd5b80635204281c1461086f57806352d8bfc21461088e57806353433643146108a25780635358fbda146108c1578063589ff76c146108d4575f80fd5b806337b12b5f1161045e5780633f214bb2116104185780633f214bb21461078657806340044801146107a557806347faf311146107c45780634febc81b146107f757806350388cb6146108235780635097ef5914610850575f80fd5b806337b12b5f146106a4578063388dd1d1146106c3578063389ed267146106e25780633dbe8b5a146107155780633df6c438146107345780633f04f0c814610753575f80fd5b80631b40b231116104af5780631b40b231146105a3578063248a9ca3146105c257806326a666e4146105fc5780632de03aa1146106335780632f2ff15d1461066657806336568abe14610685575f80fd5b806301ffc9a7146104eb578063046f7da21461051f57806308a679ad14610535578063157a039b1461055457806315dae03e14610567575b5f80fd5b3480156104f6575f80fd5b5061050a6105053660046152b6565b6112fd565b60405190151581526020015b60405180910390f35b34801561052a575f80fd5b50610533611327565b005b348015610540575f80fd5b5061053361054f3660046152dd565b61135c565b6105336105623660046153b4565b6114b4565b348015610572575f80fd5b507f636f6d6d756e6974792d6f6e636861696e2d76310000000000000000000000005b604051908152602001610516565b3480156105ae575f80fd5b506105336105bd366004615476565b611638565b3480156105cd575f80fd5b506105956105dc3660046154a4565b5f9081525f80516020615f14833981519152602052604090206001015490565b348015610607575f80fd5b5060045461061b906001600160a01b031681565b6040516001600160a01b039091168152602001610516565b34801561063e575f80fd5b506105957f2fc10cc8ae19568712f7a176fb4978616a610650813c9d05326c34abb62749c781565b348015610671575f80fd5b50610533610680366004615476565b6116b1565b348015610690575f80fd5b5061053361069f366004615476565b6116e1565b3480156106af575f80fd5b506105336106be3660046154bb565b611719565b3480156106ce575f80fd5b506105336106dd3660046152dd565b61187e565b3480156106ed575f80fd5b506105957f139c2898040ef16910dc9f44dc697df79363da767d8bc92f2e310312b816e46d81565b348015610720575f80fd5b5061050a61072f3660046154f9565b611985565b34801561073f575f80fd5b5061053361074e366004615519565b6119ab565b34801561075e575f80fd5b506105957fe85fdec10fe0f93d0792364051df7c3d73e37c17b3a954bffe593960e3cd301281565b348015610791575f80fd5b506105336107a0366004615584565b611a3a565b3480156107b0575f80fd5b506105336107bf3660046154f9565b611ab2565b3480156107cf575f80fd5b506105957f000000000000000000000000000000000000000000000000000000000000000a81565b348015610802575f80fd5b506108166108113660046154f9565b611b8b565b60405161051691906155b7565b34801561082e575f80fd5b5061084261083d3660046152dd565b611c73565b60405161051692919061563d565b34801561085b575f80fd5b5061053361086a366004615519565b611ca5565b34801561087a575f80fd5b506105336108893660046154a4565b611cfb565b348015610899575f80fd5b50610533611d5e565b3480156108ad575f80fd5b5061050a6108bc3660046154f9565b611dba565b6105336108cf3660046154a4565b611dca565b3480156108df575f80fd5b50610595611e42565b3480156108f3575f80fd5b506109076109023660046152dd565b611e70565b6040516105169190615661565b34801561091f575f80fd5b50610533611e90565b348015610933575f80fd5b50610533610942366004615673565b611f7e565b348015610952575f80fd5b5061059561096136600461569d565b611fcd565b348015610971575f80fd5b5061050a6109803660046154a4565b600954600160c01b90046001600160401b03161190565b3480156109a2575f80fd5b50610b256109b13660046154a4565b604080516101e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081018290526101c0810191909152505f9081526006602090815260409182902082516101e081018452815463ffffffff8082168352600160201b808304821695840195909552600160401b808304821696840196909652600160601b820481166060840152600160801b820481166080840152600160a01b808304821660a0850152600160c01b8304821660c085015260ff600160e01b909304831660e085015260018501548083166101008601529586049091166101208401526001600160a01b0395909404851661014083015260028301548516610160830152600383015485166101808301526004909201549384166101a08201529190920490911615156101c082015290565b60405161051691906156c3565b348015610b3d575f80fd5b50610533610b4c3660046157f0565b611fea565b348015610b5c575f80fd5b50610533610b6b3660046154a4565b612157565b348015610b7b575f80fd5b50610533610b8a3660046154a4565b612196565b610533610b9d3660046154a4565b6121d5565b348015610bad575f80fd5b50610595610bbc3660046154a4565b612211565b348015610bcc575f80fd5b50610533610bdb366004615476565b612293565b348015610beb575f80fd5b506105955f80516020615ef483398151915281565b348015610c0b575f80fd5b50610533610c1a366004615673565b6122e1565b348015610c2a575f80fd5b50610533610c39366004615519565b612330565b348015610c49575f80fd5b50600954600160c01b90046001600160401b0316610595565b348015610c6d575f80fd5b506105957f59911a6aa08a72fe3824aec4500dc42335c6d0702b6d5c5c72ceb265a0de930281565b348015610ca0575f80fd5b50610533610caf366004615673565b612386565b348015610cbf575f80fd5b50610533610cce3660046152dd565b6123d5565b348015610cde575f80fd5b50610533610ced366004615476565b612583565b348015610cfd575f80fd5b50610533610d0c3660046154a4565b6125d1565b348015610d1c575f80fd5b50610595610d2b3660046154a4565b5f9081526006602052604090205463ffffffff600160201b82048116918116919091031690565b348015610d5d575f80fd5b5061061b610d6c3660046154f9565b6126f9565b348015610d7c575f80fd5b50610533612731565b348015610d90575f80fd5b5061050a610d9f366004615476565b612751565b348015610daf575f80fd5b50610533610dbe3660046158c6565b612787565b348015610dce575f80fd5b5060035461061b906001600160a01b031681565b348015610ded575f80fd5b50600954604080516001600160401b03600160401b8404811682528084166020830152600160801b90930490921690820152606001610516565b348015610e32575f80fd5b50610533610e41366004615954565b612886565b348015610e51575f80fd5b50610533610e60366004615954565b6128e8565b348015610e70575f80fd5b50610533610e7f3660046158c6565b612940565b348015610e8f575f80fd5b506105955f81565b348015610ea2575f80fd5b50610533610eb13660046154f9565b6129fa565b348015610ec1575f80fd5b506105955f1981565b348015610ed5575f80fd5b50610533610ee43660046157f0565b612a2a565b348015610ef4575f80fd5b506105957fb3e25b5404b87e5a838579cb5d7481d61ad96ee284d38ec1e97c07ba64e7f6fc81565b348015610f27575f80fd5b50610533610f363660046154a4565b612b4e565b348015610f46575f80fd5b5061050a612b9d565b348015610f5a575f80fd5b50610f6e610f693660046154a4565b612bcd565b604080519889526020890197909752958701949094526060860192909252608085015260a084015260c083015260e082015261010001610516565b348015610fb4575f80fd5b50610533610fc3366004615954565b612d38565b348015610fd3575f80fd5b50610533610fe23660046154a4565b612e83565b348015610ff2575f80fd5b506105957f000000000000000000000000000000000000000000000000016345785d8a000081565b348015611025575f80fd5b506105336110343660046159ba565b612eba565b348015611044575f80fd5b50610842611053366004615a0a565b6130fb565b348015611063575f80fd5b506105956110723660046154a4565b613413565b348015611082575f80fd5b506105957f79dfcec784e591aafcf60db7db7b029a5c8b12aac4afd4e8c4eb740430405fa681565b3480156110b5575f80fd5b50600554610595565b3480156110c9575f80fd5b50610533613451565b3480156110dd575f80fd5b506105336110ec366004615476565b6134e4565b3480156110fc575f80fd5b506105957f0000000000000000000000000000000000000000000000000de0b6b3a764000081565b34801561112f575f80fd5b506105955f5481565b348015611143575f80fd5b5061061b7f00000000000000000000000028fab2059c713a7f9d8c86db49f9bb0e96af1ef881565b348015611176575f80fd5b5061061b7f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c9503481565b3480156111a9575f80fd5b506105336111b8366004615584565b613514565b3480156111c8575f80fd5b5060045461050a90600160a01b900460ff1681565b3480156111e8575f80fd5b506105957f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea0981565b34801561121b575f80fd5b50610533613553565b34801561122f575f80fd5b5061053361123e3660046152dd565b61356a565b34801561124e575f80fd5b5061053361125d3660046154a4565b61359f565b34801561126d575f80fd5b5061053361127c366004615a65565b6135d2565b34801561128c575f80fd5b506001546112ab906001600160801b0380821691600160801b90041682565b604080516001600160801b03938416815292909116602083015201610516565b3480156112d6575f80fd5b506105336112e53660046154f9565b613849565b6105336112f8366004615aa1565b6139cc565b5f6001600160e01b03198216635a05180f60e01b1480611321575061132182613ade565b92915050565b7f2fc10cc8ae19568712f7a176fb4978616a610650813c9d05326c34abb62749c761135181613b12565b611359613b1c565b50565b5f80516020615ef483398151915261137381613b12565b60ff8311156113955760405163b4fa3fb360e01b815260040160405180910390fd5b63ffffffff8211156113ba5760405163b4fa3fb360e01b815260040160405180910390fd5b6113c384613b71565b5f8481526006602052604090208054600160e01b900460ff16841480156113f757508054600160c01b900463ffffffff1683145b1561140257506114ae565b8054600160e01b900460ff16841461142b57805460ff60e01b1916600160e01b60ff8616021781555b8054600160c01b900463ffffffff16831461145d57805463ffffffff60c01b1916600160c01b63ffffffff8516021781555b604080518581526020810185905286917ff92eb109ce5b449e9b121c352c6aeb4319538a90738cb95d84f08e41274e92d2910160405180910390a26114a4855f6001613ba7565b6114ac613e23565b505b50505050565b6114bc613e63565b5f6114c985838686613e8b565b600354604051636e13f09960e01b8152600481018390529192506001600160a01b031690630f23e742908c908390636e13f099906024015f60405180830381865afa15801561151a573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f191682016040526115419190810190615b89565b6040518363ffffffff1660e01b815260040161155e929190615c5b565b602060405180830381865afa158015611579573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061159d9190615cc3565b34146115bc5760405163162908e360e11b815260040160405180910390fd5b600354604051630b96641560e21b8152336004820152602481018390526001600160a01b0390911690632e5990549034906044015f604051808303818588803b158015611607575f80fd5b505af1158015611619573d5f803e3d5ffd5b505050505061162c818b8b8b8b8b614214565b50505050505050505050565b60405162d74f0b60e71b815260066004820152602481018390526001600160a01b038216604482015273f8e5de8baf8ad7c93dcb61d13d00eb3d57131c7290636ba78580906064015b5f6040518083038186803b158015611697575f80fd5b505af41580156116a9573d5f803e3d5ffd5b505050505050565b5f8281525f80516020615f1483398151915260205260409020600101546116d781613b12565b6114ae8383614372565b6001600160a01b038116331461170a5760405163334bd91960e11b815260040160405180910390fd5b61171482826143be565b505050565b7fe85fdec10fe0f93d0792364051df7c3d73e37c17b3a954bffe593960e3cd301261174381613b12565b5f5b828110156114ae575f84848381811061176057611760615cda565b90506020020135905061177281613b71565b6003546040516325d9153960e11b8152600481018390525f916001600160a01b031690634bb22a72906024016020604051808303815f875af11580156117ba573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117de9190615cc3565b9050801561184b5760035460405163449add1b60e01b8152600481018490526001600160a01b039091169063449add1b906024015f604051808303815f87803b158015611829575f80fd5b505af115801561183b573d5f803e3d5ffd5b5050505061184b8260015f613ba7565b60405182907ef4fe19c0404d2fbb58da6f646c0a3ee5a6994a034213bbd22b072ed1ca5c27905f90a25050600101611745565b7f59911a6aa08a72fe3824aec4500dc42335c6d0702b6d5c5c72ceb265a0de93026118a881613b12565b6118b184613b71565b6003546001600160a01b031663dcab7f83856118ed7f000000000000000000000000000000000000000000000000016345785d8a000086615d02565b6040516001600160e01b031960e085901b168152600481019290925260248201526044015f604051808303815f87803b158015611928575f80fd5b505af115801561193a573d5f803e3d5ffd5b505060408051868152602081018690528793507feec4d6dbe34149c6728a9638eca869d0e5a7fcd85c7a96178f7e9780b4b7fe4b92500160405180910390a26114ae8460015f613ba7565b5f600881608085901b84175b815260208101919091526040015f205460ff169392505050565b6119b48561440a565b600380545f878152600660205260409081902090920154915163cc810cb960e01b81526001600160a01b039182169263cc810cb992611a01928a928a921690899089908990600401615d45565b5f604051808303815f87803b158015611a18575f80fd5b505af1158015611a2a573d5f803e3d5ffd5b505050506114ac85600180613ba7565b611a4383613b71565b600354604051637bcb377f60e11b81526001600160a01b039091169063f7966efe90611a79903390879087908790600401615d85565b5f604051808303815f87803b158015611a90575f80fd5b505af1158015611aa2573d5f803e3d5ffd5b5050505061171483600180613ba7565b7f59911a6aa08a72fe3824aec4500dc42335c6d0702b6d5c5c72ceb265a0de9302611adc81613b12565b611ae583613b71565b60035460405163d963ae5560e01b815260048101859052602481018490526001600160a01b039091169063d963ae55906044015f604051808303815f87803b158015611b2f575f80fd5b505af1158015611b41573d5f803e3d5ffd5b50505050827f1e7ebd3c5f4de9502000b6f7e6e7cf5d4ecb27d6fe1778e43fb9d1d0ca87d0e783604051611b7791815260200190565b60405180910390a261171483600180613ba7565b600954606090600160c01b90046001600160401b03168084101580611bae575082155b15611bc8575050604080515f815260208101909152611321565b5f611bd38583615ded565b8410611be857611be38583615ded565b611bea565b835b9050806001600160401b03811115611c0457611c04615b1d565b604051908082528060200260200182016040528015611c2d578160200160208202803683370190505b5092505f5b8351811015611c6a57611c458187615d02565b848281518110611c5757611c57615cda565b6020908102919091010152600101611c32565b50505092915050565b606080611c81858585614498565b611c8a836144d5565b9092509050611c9d85858585855f61457b565b935093915050565b611cae8561440a565b600380545f87815260066020526040908190209092015491516370903eb960e01b81526001600160a01b03918216926370903eb992611a01928a928a921690899089908990600401615d45565b604051631f46d51760e01b8152600660048201526024810182905273f8e5de8baf8ad7c93dcb61d13d00eb3d57131c7290631f46d517906044015b5f6040518083038186803b158015611d4c575f80fd5b505af41580156114ac573d5f803e3d5ffd5b611d66614609565b73a74528edc289b1a597faf83fcff7eff871cc01d96352d8bfc26040518163ffffffff1660e01b81526004015f6040518083038186803b158015611da8575f80fd5b505af41580156114ae573d5f803e3d5ffd5b5f600781608085901b8417611991565b611dd381613b71565b600354604051630b96641560e21b8152336004820152602481018390526001600160a01b0390911690632e5990549034906044015b5f604051808303818588803b158015611e1f575f80fd5b505af1158015611e31573d5f803e3d5ffd5b505050505061135981600180613ba7565b5f611e6b7fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a025490565b905090565b6060611e7d848484614498565b611e88848484614632565b949350505050565b611e98614609565b604051633d7ad0b760e21b815230600482015273a74528edc289b1a597faf83fcff7eff871cc01d9906389ad9443907f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c95034906001600160a01b0382169063f5eb42dc90602401602060405180830381865afa158015611f18573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611f3c9190615cc3565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044015f6040518083038186803b158015611da8575f80fd5b611f86614609565b604051635c654ad960e01b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d990635c654ad990604401611681565b6001600160801b0381165f90815260026020526040812054611321565b611ff2613e63565b5f611fff86838686613e8b565b600354604051636e13f09960e01b8152600481018390529192505f916001600160a01b0390911690630f23e742908e908390636e13f099906024015f60405180830381865afa158015612054573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261207b9190810190615b89565b6040518363ffffffff1660e01b8152600401612098929190615c5b565b602060405180830381865afa1580156120b3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120d79190615cc3565b60035460405163263f69e960e11b81529192506001600160a01b031690634c7ed3d29061210e903390869086908c90600401615d85565b5f604051808303815f87803b158015612125575f80fd5b505af1158015612137573d5f803e3d5ffd5b50505050612149828d8d8d8d8d614214565b505050505050505050505050565b60405163612b8c3b60e11b8152600660048201526024810182905273f8e5de8baf8ad7c93dcb61d13d00eb3d57131c729063c257187690604401611d36565b60405163c990450f60e01b8152600660048201526024810182905273f8e5de8baf8ad7c93dcb61d13d00eb3d57131c729063c990450f90604401611d36565b6121de81613b71565b6003546040516315b5c47760e01b8152600481018390526001600160a01b03909116906315b5c477903490602401611e08565b6040516351fbfaa560e11b81526001600482015260066024820152604481018290525f90739031730603ea1a523b34d4b04b81ea7a08db0fc49063a3f7f54a90606401602060405180830381865af415801561226f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113219190615cc3565b6040516317a9a2c160e11b815260066004820152602481018390526001600160a01b038216604482015273f8e5de8baf8ad7c93dcb61d13d00eb3d57131c7290632f53458290606401611681565b6122e9614609565b6040516340cea66360e11b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d99063819d4cc690604401611681565b6123398561440a565b600380545f878152600660205260409081902090920154915163f939122360e01b81526001600160a01b039182169263f939122392611a01928a928a921690899089908990600401615d45565b61238e614609565b604051638980f11f60e01b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d990638980f11f90604401611681565b6123de836146cd565b5f8381526006602052604090208054600160401b900463ffffffff1683101561241a57604051635caf530f60e11b815260040160405180910390fd5b80545f906124339086908690869063ffffffff16614740565b90505f835f546124439190615e00565b905080156124d657600354604051632207e80f60e21b815260048101889052602481018390526001600160a01b039091169063881fa03c906044015f604051808303815f87803b158015612495575f80fd5b505af11580156124a7573d5f803e3d5ffd5b50506040518892507f1cbb8dafbedbdf4f813a8ed1f50d871def63e1104f8729b677af57905eda90f691505f90a25b825463ffffffff191663ffffffff831617835560405182815286907fdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f09060200160405180910390a2825463ffffffff60601b1916600160601b63ffffffff84160217835560405182815286907f947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd9060200160405180910390a261257b865f6001613ba7565b6116a9613e23565b604051632a5a705b60e01b815260066004820152602481018390526001600160a01b038216604482015273f8e5de8baf8ad7c93dcb61d13d00eb3d57131c7290632a5a705b90606401611681565b5f80516020615ef48339815191526125e881613b12565b7f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950346001600160a01b0316638fcb4e5b60035f9054906101000a90046001600160a01b03166001600160a01b0316630d43e8ad6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612667573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061268b9190615e17565b6040516001600160e01b031960e084901b1681526001600160a01b039091166004820152602481018590526044016020604051808303815f875af11580156126d5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117149190615cc3565b5f8281527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000602081905260408220611e8890846148ca565b5f80516020615ef483398151915261274881613b12565b6113595f6148d5565b5f9182525f80516020615f14833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b61278f613e63565b612798876146cd565b6003546040516358a46db560e11b815260048101899052602481018890525f916001600160a01b03169063b148db6a90604401602060405180830381865afa1580156127e6573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061280a9190615cc3565b60035460405163263f69e960e11b81529192506001600160a01b031690634c7ed3d2906128419033908c9086908890600401615d85565b5f604051808303815f87803b158015612858575f80fd5b505af115801561286a573d5f803e3d5ffd5b5050505061287c888888888888614214565b5050505050505050565b5f80516020615ef483398151915261289d81613b12565b5f6128aa86868686614910565b90505f5b818110156128df576008810287013560c01c6010820286013560801c6128d582825f614983565b50506001016128ae565b506116a9613e23565b5f80516020615ef48339815191526128ff81613b12565b5f61290c86868686614910565b90505f5b818110156128df576008810287013560c01c6010820286013560801c6129368282614aa4565b5050600101612910565b612948613e63565b612951876146cd565b600354604051632884698160e01b815260048101899052602481018890525f916001600160a01b031690632884698190604401602060405180830381865afa15801561299f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129c39190615cc3565b600354604051637bcb377f60e11b81529192506001600160a01b03169063f7966efe906128419033908c9086908890600401615d85565b5f80516020615ef4833981519152612a1181613b12565b604051630280e1e560e61b815260040160405180910390fd5b612a32613e63565b5f612a3f86838686613e8b565b600354604051636e13f09960e01b8152600481018390529192505f916001600160a01b0390911690639a4df8f0908e908390636e13f099906024015f60405180830381865afa158015612a94573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052612abb9190810190615b89565b6040518363ffffffff1660e01b8152600401612ad8929190615c5b565b602060405180830381865afa158015612af3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612b179190615cc3565b600354604051637bcb377f60e11b81529192506001600160a01b03169063f7966efe9061210e903390869086908c90600401615d85565b612b57816146cd565b604051633f58c75d60e21b8152600160048201526006602482015260448101829052739031730603ea1a523b34d4b04b81ea7a08db0fc49063fd631d7490606401611d36565b5f612bc67fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a025490565b4210905090565b5f818152600660205260408082206003549151634e28b08160e11b815260048101859052839283928392839283928392839283916001600160a01b0390911690639c51610290602401602060405180830381865afa158015612c31573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612c559190615cc3565b90505f81118015612c7157508154600160e01b900460ff166002145b15612cac57815460029a50612ca59063ffffffff600160c01b8204811691600160201b810482169082160316839003614c17565b9850612cf6565b8015612cd857815460029a5063ffffffff600160201b8204811691811691909103168190039850612cf6565b8154600160e01b810460ff169a50600160c01b900463ffffffff1698505b508054600190910154989a97995063ffffffff600160801b82048116995f998a99509082169750600160401b830482169650600160a01b909204169350915050565b5f80516020615ef4833981519152612d4f81613b12565b5f612d5c86868686614910565b90505f5b818110156128df576008810287013560c01c6010820286013560801c612d8582613b71565b5f8281526006602052604090208054600160601b900463ffffffff168210612dc0576040516388e1a28160e01b815260040160405180910390fd5b8054600160401b900463ffffffff16821015612def576040516388e1a28160e01b815260040160405180910390fd5b805463ffffffff60601b1916600160601b63ffffffff84160217815560405182815283907f947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd9060200160405180910390a260405183907fe5725d045d5c47bd1483feba445e395dc8647486963e6d54aad9ed03ff7d6ce6905f90a2612e75835f80613ba7565b505050806001019050612d60565b7f79dfcec784e591aafcf60db7db7b029a5c8b12aac4afd4e8c4eb740430405fa6612ead81613b12565b612eb6826148d5565b5050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f81158015612efe5750825b90505f826001600160401b03166001148015612f195750303b155b905081158015612f27575080155b15612f455760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315612f6f57845460ff60401b1916600160401b1785555b6001600160a01b038916612f96576040516368ea2bc160e01b815260040160405180910390fd5b6001600160a01b038616612fbd57604051633ef39b8160e01b815260040160405180910390fd5b612fc5614c2c565b600380546001600160a01b03808c166001600160a01b03199283161790925560048054928b1692909116919091179055612fff5f87614372565b506130965f80516020615ef48339815191527f00000000000000000000000028fab2059c713a7f9d8c86db49f9bb0e96af1ef86001600160a01b031663ef6c064c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561306d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130919190615e17565b614372565b506130a0876148d5565b6130aa5f19614c34565b83156130f057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050565b6060805f80516020615ef483398151915261311581613b12565b61311e866144d5565b9093509150851561340a576001546001600160801b03165f908152600260205260408120548791905b8015613385575f6131588260c01c90565b6001600160401b039081165f8181526006602052604081208054929450608086901c90931692916131a09061319a90600160a01b900463ffffffff1685614c17565b88614c17565b9050808711806131af57508281145b156131f0576001808301805463ffffffff600160201b80830482168890039091160267ffffffff00000000199091161790556131ea90614c83565b5061324e565b60018201805463ffffffff600160201b808304821685900382160267ffffffff00000000199092169190911790915561322f90869083860390614ce216565b6001546001600160801b03165f90815260026020526040902081905594505b805f0361325e5750505050613365565b815461327c908590600160401b900463ffffffff16838d8d8b61457b565b815463ffffffff600160401b80830482168401821681026bffffffff000000000000000019909316929092178085556040519290041681529581019584907f24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b789060200160405180910390a2815463ffffffff600160a01b808304821684900391821690810263ffffffff60a01b199093169290921784556040519182529085907ff9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed339060200160405180910390a28188039750875f0361335f575050505050613385565b50505050505b506001546001600160801b03165f90815260026020526040902054613147565b508781146133a657604051630bc9ea5560e21b815260040160405180910390fd5b600980546001600160401b03600160801b80830482168c9003821602808216828416178c0190911667ffffffffffffffff1990911677ffffffffffffffff0000000000000000ffffffffffffffff1990921691909117179055613407613e23565b50505b50935093915050565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200060208190526040822061344a90614d07565b9392505050565b7f79dfcec784e591aafcf60db7db7b029a5c8b12aac4afd4e8c4eb740430405fa661347b81613b12565b600454600160a01b900460ff16156134a65760405163ef65161f60e01b815260040160405180910390fd5b6004805460ff60a01b1916600160a01b1790556040517fe5eb57aa4d841adeece4ac87bd294965df4a894f0aa24db4a4a55a39ab101d6e905f90a150565b5f8281525f80516020615f14833981519152602052604090206001015461350a81613b12565b6114ae83836143be565b61351d83613b71565b60035460405163263f69e960e11b81526001600160a01b0390911690634c7ed3d290611a79903390879087908790600401615d85565b5f80516020615ef483398151915261135981613b12565b5f80516020615ef483398151915261358181613b12565b61358d84846001614983565b6135978483614aa4565b6114ae613e23565b7f139c2898040ef16910dc9f44dc697df79363da767d8bc92f2e310312b816e46d6135c981613b12565b612eb682614c34565b7f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea096135fc81613b12565b61360585613b71565b5f8581526006602052604090208054600160401b900463ffffffff16851061364057604051635caf530f60e11b815260040160405180910390fd5b608086901b85175f8181526007602052604090205460ff161561367657604051639fbfc58960e01b815260040160405180910390fd5b5f8181526007602052604090819020805460ff19166001908117909155835463ffffffff600160201b80830482169093011690910267ffffffff00000000199091161783555187907fcb2f99f65711a7d6df7f552255b910bf59f09fcd5935f44c170b4cb0d1b50995906136f69089908990918252602082015260400190565b60405180910390a283156137b7575f8181526008602052604090205460ff1615613744577f0000000000000000000000000000000000000000000000000de0b6b3a76400008501945061375d565b5f818152600860205260409020805460ff191660011790555b60035460405163449add1b60e01b8152600481018990526001600160a01b039091169063449add1b906024015f604051808303815f87803b1580156137a0575f80fd5b505af11580156137b2573d5f803e3d5ffd5b505050505b846801bc16d674ec80000011156138345760035460405163e5220e3f60e01b8152600481018990526801bc16d674ec80000087900360248201526001600160a01b039091169063e5220e3f906044015f604051808303815f87803b15801561381d575f80fd5b505af115801561382f573d5f803e3d5ffd5b505050505b61384087600180613ba7565b50505050505050565b7f0ce23c3e399818cfee81a7ab0880f714e53d7672b08df0fa62f2843416e1ea0961387381613b12565b61387c83613b71565b5f8381526006602052604090208054600160401b900463ffffffff1683106138b757604051635caf530f60e11b815260040160405180910390fd5b608084901b83175f8181526008602052604090205460ff16156138ed57604051639fbfc58960e01b815260040160405180910390fd5b5f8181526008602052604090819020805460ff191660011790555185907fd34db8e8c0ddbc9c7b6dd8c397623dfbe01929e41e527540bff8794685d9b407906139399087815260200190565b60405180910390a260035460405163e5220e3f60e01b8152600481018790527f0000000000000000000000000000000000000000000000000de0b6b3a764000060248201526001600160a01b039091169063e5220e3f906044015f604051808303815f87803b1580156139aa575f80fd5b505af11580156139bc573d5f803e3d5ffd5b505050506114ac8560015f613ba7565b6139d4613e63565b6139dd866146cd565b6003546040516358a46db560e11b815260048101889052602481018790526001600160a01b039091169063b148db6a90604401602060405180830381865afa158015613a2b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613a4f9190615cc3565b3414613a6e5760405163162908e360e11b815260040160405180910390fd5b600354604051630b96641560e21b8152336004820152602481018890526001600160a01b0390911690632e5990549034906044015f604051808303818588803b158015613ab9575f80fd5b505af1158015613acb573d5f803e3d5ffd5b50505050506116a9868686868686614214565b5f6001600160e01b03198216637965db0b60e01b148061132157506301ffc9a760e01b6001600160e01b0319831614611321565b6113598133614d10565b613b24614d4d565b427fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a02556040517f62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9905f90a1565b600954600160c01b90046001600160401b0316811015613b8e5750565b604051633ed893db60e21b815260040160405180910390fd5b5f8381526006602052604081208054909190613bd99063ffffffff600160401b8204811691600160601b900416615e32565b6003546040516301a5e9e360e01b81526004810188905263ffffffff9290921692505f916001600160a01b03909116906301a5e9e390602401602060405180830381865afa158015613c2d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c519190615cc3565b905081811115613c63575f9150613c69565b80820391505b8254600160801b900463ffffffff1615801590613c8557505f82115b15613c8e575f91505b8254600160e01b900460ff1615801590613ca757505f82115b15613cff57825463ffffffff600160201b82048116600160401b8304821603811691613cfb91600160c01b909104168210613ce2575f613cf5565b8454600160c01b900463ffffffff168290035b84614c17565b9250505b8254600160a01b900463ffffffff1682146116a95782546009805467ffffffffffffffff60801b198116600160a01b9384900463ffffffff908116600160801b938490046001600160401b039081169190910388011690920217909155845463ffffffff60a01b191690841690910217835560405182815286907ff9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed339060200160405180910390a28415613db457613db4613e23565b83156116a957604051633f58c75d60e21b8152600160048201526006602482015260448101879052739031730603ea1a523b34d4b04b81ea7a08db0fc49063fd631d74906064015f6040518083038186803b158015613e11575f80fd5b505af415801561162c573d5f803e3d5ffd5b60058054600101908190556040519081527f7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa9060200160405180910390a1565b613e6b612b9d565b15613e8957604051630286f07360e31b815260040160405180910390fd5b565b6004545f90600160a01b900460ff16613ed257811580613eb457506004546001600160a01b0316155b15613ed25760405163084a55b960e41b815260040160405180910390fd5b50600954600160c01b90046001600160401b03165f81815260066020908152604082209190613f0390880188615e4f565b6001600160a01b031614613f2357613f1e6020870187615e4f565b613f25565b335b6001820180546001600160a01b0392909216600160401b027fffffffff0000000000000000000000000000000000000000ffffffffffffffff9092169190911790555f613f786040880160208901615e4f565b6001600160a01b031614613f9b57613f966040870160208801615e4f565b613f9d565b335b6003820180546001600160a01b0319166001600160a01b0392909216919091179055613fcf6060870160408801615e6a565b1561400357613fe46060870160408801615e6a565b600482018054911515600160a01b0260ff60a01b199092169190911790555b6009805460016001600160401b03600160c01b80840482168301909116026001600160c01b03909216919091179091556003820154908201546040516001600160a01b0392831692600160401b9092049091169084907ff35982c84fdc94f58d48e901c54c615804cf7d7939b9b8f76ce4d459354e6363905f90a46001600160a01b038516156140c3576040516001600160a01b0386169083907f67334334c388385e5f244703f8a8b28b7f4ffe52909130aca69bc62a8e27f09a905f90a35b82158015906140dc57506004546001600160a01b031615155b1561420b576004805460405163076123b360e21b81526001600160a01b0390911691631d848ecc916141149133918991899101615e83565b5f604051808303815f87803b15801561412b575f80fd5b505af115801561413d573d5f803e3d5ffd5b5050600354600480546040805163464b6c0d60e11b815290516001600160a01b03948516965063b2d03e4d9550889490921692638c96d81a9282820192602092908290030181865afa158015614195573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906141b99190615cc3565b6040516001600160e01b031960e085901b168152600481019290925260248201526044015f604051808303815f87803b1580156141f4575f80fd5b505af1158015614206573d5f803e3d5ffd5b505050505b50949350505050565b5f868152600660205260409020805460045463ffffffff90911690600160a01b900460ff1615801561426757507f000000000000000000000000000000000000000000000000000000000000000a878201115b15614285576040516347f1bdb360e11b815260040160405180910390fd5b61429488828989898989614d72565b50815463ffffffff600160601b8204811691160361430f57815463ffffffff600160601b80830482168a018216810263ffffffff60601b199093169290921780855560405192900416815288907f947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd9060200160405180910390a25b815463ffffffff80821689011663ffffffff199091168117835560405190815288907fdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f09060200160405180910390a261436a885f6001613ba7565b61287c613e23565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e823717059320008161439f8585614f06565b90508015611e88575f85815260208390526040902061420b9085614fae565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000816143eb8585614fc2565b90508015611e88575f85815260208390526040902061420b908561503b565b5f8181526006602052604090206001810154600160401b90046001600160a01b031661444957604051633ed893db60e21b815260040160405180910390fd5b6001810154600160401b90046001600160a01b0316331480159061447a575060038101546001600160a01b03163314155b15612eb65760405163743a3f7960e11b815260040160405180910390fd5b5f8381526006602052604090205463ffffffff166144b68284615d02565b111561171457604051635caf530f60e11b815260040160405180910390fd5b6060806144e3603084615e00565b6001600160401b038111156144fa576144fa615b1d565b6040519080825280601f01601f191660200182016040528015614524576020820181803683370190505b50614530606085615e00565b6001600160401b0381111561454757614547615b1d565b6040519080825280601f01601f191660200182016040528015614571576020820181803683370190505b5091509150915091565b5f805b8581101561287c576145bc88614594838a615d02565b7f059e9c54cf92ba46cc39c6b4acd51d5116e9d49fabee6193530ea918b54be94a919061504f565b60018082015460801c85840160308181028a01908101929092528354602092830152600284015460609182028901928301526003840154604083015260048401549101529092500161457e565b613e897fb3e25b5404b87e5a838579cb5d7481d61ad96ee284d38ec1e97c07ba64e7f6fc613b12565b60605f614640603084615e00565b6001600160401b0381111561465757614657615b1d565b6040519080825280601f01601f191660200182016040528015614681576020820181803683370190505b5091505f5b838110156146c45761469c866145948388615d02565b9150603081026020840101600183015460801c60108201528254815250600181019050614686565b50509392505050565b5f8181526006602052604090206001810154600160401b90046001600160a01b031661470c57604051633ed893db60e21b815260040160405180910390fd5b6001810154600160401b90046001600160a01b03163314612eb65760405163743a3f7960e11b815260040160405180910390fd5b5f8215806147565750816147548486615d02565b115b80614764575063ffffffff82115b156147825760405163575697ff60e01b815260040160405180910390fd5b604080516030808252606082019092525f91829182918291906020820181803683370190505090508787015b888111156148bc576147e47f059e9c54cf92ba46cc39c6b4acd51d5116e9d49fabee6193530ea918b54be94a8b5f19840161504f565b9450600185015460801c60308301528454602083015286811015614856576148307f059e9c54cf92ba46cc39c6b4acd51d5116e9d49fabee6193530ea918b54be94a8b5f198a0161504f565b93505f92505b6005831015614852578284015483860155600183019250614836565b8394505b5f92505b6005831015614873575f8386015560018301925061485a565b600187039650600181039050897fea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9836040516148af9190615661565b60405180910390a26147ae565b509498975050505050505050565b5f61344a8383615086565b5f8190556040518181527f699ec9c671aad1f3dcc15e571375584a1d6fb7176afd545d14467fd31477e98e906020015b60405180910390a150565b5f61491c600885615eb9565b614927601084615eb9565b14158061493d575061493a600885615ecc565b15155b80614951575061494e601083615ecc565b15155b1561496f5760405163319c9a2160e21b815260040160405180910390fd5b61497a600885615eb9565b95945050505050565b61498c83613b71565b5f838152600660205260409020600181015463ffffffff1683036149b05750505050565b8054600160401b900463ffffffff168311156149df5760405163cc11217f60e01b815260040160405180910390fd5b811580156149f65750600181015463ffffffff1683105b15614a14576040516371a4bd1560e01b815260040160405180910390fd5b6001810180546009805463ffffffff9283166001600160401b03600160401b808404821692909203890116026fffffffffffffffff000000000000000019909116179055815490851663ffffffff1990911617905560405183815284907f0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c2201669060200160405180910390a250505050565b614aad82613b71565b5f8281526006602052604090208054600160801b900463ffffffff168203614ad457505050565b6001810154815463ffffffff918216600160401b90910482160316821115614b0f57604051636af5e8d960e11b815260040160405180910390fd5b805463ffffffff60801b1916600160801b63ffffffff84160217815560405182815283907fb4f5879eca27b32881cec7907d1310378e9b4c79927062fb7d4a321434b5b04a9060200160405180910390a25f82118015614b7c57508054600160a01b900463ffffffff1615155b15614c0c5780546009805467ffffffffffffffff60801b198116600160a01b90930463ffffffff16600160801b918290046001600160401b03908116919091031602919091179055805463ffffffff60a01b191681556040515f815283907ff9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed339060200160405180910390a2505050565b611714835f80613ba7565b5f818310614c25578161344a565b5090919050565b613e896150ac565b614c3c613e63565b805f03614c5c5760405163ad58bfc760e01b815260040160405180910390fd5b5f5f198203614c6d57505f19614c7a565b614c778242615d02565b90505b612eb6816150f5565b80546001600160801b03165f90815260018201602052604090205480614cbc576040516363c3654960e01b815260040160405180910390fd5b81546fffffffffffffffffffffffffffffffff19166001600160801b0382161790915590565b60801b67ffffffffffffffff60801b1667ffffffffffffffff60801b19919091161790565b5f611321825490565b614d1a8282612751565b612eb65760405163e2517d3f60e01b81526001600160a01b03821660048201526024810183905260440160405180910390fd5b614d55612b9d565b613e895760405163b047186b60e01b815260040160405180910390fd5b5f851580614d8c575063ffffffff614d8a8789615d02565b115b15614daa5760405163575697ff60e01b815260040160405180910390fd5b6030860284141580614dbf5750606086028214155b15614ddd5760405163251f56a160e21b815260040160405180910390fd5b604080516030808252606082019092525f91829182916020820181803683370190505090505f5b89811015614ef657614e377f059e9c54cf92ba46cc39c6b4acd51d5116e9d49fabee6193530ea918b54be94a8d8d61504f565b60308281028b0160108101359185018290523560208501819052919550171592508215614e7757604051630f35a7eb60e21b815260040160405180910390fd5b60208201518455603082015160801b60018501556060810287018035600286015560208101356003860155604081013560048601555060018101905060018b019a508b7fc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a8983604051614ee99190615661565b60405180910390a2614e04565b50989a9950505050505050505050565b5f5f80516020615f14833981519152614f1f8484612751565b614f9e575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055614f543390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050611321565b5f915050611321565b5092915050565b5f61344a836001600160a01b038416615190565b5f5f80516020615f14833981519152614fdb8484612751565b15614f9e575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050611321565b5f61344a836001600160a01b0384166151dc565b6040805160208082019590955280820193909352606080840192909252805180840390920182526080909201909152805191012090565b5f825f01828154811061509b5761509b615cda565b905f5260205f200154905092915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16613e8957604051631afcd79f60e31b815260040160405180910390fd5b61511e7fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a02829055565b5f198103615157576040515f1981527f32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e90602001614905565b7f32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e6151824283615ded565b604051908152602001614905565b5f8181526001830160205260408120546151d557508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155611321565b505f611321565b5f8181526001830160205260408120548015614f9e575f6151fe600183615ded565b85549091505f9061521190600190615ded565b9050808214615270575f865f01828154811061522f5761522f615cda565b905f5260205f200154905080875f01848154811061524f5761524f615cda565b5f918252602080832090910192909255918252600188019052604090208390555b855486908061528157615281615edf565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050611321565b5f602082840312156152c6575f80fd5b81356001600160e01b03198116811461344a575f80fd5b5f805f606084860312156152ef575f80fd5b505081359360208301359350604090920135919050565b5f8083601f840112615316575f80fd5b5081356001600160401b0381111561532c575f80fd5b602083019150836020828501011115615343575f80fd5b9250929050565b5f6060828403121561535a575f80fd5b50919050565b5f8083601f840112615370575f80fd5b5081356001600160401b03811115615386575f80fd5b6020830191508360208260051b8501011115615343575f80fd5b6001600160a01b0381168114611359575f80fd5b5f805f805f805f805f6101008a8c0312156153cd575f80fd5b8935985060208a01356001600160401b03808211156153ea575f80fd5b6153f68d838e01615306565b909a50985060408c013591508082111561540e575f80fd5b61541a8d838e01615306565b909850965086915061542f8d60608e0161534a565b955060c08c0135915080821115615444575f80fd5b506154518c828d01615360565b90945092505060e08a0135615465816153a0565b809150509295985092959850929598565b5f8060408385031215615487575f80fd5b823591506020830135615499816153a0565b809150509250929050565b5f602082840312156154b4575f80fd5b5035919050565b5f80602083850312156154cc575f80fd5b82356001600160401b038111156154e1575f80fd5b6154ed85828601615360565b90969095509350505050565b5f806040838503121561550a575f80fd5b50508035926020909101359150565b5f805f805f6080868803121561552d575f80fd5b85359450602086013593506040860135925060608601356001600160401b03811115615557575f80fd5b61556388828901615360565b969995985093965092949392505050565b5f60a0828403121561535a575f80fd5b5f805f60e08486031215615596575f80fd5b83359250602084013591506155ae8560408601615574565b90509250925092565b602080825282518282018190525f9190848201906040850190845b818110156155ee578351835292840192918401916001016155d2565b50909695505050505050565b5f81518084525f5b8181101561561e57602081850181015186830182015201615602565b505f602082860101526020601f19601f83011685010191505092915050565b604081525f61564f60408301856155fa565b828103602084015261497a81856155fa565b602081525f61344a60208301846155fa565b5f8060408385031215615684575f80fd5b823561568f816153a0565b946020939093013593505050565b5f602082840312156156ad575f80fd5b81356001600160801b038116811461344a575f80fd5b815163ffffffff1681526101e0810160208301516156e9602084018263ffffffff169052565b506040830151615701604084018263ffffffff169052565b506060830151615719606084018263ffffffff169052565b506080830151615731608084018263ffffffff169052565b5060a083015161574960a084018263ffffffff169052565b5060c083015161576160c084018263ffffffff169052565b5060e083015161577660e084018260ff169052565b506101008381015163ffffffff908116918401919091526101208085015190911690830152610140808401516001600160a01b039081169184019190915261016080850151821690840152610180808501518216908401526101a080850151909116908301526101c0928301511515929091019190915290565b5f805f805f805f805f806101a08b8d03121561580a575f80fd5b8a35995060208b01356001600160401b0380821115615827575f80fd5b6158338e838f01615306565b909b50995060408d013591508082111561584b575f80fd5b6158578e838f01615306565b909950975087915061586c8e60608f0161534a565b965061587b8e60c08f01615574565b95506101608d0135915080821115615891575f80fd5b5061589e8d828e01615360565b9094509250506101808b01356158b3816153a0565b809150509295989b9194979a5092959850565b5f805f805f805f610120888a0312156158dd575f80fd5b873596506020880135955060408801356001600160401b0380821115615901575f80fd5b61590d8b838c01615306565b909750955060608a0135915080821115615925575f80fd5b506159328a828b01615306565b909450925061594690508960808a01615574565b905092959891949750929550565b5f805f8060408587031215615967575f80fd5b84356001600160401b038082111561597d575f80fd5b61598988838901615306565b909650945060208701359150808211156159a1575f80fd5b506159ae87828801615306565b95989497509550505050565b5f805f80608085870312156159cd575f80fd5b84356159d8816153a0565b935060208501356159e8816153a0565b92506040850135915060608501356159ff816153a0565b939692955090935050565b5f805f60408486031215615a1c575f80fd5b8335925060208401356001600160401b03811115615a38575f80fd5b615a4486828701615306565b9497909650939450505050565b80358015158114615a60575f80fd5b919050565b5f805f8060808587031215615a78575f80fd5b843593506020850135925060408501359150615a9660608601615a51565b905092959194509250565b5f805f805f8060808789031215615ab6575f80fd5b863595506020870135945060408701356001600160401b0380821115615ada575f80fd5b615ae68a838b01615306565b90965094506060890135915080821115615afe575f80fd5b50615b0b89828a01615306565b979a9699509497509295939492505050565b634e487b7160e01b5f52604160045260245ffd5b604080519081016001600160401b0381118282101715615b5357615b53615b1d565b60405290565b604051601f8201601f191681016001600160401b0381118282101715615b8157615b81615b1d565b604052919050565b5f6020808385031215615b9a575f80fd5b82516001600160401b0380821115615bb0575f80fd5b9084019060408287031215615bc3575f80fd5b615bcb615b31565b825182811115615bd9575f80fd5b8301601f81018813615be9575f80fd5b805183811115615bfb57615bfb615b1d565b8060051b9350615c0c868501615b59565b818152938201860193868101908a861115615c25575f80fd5b928701925b85841015615c4357835182529287019290870190615c2a565b84525050509183015192820192909252949350505050565b8281525f60206040602084015260808301845160408086015281815180845260a0870191506020830193505f92505b80831015615caa5783518252928401926001929092019190840190615c8a565b5060208701516060870152809450505050509392505050565b5f60208284031215615cd3575f80fd5b5051919050565b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601160045260245ffd5b8082018082111561132157611321615cee565b8183525f6001600160fb1b03831115615d2c575f80fd5b8260051b80836020870137939093016020019392505050565b8681528560208201526001600160a01b038516604082015283606082015260a060808201525f615d7960a083018486615d15565b98975050505050505050565b5f610100820190506001600160a01b03861682528460208301528360408301528235606083015260208301356080830152604083013560ff8116808214615dca575f80fd5b60a084015250606083013560c083015260809092013560e0909101529392505050565b8181038181111561132157611321615cee565b808202811582820484141761132157611321615cee565b5f60208284031215615e27575f80fd5b815161344a816153a0565b63ffffffff828116828216039080821115614fa757614fa7615cee565b5f60208284031215615e5f575f80fd5b813561344a816153a0565b5f60208284031215615e7a575f80fd5b61344a82615a51565b6001600160a01b0384168152604060208201525f61497a604083018486615d15565b634e487b7160e01b5f52601260045260245ffd5b5f82615ec757615ec7615ea5565b500490565b5f82615eda57615eda615ea5565b500690565b634e487b7160e01b5f52603160045260245ffdfebb75b874360e0bfd87f964eadd8276d8efb7c942134fc329b513032d0803e0c602dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a164736f6c6343000818000a",
}

// CSModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use CSModuleMetaData.ABI instead.
var CSModuleABI = CSModuleMetaData.ABI

// CSModuleBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CSModuleMetaData.Bin instead.
var CSModuleBin = CSModuleMetaData.Bin

// DeployCSModule deploys a new Ethereum contract, binding an instance of CSModule to it.
func DeployCSModule(auth *bind.TransactOpts, backend bind.ContractBackend, moduleType [32]byte, minSlashingPenaltyQuotient *big.Int, elRewardsStealingFine *big.Int, maxKeysPerOperatorEA *big.Int, lidoLocator common.Address) (common.Address, *types.Transaction, *CSModule, error) {
	parsed, err := CSModuleMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CSModuleBin), backend, moduleType, minSlashingPenaltyQuotient, elRewardsStealingFine, maxKeysPerOperatorEA, lidoLocator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CSModule{CSModuleCaller: CSModuleCaller{contract: contract}, CSModuleTransactor: CSModuleTransactor{contract: contract}, CSModuleFilterer: CSModuleFilterer{contract: contract}}, nil
}

// CSModule is an auto generated Go binding around an Ethereum contract.
type CSModule struct {
	CSModuleCaller     // Read-only binding to the contract
	CSModuleTransactor // Write-only binding to the contract
	CSModuleFilterer   // Log filterer for contract events
}

// CSModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CSModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CSModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CSModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CSModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CSModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CSModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CSModuleSession struct {
	Contract     *CSModule         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CSModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CSModuleCallerSession struct {
	Contract *CSModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CSModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CSModuleTransactorSession struct {
	Contract     *CSModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CSModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CSModuleRaw struct {
	Contract *CSModule // Generic contract binding to access the raw methods on
}

// CSModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CSModuleCallerRaw struct {
	Contract *CSModuleCaller // Generic read-only contract binding to access the raw methods on
}

// CSModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CSModuleTransactorRaw struct {
	Contract *CSModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCSModule creates a new instance of CSModule, bound to a specific deployed contract.
func NewCSModule(address common.Address, backend bind.ContractBackend) (*CSModule, error) {
	contract, err := bindCSModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CSModule{CSModuleCaller: CSModuleCaller{contract: contract}, CSModuleTransactor: CSModuleTransactor{contract: contract}, CSModuleFilterer: CSModuleFilterer{contract: contract}}, nil
}

// NewCSModuleCaller creates a new read-only instance of CSModule, bound to a specific deployed contract.
func NewCSModuleCaller(address common.Address, caller bind.ContractCaller) (*CSModuleCaller, error) {
	contract, err := bindCSModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CSModuleCaller{contract: contract}, nil
}

// NewCSModuleTransactor creates a new write-only instance of CSModule, bound to a specific deployed contract.
func NewCSModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*CSModuleTransactor, error) {
	contract, err := bindCSModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CSModuleTransactor{contract: contract}, nil
}

// NewCSModuleFilterer creates a new log filterer instance of CSModule, bound to a specific deployed contract.
func NewCSModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*CSModuleFilterer, error) {
	contract, err := bindCSModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CSModuleFilterer{contract: contract}, nil
}

// bindCSModule binds a generic wrapper to an already deployed contract.
func bindCSModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CSModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CSModule *CSModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CSModule.Contract.CSModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CSModule *CSModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSModule.Contract.CSModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CSModule *CSModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CSModule.Contract.CSModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CSModule *CSModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CSModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CSModule *CSModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CSModule *CSModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CSModule.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CSModule *CSModuleSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CSModule.Contract.DEFAULTADMINROLE(&_CSModule.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CSModule.Contract.DEFAULTADMINROLE(&_CSModule.CallOpts)
}

// ELREWARDSSTEALINGFINE is a free data retrieval call binding the contract method 0xbdac46a2.
//
// Solidity: function EL_REWARDS_STEALING_FINE() view returns(uint256)
func (_CSModule *CSModuleCaller) ELREWARDSSTEALINGFINE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "EL_REWARDS_STEALING_FINE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ELREWARDSSTEALINGFINE is a free data retrieval call binding the contract method 0xbdac46a2.
//
// Solidity: function EL_REWARDS_STEALING_FINE() view returns(uint256)
func (_CSModule *CSModuleSession) ELREWARDSSTEALINGFINE() (*big.Int, error) {
	return _CSModule.Contract.ELREWARDSSTEALINGFINE(&_CSModule.CallOpts)
}

// ELREWARDSSTEALINGFINE is a free data retrieval call binding the contract method 0xbdac46a2.
//
// Solidity: function EL_REWARDS_STEALING_FINE() view returns(uint256)
func (_CSModule *CSModuleCallerSession) ELREWARDSSTEALINGFINE() (*big.Int, error) {
	return _CSModule.Contract.ELREWARDSSTEALINGFINE(&_CSModule.CallOpts)
}

// INITIALSLASHINGPENALTY is a free data retrieval call binding the contract method 0xd6477919.
//
// Solidity: function INITIAL_SLASHING_PENALTY() view returns(uint256)
func (_CSModule *CSModuleCaller) INITIALSLASHINGPENALTY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "INITIAL_SLASHING_PENALTY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITIALSLASHINGPENALTY is a free data retrieval call binding the contract method 0xd6477919.
//
// Solidity: function INITIAL_SLASHING_PENALTY() view returns(uint256)
func (_CSModule *CSModuleSession) INITIALSLASHINGPENALTY() (*big.Int, error) {
	return _CSModule.Contract.INITIALSLASHINGPENALTY(&_CSModule.CallOpts)
}

// INITIALSLASHINGPENALTY is a free data retrieval call binding the contract method 0xd6477919.
//
// Solidity: function INITIAL_SLASHING_PENALTY() view returns(uint256)
func (_CSModule *CSModuleCallerSession) INITIALSLASHINGPENALTY() (*big.Int, error) {
	return _CSModule.Contract.INITIALSLASHINGPENALTY(&_CSModule.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_CSModule *CSModuleCaller) LIDOLOCATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "LIDO_LOCATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_CSModule *CSModuleSession) LIDOLOCATOR() (common.Address, error) {
	return _CSModule.Contract.LIDOLOCATOR(&_CSModule.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_CSModule *CSModuleCallerSession) LIDOLOCATOR() (common.Address, error) {
	return _CSModule.Contract.LIDOLOCATOR(&_CSModule.CallOpts)
}

// MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE is a free data retrieval call binding the contract method 0x47faf311.
//
// Solidity: function MAX_SIGNING_KEYS_PER_OPERATOR_BEFORE_PUBLIC_RELEASE() view returns(uint256)
func (_CSModule *CSModuleCaller) MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "MAX_SIGNING_KEYS_PER_OPERATOR_BEFORE_PUBLIC_RELEASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE is a free data retrieval call binding the contract method 0x47faf311.
//
// Solidity: function MAX_SIGNING_KEYS_PER_OPERATOR_BEFORE_PUBLIC_RELEASE() view returns(uint256)
func (_CSModule *CSModuleSession) MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE() (*big.Int, error) {
	return _CSModule.Contract.MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE(&_CSModule.CallOpts)
}

// MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE is a free data retrieval call binding the contract method 0x47faf311.
//
// Solidity: function MAX_SIGNING_KEYS_PER_OPERATOR_BEFORE_PUBLIC_RELEASE() view returns(uint256)
func (_CSModule *CSModuleCallerSession) MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE() (*big.Int, error) {
	return _CSModule.Contract.MAXSIGNINGKEYSPEROPERATORBEFOREPUBLICRELEASE(&_CSModule.CallOpts)
}

// MODULEMANAGERROLE is a free data retrieval call binding the contract method 0xcb17ed3e.
//
// Solidity: function MODULE_MANAGER_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCaller) MODULEMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "MODULE_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MODULEMANAGERROLE is a free data retrieval call binding the contract method 0xcb17ed3e.
//
// Solidity: function MODULE_MANAGER_ROLE() view returns(bytes32)
func (_CSModule *CSModuleSession) MODULEMANAGERROLE() ([32]byte, error) {
	return _CSModule.Contract.MODULEMANAGERROLE(&_CSModule.CallOpts)
}

// MODULEMANAGERROLE is a free data retrieval call binding the contract method 0xcb17ed3e.
//
// Solidity: function MODULE_MANAGER_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCallerSession) MODULEMANAGERROLE() ([32]byte, error) {
	return _CSModule.Contract.MODULEMANAGERROLE(&_CSModule.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_CSModule *CSModuleCaller) PAUSEINFINITELY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "PAUSE_INFINITELY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_CSModule *CSModuleSession) PAUSEINFINITELY() (*big.Int, error) {
	return _CSModule.Contract.PAUSEINFINITELY(&_CSModule.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_CSModule *CSModuleCallerSession) PAUSEINFINITELY() (*big.Int, error) {
	return _CSModule.Contract.PAUSEINFINITELY(&_CSModule.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_CSModule *CSModuleSession) PAUSEROLE() ([32]byte, error) {
	return _CSModule.Contract.PAUSEROLE(&_CSModule.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCallerSession) PAUSEROLE() ([32]byte, error) {
	return _CSModule.Contract.PAUSEROLE(&_CSModule.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCaller) RECOVERERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "RECOVERER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_CSModule *CSModuleSession) RECOVERERROLE() ([32]byte, error) {
	return _CSModule.Contract.RECOVERERROLE(&_CSModule.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCallerSession) RECOVERERROLE() ([32]byte, error) {
	return _CSModule.Contract.RECOVERERROLE(&_CSModule.CallOpts)
}

// REPORTELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x8573e351.
//
// Solidity: function REPORT_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCaller) REPORTELREWARDSSTEALINGPENALTYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "REPORT_EL_REWARDS_STEALING_PENALTY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// REPORTELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x8573e351.
//
// Solidity: function REPORT_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_CSModule *CSModuleSession) REPORTELREWARDSSTEALINGPENALTYROLE() ([32]byte, error) {
	return _CSModule.Contract.REPORTELREWARDSSTEALINGPENALTYROLE(&_CSModule.CallOpts)
}

// REPORTELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x8573e351.
//
// Solidity: function REPORT_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCallerSession) REPORTELREWARDSSTEALINGPENALTYROLE() ([32]byte, error) {
	return _CSModule.Contract.REPORTELREWARDSSTEALINGPENALTYROLE(&_CSModule.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_CSModule *CSModuleSession) RESUMEROLE() ([32]byte, error) {
	return _CSModule.Contract.RESUMEROLE(&_CSModule.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCallerSession) RESUMEROLE() ([32]byte, error) {
	return _CSModule.Contract.RESUMEROLE(&_CSModule.CallOpts)
}

// SETTLEELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x3f04f0c8.
//
// Solidity: function SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCaller) SETTLEELREWARDSSTEALINGPENALTYROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETTLEELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x3f04f0c8.
//
// Solidity: function SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_CSModule *CSModuleSession) SETTLEELREWARDSSTEALINGPENALTYROLE() ([32]byte, error) {
	return _CSModule.Contract.SETTLEELREWARDSSTEALINGPENALTYROLE(&_CSModule.CallOpts)
}

// SETTLEELREWARDSSTEALINGPENALTYROLE is a free data retrieval call binding the contract method 0x3f04f0c8.
//
// Solidity: function SETTLE_EL_REWARDS_STEALING_PENALTY_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCallerSession) SETTLEELREWARDSSTEALINGPENALTYROLE() ([32]byte, error) {
	return _CSModule.Contract.SETTLEELREWARDSSTEALINGPENALTYROLE(&_CSModule.CallOpts)
}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCaller) STAKINGROUTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "STAKING_ROUTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_CSModule *CSModuleSession) STAKINGROUTERROLE() ([32]byte, error) {
	return _CSModule.Contract.STAKINGROUTERROLE(&_CSModule.CallOpts)
}

// STAKINGROUTERROLE is a free data retrieval call binding the contract method 0x80231f15.
//
// Solidity: function STAKING_ROUTER_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCallerSession) STAKINGROUTERROLE() ([32]byte, error) {
	return _CSModule.Contract.STAKINGROUTERROLE(&_CSModule.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_CSModule *CSModuleCaller) STETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "STETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_CSModule *CSModuleSession) STETH() (common.Address, error) {
	return _CSModule.Contract.STETH(&_CSModule.CallOpts)
}

// STETH is a free data retrieval call binding the contract method 0xe00bfe50.
//
// Solidity: function STETH() view returns(address)
func (_CSModule *CSModuleCallerSession) STETH() (common.Address, error) {
	return _CSModule.Contract.STETH(&_CSModule.CallOpts)
}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCaller) VERIFIERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "VERIFIER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_CSModule *CSModuleSession) VERIFIERROLE() ([32]byte, error) {
	return _CSModule.Contract.VERIFIERROLE(&_CSModule.CallOpts)
}

// VERIFIERROLE is a free data retrieval call binding the contract method 0xe7705db6.
//
// Solidity: function VERIFIER_ROLE() view returns(bytes32)
func (_CSModule *CSModuleCallerSession) VERIFIERROLE() ([32]byte, error) {
	return _CSModule.Contract.VERIFIERROLE(&_CSModule.CallOpts)
}

// Accounting is a free data retrieval call binding the contract method 0x9624e83e.
//
// Solidity: function accounting() view returns(address)
func (_CSModule *CSModuleCaller) Accounting(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "accounting")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Accounting is a free data retrieval call binding the contract method 0x9624e83e.
//
// Solidity: function accounting() view returns(address)
func (_CSModule *CSModuleSession) Accounting() (common.Address, error) {
	return _CSModule.Contract.Accounting(&_CSModule.CallOpts)
}

// Accounting is a free data retrieval call binding the contract method 0x9624e83e.
//
// Solidity: function accounting() view returns(address)
func (_CSModule *CSModuleCallerSession) Accounting() (common.Address, error) {
	return _CSModule.Contract.Accounting(&_CSModule.CallOpts)
}

// DepositQueue is a free data retrieval call binding the contract method 0xf617eecc.
//
// Solidity: function depositQueue() view returns(uint128 head, uint128 tail)
func (_CSModule *CSModuleCaller) DepositQueue(opts *bind.CallOpts) (struct {
	Head *big.Int
	Tail *big.Int
}, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "depositQueue")

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

// DepositQueue is a free data retrieval call binding the contract method 0xf617eecc.
//
// Solidity: function depositQueue() view returns(uint128 head, uint128 tail)
func (_CSModule *CSModuleSession) DepositQueue() (struct {
	Head *big.Int
	Tail *big.Int
}, error) {
	return _CSModule.Contract.DepositQueue(&_CSModule.CallOpts)
}

// DepositQueue is a free data retrieval call binding the contract method 0xf617eecc.
//
// Solidity: function depositQueue() view returns(uint128 head, uint128 tail)
func (_CSModule *CSModuleCallerSession) DepositQueue() (struct {
	Head *big.Int
	Tail *big.Int
}, error) {
	return _CSModule.Contract.DepositQueue(&_CSModule.CallOpts)
}

// DepositQueueItem is a free data retrieval call binding the contract method 0x5e169fb8.
//
// Solidity: function depositQueueItem(uint128 index) view returns(uint256)
func (_CSModule *CSModuleCaller) DepositQueueItem(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "depositQueueItem", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DepositQueueItem is a free data retrieval call binding the contract method 0x5e169fb8.
//
// Solidity: function depositQueueItem(uint128 index) view returns(uint256)
func (_CSModule *CSModuleSession) DepositQueueItem(index *big.Int) (*big.Int, error) {
	return _CSModule.Contract.DepositQueueItem(&_CSModule.CallOpts, index)
}

// DepositQueueItem is a free data retrieval call binding the contract method 0x5e169fb8.
//
// Solidity: function depositQueueItem(uint128 index) view returns(uint256)
func (_CSModule *CSModuleCallerSession) DepositQueueItem(index *big.Int) (*big.Int, error) {
	return _CSModule.Contract.DepositQueueItem(&_CSModule.CallOpts, index)
}

// EarlyAdoption is a free data retrieval call binding the contract method 0x26a666e4.
//
// Solidity: function earlyAdoption() view returns(address)
func (_CSModule *CSModuleCaller) EarlyAdoption(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "earlyAdoption")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// EarlyAdoption is a free data retrieval call binding the contract method 0x26a666e4.
//
// Solidity: function earlyAdoption() view returns(address)
func (_CSModule *CSModuleSession) EarlyAdoption() (common.Address, error) {
	return _CSModule.Contract.EarlyAdoption(&_CSModule.CallOpts)
}

// EarlyAdoption is a free data retrieval call binding the contract method 0x26a666e4.
//
// Solidity: function earlyAdoption() view returns(address)
func (_CSModule *CSModuleCallerSession) EarlyAdoption() (common.Address, error) {
	return _CSModule.Contract.EarlyAdoption(&_CSModule.CallOpts)
}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_CSModule *CSModuleCaller) GetActiveNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getActiveNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_CSModule *CSModuleSession) GetActiveNodeOperatorsCount() (*big.Int, error) {
	return _CSModule.Contract.GetActiveNodeOperatorsCount(&_CSModule.CallOpts)
}

// GetActiveNodeOperatorsCount is a free data retrieval call binding the contract method 0x8469cbd3.
//
// Solidity: function getActiveNodeOperatorsCount() view returns(uint256)
func (_CSModule *CSModuleCallerSession) GetActiveNodeOperatorsCount() (*big.Int, error) {
	return _CSModule.Contract.GetActiveNodeOperatorsCount(&_CSModule.CallOpts)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x65c14dc7.
//
// Solidity: function getNodeOperator(uint256 nodeOperatorId) view returns((uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint8,uint32,uint32,address,address,address,address,bool))
func (_CSModule *CSModuleCaller) GetNodeOperator(opts *bind.CallOpts, nodeOperatorId *big.Int) (NodeOperator, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getNodeOperator", nodeOperatorId)

	if err != nil {
		return *new(NodeOperator), err
	}

	out0 := *abi.ConvertType(out[0], new(NodeOperator)).(*NodeOperator)

	return out0, err

}

// GetNodeOperator is a free data retrieval call binding the contract method 0x65c14dc7.
//
// Solidity: function getNodeOperator(uint256 nodeOperatorId) view returns((uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint8,uint32,uint32,address,address,address,address,bool))
func (_CSModule *CSModuleSession) GetNodeOperator(nodeOperatorId *big.Int) (NodeOperator, error) {
	return _CSModule.Contract.GetNodeOperator(&_CSModule.CallOpts, nodeOperatorId)
}

// GetNodeOperator is a free data retrieval call binding the contract method 0x65c14dc7.
//
// Solidity: function getNodeOperator(uint256 nodeOperatorId) view returns((uint32,uint32,uint32,uint32,uint32,uint32,uint32,uint8,uint32,uint32,address,address,address,address,bool))
func (_CSModule *CSModuleCallerSession) GetNodeOperator(nodeOperatorId *big.Int) (NodeOperator, error) {
	return _CSModule.Contract.GetNodeOperator(&_CSModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 offset, uint256 limit) view returns(uint256[] nodeOperatorIds)
func (_CSModule *CSModuleCaller) GetNodeOperatorIds(opts *bind.CallOpts, offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getNodeOperatorIds", offset, limit)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 offset, uint256 limit) view returns(uint256[] nodeOperatorIds)
func (_CSModule *CSModuleSession) GetNodeOperatorIds(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _CSModule.Contract.GetNodeOperatorIds(&_CSModule.CallOpts, offset, limit)
}

// GetNodeOperatorIds is a free data retrieval call binding the contract method 0x4febc81b.
//
// Solidity: function getNodeOperatorIds(uint256 offset, uint256 limit) view returns(uint256[] nodeOperatorIds)
func (_CSModule *CSModuleCallerSession) GetNodeOperatorIds(offset *big.Int, limit *big.Int) ([]*big.Int, error) {
	return _CSModule.Contract.GetNodeOperatorIds(&_CSModule.CallOpts, offset, limit)
}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 nodeOperatorId) view returns(bool)
func (_CSModule *CSModuleCaller) GetNodeOperatorIsActive(opts *bind.CallOpts, nodeOperatorId *big.Int) (bool, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getNodeOperatorIsActive", nodeOperatorId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 nodeOperatorId) view returns(bool)
func (_CSModule *CSModuleSession) GetNodeOperatorIsActive(nodeOperatorId *big.Int) (bool, error) {
	return _CSModule.Contract.GetNodeOperatorIsActive(&_CSModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorIsActive is a free data retrieval call binding the contract method 0x5e2fb908.
//
// Solidity: function getNodeOperatorIsActive(uint256 nodeOperatorId) view returns(bool)
func (_CSModule *CSModuleCallerSession) GetNodeOperatorIsActive(nodeOperatorId *big.Int) (bool, error) {
	return _CSModule.Contract.GetNodeOperatorIsActive(&_CSModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorNonWithdrawnKeys is a free data retrieval call binding the contract method 0x8ec69028.
//
// Solidity: function getNodeOperatorNonWithdrawnKeys(uint256 nodeOperatorId) view returns(uint256)
func (_CSModule *CSModuleCaller) GetNodeOperatorNonWithdrawnKeys(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getNodeOperatorNonWithdrawnKeys", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorNonWithdrawnKeys is a free data retrieval call binding the contract method 0x8ec69028.
//
// Solidity: function getNodeOperatorNonWithdrawnKeys(uint256 nodeOperatorId) view returns(uint256)
func (_CSModule *CSModuleSession) GetNodeOperatorNonWithdrawnKeys(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSModule.Contract.GetNodeOperatorNonWithdrawnKeys(&_CSModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorNonWithdrawnKeys is a free data retrieval call binding the contract method 0x8ec69028.
//
// Solidity: function getNodeOperatorNonWithdrawnKeys(uint256 nodeOperatorId) view returns(uint256)
func (_CSModule *CSModuleCallerSession) GetNodeOperatorNonWithdrawnKeys(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSModule.Contract.GetNodeOperatorNonWithdrawnKeys(&_CSModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xb3076c3c.
//
// Solidity: function getNodeOperatorSummary(uint256 nodeOperatorId) view returns(uint256 targetLimitMode, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_CSModule *CSModuleCaller) GetNodeOperatorSummary(opts *bind.CallOpts, nodeOperatorId *big.Int) (struct {
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
	err := _CSModule.contract.Call(opts, &out, "getNodeOperatorSummary", nodeOperatorId)

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
func (_CSModule *CSModuleSession) GetNodeOperatorSummary(nodeOperatorId *big.Int) (struct {
	TargetLimitMode            *big.Int
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _CSModule.Contract.GetNodeOperatorSummary(&_CSModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorSummary is a free data retrieval call binding the contract method 0xb3076c3c.
//
// Solidity: function getNodeOperatorSummary(uint256 nodeOperatorId) view returns(uint256 targetLimitMode, uint256 targetValidatorsCount, uint256 stuckValidatorsCount, uint256 refundedValidatorsCount, uint256 stuckPenaltyEndTimestamp, uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_CSModule *CSModuleCallerSession) GetNodeOperatorSummary(nodeOperatorId *big.Int) (struct {
	TargetLimitMode            *big.Int
	TargetValidatorsCount      *big.Int
	StuckValidatorsCount       *big.Int
	RefundedValidatorsCount    *big.Int
	StuckPenaltyEndTimestamp   *big.Int
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _CSModule.Contract.GetNodeOperatorSummary(&_CSModule.CallOpts, nodeOperatorId)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_CSModule *CSModuleCaller) GetNodeOperatorsCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getNodeOperatorsCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_CSModule *CSModuleSession) GetNodeOperatorsCount() (*big.Int, error) {
	return _CSModule.Contract.GetNodeOperatorsCount(&_CSModule.CallOpts)
}

// GetNodeOperatorsCount is a free data retrieval call binding the contract method 0xa70c70e4.
//
// Solidity: function getNodeOperatorsCount() view returns(uint256)
func (_CSModule *CSModuleCallerSession) GetNodeOperatorsCount() (*big.Int, error) {
	return _CSModule.Contract.GetNodeOperatorsCount(&_CSModule.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_CSModule *CSModuleCaller) GetNonce(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getNonce")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_CSModule *CSModuleSession) GetNonce() (*big.Int, error) {
	return _CSModule.Contract.GetNonce(&_CSModule.CallOpts)
}

// GetNonce is a free data retrieval call binding the contract method 0xd087d288.
//
// Solidity: function getNonce() view returns(uint256)
func (_CSModule *CSModuleCallerSession) GetNonce() (*big.Int, error) {
	return _CSModule.Contract.GetNonce(&_CSModule.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_CSModule *CSModuleCaller) GetResumeSinceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getResumeSinceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_CSModule *CSModuleSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _CSModule.Contract.GetResumeSinceTimestamp(&_CSModule.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_CSModule *CSModuleCallerSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _CSModule.Contract.GetResumeSinceTimestamp(&_CSModule.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CSModule *CSModuleCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CSModule *CSModuleSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CSModule.Contract.GetRoleAdmin(&_CSModule.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CSModule *CSModuleCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CSModule.Contract.GetRoleAdmin(&_CSModule.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CSModule *CSModuleCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CSModule *CSModuleSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CSModule.Contract.GetRoleMember(&_CSModule.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CSModule *CSModuleCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CSModule.Contract.GetRoleMember(&_CSModule.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CSModule *CSModuleCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CSModule *CSModuleSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CSModule.Contract.GetRoleMemberCount(&_CSModule.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CSModule *CSModuleCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CSModule.Contract.GetRoleMemberCount(&_CSModule.CallOpts, role)
}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes)
func (_CSModule *CSModuleCaller) GetSigningKeys(opts *bind.CallOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]byte, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getSigningKeys", nodeOperatorId, startIndex, keysCount)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes)
func (_CSModule *CSModuleSession) GetSigningKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]byte, error) {
	return _CSModule.Contract.GetSigningKeys(&_CSModule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetSigningKeys is a free data retrieval call binding the contract method 0x59e25c12.
//
// Solidity: function getSigningKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes)
func (_CSModule *CSModuleCallerSession) GetSigningKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) ([]byte, error) {
	return _CSModule.Contract.GetSigningKeys(&_CSModule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetSigningKeysWithSignatures is a free data retrieval call binding the contract method 0x50388cb6.
//
// Solidity: function getSigningKeysWithSignatures(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys, bytes signatures)
func (_CSModule *CSModuleCaller) GetSigningKeysWithSignatures(opts *bind.CallOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (struct {
	Keys       []byte
	Signatures []byte
}, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getSigningKeysWithSignatures", nodeOperatorId, startIndex, keysCount)

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
func (_CSModule *CSModuleSession) GetSigningKeysWithSignatures(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (struct {
	Keys       []byte
	Signatures []byte
}, error) {
	return _CSModule.Contract.GetSigningKeysWithSignatures(&_CSModule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetSigningKeysWithSignatures is a free data retrieval call binding the contract method 0x50388cb6.
//
// Solidity: function getSigningKeysWithSignatures(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) view returns(bytes keys, bytes signatures)
func (_CSModule *CSModuleCallerSession) GetSigningKeysWithSignatures(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (struct {
	Keys       []byte
	Signatures []byte
}, error) {
	return _CSModule.Contract.GetSigningKeysWithSignatures(&_CSModule.CallOpts, nodeOperatorId, startIndex, keysCount)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x9abddf09.
//
// Solidity: function getStakingModuleSummary() view returns(uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_CSModule *CSModuleCaller) GetStakingModuleSummary(opts *bind.CallOpts) (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getStakingModuleSummary")

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
func (_CSModule *CSModuleSession) GetStakingModuleSummary() (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _CSModule.Contract.GetStakingModuleSummary(&_CSModule.CallOpts)
}

// GetStakingModuleSummary is a free data retrieval call binding the contract method 0x9abddf09.
//
// Solidity: function getStakingModuleSummary() view returns(uint256 totalExitedValidators, uint256 totalDepositedValidators, uint256 depositableValidatorsCount)
func (_CSModule *CSModuleCallerSession) GetStakingModuleSummary() (struct {
	TotalExitedValidators      *big.Int
	TotalDepositedValidators   *big.Int
	DepositableValidatorsCount *big.Int
}, error) {
	return _CSModule.Contract.GetStakingModuleSummary(&_CSModule.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_CSModule *CSModuleCaller) GetType(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "getType")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_CSModule *CSModuleSession) GetType() ([32]byte, error) {
	return _CSModule.Contract.GetType(&_CSModule.CallOpts)
}

// GetType is a free data retrieval call binding the contract method 0x15dae03e.
//
// Solidity: function getType() view returns(bytes32)
func (_CSModule *CSModuleCallerSession) GetType() ([32]byte, error) {
	return _CSModule.Contract.GetType(&_CSModule.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CSModule *CSModuleCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CSModule *CSModuleSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CSModule.Contract.HasRole(&_CSModule.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CSModule *CSModuleCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CSModule.Contract.HasRole(&_CSModule.CallOpts, role, account)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CSModule *CSModuleCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CSModule *CSModuleSession) IsPaused() (bool, error) {
	return _CSModule.Contract.IsPaused(&_CSModule.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CSModule *CSModuleCallerSession) IsPaused() (bool, error) {
	return _CSModule.Contract.IsPaused(&_CSModule.CallOpts)
}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0x3dbe8b5a.
//
// Solidity: function isValidatorSlashed(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_CSModule *CSModuleCaller) IsValidatorSlashed(opts *bind.CallOpts, nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "isValidatorSlashed", nodeOperatorId, keyIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0x3dbe8b5a.
//
// Solidity: function isValidatorSlashed(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_CSModule *CSModuleSession) IsValidatorSlashed(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _CSModule.Contract.IsValidatorSlashed(&_CSModule.CallOpts, nodeOperatorId, keyIndex)
}

// IsValidatorSlashed is a free data retrieval call binding the contract method 0x3dbe8b5a.
//
// Solidity: function isValidatorSlashed(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_CSModule *CSModuleCallerSession) IsValidatorSlashed(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _CSModule.Contract.IsValidatorSlashed(&_CSModule.CallOpts, nodeOperatorId, keyIndex)
}

// IsValidatorWithdrawn is a free data retrieval call binding the contract method 0x53433643.
//
// Solidity: function isValidatorWithdrawn(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_CSModule *CSModuleCaller) IsValidatorWithdrawn(opts *bind.CallOpts, nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "isValidatorWithdrawn", nodeOperatorId, keyIndex)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsValidatorWithdrawn is a free data retrieval call binding the contract method 0x53433643.
//
// Solidity: function isValidatorWithdrawn(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_CSModule *CSModuleSession) IsValidatorWithdrawn(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _CSModule.Contract.IsValidatorWithdrawn(&_CSModule.CallOpts, nodeOperatorId, keyIndex)
}

// IsValidatorWithdrawn is a free data retrieval call binding the contract method 0x53433643.
//
// Solidity: function isValidatorWithdrawn(uint256 nodeOperatorId, uint256 keyIndex) view returns(bool)
func (_CSModule *CSModuleCallerSession) IsValidatorWithdrawn(nodeOperatorId *big.Int, keyIndex *big.Int) (bool, error) {
	return _CSModule.Contract.IsValidatorWithdrawn(&_CSModule.CallOpts, nodeOperatorId, keyIndex)
}

// KeyRemovalCharge is a free data retrieval call binding the contract method 0xd9df8c92.
//
// Solidity: function keyRemovalCharge() view returns(uint256)
func (_CSModule *CSModuleCaller) KeyRemovalCharge(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "keyRemovalCharge")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// KeyRemovalCharge is a free data retrieval call binding the contract method 0xd9df8c92.
//
// Solidity: function keyRemovalCharge() view returns(uint256)
func (_CSModule *CSModuleSession) KeyRemovalCharge() (*big.Int, error) {
	return _CSModule.Contract.KeyRemovalCharge(&_CSModule.CallOpts)
}

// KeyRemovalCharge is a free data retrieval call binding the contract method 0xd9df8c92.
//
// Solidity: function keyRemovalCharge() view returns(uint256)
func (_CSModule *CSModuleCallerSession) KeyRemovalCharge() (*big.Int, error) {
	return _CSModule.Contract.KeyRemovalCharge(&_CSModule.CallOpts)
}

// PublicRelease is a free data retrieval call binding the contract method 0xe21a430b.
//
// Solidity: function publicRelease() view returns(bool)
func (_CSModule *CSModuleCaller) PublicRelease(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "publicRelease")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PublicRelease is a free data retrieval call binding the contract method 0xe21a430b.
//
// Solidity: function publicRelease() view returns(bool)
func (_CSModule *CSModuleSession) PublicRelease() (bool, error) {
	return _CSModule.Contract.PublicRelease(&_CSModule.CallOpts)
}

// PublicRelease is a free data retrieval call binding the contract method 0xe21a430b.
//
// Solidity: function publicRelease() view returns(bool)
func (_CSModule *CSModuleCallerSession) PublicRelease() (bool, error) {
	return _CSModule.Contract.PublicRelease(&_CSModule.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CSModule *CSModuleCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CSModule.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CSModule *CSModuleSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CSModule.Contract.SupportsInterface(&_CSModule.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CSModule *CSModuleCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CSModule.Contract.SupportsInterface(&_CSModule.CallOpts, interfaceId)
}

// ActivatePublicRelease is a paid mutator transaction binding the contract method 0xd3931457.
//
// Solidity: function activatePublicRelease() returns()
func (_CSModule *CSModuleTransactor) ActivatePublicRelease(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "activatePublicRelease")
}

// ActivatePublicRelease is a paid mutator transaction binding the contract method 0xd3931457.
//
// Solidity: function activatePublicRelease() returns()
func (_CSModule *CSModuleSession) ActivatePublicRelease() (*types.Transaction, error) {
	return _CSModule.Contract.ActivatePublicRelease(&_CSModule.TransactOpts)
}

// ActivatePublicRelease is a paid mutator transaction binding the contract method 0xd3931457.
//
// Solidity: function activatePublicRelease() returns()
func (_CSModule *CSModuleTransactorSession) ActivatePublicRelease() (*types.Transaction, error) {
	return _CSModule.Contract.ActivatePublicRelease(&_CSModule.TransactOpts)
}

// AddNodeOperatorETH is a paid mutator transaction binding the contract method 0x157a039b.
//
// Solidity: function addNodeOperatorETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, bytes32[] eaProof, address referrer) payable returns()
func (_CSModule *CSModuleTransactor) AddNodeOperatorETH(opts *bind.TransactOpts, keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "addNodeOperatorETH", keysCount, publicKeys, signatures, managementProperties, eaProof, referrer)
}

// AddNodeOperatorETH is a paid mutator transaction binding the contract method 0x157a039b.
//
// Solidity: function addNodeOperatorETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, bytes32[] eaProof, address referrer) payable returns()
func (_CSModule *CSModuleSession) AddNodeOperatorETH(keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.AddNodeOperatorETH(&_CSModule.TransactOpts, keysCount, publicKeys, signatures, managementProperties, eaProof, referrer)
}

// AddNodeOperatorETH is a paid mutator transaction binding the contract method 0x157a039b.
//
// Solidity: function addNodeOperatorETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, bytes32[] eaProof, address referrer) payable returns()
func (_CSModule *CSModuleTransactorSession) AddNodeOperatorETH(keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.AddNodeOperatorETH(&_CSModule.TransactOpts, keysCount, publicKeys, signatures, managementProperties, eaProof, referrer)
}

// AddNodeOperatorStETH is a paid mutator transaction binding the contract method 0x6a5f2c4a.
//
// Solidity: function addNodeOperatorStETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, (uint256,uint256,uint8,bytes32,bytes32) permit, bytes32[] eaProof, address referrer) returns()
func (_CSModule *CSModuleTransactor) AddNodeOperatorStETH(opts *bind.TransactOpts, keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, permit ICSAccountingPermitInput, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "addNodeOperatorStETH", keysCount, publicKeys, signatures, managementProperties, permit, eaProof, referrer)
}

// AddNodeOperatorStETH is a paid mutator transaction binding the contract method 0x6a5f2c4a.
//
// Solidity: function addNodeOperatorStETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, (uint256,uint256,uint8,bytes32,bytes32) permit, bytes32[] eaProof, address referrer) returns()
func (_CSModule *CSModuleSession) AddNodeOperatorStETH(keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, permit ICSAccountingPermitInput, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.AddNodeOperatorStETH(&_CSModule.TransactOpts, keysCount, publicKeys, signatures, managementProperties, permit, eaProof, referrer)
}

// AddNodeOperatorStETH is a paid mutator transaction binding the contract method 0x6a5f2c4a.
//
// Solidity: function addNodeOperatorStETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, (uint256,uint256,uint8,bytes32,bytes32) permit, bytes32[] eaProof, address referrer) returns()
func (_CSModule *CSModuleTransactorSession) AddNodeOperatorStETH(keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, permit ICSAccountingPermitInput, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.AddNodeOperatorStETH(&_CSModule.TransactOpts, keysCount, publicKeys, signatures, managementProperties, permit, eaProof, referrer)
}

// AddNodeOperatorWstETH is a paid mutator transaction binding the contract method 0xacc446eb.
//
// Solidity: function addNodeOperatorWstETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, (uint256,uint256,uint8,bytes32,bytes32) permit, bytes32[] eaProof, address referrer) returns()
func (_CSModule *CSModuleTransactor) AddNodeOperatorWstETH(opts *bind.TransactOpts, keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, permit ICSAccountingPermitInput, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "addNodeOperatorWstETH", keysCount, publicKeys, signatures, managementProperties, permit, eaProof, referrer)
}

// AddNodeOperatorWstETH is a paid mutator transaction binding the contract method 0xacc446eb.
//
// Solidity: function addNodeOperatorWstETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, (uint256,uint256,uint8,bytes32,bytes32) permit, bytes32[] eaProof, address referrer) returns()
func (_CSModule *CSModuleSession) AddNodeOperatorWstETH(keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, permit ICSAccountingPermitInput, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.AddNodeOperatorWstETH(&_CSModule.TransactOpts, keysCount, publicKeys, signatures, managementProperties, permit, eaProof, referrer)
}

// AddNodeOperatorWstETH is a paid mutator transaction binding the contract method 0xacc446eb.
//
// Solidity: function addNodeOperatorWstETH(uint256 keysCount, bytes publicKeys, bytes signatures, (address,address,bool) managementProperties, (uint256,uint256,uint8,bytes32,bytes32) permit, bytes32[] eaProof, address referrer) returns()
func (_CSModule *CSModuleTransactorSession) AddNodeOperatorWstETH(keysCount *big.Int, publicKeys []byte, signatures []byte, managementProperties NodeOperatorManagementProperties, permit ICSAccountingPermitInput, eaProof [][32]byte, referrer common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.AddNodeOperatorWstETH(&_CSModule.TransactOpts, keysCount, publicKeys, signatures, managementProperties, permit, eaProof, referrer)
}

// AddValidatorKeysETH is a paid mutator transaction binding the contract method 0xfe7ed3cd.
//
// Solidity: function addValidatorKeysETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures) payable returns()
func (_CSModule *CSModuleTransactor) AddValidatorKeysETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "addValidatorKeysETH", nodeOperatorId, keysCount, publicKeys, signatures)
}

// AddValidatorKeysETH is a paid mutator transaction binding the contract method 0xfe7ed3cd.
//
// Solidity: function addValidatorKeysETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures) payable returns()
func (_CSModule *CSModuleSession) AddValidatorKeysETH(nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte) (*types.Transaction, error) {
	return _CSModule.Contract.AddValidatorKeysETH(&_CSModule.TransactOpts, nodeOperatorId, keysCount, publicKeys, signatures)
}

// AddValidatorKeysETH is a paid mutator transaction binding the contract method 0xfe7ed3cd.
//
// Solidity: function addValidatorKeysETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures) payable returns()
func (_CSModule *CSModuleTransactorSession) AddValidatorKeysETH(nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte) (*types.Transaction, error) {
	return _CSModule.Contract.AddValidatorKeysETH(&_CSModule.TransactOpts, nodeOperatorId, keysCount, publicKeys, signatures)
}

// AddValidatorKeysStETH is a paid mutator transaction binding the contract method 0x946654ad.
//
// Solidity: function addValidatorKeysStETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSModule *CSModuleTransactor) AddValidatorKeysStETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "addValidatorKeysStETH", nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysStETH is a paid mutator transaction binding the contract method 0x946654ad.
//
// Solidity: function addValidatorKeysStETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSModule *CSModuleSession) AddValidatorKeysStETH(nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSModule.Contract.AddValidatorKeysStETH(&_CSModule.TransactOpts, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysStETH is a paid mutator transaction binding the contract method 0x946654ad.
//
// Solidity: function addValidatorKeysStETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSModule *CSModuleTransactorSession) AddValidatorKeysStETH(nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSModule.Contract.AddValidatorKeysStETH(&_CSModule.TransactOpts, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysWstETH is a paid mutator transaction binding the contract method 0x9ec3c24c.
//
// Solidity: function addValidatorKeysWstETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSModule *CSModuleTransactor) AddValidatorKeysWstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "addValidatorKeysWstETH", nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysWstETH is a paid mutator transaction binding the contract method 0x9ec3c24c.
//
// Solidity: function addValidatorKeysWstETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSModule *CSModuleSession) AddValidatorKeysWstETH(nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSModule.Contract.AddValidatorKeysWstETH(&_CSModule.TransactOpts, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// AddValidatorKeysWstETH is a paid mutator transaction binding the contract method 0x9ec3c24c.
//
// Solidity: function addValidatorKeysWstETH(uint256 nodeOperatorId, uint256 keysCount, bytes publicKeys, bytes signatures, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSModule *CSModuleTransactorSession) AddValidatorKeysWstETH(nodeOperatorId *big.Int, keysCount *big.Int, publicKeys []byte, signatures []byte, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSModule.Contract.AddValidatorKeysWstETH(&_CSModule.TransactOpts, nodeOperatorId, keysCount, publicKeys, signatures, permit)
}

// CancelELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x40044801.
//
// Solidity: function cancelELRewardsStealingPenalty(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSModule *CSModuleTransactor) CancelELRewardsStealingPenalty(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "cancelELRewardsStealingPenalty", nodeOperatorId, amount)
}

// CancelELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x40044801.
//
// Solidity: function cancelELRewardsStealingPenalty(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSModule *CSModuleSession) CancelELRewardsStealingPenalty(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.CancelELRewardsStealingPenalty(&_CSModule.TransactOpts, nodeOperatorId, amount)
}

// CancelELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x40044801.
//
// Solidity: function cancelELRewardsStealingPenalty(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSModule *CSModuleTransactorSession) CancelELRewardsStealingPenalty(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.CancelELRewardsStealingPenalty(&_CSModule.TransactOpts, nodeOperatorId, amount)
}

// ChangeNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x75a401da.
//
// Solidity: function changeNodeOperatorRewardAddress(uint256 nodeOperatorId, address newAddress) returns()
func (_CSModule *CSModuleTransactor) ChangeNodeOperatorRewardAddress(opts *bind.TransactOpts, nodeOperatorId *big.Int, newAddress common.Address) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "changeNodeOperatorRewardAddress", nodeOperatorId, newAddress)
}

// ChangeNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x75a401da.
//
// Solidity: function changeNodeOperatorRewardAddress(uint256 nodeOperatorId, address newAddress) returns()
func (_CSModule *CSModuleSession) ChangeNodeOperatorRewardAddress(nodeOperatorId *big.Int, newAddress common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.ChangeNodeOperatorRewardAddress(&_CSModule.TransactOpts, nodeOperatorId, newAddress)
}

// ChangeNodeOperatorRewardAddress is a paid mutator transaction binding the contract method 0x75a401da.
//
// Solidity: function changeNodeOperatorRewardAddress(uint256 nodeOperatorId, address newAddress) returns()
func (_CSModule *CSModuleTransactorSession) ChangeNodeOperatorRewardAddress(nodeOperatorId *big.Int, newAddress common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.ChangeNodeOperatorRewardAddress(&_CSModule.TransactOpts, nodeOperatorId, newAddress)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0x8409d4fe.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSModule *CSModuleTransactor) ClaimRewardsStETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, stETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "claimRewardsStETH", nodeOperatorId, stETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0x8409d4fe.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSModule *CSModuleSession) ClaimRewardsStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSModule.Contract.ClaimRewardsStETH(&_CSModule.TransactOpts, nodeOperatorId, stETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0x8409d4fe.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSModule *CSModuleTransactorSession) ClaimRewardsStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSModule.Contract.ClaimRewardsStETH(&_CSModule.TransactOpts, nodeOperatorId, stETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0x3df6c438.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stEthAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSModule *CSModuleTransactor) ClaimRewardsUnstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, stEthAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "claimRewardsUnstETH", nodeOperatorId, stEthAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0x3df6c438.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stEthAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSModule *CSModuleSession) ClaimRewardsUnstETH(nodeOperatorId *big.Int, stEthAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSModule.Contract.ClaimRewardsUnstETH(&_CSModule.TransactOpts, nodeOperatorId, stEthAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0x3df6c438.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stEthAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSModule *CSModuleTransactorSession) ClaimRewardsUnstETH(nodeOperatorId *big.Int, stEthAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSModule.Contract.ClaimRewardsUnstETH(&_CSModule.TransactOpts, nodeOperatorId, stEthAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x5097ef59.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSModule *CSModuleTransactor) ClaimRewardsWstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, wstETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "claimRewardsWstETH", nodeOperatorId, wstETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x5097ef59.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSModule *CSModuleSession) ClaimRewardsWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSModule.Contract.ClaimRewardsWstETH(&_CSModule.TransactOpts, nodeOperatorId, wstETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x5097ef59.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSModule *CSModuleTransactorSession) ClaimRewardsWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSModule.Contract.ClaimRewardsWstETH(&_CSModule.TransactOpts, nodeOperatorId, wstETHAmount, cumulativeFeeShares, rewardsProof)
}

// CleanDepositQueue is a paid mutator transaction binding the contract method 0x735dfa28.
//
// Solidity: function cleanDepositQueue(uint256 maxItems) returns(uint256)
func (_CSModule *CSModuleTransactor) CleanDepositQueue(opts *bind.TransactOpts, maxItems *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "cleanDepositQueue", maxItems)
}

// CleanDepositQueue is a paid mutator transaction binding the contract method 0x735dfa28.
//
// Solidity: function cleanDepositQueue(uint256 maxItems) returns(uint256)
func (_CSModule *CSModuleSession) CleanDepositQueue(maxItems *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.CleanDepositQueue(&_CSModule.TransactOpts, maxItems)
}

// CleanDepositQueue is a paid mutator transaction binding the contract method 0x735dfa28.
//
// Solidity: function cleanDepositQueue(uint256 maxItems) returns(uint256)
func (_CSModule *CSModuleTransactorSession) CleanDepositQueue(maxItems *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.CleanDepositQueue(&_CSModule.TransactOpts, maxItems)
}

// CompensateELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x6efe37a2.
//
// Solidity: function compensateELRewardsStealingPenalty(uint256 nodeOperatorId) payable returns()
func (_CSModule *CSModuleTransactor) CompensateELRewardsStealingPenalty(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "compensateELRewardsStealingPenalty", nodeOperatorId)
}

// CompensateELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x6efe37a2.
//
// Solidity: function compensateELRewardsStealingPenalty(uint256 nodeOperatorId) payable returns()
func (_CSModule *CSModuleSession) CompensateELRewardsStealingPenalty(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.CompensateELRewardsStealingPenalty(&_CSModule.TransactOpts, nodeOperatorId)
}

// CompensateELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x6efe37a2.
//
// Solidity: function compensateELRewardsStealingPenalty(uint256 nodeOperatorId) payable returns()
func (_CSModule *CSModuleTransactorSession) CompensateELRewardsStealingPenalty(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.CompensateELRewardsStealingPenalty(&_CSModule.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x6bb1bfdf.
//
// Solidity: function confirmNodeOperatorManagerAddressChange(uint256 nodeOperatorId) returns()
func (_CSModule *CSModuleTransactor) ConfirmNodeOperatorManagerAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "confirmNodeOperatorManagerAddressChange", nodeOperatorId)
}

// ConfirmNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x6bb1bfdf.
//
// Solidity: function confirmNodeOperatorManagerAddressChange(uint256 nodeOperatorId) returns()
func (_CSModule *CSModuleSession) ConfirmNodeOperatorManagerAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.ConfirmNodeOperatorManagerAddressChange(&_CSModule.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x6bb1bfdf.
//
// Solidity: function confirmNodeOperatorManagerAddressChange(uint256 nodeOperatorId) returns()
func (_CSModule *CSModuleTransactorSession) ConfirmNodeOperatorManagerAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.ConfirmNodeOperatorManagerAddressChange(&_CSModule.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x5204281c.
//
// Solidity: function confirmNodeOperatorRewardAddressChange(uint256 nodeOperatorId) returns()
func (_CSModule *CSModuleTransactor) ConfirmNodeOperatorRewardAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "confirmNodeOperatorRewardAddressChange", nodeOperatorId)
}

// ConfirmNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x5204281c.
//
// Solidity: function confirmNodeOperatorRewardAddressChange(uint256 nodeOperatorId) returns()
func (_CSModule *CSModuleSession) ConfirmNodeOperatorRewardAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.ConfirmNodeOperatorRewardAddressChange(&_CSModule.TransactOpts, nodeOperatorId)
}

// ConfirmNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x5204281c.
//
// Solidity: function confirmNodeOperatorRewardAddressChange(uint256 nodeOperatorId) returns()
func (_CSModule *CSModuleTransactorSession) ConfirmNodeOperatorRewardAddressChange(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.ConfirmNodeOperatorRewardAddressChange(&_CSModule.TransactOpts, nodeOperatorId)
}

// DecreaseVettedSigningKeysCount is a paid mutator transaction binding the contract method 0xb643189b.
//
// Solidity: function decreaseVettedSigningKeysCount(bytes nodeOperatorIds, bytes vettedSigningKeysCounts) returns()
func (_CSModule *CSModuleTransactor) DecreaseVettedSigningKeysCount(opts *bind.TransactOpts, nodeOperatorIds []byte, vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "decreaseVettedSigningKeysCount", nodeOperatorIds, vettedSigningKeysCounts)
}

// DecreaseVettedSigningKeysCount is a paid mutator transaction binding the contract method 0xb643189b.
//
// Solidity: function decreaseVettedSigningKeysCount(bytes nodeOperatorIds, bytes vettedSigningKeysCounts) returns()
func (_CSModule *CSModuleSession) DecreaseVettedSigningKeysCount(nodeOperatorIds []byte, vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _CSModule.Contract.DecreaseVettedSigningKeysCount(&_CSModule.TransactOpts, nodeOperatorIds, vettedSigningKeysCounts)
}

// DecreaseVettedSigningKeysCount is a paid mutator transaction binding the contract method 0xb643189b.
//
// Solidity: function decreaseVettedSigningKeysCount(bytes nodeOperatorIds, bytes vettedSigningKeysCounts) returns()
func (_CSModule *CSModuleTransactorSession) DecreaseVettedSigningKeysCount(nodeOperatorIds []byte, vettedSigningKeysCounts []byte) (*types.Transaction, error) {
	return _CSModule.Contract.DecreaseVettedSigningKeysCount(&_CSModule.TransactOpts, nodeOperatorIds, vettedSigningKeysCounts)
}

// DepositETH is a paid mutator transaction binding the contract method 0x5358fbda.
//
// Solidity: function depositETH(uint256 nodeOperatorId) payable returns()
func (_CSModule *CSModuleTransactor) DepositETH(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "depositETH", nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x5358fbda.
//
// Solidity: function depositETH(uint256 nodeOperatorId) payable returns()
func (_CSModule *CSModuleSession) DepositETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.DepositETH(&_CSModule.TransactOpts, nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x5358fbda.
//
// Solidity: function depositETH(uint256 nodeOperatorId) payable returns()
func (_CSModule *CSModuleTransactorSession) DepositETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.DepositETH(&_CSModule.TransactOpts, nodeOperatorId)
}

// DepositStETH is a paid mutator transaction binding the contract method 0xe1aa105d.
//
// Solidity: function depositStETH(uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSModule *CSModuleTransactor) DepositStETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "depositStETH", nodeOperatorId, stETHAmount, permit)
}

// DepositStETH is a paid mutator transaction binding the contract method 0xe1aa105d.
//
// Solidity: function depositStETH(uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSModule *CSModuleSession) DepositStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSModule.Contract.DepositStETH(&_CSModule.TransactOpts, nodeOperatorId, stETHAmount, permit)
}

// DepositStETH is a paid mutator transaction binding the contract method 0xe1aa105d.
//
// Solidity: function depositStETH(uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSModule *CSModuleTransactorSession) DepositStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSModule.Contract.DepositStETH(&_CSModule.TransactOpts, nodeOperatorId, stETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0x3f214bb2.
//
// Solidity: function depositWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSModule *CSModuleTransactor) DepositWstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "depositWstETH", nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0x3f214bb2.
//
// Solidity: function depositWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSModule *CSModuleSession) DepositWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSModule.Contract.DepositWstETH(&_CSModule.TransactOpts, nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0x3f214bb2.
//
// Solidity: function depositWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSModule *CSModuleTransactorSession) DepositWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSModule.Contract.DepositWstETH(&_CSModule.TransactOpts, nodeOperatorId, wstETHAmount, permit)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CSModule *CSModuleTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CSModule *CSModuleSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.GrantRole(&_CSModule.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CSModule *CSModuleTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.GrantRole(&_CSModule.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address _accounting, address _earlyAdoption, uint256 _keyRemovalCharge, address admin) returns()
func (_CSModule *CSModuleTransactor) Initialize(opts *bind.TransactOpts, _accounting common.Address, _earlyAdoption common.Address, _keyRemovalCharge *big.Int, admin common.Address) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "initialize", _accounting, _earlyAdoption, _keyRemovalCharge, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address _accounting, address _earlyAdoption, uint256 _keyRemovalCharge, address admin) returns()
func (_CSModule *CSModuleSession) Initialize(_accounting common.Address, _earlyAdoption common.Address, _keyRemovalCharge *big.Int, admin common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.Initialize(&_CSModule.TransactOpts, _accounting, _earlyAdoption, _keyRemovalCharge, admin)
}

// Initialize is a paid mutator transaction binding the contract method 0xbe203094.
//
// Solidity: function initialize(address _accounting, address _earlyAdoption, uint256 _keyRemovalCharge, address admin) returns()
func (_CSModule *CSModuleTransactorSession) Initialize(_accounting common.Address, _earlyAdoption common.Address, _keyRemovalCharge *big.Int, admin common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.Initialize(&_CSModule.TransactOpts, _accounting, _earlyAdoption, _keyRemovalCharge, admin)
}

// NormalizeQueue is a paid mutator transaction binding the contract method 0xb1520dc5.
//
// Solidity: function normalizeQueue(uint256 nodeOperatorId) returns()
func (_CSModule *CSModuleTransactor) NormalizeQueue(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "normalizeQueue", nodeOperatorId)
}

// NormalizeQueue is a paid mutator transaction binding the contract method 0xb1520dc5.
//
// Solidity: function normalizeQueue(uint256 nodeOperatorId) returns()
func (_CSModule *CSModuleSession) NormalizeQueue(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.NormalizeQueue(&_CSModule.TransactOpts, nodeOperatorId)
}

// NormalizeQueue is a paid mutator transaction binding the contract method 0xb1520dc5.
//
// Solidity: function normalizeQueue(uint256 nodeOperatorId) returns()
func (_CSModule *CSModuleTransactorSession) NormalizeQueue(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.NormalizeQueue(&_CSModule.TransactOpts, nodeOperatorId)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_CSModule *CSModuleTransactor) ObtainDepositData(opts *bind.TransactOpts, depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "obtainDepositData", depositsCount, arg1)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_CSModule *CSModuleSession) ObtainDepositData(depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _CSModule.Contract.ObtainDepositData(&_CSModule.TransactOpts, depositsCount, arg1)
}

// ObtainDepositData is a paid mutator transaction binding the contract method 0xbee41b58.
//
// Solidity: function obtainDepositData(uint256 depositsCount, bytes ) returns(bytes publicKeys, bytes signatures)
func (_CSModule *CSModuleTransactorSession) ObtainDepositData(depositsCount *big.Int, arg1 []byte) (*types.Transaction, error) {
	return _CSModule.Contract.ObtainDepositData(&_CSModule.TransactOpts, depositsCount, arg1)
}

// OnExitedAndStuckValidatorsCountsUpdated is a paid mutator transaction binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() returns()
func (_CSModule *CSModuleTransactor) OnExitedAndStuckValidatorsCountsUpdated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "onExitedAndStuckValidatorsCountsUpdated")
}

// OnExitedAndStuckValidatorsCountsUpdated is a paid mutator transaction binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() returns()
func (_CSModule *CSModuleSession) OnExitedAndStuckValidatorsCountsUpdated() (*types.Transaction, error) {
	return _CSModule.Contract.OnExitedAndStuckValidatorsCountsUpdated(&_CSModule.TransactOpts)
}

// OnExitedAndStuckValidatorsCountsUpdated is a paid mutator transaction binding the contract method 0xe864299e.
//
// Solidity: function onExitedAndStuckValidatorsCountsUpdated() returns()
func (_CSModule *CSModuleTransactorSession) OnExitedAndStuckValidatorsCountsUpdated() (*types.Transaction, error) {
	return _CSModule.Contract.OnExitedAndStuckValidatorsCountsUpdated(&_CSModule.TransactOpts)
}

// OnRewardsMinted is a paid mutator transaction binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 totalShares) returns()
func (_CSModule *CSModuleTransactor) OnRewardsMinted(opts *bind.TransactOpts, totalShares *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "onRewardsMinted", totalShares)
}

// OnRewardsMinted is a paid mutator transaction binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 totalShares) returns()
func (_CSModule *CSModuleSession) OnRewardsMinted(totalShares *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.OnRewardsMinted(&_CSModule.TransactOpts, totalShares)
}

// OnRewardsMinted is a paid mutator transaction binding the contract method 0x8d7e4017.
//
// Solidity: function onRewardsMinted(uint256 totalShares) returns()
func (_CSModule *CSModuleTransactorSession) OnRewardsMinted(totalShares *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.OnRewardsMinted(&_CSModule.TransactOpts, totalShares)
}

// OnWithdrawalCredentialsChanged is a paid mutator transaction binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() returns()
func (_CSModule *CSModuleTransactor) OnWithdrawalCredentialsChanged(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "onWithdrawalCredentialsChanged")
}

// OnWithdrawalCredentialsChanged is a paid mutator transaction binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() returns()
func (_CSModule *CSModuleSession) OnWithdrawalCredentialsChanged() (*types.Transaction, error) {
	return _CSModule.Contract.OnWithdrawalCredentialsChanged(&_CSModule.TransactOpts)
}

// OnWithdrawalCredentialsChanged is a paid mutator transaction binding the contract method 0x90c09bdb.
//
// Solidity: function onWithdrawalCredentialsChanged() returns()
func (_CSModule *CSModuleTransactorSession) OnWithdrawalCredentialsChanged() (*types.Transaction, error) {
	return _CSModule.Contract.OnWithdrawalCredentialsChanged(&_CSModule.TransactOpts)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_CSModule *CSModuleTransactor) PauseFor(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "pauseFor", duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_CSModule *CSModuleSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.PauseFor(&_CSModule.TransactOpts, duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_CSModule *CSModuleTransactorSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.PauseFor(&_CSModule.TransactOpts, duration)
}

// ProposeNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x8cabe959.
//
// Solidity: function proposeNodeOperatorManagerAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_CSModule *CSModuleTransactor) ProposeNodeOperatorManagerAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "proposeNodeOperatorManagerAddressChange", nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x8cabe959.
//
// Solidity: function proposeNodeOperatorManagerAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_CSModule *CSModuleSession) ProposeNodeOperatorManagerAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.ProposeNodeOperatorManagerAddressChange(&_CSModule.TransactOpts, nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorManagerAddressChange is a paid mutator transaction binding the contract method 0x8cabe959.
//
// Solidity: function proposeNodeOperatorManagerAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_CSModule *CSModuleTransactorSession) ProposeNodeOperatorManagerAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.ProposeNodeOperatorManagerAddressChange(&_CSModule.TransactOpts, nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x1b40b231.
//
// Solidity: function proposeNodeOperatorRewardAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_CSModule *CSModuleTransactor) ProposeNodeOperatorRewardAddressChange(opts *bind.TransactOpts, nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "proposeNodeOperatorRewardAddressChange", nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x1b40b231.
//
// Solidity: function proposeNodeOperatorRewardAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_CSModule *CSModuleSession) ProposeNodeOperatorRewardAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.ProposeNodeOperatorRewardAddressChange(&_CSModule.TransactOpts, nodeOperatorId, proposedAddress)
}

// ProposeNodeOperatorRewardAddressChange is a paid mutator transaction binding the contract method 0x1b40b231.
//
// Solidity: function proposeNodeOperatorRewardAddressChange(uint256 nodeOperatorId, address proposedAddress) returns()
func (_CSModule *CSModuleTransactorSession) ProposeNodeOperatorRewardAddressChange(nodeOperatorId *big.Int, proposedAddress common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.ProposeNodeOperatorRewardAddressChange(&_CSModule.TransactOpts, nodeOperatorId, proposedAddress)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_CSModule *CSModuleTransactor) RecoverERC1155(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "recoverERC1155", token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_CSModule *CSModuleSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.RecoverERC1155(&_CSModule.TransactOpts, token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_CSModule *CSModuleTransactorSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.RecoverERC1155(&_CSModule.TransactOpts, token, tokenId)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_CSModule *CSModuleTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "recoverERC20", token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_CSModule *CSModuleSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.RecoverERC20(&_CSModule.TransactOpts, token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_CSModule *CSModuleTransactorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.RecoverERC20(&_CSModule.TransactOpts, token, amount)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_CSModule *CSModuleTransactor) RecoverERC721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "recoverERC721", token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_CSModule *CSModuleSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.RecoverERC721(&_CSModule.TransactOpts, token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_CSModule *CSModuleTransactorSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.RecoverERC721(&_CSModule.TransactOpts, token, tokenId)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_CSModule *CSModuleTransactor) RecoverEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "recoverEther")
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_CSModule *CSModuleSession) RecoverEther() (*types.Transaction, error) {
	return _CSModule.Contract.RecoverEther(&_CSModule.TransactOpts)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_CSModule *CSModuleTransactorSession) RecoverEther() (*types.Transaction, error) {
	return _CSModule.Contract.RecoverEther(&_CSModule.TransactOpts)
}

// RecoverStETHShares is a paid mutator transaction binding the contract method 0x5a73bdc8.
//
// Solidity: function recoverStETHShares() returns()
func (_CSModule *CSModuleTransactor) RecoverStETHShares(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "recoverStETHShares")
}

// RecoverStETHShares is a paid mutator transaction binding the contract method 0x5a73bdc8.
//
// Solidity: function recoverStETHShares() returns()
func (_CSModule *CSModuleSession) RecoverStETHShares() (*types.Transaction, error) {
	return _CSModule.Contract.RecoverStETHShares(&_CSModule.TransactOpts)
}

// RecoverStETHShares is a paid mutator transaction binding the contract method 0x5a73bdc8.
//
// Solidity: function recoverStETHShares() returns()
func (_CSModule *CSModuleTransactorSession) RecoverStETHShares() (*types.Transaction, error) {
	return _CSModule.Contract.RecoverStETHShares(&_CSModule.TransactOpts)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0x8b3ac71d.
//
// Solidity: function removeKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) returns()
func (_CSModule *CSModuleTransactor) RemoveKeys(opts *bind.TransactOpts, nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "removeKeys", nodeOperatorId, startIndex, keysCount)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0x8b3ac71d.
//
// Solidity: function removeKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) returns()
func (_CSModule *CSModuleSession) RemoveKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.RemoveKeys(&_CSModule.TransactOpts, nodeOperatorId, startIndex, keysCount)
}

// RemoveKeys is a paid mutator transaction binding the contract method 0x8b3ac71d.
//
// Solidity: function removeKeys(uint256 nodeOperatorId, uint256 startIndex, uint256 keysCount) returns()
func (_CSModule *CSModuleTransactorSession) RemoveKeys(nodeOperatorId *big.Int, startIndex *big.Int, keysCount *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.RemoveKeys(&_CSModule.TransactOpts, nodeOperatorId, startIndex, keysCount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CSModule *CSModuleTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CSModule *CSModuleSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.RenounceRole(&_CSModule.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CSModule *CSModuleTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.RenounceRole(&_CSModule.TransactOpts, role, callerConfirmation)
}

// ReportELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x388dd1d1.
//
// Solidity: function reportELRewardsStealingPenalty(uint256 nodeOperatorId, bytes32 blockHash, uint256 amount) returns()
func (_CSModule *CSModuleTransactor) ReportELRewardsStealingPenalty(opts *bind.TransactOpts, nodeOperatorId *big.Int, blockHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "reportELRewardsStealingPenalty", nodeOperatorId, blockHash, amount)
}

// ReportELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x388dd1d1.
//
// Solidity: function reportELRewardsStealingPenalty(uint256 nodeOperatorId, bytes32 blockHash, uint256 amount) returns()
func (_CSModule *CSModuleSession) ReportELRewardsStealingPenalty(nodeOperatorId *big.Int, blockHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.ReportELRewardsStealingPenalty(&_CSModule.TransactOpts, nodeOperatorId, blockHash, amount)
}

// ReportELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x388dd1d1.
//
// Solidity: function reportELRewardsStealingPenalty(uint256 nodeOperatorId, bytes32 blockHash, uint256 amount) returns()
func (_CSModule *CSModuleTransactorSession) ReportELRewardsStealingPenalty(nodeOperatorId *big.Int, blockHash [32]byte, amount *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.ReportELRewardsStealingPenalty(&_CSModule.TransactOpts, nodeOperatorId, blockHash, amount)
}

// ResetNodeOperatorManagerAddress is a paid mutator transaction binding the contract method 0x6a6304cc.
//
// Solidity: function resetNodeOperatorManagerAddress(uint256 nodeOperatorId) returns()
func (_CSModule *CSModuleTransactor) ResetNodeOperatorManagerAddress(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "resetNodeOperatorManagerAddress", nodeOperatorId)
}

// ResetNodeOperatorManagerAddress is a paid mutator transaction binding the contract method 0x6a6304cc.
//
// Solidity: function resetNodeOperatorManagerAddress(uint256 nodeOperatorId) returns()
func (_CSModule *CSModuleSession) ResetNodeOperatorManagerAddress(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.ResetNodeOperatorManagerAddress(&_CSModule.TransactOpts, nodeOperatorId)
}

// ResetNodeOperatorManagerAddress is a paid mutator transaction binding the contract method 0x6a6304cc.
//
// Solidity: function resetNodeOperatorManagerAddress(uint256 nodeOperatorId) returns()
func (_CSModule *CSModuleTransactorSession) ResetNodeOperatorManagerAddress(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.ResetNodeOperatorManagerAddress(&_CSModule.TransactOpts, nodeOperatorId)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_CSModule *CSModuleTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_CSModule *CSModuleSession) Resume() (*types.Transaction, error) {
	return _CSModule.Contract.Resume(&_CSModule.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_CSModule *CSModuleTransactorSession) Resume() (*types.Transaction, error) {
	return _CSModule.Contract.Resume(&_CSModule.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CSModule *CSModuleTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CSModule *CSModuleSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.RevokeRole(&_CSModule.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CSModule *CSModuleTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CSModule.Contract.RevokeRole(&_CSModule.TransactOpts, role, account)
}

// SetKeyRemovalCharge is a paid mutator transaction binding the contract method 0xba1557ae.
//
// Solidity: function setKeyRemovalCharge(uint256 amount) returns()
func (_CSModule *CSModuleTransactor) SetKeyRemovalCharge(opts *bind.TransactOpts, amount *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "setKeyRemovalCharge", amount)
}

// SetKeyRemovalCharge is a paid mutator transaction binding the contract method 0xba1557ae.
//
// Solidity: function setKeyRemovalCharge(uint256 amount) returns()
func (_CSModule *CSModuleSession) SetKeyRemovalCharge(amount *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.SetKeyRemovalCharge(&_CSModule.TransactOpts, amount)
}

// SetKeyRemovalCharge is a paid mutator transaction binding the contract method 0xba1557ae.
//
// Solidity: function setKeyRemovalCharge(uint256 amount) returns()
func (_CSModule *CSModuleTransactorSession) SetKeyRemovalCharge(amount *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.SetKeyRemovalCharge(&_CSModule.TransactOpts, amount)
}

// SettleELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x37b12b5f.
//
// Solidity: function settleELRewardsStealingPenalty(uint256[] nodeOperatorIds) returns()
func (_CSModule *CSModuleTransactor) SettleELRewardsStealingPenalty(opts *bind.TransactOpts, nodeOperatorIds []*big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "settleELRewardsStealingPenalty", nodeOperatorIds)
}

// SettleELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x37b12b5f.
//
// Solidity: function settleELRewardsStealingPenalty(uint256[] nodeOperatorIds) returns()
func (_CSModule *CSModuleSession) SettleELRewardsStealingPenalty(nodeOperatorIds []*big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.SettleELRewardsStealingPenalty(&_CSModule.TransactOpts, nodeOperatorIds)
}

// SettleELRewardsStealingPenalty is a paid mutator transaction binding the contract method 0x37b12b5f.
//
// Solidity: function settleELRewardsStealingPenalty(uint256[] nodeOperatorIds) returns()
func (_CSModule *CSModuleTransactorSession) SettleELRewardsStealingPenalty(nodeOperatorIds []*big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.SettleELRewardsStealingPenalty(&_CSModule.TransactOpts, nodeOperatorIds)
}

// SubmitInitialSlashing is a paid mutator transaction binding the contract method 0xf96d3952.
//
// Solidity: function submitInitialSlashing(uint256 nodeOperatorId, uint256 keyIndex) returns()
func (_CSModule *CSModuleTransactor) SubmitInitialSlashing(opts *bind.TransactOpts, nodeOperatorId *big.Int, keyIndex *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "submitInitialSlashing", nodeOperatorId, keyIndex)
}

// SubmitInitialSlashing is a paid mutator transaction binding the contract method 0xf96d3952.
//
// Solidity: function submitInitialSlashing(uint256 nodeOperatorId, uint256 keyIndex) returns()
func (_CSModule *CSModuleSession) SubmitInitialSlashing(nodeOperatorId *big.Int, keyIndex *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.SubmitInitialSlashing(&_CSModule.TransactOpts, nodeOperatorId, keyIndex)
}

// SubmitInitialSlashing is a paid mutator transaction binding the contract method 0xf96d3952.
//
// Solidity: function submitInitialSlashing(uint256 nodeOperatorId, uint256 keyIndex) returns()
func (_CSModule *CSModuleTransactorSession) SubmitInitialSlashing(nodeOperatorId *big.Int, keyIndex *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.SubmitInitialSlashing(&_CSModule.TransactOpts, nodeOperatorId, keyIndex)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0xf408b551.
//
// Solidity: function submitWithdrawal(uint256 nodeOperatorId, uint256 keyIndex, uint256 amount, bool isSlashed) returns()
func (_CSModule *CSModuleTransactor) SubmitWithdrawal(opts *bind.TransactOpts, nodeOperatorId *big.Int, keyIndex *big.Int, amount *big.Int, isSlashed bool) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "submitWithdrawal", nodeOperatorId, keyIndex, amount, isSlashed)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0xf408b551.
//
// Solidity: function submitWithdrawal(uint256 nodeOperatorId, uint256 keyIndex, uint256 amount, bool isSlashed) returns()
func (_CSModule *CSModuleSession) SubmitWithdrawal(nodeOperatorId *big.Int, keyIndex *big.Int, amount *big.Int, isSlashed bool) (*types.Transaction, error) {
	return _CSModule.Contract.SubmitWithdrawal(&_CSModule.TransactOpts, nodeOperatorId, keyIndex, amount, isSlashed)
}

// SubmitWithdrawal is a paid mutator transaction binding the contract method 0xf408b551.
//
// Solidity: function submitWithdrawal(uint256 nodeOperatorId, uint256 keyIndex, uint256 amount, bool isSlashed) returns()
func (_CSModule *CSModuleTransactorSession) SubmitWithdrawal(nodeOperatorId *big.Int, keyIndex *big.Int, amount *big.Int, isSlashed bool) (*types.Transaction, error) {
	return _CSModule.Contract.SubmitWithdrawal(&_CSModule.TransactOpts, nodeOperatorId, keyIndex, amount, isSlashed)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0xf2e2ca63.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 nodeOperatorId, uint256 exitedValidatorsKeysCount, uint256 stuckValidatorsKeysCount) returns()
func (_CSModule *CSModuleTransactor) UnsafeUpdateValidatorsCount(opts *bind.TransactOpts, nodeOperatorId *big.Int, exitedValidatorsKeysCount *big.Int, stuckValidatorsKeysCount *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "unsafeUpdateValidatorsCount", nodeOperatorId, exitedValidatorsKeysCount, stuckValidatorsKeysCount)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0xf2e2ca63.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 nodeOperatorId, uint256 exitedValidatorsKeysCount, uint256 stuckValidatorsKeysCount) returns()
func (_CSModule *CSModuleSession) UnsafeUpdateValidatorsCount(nodeOperatorId *big.Int, exitedValidatorsKeysCount *big.Int, stuckValidatorsKeysCount *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.UnsafeUpdateValidatorsCount(&_CSModule.TransactOpts, nodeOperatorId, exitedValidatorsKeysCount, stuckValidatorsKeysCount)
}

// UnsafeUpdateValidatorsCount is a paid mutator transaction binding the contract method 0xf2e2ca63.
//
// Solidity: function unsafeUpdateValidatorsCount(uint256 nodeOperatorId, uint256 exitedValidatorsKeysCount, uint256 stuckValidatorsKeysCount) returns()
func (_CSModule *CSModuleTransactorSession) UnsafeUpdateValidatorsCount(nodeOperatorId *big.Int, exitedValidatorsKeysCount *big.Int, stuckValidatorsKeysCount *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.UnsafeUpdateValidatorsCount(&_CSModule.TransactOpts, nodeOperatorId, exitedValidatorsKeysCount, stuckValidatorsKeysCount)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes nodeOperatorIds, bytes exitedValidatorsCounts) returns()
func (_CSModule *CSModuleTransactor) UpdateExitedValidatorsCount(opts *bind.TransactOpts, nodeOperatorIds []byte, exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "updateExitedValidatorsCount", nodeOperatorIds, exitedValidatorsCounts)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes nodeOperatorIds, bytes exitedValidatorsCounts) returns()
func (_CSModule *CSModuleSession) UpdateExitedValidatorsCount(nodeOperatorIds []byte, exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _CSModule.Contract.UpdateExitedValidatorsCount(&_CSModule.TransactOpts, nodeOperatorIds, exitedValidatorsCounts)
}

// UpdateExitedValidatorsCount is a paid mutator transaction binding the contract method 0x9b00c146.
//
// Solidity: function updateExitedValidatorsCount(bytes nodeOperatorIds, bytes exitedValidatorsCounts) returns()
func (_CSModule *CSModuleTransactorSession) UpdateExitedValidatorsCount(nodeOperatorIds []byte, exitedValidatorsCounts []byte) (*types.Transaction, error) {
	return _CSModule.Contract.UpdateExitedValidatorsCount(&_CSModule.TransactOpts, nodeOperatorIds, exitedValidatorsCounts)
}

// UpdateRefundedValidatorsCount is a paid mutator transaction binding the contract method 0xa2e080f1.
//
// Solidity: function updateRefundedValidatorsCount(uint256 , uint256 ) returns()
func (_CSModule *CSModuleTransactor) UpdateRefundedValidatorsCount(opts *bind.TransactOpts, arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "updateRefundedValidatorsCount", arg0, arg1)
}

// UpdateRefundedValidatorsCount is a paid mutator transaction binding the contract method 0xa2e080f1.
//
// Solidity: function updateRefundedValidatorsCount(uint256 , uint256 ) returns()
func (_CSModule *CSModuleSession) UpdateRefundedValidatorsCount(arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.UpdateRefundedValidatorsCount(&_CSModule.TransactOpts, arg0, arg1)
}

// UpdateRefundedValidatorsCount is a paid mutator transaction binding the contract method 0xa2e080f1.
//
// Solidity: function updateRefundedValidatorsCount(uint256 , uint256 ) returns()
func (_CSModule *CSModuleTransactorSession) UpdateRefundedValidatorsCount(arg0 *big.Int, arg1 *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.UpdateRefundedValidatorsCount(&_CSModule.TransactOpts, arg0, arg1)
}

// UpdateStuckValidatorsCount is a paid mutator transaction binding the contract method 0x9b3d1900.
//
// Solidity: function updateStuckValidatorsCount(bytes nodeOperatorIds, bytes stuckValidatorsCounts) returns()
func (_CSModule *CSModuleTransactor) UpdateStuckValidatorsCount(opts *bind.TransactOpts, nodeOperatorIds []byte, stuckValidatorsCounts []byte) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "updateStuckValidatorsCount", nodeOperatorIds, stuckValidatorsCounts)
}

// UpdateStuckValidatorsCount is a paid mutator transaction binding the contract method 0x9b3d1900.
//
// Solidity: function updateStuckValidatorsCount(bytes nodeOperatorIds, bytes stuckValidatorsCounts) returns()
func (_CSModule *CSModuleSession) UpdateStuckValidatorsCount(nodeOperatorIds []byte, stuckValidatorsCounts []byte) (*types.Transaction, error) {
	return _CSModule.Contract.UpdateStuckValidatorsCount(&_CSModule.TransactOpts, nodeOperatorIds, stuckValidatorsCounts)
}

// UpdateStuckValidatorsCount is a paid mutator transaction binding the contract method 0x9b3d1900.
//
// Solidity: function updateStuckValidatorsCount(bytes nodeOperatorIds, bytes stuckValidatorsCounts) returns()
func (_CSModule *CSModuleTransactorSession) UpdateStuckValidatorsCount(nodeOperatorIds []byte, stuckValidatorsCounts []byte) (*types.Transaction, error) {
	return _CSModule.Contract.UpdateStuckValidatorsCount(&_CSModule.TransactOpts, nodeOperatorIds, stuckValidatorsCounts)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x08a679ad.
//
// Solidity: function updateTargetValidatorsLimits(uint256 nodeOperatorId, uint256 targetLimitMode, uint256 targetLimit) returns()
func (_CSModule *CSModuleTransactor) UpdateTargetValidatorsLimits(opts *bind.TransactOpts, nodeOperatorId *big.Int, targetLimitMode *big.Int, targetLimit *big.Int) (*types.Transaction, error) {
	return _CSModule.contract.Transact(opts, "updateTargetValidatorsLimits", nodeOperatorId, targetLimitMode, targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x08a679ad.
//
// Solidity: function updateTargetValidatorsLimits(uint256 nodeOperatorId, uint256 targetLimitMode, uint256 targetLimit) returns()
func (_CSModule *CSModuleSession) UpdateTargetValidatorsLimits(nodeOperatorId *big.Int, targetLimitMode *big.Int, targetLimit *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.UpdateTargetValidatorsLimits(&_CSModule.TransactOpts, nodeOperatorId, targetLimitMode, targetLimit)
}

// UpdateTargetValidatorsLimits is a paid mutator transaction binding the contract method 0x08a679ad.
//
// Solidity: function updateTargetValidatorsLimits(uint256 nodeOperatorId, uint256 targetLimitMode, uint256 targetLimit) returns()
func (_CSModule *CSModuleTransactorSession) UpdateTargetValidatorsLimits(nodeOperatorId *big.Int, targetLimitMode *big.Int, targetLimit *big.Int) (*types.Transaction, error) {
	return _CSModule.Contract.UpdateTargetValidatorsLimits(&_CSModule.TransactOpts, nodeOperatorId, targetLimitMode, targetLimit)
}

// CSModuleBatchEnqueuedIterator is returned from FilterBatchEnqueued and is used to iterate over the raw logs and unpacked data for BatchEnqueued events raised by the CSModule contract.
type CSModuleBatchEnqueuedIterator struct {
	Event *CSModuleBatchEnqueued // Event containing the contract specifics and raw log

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
func (it *CSModuleBatchEnqueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleBatchEnqueued)
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
		it.Event = new(CSModuleBatchEnqueued)
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
func (it *CSModuleBatchEnqueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleBatchEnqueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleBatchEnqueued represents a BatchEnqueued event raised by the CSModule contract.
type CSModuleBatchEnqueued struct {
	NodeOperatorId *big.Int
	Count          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBatchEnqueued is a free log retrieval operation binding the contract event 0x162b3db9d9ca7d0abe51ad8229dc058550a74b769457fd055579b5bdc5492536.
//
// Solidity: event BatchEnqueued(uint256 indexed nodeOperatorId, uint256 count)
func (_CSModule *CSModuleFilterer) FilterBatchEnqueued(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleBatchEnqueuedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "BatchEnqueued", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleBatchEnqueuedIterator{contract: _CSModule.contract, event: "BatchEnqueued", logs: logs, sub: sub}, nil
}

// WatchBatchEnqueued is a free log subscription operation binding the contract event 0x162b3db9d9ca7d0abe51ad8229dc058550a74b769457fd055579b5bdc5492536.
//
// Solidity: event BatchEnqueued(uint256 indexed nodeOperatorId, uint256 count)
func (_CSModule *CSModuleFilterer) WatchBatchEnqueued(opts *bind.WatchOpts, sink chan<- *CSModuleBatchEnqueued, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "BatchEnqueued", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleBatchEnqueued)
				if err := _CSModule.contract.UnpackLog(event, "BatchEnqueued", log); err != nil {
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

// ParseBatchEnqueued is a log parse operation binding the contract event 0x162b3db9d9ca7d0abe51ad8229dc058550a74b769457fd055579b5bdc5492536.
//
// Solidity: event BatchEnqueued(uint256 indexed nodeOperatorId, uint256 count)
func (_CSModule *CSModuleFilterer) ParseBatchEnqueued(log types.Log) (*CSModuleBatchEnqueued, error) {
	event := new(CSModuleBatchEnqueued)
	if err := _CSModule.contract.UnpackLog(event, "BatchEnqueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleDepositableSigningKeysCountChangedIterator is returned from FilterDepositableSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for DepositableSigningKeysCountChanged events raised by the CSModule contract.
type CSModuleDepositableSigningKeysCountChangedIterator struct {
	Event *CSModuleDepositableSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CSModuleDepositableSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleDepositableSigningKeysCountChanged)
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
		it.Event = new(CSModuleDepositableSigningKeysCountChanged)
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
func (it *CSModuleDepositableSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleDepositableSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleDepositableSigningKeysCountChanged represents a DepositableSigningKeysCountChanged event raised by the CSModule contract.
type CSModuleDepositableSigningKeysCountChanged struct {
	NodeOperatorId       *big.Int
	DepositableKeysCount *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterDepositableSigningKeysCountChanged is a free log retrieval operation binding the contract event 0xf9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed33.
//
// Solidity: event DepositableSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositableKeysCount)
func (_CSModule *CSModuleFilterer) FilterDepositableSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleDepositableSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "DepositableSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleDepositableSigningKeysCountChangedIterator{contract: _CSModule.contract, event: "DepositableSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchDepositableSigningKeysCountChanged is a free log subscription operation binding the contract event 0xf9109091b368cedad2edff45414eef892edd6b4fe80084bd590aa8f8def8ed33.
//
// Solidity: event DepositableSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositableKeysCount)
func (_CSModule *CSModuleFilterer) WatchDepositableSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CSModuleDepositableSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "DepositableSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleDepositableSigningKeysCountChanged)
				if err := _CSModule.contract.UnpackLog(event, "DepositableSigningKeysCountChanged", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseDepositableSigningKeysCountChanged(log types.Log) (*CSModuleDepositableSigningKeysCountChanged, error) {
	event := new(CSModuleDepositableSigningKeysCountChanged)
	if err := _CSModule.contract.UnpackLog(event, "DepositableSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleDepositedSigningKeysCountChangedIterator is returned from FilterDepositedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for DepositedSigningKeysCountChanged events raised by the CSModule contract.
type CSModuleDepositedSigningKeysCountChangedIterator struct {
	Event *CSModuleDepositedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CSModuleDepositedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleDepositedSigningKeysCountChanged)
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
		it.Event = new(CSModuleDepositedSigningKeysCountChanged)
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
func (it *CSModuleDepositedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleDepositedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleDepositedSigningKeysCountChanged represents a DepositedSigningKeysCountChanged event raised by the CSModule contract.
type CSModuleDepositedSigningKeysCountChanged struct {
	NodeOperatorId     *big.Int
	DepositedKeysCount *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterDepositedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b78.
//
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedKeysCount)
func (_CSModule *CSModuleFilterer) FilterDepositedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleDepositedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "DepositedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleDepositedSigningKeysCountChangedIterator{contract: _CSModule.contract, event: "DepositedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchDepositedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x24eb1c9e765ba41accf9437300ea91ece5ed3f897ec3cdee0e9debd7fe309b78.
//
// Solidity: event DepositedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 depositedKeysCount)
func (_CSModule *CSModuleFilterer) WatchDepositedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CSModuleDepositedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "DepositedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleDepositedSigningKeysCountChanged)
				if err := _CSModule.contract.UnpackLog(event, "DepositedSigningKeysCountChanged", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseDepositedSigningKeysCountChanged(log types.Log) (*CSModuleDepositedSigningKeysCountChanged, error) {
	event := new(CSModuleDepositedSigningKeysCountChanged)
	if err := _CSModule.contract.UnpackLog(event, "DepositedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleELRewardsStealingPenaltyCancelledIterator is returned from FilterELRewardsStealingPenaltyCancelled and is used to iterate over the raw logs and unpacked data for ELRewardsStealingPenaltyCancelled events raised by the CSModule contract.
type CSModuleELRewardsStealingPenaltyCancelledIterator struct {
	Event *CSModuleELRewardsStealingPenaltyCancelled // Event containing the contract specifics and raw log

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
func (it *CSModuleELRewardsStealingPenaltyCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleELRewardsStealingPenaltyCancelled)
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
		it.Event = new(CSModuleELRewardsStealingPenaltyCancelled)
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
func (it *CSModuleELRewardsStealingPenaltyCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleELRewardsStealingPenaltyCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleELRewardsStealingPenaltyCancelled represents a ELRewardsStealingPenaltyCancelled event raised by the CSModule contract.
type CSModuleELRewardsStealingPenaltyCancelled struct {
	NodeOperatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterELRewardsStealingPenaltyCancelled is a free log retrieval operation binding the contract event 0x1e7ebd3c5f4de9502000b6f7e6e7cf5d4ecb27d6fe1778e43fb9d1d0ca87d0e7.
//
// Solidity: event ELRewardsStealingPenaltyCancelled(uint256 indexed nodeOperatorId, uint256 amount)
func (_CSModule *CSModuleFilterer) FilterELRewardsStealingPenaltyCancelled(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleELRewardsStealingPenaltyCancelledIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "ELRewardsStealingPenaltyCancelled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleELRewardsStealingPenaltyCancelledIterator{contract: _CSModule.contract, event: "ELRewardsStealingPenaltyCancelled", logs: logs, sub: sub}, nil
}

// WatchELRewardsStealingPenaltyCancelled is a free log subscription operation binding the contract event 0x1e7ebd3c5f4de9502000b6f7e6e7cf5d4ecb27d6fe1778e43fb9d1d0ca87d0e7.
//
// Solidity: event ELRewardsStealingPenaltyCancelled(uint256 indexed nodeOperatorId, uint256 amount)
func (_CSModule *CSModuleFilterer) WatchELRewardsStealingPenaltyCancelled(opts *bind.WatchOpts, sink chan<- *CSModuleELRewardsStealingPenaltyCancelled, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "ELRewardsStealingPenaltyCancelled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleELRewardsStealingPenaltyCancelled)
				if err := _CSModule.contract.UnpackLog(event, "ELRewardsStealingPenaltyCancelled", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseELRewardsStealingPenaltyCancelled(log types.Log) (*CSModuleELRewardsStealingPenaltyCancelled, error) {
	event := new(CSModuleELRewardsStealingPenaltyCancelled)
	if err := _CSModule.contract.UnpackLog(event, "ELRewardsStealingPenaltyCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleELRewardsStealingPenaltyReportedIterator is returned from FilterELRewardsStealingPenaltyReported and is used to iterate over the raw logs and unpacked data for ELRewardsStealingPenaltyReported events raised by the CSModule contract.
type CSModuleELRewardsStealingPenaltyReportedIterator struct {
	Event *CSModuleELRewardsStealingPenaltyReported // Event containing the contract specifics and raw log

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
func (it *CSModuleELRewardsStealingPenaltyReportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleELRewardsStealingPenaltyReported)
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
		it.Event = new(CSModuleELRewardsStealingPenaltyReported)
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
func (it *CSModuleELRewardsStealingPenaltyReportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleELRewardsStealingPenaltyReportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleELRewardsStealingPenaltyReported represents a ELRewardsStealingPenaltyReported event raised by the CSModule contract.
type CSModuleELRewardsStealingPenaltyReported struct {
	NodeOperatorId    *big.Int
	ProposedBlockHash [32]byte
	StolenAmount      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterELRewardsStealingPenaltyReported is a free log retrieval operation binding the contract event 0xeec4d6dbe34149c6728a9638eca869d0e5a7fcd85c7a96178f7e9780b4b7fe4b.
//
// Solidity: event ELRewardsStealingPenaltyReported(uint256 indexed nodeOperatorId, bytes32 proposedBlockHash, uint256 stolenAmount)
func (_CSModule *CSModuleFilterer) FilterELRewardsStealingPenaltyReported(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleELRewardsStealingPenaltyReportedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "ELRewardsStealingPenaltyReported", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleELRewardsStealingPenaltyReportedIterator{contract: _CSModule.contract, event: "ELRewardsStealingPenaltyReported", logs: logs, sub: sub}, nil
}

// WatchELRewardsStealingPenaltyReported is a free log subscription operation binding the contract event 0xeec4d6dbe34149c6728a9638eca869d0e5a7fcd85c7a96178f7e9780b4b7fe4b.
//
// Solidity: event ELRewardsStealingPenaltyReported(uint256 indexed nodeOperatorId, bytes32 proposedBlockHash, uint256 stolenAmount)
func (_CSModule *CSModuleFilterer) WatchELRewardsStealingPenaltyReported(opts *bind.WatchOpts, sink chan<- *CSModuleELRewardsStealingPenaltyReported, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "ELRewardsStealingPenaltyReported", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleELRewardsStealingPenaltyReported)
				if err := _CSModule.contract.UnpackLog(event, "ELRewardsStealingPenaltyReported", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseELRewardsStealingPenaltyReported(log types.Log) (*CSModuleELRewardsStealingPenaltyReported, error) {
	event := new(CSModuleELRewardsStealingPenaltyReported)
	if err := _CSModule.contract.UnpackLog(event, "ELRewardsStealingPenaltyReported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleELRewardsStealingPenaltySettledIterator is returned from FilterELRewardsStealingPenaltySettled and is used to iterate over the raw logs and unpacked data for ELRewardsStealingPenaltySettled events raised by the CSModule contract.
type CSModuleELRewardsStealingPenaltySettledIterator struct {
	Event *CSModuleELRewardsStealingPenaltySettled // Event containing the contract specifics and raw log

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
func (it *CSModuleELRewardsStealingPenaltySettledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleELRewardsStealingPenaltySettled)
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
		it.Event = new(CSModuleELRewardsStealingPenaltySettled)
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
func (it *CSModuleELRewardsStealingPenaltySettledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleELRewardsStealingPenaltySettledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleELRewardsStealingPenaltySettled represents a ELRewardsStealingPenaltySettled event raised by the CSModule contract.
type CSModuleELRewardsStealingPenaltySettled struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterELRewardsStealingPenaltySettled is a free log retrieval operation binding the contract event 0x00f4fe19c0404d2fbb58da6f646c0a3ee5a6994a034213bbd22b072ed1ca5c27.
//
// Solidity: event ELRewardsStealingPenaltySettled(uint256 indexed nodeOperatorId)
func (_CSModule *CSModuleFilterer) FilterELRewardsStealingPenaltySettled(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleELRewardsStealingPenaltySettledIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "ELRewardsStealingPenaltySettled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleELRewardsStealingPenaltySettledIterator{contract: _CSModule.contract, event: "ELRewardsStealingPenaltySettled", logs: logs, sub: sub}, nil
}

// WatchELRewardsStealingPenaltySettled is a free log subscription operation binding the contract event 0x00f4fe19c0404d2fbb58da6f646c0a3ee5a6994a034213bbd22b072ed1ca5c27.
//
// Solidity: event ELRewardsStealingPenaltySettled(uint256 indexed nodeOperatorId)
func (_CSModule *CSModuleFilterer) WatchELRewardsStealingPenaltySettled(opts *bind.WatchOpts, sink chan<- *CSModuleELRewardsStealingPenaltySettled, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "ELRewardsStealingPenaltySettled", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleELRewardsStealingPenaltySettled)
				if err := _CSModule.contract.UnpackLog(event, "ELRewardsStealingPenaltySettled", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseELRewardsStealingPenaltySettled(log types.Log) (*CSModuleELRewardsStealingPenaltySettled, error) {
	event := new(CSModuleELRewardsStealingPenaltySettled)
	if err := _CSModule.contract.UnpackLog(event, "ELRewardsStealingPenaltySettled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleERC1155RecoveredIterator is returned from FilterERC1155Recovered and is used to iterate over the raw logs and unpacked data for ERC1155Recovered events raised by the CSModule contract.
type CSModuleERC1155RecoveredIterator struct {
	Event *CSModuleERC1155Recovered // Event containing the contract specifics and raw log

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
func (it *CSModuleERC1155RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleERC1155Recovered)
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
		it.Event = new(CSModuleERC1155Recovered)
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
func (it *CSModuleERC1155RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleERC1155RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleERC1155Recovered represents a ERC1155Recovered event raised by the CSModule contract.
type CSModuleERC1155Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC1155Recovered is a free log retrieval operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_CSModule *CSModuleFilterer) FilterERC1155Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CSModuleERC1155RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleERC1155RecoveredIterator{contract: _CSModule.contract, event: "ERC1155Recovered", logs: logs, sub: sub}, nil
}

// WatchERC1155Recovered is a free log subscription operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_CSModule *CSModuleFilterer) WatchERC1155Recovered(opts *bind.WatchOpts, sink chan<- *CSModuleERC1155Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleERC1155Recovered)
				if err := _CSModule.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseERC1155Recovered(log types.Log) (*CSModuleERC1155Recovered, error) {
	event := new(CSModuleERC1155Recovered)
	if err := _CSModule.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleERC20RecoveredIterator is returned from FilterERC20Recovered and is used to iterate over the raw logs and unpacked data for ERC20Recovered events raised by the CSModule contract.
type CSModuleERC20RecoveredIterator struct {
	Event *CSModuleERC20Recovered // Event containing the contract specifics and raw log

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
func (it *CSModuleERC20RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleERC20Recovered)
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
		it.Event = new(CSModuleERC20Recovered)
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
func (it *CSModuleERC20RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleERC20RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleERC20Recovered represents a ERC20Recovered event raised by the CSModule contract.
type CSModuleERC20Recovered struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Recovered is a free log retrieval operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_CSModule *CSModuleFilterer) FilterERC20Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CSModuleERC20RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleERC20RecoveredIterator{contract: _CSModule.contract, event: "ERC20Recovered", logs: logs, sub: sub}, nil
}

// WatchERC20Recovered is a free log subscription operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_CSModule *CSModuleFilterer) WatchERC20Recovered(opts *bind.WatchOpts, sink chan<- *CSModuleERC20Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleERC20Recovered)
				if err := _CSModule.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseERC20Recovered(log types.Log) (*CSModuleERC20Recovered, error) {
	event := new(CSModuleERC20Recovered)
	if err := _CSModule.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleERC721RecoveredIterator is returned from FilterERC721Recovered and is used to iterate over the raw logs and unpacked data for ERC721Recovered events raised by the CSModule contract.
type CSModuleERC721RecoveredIterator struct {
	Event *CSModuleERC721Recovered // Event containing the contract specifics and raw log

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
func (it *CSModuleERC721RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleERC721Recovered)
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
		it.Event = new(CSModuleERC721Recovered)
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
func (it *CSModuleERC721RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleERC721RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleERC721Recovered represents a ERC721Recovered event raised by the CSModule contract.
type CSModuleERC721Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721Recovered is a free log retrieval operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_CSModule *CSModuleFilterer) FilterERC721Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CSModuleERC721RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleERC721RecoveredIterator{contract: _CSModule.contract, event: "ERC721Recovered", logs: logs, sub: sub}, nil
}

// WatchERC721Recovered is a free log subscription operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_CSModule *CSModuleFilterer) WatchERC721Recovered(opts *bind.WatchOpts, sink chan<- *CSModuleERC721Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleERC721Recovered)
				if err := _CSModule.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseERC721Recovered(log types.Log) (*CSModuleERC721Recovered, error) {
	event := new(CSModuleERC721Recovered)
	if err := _CSModule.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleEtherRecoveredIterator is returned from FilterEtherRecovered and is used to iterate over the raw logs and unpacked data for EtherRecovered events raised by the CSModule contract.
type CSModuleEtherRecoveredIterator struct {
	Event *CSModuleEtherRecovered // Event containing the contract specifics and raw log

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
func (it *CSModuleEtherRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleEtherRecovered)
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
		it.Event = new(CSModuleEtherRecovered)
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
func (it *CSModuleEtherRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleEtherRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleEtherRecovered represents a EtherRecovered event raised by the CSModule contract.
type CSModuleEtherRecovered struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEtherRecovered is a free log retrieval operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_CSModule *CSModuleFilterer) FilterEtherRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CSModuleEtherRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleEtherRecoveredIterator{contract: _CSModule.contract, event: "EtherRecovered", logs: logs, sub: sub}, nil
}

// WatchEtherRecovered is a free log subscription operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_CSModule *CSModuleFilterer) WatchEtherRecovered(opts *bind.WatchOpts, sink chan<- *CSModuleEtherRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleEtherRecovered)
				if err := _CSModule.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseEtherRecovered(log types.Log) (*CSModuleEtherRecovered, error) {
	event := new(CSModuleEtherRecovered)
	if err := _CSModule.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleExitedSigningKeysCountChangedIterator is returned from FilterExitedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for ExitedSigningKeysCountChanged events raised by the CSModule contract.
type CSModuleExitedSigningKeysCountChangedIterator struct {
	Event *CSModuleExitedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CSModuleExitedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleExitedSigningKeysCountChanged)
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
		it.Event = new(CSModuleExitedSigningKeysCountChanged)
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
func (it *CSModuleExitedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleExitedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleExitedSigningKeysCountChanged represents a ExitedSigningKeysCountChanged event raised by the CSModule contract.
type CSModuleExitedSigningKeysCountChanged struct {
	NodeOperatorId  *big.Int
	ExitedKeysCount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterExitedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c220166.
//
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedKeysCount)
func (_CSModule *CSModuleFilterer) FilterExitedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleExitedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "ExitedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleExitedSigningKeysCountChangedIterator{contract: _CSModule.contract, event: "ExitedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchExitedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x0f67960648751434ae86bf350db61194f387fda387e7f568b0ccd0ae0c220166.
//
// Solidity: event ExitedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 exitedKeysCount)
func (_CSModule *CSModuleFilterer) WatchExitedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CSModuleExitedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "ExitedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleExitedSigningKeysCountChanged)
				if err := _CSModule.contract.UnpackLog(event, "ExitedSigningKeysCountChanged", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseExitedSigningKeysCountChanged(log types.Log) (*CSModuleExitedSigningKeysCountChanged, error) {
	event := new(CSModuleExitedSigningKeysCountChanged)
	if err := _CSModule.contract.UnpackLog(event, "ExitedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleInitialSlashingSubmittedIterator is returned from FilterInitialSlashingSubmitted and is used to iterate over the raw logs and unpacked data for InitialSlashingSubmitted events raised by the CSModule contract.
type CSModuleInitialSlashingSubmittedIterator struct {
	Event *CSModuleInitialSlashingSubmitted // Event containing the contract specifics and raw log

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
func (it *CSModuleInitialSlashingSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleInitialSlashingSubmitted)
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
		it.Event = new(CSModuleInitialSlashingSubmitted)
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
func (it *CSModuleInitialSlashingSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleInitialSlashingSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleInitialSlashingSubmitted represents a InitialSlashingSubmitted event raised by the CSModule contract.
type CSModuleInitialSlashingSubmitted struct {
	NodeOperatorId *big.Int
	KeyIndex       *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterInitialSlashingSubmitted is a free log retrieval operation binding the contract event 0xd34db8e8c0ddbc9c7b6dd8c397623dfbe01929e41e527540bff8794685d9b407.
//
// Solidity: event InitialSlashingSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex)
func (_CSModule *CSModuleFilterer) FilterInitialSlashingSubmitted(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleInitialSlashingSubmittedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "InitialSlashingSubmitted", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleInitialSlashingSubmittedIterator{contract: _CSModule.contract, event: "InitialSlashingSubmitted", logs: logs, sub: sub}, nil
}

// WatchInitialSlashingSubmitted is a free log subscription operation binding the contract event 0xd34db8e8c0ddbc9c7b6dd8c397623dfbe01929e41e527540bff8794685d9b407.
//
// Solidity: event InitialSlashingSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex)
func (_CSModule *CSModuleFilterer) WatchInitialSlashingSubmitted(opts *bind.WatchOpts, sink chan<- *CSModuleInitialSlashingSubmitted, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "InitialSlashingSubmitted", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleInitialSlashingSubmitted)
				if err := _CSModule.contract.UnpackLog(event, "InitialSlashingSubmitted", log); err != nil {
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

// ParseInitialSlashingSubmitted is a log parse operation binding the contract event 0xd34db8e8c0ddbc9c7b6dd8c397623dfbe01929e41e527540bff8794685d9b407.
//
// Solidity: event InitialSlashingSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex)
func (_CSModule *CSModuleFilterer) ParseInitialSlashingSubmitted(log types.Log) (*CSModuleInitialSlashingSubmitted, error) {
	event := new(CSModuleInitialSlashingSubmitted)
	if err := _CSModule.contract.UnpackLog(event, "InitialSlashingSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CSModule contract.
type CSModuleInitializedIterator struct {
	Event *CSModuleInitialized // Event containing the contract specifics and raw log

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
func (it *CSModuleInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleInitialized)
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
		it.Event = new(CSModuleInitialized)
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
func (it *CSModuleInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleInitialized represents a Initialized event raised by the CSModule contract.
type CSModuleInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CSModule *CSModuleFilterer) FilterInitialized(opts *bind.FilterOpts) (*CSModuleInitializedIterator, error) {

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CSModuleInitializedIterator{contract: _CSModule.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CSModule *CSModuleFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CSModuleInitialized) (event.Subscription, error) {

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleInitialized)
				if err := _CSModule.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseInitialized(log types.Log) (*CSModuleInitialized, error) {
	event := new(CSModuleInitialized)
	if err := _CSModule.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleKeyRemovalChargeAppliedIterator is returned from FilterKeyRemovalChargeApplied and is used to iterate over the raw logs and unpacked data for KeyRemovalChargeApplied events raised by the CSModule contract.
type CSModuleKeyRemovalChargeAppliedIterator struct {
	Event *CSModuleKeyRemovalChargeApplied // Event containing the contract specifics and raw log

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
func (it *CSModuleKeyRemovalChargeAppliedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleKeyRemovalChargeApplied)
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
		it.Event = new(CSModuleKeyRemovalChargeApplied)
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
func (it *CSModuleKeyRemovalChargeAppliedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleKeyRemovalChargeAppliedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleKeyRemovalChargeApplied represents a KeyRemovalChargeApplied event raised by the CSModule contract.
type CSModuleKeyRemovalChargeApplied struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterKeyRemovalChargeApplied is a free log retrieval operation binding the contract event 0x1cbb8dafbedbdf4f813a8ed1f50d871def63e1104f8729b677af57905eda90f6.
//
// Solidity: event KeyRemovalChargeApplied(uint256 indexed nodeOperatorId)
func (_CSModule *CSModuleFilterer) FilterKeyRemovalChargeApplied(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleKeyRemovalChargeAppliedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "KeyRemovalChargeApplied", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleKeyRemovalChargeAppliedIterator{contract: _CSModule.contract, event: "KeyRemovalChargeApplied", logs: logs, sub: sub}, nil
}

// WatchKeyRemovalChargeApplied is a free log subscription operation binding the contract event 0x1cbb8dafbedbdf4f813a8ed1f50d871def63e1104f8729b677af57905eda90f6.
//
// Solidity: event KeyRemovalChargeApplied(uint256 indexed nodeOperatorId)
func (_CSModule *CSModuleFilterer) WatchKeyRemovalChargeApplied(opts *bind.WatchOpts, sink chan<- *CSModuleKeyRemovalChargeApplied, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "KeyRemovalChargeApplied", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleKeyRemovalChargeApplied)
				if err := _CSModule.contract.UnpackLog(event, "KeyRemovalChargeApplied", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseKeyRemovalChargeApplied(log types.Log) (*CSModuleKeyRemovalChargeApplied, error) {
	event := new(CSModuleKeyRemovalChargeApplied)
	if err := _CSModule.contract.UnpackLog(event, "KeyRemovalChargeApplied", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleKeyRemovalChargeSetIterator is returned from FilterKeyRemovalChargeSet and is used to iterate over the raw logs and unpacked data for KeyRemovalChargeSet events raised by the CSModule contract.
type CSModuleKeyRemovalChargeSetIterator struct {
	Event *CSModuleKeyRemovalChargeSet // Event containing the contract specifics and raw log

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
func (it *CSModuleKeyRemovalChargeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleKeyRemovalChargeSet)
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
		it.Event = new(CSModuleKeyRemovalChargeSet)
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
func (it *CSModuleKeyRemovalChargeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleKeyRemovalChargeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleKeyRemovalChargeSet represents a KeyRemovalChargeSet event raised by the CSModule contract.
type CSModuleKeyRemovalChargeSet struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterKeyRemovalChargeSet is a free log retrieval operation binding the contract event 0x699ec9c671aad1f3dcc15e571375584a1d6fb7176afd545d14467fd31477e98e.
//
// Solidity: event KeyRemovalChargeSet(uint256 amount)
func (_CSModule *CSModuleFilterer) FilterKeyRemovalChargeSet(opts *bind.FilterOpts) (*CSModuleKeyRemovalChargeSetIterator, error) {

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "KeyRemovalChargeSet")
	if err != nil {
		return nil, err
	}
	return &CSModuleKeyRemovalChargeSetIterator{contract: _CSModule.contract, event: "KeyRemovalChargeSet", logs: logs, sub: sub}, nil
}

// WatchKeyRemovalChargeSet is a free log subscription operation binding the contract event 0x699ec9c671aad1f3dcc15e571375584a1d6fb7176afd545d14467fd31477e98e.
//
// Solidity: event KeyRemovalChargeSet(uint256 amount)
func (_CSModule *CSModuleFilterer) WatchKeyRemovalChargeSet(opts *bind.WatchOpts, sink chan<- *CSModuleKeyRemovalChargeSet) (event.Subscription, error) {

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "KeyRemovalChargeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleKeyRemovalChargeSet)
				if err := _CSModule.contract.UnpackLog(event, "KeyRemovalChargeSet", log); err != nil {
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

// ParseKeyRemovalChargeSet is a log parse operation binding the contract event 0x699ec9c671aad1f3dcc15e571375584a1d6fb7176afd545d14467fd31477e98e.
//
// Solidity: event KeyRemovalChargeSet(uint256 amount)
func (_CSModule *CSModuleFilterer) ParseKeyRemovalChargeSet(log types.Log) (*CSModuleKeyRemovalChargeSet, error) {
	event := new(CSModuleKeyRemovalChargeSet)
	if err := _CSModule.contract.UnpackLog(event, "KeyRemovalChargeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleNodeOperatorAddedIterator is returned from FilterNodeOperatorAdded and is used to iterate over the raw logs and unpacked data for NodeOperatorAdded events raised by the CSModule contract.
type CSModuleNodeOperatorAddedIterator struct {
	Event *CSModuleNodeOperatorAdded // Event containing the contract specifics and raw log

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
func (it *CSModuleNodeOperatorAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleNodeOperatorAdded)
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
		it.Event = new(CSModuleNodeOperatorAdded)
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
func (it *CSModuleNodeOperatorAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleNodeOperatorAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleNodeOperatorAdded represents a NodeOperatorAdded event raised by the CSModule contract.
type CSModuleNodeOperatorAdded struct {
	NodeOperatorId *big.Int
	ManagerAddress common.Address
	RewardAddress  common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorAdded is a free log retrieval operation binding the contract event 0xf35982c84fdc94f58d48e901c54c615804cf7d7939b9b8f76ce4d459354e6363.
//
// Solidity: event NodeOperatorAdded(uint256 indexed nodeOperatorId, address indexed managerAddress, address indexed rewardAddress)
func (_CSModule *CSModuleFilterer) FilterNodeOperatorAdded(opts *bind.FilterOpts, nodeOperatorId []*big.Int, managerAddress []common.Address, rewardAddress []common.Address) (*CSModuleNodeOperatorAddedIterator, error) {

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

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "NodeOperatorAdded", nodeOperatorIdRule, managerAddressRule, rewardAddressRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleNodeOperatorAddedIterator{contract: _CSModule.contract, event: "NodeOperatorAdded", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorAdded is a free log subscription operation binding the contract event 0xf35982c84fdc94f58d48e901c54c615804cf7d7939b9b8f76ce4d459354e6363.
//
// Solidity: event NodeOperatorAdded(uint256 indexed nodeOperatorId, address indexed managerAddress, address indexed rewardAddress)
func (_CSModule *CSModuleFilterer) WatchNodeOperatorAdded(opts *bind.WatchOpts, sink chan<- *CSModuleNodeOperatorAdded, nodeOperatorId []*big.Int, managerAddress []common.Address, rewardAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "NodeOperatorAdded", nodeOperatorIdRule, managerAddressRule, rewardAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleNodeOperatorAdded)
				if err := _CSModule.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
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

// ParseNodeOperatorAdded is a log parse operation binding the contract event 0xf35982c84fdc94f58d48e901c54c615804cf7d7939b9b8f76ce4d459354e6363.
//
// Solidity: event NodeOperatorAdded(uint256 indexed nodeOperatorId, address indexed managerAddress, address indexed rewardAddress)
func (_CSModule *CSModuleFilterer) ParseNodeOperatorAdded(log types.Log) (*CSModuleNodeOperatorAdded, error) {
	event := new(CSModuleNodeOperatorAdded)
	if err := _CSModule.contract.UnpackLog(event, "NodeOperatorAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleNodeOperatorManagerAddressChangeProposedIterator is returned from FilterNodeOperatorManagerAddressChangeProposed and is used to iterate over the raw logs and unpacked data for NodeOperatorManagerAddressChangeProposed events raised by the CSModule contract.
type CSModuleNodeOperatorManagerAddressChangeProposedIterator struct {
	Event *CSModuleNodeOperatorManagerAddressChangeProposed // Event containing the contract specifics and raw log

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
func (it *CSModuleNodeOperatorManagerAddressChangeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleNodeOperatorManagerAddressChangeProposed)
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
		it.Event = new(CSModuleNodeOperatorManagerAddressChangeProposed)
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
func (it *CSModuleNodeOperatorManagerAddressChangeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleNodeOperatorManagerAddressChangeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleNodeOperatorManagerAddressChangeProposed represents a NodeOperatorManagerAddressChangeProposed event raised by the CSModule contract.
type CSModuleNodeOperatorManagerAddressChangeProposed struct {
	NodeOperatorId     *big.Int
	OldProposedAddress common.Address
	NewProposedAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorManagerAddressChangeProposed is a free log retrieval operation binding the contract event 0x4048f15a706950765ca59f99d0fa6fe8edaaa3f3e3d0337417082e2131df82fb.
//
// Solidity: event NodeOperatorManagerAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_CSModule *CSModuleFilterer) FilterNodeOperatorManagerAddressChangeProposed(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (*CSModuleNodeOperatorManagerAddressChangeProposedIterator, error) {

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

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "NodeOperatorManagerAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleNodeOperatorManagerAddressChangeProposedIterator{contract: _CSModule.contract, event: "NodeOperatorManagerAddressChangeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorManagerAddressChangeProposed is a free log subscription operation binding the contract event 0x4048f15a706950765ca59f99d0fa6fe8edaaa3f3e3d0337417082e2131df82fb.
//
// Solidity: event NodeOperatorManagerAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_CSModule *CSModuleFilterer) WatchNodeOperatorManagerAddressChangeProposed(opts *bind.WatchOpts, sink chan<- *CSModuleNodeOperatorManagerAddressChangeProposed, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "NodeOperatorManagerAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleNodeOperatorManagerAddressChangeProposed)
				if err := _CSModule.contract.UnpackLog(event, "NodeOperatorManagerAddressChangeProposed", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseNodeOperatorManagerAddressChangeProposed(log types.Log) (*CSModuleNodeOperatorManagerAddressChangeProposed, error) {
	event := new(CSModuleNodeOperatorManagerAddressChangeProposed)
	if err := _CSModule.contract.UnpackLog(event, "NodeOperatorManagerAddressChangeProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleNodeOperatorManagerAddressChangedIterator is returned from FilterNodeOperatorManagerAddressChanged and is used to iterate over the raw logs and unpacked data for NodeOperatorManagerAddressChanged events raised by the CSModule contract.
type CSModuleNodeOperatorManagerAddressChangedIterator struct {
	Event *CSModuleNodeOperatorManagerAddressChanged // Event containing the contract specifics and raw log

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
func (it *CSModuleNodeOperatorManagerAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleNodeOperatorManagerAddressChanged)
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
		it.Event = new(CSModuleNodeOperatorManagerAddressChanged)
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
func (it *CSModuleNodeOperatorManagerAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleNodeOperatorManagerAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleNodeOperatorManagerAddressChanged represents a NodeOperatorManagerAddressChanged event raised by the CSModule contract.
type CSModuleNodeOperatorManagerAddressChanged struct {
	NodeOperatorId *big.Int
	OldAddress     common.Address
	NewAddress     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorManagerAddressChanged is a free log retrieval operation binding the contract event 0x862021f23449d6e8516867bd839be15a3d8698a7561c5c2c35069074b7e91e61.
//
// Solidity: event NodeOperatorManagerAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_CSModule *CSModuleFilterer) FilterNodeOperatorManagerAddressChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (*CSModuleNodeOperatorManagerAddressChangedIterator, error) {

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

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "NodeOperatorManagerAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleNodeOperatorManagerAddressChangedIterator{contract: _CSModule.contract, event: "NodeOperatorManagerAddressChanged", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorManagerAddressChanged is a free log subscription operation binding the contract event 0x862021f23449d6e8516867bd839be15a3d8698a7561c5c2c35069074b7e91e61.
//
// Solidity: event NodeOperatorManagerAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_CSModule *CSModuleFilterer) WatchNodeOperatorManagerAddressChanged(opts *bind.WatchOpts, sink chan<- *CSModuleNodeOperatorManagerAddressChanged, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "NodeOperatorManagerAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleNodeOperatorManagerAddressChanged)
				if err := _CSModule.contract.UnpackLog(event, "NodeOperatorManagerAddressChanged", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseNodeOperatorManagerAddressChanged(log types.Log) (*CSModuleNodeOperatorManagerAddressChanged, error) {
	event := new(CSModuleNodeOperatorManagerAddressChanged)
	if err := _CSModule.contract.UnpackLog(event, "NodeOperatorManagerAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleNodeOperatorRewardAddressChangeProposedIterator is returned from FilterNodeOperatorRewardAddressChangeProposed and is used to iterate over the raw logs and unpacked data for NodeOperatorRewardAddressChangeProposed events raised by the CSModule contract.
type CSModuleNodeOperatorRewardAddressChangeProposedIterator struct {
	Event *CSModuleNodeOperatorRewardAddressChangeProposed // Event containing the contract specifics and raw log

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
func (it *CSModuleNodeOperatorRewardAddressChangeProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleNodeOperatorRewardAddressChangeProposed)
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
		it.Event = new(CSModuleNodeOperatorRewardAddressChangeProposed)
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
func (it *CSModuleNodeOperatorRewardAddressChangeProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleNodeOperatorRewardAddressChangeProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleNodeOperatorRewardAddressChangeProposed represents a NodeOperatorRewardAddressChangeProposed event raised by the CSModule contract.
type CSModuleNodeOperatorRewardAddressChangeProposed struct {
	NodeOperatorId     *big.Int
	OldProposedAddress common.Address
	NewProposedAddress common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRewardAddressChangeProposed is a free log retrieval operation binding the contract event 0xb5878cdb1d66f971efe3b138a71c64bc5bc519314db2533e0e4cde954409ea5a.
//
// Solidity: event NodeOperatorRewardAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_CSModule *CSModuleFilterer) FilterNodeOperatorRewardAddressChangeProposed(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (*CSModuleNodeOperatorRewardAddressChangeProposedIterator, error) {

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

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "NodeOperatorRewardAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleNodeOperatorRewardAddressChangeProposedIterator{contract: _CSModule.contract, event: "NodeOperatorRewardAddressChangeProposed", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRewardAddressChangeProposed is a free log subscription operation binding the contract event 0xb5878cdb1d66f971efe3b138a71c64bc5bc519314db2533e0e4cde954409ea5a.
//
// Solidity: event NodeOperatorRewardAddressChangeProposed(uint256 indexed nodeOperatorId, address indexed oldProposedAddress, address indexed newProposedAddress)
func (_CSModule *CSModuleFilterer) WatchNodeOperatorRewardAddressChangeProposed(opts *bind.WatchOpts, sink chan<- *CSModuleNodeOperatorRewardAddressChangeProposed, nodeOperatorId []*big.Int, oldProposedAddress []common.Address, newProposedAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "NodeOperatorRewardAddressChangeProposed", nodeOperatorIdRule, oldProposedAddressRule, newProposedAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleNodeOperatorRewardAddressChangeProposed)
				if err := _CSModule.contract.UnpackLog(event, "NodeOperatorRewardAddressChangeProposed", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseNodeOperatorRewardAddressChangeProposed(log types.Log) (*CSModuleNodeOperatorRewardAddressChangeProposed, error) {
	event := new(CSModuleNodeOperatorRewardAddressChangeProposed)
	if err := _CSModule.contract.UnpackLog(event, "NodeOperatorRewardAddressChangeProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleNodeOperatorRewardAddressChangedIterator is returned from FilterNodeOperatorRewardAddressChanged and is used to iterate over the raw logs and unpacked data for NodeOperatorRewardAddressChanged events raised by the CSModule contract.
type CSModuleNodeOperatorRewardAddressChangedIterator struct {
	Event *CSModuleNodeOperatorRewardAddressChanged // Event containing the contract specifics and raw log

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
func (it *CSModuleNodeOperatorRewardAddressChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleNodeOperatorRewardAddressChanged)
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
		it.Event = new(CSModuleNodeOperatorRewardAddressChanged)
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
func (it *CSModuleNodeOperatorRewardAddressChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleNodeOperatorRewardAddressChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleNodeOperatorRewardAddressChanged represents a NodeOperatorRewardAddressChanged event raised by the CSModule contract.
type CSModuleNodeOperatorRewardAddressChanged struct {
	NodeOperatorId *big.Int
	OldAddress     common.Address
	NewAddress     common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNodeOperatorRewardAddressChanged is a free log retrieval operation binding the contract event 0x069ac7cd8230db015b7250c8e5425149cf1a3e912d9569f497165e55b3b6b7b2.
//
// Solidity: event NodeOperatorRewardAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_CSModule *CSModuleFilterer) FilterNodeOperatorRewardAddressChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (*CSModuleNodeOperatorRewardAddressChangedIterator, error) {

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

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "NodeOperatorRewardAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleNodeOperatorRewardAddressChangedIterator{contract: _CSModule.contract, event: "NodeOperatorRewardAddressChanged", logs: logs, sub: sub}, nil
}

// WatchNodeOperatorRewardAddressChanged is a free log subscription operation binding the contract event 0x069ac7cd8230db015b7250c8e5425149cf1a3e912d9569f497165e55b3b6b7b2.
//
// Solidity: event NodeOperatorRewardAddressChanged(uint256 indexed nodeOperatorId, address indexed oldAddress, address indexed newAddress)
func (_CSModule *CSModuleFilterer) WatchNodeOperatorRewardAddressChanged(opts *bind.WatchOpts, sink chan<- *CSModuleNodeOperatorRewardAddressChanged, nodeOperatorId []*big.Int, oldAddress []common.Address, newAddress []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "NodeOperatorRewardAddressChanged", nodeOperatorIdRule, oldAddressRule, newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleNodeOperatorRewardAddressChanged)
				if err := _CSModule.contract.UnpackLog(event, "NodeOperatorRewardAddressChanged", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseNodeOperatorRewardAddressChanged(log types.Log) (*CSModuleNodeOperatorRewardAddressChanged, error) {
	event := new(CSModuleNodeOperatorRewardAddressChanged)
	if err := _CSModule.contract.UnpackLog(event, "NodeOperatorRewardAddressChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleNonceChangedIterator is returned from FilterNonceChanged and is used to iterate over the raw logs and unpacked data for NonceChanged events raised by the CSModule contract.
type CSModuleNonceChangedIterator struct {
	Event *CSModuleNonceChanged // Event containing the contract specifics and raw log

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
func (it *CSModuleNonceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleNonceChanged)
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
		it.Event = new(CSModuleNonceChanged)
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
func (it *CSModuleNonceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleNonceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleNonceChanged represents a NonceChanged event raised by the CSModule contract.
type CSModuleNonceChanged struct {
	Nonce *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterNonceChanged is a free log retrieval operation binding the contract event 0x7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa.
//
// Solidity: event NonceChanged(uint256 nonce)
func (_CSModule *CSModuleFilterer) FilterNonceChanged(opts *bind.FilterOpts) (*CSModuleNonceChangedIterator, error) {

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "NonceChanged")
	if err != nil {
		return nil, err
	}
	return &CSModuleNonceChangedIterator{contract: _CSModule.contract, event: "NonceChanged", logs: logs, sub: sub}, nil
}

// WatchNonceChanged is a free log subscription operation binding the contract event 0x7220970e1f1f12864ecccd8942690a837c7a8dd45d158cb891eb45a8a69134aa.
//
// Solidity: event NonceChanged(uint256 nonce)
func (_CSModule *CSModuleFilterer) WatchNonceChanged(opts *bind.WatchOpts, sink chan<- *CSModuleNonceChanged) (event.Subscription, error) {

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "NonceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleNonceChanged)
				if err := _CSModule.contract.UnpackLog(event, "NonceChanged", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseNonceChanged(log types.Log) (*CSModuleNonceChanged, error) {
	event := new(CSModuleNonceChanged)
	if err := _CSModule.contract.UnpackLog(event, "NonceChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModulePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CSModule contract.
type CSModulePausedIterator struct {
	Event *CSModulePaused // Event containing the contract specifics and raw log

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
func (it *CSModulePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModulePaused)
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
		it.Event = new(CSModulePaused)
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
func (it *CSModulePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModulePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModulePaused represents a Paused event raised by the CSModule contract.
type CSModulePaused struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_CSModule *CSModuleFilterer) FilterPaused(opts *bind.FilterOpts) (*CSModulePausedIterator, error) {

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CSModulePausedIterator{contract: _CSModule.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_CSModule *CSModuleFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CSModulePaused) (event.Subscription, error) {

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModulePaused)
				if err := _CSModule.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParsePaused(log types.Log) (*CSModulePaused, error) {
	event := new(CSModulePaused)
	if err := _CSModule.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModulePublicReleaseIterator is returned from FilterPublicRelease and is used to iterate over the raw logs and unpacked data for PublicRelease events raised by the CSModule contract.
type CSModulePublicReleaseIterator struct {
	Event *CSModulePublicRelease // Event containing the contract specifics and raw log

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
func (it *CSModulePublicReleaseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModulePublicRelease)
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
		it.Event = new(CSModulePublicRelease)
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
func (it *CSModulePublicReleaseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModulePublicReleaseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModulePublicRelease represents a PublicRelease event raised by the CSModule contract.
type CSModulePublicRelease struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPublicRelease is a free log retrieval operation binding the contract event 0xe5eb57aa4d841adeece4ac87bd294965df4a894f0aa24db4a4a55a39ab101d6e.
//
// Solidity: event PublicRelease()
func (_CSModule *CSModuleFilterer) FilterPublicRelease(opts *bind.FilterOpts) (*CSModulePublicReleaseIterator, error) {

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "PublicRelease")
	if err != nil {
		return nil, err
	}
	return &CSModulePublicReleaseIterator{contract: _CSModule.contract, event: "PublicRelease", logs: logs, sub: sub}, nil
}

// WatchPublicRelease is a free log subscription operation binding the contract event 0xe5eb57aa4d841adeece4ac87bd294965df4a894f0aa24db4a4a55a39ab101d6e.
//
// Solidity: event PublicRelease()
func (_CSModule *CSModuleFilterer) WatchPublicRelease(opts *bind.WatchOpts, sink chan<- *CSModulePublicRelease) (event.Subscription, error) {

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "PublicRelease")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModulePublicRelease)
				if err := _CSModule.contract.UnpackLog(event, "PublicRelease", log); err != nil {
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

// ParsePublicRelease is a log parse operation binding the contract event 0xe5eb57aa4d841adeece4ac87bd294965df4a894f0aa24db4a4a55a39ab101d6e.
//
// Solidity: event PublicRelease()
func (_CSModule *CSModuleFilterer) ParsePublicRelease(log types.Log) (*CSModulePublicRelease, error) {
	event := new(CSModulePublicRelease)
	if err := _CSModule.contract.UnpackLog(event, "PublicRelease", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleReferrerSetIterator is returned from FilterReferrerSet and is used to iterate over the raw logs and unpacked data for ReferrerSet events raised by the CSModule contract.
type CSModuleReferrerSetIterator struct {
	Event *CSModuleReferrerSet // Event containing the contract specifics and raw log

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
func (it *CSModuleReferrerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleReferrerSet)
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
		it.Event = new(CSModuleReferrerSet)
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
func (it *CSModuleReferrerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleReferrerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleReferrerSet represents a ReferrerSet event raised by the CSModule contract.
type CSModuleReferrerSet struct {
	NodeOperatorId *big.Int
	Referrer       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterReferrerSet is a free log retrieval operation binding the contract event 0x67334334c388385e5f244703f8a8b28b7f4ffe52909130aca69bc62a8e27f09a.
//
// Solidity: event ReferrerSet(uint256 indexed nodeOperatorId, address indexed referrer)
func (_CSModule *CSModuleFilterer) FilterReferrerSet(opts *bind.FilterOpts, nodeOperatorId []*big.Int, referrer []common.Address) (*CSModuleReferrerSetIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "ReferrerSet", nodeOperatorIdRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleReferrerSetIterator{contract: _CSModule.contract, event: "ReferrerSet", logs: logs, sub: sub}, nil
}

// WatchReferrerSet is a free log subscription operation binding the contract event 0x67334334c388385e5f244703f8a8b28b7f4ffe52909130aca69bc62a8e27f09a.
//
// Solidity: event ReferrerSet(uint256 indexed nodeOperatorId, address indexed referrer)
func (_CSModule *CSModuleFilterer) WatchReferrerSet(opts *bind.WatchOpts, sink chan<- *CSModuleReferrerSet, nodeOperatorId []*big.Int, referrer []common.Address) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "ReferrerSet", nodeOperatorIdRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleReferrerSet)
				if err := _CSModule.contract.UnpackLog(event, "ReferrerSet", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseReferrerSet(log types.Log) (*CSModuleReferrerSet, error) {
	event := new(CSModuleReferrerSet)
	if err := _CSModule.contract.UnpackLog(event, "ReferrerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the CSModule contract.
type CSModuleResumedIterator struct {
	Event *CSModuleResumed // Event containing the contract specifics and raw log

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
func (it *CSModuleResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleResumed)
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
		it.Event = new(CSModuleResumed)
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
func (it *CSModuleResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleResumed represents a Resumed event raised by the CSModule contract.
type CSModuleResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_CSModule *CSModuleFilterer) FilterResumed(opts *bind.FilterOpts) (*CSModuleResumedIterator, error) {

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &CSModuleResumedIterator{contract: _CSModule.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_CSModule *CSModuleFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *CSModuleResumed) (event.Subscription, error) {

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleResumed)
				if err := _CSModule.contract.UnpackLog(event, "Resumed", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseResumed(log types.Log) (*CSModuleResumed, error) {
	event := new(CSModuleResumed)
	if err := _CSModule.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CSModule contract.
type CSModuleRoleAdminChangedIterator struct {
	Event *CSModuleRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CSModuleRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleRoleAdminChanged)
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
		it.Event = new(CSModuleRoleAdminChanged)
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
func (it *CSModuleRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleRoleAdminChanged represents a RoleAdminChanged event raised by the CSModule contract.
type CSModuleRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CSModule *CSModuleFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CSModuleRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleRoleAdminChangedIterator{contract: _CSModule.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CSModule *CSModuleFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CSModuleRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleRoleAdminChanged)
				if err := _CSModule.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseRoleAdminChanged(log types.Log) (*CSModuleRoleAdminChanged, error) {
	event := new(CSModuleRoleAdminChanged)
	if err := _CSModule.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CSModule contract.
type CSModuleRoleGrantedIterator struct {
	Event *CSModuleRoleGranted // Event containing the contract specifics and raw log

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
func (it *CSModuleRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleRoleGranted)
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
		it.Event = new(CSModuleRoleGranted)
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
func (it *CSModuleRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleRoleGranted represents a RoleGranted event raised by the CSModule contract.
type CSModuleRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CSModule *CSModuleFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CSModuleRoleGrantedIterator, error) {

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

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleRoleGrantedIterator{contract: _CSModule.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CSModule *CSModuleFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CSModuleRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleRoleGranted)
				if err := _CSModule.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseRoleGranted(log types.Log) (*CSModuleRoleGranted, error) {
	event := new(CSModuleRoleGranted)
	if err := _CSModule.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CSModule contract.
type CSModuleRoleRevokedIterator struct {
	Event *CSModuleRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CSModuleRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleRoleRevoked)
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
		it.Event = new(CSModuleRoleRevoked)
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
func (it *CSModuleRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleRoleRevoked represents a RoleRevoked event raised by the CSModule contract.
type CSModuleRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CSModule *CSModuleFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CSModuleRoleRevokedIterator, error) {

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

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleRoleRevokedIterator{contract: _CSModule.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CSModule *CSModuleFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CSModuleRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleRoleRevoked)
				if err := _CSModule.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseRoleRevoked(log types.Log) (*CSModuleRoleRevoked, error) {
	event := new(CSModuleRoleRevoked)
	if err := _CSModule.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleSigningKeyAddedIterator is returned from FilterSigningKeyAdded and is used to iterate over the raw logs and unpacked data for SigningKeyAdded events raised by the CSModule contract.
type CSModuleSigningKeyAddedIterator struct {
	Event *CSModuleSigningKeyAdded // Event containing the contract specifics and raw log

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
func (it *CSModuleSigningKeyAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleSigningKeyAdded)
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
		it.Event = new(CSModuleSigningKeyAdded)
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
func (it *CSModuleSigningKeyAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleSigningKeyAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleSigningKeyAdded represents a SigningKeyAdded event raised by the CSModule contract.
type CSModuleSigningKeyAdded struct {
	NodeOperatorId *big.Int
	Pubkey         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSigningKeyAdded is a free log retrieval operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed nodeOperatorId, bytes pubkey)
func (_CSModule *CSModuleFilterer) FilterSigningKeyAdded(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleSigningKeyAddedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "SigningKeyAdded", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleSigningKeyAddedIterator{contract: _CSModule.contract, event: "SigningKeyAdded", logs: logs, sub: sub}, nil
}

// WatchSigningKeyAdded is a free log subscription operation binding the contract event 0xc77a17d6b857abe6d6e6c37301621bc72c4dd52fa8830fb54dfa715c04911a89.
//
// Solidity: event SigningKeyAdded(uint256 indexed nodeOperatorId, bytes pubkey)
func (_CSModule *CSModuleFilterer) WatchSigningKeyAdded(opts *bind.WatchOpts, sink chan<- *CSModuleSigningKeyAdded, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "SigningKeyAdded", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleSigningKeyAdded)
				if err := _CSModule.contract.UnpackLog(event, "SigningKeyAdded", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseSigningKeyAdded(log types.Log) (*CSModuleSigningKeyAdded, error) {
	event := new(CSModuleSigningKeyAdded)
	if err := _CSModule.contract.UnpackLog(event, "SigningKeyAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleSigningKeyRemovedIterator is returned from FilterSigningKeyRemoved and is used to iterate over the raw logs and unpacked data for SigningKeyRemoved events raised by the CSModule contract.
type CSModuleSigningKeyRemovedIterator struct {
	Event *CSModuleSigningKeyRemoved // Event containing the contract specifics and raw log

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
func (it *CSModuleSigningKeyRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleSigningKeyRemoved)
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
		it.Event = new(CSModuleSigningKeyRemoved)
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
func (it *CSModuleSigningKeyRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleSigningKeyRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleSigningKeyRemoved represents a SigningKeyRemoved event raised by the CSModule contract.
type CSModuleSigningKeyRemoved struct {
	NodeOperatorId *big.Int
	Pubkey         []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterSigningKeyRemoved is a free log retrieval operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed nodeOperatorId, bytes pubkey)
func (_CSModule *CSModuleFilterer) FilterSigningKeyRemoved(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleSigningKeyRemovedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "SigningKeyRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleSigningKeyRemovedIterator{contract: _CSModule.contract, event: "SigningKeyRemoved", logs: logs, sub: sub}, nil
}

// WatchSigningKeyRemoved is a free log subscription operation binding the contract event 0xea4b75aaf57196f73d338cadf79ecd0a437902e2dd0d2c4c2cf3ea71b8ab27b9.
//
// Solidity: event SigningKeyRemoved(uint256 indexed nodeOperatorId, bytes pubkey)
func (_CSModule *CSModuleFilterer) WatchSigningKeyRemoved(opts *bind.WatchOpts, sink chan<- *CSModuleSigningKeyRemoved, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "SigningKeyRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleSigningKeyRemoved)
				if err := _CSModule.contract.UnpackLog(event, "SigningKeyRemoved", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseSigningKeyRemoved(log types.Log) (*CSModuleSigningKeyRemoved, error) {
	event := new(CSModuleSigningKeyRemoved)
	if err := _CSModule.contract.UnpackLog(event, "SigningKeyRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleStETHSharesRecoveredIterator is returned from FilterStETHSharesRecovered and is used to iterate over the raw logs and unpacked data for StETHSharesRecovered events raised by the CSModule contract.
type CSModuleStETHSharesRecoveredIterator struct {
	Event *CSModuleStETHSharesRecovered // Event containing the contract specifics and raw log

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
func (it *CSModuleStETHSharesRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleStETHSharesRecovered)
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
		it.Event = new(CSModuleStETHSharesRecovered)
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
func (it *CSModuleStETHSharesRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleStETHSharesRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleStETHSharesRecovered represents a StETHSharesRecovered event raised by the CSModule contract.
type CSModuleStETHSharesRecovered struct {
	Recipient common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStETHSharesRecovered is a free log retrieval operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_CSModule *CSModuleFilterer) FilterStETHSharesRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CSModuleStETHSharesRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleStETHSharesRecoveredIterator{contract: _CSModule.contract, event: "StETHSharesRecovered", logs: logs, sub: sub}, nil
}

// WatchStETHSharesRecovered is a free log subscription operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_CSModule *CSModuleFilterer) WatchStETHSharesRecovered(opts *bind.WatchOpts, sink chan<- *CSModuleStETHSharesRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleStETHSharesRecovered)
				if err := _CSModule.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseStETHSharesRecovered(log types.Log) (*CSModuleStETHSharesRecovered, error) {
	event := new(CSModuleStETHSharesRecovered)
	if err := _CSModule.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleStuckSigningKeysCountChangedIterator is returned from FilterStuckSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for StuckSigningKeysCountChanged events raised by the CSModule contract.
type CSModuleStuckSigningKeysCountChangedIterator struct {
	Event *CSModuleStuckSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CSModuleStuckSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleStuckSigningKeysCountChanged)
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
		it.Event = new(CSModuleStuckSigningKeysCountChanged)
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
func (it *CSModuleStuckSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleStuckSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleStuckSigningKeysCountChanged represents a StuckSigningKeysCountChanged event raised by the CSModule contract.
type CSModuleStuckSigningKeysCountChanged struct {
	NodeOperatorId *big.Int
	StuckKeysCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterStuckSigningKeysCountChanged is a free log retrieval operation binding the contract event 0xb4f5879eca27b32881cec7907d1310378e9b4c79927062fb7d4a321434b5b04a.
//
// Solidity: event StuckSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 stuckKeysCount)
func (_CSModule *CSModuleFilterer) FilterStuckSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleStuckSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "StuckSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleStuckSigningKeysCountChangedIterator{contract: _CSModule.contract, event: "StuckSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchStuckSigningKeysCountChanged is a free log subscription operation binding the contract event 0xb4f5879eca27b32881cec7907d1310378e9b4c79927062fb7d4a321434b5b04a.
//
// Solidity: event StuckSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 stuckKeysCount)
func (_CSModule *CSModuleFilterer) WatchStuckSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CSModuleStuckSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "StuckSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleStuckSigningKeysCountChanged)
				if err := _CSModule.contract.UnpackLog(event, "StuckSigningKeysCountChanged", log); err != nil {
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

// ParseStuckSigningKeysCountChanged is a log parse operation binding the contract event 0xb4f5879eca27b32881cec7907d1310378e9b4c79927062fb7d4a321434b5b04a.
//
// Solidity: event StuckSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 stuckKeysCount)
func (_CSModule *CSModuleFilterer) ParseStuckSigningKeysCountChanged(log types.Log) (*CSModuleStuckSigningKeysCountChanged, error) {
	event := new(CSModuleStuckSigningKeysCountChanged)
	if err := _CSModule.contract.UnpackLog(event, "StuckSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleTargetValidatorsCountChangedIterator is returned from FilterTargetValidatorsCountChanged and is used to iterate over the raw logs and unpacked data for TargetValidatorsCountChanged events raised by the CSModule contract.
type CSModuleTargetValidatorsCountChangedIterator struct {
	Event *CSModuleTargetValidatorsCountChanged // Event containing the contract specifics and raw log

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
func (it *CSModuleTargetValidatorsCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleTargetValidatorsCountChanged)
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
		it.Event = new(CSModuleTargetValidatorsCountChanged)
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
func (it *CSModuleTargetValidatorsCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleTargetValidatorsCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleTargetValidatorsCountChanged represents a TargetValidatorsCountChanged event raised by the CSModule contract.
type CSModuleTargetValidatorsCountChanged struct {
	NodeOperatorId        *big.Int
	TargetLimitMode       *big.Int
	TargetValidatorsCount *big.Int
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterTargetValidatorsCountChanged is a free log retrieval operation binding the contract event 0xf92eb109ce5b449e9b121c352c6aeb4319538a90738cb95d84f08e41274e92d2.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetLimitMode, uint256 targetValidatorsCount)
func (_CSModule *CSModuleFilterer) FilterTargetValidatorsCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleTargetValidatorsCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "TargetValidatorsCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleTargetValidatorsCountChangedIterator{contract: _CSModule.contract, event: "TargetValidatorsCountChanged", logs: logs, sub: sub}, nil
}

// WatchTargetValidatorsCountChanged is a free log subscription operation binding the contract event 0xf92eb109ce5b449e9b121c352c6aeb4319538a90738cb95d84f08e41274e92d2.
//
// Solidity: event TargetValidatorsCountChanged(uint256 indexed nodeOperatorId, uint256 targetLimitMode, uint256 targetValidatorsCount)
func (_CSModule *CSModuleFilterer) WatchTargetValidatorsCountChanged(opts *bind.WatchOpts, sink chan<- *CSModuleTargetValidatorsCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "TargetValidatorsCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleTargetValidatorsCountChanged)
				if err := _CSModule.contract.UnpackLog(event, "TargetValidatorsCountChanged", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseTargetValidatorsCountChanged(log types.Log) (*CSModuleTargetValidatorsCountChanged, error) {
	event := new(CSModuleTargetValidatorsCountChanged)
	if err := _CSModule.contract.UnpackLog(event, "TargetValidatorsCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleTotalSigningKeysCountChangedIterator is returned from FilterTotalSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for TotalSigningKeysCountChanged events raised by the CSModule contract.
type CSModuleTotalSigningKeysCountChangedIterator struct {
	Event *CSModuleTotalSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CSModuleTotalSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleTotalSigningKeysCountChanged)
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
		it.Event = new(CSModuleTotalSigningKeysCountChanged)
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
func (it *CSModuleTotalSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleTotalSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleTotalSigningKeysCountChanged represents a TotalSigningKeysCountChanged event raised by the CSModule contract.
type CSModuleTotalSigningKeysCountChanged struct {
	NodeOperatorId *big.Int
	TotalKeysCount *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTotalSigningKeysCountChanged is a free log retrieval operation binding the contract event 0xdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f0.
//
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalKeysCount)
func (_CSModule *CSModuleFilterer) FilterTotalSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleTotalSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "TotalSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleTotalSigningKeysCountChangedIterator{contract: _CSModule.contract, event: "TotalSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchTotalSigningKeysCountChanged is a free log subscription operation binding the contract event 0xdd01838a366ae4dc9a86e1922512c0716abebc9a440baae0e22d2dec578223f0.
//
// Solidity: event TotalSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 totalKeysCount)
func (_CSModule *CSModuleFilterer) WatchTotalSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CSModuleTotalSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "TotalSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleTotalSigningKeysCountChanged)
				if err := _CSModule.contract.UnpackLog(event, "TotalSigningKeysCountChanged", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseTotalSigningKeysCountChanged(log types.Log) (*CSModuleTotalSigningKeysCountChanged, error) {
	event := new(CSModuleTotalSigningKeysCountChanged)
	if err := _CSModule.contract.UnpackLog(event, "TotalSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleVettedSigningKeysCountChangedIterator is returned from FilterVettedSigningKeysCountChanged and is used to iterate over the raw logs and unpacked data for VettedSigningKeysCountChanged events raised by the CSModule contract.
type CSModuleVettedSigningKeysCountChangedIterator struct {
	Event *CSModuleVettedSigningKeysCountChanged // Event containing the contract specifics and raw log

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
func (it *CSModuleVettedSigningKeysCountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleVettedSigningKeysCountChanged)
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
		it.Event = new(CSModuleVettedSigningKeysCountChanged)
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
func (it *CSModuleVettedSigningKeysCountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleVettedSigningKeysCountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleVettedSigningKeysCountChanged represents a VettedSigningKeysCountChanged event raised by the CSModule contract.
type CSModuleVettedSigningKeysCountChanged struct {
	NodeOperatorId  *big.Int
	VettedKeysCount *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVettedSigningKeysCountChanged is a free log retrieval operation binding the contract event 0x947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd.
//
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 vettedKeysCount)
func (_CSModule *CSModuleFilterer) FilterVettedSigningKeysCountChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleVettedSigningKeysCountChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "VettedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleVettedSigningKeysCountChangedIterator{contract: _CSModule.contract, event: "VettedSigningKeysCountChanged", logs: logs, sub: sub}, nil
}

// WatchVettedSigningKeysCountChanged is a free log subscription operation binding the contract event 0x947f955eec7e1f626bee3afd2aa47b5de04ddcdd3fe78dc8838213015ef58dfd.
//
// Solidity: event VettedSigningKeysCountChanged(uint256 indexed nodeOperatorId, uint256 vettedKeysCount)
func (_CSModule *CSModuleFilterer) WatchVettedSigningKeysCountChanged(opts *bind.WatchOpts, sink chan<- *CSModuleVettedSigningKeysCountChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "VettedSigningKeysCountChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleVettedSigningKeysCountChanged)
				if err := _CSModule.contract.UnpackLog(event, "VettedSigningKeysCountChanged", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseVettedSigningKeysCountChanged(log types.Log) (*CSModuleVettedSigningKeysCountChanged, error) {
	event := new(CSModuleVettedSigningKeysCountChanged)
	if err := _CSModule.contract.UnpackLog(event, "VettedSigningKeysCountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleVettedSigningKeysCountDecreasedIterator is returned from FilterVettedSigningKeysCountDecreased and is used to iterate over the raw logs and unpacked data for VettedSigningKeysCountDecreased events raised by the CSModule contract.
type CSModuleVettedSigningKeysCountDecreasedIterator struct {
	Event *CSModuleVettedSigningKeysCountDecreased // Event containing the contract specifics and raw log

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
func (it *CSModuleVettedSigningKeysCountDecreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleVettedSigningKeysCountDecreased)
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
		it.Event = new(CSModuleVettedSigningKeysCountDecreased)
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
func (it *CSModuleVettedSigningKeysCountDecreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleVettedSigningKeysCountDecreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleVettedSigningKeysCountDecreased represents a VettedSigningKeysCountDecreased event raised by the CSModule contract.
type CSModuleVettedSigningKeysCountDecreased struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVettedSigningKeysCountDecreased is a free log retrieval operation binding the contract event 0xe5725d045d5c47bd1483feba445e395dc8647486963e6d54aad9ed03ff7d6ce6.
//
// Solidity: event VettedSigningKeysCountDecreased(uint256 indexed nodeOperatorId)
func (_CSModule *CSModuleFilterer) FilterVettedSigningKeysCountDecreased(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleVettedSigningKeysCountDecreasedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "VettedSigningKeysCountDecreased", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleVettedSigningKeysCountDecreasedIterator{contract: _CSModule.contract, event: "VettedSigningKeysCountDecreased", logs: logs, sub: sub}, nil
}

// WatchVettedSigningKeysCountDecreased is a free log subscription operation binding the contract event 0xe5725d045d5c47bd1483feba445e395dc8647486963e6d54aad9ed03ff7d6ce6.
//
// Solidity: event VettedSigningKeysCountDecreased(uint256 indexed nodeOperatorId)
func (_CSModule *CSModuleFilterer) WatchVettedSigningKeysCountDecreased(opts *bind.WatchOpts, sink chan<- *CSModuleVettedSigningKeysCountDecreased, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "VettedSigningKeysCountDecreased", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleVettedSigningKeysCountDecreased)
				if err := _CSModule.contract.UnpackLog(event, "VettedSigningKeysCountDecreased", log); err != nil {
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
func (_CSModule *CSModuleFilterer) ParseVettedSigningKeysCountDecreased(log types.Log) (*CSModuleVettedSigningKeysCountDecreased, error) {
	event := new(CSModuleVettedSigningKeysCountDecreased)
	if err := _CSModule.contract.UnpackLog(event, "VettedSigningKeysCountDecreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSModuleWithdrawalSubmittedIterator is returned from FilterWithdrawalSubmitted and is used to iterate over the raw logs and unpacked data for WithdrawalSubmitted events raised by the CSModule contract.
type CSModuleWithdrawalSubmittedIterator struct {
	Event *CSModuleWithdrawalSubmitted // Event containing the contract specifics and raw log

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
func (it *CSModuleWithdrawalSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSModuleWithdrawalSubmitted)
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
		it.Event = new(CSModuleWithdrawalSubmitted)
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
func (it *CSModuleWithdrawalSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSModuleWithdrawalSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSModuleWithdrawalSubmitted represents a WithdrawalSubmitted event raised by the CSModule contract.
type CSModuleWithdrawalSubmitted struct {
	NodeOperatorId *big.Int
	KeyIndex       *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterWithdrawalSubmitted is a free log retrieval operation binding the contract event 0xcb2f99f65711a7d6df7f552255b910bf59f09fcd5935f44c170b4cb0d1b50995.
//
// Solidity: event WithdrawalSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex, uint256 amount)
func (_CSModule *CSModuleFilterer) FilterWithdrawalSubmitted(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSModuleWithdrawalSubmittedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.FilterLogs(opts, "WithdrawalSubmitted", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSModuleWithdrawalSubmittedIterator{contract: _CSModule.contract, event: "WithdrawalSubmitted", logs: logs, sub: sub}, nil
}

// WatchWithdrawalSubmitted is a free log subscription operation binding the contract event 0xcb2f99f65711a7d6df7f552255b910bf59f09fcd5935f44c170b4cb0d1b50995.
//
// Solidity: event WithdrawalSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex, uint256 amount)
func (_CSModule *CSModuleFilterer) WatchWithdrawalSubmitted(opts *bind.WatchOpts, sink chan<- *CSModuleWithdrawalSubmitted, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSModule.contract.WatchLogs(opts, "WithdrawalSubmitted", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSModuleWithdrawalSubmitted)
				if err := _CSModule.contract.UnpackLog(event, "WithdrawalSubmitted", log); err != nil {
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

// ParseWithdrawalSubmitted is a log parse operation binding the contract event 0xcb2f99f65711a7d6df7f552255b910bf59f09fcd5935f44c170b4cb0d1b50995.
//
// Solidity: event WithdrawalSubmitted(uint256 indexed nodeOperatorId, uint256 keyIndex, uint256 amount)
func (_CSModule *CSModuleFilterer) ParseWithdrawalSubmitted(log types.Log) (*CSModuleWithdrawalSubmitted, error) {
	event := new(CSModuleWithdrawalSubmitted)
	if err := _CSModule.contract.UnpackLog(event, "WithdrawalSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
