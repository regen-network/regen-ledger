package server

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/group"
)

// NewHandler creates an sdk.Handler for all the group type messages
func NewHandler(storeKey sdk.StoreKey, paramSpace paramstypes.Subspace,
	router sdk.Router, cdc codec.Marshaler) sdk.Handler {
	impl := newServer(storeKey, paramSpace, router, cdc)

	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		regenCtx := types.Context{Context: ctx}

		switch msg := msg.(type) {
		case *group.MsgCreateGroupRequest:
			res, err := impl.CreateGroup(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgUpdateGroupMembersRequest:
			res, err := impl.UpdateGroupMembers(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgUpdateGroupAdminRequest:
			res, err := impl.UpdateGroupAdmin(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgUpdateGroupCommentRequest:
			res, err := impl.UpdateGroupComment(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgCreateGroupAccountRequest:
			res, err := impl.CreateGroupAccount(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgUpdateGroupAccountAdminRequest:
			res, err := impl.UpdateGroupAccountAdmin(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgUpdateGroupAccountDecisionPolicyRequest:
			res, err := impl.UpdateGroupAccountDecisionPolicy(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgUpdateGroupAccountCommentRequest:
			res, err := impl.UpdateGroupAccountComment(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgCreateProposalRequest:
			res, err := impl.CreateProposal(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgVoteRequest:
			res, err := impl.Vote(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgExecRequest:
			res, err := impl.Exec(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, errors.Wrapf(errors.ErrUnknownRequest, "unrecognized %s message type: %T", group.ModuleName, msg)
		}
	}
}
