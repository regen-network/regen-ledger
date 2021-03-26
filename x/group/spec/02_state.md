<!--
order: 2
-->

# State

The `group` module uses the `orm` package which provides table storage with support for
primary keys and secondary indexes. `orm` also defines `Sequence` which is a persistent unique key generator based on a counter that can be used along with `Table`s.

Here's the list of tables and associated sequences and indexes stored as part of the `group` module.

## Group Table

The `groupTable` stores `GroupInfo`: `0x0 | []byte(group_id) -> ProtocolBuffer(GroupInfo)`

### groupSeq

The value of `groupSeq` is incremented when creating a new group and corresponds to the new `GroupId`: `0x1 | 0x1 -> BigEndian`

### groupByAdminIndex

`groupByAdminIndex` allows to retrieve groups by admin address using prefix `0x2`.

## Group Member Table

The `groupMemberTable` stores `GroupMember`s: `0x10 | []byte(GroupId) | []byte(member.Address) -> ProtocolBuffer(GroupMember)`

### groupMemberByGroupIndex

`groupMemberByGroupIndex` allows to retrieve group members by group id using prefix `0x11`.

### groupMemberByMemberIndex

`groupMemberByMemberIndex` allows to retrieve group members by member address using prefix `0x12`.

## Group Account Table

The `groupAccountTable` stores `GroupAccountInfo`: `0x20 | []byte(Address) -> ProtocolBuffer(GroupAccountInfo)`

### groupAccountSeq

The value of `groupAccountSeq` is incremented when creating a new group account and is used to generate the new group account `Address`:
`0x21 | 0x1 -> BigEndian`

### groupAccountByGroupIndex

`groupAccountByGroupIndex` allows to retrieve group accounts by group id using prefix `0x22`.
`0x22 | ->`

### groupAccountByAdminIndex

`groupAccountByAdminIndex` allows to retrieve group accounts by admin address using prefix `0x23`.

## Proposal Table

The `proposalTable` stores `Proposal`s: `0x30 | []byte(ProposalId) -> ProtocolBuffer(Proposal)`

### proposalSeq

The value of `proposalSeq` is incremented when creating a new proposal and corresponds to the new `ProposalId`: `0x31 | 0x1 -> BigEndian`

### proposalByGroupAccountIndex

`proposalByGroupAccountIndex` allows to retrieve proposals by group account address using prefix `0x32`.

### proposalByProposerIndex

`proposalByProposerIndex` allows to retrieve proposals by proposer address using prefix `0x33`.

## Vote Table

The `voteTable` stores `Vote`s: `0x40 | []byte(ProposalId) | []byte(voter.Address) -> ProtocolBuffer(Vote)`

### voteByProposalIndex

`voteByProposalIndex` allows to retrieve votes by proposal id using prefix `0x41`.

### voteByVoterIndex

`voteByVoterIndex` allows to retrieve votes by voter address using prefix `0x42`.
