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
	Intervals []ICSBondCurveBondCurveInterval
}

// ICSBondCurveBondCurveInterval is an auto generated low-level Go binding around an user-defined struct.
type ICSBondCurveBondCurveInterval struct {
	MinKeysCount *big.Int
	MinBond      *big.Int
	Trend        *big.Int
}

// ICSBondCurveBondCurveIntervalInput is an auto generated low-level Go binding around an user-defined struct.
type ICSBondCurveBondCurveIntervalInput struct {
	MinKeysCount *big.Int
	Trend        *big.Int
}

// ICSBondLockBondLock is an auto generated low-level Go binding around an user-defined struct.
type ICSBondLockBondLock struct {
	Amount *big.Int
	Until  *big.Int
}

// CsaccountingMetaData contains all meta data concerning the Csaccounting contract.
var CsaccountingMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lidoLocator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"module\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeDistributor\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minBondLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxBondLockPeriod\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AccessControlBadConfirmation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"neededRole\",\"type\":\"bytes32\"}],\"name\":\"AccessControlUnauthorizedAccount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ElRewardsVaultReceiveFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedToSendEther\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondCurveId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondCurveLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondCurveValues\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondCurvesLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondLockAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidBondLockPeriod\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitialization\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidInitializationCurveId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NodeOperatorDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotAllowedToRecover\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotInitializing\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NothingToClaim\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseUntilMustBeInFuture\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedExpected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResumedExpected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotEligible\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotModule\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroAdminAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroChargePenaltyRecipientAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFeeDistributorAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroLocatorAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroModuleAddress\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPauseDuration\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountToBurn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedAmount\",\"type\":\"uint256\"}],\"name\":\"BondBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toChargeAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chargedAmount\",\"type\":\"uint256\"}],\"name\":\"BondCharged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondClaimedStETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"name\":\"BondClaimedUnstETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondClaimedWstETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeysCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structICSBondCurve.BondCurveIntervalInput[]\",\"name\":\"bondCurveIntervals\",\"type\":\"tuple[]\"}],\"name\":\"BondCurveAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"BondCurveSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeysCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"indexed\":false,\"internalType\":\"structICSBondCurve.BondCurveIntervalInput[]\",\"name\":\"bondCurveIntervals\",\"type\":\"tuple[]\"}],\"name\":\"BondCurveUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondDepositedETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondDepositedStETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondDepositedWstETH\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"until\",\"type\":\"uint256\"}],\"name\":\"BondLockChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"BondLockCompensated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"BondLockPeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"BondLockRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"chargePenaltyRecipient\",\"type\":\"address\"}],\"name\":\"ChargePenaltyRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC1155Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"ERC20Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"ERC721Recovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"EtherRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"version\",\"type\":\"uint64\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"shares\",\"type\":\"uint256\"}],\"name\":\"StETHSharesRecovered\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_BOND_CURVE_ID\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"FEE_DISTRIBUTOR\",\"outputs\":[{\"internalType\":\"contractICSFeeDistributor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIDO\",\"outputs\":[{\"internalType\":\"contractILido\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LIDO_LOCATOR\",\"outputs\":[{\"internalType\":\"contractILidoLocator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_BOND_CURVES_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_BOND_LOCK_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_CURVE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_BOND_LOCK_PERIOD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_CURVE_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MODULE\",\"outputs\":[{\"internalType\":\"contractICSModule\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_INFINITELY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RECOVERER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SET_BOND_CURVE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WITHDRAWAL_QUEUE\",\"outputs\":[{\"internalType\":\"contractIWithdrawalQueue\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WSTETH\",\"outputs\":[{\"internalType\":\"contractIWstETH\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeysCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"internalType\":\"structICSBondCurve.BondCurveIntervalInput[]\",\"name\":\"bondCurve\",\"type\":\"tuple[]\"}],\"name\":\"addBondCurve\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"chargeFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"chargePenaltyRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stETHAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimRewardsStETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimedShares\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stETHAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimRewardsUnstETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"requestId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wstETHAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"claimRewardsWstETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimedWstETH\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"compensateLockedBondETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"depositETH\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stETHAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"depositStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stETHAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"depositStETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wstETHAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"depositWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"wstETHAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"internalType\":\"structICSAccounting.PermitInput\",\"name\":\"permit\",\"type\":\"tuple\"}],\"name\":\"depositWstETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeDistributor\",\"outputs\":[{\"internalType\":\"contractICSFeeDistributor\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeysCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"internalType\":\"structICSBondCurve.BondCurveIntervalInput[][]\",\"name\":\"bondCurvesInputs\",\"type\":\"tuple[][]\"}],\"name\":\"finalizeUpgradeV2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getActualLockedBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getBond\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keys\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getBondAmountByKeysCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"keysCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getBondAmountByKeysCountWstETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getBondCurve\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeysCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"internalType\":\"structICSBondCurve.BondCurveInterval[]\",\"name\":\"intervals\",\"type\":\"tuple[]\"}],\"internalType\":\"structICSBondCurve.BondCurve\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getBondCurveId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBondLockPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getBondShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getBondSummary\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getBondSummaryShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"current\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"required\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getClaimableBondShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"getClaimableRewardsAndBondShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimableShares\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getCurveInfo\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeysCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minBond\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"internalType\":\"structICSBondCurve.BondCurveInterval[]\",\"name\":\"intervals\",\"type\":\"tuple[]\"}],\"internalType\":\"structICSBondCurve.BondCurve\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurvesCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitializedVersion\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"getKeysCountByBondAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getLockedBondInfo\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"until\",\"type\":\"uint128\"}],\"internalType\":\"structICSBondLock.BondLock\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalKeys\",\"type\":\"uint256\"}],\"name\":\"getRequiredBondForNextKeys\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"additionalKeys\",\"type\":\"uint256\"}],\"name\":\"getRequiredBondForNextKeysWstETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResumeSinceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getUnbondedKeysCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"getUnbondedKeysCountToEject\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeysCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"internalType\":\"structICSBondCurve.BondCurveIntervalInput[]\",\"name\":\"bondCurve\",\"type\":\"tuple[]\"},{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"bondLockPeriod\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_chargePenaltyRecipient\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"lockBondETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"pauseFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"penalize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"cumulativeFeeShares\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"rewardsProof\",\"type\":\"bytes32[]\"}],\"name\":\"pullFeeRewards\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC1155\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"recoverERC20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"recoverERC721\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverEther\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"recoverStETHShares\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"releaseLockedBondETH\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renewBurnerAllowance\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"callerConfirmation\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"}],\"name\":\"setBondCurve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"period\",\"type\":\"uint256\"}],\"name\":\"setBondLockPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_chargePenaltyRecipient\",\"type\":\"address\"}],\"name\":\"setChargePenaltyRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"}],\"name\":\"settleLockedBondETH\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"applied\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalBondShares\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"curveId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint256\",\"name\":\"minKeysCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"trend\",\"type\":\"uint256\"}],\"internalType\":\"structICSBondCurve.BondCurveIntervalInput[]\",\"name\":\"bondCurve\",\"type\":\"tuple[]\"}],\"name\":\"updateBondCurve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405260043610610484575f3560e01c80638331618411610257578063cb11c5271161013f578063dcab7f83116100be578063ef038a5411610083578063ef038a5414610f95578063f3efecc414610fb4578063f3f449c714610fc8578063f428d69614610fe7578063f7966efe14611006578063fee6380514611025575f80fd5b8063dcab7f8314610ee6578063e1aa105d14610f05578063e3c1962414610f24578063e5220e3f14610f57578063ead42a6914610f76575f80fd5b8063d90f35dc11610104578063d90f35dc14610e23578063d963ae5514610e42578063d9fb643a14610e61578063dbba4b4814610e94578063dc38ea3d14610ec7575f80fd5b8063cb11c52714610d93578063ce19793f14610da7578063d2896ff414610dc6578063d547741f14610de5578063d8fe764214610e04575f80fd5b8063a217fddf116101d6578063b148db6a1161019b578063b148db6a14610cd7578063b187bd2614610cf6578063b2d03e4d14610d0a578063b3c6501514610d29578063b5b624bf14610d55578063ca15c87314610d74575f80fd5b8063a217fddf14610840578063a302ee3814610c5d578063a41a7f8b14610c71578063ac1781c814610c85578063acf1c94814610ca4575f80fd5b80638f6549ae1161021c5780638f6549ae14610bad5780639010d07c14610be157806391d1485414610c005780639b4c6c2714610c1f5780639c51610214610c3e575f80fd5b80638331618414610a875780638409d4fe14610b1d578063881fa03c14610b3c5780638980f11f14610b5b5780638b21f17014610b7a575f80fd5b80633f214bb21161037a578063546da24f116102f9578063684b21c9116102be578063684b21c9146109845780636910dcce146109a3578063699340f4146109d65780636e13f09914610a0957806374d70aea14610a35578063819d4cc614610a68575f80fd5b8063546da24f146108eb57806356022ae31461090a578063589ff76c1461093d5780635a73bdc8146109515780635c654ad914610965575f80fd5b80634bb22a721161033f5780634bb22a72146108675780634c7ed3d2146108865780635097ef59146108a557806352d8bfc2146108c45780635358fbda146108d8575f80fd5b80633f214bb2146107cf578063433cd6c3146107ee5780634342b3c11461080d578063443fbfef146108405780634b2ce9fe14610853575f80fd5b8063165123dd116104065780632e599054116103cb5780632e5990541461072c5780632f2ff15d1461073f57806336568abe1461075e578063389ed2671461077d5780633df6c438146107b0575f80fd5b8063165123dd1461064e5780631ce7cb8f1461066d578063248a9ca3146106a057806328846981146106da5780632de03aa1146106f9575f80fd5b8063094d3a341161044c578063094d3a34146105805780630b3d765a146105cb5780630d43e8ad146105ea57806313d1234b1461061c57806315b5c4771461063b575f80fd5b806301a5e9e31461048857806301ffc9a7146104ba578063046f7da2146104e95780630569b947146104ff57806306cd0e9014610549575b5f80fd5b348015610493575f80fd5b506104a76104a2366004615010565b611058565b6040519081526020015b60405180910390f35b3480156104c5575f80fd5b506104d96104d4366004615027565b61106a565b60405190151581526020016104b1565b3480156104f4575f80fd5b506104fd61108e565b005b34801561050a575f80fd5b506104a7610519366004615010565b5f9081527f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe1501602052604090205490565b348015610554575f80fd5b506104a7610563366004615010565b5f9081525f805160206157b2833981519152602052604090205490565b34801561058b575f80fd5b506105b37f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f81565b6040516001600160a01b0390911681526020016104b1565b3480156105d6575f80fd5b506104fd6105e5366004615095565b6110c3565b3480156105f5575f80fd5b507f000000000000000000000000d99cc66fec647e68294c6477b40fc7e0f6f618d06105b3565b348015610627575f80fd5b506104a76106363660046150d3565b61122f565b6104fd610649366004615010565b611249565b348015610659575f80fd5b506001546105b3906001600160a01b031681565b348015610678575f80fd5b506104a77f000000000000000000000000000000000000000000000000000000000024ea0081565b3480156106ab575f80fd5b506104a76106ba366004615010565b5f9081525f80516020615812833981519152602052604090206001015490565b3480156106e5575f80fd5b506104a76106f43660046150d3565b6113c9565b348015610704575f80fd5b506104a77f2fc10cc8ae19568712f7a176fb4978616a610650813c9d05326c34abb62749c781565b6104fd61073a366004615107565b6113d7565b34801561074a575f80fd5b506104fd610759366004615131565b611436565b348015610769575f80fd5b506104fd610778366004615131565b61146c565b348015610788575f80fd5b506104a77f139c2898040ef16910dc9f44dc697df79363da767d8bc92f2e310312b816e46d81565b3480156107bb575f80fd5b506104a76107ca36600461515f565b6114a4565b3480156107da575f80fd5b506104fd6107e93660046151ca565b6115e4565b3480156107f9575f80fd5b506104fd6108083660046151fd565b611684565b348015610818575f80fd5b506104a77f645c9e6d2a86805cb5a28b1e4751c0dab493df7cf935070ce405489ba1a7bf7281565b34801561084b575f80fd5b506104a75f81565b34801561085e575f80fd5b506104a7606481565b348015610872575f80fd5b506104d9610881366004615010565b611697565b348015610891575f80fd5b506104fd6108a0366004615218565b611713565b3480156108b0575f80fd5b506104a76108bf36600461515f565b611779565b3480156108cf575f80fd5b506104fd611837565b6104fd6108e6366004615010565b611893565b3480156108f6575f80fd5b506104a76109053660046150d3565b611926565b348015610915575f80fd5b507f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe1502546104a7565b348015610948575f80fd5b506104a7611939565b34801561095c575f80fd5b506104fd611967565b348015610970575f80fd5b506104fd61097f366004615107565b611ab5565b34801561098f575f80fd5b506104fd61099e366004615010565b611b30565b3480156109ae575f80fd5b506105b37f000000000000000000000000d99cc66fec647e68294c6477b40fc7e0f6f618d081565b3480156109e1575f80fd5b506105b37f000000000000000000000000889edc2edab5f40e902b864ad4d7ade8e412f9b181565b348015610a14575f80fd5b50610a28610a23366004615010565b611b43565b6040516104b1919061525f565b348015610a40575f80fd5b507f23f334b9eb5378c2a1573857b8f9d9ca79959360a69e73d3f16848e56ec92101546104a7565b348015610a73575f80fd5b506104fd610a82366004615107565b611c0c565b348015610a92575f80fd5b50610af6610aa1366004615010565b6040805180820182525f80825260209182018190529283525f805160206157f28339815191528152918190208151808301909252546001600160801b038082168352600160801b909104169181019190915290565b6040805182516001600160801b0390811682526020938401511692810192909252016104b1565b348015610b28575f80fd5b506104a7610b3736600461515f565b611c5b565b348015610b47575f80fd5b506104fd610b563660046150d3565b611d19565b348015610b66575f80fd5b506104fd610b75366004615107565b611d7b565b348015610b85575f80fd5b506105b37f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe8481565b348015610bb8575f80fd5b50610bcc610bc7366004615010565b611e1c565b604080519283526020830191909152016104b1565b348015610bec575f80fd5b506105b3610bfb3660046150d3565b611e48565b348015610c0b575f80fd5b506104d9610c1a366004615131565b611e88565b348015610c2a575f80fd5b506104fd610c393660046152c3565b611ebe565b348015610c49575f80fd5b506104a7610c58366004615010565b611f4e565b348015610c68575f80fd5b506104a75f1981565b348015610c7c575f80fd5b506104a7611f59565b348015610c90575f80fd5b506104a7610c9f366004615010565b611f82565b348015610caf575f80fd5b506104a77fb3e25b5404b87e5a838579cb5d7481d61ad96ee284d38ec1e97c07ba64e7f6fc81565b348015610ce2575f80fd5b506104a7610cf13660046150d3565b611f8c565b348015610d01575f80fd5b506104d9611fc0565b348015610d15575f80fd5b506104fd610d243660046150d3565b611ff0565b348015610d34575f80fd5b50610d3d61202d565b6040516001600160401b0390911681526020016104b1565b348015610d60575f80fd5b50610a28610d6f366004615010565b61205f565b348015610d7f575f80fd5b506104a7610d8e366004615010565b612077565b348015610d9e575f80fd5b506104a7600181565b348015610db2575f80fd5b50610bcc610dc1366004615010565b6120ae565b348015610dd1575f80fd5b506104fd610de0366004615351565b6120c5565b348015610df0575f80fd5b506104fd610dff366004615131565b612473565b348015610e0f575f80fd5b506104a7610e1e366004615010565b6124a3565b348015610e2e575f80fd5b506104a7610e3d3660046152c3565b6124c6565b348015610e4d575f80fd5b506104fd610e5c3660046150d3565b61259a565b348015610e6c575f80fd5b506105b37f0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca081565b348015610e9f575f80fd5b506105b37f000000000000000000000000c1d0b3de6792bf6b4b37eccdcc24e45978cfd2eb81565b348015610ed2575f80fd5b506104a7610ee13660046150d3565b6125ed565b348015610ef1575f80fd5b506104fd610f003660046150d3565b612600565b348015610f10575f80fd5b506104fd610f1f3660046151ca565b612653565b348015610f2f575f80fd5b506104a77f0000000000000000000000000000000000000000000000000000000001e1338081565b348015610f62575f80fd5b506104fd610f713660046150d3565b612679565b348015610f81575f80fd5b506104a7610f90366004615010565b6126cc565b348015610fa0575f80fd5b506104fd610faf3660046153bf565b612720565b348015610fbf575f80fd5b506104fd612755565b348015610fd3575f80fd5b506104fd610fe2366004615010565b612872565b348015610ff2575f80fd5b506104a7611001366004615406565b6128a5565b348015611011575f80fd5b506104fd611020366004615218565b6128da565b348015611030575f80fd5b506104a77fd35e4a788498271198ec69c34f1dc762a1eee8200c111f598da1b3dde946783d81565b5f611064826001612940565b92915050565b5f6001600160e01b03198216635a05180f60e01b1480611064575061106482612a59565b7f2fc10cc8ae19568712f7a176fb4978616a610650813c9d05326c34abb62749c76110b881612a8d565b6110c0612a97565b50565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff168061110c575080546001600160401b03808416911610155b1561112a5760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b1781555f8055611156612aec565b83146111755760405163ed0f2e3b60e01b815260040160405180910390fd5b6111a184845f81811061118a5761118a615438565b905060200281019061119c919061544c565b612b00565b60015b838110156111e1576111d88585838181106111c1576111c1615438565b90506020028101906111d3919061544c565b612b34565b506001016111a4565b50805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050565b5f61124261123d8484611926565b612bb3565b9392505050565b336001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f1614611292576040516303f249a160e51b815260040160405180910390fd5b5f7f000000000000000000000000c1d0b3de6792bf6b4b37eccdcc24e45978cfd2eb6001600160a01b031663e441d25f6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112ef573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113139190615491565b6001600160a01b0316346040515f6040518083038185875af1925050503d805f811461135a576040519150601f19603f3d011682016040523d82523d5f602084013e61135f565b606091505b5050905080611381576040516324f09be760e21b815260040160405180910390fd5b61138b8234612c4a565b817fb6ee6e3aae6776519627b46786a622b642c38cabfe4c97cb34054fd63fc11a23346040516113bd91815260200190565b60405180910390a25050565b5f61124261123d8484611f8c565b6113df612cd0565b336001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f1614611428576040516303f249a160e51b815260040160405180910390fd5b6114328282612cf8565b5050565b5f8281525f80516020615812833981519152602052604090206001015461145c81612a8d565b6114668383612de5565b50505050565b6001600160a01b03811633146114955760405163334bd91960e11b815260040160405180910390fd5b61149f8282612e3a565b505050565b5f6114ad612cd0565b6040516324cdc74d60e11b8152600481018790525f907f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f6001600160a01b03169063499b8e9a90602401606060405180830381865afa158015611512573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906115369190615504565b905061154181612e86565b82156115535761155387868686612ef3565b61156287878360200151612f93565b6040516308eab3cd60e41b8152600481018990529092507f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f6001600160a01b031690638eab3cd0906024015f604051808303815f87803b1580156115c4575f80fd5b505af11580156115d6573d5f803e3d5ffd5b505050505095945050505050565b6115ec612cd0565b6115f583613271565b6115ff3382613314565b61160a338484613483565b6040516308eab3cd60e41b8152600481018490527f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f6001600160a01b031690638eab3cd0906024015f604051808303815f87803b158015611669575f80fd5b505af115801561167b573d5f803e3d5ffd5b50505050505050565b5f61168e81612a8d565b6114328261370e565b5f336001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f16146116e1576040516303f249a160e51b815260040160405180910390fd5b505f806116ed836126cc565b9050801561170d576116ff8382613797565b611708836138ed565b600191505b50919050565b61171b612cd0565b336001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f1614611764576040516303f249a160e51b815260040160405180910390fd5b61176e8482613933565b611466848484613a20565b5f611782612cd0565b6040516324cdc74d60e11b8152600481018790525f907f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f6001600160a01b03169063499b8e9a90602401606060405180830381865afa1580156117e7573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061180b9190615504565b905061181681612e86565b82156118285761182887868686612ef3565b61156287878360200151613b1d565b61183f613de3565b73a74528edc289b1a597faf83fcff7eff871cc01d96352d8bfc26040518163ffffffff1660e01b81526004015f6040518083038186803b158015611881575f80fd5b505af4158015611466573d5f803e3d5ffd5b61189b612cd0565b6118a481613271565b6118ae3382612cf8565b6040516308eab3cd60e41b8152600481018290527f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f6001600160a01b031690638eab3cd0906024015f604051808303815f87803b15801561190d575f80fd5b505af115801561191f573d5f803e3d5ffd5b5050505050565b5f6112428361193484613e0c565b613e87565b5f6119627fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a025490565b905090565b61196f613de3565b5f6119987f23f334b9eb5378c2a1573857b8f9d9ca79959360a69e73d3f16848e56ec921015490565b604051633d7ad0b760e21b81523060048201527f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b03169063f5eb42dc90602401602060405180830381865afa1580156119fa573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a1e9190615571565b611a28919061559c565b6040516389ad944360e01b81526001600160a01b037f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe841660048201526024810182905290915073a74528edc289b1a597faf83fcff7eff871cc01d9906389ad9443906044015f6040518083038186803b158015611aa3575f80fd5b505af415801561191f573d5f803e3d5ffd5b611abd613de3565b604051635c654ad960e01b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d990635c654ad9906044015b5f6040518083038186803b158015611b16575f80fd5b505af4158015611b28573d5f803e3d5ffd5b505050505050565b5f611b3a81612a8d565b61143282613f31565b604080516020810190915260608152611b8e611b89835f9081527f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe1501602052604090205490565b613e0c565b6040805182546020818102830184018452820181815291939284929184915f9085015b82821015611bfe578382905f5260205f2090600302016040518060600160405290815f82015481526020016001820154815260200160028201548152505081526020019060010190611bb1565b505050915250909392505050565b611c14613de3565b6040516340cea66360e11b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d99063819d4cc690604401611b00565b5f611c64612cd0565b6040516324cdc74d60e11b8152600481018790525f907f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f6001600160a01b03169063499b8e9a90602401606060405180830381865afa158015611cc9573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611ced9190615504565b9050611cf881612e86565b8215611d0a57611d0a87868686612ef3565b61156287878360200151613fef565b336001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f1614611d62576040516303f249a160e51b815260040160405180910390fd5b60015461143290839083906001600160a01b031661412a565b611d83613de3565b7f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b0316826001600160a01b031603611dd5576040516319efe5d760e21b815260040160405180910390fd5b604051638980f11f60e01b81526001600160a01b03831660048201526024810182905273a74528edc289b1a597faf83fcff7eff871cc01d990638980f11f90604401611b00565b5f8181525f805160206157b2833981519152602052604081205490611e418382614231565b9050915091565b5f8281527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000602081905260408220611e80908461423f565b949350505050565b5f9182525f80516020615812833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b611ec784613271565b611ed384848484612ef3565b6040516308eab3cd60e41b8152600481018590527f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f6001600160a01b031690638eab3cd0906024015f604051808303815f87803b158015611f32575f80fd5b505af1158015611f44573d5f803e3d5ffd5b5050505050505050565b5f611064825f612940565b5f7f78c5a36767279da056404c09083fca30cf3ea61c442cfaba6669f76a37393f005b54919050565b5f6110648261424a565b5f80611f97846124a3565b90505f611fa4858561426f565b9050818111611fb3575f611fb7565b8181035b95945050505050565b5f611fe97fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a025490565b4210905090565b7f645c9e6d2a86805cb5a28b1e4751c0dab493df7cf935070ce405489ba1a7bf7261201a81612a8d565b61202383613271565b61160a8383614363565b5f6119627ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00546001600160401b031690565b604080516020810190915260608152611b8e82613e0c565b5f8181527fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e82371705932000602081905260408220611242906143fd565b5f806120b9836124a3565b9150611e41835f61426f565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805460029190600160401b900460ff168061210e575080546001600160401b03808416911610155b1561212c5760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b178155612155614406565b61215f8787612b00565b6121688461440e565b6001600160a01b03851661218f57604051633ef39b8160e01b815260040160405180910390fd5b6121995f86612de5565b506121a38361370e565b60405163095ea7b360e01b81526001600160a01b037f0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca0811660048301525f1960248301527f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe84169063095ea7b3906044016020604051808303815f875af115801561222f573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061225391906155af565b5060405163095ea7b360e01b81526001600160a01b037f000000000000000000000000889edc2edab5f40e902b864ad4d7ade8e412f9b1811660048301525f1960248301527f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe84169063095ea7b3906044016020604051808303815f875af11580156122e0573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061230491906155af565b507f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b031663095ea7b37f000000000000000000000000c1d0b3de6792bf6b4b37eccdcc24e45978cfd2eb6001600160a01b03166327810b6e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015612390573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123b49190615491565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201525f1960248201526044016020604051808303815f875af11580156123fe573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061242291906155af565b50805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a150505050505050565b5f8281525f80516020615812833981519152602052604090206001015461249981612a8d565b6114668383612e3a565b5f8181525f805160206157b283398151915260205260408120546110649061441f565b5f807f000000000000000000000000d99cc66fec647e68294c6477b40fc7e0f6f618d06001600160a01b0316635e8e8f6f878787876040518563ffffffff1660e01b815260040161251a94939291906155c8565b602060405180830381865afa158015612535573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906125599190615571565b90505f8061256688611e1c565b9092509050612575838361560f565b9150808211612584575f61258e565b61258e818361559c565b98975050505050505050565b336001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f16146125e3576040516303f249a160e51b815260040160405180910390fd5b6114328282612c4a565b5f611242836125fb84613e0c565b61447b565b336001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f1614612649576040516303f249a160e51b815260040160405180910390fd5b61143282826145a5565b61265b612cd0565b61266483613271565b61266e3382613933565b61160a338484613a20565b336001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f16146126c2576040516303f249a160e51b815260040160405180910390fd5b6114328282613797565b5f8181525f805160206157f283398151915260205260408120805442600160801b9091046001600160801b031611612704575f612710565b80546001600160801b03165b6001600160801b03169392505050565b7fd35e4a788498271198ec69c34f1dc762a1eee8200c111f598da1b3dde946783d61274a81612a8d565b61146684848461464e565b7f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b031663095ea7b37f000000000000000000000000c1d0b3de6792bf6b4b37eccdcc24e45978cfd2eb6001600160a01b03166327810b6e6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156127e0573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906128049190615491565b6040516001600160e01b031960e084901b1681526001600160a01b0390911660048201525f1960248201526044016020604051808303815f875af115801561284e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110c091906155af565b7f139c2898040ef16910dc9f44dc697df79363da767d8bc92f2e310312b816e46d61289c81612a8d565b61143282614730565b5f7fd35e4a788498271198ec69c34f1dc762a1eee8200c111f598da1b3dde946783d6128d081612a8d565b611e808484612b34565b6128e2612cd0565b336001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f161461292b576040516303f249a160e51b815260040160405180910390fd5b6129358482613314565b611466848484613483565b6040516311d8d20560e31b8152600481018390525f9081906001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f1690638ec6902890602401602060405180830381865afa1580156129a7573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906129cb9190615571565b90505f6129d7856124a3565b90508315612a01575f6129e9866126cc565b9050818111156129fe57829350505050611064565b90035b5f8581527f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe15016020526040812054612a3c90600a8401906125ed565b9050808311612a4b575f612a4f565b8083035b9695505050505050565b5f6001600160e01b03198216637965db0b60e01b148061106457506301ffc9a760e01b6001600160e01b0319831614611064565b6110c0813361477f565b612a9f6147bd565b427fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a02556040517f62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9905f90a1565b5f5f805160206157d2833981519152611f7c565b612b086147e2565b5f612b138383612b34565b9050801561149f576040516320315aa760e01b815260040160405180910390fd5b5f5f805160206157d2833981519152612b4d848461482b565b6002810180546001810182555f9182526020909120909250612b729083018585614973565b817f707691ca33c3fcf1738eeb4c10826bd3030b3687166d6de80eb5896067fd21598585604051612ba4929190615622565b60405180910390a25092915050565b5f815f03612bc257505f919050565b604051631920845160e01b8152600481018390527f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b0316906319208451906024015b602060405180830381865afa158015612c26573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110649190615571565b805f03612c6a57604051633649e09760e11b815260040160405180910390fd5b5f612c74836126cc565b905081811015612c9757604051633649e09760e11b815260040160405180910390fd5b5f8381525f805160206157f2833981519152602052604090205461149f90849084840390600160801b90046001600160801b0316614ae1565b612cd8611fc0565b15612cf657604051630286f07360e31b815260040160405180910390fd5b565b345f03612d03575050565b60405163a1903eab60e01b81525f60048201819052906001600160a01b037f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe84169063a1903eab90349060240160206040518083038185885af1158015612d6b573d5f803e3d5ffd5b50505050506040513d601f19601f82011682018060405250810190612d909190615571565b9050612d9c8282614b95565b604080516001600160a01b038516815234602082015283917f16ec5116295424dec7fd52c87d9971a963ea7f59f741ad9ad468f0312055dc4991015b60405180910390a2505050565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200081612e128585614bea565b90508015611e80575f858152602083905260409020612e319085614c92565b50949350505050565b5f7fc1f6fe24621ce81ec5827caf0253cadb74709b061630e6b55e8237170593200081612e678585614ca6565b90508015611e80575f858152602083905260409020612e319085614d1f565b80516001600160a01b0316612eae57604051633ed893db60e21b815260040160405180910390fd5b80516001600160a01b0316331480612ed2575060208101516001600160a01b031633145b15612eda5750565b60405163743a3f7960e11b815260040160405180910390fd5b6040516321893f7b60e01b81525f906001600160a01b037f000000000000000000000000d99cc66fec647e68294c6477b40fc7e0f6f618d016906321893f7b90612f479088908890889088906004016155c8565b6020604051808303815f875af1158015612f63573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612f879190615571565b905061191f8582614b95565b5f80612f9e8561424a565b90505f612faa8261441f565b8510612fb65781612fbf565b612fbf85612bb3565b9050805f03612fe1576040516312d37ee560e31b815260040160405180910390fd5b6040805160018082528183019092525f916020808301908036833701905050905061300b8261441f565b815f8151811061301d5761301d615438565b6020908102919091010152604051633d7ad0b760e21b81523060048201525f907f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b03169063f5eb42dc90602401602060405180830381865afa15801561308c573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906130b09190615571565b604051636b34082160e11b81529091506001600160a01b037f000000000000000000000000889edc2edab5f40e902b864ad4d7ade8e412f9b1169063d6681042906131019085908a90600401615668565b5f604051808303815f875af115801561311c573d5f803e3d5ffd5b505050506040513d5f823e601f3d908101601f1916820160405261314391908101906156be565b5f8151811061315457613154615438565b6020908102919091010151604051633d7ad0b760e21b81523060048201529095505f906001600160a01b037f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe84169063f5eb42dc90602401602060405180830381865afa1580156131c6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906131ea9190615571565b90506131ff896131fa838561559c565b614d33565b887f26673a9d018b21192d08ee14377b798f11b9e5b15ea1559c110265716b8985b588855f8151811061323457613234615438565b602090810291909101810151604080516001600160a01b03909416845291830152810189905260600160405180910390a250505050509392505050565b7f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f6001600160a01b031663a70c70e46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156132cd573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906132f19190615571565b8110156132fb5750565b604051633ed893db60e21b815260040160405180910390fd5b8035158015906133b45750604051636eb1769f60e11b81526001600160a01b0383811660048301523060248301528235917f0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca09091169063dd62ed3e90604401602060405180830381865afa15801561338e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906133b29190615571565b105b15611432576001600160a01b037f0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca01663d505accf8330843560208601356134016060880160408901615752565b6040516001600160e01b031960e088901b1681526001600160a01b0395861660048201529490931660248501526044840191909152606483015260ff166084820152606084013560a4820152608084013560c482015260e4015f604051808303815f87803b158015613471575f80fd5b505af1158015611b28573d5f803e3d5ffd5b805f0361348f57505050565b6040516323b872dd60e01b81526001600160a01b038481166004830152306024830152604482018390527f0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca016906323b872dd906064015f604051808303815f87803b1580156134fc575f80fd5b505af115801561350e573d5f803e3d5ffd5b5050604051633d7ad0b760e21b81523060048201525f92507f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b0316915063f5eb42dc90602401602060405180830381865afa158015613576573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061359a9190615571565b604051636f074d1f60e11b8152600481018490529091507f0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca06001600160a01b03169063de0e9a3e906024016020604051808303815f875af1158015613601573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136259190615571565b50604051633d7ad0b760e21b81523060048201525f907f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b03169063f5eb42dc90602401602060405180830381865afa15801561368a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906136ae9190615571565b90506136c3846136be848461559c565b614b95565b604080516001600160a01b03871681526020810185905285917f6576bbc9c5b478bf9717dc3d2bcb485e5ff0727df77c72558727597f3609d3f1910160405180910390a25050505050565b6001600160a01b03811661373557604051631279f7c160e21b815260040160405180910390fd5b6001805473ffffffffffffffffffffffffffffffffffffffff19166001600160a01b0383169081179091556040519081527f4beaaee83871b066b675515d6a53567e76411f60266703cef934a01905a4d832906020015b60405180910390a150565b5f6137a182612bb3565b90505f6137ae8483614d83565b9050805f036137bd5750505050565b7f000000000000000000000000c1d0b3de6792bf6b4b37eccdcc24e45978cfd2eb6001600160a01b03166327810b6e6040518163ffffffff1660e01b8152600401602060405180830381865afa158015613819573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061383d9190615491565b6040516308c2292560e31b8152306004820152602481018390526001600160a01b0391909116906346114928906044015f604051808303815f87803b158015613884575f80fd5b505af1158015613896573d5f803e3d5ffd5b50505050837f4da924ae7845fe96897faab524b536685b8bbc4d82fbb45c10d941e0f86ade0f6138c58461441f565b6138ce8461441f565b604080519283526020830191909152015b60405180910390a250505050565b5f8181525f805160206157f283398151915260205260408082208290555182917f844ae6b00e8a437dcdde1a634feab3273e08bb5c274a4be3b195b308ae0ba20a91a250565b8035158015906139d35750604051636eb1769f60e11b81526001600160a01b0383811660048301523060248301528235917f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe849091169063dd62ed3e90604401602060405180830381865afa1580156139ad573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906139d19190615571565b105b15611432576001600160a01b037f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe841663d505accf8330843560208601356134016060880160408901615752565b805f03613a2c57505050565b5f613a3682612bb3565b604051636d78045960e01b81526001600160a01b038681166004830152306024830152604482018390529192507f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe8490911690636d780459906064016020604051808303815f875af1158015613aad573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613ad19190615571565b50613adc8382614b95565b604080516001600160a01b03861681526020810184905284917fee31ebba29fd5471227e42fd8ca621a892d689901892cb8febb03fe802c3214b91016138df565b5f80613b288561424a565b90505f818510613b385781613b3a565b845b9050805f03613b5c576040516312d37ee560e31b815260040160405180910390fd5b604051633d7ad0b760e21b81523060048201525f907f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b03169063f5eb42dc90602401602060405180830381865afa158015613bc0573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613be49190615571565b90507f0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca06001600160a01b031663ea598cb0613c1e8461441f565b6040518263ffffffff1660e01b8152600401613c3c91815260200190565b6020604051808303815f875af1158015613c58573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613c7c9190615571565b604051633d7ad0b760e21b81523060048201529094505f906001600160a01b037f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe84169063f5eb42dc90602401602060405180830381865afa158015613ce3573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190613d079190615571565b9050613d17886131fa838561559c565b60405163a9059cbb60e01b81526001600160a01b038781166004830152602482018790527f0000000000000000000000007f39c581f595b53c5cb19bd0b3f8da6c935e2ca0169063a9059cbb906044015f604051808303815f87803b158015613d7e575f80fd5b505af1158015613d90573d5f803e3d5ffd5b5050604080516001600160a01b038a168152602081018990528b93507fe6a8c06447e05a412e5e9581e088941f3994db3d8a9bfd3275b38d77acacafac92500160405180910390a2505050509392505050565b612cf67fb3e25b5404b87e5a838579cb5d7481d61ad96ee284d38ec1e97c07ba64e7f6fc612a8d565b7f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe1502545f905f805160206157d2833981519152905f1901831115613e62576040516331e784e960e11b815260040160405180910390fd5b806002018381548110613e7757613e77615438565b905f5260205f2001915050919050565b5f81838203613e99575f915050611064565b80545f905f19015b80821015613ef2575f6002600184840101049050838181548110613ec757613ec7615438565b905f5260205f2090600302015f0154871015613ee857600181039150613eec565b8092505b50613ea1565b5f838381548110613f0557613f05615438565b905f5260205f20906003020190508060020154815f015488030281600101540194505050505092915050565b7f000000000000000000000000000000000000000000000000000000000024ea00811080613f7e57507f0000000000000000000000000000000000000000000000000000000001e1338081115b15613f9c5760405163326bd91760e01b815260040160405180910390fd5b807f78c5a36767279da056404c09083fca30cf3ea61c442cfaba6669f76a37393f00556040518181527fd117ae9105bfc4a5acf683370984ce7aea9498aa2849fc0851e0b012552b31039060200161378c565b5f80613ffa8561424a565b90506140058161441f565b8410614011578061401a565b61401a84612bb3565b9150815f0361403c576040516312d37ee560e31b815260040160405180910390fd5b6140468583614d33565b604051638fcb4e5b60e01b81526001600160a01b038481166004830152602482018490525f917f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe8490911690638fcb4e5b906044016020604051808303815f875af11580156140b6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906140da9190615571565b604080516001600160a01b03871681526020810183905291925087917f3e3a1398fe71575ed0c17a80cd9d46ad684c2c75c2fad7b0e7dde15e78ab22d3910160405180910390a250509392505050565b5f61413483612bb3565b90505f6141418583614d83565b9050805f03614151575050505050565b604051638fcb4e5b60e01b81526001600160a01b038481166004830152602482018390525f917f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe8490911690638fcb4e5b906044016020604051808303815f875af11580156141c1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906141e59190615571565b9050857f8615528474a7bb3a28d9971535d956b79242b8e8fcfb27f3e331270fff088afd6142128561441f565b60408051918252602082018590520160405180910390a2505050505050565b5f61124261123d848461426f565b5f6112428383614db8565b5f805f61425684611e1c565b91509150808211614267575f611e80565b900392915050565b5f8281527f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe1501602052604081205481906040516311d8d20560e31b8152600481018690529091505f906001600160a01b037f000000000000000000000000da7de2ecddfccc6c3af10108db212acbbf9ea83f1690638ec6902890602401602060405180830381865afa158015614306573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061432a9190615571565b90505f61434061433a868461560f565b84611926565b90505f61434c876126cc565b9050614358818361560f565b979650505050505050565b7f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe1502545f805160206157d2833981519152905f19018211156143b7576040516331e784e960e11b815260040160405180910390fd5b5f838152600182016020526040908190208390555183907f4642db1736894887bc907d721f20af84d3e585a0a3cea90f41b78b2aa906541b90612dd89085815260200190565b5f611064825490565b612cf66147e2565b6144166147e2565b6110c081613f31565b5f815f0361442e57505f919050565b604051630f451f7160e31b8152600481018390527f000000000000000000000000ae7ab96520de3a18e5e111b5eaab095312d7fe846001600160a01b031690637a28fb8890602401612c0b565b80545f9082908190839061449157614491615438565b905f5260205f209060030201600101548410156144b1575f915050611064565b80545f905f19015b8082101561450b575f60026001848401010490508381815481106144df576144df615438565b905f5260205f2090600302016001015487101561450157600181039150614505565b8092505b506144b9565b82545f905f190183101561455d5783836001018154811061452e5761452e615438565b905f5260205f2090600302019050806002015481600101540387111561455d57545f1901935061106492505050565b83838154811061456f5761456f615438565b905f5260205f20906003020190508060020154816001015488038161459657614596615772565b91549104019695505050505050565b7f78c5a36767279da056404c09083fca30cf3ea61c442cfaba6669f76a37393f005f8290036145e757604051633649e09760e11b815260040160405180910390fd5b5f8381526001820160209081526040918290208251808401909352546001600160801b038082168452600160801b9091041690820181905242101561463e57805161463b906001600160801b03168461560f565b92505b6114668484845f01544201614ae1565b7f8f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe1502545f805160206157d2833981519152905f19018411156146a2576040516331e784e960e11b815260040160405180910390fd5b6146ac838361482b565b8060020184815481106146c1576146c1615438565b5f9182526020822001906146d58282614fd6565b50506146fe8160020185815481106146ef576146ef615438565b905f5260205f20018484614973565b837f77c7f59d9ea0a6ee0417e777c399834e7ce0647a7ece2b12f4dbff0a6a1980c884846040516138df929190615622565b614738612cd0565b805f036147585760405163ad58bfc760e01b815260040160405180910390fd5b5f5f19820361476957505f19614776565b614773824261560f565b90505b61143281614dde565b6147898282611e88565b6114325760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044015b60405180910390fd5b6147c5611fc0565b612cf65760405163b047186b60e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16612cf657604051631afcd79f60e31b815260040160405180910390fd5b600181108061483a5750606481115b1561485857604051638326bf5360e01b815260040160405180910390fd5b81815f81811061486a5761486a615438565b9050604002015f0135600114614893576040516302527aef60e21b815260040160405180910390fd5b81815f8181106148a5576148a5615438565b905060400201602001355f036148ce576040516302527aef60e21b815260040160405180910390fd5b60015b8181101561149f578282600183038181106148ee576148ee615438565b9050604002015f013583838381811061490957614909615438565b9050604002015f013511614930576040516302527aef60e21b815260040160405180910390fd5b82828281811061494257614942615438565b905060400201602001355f0361496b576040516302527aef60e21b815260040160405180910390fd5b6001016148d1565b82546001810184555f84815260208120600390920290910190839083908161499d5761499d615438565b604002919091013582555082825f816149b8576149b8615438565b90506040020160200135816002018190555082825f8181106149dc576149dc615438565b9050604002016020013581600101819055505f600190505b8281101561191f5784546001810186555f8681526020902060039091020191848483818110614a2557614a25615438565b6040029190910135845550848483818110614a4257614a42615438565b90506040020160200135836002018190555080600201546001825f0154878786818110614a7157614a71615438565b9050604002015f0135614a84919061559c565b614a8e919061559c565b614a989190615786565b8160010154868685818110614aaf57614aaf615438565b90506040020160200135614ac3919061560f565b614acd919061560f565b8360010181905550508060010190506149f4565b815f03614af15761149f836138ed565b6040518060400160405280614b0584614e79565b6001600160801b03168152602001614b1c83614e79565b6001600160801b039081169091525f8581525f805160206157f283398151915260209081526040918290208451948201518416600160801b029490931693909317909155805184815291820183905284917f69a153d448f54b17f05cf3b268a2efab87c94a4727d108c4ca4aa3e5d65113de9101612dd8565b805f03614ba0575050565b5f9182525f805160206157b283398151915260205260409091208054820190557f23f334b9eb5378c2a1573857b8f9d9ca79959360a69e73d3f16848e56ec9210180549091019055565b5f5f80516020615812833981519152614c038484611e88565b614c82575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055614c383390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a46001915050611064565b5f915050611064565b5092915050565b5f611242836001600160a01b038416614eb0565b5f5f80516020615812833981519152614cbf8484611e88565b15614c82575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a46001915050611064565b5f611242836001600160a01b038416614efc565b5f8281525f805160206157b2833981519152602081905260408220805491928492614d5f90849061559c565b9250508190555081816001015f828254614d79919061559c565b9091555050505050565b5f8281525f805160206157b28339815191526020526040812054808310614daa5780614dac565b825b9150614c8b8483614d33565b5f825f018281548110614dcd57614dcd615438565b905f5260205f200154905092915050565b614e077fe8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a02829055565b5f198103614e40576040515f1981527f32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e9060200161378c565b7f32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e614e6b428361559c565b60405190815260200161378c565b5f6001600160801b03821115614eac576040516306dfcc6560e41b815260806004820152602481018390526044016147b4565b5090565b5f818152600183016020526040812054614ef557508154600181810184555f848152602080822090930184905584548482528286019093526040902091909155611064565b505f611064565b5f8181526001830160205260408120548015614c82575f614f1e60018361559c565b85549091505f90614f319060019061559c565b9050808214614f90575f865f018281548110614f4f57614f4f615438565b905f5260205f200154905080875f018481548110614f6f57614f6f615438565b5f918252602080832090910192909255918252600188019052604090208390555b8554869080614fa157614fa161579d565b600190038181905f5260205f20015f90559055856001015f8681526020019081526020015f205f905560019350505050611064565b5080545f8255600302905f5260205f20908101906110c091905b80821115614eac575f808255600182018190556002820155600301614ff0565b5f60208284031215615020575f80fd5b5035919050565b5f60208284031215615037575f80fd5b81356001600160e01b031981168114611242575f80fd5b5f8083601f84011261505e575f80fd5b5081356001600160401b03811115615074575f80fd5b6020830191508360208260051b850101111561508e575f80fd5b9250929050565b5f80602083850312156150a6575f80fd5b82356001600160401b038111156150bb575f80fd5b6150c78582860161504e565b90969095509350505050565b5f80604083850312156150e4575f80fd5b50508035926020909101359150565b6001600160a01b03811681146110c0575f80fd5b5f8060408385031215615118575f80fd5b8235615123816150f3565b946020939093013593505050565b5f8060408385031215615142575f80fd5b823591506020830135615154816150f3565b809150509250929050565b5f805f805f60808688031215615173575f80fd5b85359450602086013593506040860135925060608601356001600160401b0381111561519d575f80fd5b6151a98882890161504e565b969995985093965092949392505050565b5f60a0828403121561170d575f80fd5b5f805f60e084860312156151dc575f80fd5b83359250602084013591506151f485604086016151ba565b90509250925092565b5f6020828403121561520d575f80fd5b8135611242816150f3565b5f805f80610100858703121561522c575f80fd5b8435615237816150f3565b9350602085013592506040850135915061525486606087016151ba565b905092959194509250565b5f60208083526040808401855183848701528181518084526060935060608801915085830192505f5b818110156152b5578351805184528781015188850152860151868401529286019291840191600101615288565b509098975050505050505050565b5f805f80606085870312156152d6575f80fd5b843593506020850135925060408501356001600160401b038111156152f9575f80fd5b6153058782880161504e565b95989497509550505050565b5f8083601f840112615321575f80fd5b5081356001600160401b03811115615337575f80fd5b6020830191508360208260061b850101111561508e575f80fd5b5f805f805f60808688031215615365575f80fd5b85356001600160401b0381111561537a575f80fd5b61538688828901615311565b909650945050602086013561539a816150f3565b92506040860135915060608601356153b1816150f3565b809150509295509295909350565b5f805f604084860312156153d1575f80fd5b8335925060208401356001600160401b038111156153ed575f80fd5b6153f986828701615311565b9497909650939450505050565b5f8060208385031215615417575f80fd5b82356001600160401b0381111561542c575f80fd5b6150c785828601615311565b634e487b7160e01b5f52603260045260245ffd5b5f808335601e19843603018112615461575f80fd5b8301803591506001600160401b0382111561547a575f80fd5b6020019150600681901b360382131561508e575f80fd5b5f602082840312156154a1575f80fd5b8151611242816150f3565b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f191681016001600160401b03811182821017156154e8576154e86154ac565b604052919050565b805180151581146154ff575f80fd5b919050565b5f60608284031215615514575f80fd5b604051606081018181106001600160401b0382111715615536576155366154ac565b6040528251615544816150f3565b81526020830151615554816150f3565b6020820152615565604084016154f0565b60408201529392505050565b5f60208284031215615581575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b8181038181111561106457611064615588565b5f602082840312156155bf575f80fd5b611242826154f0565b8481526020810184905260606040820181905281018290525f6001600160fb1b038311156155f4575f80fd5b8260051b808560808501379190910160800195945050505050565b8082018082111561106457611064615588565b60208082528181018390525f90604080840186845b8781101561565b578135835284820135858401529183019190830190600101615637565b5090979650505050505050565b604080825283519082018190525f906020906060840190828701845b828110156156a057815184529284019290840190600101615684565b50505080925050506001600160a01b03831660208301529392505050565b5f60208083850312156156cf575f80fd5b82516001600160401b03808211156156e5575f80fd5b818501915085601f8301126156f8575f80fd5b81518181111561570a5761570a6154ac565b8060051b915061571b8483016154c0565b8181529183018401918481019088841115615734575f80fd5b938501935b8385101561258e57845182529385019390850190615739565b5f60208284031215615762575f80fd5b813560ff81168114611242575f80fd5b634e487b7160e01b5f52601260045260245ffd5b808202811582820484141761106457611064615588565b634e487b7160e01b5f52603160045260245ffdfe23f334b9eb5378c2a1573857b8f9d9ca79959360a69e73d3f16848e56ec921008f22e270e477f5becb8793b61d439ab7ae990ed8eba045eb72061c0e6cfe150078c5a36767279da056404c09083fca30cf3ea61c442cfaba6669f76a37393f0102dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800a164736f6c6343000818000a",
}

