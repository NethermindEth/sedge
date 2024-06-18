package mevboostrelaylist

import (
	"io"
	"log"
	"math/big"
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient/simulated"
	"github.com/ethereum/go-ethereum/params"
	"github.com/stretchr/testify/suite"
)

type MEVBoostRelayAllowedListTestSuite struct {
	suite.Suite
	auth            *bind.TransactOpts
	address         common.Address
	gAlloc          types.GenesisAlloc
	sim             *simulated.Backend
	contractAddress common.Address
	tx              *types.Transaction
	apiContract     *Api
	callOpts        *bind.CallOpts
}

func TestMEVBoostRelayAllowedListSuite(t *testing.T) {
	suite.Run(t, new(MEVBoostRelayAllowedListTestSuite))
}

func (s *MEVBoostRelayAllowedListTestSuite) SetupSuite() {
	// Silence logger
	log.SetOutput(io.Discard)

	key, err := crypto.GenerateKey()
	if err != nil {
		s.T().Fatalf("Failed to generate private key: %v", err)
	}

	chainID := big.NewInt(params.AllEthashProtocolChanges.ChainID.Int64())

	s.auth, err = bind.NewKeyedTransactorWithChainID(key, chainID)
	if err != nil {
		s.T().Fatalf("Failed to create transactor with chain ID: %v", err)
	}

	s.address = s.auth.From

	s.gAlloc = types.GenesisAlloc{
		s.address: {Balance: big.NewInt(9e18)},
	}

	s.sim = simulated.NewBackend(s.gAlloc)

	s.callOpts = &bind.CallOpts{
		Pending: false,
		From:    s.address,
		Context: nil,
	}

	contractAddress, tx, apiContract, err := DeployApi(s.auth, s.sim.Client(), s.address)
	if err != nil {
		s.T().Fatalf("Failed to deploy Api contract: %v", err)
	}
	s.sim.Commit()
	s.contractAddress = contractAddress
	s.tx = tx
	s.apiContract = apiContract
}

func (s *MEVBoostRelayAllowedListTestSuite) TearDownSuite() {
	// Cleanup resources if necessary
	if s.sim != nil {
		s.sim.Close()
	}
}

func (s *MEVBoostRelayAllowedListTestSuite) TestGetOwner() {
	ownerAddress, err := s.apiContract.GetOwner(s.callOpts)
	if err != nil {
		s.T().Fatalf("Failed to call GetOwner: %v", err)
	}
	if ownerAddress != s.address {
		s.T().Errorf("Incorrect owner address. Expected: %s, Got: %s", s.address.Hex(), ownerAddress.Hex())
	}
}

func (s *MEVBoostRelayAllowedListTestSuite) TestSetAndGetManager() {
	newManagerAddress := common.HexToAddress("0xabcdefabcdefabcdefabcdefabcdefabcdef1234")

	_, err := s.apiContract.SetManager(s.auth, newManagerAddress)
	if err != nil {
		s.T().Fatalf("Failed to call SetManager: %v", err)
	}
	s.sim.Commit()

	managerAddress, err := s.apiContract.GetManager(s.callOpts)
	if err != nil {
		s.T().Fatalf("Failed to call GetManager: %v", err)
	}
	if managerAddress != newManagerAddress {
		s.T().Errorf("Incorrect manager address. Expected: %s, Got: %s", newManagerAddress.Hex(), managerAddress.Hex())
	}
}

func (s *MEVBoostRelayAllowedListTestSuite) TestDismissManager() {
	newManagerAddress := common.HexToAddress("0x1234567890123456789012345678901234567890")

	_, err := s.apiContract.SetManager(s.auth, newManagerAddress)
	if err != nil {
		s.T().Fatalf("Failed to call SetManager: %v", err)
	}
	s.sim.Commit()

	_, err = s.apiContract.DismissManager(s.auth)
	if err != nil {
		s.T().Fatalf("Failed to call DismissManager: %v", err)
	}
	s.sim.Commit()

	address, _ := s.apiContract.GetManager(s.callOpts)
	if address != common.BigToAddress(big.NewInt(0)) {
		s.T().Fatalf("Expected GetManager to return address(0) after dismissing manager. Expected: %s, Got: %s", common.BigToAddress(big.NewInt(0)).Hex(), address.Hex())
	}
}

func (s *MEVBoostRelayAllowedListTestSuite) TestGetRelaysAmount() {
	amount, err := s.apiContract.GetRelaysAmount(s.callOpts)
	if err != nil {
		s.T().Fatalf("Failed to call GetRelaysAmount: %v", err)
	}

	relays, err := s.apiContract.GetRelays(s.callOpts)
	if err != nil {
		s.T().Fatalf("Failed to call GetRelays: %v", err)
	}
	expectedAmount := big.NewInt(int64(len(relays)))
	if amount.Cmp(expectedAmount) != 0 {
		s.T().Errorf("Incorrect relays amount. Expected: %s, Got: %s", expectedAmount, amount)
	}
}

func (s *MEVBoostRelayAllowedListTestSuite) TestGetRelayByUri() {
	expectedRelay := Struct0{
		Uri:         "example_uri",
		Operator:    "example_operator",
		IsMandatory: true,
		Description: "example_description",
	}
	_, err := s.apiContract.AddRelay(s.auth, expectedRelay.Uri, expectedRelay.Operator, expectedRelay.IsMandatory, expectedRelay.Description)
	if err != nil {
		s.T().Fatalf("Failed to call AddRelay: %v", err)
	}
	s.sim.Commit()

	relay, err := s.apiContract.GetRelayByUri(s.callOpts, expectedRelay.Uri)
	if err != nil {
		s.T().Fatalf("Failed to call GetRelayByUri: %v", err)
	}
	if !reflect.DeepEqual(relay, expectedRelay) {
		s.T().Fatalf("Returned relay does not match expected data. Expected: %+v, Got: %+v", expectedRelay, relay)
	}
}

func (s *MEVBoostRelayAllowedListTestSuite) TestGetAllowedListVersion() {
	version, err := s.apiContract.GetAllowedListVersion(s.callOpts)
	if err != nil {
		s.T().Fatalf("Failed to call GetAllowedListVersion: %v", err)
	}

	addedRelay := Struct0{
		Uri:         "example__uri",
		Operator:    "example_op",
		IsMandatory: true,
		Description: "example",
	}
	_, err = s.apiContract.AddRelay(s.auth, addedRelay.Uri, addedRelay.Operator, addedRelay.IsMandatory, addedRelay.Description)
	if err != nil {
		s.T().Fatalf("Failed to call AddRelay: %v", err)
	}
	s.sim.Commit()

	updatedVersion, err := s.apiContract.GetAllowedListVersion(s.callOpts)
	if err != nil {
		s.T().Fatalf("Failed to call GetAllowedListVersion: %v", err)
	}
	expectedVersion := version.Add(big.NewInt(1), version)
	if updatedVersion.Cmp(expectedVersion) != 0 {
		s.T().Errorf("Incorrect relays amount. Expected: %s, Got: %s", expectedVersion, updatedVersion)
	}
}
