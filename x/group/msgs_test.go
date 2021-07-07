package group

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMsgCreateGroupValidation(t *testing.T) {
	_, _, myAddr := testdata.KeyTestPubAddr()
	_, _, myOtherAddr := testdata.KeyTestPubAddr()

	specs := map[string]struct {
		src    MsgCreateGroupRequest
		expErr bool
	}{
		"all good with minimum fields set": {
			src: MsgCreateGroupRequest{Admin: myAddr.String()},
		},
		"all good with a member": {
			src: MsgCreateGroupRequest{
				Admin: myAddr.String(),
				Members: []Member{
					{Address: myAddr.String(), Weight: "1"},
				},
			},
		},
		"all good with multiple members": {
			src: MsgCreateGroupRequest{
				Admin: myAddr.String(),
				Members: []Member{
					{Address: myAddr.String(), Weight: "1"},
					{Address: myOtherAddr.String(), Weight: "2"},
				},
			},
		},
		"admin required": {
			src:    MsgCreateGroupRequest{},
			expErr: true,
		},
		"valid admin required": {
			src: MsgCreateGroupRequest{
				Admin: "invalid-address",
			},
			expErr: true,
		},
		"duplicate member addresses not allowed": {
			src: MsgCreateGroupRequest{
				Admin: myAddr.String(),
				Members: []Member{
					{Address: myAddr.String(), Weight: "1"},
					{Address: myAddr.String(), Weight: "2"},
				},
			},
			expErr: true,
		},
		"negative member's weight not allowed": {
			src: MsgCreateGroupRequest{
				Admin: myAddr.String(),
				Members: []Member{
					{Address: myAddr.String(), Weight: "-1"},
				},
			},
			expErr: true,
		},
		"empty member's weight not allowed": {
			src: MsgCreateGroupRequest{
				Admin:   myAddr.String(),
				Members: []Member{{Address: myAddr.String()}},
			},
			expErr: true,
		},
		"zero member's weight not allowed": {
			src: MsgCreateGroupRequest{
				Admin:   myAddr.String(),
				Members: []Member{{Address: myAddr.String(), Weight: "0"}},
			},
			expErr: true,
		},
		"member address required": {
			src: MsgCreateGroupRequest{
				Admin: myAddr.String(),
				Members: []Member{
					{Weight: "1"},
				},
			},
			expErr: true,
		},
		"valid member address required": {
			src: MsgCreateGroupRequest{
				Admin: myAddr.String(),
				Members: []Member{
					{Address: "invalid-address", Weight: "1"},
				},
			},
			expErr: true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			err := spec.src.ValidateBasic()
			if spec.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgCreateGroupSigner(t *testing.T) {
	_, _, myAddr := testdata.KeyTestPubAddr()
	assert.Equal(t, []sdk.AccAddress{myAddr}, MsgCreateGroupRequest{Admin: myAddr.String()}.GetSigners())
}

func TestMsgCreateGroupAccount(t *testing.T) {
	_, _, myAddr := testdata.KeyTestPubAddr()

	specs := map[string]struct {
		admin     sdk.AccAddress
		group     uint64
		threshold string
		timeout   proto.Duration
		expErr    bool
	}{
		"all good with minimum fields set": {
			admin:     myAddr,
			group:     1,
			threshold: "1",
			timeout:   proto.Duration{Seconds: 1},
		},
		"zero threshold not allowed": {
			admin:     myAddr,
			group:     1,
			threshold: "0",
			timeout:   proto.Duration{Seconds: 1},
			expErr:    true,
		},
		"admin required": {
			group:     1,
			threshold: "1",
			timeout:   proto.Duration{Seconds: 1},
			expErr:    true,
		},
		"group required": {
			admin:     myAddr,
			threshold: "1",
			timeout:   proto.Duration{Seconds: 1},
			expErr:    true,
		},
		"decision policy required": {
			admin:  myAddr,
			group:  1,
			expErr: true,
		},
		"decision policy without timeout": {
			admin:     myAddr,
			group:     1,
			threshold: "1",
			expErr:    true,
		},
		"decision policy with invalid timeout": {
			admin:     myAddr,
			group:     1,
			threshold: "1",
			timeout:   proto.Duration{Seconds: -1},
			expErr:    true,
		},
		"decision policy without threshold": {
			admin:   myAddr,
			group:   1,
			timeout: proto.Duration{Seconds: 1},
			expErr:  true,
		},
		"decision policy with negative threshold": {
			admin:     myAddr,
			group:     1,
			threshold: "-1",
			timeout:   proto.Duration{Seconds: 1},
			expErr:    true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			m, err := NewMsgCreateGroupAccountRequest(
				spec.admin,
				spec.group,
				nil,
				&ThresholdDecisionPolicy{
					Threshold: spec.threshold,
					Timeout:   spec.timeout,
				},
			)
			require.NoError(t, err)

			if spec.expErr {
				require.Error(t, m.ValidateBasic())
			} else {
				require.NoError(t, m.ValidateBasic())
			}
		})
	}
}

func TestMsgCreateProposalRequest(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	groupAccAddr := addr.String()

	_, _, addr = testdata.KeyTestPubAddr()
	memberAddr := addr.String()

	specs := map[string]struct {
		src    MsgCreateProposalRequest
		expErr bool
	}{
		"all good with minimum fields set": {
			src: MsgCreateProposalRequest{
				Address:   groupAccAddr,
				Proposers: []string{memberAddr},
			},
		},
		"group account required": {
			src: MsgCreateProposalRequest{
				Proposers: []string{memberAddr},
			},
			expErr: true,
		},
		"proposers required": {
			src: MsgCreateProposalRequest{
				Address: groupAccAddr,
			},
			expErr: true,
		},
		"valid proposer address required": {
			src: MsgCreateProposalRequest{
				Address:   groupAccAddr,
				Proposers: []string{"invalid-member-address"},
			},
			expErr: true,
		},
		"no duplicate proposers": {
			src: MsgCreateProposalRequest{
				Address:   groupAccAddr,
				Proposers: []string{memberAddr, memberAddr},
			},
			expErr: true,
		},
		"empty proposer address not allowed": {
			src: MsgCreateProposalRequest{
				Address:   groupAccAddr,
				Proposers: []string{memberAddr, ""},
			},
			expErr: true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			err := spec.src.ValidateBasic()
			if spec.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgVote(t *testing.T) {
	_, _, addr := testdata.KeyTestPubAddr()
	memberAddr := addr.String()

	specs := map[string]struct {
		src    MsgVoteRequest
		expErr bool
	}{
		"all good with minimum fields set": {
			src: MsgVoteRequest{
				ProposalId: 1,
				Choice:     Choice_CHOICE_YES,
				Voter:      memberAddr,
			},
		},
		"proposal required": {
			src: MsgVoteRequest{
				Choice: Choice_CHOICE_YES,
				Voter:  memberAddr,
			},
			expErr: true,
		},
		"choice required": {
			src: MsgVoteRequest{
				ProposalId: 1,
				Voter:      memberAddr,
			},
			expErr: true,
		},
		"valid choice required": {
			src: MsgVoteRequest{
				ProposalId: 1,
				Choice:     5,
				Voter:      memberAddr,
			},
			expErr: true,
		},
		"voter required": {
			src: MsgVoteRequest{
				ProposalId: 1,
				Choice:     Choice_CHOICE_YES,
			},
			expErr: true,
		},
		"valid voter address required": {
			src: MsgVoteRequest{
				ProposalId: 1,
				Choice:     Choice_CHOICE_YES,
				Voter:      "invalid-member-address",
			},
			expErr: true,
		},
		"empty voters address not allowed": {
			src: MsgVoteRequest{
				ProposalId: 1,
				Choice:     Choice_CHOICE_YES,
				Voter:      "",
			},
			expErr: true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			err := spec.src.ValidateBasic()
			if spec.expErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
