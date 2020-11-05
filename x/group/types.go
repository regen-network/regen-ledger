package group

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/params"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/regen-network/regen-ledger/orm"
	proto "github.com/gogo/protobuf/proto"
	"github.com/gogo/protobuf/types"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"

	"time"
)

type GroupID uint64

func (p GroupID) Uint64() uint64 {
	return uint64(p)
}

func (p GroupID) Empty() bool {
	return p == 0
}

func (g GroupID) Bytes() []byte {
	return orm.EncodeSequence(uint64(g))
}

type ProposalID uint64

func (p ProposalID) Bytes() []byte {
	return orm.EncodeSequence(uint64(p))
}

func (p ProposalID) Uint64() uint64 {
	return uint64(p)
}
func (p ProposalID) Empty() bool {
	return p == 0
}

type DecisionPolicyResult struct {
	Allow bool
	Final bool
}

// DecisionPolicy is the persistent set of rules to determine the result of election on a proposal.
type DecisionPolicy interface {
	orm.Persistent
	orm.Validateable
	GetTimeout() types.Duration
	Allow(tally Tally, totalPower sdk.Dec, votingDuration time.Duration) (DecisionPolicyResult, error)
	Validate(g GroupMetadata) error
}

// Implements DecisionPolicy Interface
var _ DecisionPolicy = &ThresholdDecisionPolicy{}

// NewThresholdDecisionPolicy creates a threshold DecisionPolicy
func NewThresholdDecisionPolicy(threshold sdk.Dec, timeout types.Duration) DecisionPolicy {
	return &ThresholdDecisionPolicy{threshold, timeout}
}

// Allow allows a proposal to pass when the tally of yes votes equals or exceeds the threshold before the timeout.
func (p ThresholdDecisionPolicy) Allow(tally Tally, totalPower sdk.Dec, votingDuration time.Duration) (DecisionPolicyResult, error) {
	timeout, err := types.DurationFromProto(&p.Timeout)
	if err != nil {
		return DecisionPolicyResult{}, err
	}
	if timeout <= votingDuration {
		return DecisionPolicyResult{Allow: false, Final: true}, nil
	}
	if tally.YesCount.GTE(p.Threshold) {
		return DecisionPolicyResult{Allow: true, Final: true}, nil
	}
	undecided := totalPower.Sub(tally.TotalCounts())
	if tally.YesCount.Add(undecided).LT(p.Threshold) {
		return DecisionPolicyResult{Allow: false, Final: true}, nil
	}
	return DecisionPolicyResult{Allow: false, Final: false}, nil
}

// GetThreshold returns the policy threshold
func (p *ThresholdDecisionPolicy) GetThreshold() sdk.Dec {
	return p.Threshold
}

// Validate returns an error if policy threshold is greater than the total group weight
func (p *ThresholdDecisionPolicy) Validate(g GroupMetadata) error {
	if p.GetThreshold().GT(g.TotalWeight) {
		return errors.Wrap(ErrInvalid, "policy threshold should not be greater than the total group weight")
	}
	return nil
}

func (p ThresholdDecisionPolicy) ValidateBasic() error {
	if p.Threshold.IsNil() {
		return errors.Wrap(ErrEmpty, "threshold")
	}
	if p.Threshold.LT(sdk.OneDec()) {
		return errors.Wrap(ErrInvalid, "threshold")
	}
	timeout, err := types.DurationFromProto(&p.Timeout)
	if err != nil {
		return errors.Wrap(err, "timeout")
	}

	if timeout <= time.Nanosecond {
		return errors.Wrap(ErrInvalid, "timeout")
	}
	return nil
}

func (g GroupMember) NaturalKey() []byte {
	result := make([]byte, 8, 8+len(g.Member))
	copy(result[0:8], g.Group.Bytes())
	result = append(result, g.Member...)
	return result
}

func (g GroupAccountMetadata) NaturalKey() []byte {
	return g.GroupAccount
}

var _ orm.Validateable = GroupAccountMetadata{}

// NewGroupAccountMetadata creates a new GroupAccountMetadata instance
func NewGroupAccountMetadata(groupAccount sdk.AccAddress, group GroupID, admin sdk.AccAddress, comment string, version uint64, decisionPolicy DecisionPolicy) (GroupAccountMetadata, error) {
	p := GroupAccountMetadata{
		GroupAccount: groupAccount,
		Group:        group,
		Admin:        admin,
		Comment:      comment,
		Version:      version,
	}

	msg, ok := decisionPolicy.(proto.Message)
	if !ok {
		return GroupAccountMetadata{}, fmt.Errorf("%T does not implement proto.Message", decisionPolicy)
	}

	any, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return GroupAccountMetadata{}, err
	}

	p.DecisionPolicy = any
	return p, nil
}

