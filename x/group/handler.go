package group

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/pkg/errors"
)

// NewHandler creates a new message handler.
func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case MsgCreateGroup:
			return handleMsgCreateGroup(ctx, k, msg)
		case MsgUpdateGroupAdmin:
			return handleMsgUpdateGroupAdmin(ctx, k, msg)
		case MsgUpdateGroupComment:
			return handleMsgUpdateGroupComment(ctx, k, msg)
		case MsgUpdateGroupMembers:
			return handleMsgUpdateGroupMembers(ctx, k, msg)
		case MsgCreateGroupAccount:
			return handleMsgCreateGroupAccount(ctx, k, msg)
		case MsgVote:
			return handleMsgVote(ctx, k, msg)
		case MsgExec:
			return handleMsgExec(ctx, k, msg)
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized group message type: %T", msg)
		}
	}
}

// TODO: Do we want to introduce any new event types?

func handleMsgVote(ctx sdk.Context, k Keeper, msg MsgVote) (*sdk.Result, error) {
	if err := k.Vote(ctx, msg.Proposal, msg.Voters, msg.Choice, msg.Comment); err != nil {
		return nil, err
	}
	// todo: event?
	return &sdk.Result{
		Log:    fmt.Sprintf("Voted for proposal: %d", msg.Proposal),
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}

func handleMsgExec(ctx sdk.Context, k Keeper, msg MsgExec) (*sdk.Result, error) {
	if err := k.ExecProposal(ctx, msg.Proposal); err != nil {
		return nil, err
	}
	// todo: event?
	return &sdk.Result{
		Log:    fmt.Sprintf("Executed proposal: %d", msg.Proposal),
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}

func handleMsgCreateGroupAccount(ctx sdk.Context, k Keeper, msg MsgCreateGroupAccount) (*sdk.Result, error) {
	decisionPolicy := msg.GetDecisionPolicy()
	acc, err := k.CreateGroupAccount(ctx, msg.GetAdmin(), msg.GetGroup(), decisionPolicy, msg.GetComment())
	if err != nil {
		return nil, errors.Wrap(err, "create group account")
	}
	return buildGroupAccountResult(ctx, msg.GetAdmin(), acc, "created")
}

func buildGroupAccountResult(ctx sdk.Context, admin sdk.AccAddress, acc sdk.AccAddress, note string) (*sdk.Result, error) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			sdk.EventTypeMessage,
			sdk.NewAttribute(sdk.AttributeKeyModule, ModuleName),
			sdk.NewAttribute(sdk.AttributeKeySender, admin.String()),
		),
	)
	return &sdk.Result{
		Data:   acc.Bytes(),
		Log:    fmt.Sprintf("Group account %s %s", acc.String(), note),
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}
