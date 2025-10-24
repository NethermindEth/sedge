// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vebo

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

// ValidatorsExitBusExitRequestsData is an auto generated low-level Go binding around an user-defined struct.
type ValidatorsExitBusExitRequestsData struct {
	Data       []byte
	DataFormat *big.Int
}

// ValidatorsExitBusOracleProcessingState is an auto generated low-level Go binding around an user-defined struct.
type ValidatorsExitBusOracleProcessingState struct {
	CurrentFrameRefSlot    *big.Int
	ProcessingDeadlineTime *big.Int
	DataHash               [32]byte
	DataSubmitted          bool
	DataFormat             *big.Int
	RequestsCount          *big.Int
	RequestsSubmitted      *big.Int
}

// ValidatorsExitBusOracleReportData is an auto generated low-level Go binding around an user-defined struct.
type ValidatorsExitBusOracleReportData struct {
	ConsensusVersion *big.Int
	RefSlot          *big.Int
	RequestsCount    *big.Int
	DataFormat       *big.Int
	Data             []byte
}

// VeboMetaData contains all meta data concerning the Vebo contract.
var VeboMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"secondsPerSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"genesisTime\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"lidoLocator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AddressCannotBeSame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AddressCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AdminCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exitDataIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"}],\"name\":\"ExitDataIndexOutOfRange\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitHashAlreadySubmitted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ExitHashNotSubmitted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"remainingLimit\",\"type\":\"uint256\"}],\"name\":\"ExitRequestsLimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"HashCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"initialRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"InitialRefSlotCannotBeLessThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidContractVersionIncrement\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidExitDataIndexSortOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidModuleId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestsDataLength\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRequestsDataSortOrder\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LimitExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NoConsensusReportToProcess\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonZeroContractVersionOnInit\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseUntilMustBeInFuture\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PausedExpected\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"ProcessingDeadlineMissed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RefSlotAlreadyProcessing\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotCannotDecrease\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingRefSlot\",\"type\":\"uint256\"}],\"name\":\"RefSlotMustBeGreaterThanProcessingOne\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestsAlreadyDelivered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RequestsNotDelivered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ResumedExpected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SecondsPerSlotCannotBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderIsNotTheConsensusContract\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SenderNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLargeExitsPerFrame\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLargeFrameDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TooLargeMaxExitRequestsLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxRequestsPerReport\",\"type\":\"uint256\"}],\"name\":\"TooManyExitRequestsInReport\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedChainConfig\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expectedVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receivedVersion\",\"type\":\"uint256\"}],\"name\":\"UnexpectedConsensusVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"expected\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"received\",\"type\":\"uint256\"}],\"name\":\"UnexpectedContractVersion\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"consensusHash\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"receivedHash\",\"type\":\"bytes32\"}],\"name\":\"UnexpectedDataHash\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"consensusRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataRefSlot\",\"type\":\"uint256\"}],\"name\":\"UnexpectedRefSlot\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnexpectedRequestsDataLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"format\",\"type\":\"uint256\"}],\"name\":\"UnsupportedRequestsDataFormat\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VersionCannotBeSame\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"ZeroArgument\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroFrameDuration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroPauseDuration\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"prevAddr\",\"type\":\"address\"}],\"name\":\"ConsensusHashContractSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"prevVersion\",\"type\":\"uint256\"}],\"name\":\"ConsensusVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"ContractVersionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRequestsHash\",\"type\":\"bytes32\"}],\"name\":\"ExitDataProcessing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxExitRequestsLimit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"exitsPerFrame\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"frameDurationInSec\",\"type\":\"uint256\"}],\"name\":\"ExitRequestsLimitSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ProcessingStarted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"ReportDiscarded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"}],\"name\":\"ReportSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"exitRequestsHash\",\"type\":\"bytes32\"}],\"name\":\"RequestsHashSubmitted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Resumed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxValidatorsPerReport\",\"type\":\"uint256\"}],\"name\":\"SetMaxValidatorsPerReport\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"stakingModuleId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"nodeOperatorId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"validatorIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"validatorPubkey\",\"type\":\"bytes\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"ValidatorExitRequest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestsProcessed\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"}],\"name\":\"WarnDataIncompleteProcessing\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"WarnProcessingMissed\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"DATA_FORMAT_LIST\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXIT_REQUEST_LIMIT_MANAGER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXIT_TYPE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GENESIS_TIME\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_CONSENSUS_CONTRACT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MANAGE_CONSENSUS_VERSION_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_INFINITELY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSE_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RESUME_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SECONDS_PER_SLOT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMIT_DATA_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SUBMIT_REPORT_HASH_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"}],\"name\":\"discardConsensusReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxValidatorsPerReport\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExitRequestsLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"frameDurationInSec\",\"type\":\"uint256\"}],\"name\":\"finalizeUpgrade_v2\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusReport\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"processingStarted\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getConsensusVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getContractVersion\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"exitRequestsHash\",\"type\":\"bytes32\"}],\"name\":\"getDeliveryTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"deliveryDateTimestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getExitRequestLimitFullInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxExitRequestsLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"frameDurationInSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"prevExitRequestsLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentExitRequestsLimit\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastProcessingRefSlot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxValidatorsPerReport\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProcessingState\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"currentFrameRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"processingDeadlineTime\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"dataHash\",\"type\":\"bytes32\"},{\"internalType\":\"bool\",\"name\":\"dataSubmitted\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"dataFormat\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestsSubmitted\",\"type\":\"uint256\"}],\"internalType\":\"structValidatorsExitBusOracle.ProcessingState\",\"name\":\"result\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResumeSinceTimestamp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalRequestsProcessed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"consensusContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"lastProcessingRefSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxValidatorsPerRequest\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxExitRequestsLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"frameDurationInSec\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"pauseFor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_pauseUntilInclusive\",\"type\":\"uint256\"}],\"name\":\"pauseUntil\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"resume\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"setConsensusContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"version\",\"type\":\"uint256\"}],\"name\":\"setConsensusVersion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxExitRequestsLimit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exitsPerFrame\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"frameDurationInSec\",\"type\":\"uint256\"}],\"name\":\"setExitRequestLimit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"maxRequests\",\"type\":\"uint256\"}],\"name\":\"setMaxValidatorsPerReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"reportHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"}],\"name\":\"submitConsensusReport\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"dataFormat\",\"type\":\"uint256\"}],\"internalType\":\"structValidatorsExitBus.ExitRequestsData\",\"name\":\"request\",\"type\":\"tuple\"}],\"name\":\"submitExitRequestsData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"exitRequestsHash\",\"type\":\"bytes32\"}],\"name\":\"submitExitRequestsHash\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"consensusVersion\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"refSlot\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"requestsCount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"dataFormat\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"internalType\":\"structValidatorsExitBusOracle.ReportData\",\"name\":\"data\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"contractVersion\",\"type\":\"uint256\"}],\"name\":\"submitReportData\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"dataFormat\",\"type\":\"uint256\"}],\"internalType\":\"structValidatorsExitBus.ExitRequestsData\",\"name\":\"exitsData\",\"type\":\"tuple\"},{\"internalType\":\"uint256[]\",\"name\":\"exitDataIndexes\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"refundRecipient\",\"type\":\"address\"}],\"name\":\"triggerExits\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"exitRequests\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"dataFormat\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"unpackExitRequest\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"pubkey\",\"type\":\"bytes\"},{\"internalType\":\"uint256\",\"name\":\"nodeOpId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"moduleId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"valIndex\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"}]",
	Bin: "0x6080604052600436106102ae5760003560e01c80639010d07c11610175578063b6b764b2116100dc578063d438121711610095578063e2793e721161006f578063e2793e7214610931578063f288246114610946578063f3f449c71461097a578063ff406d191461099a57600080fd5b8063d4381217146108dc578063d547741f146108fc578063e271b7741461091c57600080fd5b8063b6b764b2146107f6578063b8fe0ad014610833578063c1f665bc14610853578063c469c30714610868578063ca15c87314610888578063d072f014146108a857600080fd5b8063a52289bf1161012e578063a52289bf14610719578063ab53ac4814610739578063abe9cfc81461076d578063ad5cac4e1461078d578063b187bd26146107c1578063b1b19f57146107d657600080fd5b80639010d07c1461066757806391d14854146106875780639cc23c79146106a7578063a217fddf146106db578063a2ab7065146106f0578063a302ee381461070357600080fd5b806346e1f576116102195780637dad759d116101d25780637dad759d146105485780638aa10435146105785780638ba796af1461058d5780638d591474146105ad5780638f55b571146105cd5780638f7797c2146105fa57600080fd5b806346e1f5761461047357806356254a97146104a7578063589ff76c146104c75780635be20425146104dc57806360d64d38146104f15780636f2c322d1461052857600080fd5b80632de03aa11161026b5780632de03aa1146103825780632f2ff15d146103b6578063304b9071146103d65780633584d59c1461040a57806336568abe1461041f578063389ed2671461043f57600080fd5b806301ffc9a7146102b3578063046f7da2146102e8578063063f36ad146102ff57806306e413891461031f578063248a9ca314610342578063294492c814610362575b600080fd5b3480156102bf57600080fd5b506102d36102ce366004613891565b6109ba565b60405190151581526020015b60405180910390f35b3480156102f457600080fd5b506102fd6109e5565b005b34801561030b57600080fd5b506102fd61031a3660046138bb565b610a23565b34801561032b57600080fd5b50610334600281565b6040519081526020016102df565b34801561034e57600080fd5b5061033461035d3660046138e7565b610c1b565b34801561036e57600080fd5b506102fd61037d366004613900565b610c3d565b34801561038e57600080fd5b506103347f2fc10cc8ae19568712f7a176fb4978616a610650813c9d05326c34abb62749c781565b3480156103c257600080fd5b506102fd6103d136600461396e565b610d28565b3480156103e257600080fd5b506103347f000000000000000000000000000000000000000000000000000000000000000c81565b34801561041657600080fd5b50610334610d4a565b34801561042b57600080fd5b506102fd61043a36600461396e565b610d67565b34801561044b57600080fd5b506103347f139c2898040ef16910dc9f44dc697df79363da767d8bc92f2e310312b816e46d81565b34801561047f57600080fd5b506103347f65fa0c17458517c727737e4153dd477fa3e328cf706640b0f68b1a285c5990da81565b3480156104b357600080fd5b506102fd6104c23660046138bb565b610de5565b3480156104d357600080fd5b50610334610e21565b3480156104e857600080fd5b50610334610e39565b3480156104fd57600080fd5b50610506610e51565b60408051948552602085019390935291830152151560608201526080016102df565b34801561053457600080fd5b506102fd6105433660046138e7565b610ef9565b34801561055457600080fd5b5061056861056336600461399e565b610f2d565b6040516102df9493929190613a74565b34801561058457600080fd5b50610334610faf565b34801561059957600080fd5b506102fd6105a8366004613aa3565b610fd9565b3480156105b957600080fd5b506102fd6105c83660046138e7565b611037565b3480156105d957600080fd5b506105e261106b565b6040516001600160a01b0390911681526020016102df565b34801561060657600080fd5b5061060f611083565b6040516102df9190600060e0820190508251825260208301516020830152604083015160408301526060830151151560608301526080830151608083015260a083015160a083015260c083015160c083015292915050565b34801561067357600080fd5b506105e2610682366004613b0e565b61121a565b34801561069357600080fd5b506102d36106a236600461396e565b611246565b3480156106b357600080fd5b506103347fc31b1e4b732c5173dc51d519dfa432bad95550ecc4b0f9a61c2a558a2a8e434181565b3480156106e757600080fd5b50610334600081565b6102fd6106fe366004613b48565b61127e565b34801561070f57600080fd5b5061033460001981565b34801561072557600080fd5b506103346107343660046138e7565b6116bb565b34801561074557600080fd5b506103347f9c616dd118785b2e2fccf45a4ff151a335ff7b6a84cd1c4d7fd9f97f39ea934281565b34801561077957600080fd5b506102fd6107883660046138e7565b6116ff565b34801561079957600080fd5b506103347f04a0afbbd09d5ad397fc858789da4f8edd59f5ca5098d70faa490babee945c3b81565b3480156107cd57600080fd5b506102d3611733565b3480156107e257600080fd5b506102fd6107f13660046138e7565b611752565b34801561080257600080fd5b5061080b61179d565b604080519586526020860194909452928401919091526060830152608082015260a0016102df565b34801561083f57600080fd5b506102fd61084e366004613bf2565b611810565b34801561085f57600080fd5b506103346119b8565b34801561087457600080fd5b506102fd610883366004613c2e565b6119c2565b34801561089457600080fd5b506103346108a33660046138e7565b611a0c565b3480156108b457600080fd5b506103347f22ebb4dbafb72948800c1e1afa1688772a1a4cfc54d5ebfcec8163b1139c082e81565b3480156108e857600080fd5b506102fd6108f73660046138e7565b611a30565b34801561090857600080fd5b506102fd61091736600461396e565b611b73565b34801561092857600080fd5b50610334600181565b34801561093d57600080fd5b50610334611b90565b34801561095257600080fd5b506103347f000000000000000000000000000000000000000000000000000000005fc6305781565b34801561098657600080fd5b506102fd6109953660046138e7565b611ba8565b3480156109a657600080fd5b506102fd6109b5366004613c4b565b611bdc565b60006001600160e01b03198216635a05180f60e01b14806109df57506109df82611bec565b92915050565b6109ed611c21565b7f2fc10cc8ae19568712f7a176fb4978616a610650813c9d05326c34abb62749c7610a188133611c48565b610a20611cac565b50565b610a2b611cf0565b60008051602061420b833981519152546001600160401b031680831015610a745760405163431d301760e11b815260048101849052602481018290526044015b60405180910390fd5b6000610a8c6000805160206141cb8339815191525490565b9050808411610ab8576040516360a41e4960e01b81526004810185905260248101829052604401610a6b565b82421115610adc5760405163537bacdf60e11b815260048101849052602401610a6b565b818414158015610aec5750818114155b15610b1d5760405182907f800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c3890600090a25b84610b3b57604051635b18a69f60e11b815260040160405180910390fd5b604080518681526020810185905285917faed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a910160405180910390a260006040518060600160405280878152602001610b9287611d31565b6001600160401b03168152602001610ba986611d31565b6001600160401b03908116909152815160008051602061410b83398151915255602082015160008051602061420b833981519152805460408501518416600160401b026001600160801b031990911692909316919091179190911790559050610c13818484611d9d565b505050505050565b60009081526000805160206141ab833981519152602052604090206001015490565b610c45611e87565b610c4d611ead565b610c5681611f09565b6000610c656080840184613c7d565b8460600135604051602001610c7c93929190613cf3565b604051602081830303815290604052805190602001209050600083604051602001610ca79190613d17565b604051602081830303815290604052805190602001209050610cd28460200135856000013583611f3f565b610cda612037565b50610ce5828461215d565b610cee84612168565b6040518281527f01b8de053572c3c2104259b555c485ccac8017196b3471e8483b7e96f071608a906020015b60405180910390a150505050565b610d3182610c1b565b610d3b8133611c48565b610d458383612429565b505050565b6000610d626000805160206141cb8339815191525490565b905090565b6001600160a01b0381163314610dd75760405162461bcd60e51b815260206004820152602f60248201527f416363657373436f6e74726f6c3a2063616e206f6e6c792072656e6f756e636560448201526e103937b632b9903337b91039b2b63360891b6064820152608401610a6b565b610de18282612458565b5050565b7f9c616dd118785b2e2fccf45a4ff151a335ff7b6a84cd1c4d7fd9f97f39ea9342610e108133611c48565b610e1b848484612487565b50505050565b6000610d626000805160206141eb8339815191525490565b6000610d6260008051602061422b8339815191525490565b60008080808060008051602061410b83398151915260408051606081018252825481526001909201546001600160401b038082166020850152600160401b909104169082015290506000610eb16000805160206141cb8339815191525490565b82516020840151604085015192935090918215801590610edd57508385602001516001600160401b0316145b92996001600160401b0392831699509116965090945092505050565b7f9c616dd118785b2e2fccf45a4ff151a335ff7b6a84cd1c4d7fd9f97f39ea9342610f248133611c48565b610de18261259e565b60606000806000610f3f888888612645565b610f4a604088613dcc565b8510610f7e5784610f5c604089613dcc565b6040516394a0fcb760e01b815260048101929092526024820152604401610a6b565b6000610f8b89898861269d565b6040810151815160208301516060909301519c909b50919950975095505050505050565b6000610d627f4dd0f6662ba1d6b081f08b350f5e9a6a7b15cf586926ba66f753594928fa64a65490565b6001600160a01b03881661100057604051636b35b1b760e01b815260040160405180910390fd5b61100b600089612744565b61101660001961274e565b61102187878761279f565b61102d84848484612812565b5050505050505050565b7fc31b1e4b732c5173dc51d519dfa432bad95550ecc4b0f9a61c2a558a2a8e43416110628133611c48565b610de182612830565b6000610d6260008051602061416b8339815191525490565b6110c86040518060e001604052806000815260200160008152602001600080191681526020016000151581526020016000815260200160008152602001600081525090565b6040805160608101825260008051602061410b83398151915254815260008051602061420b833981519152546001600160401b038082166020840152600160401b909104169181019190915261111c6128b3565b82528051158061113d575080602001516001600160401b0316826000015114155b15611146575090565b6040818101516001600160401b0316602084015281519083015260007ff54f01aac0787b485340ed16cefe4fba326c1674376c8dcd7c2a644b4643792f6040805160808101825291546001600160401b03808216808552600160401b830482166020860152600160801b830490911692840192909252600160c01b900461ffff1660608084019190915285519091149085018190529091506111e757505090565b606081015161ffff16608084015260208101516001600160401b0390811660a08501526040909101511660c08301525090565b600082815260008051602061418b8339815191526020526040812061123f9083612948565b9392505050565b60009182526000805160206141ab833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b611286611e87565b60006112923447613de0565b9050346112ce576040516356e4289360e01b81526020600482015260096024820152686d73672e76616c756560b81b6044820152606401610a6b565b8261130e576040516356e4289360e01b815260206004820152600f60248201526e6578697444617461496e646578657360881b6044820152606401610a6b565b6001600160a01b038216611320573391505b600060008051602061412b833981519152600061133d8880613c7d565b896020013560405160200161135493929190613cf3565b604051602081830303815290604052805190602001208152602001908152602001600020905061138381612954565b61138c81612979565b6113a36113998780613c7d565b8860200135612645565b6000846001600160401b038111156113bd576113bd613df7565b60405190808252806020026020018201604052801561141257816020015b6113ff60405180606001604052806000815260200160008152602001606081525090565b8152602001906001900390816113db5790505b509050600019600060406114268a80613c7d565b611431929150613dcc565b905060005b878110156115ac578189898381811061145157611451613e0d565b905060200201351061149b5788888281811061146f5761146f613e0d565b90506020020135826040516394a0fcb760e01b8152600401610a6b929190918252602082015260400190565b6000811180156114c35750828989838181106114b9576114b9613e0d565b9050602002013511155b156114e1576040516307032e6360e41b815260040160405180910390fd5b8888828181106114f3576114f3613e0d565b602002919091013593506000905061152c61150e8c80613c7d565b8c8c8681811061152057611520613e0d565b9050602002013561269d565b905080602001516000141561155457604051634632571560e01b815260040160405180910390fd5b60405180606001604052808260200151815260200182600001518152602001826060015181525085838151811061158d5761158d613e0d565b60200260200101819052505080806115a490613e23565b915050611436565b507f000000000000000000000000c1d0b3de6792bf6b4b37eccdcc24e45978cfd2eb6001600160a01b031663d6dff5806040518163ffffffff1660e01b815260040160206040518083038186803b15801561160657600080fd5b505afa15801561161a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061163e9190613e3e565b6001600160a01b031663138b1b1534858960026040518563ffffffff1660e01b815260040161166f93929190613e5b565b6000604051808303818588803b15801561168857600080fd5b505af115801561169c573d6000803e3d6000fd5b5050505050505050508047146116b4576116b4613efa565b5050505050565b600081815260008051602061412b8339815191526020819052604082206116e181612954565b6116ea81612979565b54600160201b900463ffffffff169392505050565b7f139c2898040ef16910dc9f44dc697df79363da767d8bc92f2e310312b816e46d61172a8133611c48565b610de1826129a5565b600061174b6000805160206141eb8339815191525490565b4210905090565b61175a611e87565b7f22ebb4dbafb72948800c1e1afa1688772a1a4cfc54d5ebfcec8163b1139c082e6117858133611c48565b600061178f610faf565b9050610d45838260006129f8565b6000806000806000806117bd60008051602061414b833981519152612ab9565b805160808201516060830151602084015163ffffffff9384169a509183169850821696501693509050856117f357600019611806565b6118068163ffffffff42811690612b3a16565b9150509091929394565b611818611e87565b60006118248280613c7d565b836020013560405160200161183b93929190613cf3565b604051602081830303815290604052805190602001209050600061186a60008051602061412b83398151915290565b600083815260209190915260409020905061188481612954565b61188d81612bfe565b6118a461189a8480613c7d565b8560200135612645565b80546118b59063ffffffff16611f09565b600060406118c38580613c7d565b6118ce929150613dcc565b905060006118da612c2b565b90508082111561190a576040516001620518fd60e41b031981526004810183905260248101829052604401610a6b565b61191382612c55565b6119256119208680613c7d565b612ceb565b61195b8261193f6000805160206140eb8339815191525490565b6119499190613f10565b6000805160206140eb83398151915255565b61197e83805467ffffffff000000001916600160201b63ffffffff421602179055565b6040518481527f01b8de053572c3c2104259b555c485ccac8017196b3471e8483b7e96f071608a9060200160405180910390a15050505050565b6000610d62612c2b565b7f04a0afbbd09d5ad397fc858789da4f8edd59f5ca5098d70faa490babee945c3b6119ed8133611c48565b610de182611a076000805160206141cb8339815191525490565b612dcc565b600081815260008051602061418b833981519152602052604081206109df90613025565b611a38611cf0565b6040805160608101825260008051602061410b83398151915254815260008051602061420b833981519152546001600160401b0380821660208401819052600160401b909204169282019290925290821015611abf57602081015160405163431d301760e11b8152600481018490526001600160401b039091166024820152604401610a6b565b80602001516001600160401b0316821115611ad8575050565b6000611af06000805160206141cb8339815191525490565b9050808311611b11576040516252e2c960e41b815260040160405180910390fd5b600060008051602061410b8339815191525581602001516001600160401b03167fe21266bc27ee721ac10034efaf7fd724656ef471c75b8402cd8f07850af6b6768360000151604051611b6691815260200190565b60405180910390a2505050565b611b7c82610c1b565b611b868133611c48565b610d458383612458565b6000610d626000805160206140eb8339815191525490565b7f139c2898040ef16910dc9f44dc697df79363da767d8bc92f2e310312b816e46d611bd38133611c48565b610de18261274e565b610e1b84848484612812565b9055565b60006001600160e01b03198216637965db0b60e01b14806109df57506301ffc9a760e01b6001600160e01b03198316146109df565b611c29611733565b611c465760405163b047186b60e01b815260040160405180910390fd5b565b611c528282611246565b610de157611c6a816001600160a01b0316601461302f565b611c7583602061302f565b604051602001611c86929190613f28565b60408051601f198184030181529082905262461bcd60e51b8252610a6b91600401613f9d565b611cb4611c21565b426000805160206141eb833981519152556040517f62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f990600090a1565b60008051602061416b833981519152546001600160a01b0316336001600160a01b031614611c465760405163fef4d83160e01b815260040160405180910390fd5b60006001600160401b03821115611d995760405162461bcd60e51b815260206004820152602660248201527f53616665436173743a2076616c756520646f65736e27742066697420696e203660448201526534206269747360d01b6064820152608401610a6b565b5090565b604080516080810182527ff54f01aac0787b485340ed16cefe4fba326c1674376c8dcd7c2a644b4643792f546001600160401b03808216808452600160401b830482166020850152600160801b830490911693830193909352600160c01b900461ffff1660608201529082148015611e2e575080602001516001600160401b031681604001516001600160401b0316105b15610e1b5760408082015160208084015183516001600160401b03938416815292169082015283917fefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4910160405180910390a250505050565b611e8f611733565b15611c4657604051630286f07360e31b815260040160405180910390fd5b33611ed87f65fa0c17458517c727737e4153dd477fa3e328cf706640b0f68b1a285c5990da82611246565b158015611eeb5750611ee9816131ca565b155b15610a20576040516323dada5360e01b815260040160405180910390fd5b6000611f13610faf565b9050808214610de1576040516303abe78360e21b81526004810182905260248101839052604401610a6b565b6040805160608101825260008051602061410b83398151915254815260008051602061420b833981519152546001600160401b0380821660208401819052600160401b9092041692820192909252908414611fc557602081015160405163490b8d4560e11b81526001600160401b03909116600482015260248101859052604401610a6b565b6000611fdd60008051602061422b8339815191525490565b905080841461200957604051632a37dd3d60e11b81526004810182905260248101859052604401610a6b565b815183146116b457815160405163642c75c760e11b8152600481019190915260248101849052604401610a6b565b6040805160608101825260008051602061410b8339815191525480825260008051602061420b833981519152546001600160401b038082166020850152600160401b90910416928201929092526000916120a4576040516364dfc18f60e01b815260040160405180910390fd5b6120ba81604001516001600160401b0316613260565b60006120d26000805160206141cb8339815191525490565b905081602001516001600160401b0316811415612101576040516252e2c960e41b815260040160405180910390fd5b6020828101516001600160401b03166000805160206141cb833981519152819055835160405190815290917ff73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a910160405180910390a292915050565b610de1828242613284565b600181606001351461219357604051630f542bef60e21b815260608201356004820152602401610a6b565b60406121a26080830183613c7d565b6121ad929150613fb0565b156121cb57604051630260e4e160e41b815260040160405180910390fd5b604080820135906121df6080840184613c7d565b6121ea929150613dcc565b146122085760405163f34afee160e01b815260040160405180910390fd5b7f000000000000000000000000c1d0b3de6792bf6b4b37eccdcc24e45978cfd2eb6001600160a01b031663f5e6d50f6040518163ffffffff1660e01b815260040160206040518083038186803b15801561226157600080fd5b505afa158015612275573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906122999190613e3e565b6001600160a01b0316633e0865dd82604001356040518263ffffffff1660e01b81526004016122ca91815260200190565b60006040518083038186803b1580156122e257600080fd5b505afa1580156122f6573d6000803e3d6000fd5b5061230c92506119209150506080830183613c7d565b60405180608001604052806123248360200135611d31565b6001600160401b0316815260200161233f8360400135611d31565b6001600160401b0316815260200161235a8360400135611d31565b6001600160401b0316815260016020909101527ff54f01aac0787b485340ed16cefe4fba326c1674376c8dcd7c2a644b4643792f81518154602084015160408086015160609096015161ffff16600160c01b0261ffff60c01b196001600160401b03978816600160801b021669ffffffffffffffffffff60801b19938816600160401b026001600160801b0319909516979095169690961792909217169190911792909217905581013561240b5750565b610a20816040013561193f6000805160206140eb8339815191525490565b61243382826132bb565b600082815260008051602061418b83398151915260205260409020610d459082613331565b6124628282613346565b600082815260008051602061418b83398151915260205260409020610d4590826133ba565b63ffffffff42166125606124b8858585856124af60008051602061414b833981519152612ab9565b939291906133cf565b60008051602061414b833981519152908051825460208301516040840151606085015160809095015163ffffffff908116600160801b0263ffffffff60801b19968216600160601b0263ffffffff60601b19938316600160401b02939093166fffffffffffffffff000000000000000019948316600160201b0267ffffffffffffffff1990961692909616919091179390931791909116929092179190911791909116179055565b60408051858152602081018590529081018390527f3119d910326e0f179e121df55f23f45b8a5022ff10c73c02aabf2b48ae36070a90606001610d1a565b806125e5576040516356e4289360e01b81526020600482015260166024820152751b585e15985b1a59185d1bdc9cd4195c94995c1bdc9d60521b6044820152606401610a6b565b61260e7f4f034b2ceac9c934b225eea10a2de790b8651266f9b21cd9dbcb244c075e324c829055565b6040518181527f9b17a153b6e933d8497c6b713fbd70c893891d75639ede17ce6e4cea08e7cfc3906020015b60405180910390a150565b6001811461266957604051630f542bef60e21b815260048101829052602401610a6b565b81158061267f575061267c604083613fb0565b15155b15610d4557604051630260e4e160e41b815260040160405180910390fd5b6126c86040518060800160405280600081526020016000815260200160008152602001606081525090565b6001600160401b03604083810286018035608081901c9384168584015264ffffffffff60c082901c16855260e81c60208501528151603080825260608201909352909291600091906020820181803683375050506040860288019350905060108301602082016030828237505060608401525090949350505050565b610de18282612429565b612756611e87565b806127745760405163ad58bfc760e01b815260040160405180910390fd5b60006000198214156127895750600019612796565b6127938242613f10565b90505b610de18161351e565b6127a960016135aa565b6127b38382612dcc565b6127bc82612830565b6127d36000805160206141cb833981519152829055565b6127dc81611d31565b60008051602061410b833981519152600101805467ffffffffffffffff19166001600160401b0392909216919091179055505050565b61281c60026135d9565b6128258461259e565b610e1b838383612487565b600061284860008051602061422b8339815191525490565b90508082141561286b57604051631d7c761b60e21b815260040160405180910390fd5b61288260008051602061422b833981519152839055565b604051819083907ffa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a89590600090a35050565b6000806128cc60008051602061416b8339815191525490565b90506000816001600160a01b03166372f79b136040518163ffffffff1660e01b8152600401604080518083038186803b15801561290857600080fd5b505afa15801561291c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906129409190613fc4565b509392505050565b600061123f838361360b565b805463ffffffff16610a2057604051633090184b60e11b815260040160405180910390fd5b8054600160201b900463ffffffff16610a20576040516379eeb04b60e11b815260040160405180910390fd5b6129ad611e87565b428110156129ce576040516339e2ec5360e11b815260040160405180910390fd5b600060001982146129eb576129e4826001613f10565b9050612796565b50600019610de18161351e565b600083815260008051602061412b833981519152602081905260409091205463ffffffff1615612a3b576040516334dd7fc760e11b815260040160405180910390fd5b60408051808201825263ffffffff808616825284811660208084019182526000898152908690528490209251835491518316600160201b0267ffffffffffffffff19909216921691909117179055517f76d8359ea28964b79f7fa8bb502ec325fd0d1e956c42a0436940e35d0e99f2de90610d1a9086815260200190565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152816040805160a081018252915463ffffffff8082168452600160201b820481166020850152600160401b8204811692840192909252600160601b810482166060840152600160801b900416608082015292915050565b600080836040015163ffffffff1683612b539190613de0565b9050836060015163ffffffff16811080612b755750608084015163ffffffff16155b15612b8c575050602082015163ffffffff166109df565b6000846060015163ffffffff1682612ba49190613dcc565b90506000856080015163ffffffff1682612bbe9190613fe8565b9050600081876020015163ffffffff16612bd89190613f10565b875190915063ffffffff16811115612bf45750855163ffffffff165b9695505050505050565b8054600160201b900463ffffffff1615610a2057604051631cbf22fb60e21b815260040160405180910390fd5b6000610d627f4f034b2ceac9c934b225eea10a2de790b8651266f9b21cd9dbcb244c075e324c5490565b6000612c6e60008051602061414b833981519152612ab9565b9050612c80815163ffffffff16151590565b612c88575050565b6000612c9d8263ffffffff42811690612b3a16565b905080831115612cca5760405163106865a560e31b81526004810184905260248101829052604401610a6b565b610d456124b8612cda8584613de0565b849063ffffffff4281169061363516565b81818101600063ffffffff4216366030838080805b888a1015612dbe5760408a019960108101965035608081901c945060e81c925082612d3e57604051634632571560e01b815260040160405180910390fd5b878411612d5e576040516362851b8960e01b815260040160405180910390fd5b839050604084901c64ffffffffff169150839750806001600160401b031682847f96395f55c4997466e5035d777f0e1ba82b8cae217aaad05cf07839eb7c75bcf289898c604051612db193929190613cf3565b60405180910390a4612d00565b505050505050505050505050565b6001600160a01b038216612df3576040516303988b8160e61b815260040160405180910390fd5b6000612e0b60008051602061416b8339815191525490565b9050806001600160a01b0316836001600160a01b03161415612e40576040516321a55ce160e11b815260040160405180910390fd5b600080846001600160a01b031663606c0c946040518163ffffffff1660e01b815260040160606040518083038186803b158015612e7c57600080fd5b505afa158015612e90573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612eb49190614007565b92509250507f000000000000000000000000000000000000000000000000000000000000000c82141580612f0857507f000000000000000000000000000000000000000000000000000000005fc630578114155b15612f2657604051635401d0a160e11b815260040160405180910390fd5b6000856001600160a01b0316636095012f6040518163ffffffff1660e01b815260040160206040518083038186803b158015612f6157600080fd5b505afa158015612f75573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190612f999190614035565b905084811015612fc657604051631e779ad160e11b81526004810182905260248101869052604401610a6b565b612fdd60008051602061416b833981519152879055565b836001600160a01b0316866001600160a01b03167f25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c60405160405180910390a3505050505050565b60006109df825490565b6060600061303e836002613fe8565b613049906002613f10565b6001600160401b0381111561306057613060613df7565b6040519080825280601f01601f19166020018201604052801561308a576020820181803683370190505b509050600360fc1b816000815181106130a5576130a5613e0d565b60200101906001600160f81b031916908160001a905350600f60fb1b816001815181106130d4576130d4613e0d565b60200101906001600160f81b031916908160001a90535060006130f8846002613fe8565b613103906001613f10565b90505b600181111561317b576f181899199a1a9b1b9c1cb0b131b232b360811b85600f166010811061313757613137613e0d565b1a60f81b82828151811061314d5761314d613e0d565b60200101906001600160f81b031916908160001a90535060049490941c936131748161404e565b9050613106565b50831561123f5760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610a6b565b6000806131e360008051602061416b8339815191525490565b604051631951c03760e01b81526001600160a01b03858116600483015291925090821690631951c0379060240160206040518083038186803b15801561322857600080fd5b505afa15801561323c573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061123f9190614065565b80421115610a205760405163537bacdf60e11b815260048101829052602401610a6b565b600083815260008051602061412b8339815191526020819052604090912054600160201b900463ffffffff1615612a3b5750505050565b6132c58282611246565b610de15760008281526000805160206141ab833981519152602090815260408083206001600160a01b0385168085529252808320805460ff1916600117905551339285917f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d9190a45050565b600061123f836001600160a01b0384166136f6565b6133508282611246565b15610de15760008281526000805160206141ab833981519152602090815260408083206001600160a01b0385168085529252808320805460ff1916905551339285917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a45050565b600061123f836001600160a01b038416613745565b6040805160a08101825260008082526020820181905291810182905260608101829052608081019190915263ffffffff85111561341f57604051635752823560e11b815260040160405180910390fd5b63ffffffff8311156134445760405163bbdd2da360e01b815260040160405180910390fd5b848411156134655760405163528f486360e01b815260040160405180910390fd5b8261348357604051636765a75d60e01b815260040160405180910390fd5b855163ffffffff166134a05763ffffffff851660208701526134fa565b60006134ac8784612b3a565b905060008188600001516134c09190614087565b9050868163ffffffff16106134db57600060208901526134f7565b6134eb63ffffffff821688613de0565b63ffffffff1660208901525b50505b5063ffffffff92831660808601529082166060850152918116835216604082015290565b6135356000805160206141eb833981519152829055565b6000198114156135715760405160001981527f32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e9060200161263a565b7f32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e61359c4283613de0565b60405190815260200161263a565b6135b2610faf565b156135d05760405163184e52a160e21b815260040160405180910390fd5b610a2081613838565b6135e1610faf565b6135ec906001613f10565b81146135d05760405163167679d560e01b815260040160405180910390fd5b600082600001828154811061362257613622613e0d565b9060005260206000200154905092915050565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152835163ffffffff1683111561368857604051631930e3c960e11b815260040160405180910390fd5b6000846040015163ffffffff16836136a09190613de0565b9050846060015163ffffffff16816136b89190613fb0565b6136c29082613de0565b63ffffffff8516602087015260408601805191925082916136e49083906140ac565b63ffffffff1690525093949350505050565b600081815260018301602052604081205461373d575081546001818101845560008481526020808220909301849055845484825282860190935260409020919091556109df565b5060006109df565b6000818152600183016020526040812054801561382e576000613769600183613de0565b855490915060009061377d90600190613de0565b90508181146137e257600086600001828154811061379d5761379d613e0d565b90600052602060002001549050808760000184815481106137c0576137c0613e0d565b6000918252602080832090910192909255918252600188019052604090208390555b85548690806137f3576137f36140d4565b6001900381819060005260206000200160009055905585600101600086815260200190815260200160002060009055600193505050506109df565b60009150506109df565b6138617f4dd0f6662ba1d6b081f08b350f5e9a6a7b15cf586926ba66f753594928fa64a6829055565b6040518181527ffddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb9060200161263a565b6000602082840312156138a357600080fd5b81356001600160e01b03198116811461123f57600080fd5b6000806000606084860312156138d057600080fd5b505081359360208301359350604090920135919050565b6000602082840312156138f957600080fd5b5035919050565b6000806040838503121561391357600080fd5b82356001600160401b0381111561392957600080fd5b830160a0818603121561393b57600080fd5b946020939093013593505050565b6001600160a01b0381168114610a2057600080fd5b803561396981613949565b919050565b6000806040838503121561398157600080fd5b82359150602083013561399381613949565b809150509250929050565b600080600080606085870312156139b457600080fd5b84356001600160401b03808211156139cb57600080fd5b818701915087601f8301126139df57600080fd5b8135818111156139ee57600080fd5b886020828501011115613a0057600080fd5b6020928301999098509187013596604001359550909350505050565b60005b83811015613a37578181015183820152602001613a1f565b83811115610e1b5750506000910152565b60008151808452613a60816020860160208601613a1c565b601f01601f19169290920160200192915050565b608081526000613a876080830187613a48565b6020830195909552506040810192909252606090910152919050565b600080600080600080600080610100898b031215613ac057600080fd5b8835613acb81613949565b97506020890135613adb81613949565b979a9799505050506040860135956060810135956080820135955060a0820135945060c0820135935060e0909101359150565b60008060408385031215613b2157600080fd5b50508035926020909101359150565b600060408284031215613b4257600080fd5b50919050565b60008060008060608587031215613b5e57600080fd5b84356001600160401b0380821115613b7557600080fd5b613b8188838901613b30565b95506020870135915080821115613b9757600080fd5b818701915087601f830112613bab57600080fd5b813581811115613bba57600080fd5b8860208260051b8501011115613bcf57600080fd5b602083019550809450505050613be76040860161395e565b905092959194509250565b600060208284031215613c0457600080fd5b81356001600160401b03811115613c1a57600080fd5b613c2684828501613b30565b949350505050565b600060208284031215613c4057600080fd5b813561123f81613949565b60008060008060808587031215613c6157600080fd5b5050823594602084013594506040840135936060013592509050565b6000808335601e19843603018112613c9457600080fd5b8301803591506001600160401b03821115613cae57600080fd5b602001915036819003821315613cc357600080fd5b9250929050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b604081526000613d07604083018587613cca565b9050826020830152949350505050565b602081528135602082015260208201356040820152604082013560608201526060820135608082015260006080830135601e19843603018112613d5957600080fd5b830180356001600160401b03811115613d7157600080fd5b803603851315613d8057600080fd5b60a080850152613d9760c085018260208501613cca565b95945050505050565b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600082613ddb57613ddb613da0565b500490565b600082821015613df257613df2613db6565b500390565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b6000600019821415613e3757613e37613db6565b5060010190565b600060208284031215613e5057600080fd5b815161123f81613949565b6000606080830181845280875180835260808601915060808160051b87010192506020808a0160005b83811015613ece57888603607f190185528151805187528381015184880152604090810151908701889052613ebb88880182613a48565b9650509382019390820190600101613e84565b5050839550613ee78188018a6001600160a01b03169052565b5050505050826040830152949350505050565b634e487b7160e01b600052600160045260246000fd5b60008219821115613f2357613f23613db6565b500190565b7f416363657373436f6e74726f6c3a206163636f756e7420000000000000000000815260008351613f60816017850160208801613a1c565b7001034b99036b4b9b9b4b733903937b6329607d1b6017918401918201528351613f91816028840160208801613a1c565b01602801949350505050565b60208152600061123f6020830184613a48565b600082613fbf57613fbf613da0565b500690565b60008060408385031215613fd757600080fd5b505080516020909101519092909150565b600081600019048311821515161561400257614002613db6565b500290565b60008060006060848603121561401c57600080fd5b8351925060208401519150604084015190509250925092565b60006020828403121561404757600080fd5b5051919050565b60008161405d5761405d613db6565b506000190190565b60006020828403121561407757600080fd5b8151801515811461123f57600080fd5b600063ffffffff838116908316818110156140a4576140a4613db6565b039392505050565b600063ffffffff8083168185168083038211156140cb576140cb613db6565b01949350505050565b634e487b7160e01b600052603160045260246000fdfe423c0a70d629d0b16eb0cfb674ba25f8352fe47057f0f4af829a850a22c6cc4a9d565e483b8608dc09e04eff85533859683d2eeaa6ebc28af53a92d7dba3eea6dce71a9cf7bc22bcecc928846e7ac6e9b30fcfbdb2d141112104f53e9c26d37f572f3e3fa8126dcb14368767182c30d88ccaa062c1c9dea242e8fca84104691eb0e01b719c2c32a677822ce1584cb6a66e576ee3c2c506b9621dbe626355aa658f8c450dae5029cd48cd91dd9db65da48fb742893edfc7941250f6721d93cbbe9a627a5d4aa7c17f87ff26e3fe9a42c2b6c559e8b41a42282d0ecebb17c0e4d3c9bdcd6eb2e956ecf03d8d27bee4c163f9b5c078aa69020d618e76513b5d0a94e8b012900cb200ee5dfc3b895a32791b67d12891b09f117814f167a237783a029d565e483b8608dc09e04eff85533859683d2eeaa6ebc28af53a92d7dba3eea72767d6892477f8d2750fb44e817c9aed93d34d3c6be4101ed58bcac692c99e9ca2646970667358221220bbb7dc2077ee18343d93363a94175927c00cf9cf4ca2f750c24af87082a2726b64736f6c63430008090033",
}

