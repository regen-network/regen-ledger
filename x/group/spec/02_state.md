# State

The `group` module uses the `orm` package which provides table storage with support for
primary keys and secondary indexes. `orm` also defines `Sequence` which is a persistent unique key generator based on a counter that can be used along with `Table`s.

Here's the list of tables and associated sequences and indexes stored as part of the `group` module.

## Group Table

The `groupTable` stores `GroupInfo`: `0x0 | []byte(GroupId) -> ProtocolBuffer(GroupInfo)`.

### groupSeq

The value of `groupSeq` is incremented when creating a new group and corresponds to the new `GroupId`: `0x1 | 0x1 -> BigEndian`.

The second `0x1` corresponds to the ORM `sequenceStorageKey`.

### groupByAdminIndex

`groupByAdminIndex` allows to retrieve groups by admin address:
`0x2 | []byte(group.Admin) | []byte(GroupId) -> []byte()`.

## Group Member Table

The `groupMemberTable` stores `GroupMember`s: `0x10 | []byte(GroupId) | []byte(member.Address) -> ProtocolBuffer(GroupMember)`.

The `groupMemberTable` is a primary key table and its `PrimaryKey` is given by
`[]byte(GroupId) | []byte(member.Address)` which is used by the following indexes.

### groupMemberByGroupIndex

`groupMemberByGroupIndex` allows to retrieve group members by group id:
`0x11 | []byte(GroupId) | PrimaryKey | byte(len(PrimaryKey)) -> []byte()`.

### groupMemberByMemberIndex

`groupMemberByMemberIndex` allows to retrieve group members by member address:
`0x12 | []byte(member.Address) | PrimaryKey | byte(len(PrimaryKey)) -> []byte()`.

## Group Account Table

The `groupAccountTable` stores `GroupAccountInfo`: `0x20 | []byte(Address) -> ProtocolBuffer(GroupAccountInfo)`.

The `groupAccountTable` is a primary key table and its `PrimaryKey` is given by
`[]byte(Address)` which is used by the following indexes.

### groupAccountSeq

The value of `groupAccountSeq` is incremented when creating a new group account and is used to generate the new group account `Address`:
`0x21 | 0x1 -> BigEndian`.

The second `0x1` corresponds to the ORM `sequenceStorageKey`.

### groupAccountByGroupIndex

`groupAccountByGroupIndex` allows to retrieve group accounts by group id:
`0x22 | []byte(GroupId) | PrimaryKey | byte(len(PrimaryKey)) -> []byte()`.

### groupAccountByAdminIndex

`groupAccountByAdminIndex` allows to retrieve group accounts by admin address:
`0x23 | []byte(Address) | PrimaryKey | byte(len(PrimaryKey)) -> []byte()`.

## Proposal Table

The `proposalTable` stores `Proposal`s: `0x30 | []byte(ProposalId) -> ProtocolBuffer(Proposal)`.

### proposalSeq

The value of `proposalSeq` is incremented when creating a new proposal and corresponds to the new `ProposalId`: `0x31 | 0x1 -> BigEndian`.

The second `0x1` corresponds to the ORM `sequenceStorageKey`.

### proposalByGroupAccountIndex

`proposalByGroupAccountIndex` allows to retrieve proposals by group account address:
`0x32 | []byte(account.Address) | []byte(ProposalId) -> []byte()`.

### proposalByProposerIndex

`proposalByProposerIndex` allows to retrieve proposals by proposer address:
`0x33 | []byte(proposer.Address) | []byte(ProposalId) -> []byte()`.

## Vote Table

The `voteTable` stores `Vote`s: `0x40 | []byte(ProposalId) | []byte(voter.Address) -> ProtocolBuffer(Vote)`.

The `voteTable` is a primary key table and its `PrimaryKey` is given by
`[]byte(ProposalId) | []byte(voter.Address)` which is used by the following indexes.

### voteByProposalIndex

`voteByProposalIndex` allows to retrieve votes by proposal id:
`0x41 | []byte(ProposalId) | PrimaryKey | byte(len(PrimaryKey)) -> []byte()`.

### voteByVoterIndex

`voteByVoterIndex` allows to retrieve votes by voter address:
`0x42 | []byte(voter.Address) | PrimaryKey | byte(len(PrimaryKey)) -> []byte()`.