func (g GroupAccountMetadata) GetDecisionPolicy() DecisionPolicy {
	decisionPolicy, ok := g.DecisionPolicy.GetCachedValue().(DecisionPolicy)
	if !ok {
		return nil
	}
	return decisionPolicy
}

func (g GroupAccountMetadata) ValidateBasic() error {
	if g.Admin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "admin")
	}
	if err := sdk.VerifyAddressFormat(g.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}
	if g.GroupAccount.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "group account")
	}
	if err := sdk.VerifyAddressFormat(g.GroupAccount); err != nil {
		return sdkerrors.Wrap(err, "group account")
	}

	if g.Group == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group")
	}
	if g.Version == 0 {
		return sdkerrors.Wrap(ErrEmpty, "version")
	}
	policy := g.GetDecisionPolicy()

	if policy == nil {
		return errors.Wrap(ErrEmpty, "policy")
	}
	if err := policy.ValidateBasic(); err != nil {
		return errors.Wrap(err, "policy")
	}
	return nil
}

func (v Vote) NaturalKey() []byte {
	result := make([]byte, 8, 8+len(v.Voter))
	copy(result[0:8], v.Proposal.Bytes())
	result = append(result, v.Voter...)
	return result
}

var _ orm.Validateable = Vote{}

func (v Vote) ValidateBasic() error {
	if len(v.Voter) == 0 {
		return errors.Wrap(ErrEmpty, "voter")
	}
	if err := sdk.VerifyAddressFormat(v.Voter); err != nil {
		return sdkerrors.Wrap(err, "voter")
	}
	if v.Proposal == 0 {
		return errors.Wrap(ErrEmpty, "proposal")
	}
	if v.Choice == Choice_UNKNOWN {
		return errors.Wrap(ErrEmpty, "choice")
	}
	if _, ok := Choice_name[int32(v.Choice)]; !ok {
		return errors.Wrap(ErrInvalid, "choice")
	}
	t, err := types.TimestampFromProto(&v.SubmittedAt)
	if err != nil {
		return errors.Wrap(err, "submitted at")
	}
	if t.IsZero() {
		return errors.Wrap(ErrEmpty, "submitted at")
	}
	return nil
}

const defaultMaxCommentLength = 255

// Parameter keys
var (
	ParamMaxCommentLength = []byte("MaxCommentLength")
)

// DefaultParams returns the default parameters for the group module.
func DefaultParams() Params {
	return Params{
		MaxCommentLength: defaultMaxCommentLength,
	}
}

func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		params.NewParamSetPair(ParamMaxCommentLength, &p.MaxCommentLength, noopValidator()),
	}
}
func (p Params) Validate() error {
	return nil
}

func noopValidator() paramtypes.ValueValidatorFn {
	return func(value interface{}) error { return nil }
}

var _ orm.Validateable = GroupMetadata{}

func (m GroupMetadata) ValidateBasic() error {
	if m.Group.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "group")
	}
	if m.Admin.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "admin")
	}
	if err := sdk.VerifyAddressFormat(m.Admin); err != nil {
		return sdkerrors.Wrap(err, "admin")
	}
	if m.TotalWeight.IsNil() || m.TotalWeight.LT(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalid, "total weight")
	}
	if m.Version == 0 {
		return sdkerrors.Wrap(ErrEmpty, "version")
	}
	return nil
}

var _ orm.Validateable = GroupMember{}

func (m GroupMember) ValidateBasic() error {
	if m.Group.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "group")
	}
	if m.Member.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "address")
	}
	if m.Weight.IsNil() || m.Weight.LTE(sdk.ZeroDec()) {
		return sdkerrors.Wrap(ErrInvalid, "member power")
	}
	if err := sdk.VerifyAddressFormat(m.Member); err != nil {
		return sdkerrors.Wrap(err, "address")
	}
	return nil
}

var _ orm.Validateable = ProposalBase{}

