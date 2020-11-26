 <!-- This file is auto-generated. Please do not modify it yourself. -->
# Protobuf Documentation
<a name="top"></a>

## Table of Contents

- [regen/group/v1alpha1/events.proto](#regen/group/v1alpha1/events.proto)
    - [EventCreateGroup](#regen.group.v1alpha1.EventCreateGroup)
    - [EventCreateGroupAccount](#regen.group.v1alpha1.EventCreateGroupAccount)
    - [EventUpdateGroup](#regen.group.v1alpha1.EventUpdateGroup)
    - [EventUpdateGroupAccount](#regen.group.v1alpha1.EventUpdateGroupAccount)
  
- [regen/group/v1alpha1/genesis.proto](#regen/group/v1alpha1/genesis.proto)
    - [GenesisState](#regen.group.v1alpha1.GenesisState)
    - [Params](#regen.group.v1alpha1.Params)
  
- [regen/group/v1alpha1/types.proto](#regen/group/v1alpha1/types.proto)
    - [GroupAccountMetadata](#regen.group.v1alpha1.GroupAccountMetadata)
    - [GroupMember](#regen.group.v1alpha1.GroupMember)
    - [GroupMetadata](#regen.group.v1alpha1.GroupMetadata)
    - [Member](#regen.group.v1alpha1.Member)
    - [Proposal](#regen.group.v1alpha1.Proposal)
    - [Tally](#regen.group.v1alpha1.Tally)
    - [ThresholdDecisionPolicy](#regen.group.v1alpha1.ThresholdDecisionPolicy)
    - [Vote](#regen.group.v1alpha1.Vote)
  
    - [Choice](#regen.group.v1alpha1.Choice)
    - [Proposal.ExecutorResult](#regen.group.v1alpha1.Proposal.ExecutorResult)
    - [Proposal.Result](#regen.group.v1alpha1.Proposal.Result)
    - [Proposal.Status](#regen.group.v1alpha1.Proposal.Status)
  
- [regen/group/v1alpha1/tx.proto](#regen/group/v1alpha1/tx.proto)
    - [MsgCreateGroupAccountRequest](#regen.group.v1alpha1.MsgCreateGroupAccountRequest)
    - [MsgCreateGroupAccountResponse](#regen.group.v1alpha1.MsgCreateGroupAccountResponse)
    - [MsgCreateGroupRequest](#regen.group.v1alpha1.MsgCreateGroupRequest)
    - [MsgCreateGroupResponse](#regen.group.v1alpha1.MsgCreateGroupResponse)
    - [MsgCreateProposalRequest](#regen.group.v1alpha1.MsgCreateProposalRequest)
    - [MsgCreateProposalResponse](#regen.group.v1alpha1.MsgCreateProposalResponse)
    - [MsgExecRequest](#regen.group.v1alpha1.MsgExecRequest)
    - [MsgExecResponse](#regen.group.v1alpha1.MsgExecResponse)
    - [MsgUpdateGroupAccountAdminRequest](#regen.group.v1alpha1.MsgUpdateGroupAccountAdminRequest)
    - [MsgUpdateGroupAccountAdminResponse](#regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse)
    - [MsgUpdateGroupAccountCommentRequest](#regen.group.v1alpha1.MsgUpdateGroupAccountCommentRequest)
    - [MsgUpdateGroupAccountCommentResponse](#regen.group.v1alpha1.MsgUpdateGroupAccountCommentResponse)
    - [MsgUpdateGroupAccountDecisionPolicyRequest](#regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyRequest)
    - [MsgUpdateGroupAccountDecisionPolicyResponse](#regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse)
    - [MsgUpdateGroupAdminRequest](#regen.group.v1alpha1.MsgUpdateGroupAdminRequest)
    - [MsgUpdateGroupAdminResponse](#regen.group.v1alpha1.MsgUpdateGroupAdminResponse)
    - [MsgUpdateGroupCommentRequest](#regen.group.v1alpha1.MsgUpdateGroupCommentRequest)
    - [MsgUpdateGroupCommentResponse](#regen.group.v1alpha1.MsgUpdateGroupCommentResponse)
    - [MsgUpdateGroupMembersRequest](#regen.group.v1alpha1.MsgUpdateGroupMembersRequest)
    - [MsgUpdateGroupMembersResponse](#regen.group.v1alpha1.MsgUpdateGroupMembersResponse)
    - [MsgVoteRequest](#regen.group.v1alpha1.MsgVoteRequest)
    - [MsgVoteResponse](#regen.group.v1alpha1.MsgVoteResponse)
  
    - [Msg](#regen.group.v1alpha1.Msg)
  
- [Scalar Value Types](#scalar-value-types)



<a name="regen/group/v1alpha1/events.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/group/v1alpha1/events.proto



<a name="regen.group.v1alpha1.EventCreateGroup"></a>

### EventCreateGroup
EventCreateGroup is an event emitted when a group is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [string](#string) |  | group is the unique ID of the group. |
| admin | [string](#string) |  | admin is the admin address of the group. |






<a name="regen.group.v1alpha1.EventCreateGroupAccount"></a>

### EventCreateGroupAccount
EventCreateGroupAccount is an event emitted when a group account is created.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_account | [string](#string) |  | group_account is the address of the group account. |
| admin | [string](#string) |  | admin is the admin address of the group. |






<a name="regen.group.v1alpha1.EventUpdateGroup"></a>

### EventUpdateGroup
EventUpdateGroup is an event emitted when a group is updated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [string](#string) |  | group is the unique ID of the group. |
| admin | [string](#string) |  | admin is the admin address of the group. |






<a name="regen.group.v1alpha1.EventUpdateGroupAccount"></a>

### EventUpdateGroupAccount
EventUpdateGroupAccount is an event emitted when a group account is updated.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_account | [string](#string) |  | group_account is the address of the group account. |
| admin | [string](#string) |  | admin is the admin address of the group. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/group/v1alpha1/genesis.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/group/v1alpha1/genesis.proto



<a name="regen.group.v1alpha1.GenesisState"></a>

### GenesisState
GenesisState defines the group module's genesis state.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| params | [Params](#regen.group.v1alpha1.Params) |  | params are the genesis state Params. |






<a name="regen.group.v1alpha1.Params"></a>

### Params
Params defines the set of configurable parameters.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| max_comment_length | [uint32](#uint32) |  | max_comment_length defines the maximum comment length of the group module genesis params. |





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/group/v1alpha1/types.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/group/v1alpha1/types.proto



<a name="regen.group.v1alpha1.GroupAccountMetadata"></a>

### GroupAccountMetadata
GroupAccountMetadata specifies group account metadata.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_account | [bytes](#bytes) |  | group_account is the group account address. |
| group | [uint64](#uint64) |  | group is the unique ID of the group. |
| admin | [bytes](#bytes) |  | admin is the account address of the group admin. |
| comment | [string](#string) |  | comment is the group account's comment. |
| version | [uint64](#uint64) |  | version is used to track changes to a group's GroupAccountMetadata structure that would create a different result on a running proposal. |
| decision_policy | [google.protobuf.Any](#google.protobuf.Any) |  | decision_policy specifies the group account's decision policy. |






<a name="regen.group.v1alpha1.GroupMember"></a>

### GroupMember
GroupMember represents the relationship between a group and a member.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [uint64](#uint64) |  | group is the unique ID of the group. |
| member | [bytes](#bytes) |  | member is the account address of the group member. todo: @aaronc field has different name in `Member.address`. Can we unify this? |
| weight | [string](#string) |  | weight is the power of the group member. todo: @aaronc it is `Member.power`. Can we unify this? |
| comment | [string](#string) |  | comment is the member's comment.

TODO: @aaronc was comment missing by intention? |






<a name="regen.group.v1alpha1.GroupMetadata"></a>

### GroupMetadata
GroupMetadata represents a group's metadata


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [uint64](#uint64) |  | group is the unique ID of the group. |
| admin | [bytes](#bytes) |  | admin is the account address of the group's admin. |
| comment | [string](#string) |  | comment is the group's comment. |
| version | [uint64](#uint64) |  | version is used to track changes to a group's membership structure that would break existing proposals. Whenever any members power is changed, or any member is added or removed this version is incremented and will cause proposals based on older versions of this group to fail |
| total_weight | [string](#string) |  | total_weight is the sum of the group members' powers. |






<a name="regen.group.v1alpha1.Member"></a>

### Member
Member represents a group member with an account address,
non-zero power and a comment.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| address | [bytes](#bytes) |  | address is the member's account address. |
| power | [string](#string) |  | power is the member's power that should be greater than 0. |
| comment | [string](#string) |  | comment is the member's comment. |






<a name="regen.group.v1alpha1.Proposal"></a>

### Proposal
Proposal defines the core field members of a group proposal.
Any member of a group can submit a proposal for a group account to decide upon.
A proposal consists of a set of `sdk.Msg`s that will be executed if the proposal
passes as well as any comment associated with the proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_account | [bytes](#bytes) |  | group_account is the group account address. |
| comment | [string](#string) |  | comment is the proposal's comment. |
| proposers | [bytes](#bytes) | repeated | proposers are the account addresses of the proposers. |
| submitted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | submitted_at is a timestamp specifying when a proposal was submitted. |
| group_version | [uint64](#uint64) |  | group_version tracks the version of the group that this proposal corresponds to. When group membership is changed, existing proposals for prior group versions will become invalid. |
| group_account_version | [uint64](#uint64) |  | group_account_version tracks the version of the group account that this proposal corresponds to. When a decision policy is changed, an existing proposals for prior policy versions will become invalid. |
| status | [Proposal.Status](#regen.group.v1alpha1.Proposal.Status) |  | Status represents the high level position in the life cycle of the proposal. Initial value is Submitted. |
| result | [Proposal.Result](#regen.group.v1alpha1.Proposal.Result) |  | result is the final result based on the votes and election rule. Initial value is Undefined. The result is persisted so that clients can always rely on this state and not have to replicate the logic. |
| vote_state | [Tally](#regen.group.v1alpha1.Tally) |  | vote_state contains the sums of all weighted votes for this proposal. |
| timeout | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | timeout is the timestamp of the block where the proposal execution times out. Header times of the votes and execution messages must be before this end time to be included in the election. After the timeout timestamp the proposal can not be executed anymore and should be considered pending delete. |
| executor_result | [Proposal.ExecutorResult](#regen.group.v1alpha1.Proposal.ExecutorResult) |  | executor_result is the final result based on the votes and election rule. Initial value is NotRun. |
| msgs | [google.protobuf.Any](#google.protobuf.Any) | repeated | msgs is a list of Msgs that will be executed if the proposal passes. |






<a name="regen.group.v1alpha1.Tally"></a>

### Tally
Tally represents the sum of weighted votes.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| yes_count | [string](#string) |  | yes_count is the weighted sum of yes votes. |
| no_count | [string](#string) |  | no_count is the weighted sum of no votes. |
| abstain_count | [string](#string) |  | abstain_count is the weighted sum of abstainers |
| veto_count | [string](#string) |  | veto_count is the weighted sum of vetoes. |






<a name="regen.group.v1alpha1.ThresholdDecisionPolicy"></a>

### ThresholdDecisionPolicy
ThresholdDecisionPolicy implements the DecisionPolicy interface


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| threshold | [string](#string) |  | threshold is a fix weight value that must be met or exceeded for a proposal to succeed. |
| timeout | [google.protobuf.Duration](#google.protobuf.Duration) |  | timeout is the duration from submission of a proposal to the end of voting period Within this times votes and exec messages can be submitted. |






<a name="regen.group.v1alpha1.Vote"></a>

### Vote
Vote represents a vote for a proposal.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| proposal | [uint64](#uint64) |  | proposal is the unique ID of the proposal. |
| voter | [bytes](#bytes) |  | voter is the account address of the voter. |
| choice | [Choice](#regen.group.v1alpha1.Choice) |  | choice is the voter's choice on the proposal. |
| comment | [string](#string) |  | comment's is the vote's comment. |
| submitted_at | [google.protobuf.Timestamp](#google.protobuf.Timestamp) |  | submitted_at is the timestamp when the vote was submitted. |





 <!-- end messages -->


<a name="regen.group.v1alpha1.Choice"></a>

### Choice
Choice defines available types of choices for voting.

| Name | Number | Description |
| ---- | ------ | ----------- |
| CHOICE_UNSPECIFIED | 0 | CHOICE_UNSPECIFIED defines a no-op voting choice. |
| CHOICE_NO | 1 | CHOICE_NO defines a no voting choice. |
| CHOICE_YES | 2 | CHOICE_YES defines a yes voting choice. |
| CHOICE_ABSTAIN | 3 | CHOICE_ABSTAIN defines an abstaining voting choice. |
| CHOICE_VETO | 4 | CHOICE_VETO defines a voting choice with veto. |



<a name="regen.group.v1alpha1.Proposal.ExecutorResult"></a>

### Proposal.ExecutorResult
ExecutorResult defines types of proposal executor results.

| Name | Number | Description |
| ---- | ------ | ----------- |
| EXECUTOR_RESULT_UNSPECIFIED | 0 | An empty value is not allowed. |
| EXECUTOR_RESULT_NOT_RUN | 1 | We have not yet run the executor. |
| EXECUTOR_RESULT_SUCCESS | 2 | The executor was successful and proposed action updated state. |
| EXECUTOR_RESULT_FAILURE | 3 | The executor returned an error and proposed action didn't update state. |



<a name="regen.group.v1alpha1.Proposal.Result"></a>

### Proposal.Result
Result defines types of proposal results.

| Name | Number | Description |
| ---- | ------ | ----------- |
| RESULT_UNSPECIFIED | 0 | An empty value is invalid and not allowed |
| RESULT_UNDEFINED | 1 | Until a final tally has happened the status is undefined |
| RESULT_ACCEPTED | 2 | Final result of the tally |
| RESULT_REJECTED | 3 | Final result of the tally |



<a name="regen.group.v1alpha1.Proposal.Status"></a>

### Proposal.Status
Status defines proposal statuses.

| Name | Number | Description |
| ---- | ------ | ----------- |
| STATUS_UNSPECIFIED | 0 | An empty value is invalid and not allowed. |
| STATUS_SUBMITTED | 1 | Initial status of a proposal when persisted. |
| STATUS_CLOSED | 2 | Final status of a proposal when the final tally was executed. |
| STATUS_ABORTED | 3 | Final status of a proposal when the group was modified before the final tally. |


 <!-- end enums -->

 <!-- end HasExtensions -->

 <!-- end services -->



<a name="regen/group/v1alpha1/tx.proto"></a>
<p align="right"><a href="#top">Top</a></p>

## regen/group/v1alpha1/tx.proto



<a name="regen.group.v1alpha1.MsgCreateGroupAccountRequest"></a>

### MsgCreateGroupAccountRequest
MsgCreateGroupAccountRequest is the Msg/CreateGroupAccount request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [bytes](#bytes) |  | admin is the account address of the group admin. |
| group | [uint64](#uint64) |  | group is the unique ID of the group. |
| comment | [string](#string) |  | comment is the group account's comment. |
| decision_policy | [google.protobuf.Any](#google.protobuf.Any) |  | decision_policy specifies the group account's decision policy. |






<a name="regen.group.v1alpha1.MsgCreateGroupAccountResponse"></a>

### MsgCreateGroupAccountResponse
MsgCreateGroupAccountResponse is the Msg/CreateGroupAccount response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_account | [bytes](#bytes) |  | group_account is the account address of the newly created group account. |






<a name="regen.group.v1alpha1.MsgCreateGroupRequest"></a>

### MsgCreateGroupRequest
MsgCreateGroupRequest is the Msg/CreateGroup request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [bytes](#bytes) |  | admin is the account address of the group admin. |
| members | [Member](#regen.group.v1alpha1.Member) | repeated | members defines the group members. |
| comment | [string](#string) |  | group is the group's comment. |






<a name="regen.group.v1alpha1.MsgCreateGroupResponse"></a>

### MsgCreateGroupResponse
MsgCreateGroupResponse is the Msg/CreateGroup response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [uint64](#uint64) |  | group is the unique ID of the newly created group. |






<a name="regen.group.v1alpha1.MsgCreateProposalRequest"></a>

### MsgCreateProposalRequest
MsgCreateProposalRequest is the Msg/CreateProposal request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_account | [bytes](#bytes) |  | group_account is the group account address. |
| proposers | [bytes](#bytes) | repeated | proposers are the account addresses of the proposers. Proposers signatures will be counted as yes votes. |
| comment | [string](#string) |  | comment is the proposal's comment. |
| msgs | [google.protobuf.Any](#google.protobuf.Any) | repeated | msgs is a list of Msgs that will be executed if the proposal passes. |






<a name="regen.group.v1alpha1.MsgCreateProposalResponse"></a>

### MsgCreateProposalResponse
MsgCreateProposalResponse is the Msg/CreateProposal response type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| proposal | [uint64](#uint64) |  | proposal is the unique ID of the proposal. |






<a name="regen.group.v1alpha1.MsgExecRequest"></a>

### MsgExecRequest
MsgExecRequest is the Msg/Exec request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| proposal | [uint64](#uint64) |  | proposal is the unique ID of the proposal. |
| signer | [bytes](#bytes) |  | signer is the account address used to execute the proposal. |






<a name="regen.group.v1alpha1.MsgExecResponse"></a>

### MsgExecResponse
MsgExecResponse is the Msg/Exec request type.






<a name="regen.group.v1alpha1.MsgUpdateGroupAccountAdminRequest"></a>

### MsgUpdateGroupAccountAdminRequest
MsgUpdateGroupAccountAdminRequest is the Msg/UpdateGroupAccountAdmin request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [bytes](#bytes) |  | admin is the account address of the group admin. |
| group_account | [bytes](#bytes) |  | group_account is the group account address. |
| new_admin | [bytes](#bytes) |  | new_admin is the new group account admin. |






<a name="regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse"></a>

### MsgUpdateGroupAccountAdminResponse
MsgUpdateGroupAccountAdminResponse is the Msg/UpdateGroupAccountAdmin response type.






<a name="regen.group.v1alpha1.MsgUpdateGroupAccountCommentRequest"></a>

### MsgUpdateGroupAccountCommentRequest
MsgUpdateGroupAccountCommentRequest is the Msg/UpdateGroupAccountComment request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [bytes](#bytes) |  | admin is the account address of the group admin. |
| group_account | [bytes](#bytes) |  | group_account is the group account address. |
| comment | [string](#string) |  | comment is the updated group account comment. |






<a name="regen.group.v1alpha1.MsgUpdateGroupAccountCommentResponse"></a>

### MsgUpdateGroupAccountCommentResponse
MsgUpdateGroupAccountCommentResponse is the Msg/UpdateGroupAccountComment response type.






<a name="regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyRequest"></a>

### MsgUpdateGroupAccountDecisionPolicyRequest
MsgUpdateGroupAccountDecisionPolicyRequest is the Msg/UpdateGroupAccountDecisionPolicy request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [bytes](#bytes) |  | admin is the account address of the group admin. |
| group_account | [bytes](#bytes) |  | group_account is the group account address. |
| decision_policy | [google.protobuf.Any](#google.protobuf.Any) |  | decision_policy is the updated group account decision policy. |






<a name="regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse"></a>

### MsgUpdateGroupAccountDecisionPolicyResponse
MsgUpdateGroupAccountDecisionPolicyResponse is the Msg/UpdateGroupAccountDecisionPolicy response type.






<a name="regen.group.v1alpha1.MsgUpdateGroupAdminRequest"></a>

### MsgUpdateGroupAdminRequest
MsgUpdateGroupAdminRequest is the Msg/UpdateGroupAdmin request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [bytes](#bytes) |  | admin is the current account address of the group admin. |
| group | [uint64](#uint64) |  | group is the unique ID of the group. |
| new_admin | [bytes](#bytes) |  | new_admin is the group new admin account address. |






<a name="regen.group.v1alpha1.MsgUpdateGroupAdminResponse"></a>

### MsgUpdateGroupAdminResponse
MsgUpdateGroupAdminResponse is the Msg/UpdateGroupAdmin response type.






<a name="regen.group.v1alpha1.MsgUpdateGroupCommentRequest"></a>

### MsgUpdateGroupCommentRequest
MsgUpdateGroupCommentRequest is the Msg/UpdateGroupComment request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [bytes](#bytes) |  | admin is the account address of the group admin. |
| group | [uint64](#uint64) |  | group is the unique ID of the group. |
| comment | [string](#string) |  | comment is the updated group's comment. |






<a name="regen.group.v1alpha1.MsgUpdateGroupCommentResponse"></a>

### MsgUpdateGroupCommentResponse
MsgUpdateGroupCommentResponse is the Msg/UpdateGroupComment response type.






<a name="regen.group.v1alpha1.MsgUpdateGroupMembersRequest"></a>

### MsgUpdateGroupMembersRequest
MsgUpdateGroupMembersRequest is the Msg/UpdateGroupMembers request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| admin | [bytes](#bytes) |  | admin is the account address of the group admin. |
| group | [uint64](#uint64) |  | group is the unique ID of the group. |
| member_updates | [Member](#regen.group.v1alpha1.Member) | repeated | member_updates is the list of members to update, set power to 0 to remove a member. |






<a name="regen.group.v1alpha1.MsgUpdateGroupMembersResponse"></a>

### MsgUpdateGroupMembersResponse
MsgUpdateGroupMembersResponse is the Msg/UpdateGroupMembers response type.






<a name="regen.group.v1alpha1.MsgVoteRequest"></a>

### MsgVoteRequest
MsgVoteRequest is the Msg/Vote request type.


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| proposal | [uint64](#uint64) |  | proposal is the unique ID of the proposal. |
| voters | [bytes](#bytes) | repeated | voters is the lists of voters' account addresses. |
| choice | [Choice](#regen.group.v1alpha1.Choice) |  | choice is the voters' choice on the proposal. |
| comment | [string](#string) |  | comment's is the vote's comment. |






<a name="regen.group.v1alpha1.MsgVoteResponse"></a>

### MsgVoteResponse
MsgVoteResponse is the Msg/Vote response type.





 <!-- end messages -->

 <!-- end enums -->

 <!-- end HasExtensions -->


<a name="regen.group.v1alpha1.Msg"></a>

### Msg
Msg is the regen.group.v1alpha1 Msg service.

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateGroup | [MsgCreateGroupRequest](#regen.group.v1alpha1.MsgCreateGroupRequest) | [MsgCreateGroupResponse](#regen.group.v1alpha1.MsgCreateGroupResponse) | CreateGroup creates a new group with an admin account address, a list of members and an comment. |
| UpdateGroupMembers | [MsgUpdateGroupMembersRequest](#regen.group.v1alpha1.MsgUpdateGroupMembersRequest) | [MsgUpdateGroupMembersResponse](#regen.group.v1alpha1.MsgUpdateGroupMembersResponse) | UpdateGroupMembers updates the group members with given group id and admin address. |
| UpdateGroupAdmin | [MsgUpdateGroupAdminRequest](#regen.group.v1alpha1.MsgUpdateGroupAdminRequest) | [MsgUpdateGroupAdminResponse](#regen.group.v1alpha1.MsgUpdateGroupAdminResponse) | UpdateGroupAdmin updates the group admin with given group id and previous admin address. |
| UpdateGroupComment | [MsgUpdateGroupCommentRequest](#regen.group.v1alpha1.MsgUpdateGroupCommentRequest) | [MsgUpdateGroupCommentResponse](#regen.group.v1alpha1.MsgUpdateGroupCommentResponse) | UpdateGroupComment updates the group comment with given group id and admin address. |
| CreateGroupAccount | [MsgCreateGroupAccountRequest](#regen.group.v1alpha1.MsgCreateGroupAccountRequest) | [MsgCreateGroupAccountResponse](#regen.group.v1alpha1.MsgCreateGroupAccountResponse) | CreateGroupAccount creates a new group account using given DecisionPolicy. |
| UpdateGroupAccountAdmin | [MsgUpdateGroupAccountAdminRequest](#regen.group.v1alpha1.MsgUpdateGroupAccountAdminRequest) | [MsgUpdateGroupAccountAdminResponse](#regen.group.v1alpha1.MsgUpdateGroupAccountAdminResponse) | UpdateGroupAccountAdmin updates a group account admin. |
| UpdateGroupAccountDecisionPolicy | [MsgUpdateGroupAccountDecisionPolicyRequest](#regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyRequest) | [MsgUpdateGroupAccountDecisionPolicyResponse](#regen.group.v1alpha1.MsgUpdateGroupAccountDecisionPolicyResponse) | UpdateGroupAccountDecisionPolicy allows a group account decision policy to be updated. |
| UpdateGroupAccountComment | [MsgUpdateGroupAccountCommentRequest](#regen.group.v1alpha1.MsgUpdateGroupAccountCommentRequest) | [MsgUpdateGroupAccountCommentResponse](#regen.group.v1alpha1.MsgUpdateGroupAccountCommentResponse) | UpdateGroupAccountComment updates a group account comment. |
| CreateProposal | [MsgCreateProposalRequest](#regen.group.v1alpha1.MsgCreateProposalRequest) | [MsgCreateProposalResponse](#regen.group.v1alpha1.MsgCreateProposalResponse) | CreateProposal submits a new proposal. |
| Vote | [MsgVoteRequest](#regen.group.v1alpha1.MsgVoteRequest) | [MsgVoteResponse](#regen.group.v1alpha1.MsgVoteResponse) | Vote allows voters to vote on a proposal. |
| Exec | [MsgExecRequest](#regen.group.v1alpha1.MsgExecRequest) | [MsgExecResponse](#regen.group.v1alpha1.MsgExecResponse) | Exec executes a proposal. |

 <!-- end services -->



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

