package core

import (
	"context"
	"encoding/hex"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	bridgev1 "github.com/regen-network/regen-ledger/api/axelar/bridge/v1"
	"github.com/regen-network/regen-ledger/x/axelarbridge"
)

func NewRegenBridgeHandler(cdc codec.Codec, router *baseapp.MsgServiceRouter) axelarbridge.HandlerMap {
	return axelarbridge.HandlerMap{
		"regen_toucan_bridge": func(ctx context.Context, event bridgev1.Event) error {
			if event.Sender != "0x<harcoded address of regen smart contract on Polygon>" {
				return sdkerrors.ErrInvalidRequest.Wrap("unknown sender address")
			}

			var payload ToucanPayload
			err := cdc.Unmarshal(event.Payload, &payload)
			if err != nil {
				return err
			}

			batchDenom, err := convertERC20ToBatchDenom(payload.Tco2ContractAddress)
			if err != nil {
				return err
			}

			msg := MsgMintBatchCredits{
				Issuer:     "", // TODO
				BatchDenom: batchDenom,
				Issuance: []*BatchIssuance{{
					Recipient: payload.Recipient,
				}},
				OriginTx: &OriginTx{Id: hex.EncodeToString(event.SrcTxId), Typ: event.SrcChain},
				Note:     payload.Note,
			}

			h := router.Handler(&msg)
			if h == nil {
				return sdkerrors.ErrInvalidRequest.Wrap("empty handler")
			}

			_, err = h(sdk.UnwrapSDKContext(ctx), &msg)
			return err
		},
	}
}

func convertERC20ToBatchDenom(tc02Token string) (string, error) {
	switch tc02Token {
	case "0x<toucan>":
		return "C01-20190101-20210101-008", nil
		// TODO Other cases to be updates by governance.
	}

	return "", sdkerrors.ErrUnknownRequest.Wrapf("unknown tco2 token %s", tc02Token)
}
