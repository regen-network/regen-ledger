package server

import (
	"strconv"

	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/regen-network/regen-ledger/x/group"
)

type resTally struct {
	YesCount     int64
	NoCount      int64
	AbstainCount int64
	VetoCount    int64
}

func resultVoteState(v group.Tally) resTally {
	return resTally{
		YesCount:     strToInt(v.YesCount),
		NoCount:      strToInt(v.NoCount),
		AbstainCount: strToInt(v.AbstainCount),
		VetoCount:    strToInt(v.VetoCount),
	}
}

func strToInt(str string) int64 {
	integer, _ := strconv.ParseInt(str, 10, 64)
	return integer
}

func (s serverImpl) RegisterInvariants(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(group.ModuleName, "Tally-Votes", s.TallyVotesInvariant())
}

func (s serverImpl) AllInvariants() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := s.TallyVotesInvariant()(ctx)
		if stop {
			return res, stop
		}
		return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", "\tTallyVoteSums is failed"), false
	}
}

func (s serverImpl) TallyVotesInvariant() sdk.Invariant {
	return func(sdkCtx sdk.Context) (string, bool) {
		ctx := types.Context{Context: sdkCtx}
		var pageReq = query.PageRequest{
			Offset:     0,
			Limit:      100,
			CountTotal: true,
		}
		it, err := s.getAllProposals(ctx, &pageReq)
		if err != nil {
			return "No proposals found", false
		}

		var Proposals2 []*group.Proposal
		_, err = orm.Paginate(it, &pageReq, &Proposals2)
		if err != nil {
			return "", false
		}
		sdkCtx = sdkCtx.WithBlockHeight(ctx.BlockHeight() - 1)
		it2, err := s.getAllProposals(types.Context{Context: sdkCtx}, &pageReq)
		if err != nil {
			return "No proposals found", false
		}

		var Proposals []*group.Proposal
		_, err = orm.Paginate(it2, &pageReq, &Proposals)
		if err != nil {
			return "", false
		}
		for i := 0; i < len(Proposals) && i < len(Proposals2); i++ {
			if int32(Proposals[i].Status) == 1 && int32(Proposals2[i].Status) == 1 {
				var voteState1 = resultVoteState(Proposals[i].VoteState)
				var voteState2 = resultVoteState(Proposals2[i].VoteState)
				if (voteState1.YesCount > voteState2.YesCount) || (voteState1.NoCount > voteState2.NoCount) || (voteState1.AbstainCount > voteState2.AbstainCount) || (voteState1.VetoCount > voteState2.VetoCount) {
					return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", "\tTallyVoteSums is failed"), true
				}
			}
		}
		return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", "\tTallyVoteSums is passed"), false

	}
}
