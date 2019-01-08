package proposal

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"golang.org/x/crypto/blake2b"
)

type Keeper struct {
	storeKey sdk.StoreKey
	handler  ProposalHandler
	cdc      *codec.Codec
}

func NewKeeper(storeKey sdk.StoreKey, handler ProposalHandler, cdc *codec.Codec) Keeper {
	return Keeper{storeKey: storeKey, handler: handler, cdc: cdc}
}

func (keeper Keeper) Propose(ctx sdk.Context, proposer sdk.AccAddress, action ProposalAction) sdk.Result {
	if !keeper.handler.CanHandle(action) {
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
			Log:  "can't handle this proposal type",
		}
	}

	store := ctx.KVStore(keeper.storeKey)
	hashBz := blake2b.Sum256(action.GetSignBytes())
	id := hashBz[:]
	if store.Has(id) {
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
			Log:  fmt.Sprintf("proposal %s already exists", hex.EncodeToString(id)),
		}
	}

	prop := Proposal{
		Proposer:  proposer,
		Action:    action,
		Approvers: []sdk.AccAddress{proposer},
	}

	keeper.storeProposal(ctx, id, &prop)

	return sdk.Result{Code: sdk.CodeOK, Data: id}
}

func (keeper Keeper) storeProposal(ctx sdk.Context, id []byte, proposal *Proposal) {
	store := ctx.KVStore(keeper.storeKey)
	bz, err := keeper.cdc.MarshalBinaryBare(proposal)
	if err != nil {
		panic(err)
	}

	store.Set(id, bz)
}

func (keeper Keeper) GetProposal(ctx sdk.Context, id []byte) (proposal *Proposal, err sdk.Error) {
	store := ctx.KVStore(keeper.storeKey)
	bz := store.Get(id)
	proposal = &Proposal{}
	marshalErr := keeper.cdc.UnmarshalBinaryBare(bz, proposal)
	if marshalErr != nil {
		return proposal, sdk.ErrUnknownRequest(marshalErr.Error())
	}
	return proposal, nil
}

func (keeper Keeper) Vote(ctx sdk.Context, proposalId []byte, voter sdk.AccAddress, yesNo bool) sdk.Result {
	proposal, err := keeper.GetProposal(ctx, proposalId)

	if err != nil {
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
			Log:  "can't find proposal",
		}
	}

	var newVotes []sdk.AccAddress
	votes := proposal.Approvers
	nVotes := len(votes)

	if yesNo {
		newVotes = make([]sdk.AccAddress, nVotes+1)
		for i := 0; i < nVotes; i++ {
			oldVoter := votes[i]
			if bytes.Equal(voter, oldVoter) {
				// Already voted YES
				return sdk.Result{
					Code: sdk.CodeUnknownRequest,
					Log:  "already voted yes",
				}
			}
			newVotes[i] = oldVoter
		}
		newVotes[nVotes] = voter
	} else {
		newVotes = make([]sdk.AccAddress, nVotes)
		didntVote := true
		j := 0
		for i := 0; i < nVotes; i++ {
			oldVoter := votes[i]
			if bytes.Equal(voter, oldVoter) {
				didntVote = false
			} else {
				newVotes[j] = oldVoter
				j++
			}
		}
		if didntVote {
			return sdk.Result{
				Code: sdk.CodeUnknownRequest,
				Log:  "didn't vote yes previously",
			}
		}
		if j != nVotes-1 {
			panic("unexpected vote count")
		}
		newVotes = newVotes[:j]
	}

	newProp := Proposal{
		Proposer:  proposal.Proposer,
		Action:    proposal.Action,
		Approvers: newVotes,
	}

	keeper.storeProposal(ctx, proposalId, &newProp)

	return sdk.Result{Code: sdk.CodeOK}
}

func (keeper Keeper) TryExecute(ctx sdk.Context, proposalId []byte) sdk.Result {
	proposal, err := keeper.GetProposal(ctx, proposalId)

	if err != nil {
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
			Log:  "can't find proposal",
		}
	}

	res := keeper.handler.Handle(ctx, proposal.Action, proposal.Approvers)

	if res.Code == sdk.CodeOK {
		store := ctx.KVStore(keeper.storeKey)
		store.Delete(proposalId)
	}

	return res
}

func (keeper Keeper) Withdraw(ctx sdk.Context, proposalId []byte, proposer sdk.AccAddress) sdk.Result {
	proposal, err := keeper.GetProposal(ctx, proposalId)

	if err != nil {
		return sdk.Result{
			Code: sdk.CodeUnknownRequest,
			Log:  "can't find proposal",
		}
	}

	if !bytes.Equal(proposer, proposal.Proposer) {
		return sdk.Result{
			Code: sdk.CodeUnauthorized,
			Log:  "you didn't propose this",
		}
	}

	store := ctx.KVStore(keeper.storeKey)
	store.Delete(proposalId)

	return sdk.Result{Code: sdk.CodeOK}
}