// CsaccountingABI is the input ABI used to generate the binding from.
// Deprecated: Use CsaccountingMetaData.ABI instead.
var CsaccountingABI = CsaccountingMetaData.ABI

// CsaccountingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CsaccountingMetaData.Bin instead.
var CsaccountingBin = CsaccountingMetaData.Bin

// DeployCsaccounting deploys a new Ethereum contract, binding an instance of Csaccounting to it.
func DeployCsaccounting(auth *bind.TransactOpts, backend bind.ContractBackend, lidoLocator common.Address, module common.Address, _feeDistributor common.Address, minBondLockPeriod *big.Int, maxBondLockPeriod *big.Int) (common.Address, *types.Transaction, *Csaccounting, error) {
	parsed, err := CsaccountingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CsaccountingBin), backend, lidoLocator, module, _feeDistributor, minBondLockPeriod, maxBondLockPeriod)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Csaccounting{CsaccountingCaller: CsaccountingCaller{contract: contract}, CsaccountingTransactor: CsaccountingTransactor{contract: contract}, CsaccountingFilterer: CsaccountingFilterer{contract: contract}}, nil
}

// Csaccounting is an auto generated Go binding around an Ethereum contract.
type Csaccounting struct {
	CsaccountingCaller     // Read-only binding to the contract
	CsaccountingTransactor // Write-only binding to the contract
	CsaccountingFilterer   // Log filterer for contract events
}

