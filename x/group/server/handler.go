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
		case *group.MsgCreateGroup:
			res, err := impl.CreateGroup(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgUpdateGroupMembers:
			res, err := impl.UpdateGroupMembers(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgUpdateGroupAdmin:
			res, err := impl.UpdateGroupAdmin(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgUpdateGroupMetadata:
			res, err := impl.UpdateGroupMetadata(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgCreateGroupAccount:
			res, err := impl.CreateGroupAccount(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgUpdateGroupAccountAdmin:
			res, err := impl.UpdateGroupAccountAdmin(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgUpdateGroupAccountDecisionPolicy:
			res, err := impl.UpdateGroupAccountDecisionPolicy(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgUpdateGroupAccountMetadata:
			res, err := impl.UpdateGroupAccountMetadata(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgCreateProposal:
			res, err := impl.CreateProposal(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgVote:
			res, err := impl.Vote(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		case *group.MsgExec:
			res, err := impl.Exec(regenCtx, msg)
			return sdk.WrapServiceResult(ctx, res, err)

		default:
			return nil, errors.Wrapf(errors.ErrUnknownRequest, "unrecognized %s message type: %T", group.ModuleName, msg)
		}
	}
}
