// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package csaccounting

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

// ICSBondCurveBondCurve is an auto generated low-level Go binding around an user-defined struct.
type ICSBondCurveBondCurve struct {
	Points []*big.Int
	Trend  *big.Int
}

// ICSBondLockBondLock is an auto generated low-level Go binding around an user-defined struct.
type ICSBondLockBondLock struct {
	Amount         *big.Int
	RetentionUntil *big.Int
}

// CSAccountingMetaData contains all meta data concerning the CSAccounting contract.
var CSAccountingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lidoLocator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"communityStakingModule\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxCurveLength\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBondLockRetentionPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBondLockRetentionPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ElRewardsVaultReceiveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondCurveId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondCurveLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondCurveMaxLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondCurveValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondLockAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondLockRetentionPeriod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialisationCurveId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeOperatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToRecover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToClaim\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseUntilMustBeInFuture\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedExpected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResumedExpected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotCSM\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroChargePenaltyRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFeeDistributorAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroLocatorAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroModuleAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPauseDuration\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toBurnAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"}],\"name\":\"BondBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChargeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chargedAmount\",\"type\":\"uint256\"}],\"name\":\"BondCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondClaimedStETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"BondClaimedUnstETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondClaimedWstETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"bondCurve\",\"type\":\"uint256[]\"}],\"name\":\"BondCurveAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"BondCurveSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"bondCurve\",\"type\":\"uint256[]\"}],\"name\":\"BondCurveUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondDepositedETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondDepositedStETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondDepositedWstETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"retentionUntil\",\"type\":\"uint256\"}],\"name\":\"BondLockChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondLockCompensated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"BondLockRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"retentionPeriod\",\"type\":\"uint256\"}],\"name\":\"BondLockRetentionPeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"chargePenaltyRecipient\",\"type\":\"address\"}],\"name\":\"ChargePenaltyRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"StETHSharesRecovered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ACCOUNTING_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CSM\",\"outputs\":[{\"internalType\":\"contractICSModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_BOND_CURVE_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIDO\",\"outputs\":[{\"internalType\":\"contractILido\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIDO_LOCATOR\",\"outputs\":[{\"internalType\":\"contractILidoLocator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_BOND_CURVES_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BOND_LOCK_RETENTION_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CURVE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BOND_LOCK_RETENTION_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_CURVE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_INFINITELY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECOVERER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESET_BOND_CURVE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_BOND_CURVE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_QUEUE\",\"outputs\":[{\"internalType\":\"contractIWithdrawalQueue\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WSTETH\",\"outputs\":[{\"internalType\":\"contractIWstETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"bondCurve\",\"type\":\"uint256[]\"}],\"name\":\"addBondCurve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"chargeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chargePenaltyRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stETHAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimRewardsStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stEthAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimRewardsUnstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wstETHAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"rewardAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimRewardsWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"compensateLockedBondETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stETHAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"depositStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wstETHAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"depositWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDistributor\",\"outputs\":[{\"internalType\":\"contractICSFeeDistributor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getActualLockedBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keys\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"points\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"internalType\":\"structICSBondCurve.BondCurve\",\"name\":\"curve\",\"type\":\"tuple\"}],\"name\":\"getBondAmountByKeysCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keys\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getBondAmountByKeysCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getBondAmountByKeysCountWstETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"points\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"internalType\":\"structICSBondCurve.BondCurve\",\"name\":\"curve\",\"type\":\"tuple\"}],\"name\":\"getBondAmountByKeysCountWstETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getBondCurve\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"points\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"internalType\":\"structICSBondCurve.BondCurve\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getBondCurveId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondLockRetentionPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getBondShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getBondSummary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getBondSummaryShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getCurveInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"points\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"internalType\":\"structICSBondCurve.BondCurve\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256[]\",\"name\":\"points\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"internalType\":\"structICSBondCurve.BondCurve\",\"name\":\"curve\",\"type\":\"tuple\"}],\"name\":\"getKeysCountByBondAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getKeysCountByBondAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getLockedBondInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"retentionUntil\",\"type\":\"uint128\"}],\"internalType\":\"structICSBondLock.BondLock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalKeys\",\"type\":\"uint256\"}],\"name\":\"getRequiredBondForNextKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalKeys\",\"type\":\"uint256\"}],\"name\":\"getRequiredBondForNextKeysWstETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResumeSinceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getUnbondedKeysCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getUnbondedKeysCountToEject\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256[]\",\"name\":\"bondCurve\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeDistributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bondLockRetentionPeriod\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_chargePenaltyRecipient\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockBondETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"pauseFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"penalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"pullFeeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverStETHShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"releaseLockedBondETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renewBurnerAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"resetBondCurve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"setBondCurve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chargePenaltyRecipient\",\"type\":\"address\"}],\"name\":\"setChargePenaltyRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"retention\",\"type\":\"uint256\"}],\"name\":\"setLockedBondRetentionPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"settleLockedBondETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"settledAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBondShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"bondCurve\",\"type\":\"uint256[]\"}],\"name\":\"updateBondCurve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260043610610463575f3560e01c80638980f11f11610241578063cb11c52711610134578063dcab7f83116100b3578063f3f449c711610078578063f3f449c714610f3c578063f7966efe14610f5b578063f939122314610f7a578063fab382f114610f99578063fee6380514610fb8575f80fd5b8063dcab7f8314610e98578063def82d0214610eb7578063e5220e3f14610eea578063ead42a6914610f09578063f3efecc414610f28575f80fd5b8063d8fe7642116100f9578063d8fe764214610dd5578063d963ae5514610df4578063d9fb643a14610e13578063dbba4b4814610e46578063dc38ea3d14610e79575f80fd5b8063cb11c52714610d45578063cc810cb914610d59578063ce19793f14610d78578063d2fa16a614610d97578063d547741f14610db6575f80fd5b80639c516102116101c0578063b148db6a11610185578063b148db6a14610cb5578063b187bd2614610cd4578063b2d03e4d14610ce8578063b5b624bf14610d07578063ca15c87314610d26575f80fd5b80639c51610214610c1c578063a217fddf146107a1578063a302ee3814610c3b578063acf1c94814610c4f578063ae84975614610c82575f80fd5b80639010d07c116102065780639010d07c14610b8157806391d1485414610ba05780639996522514610bbf5780639a4df8f014610bde5780639b4c6c2714610bfd575f80fd5b80638980f11f14610a955780638b21f17014610ab45780638de2b27214610ae75780638ed5c5d714610b1a5780638f6549ae14610b4d575f80fd5b80634342b3c111610359578063589ff76c116102d857806370903eb91161029d57806370903eb91461096f57806374d70aea1461098e578063819d4cc6146109c157806383316184146109e0578063881fa03c14610a76575f80fd5b8063589ff76c146108c95780635a73bdc8146108dd5780635c654ad9146108f1578063699340f4146109105780636e13f09914610943575f80fd5b80634c7ed3d21161031e5780634c7ed3d214610825578063526352fc1461084457806352d8bfc214610877578063546da24f1461088b578063573b6245146108aa575f80fd5b80634342b3c11461076e578063443fbfef146107a1578063449add1b146107b45780634b2ce9fe146107d35780634bb22a7214610806575f80fd5b8063165123dd116103e55780632e599054116103aa5780632e599054146106cb5780632f2ff15d146106de57806336568abe146106fd578063389ed2671461071c578063433cd6c31461074f575f80fd5b8063165123dd146105ed57806321d439d51461060c578063248a9ca31461063f57806328846981146106795780632de03aa114610698575f80fd5b806306cd0e901161042b57806306cd0e90146105475780630d43e8ad146105665780630f23e7421461059c57806313d1234b146105bb57806315b5c477146105da575f80fd5b8063019c1a4f1461046757806301a5e9e31461048857806301ffc9a7146104ba578063046f7da2146104e95780630569b947146104fd575b5f80fd5b348015610472575f80fd5b50610486610481366004614c69565b610feb565b005b348015610493575f80fd5b506104a76104a2366004614cb1565b611026565b6040519081526020015b60405180910390f35b3480156104c5575f80fd5b506104d96104d4366004614cc8565b611038565b60405190151581526020016104b1565b3480156104f4575f80fd5b5061048661105c565b348015610508575f80fd5b506104a7610517366004614cb1565b5f9081527f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe1501602052604090205490565b348015610552575f80fd5b506104a7610561366004614cb1565b611091565b348015610571575f80fd5b505f54610584906001600160a01b031681565b6040516001600160a01b0390911681526020016104b1565b3480156105a7575f80fd5b506104a76105b6366004614d80565b6110c1565b3480156105c6575f80fd5b506104a76105d5366004614e55565b611140565b6104866105e8366004614cb1565b6111de565b3480156105f8575f80fd5b50600154610584906001600160a01b031681565b348015610617575f80fd5b506104a77fb5dffea014b759c493d63b1edaceb942631d6468998125e1b4fe427c9908213481565b34801561064a575f80fd5b506104a7610659366004614cb1565b5f9081525f80516020615385833981519152602052604090206001015490565b348015610684575f80fd5b506104a7610693366004614e55565b61135e565b3480156106a3575f80fd5b506104a77f2fc10cc8ae19568712f7a176fb4978616a610650813c9d05326c34abb62749c781565b6104866106d9366004614e89565b611398565b3480156106e9575f80fd5b506104866106f8366004614eb3565b6113f7565b348015610708575f80fd5b50610486610717366004614eb3565b611427565b348015610727575f80fd5b506104a77f139c2898040ef16910dc9f44dc697df79363da767d8bc92f2e310312b816e46d81565b34801561075a575f80fd5b50610486610769366004614ee1565b61145f565b348015610779575f80fd5b506104a77f645c9e6d2a86805cb5a28b1e4751c0dab493df7cf935070ce405489ba1a7bf7281565b3480156107ac575f80fd5b506104a75f81565b3480156107bf575f80fd5b506104866107ce366004614cb1565b611492565b3480156107de575f80fd5b506104a77f000000000000000000000000000000000000000000000000000000000000000a81565b348015610811575f80fd5b506104a7610820366004614cb1565b6114ce565b348015610830575f80fd5b5061048661083f366004614efc565b611546565b34801561084f575f80fd5b506104a77f0000000000000000000000000000000000000000000000000000000001e1338081565b348015610882575f80fd5b50610486611716565b348015610896575f80fd5b506104a76108a5366004614e55565b611772565b3480156108b5575f80fd5b506104a76108c4366004614f4b565b611780565b3480156108d4575f80fd5b506104a76117b5565b3480156108e8575f80fd5b506104866117e3565b3480156108fc575f80fd5b5061048661090b366004614e89565b611938565b34801561091b575f80fd5b506105847f000000000000000000000000c7cc160b58f8bb0bac94b80847e2cf2800565c5081565b34801561094e575f80fd5b5061096261095d366004614cb1565b6119b3565b6040516104b19190614f8a565b34801561097a575f80fd5b50610486610989366004614fe5565b6119fe565b348015610999575f80fd5b507f23f334b9eb5378c2a1573857b8f9d9ca79959360a69e73d3f16848e56ec92101546104a7565b3480156109cc575f80fd5b506104866109db366004614e89565b611a6c565b3480156109eb575f80fd5b50610a4f6109fa366004614cb1565b6040805180820182525f80825260209182018190529283525f805160206153658339815191528152918190208151808301909252546001600160801b038082168352600160801b909104169181019190915290565b6040805182516001600160801b0390811682526020938401511692810192909252016104b1565b348015610a81575f80fd5b50610486610a90366004614e55565b611abb565b348015610aa0575f80fd5b50610486610aaf366004614e89565b611b1d565b348015610abf575f80fd5b506105847f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c9503481565b348015610af2575f80fd5b506105847f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f81565b348015610b25575f80fd5b506104a77f40579467dba486691cc62fd8536d22c6d4dc9cdc7bc716ef2518422aa554c09881565b348015610b58575f80fd5b50610b6c610b67366004614cb1565b611bbe565b604080519283526020830191909152016104b1565b348015610b8c575f80fd5b50610584610b9b366004614e55565b611c78565b348015610bab575f80fd5b506104d9610bba366004614eb3565b611cb0565b348015610bca575f80fd5b50610486610bd9366004614cb1565b611ce6565b348015610be9575f80fd5b506104a7610bf8366004614d80565b611d19565b348015610c08575f80fd5b50610486610c17366004615053565b611d53565b348015610c27575f80fd5b506104a7610c36366004614cb1565b611d68565b348015610c46575f80fd5b506104a75f1981565b348015610c5a575f80fd5b506104a77fb3e25b5404b87e5a838579cb5d7481d61ad96ee284d38ec1e97c07ba64e7f6fc81565b348015610c8d575f80fd5b506104a77f000000000000000000000000000000000000000000000000000000000000000081565b348015610cc0575f80fd5b506104a7610ccf366004614e55565b611d73565b348015610cdf575f80fd5b506104d9611e4f565b348015610cf3575f80fd5b50610486610d02366004614e55565b611e7f565b348015610d12575f80fd5b50610962610d21366004614cb1565b611ebc565b348015610d31575f80fd5b506104a7610d40366004614cb1565b611f80565b348015610d50575f80fd5b506104a7600181565b348015610d64575f80fd5b50610486610d73366004614fe5565b611fb7565b348015610d83575f80fd5b50610b6c610d92366004614cb1565b612025565b348015610da2575f80fd5b506104a7610db1366004614d80565b6120d7565b348015610dc1575f80fd5b50610486610dd0366004614eb3565b61213b565b348015610de0575f80fd5b506104a7610def366004614cb1565b61216b565b348015610dff575f80fd5b50610486610e0e366004614e55565b61217d565b348015610e1e575f80fd5b506105847f0000000000000000000000008d09a4502cc8cf1547ad300e066060d043f6982d81565b348015610e51575f80fd5b506105847f00000000000000000000000028fab2059c713a7f9d8c86db49f9bb0e96af1ef881565b348015610e84575f80fd5b506104a7610e93366004614e55565b6121d0565b348015610ea3575f80fd5b50610486610eb2366004614e55565b6121de565b348015610ec2575f80fd5b507f78c5a36767279da056404c09083fca30cf3ea61c442cfaba6669f76a37393f00546104a7565b348015610ef5575f80fd5b50610486610f04366004614e55565b612231565b348015610f14575f80fd5b506104a7610f23366004614cb1565b612284565b348015610f33575f80fd5b506104866122d8565b348015610f47575f80fd5b50610486610f56366004614cb1565b6123f5565b348015610f66575f80fd5b50610486610f75366004614efc565b612428565b348015610f85575f80fd5b50610486610f94366004614fe5565b6125f8565b348015610fa4575f80fd5b50610486610fb33660046150a2565b612666565b348015610fc3575f80fd5b506104a77fd35e4a788498271198ec69c34f1dc762a1eee8200c111f598da1b3dde946783d81565b7fd35e4a788498271198ec69c34f1dc762a1eee8200c111f598da1b3dde946783d61101581612b1d565b611020848484612b27565b50505050565b5f611032826001612c92565b92915050565b5f6001600160e01b03198216635a05180f60e01b1480611032575061103282612d80565b7f2fc10cc8ae19568712f7a176fb4978616a610650813c9d05326c34abb62749c761108681612b1d565b61108e612db4565b50565b5f9081527f23f334b9eb5378c2a1573857b8f9d9ca79959360a69e73d3f16848e56ec92100602052604090205490565b5f825f036110d057505f611032565b8151518084116110fa576110f56110e8600186615136565b8451602091820201015190565b611138565b60208301516111098286615136565b6111139190615149565b61112e611121600184615136565b8551602091820201015190565b6111389190615160565b949350505050565b5f7f0000000000000000000000008d09a4502cc8cf1547ad300e066060d043f6982d6001600160a01b031663b0e3890061117a8585611772565b6040518263ffffffff1660e01b815260040161119891815260200190565b602060405180830381865afa1580156111b3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111d79190615173565b9392505050565b336001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f161461122757604051633bebb4c160e11b815260040160405180910390fd5b5f7f00000000000000000000000028fab2059c713a7f9d8c86db49f9bb0e96af1ef86001600160a01b031663e441d25f6040518163ffffffff1660e01b8152600401602060405180830381865afa158015611284573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112a8919061518a565b6001600160a01b0316346040515f6040518083038185875af1925050503d805f81146112ef576040519150601f19603f3d011682016040523d82523d5f602084013e6112f4565b606091505b5050905080611316576040516324f09be760e21b815260040160405180910390fd5b6113208234612e09565b817fb6ee6e3aae6776519627b46786a622b642c38cabfe4c97cb34054fd63fc11a233460405161135291815260200190565b60405180910390a25050565b5f7f0000000000000000000000008d09a4502cc8cf1547ad300e066060d043f6982d6001600160a01b031663b0e3890061117a8585611d73565b6113a0612e8f565b336001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f16146113e957604051633bebb4c160e11b815260040160405180910390fd5b6113f38282612eb7565b5050565b5f8281525f80516020615385833981519152602052604090206001015461141d81612b1d565b6110208383612fa4565b6001600160a01b03811633146114505760405163334bd91960e11b815260040160405180910390fd5b61145a8282612ff9565b505050565b7f40579467dba486691cc62fd8536d22c6d4dc9cdc7bc716ef2518422aa554c09861148981612b1d565b6113f382613045565b7fb5dffea014b759c493d63b1edaceb942631d6468998125e1b4fe427c990821346114bc81612b1d565b6114c5826130c1565b6113f382613164565b5f336001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f161461151857604051633bebb4c160e11b815260040160405180910390fd5b5f61152283612284565b905080156115375761153483826131ca565b91505b61154083613316565b50919050565b61154e612e8f565b336001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f161461159757604051633bebb4c160e11b815260040160405180910390fd5b8035158015906116375750604051636eb1769f60e11b81526001600160a01b0385811660048301523060248301528235917f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950349091169063dd62ed3e90604401602060405180830381865afa158015611611573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116359190615173565b105b1561170b576001600160a01b037f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950341663d505accf85308435602086013561168460608801604089016151a5565b6040516001600160e01b031960e088901b1681526001600160a01b0395861660048201529490931660248501526044840191909152606483015260ff166084820152606084013560a4820152608084013560c482015260e4015f604051808303815f87803b1580156116f4575f80fd5b505af1158015611706573d5f803e3d5ffd5b505050505b61102084848461335c565b61171e613459565b73a74528edc289b1a597faf83fcff7eff871cc01d96352d8bfc26040518163ffffffff1660e01b81526004015f6040518083038186803b158015611760575f80fd5b505af4158015611020573d5f803e3d5ffd5b5f6111d7836105b684611ebc565b5f7fd35e4a788498271198ec69c34f1dc762a1eee8200c111f598da1b3dde946783d6117ab81612b1d565b6111388484613482565b5f6117de7fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a025490565b905090565b6117eb613459565b5f6118147f23f334b9eb5378c2a1573857b8f9d9ca79959360a69e73d3f16848e56ec921015490565b604051633d7ad0b760e21b81523060048201527f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950346001600160a01b03169063f5eb42dc90602401602060405180830381865afa158015611876573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061189a9190615173565b6118a49190615136565b6040516389ad944360e01b81526001600160a01b037f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950341660048201526024810182905290915073a74528edc289b1a597faf83fcff7eff871cc01d9906389ad9443906044015f6040518083038186803b15801561191f575f80fd5b505af4158015611931573d5f803e3d5ffd5b5050505050565b611940613459565b604051635c654ad960e01b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d990635c654ad9906044015b5f6040518083038186803b158015611999575f80fd5b505af41580156119ab573d5f803e3d5ffd5b505050505050565b60408051808201909152606081525f6020820152611032610d21835f9081527f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe1501602052604090205490565b611a06612e8f565b336001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f1614611a4f57604051633bebb4c160e11b815260040160405180910390fd5b8015611a6157611a61868484846135be565b6119ab868686613640565b611a74613459565b6040516340cea66360e11b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d99063819d4cc690604401611983565b336001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f1614611b0457604051633bebb4c160e11b815260040160405180910390fd5b60015461145a90839083906001600160a01b0316613903565b611b25613459565b7f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950346001600160a01b0316826001600160a01b031603611b77576040516319efe5d760e21b815260040160405180910390fd5b604051638980f11f60e01b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d990638980f11f90604401611983565b5f80611bc983611091565b9150611c71611bd784612284565b6040516311d8d20560e31b815260048101869052611c6b907f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f6001600160a01b031690638ec6902890602401602060405180830381865afa158015611c3e573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611c629190615173565b6105b6876119b3565b01613a04565b9050915091565b5f8281527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e823717059320006020819052604082206111389084613a8e565b5f9182525f80516020615385833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b7f40579467dba486691cc62fd8536d22c6d4dc9cdc7bc716ef2518422aa554c098611d1081612b1d565b6113f382613a99565b5f7f0000000000000000000000008d09a4502cc8cf1547ad300e066060d043f6982d6001600160a01b031663b0e3890061117a85856110c1565b611d5c846130c1565b611020848484846135be565b5f611032825f612c92565b5f80611d7e8461216b565b90505f611e1c847f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f6001600160a01b0316638ec69028886040518263ffffffff1660e01b8152600401611dd391815260200190565b602060405180830381865afa158015611dee573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611e129190615173565b611c629190615160565b90505f611e2886612284565b611e329083615160565b9050828111611e41575f611e45565b8281035b9695505050505050565b5f611e787fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a025490565b4210905090565b7f645c9e6d2a86805cb5a28b1e4751c0dab493df7cf935070ce405489ba1a7bf72611ea981612b1d565b611eb2836130c1565b61145a8383613b57565b60408051808201909152606081525f60208201527f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe1500805483908110611f0357611f036151c5565b905f5260205f2090600202016040518060400160405290815f8201805480602002602001604051908101604052809291908181526020018280548015611f6657602002820191905f5260205f20905b815481526020019060010190808311611f52575b505050505081526020016001820154815250509050919050565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e823717059320006020819052604082206111d790613be3565b611fbf612e8f565b336001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f161461200857604051633bebb4c160e11b815260040160405180910390fd5b801561201a5761201a868484846135be565b6119ab868686613bec565b5f806120308361216b565b915061203b83612284565b6040516311d8d20560e31b8152600481018590526120cf907f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f6001600160a01b031690638ec6902890602401602060405180830381865afa1580156120a2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906120c69190615173565b6105b6866119b3565b019050915091565b8051602001515f908310156120ed57505f611032565b8151515f6120ff611121600184615136565b90508085106121255783602001518186038161211d5761211d6151d9565b048201612132565b61213285855f0151613ec0565b95945050505050565b5f8281525f80516020615385833981519152602052604090206001015461216181612b1d565b6110208383612ff9565b5f61103261217883611091565b613f2c565b336001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f16146121c657604051633bebb4c160e11b815260040160405180910390fd5b6113f38282612e09565b5f6111d783610db184611ebc565b336001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f161461222757604051633bebb4c160e11b815260040160405180910390fd5b6113f38282613f7b565b336001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f161461227a57604051633bebb4c160e11b815260040160405180910390fd5b61145a82826131ca565b5f8181525f8051602061536583398151915260205260408120805442600160801b9091046001600160801b0316116122bc575f6122c8565b80546001600160801b03165b6001600160801b03169392505050565b7f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950346001600160a01b031663095ea7b37f00000000000000000000000028fab2059c713a7f9d8c86db49f9bb0e96af1ef86001600160a01b03166327810b6e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612363573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612387919061518a565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201525f1960248201526044016020604051808303815f875af11580156123d1573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061108e91906151ed565b7f139c2898040ef16910dc9f44dc697df79363da767d8bc92f2e310312b816e46d61241f81612b1d565b6113f382614014565b612430612e8f565b336001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f161461247957604051633bebb4c160e11b815260040160405180910390fd5b8035158015906125195750604051636eb1769f60e11b81526001600160a01b0385811660048301523060248301528235917f0000000000000000000000008d09a4502cc8cf1547ad300e066060d043f6982d9091169063dd62ed3e90604401602060405180830381865afa1580156124f3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125179190615173565b105b156125ed576001600160a01b037f0000000000000000000000008d09a4502cc8cf1547ad300e066060d043f6982d1663d505accf85308435602086013561256660608801604089016151a5565b6040516001600160e01b031960e088901b1681526001600160a01b0395861660048201529490931660248501526044840191909152606483015260ff166084820152606084013560a4820152608084013560c482015260e4015f604051808303815f87803b1580156125d6575f80fd5b505af11580156125e8573d5f803e3d5ffd5b505050505b611020848484614063565b612600612e8f565b336001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f161461264957604051633bebb4c160e11b815260040160405180910390fd5b801561265b5761265b868484846135be565b6119ab8686866142e8565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f811580156126ab5750825b90505f8267ffffffffffffffff1660011480156126c75750303b155b9050811580156126d5575080155b156126f35760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561271d57845460ff60401b1916600160401b1785555b61272561441b565b61272f8b8b614423565b61273887614457565b6001600160a01b03891661275f57604051633ef39b8160e01b815260040160405180910390fd5b6001600160a01b0388166127865760405163658b92ad60e11b815260040160405180910390fd5b6127905f8a612fa4565b506127db7f645c9e6d2a86805cb5a28b1e4751c0dab493df7cf935070ce405489ba1a7bf727f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f612fa4565b506128267fb5dffea014b759c493d63b1edaceb942631d6468998125e1b4fe427c990821347f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f612fa4565b505f80546001600160a01b0319166001600160a01b038a1617905561284a86613045565b60405163095ea7b360e01b81526001600160a01b037f0000000000000000000000008d09a4502cc8cf1547ad300e066060d043f6982d811660048301525f1960248301527f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c95034169063095ea7b3906044016020604051808303815f875af11580156128d6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128fa91906151ed565b5060405163095ea7b360e01b81526001600160a01b037f000000000000000000000000c7cc160b58f8bb0bac94b80847e2cf2800565c50811660048301525f1960248301527f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c95034169063095ea7b3906044016020604051808303815f875af1158015612987573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129ab91906151ed565b507f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950346001600160a01b031663095ea7b37f00000000000000000000000028fab2059c713a7f9d8c86db49f9bb0e96af1ef86001600160a01b03166327810b6e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612a37573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612a5b919061518a565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201525f1960248201526044016020604051808303815f875af1158015612aa5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612ac991906151ed565b508315612b1057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b61108e8133614468565b7f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe150080545f1901841115612b6d576040516331e784e960e11b815260040160405180910390fd5b612b7783836144a6565b5f60018311612b86575f612ba4565b83836001198101818110612b9c57612b9c6151c5565b905060200201355b84845f198101818110612bb957612bb96151c5565b9050602002013503905060405180604001604052808585808060200260200160405190810160405280939291908181526020018383602002808284375f920191909152505050908252506020018290528254839087908110612c1d57612c1d6151c5565b905f5260205f2090600202015f820151815f019080519060200190612c43929190614bcd565b506020820151816001015590505050837f53da7af401538204fd91f2946f2fe85d05224d2cc766fd7aa9fbd8bf4fb4ce9f8484604051612c8492919061523c565b60405180910390a250505050565b6040516311d8d20560e31b8152600481018390525f9081906001600160a01b037f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f1690638ec6902890602401602060405180830381865afa158015612cf9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612d1d9190615173565b90505f612d2c61217886611091565b600a0190508315612d58575f612d4186612284565b9050808211612d5557829350505050611032565b90035b5f612d6682610db1886119b3565b9050808311612d75575f611e45565b909103949350505050565b5f6001600160e01b03198216637965db0b60e01b148061103257506301ffc9a760e01b6001600160e01b0319831614611032565b612dbc614590565b427fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a02556040517f62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9905f90a1565b5f612e1383612284565b9050815f03612e3557604051633649e09760e11b815260040160405180910390fd5b81811015612e5657604051633649e09760e11b815260040160405180910390fd5b5f8381525f80516020615365833981519152602052604090205461145a90849084840390600160801b90046001600160801b03166145b5565b612e97611e4f565b15612eb557604051630286f07360e31b815260040160405180910390fd5b565b345f03612ec2575050565b60405163a1903eab60e01b81525f60048201819052906001600160a01b037f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c95034169063a1903eab90349060240160206040518083038185885af1158015612f2a573d5f803e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190612f4f9190615173565b9050612f5b8282614669565b604080516001600160a01b038516815234602082015283917f16ec5116295424dec7fd52c87d9971a963ea7f59f741ad9ad468f0312055dc4991015b60405180910390a2505050565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200081612fd185856146d1565b90508015611138575f858152602083905260409020612ff09085614779565b50949350505050565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200081613026858561478d565b90508015611138575f858152602083905260409020612ff09085614806565b6001600160a01b03811661306c57604051631279f7c160e21b815260040160405180910390fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040519081527f4beaaee83871b066b675515d6a53567e76411f60266703cef934a01905a4d832906020015b60405180910390a150565b7f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f6001600160a01b031663a70c70e46040518163ffffffff1660e01b8152600401602060405180830381865afa15801561311d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131419190615173565b81101561314b5750565b604051633ed893db60e21b815260040160405180910390fd5b5f8181527f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe1501602090815260408083208390555191825282917f4642db1736894887bc907d721f20af84d3e585a0a3cea90f41b78b2aa906541b910160405180910390a250565b5f806131d583613a04565b90505f6131e2858361481a565b90507f00000000000000000000000028fab2059c713a7f9d8c86db49f9bb0e96af1ef86001600160a01b03166327810b6e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613240573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613264919061518a565b6040516308c2292560e31b8152306004820152602481018390526001600160a01b0391909116906346114928906044015f604051808303815f87803b1580156132ab575f80fd5b505af11580156132bd573d5f803e3d5ffd5b505050506132ca81613f2c565b9250847f4da924ae7845fe96897faab524b536685b8bbc4d82fbb45c10d941e0f86ade0f6132f784613f2c565b60408051918252602082018790520160405180910390a2505092915050565b5f8181525f8051602061536583398151915260205260408082208290555182917f844ae6b00e8a437dcdde1a634feab3273e08bb5c274a4be3b195b308ae0ba20a91a250565b805f0361336857505050565b5f61337282613a04565b604051636d78045960e01b81526001600160a01b038681166004830152306024830152604482018390529192507f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c9503490911690636d780459906064016020604051808303815f875af11580156133e9573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061340d9190615173565b506134188382614669565b604080516001600160a01b03861681526020810184905284917fee31ebba29fd5471227e42fd8ca621a892d689901892cb8febb03fe802c3214b9101612c84565b612eb57fb3e25b5404b87e5a838579cb5d7481d61ad96ee284d38ec1e97c07ba64e7f6fc612b1d565b5f7f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe15006134ae84846144a6565b5f600184116134bd575f6134db565b848460011981018181106134d3576134d36151c5565b905060200201355b85855f1981018181106134f0576134f06151c5565b90506020020135039050815f0160405180604001604052808787808060200260200160405190810160405280939291908181526020018383602002808284375f920182905250938552505050602091820185905283546001810185559381528190208251805193946002029091019261356c9284920190614bcd565b506020820151816001015550507f1fb1d9b944dd7015e95b7b7a9623c45792e4532badcf9c6e7a284d7d4d0570f085856040516135aa92919061523c565b60405180910390a150545f19019392505050565b5f80546040516321893f7b60e01b81526001600160a01b03909116906321893f7b906135f490889088908890889060040161524f565b6020604051808303815f875af1158015613610573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136349190615173565b90506119318582614669565b5f61364a84614842565b90505f81841061365a578161365c565b835b9050805f0361367e576040516312d37ee560e31b815260040160405180910390fd5b604051633d7ad0b760e21b81523060048201525f907f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950346001600160a01b03169063f5eb42dc90602401602060405180830381865afa1580156136e2573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906137069190615173565b90505f7f0000000000000000000000008d09a4502cc8cf1547ad300e066060d043f6982d6001600160a01b031663ea598cb061374185613f2c565b6040518263ffffffff1660e01b815260040161375f91815260200190565b6020604051808303815f875af115801561377b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061379f9190615173565b604051633d7ad0b760e21b81523060048201529091505f906001600160a01b037f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c95034169063f5eb42dc90602401602060405180830381865afa158015613806573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061382a9190615173565b905061383888828503614907565b60405163a9059cbb60e01b81526001600160a01b038781166004830152602482018490527f0000000000000000000000008d09a4502cc8cf1547ad300e066060d043f6982d169063a9059cbb906044015f604051808303815f87803b15801561389f575f80fd5b505af11580156138b1573d5f803e3d5ffd5b5050604080516001600160a01b038a168152602081018690528b93507fe6a8c06447e05a412e5e9581e088941f3994db3d8a9bfd3275b38d77acacafac92500160405180910390a25050505050505050565b5f8061390e84613a04565b905061391a858261481a565b604051638fcb4e5b60e01b81526001600160a01b038581166004830152602482018390529193507f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c9503490911690638fcb4e5b906044016020604051808303815f875af115801561398b573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906139af9190615173565b50847f8615528474a7bb3a28d9971535d956b79242b8e8fcfb27f3e331270fff088afd6139db83613f2c565b6139e485613f2c565b6040805192835260208301919091520160405180910390a2509392505050565b604051631920845160e01b8152600481018290525f907f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950346001600160a01b0316906319208451906024015b602060405180830381865afa158015613a6a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110329190615173565b5f6111d78383614966565b7f0000000000000000000000000000000000000000000000000000000000000000811080613ae657507f0000000000000000000000000000000000000000000000000000000001e1338081115b15613b045760405163dee7108760e01b815260040160405180910390fd5b807f78c5a36767279da056404c09083fca30cf3ea61c442cfaba6669f76a37393f00556040518181527fdaf5eddbe9ed0768e54cc8f739a9cb86c57fc70da07eff01d9ba886f21a7a4b3906020016130b6565b7f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe150080545f1901821115613b9d576040516331e784e960e11b815260040160405180910390fd5b5f838152600182016020526040908190208390555183907f4642db1736894887bc907d721f20af84d3e585a0a3cea90f41b78b2aa906541b90612f979085815260200190565b5f611032825490565b5f613bf684614842565b90505f613c0282613f2c565b8410613c0e5781613c17565b613c1784613a04565b9050805f03613c39576040516312d37ee560e31b815260040160405180910390fd5b6040805160018082528183019092525f9160208083019080368337019050509050613c6382613f2c565b815f81518110613c7557613c756151c5565b6020908102919091010152604051633d7ad0b760e21b81523060048201525f907f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950346001600160a01b03169063f5eb42dc90602401602060405180830381865afa158015613ce4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613d089190615173565b90505f7f000000000000000000000000c7cc160b58f8bb0bac94b80847e2cf2800565c506001600160a01b031663d668104284886040518363ffffffff1660e01b8152600401613d5992919061526e565b5f604051808303815f875af1158015613d74573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f19168201604052613d9b91908101906152c4565b604051633d7ad0b760e21b81523060048201529091505f906001600160a01b037f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c95034169063f5eb42dc90602401602060405180830381865afa158015613e02573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613e269190615173565b9050613e3489828503614907565b887f26673a9d018b21192d08ee14377b798f11b9e5b15ea1559c110265716b8985b588865f81518110613e6957613e696151c5565b6020026020010151855f81518110613e8357613e836151c5565b602090810291909101810151604080516001600160a01b0390951685529184019290925282015260600160405180910390a2505050505050505050565b80515f9081906001190181805b828411613f21575050600282820104602081810286010151808703613efa57506001019250611032915050565b80871015613f0d57600182039250613ecd565b80871115613f1c578160010193505b613ecd565b509195945050505050565b604051630f451f7160e31b8152600481018290525f907f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950346001600160a01b031690637a28fb8890602401613a4f565b7f78c5a36767279da056404c09083fca30cf3ea61c442cfaba6669f76a37393f005f829003613fbd57604051633649e09760e11b815260040160405180910390fd5b5f83815260018201602052604090205442600160801b9091046001600160801b03161115614004575f8381526001820160205260409020546001600160801b031691909101905b61145a8383835f015442016145b5565b61401c612e8f565b805f0361403c5760405163ad58bfc760e01b815260040160405180910390fd5b5f5f19820361404d57505f1961405a565b6140578242615160565b90505b6113f38161498c565b805f0361406f57505050565b6040516323b872dd60e01b81526001600160a01b038481166004830152306024830152604482018390527f0000000000000000000000008d09a4502cc8cf1547ad300e066060d043f6982d16906323b872dd906064015f604051808303815f87803b1580156140dc575f80fd5b505af11580156140ee573d5f803e3d5ffd5b5050604051633d7ad0b760e21b81523060048201525f92507f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950346001600160a01b0316915063f5eb42dc90602401602060405180830381865afa158015614156573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061417a9190615173565b604051636f074d1f60e11b8152600481018490529091507f0000000000000000000000008d09a4502cc8cf1547ad300e066060d043f6982d6001600160a01b03169063de0e9a3e906024016020604051808303815f875af11580156141e1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906142059190615173565b50604051633d7ad0b760e21b81523060048201525f907f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950346001600160a01b03169063f5eb42dc90602401602060405180830381865afa15801561426a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061428e9190615173565b905061429c84838303614669565b604080516001600160a01b03871681526020810185905285917f6576bbc9c5b478bf9717dc3d2bcb485e5ff0727df77c72558727597f3609d3f191015b60405180910390a25050505050565b5f6142f284614842565b90505f6142fe82613f2c565b841061430a5781614313565b61431384613a04565b9050805f03614335576040516312d37ee560e31b815260040160405180910390fd5b61433f8582614907565b604051638fcb4e5b60e01b81526001600160a01b038481166004830152602482018390527f0000000000000000000000003f1c547b21f65e10480de3ad8e19faac46c950341690638fcb4e5b906044016020604051808303815f875af11580156143ab573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906143cf9190615173565b50847f3e3a1398fe71575ed0c17a80cd9d46ad684c2c75c2fad7b0e7dde15e78ab22d3846143fc84613f2c565b604080516001600160a01b0390931683526020830191909152016142d9565b612eb5614a27565b61442b614a27565b5f6144368383613482565b9050801561145a57604051634273eaaf60e11b815260040160405180910390fd5b61445f614a27565b61108e81613a99565b6144728282611cb0565b6113f35760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044015b60405180910390fd5b60018110806144d457507f000000000000000000000000000000000000000000000000000000000000000a81115b156144f257604051638326bf5360e01b815260040160405180910390fd5b81815f818110614504576145046151c5565b905060200201355f0361452a576040516302527aef60e21b815260040160405180910390fd5b60015b8181101561145a5782826001830381811061454a5761454a6151c5565b90506020020135838383818110614563576145636151c5565b9050602002013511614588576040516302527aef60e21b815260040160405180910390fd5b60010161452d565b614598611e4f565b612eb55760405163b047186b60e01b815260040160405180910390fd5b815f036145c55761145a83613316565b60405180604001604052806145d984614a70565b6001600160801b031681526020016145f083614a70565b6001600160801b039081169091525f8581525f8051602061536583398151915260209081526040918290208451948201518416600160801b029490931693909317909155805184815291820183905284917f69a153d448f54b17f05cf3b268a2efab87c94a4727d108c4ca4aa3e5d65113de9101612f97565b805f03614674575050565b5f9182527f23f334b9eb5378c2a1573857b8f9d9ca79959360a69e73d3f16848e56ec9210060205260409091208054820190557f23f334b9eb5378c2a1573857b8f9d9ca79959360a69e73d3f16848e56ec9210180549091019055565b5f5f805160206153858339815191526146ea8484611cb0565b614769575f848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561471f3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050611032565b5f915050611032565b5092915050565b5f6111d7836001600160a01b038416614aa7565b5f5f805160206153858339815191526147a68484611cb0565b15614769575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050611032565b5f6111d7836001600160a01b038416614af3565b5f8061482584611091565b90508083106148345780614836565b825b91506147728483614907565b5f8061484d83611091565b90505f6148f061485c85612284565b6040516311d8d20560e31b815260048101879052611c6b907f0000000000000000000000004562c3e63c2e586cd1651b958c22f88135acad4f6001600160a01b031690638ec6902890602401602060405180830381865afa1580156148c3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906148e79190615173565b6105b6886119b3565b90508082116148ff575f611138565b900392915050565b5f9182527f23f334b9eb5378c2a1573857b8f9d9ca79959360a69e73d3f16848e56ec921006020526040909120805482900390557f23f334b9eb5378c2a1573857b8f9d9ca79959360a69e73d3f16848e56ec921018054919091039055565b5f825f01828154811061497b5761497b6151c5565b905f5260205f200154905092915050565b6149b57fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a02829055565b5f1981036149ee576040515f1981527f32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e906020016130b6565b7f32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e614a194283615136565b6040519081526020016130b6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16612eb557604051631afcd79f60e31b815260040160405180910390fd5b5f6001600160801b03821115614aa3576040516306dfcc6560e41b8152608060048201526024810183905260440161449d565b5090565b5f818152600183016020526040812054614aec57508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155611032565b505f611032565b5f8181526001830160205260408120548015614769575f614b15600183615136565b85549091505f90614b2890600190615136565b9050808214614b87575f865f018281548110614b4657614b466151c5565b905f5260205f200154905080875f018481548110614b6657614b666151c5565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080614b9857614b98615350565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050611032565b828054828255905f5260205f20908101928215614c06579160200282015b82811115614c06578251825591602001919060010190614beb565b50614aa39291505b80821115614aa3575f8155600101614c0e565b5f8083601f840112614c31575f80fd5b50813567ffffffffffffffff811115614c48575f80fd5b6020830191508360208260051b8501011115614c62575f80fd5b9250929050565b5f805f60408486031215614c7b575f80fd5b83359250602084013567ffffffffffffffff811115614c98575f80fd5b614ca486828701614c21565b9497909650939450505050565b5f60208284031215614cc1575f80fd5b5035919050565b5f60208284031215614cd8575f80fd5b81356001600160e01b0319811681146111d7575f80fd5b634e487b7160e01b5f52604160045260245ffd5b6040805190810167ffffffffffffffff81118282101715614d2657614d26614cef565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715614d5557614d55614cef565b604052919050565b5f67ffffffffffffffff821115614d7657614d76614cef565b5060051b60200190565b5f8060408385031215614d91575f80fd5b8235915060208084013567ffffffffffffffff80821115614db0575f80fd5b9085019060408288031215614dc3575f80fd5b614dcb614d03565b823582811115614dd9575f80fd5b83019150601f82018813614deb575f80fd5b8135614dfe614df982614d5d565b614d2c565b81815260059190911b8301850190858101908a831115614e1c575f80fd5b938601935b82851015614e3a57843582529386019390860190614e21565b83525050918301359282019290925292959294509192505050565b5f8060408385031215614e66575f80fd5b50508035926020909101359150565b6001600160a01b038116811461108e575f80fd5b5f8060408385031215614e9a575f80fd5b8235614ea581614e75565b946020939093013593505050565b5f8060408385031215614ec4575f80fd5b823591506020830135614ed681614e75565b809150509250929050565b5f60208284031215614ef1575f80fd5b81356111d781614e75565b5f805f80848603610100811215614f11575f80fd5b8535614f1c81614e75565b9450602086013593506040860135925060a0605f1982011215614f3d575f80fd5b509295919450926060019150565b5f8060208385031215614f5c575f80fd5b823567ffffffffffffffff811115614f72575f80fd5b614f7e85828601614c21565b90969095509350505050565b602080825282516040838301528051606084018190525f9291820190839060808601905b80831015614fce5783518252928401926001929092019190840190614fae565b508387015160408701528094505050505092915050565b5f805f805f8060a08789031215614ffa575f80fd5b8635955060208701359450604087013561501381614e75565b935060608701359250608087013567ffffffffffffffff811115615035575f80fd5b61504189828a01614c21565b979a9699509497509295939492505050565b5f805f8060608587031215615066575f80fd5b8435935060208501359250604085013567ffffffffffffffff81111561508a575f80fd5b61509687828801614c21565b95989497509550505050565b5f805f805f8060a087890312156150b7575f80fd5b863567ffffffffffffffff8111156150cd575f80fd5b6150d989828a01614c21565b90975095505060208701356150ed81614e75565b935060408701356150fd81614e75565b925060608701359150608087013561511481614e75565b809150509295509295509295565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561103257611032615122565b808202811582820484141761103257611032615122565b8082018082111561103257611032615122565b5f60208284031215615183575f80fd5b5051919050565b5f6020828403121561519a575f80fd5b81516111d781614e75565b5f602082840312156151b5575f80fd5b813560ff811681146111d7575f80fd5b634e487b7160e01b5f52603260045260245ffd5b634e487b7160e01b5f52601260045260245ffd5b5f602082840312156151fd575f80fd5b815180151581146111d7575f80fd5b8183525f6001600160fb1b03831115615223575f80fd5b8260051b80836020870137939093016020019392505050565b602081525f61113860208301848661520c565b848152836020820152606060408201525f611e4560608301848661520c565b604080825283519082018190525f906020906060840190828701845b828110156152a65781518452928401929084019060010161528a565b50505080925050506001600160a01b03831660208301529392505050565b5f60208083850312156152d5575f80fd5b825167ffffffffffffffff8111156152eb575f80fd5b8301601f810185136152fb575f80fd5b8051615309614df982614d5d565b81815260059190911b82018301908381019087831115615327575f80fd5b928401925b828410156153455783518252928401929084019061532c565b979650505050505050565b634e487b7160e01b5f52603160045260245ffdfe78c5a36767279da056404c09083fca30cf3ea61c442cfaba6669f76a37393f0102dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a164736f6c6343000818000a",
}