// VeboABI is the input ABI used to generate the binding from.
// Deprecated: Use VeboMetaData.ABI instead.
var VeboABI = VeboMetaData.ABI

// VeboBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VeboMetaData.Bin instead.
var VeboBin = VeboMetaData.Bin

// DeployVebo deploys a new Ethereum contract, binding an instance of Vebo to it.
func DeployVebo(auth *bind.TransactOpts, backend bind.ContractBackend, secondsPerSlot *big.Int, genesisTime *big.Int, lidoLocator common.Address) (common.Address, *types.Transaction, *Vebo, error) {
	parsed, err := VeboMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VeboBin), backend, secondsPerSlot, genesisTime, lidoLocator)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vebo{VeboCaller: VeboCaller{contract: contract}, VeboTransactor: VeboTransactor{contract: contract}, VeboFilterer: VeboFilterer{contract: contract}}, nil
}

// Vebo is an auto generated Go binding around an Ethereum contract.
type Vebo struct {
	VeboCaller     // Read-only binding to the contract
	VeboTransactor // Write-only binding to the contract
	VeboFilterer   // Log filterer for contract events
}

// VeboCaller is an auto generated read-only Go binding around an Ethereum contract.
type VeboCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeboTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VeboTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeboFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VeboFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VeboSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VeboSession struct {
	Contract     *Vebo             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VeboCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VeboCallerSession struct {
	Contract *VeboCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VeboTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VeboTransactorSession struct {
	Contract     *VeboTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VeboRaw is an auto generated low-level Go binding around an Ethereum contract.
type VeboRaw struct {
	Contract *Vebo // Generic contract binding to access the raw methods on
}

// VeboCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VeboCallerRaw struct {
	Contract *VeboCaller // Generic read-only contract binding to access the raw methods on
}

// VeboTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VeboTransactorRaw struct {
	Contract *VeboTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVebo creates a new instance of Vebo, bound to a specific deployed contract.
func NewVebo(address common.Address, backend bind.ContractBackend) (*Vebo, error) {
	contract, err := bindVebo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vebo{VeboCaller: VeboCaller{contract: contract}, VeboTransactor: VeboTransactor{contract: contract}, VeboFilterer: VeboFilterer{contract: contract}}, nil
}

// NewVeboCaller creates a new read-only instance of Vebo, bound to a specific deployed contract.
func NewVeboCaller(address common.Address, caller bind.ContractCaller) (*VeboCaller, error) {
	contract, err := bindVebo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VeboCaller{contract: contract}, nil
}

// NewVeboTransactor creates a new write-only instance of Vebo, bound to a specific deployed contract.
func NewVeboTransactor(address common.Address, transactor bind.ContractTransactor) (*VeboTransactor, error) {
	contract, err := bindVebo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VeboTransactor{contract: contract}, nil
}

// NewVeboFilterer creates a new log filterer instance of Vebo, bound to a specific deployed contract.
func NewVeboFilterer(address common.Address, filterer bind.ContractFilterer) (*VeboFilterer, error) {
	contract, err := bindVebo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VeboFilterer{contract: contract}, nil
}

// bindVebo binds a generic wrapper to an already deployed contract.
func bindVebo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VeboMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vebo *VeboRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vebo.Contract.VeboCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vebo *VeboRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vebo.Contract.VeboTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vebo *VeboRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vebo.Contract.VeboTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vebo *VeboCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vebo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vebo *VeboTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vebo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vebo *VeboTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vebo.Contract.contract.Transact(opts, method, params...)
}

// DATAFORMATLIST is a free data retrieval call binding the contract method 0xe271b774.
//
// Solidity: function DATA_FORMAT_LIST() view returns(uint256)
func (_Vebo *VeboCaller) DATAFORMATLIST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "DATA_FORMAT_LIST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DATAFORMATLIST is a free data retrieval call binding the contract method 0xe271b774.
//
// Solidity: function DATA_FORMAT_LIST() view returns(uint256)
func (_Vebo *VeboSession) DATAFORMATLIST() (*big.Int, error) {
	return _Vebo.Contract.DATAFORMATLIST(&_Vebo.CallOpts)
}

// DATAFORMATLIST is a free data retrieval call binding the contract method 0xe271b774.
//
// Solidity: function DATA_FORMAT_LIST() view returns(uint256)
func (_Vebo *VeboCallerSession) DATAFORMATLIST() (*big.Int, error) {
	return _Vebo.Contract.DATAFORMATLIST(&_Vebo.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vebo.Contract.DEFAULTADMINROLE(&_Vebo.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Vebo.Contract.DEFAULTADMINROLE(&_Vebo.CallOpts)
}

// EXITREQUESTLIMITMANAGERROLE is a free data retrieval call binding the contract method 0xab53ac48.
//
// Solidity: function EXIT_REQUEST_LIMIT_MANAGER_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) EXITREQUESTLIMITMANAGERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "EXIT_REQUEST_LIMIT_MANAGER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXITREQUESTLIMITMANAGERROLE is a free data retrieval call binding the contract method 0xab53ac48.
//
// Solidity: function EXIT_REQUEST_LIMIT_MANAGER_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) EXITREQUESTLIMITMANAGERROLE() ([32]byte, error) {
	return _Vebo.Contract.EXITREQUESTLIMITMANAGERROLE(&_Vebo.CallOpts)
}

// EXITREQUESTLIMITMANAGERROLE is a free data retrieval call binding the contract method 0xab53ac48.
//
// Solidity: function EXIT_REQUEST_LIMIT_MANAGER_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) EXITREQUESTLIMITMANAGERROLE() ([32]byte, error) {
	return _Vebo.Contract.EXITREQUESTLIMITMANAGERROLE(&_Vebo.CallOpts)
}

// EXITTYPE is a free data retrieval call binding the contract method 0x06e41389.
//
// Solidity: function EXIT_TYPE() view returns(uint256)
func (_Vebo *VeboCaller) EXITTYPE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "EXIT_TYPE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITTYPE is a free data retrieval call binding the contract method 0x06e41389.
//
// Solidity: function EXIT_TYPE() view returns(uint256)
func (_Vebo *VeboSession) EXITTYPE() (*big.Int, error) {
	return _Vebo.Contract.EXITTYPE(&_Vebo.CallOpts)
}

// EXITTYPE is a free data retrieval call binding the contract method 0x06e41389.
//
// Solidity: function EXIT_TYPE() view returns(uint256)
func (_Vebo *VeboCallerSession) EXITTYPE() (*big.Int, error) {
	return _Vebo.Contract.EXITTYPE(&_Vebo.CallOpts)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_Vebo *VeboCaller) GENESISTIME(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "GENESIS_TIME")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_Vebo *VeboSession) GENESISTIME() (*big.Int, error) {
	return _Vebo.Contract.GENESISTIME(&_Vebo.CallOpts)
}