// CsaccountingCaller is an auto generated read-only Go binding around an Ethereum contract.
type CsaccountingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsaccountingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CsaccountingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsaccountingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CsaccountingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CsaccountingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CsaccountingSession struct {
	Contract     *Csaccounting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CsaccountingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CsaccountingCallerSession struct {
	Contract *CsaccountingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// CsaccountingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CsaccountingTransactorSession struct {
	Contract     *CsaccountingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// CsaccountingRaw is an auto generated low-level Go binding around an Ethereum contract.
type CsaccountingRaw struct {
	Contract *Csaccounting // Generic contract binding to access the raw methods on
}

// CsaccountingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CsaccountingCallerRaw struct {
	Contract *CsaccountingCaller // Generic read-only contract binding to access the raw methods on
}

// CsaccountingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CsaccountingTransactorRaw struct {
	Contract *CsaccountingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCsaccounting creates a new instance of Csaccounting, bound to a specific deployed contract.
func NewCsaccounting(address common.Address, backend bind.ContractBackend) (*Csaccounting, error) {
	contract, err := bindCsaccounting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Csaccounting{CsaccountingCaller: CsaccountingCaller{contract: contract}, CsaccountingTransactor: CsaccountingTransactor{contract: contract}, CsaccountingFilterer: CsaccountingFilterer{contract: contract}}, nil
}

// NewCsaccountingCaller creates a new read-only instance of Csaccounting, bound to a specific deployed contract.
func NewCsaccountingCaller(address common.Address, caller bind.ContractCaller) (*CsaccountingCaller, error) {
	contract, err := bindCsaccounting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CsaccountingCaller{contract: contract}, nil
}

// NewCsaccountingTransactor creates a new write-only instance of Csaccounting, bound to a specific deployed contract.
func NewCsaccountingTransactor(address common.Address, transactor bind.ContractTransactor) (*CsaccountingTransactor, error) {
	contract, err := bindCsaccounting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CsaccountingTransactor{contract: contract}, nil
}

// NewCsaccountingFilterer creates a new log filterer instance of Csaccounting, bound to a specific deployed contract.
func NewCsaccountingFilterer(address common.Address, filterer bind.ContractFilterer) (*CsaccountingFilterer, error) {
	contract, err := bindCsaccounting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CsaccountingFilterer{contract: contract}, nil
}

// bindCsaccounting binds a generic wrapper to an already deployed contract.
func bindCsaccounting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CsaccountingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csaccounting *CsaccountingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csaccounting.Contract.CsaccountingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csaccounting *CsaccountingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csaccounting.Contract.CsaccountingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csaccounting *CsaccountingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csaccounting.Contract.CsaccountingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Csaccounting *CsaccountingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Csaccounting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Csaccounting *CsaccountingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csaccounting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Csaccounting *CsaccountingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Csaccounting.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Csaccounting.Contract.DEFAULTADMINROLE(&_Csaccounting.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Csaccounting.Contract.DEFAULTADMINROLE(&_Csaccounting.CallOpts)
}

// DEFAULTBONDCURVEID is a free data retrieval call binding the contract method 0x443fbfef.
//
// Solidity: function DEFAULT_BOND_CURVE_ID() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) DEFAULTBONDCURVEID(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "DEFAULT_BOND_CURVE_ID")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULTBONDCURVEID is a free data retrieval call binding the contract method 0x443fbfef.
//
// Solidity: function DEFAULT_BOND_CURVE_ID() view returns(uint256)
func (_Csaccounting *CsaccountingSession) DEFAULTBONDCURVEID() (*big.Int, error) {
	return _Csaccounting.Contract.DEFAULTBONDCURVEID(&_Csaccounting.CallOpts)
}

// DEFAULTBONDCURVEID is a free data retrieval call binding the contract method 0x443fbfef.
//
// Solidity: function DEFAULT_BOND_CURVE_ID() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) DEFAULTBONDCURVEID() (*big.Int, error) {
	return _Csaccounting.Contract.DEFAULTBONDCURVEID(&_Csaccounting.CallOpts)
}

// FEEDISTRIBUTOR is a free data retrieval call binding the contract method 0x6910dcce.
//
// Solidity: function FEE_DISTRIBUTOR() view returns(address)
func (_Csaccounting *CsaccountingCaller) FEEDISTRIBUTOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "FEE_DISTRIBUTOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FEEDISTRIBUTOR is a free data retrieval call binding the contract method 0x6910dcce.
//
// Solidity: function FEE_DISTRIBUTOR() view returns(address)
func (_Csaccounting *CsaccountingSession) FEEDISTRIBUTOR() (common.Address, error) {
	return _Csaccounting.Contract.FEEDISTRIBUTOR(&_Csaccounting.CallOpts)
}

// FEEDISTRIBUTOR is a free data retrieval call binding the contract method 0x6910dcce.
//
// Solidity: function FEE_DISTRIBUTOR() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) FEEDISTRIBUTOR() (common.Address, error) {
	return _Csaccounting.Contract.FEEDISTRIBUTOR(&_Csaccounting.CallOpts)
}

// LIDO is a free data retrieval call binding the contract method 0x8b21f170.
//
// Solidity: function LIDO() view returns(address)
func (_Csaccounting *CsaccountingCaller) LIDO(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "LIDO")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIDO is a free data retrieval call binding the contract method 0x8b21f170.
//
// Solidity: function LIDO() view returns(address)
func (_Csaccounting *CsaccountingSession) LIDO() (common.Address, error) {
	return _Csaccounting.Contract.LIDO(&_Csaccounting.CallOpts)
}

// LIDO is a free data retrieval call binding the contract method 0x8b21f170.
//
// Solidity: function LIDO() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) LIDO() (common.Address, error) {
	return _Csaccounting.Contract.LIDO(&_Csaccounting.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_Csaccounting *CsaccountingCaller) LIDOLOCATOR(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "LIDO_LOCATOR")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_Csaccounting *CsaccountingSession) LIDOLOCATOR() (common.Address, error) {
	return _Csaccounting.Contract.LIDOLOCATOR(&_Csaccounting.CallOpts)
}

// LIDOLOCATOR is a free data retrieval call binding the contract method 0xdbba4b48.
//
// Solidity: function LIDO_LOCATOR() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) LIDOLOCATOR() (common.Address, error) {
	return _Csaccounting.Contract.LIDOLOCATOR(&_Csaccounting.CallOpts)
}

// MANAGEBONDCURVESROLE is a free data retrieval call binding the contract method 0xfee63805.
//
// Solidity: function MANAGE_BOND_CURVES_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) MANAGEBONDCURVESROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "MANAGE_BOND_CURVES_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGEBONDCURVESROLE is a free data retrieval call binding the contract method 0xfee63805.
//
// Solidity: function MANAGE_BOND_CURVES_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) MANAGEBONDCURVESROLE() ([32]byte, error) {
	return _Csaccounting.Contract.MANAGEBONDCURVESROLE(&_Csaccounting.CallOpts)
}

// MANAGEBONDCURVESROLE is a free data retrieval call binding the contract method 0xfee63805.
//
// Solidity: function MANAGE_BOND_CURVES_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) MANAGEBONDCURVESROLE() ([32]byte, error) {
	return _Csaccounting.Contract.MANAGEBONDCURVESROLE(&_Csaccounting.CallOpts)
}

// MAXBONDLOCKPERIOD is a free data retrieval call binding the contract method 0xe3c19624.
//
// Solidity: function MAX_BOND_LOCK_PERIOD() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) MAXBONDLOCKPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "MAX_BOND_LOCK_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBONDLOCKPERIOD is a free data retrieval call binding the contract method 0xe3c19624.
//
// Solidity: function MAX_BOND_LOCK_PERIOD() view returns(uint256)
func (_Csaccounting *CsaccountingSession) MAXBONDLOCKPERIOD() (*big.Int, error) {
	return _Csaccounting.Contract.MAXBONDLOCKPERIOD(&_Csaccounting.CallOpts)
}

// MAXBONDLOCKPERIOD is a free data retrieval call binding the contract method 0xe3c19624.
//
// Solidity: function MAX_BOND_LOCK_PERIOD() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) MAXBONDLOCKPERIOD() (*big.Int, error) {
	return _Csaccounting.Contract.MAXBONDLOCKPERIOD(&_Csaccounting.CallOpts)
}

// MAXCURVELENGTH is a free data retrieval call binding the contract method 0x4b2ce9fe.
//
// Solidity: function MAX_CURVE_LENGTH() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) MAXCURVELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "MAX_CURVE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXCURVELENGTH is a free data retrieval call binding the contract method 0x4b2ce9fe.
//
// Solidity: function MAX_CURVE_LENGTH() view returns(uint256)
func (_Csaccounting *CsaccountingSession) MAXCURVELENGTH() (*big.Int, error) {
	return _Csaccounting.Contract.MAXCURVELENGTH(&_Csaccounting.CallOpts)
}

// MAXCURVELENGTH is a free data retrieval call binding the contract method 0x4b2ce9fe.
//
// Solidity: function MAX_CURVE_LENGTH() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) MAXCURVELENGTH() (*big.Int, error) {
	return _Csaccounting.Contract.MAXCURVELENGTH(&_Csaccounting.CallOpts)
}

// MINBONDLOCKPERIOD is a free data retrieval call binding the contract method 0x1ce7cb8f.
//
// Solidity: function MIN_BOND_LOCK_PERIOD() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) MINBONDLOCKPERIOD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "MIN_BOND_LOCK_PERIOD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBONDLOCKPERIOD is a free data retrieval call binding the contract method 0x1ce7cb8f.
//
// Solidity: function MIN_BOND_LOCK_PERIOD() view returns(uint256)
func (_Csaccounting *CsaccountingSession) MINBONDLOCKPERIOD() (*big.Int, error) {
	return _Csaccounting.Contract.MINBONDLOCKPERIOD(&_Csaccounting.CallOpts)
}