// CSAccountingABI is the input ABI used to generate the binding from.
// Deprecated: Use CSAccountingMetaData.ABI instead.
var CSAccountingABI = CSAccountingMetaData.ABI

// CSAccountingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CSAccountingMetaData.Bin instead.
var CSAccountingBin = CSAccountingMetaData.Bin

// DeployCSAccounting deploys a new Ethereum contract, binding an instance of CSAccounting to it.
func DeployCSAccounting(auth *bind.TransactOpts, backend bind.ContractBackend, lidoLocator common.Address, communityStakingModule common.Address, maxCurveLength *big.Int, minBondLockRetentionPeriod *big.Int, maxBondLockRetentionPeriod *big.Int) (common.Address, *types.Transaction, *CSAccounting, error) {
	parsed, err := CSAccountingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CSAccountingBin), backend, lidoLocator, communityStakingModule, maxCurveLength, minBondLockRetentionPeriod, maxBondLockRetentionPeriod)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CSAccounting{CSAccountingCaller: CSAccountingCaller{contract: contract}, CSAccountingTransactor: CSAccountingTransactor{contract: contract}, CSAccountingFilterer: CSAccountingFilterer{contract: contract}}, nil
}

// CSAccounting is an auto generated Go binding around an Ethereum contract.
type CSAccounting struct {
	CSAccountingCaller     // Read-only binding to the contract
	CSAccountingTransactor // Write-only binding to the contract
	CSAccountingFilterer   // Log filterer for contract events
}

// CSAccountingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CSAccountingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CSAccountingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CSAccountingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CSAccountingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CSAccountingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CSAccountingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CSAccountingSession struct {
	Contract     *CSAccounting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CSAccountingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CSAccountingCallerSession struct {
	Contract *CSAccountingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CSAccountingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CSAccountingTransactorSession struct {
	Contract     *CSAccountingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CSAccountingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CSAccountingRaw struct {
	Contract *CSAccounting // Generic contract binding to access the raw methods on
}

// CSAccountingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CSAccountingCallerRaw struct {
	Contract *CSAccountingCaller // Generic read-only contract binding to access the raw methods on
}

// CSAccountingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CSAccountingTransactorRaw struct {
	Contract *CSAccountingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCSAccounting creates a new instance of CSAccounting, bound to a specific deployed contract.
func NewCSAccounting(address common.Address, backend bind.ContractBackend) (*CSAccounting, error) {
	contract, err := bindCSAccounting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CSAccounting{CSAccountingCaller: CSAccountingCaller{contract: contract}, CSAccountingTransactor: CSAccountingTransactor{contract: contract}, CSAccountingFilterer: CSAccountingFilterer{contract: contract}}, nil
}

// NewCSAccountingCaller creates a new read-only instance of CSAccounting, bound to a specific deployed contract.
func NewCSAccountingCaller(address common.Address, caller bind.ContractCaller) (*CSAccountingCaller, error) {
	contract, err := bindCSAccounting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CSAccountingCaller{contract: contract}, nil
}

// NewCSAccountingTransactor creates a new write-only instance of CSAccounting, bound to a specific deployed contract.
func NewCSAccountingTransactor(address common.Address, transactor bind.ContractTransactor) (*CSAccountingTransactor, error) {
	contract, err := bindCSAccounting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CSAccountingTransactor{contract: contract}, nil
}

// NewCSAccountingFilterer creates a new log filterer instance of CSAccounting, bound to a specific deployed contract.
func NewCSAccountingFilterer(address common.Address, filterer bind.ContractFilterer) (*CSAccountingFilterer, error) {
	contract, err := bindCSAccounting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CSAccountingFilterer{contract: contract}, nil
}

// bindCSAccounting binds a generic wrapper to an already deployed contract.
func bindCSAccounting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CSAccountingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CSAccounting *CSAccountingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CSAccounting.Contract.CSAccountingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CSAccounting *CSAccountingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSAccounting.Contract.CSAccountingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CSAccounting *CSAccountingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CSAccounting.Contract.CSAccountingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CSAccounting *CSAccountingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CSAccounting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CSAccounting *CSAccountingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSAccounting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CSAccounting *CSAccountingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CSAccounting.Contract.contract.Transact(opts, method, params...)
}

// ACCOUNTINGMANAGERROLE is a free data retrieval call binding the contract method 0x8ed5c5d7.
//
// Solidity: function ACCOUNTING_MANAGER_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCaller) ACCOUNTINGMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "ACCOUNTING_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACCOUNTINGMANAGERROLE is a free data retrieval call binding the contract method 0x8ed5c5d7.
//
// Solidity: function ACCOUNTING_MANAGER_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingSession) ACCOUNTINGMANAGERROLE() ([32]byte, error) {
	return _CSAccounting.Contract.ACCOUNTINGMANAGERROLE(&_CSAccounting.CallOpts)
}

// ACCOUNTINGMANAGERROLE is a free data retrieval call binding the contract method 0x8ed5c5d7.
//
// Solidity: function ACCOUNTING_MANAGER_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCallerSession) ACCOUNTINGMANAGERROLE() ([32]byte, error) {
	return _CSAccounting.Contract.ACCOUNTINGMANAGERROLE(&_CSAccounting.CallOpts)
}

// CSM is a free data retrieval call binding the contract method 0x8de2b272.
//
// Solidity: function CSM() view returns(address)
func (_CSAccounting *CSAccountingCaller) CSM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "CSM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CSM is a free data retrieval call binding the contract method 0x8de2b272.
//
// Solidity: function CSM() view returns(address)
func (_CSAccounting *CSAccountingSession) CSM() (common.Address, error) {
	return _CSAccounting.Contract.CSM(&_CSAccounting.CallOpts)
}

// CSM is a free data retrieval call binding the contract method 0x8de2b272.
//
// Solidity: function CSM() view returns(address)
func (_CSAccounting *CSAccountingCallerSession) CSM() (common.Address, error) {
	return _CSAccounting.Contract.CSM(&_CSAccounting.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CSAccounting.Contract.DEFAULTADMINROLE(&_CSAccounting.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _CSAccounting.Contract.DEFAULTADMINROLE(&_CSAccounting.CallOpts)
}

// DEFAULTBONDCURVEID is a free data retrieval call binding the contract method 0x443fbfef.
//
// Solidity: function DEFAULT_BOND_CURVE_ID() view returns(uint256)
func (_CSAccounting *CSAccountingCaller) DEFAULTBONDCURVEID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "DEFAULT_BOND_CURVE_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTBONDCURVEID is a free data retrieval call binding the contract method 0x443fbfef.
//
// Solidity: function DEFAULT_BOND_CURVE_ID() view returns(uint256)
func (_CSAccounting *CSAccountingSession) DEFAULTBONDCURVEID() (*big.Int, error) {
	return _CSAccounting.Contract.DEFAULTBONDCURVEID(&_CSAccounting.CallOpts)
}

// DEFAULTBONDCURVEID is a free data retrieval call binding the contract method 0x443fbfef.
//
// Solidity: function DEFAULT_BOND_CURVE_ID() view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) DEFAULTBONDCURVEID() (*big.Int, error) {
	return _CSAccounting.Contract.DEFAULTBONDCURVEID(&_CSAccounting.CallOpts)
}

// LIDO is a free data retrieval call binding the contract method 0x8b21f170.
//
// Solidity: function LIDO() view returns(address)
func (_CSAccounting *CSAccountingCaller) LIDO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "LIDO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIDO is a free data retrieval call binding the contract method 0x8b21f170.
//
// Solidity: function LIDO() view returns(address)
func (_CSAccounting *CSAccountingSession) LIDO() (common.Address, error) {
	return _CSAccounting.Contract.LIDO(&_CSAccounting.CallOpts)
}

// LIDO is a free data retrieval call binding the contract method 0x8b21f170.
//
// Solidity: function LIDO() view returns(address)
func (_CSAccounting *CSAccountingCallerSession) LIDO() (common.Address, error) {
	return _CSAccounting.Contract.LIDO(&_CSAccounting.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_CSAccounting *CSAccountingCaller) LIDOLOCATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "LIDO_LOCATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_CSAccounting *CSAccountingSession) LIDOLOCATOR() (common.Address, error) {
	return _CSAccounting.Contract.LIDOLOCATOR(&_CSAccounting.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_CSAccounting *CSAccountingCallerSession) LIDOLOCATOR() (common.Address, error) {
	return _CSAccounting.Contract.LIDOLOCATOR(&_CSAccounting.CallOpts)
}

// MANAGEBONDCURVESROLE is a free data retrieval call binding the contract method 0xfee63805.
//
// Solidity: function MANAGE_BOND_CURVES_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCaller) MANAGEBONDCURVESROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "MANAGE_BOND_CURVES_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEBONDCURVESROLE is a free data retrieval call binding the contract method 0xfee63805.
//
// Solidity: function MANAGE_BOND_CURVES_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingSession) MANAGEBONDCURVESROLE() ([32]byte, error) {
	return _CSAccounting.Contract.MANAGEBONDCURVESROLE(&_CSAccounting.CallOpts)
}

// MANAGEBONDCURVESROLE is a free data retrieval call binding the contract method 0xfee63805.
//
// Solidity: function MANAGE_BOND_CURVES_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCallerSession) MANAGEBONDCURVESROLE() ([32]byte, error) {
	return _CSAccounting.Contract.MANAGEBONDCURVESROLE(&_CSAccounting.CallOpts)
}

// MAXBONDLOCKRETENTIONPERIOD is a free data retrieval call binding the contract method 0x526352fc.
//
// Solidity: function MAX_BOND_LOCK_RETENTION_PERIOD() view returns(uint256)
func (_CSAccounting *CSAccountingCaller) MAXBONDLOCKRETENTIONPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "MAX_BOND_LOCK_RETENTION_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBONDLOCKRETENTIONPERIOD is a free data retrieval call binding the contract method 0x526352fc.
//
// Solidity: function MAX_BOND_LOCK_RETENTION_PERIOD() view returns(uint256)
func (_CSAccounting *CSAccountingSession) MAXBONDLOCKRETENTIONPERIOD() (*big.Int, error) {
	return _CSAccounting.Contract.MAXBONDLOCKRETENTIONPERIOD(&_CSAccounting.CallOpts)
}

// MAXBONDLOCKRETENTIONPERIOD is a free data retrieval call binding the contract method 0x526352fc.
//
// Solidity: function MAX_BOND_LOCK_RETENTION_PERIOD() view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) MAXBONDLOCKRETENTIONPERIOD() (*big.Int, error) {
	return _CSAccounting.Contract.MAXBONDLOCKRETENTIONPERIOD(&_CSAccounting.CallOpts)
}

// MAXCURVELENGTH is a free data retrieval call binding the contract method 0x4b2ce9fe.
//
// Solidity: function MAX_CURVE_LENGTH() view returns(uint256)
func (_CSAccounting *CSAccountingCaller) MAXCURVELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "MAX_CURVE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCURVELENGTH is a free data retrieval call binding the contract method 0x4b2ce9fe.
//
// Solidity: function MAX_CURVE_LENGTH() view returns(uint256)
func (_CSAccounting *CSAccountingSession) MAXCURVELENGTH() (*big.Int, error) {
	return _CSAccounting.Contract.MAXCURVELENGTH(&_CSAccounting.CallOpts)
}

// MAXCURVELENGTH is a free data retrieval call binding the contract method 0x4b2ce9fe.
//
// Solidity: function MAX_CURVE_LENGTH() view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) MAXCURVELENGTH() (*big.Int, error) {
	return _CSAccounting.Contract.MAXCURVELENGTH(&_CSAccounting.CallOpts)
}

// MINBONDLOCKRETENTIONPERIOD is a free data retrieval call binding the contract method 0xae849756.
//
// Solidity: function MIN_BOND_LOCK_RETENTION_PERIOD() view returns(uint256)
func (_CSAccounting *CSAccountingCaller) MINBONDLOCKRETENTIONPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "MIN_BOND_LOCK_RETENTION_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBONDLOCKRETENTIONPERIOD is a free data retrieval call binding the contract method 0xae849756.
//
// Solidity: function MIN_BOND_LOCK_RETENTION_PERIOD() view returns(uint256)
func (_CSAccounting *CSAccountingSession) MINBONDLOCKRETENTIONPERIOD() (*big.Int, error) {
	return _CSAccounting.Contract.MINBONDLOCKRETENTIONPERIOD(&_CSAccounting.CallOpts)
}

// MINBONDLOCKRETENTIONPERIOD is a free data retrieval call binding the contract method 0xae849756.
//
// Solidity: function MIN_BOND_LOCK_RETENTION_PERIOD() view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) MINBONDLOCKRETENTIONPERIOD() (*big.Int, error) {
	return _CSAccounting.Contract.MINBONDLOCKRETENTIONPERIOD(&_CSAccounting.CallOpts)
}

// MINCURVELENGTH is a free data retrieval call binding the contract method 0xcb11c527.
//
// Solidity: function MIN_CURVE_LENGTH() view returns(uint256)
func (_CSAccounting *CSAccountingCaller) MINCURVELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "MIN_CURVE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINCURVELENGTH is a free data retrieval call binding the contract method 0xcb11c527.
//
// Solidity: function MIN_CURVE_LENGTH() view returns(uint256)
func (_CSAccounting *CSAccountingSession) MINCURVELENGTH() (*big.Int, error) {
	return _CSAccounting.Contract.MINCURVELENGTH(&_CSAccounting.CallOpts)
}

