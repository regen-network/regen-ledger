package group

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	proto "github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestThresholdDecisionPolicy(t *testing.T) {
	specs := map[string]struct {
		srcPolicy         ThresholdDecisionPolicy
		srcTally          Tally
		srcTotalPower     sdk.Dec
		srcVotingDuration time.Duration
		expResult         DecisionPolicyResult
		expErr            error
	}{
		"accept when yes count greater than threshold": {
			srcPolicy: ThresholdDecisionPolicy{
				Threshold: sdk.OneDec(),
				Timeout:   proto.Duration{Seconds: 1},
			},
			srcTally:          Tally{YesCount: sdk.NewDec(2)},
			srcTotalPower:     sdk.NewDec(3),
			srcVotingDuration: time.Millisecond,
			expResult:         DecisionPolicyResult{Allow: true, Final: true},
		},
		"accept when yes count equal to threshold": {
			srcPolicy: ThresholdDecisionPolicy{
				Threshold: sdk.OneDec(),
				Timeout:   proto.Duration{Seconds: 1},
			},
			srcTally:          Tally{YesCount: sdk.OneDec(), NoCount: sdk.ZeroDec(), AbstainCount: sdk.ZeroDec(), VetoCount: sdk.ZeroDec()},
			srcTotalPower:     sdk.NewDec(3),
			srcVotingDuration: time.Millisecond,
			expResult:         DecisionPolicyResult{Allow: true, Final: true},
		},
		"reject when yes count lower to threshold": {
			srcPolicy: ThresholdDecisionPolicy{
				Threshold: sdk.OneDec(),
				Timeout:   proto.Duration{Seconds: 1},
			},
			srcTally:          Tally{YesCount: sdk.ZeroDec(), NoCount: sdk.ZeroDec(), AbstainCount: sdk.ZeroDec(), VetoCount: sdk.ZeroDec()},
			srcTotalPower:     sdk.NewDec(3),
			srcVotingDuration: time.Millisecond,
			expResult:         DecisionPolicyResult{Allow: false, Final: false},
		},
		"reject as final when remaining votes can't cross threshold": {
			srcPolicy: ThresholdDecisionPolicy{
				Threshold: sdk.NewDec(2),
				Timeout:   proto.Duration{Seconds: 1},
			},
			srcTally:          Tally{YesCount: sdk.ZeroDec(), NoCount: sdk.NewDec(2), AbstainCount: sdk.ZeroDec(), VetoCount: sdk.ZeroDec()},
			srcTotalPower:     sdk.NewDec(3),
			srcVotingDuration: time.Millisecond,
			expResult:         DecisionPolicyResult{Allow: false, Final: true},
		},
		"expired when on timeout": {
			srcPolicy: ThresholdDecisionPolicy{
				Threshold: sdk.OneDec(),
				Timeout:   proto.Duration{Seconds: 1},
			},
			srcTally:          Tally{YesCount: sdk.NewDec(2)},
			srcTotalPower:     sdk.NewDec(3),
			srcVotingDuration: time.Second,
			expResult:         DecisionPolicyResult{Allow: false, Final: true},
		},
		"expired when after timeout": {
			srcPolicy: ThresholdDecisionPolicy{
				Threshold: sdk.OneDec(),
				Timeout:   proto.Duration{Seconds: 1},
			},
			srcTally:          Tally{YesCount: sdk.NewDec(2)},
			srcTotalPower:     sdk.NewDec(3),
			srcVotingDuration: time.Second + time.Nanosecond,
			expResult:         DecisionPolicyResult{Allow: false, Final: true},
		},
		"abstain has no impact": {
			srcPolicy: ThresholdDecisionPolicy{
				Threshold: sdk.OneDec(),
				Timeout:   proto.Duration{Seconds: 1},
			},
			srcTally:          Tally{YesCount: sdk.ZeroDec(), NoCount: sdk.ZeroDec(), AbstainCount: sdk.OneDec(), VetoCount: sdk.ZeroDec()},
			srcTotalPower:     sdk.NewDec(3),
			srcVotingDuration: time.Millisecond,
			expResult:         DecisionPolicyResult{Allow: false, Final: false},
		},
		"veto same as no": {
			srcPolicy: ThresholdDecisionPolicy{
				Threshold: sdk.OneDec(),
				Timeout:   proto.Duration{Seconds: 1},
			},
			srcTally:          Tally{YesCount: sdk.ZeroDec(), NoCount: sdk.ZeroDec(), AbstainCount: sdk.ZeroDec(), VetoCount: sdk.NewDec(2)},
			srcTotalPower:     sdk.NewDec(3),
			srcVotingDuration: time.Millisecond,
			expResult:         DecisionPolicyResult{Allow: false, Final: false},
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			res, err := spec.srcPolicy.Allow(spec.srcTally, spec.srcTotalPower, spec.srcVotingDuration)
			if spec.expErr != nil {
				require.Error(t, err)
				return
			}
			require.NoError(t, err)
			assert.Equal(t, spec.expResult, res)
		})
	}
}