// MINBONDLOCKPERIOD is a free data retrieval call binding the contract method 0x1ce7cb8f.
//
// Solidity: function MIN_BOND_LOCK_PERIOD() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) MINBONDLOCKPERIOD() (*big.Int, error) {
	return _Csaccounting.Contract.MINBONDLOCKPERIOD(&_Csaccounting.CallOpts)
}

// MINCURVELENGTH is a free data retrieval call binding the contract method 0xcb11c527.
//
// Solidity: function MIN_CURVE_LENGTH() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) MINCURVELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "MIN_CURVE_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINCURVELENGTH is a free data retrieval call binding the contract method 0xcb11c527.
//
// Solidity: function MIN_CURVE_LENGTH() view returns(uint256)
func (_Csaccounting *CsaccountingSession) MINCURVELENGTH() (*big.Int, error) {
	return _Csaccounting.Contract.MINCURVELENGTH(&_Csaccounting.CallOpts)
}

// MINCURVELENGTH is a free data retrieval call binding the contract method 0xcb11c527.
//
// Solidity: function MIN_CURVE_LENGTH() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) MINCURVELENGTH() (*big.Int, error) {
	return _Csaccounting.Contract.MINCURVELENGTH(&_Csaccounting.CallOpts)
}

// MODULE is a free data retrieval call binding the contract method 0x094d3a34.
//
// Solidity: function MODULE() view returns(address)
func (_Csaccounting *CsaccountingCaller) MODULE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "MODULE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MODULE is a free data retrieval call binding the contract method 0x094d3a34.
//
// Solidity: function MODULE() view returns(address)
func (_Csaccounting *CsaccountingSession) MODULE() (common.Address, error) {
	return _Csaccounting.Contract.MODULE(&_Csaccounting.CallOpts)
}

// MODULE is a free data retrieval call binding the contract method 0x094d3a34.
//
// Solidity: function MODULE() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) MODULE() (common.Address, error) {
	return _Csaccounting.Contract.MODULE(&_Csaccounting.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) PAUSEINFINITELY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "PAUSE_INFINITELY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Csaccounting *CsaccountingSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Csaccounting.Contract.PAUSEINFINITELY(&_Csaccounting.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Csaccounting.Contract.PAUSEINFINITELY(&_Csaccounting.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) PAUSEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.PAUSEROLE(&_Csaccounting.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) PAUSEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.PAUSEROLE(&_Csaccounting.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) RECOVERERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "RECOVERER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) RECOVERERROLE() ([32]byte, error) {
	return _Csaccounting.Contract.RECOVERERROLE(&_Csaccounting.CallOpts)
}

// RECOVERERROLE is a free data retrieval call binding the contract method 0xacf1c948.
//
// Solidity: function RECOVERER_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) RECOVERERROLE() ([32]byte, error) {
	return _Csaccounting.Contract.RECOVERERROLE(&_Csaccounting.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) RESUMEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.RESUMEROLE(&_Csaccounting.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) RESUMEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.RESUMEROLE(&_Csaccounting.CallOpts)
}

// SETBONDCURVEROLE is a free data retrieval call binding the contract method 0x4342b3c1.
//
// Solidity: function SET_BOND_CURVE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) SETBONDCURVEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "SET_BOND_CURVE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SETBONDCURVEROLE is a free data retrieval call binding the contract method 0x4342b3c1.
//
// Solidity: function SET_BOND_CURVE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingSession) SETBONDCURVEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.SETBONDCURVEROLE(&_Csaccounting.CallOpts)
}

// SETBONDCURVEROLE is a free data retrieval call binding the contract method 0x4342b3c1.
//
// Solidity: function SET_BOND_CURVE_ROLE() view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) SETBONDCURVEROLE() ([32]byte, error) {
	return _Csaccounting.Contract.SETBONDCURVEROLE(&_Csaccounting.CallOpts)
}

// WITHDRAWALQUEUE is a free data retrieval call binding the contract method 0x699340f4.
//
// Solidity: function WITHDRAWAL_QUEUE() view returns(address)
func (_Csaccounting *CsaccountingCaller) WITHDRAWALQUEUE(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "WITHDRAWAL_QUEUE")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WITHDRAWALQUEUE is a free data retrieval call binding the contract method 0x699340f4.
//
// Solidity: function WITHDRAWAL_QUEUE() view returns(address)
func (_Csaccounting *CsaccountingSession) WITHDRAWALQUEUE() (common.Address, error) {
	return _Csaccounting.Contract.WITHDRAWALQUEUE(&_Csaccounting.CallOpts)
}

// WITHDRAWALQUEUE is a free data retrieval call binding the contract method 0x699340f4.
//
// Solidity: function WITHDRAWAL_QUEUE() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) WITHDRAWALQUEUE() (common.Address, error) {
	return _Csaccounting.Contract.WITHDRAWALQUEUE(&_Csaccounting.CallOpts)
}

// WSTETH is a free data retrieval call binding the contract method 0xd9fb643a.
//
// Solidity: function WSTETH() view returns(address)
func (_Csaccounting *CsaccountingCaller) WSTETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "WSTETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WSTETH is a free data retrieval call binding the contract method 0xd9fb643a.
//
// Solidity: function WSTETH() view returns(address)
func (_Csaccounting *CsaccountingSession) WSTETH() (common.Address, error) {
	return _Csaccounting.Contract.WSTETH(&_Csaccounting.CallOpts)
}

// WSTETH is a free data retrieval call binding the contract method 0xd9fb643a.
//
// Solidity: function WSTETH() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) WSTETH() (common.Address, error) {
	return _Csaccounting.Contract.WSTETH(&_Csaccounting.CallOpts)
}

// ChargePenaltyRecipient is a free data retrieval call binding the contract method 0x165123dd.
//
// Solidity: function chargePenaltyRecipient() view returns(address)
func (_Csaccounting *CsaccountingCaller) ChargePenaltyRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "chargePenaltyRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ChargePenaltyRecipient is a free data retrieval call binding the contract method 0x165123dd.
//
// Solidity: function chargePenaltyRecipient() view returns(address)
func (_Csaccounting *CsaccountingSession) ChargePenaltyRecipient() (common.Address, error) {
	return _Csaccounting.Contract.ChargePenaltyRecipient(&_Csaccounting.CallOpts)
}

// ChargePenaltyRecipient is a free data retrieval call binding the contract method 0x165123dd.
//
// Solidity: function chargePenaltyRecipient() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) ChargePenaltyRecipient() (common.Address, error) {
	return _Csaccounting.Contract.ChargePenaltyRecipient(&_Csaccounting.CallOpts)
}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_Csaccounting *CsaccountingCaller) FeeDistributor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "feeDistributor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_Csaccounting *CsaccountingSession) FeeDistributor() (common.Address, error) {
	return _Csaccounting.Contract.FeeDistributor(&_Csaccounting.CallOpts)
}

// FeeDistributor is a free data retrieval call binding the contract method 0x0d43e8ad.
//
// Solidity: function feeDistributor() view returns(address)
func (_Csaccounting *CsaccountingCallerSession) FeeDistributor() (common.Address, error) {
	return _Csaccounting.Contract.FeeDistributor(&_Csaccounting.CallOpts)
}

// GetActualLockedBond is a free data retrieval call binding the contract method 0xead42a69.
//
// Solidity: function getActualLockedBond(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetActualLockedBond(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getActualLockedBond", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActualLockedBond is a free data retrieval call binding the contract method 0xead42a69.
//
// Solidity: function getActualLockedBond(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetActualLockedBond(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetActualLockedBond(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetActualLockedBond is a free data retrieval call binding the contract method 0xead42a69.
//
// Solidity: function getActualLockedBond(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetActualLockedBond(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetActualLockedBond(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBond(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBond", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBond(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBond(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBond is a free data retrieval call binding the contract method 0xd8fe7642.
//
// Solidity: function getBond(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBond(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBond(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondAmountByKeysCount is a free data retrieval call binding the contract method 0x546da24f.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBondAmountByKeysCount(opts *bind.CallOpts, keys *big.Int, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondAmountByKeysCount", keys, curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondAmountByKeysCount is a free data retrieval call binding the contract method 0x546da24f.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBondAmountByKeysCount(keys *big.Int, curveId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondAmountByKeysCount(&_Csaccounting.CallOpts, keys, curveId)
}

// GetBondAmountByKeysCount is a free data retrieval call binding the contract method 0x546da24f.
//
// Solidity: function getBondAmountByKeysCount(uint256 keys, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBondAmountByKeysCount(keys *big.Int, curveId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondAmountByKeysCount(&_Csaccounting.CallOpts, keys, curveId)
}

// GetBondAmountByKeysCountWstETH is a free data retrieval call binding the contract method 0x13d1234b.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBondAmountByKeysCountWstETH(opts *bind.CallOpts, keysCount *big.Int, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondAmountByKeysCountWstETH", keysCount, curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondAmountByKeysCountWstETH is a free data retrieval call binding the contract method 0x13d1234b.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBondAmountByKeysCountWstETH(keysCount *big.Int, curveId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondAmountByKeysCountWstETH(&_Csaccounting.CallOpts, keysCount, curveId)
}

// GetBondAmountByKeysCountWstETH is a free data retrieval call binding the contract method 0x13d1234b.
//
// Solidity: function getBondAmountByKeysCountWstETH(uint256 keysCount, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBondAmountByKeysCountWstETH(keysCount *big.Int, curveId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondAmountByKeysCountWstETH(&_Csaccounting.CallOpts, keysCount, curveId)
}

// GetBondCurve is a free data retrieval call binding the contract method 0x6e13f099.
//
// Solidity: function getBondCurve(uint256 nodeOperatorId) view returns(((uint256,uint256,uint256)[]))
func (_Csaccounting *CsaccountingCaller) GetBondCurve(opts *bind.CallOpts, nodeOperatorId *big.Int) (ICSBondCurveBondCurve, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondCurve", nodeOperatorId)

	if err != nil {
		return *new(ICSBondCurveBondCurve), err
	}

	out0 := *abi.ConvertType(out[0], new(ICSBondCurveBondCurve)).(*ICSBondCurveBondCurve)

	return out0, err

}