// MINCURVELENGTH is a free data retrieval call binding the contract method 0xcb11c527.
//
// Solidity: function MIN_CURVE_LENGTH() view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) MINCURVELENGTH() (*big.Int, error) {
	return _CSAccounting.Contract.MINCURVELENGTH(&_CSAccounting.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_CSAccounting *CSAccountingCaller) PAUSEINFINITELY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "PAUSE_INFINITELY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_CSAccounting *CSAccountingSession) PAUSEINFINITELY() (*big.Int, error) {
	return _CSAccounting.Contract.PAUSEINFINITELY(&_CSAccounting.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) PAUSEINFINITELY() (*big.Int, error) {
	return _CSAccounting.Contract.PAUSEINFINITELY(&_CSAccounting.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingSession) PAUSEROLE() ([32]byte, error) {
	return _CSAccounting.Contract.PAUSEROLE(&_CSAccounting.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCallerSession) PAUSEROLE() ([32]byte, error) {
	return _CSAccounting.Contract.PAUSEROLE(&_CSAccounting.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCaller) RECOVERERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "RECOVERER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingSession) RECOVERERROLE() ([32]byte, error) {
	return _CSAccounting.Contract.RECOVERERROLE(&_CSAccounting.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCallerSession) RECOVERERROLE() ([32]byte, error) {
	return _CSAccounting.Contract.RECOVERERROLE(&_CSAccounting.CallOpts)
}

// RESETBONDCURVEROLE is a free data retrieval call binding the contract method 0x21d439d5.
//
// Solidity: function RESET_BOND_CURVE_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCaller) RESETBONDCURVEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "RESET_BOND_CURVE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESETBONDCURVEROLE is a free data retrieval call binding the contract method 0x21d439d5.
//
// Solidity: function RESET_BOND_CURVE_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingSession) RESETBONDCURVEROLE() ([32]byte, error) {
	return _CSAccounting.Contract.RESETBONDCURVEROLE(&_CSAccounting.CallOpts)
}

// RESETBONDCURVEROLE is a free data retrieval call binding the contract method 0x21d439d5.
//
// Solidity: function RESET_BOND_CURVE_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCallerSession) RESETBONDCURVEROLE() ([32]byte, error) {
	return _CSAccounting.Contract.RESETBONDCURVEROLE(&_CSAccounting.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingSession) RESUMEROLE() ([32]byte, error) {
	return _CSAccounting.Contract.RESUMEROLE(&_CSAccounting.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCallerSession) RESUMEROLE() ([32]byte, error) {
	return _CSAccounting.Contract.RESUMEROLE(&_CSAccounting.CallOpts)
}

// SETBONDCURVEROLE is a free data retrieval call binding the contract method 0x4342b3c1.
//
// Solidity: function SET_BOND_CURVE_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCaller) SETBONDCURVEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "SET_BOND_CURVE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETBONDCURVEROLE is a free data retrieval call binding the contract method 0x4342b3c1.
//
// Solidity: function SET_BOND_CURVE_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingSession) SETBONDCURVEROLE() ([32]byte, error) {
	return _CSAccounting.Contract.SETBONDCURVEROLE(&_CSAccounting.CallOpts)
}

// SETBONDCURVEROLE is a free data retrieval call binding the contract method 0x4342b3c1.
//
// Solidity: function SET_BOND_CURVE_ROLE() view returns(bytes32)
func (_CSAccounting *CSAccountingCallerSession) SETBONDCURVEROLE() ([32]byte, error) {
	return _CSAccounting.Contract.SETBONDCURVEROLE(&_CSAccounting.CallOpts)
}

// WITHDRAWALQUEUE is a free data retrieval call binding the contract method 0x699340f4.
//
// Solidity: function WITHDRAWAL_QUEUE() view returns(address)
func (_CSAccounting *CSAccountingCaller) WITHDRAWALQUEUE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "WITHDRAWAL_QUEUE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WITHDRAWALQUEUE is a free data retrieval call binding the contract method 0x699340f4.
//
// Solidity: function WITHDRAWAL_QUEUE() view returns(address)
func (_CSAccounting *CSAccountingSession) WITHDRAWALQUEUE() (common.Address, error) {
	return _CSAccounting.Contract.WITHDRAWALQUEUE(&_CSAccounting.CallOpts)
}

// WITHDRAWALQUEUE is a free data retrieval call binding the contract method 0x699340f4.
//
// Solidity: function WITHDRAWAL_QUEUE() view returns(address)
func (_CSAccounting *CSAccountingCallerSession) WITHDRAWALQUEUE() (common.Address, error) {
	return _CSAccounting.Contract.WITHDRAWALQUEUE(&_CSAccounting.CallOpts)
}

// WSTETH is a free data retrieval call binding the contract method 0xd9fb643a.
//
// Solidity: function WSTETH() view returns(address)
func (_CSAccounting *CSAccountingCaller) WSTETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "WSTETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WSTETH is a free data retrieval call binding the contract method 0xd9fb643a.
//
// Solidity: function WSTETH() view returns(address)
func (_CSAccounting *CSAccountingSession) WSTETH() (common.Address, error) {
	return _CSAccounting.Contract.WSTETH(&_CSAccounting.CallOpts)
}

// WSTETH is a free data retrieval call binding the contract method 0xd9fb643a.
//
// Solidity: function WSTETH() view returns(address)
func (_CSAccounting *CSAccountingCallerSession) WSTETH() (common.Address, error) {
	return _CSAccounting.Contract.WSTETH(&_CSAccounting.CallOpts)
}

// ChargePenaltyRecipient is a free data retrieval call binding the contract method 0x165123dd.
//
// Solidity: function chargePenaltyRecipient() view returns(address)
func (_CSAccounting *CSAccountingCaller) ChargePenaltyRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "chargePenaltyRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChargePenaltyRecipient is a free data retrieval call binding the contract method 0x165123dd.
//
// Solidity: function chargePenaltyRecipient() view returns(address)
func (_CSAccounting *CSAccountingSession) ChargePenaltyRecipient() (common.Address, error) {
	return _CSAccounting.Contract.ChargePenaltyRecipient(&_CSAccounting.CallOpts)
}

// ChargePenaltyRecipient is a free data retrieval call binding the contract method 0x165123dd.
//
// Solidity: function chargePenaltyRecipient() view returns(address)
func (_CSAccounting *CSAccountingCallerSession) ChargePenaltyRecipient() (common.Address, error) {
	return _CSAccounting.Contract.ChargePenaltyRecipient(&_CSAccounting.CallOpts)
}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_CSAccounting *CSAccountingCaller) FeeDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "feeDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_CSAccounting *CSAccountingSession) FeeDistributor() (common.Address, error) {
	return _CSAccounting.Contract.FeeDistributor(&_CSAccounting.CallOpts)
}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_CSAccounting *CSAccountingCallerSession) FeeDistributor() (common.Address, error) {
	return _CSAccounting.Contract.FeeDistributor(&_CSAccounting.CallOpts)
}

// GetActualLockedBond is a free data retrieval call binding the contract method 0xead42a69.
//
// Solidity: function getActualLockedBond(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetActualLockedBond(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getActualLockedBond", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActualLockedBond is a free data retrieval call binding the contract method 0xead42a69.
//
// Solidity: function getActualLockedBond(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetActualLockedBond(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetActualLockedBond(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetActualLockedBond is a free data retrieval call binding the contract method 0xead42a69.
//
// Solidity: function getActualLockedBond(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetActualLockedBond(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetActualLockedBond(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetBond(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getBond", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetBond(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetBond(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetBond(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetBond(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetBondAmountByKeysCount is a free data retrieval call binding the contract method 0x0f23e742.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, (uint256[],uint256) curve) pure returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetBondAmountByKeysCount(opts *bind.CallOpts, keys *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getBondAmountByKeysCount", keys, curve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondAmountByKeysCount is a free data retrieval call binding the contract method 0x0f23e742.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, (uint256[],uint256) curve) pure returns(uint256)
func (_CSAccounting *CSAccountingSession) GetBondAmountByKeysCount(keys *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	return _CSAccounting.Contract.GetBondAmountByKeysCount(&_CSAccounting.CallOpts, keys, curve)
}

// GetBondAmountByKeysCount is a free data retrieval call binding the contract method 0x0f23e742.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, (uint256[],uint256) curve) pure returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetBondAmountByKeysCount(keys *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	return _CSAccounting.Contract.GetBondAmountByKeysCount(&_CSAccounting.CallOpts, keys, curve)
}

// GetBondAmountByKeysCount0 is a free data retrieval call binding the contract method 0x546da24f.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, uint256 curveId) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetBondAmountByKeysCount0(opts *bind.CallOpts, keys *big.Int, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getBondAmountByKeysCount0", keys, curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondAmountByKeysCount0 is a free data retrieval call binding the contract method 0x546da24f.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, uint256 curveId) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetBondAmountByKeysCount0(keys *big.Int, curveId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetBondAmountByKeysCount0(&_CSAccounting.CallOpts, keys, curveId)
}

// GetBondAmountByKeysCount0 is a free data retrieval call binding the contract method 0x546da24f.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, uint256 curveId) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetBondAmountByKeysCount0(keys *big.Int, curveId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetBondAmountByKeysCount0(&_CSAccounting.CallOpts, keys, curveId)
}

// GetBondAmountByKeysCountWstETH is a free data retrieval call binding the contract method 0x13d1234b.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, uint256 curveId) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetBondAmountByKeysCountWstETH(opts *bind.CallOpts, keysCount *big.Int, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getBondAmountByKeysCountWstETH", keysCount, curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondAmountByKeysCountWstETH is a free data retrieval call binding the contract method 0x13d1234b.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, uint256 curveId) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetBondAmountByKeysCountWstETH(keysCount *big.Int, curveId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetBondAmountByKeysCountWstETH(&_CSAccounting.CallOpts, keysCount, curveId)
}

// GetBondAmountByKeysCountWstETH is a free data retrieval call binding the contract method 0x13d1234b.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, uint256 curveId) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetBondAmountByKeysCountWstETH(keysCount *big.Int, curveId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetBondAmountByKeysCountWstETH(&_CSAccounting.CallOpts, keysCount, curveId)
}

// GetBondAmountByKeysCountWstETH0 is a free data retrieval call binding the contract method 0x9a4df8f0.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, (uint256[],uint256) curve) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetBondAmountByKeysCountWstETH0(opts *bind.CallOpts, keysCount *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getBondAmountByKeysCountWstETH0", keysCount, curve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondAmountByKeysCountWstETH0 is a free data retrieval call binding the contract method 0x9a4df8f0.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, (uint256[],uint256) curve) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetBondAmountByKeysCountWstETH0(keysCount *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	return _CSAccounting.Contract.GetBondAmountByKeysCountWstETH0(&_CSAccounting.CallOpts, keysCount, curve)
}

// GetBondAmountByKeysCountWstETH0 is a free data retrieval call binding the contract method 0x9a4df8f0.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, (uint256[],uint256) curve) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetBondAmountByKeysCountWstETH0(keysCount *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	return _CSAccounting.Contract.GetBondAmountByKeysCountWstETH0(&_CSAccounting.CallOpts, keysCount, curve)
}

// GetBondCurve is a free data retrieval call binding the contract method 0x6e13f099.
//
// Solidity: function getBondCurve(uint256 nodeOperatorId) view returns((uint256[],uint256))
func (_CSAccounting *CSAccountingCaller) GetBondCurve(opts *bind.CallOpts, nodeOperatorId *big.Int) (ICSBondCurveBondCurve, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getBondCurve", nodeOperatorId)

	if err != nil {
		return *new(ICSBondCurveBondCurve), err
	}

	out0 := *abi.ConvertType(out[0], new(ICSBondCurveBondCurve)).(*ICSBondCurveBondCurve)

	return out0, err

}