func (p ProposalBase) ValidateBasic() error {
	if p.GroupAccount.Empty() {
		return sdkerrors.Wrap(ErrEmpty, "group account")
	}
	if err := sdk.VerifyAddressFormat(p.GroupAccount); err != nil {
		return sdkerrors.Wrap(err, "group account")
	}
	if len(p.Proposers) == 0 {
		return sdkerrors.Wrap(ErrEmpty, "proposers")
	}
	if err := AccAddresses(p.Proposers).ValidateBasic(); err != nil {
		return sdkerrors.Wrap(err, "proposers")
	}
	if p.SubmittedAt.Seconds == 0 && p.SubmittedAt.Nanos == 0 {
		return sdkerrors.Wrap(ErrEmpty, "submitted at")
	}
	if p.GroupVersion == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group version")
	}
	if p.GroupAccountVersion == 0 {
		return sdkerrors.Wrap(ErrEmpty, "group account version")
	}
	if p.Status == ProposalStatusInvalid {
		return sdkerrors.Wrap(ErrEmpty, "status")
	}
	if _, ok := ProposalBase_Status_name[int32(p.Status)]; !ok {
		return sdkerrors.Wrap(ErrInvalid, "status")
	}
	if p.Result == ProposalResultInvalid {
		return sdkerrors.Wrap(ErrEmpty, "result")
	}
	if _, ok := ProposalBase_Result_name[int32(p.Result)]; !ok {
		return sdkerrors.Wrap(ErrInvalid, "result")
	}
	if p.ExecutorResult == ProposalExecutorResultInvalid {
		return sdkerrors.Wrap(ErrEmpty, "executor result")
	}
	if _, ok := ProposalBase_ExecutorResult_name[int32(p.ExecutorResult)]; !ok {
		return sdkerrors.Wrap(ErrInvalid, "executor result")
	}
	if err := p.VoteState.ValidateBasic(); err != nil {
		return errors.Wrap(err, "vote state")
	}
	if p.Timeout.Seconds == 0 && p.Timeout.Nanos == 0 {
		return sdkerrors.Wrap(ErrEmpty, "timeout")
	}
	return nil
}

func (t *Tally) Sub(vote Vote, weight sdk.Dec) error {
	if weight.LTE(sdk.ZeroDec()) {
		return errors.Wrap(ErrInvalid, "weight must be greater than 0")
	}
	switch vote.Choice {
	case Choice_YES:
		t.YesCount = t.YesCount.Sub(weight)
	case Choice_NO:
		t.NoCount = t.NoCount.Sub(weight)
	case Choice_ABSTAIN:
		t.AbstainCount = t.AbstainCount.Sub(weight)
	case Choice_VETO:
		t.VetoCount = t.VetoCount.Sub(weight)
	default:
		return errors.Wrapf(ErrInvalid, "unknown choice %s", vote.Choice.String())
	}
	return nil
}

func (t *Tally) Add(vote Vote, weight sdk.Dec) error {
	if weight.LTE(sdk.ZeroDec()) {
		return errors.Wrap(ErrInvalid, "weight must be greater than 0")
	}
	switch vote.Choice {
	case Choice_YES:
		t.YesCount = t.YesCount.Add(weight)
	case Choice_NO:
		t.NoCount = t.NoCount.Add(weight)
	case Choice_ABSTAIN:
		t.AbstainCount = t.AbstainCount.Add(weight)
	case Choice_VETO:
		t.VetoCount = t.VetoCount.Add(weight)
	default:
		return errors.Wrapf(ErrInvalid, "unknown choice %s", vote.Choice.String())
	}
	return nil
}

// TotalCounts is the sum of all weights.
func (t Tally) TotalCounts() sdk.Dec {
	return t.YesCount.Add(t.NoCount).Add(t.AbstainCount).Add(t.VetoCount)
}

func (t Tally) ValidateBasic() error {
	switch {
	case t.YesCount.IsNil():
		return errors.Wrap(ErrInvalid, "yes count nil")
	case t.YesCount.LT(sdk.ZeroDec()):
		return errors.Wrap(ErrInvalid, "yes count negative")
	case t.NoCount.IsNil():
		return errors.Wrap(ErrInvalid, "no count nil")
	case t.NoCount.LT(sdk.ZeroDec()):
		return errors.Wrap(ErrInvalid, "no count negative")
	case t.AbstainCount.IsNil():
		return errors.Wrap(ErrInvalid, "abstain count nil")
	case t.AbstainCount.LT(sdk.ZeroDec()):
		return errors.Wrap(ErrInvalid, "abstain count negative")
	case t.VetoCount.IsNil():
		return errors.Wrap(ErrInvalid, "veto count nil")
	case t.VetoCount.LT(sdk.ZeroDec()):
		return errors.Wrap(ErrInvalid, "veto count negative")
	}
	return nil
}

func (g GenesisState) String() string {
	out, _ := yaml.Marshal(g)
	return string(out)
}