// GetBondCurve is a free data retrieval call binding the contract method 0x6e13f099.
//
// Solidity: function getBondCurve(uint256 nodeOperatorId) view returns(((uint256,uint256,uint256)[]))
func (_Csaccounting *CsaccountingSession) GetBondCurve(nodeOperatorId *big.Int) (ICSBondCurveBondCurve, error) {
	return _Csaccounting.Contract.GetBondCurve(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondCurve is a free data retrieval call binding the contract method 0x6e13f099.
//
// Solidity: function getBondCurve(uint256 nodeOperatorId) view returns(((uint256,uint256,uint256)[]))
func (_Csaccounting *CsaccountingCallerSession) GetBondCurve(nodeOperatorId *big.Int) (ICSBondCurveBondCurve, error) {
	return _Csaccounting.Contract.GetBondCurve(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondCurveId is a free data retrieval call binding the contract method 0x0569b947.
//
// Solidity: function getBondCurveId(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBondCurveId(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondCurveId", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondCurveId is a free data retrieval call binding the contract method 0x0569b947.
//
// Solidity: function getBondCurveId(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBondCurveId(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondCurveId(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondCurveId is a free data retrieval call binding the contract method 0x0569b947.
//
// Solidity: function getBondCurveId(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBondCurveId(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondCurveId(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondLockPeriod is a free data retrieval call binding the contract method 0xa41a7f8b.
//
// Solidity: function getBondLockPeriod() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBondLockPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondLockPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondLockPeriod is a free data retrieval call binding the contract method 0xa41a7f8b.
//
// Solidity: function getBondLockPeriod() view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBondLockPeriod() (*big.Int, error) {
	return _Csaccounting.Contract.GetBondLockPeriod(&_Csaccounting.CallOpts)
}

// GetBondLockPeriod is a free data retrieval call binding the contract method 0xa41a7f8b.
//
// Solidity: function getBondLockPeriod() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBondLockPeriod() (*big.Int, error) {
	return _Csaccounting.Contract.GetBondLockPeriod(&_Csaccounting.CallOpts)
}

// GetBondShares is a free data retrieval call binding the contract method 0x06cd0e90.
//
// Solidity: function getBondShares(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetBondShares(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondShares", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBondShares is a free data retrieval call binding the contract method 0x06cd0e90.
//
// Solidity: function getBondShares(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetBondShares(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondShares(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondShares is a free data retrieval call binding the contract method 0x06cd0e90.
//
// Solidity: function getBondShares(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetBondShares(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetBondShares(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondSummary is a free data retrieval call binding the contract method 0xce19793f.
//
// Solidity: function getBondSummary(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_Csaccounting *CsaccountingCaller) GetBondSummary(opts *bind.CallOpts, nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondSummary", nodeOperatorId)

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
func (_Csaccounting *CsaccountingSession) GetBondSummary(nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	return _Csaccounting.Contract.GetBondSummary(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondSummary is a free data retrieval call binding the contract method 0xce19793f.
//
// Solidity: function getBondSummary(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_Csaccounting *CsaccountingCallerSession) GetBondSummary(nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	return _Csaccounting.Contract.GetBondSummary(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondSummaryShares is a free data retrieval call binding the contract method 0x8f6549ae.
//
// Solidity: function getBondSummaryShares(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_Csaccounting *CsaccountingCaller) GetBondSummaryShares(opts *bind.CallOpts, nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getBondSummaryShares", nodeOperatorId)

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
func (_Csaccounting *CsaccountingSession) GetBondSummaryShares(nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	return _Csaccounting.Contract.GetBondSummaryShares(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetBondSummaryShares is a free data retrieval call binding the contract method 0x8f6549ae.
//
// Solidity: function getBondSummaryShares(uint256 nodeOperatorId) view returns(uint256 current, uint256 required)
func (_Csaccounting *CsaccountingCallerSession) GetBondSummaryShares(nodeOperatorId *big.Int) (struct {
	Current  *big.Int
	Required *big.Int
}, error) {
	return _Csaccounting.Contract.GetBondSummaryShares(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetClaimableBondShares is a free data retrieval call binding the contract method 0xac1781c8.
//
// Solidity: function getClaimableBondShares(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetClaimableBondShares(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getClaimableBondShares", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableBondShares is a free data retrieval call binding the contract method 0xac1781c8.
//
// Solidity: function getClaimableBondShares(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetClaimableBondShares(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetClaimableBondShares(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetClaimableBondShares is a free data retrieval call binding the contract method 0xac1781c8.
//
// Solidity: function getClaimableBondShares(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetClaimableBondShares(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetClaimableBondShares(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetClaimableRewardsAndBondShares is a free data retrieval call binding the contract method 0xd90f35dc.
//
// Solidity: function getClaimableRewardsAndBondShares(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] rewardsProof) view returns(uint256 claimableShares)
func (_Csaccounting *CsaccountingCaller) GetClaimableRewardsAndBondShares(opts *bind.CallOpts, nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getClaimableRewardsAndBondShares", nodeOperatorId, cumulativeFeeShares, rewardsProof)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetClaimableRewardsAndBondShares is a free data retrieval call binding the contract method 0xd90f35dc.
//
// Solidity: function getClaimableRewardsAndBondShares(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] rewardsProof) view returns(uint256 claimableShares)
func (_Csaccounting *CsaccountingSession) GetClaimableRewardsAndBondShares(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*big.Int, error) {
	return _Csaccounting.Contract.GetClaimableRewardsAndBondShares(&_Csaccounting.CallOpts, nodeOperatorId, cumulativeFeeShares, rewardsProof)
}

// GetClaimableRewardsAndBondShares is a free data retrieval call binding the contract method 0xd90f35dc.
//
// Solidity: function getClaimableRewardsAndBondShares(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] rewardsProof) view returns(uint256 claimableShares)
func (_Csaccounting *CsaccountingCallerSession) GetClaimableRewardsAndBondShares(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*big.Int, error) {
	return _Csaccounting.Contract.GetClaimableRewardsAndBondShares(&_Csaccounting.CallOpts, nodeOperatorId, cumulativeFeeShares, rewardsProof)
}

// GetCurveInfo is a free data retrieval call binding the contract method 0xb5b624bf.
//
// Solidity: function getCurveInfo(uint256 curveId) view returns(((uint256,uint256,uint256)[]))
func (_Csaccounting *CsaccountingCaller) GetCurveInfo(opts *bind.CallOpts, curveId *big.Int) (ICSBondCurveBondCurve, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getCurveInfo", curveId)

	if err != nil {
		return *new(ICSBondCurveBondCurve), err
	}

	out0 := *abi.ConvertType(out[0], new(ICSBondCurveBondCurve)).(*ICSBondCurveBondCurve)

	return out0, err

}

// GetCurveInfo is a free data retrieval call binding the contract method 0xb5b624bf.
//
// Solidity: function getCurveInfo(uint256 curveId) view returns(((uint256,uint256,uint256)[]))
func (_Csaccounting *CsaccountingSession) GetCurveInfo(curveId *big.Int) (ICSBondCurveBondCurve, error) {
	return _Csaccounting.Contract.GetCurveInfo(&_Csaccounting.CallOpts, curveId)
}

// GetCurveInfo is a free data retrieval call binding the contract method 0xb5b624bf.
//
// Solidity: function getCurveInfo(uint256 curveId) view returns(((uint256,uint256,uint256)[]))
func (_Csaccounting *CsaccountingCallerSession) GetCurveInfo(curveId *big.Int) (ICSBondCurveBondCurve, error) {
	return _Csaccounting.Contract.GetCurveInfo(&_Csaccounting.CallOpts, curveId)
}

// GetCurvesCount is a free data retrieval call binding the contract method 0x56022ae3.
//
// Solidity: function getCurvesCount() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetCurvesCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getCurvesCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurvesCount is a free data retrieval call binding the contract method 0x56022ae3.
//
// Solidity: function getCurvesCount() view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetCurvesCount() (*big.Int, error) {
	return _Csaccounting.Contract.GetCurvesCount(&_Csaccounting.CallOpts)
}

// GetCurvesCount is a free data retrieval call binding the contract method 0x56022ae3.
//
// Solidity: function getCurvesCount() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetCurvesCount() (*big.Int, error) {
	return _Csaccounting.Contract.GetCurvesCount(&_Csaccounting.CallOpts)
}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_Csaccounting *CsaccountingCaller) GetInitializedVersion(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getInitializedVersion")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_Csaccounting *CsaccountingSession) GetInitializedVersion() (uint64, error) {
	return _Csaccounting.Contract.GetInitializedVersion(&_Csaccounting.CallOpts)
}

// GetInitializedVersion is a free data retrieval call binding the contract method 0xb3c65015.
//
// Solidity: function getInitializedVersion() view returns(uint64)
func (_Csaccounting *CsaccountingCallerSession) GetInitializedVersion() (uint64, error) {
	return _Csaccounting.Contract.GetInitializedVersion(&_Csaccounting.CallOpts)
}

// GetKeysCountByBondAmount is a free data retrieval call binding the contract method 0xdc38ea3d.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetKeysCountByBondAmount(opts *bind.CallOpts, amount *big.Int, curveId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getKeysCountByBondAmount", amount, curveId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetKeysCountByBondAmount is a free data retrieval call binding the contract method 0xdc38ea3d.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetKeysCountByBondAmount(amount *big.Int, curveId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetKeysCountByBondAmount(&_Csaccounting.CallOpts, amount, curveId)
}

// GetKeysCountByBondAmount is a free data retrieval call binding the contract method 0xdc38ea3d.
//
// Solidity: function getKeysCountByBondAmount(uint256 amount, uint256 curveId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetKeysCountByBondAmount(amount *big.Int, curveId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetKeysCountByBondAmount(&_Csaccounting.CallOpts, amount, curveId)
}

// GetLockedBondInfo is a free data retrieval call binding the contract method 0x83316184.
//
// Solidity: function getLockedBondInfo(uint256 nodeOperatorId) view returns((uint128,uint128))
func (_Csaccounting *CsaccountingCaller) GetLockedBondInfo(opts *bind.CallOpts, nodeOperatorId *big.Int) (ICSBondLockBondLock, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getLockedBondInfo", nodeOperatorId)

	if err != nil {
		return *new(ICSBondLockBondLock), err
	}

	out0 := *abi.ConvertType(out[0], new(ICSBondLockBondLock)).(*ICSBondLockBondLock)

	return out0, err

}

// GetLockedBondInfo is a free data retrieval call binding the contract method 0x83316184.
//
// Solidity: function getLockedBondInfo(uint256 nodeOperatorId) view returns((uint128,uint128))
func (_Csaccounting *CsaccountingSession) GetLockedBondInfo(nodeOperatorId *big.Int) (ICSBondLockBondLock, error) {
	return _Csaccounting.Contract.GetLockedBondInfo(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetLockedBondInfo is a free data retrieval call binding the contract method 0x83316184.
//
// Solidity: function getLockedBondInfo(uint256 nodeOperatorId) view returns((uint128,uint128))
func (_Csaccounting *CsaccountingCallerSession) GetLockedBondInfo(nodeOperatorId *big.Int) (ICSBondLockBondLock, error) {
	return _Csaccounting.Contract.GetLockedBondInfo(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetRequiredBondForNextKeys is a free data retrieval call binding the contract method 0xb148db6a.
//
// Solidity: function getRequiredBondForNextKeys(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetRequiredBondForNextKeys(opts *bind.CallOpts, nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getRequiredBondForNextKeys", nodeOperatorId, additionalKeys)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredBondForNextKeys is a free data retrieval call binding the contract method 0xb148db6a.
//
// Solidity: function getRequiredBondForNextKeys(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetRequiredBondForNextKeys(nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetRequiredBondForNextKeys(&_Csaccounting.CallOpts, nodeOperatorId, additionalKeys)
}

// GetRequiredBondForNextKeys is a free data retrieval call binding the contract method 0xb148db6a.
//
// Solidity: function getRequiredBondForNextKeys(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetRequiredBondForNextKeys(nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetRequiredBondForNextKeys(&_Csaccounting.CallOpts, nodeOperatorId, additionalKeys)
}

// GetRequiredBondForNextKeysWstETH is a free data retrieval call binding the contract method 0x28846981.
//
// Solidity: function getRequiredBondForNextKeysWstETH(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetRequiredBondForNextKeysWstETH(opts *bind.CallOpts, nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getRequiredBondForNextKeysWstETH", nodeOperatorId, additionalKeys)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRequiredBondForNextKeysWstETH is a free data retrieval call binding the contract method 0x28846981.
//
// Solidity: function getRequiredBondForNextKeysWstETH(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetRequiredBondForNextKeysWstETH(nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetRequiredBondForNextKeysWstETH(&_Csaccounting.CallOpts, nodeOperatorId, additionalKeys)
}

// GetRequiredBondForNextKeysWstETH is a free data retrieval call binding the contract method 0x28846981.
//
// Solidity: function getRequiredBondForNextKeysWstETH(uint256 nodeOperatorId, uint256 additionalKeys) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetRequiredBondForNextKeysWstETH(nodeOperatorId *big.Int, additionalKeys *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetRequiredBondForNextKeysWstETH(&_Csaccounting.CallOpts, nodeOperatorId, additionalKeys)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetResumeSinceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getResumeSinceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Csaccounting.Contract.GetResumeSinceTimestamp(&_Csaccounting.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Csaccounting.Contract.GetResumeSinceTimestamp(&_Csaccounting.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csaccounting *CsaccountingCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csaccounting *CsaccountingSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Csaccounting.Contract.GetRoleAdmin(&_Csaccounting.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Csaccounting *CsaccountingCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Csaccounting.Contract.GetRoleAdmin(&_Csaccounting.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csaccounting *CsaccountingCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csaccounting *CsaccountingSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Csaccounting.Contract.GetRoleMember(&_Csaccounting.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Csaccounting *CsaccountingCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Csaccounting.Contract.GetRoleMember(&_Csaccounting.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Csaccounting.Contract.GetRoleMemberCount(&_Csaccounting.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Csaccounting.Contract.GetRoleMemberCount(&_Csaccounting.CallOpts, role)
}

// GetUnbondedKeysCount is a free data retrieval call binding the contract method 0x01a5e9e3.
//
// Solidity: function getUnbondedKeysCount(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetUnbondedKeysCount(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getUnbondedKeysCount", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnbondedKeysCount is a free data retrieval call binding the contract method 0x01a5e9e3.
//
// Solidity: function getUnbondedKeysCount(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetUnbondedKeysCount(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetUnbondedKeysCount(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetUnbondedKeysCount is a free data retrieval call binding the contract method 0x01a5e9e3.
//
// Solidity: function getUnbondedKeysCount(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetUnbondedKeysCount(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetUnbondedKeysCount(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetUnbondedKeysCountToEject is a free data retrieval call binding the contract method 0x9c516102.
//
// Solidity: function getUnbondedKeysCountToEject(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCaller) GetUnbondedKeysCountToEject(opts *bind.CallOpts, nodeOperatorId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "getUnbondedKeysCountToEject", nodeOperatorId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUnbondedKeysCountToEject is a free data retrieval call binding the contract method 0x9c516102.
//
// Solidity: function getUnbondedKeysCountToEject(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingSession) GetUnbondedKeysCountToEject(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetUnbondedKeysCountToEject(&_Csaccounting.CallOpts, nodeOperatorId)
}

// GetUnbondedKeysCountToEject is a free data retrieval call binding the contract method 0x9c516102.
//
// Solidity: function getUnbondedKeysCountToEject(uint256 nodeOperatorId) view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) GetUnbondedKeysCountToEject(nodeOperatorId *big.Int) (*big.Int, error) {
	return _Csaccounting.Contract.GetUnbondedKeysCountToEject(&_Csaccounting.CallOpts, nodeOperatorId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csaccounting *CsaccountingCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csaccounting *CsaccountingSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Csaccounting.Contract.HasRole(&_Csaccounting.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Csaccounting *CsaccountingCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Csaccounting.Contract.HasRole(&_Csaccounting.CallOpts, role, account)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Csaccounting *CsaccountingCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Csaccounting *CsaccountingSession) IsPaused() (bool, error) {
	return _Csaccounting.Contract.IsPaused(&_Csaccounting.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Csaccounting *CsaccountingCallerSession) IsPaused() (bool, error) {
	return _Csaccounting.Contract.IsPaused(&_Csaccounting.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csaccounting *CsaccountingCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csaccounting *CsaccountingSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Csaccounting.Contract.SupportsInterface(&_Csaccounting.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Csaccounting *CsaccountingCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Csaccounting.Contract.SupportsInterface(&_Csaccounting.CallOpts, interfaceId)
}

// TotalBondShares is a free data retrieval call binding the contract method 0x74d70aea.
//
// Solidity: function totalBondShares() view returns(uint256)
func (_Csaccounting *CsaccountingCaller) TotalBondShares(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Csaccounting.contract.Call(opts, &out, "totalBondShares")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBondShares is a free data retrieval call binding the contract method 0x74d70aea.
//
// Solidity: function totalBondShares() view returns(uint256)
func (_Csaccounting *CsaccountingSession) TotalBondShares() (*big.Int, error) {
	return _Csaccounting.Contract.TotalBondShares(&_Csaccounting.CallOpts)
}

// TotalBondShares is a free data retrieval call binding the contract method 0x74d70aea.
//
// Solidity: function totalBondShares() view returns(uint256)
func (_Csaccounting *CsaccountingCallerSession) TotalBondShares() (*big.Int, error) {
	return _Csaccounting.Contract.TotalBondShares(&_Csaccounting.CallOpts)
}

// AddBondCurve is a paid mutator transaction binding the contract method 0xf428d696.
//
// Solidity: function addBondCurve((uint256,uint256)[] bondCurve) returns(uint256 id)
func (_Csaccounting *CsaccountingTransactor) AddBondCurve(opts *bind.TransactOpts, bondCurve []ICSBondCurveBondCurveIntervalInput) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "addBondCurve", bondCurve)
}

// AddBondCurve is a paid mutator transaction binding the contract method 0xf428d696.
//
// Solidity: function addBondCurve((uint256,uint256)[] bondCurve) returns(uint256 id)
func (_Csaccounting *CsaccountingSession) AddBondCurve(bondCurve []ICSBondCurveBondCurveIntervalInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.AddBondCurve(&_Csaccounting.TransactOpts, bondCurve)
}

// AddBondCurve is a paid mutator transaction binding the contract method 0xf428d696.
//
// Solidity: function addBondCurve((uint256,uint256)[] bondCurve) returns(uint256 id)
func (_Csaccounting *CsaccountingTransactorSession) AddBondCurve(bondCurve []ICSBondCurveBondCurveIntervalInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.AddBondCurve(&_Csaccounting.TransactOpts, bondCurve)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x881fa03c.
//
// Solidity: function chargeFee(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactor) ChargeFee(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "chargeFee", nodeOperatorId, amount)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x881fa03c.
//
// Solidity: function chargeFee(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingSession) ChargeFee(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.ChargeFee(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// ChargeFee is a paid mutator transaction binding the contract method 0x881fa03c.
//
// Solidity: function chargeFee(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactorSession) ChargeFee(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.ChargeFee(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0x8409d4fe.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns(uint256 claimedShares)
func (_Csaccounting *CsaccountingTransactor) ClaimRewardsStETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, stETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "claimRewardsStETH", nodeOperatorId, stETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0x8409d4fe.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns(uint256 claimedShares)
func (_Csaccounting *CsaccountingSession) ClaimRewardsStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.ClaimRewardsStETH(&_Csaccounting.TransactOpts, nodeOperatorId, stETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsStETH is a paid mutator transaction binding the contract method 0x8409d4fe.
//
// Solidity: function claimRewardsStETH(uint256 nodeOperatorId, uint256 stETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns(uint256 claimedShares)
func (_Csaccounting *CsaccountingTransactorSession) ClaimRewardsStETH(nodeOperatorId *big.Int, stETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.ClaimRewardsStETH(&_Csaccounting.TransactOpts, nodeOperatorId, stETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0x3df6c438.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns(uint256 requestId)
func (_Csaccounting *CsaccountingTransactor) ClaimRewardsUnstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, stETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "claimRewardsUnstETH", nodeOperatorId, stETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0x3df6c438.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns(uint256 requestId)
func (_Csaccounting *CsaccountingSession) ClaimRewardsUnstETH(nodeOperatorId *big.Int, stETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.ClaimRewardsUnstETH(&_Csaccounting.TransactOpts, nodeOperatorId, stETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsUnstETH is a paid mutator transaction binding the contract method 0x3df6c438.
//
// Solidity: function claimRewardsUnstETH(uint256 nodeOperatorId, uint256 stETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns(uint256 requestId)
func (_Csaccounting *CsaccountingTransactorSession) ClaimRewardsUnstETH(nodeOperatorId *big.Int, stETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.ClaimRewardsUnstETH(&_Csaccounting.TransactOpts, nodeOperatorId, stETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x5097ef59.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns(uint256 claimedWstETH)
func (_Csaccounting *CsaccountingTransactor) ClaimRewardsWstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, wstETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "claimRewardsWstETH", nodeOperatorId, wstETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x5097ef59.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns(uint256 claimedWstETH)
func (_Csaccounting *CsaccountingSession) ClaimRewardsWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.ClaimRewardsWstETH(&_Csaccounting.TransactOpts, nodeOperatorId, wstETHAmount, cumulativeFeeShares, rewardsProof)
}

// ClaimRewardsWstETH is a paid mutator transaction binding the contract method 0x5097ef59.
//
// Solidity: function claimRewardsWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns(uint256 claimedWstETH)
func (_Csaccounting *CsaccountingTransactorSession) ClaimRewardsWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.ClaimRewardsWstETH(&_Csaccounting.TransactOpts, nodeOperatorId, wstETHAmount, cumulativeFeeShares, rewardsProof)
}

// CompensateLockedBondETH is a paid mutator transaction binding the contract method 0x15b5c477.
//
// Solidity: function compensateLockedBondETH(uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingTransactor) CompensateLockedBondETH(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "compensateLockedBondETH", nodeOperatorId)
}

// CompensateLockedBondETH is a paid mutator transaction binding the contract method 0x15b5c477.
//
// Solidity: function compensateLockedBondETH(uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingSession) CompensateLockedBondETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.CompensateLockedBondETH(&_Csaccounting.TransactOpts, nodeOperatorId)
}

// CompensateLockedBondETH is a paid mutator transaction binding the contract method 0x15b5c477.
//
// Solidity: function compensateLockedBondETH(uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingTransactorSession) CompensateLockedBondETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.CompensateLockedBondETH(&_Csaccounting.TransactOpts, nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address from, uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingTransactor) DepositETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "depositETH", from, nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address from, uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingSession) DepositETH(from common.Address, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositETH(&_Csaccounting.TransactOpts, from, nodeOperatorId)
}

// DepositETH is a paid mutator transaction binding the contract method 0x2e599054.
//
// Solidity: function depositETH(address from, uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingTransactorSession) DepositETH(from common.Address, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositETH(&_Csaccounting.TransactOpts, from, nodeOperatorId)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0x5358fbda.
//
// Solidity: function depositETH(uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingTransactor) DepositETH0(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "depositETH0", nodeOperatorId)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0x5358fbda.
//
// Solidity: function depositETH(uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingSession) DepositETH0(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositETH0(&_Csaccounting.TransactOpts, nodeOperatorId)
}

// DepositETH0 is a paid mutator transaction binding the contract method 0x5358fbda.
//
// Solidity: function depositETH(uint256 nodeOperatorId) payable returns()
func (_Csaccounting *CsaccountingTransactorSession) DepositETH0(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositETH0(&_Csaccounting.TransactOpts, nodeOperatorId)
}

// DepositStETH is a paid mutator transaction binding the contract method 0x4c7ed3d2.
//
// Solidity: function depositStETH(address from, uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingTransactor) DepositStETH(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "depositStETH", from, nodeOperatorId, stETHAmount, permit)
}

// DepositStETH is a paid mutator transaction binding the contract method 0x4c7ed3d2.
//
// Solidity: function depositStETH(address from, uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingSession) DepositStETH(from common.Address, nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositStETH(&_Csaccounting.TransactOpts, from, nodeOperatorId, stETHAmount, permit)
}

// DepositStETH is a paid mutator transaction binding the contract method 0x4c7ed3d2.
//
// Solidity: function depositStETH(address from, uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingTransactorSession) DepositStETH(from common.Address, nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositStETH(&_Csaccounting.TransactOpts, from, nodeOperatorId, stETHAmount, permit)
}

// DepositStETH0 is a paid mutator transaction binding the contract method 0xe1aa105d.
//
// Solidity: function depositStETH(uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingTransactor) DepositStETH0(opts *bind.TransactOpts, nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "depositStETH0", nodeOperatorId, stETHAmount, permit)
}

// DepositStETH0 is a paid mutator transaction binding the contract method 0xe1aa105d.
//
// Solidity: function depositStETH(uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingSession) DepositStETH0(nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositStETH0(&_Csaccounting.TransactOpts, nodeOperatorId, stETHAmount, permit)
}

// DepositStETH0 is a paid mutator transaction binding the contract method 0xe1aa105d.
//
// Solidity: function depositStETH(uint256 nodeOperatorId, uint256 stETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingTransactorSession) DepositStETH0(nodeOperatorId *big.Int, stETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositStETH0(&_Csaccounting.TransactOpts, nodeOperatorId, stETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0x3f214bb2.
//
// Solidity: function depositWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingTransactor) DepositWstETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "depositWstETH", nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0x3f214bb2.
//
// Solidity: function depositWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingSession) DepositWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositWstETH(&_Csaccounting.TransactOpts, nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH is a paid mutator transaction binding the contract method 0x3f214bb2.
//
// Solidity: function depositWstETH(uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingTransactorSession) DepositWstETH(nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositWstETH(&_Csaccounting.TransactOpts, nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH0 is a paid mutator transaction binding the contract method 0xf7966efe.
//
// Solidity: function depositWstETH(address from, uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingTransactor) DepositWstETH0(opts *bind.TransactOpts, from common.Address, nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "depositWstETH0", from, nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH0 is a paid mutator transaction binding the contract method 0xf7966efe.
//
// Solidity: function depositWstETH(address from, uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingSession) DepositWstETH0(from common.Address, nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositWstETH0(&_Csaccounting.TransactOpts, from, nodeOperatorId, wstETHAmount, permit)
}

// DepositWstETH0 is a paid mutator transaction binding the contract method 0xf7966efe.
//
// Solidity: function depositWstETH(address from, uint256 nodeOperatorId, uint256 wstETHAmount, (uint256,uint256,uint8,bytes32,bytes32) permit) returns()
func (_Csaccounting *CsaccountingTransactorSession) DepositWstETH0(from common.Address, nodeOperatorId *big.Int, wstETHAmount *big.Int, permit ICSAccountingPermitInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.DepositWstETH0(&_Csaccounting.TransactOpts, from, nodeOperatorId, wstETHAmount, permit)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x0b3d765a.
//
// Solidity: function finalizeUpgradeV2((uint256,uint256)[][] bondCurvesInputs) returns()
func (_Csaccounting *CsaccountingTransactor) FinalizeUpgradeV2(opts *bind.TransactOpts, bondCurvesInputs [][]ICSBondCurveBondCurveIntervalInput) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "finalizeUpgradeV2", bondCurvesInputs)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x0b3d765a.
//
// Solidity: function finalizeUpgradeV2((uint256,uint256)[][] bondCurvesInputs) returns()
func (_Csaccounting *CsaccountingSession) FinalizeUpgradeV2(bondCurvesInputs [][]ICSBondCurveBondCurveIntervalInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.FinalizeUpgradeV2(&_Csaccounting.TransactOpts, bondCurvesInputs)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0x0b3d765a.
//
// Solidity: function finalizeUpgradeV2((uint256,uint256)[][] bondCurvesInputs) returns()
func (_Csaccounting *CsaccountingTransactorSession) FinalizeUpgradeV2(bondCurvesInputs [][]ICSBondCurveBondCurveIntervalInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.FinalizeUpgradeV2(&_Csaccounting.TransactOpts, bondCurvesInputs)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csaccounting *CsaccountingTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csaccounting *CsaccountingSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.GrantRole(&_Csaccounting.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Csaccounting *CsaccountingTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.GrantRole(&_Csaccounting.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2896ff4.
//
// Solidity: function initialize((uint256,uint256)[] bondCurve, address admin, uint256 bondLockPeriod, address _chargePenaltyRecipient) returns()
func (_Csaccounting *CsaccountingTransactor) Initialize(opts *bind.TransactOpts, bondCurve []ICSBondCurveBondCurveIntervalInput, admin common.Address, bondLockPeriod *big.Int, _chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "initialize", bondCurve, admin, bondLockPeriod, _chargePenaltyRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2896ff4.
//
// Solidity: function initialize((uint256,uint256)[] bondCurve, address admin, uint256 bondLockPeriod, address _chargePenaltyRecipient) returns()
func (_Csaccounting *CsaccountingSession) Initialize(bondCurve []ICSBondCurveBondCurveIntervalInput, admin common.Address, bondLockPeriod *big.Int, _chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.Initialize(&_Csaccounting.TransactOpts, bondCurve, admin, bondLockPeriod, _chargePenaltyRecipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xd2896ff4.
//
// Solidity: function initialize((uint256,uint256)[] bondCurve, address admin, uint256 bondLockPeriod, address _chargePenaltyRecipient) returns()
func (_Csaccounting *CsaccountingTransactorSession) Initialize(bondCurve []ICSBondCurveBondCurveIntervalInput, admin common.Address, bondLockPeriod *big.Int, _chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.Initialize(&_Csaccounting.TransactOpts, bondCurve, admin, bondLockPeriod, _chargePenaltyRecipient)
}

// LockBondETH is a paid mutator transaction binding the contract method 0xdcab7f83.
//
// Solidity: function lockBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactor) LockBondETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "lockBondETH", nodeOperatorId, amount)
}

// LockBondETH is a paid mutator transaction binding the contract method 0xdcab7f83.
//
// Solidity: function lockBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingSession) LockBondETH(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.LockBondETH(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// LockBondETH is a paid mutator transaction binding the contract method 0xdcab7f83.
//
// Solidity: function lockBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactorSession) LockBondETH(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.LockBondETH(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Csaccounting *CsaccountingTransactor) PauseFor(opts *bind.TransactOpts, duration *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "pauseFor", duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Csaccounting *CsaccountingSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.PauseFor(&_Csaccounting.TransactOpts, duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 duration) returns()
func (_Csaccounting *CsaccountingTransactorSession) PauseFor(duration *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.PauseFor(&_Csaccounting.TransactOpts, duration)
}

// Penalize is a paid mutator transaction binding the contract method 0xe5220e3f.
//
// Solidity: function penalize(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactor) Penalize(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "penalize", nodeOperatorId, amount)
}

// Penalize is a paid mutator transaction binding the contract method 0xe5220e3f.
//
// Solidity: function penalize(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingSession) Penalize(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.Penalize(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// Penalize is a paid mutator transaction binding the contract method 0xe5220e3f.
//
// Solidity: function penalize(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactorSession) Penalize(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.Penalize(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// PullFeeRewards is a paid mutator transaction binding the contract method 0x9b4c6c27.
//
// Solidity: function pullFeeRewards(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingTransactor) PullFeeRewards(opts *bind.TransactOpts, nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "pullFeeRewards", nodeOperatorId, cumulativeFeeShares, rewardsProof)
}

// PullFeeRewards is a paid mutator transaction binding the contract method 0x9b4c6c27.
//
// Solidity: function pullFeeRewards(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingSession) PullFeeRewards(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.PullFeeRewards(&_Csaccounting.TransactOpts, nodeOperatorId, cumulativeFeeShares, rewardsProof)
}

// PullFeeRewards is a paid mutator transaction binding the contract method 0x9b4c6c27.
//
// Solidity: function pullFeeRewards(uint256 nodeOperatorId, uint256 cumulativeFeeShares, bytes32[] rewardsProof) returns()
func (_Csaccounting *CsaccountingTransactorSession) PullFeeRewards(nodeOperatorId *big.Int, cumulativeFeeShares *big.Int, rewardsProof [][32]byte) (*types.Transaction, error) {
	return _Csaccounting.Contract.PullFeeRewards(&_Csaccounting.TransactOpts, nodeOperatorId, cumulativeFeeShares, rewardsProof)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csaccounting *CsaccountingTransactor) RecoverERC1155(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "recoverERC1155", token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csaccounting *CsaccountingSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverERC1155(&_Csaccounting.TransactOpts, token, tokenId)
}

// RecoverERC1155 is a paid mutator transaction binding the contract method 0x5c654ad9.
//
// Solidity: function recoverERC1155(address token, uint256 tokenId) returns()
func (_Csaccounting *CsaccountingTransactorSession) RecoverERC1155(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverERC1155(&_Csaccounting.TransactOpts, token, tokenId)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactor) RecoverERC20(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "recoverERC20", token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csaccounting *CsaccountingSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverERC20(&_Csaccounting.TransactOpts, token, amount)
}

// RecoverERC20 is a paid mutator transaction binding the contract method 0x8980f11f.
//
// Solidity: function recoverERC20(address token, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactorSession) RecoverERC20(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverERC20(&_Csaccounting.TransactOpts, token, amount)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csaccounting *CsaccountingTransactor) RecoverERC721(opts *bind.TransactOpts, token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "recoverERC721", token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csaccounting *CsaccountingSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverERC721(&_Csaccounting.TransactOpts, token, tokenId)
}

// RecoverERC721 is a paid mutator transaction binding the contract method 0x819d4cc6.
//
// Solidity: function recoverERC721(address token, uint256 tokenId) returns()
func (_Csaccounting *CsaccountingTransactorSession) RecoverERC721(token common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverERC721(&_Csaccounting.TransactOpts, token, tokenId)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csaccounting *CsaccountingTransactor) RecoverEther(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "recoverEther")
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csaccounting *CsaccountingSession) RecoverEther() (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverEther(&_Csaccounting.TransactOpts)
}

// RecoverEther is a paid mutator transaction binding the contract method 0x52d8bfc2.
//
// Solidity: function recoverEther() returns()
func (_Csaccounting *CsaccountingTransactorSession) RecoverEther() (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverEther(&_Csaccounting.TransactOpts)
}

// RecoverStETHShares is a paid mutator transaction binding the contract method 0x5a73bdc8.
//
// Solidity: function recoverStETHShares() returns()
func (_Csaccounting *CsaccountingTransactor) RecoverStETHShares(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "recoverStETHShares")
}

// RecoverStETHShares is a paid mutator transaction binding the contract method 0x5a73bdc8.
//
// Solidity: function recoverStETHShares() returns()
func (_Csaccounting *CsaccountingSession) RecoverStETHShares() (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverStETHShares(&_Csaccounting.TransactOpts)
}

// RecoverStETHShares is a paid mutator transaction binding the contract method 0x5a73bdc8.
//
// Solidity: function recoverStETHShares() returns()
func (_Csaccounting *CsaccountingTransactorSession) RecoverStETHShares() (*types.Transaction, error) {
	return _Csaccounting.Contract.RecoverStETHShares(&_Csaccounting.TransactOpts)
}

// ReleaseLockedBondETH is a paid mutator transaction binding the contract method 0xd963ae55.
//
// Solidity: function releaseLockedBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactor) ReleaseLockedBondETH(opts *bind.TransactOpts, nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "releaseLockedBondETH", nodeOperatorId, amount)
}

// ReleaseLockedBondETH is a paid mutator transaction binding the contract method 0xd963ae55.
//
// Solidity: function releaseLockedBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingSession) ReleaseLockedBondETH(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.ReleaseLockedBondETH(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// ReleaseLockedBondETH is a paid mutator transaction binding the contract method 0xd963ae55.
//
// Solidity: function releaseLockedBondETH(uint256 nodeOperatorId, uint256 amount) returns()
func (_Csaccounting *CsaccountingTransactorSession) ReleaseLockedBondETH(nodeOperatorId *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.ReleaseLockedBondETH(&_Csaccounting.TransactOpts, nodeOperatorId, amount)
}

// RenewBurnerAllowance is a paid mutator transaction binding the contract method 0xf3efecc4.
//
// Solidity: function renewBurnerAllowance() returns()
func (_Csaccounting *CsaccountingTransactor) RenewBurnerAllowance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "renewBurnerAllowance")
}