// GENESISTIME is a free data retrieval call binding the contract method 0xf2882461.
//
// Solidity: function GENESIS_TIME() view returns(uint256)
func (_Vebo *VeboCallerSession) GENESISTIME() (*big.Int, error) {
	return _Vebo.Contract.GENESISTIME(&_Vebo.CallOpts)
}

// MANAGECONSENSUSCONTRACTROLE is a free data retrieval call binding the contract method 0xad5cac4e.
//
// Solidity: function MANAGE_CONSENSUS_CONTRACT_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) MANAGECONSENSUSCONTRACTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "MANAGE_CONSENSUS_CONTRACT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGECONSENSUSCONTRACTROLE is a free data retrieval call binding the contract method 0xad5cac4e.
//
// Solidity: function MANAGE_CONSENSUS_CONTRACT_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) MANAGECONSENSUSCONTRACTROLE() ([32]byte, error) {
	return _Vebo.Contract.MANAGECONSENSUSCONTRACTROLE(&_Vebo.CallOpts)
}

// MANAGECONSENSUSCONTRACTROLE is a free data retrieval call binding the contract method 0xad5cac4e.
//
// Solidity: function MANAGE_CONSENSUS_CONTRACT_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) MANAGECONSENSUSCONTRACTROLE() ([32]byte, error) {
	return _Vebo.Contract.MANAGECONSENSUSCONTRACTROLE(&_Vebo.CallOpts)
}

