package group

import (
	"errors"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/orm"
)

func NewHandler(k Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		logger := ctx.Logger().With("module", fmt.Sprintf("x/%s", ModuleName))
		switch msg := msg.(type) {
		case *MsgPropose:
			return handleMsgPropose(ctx, k, msg)
		case *MsgAlwaysSucceed:
			logger.Info("executed MsgAlwaysSucceed msg")
			return &sdk.Result{
				Data:   nil,
				Log:    "MsgAlwaysSucceed executed",
				Events: ctx.EventManager().ABCIEvents(),
			}, nil
		case *MsgAlwaysFail:
			logger.Info("executed MsgAlwaysFail msg")
			return nil, errors.New("execution of MsgAlwaysFail testdata always fails")
		case *MsgSetValue:
			logger.Info("executed MsgSetValue msg")
			k.SetValue(ctx, msg.Value)
			return &sdk.Result{
				Data:   []byte(msg.Value),
				Log:    "MsgSetValue executed",
				Events: ctx.EventManager().ABCIEvents(),
			}, nil
		case *MsgIncCounter:
			logger.Info("executed MsgIncCounter msg")
			return &sdk.Result{
				Data:   k.IncCounter(ctx),
				Log:    "MsgIncCounter executed",
				Events: ctx.EventManager().ABCIEvents(),
			}, nil
		case *MsgConditional:
			logger.Info("executed MsgConditional msg")
			if k.GetCounter(ctx) != msg.ExpectedCounter {
				return nil, errors.New("counter condition not matched")
			}
			return &sdk.Result{
				Data:   orm.EncodeSequence(msg.ExpectedCounter),
				Log:    "MsgConditional executed",
				Events: ctx.EventManager().ABCIEvents(),
			}, nil
		case *MsgAuthenticate:
			logger.Info("executed MsgAuthenticate msg")
			return &sdk.Result{
				Data:   nil,
				Log:    "MsgAuthenticate executed",
				Events: ctx.EventManager().ABCIEvents(),
			}, nil
		default:
			return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized message type: %T", msg)
		}
	}
}

func handleMsgPropose(ctx sdk.Context, k Keeper, msg *MsgPropose) (*sdk.Result, error) {
	id, err := k.CreateProposal(ctx, msg.Base.GroupAccount, msg.Base.Proposers, msg.Base.Comment, msg.GetMsgs())
	if err != nil {
		return nil, err
	}
	return &sdk.Result{
		Data:   id.Bytes(),
		Log:    fmt.Sprintf("Proposal created :%d", id),
		Events: ctx.EventManager().ABCIEvents(),
	}, nil
}