// RenewBurnerAllowance is a paid mutator transaction binding the contract method 0xf3efecc4.
//
// Solidity: function renewBurnerAllowance() returns()
func (_Csaccounting *CsaccountingSession) RenewBurnerAllowance() (*types.Transaction, error) {
	return _Csaccounting.Contract.RenewBurnerAllowance(&_Csaccounting.TransactOpts)
}

// RenewBurnerAllowance is a paid mutator transaction binding the contract method 0xf3efecc4.
//
// Solidity: function renewBurnerAllowance() returns()
func (_Csaccounting *CsaccountingTransactorSession) RenewBurnerAllowance() (*types.Transaction, error) {
	return _Csaccounting.Contract.RenewBurnerAllowance(&_Csaccounting.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csaccounting *CsaccountingTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csaccounting *CsaccountingSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.RenounceRole(&_Csaccounting.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Csaccounting *CsaccountingTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.RenounceRole(&_Csaccounting.TransactOpts, role, callerConfirmation)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Csaccounting *CsaccountingTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Csaccounting *CsaccountingSession) Resume() (*types.Transaction, error) {
	return _Csaccounting.Contract.Resume(&_Csaccounting.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Csaccounting *CsaccountingTransactorSession) Resume() (*types.Transaction, error) {
	return _Csaccounting.Contract.Resume(&_Csaccounting.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csaccounting *CsaccountingTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csaccounting *CsaccountingSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.RevokeRole(&_Csaccounting.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Csaccounting *CsaccountingTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.RevokeRole(&_Csaccounting.TransactOpts, role, account)
}

// SetBondCurve is a paid mutator transaction binding the contract method 0xb2d03e4d.
//
// Solidity: function setBondCurve(uint256 nodeOperatorId, uint256 curveId) returns()
func (_Csaccounting *CsaccountingTransactor) SetBondCurve(opts *bind.TransactOpts, nodeOperatorId *big.Int, curveId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "setBondCurve", nodeOperatorId, curveId)
}

// SetBondCurve is a paid mutator transaction binding the contract method 0xb2d03e4d.
//
// Solidity: function setBondCurve(uint256 nodeOperatorId, uint256 curveId) returns()
func (_Csaccounting *CsaccountingSession) SetBondCurve(nodeOperatorId *big.Int, curveId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.SetBondCurve(&_Csaccounting.TransactOpts, nodeOperatorId, curveId)
}

// SetBondCurve is a paid mutator transaction binding the contract method 0xb2d03e4d.
//
// Solidity: function setBondCurve(uint256 nodeOperatorId, uint256 curveId) returns()
func (_Csaccounting *CsaccountingTransactorSession) SetBondCurve(nodeOperatorId *big.Int, curveId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.SetBondCurve(&_Csaccounting.TransactOpts, nodeOperatorId, curveId)
}

// SetBondLockPeriod is a paid mutator transaction binding the contract method 0x684b21c9.
//
// Solidity: function setBondLockPeriod(uint256 period) returns()
func (_Csaccounting *CsaccountingTransactor) SetBondLockPeriod(opts *bind.TransactOpts, period *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "setBondLockPeriod", period)
}

// SetBondLockPeriod is a paid mutator transaction binding the contract method 0x684b21c9.
//
// Solidity: function setBondLockPeriod(uint256 period) returns()
func (_Csaccounting *CsaccountingSession) SetBondLockPeriod(period *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.SetBondLockPeriod(&_Csaccounting.TransactOpts, period)
}

// SetBondLockPeriod is a paid mutator transaction binding the contract method 0x684b21c9.
//
// Solidity: function setBondLockPeriod(uint256 period) returns()
func (_Csaccounting *CsaccountingTransactorSession) SetBondLockPeriod(period *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.SetBondLockPeriod(&_Csaccounting.TransactOpts, period)
}

// SetChargePenaltyRecipient is a paid mutator transaction binding the contract method 0x433cd6c3.
//
// Solidity: function setChargePenaltyRecipient(address _chargePenaltyRecipient) returns()
func (_Csaccounting *CsaccountingTransactor) SetChargePenaltyRecipient(opts *bind.TransactOpts, _chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "setChargePenaltyRecipient", _chargePenaltyRecipient)
}

// SetChargePenaltyRecipient is a paid mutator transaction binding the contract method 0x433cd6c3.
//
// Solidity: function setChargePenaltyRecipient(address _chargePenaltyRecipient) returns()
func (_Csaccounting *CsaccountingSession) SetChargePenaltyRecipient(_chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.SetChargePenaltyRecipient(&_Csaccounting.TransactOpts, _chargePenaltyRecipient)
}

// SetChargePenaltyRecipient is a paid mutator transaction binding the contract method 0x433cd6c3.
//
// Solidity: function setChargePenaltyRecipient(address _chargePenaltyRecipient) returns()
func (_Csaccounting *CsaccountingTransactorSession) SetChargePenaltyRecipient(_chargePenaltyRecipient common.Address) (*types.Transaction, error) {
	return _Csaccounting.Contract.SetChargePenaltyRecipient(&_Csaccounting.TransactOpts, _chargePenaltyRecipient)
}

// SettleLockedBondETH is a paid mutator transaction binding the contract method 0x4bb22a72.
//
// Solidity: function settleLockedBondETH(uint256 nodeOperatorId) returns(bool applied)
func (_Csaccounting *CsaccountingTransactor) SettleLockedBondETH(opts *bind.TransactOpts, nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "settleLockedBondETH", nodeOperatorId)
}

// SettleLockedBondETH is a paid mutator transaction binding the contract method 0x4bb22a72.
//
// Solidity: function settleLockedBondETH(uint256 nodeOperatorId) returns(bool applied)
func (_Csaccounting *CsaccountingSession) SettleLockedBondETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.SettleLockedBondETH(&_Csaccounting.TransactOpts, nodeOperatorId)
}

// SettleLockedBondETH is a paid mutator transaction binding the contract method 0x4bb22a72.
//
// Solidity: function settleLockedBondETH(uint256 nodeOperatorId) returns(bool applied)
func (_Csaccounting *CsaccountingTransactorSession) SettleLockedBondETH(nodeOperatorId *big.Int) (*types.Transaction, error) {
	return _Csaccounting.Contract.SettleLockedBondETH(&_Csaccounting.TransactOpts, nodeOperatorId)
}

// UpdateBondCurve is a paid mutator transaction binding the contract method 0xef038a54.
//
// Solidity: function updateBondCurve(uint256 curveId, (uint256,uint256)[] bondCurve) returns()
func (_Csaccounting *CsaccountingTransactor) UpdateBondCurve(opts *bind.TransactOpts, curveId *big.Int, bondCurve []ICSBondCurveBondCurveIntervalInput) (*types.Transaction, error) {
	return _Csaccounting.contract.Transact(opts, "updateBondCurve", curveId, bondCurve)
}

// UpdateBondCurve is a paid mutator transaction binding the contract method 0xef038a54.
//
// Solidity: function updateBondCurve(uint256 curveId, (uint256,uint256)[] bondCurve) returns()
func (_Csaccounting *CsaccountingSession) UpdateBondCurve(curveId *big.Int, bondCurve []ICSBondCurveBondCurveIntervalInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.UpdateBondCurve(&_Csaccounting.TransactOpts, curveId, bondCurve)
}

// UpdateBondCurve is a paid mutator transaction binding the contract method 0xef038a54.
//
// Solidity: function updateBondCurve(uint256 curveId, (uint256,uint256)[] bondCurve) returns()
func (_Csaccounting *CsaccountingTransactorSession) UpdateBondCurve(curveId *big.Int, bondCurve []ICSBondCurveBondCurveIntervalInput) (*types.Transaction, error) {
	return _Csaccounting.Contract.UpdateBondCurve(&_Csaccounting.TransactOpts, curveId, bondCurve)
}