// MANAGECONSENSUSVERSIONROLE is a free data retrieval call binding the contract method 0x9cc23c79.
//
// Solidity: function MANAGE_CONSENSUS_VERSION_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) MANAGECONSENSUSVERSIONROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "MANAGE_CONSENSUS_VERSION_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MANAGECONSENSUSVERSIONROLE is a free data retrieval call binding the contract method 0x9cc23c79.
//
// Solidity: function MANAGE_CONSENSUS_VERSION_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) MANAGECONSENSUSVERSIONROLE() ([32]byte, error) {
	return _Vebo.Contract.MANAGECONSENSUSVERSIONROLE(&_Vebo.CallOpts)
}

// MANAGECONSENSUSVERSIONROLE is a free data retrieval call binding the contract method 0x9cc23c79.
//
// Solidity: function MANAGE_CONSENSUS_VERSION_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) MANAGECONSENSUSVERSIONROLE() ([32]byte, error) {
	return _Vebo.Contract.MANAGECONSENSUSVERSIONROLE(&_Vebo.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Vebo *VeboCaller) PAUSEINFINITELY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "PAUSE_INFINITELY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Vebo *VeboSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Vebo.Contract.PAUSEINFINITELY(&_Vebo.CallOpts)
}

// PAUSEINFINITELY is a free data retrieval call binding the contract method 0xa302ee38.
//
// Solidity: function PAUSE_INFINITELY() view returns(uint256)
func (_Vebo *VeboCallerSession) PAUSEINFINITELY() (*big.Int, error) {
	return _Vebo.Contract.PAUSEINFINITELY(&_Vebo.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) PAUSEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "PAUSE_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) PAUSEROLE() ([32]byte, error) {
	return _Vebo.Contract.PAUSEROLE(&_Vebo.CallOpts)
}

// PAUSEROLE is a free data retrieval call binding the contract method 0x389ed267.
//
// Solidity: function PAUSE_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) PAUSEROLE() ([32]byte, error) {
	return _Vebo.Contract.PAUSEROLE(&_Vebo.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) RESUMEROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "RESUME_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) RESUMEROLE() ([32]byte, error) {
	return _Vebo.Contract.RESUMEROLE(&_Vebo.CallOpts)
}

// RESUMEROLE is a free data retrieval call binding the contract method 0x2de03aa1.
//
// Solidity: function RESUME_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) RESUMEROLE() ([32]byte, error) {
	return _Vebo.Contract.RESUMEROLE(&_Vebo.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_Vebo *VeboCaller) SECONDSPERSLOT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "SECONDS_PER_SLOT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_Vebo *VeboSession) SECONDSPERSLOT() (*big.Int, error) {
	return _Vebo.Contract.SECONDSPERSLOT(&_Vebo.CallOpts)
}

// SECONDSPERSLOT is a free data retrieval call binding the contract method 0x304b9071.
//
// Solidity: function SECONDS_PER_SLOT() view returns(uint256)
func (_Vebo *VeboCallerSession) SECONDSPERSLOT() (*big.Int, error) {
	return _Vebo.Contract.SECONDSPERSLOT(&_Vebo.CallOpts)
}

// SUBMITDATAROLE is a free data retrieval call binding the contract method 0x46e1f576.
//
// Solidity: function SUBMIT_DATA_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) SUBMITDATAROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "SUBMIT_DATA_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUBMITDATAROLE is a free data retrieval call binding the contract method 0x46e1f576.
//
// Solidity: function SUBMIT_DATA_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) SUBMITDATAROLE() ([32]byte, error) {
	return _Vebo.Contract.SUBMITDATAROLE(&_Vebo.CallOpts)
}

// SUBMITDATAROLE is a free data retrieval call binding the contract method 0x46e1f576.
//
// Solidity: function SUBMIT_DATA_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) SUBMITDATAROLE() ([32]byte, error) {
	return _Vebo.Contract.SUBMITDATAROLE(&_Vebo.CallOpts)
}

// SUBMITREPORTHASHROLE is a free data retrieval call binding the contract method 0xd072f014.
//
// Solidity: function SUBMIT_REPORT_HASH_ROLE() view returns(bytes32)
func (_Vebo *VeboCaller) SUBMITREPORTHASHROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "SUBMIT_REPORT_HASH_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SUBMITREPORTHASHROLE is a free data retrieval call binding the contract method 0xd072f014.
//
// Solidity: function SUBMIT_REPORT_HASH_ROLE() view returns(bytes32)
func (_Vebo *VeboSession) SUBMITREPORTHASHROLE() ([32]byte, error) {
	return _Vebo.Contract.SUBMITREPORTHASHROLE(&_Vebo.CallOpts)
}

// SUBMITREPORTHASHROLE is a free data retrieval call binding the contract method 0xd072f014.
//
// Solidity: function SUBMIT_REPORT_HASH_ROLE() view returns(bytes32)
func (_Vebo *VeboCallerSession) SUBMITREPORTHASHROLE() ([32]byte, error) {
	return _Vebo.Contract.SUBMITREPORTHASHROLE(&_Vebo.CallOpts)
}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_Vebo *VeboCaller) GetConsensusContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getConsensusContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_Vebo *VeboSession) GetConsensusContract() (common.Address, error) {
	return _Vebo.Contract.GetConsensusContract(&_Vebo.CallOpts)
}

// GetConsensusContract is a free data retrieval call binding the contract method 0x8f55b571.
//
// Solidity: function getConsensusContract() view returns(address)
func (_Vebo *VeboCallerSession) GetConsensusContract() (common.Address, error) {
	return _Vebo.Contract.GetConsensusContract(&_Vebo.CallOpts)
}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_Vebo *VeboCaller) GetConsensusReport(opts *bind.CallOpts) (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getConsensusReport")

	outstruct := new(struct {
		Hash                   [32]byte
		RefSlot                *big.Int
		ProcessingDeadlineTime *big.Int
		ProcessingStarted      bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Hash = *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)
	outstruct.RefSlot = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProcessingDeadlineTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ProcessingStarted = *abi.ConvertType(out[3], new(bool)).(*bool)

	return *outstruct, err

}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_Vebo *VeboSession) GetConsensusReport() (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	return _Vebo.Contract.GetConsensusReport(&_Vebo.CallOpts)
}

// GetConsensusReport is a free data retrieval call binding the contract method 0x60d64d38.
//
// Solidity: function getConsensusReport() view returns(bytes32 hash, uint256 refSlot, uint256 processingDeadlineTime, bool processingStarted)
func (_Vebo *VeboCallerSession) GetConsensusReport() (struct {
	Hash                   [32]byte
	RefSlot                *big.Int
	ProcessingDeadlineTime *big.Int
	ProcessingStarted      bool
}, error) {
	return _Vebo.Contract.GetConsensusReport(&_Vebo.CallOpts)
}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_Vebo *VeboCaller) GetConsensusVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getConsensusVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_Vebo *VeboSession) GetConsensusVersion() (*big.Int, error) {
	return _Vebo.Contract.GetConsensusVersion(&_Vebo.CallOpts)
}