// GetBondCurve is a free data retrieval call binding the contract method 0x6e13f099.
//
// Solidity: function getBondCurve(uint256 nodeOperatorId) view returns((uint256[],uint256))
func (_CSAccounting *CSAccountingSession) GetBondCurve(nodeOperatorId *big.Int) (ICSBondCurveBondCurve, error) {
	return _CSAccounting.Contract.GetBondCurve(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetBondCurve is a free data retrieval call binding the contract method 0x6e13f099.
//
// Solidity: function getBondCurve(uint256 nodeOperatorId) view returns((uint256[],uint256))
func (_CSAccounting *CSAccountingCallerSession) GetBondCurve(nodeOperatorId *big.Int) (ICSBondCurveBondCurve, error) {
	return _CSAccounting.Contract.GetBondCurve(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetBondCurveId is a free data retrieval call binding the contract method 0x0569b947.
//
// Solidity: function getBondCurveId(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetBondCurveId(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getBondCurveId", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondCurveId is a free data retrieval call binding the contract method 0x0569b947.
//
// Solidity: function getBondCurveId(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetBondCurveId(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetBondCurveId(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetBondCurveId is a free data retrieval call binding the contract method 0x0569b947.
//
// Solidity: function getBondCurveId(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetBondCurveId(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetBondCurveId(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetBondLockRetentionPeriod is a free data retrieval call binding the contract method 0xdef82d02.
//
// Solidity: function getBondLockRetentionPeriod() view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetBondLockRetentionPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getBondLockRetentionPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondLockRetentionPeriod is a free data retrieval call binding the contract method 0xdef82d02.
//
// Solidity: function getBondLockRetentionPeriod() view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetBondLockRetentionPeriod() (*big.Int, error) {
	return _CSAccounting.Contract.GetBondLockRetentionPeriod(&_CSAccounting.CallOpts)
}

// GetBondLockRetentionPeriod is a free data retrieval call binding the contract method 0xdef82d02.
//
// Solidity: function getBondLockRetentionPeriod() view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetBondLockRetentionPeriod() (*big.Int, error) {
	return _CSAccounting.Contract.GetBondLockRetentionPeriod(&_CSAccounting.CallOpts)
}

// GetBondShares is a free data retrieval call binding the contract method 0x06cd0e90.
//
// Solidity: function getBondShares(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetBondShares(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getBondShares", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondShares is a free data retrieval call binding the contract method 0x06cd0e90.
//
// Solidity: function getBondShares(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetBondShares(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetBondShares(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetBondShares is a free data retrieval call binding the contract method 0x06cd0e90.
//
// Solidity: function getBondShares(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetBondShares(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetBondShares(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetBondSummary is a free data retrieval call binding the contract method 0xce19793f.
//
// Solidity: function getBondSummary(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_CSAccounting *CSAccountingCaller) GetBondSummary(opts *bind.CallOpts, nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getBondSummary", nodeOperatorId)

	outstruct := new(struct {
		Current  *big.Int
		Required *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Current = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Required = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBondSummary is a free data retrieval call binding the contract method 0xce19793f.
//
// Solidity: function getBondSummary(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_CSAccounting *CSAccountingSession) GetBondSummary(nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	return _CSAccounting.Contract.GetBondSummary(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetBondSummary is a free data retrieval call binding the contract method 0xce19793f.
//
// Solidity: function getBondSummary(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_CSAccounting *CSAccountingCallerSession) GetBondSummary(nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	return _CSAccounting.Contract.GetBondSummary(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetBondSummaryShares is a free data retrieval call binding the contract method 0x8f6549ae.
//
// Solidity: function getBondSummaryShares(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_CSAccounting *CSAccountingCaller) GetBondSummaryShares(opts *bind.CallOpts, nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getBondSummaryShares", nodeOperatorId)

	outstruct := new(struct {
		Current  *big.Int
		Required *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Current = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Required = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBondSummaryShares is a free data retrieval call binding the contract method 0x8f6549ae.
//
// Solidity: function getBondSummaryShares(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_CSAccounting *CSAccountingSession) GetBondSummaryShares(nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	return _CSAccounting.Contract.GetBondSummaryShares(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetBondSummaryShares is a free data retrieval call binding the contract method 0x8f6549ae.
//
// Solidity: function getBondSummaryShares(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_CSAccounting *CSAccountingCallerSession) GetBondSummaryShares(nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	return _CSAccounting.Contract.GetBondSummaryShares(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetCurveInfo is a free data retrieval call binding the contract method 0xb5b624bf.
//
// Solidity: function getCurveInfo(uint256 curveId) view returns((uint256[],uint256))
func (_CSAccounting *CSAccountingCaller) GetCurveInfo(opts *bind.CallOpts, curveId *big.Int) (ICSBondCurveBondCurve, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getCurveInfo", curveId)

	if err != nil {
		return *new(ICSBondCurveBondCurve), err
	}

	out0 := *abi.ConvertType(out[0], new(ICSBondCurveBondCurve)).(*ICSBondCurveBondCurve)

	return out0, err

}

// GetCurveInfo is a free data retrieval call binding the contract method 0xb5b624bf.
//
// Solidity: function getCurveInfo(uint256 curveId) view returns((uint256[],uint256))
func (_CSAccounting *CSAccountingSession) GetCurveInfo(curveId *big.Int) (ICSBondCurveBondCurve, error) {
	return _CSAccounting.Contract.GetCurveInfo(&_CSAccounting.CallOpts, curveId)
}

// GetCurveInfo is a free data retrieval call binding the contract method 0xb5b624bf.
//
// Solidity: function getCurveInfo(uint256 curveId) view returns((uint256[],uint256))
func (_CSAccounting *CSAccountingCallerSession) GetCurveInfo(curveId *big.Int) (ICSBondCurveBondCurve, error) {
	return _CSAccounting.Contract.GetCurveInfo(&_CSAccounting.CallOpts, curveId)
}

// GetKeysCountByBondAmount is a free data retrieval call binding the contract method 0xd2fa16a6.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, (uint256[],uint256) curve) pure returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetKeysCountByBondAmount(opts *bind.CallOpts, amount *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getKeysCountByBondAmount", amount, curve)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeysCountByBondAmount is a free data retrieval call binding the contract method 0xd2fa16a6.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, (uint256[],uint256) curve) pure returns(uint256)
func (_CSAccounting *CSAccountingSession) GetKeysCountByBondAmount(amount *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	return _CSAccounting.Contract.GetKeysCountByBondAmount(&_CSAccounting.CallOpts, amount, curve)
}

// GetKeysCountByBondAmount is a free data retrieval call binding the contract method 0xd2fa16a6.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, (uint256[],uint256) curve) pure returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetKeysCountByBondAmount(amount *big.Int, curve ICSBondCurveBondCurve) (*big.Int, error) {
	return _CSAccounting.Contract.GetKeysCountByBondAmount(&_CSAccounting.CallOpts, amount, curve)
}

// GetKeysCountByBondAmount0 is a free data retrieval call binding the contract method 0xdc38ea3d.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, uint256 curveId) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetKeysCountByBondAmount0(opts *bind.CallOpts, amount *big.Int, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getKeysCountByBondAmount0", amount, curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeysCountByBondAmount0 is a free data retrieval call binding the contract method 0xdc38ea3d.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, uint256 curveId) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetKeysCountByBondAmount0(amount *big.Int, curveId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetKeysCountByBondAmount0(&_CSAccounting.CallOpts, amount, curveId)
}

// GetKeysCountByBondAmount0 is a free data retrieval call binding the contract method 0xdc38ea3d.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, uint256 curveId) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetKeysCountByBondAmount0(amount *big.Int, curveId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetKeysCountByBondAmount0(&_CSAccounting.CallOpts, amount, curveId)
}

// GetLockedBondInfo is a free data retrieval call binding the contract method 0x83316184.
//
// Solidity: function getLockedBondInfo(uint256 nodeOperatorId) view returns((uint128,uint128))
func (_CSAccounting *CSAccountingCaller) GetLockedBondInfo(opts *bind.CallOpts, nodeOperatorId *big.Int) (ICSBondLockBondLock, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getLockedBondInfo", nodeOperatorId)

	if err != nil {
		return *new(ICSBondLockBondLock), err
	}

	out0 := *abi.ConvertType(out[0], new(ICSBondLockBondLock)).(*ICSBondLockBondLock)

	return out0, err

}

// GetLockedBondInfo is a free data retrieval call binding the contract method 0x83316184.
//
// Solidity: function getLockedBondInfo(uint256 nodeOperatorId) view returns((uint128,uint128))
func (_CSAccounting *CSAccountingSession) GetLockedBondInfo(nodeOperatorId *big.Int) (ICSBondLockBondLock, error) {
	return _CSAccounting.Contract.GetLockedBondInfo(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetLockedBondInfo is a free data retrieval call binding the contract method 0x83316184.
//
// Solidity: function getLockedBondInfo(uint256 nodeOperatorId) view returns((uint128,uint128))
func (_CSAccounting *CSAccountingCallerSession) GetLockedBondInfo(nodeOperatorId *big.Int) (ICSBondLockBondLock, error) {
	return _CSAccounting.Contract.GetLockedBondInfo(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetRequiredBondForNextKeys is a free data retrieval call binding the contract method 0xb148db6a.
//
// Solidity: function getRequiredBondForNextKeys(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetRequiredBondForNextKeys(opts *bind.CallOpts, nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getRequiredBondForNextKeys", nodeOperatorId, additionalKeys)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredBondForNextKeys is a free data retrieval call binding the contract method 0xb148db6a.
//
// Solidity: function getRequiredBondForNextKeys(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetRequiredBondForNextKeys(nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetRequiredBondForNextKeys(&_CSAccounting.CallOpts, nodeOperatorId, additionalKeys)
}

// GetRequiredBondForNextKeys is a free data retrieval call binding the contract method 0xb148db6a.
//
// Solidity: function getRequiredBondForNextKeys(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetRequiredBondForNextKeys(nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetRequiredBondForNextKeys(&_CSAccounting.CallOpts, nodeOperatorId, additionalKeys)
}

// GetRequiredBondForNextKeysWstETH is a free data retrieval call binding the contract method 0x28846981.
//
// Solidity: function getRequiredBondForNextKeysWstETH(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetRequiredBondForNextKeysWstETH(opts *bind.CallOpts, nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getRequiredBondForNextKeysWstETH", nodeOperatorId, additionalKeys)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredBondForNextKeysWstETH is a free data retrieval call binding the contract method 0x28846981.
//
// Solidity: function getRequiredBondForNextKeysWstETH(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetRequiredBondForNextKeysWstETH(nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetRequiredBondForNextKeysWstETH(&_CSAccounting.CallOpts, nodeOperatorId, additionalKeys)
}

// GetRequiredBondForNextKeysWstETH is a free data retrieval call binding the contract method 0x28846981.
//
// Solidity: function getRequiredBondForNextKeysWstETH(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetRequiredBondForNextKeysWstETH(nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetRequiredBondForNextKeysWstETH(&_CSAccounting.CallOpts, nodeOperatorId, additionalKeys)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetResumeSinceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getResumeSinceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _CSAccounting.Contract.GetResumeSinceTimestamp(&_CSAccounting.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _CSAccounting.Contract.GetResumeSinceTimestamp(&_CSAccounting.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CSAccounting *CSAccountingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CSAccounting *CSAccountingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CSAccounting.Contract.GetRoleAdmin(&_CSAccounting.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_CSAccounting *CSAccountingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _CSAccounting.Contract.GetRoleAdmin(&_CSAccounting.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CSAccounting *CSAccountingCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CSAccounting *CSAccountingSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CSAccounting.Contract.GetRoleMember(&_CSAccounting.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_CSAccounting *CSAccountingCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _CSAccounting.Contract.GetRoleMember(&_CSAccounting.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CSAccounting.Contract.GetRoleMemberCount(&_CSAccounting.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _CSAccounting.Contract.GetRoleMemberCount(&_CSAccounting.CallOpts, role)
}

// GetUnbondedKeysCount is a free data retrieval call binding the contract method 0x01a5e9e3.
//
// Solidity: function getUnbondedKeysCount(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetUnbondedKeysCount(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getUnbondedKeysCount", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnbondedKeysCount is a free data retrieval call binding the contract method 0x01a5e9e3.
//
// Solidity: function getUnbondedKeysCount(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetUnbondedKeysCount(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetUnbondedKeysCount(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetUnbondedKeysCount is a free data retrieval call binding the contract method 0x01a5e9e3.
//
// Solidity: function getUnbondedKeysCount(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetUnbondedKeysCount(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetUnbondedKeysCount(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetUnbondedKeysCountToEject is a free data retrieval call binding the contract method 0x9c516102.
//
// Solidity: function getUnbondedKeysCountToEject(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingCaller) GetUnbondedKeysCountToEject(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "getUnbondedKeysCountToEject", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnbondedKeysCountToEject is a free data retrieval call binding the contract method 0x9c516102.
//
// Solidity: function getUnbondedKeysCountToEject(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingSession) GetUnbondedKeysCountToEject(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetUnbondedKeysCountToEject(&_CSAccounting.CallOpts, nodeOperatorId)
}

// GetUnbondedKeysCountToEject is a free data retrieval call binding the contract method 0x9c516102.
//
// Solidity: function getUnbondedKeysCountToEject(uint256 nodeOperatorId) view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) GetUnbondedKeysCountToEject(nodeOperatorId *big.Int) (*big.Int, error) {
	return _CSAccounting.Contract.GetUnbondedKeysCountToEject(&_CSAccounting.CallOpts, nodeOperatorId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CSAccounting *CSAccountingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CSAccounting *CSAccountingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CSAccounting.Contract.HasRole(&_CSAccounting.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_CSAccounting *CSAccountingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _CSAccounting.Contract.HasRole(&_CSAccounting.CallOpts, role, account)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CSAccounting *CSAccountingCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CSAccounting *CSAccountingSession) IsPaused() (bool, error) {
	return _CSAccounting.Contract.IsPaused(&_CSAccounting.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_CSAccounting *CSAccountingCallerSession) IsPaused() (bool, error) {
	return _CSAccounting.Contract.IsPaused(&_CSAccounting.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CSAccounting *CSAccountingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CSAccounting *CSAccountingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CSAccounting.Contract.SupportsInterface(&_CSAccounting.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_CSAccounting *CSAccountingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _CSAccounting.Contract.SupportsInterface(&_CSAccounting.CallOpts, interfaceId)
}

// TotalBondShares is a free data retrieval call binding the contract method 0x74d70aea.
//
// Solidity: function totalBondShares() view returns(uint256)
func (_CSAccounting *CSAccountingCaller) TotalBondShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CSAccounting.contract.Call(opts, &out, "totalBondShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBondShares is a free data retrieval call binding the contract method 0x74d70aea.
//
// Solidity: function totalBondShares() view returns(uint256)
func (_CSAccounting *CSAccountingSession) TotalBondShares() (*big.Int, error) {
	return _CSAccounting.Contract.TotalBondShares(&_CSAccounting.CallOpts)
}

// TotalBondShares is a free data retrieval call binding the contract method 0x74d70aea.
//
// Solidity: function totalBondShares() view returns(uint256)
func (_CSAccounting *CSAccountingCallerSession) TotalBondShares() (*big.Int, error) {
	return _CSAccounting.Contract.TotalBondShares(&_CSAccounting.CallOpts)
}

// AddBondCurve is a paid mutator transaction binding the contract method 0x573b6245.
//
// Solidity: function addBondCurve(uint256[] bondCurve) returns(uint256 id)
func (_CSAccounting *CSAccountingTransactor) AddBondCurve(opts *bind.TransactOpts, bondCurve []*big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "addBondCurve", bondCurve)
}

// AddBondCurve is a paid mutator transaction binding the contract method 0x573b6245.
//
// Solidity: function addBondCurve(uint256[] bondCurve) returns(uint256 id)
func (_CSAccounting *CSAccountingSession) AddBondCurve(bondCurve []*big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.AddBondCurve(&_CSAccounting.TransactOpts, bondCurve)
}

// AddBondCurve is a paid mutator transaction binding the contract method 0x573b6245.
//
// Solidity: function addBondCurve(uint256[] bondCurve) returns(uint256 id)
func (_CSAccounting *CSAccountingTransactorSession) AddBondCurve(bondCurve []*big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.AddBondCurve(&_CSAccounting.TransactOpts, bondCurve)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x881fa03c.
//
// Solidity: function chargeFee(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSAccounting *CSAccountingTransactor) ChargeFee(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "chargeFee", nodeOperatorId, amount)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x881fa03c.
//
// Solidity: function chargeFee(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSAccounting *CSAccountingSession) ChargeFee(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.ChargeFee(&_CSAccounting.TransactOpts, nodeOperatorId, amount)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x881fa03c.
//
// Solidity: function chargeFee(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSAccounting *CSAccountingTransactorSession) ChargeFee(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.ChargeFee(&_CSAccounting.TransactOpts, nodeOperatorId, amount)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0xf9391223.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSAccounting *CSAccountingTransactor) ClaimRewardsStETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, stETHAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "claimRewardsStETH", nodeOperatorId, stETHAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0xf9391223.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSAccounting *CSAccountingSession) ClaimRewardsStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSAccounting.Contract.ClaimRewardsStETH(&_CSAccounting.TransactOpts, nodeOperatorId, stETHAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0xf9391223.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSAccounting *CSAccountingTransactorSession) ClaimRewardsStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSAccounting.Contract.ClaimRewardsStETH(&_CSAccounting.TransactOpts, nodeOperatorId, stETHAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0xcc810cb9.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stEthAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSAccounting *CSAccountingTransactor) ClaimRewardsUnstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, stEthAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "claimRewardsUnstETH", nodeOperatorId, stEthAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0xcc810cb9.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stEthAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSAccounting *CSAccountingSession) ClaimRewardsUnstETH(nodeOperatorId *big.Int, stEthAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSAccounting.Contract.ClaimRewardsUnstETH(&_CSAccounting.TransactOpts, nodeOperatorId, stEthAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0xcc810cb9.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stEthAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSAccounting *CSAccountingTransactorSession) ClaimRewardsUnstETH(nodeOperatorId *big.Int, stEthAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSAccounting.Contract.ClaimRewardsUnstETH(&_CSAccounting.TransactOpts, nodeOperatorId, stEthAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x70903eb9.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSAccounting *CSAccountingTransactor) ClaimRewardsWstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, wstETHAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "claimRewardsWstETH", nodeOperatorId, wstETHAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x70903eb9.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSAccounting *CSAccountingSession) ClaimRewardsWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSAccounting.Contract.ClaimRewardsWstETH(&_CSAccounting.TransactOpts, nodeOperatorId, wstETHAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x70903eb9.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, address rewardAddress, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSAccounting *CSAccountingTransactorSession) ClaimRewardsWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, rewardAddress common.Address, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSAccounting.Contract.ClaimRewardsWstETH(&_CSAccounting.TransactOpts, nodeOperatorId, wstETHAmount, rewardAddress, cumulativeFeeShares, rewardsProof)
}

// CompensateLockedBondETH is a paid mutator transaction binding the contract method 0x15b5c477.
//
// Solidity: function compensateLockedBondETH(uint256 nodeOperatorId) payable returns()
func (_CSAccounting *CSAccountingTransactor) CompensateLockedBondETH(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "compensateLockedBondETH", nodeOperatorId)
}

// CompensateLockedBondETH is a paid mutator transaction binding the contract method 0x15b5c477.
//
// Solidity: function compensateLockedBondETH(uint256 nodeOperatorId) payable returns()
func (_CSAccounting *CSAccountingSession) CompensateLockedBondETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.CompensateLockedBondETH(&_CSAccounting.TransactOpts, nodeOperatorId)
}

// CompensateLockedBondETH is a paid mutator transaction binding the contract method 0x15b5c477.
//
// Solidity: function compensateLockedBondETH(uint256 nodeOperatorId) payable returns()
func (_CSAccounting *CSAccountingTransactorSession) CompensateLockedBondETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.CompensateLockedBondETH(&_CSAccounting.TransactOpts, nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address from, uint256 nodeOperatorId) payable returns()
func (_CSAccounting *CSAccountingTransactor) DepositETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "depositETH", from, nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address from, uint256 nodeOperatorId) payable returns()
func (_CSAccounting *CSAccountingSession) DepositETH(from common.Address, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.DepositETH(&_CSAccounting.TransactOpts, from, nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address from, uint256 nodeOperatorId) payable returns()
func (_CSAccounting *CSAccountingTransactorSession) DepositETH(from common.Address, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.DepositETH(&_CSAccounting.TransactOpts, from, nodeOperatorId)
}

// DepositStETH is a paid mutator transaction binding the contract method 0x4c7ed3d2.
//
// Solidity: function depositStETH(address from, uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSAccounting *CSAccountingTransactor) DepositStETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "depositStETH", from, nodeOperatorId, stETHAmount, permit)
}

// DepositStETH is a paid mutator transaction binding the contract method 0x4c7ed3d2.
//
// Solidity: function depositStETH(address from, uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSAccounting *CSAccountingSession) DepositStETH(from common.Address, nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSAccounting.Contract.DepositStETH(&_CSAccounting.TransactOpts, from, nodeOperatorId, stETHAmount, permit)
}

// DepositStETH is a paid mutator transaction binding the contract method 0x4c7ed3d2.
//
// Solidity: function depositStETH(address from, uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSAccounting *CSAccountingTransactorSession) DepositStETH(from common.Address, nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSAccounting.Contract.DepositStETH(&_CSAccounting.TransactOpts, from, nodeOperatorId, stETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0xf7966efe.
//
// Solidity: function depositWstETH(address from, uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSAccounting *CSAccountingTransactor) DepositWstETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "depositWstETH", from, nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0xf7966efe.
//
// Solidity: function depositWstETH(address from, uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSAccounting *CSAccountingSession) DepositWstETH(from common.Address, nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSAccounting.Contract.DepositWstETH(&_CSAccounting.TransactOpts, from, nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0xf7966efe.
//
// Solidity: function depositWstETH(address from, uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_CSAccounting *CSAccountingTransactorSession) DepositWstETH(from common.Address, nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _CSAccounting.Contract.DepositWstETH(&_CSAccounting.TransactOpts, from, nodeOperatorId, wstETHAmount, permit)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CSAccounting *CSAccountingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CSAccounting *CSAccountingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CSAccounting.Contract.GrantRole(&_CSAccounting.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_CSAccounting *CSAccountingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CSAccounting.Contract.GrantRole(&_CSAccounting.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xfab382f1.
//
// Solidity: function initialize(uint256[] bondCurve, address admin, address _feeDistributor, uint256 bondLockRetentionPeriod, address _chargePenaltyRecipient) returns()
func (_CSAccounting *CSAccountingTransactor) Initialize(opts *bind.TransactOpts, bondCurve []*big.Int, admin common.Address, _feeDistributor common.Address, bondLockRetentionPeriod *big.Int, _chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "initialize", bondCurve, admin, _feeDistributor, bondLockRetentionPeriod, _chargePenaltyRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xfab382f1.
//
// Solidity: function initialize(uint256[] bondCurve, address admin, address _feeDistributor, uint256 bondLockRetentionPeriod, address _chargePenaltyRecipient) returns()
func (_CSAccounting *CSAccountingSession) Initialize(bondCurve []*big.Int, admin common.Address, _feeDistributor common.Address, bondLockRetentionPeriod *big.Int, _chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _CSAccounting.Contract.Initialize(&_CSAccounting.TransactOpts, bondCurve, admin, _feeDistributor, bondLockRetentionPeriod, _chargePenaltyRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xfab382f1.
//
// Solidity: function initialize(uint256[] bondCurve, address admin, address _feeDistributor, uint256 bondLockRetentionPeriod, address _chargePenaltyRecipient) returns()
func (_CSAccounting *CSAccountingTransactorSession) Initialize(bondCurve []*big.Int, admin common.Address, _feeDistributor common.Address, bondLockRetentionPeriod *big.Int, _chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _CSAccounting.Contract.Initialize(&_CSAccounting.TransactOpts, bondCurve, admin, _feeDistributor, bondLockRetentionPeriod, _chargePenaltyRecipient)
}

// LockBondETH is a paid mutator transaction binding the contract method 0xdcab7f83.
//
// Solidity: function lockBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSAccounting *CSAccountingTransactor) LockBondETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "lockBondETH", nodeOperatorId, amount)
}

// LockBondETH is a paid mutator transaction binding the contract method 0xdcab7f83.
//
// Solidity: function lockBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSAccounting *CSAccountingSession) LockBondETH(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.LockBondETH(&_CSAccounting.TransactOpts, nodeOperatorId, amount)
}

// LockBondETH is a paid mutator transaction binding the contract method 0xdcab7f83.
//
// Solidity: function lockBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSAccounting *CSAccountingTransactorSession) LockBondETH(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.LockBondETH(&_CSAccounting.TransactOpts, nodeOperatorId, amount)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_CSAccounting *CSAccountingTransactor) PauseFor(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "pauseFor", duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_CSAccounting *CSAccountingSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.PauseFor(&_CSAccounting.TransactOpts, duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_CSAccounting *CSAccountingTransactorSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.PauseFor(&_CSAccounting.TransactOpts, duration)
}

// Penalize is a paid mutator transaction binding the contract method 0xe5220e3f.
//
// Solidity: function penalize(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSAccounting *CSAccountingTransactor) Penalize(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "penalize", nodeOperatorId, amount)
}

// Penalize is a paid mutator transaction binding the contract method 0xe5220e3f.
//
// Solidity: function penalize(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSAccounting *CSAccountingSession) Penalize(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.Penalize(&_CSAccounting.TransactOpts, nodeOperatorId, amount)
}

// Penalize is a paid mutator transaction binding the contract method 0xe5220e3f.
//
// Solidity: function penalize(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSAccounting *CSAccountingTransactorSession) Penalize(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.Penalize(&_CSAccounting.TransactOpts, nodeOperatorId, amount)
}

// PullFeeRewards is a paid mutator transaction binding the contract method 0x9b4c6c27.
//
// Solidity: function pullFeeRewards(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSAccounting *CSAccountingTransactor) PullFeeRewards(opts *bind.TransactOpts, nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "pullFeeRewards", nodeOperatorId, cumulativeFeeShares, rewardsProof)
}

// PullFeeRewards is a paid mutator transaction binding the contract method 0x9b4c6c27.
//
// Solidity: function pullFeeRewards(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSAccounting *CSAccountingSession) PullFeeRewards(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSAccounting.Contract.PullFeeRewards(&_CSAccounting.TransactOpts, nodeOperatorId, cumulativeFeeShares, rewardsProof)
}

// PullFeeRewards is a paid mutator transaction binding the contract method 0x9b4c6c27.
//
// Solidity: function pullFeeRewards(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_CSAccounting *CSAccountingTransactorSession) PullFeeRewards(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _CSAccounting.Contract.PullFeeRewards(&_CSAccounting.TransactOpts, nodeOperatorId, cumulativeFeeShares, rewardsProof)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_CSAccounting *CSAccountingTransactor) RecoverERC1155(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "recoverERC1155", token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_CSAccounting *CSAccountingSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.RecoverERC1155(&_CSAccounting.TransactOpts, token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_CSAccounting *CSAccountingTransactorSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.RecoverERC1155(&_CSAccounting.TransactOpts, token, tokenId)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_CSAccounting *CSAccountingTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "recoverERC20", token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_CSAccounting *CSAccountingSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.RecoverERC20(&_CSAccounting.TransactOpts, token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_CSAccounting *CSAccountingTransactorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.RecoverERC20(&_CSAccounting.TransactOpts, token, amount)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_CSAccounting *CSAccountingTransactor) RecoverERC721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "recoverERC721", token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_CSAccounting *CSAccountingSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.RecoverERC721(&_CSAccounting.TransactOpts, token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_CSAccounting *CSAccountingTransactorSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.RecoverERC721(&_CSAccounting.TransactOpts, token, tokenId)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_CSAccounting *CSAccountingTransactor) RecoverEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "recoverEther")
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_CSAccounting *CSAccountingSession) RecoverEther() (*types.Transaction, error) {
	return _CSAccounting.Contract.RecoverEther(&_CSAccounting.TransactOpts)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_CSAccounting *CSAccountingTransactorSession) RecoverEther() (*types.Transaction, error) {
	return _CSAccounting.Contract.RecoverEther(&_CSAccounting.TransactOpts)
}

// RecoverStETHShares is a paid mutator transaction binding the contract method 0x5a73bdc8.
//
// Solidity: function recoverStETHShares() returns()
func (_CSAccounting *CSAccountingTransactor) RecoverStETHShares(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "recoverStETHShares")
}