// CsaccountingBondBurnedIterator is returned from FilterBondBurned and is used to iterate over the raw logs and unpacked data for BondBurned events raised by the Csaccounting contract.
type CsaccountingBondBurnedIterator struct {
	Event *CsaccountingBondBurned // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondBurned)
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
		it.Event = new(CsaccountingBondBurned)
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
func (it *CsaccountingBondBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondBurned represents a BondBurned event raised by the Csaccounting contract.
type CsaccountingBondBurned struct {
	NodeOperatorId *big.Int
	AmountToBurn   *big.Int
	BurnedAmount   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondBurned is a free log retrieval operation binding the contract event 0x4da924ae7845fe96897faab524b536685b8bbc4d82fbb45c10d941e0f86ade0f.
//
// Solidity: event BondBurned(uint256 indexed nodeOperatorId, uint256 amountToBurn, uint256 burnedAmount)
func (_Csaccounting *CsaccountingFilterer) FilterBondBurned(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondBurnedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondBurned", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondBurnedIterator{contract: _Csaccounting.contract, event: "BondBurned", logs: logs, sub: sub}, nil
}

// WatchBondBurned is a free log subscription operation binding the contract event 0x4da924ae7845fe96897faab524b536685b8bbc4d82fbb45c10d941e0f86ade0f.
//
// Solidity: event BondBurned(uint256 indexed nodeOperatorId, uint256 amountToBurn, uint256 burnedAmount)
func (_Csaccounting *CsaccountingFilterer) WatchBondBurned(opts *bind.WatchOpts, sink chan<- *CsaccountingBondBurned, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondBurned", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondBurned)
				if err := _Csaccounting.contract.UnpackLog(event, "BondBurned", log); err != nil {
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
// Solidity: event BondBurned(uint256 indexed nodeOperatorId, uint256 amountToBurn, uint256 burnedAmount)
func (_Csaccounting *CsaccountingFilterer) ParseBondBurned(log types.Log) (*CsaccountingBondBurned, error) {
	event := new(CsaccountingBondBurned)
	if err := _Csaccounting.contract.UnpackLog(event, "BondBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondChargedIterator is returned from FilterBondCharged and is used to iterate over the raw logs and unpacked data for BondCharged events raised by the Csaccounting contract.
type CsaccountingBondChargedIterator struct {
	Event *CsaccountingBondCharged // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondChargedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondCharged)
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
		it.Event = new(CsaccountingBondCharged)
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
func (it *CsaccountingBondChargedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondChargedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondCharged represents a BondCharged event raised by the Csaccounting contract.
type CsaccountingBondCharged struct {
	NodeOperatorId *big.Int
	ToChargeAmount *big.Int
	ChargedAmount  *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondCharged is a free log retrieval operation binding the contract event 0x8615528474a7bb3a28d9971535d956b79242b8e8fcfb27f3e331270fff088afd.
//
// Solidity: event BondCharged(uint256 indexed nodeOperatorId, uint256 toChargeAmount, uint256 chargedAmount)
func (_Csaccounting *CsaccountingFilterer) FilterBondCharged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondChargedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondCharged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondChargedIterator{contract: _Csaccounting.contract, event: "BondCharged", logs: logs, sub: sub}, nil
}

// WatchBondCharged is a free log subscription operation binding the contract event 0x8615528474a7bb3a28d9971535d956b79242b8e8fcfb27f3e331270fff088afd.
//
// Solidity: event BondCharged(uint256 indexed nodeOperatorId, uint256 toChargeAmount, uint256 chargedAmount)
func (_Csaccounting *CsaccountingFilterer) WatchBondCharged(opts *bind.WatchOpts, sink chan<- *CsaccountingBondCharged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondCharged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondCharged)
				if err := _Csaccounting.contract.UnpackLog(event, "BondCharged", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseBondCharged(log types.Log) (*CsaccountingBondCharged, error) {
	event := new(CsaccountingBondCharged)
	if err := _Csaccounting.contract.UnpackLog(event, "BondCharged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondClaimedStETHIterator is returned from FilterBondClaimedStETH and is used to iterate over the raw logs and unpacked data for BondClaimedStETH events raised by the Csaccounting contract.
type CsaccountingBondClaimedStETHIterator struct {
	Event *CsaccountingBondClaimedStETH // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondClaimedStETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondClaimedStETH)
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
		it.Event = new(CsaccountingBondClaimedStETH)
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
func (it *CsaccountingBondClaimedStETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondClaimedStETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondClaimedStETH represents a BondClaimedStETH event raised by the Csaccounting contract.
type CsaccountingBondClaimedStETH struct {
	NodeOperatorId *big.Int
	To             common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondClaimedStETH is a free log retrieval operation binding the contract event 0x3e3a1398fe71575ed0c17a80cd9d46ad684c2c75c2fad7b0e7dde15e78ab22d3.
//
// Solidity: event BondClaimedStETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterBondClaimedStETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondClaimedStETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondClaimedStETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondClaimedStETHIterator{contract: _Csaccounting.contract, event: "BondClaimedStETH", logs: logs, sub: sub}, nil
}

// WatchBondClaimedStETH is a free log subscription operation binding the contract event 0x3e3a1398fe71575ed0c17a80cd9d46ad684c2c75c2fad7b0e7dde15e78ab22d3.
//
// Solidity: event BondClaimedStETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchBondClaimedStETH(opts *bind.WatchOpts, sink chan<- *CsaccountingBondClaimedStETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondClaimedStETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondClaimedStETH)
				if err := _Csaccounting.contract.UnpackLog(event, "BondClaimedStETH", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseBondClaimedStETH(log types.Log) (*CsaccountingBondClaimedStETH, error) {
	event := new(CsaccountingBondClaimedStETH)
	if err := _Csaccounting.contract.UnpackLog(event, "BondClaimedStETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondClaimedUnstETHIterator is returned from FilterBondClaimedUnstETH and is used to iterate over the raw logs and unpacked data for BondClaimedUnstETH events raised by the Csaccounting contract.
type CsaccountingBondClaimedUnstETHIterator struct {
	Event *CsaccountingBondClaimedUnstETH // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondClaimedUnstETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondClaimedUnstETH)
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
		it.Event = new(CsaccountingBondClaimedUnstETH)
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
func (it *CsaccountingBondClaimedUnstETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondClaimedUnstETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondClaimedUnstETH represents a BondClaimedUnstETH event raised by the Csaccounting contract.
type CsaccountingBondClaimedUnstETH struct {
	NodeOperatorId *big.Int
	To             common.Address
	Amount         *big.Int
	RequestId      *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondClaimedUnstETH is a free log retrieval operation binding the contract event 0x26673a9d018b21192d08ee14377b798f11b9e5b15ea1559c110265716b8985b5.
//
// Solidity: event BondClaimedUnstETH(uint256 indexed nodeOperatorId, address to, uint256 amount, uint256 requestId)
func (_Csaccounting *CsaccountingFilterer) FilterBondClaimedUnstETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondClaimedUnstETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondClaimedUnstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondClaimedUnstETHIterator{contract: _Csaccounting.contract, event: "BondClaimedUnstETH", logs: logs, sub: sub}, nil
}

// WatchBondClaimedUnstETH is a free log subscription operation binding the contract event 0x26673a9d018b21192d08ee14377b798f11b9e5b15ea1559c110265716b8985b5.
//
// Solidity: event BondClaimedUnstETH(uint256 indexed nodeOperatorId, address to, uint256 amount, uint256 requestId)
func (_Csaccounting *CsaccountingFilterer) WatchBondClaimedUnstETH(opts *bind.WatchOpts, sink chan<- *CsaccountingBondClaimedUnstETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondClaimedUnstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondClaimedUnstETH)
				if err := _Csaccounting.contract.UnpackLog(event, "BondClaimedUnstETH", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseBondClaimedUnstETH(log types.Log) (*CsaccountingBondClaimedUnstETH, error) {
	event := new(CsaccountingBondClaimedUnstETH)
	if err := _Csaccounting.contract.UnpackLog(event, "BondClaimedUnstETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondClaimedWstETHIterator is returned from FilterBondClaimedWstETH and is used to iterate over the raw logs and unpacked data for BondClaimedWstETH events raised by the Csaccounting contract.
type CsaccountingBondClaimedWstETHIterator struct {
	Event *CsaccountingBondClaimedWstETH // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondClaimedWstETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondClaimedWstETH)
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
		it.Event = new(CsaccountingBondClaimedWstETH)
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
func (it *CsaccountingBondClaimedWstETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondClaimedWstETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondClaimedWstETH represents a BondClaimedWstETH event raised by the Csaccounting contract.
type CsaccountingBondClaimedWstETH struct {
	NodeOperatorId *big.Int
	To             common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondClaimedWstETH is a free log retrieval operation binding the contract event 0xe6a8c06447e05a412e5e9581e088941f3994db3d8a9bfd3275b38d77acacafac.
//
// Solidity: event BondClaimedWstETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterBondClaimedWstETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondClaimedWstETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondClaimedWstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondClaimedWstETHIterator{contract: _Csaccounting.contract, event: "BondClaimedWstETH", logs: logs, sub: sub}, nil
}

// WatchBondClaimedWstETH is a free log subscription operation binding the contract event 0xe6a8c06447e05a412e5e9581e088941f3994db3d8a9bfd3275b38d77acacafac.
//
// Solidity: event BondClaimedWstETH(uint256 indexed nodeOperatorId, address to, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchBondClaimedWstETH(opts *bind.WatchOpts, sink chan<- *CsaccountingBondClaimedWstETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondClaimedWstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondClaimedWstETH)
				if err := _Csaccounting.contract.UnpackLog(event, "BondClaimedWstETH", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseBondClaimedWstETH(log types.Log) (*CsaccountingBondClaimedWstETH, error) {
	event := new(CsaccountingBondClaimedWstETH)
	if err := _Csaccounting.contract.UnpackLog(event, "BondClaimedWstETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondCurveAddedIterator is returned from FilterBondCurveAdded and is used to iterate over the raw logs and unpacked data for BondCurveAdded events raised by the Csaccounting contract.
type CsaccountingBondCurveAddedIterator struct {
	Event *CsaccountingBondCurveAdded // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondCurveAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondCurveAdded)
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
		it.Event = new(CsaccountingBondCurveAdded)
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
func (it *CsaccountingBondCurveAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondCurveAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondCurveAdded represents a BondCurveAdded event raised by the Csaccounting contract.
type CsaccountingBondCurveAdded struct {
	CurveId            *big.Int
	BondCurveIntervals []ICSBondCurveBondCurveIntervalInput
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBondCurveAdded is a free log retrieval operation binding the contract event 0x707691ca33c3fcf1738eeb4c10826bd3030b3687166d6de80eb5896067fd2159.
//
// Solidity: event BondCurveAdded(uint256 indexed curveId, (uint256,uint256)[] bondCurveIntervals)
func (_Csaccounting *CsaccountingFilterer) FilterBondCurveAdded(opts *bind.FilterOpts, curveId []*big.Int) (*CsaccountingBondCurveAddedIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondCurveAdded", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondCurveAddedIterator{contract: _Csaccounting.contract, event: "BondCurveAdded", logs: logs, sub: sub}, nil
}

// WatchBondCurveAdded is a free log subscription operation binding the contract event 0x707691ca33c3fcf1738eeb4c10826bd3030b3687166d6de80eb5896067fd2159.
//
// Solidity: event BondCurveAdded(uint256 indexed curveId, (uint256,uint256)[] bondCurveIntervals)
func (_Csaccounting *CsaccountingFilterer) WatchBondCurveAdded(opts *bind.WatchOpts, sink chan<- *CsaccountingBondCurveAdded, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondCurveAdded", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondCurveAdded)
				if err := _Csaccounting.contract.UnpackLog(event, "BondCurveAdded", log); err != nil {
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

// ParseBondCurveAdded is a log parse operation binding the contract event 0x707691ca33c3fcf1738eeb4c10826bd3030b3687166d6de80eb5896067fd2159.
//
// Solidity: event BondCurveAdded(uint256 indexed curveId, (uint256,uint256)[] bondCurveIntervals)
func (_Csaccounting *CsaccountingFilterer) ParseBondCurveAdded(log types.Log) (*CsaccountingBondCurveAdded, error) {
	event := new(CsaccountingBondCurveAdded)
	if err := _Csaccounting.contract.UnpackLog(event, "BondCurveAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondCurveSetIterator is returned from FilterBondCurveSet and is used to iterate over the raw logs and unpacked data for BondCurveSet events raised by the Csaccounting contract.
type CsaccountingBondCurveSetIterator struct {
	Event *CsaccountingBondCurveSet // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondCurveSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondCurveSet)
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
		it.Event = new(CsaccountingBondCurveSet)
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
func (it *CsaccountingBondCurveSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondCurveSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondCurveSet represents a BondCurveSet event raised by the Csaccounting contract.
type CsaccountingBondCurveSet struct {
	NodeOperatorId *big.Int
	CurveId        *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondCurveSet is a free log retrieval operation binding the contract event 0x4642db1736894887bc907d721f20af84d3e585a0a3cea90f41b78b2aa906541b.
//
// Solidity: event BondCurveSet(uint256 indexed nodeOperatorId, uint256 curveId)
func (_Csaccounting *CsaccountingFilterer) FilterBondCurveSet(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondCurveSetIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondCurveSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondCurveSetIterator{contract: _Csaccounting.contract, event: "BondCurveSet", logs: logs, sub: sub}, nil
}

// WatchBondCurveSet is a free log subscription operation binding the contract event 0x4642db1736894887bc907d721f20af84d3e585a0a3cea90f41b78b2aa906541b.
//
// Solidity: event BondCurveSet(uint256 indexed nodeOperatorId, uint256 curveId)
func (_Csaccounting *CsaccountingFilterer) WatchBondCurveSet(opts *bind.WatchOpts, sink chan<- *CsaccountingBondCurveSet, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondCurveSet", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondCurveSet)
				if err := _Csaccounting.contract.UnpackLog(event, "BondCurveSet", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseBondCurveSet(log types.Log) (*CsaccountingBondCurveSet, error) {
	event := new(CsaccountingBondCurveSet)
	if err := _Csaccounting.contract.UnpackLog(event, "BondCurveSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondCurveUpdatedIterator is returned from FilterBondCurveUpdated and is used to iterate over the raw logs and unpacked data for BondCurveUpdated events raised by the Csaccounting contract.
type CsaccountingBondCurveUpdatedIterator struct {
	Event *CsaccountingBondCurveUpdated // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondCurveUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondCurveUpdated)
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
		it.Event = new(CsaccountingBondCurveUpdated)
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
func (it *CsaccountingBondCurveUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondCurveUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondCurveUpdated represents a BondCurveUpdated event raised by the Csaccounting contract.
type CsaccountingBondCurveUpdated struct {
	CurveId            *big.Int
	BondCurveIntervals []ICSBondCurveBondCurveIntervalInput
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterBondCurveUpdated is a free log retrieval operation binding the contract event 0x77c7f59d9ea0a6ee0417e777c399834e7ce0647a7ece2b12f4dbff0a6a1980c8.
//
// Solidity: event BondCurveUpdated(uint256 indexed curveId, (uint256,uint256)[] bondCurveIntervals)
func (_Csaccounting *CsaccountingFilterer) FilterBondCurveUpdated(opts *bind.FilterOpts, curveId []*big.Int) (*CsaccountingBondCurveUpdatedIterator, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondCurveUpdated", curveIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondCurveUpdatedIterator{contract: _Csaccounting.contract, event: "BondCurveUpdated", logs: logs, sub: sub}, nil
}

// WatchBondCurveUpdated is a free log subscription operation binding the contract event 0x77c7f59d9ea0a6ee0417e777c399834e7ce0647a7ece2b12f4dbff0a6a1980c8.
//
// Solidity: event BondCurveUpdated(uint256 indexed curveId, (uint256,uint256)[] bondCurveIntervals)
func (_Csaccounting *CsaccountingFilterer) WatchBondCurveUpdated(opts *bind.WatchOpts, sink chan<- *CsaccountingBondCurveUpdated, curveId []*big.Int) (event.Subscription, error) {

	var curveIdRule []interface{}
	for _, curveIdItem := range curveId {
		curveIdRule = append(curveIdRule, curveIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondCurveUpdated", curveIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondCurveUpdated)
				if err := _Csaccounting.contract.UnpackLog(event, "BondCurveUpdated", log); err != nil {
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

// ParseBondCurveUpdated is a log parse operation binding the contract event 0x77c7f59d9ea0a6ee0417e777c399834e7ce0647a7ece2b12f4dbff0a6a1980c8.
//
// Solidity: event BondCurveUpdated(uint256 indexed curveId, (uint256,uint256)[] bondCurveIntervals)
func (_Csaccounting *CsaccountingFilterer) ParseBondCurveUpdated(log types.Log) (*CsaccountingBondCurveUpdated, error) {
	event := new(CsaccountingBondCurveUpdated)
	if err := _Csaccounting.contract.UnpackLog(event, "BondCurveUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondDepositedETHIterator is returned from FilterBondDepositedETH and is used to iterate over the raw logs and unpacked data for BondDepositedETH events raised by the Csaccounting contract.
type CsaccountingBondDepositedETHIterator struct {
	Event *CsaccountingBondDepositedETH // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondDepositedETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondDepositedETH)
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
		it.Event = new(CsaccountingBondDepositedETH)
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
func (it *CsaccountingBondDepositedETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondDepositedETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondDepositedETH represents a BondDepositedETH event raised by the Csaccounting contract.
type CsaccountingBondDepositedETH struct {
	NodeOperatorId *big.Int
	From           common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondDepositedETH is a free log retrieval operation binding the contract event 0x16ec5116295424dec7fd52c87d9971a963ea7f59f741ad9ad468f0312055dc49.
//
// Solidity: event BondDepositedETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterBondDepositedETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondDepositedETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondDepositedETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondDepositedETHIterator{contract: _Csaccounting.contract, event: "BondDepositedETH", logs: logs, sub: sub}, nil
}

// WatchBondDepositedETH is a free log subscription operation binding the contract event 0x16ec5116295424dec7fd52c87d9971a963ea7f59f741ad9ad468f0312055dc49.
//
// Solidity: event BondDepositedETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchBondDepositedETH(opts *bind.WatchOpts, sink chan<- *CsaccountingBondDepositedETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondDepositedETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondDepositedETH)
				if err := _Csaccounting.contract.UnpackLog(event, "BondDepositedETH", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseBondDepositedETH(log types.Log) (*CsaccountingBondDepositedETH, error) {
	event := new(CsaccountingBondDepositedETH)
	if err := _Csaccounting.contract.UnpackLog(event, "BondDepositedETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondDepositedStETHIterator is returned from FilterBondDepositedStETH and is used to iterate over the raw logs and unpacked data for BondDepositedStETH events raised by the Csaccounting contract.
type CsaccountingBondDepositedStETHIterator struct {
	Event *CsaccountingBondDepositedStETH // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondDepositedStETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondDepositedStETH)
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
		it.Event = new(CsaccountingBondDepositedStETH)
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
func (it *CsaccountingBondDepositedStETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondDepositedStETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondDepositedStETH represents a BondDepositedStETH event raised by the Csaccounting contract.
type CsaccountingBondDepositedStETH struct {
	NodeOperatorId *big.Int
	From           common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondDepositedStETH is a free log retrieval operation binding the contract event 0xee31ebba29fd5471227e42fd8ca621a892d689901892cb8febb03fe802c3214b.
//
// Solidity: event BondDepositedStETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterBondDepositedStETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondDepositedStETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondDepositedStETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondDepositedStETHIterator{contract: _Csaccounting.contract, event: "BondDepositedStETH", logs: logs, sub: sub}, nil
}

// WatchBondDepositedStETH is a free log subscription operation binding the contract event 0xee31ebba29fd5471227e42fd8ca621a892d689901892cb8febb03fe802c3214b.
//
// Solidity: event BondDepositedStETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchBondDepositedStETH(opts *bind.WatchOpts, sink chan<- *CsaccountingBondDepositedStETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondDepositedStETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondDepositedStETH)
				if err := _Csaccounting.contract.UnpackLog(event, "BondDepositedStETH", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseBondDepositedStETH(log types.Log) (*CsaccountingBondDepositedStETH, error) {
	event := new(CsaccountingBondDepositedStETH)
	if err := _Csaccounting.contract.UnpackLog(event, "BondDepositedStETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondDepositedWstETHIterator is returned from FilterBondDepositedWstETH and is used to iterate over the raw logs and unpacked data for BondDepositedWstETH events raised by the Csaccounting contract.
type CsaccountingBondDepositedWstETHIterator struct {
	Event *CsaccountingBondDepositedWstETH // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondDepositedWstETHIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondDepositedWstETH)
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
		it.Event = new(CsaccountingBondDepositedWstETH)
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
func (it *CsaccountingBondDepositedWstETHIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondDepositedWstETHIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondDepositedWstETH represents a BondDepositedWstETH event raised by the Csaccounting contract.
type CsaccountingBondDepositedWstETH struct {
	NodeOperatorId *big.Int
	From           common.Address
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondDepositedWstETH is a free log retrieval operation binding the contract event 0x6576bbc9c5b478bf9717dc3d2bcb485e5ff0727df77c72558727597f3609d3f1.
//
// Solidity: event BondDepositedWstETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterBondDepositedWstETH(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondDepositedWstETHIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondDepositedWstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondDepositedWstETHIterator{contract: _Csaccounting.contract, event: "BondDepositedWstETH", logs: logs, sub: sub}, nil
}

// WatchBondDepositedWstETH is a free log subscription operation binding the contract event 0x6576bbc9c5b478bf9717dc3d2bcb485e5ff0727df77c72558727597f3609d3f1.
//
// Solidity: event BondDepositedWstETH(uint256 indexed nodeOperatorId, address from, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchBondDepositedWstETH(opts *bind.WatchOpts, sink chan<- *CsaccountingBondDepositedWstETH, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondDepositedWstETH", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondDepositedWstETH)
				if err := _Csaccounting.contract.UnpackLog(event, "BondDepositedWstETH", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseBondDepositedWstETH(log types.Log) (*CsaccountingBondDepositedWstETH, error) {
	event := new(CsaccountingBondDepositedWstETH)
	if err := _Csaccounting.contract.UnpackLog(event, "BondDepositedWstETH", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondLockChangedIterator is returned from FilterBondLockChanged and is used to iterate over the raw logs and unpacked data for BondLockChanged events raised by the Csaccounting contract.
type CsaccountingBondLockChangedIterator struct {
	Event *CsaccountingBondLockChanged // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondLockChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondLockChanged)
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
		it.Event = new(CsaccountingBondLockChanged)
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
func (it *CsaccountingBondLockChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondLockChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondLockChanged represents a BondLockChanged event raised by the Csaccounting contract.
type CsaccountingBondLockChanged struct {
	NodeOperatorId *big.Int
	NewAmount      *big.Int
	Until          *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondLockChanged is a free log retrieval operation binding the contract event 0x69a153d448f54b17f05cf3b268a2efab87c94a4727d108c4ca4aa3e5d65113de.
//
// Solidity: event BondLockChanged(uint256 indexed nodeOperatorId, uint256 newAmount, uint256 until)
func (_Csaccounting *CsaccountingFilterer) FilterBondLockChanged(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondLockChangedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondLockChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondLockChangedIterator{contract: _Csaccounting.contract, event: "BondLockChanged", logs: logs, sub: sub}, nil
}

// WatchBondLockChanged is a free log subscription operation binding the contract event 0x69a153d448f54b17f05cf3b268a2efab87c94a4727d108c4ca4aa3e5d65113de.
//
// Solidity: event BondLockChanged(uint256 indexed nodeOperatorId, uint256 newAmount, uint256 until)
func (_Csaccounting *CsaccountingFilterer) WatchBondLockChanged(opts *bind.WatchOpts, sink chan<- *CsaccountingBondLockChanged, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondLockChanged", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondLockChanged)
				if err := _Csaccounting.contract.UnpackLog(event, "BondLockChanged", log); err != nil {
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
// Solidity: event BondLockChanged(uint256 indexed nodeOperatorId, uint256 newAmount, uint256 until)
func (_Csaccounting *CsaccountingFilterer) ParseBondLockChanged(log types.Log) (*CsaccountingBondLockChanged, error) {
	event := new(CsaccountingBondLockChanged)
	if err := _Csaccounting.contract.UnpackLog(event, "BondLockChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondLockCompensatedIterator is returned from FilterBondLockCompensated and is used to iterate over the raw logs and unpacked data for BondLockCompensated events raised by the Csaccounting contract.
type CsaccountingBondLockCompensatedIterator struct {
	Event *CsaccountingBondLockCompensated // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondLockCompensatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondLockCompensated)
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
		it.Event = new(CsaccountingBondLockCompensated)
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
func (it *CsaccountingBondLockCompensatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondLockCompensatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondLockCompensated represents a BondLockCompensated event raised by the Csaccounting contract.
type CsaccountingBondLockCompensated struct {
	NodeOperatorId *big.Int
	Amount         *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondLockCompensated is a free log retrieval operation binding the contract event 0xb6ee6e3aae6776519627b46786a622b642c38cabfe4c97cb34054fd63fc11a23.
//
// Solidity: event BondLockCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterBondLockCompensated(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondLockCompensatedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondLockCompensated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondLockCompensatedIterator{contract: _Csaccounting.contract, event: "BondLockCompensated", logs: logs, sub: sub}, nil
}

// WatchBondLockCompensated is a free log subscription operation binding the contract event 0xb6ee6e3aae6776519627b46786a622b642c38cabfe4c97cb34054fd63fc11a23.
//
// Solidity: event BondLockCompensated(uint256 indexed nodeOperatorId, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchBondLockCompensated(opts *bind.WatchOpts, sink chan<- *CsaccountingBondLockCompensated, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondLockCompensated", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondLockCompensated)
				if err := _Csaccounting.contract.UnpackLog(event, "BondLockCompensated", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseBondLockCompensated(log types.Log) (*CsaccountingBondLockCompensated, error) {
	event := new(CsaccountingBondLockCompensated)
	if err := _Csaccounting.contract.UnpackLog(event, "BondLockCompensated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondLockPeriodChangedIterator is returned from FilterBondLockPeriodChanged and is used to iterate over the raw logs and unpacked data for BondLockPeriodChanged events raised by the Csaccounting contract.
type CsaccountingBondLockPeriodChangedIterator struct {
	Event *CsaccountingBondLockPeriodChanged // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondLockPeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondLockPeriodChanged)
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
		it.Event = new(CsaccountingBondLockPeriodChanged)
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
func (it *CsaccountingBondLockPeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondLockPeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondLockPeriodChanged represents a BondLockPeriodChanged event raised by the Csaccounting contract.
type CsaccountingBondLockPeriodChanged struct {
	Period *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBondLockPeriodChanged is a free log retrieval operation binding the contract event 0xd117ae9105bfc4a5acf683370984ce7aea9498aa2849fc0851e0b012552b3103.
//
// Solidity: event BondLockPeriodChanged(uint256 period)
func (_Csaccounting *CsaccountingFilterer) FilterBondLockPeriodChanged(opts *bind.FilterOpts) (*CsaccountingBondLockPeriodChangedIterator, error) {

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondLockPeriodChanged")
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondLockPeriodChangedIterator{contract: _Csaccounting.contract, event: "BondLockPeriodChanged", logs: logs, sub: sub}, nil
}

// WatchBondLockPeriodChanged is a free log subscription operation binding the contract event 0xd117ae9105bfc4a5acf683370984ce7aea9498aa2849fc0851e0b012552b3103.
//
// Solidity: event BondLockPeriodChanged(uint256 period)
func (_Csaccounting *CsaccountingFilterer) WatchBondLockPeriodChanged(opts *bind.WatchOpts, sink chan<- *CsaccountingBondLockPeriodChanged) (event.Subscription, error) {

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondLockPeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondLockPeriodChanged)
				if err := _Csaccounting.contract.UnpackLog(event, "BondLockPeriodChanged", log); err != nil {
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

// ParseBondLockPeriodChanged is a log parse operation binding the contract event 0xd117ae9105bfc4a5acf683370984ce7aea9498aa2849fc0851e0b012552b3103.
//
// Solidity: event BondLockPeriodChanged(uint256 period)
func (_Csaccounting *CsaccountingFilterer) ParseBondLockPeriodChanged(log types.Log) (*CsaccountingBondLockPeriodChanged, error) {
	event := new(CsaccountingBondLockPeriodChanged)
	if err := _Csaccounting.contract.UnpackLog(event, "BondLockPeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingBondLockRemovedIterator is returned from FilterBondLockRemoved and is used to iterate over the raw logs and unpacked data for BondLockRemoved events raised by the Csaccounting contract.
type CsaccountingBondLockRemovedIterator struct {
	Event *CsaccountingBondLockRemoved // Event containing the contract specifics and raw log

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
func (it *CsaccountingBondLockRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingBondLockRemoved)
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
		it.Event = new(CsaccountingBondLockRemoved)
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
func (it *CsaccountingBondLockRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingBondLockRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingBondLockRemoved represents a BondLockRemoved event raised by the Csaccounting contract.
type CsaccountingBondLockRemoved struct {
	NodeOperatorId *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondLockRemoved is a free log retrieval operation binding the contract event 0x844ae6b00e8a437dcdde1a634feab3273e08bb5c274a4be3b195b308ae0ba20a.
//
// Solidity: event BondLockRemoved(uint256 indexed nodeOperatorId)
func (_Csaccounting *CsaccountingFilterer) FilterBondLockRemoved(opts *bind.FilterOpts, nodeOperatorId []*big.Int) (*CsaccountingBondLockRemovedIterator, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "BondLockRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingBondLockRemovedIterator{contract: _Csaccounting.contract, event: "BondLockRemoved", logs: logs, sub: sub}, nil
}

// WatchBondLockRemoved is a free log subscription operation binding the contract event 0x844ae6b00e8a437dcdde1a634feab3273e08bb5c274a4be3b195b308ae0ba20a.
//
// Solidity: event BondLockRemoved(uint256 indexed nodeOperatorId)
func (_Csaccounting *CsaccountingFilterer) WatchBondLockRemoved(opts *bind.WatchOpts, sink chan<- *CsaccountingBondLockRemoved, nodeOperatorId []*big.Int) (event.Subscription, error) {

	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "BondLockRemoved", nodeOperatorIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingBondLockRemoved)
				if err := _Csaccounting.contract.UnpackLog(event, "BondLockRemoved", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseBondLockRemoved(log types.Log) (*CsaccountingBondLockRemoved, error) {
	event := new(CsaccountingBondLockRemoved)
	if err := _Csaccounting.contract.UnpackLog(event, "BondLockRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingChargePenaltyRecipientSetIterator is returned from FilterChargePenaltyRecipientSet and is used to iterate over the raw logs and unpacked data for ChargePenaltyRecipientSet events raised by the Csaccounting contract.
type CsaccountingChargePenaltyRecipientSetIterator struct {
	Event *CsaccountingChargePenaltyRecipientSet // Event containing the contract specifics and raw log

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
func (it *CsaccountingChargePenaltyRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingChargePenaltyRecipientSet)
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
		it.Event = new(CsaccountingChargePenaltyRecipientSet)
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
func (it *CsaccountingChargePenaltyRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingChargePenaltyRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingChargePenaltyRecipientSet represents a ChargePenaltyRecipientSet event raised by the Csaccounting contract.
type CsaccountingChargePenaltyRecipientSet struct {
	ChargePenaltyRecipient common.Address
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterChargePenaltyRecipientSet is a free log retrieval operation binding the contract event 0x4beaaee83871b066b675515d6a53567e76411f60266703cef934a01905a4d832.
//
// Solidity: event ChargePenaltyRecipientSet(address chargePenaltyRecipient)
func (_Csaccounting *CsaccountingFilterer) FilterChargePenaltyRecipientSet(opts *bind.FilterOpts) (*CsaccountingChargePenaltyRecipientSetIterator, error) {

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "ChargePenaltyRecipientSet")
	if err != nil {
		return nil, err
	}
	return &CsaccountingChargePenaltyRecipientSetIterator{contract: _Csaccounting.contract, event: "ChargePenaltyRecipientSet", logs: logs, sub: sub}, nil
}

// WatchChargePenaltyRecipientSet is a free log subscription operation binding the contract event 0x4beaaee83871b066b675515d6a53567e76411f60266703cef934a01905a4d832.
//
// Solidity: event ChargePenaltyRecipientSet(address chargePenaltyRecipient)
func (_Csaccounting *CsaccountingFilterer) WatchChargePenaltyRecipientSet(opts *bind.WatchOpts, sink chan<- *CsaccountingChargePenaltyRecipientSet) (event.Subscription, error) {

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "ChargePenaltyRecipientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingChargePenaltyRecipientSet)
				if err := _Csaccounting.contract.UnpackLog(event, "ChargePenaltyRecipientSet", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseChargePenaltyRecipientSet(log types.Log) (*CsaccountingChargePenaltyRecipientSet, error) {
	event := new(CsaccountingChargePenaltyRecipientSet)
	if err := _Csaccounting.contract.UnpackLog(event, "ChargePenaltyRecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingERC1155RecoveredIterator is returned from FilterERC1155Recovered and is used to iterate over the raw logs and unpacked data for ERC1155Recovered events raised by the Csaccounting contract.
type CsaccountingERC1155RecoveredIterator struct {
	Event *CsaccountingERC1155Recovered // Event containing the contract specifics and raw log

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
func (it *CsaccountingERC1155RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingERC1155Recovered)
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
		it.Event = new(CsaccountingERC1155Recovered)
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
func (it *CsaccountingERC1155RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingERC1155RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingERC1155Recovered represents a ERC1155Recovered event raised by the Csaccounting contract.
type CsaccountingERC1155Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC1155Recovered is a free log retrieval operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterERC1155Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsaccountingERC1155RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingERC1155RecoveredIterator{contract: _Csaccounting.contract, event: "ERC1155Recovered", logs: logs, sub: sub}, nil
}

// WatchERC1155Recovered is a free log subscription operation binding the contract event 0x5cf02e753b3eb0f4bee4460a72817d8e5e3c75cd4d65c1d0b06dca88b8032936.
//
// Solidity: event ERC1155Recovered(address indexed token, uint256 tokenId, address indexed recipient, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchERC1155Recovered(opts *bind.WatchOpts, sink chan<- *CsaccountingERC1155Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "ERC1155Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingERC1155Recovered)
				if err := _Csaccounting.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseERC1155Recovered(log types.Log) (*CsaccountingERC1155Recovered, error) {
	event := new(CsaccountingERC1155Recovered)
	if err := _Csaccounting.contract.UnpackLog(event, "ERC1155Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingERC20RecoveredIterator is returned from FilterERC20Recovered and is used to iterate over the raw logs and unpacked data for ERC20Recovered events raised by the Csaccounting contract.
type CsaccountingERC20RecoveredIterator struct {
	Event *CsaccountingERC20Recovered // Event containing the contract specifics and raw log

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
func (it *CsaccountingERC20RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingERC20Recovered)
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
		it.Event = new(CsaccountingERC20Recovered)
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
func (it *CsaccountingERC20RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingERC20RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingERC20Recovered represents a ERC20Recovered event raised by the Csaccounting contract.
type CsaccountingERC20Recovered struct {
	Token     common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC20Recovered is a free log retrieval operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterERC20Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsaccountingERC20RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingERC20RecoveredIterator{contract: _Csaccounting.contract, event: "ERC20Recovered", logs: logs, sub: sub}, nil
}

// WatchERC20Recovered is a free log subscription operation binding the contract event 0xaca8fb252cde442184e5f10e0f2e6e4029e8cd7717cae63559079610702436aa.
//
// Solidity: event ERC20Recovered(address indexed token, address indexed recipient, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchERC20Recovered(opts *bind.WatchOpts, sink chan<- *CsaccountingERC20Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "ERC20Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingERC20Recovered)
				if err := _Csaccounting.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseERC20Recovered(log types.Log) (*CsaccountingERC20Recovered, error) {
	event := new(CsaccountingERC20Recovered)
	if err := _Csaccounting.contract.UnpackLog(event, "ERC20Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingERC721RecoveredIterator is returned from FilterERC721Recovered and is used to iterate over the raw logs and unpacked data for ERC721Recovered events raised by the Csaccounting contract.
type CsaccountingERC721RecoveredIterator struct {
	Event *CsaccountingERC721Recovered // Event containing the contract specifics and raw log

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
func (it *CsaccountingERC721RecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingERC721Recovered)
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
		it.Event = new(CsaccountingERC721Recovered)
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
func (it *CsaccountingERC721RecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingERC721RecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingERC721Recovered represents a ERC721Recovered event raised by the Csaccounting contract.
type CsaccountingERC721Recovered struct {
	Token     common.Address
	TokenId   *big.Int
	Recipient common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterERC721Recovered is a free log retrieval operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csaccounting *CsaccountingFilterer) FilterERC721Recovered(opts *bind.FilterOpts, token []common.Address, recipient []common.Address) (*CsaccountingERC721RecoveredIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingERC721RecoveredIterator{contract: _Csaccounting.contract, event: "ERC721Recovered", logs: logs, sub: sub}, nil
}

// WatchERC721Recovered is a free log subscription operation binding the contract event 0x8166bf75d2ff2fa3c8f3c44410540bf42e9a5359b48409e8d660291dc9f788c8.
//
// Solidity: event ERC721Recovered(address indexed token, uint256 tokenId, address indexed recipient)
func (_Csaccounting *CsaccountingFilterer) WatchERC721Recovered(opts *bind.WatchOpts, sink chan<- *CsaccountingERC721Recovered, token []common.Address, recipient []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "ERC721Recovered", tokenRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingERC721Recovered)
				if err := _Csaccounting.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseERC721Recovered(log types.Log) (*CsaccountingERC721Recovered, error) {
	event := new(CsaccountingERC721Recovered)
	if err := _Csaccounting.contract.UnpackLog(event, "ERC721Recovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingEtherRecoveredIterator is returned from FilterEtherRecovered and is used to iterate over the raw logs and unpacked data for EtherRecovered events raised by the Csaccounting contract.
type CsaccountingEtherRecoveredIterator struct {
	Event *CsaccountingEtherRecovered // Event containing the contract specifics and raw log

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
func (it *CsaccountingEtherRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingEtherRecovered)
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
		it.Event = new(CsaccountingEtherRecovered)
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
func (it *CsaccountingEtherRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingEtherRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingEtherRecovered represents a EtherRecovered event raised by the Csaccounting contract.
type CsaccountingEtherRecovered struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEtherRecovered is a free log retrieval operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) FilterEtherRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CsaccountingEtherRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingEtherRecoveredIterator{contract: _Csaccounting.contract, event: "EtherRecovered", logs: logs, sub: sub}, nil
}

// WatchEtherRecovered is a free log subscription operation binding the contract event 0x8e274e42262a7f013b700b35c2b4629ccce1702f8fe83f8dfb7eacbb26a4382c.
//
// Solidity: event EtherRecovered(address indexed recipient, uint256 amount)
func (_Csaccounting *CsaccountingFilterer) WatchEtherRecovered(opts *bind.WatchOpts, sink chan<- *CsaccountingEtherRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "EtherRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingEtherRecovered)
				if err := _Csaccounting.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseEtherRecovered(log types.Log) (*CsaccountingEtherRecovered, error) {
	event := new(CsaccountingEtherRecovered)
	if err := _Csaccounting.contract.UnpackLog(event, "EtherRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Csaccounting contract.
type CsaccountingInitializedIterator struct {
	Event *CsaccountingInitialized // Event containing the contract specifics and raw log

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
func (it *CsaccountingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingInitialized)
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
		it.Event = new(CsaccountingInitialized)
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
func (it *CsaccountingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingInitialized represents a Initialized event raised by the Csaccounting contract.
type CsaccountingInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csaccounting *CsaccountingFilterer) FilterInitialized(opts *bind.FilterOpts) (*CsaccountingInitializedIterator, error) {

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &CsaccountingInitializedIterator{contract: _Csaccounting.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Csaccounting *CsaccountingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *CsaccountingInitialized) (event.Subscription, error) {

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingInitialized)
				if err := _Csaccounting.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseInitialized(log types.Log) (*CsaccountingInitialized, error) {
	event := new(CsaccountingInitialized)
	if err := _Csaccounting.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Csaccounting contract.
type CsaccountingPausedIterator struct {
	Event *CsaccountingPaused // Event containing the contract specifics and raw log

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
func (it *CsaccountingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingPaused)
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
		it.Event = new(CsaccountingPaused)
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
func (it *CsaccountingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingPaused represents a Paused event raised by the Csaccounting contract.
type CsaccountingPaused struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Csaccounting *CsaccountingFilterer) FilterPaused(opts *bind.FilterOpts) (*CsaccountingPausedIterator, error) {

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &CsaccountingPausedIterator{contract: _Csaccounting.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Csaccounting *CsaccountingFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *CsaccountingPaused) (event.Subscription, error) {

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingPaused)
				if err := _Csaccounting.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParsePaused(log types.Log) (*CsaccountingPaused, error) {
	event := new(CsaccountingPaused)
	if err := _Csaccounting.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Csaccounting contract.
type CsaccountingResumedIterator struct {
	Event *CsaccountingResumed // Event containing the contract specifics and raw log

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
func (it *CsaccountingResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingResumed)
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
		it.Event = new(CsaccountingResumed)
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
func (it *CsaccountingResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingResumed represents a Resumed event raised by the Csaccounting contract.
type CsaccountingResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Csaccounting *CsaccountingFilterer) FilterResumed(opts *bind.FilterOpts) (*CsaccountingResumedIterator, error) {

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &CsaccountingResumedIterator{contract: _Csaccounting.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Csaccounting *CsaccountingFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *CsaccountingResumed) (event.Subscription, error) {

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingResumed)
				if err := _Csaccounting.contract.UnpackLog(event, "Resumed", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseResumed(log types.Log) (*CsaccountingResumed, error) {
	event := new(CsaccountingResumed)
	if err := _Csaccounting.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Csaccounting contract.
type CsaccountingRoleAdminChangedIterator struct {
	Event *CsaccountingRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *CsaccountingRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingRoleAdminChanged)
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
		it.Event = new(CsaccountingRoleAdminChanged)
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
func (it *CsaccountingRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingRoleAdminChanged represents a RoleAdminChanged event raised by the Csaccounting contract.
type CsaccountingRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Csaccounting *CsaccountingFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*CsaccountingRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingRoleAdminChangedIterator{contract: _Csaccounting.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Csaccounting *CsaccountingFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *CsaccountingRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingRoleAdminChanged)
				if err := _Csaccounting.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseRoleAdminChanged(log types.Log) (*CsaccountingRoleAdminChanged, error) {
	event := new(CsaccountingRoleAdminChanged)
	if err := _Csaccounting.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Csaccounting contract.
type CsaccountingRoleGrantedIterator struct {
	Event *CsaccountingRoleGranted // Event containing the contract specifics and raw log

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
func (it *CsaccountingRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingRoleGranted)
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
		it.Event = new(CsaccountingRoleGranted)
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
func (it *CsaccountingRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingRoleGranted represents a RoleGranted event raised by the Csaccounting contract.
type CsaccountingRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csaccounting *CsaccountingFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CsaccountingRoleGrantedIterator, error) {

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

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingRoleGrantedIterator{contract: _Csaccounting.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csaccounting *CsaccountingFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *CsaccountingRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingRoleGranted)
				if err := _Csaccounting.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseRoleGranted(log types.Log) (*CsaccountingRoleGranted, error) {
	event := new(CsaccountingRoleGranted)
	if err := _Csaccounting.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Csaccounting contract.
type CsaccountingRoleRevokedIterator struct {
	Event *CsaccountingRoleRevoked // Event containing the contract specifics and raw log

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
func (it *CsaccountingRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingRoleRevoked)
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
		it.Event = new(CsaccountingRoleRevoked)
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
func (it *CsaccountingRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingRoleRevoked represents a RoleRevoked event raised by the Csaccounting contract.
type CsaccountingRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csaccounting *CsaccountingFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*CsaccountingRoleRevokedIterator, error) {

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

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingRoleRevokedIterator{contract: _Csaccounting.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Csaccounting *CsaccountingFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *CsaccountingRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingRoleRevoked)
				if err := _Csaccounting.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseRoleRevoked(log types.Log) (*CsaccountingRoleRevoked, error) {
	event := new(CsaccountingRoleRevoked)
	if err := _Csaccounting.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CsaccountingStETHSharesRecoveredIterator is returned from FilterStETHSharesRecovered and is used to iterate over the raw logs and unpacked data for StETHSharesRecovered events raised by the Csaccounting contract.
type CsaccountingStETHSharesRecoveredIterator struct {
	Event *CsaccountingStETHSharesRecovered // Event containing the contract specifics and raw log

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
func (it *CsaccountingStETHSharesRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CsaccountingStETHSharesRecovered)
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
		it.Event = new(CsaccountingStETHSharesRecovered)
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
func (it *CsaccountingStETHSharesRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CsaccountingStETHSharesRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CsaccountingStETHSharesRecovered represents a StETHSharesRecovered event raised by the Csaccounting contract.
type CsaccountingStETHSharesRecovered struct {
	Recipient common.Address
	Shares    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterStETHSharesRecovered is a free log retrieval operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csaccounting *CsaccountingFilterer) FilterStETHSharesRecovered(opts *bind.FilterOpts, recipient []common.Address) (*CsaccountingStETHSharesRecoveredIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.FilterLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return &CsaccountingStETHSharesRecoveredIterator{contract: _Csaccounting.contract, event: "StETHSharesRecovered", logs: logs, sub: sub}, nil
}

// WatchStETHSharesRecovered is a free log subscription operation binding the contract event 0x426e7e0100db57255d4af4a46cd49552ef74f5f002bbdc8d4ebb6371c0070a02.
//
// Solidity: event StETHSharesRecovered(address indexed recipient, uint256 shares)
func (_Csaccounting *CsaccountingFilterer) WatchStETHSharesRecovered(opts *bind.WatchOpts, sink chan<- *CsaccountingStETHSharesRecovered, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _Csaccounting.contract.WatchLogs(opts, "StETHSharesRecovered", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CsaccountingStETHSharesRecovered)
				if err := _Csaccounting.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
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
func (_Csaccounting *CsaccountingFilterer) ParseStETHSharesRecovered(log types.Log) (*CsaccountingStETHSharesRecovered, error) {
	event := new(CsaccountingStETHSharesRecovered)
	if err := _Csaccounting.contract.UnpackLog(event, "StETHSharesRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