func TestThresholdDecisionPolicyValidation(t *testing.T) {
	maxSeconds := int64(10000 * 365.25 * 24 * 60 * 60)
	specs := map[string]struct {
		src    ThresholdDecisionPolicy
		expErr bool
	}{
		"all good": {src: ThresholdDecisionPolicy{
			Threshold: sdk.OneDec(),
			Timeout:   proto.Duration{Seconds: 1},
		}},
		"threshold missing": {src: ThresholdDecisionPolicy{
			Timeout: proto.Duration{Seconds: 1},
		},
			expErr: true,
		},
		"timeout missing": {src: ThresholdDecisionPolicy{
			Threshold: sdk.OneDec(),
		},
			expErr: true,
		},
		"duration out of limit": {src: ThresholdDecisionPolicy{
			Threshold: sdk.OneDec(),
			Timeout:   proto.Duration{Seconds: maxSeconds + 1},
		},
			expErr: true,
		},
		"no negative thresholds": {src: ThresholdDecisionPolicy{
			Threshold: sdk.NewDec(-1),
			Timeout:   proto.Duration{Seconds: 1},
		},
			expErr: true,
		},
		"no empty thresholds": {src: ThresholdDecisionPolicy{
			Timeout: proto.Duration{Seconds: 1},
		},
			expErr: true,
		},
		"no zero thresholds": {src: ThresholdDecisionPolicy{
			Timeout:   proto.Duration{Seconds: 1},
			Threshold: sdk.ZeroDec(),
		},
			expErr: true,
		},
		"no negative timeouts": {src: ThresholdDecisionPolicy{
			Threshold: sdk.OneDec(),
			Timeout:   proto.Duration{Seconds: -1},
		},
			expErr: true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			err := spec.src.ValidateBasic()
			assert.Equal(t, spec.expErr, err != nil, err)
		})
	}
}

func TestVoteNaturalKey(t *testing.T) {
	v := Vote{
		Proposal: 1,
		Voter:    []byte{0xff, 0xfe},
	}
	assert.Equal(t, []byte{0, 0, 0, 0, 0, 0, 0, 1, 0xff, 0xfe}, v.NaturalKey())
}

