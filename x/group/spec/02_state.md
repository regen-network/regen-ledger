<!--
order: 2
-->

# State

The `group` module uses the `orm` package to 

## Group Table

`0x0 | ->`

### groupSeq

`0x1 -> ProtocolBuffer(uint64)`

### groupByAdminIndex

`0x2 | ->`

## Group Member Table

`0x10 | ->`

### groupMemberByGroupIndex

`0x11 | ->`

### groupMemberByMemberIndex

`0x12 | ->`


## Group Account Table

`0x20 | ->`

### groupAccountSeq

`0x21 ->`

### groupAccountByGroupIndex

`0x22 | ->`

### groupAccountByAdminIndex

`0x22 | ->`


## Proposal Table

`0x30 | ->`

### proposalSeq

`0x31 ->`

### proposalByGroupAccountIndex

`0x32 | ->`

### proposalByProposerIndex

`0x33 | ->`


## Vote Table

`0x40 | ->`

### voteByProposalIndex

`0x41 | ->`

### voteByVoterIndex

`0x42 | ->`

