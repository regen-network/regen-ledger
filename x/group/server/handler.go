package server

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/types"
	servermodule "github.com/regen-network/regen-ledger/types/module/server"
	"github.com/regen-network/regen-ledger/x/group"
	"github.com/regen-network/regen-ledger/x/group/exported"
)

// NewHandler creates an sdk.Handler for all the group type messages.
// This is needed for supporting amino-json signing.
func NewHandler(configurator servermodule.Configurator, accountKeeper exported.AccountKeeper, bankKeeper exported.BankKeeper) sdk.Handler {
	impl := newServer(configurator.ModuleKey(), accountKeeper, bankKeeper, configurator.Marshaler())

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

		case *group.MsgUpdateGroupMetadataRequest:
			res, err := impl.UpdateGroupMetadata(regenCtx, msg)
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

		case *group.MsgUpdateGroupAccountMetadataRequest:
			res, err := impl.UpdateGroupAccountMetadata(regenCtx, msg)
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
