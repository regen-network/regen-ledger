package core

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/gocuke"
	"github.com/tendermint/tendermint/libs/rand"
	"gotest.tools/v3/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type bridgeReceiveSuite struct {
	*baseSuite
	// CONSTANTS: DO NOT CHANGE
	// the test steps expect these variables
	// to always remain the same.
	validClassId    string
	validProjectId  string
	validBatchDenom string

	// fields that may be altered by test steps
	classId, projectId, batchDenom string
	bridgeAddr                     sdk.AccAddress
	recipient                      sdk.AccAddress
	amount                         string
	referenceId                    string
	msg                            core.MsgBridgeReceive
	res                            core.MsgBridgeReceiveResponse
}

func TestBridgeReceive(t *testing.T) {
	t.Parallel()
	gocuke.NewRunner(t, &bridgeReceiveSuite{}).Path("./features/msg_bridge_receive.feature").Run()
}

func (s *bridgeReceiveSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.validClassId, s.validProjectId, s.validBatchDenom = s.setupClassProjectBatch(t)
}

func (s *bridgeReceiveSuite) AValidClassId() {
	s.classId = s.validClassId
}

func (s *bridgeReceiveSuite) TheBridgeAddressIsAnIssuer(a string) {
	class, err := s.stateStore.ClassTable().GetById(s.ctx, s.classId)
	assert.NilError(s.t, err)
	s.bridgeAddr = sdk.AccAddress(a)
	err = s.stateStore.ClassIssuerTable().Insert(s.ctx, &api.ClassIssuer{
		ClassKey: class.Key,
		Issuer:   s.bridgeAddr,
	})
}

func (s *bridgeReceiveSuite) AValidBridgeMsg() {
	s.recipient = sdk.AccAddress(rand.Str(5))
	s.amount = "3"
	s.referenceId = "VCS-001"
	start, end := time.Now(), time.Now()
	s.msg = core.MsgBridgeReceive{
		Issuer: s.bridgeAddr.String(),
		Batch: &core.MsgBridgeReceive_Batch{
			Recipient: s.recipient.String(),
			Amount:    s.amount,
			StartDate: &start,
			EndDate:   &end,
			Metadata:  "hello",
		},
		Project: &core.MsgBridgeReceive_Project{
			ReferenceId:  s.referenceId,
			Jurisdiction: "US-OR",
			Metadata:     "hi",
		},
		ClassId: s.classId,
		OriginTx: &core.OriginTx{
			Id:     "0x12345",
			Source: "polygon",
		},
		Note: "note",
	}
}

func (s *bridgeReceiveSuite) TheTransactionSucceeds() {
	res, err := s.k.BridgeReceive(s.ctx, &s.msg)
	assert.NilError(s.t, err)
	s.res = *res
}

func (s *bridgeReceiveSuite) TheBridgeAddressIsNotAnIssuer(a string) {
	s.bridgeAddr = sdk.AccAddress([]byte(a))
}

func (s *bridgeReceiveSuite) AnInvalidClassId() {
	s.classId = "INV001"
}

func (s *bridgeReceiveSuite) TheBridgeAddress(a string) {
	s.bridgeAddr = sdk.AccAddress([]byte(a))
}

func (s *bridgeReceiveSuite) ANewReferenceId() {
	s.referenceId = rand.Str(5)

	// sanity check that this ref ID hasn't been used yet.
	it, err := s.stateStore.ProjectTable().List(s.ctx, api.ProjectReferenceIdIndexKey{}.WithReferenceId(s.referenceId))
	assert.NilError(s.t, err)
	assert.Equal(s.t, false, it.Next())
	it.Close()
}

func (s *bridgeReceiveSuite) ANewProjectIsCreated() {
	project, err := s.stateStore.ProjectTable().GetById(s.ctx, s.res.ProjectId)
	assert.NilError(s.t, err)
	assert.Equal(s.t, project.ReferenceId, s.referenceId)
}

func (s *bridgeReceiveSuite) TheTransactionFailsWith(a string) {
	_, err := s.k.BridgeReceive(s.ctx, &s.msg)
	assert.ErrorContains(s.t, err, a)
}