func TestGroupMetadataValidation(t *testing.T) {
	specs := map[string]struct {
		src    GroupMetadata
		expErr bool
	}{
		"all good": {
			src: GroupMetadata{
				Group:       1,
				Admin:       []byte("valid--admin-address"),
				Comment:     "any",
				Version:     1,
				TotalWeight: sdk.ZeroDec(),
			},
		},
		"invalid group": {
			src: GroupMetadata{
				Admin:       []byte("valid--admin-address"),
				Comment:     "any",
				Version:     1,
				TotalWeight: sdk.ZeroDec(),
			},
			expErr: true,
		},
		"invalid admin": {
			src: GroupMetadata{
				Group:       1,
				Admin:       []byte(""),
				Comment:     "any",
				Version:     1,
				TotalWeight: sdk.ZeroDec(),
			},
			expErr: true,
		},
		"invalid version": {
			src: GroupMetadata{
				Group:       1,
				Admin:       []byte("valid--admin-address"),
				Comment:     "any",
				TotalWeight: sdk.ZeroDec(),
			},
			expErr: true,
		},
		"unset total weight": {
			src: GroupMetadata{
				Group:   1,
				Admin:   []byte("valid--admin-address"),
				Comment: "any",
				Version: 1,
			},
			expErr: true,
		},
		"negative total weight": {
			src: GroupMetadata{
				Group:       1,
				Admin:       []byte("valid--admin-address"),
				Comment:     "any",
				Version:     1,
				TotalWeight: sdk.NewDec(-1),
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

func TestGroupMemberValidation(t *testing.T) {
	specs := map[string]struct {
		src    GroupMember
		expErr bool
	}{
		"all good": {
			src: GroupMember{
				Group:   1,
				Member:  []byte("valid-member-address"),
				Weight:  sdk.OneDec(),
				Comment: "any",
			},
		},
		"invalid group": {
			src: GroupMember{
				Group:   0,
				Member:  []byte("valid-member-address"),
				Weight:  sdk.OneDec(),
				Comment: "any",
			},
			expErr: true,
		},
		"invalid address": {
			src: GroupMember{
				Group:   1,
				Member:  []byte("invalid-member-address"),
				Weight:  sdk.OneDec(),
				Comment: "any",
			},
			expErr: true,
		},
		"empy address": {
			src: GroupMember{
				Group:   1,
				Weight:  sdk.OneDec(),
				Comment: "any",
			},
			expErr: true,
		},
		"invalid weight": {
			src: GroupMember{
				Group:   1,
				Member:  []byte("valid-member-address"),
				Weight:  sdk.ZeroDec(),
				Comment: "any",
			},
			expErr: true,
		},
		"nil weight": {
			src: GroupMember{
				Group:   1,
				Member:  []byte("valid-member-address"),
				Comment: "any",
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

func TestGroupAccountMetadata(t *testing.T) {
	specs := map[string]struct {
		groupAccount sdk.AccAddress
		group        ID
		admin        sdk.AccAddress
		comment      string
		version      uint64
		threshold    sdk.Dec
		timeout      proto.Duration
		expErr       bool
	}{
		"all good": {
			group:        1,
			groupAccount: []byte("valid--group-address"),
			admin:        []byte("valid--admin-address"),
			comment:      "any",
			version:      1,
			threshold:    sdk.OneDec(),
			timeout:      proto.Duration{Seconds: 1},
		},
		"invalid group": {
			group:        0,
			groupAccount: []byte("valid--group-address"),
			admin:        []byte("valid--admin-address"),
			comment:      "any",
			version:      1,
			threshold:    sdk.OneDec(),
			timeout:      proto.Duration{Seconds: 1},
			expErr:       true,
		},
		"invalid group account address": {
			group:        1,
			groupAccount: []byte("any-invalid-group-address"),
			admin:        []byte("valid--admin-address"),
			comment:      "any",
			version:      1,
			threshold:    sdk.OneDec(),
			timeout:      proto.Duration{Seconds: 1},
			expErr:       true,
		},
		"empty group account address": {
			group:     1,
			admin:     []byte("valid--admin-address"),
			comment:   "any",
			version:   1,
			threshold: sdk.OneDec(),
			timeout:   proto.Duration{Seconds: 1},
			expErr:    true,
		},
		"empty admin account address": {
			group:        1,
			groupAccount: []byte("valid--group-address"),
			comment:      "any",
			version:      1,
			threshold:    sdk.OneDec(),
			timeout:      proto.Duration{Seconds: 1},
			expErr:       true,
		},
		"invalid admin account address": {
			group:        1,
			groupAccount: []byte("valid--group-address"),
			admin:        []byte("any-invalid-admin-address"),
			comment:      "any",
			version:      1,
			threshold:    sdk.OneDec(),
			timeout:      proto.Duration{Seconds: 1},
			expErr:       true,
		},
		"empty version number": {
			group:        1,
			groupAccount: []byte("valid--group-address"),
			admin:        []byte("valid--admin-address"),
			comment:      "any",
			threshold:    sdk.OneDec(),
			timeout:      proto.Duration{Seconds: 1},
			expErr:       true,
		},
		"missing decision policy": {
			group:        1,
			groupAccount: []byte("valid--group-address"),
			admin:        []byte("valid--admin-address"),
			comment:      "any",
			version:      1,
			expErr:       true,
		},
		"missing decision policy timeout": {
			group:        1,
			groupAccount: []byte("valid--group-address"),
			admin:        []byte("valid--admin-address"),
			comment:      "any",
			version:      1,
			threshold:    sdk.OneDec(),
			expErr:       true,
		},
		"decision policy with invalid timeout": {
			group:        1,
			groupAccount: []byte("valid--group-address"),
			admin:        []byte("valid--admin-address"),
			comment:      "any",
			version:      1,
			threshold:    sdk.OneDec(),
			timeout:      proto.Duration{Seconds: -1},
			expErr:       true,
		},
		"missing decision policy threshold": {
			group:        1,
			groupAccount: []byte("valid--group-address"),
			admin:        []byte("valid--admin-address"),
			comment:      "any",
			version:      1,
			timeout:      proto.Duration{Seconds: 1},
			expErr:       true,
		},
		"decision policy with negative threshold": {
			group:        1,
			groupAccount: []byte("valid--group-address"),
			admin:        []byte("valid--admin-address"),
			comment:      "any",
			version:      1,
			threshold:    sdk.NewDec(-1),
			timeout:      proto.Duration{Seconds: 1},
			expErr:       true,
		},
		"decision policy with zero threshold": {
			group:        1,
			groupAccount: []byte("valid--group-address"),
			admin:        []byte("valid--admin-address"),
			comment:      "any",
			version:      1,
			threshold:    sdk.ZeroDec(),
			timeout:      proto.Duration{Seconds: 1},
			expErr:       true,
		},
	}
	for msg, spec := range specs {
		t.Run(msg, func(t *testing.T) {
			m, err := NewGroupAccountMetadata(
				spec.groupAccount,
				spec.group,
				spec.admin,
				spec.comment,
				spec.version,
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
