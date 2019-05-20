package oracle

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
	"time"
)

type MsgRequestCompute struct {
	// Invocation is the descriptor of the function to invoke. Invocation specifies
	Invocation Invocation
	// Fee is the fee to pay the oracle(s) for running the function
	Fee    sdk.Coins
	Params InvocationParams
	// Requester is the address requesting computation
	Requester sdk.AccAddress
}

type InvocationParams struct {
	MinOracles         int
	RandomAuditPercent int
	MinBond            sdk.Int
	ChallengeWindow    time.Duration
}

type Invocation struct {
	// Function is the specification of the function to run
	Function types.DataAddress
	// Input is the function input
	Input types.DataAddress
	// BlockHeight is the height up to which the function has access to the
	// block-chain state database
	BlockHeight int64
	// Authorization is the value to be passed as the HTTP Authorization header
	// when private data stores are queried. It should be left blank to allow
	// only public data. This value should ensure that functions invocation is
	// purely deterministic even up to access to private data. More detailed
	// specification TBD
	Authorization string
}

type MsgCommitResult struct {
	Hash      []byte
	FlagAudit bool
	Oracle    sdk.AccAddress
}

type Vote struct {
	Oracle sdk.AccAddress
	Agree  bool
}

type MsgVoteResult struct {
	Result types.DataAddress
	Votes  []Vote
	Oracle sdk.AccAddress
}