// RecoverStETHShares is a paid mutator transaction binding the contract method 0x5a73bdc8.
//
// Solidity: function recoverStETHShares() returns()
func (_CSAccounting *CSAccountingSession) RecoverStETHShares() (*types.Transaction, error) {
	return _CSAccounting.Contract.RecoverStETHShares(&_CSAccounting.TransactOpts)
}

// RecoverStETHShares is a paid mutator transaction binding the contract method 0x5a73bdc8.
//
// Solidity: function recoverStETHShares() returns()
func (_CSAccounting *CSAccountingTransactorSession) RecoverStETHShares() (*types.Transaction, error) {
	return _CSAccounting.Contract.RecoverStETHShares(&_CSAccounting.TransactOpts)
}

// ReleaseLockedBondETH is a paid mutator transaction binding the contract method 0xd963ae55.
//
// Solidity: function releaseLockedBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSAccounting *CSAccountingTransactor) ReleaseLockedBondETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "releaseLockedBondETH", nodeOperatorId, amount)
}

// ReleaseLockedBondETH is a paid mutator transaction binding the contract method 0xd963ae55.
//
// Solidity: function releaseLockedBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSAccounting *CSAccountingSession) ReleaseLockedBondETH(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.ReleaseLockedBondETH(&_CSAccounting.TransactOpts, nodeOperatorId, amount)
}

// ReleaseLockedBondETH is a paid mutator transaction binding the contract method 0xd963ae55.
//
// Solidity: function releaseLockedBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_CSAccounting *CSAccountingTransactorSession) ReleaseLockedBondETH(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.ReleaseLockedBondETH(&_CSAccounting.TransactOpts, nodeOperatorId, amount)
}

// RenewBurnerAllowance is a paid mutator transaction binding the contract method 0xf3efecc4.
//
// Solidity: function renewBurnerAllowance() returns()
func (_CSAccounting *CSAccountingTransactor) RenewBurnerAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "renewBurnerAllowance")
}

// RenewBurnerAllowance is a paid mutator transaction binding the contract method 0xf3efecc4.
//
// Solidity: function renewBurnerAllowance() returns()
func (_CSAccounting *CSAccountingSession) RenewBurnerAllowance() (*types.Transaction, error) {
	return _CSAccounting.Contract.RenewBurnerAllowance(&_CSAccounting.TransactOpts)
}

// RenewBurnerAllowance is a paid mutator transaction binding the contract method 0xf3efecc4.
//
// Solidity: function renewBurnerAllowance() returns()
func (_CSAccounting *CSAccountingTransactorSession) RenewBurnerAllowance() (*types.Transaction, error) {
	return _CSAccounting.Contract.RenewBurnerAllowance(&_CSAccounting.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CSAccounting *CSAccountingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CSAccounting *CSAccountingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CSAccounting.Contract.RenounceRole(&_CSAccounting.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_CSAccounting *CSAccountingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _CSAccounting.Contract.RenounceRole(&_CSAccounting.TransactOpts, role, callerConfirmation)
}

// ResetBondCurve is a paid mutator transaction binding the contract method 0x449add1b.
//
// Solidity: function resetBondCurve(uint256 nodeOperatorId) returns()
func (_CSAccounting *CSAccountingTransactor) ResetBondCurve(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "resetBondCurve", nodeOperatorId)
}

// ResetBondCurve is a paid mutator transaction binding the contract method 0x449add1b.
//
// Solidity: function resetBondCurve(uint256 nodeOperatorId) returns()
func (_CSAccounting *CSAccountingSession) ResetBondCurve(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.ResetBondCurve(&_CSAccounting.TransactOpts, nodeOperatorId)
}

// ResetBondCurve is a paid mutator transaction binding the contract method 0x449add1b.
//
// Solidity: function resetBondCurve(uint256 nodeOperatorId) returns()
func (_CSAccounting *CSAccountingTransactorSession) ResetBondCurve(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.ResetBondCurve(&_CSAccounting.TransactOpts, nodeOperatorId)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_CSAccounting *CSAccountingTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_CSAccounting *CSAccountingSession) Resume() (*types.Transaction, error) {
	return _CSAccounting.Contract.Resume(&_CSAccounting.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_CSAccounting *CSAccountingTransactorSession) Resume() (*types.Transaction, error) {
	return _CSAccounting.Contract.Resume(&_CSAccounting.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CSAccounting *CSAccountingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CSAccounting *CSAccountingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CSAccounting.Contract.RevokeRole(&_CSAccounting.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_CSAccounting *CSAccountingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _CSAccounting.Contract.RevokeRole(&_CSAccounting.TransactOpts, role, account)
}

// SetBondCurve is a paid mutator transaction binding the contract method 0xb2d03e4d.
//
// Solidity: function setBondCurve(uint256 nodeOperatorId, uint256 curveId) returns()
func (_CSAccounting *CSAccountingTransactor) SetBondCurve(opts *bind.TransactOpts, nodeOperatorId *big.Int, curveId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "setBondCurve", nodeOperatorId, curveId)
}

// SetBondCurve is a paid mutator transaction binding the contract method 0xb2d03e4d.
//
// Solidity: function setBondCurve(uint256 nodeOperatorId, uint256 curveId) returns()
func (_CSAccounting *CSAccountingSession) SetBondCurve(nodeOperatorId *big.Int, curveId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.SetBondCurve(&_CSAccounting.TransactOpts, nodeOperatorId, curveId)
}

// SetBondCurve is a paid mutator transaction binding the contract method 0xb2d03e4d.
//
// Solidity: function setBondCurve(uint256 nodeOperatorId, uint256 curveId) returns()
func (_CSAccounting *CSAccountingTransactorSession) SetBondCurve(nodeOperatorId *big.Int, curveId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.SetBondCurve(&_CSAccounting.TransactOpts, nodeOperatorId, curveId)
}

// SetChargePenaltyRecipient is a paid mutator transaction binding the contract method 0x433cd6c3.
//
// Solidity: function setChargePenaltyRecipient(address _chargePenaltyRecipient) returns()
func (_CSAccounting *CSAccountingTransactor) SetChargePenaltyRecipient(opts *bind.TransactOpts, _chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "setChargePenaltyRecipient", _chargePenaltyRecipient)
}

// SetChargePenaltyRecipient is a paid mutator transaction binding the contract method 0x433cd6c3.
//
// Solidity: function setChargePenaltyRecipient(address _chargePenaltyRecipient) returns()
func (_CSAccounting *CSAccountingSession) SetChargePenaltyRecipient(_chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _CSAccounting.Contract.SetChargePenaltyRecipient(&_CSAccounting.TransactOpts, _chargePenaltyRecipient)
}

// SetChargePenaltyRecipient is a paid mutator transaction binding the contract method 0x433cd6c3.
//
// Solidity: function setChargePenaltyRecipient(address _chargePenaltyRecipient) returns()
func (_CSAccounting *CSAccountingTransactorSession) SetChargePenaltyRecipient(_chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _CSAccounting.Contract.SetChargePenaltyRecipient(&_CSAccounting.TransactOpts, _chargePenaltyRecipient)
}

// SetLockedBondRetentionPeriod is a paid mutator transaction binding the contract method 0x99965225.
//
// Solidity: function setLockedBondRetentionPeriod(uint256 retention) returns()
func (_CSAccounting *CSAccountingTransactor) SetLockedBondRetentionPeriod(opts *bind.TransactOpts, retention *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "setLockedBondRetentionPeriod", retention)
}

// SetLockedBondRetentionPeriod is a paid mutator transaction binding the contract method 0x99965225.
//
// Solidity: function setLockedBondRetentionPeriod(uint256 retention) returns()
func (_CSAccounting *CSAccountingSession) SetLockedBondRetentionPeriod(retention *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.SetLockedBondRetentionPeriod(&_CSAccounting.TransactOpts, retention)
}

// SetLockedBondRetentionPeriod is a paid mutator transaction binding the contract method 0x99965225.
//
// Solidity: function setLockedBondRetentionPeriod(uint256 retention) returns()
func (_CSAccounting *CSAccountingTransactorSession) SetLockedBondRetentionPeriod(retention *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.SetLockedBondRetentionPeriod(&_CSAccounting.TransactOpts, retention)
}

// SettleLockedBondETH is a paid mutator transaction binding the contract method 0x4bb22a72.
//
// Solidity: function settleLockedBondETH(uint256 nodeOperatorId) returns(uint256 settledAmount)
func (_CSAccounting *CSAccountingTransactor) SettleLockedBondETH(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "settleLockedBondETH", nodeOperatorId)
}

// SettleLockedBondETH is a paid mutator transaction binding the contract method 0x4bb22a72.
//
// Solidity: function settleLockedBondETH(uint256 nodeOperatorId) returns(uint256 settledAmount)
func (_CSAccounting *CSAccountingSession) SettleLockedBondETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.SettleLockedBondETH(&_CSAccounting.TransactOpts, nodeOperatorId)
}

// SettleLockedBondETH is a paid mutator transaction binding the contract method 0x4bb22a72.
//
// Solidity: function settleLockedBondETH(uint256 nodeOperatorId) returns(uint256 settledAmount)
func (_CSAccounting *CSAccountingTransactorSession) SettleLockedBondETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.SettleLockedBondETH(&_CSAccounting.TransactOpts, nodeOperatorId)
}

// UpdateBondCurve is a paid mutator transaction binding the contract method 0x019c1a4f.
//
// Solidity: function updateBondCurve(uint256 curveId, uint256[] bondCurve) returns()
func (_CSAccounting *CSAccountingTransactor) UpdateBondCurve(opts *bind.TransactOpts, curveId *big.Int, bondCurve []*big.Int) (*types.Transaction, error) {
	return _CSAccounting.contract.Transact(opts, "updateBondCurve", curveId, bondCurve)
}

// UpdateBondCurve is a paid mutator transaction binding the contract method 0x019c1a4f.
//
// Solidity: function updateBondCurve(uint256 curveId, uint256[] bondCurve) returns()
func (_CSAccounting *CSAccountingSession) UpdateBondCurve(curveId *big.Int, bondCurve []*big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.UpdateBondCurve(&_CSAccounting.TransactOpts, curveId, bondCurve)
}

// UpdateBondCurve is a paid mutator transaction binding the contract method 0x019c1a4f.
//
// Solidity: function updateBondCurve(uint256 curveId, uint256[] bondCurve) returns()
func (_CSAccounting *CSAccountingTransactorSession) UpdateBondCurve(curveId *big.Int, bondCurve []*big.Int) (*types.Transaction, error) {
	return _CSAccounting.Contract.UpdateBondCurve(&_CSAccounting.TransactOpts, curveId, bondCurve)
}

