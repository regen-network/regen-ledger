package server

import (
	"fmt"
	"math"

	"github.com/regen-network/regen-ledger/orm"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/group"
)

func (s serverImpl) RegisterInvariants(ir sdk.InvariantRegistry) {
	ir.RegisterRoute(group.ModuleName, "Tally-Votes", s.TallyVotesInvariant())
}

func (s serverImpl) AllInvariants() sdk.Invariant {
	return func(ctx sdk.Context) (string, bool) {
		res, stop := s.TallyVotesInvariant()(ctx)
		if stop {
			return res, stop
		}
		return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", "\tvote tally sums must never have less than the block before"), false
	}
}

func (s serverImpl) TallyVotesInvariant() sdk.Invariant {
	return func(sdkCtx sdk.Context) (string, bool) {
		var proposals2 []*group.Proposal
		ctx := types.Context{Context: sdkCtx}
		it2, err := s.proposalTable.PrefixScan(ctx, 1, math.MaxUint64)
		if err != nil {
			return "start value must be less than end value in iterator", false
		}
		_, err = orm.ReadAll(it2, &proposals2)
		var proposals1 []*group.Proposal
		if ctx.BlockHeight()-1 >= 0 {
			sdkCtx = sdkCtx.WithBlockHeight(ctx.BlockHeight() - 1)
		} else {
			return "Not enough blocks to perform TallyVotesInvariant", false
		}
		it1, err := s.proposalTable.PrefixScan(sdkCtx, 1, math.MaxUint64)
		if err != nil {
			return "start value must be less than end value in iterator", false
		}
		_, err = orm.ReadAll(it1, &proposals1)

		for i := 0; i < len(proposals1) && i < len(proposals2); i++ {
			if int32(proposals1[i].Status) == 1 && int32(proposals2[i].Status) == 1 {
				yesCount1, err := proposals1[i].VoteState.GetYesCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				yesCount2, err := proposals2[i].VoteState.GetYesCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				noCount1, err := proposals1[i].VoteState.GetNoCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				noCount2, err := proposals2[i].VoteState.GetNoCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				abstainCount1, err := proposals1[i].VoteState.GetAbstainCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				abstainCount2, err := proposals2[i].VoteState.GetAbstainCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				vetoCount1, err := proposals1[i].VoteState.GetVetoCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				vetoCount2, err := proposals2[i].VoteState.GetVetoCount()
				if err != nil {
					return fmt.Sprint(err), false
				}
				if (yesCount2.Cmp(yesCount1) == -1) || (noCount2.Cmp(noCount1) == -1) || (abstainCount2.Cmp(abstainCount1) == -1) || (vetoCount2.Cmp(vetoCount1) == -1) {
					return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", "\tvote tally sums must never have less than the block before"), false
				}
			}
		}
		return sdk.FormatInvariant(group.ModuleName, "Tally-Votes", "\tTallyVotesSum is passed"), false
	}
}