// GetConsensusVersion is a free data retrieval call binding the contract method 0x5be20425.
//
// Solidity: function getConsensusVersion() view returns(uint256)
func (_Vebo *VeboCallerSession) GetConsensusVersion() (*big.Int, error) {
	return _Vebo.Contract.GetConsensusVersion(&_Vebo.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Vebo *VeboCaller) GetContractVersion(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getContractVersion")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Vebo *VeboSession) GetContractVersion() (*big.Int, error) {
	return _Vebo.Contract.GetContractVersion(&_Vebo.CallOpts)
}

// GetContractVersion is a free data retrieval call binding the contract method 0x8aa10435.
//
// Solidity: function getContractVersion() view returns(uint256)
func (_Vebo *VeboCallerSession) GetContractVersion() (*big.Int, error) {
	return _Vebo.Contract.GetContractVersion(&_Vebo.CallOpts)
}

// GetDeliveryTimestamp is a free data retrieval call binding the contract method 0xa52289bf.
//
// Solidity: function getDeliveryTimestamp(bytes32 exitRequestsHash) view returns(uint256 deliveryDateTimestamp)
func (_Vebo *VeboCaller) GetDeliveryTimestamp(opts *bind.CallOpts, exitRequestsHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getDeliveryTimestamp", exitRequestsHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDeliveryTimestamp is a free data retrieval call binding the contract method 0xa52289bf.
//
// Solidity: function getDeliveryTimestamp(bytes32 exitRequestsHash) view returns(uint256 deliveryDateTimestamp)
func (_Vebo *VeboSession) GetDeliveryTimestamp(exitRequestsHash [32]byte) (*big.Int, error) {
	return _Vebo.Contract.GetDeliveryTimestamp(&_Vebo.CallOpts, exitRequestsHash)
}

// GetDeliveryTimestamp is a free data retrieval call binding the contract method 0xa52289bf.
//
// Solidity: function getDeliveryTimestamp(bytes32 exitRequestsHash) view returns(uint256 deliveryDateTimestamp)
func (_Vebo *VeboCallerSession) GetDeliveryTimestamp(exitRequestsHash [32]byte) (*big.Int, error) {
	return _Vebo.Contract.GetDeliveryTimestamp(&_Vebo.CallOpts, exitRequestsHash)
}

// GetExitRequestLimitFullInfo is a free data retrieval call binding the contract method 0xb6b764b2.
//
// Solidity: function getExitRequestLimitFullInfo() view returns(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec, uint256 prevExitRequestsLimit, uint256 currentExitRequestsLimit)
func (_Vebo *VeboCaller) GetExitRequestLimitFullInfo(opts *bind.CallOpts) (struct {
	MaxExitRequestsLimit     *big.Int
	ExitsPerFrame            *big.Int
	FrameDurationInSec       *big.Int
	PrevExitRequestsLimit    *big.Int
	CurrentExitRequestsLimit *big.Int
}, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getExitRequestLimitFullInfo")

	outstruct := new(struct {
		MaxExitRequestsLimit     *big.Int
		ExitsPerFrame            *big.Int
		FrameDurationInSec       *big.Int
		PrevExitRequestsLimit    *big.Int
		CurrentExitRequestsLimit *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.MaxExitRequestsLimit = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ExitsPerFrame = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.FrameDurationInSec = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PrevExitRequestsLimit = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.CurrentExitRequestsLimit = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetExitRequestLimitFullInfo is a free data retrieval call binding the contract method 0xb6b764b2.
//
// Solidity: function getExitRequestLimitFullInfo() view returns(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec, uint256 prevExitRequestsLimit, uint256 currentExitRequestsLimit)
func (_Vebo *VeboSession) GetExitRequestLimitFullInfo() (struct {
	MaxExitRequestsLimit     *big.Int
	ExitsPerFrame            *big.Int
	FrameDurationInSec       *big.Int
	PrevExitRequestsLimit    *big.Int
	CurrentExitRequestsLimit *big.Int
}, error) {
	return _Vebo.Contract.GetExitRequestLimitFullInfo(&_Vebo.CallOpts)
}

// GetExitRequestLimitFullInfo is a free data retrieval call binding the contract method 0xb6b764b2.
//
// Solidity: function getExitRequestLimitFullInfo() view returns(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec, uint256 prevExitRequestsLimit, uint256 currentExitRequestsLimit)
func (_Vebo *VeboCallerSession) GetExitRequestLimitFullInfo() (struct {
	MaxExitRequestsLimit     *big.Int
	ExitsPerFrame            *big.Int
	FrameDurationInSec       *big.Int
	PrevExitRequestsLimit    *big.Int
	CurrentExitRequestsLimit *big.Int
}, error) {
	return _Vebo.Contract.GetExitRequestLimitFullInfo(&_Vebo.CallOpts)
}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_Vebo *VeboCaller) GetLastProcessingRefSlot(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getLastProcessingRefSlot")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_Vebo *VeboSession) GetLastProcessingRefSlot() (*big.Int, error) {
	return _Vebo.Contract.GetLastProcessingRefSlot(&_Vebo.CallOpts)
}

// GetLastProcessingRefSlot is a free data retrieval call binding the contract method 0x3584d59c.
//
// Solidity: function getLastProcessingRefSlot() view returns(uint256)
func (_Vebo *VeboCallerSession) GetLastProcessingRefSlot() (*big.Int, error) {
	return _Vebo.Contract.GetLastProcessingRefSlot(&_Vebo.CallOpts)
}

// GetMaxValidatorsPerReport is a free data retrieval call binding the contract method 0xc1f665bc.
//
// Solidity: function getMaxValidatorsPerReport() view returns(uint256)
func (_Vebo *VeboCaller) GetMaxValidatorsPerReport(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getMaxValidatorsPerReport")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxValidatorsPerReport is a free data retrieval call binding the contract method 0xc1f665bc.
//
// Solidity: function getMaxValidatorsPerReport() view returns(uint256)
func (_Vebo *VeboSession) GetMaxValidatorsPerReport() (*big.Int, error) {
	return _Vebo.Contract.GetMaxValidatorsPerReport(&_Vebo.CallOpts)
}

// GetMaxValidatorsPerReport is a free data retrieval call binding the contract method 0xc1f665bc.
//
// Solidity: function getMaxValidatorsPerReport() view returns(uint256)
func (_Vebo *VeboCallerSession) GetMaxValidatorsPerReport() (*big.Int, error) {
	return _Vebo.Contract.GetMaxValidatorsPerReport(&_Vebo.CallOpts)
}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256,uint256,uint256) result)
func (_Vebo *VeboCaller) GetProcessingState(opts *bind.CallOpts) (ValidatorsExitBusOracleProcessingState, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getProcessingState")

	if err != nil {
		return *new(ValidatorsExitBusOracleProcessingState), err
	}

	out0 := *abi.ConvertType(out[0], new(ValidatorsExitBusOracleProcessingState)).(*ValidatorsExitBusOracleProcessingState)

	return out0, err

}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256,uint256,uint256) result)
func (_Vebo *VeboSession) GetProcessingState() (ValidatorsExitBusOracleProcessingState, error) {
	return _Vebo.Contract.GetProcessingState(&_Vebo.CallOpts)
}

// GetProcessingState is a free data retrieval call binding the contract method 0x8f7797c2.
//
// Solidity: function getProcessingState() view returns((uint256,uint256,bytes32,bool,uint256,uint256,uint256) result)
func (_Vebo *VeboCallerSession) GetProcessingState() (ValidatorsExitBusOracleProcessingState, error) {
	return _Vebo.Contract.GetProcessingState(&_Vebo.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Vebo *VeboCaller) GetResumeSinceTimestamp(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getResumeSinceTimestamp")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Vebo *VeboSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Vebo.Contract.GetResumeSinceTimestamp(&_Vebo.CallOpts)
}

// GetResumeSinceTimestamp is a free data retrieval call binding the contract method 0x589ff76c.
//
// Solidity: function getResumeSinceTimestamp() view returns(uint256)
func (_Vebo *VeboCallerSession) GetResumeSinceTimestamp() (*big.Int, error) {
	return _Vebo.Contract.GetResumeSinceTimestamp(&_Vebo.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vebo *VeboCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vebo *VeboSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vebo.Contract.GetRoleAdmin(&_Vebo.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Vebo *VeboCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Vebo.Contract.GetRoleAdmin(&_Vebo.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Vebo *VeboCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Vebo *VeboSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Vebo.Contract.GetRoleMember(&_Vebo.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Vebo *VeboCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Vebo.Contract.GetRoleMember(&_Vebo.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Vebo *VeboCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Vebo *VeboSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Vebo.Contract.GetRoleMemberCount(&_Vebo.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Vebo *VeboCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Vebo.Contract.GetRoleMemberCount(&_Vebo.CallOpts, role)
}

// GetTotalRequestsProcessed is a free data retrieval call binding the contract method 0xe2793e72.
//
// Solidity: function getTotalRequestsProcessed() view returns(uint256)
func (_Vebo *VeboCaller) GetTotalRequestsProcessed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "getTotalRequestsProcessed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalRequestsProcessed is a free data retrieval call binding the contract method 0xe2793e72.
//
// Solidity: function getTotalRequestsProcessed() view returns(uint256)
func (_Vebo *VeboSession) GetTotalRequestsProcessed() (*big.Int, error) {
	return _Vebo.Contract.GetTotalRequestsProcessed(&_Vebo.CallOpts)
}

// GetTotalRequestsProcessed is a free data retrieval call binding the contract method 0xe2793e72.
//
// Solidity: function getTotalRequestsProcessed() view returns(uint256)
func (_Vebo *VeboCallerSession) GetTotalRequestsProcessed() (*big.Int, error) {
	return _Vebo.Contract.GetTotalRequestsProcessed(&_Vebo.CallOpts)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vebo *VeboCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vebo *VeboSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vebo.Contract.HasRole(&_Vebo.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Vebo *VeboCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Vebo.Contract.HasRole(&_Vebo.CallOpts, role, account)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Vebo *VeboCaller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Vebo *VeboSession) IsPaused() (bool, error) {
	return _Vebo.Contract.IsPaused(&_Vebo.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Vebo *VeboCallerSession) IsPaused() (bool, error) {
	return _Vebo.Contract.IsPaused(&_Vebo.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vebo *VeboCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vebo *VeboSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vebo.Contract.SupportsInterface(&_Vebo.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Vebo *VeboCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Vebo.Contract.SupportsInterface(&_Vebo.CallOpts, interfaceId)
}

// UnpackExitRequest is a free data retrieval call binding the contract method 0x7dad759d.
//
// Solidity: function unpackExitRequest(bytes exitRequests, uint256 dataFormat, uint256 index) pure returns(bytes pubkey, uint256 nodeOpId, uint256 moduleId, uint256 valIndex)
func (_Vebo *VeboCaller) UnpackExitRequest(opts *bind.CallOpts, exitRequests []byte, dataFormat *big.Int, index *big.Int) (struct {
	Pubkey   []byte
	NodeOpId *big.Int
	ModuleId *big.Int
	ValIndex *big.Int
}, error) {
	var out []interface{}
	err := _Vebo.contract.Call(opts, &out, "unpackExitRequest", exitRequests, dataFormat, index)

	outstruct := new(struct {
		Pubkey   []byte
		NodeOpId *big.Int
		ModuleId *big.Int
		ValIndex *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Pubkey = *abi.ConvertType(out[0], new([]byte)).(*[]byte)
	outstruct.NodeOpId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ModuleId = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ValIndex = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// UnpackExitRequest is a free data retrieval call binding the contract method 0x7dad759d.
//
// Solidity: function unpackExitRequest(bytes exitRequests, uint256 dataFormat, uint256 index) pure returns(bytes pubkey, uint256 nodeOpId, uint256 moduleId, uint256 valIndex)
func (_Vebo *VeboSession) UnpackExitRequest(exitRequests []byte, dataFormat *big.Int, index *big.Int) (struct {
	Pubkey   []byte
	NodeOpId *big.Int
	ModuleId *big.Int
	ValIndex *big.Int
}, error) {
	return _Vebo.Contract.UnpackExitRequest(&_Vebo.CallOpts, exitRequests, dataFormat, index)
}

// UnpackExitRequest is a free data retrieval call binding the contract method 0x7dad759d.
//
// Solidity: function unpackExitRequest(bytes exitRequests, uint256 dataFormat, uint256 index) pure returns(bytes pubkey, uint256 nodeOpId, uint256 moduleId, uint256 valIndex)
func (_Vebo *VeboCallerSession) UnpackExitRequest(exitRequests []byte, dataFormat *big.Int, index *big.Int) (struct {
	Pubkey   []byte
	NodeOpId *big.Int
	ModuleId *big.Int
	ValIndex *big.Int
}, error) {
	return _Vebo.Contract.UnpackExitRequest(&_Vebo.CallOpts, exitRequests, dataFormat, index)
}

// DiscardConsensusReport is a paid mutator transaction binding the contract method 0xd4381217.
//
// Solidity: function discardConsensusReport(uint256 refSlot) returns()
func (_Vebo *VeboTransactor) DiscardConsensusReport(opts *bind.TransactOpts, refSlot *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "discardConsensusReport", refSlot)
}

// DiscardConsensusReport is a paid mutator transaction binding the contract method 0xd4381217.
//
// Solidity: function discardConsensusReport(uint256 refSlot) returns()
func (_Vebo *VeboSession) DiscardConsensusReport(refSlot *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.DiscardConsensusReport(&_Vebo.TransactOpts, refSlot)
}

// DiscardConsensusReport is a paid mutator transaction binding the contract method 0xd4381217.
//
// Solidity: function discardConsensusReport(uint256 refSlot) returns()
func (_Vebo *VeboTransactorSession) DiscardConsensusReport(refSlot *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.DiscardConsensusReport(&_Vebo.TransactOpts, refSlot)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0xff406d19.
//
// Solidity: function finalizeUpgrade_v2(uint256 maxValidatorsPerReport, uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Vebo *VeboTransactor) FinalizeUpgradeV2(opts *bind.TransactOpts, maxValidatorsPerReport *big.Int, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "finalizeUpgrade_v2", maxValidatorsPerReport, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0xff406d19.
//
// Solidity: function finalizeUpgrade_v2(uint256 maxValidatorsPerReport, uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Vebo *VeboSession) FinalizeUpgradeV2(maxValidatorsPerReport *big.Int, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.FinalizeUpgradeV2(&_Vebo.TransactOpts, maxValidatorsPerReport, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// FinalizeUpgradeV2 is a paid mutator transaction binding the contract method 0xff406d19.
//
// Solidity: function finalizeUpgrade_v2(uint256 maxValidatorsPerReport, uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Vebo *VeboTransactorSession) FinalizeUpgradeV2(maxValidatorsPerReport *big.Int, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.FinalizeUpgradeV2(&_Vebo.TransactOpts, maxValidatorsPerReport, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vebo *VeboTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vebo *VeboSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.GrantRole(&_Vebo.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Vebo *VeboTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.GrantRole(&_Vebo.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x8ba796af.
//
// Solidity: function initialize(address admin, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, uint256 maxValidatorsPerRequest, uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Vebo *VeboTransactor) Initialize(opts *bind.TransactOpts, admin common.Address, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, maxValidatorsPerRequest *big.Int, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "initialize", admin, consensusContract, consensusVersion, lastProcessingRefSlot, maxValidatorsPerRequest, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// Initialize is a paid mutator transaction binding the contract method 0x8ba796af.
//
// Solidity: function initialize(address admin, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, uint256 maxValidatorsPerRequest, uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Vebo *VeboSession) Initialize(admin common.Address, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, maxValidatorsPerRequest *big.Int, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.Initialize(&_Vebo.TransactOpts, admin, consensusContract, consensusVersion, lastProcessingRefSlot, maxValidatorsPerRequest, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// Initialize is a paid mutator transaction binding the contract method 0x8ba796af.
//
// Solidity: function initialize(address admin, address consensusContract, uint256 consensusVersion, uint256 lastProcessingRefSlot, uint256 maxValidatorsPerRequest, uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Vebo *VeboTransactorSession) Initialize(admin common.Address, consensusContract common.Address, consensusVersion *big.Int, lastProcessingRefSlot *big.Int, maxValidatorsPerRequest *big.Int, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.Initialize(&_Vebo.TransactOpts, admin, consensusContract, consensusVersion, lastProcessingRefSlot, maxValidatorsPerRequest, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 _duration) returns()
func (_Vebo *VeboTransactor) PauseFor(opts *bind.TransactOpts, _duration *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "pauseFor", _duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 _duration) returns()
func (_Vebo *VeboSession) PauseFor(_duration *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.PauseFor(&_Vebo.TransactOpts, _duration)
}

// PauseFor is a paid mutator transaction binding the contract method 0xf3f449c7.
//
// Solidity: function pauseFor(uint256 _duration) returns()
func (_Vebo *VeboTransactorSession) PauseFor(_duration *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.PauseFor(&_Vebo.TransactOpts, _duration)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 _pauseUntilInclusive) returns()
func (_Vebo *VeboTransactor) PauseUntil(opts *bind.TransactOpts, _pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "pauseUntil", _pauseUntilInclusive)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 _pauseUntilInclusive) returns()
func (_Vebo *VeboSession) PauseUntil(_pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.PauseUntil(&_Vebo.TransactOpts, _pauseUntilInclusive)
}

// PauseUntil is a paid mutator transaction binding the contract method 0xabe9cfc8.
//
// Solidity: function pauseUntil(uint256 _pauseUntilInclusive) returns()
func (_Vebo *VeboTransactorSession) PauseUntil(_pauseUntilInclusive *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.PauseUntil(&_Vebo.TransactOpts, _pauseUntilInclusive)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vebo *VeboTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vebo *VeboSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.RenounceRole(&_Vebo.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Vebo *VeboTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.RenounceRole(&_Vebo.TransactOpts, role, account)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Vebo *VeboTransactor) Resume(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "resume")
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Vebo *VeboSession) Resume() (*types.Transaction, error) {
	return _Vebo.Contract.Resume(&_Vebo.TransactOpts)
}

// Resume is a paid mutator transaction binding the contract method 0x046f7da2.
//
// Solidity: function resume() returns()
func (_Vebo *VeboTransactorSession) Resume() (*types.Transaction, error) {
	return _Vebo.Contract.Resume(&_Vebo.TransactOpts)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vebo *VeboTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vebo *VeboSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.RevokeRole(&_Vebo.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Vebo *VeboTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.RevokeRole(&_Vebo.TransactOpts, role, account)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_Vebo *VeboTransactor) SetConsensusContract(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "setConsensusContract", addr)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_Vebo *VeboSession) SetConsensusContract(addr common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.SetConsensusContract(&_Vebo.TransactOpts, addr)
}

// SetConsensusContract is a paid mutator transaction binding the contract method 0xc469c307.
//
// Solidity: function setConsensusContract(address addr) returns()
func (_Vebo *VeboTransactorSession) SetConsensusContract(addr common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.SetConsensusContract(&_Vebo.TransactOpts, addr)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_Vebo *VeboTransactor) SetConsensusVersion(opts *bind.TransactOpts, version *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "setConsensusVersion", version)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_Vebo *VeboSession) SetConsensusVersion(version *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SetConsensusVersion(&_Vebo.TransactOpts, version)
}

// SetConsensusVersion is a paid mutator transaction binding the contract method 0x8d591474.
//
// Solidity: function setConsensusVersion(uint256 version) returns()
func (_Vebo *VeboTransactorSession) SetConsensusVersion(version *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SetConsensusVersion(&_Vebo.TransactOpts, version)
}

// SetExitRequestLimit is a paid mutator transaction binding the contract method 0x56254a97.
//
// Solidity: function setExitRequestLimit(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Vebo *VeboTransactor) SetExitRequestLimit(opts *bind.TransactOpts, maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "setExitRequestLimit", maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// SetExitRequestLimit is a paid mutator transaction binding the contract method 0x56254a97.
//
// Solidity: function setExitRequestLimit(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Vebo *VeboSession) SetExitRequestLimit(maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SetExitRequestLimit(&_Vebo.TransactOpts, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// SetExitRequestLimit is a paid mutator transaction binding the contract method 0x56254a97.
//
// Solidity: function setExitRequestLimit(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec) returns()
func (_Vebo *VeboTransactorSession) SetExitRequestLimit(maxExitRequestsLimit *big.Int, exitsPerFrame *big.Int, frameDurationInSec *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SetExitRequestLimit(&_Vebo.TransactOpts, maxExitRequestsLimit, exitsPerFrame, frameDurationInSec)
}

// SetMaxValidatorsPerReport is a paid mutator transaction binding the contract method 0x6f2c322d.
//
// Solidity: function setMaxValidatorsPerReport(uint256 maxRequests) returns()
func (_Vebo *VeboTransactor) SetMaxValidatorsPerReport(opts *bind.TransactOpts, maxRequests *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "setMaxValidatorsPerReport", maxRequests)
}

// SetMaxValidatorsPerReport is a paid mutator transaction binding the contract method 0x6f2c322d.
//
// Solidity: function setMaxValidatorsPerReport(uint256 maxRequests) returns()
func (_Vebo *VeboSession) SetMaxValidatorsPerReport(maxRequests *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SetMaxValidatorsPerReport(&_Vebo.TransactOpts, maxRequests)
}

// SetMaxValidatorsPerReport is a paid mutator transaction binding the contract method 0x6f2c322d.
//
// Solidity: function setMaxValidatorsPerReport(uint256 maxRequests) returns()
func (_Vebo *VeboTransactorSession) SetMaxValidatorsPerReport(maxRequests *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SetMaxValidatorsPerReport(&_Vebo.TransactOpts, maxRequests)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x063f36ad.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline) returns()
func (_Vebo *VeboTransactor) SubmitConsensusReport(opts *bind.TransactOpts, reportHash [32]byte, refSlot *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "submitConsensusReport", reportHash, refSlot, deadline)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x063f36ad.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline) returns()
func (_Vebo *VeboSession) SubmitConsensusReport(reportHash [32]byte, refSlot *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SubmitConsensusReport(&_Vebo.TransactOpts, reportHash, refSlot, deadline)
}

// SubmitConsensusReport is a paid mutator transaction binding the contract method 0x063f36ad.
//
// Solidity: function submitConsensusReport(bytes32 reportHash, uint256 refSlot, uint256 deadline) returns()
func (_Vebo *VeboTransactorSession) SubmitConsensusReport(reportHash [32]byte, refSlot *big.Int, deadline *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SubmitConsensusReport(&_Vebo.TransactOpts, reportHash, refSlot, deadline)
}

// SubmitExitRequestsData is a paid mutator transaction binding the contract method 0xb8fe0ad0.
//
// Solidity: function submitExitRequestsData((bytes,uint256) request) returns()
func (_Vebo *VeboTransactor) SubmitExitRequestsData(opts *bind.TransactOpts, request ValidatorsExitBusExitRequestsData) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "submitExitRequestsData", request)
}

// SubmitExitRequestsData is a paid mutator transaction binding the contract method 0xb8fe0ad0.
//
// Solidity: function submitExitRequestsData((bytes,uint256) request) returns()
func (_Vebo *VeboSession) SubmitExitRequestsData(request ValidatorsExitBusExitRequestsData) (*types.Transaction, error) {
	return _Vebo.Contract.SubmitExitRequestsData(&_Vebo.TransactOpts, request)
}

// SubmitExitRequestsData is a paid mutator transaction binding the contract method 0xb8fe0ad0.
//
// Solidity: function submitExitRequestsData((bytes,uint256) request) returns()
func (_Vebo *VeboTransactorSession) SubmitExitRequestsData(request ValidatorsExitBusExitRequestsData) (*types.Transaction, error) {
	return _Vebo.Contract.SubmitExitRequestsData(&_Vebo.TransactOpts, request)
}

// SubmitExitRequestsHash is a paid mutator transaction binding the contract method 0xb1b19f57.
//
// Solidity: function submitExitRequestsHash(bytes32 exitRequestsHash) returns()
func (_Vebo *VeboTransactor) SubmitExitRequestsHash(opts *bind.TransactOpts, exitRequestsHash [32]byte) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "submitExitRequestsHash", exitRequestsHash)
}

// SubmitExitRequestsHash is a paid mutator transaction binding the contract method 0xb1b19f57.
//
// Solidity: function submitExitRequestsHash(bytes32 exitRequestsHash) returns()
func (_Vebo *VeboSession) SubmitExitRequestsHash(exitRequestsHash [32]byte) (*types.Transaction, error) {
	return _Vebo.Contract.SubmitExitRequestsHash(&_Vebo.TransactOpts, exitRequestsHash)
}

// SubmitExitRequestsHash is a paid mutator transaction binding the contract method 0xb1b19f57.
//
// Solidity: function submitExitRequestsHash(bytes32 exitRequestsHash) returns()
func (_Vebo *VeboTransactorSession) SubmitExitRequestsHash(exitRequestsHash [32]byte) (*types.Transaction, error) {
	return _Vebo.Contract.SubmitExitRequestsHash(&_Vebo.TransactOpts, exitRequestsHash)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x294492c8.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,bytes) data, uint256 contractVersion) returns()
func (_Vebo *VeboTransactor) SubmitReportData(opts *bind.TransactOpts, data ValidatorsExitBusOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "submitReportData", data, contractVersion)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x294492c8.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,bytes) data, uint256 contractVersion) returns()
func (_Vebo *VeboSession) SubmitReportData(data ValidatorsExitBusOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SubmitReportData(&_Vebo.TransactOpts, data, contractVersion)
}

// SubmitReportData is a paid mutator transaction binding the contract method 0x294492c8.
//
// Solidity: function submitReportData((uint256,uint256,uint256,uint256,bytes) data, uint256 contractVersion) returns()
func (_Vebo *VeboTransactorSession) SubmitReportData(data ValidatorsExitBusOracleReportData, contractVersion *big.Int) (*types.Transaction, error) {
	return _Vebo.Contract.SubmitReportData(&_Vebo.TransactOpts, data, contractVersion)
}

// TriggerExits is a paid mutator transaction binding the contract method 0xa2ab7065.
//
// Solidity: function triggerExits((bytes,uint256) exitsData, uint256[] exitDataIndexes, address refundRecipient) payable returns()
func (_Vebo *VeboTransactor) TriggerExits(opts *bind.TransactOpts, exitsData ValidatorsExitBusExitRequestsData, exitDataIndexes []*big.Int, refundRecipient common.Address) (*types.Transaction, error) {
	return _Vebo.contract.Transact(opts, "triggerExits", exitsData, exitDataIndexes, refundRecipient)
}

// TriggerExits is a paid mutator transaction binding the contract method 0xa2ab7065.
//
// Solidity: function triggerExits((bytes,uint256) exitsData, uint256[] exitDataIndexes, address refundRecipient) payable returns()
func (_Vebo *VeboSession) TriggerExits(exitsData ValidatorsExitBusExitRequestsData, exitDataIndexes []*big.Int, refundRecipient common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.TriggerExits(&_Vebo.TransactOpts, exitsData, exitDataIndexes, refundRecipient)
}

// TriggerExits is a paid mutator transaction binding the contract method 0xa2ab7065.
//
// Solidity: function triggerExits((bytes,uint256) exitsData, uint256[] exitDataIndexes, address refundRecipient) payable returns()
func (_Vebo *VeboTransactorSession) TriggerExits(exitsData ValidatorsExitBusExitRequestsData, exitDataIndexes []*big.Int, refundRecipient common.Address) (*types.Transaction, error) {
	return _Vebo.Contract.TriggerExits(&_Vebo.TransactOpts, exitsData, exitDataIndexes, refundRecipient)
}

// VeboConsensusHashContractSetIterator is returned from FilterConsensusHashContractSet and is used to iterate over the raw logs and unpacked data for ConsensusHashContractSet events raised by the Vebo contract.
type VeboConsensusHashContractSetIterator struct {
	Event *VeboConsensusHashContractSet // Event containing the contract specifics and raw log

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
func (it *VeboConsensusHashContractSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboConsensusHashContractSet)
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
		it.Event = new(VeboConsensusHashContractSet)
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
func (it *VeboConsensusHashContractSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboConsensusHashContractSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboConsensusHashContractSet represents a ConsensusHashContractSet event raised by the Vebo contract.
type VeboConsensusHashContractSet struct {
	Addr     common.Address
	PrevAddr common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterConsensusHashContractSet is a free log retrieval operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_Vebo *VeboFilterer) FilterConsensusHashContractSet(opts *bind.FilterOpts, addr []common.Address, prevAddr []common.Address) (*VeboConsensusHashContractSetIterator, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var prevAddrRule []interface{}
	for _, prevAddrItem := range prevAddr {
		prevAddrRule = append(prevAddrRule, prevAddrItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ConsensusHashContractSet", addrRule, prevAddrRule)
	if err != nil {
		return nil, err
	}
	return &VeboConsensusHashContractSetIterator{contract: _Vebo.contract, event: "ConsensusHashContractSet", logs: logs, sub: sub}, nil
}

// WatchConsensusHashContractSet is a free log subscription operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_Vebo *VeboFilterer) WatchConsensusHashContractSet(opts *bind.WatchOpts, sink chan<- *VeboConsensusHashContractSet, addr []common.Address, prevAddr []common.Address) (event.Subscription, error) {

	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}
	var prevAddrRule []interface{}
	for _, prevAddrItem := range prevAddr {
		prevAddrRule = append(prevAddrRule, prevAddrItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ConsensusHashContractSet", addrRule, prevAddrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboConsensusHashContractSet)
				if err := _Vebo.contract.UnpackLog(event, "ConsensusHashContractSet", log); err != nil {
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

// ParseConsensusHashContractSet is a log parse operation binding the contract event 0x25421480fb7f52d18947876279a213696b58d7e0e5416ce5e2c9f9942661c34c.
//
// Solidity: event ConsensusHashContractSet(address indexed addr, address indexed prevAddr)
func (_Vebo *VeboFilterer) ParseConsensusHashContractSet(log types.Log) (*VeboConsensusHashContractSet, error) {
	event := new(VeboConsensusHashContractSet)
	if err := _Vebo.contract.UnpackLog(event, "ConsensusHashContractSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboConsensusVersionSetIterator is returned from FilterConsensusVersionSet and is used to iterate over the raw logs and unpacked data for ConsensusVersionSet events raised by the Vebo contract.
type VeboConsensusVersionSetIterator struct {
	Event *VeboConsensusVersionSet // Event containing the contract specifics and raw log

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
func (it *VeboConsensusVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboConsensusVersionSet)
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
		it.Event = new(VeboConsensusVersionSet)
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
func (it *VeboConsensusVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboConsensusVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboConsensusVersionSet represents a ConsensusVersionSet event raised by the Vebo contract.
type VeboConsensusVersionSet struct {
	Version     *big.Int
	PrevVersion *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterConsensusVersionSet is a free log retrieval operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_Vebo *VeboFilterer) FilterConsensusVersionSet(opts *bind.FilterOpts, version []*big.Int, prevVersion []*big.Int) (*VeboConsensusVersionSetIterator, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var prevVersionRule []interface{}
	for _, prevVersionItem := range prevVersion {
		prevVersionRule = append(prevVersionRule, prevVersionItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ConsensusVersionSet", versionRule, prevVersionRule)
	if err != nil {
		return nil, err
	}
	return &VeboConsensusVersionSetIterator{contract: _Vebo.contract, event: "ConsensusVersionSet", logs: logs, sub: sub}, nil
}

// WatchConsensusVersionSet is a free log subscription operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_Vebo *VeboFilterer) WatchConsensusVersionSet(opts *bind.WatchOpts, sink chan<- *VeboConsensusVersionSet, version []*big.Int, prevVersion []*big.Int) (event.Subscription, error) {

	var versionRule []interface{}
	for _, versionItem := range version {
		versionRule = append(versionRule, versionItem)
	}
	var prevVersionRule []interface{}
	for _, prevVersionItem := range prevVersion {
		prevVersionRule = append(prevVersionRule, prevVersionItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ConsensusVersionSet", versionRule, prevVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboConsensusVersionSet)
				if err := _Vebo.contract.UnpackLog(event, "ConsensusVersionSet", log); err != nil {
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

// ParseConsensusVersionSet is a log parse operation binding the contract event 0xfa5304972d4ec3e3207f0bbf91155a49d0dfa62488f9529403a2a49e4b29a895.
//
// Solidity: event ConsensusVersionSet(uint256 indexed version, uint256 indexed prevVersion)
func (_Vebo *VeboFilterer) ParseConsensusVersionSet(log types.Log) (*VeboConsensusVersionSet, error) {
	event := new(VeboConsensusVersionSet)
	if err := _Vebo.contract.UnpackLog(event, "ConsensusVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboContractVersionSetIterator is returned from FilterContractVersionSet and is used to iterate over the raw logs and unpacked data for ContractVersionSet events raised by the Vebo contract.
type VeboContractVersionSetIterator struct {
	Event *VeboContractVersionSet // Event containing the contract specifics and raw log

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
func (it *VeboContractVersionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboContractVersionSet)
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
		it.Event = new(VeboContractVersionSet)
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
func (it *VeboContractVersionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboContractVersionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboContractVersionSet represents a ContractVersionSet event raised by the Vebo contract.
type VeboContractVersionSet struct {
	Version *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterContractVersionSet is a free log retrieval operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Vebo *VeboFilterer) FilterContractVersionSet(opts *bind.FilterOpts) (*VeboContractVersionSetIterator, error) {

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return &VeboContractVersionSetIterator{contract: _Vebo.contract, event: "ContractVersionSet", logs: logs, sub: sub}, nil
}

// WatchContractVersionSet is a free log subscription operation binding the contract event 0xfddcded6b4f4730c226821172046b48372d3cd963c159701ae1b7c3bcac541bb.
//
// Solidity: event ContractVersionSet(uint256 version)
func (_Vebo *VeboFilterer) WatchContractVersionSet(opts *bind.WatchOpts, sink chan<- *VeboContractVersionSet) (event.Subscription, error) {

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ContractVersionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboContractVersionSet)
				if err := _Vebo.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
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
func (_Vebo *VeboFilterer) ParseContractVersionSet(log types.Log) (*VeboContractVersionSet, error) {
	event := new(VeboContractVersionSet)
	if err := _Vebo.contract.UnpackLog(event, "ContractVersionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboExitDataProcessingIterator is returned from FilterExitDataProcessing and is used to iterate over the raw logs and unpacked data for ExitDataProcessing events raised by the Vebo contract.
type VeboExitDataProcessingIterator struct {
	Event *VeboExitDataProcessing // Event containing the contract specifics and raw log

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
func (it *VeboExitDataProcessingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboExitDataProcessing)
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
		it.Event = new(VeboExitDataProcessing)
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
func (it *VeboExitDataProcessingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboExitDataProcessingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboExitDataProcessing represents a ExitDataProcessing event raised by the Vebo contract.
type VeboExitDataProcessing struct {
	ExitRequestsHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterExitDataProcessing is a free log retrieval operation binding the contract event 0x01b8de053572c3c2104259b555c485ccac8017196b3471e8483b7e96f071608a.
//
// Solidity: event ExitDataProcessing(bytes32 exitRequestsHash)
func (_Vebo *VeboFilterer) FilterExitDataProcessing(opts *bind.FilterOpts) (*VeboExitDataProcessingIterator, error) {

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ExitDataProcessing")
	if err != nil {
		return nil, err
	}
	return &VeboExitDataProcessingIterator{contract: _Vebo.contract, event: "ExitDataProcessing", logs: logs, sub: sub}, nil
}

// WatchExitDataProcessing is a free log subscription operation binding the contract event 0x01b8de053572c3c2104259b555c485ccac8017196b3471e8483b7e96f071608a.
//
// Solidity: event ExitDataProcessing(bytes32 exitRequestsHash)
func (_Vebo *VeboFilterer) WatchExitDataProcessing(opts *bind.WatchOpts, sink chan<- *VeboExitDataProcessing) (event.Subscription, error) {

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ExitDataProcessing")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboExitDataProcessing)
				if err := _Vebo.contract.UnpackLog(event, "ExitDataProcessing", log); err != nil {
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

// ParseExitDataProcessing is a log parse operation binding the contract event 0x01b8de053572c3c2104259b555c485ccac8017196b3471e8483b7e96f071608a.
//
// Solidity: event ExitDataProcessing(bytes32 exitRequestsHash)
func (_Vebo *VeboFilterer) ParseExitDataProcessing(log types.Log) (*VeboExitDataProcessing, error) {
	event := new(VeboExitDataProcessing)
	if err := _Vebo.contract.UnpackLog(event, "ExitDataProcessing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboExitRequestsLimitSetIterator is returned from FilterExitRequestsLimitSet and is used to iterate over the raw logs and unpacked data for ExitRequestsLimitSet events raised by the Vebo contract.
type VeboExitRequestsLimitSetIterator struct {
	Event *VeboExitRequestsLimitSet // Event containing the contract specifics and raw log

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
func (it *VeboExitRequestsLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboExitRequestsLimitSet)
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
		it.Event = new(VeboExitRequestsLimitSet)
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
func (it *VeboExitRequestsLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboExitRequestsLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboExitRequestsLimitSet represents a ExitRequestsLimitSet event raised by the Vebo contract.
type VeboExitRequestsLimitSet struct {
	MaxExitRequestsLimit *big.Int
	ExitsPerFrame        *big.Int
	FrameDurationInSec   *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterExitRequestsLimitSet is a free log retrieval operation binding the contract event 0x3119d910326e0f179e121df55f23f45b8a5022ff10c73c02aabf2b48ae36070a.
//
// Solidity: event ExitRequestsLimitSet(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec)
func (_Vebo *VeboFilterer) FilterExitRequestsLimitSet(opts *bind.FilterOpts) (*VeboExitRequestsLimitSetIterator, error) {

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ExitRequestsLimitSet")
	if err != nil {
		return nil, err
	}
	return &VeboExitRequestsLimitSetIterator{contract: _Vebo.contract, event: "ExitRequestsLimitSet", logs: logs, sub: sub}, nil
}

// WatchExitRequestsLimitSet is a free log subscription operation binding the contract event 0x3119d910326e0f179e121df55f23f45b8a5022ff10c73c02aabf2b48ae36070a.
//
// Solidity: event ExitRequestsLimitSet(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec)
func (_Vebo *VeboFilterer) WatchExitRequestsLimitSet(opts *bind.WatchOpts, sink chan<- *VeboExitRequestsLimitSet) (event.Subscription, error) {

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ExitRequestsLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboExitRequestsLimitSet)
				if err := _Vebo.contract.UnpackLog(event, "ExitRequestsLimitSet", log); err != nil {
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

// ParseExitRequestsLimitSet is a log parse operation binding the contract event 0x3119d910326e0f179e121df55f23f45b8a5022ff10c73c02aabf2b48ae36070a.
//
// Solidity: event ExitRequestsLimitSet(uint256 maxExitRequestsLimit, uint256 exitsPerFrame, uint256 frameDurationInSec)
func (_Vebo *VeboFilterer) ParseExitRequestsLimitSet(log types.Log) (*VeboExitRequestsLimitSet, error) {
	event := new(VeboExitRequestsLimitSet)
	if err := _Vebo.contract.UnpackLog(event, "ExitRequestsLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Vebo contract.
type VeboPausedIterator struct {
	Event *VeboPaused // Event containing the contract specifics and raw log

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
func (it *VeboPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboPaused)
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
		it.Event = new(VeboPaused)
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
func (it *VeboPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboPaused represents a Paused event raised by the Vebo contract.
type VeboPaused struct {
	Duration *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Vebo *VeboFilterer) FilterPaused(opts *bind.FilterOpts) (*VeboPausedIterator, error) {

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &VeboPausedIterator{contract: _Vebo.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x32fb7c9891bc4f963c7de9f1186d2a7755c7d6e9f4604dabe1d8bb3027c2f49e.
//
// Solidity: event Paused(uint256 duration)
func (_Vebo *VeboFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *VeboPaused) (event.Subscription, error) {

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboPaused)
				if err := _Vebo.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Vebo *VeboFilterer) ParsePaused(log types.Log) (*VeboPaused, error) {
	event := new(VeboPaused)
	if err := _Vebo.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboProcessingStartedIterator is returned from FilterProcessingStarted and is used to iterate over the raw logs and unpacked data for ProcessingStarted events raised by the Vebo contract.
type VeboProcessingStartedIterator struct {
	Event *VeboProcessingStarted // Event containing the contract specifics and raw log

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
func (it *VeboProcessingStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboProcessingStarted)
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
		it.Event = new(VeboProcessingStarted)
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
func (it *VeboProcessingStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboProcessingStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboProcessingStarted represents a ProcessingStarted event raised by the Vebo contract.
type VeboProcessingStarted struct {
	RefSlot *big.Int
	Hash    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterProcessingStarted is a free log retrieval operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_Vebo *VeboFilterer) FilterProcessingStarted(opts *bind.FilterOpts, refSlot []*big.Int) (*VeboProcessingStartedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ProcessingStarted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &VeboProcessingStartedIterator{contract: _Vebo.contract, event: "ProcessingStarted", logs: logs, sub: sub}, nil
}

// WatchProcessingStarted is a free log subscription operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_Vebo *VeboFilterer) WatchProcessingStarted(opts *bind.WatchOpts, sink chan<- *VeboProcessingStarted, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ProcessingStarted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboProcessingStarted)
				if err := _Vebo.contract.UnpackLog(event, "ProcessingStarted", log); err != nil {
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

// ParseProcessingStarted is a log parse operation binding the contract event 0xf73febded7d4502284718948a3e1d75406151c6326bde069424a584a4f6af87a.
//
// Solidity: event ProcessingStarted(uint256 indexed refSlot, bytes32 hash)
func (_Vebo *VeboFilterer) ParseProcessingStarted(log types.Log) (*VeboProcessingStarted, error) {
	event := new(VeboProcessingStarted)
	if err := _Vebo.contract.UnpackLog(event, "ProcessingStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboReportDiscardedIterator is returned from FilterReportDiscarded and is used to iterate over the raw logs and unpacked data for ReportDiscarded events raised by the Vebo contract.
type VeboReportDiscardedIterator struct {
	Event *VeboReportDiscarded // Event containing the contract specifics and raw log

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
func (it *VeboReportDiscardedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboReportDiscarded)
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
		it.Event = new(VeboReportDiscarded)
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
func (it *VeboReportDiscardedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboReportDiscardedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboReportDiscarded represents a ReportDiscarded event raised by the Vebo contract.
type VeboReportDiscarded struct {
	RefSlot *big.Int
	Hash    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReportDiscarded is a free log retrieval operation binding the contract event 0xe21266bc27ee721ac10034efaf7fd724656ef471c75b8402cd8f07850af6b676.
//
// Solidity: event ReportDiscarded(uint256 indexed refSlot, bytes32 hash)
func (_Vebo *VeboFilterer) FilterReportDiscarded(opts *bind.FilterOpts, refSlot []*big.Int) (*VeboReportDiscardedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ReportDiscarded", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &VeboReportDiscardedIterator{contract: _Vebo.contract, event: "ReportDiscarded", logs: logs, sub: sub}, nil
}

// WatchReportDiscarded is a free log subscription operation binding the contract event 0xe21266bc27ee721ac10034efaf7fd724656ef471c75b8402cd8f07850af6b676.
//
// Solidity: event ReportDiscarded(uint256 indexed refSlot, bytes32 hash)
func (_Vebo *VeboFilterer) WatchReportDiscarded(opts *bind.WatchOpts, sink chan<- *VeboReportDiscarded, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ReportDiscarded", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboReportDiscarded)
				if err := _Vebo.contract.UnpackLog(event, "ReportDiscarded", log); err != nil {
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

// ParseReportDiscarded is a log parse operation binding the contract event 0xe21266bc27ee721ac10034efaf7fd724656ef471c75b8402cd8f07850af6b676.
//
// Solidity: event ReportDiscarded(uint256 indexed refSlot, bytes32 hash)
func (_Vebo *VeboFilterer) ParseReportDiscarded(log types.Log) (*VeboReportDiscarded, error) {
	event := new(VeboReportDiscarded)
	if err := _Vebo.contract.UnpackLog(event, "ReportDiscarded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboReportSubmittedIterator is returned from FilterReportSubmitted and is used to iterate over the raw logs and unpacked data for ReportSubmitted events raised by the Vebo contract.
type VeboReportSubmittedIterator struct {
	Event *VeboReportSubmitted // Event containing the contract specifics and raw log

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
func (it *VeboReportSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboReportSubmitted)
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
		it.Event = new(VeboReportSubmitted)
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
func (it *VeboReportSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboReportSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboReportSubmitted represents a ReportSubmitted event raised by the Vebo contract.
type VeboReportSubmitted struct {
	RefSlot                *big.Int
	Hash                   [32]byte
	ProcessingDeadlineTime *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterReportSubmitted is a free log retrieval operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_Vebo *VeboFilterer) FilterReportSubmitted(opts *bind.FilterOpts, refSlot []*big.Int) (*VeboReportSubmittedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ReportSubmitted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &VeboReportSubmittedIterator{contract: _Vebo.contract, event: "ReportSubmitted", logs: logs, sub: sub}, nil
}

// WatchReportSubmitted is a free log subscription operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_Vebo *VeboFilterer) WatchReportSubmitted(opts *bind.WatchOpts, sink chan<- *VeboReportSubmitted, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ReportSubmitted", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboReportSubmitted)
				if err := _Vebo.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
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

// ParseReportSubmitted is a log parse operation binding the contract event 0xaed7d1a7a1831158dcda1e4214f5862f450bd3eb5721a5f322bf8c9fe1790b0a.
//
// Solidity: event ReportSubmitted(uint256 indexed refSlot, bytes32 hash, uint256 processingDeadlineTime)
func (_Vebo *VeboFilterer) ParseReportSubmitted(log types.Log) (*VeboReportSubmitted, error) {
	event := new(VeboReportSubmitted)
	if err := _Vebo.contract.UnpackLog(event, "ReportSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboRequestsHashSubmittedIterator is returned from FilterRequestsHashSubmitted and is used to iterate over the raw logs and unpacked data for RequestsHashSubmitted events raised by the Vebo contract.
type VeboRequestsHashSubmittedIterator struct {
	Event *VeboRequestsHashSubmitted // Event containing the contract specifics and raw log

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
func (it *VeboRequestsHashSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboRequestsHashSubmitted)
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
		it.Event = new(VeboRequestsHashSubmitted)
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
func (it *VeboRequestsHashSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboRequestsHashSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboRequestsHashSubmitted represents a RequestsHashSubmitted event raised by the Vebo contract.
type VeboRequestsHashSubmitted struct {
	ExitRequestsHash [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterRequestsHashSubmitted is a free log retrieval operation binding the contract event 0x76d8359ea28964b79f7fa8bb502ec325fd0d1e956c42a0436940e35d0e99f2de.
//
// Solidity: event RequestsHashSubmitted(bytes32 exitRequestsHash)
func (_Vebo *VeboFilterer) FilterRequestsHashSubmitted(opts *bind.FilterOpts) (*VeboRequestsHashSubmittedIterator, error) {

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "RequestsHashSubmitted")
	if err != nil {
		return nil, err
	}
	return &VeboRequestsHashSubmittedIterator{contract: _Vebo.contract, event: "RequestsHashSubmitted", logs: logs, sub: sub}, nil
}

// WatchRequestsHashSubmitted is a free log subscription operation binding the contract event 0x76d8359ea28964b79f7fa8bb502ec325fd0d1e956c42a0436940e35d0e99f2de.
//
// Solidity: event RequestsHashSubmitted(bytes32 exitRequestsHash)
func (_Vebo *VeboFilterer) WatchRequestsHashSubmitted(opts *bind.WatchOpts, sink chan<- *VeboRequestsHashSubmitted) (event.Subscription, error) {

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "RequestsHashSubmitted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboRequestsHashSubmitted)
				if err := _Vebo.contract.UnpackLog(event, "RequestsHashSubmitted", log); err != nil {
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

// ParseRequestsHashSubmitted is a log parse operation binding the contract event 0x76d8359ea28964b79f7fa8bb502ec325fd0d1e956c42a0436940e35d0e99f2de.
//
// Solidity: event RequestsHashSubmitted(bytes32 exitRequestsHash)
func (_Vebo *VeboFilterer) ParseRequestsHashSubmitted(log types.Log) (*VeboRequestsHashSubmitted, error) {
	event := new(VeboRequestsHashSubmitted)
	if err := _Vebo.contract.UnpackLog(event, "RequestsHashSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboResumedIterator is returned from FilterResumed and is used to iterate over the raw logs and unpacked data for Resumed events raised by the Vebo contract.
type VeboResumedIterator struct {
	Event *VeboResumed // Event containing the contract specifics and raw log

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
func (it *VeboResumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboResumed)
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
		it.Event = new(VeboResumed)
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
func (it *VeboResumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboResumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboResumed represents a Resumed event raised by the Vebo contract.
type VeboResumed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterResumed is a free log retrieval operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Vebo *VeboFilterer) FilterResumed(opts *bind.FilterOpts) (*VeboResumedIterator, error) {

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return &VeboResumedIterator{contract: _Vebo.contract, event: "Resumed", logs: logs, sub: sub}, nil
}

// WatchResumed is a free log subscription operation binding the contract event 0x62451d457bc659158be6e6247f56ec1df424a5c7597f71c20c2bc44e0965c8f9.
//
// Solidity: event Resumed()
func (_Vebo *VeboFilterer) WatchResumed(opts *bind.WatchOpts, sink chan<- *VeboResumed) (event.Subscription, error) {

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "Resumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboResumed)
				if err := _Vebo.contract.UnpackLog(event, "Resumed", log); err != nil {
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
func (_Vebo *VeboFilterer) ParseResumed(log types.Log) (*VeboResumed, error) {
	event := new(VeboResumed)
	if err := _Vebo.contract.UnpackLog(event, "Resumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Vebo contract.
type VeboRoleAdminChangedIterator struct {
	Event *VeboRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *VeboRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboRoleAdminChanged)
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
		it.Event = new(VeboRoleAdminChanged)
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
func (it *VeboRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboRoleAdminChanged represents a RoleAdminChanged event raised by the Vebo contract.
type VeboRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vebo *VeboFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*VeboRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &VeboRoleAdminChangedIterator{contract: _Vebo.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Vebo *VeboFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *VeboRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboRoleAdminChanged)
				if err := _Vebo.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Vebo *VeboFilterer) ParseRoleAdminChanged(log types.Log) (*VeboRoleAdminChanged, error) {
	event := new(VeboRoleAdminChanged)
	if err := _Vebo.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Vebo contract.
type VeboRoleGrantedIterator struct {
	Event *VeboRoleGranted // Event containing the contract specifics and raw log

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
func (it *VeboRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboRoleGranted)
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
		it.Event = new(VeboRoleGranted)
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
func (it *VeboRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboRoleGranted represents a RoleGranted event raised by the Vebo contract.
type VeboRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vebo *VeboFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VeboRoleGrantedIterator, error) {

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

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VeboRoleGrantedIterator{contract: _Vebo.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vebo *VeboFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *VeboRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboRoleGranted)
				if err := _Vebo.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Vebo *VeboFilterer) ParseRoleGranted(log types.Log) (*VeboRoleGranted, error) {
	event := new(VeboRoleGranted)
	if err := _Vebo.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Vebo contract.
type VeboRoleRevokedIterator struct {
	Event *VeboRoleRevoked // Event containing the contract specifics and raw log

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
func (it *VeboRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboRoleRevoked)
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
		it.Event = new(VeboRoleRevoked)
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
func (it *VeboRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboRoleRevoked represents a RoleRevoked event raised by the Vebo contract.
type VeboRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vebo *VeboFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*VeboRoleRevokedIterator, error) {

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

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &VeboRoleRevokedIterator{contract: _Vebo.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Vebo *VeboFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *VeboRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboRoleRevoked)
				if err := _Vebo.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Vebo *VeboFilterer) ParseRoleRevoked(log types.Log) (*VeboRoleRevoked, error) {
	event := new(VeboRoleRevoked)
	if err := _Vebo.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboSetMaxValidatorsPerReportIterator is returned from FilterSetMaxValidatorsPerReport and is used to iterate over the raw logs and unpacked data for SetMaxValidatorsPerReport events raised by the Vebo contract.
type VeboSetMaxValidatorsPerReportIterator struct {
	Event *VeboSetMaxValidatorsPerReport // Event containing the contract specifics and raw log

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
func (it *VeboSetMaxValidatorsPerReportIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboSetMaxValidatorsPerReport)
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
		it.Event = new(VeboSetMaxValidatorsPerReport)
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
func (it *VeboSetMaxValidatorsPerReportIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboSetMaxValidatorsPerReportIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboSetMaxValidatorsPerReport represents a SetMaxValidatorsPerReport event raised by the Vebo contract.
type VeboSetMaxValidatorsPerReport struct {
	MaxValidatorsPerReport *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterSetMaxValidatorsPerReport is a free log retrieval operation binding the contract event 0x9b17a153b6e933d8497c6b713fbd70c893891d75639ede17ce6e4cea08e7cfc3.
//
// Solidity: event SetMaxValidatorsPerReport(uint256 maxValidatorsPerReport)
func (_Vebo *VeboFilterer) FilterSetMaxValidatorsPerReport(opts *bind.FilterOpts) (*VeboSetMaxValidatorsPerReportIterator, error) {

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "SetMaxValidatorsPerReport")
	if err != nil {
		return nil, err
	}
	return &VeboSetMaxValidatorsPerReportIterator{contract: _Vebo.contract, event: "SetMaxValidatorsPerReport", logs: logs, sub: sub}, nil
}

// WatchSetMaxValidatorsPerReport is a free log subscription operation binding the contract event 0x9b17a153b6e933d8497c6b713fbd70c893891d75639ede17ce6e4cea08e7cfc3.
//
// Solidity: event SetMaxValidatorsPerReport(uint256 maxValidatorsPerReport)
func (_Vebo *VeboFilterer) WatchSetMaxValidatorsPerReport(opts *bind.WatchOpts, sink chan<- *VeboSetMaxValidatorsPerReport) (event.Subscription, error) {

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "SetMaxValidatorsPerReport")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboSetMaxValidatorsPerReport)
				if err := _Vebo.contract.UnpackLog(event, "SetMaxValidatorsPerReport", log); err != nil {
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

// ParseSetMaxValidatorsPerReport is a log parse operation binding the contract event 0x9b17a153b6e933d8497c6b713fbd70c893891d75639ede17ce6e4cea08e7cfc3.
//
// Solidity: event SetMaxValidatorsPerReport(uint256 maxValidatorsPerReport)
func (_Vebo *VeboFilterer) ParseSetMaxValidatorsPerReport(log types.Log) (*VeboSetMaxValidatorsPerReport, error) {
	event := new(VeboSetMaxValidatorsPerReport)
	if err := _Vebo.contract.UnpackLog(event, "SetMaxValidatorsPerReport", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboValidatorExitRequestIterator is returned from FilterValidatorExitRequest and is used to iterate over the raw logs and unpacked data for ValidatorExitRequest events raised by the Vebo contract.
type VeboValidatorExitRequestIterator struct {
	Event *VeboValidatorExitRequest // Event containing the contract specifics and raw log

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
func (it *VeboValidatorExitRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboValidatorExitRequest)
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
		it.Event = new(VeboValidatorExitRequest)
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
func (it *VeboValidatorExitRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboValidatorExitRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboValidatorExitRequest represents a ValidatorExitRequest event raised by the Vebo contract.
type VeboValidatorExitRequest struct {
	StakingModuleId *big.Int
	NodeOperatorId  *big.Int
	ValidatorIndex  *big.Int
	ValidatorPubkey []byte
	Timestamp       *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterValidatorExitRequest is a free log retrieval operation binding the contract event 0x96395f55c4997466e5035d777f0e1ba82b8cae217aaad05cf07839eb7c75bcf2.
//
// Solidity: event ValidatorExitRequest(uint256 indexed stakingModuleId, uint256 indexed nodeOperatorId, uint256 indexed validatorIndex, bytes validatorPubkey, uint256 timestamp)
func (_Vebo *VeboFilterer) FilterValidatorExitRequest(opts *bind.FilterOpts, stakingModuleId []*big.Int, nodeOperatorId []*big.Int, validatorIndex []*big.Int) (*VeboValidatorExitRequestIterator, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}
	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "ValidatorExitRequest", stakingModuleIdRule, nodeOperatorIdRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return &VeboValidatorExitRequestIterator{contract: _Vebo.contract, event: "ValidatorExitRequest", logs: logs, sub: sub}, nil
}

// WatchValidatorExitRequest is a free log subscription operation binding the contract event 0x96395f55c4997466e5035d777f0e1ba82b8cae217aaad05cf07839eb7c75bcf2.
//
// Solidity: event ValidatorExitRequest(uint256 indexed stakingModuleId, uint256 indexed nodeOperatorId, uint256 indexed validatorIndex, bytes validatorPubkey, uint256 timestamp)
func (_Vebo *VeboFilterer) WatchValidatorExitRequest(opts *bind.WatchOpts, sink chan<- *VeboValidatorExitRequest, stakingModuleId []*big.Int, nodeOperatorId []*big.Int, validatorIndex []*big.Int) (event.Subscription, error) {

	var stakingModuleIdRule []interface{}
	for _, stakingModuleIdItem := range stakingModuleId {
		stakingModuleIdRule = append(stakingModuleIdRule, stakingModuleIdItem)
	}
	var nodeOperatorIdRule []interface{}
	for _, nodeOperatorIdItem := range nodeOperatorId {
		nodeOperatorIdRule = append(nodeOperatorIdRule, nodeOperatorIdItem)
	}
	var validatorIndexRule []interface{}
	for _, validatorIndexItem := range validatorIndex {
		validatorIndexRule = append(validatorIndexRule, validatorIndexItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "ValidatorExitRequest", stakingModuleIdRule, nodeOperatorIdRule, validatorIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboValidatorExitRequest)
				if err := _Vebo.contract.UnpackLog(event, "ValidatorExitRequest", log); err != nil {
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

// ParseValidatorExitRequest is a log parse operation binding the contract event 0x96395f55c4997466e5035d777f0e1ba82b8cae217aaad05cf07839eb7c75bcf2.
//
// Solidity: event ValidatorExitRequest(uint256 indexed stakingModuleId, uint256 indexed nodeOperatorId, uint256 indexed validatorIndex, bytes validatorPubkey, uint256 timestamp)
func (_Vebo *VeboFilterer) ParseValidatorExitRequest(log types.Log) (*VeboValidatorExitRequest, error) {
	event := new(VeboValidatorExitRequest)
	if err := _Vebo.contract.UnpackLog(event, "ValidatorExitRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboWarnDataIncompleteProcessingIterator is returned from FilterWarnDataIncompleteProcessing and is used to iterate over the raw logs and unpacked data for WarnDataIncompleteProcessing events raised by the Vebo contract.
type VeboWarnDataIncompleteProcessingIterator struct {
	Event *VeboWarnDataIncompleteProcessing // Event containing the contract specifics and raw log

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
func (it *VeboWarnDataIncompleteProcessingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboWarnDataIncompleteProcessing)
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
		it.Event = new(VeboWarnDataIncompleteProcessing)
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
func (it *VeboWarnDataIncompleteProcessingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboWarnDataIncompleteProcessingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboWarnDataIncompleteProcessing represents a WarnDataIncompleteProcessing event raised by the Vebo contract.
type VeboWarnDataIncompleteProcessing struct {
	RefSlot           *big.Int
	RequestsProcessed *big.Int
	RequestsCount     *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterWarnDataIncompleteProcessing is a free log retrieval operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 requestsProcessed, uint256 requestsCount)
func (_Vebo *VeboFilterer) FilterWarnDataIncompleteProcessing(opts *bind.FilterOpts, refSlot []*big.Int) (*VeboWarnDataIncompleteProcessingIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "WarnDataIncompleteProcessing", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &VeboWarnDataIncompleteProcessingIterator{contract: _Vebo.contract, event: "WarnDataIncompleteProcessing", logs: logs, sub: sub}, nil
}

// WatchWarnDataIncompleteProcessing is a free log subscription operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 requestsProcessed, uint256 requestsCount)
func (_Vebo *VeboFilterer) WatchWarnDataIncompleteProcessing(opts *bind.WatchOpts, sink chan<- *VeboWarnDataIncompleteProcessing, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "WarnDataIncompleteProcessing", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboWarnDataIncompleteProcessing)
				if err := _Vebo.contract.UnpackLog(event, "WarnDataIncompleteProcessing", log); err != nil {
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

// ParseWarnDataIncompleteProcessing is a log parse operation binding the contract event 0xefc67aab43195093a8d8ed25d52281d96de480748ece2787888c586e8e1e79b4.
//
// Solidity: event WarnDataIncompleteProcessing(uint256 indexed refSlot, uint256 requestsProcessed, uint256 requestsCount)
func (_Vebo *VeboFilterer) ParseWarnDataIncompleteProcessing(log types.Log) (*VeboWarnDataIncompleteProcessing, error) {
	event := new(VeboWarnDataIncompleteProcessing)
	if err := _Vebo.contract.UnpackLog(event, "WarnDataIncompleteProcessing", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VeboWarnProcessingMissedIterator is returned from FilterWarnProcessingMissed and is used to iterate over the raw logs and unpacked data for WarnProcessingMissed events raised by the Vebo contract.
type VeboWarnProcessingMissedIterator struct {
	Event *VeboWarnProcessingMissed // Event containing the contract specifics and raw log

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
func (it *VeboWarnProcessingMissedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VeboWarnProcessingMissed)
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
		it.Event = new(VeboWarnProcessingMissed)
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
func (it *VeboWarnProcessingMissedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VeboWarnProcessingMissedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VeboWarnProcessingMissed represents a WarnProcessingMissed event raised by the Vebo contract.
type VeboWarnProcessingMissed struct {
	RefSlot *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWarnProcessingMissed is a free log retrieval operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_Vebo *VeboFilterer) FilterWarnProcessingMissed(opts *bind.FilterOpts, refSlot []*big.Int) (*VeboWarnProcessingMissedIterator, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.FilterLogs(opts, "WarnProcessingMissed", refSlotRule)
	if err != nil {
		return nil, err
	}
	return &VeboWarnProcessingMissedIterator{contract: _Vebo.contract, event: "WarnProcessingMissed", logs: logs, sub: sub}, nil
}

// WatchWarnProcessingMissed is a free log subscription operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_Vebo *VeboFilterer) WatchWarnProcessingMissed(opts *bind.WatchOpts, sink chan<- *VeboWarnProcessingMissed, refSlot []*big.Int) (event.Subscription, error) {

	var refSlotRule []interface{}
	for _, refSlotItem := range refSlot {
		refSlotRule = append(refSlotRule, refSlotItem)
	}

	logs, sub, err := _Vebo.contract.WatchLogs(opts, "WarnProcessingMissed", refSlotRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VeboWarnProcessingMissed)
				if err := _Vebo.contract.UnpackLog(event, "WarnProcessingMissed", log); err != nil {
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

// ParseWarnProcessingMissed is a log parse operation binding the contract event 0x800b849c8bf80718cf786c99d1091c079fe2c5e420a3ba7ba9b0ef8179ef2c38.
//
// Solidity: event WarnProcessingMissed(uint256 indexed refSlot)
func (_Vebo *VeboFilterer) ParseWarnProcessingMissed(log types.Log) (*VeboWarnProcessingMissed, error) {
	event := new(VeboWarnProcessingMissed)
	if err := _Vebo.contract.UnpackLog(event, "WarnProcessingMissed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