// CSAccountingBondBurnedIterator is returned from FilterBondBurned and is used to iterate over the raw logs and unpacked data for BondBurned events raised by the CSAccounting contract.
type CSAccountingBondBurnedIterator struct {
	Event *CSAccountingBondBurned // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondBurned)
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
		it.Event = new(CSAccountingBondBurned)
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
func (it *CSAccountingBondBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondBurned represents a BondBurned event raised by the CSAccounting contract.
type CSAccountingBondBurned struct {
	NodeOperatorId *big.Int
	ToBurnAmount   *big.Int
	BurnedAmount   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondBurned is a free log retrieval operation binding the contract event 0x4da924ae7845fe96897faab524b536685b8bbc4d82fbb45c10d941e0f86ade0f.
//
// Solidity: event BondBurned(uint256 indexed nodeOperatorId, uint256 toBurnAmount, uint256 burnedAmount)
func (_CSAccounting *CSAccountingFilterer) FilterBondBurned(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSAccountingBondBurnedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondBurned", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondBurnedIterator{contract: _CSAccounting.contract, event: "BondBurned", logs: logs, sub: sub}, nil
}

// WatchBondBurned is a free log subscription operation binding the contract event 0x4da924ae7845fe96897faab524b536685b8bbc4d82fbb45c10d941e0f86ade0f.
//
// Solidity: event BondBurned(uint256 indexed nodeOperatorId, uint256 toBurnAmount, uint256 burnedAmount)
func (_CSAccounting *CSAccountingFilterer) WatchBondBurned(opts *bind.WatchOpts, sink chan<- *CSAccountingBondBurned, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondBurned", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondBurned)
				if err := _CSAccounting.contract.UnpackLog(event, "BondBurned", log); err != nil {
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

// ParseBondBurned is a log parse operation binding the contract event 0x4da924ae7845fe96897faab524b536685b8bbc4d82fbb45c10d941e0f86ade0f.
//
// Solidity: event BondBurned(uint256 indexed nodeOperatorId, uint256 toBurnAmount, uint256 burnedAmount)
func (_CSAccounting *CSAccountingFilterer) ParseBondBurned(log types.Log) (*CSAccountingBondBurned, error) {
	event := new(CSAccountingBondBurned)
	if err := _CSAccounting.contract.UnpackLog(event, "BondBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondChargedIterator is returned from FilterBondCharged and is used to iterate over the raw logs and unpacked data for BondCharged events raised by the CSAccounting contract.
type CSAccountingBondChargedIterator struct {
	Event *CSAccountingBondCharged // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondCharged)
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
		it.Event = new(CSAccountingBondCharged)
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
func (it *CSAccountingBondChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondCharged represents a BondCharged event raised by the CSAccounting contract.
type CSAccountingBondCharged struct {
	NodeOperatorId *big.Int
	ToChargeAmount *big.Int
	ChargedAmount  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondCharged is a free log retrieval operation binding the contract event 0x8615528474a7bb3a28d9971535d956b79242b8e8fcfb27f3e331270fff088afd.
//
// Solidity: event BondCharged(uint256 indexed nodeOperatorId, uint256 toChargeAmount, uint256 chargedAmount)
func (_CSAccounting *CSAccountingFilterer) FilterBondCharged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSAccountingBondChargedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondCharged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondChargedIterator{contract: _CSAccounting.contract, event: "BondCharged", logs: logs, sub: sub}, nil
}

// WatchBondCharged is a free log subscription operation binding the contract event 0x8615528474a7bb3a28d9971535d956b79242b8e8fcfb27f3e331270fff088afd.
//
// Solidity: event BondCharged(uint256 indexed nodeOperatorId, uint256 toChargeAmount, uint256 chargedAmount)
func (_CSAccounting *CSAccountingFilterer) WatchBondCharged(opts *bind.WatchOpts, sink chan<- *CSAccountingBondCharged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondCharged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondCharged)
				if err := _CSAccounting.contract.UnpackLog(event, "BondCharged", log); err != nil {
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

// ParseBondCharged is a log parse operation binding the contract event 0x8615528474a7bb3a28d9971535d956b79242b8e8fcfb27f3e331270fff088afd.
//
// Solidity: event BondCharged(uint256 indexed nodeOperatorId, uint256 toChargeAmount, uint256 chargedAmount)
func (_CSAccounting *CSAccountingFilterer) ParseBondCharged(log types.Log) (*CSAccountingBondCharged, error) {
	event := new(CSAccountingBondCharged)
	if err := _CSAccounting.contract.UnpackLog(event, "BondCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondClaimedStETHIterator is returned from FilterBondClaimedStETH and is used to iterate over the raw logs and unpacked data for BondClaimedStETH events raised by the CSAccounting contract.
type CSAccountingBondClaimedStETHIterator struct {
	Event *CSAccountingBondClaimedStETH // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondClaimedStETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondClaimedStETH)
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
		it.Event = new(CSAccountingBondClaimedStETH)
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
func (it *CSAccountingBondClaimedStETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondClaimedStETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondClaimedStETH represents a BondClaimedStETH event raised by the CSAccounting contract.
type CSAccountingBondClaimedStETH struct {
	NodeOperatorId *big.Int
	To             common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondClaimedStETH is a free log retrieval operation binding the contract event 0x3e3a1398fe71575ed0c17a80cd9d46ad684c2c75c2fad7b0e7dde15e78ab22d3.
//
// Solidity: event BondClaimedStETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) FilterBondClaimedStETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSAccountingBondClaimedStETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondClaimedStETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondClaimedStETHIterator{contract: _CSAccounting.contract, event: "BondClaimedStETH", logs: logs, sub: sub}, nil
}

// WatchBondClaimedStETH is a free log subscription operation binding the contract event 0x3e3a1398fe71575ed0c17a80cd9d46ad684c2c75c2fad7b0e7dde15e78ab22d3.
//
// Solidity: event BondClaimedStETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) WatchBondClaimedStETH(opts *bind.WatchOpts, sink chan<- *CSAccountingBondClaimedStETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondClaimedStETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondClaimedStETH)
				if err := _CSAccounting.contract.UnpackLog(event, "BondClaimedStETH", log); err != nil {
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

// ParseBondClaimedStETH is a log parse operation binding the contract event 0x3e3a1398fe71575ed0c17a80cd9d46ad684c2c75c2fad7b0e7dde15e78ab22d3.
//
// Solidity: event BondClaimedStETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) ParseBondClaimedStETH(log types.Log) (*CSAccountingBondClaimedStETH, error) {
	event := new(CSAccountingBondClaimedStETH)
	if err := _CSAccounting.contract.UnpackLog(event, "BondClaimedStETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondClaimedUnstETHIterator is returned from FilterBondClaimedUnstETH and is used to iterate over the raw logs and unpacked data for BondClaimedUnstETH events raised by the CSAccounting contract.
type CSAccountingBondClaimedUnstETHIterator struct {
	Event *CSAccountingBondClaimedUnstETH // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondClaimedUnstETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondClaimedUnstETH)
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
		it.Event = new(CSAccountingBondClaimedUnstETH)
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
func (it *CSAccountingBondClaimedUnstETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondClaimedUnstETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondClaimedUnstETH represents a BondClaimedUnstETH event raised by the CSAccounting contract.
type CSAccountingBondClaimedUnstETH struct {
	NodeOperatorId *big.Int
	To             common.Address
	Amount         *big.Int
	RequestId      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondClaimedUnstETH is a free log retrieval operation binding the contract event 0x26673a9d018b21192d08ee14377b798f11b9e5b15ea1559c110265716b8985b5.
//
// Solidity: event BondClaimedUnstETH(uint256 indexed nodeOperatorId, address to, uint256 amount, uint256 requestId)
func (_CSAccounting *CSAccountingFilterer) FilterBondClaimedUnstETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSAccountingBondClaimedUnstETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondClaimedUnstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondClaimedUnstETHIterator{contract: _CSAccounting.contract, event: "BondClaimedUnstETH", logs: logs, sub: sub}, nil
}

// WatchBondClaimedUnstETH is a free log subscription operation binding the contract event 0x26673a9d018b21192d08ee14377b798f11b9e5b15ea1559c110265716b8985b5.
//
// Solidity: event BondClaimedUnstETH(uint256 indexed nodeOperatorId, address to, uint256 amount, uint256 requestId)
func (_CSAccounting *CSAccountingFilterer) WatchBondClaimedUnstETH(opts *bind.WatchOpts, sink chan<- *CSAccountingBondClaimedUnstETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondClaimedUnstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondClaimedUnstETH)
				if err := _CSAccounting.contract.UnpackLog(event, "BondClaimedUnstETH", log); err != nil {
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

// ParseBondClaimedUnstETH is a log parse operation binding the contract event 0x26673a9d018b21192d08ee14377b798f11b9e5b15ea1559c110265716b8985b5.
//
// Solidity: event BondClaimedUnstETH(uint256 indexed nodeOperatorId, address to, uint256 amount, uint256 requestId)
func (_CSAccounting *CSAccountingFilterer) ParseBondClaimedUnstETH(log types.Log) (*CSAccountingBondClaimedUnstETH, error) {
	event := new(CSAccountingBondClaimedUnstETH)
	if err := _CSAccounting.contract.UnpackLog(event, "BondClaimedUnstETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondClaimedWstETHIterator is returned from FilterBondClaimedWstETH and is used to iterate over the raw logs and unpacked data for BondClaimedWstETH events raised by the CSAccounting contract.
type CSAccountingBondClaimedWstETHIterator struct {
	Event *CSAccountingBondClaimedWstETH // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondClaimedWstETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondClaimedWstETH)
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
		it.Event = new(CSAccountingBondClaimedWstETH)
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
func (it *CSAccountingBondClaimedWstETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondClaimedWstETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondClaimedWstETH represents a BondClaimedWstETH event raised by the CSAccounting contract.
type CSAccountingBondClaimedWstETH struct {
	NodeOperatorId *big.Int
	To             common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondClaimedWstETH is a free log retrieval operation binding the contract event 0xe6a8c06447e05a412e5e9581e088941f3994db3d8a9bfd3275b38d77acacafac.
//
// Solidity: event BondClaimedWstETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) FilterBondClaimedWstETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSAccountingBondClaimedWstETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondClaimedWstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondClaimedWstETHIterator{contract: _CSAccounting.contract, event: "BondClaimedWstETH", logs: logs, sub: sub}, nil
}

// WatchBondClaimedWstETH is a free log subscription operation binding the contract event 0xe6a8c06447e05a412e5e9581e088941f3994db3d8a9bfd3275b38d77acacafac.
//
// Solidity: event BondClaimedWstETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) WatchBondClaimedWstETH(opts *bind.WatchOpts, sink chan<- *CSAccountingBondClaimedWstETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondClaimedWstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondClaimedWstETH)
				if err := _CSAccounting.contract.UnpackLog(event, "BondClaimedWstETH", log); err != nil {
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

// ParseBondClaimedWstETH is a log parse operation binding the contract event 0xe6a8c06447e05a412e5e9581e088941f3994db3d8a9bfd3275b38d77acacafac.
//
// Solidity: event BondClaimedWstETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) ParseBondClaimedWstETH(log types.Log) (*CSAccountingBondClaimedWstETH, error) {
	event := new(CSAccountingBondClaimedWstETH)
	if err := _CSAccounting.contract.UnpackLog(event, "BondClaimedWstETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondCurveAddedIterator is returned from FilterBondCurveAdded and is used to iterate over the raw logs and unpacked data for BondCurveAdded events raised by the CSAccounting contract.
type CSAccountingBondCurveAddedIterator struct {
	Event *CSAccountingBondCurveAdded // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondCurveAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondCurveAdded)
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
		it.Event = new(CSAccountingBondCurveAdded)
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
func (it *CSAccountingBondCurveAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondCurveAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondCurveAdded represents a BondCurveAdded event raised by the CSAccounting contract.
type CSAccountingBondCurveAdded struct {
	BondCurve []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBondCurveAdded is a free log retrieval operation binding the contract event 0x1fb1d9b944dd7015e95b7b7a9623c45792e4532badcf9c6e7a284d7d4d0570f0.
//
// Solidity: event BondCurveAdded(uint256[] bondCurve)
func (_CSAccounting *CSAccountingFilterer) FilterBondCurveAdded(opts *bind.FilterOpts) (*CSAccountingBondCurveAddedIterator, error) {

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondCurveAdded")
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondCurveAddedIterator{contract: _CSAccounting.contract, event: "BondCurveAdded", logs: logs, sub: sub}, nil
}

// WatchBondCurveAdded is a free log subscription operation binding the contract event 0x1fb1d9b944dd7015e95b7b7a9623c45792e4532badcf9c6e7a284d7d4d0570f0.
//
// Solidity: event BondCurveAdded(uint256[] bondCurve)
func (_CSAccounting *CSAccountingFilterer) WatchBondCurveAdded(opts *bind.WatchOpts, sink chan<- *CSAccountingBondCurveAdded) (event.Subscription, error) {

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondCurveAdded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondCurveAdded)
				if err := _CSAccounting.contract.UnpackLog(event, "BondCurveAdded", log); err != nil {
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

// ParseBondCurveAdded is a log parse operation binding the contract event 0x1fb1d9b944dd7015e95b7b7a9623c45792e4532badcf9c6e7a284d7d4d0570f0.
//
// Solidity: event BondCurveAdded(uint256[] bondCurve)
func (_CSAccounting *CSAccountingFilterer) ParseBondCurveAdded(log types.Log) (*CSAccountingBondCurveAdded, error) {
	event := new(CSAccountingBondCurveAdded)
	if err := _CSAccounting.contract.UnpackLog(event, "BondCurveAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondCurveSetIterator is returned from FilterBondCurveSet and is used to iterate over the raw logs and unpacked data for BondCurveSet events raised by the CSAccounting contract.
type CSAccountingBondCurveSetIterator struct {
	Event *CSAccountingBondCurveSet // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondCurveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondCurveSet)
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
		it.Event = new(CSAccountingBondCurveSet)
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
func (it *CSAccountingBondCurveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondCurveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondCurveSet represents a BondCurveSet event raised by the CSAccounting contract.
type CSAccountingBondCurveSet struct {
	NodeOperatorId *big.Int
	CurveId        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondCurveSet is a free log retrieval operation binding the contract event 0x4642db1736894887bc907d721f20af84d3e585a0a3cea90f41b78b2aa906541b.
//
// Solidity: event BondCurveSet(uint256 indexed nodeOperatorId, uint256 curveId)
func (_CSAccounting *CSAccountingFilterer) FilterBondCurveSet(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSAccountingBondCurveSetIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondCurveSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondCurveSetIterator{contract: _CSAccounting.contract, event: "BondCurveSet", logs: logs, sub: sub}, nil
}

// WatchBondCurveSet is a free log subscription operation binding the contract event 0x4642db1736894887bc907d721f20af84d3e585a0a3cea90f41b78b2aa906541b.
//
// Solidity: event BondCurveSet(uint256 indexed nodeOperatorId, uint256 curveId)
func (_CSAccounting *CSAccountingFilterer) WatchBondCurveSet(opts *bind.WatchOpts, sink chan<- *CSAccountingBondCurveSet, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondCurveSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondCurveSet)
				if err := _CSAccounting.contract.UnpackLog(event, "BondCurveSet", log); err != nil {
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

// ParseBondCurveSet is a log parse operation binding the contract event 0x4642db1736894887bc907d721f20af84d3e585a0a3cea90f41b78b2aa906541b.
//
// Solidity: event BondCurveSet(uint256 indexed nodeOperatorId, uint256 curveId)
func (_CSAccounting *CSAccountingFilterer) ParseBondCurveSet(log types.Log) (*CSAccountingBondCurveSet, error) {
	event := new(CSAccountingBondCurveSet)
	if err := _CSAccounting.contract.UnpackLog(event, "BondCurveSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondCurveUpdatedIterator is returned from FilterBondCurveUpdated and is used to iterate over the raw logs and unpacked data for BondCurveUpdated events raised by the CSAccounting contract.
type CSAccountingBondCurveUpdatedIterator struct {
	Event *CSAccountingBondCurveUpdated // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondCurveUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondCurveUpdated)
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
		it.Event = new(CSAccountingBondCurveUpdated)
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
func (it *CSAccountingBondCurveUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondCurveUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondCurveUpdated represents a BondCurveUpdated event raised by the CSAccounting contract.
type CSAccountingBondCurveUpdated struct {
	CurveId   *big.Int
	BondCurve []*big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBondCurveUpdated is a free log retrieval operation binding the contract event 0x53da7af401538204fd91f2946f2fe85d05224d2cc766fd7aa9fbd8bf4fb4ce9f.
//
// Solidity: event BondCurveUpdated(uint256 indexed curveId, uint256[] bondCurve)
func (_CSAccounting *CSAccountingFilterer) FilterBondCurveUpdated(opts *bind.FilterOpts, curveId []*big.Int) (*CSAccountingBondCurveUpdatedIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondCurveUpdated", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondCurveUpdatedIterator{contract: _CSAccounting.contract, event: "BondCurveUpdated", logs: logs, sub: sub}, nil
}

// WatchBondCurveUpdated is a free log subscription operation binding the contract event 0x53da7af401538204fd91f2946f2fe85d05224d2cc766fd7aa9fbd8bf4fb4ce9f.
//
// Solidity: event BondCurveUpdated(uint256 indexed curveId, uint256[] bondCurve)
func (_CSAccounting *CSAccountingFilterer) WatchBondCurveUpdated(opts *bind.WatchOpts, sink chan<- *CSAccountingBondCurveUpdated, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondCurveUpdated", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondCurveUpdated)
				if err := _CSAccounting.contract.UnpackLog(event, "BondCurveUpdated", log); err != nil {
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

// ParseBondCurveUpdated is a log parse operation binding the contract event 0x53da7af401538204fd91f2946f2fe85d05224d2cc766fd7aa9fbd8bf4fb4ce9f.
//
// Solidity: event BondCurveUpdated(uint256 indexed curveId, uint256[] bondCurve)
func (_CSAccounting *CSAccountingFilterer) ParseBondCurveUpdated(log types.Log) (*CSAccountingBondCurveUpdated, error) {
	event := new(CSAccountingBondCurveUpdated)
	if err := _CSAccounting.contract.UnpackLog(event, "BondCurveUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondDepositedETHIterator is returned from FilterBondDepositedETH and is used to iterate over the raw logs and unpacked data for BondDepositedETH events raised by the CSAccounting contract.
type CSAccountingBondDepositedETHIterator struct {
	Event *CSAccountingBondDepositedETH // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondDepositedETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondDepositedETH)
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
		it.Event = new(CSAccountingBondDepositedETH)
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
func (it *CSAccountingBondDepositedETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondDepositedETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondDepositedETH represents a BondDepositedETH event raised by the CSAccounting contract.
type CSAccountingBondDepositedETH struct {
	NodeOperatorId *big.Int
	From           common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondDepositedETH is a free log retrieval operation binding the contract event 0x16ec5116295424dec7fd52c87d9971a963ea7f59f741ad9ad468f0312055dc49.
//
// Solidity: event BondDepositedETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) FilterBondDepositedETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSAccountingBondDepositedETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondDepositedETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondDepositedETHIterator{contract: _CSAccounting.contract, event: "BondDepositedETH", logs: logs, sub: sub}, nil
}

// WatchBondDepositedETH is a free log subscription operation binding the contract event 0x16ec5116295424dec7fd52c87d9971a963ea7f59f741ad9ad468f0312055dc49.
//
// Solidity: event BondDepositedETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) WatchBondDepositedETH(opts *bind.WatchOpts, sink chan<- *CSAccountingBondDepositedETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondDepositedETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondDepositedETH)
				if err := _CSAccounting.contract.UnpackLog(event, "BondDepositedETH", log); err != nil {
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

// ParseBondDepositedETH is a log parse operation binding the contract event 0x16ec5116295424dec7fd52c87d9971a963ea7f59f741ad9ad468f0312055dc49.
//
// Solidity: event BondDepositedETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) ParseBondDepositedETH(log types.Log) (*CSAccountingBondDepositedETH, error) {
	event := new(CSAccountingBondDepositedETH)
	if err := _CSAccounting.contract.UnpackLog(event, "BondDepositedETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondDepositedStETHIterator is returned from FilterBondDepositedStETH and is used to iterate over the raw logs and unpacked data for BondDepositedStETH events raised by the CSAccounting contract.
type CSAccountingBondDepositedStETHIterator struct {
	Event *CSAccountingBondDepositedStETH // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondDepositedStETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondDepositedStETH)
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
		it.Event = new(CSAccountingBondDepositedStETH)
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
func (it *CSAccountingBondDepositedStETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondDepositedStETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondDepositedStETH represents a BondDepositedStETH event raised by the CSAccounting contract.
type CSAccountingBondDepositedStETH struct {
	NodeOperatorId *big.Int
	From           common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondDepositedStETH is a free log retrieval operation binding the contract event 0xee31ebba29fd5471227e42fd8ca621a892d689901892cb8febb03fe802c3214b.
//
// Solidity: event BondDepositedStETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) FilterBondDepositedStETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSAccountingBondDepositedStETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondDepositedStETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondDepositedStETHIterator{contract: _CSAccounting.contract, event: "BondDepositedStETH", logs: logs, sub: sub}, nil
}

// WatchBondDepositedStETH is a free log subscription operation binding the contract event 0xee31ebba29fd5471227e42fd8ca621a892d689901892cb8febb03fe802c3214b.
//
// Solidity: event BondDepositedStETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) WatchBondDepositedStETH(opts *bind.WatchOpts, sink chan<- *CSAccountingBondDepositedStETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondDepositedStETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondDepositedStETH)
				if err := _CSAccounting.contract.UnpackLog(event, "BondDepositedStETH", log); err != nil {
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

// ParseBondDepositedStETH is a log parse operation binding the contract event 0xee31ebba29fd5471227e42fd8ca621a892d689901892cb8febb03fe802c3214b.
//
// Solidity: event BondDepositedStETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) ParseBondDepositedStETH(log types.Log) (*CSAccountingBondDepositedStETH, error) {
	event := new(CSAccountingBondDepositedStETH)
	if err := _CSAccounting.contract.UnpackLog(event, "BondDepositedStETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondDepositedWstETHIterator is returned from FilterBondDepositedWstETH and is used to iterate over the raw logs and unpacked data for BondDepositedWstETH events raised by the CSAccounting contract.
type CSAccountingBondDepositedWstETHIterator struct {
	Event *CSAccountingBondDepositedWstETH // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondDepositedWstETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondDepositedWstETH)
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
		it.Event = new(CSAccountingBondDepositedWstETH)
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
func (it *CSAccountingBondDepositedWstETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondDepositedWstETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondDepositedWstETH represents a BondDepositedWstETH event raised by the CSAccounting contract.
type CSAccountingBondDepositedWstETH struct {
	NodeOperatorId *big.Int
	From           common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondDepositedWstETH is a free log retrieval operation binding the contract event 0x6576bbc9c5b478bf9717dc3d2bcb485e5ff0727df77c72558727597f3609d3f1.
//
// Solidity: event BondDepositedWstETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) FilterBondDepositedWstETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSAccountingBondDepositedWstETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondDepositedWstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondDepositedWstETHIterator{contract: _CSAccounting.contract, event: "BondDepositedWstETH", logs: logs, sub: sub}, nil
}

// WatchBondDepositedWstETH is a free log subscription operation binding the contract event 0x6576bbc9c5b478bf9717dc3d2bcb485e5ff0727df77c72558727597f3609d3f1.
//
// Solidity: event BondDepositedWstETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) WatchBondDepositedWstETH(opts *bind.WatchOpts, sink chan<- *CSAccountingBondDepositedWstETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondDepositedWstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondDepositedWstETH)
				if err := _CSAccounting.contract.UnpackLog(event, "BondDepositedWstETH", log); err != nil {
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

// ParseBondDepositedWstETH is a log parse operation binding the contract event 0x6576bbc9c5b478bf9717dc3d2bcb485e5ff0727df77c72558727597f3609d3f1.
//
// Solidity: event BondDepositedWstETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) ParseBondDepositedWstETH(log types.Log) (*CSAccountingBondDepositedWstETH, error) {
	event := new(CSAccountingBondDepositedWstETH)
	if err := _CSAccounting.contract.UnpackLog(event, "BondDepositedWstETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondLockChangedIterator is returned from FilterBondLockChanged and is used to iterate over the raw logs and unpacked data for BondLockChanged events raised by the CSAccounting contract.
type CSAccountingBondLockChangedIterator struct {
	Event *CSAccountingBondLockChanged // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondLockChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondLockChanged)
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
		it.Event = new(CSAccountingBondLockChanged)
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
func (it *CSAccountingBondLockChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondLockChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondLockChanged represents a BondLockChanged event raised by the CSAccounting contract.
type CSAccountingBondLockChanged struct {
	NodeOperatorId *big.Int
	NewAmount      *big.Int
	RetentionUntil *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondLockChanged is a free log retrieval operation binding the contract event 0x69a153d448f54b17f05cf3b268a2efab87c94a4727d108c4ca4aa3e5d65113de.
//
// Solidity: event BondLockChanged(uint256 indexed nodeOperatorId, uint256 newAmount, uint256 retentionUntil)
func (_CSAccounting *CSAccountingFilterer) FilterBondLockChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSAccountingBondLockChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondLockChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondLockChangedIterator{contract: _CSAccounting.contract, event: "BondLockChanged", logs: logs, sub: sub}, nil
}

// WatchBondLockChanged is a free log subscription operation binding the contract event 0x69a153d448f54b17f05cf3b268a2efab87c94a4727d108c4ca4aa3e5d65113de.
//
// Solidity: event BondLockChanged(uint256 indexed nodeOperatorId, uint256 newAmount, uint256 retentionUntil)
func (_CSAccounting *CSAccountingFilterer) WatchBondLockChanged(opts *bind.WatchOpts, sink chan<- *CSAccountingBondLockChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondLockChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondLockChanged)
				if err := _CSAccounting.contract.UnpackLog(event, "BondLockChanged", log); err != nil {
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

// ParseBondLockChanged is a log parse operation binding the contract event 0x69a153d448f54b17f05cf3b268a2efab87c94a4727d108c4ca4aa3e5d65113de.
//
// Solidity: event BondLockChanged(uint256 indexed nodeOperatorId, uint256 newAmount, uint256 retentionUntil)
func (_CSAccounting *CSAccountingFilterer) ParseBondLockChanged(log types.Log) (*CSAccountingBondLockChanged, error) {
	event := new(CSAccountingBondLockChanged)
	if err := _CSAccounting.contract.UnpackLog(event, "BondLockChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondLockCompensatedIterator is returned from FilterBondLockCompensated and is used to iterate over the raw logs and unpacked data for BondLockCompensated events raised by the CSAccounting contract.
type CSAccountingBondLockCompensatedIterator struct {
	Event *CSAccountingBondLockCompensated // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondLockCompensatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondLockCompensated)
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
		it.Event = new(CSAccountingBondLockCompensated)
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
func (it *CSAccountingBondLockCompensatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondLockCompensatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondLockCompensated represents a BondLockCompensated event raised by the CSAccounting contract.
type CSAccountingBondLockCompensated struct {
	NodeOperatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondLockCompensated is a free log retrieval operation binding the contract event 0xb6ee6e3aae6776519627b46786a622b642c38cabfe4c97cb34054fd63fc11a23.
//
// Solidity: event BondLockCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) FilterBondLockCompensated(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSAccountingBondLockCompensatedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondLockCompensated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondLockCompensatedIterator{contract: _CSAccounting.contract, event: "BondLockCompensated", logs: logs, sub: sub}, nil
}

// WatchBondLockCompensated is a free log subscription operation binding the contract event 0xb6ee6e3aae6776519627b46786a622b642c38cabfe4c97cb34054fd63fc11a23.
//
// Solidity: event BondLockCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) WatchBondLockCompensated(opts *bind.WatchOpts, sink chan<- *CSAccountingBondLockCompensated, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondLockCompensated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondLockCompensated)
				if err := _CSAccounting.contract.UnpackLog(event, "BondLockCompensated", log); err != nil {
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

// ParseBondLockCompensated is a log parse operation binding the contract event 0xb6ee6e3aae6776519627b46786a622b642c38cabfe4c97cb34054fd63fc11a23.
//
// Solidity: event BondLockCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) ParseBondLockCompensated(log types.Log) (*CSAccountingBondLockCompensated, error) {
	event := new(CSAccountingBondLockCompensated)
	if err := _CSAccounting.contract.UnpackLog(event, "BondLockCompensated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondLockRemovedIterator is returned from FilterBondLockRemoved and is used to iterate over the raw logs and unpacked data for BondLockRemoved events raised by the CSAccounting contract.
type CSAccountingBondLockRemovedIterator struct {
	Event *CSAccountingBondLockRemoved // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondLockRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondLockRemoved)
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
		it.Event = new(CSAccountingBondLockRemoved)
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
func (it *CSAccountingBondLockRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondLockRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondLockRemoved represents a BondLockRemoved event raised by the CSAccounting contract.
type CSAccountingBondLockRemoved struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondLockRemoved is a free log retrieval operation binding the contract event 0x844ae6b00e8a437dcdde1a634feab3273e08bb5c274a4be3b195b308ae0ba20a.
//
// Solidity: event BondLockRemoved(uint256 indexed nodeOperatorId)
func (_CSAccounting *CSAccountingFilterer) FilterBondLockRemoved(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CSAccountingBondLockRemovedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondLockRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondLockRemovedIterator{contract: _CSAccounting.contract, event: "BondLockRemoved", logs: logs, sub: sub}, nil
}

// WatchBondLockRemoved is a free log subscription operation binding the contract event 0x844ae6b00e8a437dcdde1a634feab3273e08bb5c274a4be3b195b308ae0ba20a.
//
// Solidity: event BondLockRemoved(uint256 indexed nodeOperatorId)
func (_CSAccounting *CSAccountingFilterer) WatchBondLockRemoved(opts *bind.WatchOpts, sink chan<- *CSAccountingBondLockRemoved, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondLockRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondLockRemoved)
				if err := _CSAccounting.contract.UnpackLog(event, "BondLockRemoved", log); err != nil {
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

// ParseBondLockRemoved is a log parse operation binding the contract event 0x844ae6b00e8a437dcdde1a634feab3273e08bb5c274a4be3b195b308ae0ba20a.
//
// Solidity: event BondLockRemoved(uint256 indexed nodeOperatorId)
func (_CSAccounting *CSAccountingFilterer) ParseBondLockRemoved(log types.Log) (*CSAccountingBondLockRemoved, error) {
	event := new(CSAccountingBondLockRemoved)
	if err := _CSAccounting.contract.UnpackLog(event, "BondLockRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingBondLockRetentionPeriodChangedIterator is returned from FilterBondLockRetentionPeriodChanged and is used to iterate over the raw logs and unpacked data for BondLockRetentionPeriodChanged events raised by the CSAccounting contract.
type CSAccountingBondLockRetentionPeriodChangedIterator struct {
	Event *CSAccountingBondLockRetentionPeriodChanged // Event containing the contract specifics and raw log

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
func (it *CSAccountingBondLockRetentionPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingBondLockRetentionPeriodChanged)
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
		it.Event = new(CSAccountingBondLockRetentionPeriodChanged)
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
func (it *CSAccountingBondLockRetentionPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingBondLockRetentionPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingBondLockRetentionPeriodChanged represents a BondLockRetentionPeriodChanged event raised by the CSAccounting contract.
type CSAccountingBondLockRetentionPeriodChanged struct {
	RetentionPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBondLockRetentionPeriodChanged is a free log retrieval operation binding the contract event 0xdaf5eddbe9ed0768e54cc8f739a9cb86c57fc70da07eff01d9ba886f21a7a4b3.
//
// Solidity: event BondLockRetentionPeriodChanged(uint256 retentionPeriod)
func (_CSAccounting *CSAccountingFilterer) FilterBondLockRetentionPeriodChanged(opts *bind.FilterOpts) (*CSAccountingBondLockRetentionPeriodChangedIterator, error) {

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "BondLockRetentionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &CSAccountingBondLockRetentionPeriodChangedIterator{contract: _CSAccounting.contract, event: "BondLockRetentionPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchBondLockRetentionPeriodChanged is a free log subscription operation binding the contract event 0xdaf5eddbe9ed0768e54cc8f739a9cb86c57fc70da07eff01d9ba886f21a7a4b3.
//
// Solidity: event BondLockRetentionPeriodChanged(uint256 retentionPeriod)
func (_CSAccounting *CSAccountingFilterer) WatchBondLockRetentionPeriodChanged(opts *bind.WatchOpts, sink chan<- *CSAccountingBondLockRetentionPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "BondLockRetentionPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingBondLockRetentionPeriodChanged)
				if err := _CSAccounting.contract.UnpackLog(event, "BondLockRetentionPeriodChanged", log); err != nil {
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

// ParseBondLockRetentionPeriodChanged is a log parse operation binding the contract event 0xdaf5eddbe9ed0768e54cc8f739a9cb86c57fc70da07eff01d9ba886f21a7a4b3.
//
// Solidity: event BondLockRetentionPeriodChanged(uint256 retentionPeriod)
func (_CSAccounting *CSAccountingFilterer) ParseBondLockRetentionPeriodChanged(log types.Log) (*CSAccountingBondLockRetentionPeriodChanged, error) {
	event := new(CSAccountingBondLockRetentionPeriodChanged)
	if err := _CSAccounting.contract.UnpackLog(event, "BondLockRetentionPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingChargePenaltyRecipientSetIterator is returned from FilterChargePenaltyRecipientSet and is used to iterate over the raw logs and unpacked data for ChargePenaltyRecipientSet events raised by the CSAccounting contract.
type CSAccountingChargePenaltyRecipientSetIterator struct {
	Event *CSAccountingChargePenaltyRecipientSet // Event containing the contract specifics and raw log

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
func (it *CSAccountingChargePenaltyRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingChargePenaltyRecipientSet)
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
		it.Event = new(CSAccountingChargePenaltyRecipientSet)
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
func (it *CSAccountingChargePenaltyRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingChargePenaltyRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingChargePenaltyRecipientSet represents a ChargePenaltyRecipientSet event raised by the CSAccounting contract.
type CSAccountingChargePenaltyRecipientSet struct {
	ChargePenaltyRecipient common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterChargePenaltyRecipientSet is a free log retrieval operation binding the contract event 0x4beaaee83871b066b675515d6a53567e76411f60266703cef934a01905a4d832.
//
// Solidity: event ChargePenaltyRecipientSet(address chargePenaltyRecipient)
func (_CSAccounting *CSAccountingFilterer) FilterChargePenaltyRecipientSet(opts *bind.FilterOpts) (*CSAccountingChargePenaltyRecipientSetIterator, error) {

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "ChargePenaltyRecipientSet")
	if err != nil {
		return nil, err
	}
	return &CSAccountingChargePenaltyRecipientSetIterator{contract: _CSAccounting.contract, event: "ChargePenaltyRecipientSet", logs: logs, sub: sub}, nil
}

// WatchChargePenaltyRecipientSet is a free log subscription operation binding the contract event 0x4beaaee83871b066b675515d6a53567e76411f60266703cef934a01905a4d832.
//
// Solidity: event ChargePenaltyRecipientSet(address chargePenaltyRecipient)
func (_CSAccounting *CSAccountingFilterer) WatchChargePenaltyRecipientSet(opts *bind.WatchOpts, sink chan<- *CSAccountingChargePenaltyRecipientSet) (event.Subscription, error) {

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "ChargePenaltyRecipientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingChargePenaltyRecipientSet)
				if err := _CSAccounting.contract.UnpackLog(event, "ChargePenaltyRecipientSet", log); err != nil {
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

// ParseChargePenaltyRecipientSet is a log parse operation binding the contract event 0x4beaaee83871b066b675515d6a53567e76411f60266703cef934a01905a4d832.
//
// Solidity: event ChargePenaltyRecipientSet(address chargePenaltyRecipient)
func (_CSAccounting *CSAccountingFilterer) ParseChargePenaltyRecipientSet(log types.Log) (*CSAccountingChargePenaltyRecipientSet, error) {
	event := new(CSAccountingChargePenaltyRecipientSet)
	if err := _CSAccounting.contract.UnpackLog(event, "ChargePenaltyRecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingERC1155RecoveredIterator is returned from FilterERC1155Recovered and is used to iterate over the raw logs and unpacked data for ERC1155Recovered events raised by the CSAccounting contract.
type CSAccountingERC1155RecoveredIterator struct {
	Event *CSAccountingERC1155Recovered // Event containing the contract specifics and raw log

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
func (it *CSAccountingERC1155RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingERC1155Recovered)
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
		it.Event = new(CSAccountingERC1155Recovered)
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
func (it *CSAccountingERC1155RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingERC1155RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingERC1155Recovered represents a ERC1155Recovered event raised by the CSAccounting contract.
type CSAccountingERC1155Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC1155Recovered is a free log retrieval operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) FilterERC1155Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CSAccountingERC1155RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingERC1155RecoveredIterator{contract: _CSAccounting.contract, event: "ERC1155Recovered", logs: logs, sub: sub}, nil
}

// WatchERC1155Recovered is a free log subscription operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) WatchERC1155Recovered(opts *bind.WatchOpts, sink chan<- *CSAccountingERC1155Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingERC1155Recovered)
				if err := _CSAccounting.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
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
func (_CSAccounting *CSAccountingFilterer) ParseERC1155Recovered(log types.Log) (*CSAccountingERC1155Recovered, error) {
	event := new(CSAccountingERC1155Recovered)
	if err := _CSAccounting.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingERC20RecoveredIterator is returned from FilterERC20Recovered and is used to iterate over the raw logs and unpacked data for ERC20Recovered events raised by the CSAccounting contract.
type CSAccountingERC20RecoveredIterator struct {
	Event *CSAccountingERC20Recovered // Event containing the contract specifics and raw log

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
func (it *CSAccountingERC20RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingERC20Recovered)
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
		it.Event = new(CSAccountingERC20Recovered)
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
func (it *CSAccountingERC20RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingERC20RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingERC20Recovered represents a ERC20Recovered event raised by the CSAccounting contract.
type CSAccountingERC20Recovered struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Recovered is a free log retrieval operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) FilterERC20Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CSAccountingERC20RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingERC20RecoveredIterator{contract: _CSAccounting.contract, event: "ERC20Recovered", logs: logs, sub: sub}, nil
}

// WatchERC20Recovered is a free log subscription operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) WatchERC20Recovered(opts *bind.WatchOpts, sink chan<- *CSAccountingERC20Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingERC20Recovered)
				if err := _CSAccounting.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
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
func (_CSAccounting *CSAccountingFilterer) ParseERC20Recovered(log types.Log) (*CSAccountingERC20Recovered, error) {
	event := new(CSAccountingERC20Recovered)
	if err := _CSAccounting.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingERC721RecoveredIterator is returned from FilterERC721Recovered and is used to iterate over the raw logs and unpacked data for ERC721Recovered events raised by the CSAccounting contract.
type CSAccountingERC721RecoveredIterator struct {
	Event *CSAccountingERC721Recovered // Event containing the contract specifics and raw log

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
func (it *CSAccountingERC721RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingERC721Recovered)
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
		it.Event = new(CSAccountingERC721Recovered)
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
func (it *CSAccountingERC721RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingERC721RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingERC721Recovered represents a ERC721Recovered event raised by the CSAccounting contract.
type CSAccountingERC721Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721Recovered is a free log retrieval operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_CSAccounting *CSAccountingFilterer) FilterERC721Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CSAccountingERC721RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingERC721RecoveredIterator{contract: _CSAccounting.contract, event: "ERC721Recovered", logs: logs, sub: sub}, nil
}

// WatchERC721Recovered is a free log subscription operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_CSAccounting *CSAccountingFilterer) WatchERC721Recovered(opts *bind.WatchOpts, sink chan<- *CSAccountingERC721Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingERC721Recovered)
				if err := _CSAccounting.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
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
func (_CSAccounting *CSAccountingFilterer) ParseERC721Recovered(log types.Log) (*CSAccountingERC721Recovered, error) {
	event := new(CSAccountingERC721Recovered)
	if err := _CSAccounting.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingEtherRecoveredIterator is returned from FilterEtherRecovered and is used to iterate over the raw logs and unpacked data for EtherRecovered events raised by the CSAccounting contract.
type CSAccountingEtherRecoveredIterator struct {
	Event *CSAccountingEtherRecovered // Event containing the contract specifics and raw log

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
func (it *CSAccountingEtherRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingEtherRecovered)
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
		it.Event = new(CSAccountingEtherRecovered)
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
func (it *CSAccountingEtherRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingEtherRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingEtherRecovered represents a EtherRecovered event raised by the CSAccounting contract.
type CSAccountingEtherRecovered struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEtherRecovered is a free log retrieval operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) FilterEtherRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CSAccountingEtherRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingEtherRecoveredIterator{contract: _CSAccounting.contract, event: "EtherRecovered", logs: logs, sub: sub}, nil
}

// WatchEtherRecovered is a free log subscription operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_CSAccounting *CSAccountingFilterer) WatchEtherRecovered(opts *bind.WatchOpts, sink chan<- *CSAccountingEtherRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingEtherRecovered)
				if err := _CSAccounting.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
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
func (_CSAccounting *CSAccountingFilterer) ParseEtherRecovered(log types.Log) (*CSAccountingEtherRecovered, error) {
	event := new(CSAccountingEtherRecovered)
	if err := _CSAccounting.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the CSAccounting contract.
type CSAccountingInitializedIterator struct {
	Event *CSAccountingInitialized // Event containing the contract specifics and raw log

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
func (it *CSAccountingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingInitialized)
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
		it.Event = new(CSAccountingInitialized)
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
func (it *CSAccountingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingInitialized represents a Initialized event raised by the CSAccounting contract.
type CSAccountingInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CSAccounting *CSAccountingFilterer) FilterInitialized(opts *bind.FilterOpts) (*CSAccountingInitializedIterator, error) {

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CSAccountingInitializedIterator{contract: _CSAccounting.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_CSAccounting *CSAccountingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CSAccountingInitialized) (event.Subscription, error) {

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingInitialized)
				if err := _CSAccounting.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_CSAccounting *CSAccountingFilterer) ParseInitialized(log types.Log) (*CSAccountingInitialized, error) {
	event := new(CSAccountingInitialized)
	if err := _CSAccounting.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the CSAccounting contract.
type CSAccountingPausedIterator struct {
	Event *CSAccountingPaused // Event containing the contract specifics and raw log

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
func (it *CSAccountingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingPaused)
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
		it.Event = new(CSAccountingPaused)
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
func (it *CSAccountingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingPaused represents a Paused event raised by the CSAccounting contract.
type CSAccountingPaused struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_CSAccounting *CSAccountingFilterer) FilterPaused(opts *bind.FilterOpts) (*CSAccountingPausedIterator, error) {

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CSAccountingPausedIterator{contract: _CSAccounting.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_CSAccounting *CSAccountingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CSAccountingPaused) (event.Subscription, error) {

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingPaused)
				if err := _CSAccounting.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_CSAccounting *CSAccountingFilterer) ParsePaused(log types.Log) (*CSAccountingPaused, error) {
	event := new(CSAccountingPaused)
	if err := _CSAccounting.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the CSAccounting contract.
type CSAccountingResumedIterator struct {
	Event *CSAccountingResumed // Event containing the contract specifics and raw log

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
func (it *CSAccountingResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingResumed)
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
		it.Event = new(CSAccountingResumed)
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
func (it *CSAccountingResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingResumed represents a Resumed event raised by the CSAccounting contract.
type CSAccountingResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_CSAccounting *CSAccountingFilterer) FilterResumed(opts *bind.FilterOpts) (*CSAccountingResumedIterator, error) {

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &CSAccountingResumedIterator{contract: _CSAccounting.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_CSAccounting *CSAccountingFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *CSAccountingResumed) (event.Subscription, error) {

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingResumed)
				if err := _CSAccounting.contract.UnpackLog(event, "Resumed", log); err != nil {
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
func (_CSAccounting *CSAccountingFilterer) ParseResumed(log types.Log) (*CSAccountingResumed, error) {
	event := new(CSAccountingResumed)
	if err := _CSAccounting.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the CSAccounting contract.
type CSAccountingRoleAdminChangedIterator struct {
	Event *CSAccountingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CSAccountingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingRoleAdminChanged)
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
		it.Event = new(CSAccountingRoleAdminChanged)
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
func (it *CSAccountingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingRoleAdminChanged represents a RoleAdminChanged event raised by the CSAccounting contract.
type CSAccountingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CSAccounting *CSAccountingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CSAccountingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingRoleAdminChangedIterator{contract: _CSAccounting.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_CSAccounting *CSAccountingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CSAccountingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingRoleAdminChanged)
				if err := _CSAccounting.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_CSAccounting *CSAccountingFilterer) ParseRoleAdminChanged(log types.Log) (*CSAccountingRoleAdminChanged, error) {
	event := new(CSAccountingRoleAdminChanged)
	if err := _CSAccounting.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the CSAccounting contract.
type CSAccountingRoleGrantedIterator struct {
	Event *CSAccountingRoleGranted // Event containing the contract specifics and raw log

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
func (it *CSAccountingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingRoleGranted)
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
		it.Event = new(CSAccountingRoleGranted)
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
func (it *CSAccountingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingRoleGranted represents a RoleGranted event raised by the CSAccounting contract.
type CSAccountingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CSAccounting *CSAccountingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CSAccountingRoleGrantedIterator, error) {

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

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingRoleGrantedIterator{contract: _CSAccounting.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_CSAccounting *CSAccountingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CSAccountingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingRoleGranted)
				if err := _CSAccounting.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_CSAccounting *CSAccountingFilterer) ParseRoleGranted(log types.Log) (*CSAccountingRoleGranted, error) {
	event := new(CSAccountingRoleGranted)
	if err := _CSAccounting.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the CSAccounting contract.
type CSAccountingRoleRevokedIterator struct {
	Event *CSAccountingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CSAccountingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingRoleRevoked)
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
		it.Event = new(CSAccountingRoleRevoked)
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
func (it *CSAccountingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingRoleRevoked represents a RoleRevoked event raised by the CSAccounting contract.
type CSAccountingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CSAccounting *CSAccountingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CSAccountingRoleRevokedIterator, error) {

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

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingRoleRevokedIterator{contract: _CSAccounting.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_CSAccounting *CSAccountingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CSAccountingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingRoleRevoked)
				if err := _CSAccounting.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_CSAccounting *CSAccountingFilterer) ParseRoleRevoked(log types.Log) (*CSAccountingRoleRevoked, error) {
	event := new(CSAccountingRoleRevoked)
	if err := _CSAccounting.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CSAccountingStETHSharesRecoveredIterator is returned from FilterStETHSharesRecovered and is used to iterate over the raw logs and unpacked data for StETHSharesRecovered events raised by the CSAccounting contract.
type CSAccountingStETHSharesRecoveredIterator struct {
	Event *CSAccountingStETHSharesRecovered // Event containing the contract specifics and raw log

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
func (it *CSAccountingStETHSharesRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CSAccountingStETHSharesRecovered)
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
		it.Event = new(CSAccountingStETHSharesRecovered)
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
func (it *CSAccountingStETHSharesRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CSAccountingStETHSharesRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CSAccountingStETHSharesRecovered represents a StETHSharesRecovered event raised by the CSAccounting contract.
type CSAccountingStETHSharesRecovered struct {
	Recipient common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStETHSharesRecovered is a free log retrieval operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_CSAccounting *CSAccountingFilterer) FilterStETHSharesRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CSAccountingStETHSharesRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSAccounting.contract.FilterLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CSAccountingStETHSharesRecoveredIterator{contract: _CSAccounting.contract, event: "StETHSharesRecovered", logs: logs, sub: sub}, nil
}

// WatchStETHSharesRecovered is a free log subscription operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_CSAccounting *CSAccountingFilterer) WatchStETHSharesRecovered(opts *bind.WatchOpts, sink chan<- *CSAccountingStETHSharesRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _CSAccounting.contract.WatchLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CSAccountingStETHSharesRecovered)
				if err := _CSAccounting.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
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
func (_CSAccounting *CSAccountingFilterer) ParseStETHSharesRecovered(log types.Log) (*CSAccountingStETHSharesRecovered, error) {
	event := new(CSAccountingStETHSharesRecovered)
	if err := _CSAccounting.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
