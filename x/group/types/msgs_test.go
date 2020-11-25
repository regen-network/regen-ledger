package types

import (
	"testing"

	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	"github.com/cosmos/cosmos-sdk/types"
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
			src: MsgCreateGroupRequest{Admin: myAddr},
		},
		"all good with a member": {
			src: MsgCreateGroupRequest{
				Admin: myAddr,
				Members: []Member{
					{Address: myAddr, Power: types.NewDec(1)},
				},
			},
		},
		"all good with multiple members": {
			src: MsgCreateGroupRequest{
				Admin: myAddr,
				Members: []Member{
					{Address: myAddr, Power: types.NewDec(1)},
					{Address: myOtherAddr, Power: types.NewDec(2)},
				},
			},
		},
		"admin required": {
			src:    MsgCreateGroupRequest{},
			expErr: true,
		},
		"valid admin required": {
			src: MsgCreateGroupRequest{
				Admin: []byte("invalid-address"),
			},
			expErr: true,
		},
		"duplicate member addresses not allowed": {
			src: MsgCreateGroupRequest{
				Admin: myAddr,
				Members: []Member{
					{Address: myAddr, Power: types.NewDec(1)},
					{Address: myAddr, Power: types.NewDec(2)},
				},
			},
			expErr: true,
		},
		"negative member's power not allowed": {
			src: MsgCreateGroupRequest{
				Admin: myAddr,
				Members: []Member{
					{Address: myAddr, Power: types.NewDec(-1)},
				},
			},
			expErr: true,
		},
		"empty member's power not allowed": {
			src: MsgCreateGroupRequest{
				Admin:   myAddr,
				Members: []Member{{Address: myAddr}},
			},
			expErr: true,
		},
		"zero member's power not allowed": {
			src: MsgCreateGroupRequest{
				Admin:   myAddr,
				Members: []Member{{Address: myAddr, Power: sdk.ZeroDec()}},
			},
			expErr: true,
		},
		"member address required": {
			src: MsgCreateGroupRequest{
				Admin: myAddr,
				Members: []Member{
					{Power: types.NewDec(1)},
				},
			},
			expErr: true,
		},
		"valid member address required": {
			src: MsgCreateGroupRequest{
				Admin: myAddr,
				Members: []Member{
					{Address: []byte("invalid-address"), Power: types.NewDec(1)},
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
	assert.Equal(t, []sdk.AccAddress{myAddr}, MsgCreateGroupRequest{Admin: myAddr}.GetSigners())
}

func TestMsgCreateGroupAccount(t *testing.T) {
	_, _, myAddr := testdata.KeyTestPubAddr()

	specs := map[string]struct {
		admin     sdk.AccAddress
		group     GroupID
		threshold sdk.Dec
		timeout   proto.Duration
		expErr    bool
	}{
		"all good with minimum fields set": {
			admin:     myAddr,
			group:     1,
			threshold: sdk.OneDec(),
			timeout:   proto.Duration{Seconds: 1},
		},
		"zero threshold not allowed": {
			admin:     myAddr,
			group:     1,
			threshold: sdk.ZeroDec(),
			timeout:   proto.Duration{Seconds: 1},
			expErr:    true,
		},
		"admin required": {
			group:     1,
			threshold: sdk.ZeroDec(),
			timeout:   proto.Duration{Seconds: 1},
			expErr:    true,
		},
		"valid admin required": {
			admin:     []byte("invalid-address"),
			group:     1,
			threshold: sdk.ZeroDec(),
			timeout:   proto.Duration{Seconds: 1},
			expErr:    true,
		},
		"group required": {
			admin:     myAddr,
			threshold: sdk.ZeroDec(),
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
			threshold: sdk.ZeroDec(),
			expErr:    true,
		},
		"decision policy with invalid timeout": {
			admin:     myAddr,
			group:     1,
			threshold: sdk.ZeroDec(),
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
			threshold: sdk.NewDec(-1),
			timeout:   proto.Duration{Seconds: 1},
			expErr:    true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			m, err := NewMsgCreateGroupAccount(
				spec.admin,
				spec.group,
				"any comment",
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
	specs := map[string]struct {
		src    MsgCreateProposalRequest
		expErr bool
	}{
		"all good with minimum fields set": {
			src: MsgCreateProposalRequest{
				GroupAccount: []byte("valid--group-address"),
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address")},
			},
		},
		"group account required": {
			src: MsgCreateProposalRequest{
				Proposers: []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"proposers required": {
			src: MsgCreateProposalRequest{
				GroupAccount: []byte("valid--group-address"),
			},
			expErr: true,
		},
		"valid proposer address required": {
			src: MsgCreateProposalRequest{
				GroupAccount: []byte("valid--group-address"),
				Proposers:    []sdk.AccAddress{[]byte("invalid-member-address")},
			},
			expErr: true,
		},
		"no duplicate proposers": {
			src: MsgCreateProposalRequest{
				GroupAccount: []byte("valid--group-address"),
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address"), []byte("valid-member-address")},
			},
			expErr: true,
		},
		"empty proposer address not allowed": {
			src: MsgCreateProposalRequest{
				GroupAccount: []byte("valid--group-address"),
				Proposers:    []sdk.AccAddress{[]byte("valid-member-address"), nil, []byte("other-member-address")},
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
	specs := map[string]struct {
		src    MsgVoteRequest
		expErr bool
	}{
		"all good with minimum fields set": {
			src: MsgVoteRequest{
				Proposal: 1,
				Choice:   Choice_CHOICE_YES,
				Voters:   []sdk.AccAddress{[]byte("valid-member-address")},
			},
		},
		"proposal required": {
			src: MsgVoteRequest{
				Choice: Choice_CHOICE_YES,
				Voters: []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"choice required": {
			src: MsgVoteRequest{
				Proposal: 1,
				Voters:   []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"valid choice required": {
			src: MsgVoteRequest{
				Proposal: 1,
				Choice:   5,
				Voters:   []sdk.AccAddress{[]byte("valid-member-address")},
			},
			expErr: true,
		},
		"voter required": {
			src: MsgVoteRequest{
				Proposal: 1,
				Choice:   Choice_CHOICE_YES,
			},
			expErr: true,
		},
		"valid voter address required": {
			src: MsgVoteRequest{
				Proposal: 1,
				Choice:   Choice_CHOICE_YES,
				Voters:   []sdk.AccAddress{[]byte("invalid-member-address")},
			},
			expErr: true,
		},
		"duplicate voters": {
			src: MsgVoteRequest{
				Proposal: 1,
				Choice:   Choice_CHOICE_YES,
				Voters:   []sdk.AccAddress{[]byte("valid-member-address"), []byte("valid-member-address")},
			},
			expErr: true,
		},
		"empty voters address not allowed": {
			src: MsgVoteRequest{
				Proposal: 1,
				Choice:   Choice_CHOICE_YES,
				Voters:   []sdk.AccAddress{[]byte("valid-member-address"), nil, []byte("other-member-address")},
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
