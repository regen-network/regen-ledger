package server

import (
	"encoding/binary"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/regen-network/regen-ledger/util/storehelpers"
	"github.com/regen-network/regen-ledger/x/bank"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cockroachdb/apd/v2"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/regen-network/regen-ledger/orm"

	"github.com/regen-network/regen-ledger/x/bank/math"
	"github.com/regen-network/regen-ledger/x/ecocredit"
)

func (s serverImpl) CreateClass(ctx sdk.Context, req *ecocredit.MsgCreateClassRequest) (*ecocredit.MsgCreateClassResponse, error) {
	classID := s.idSeq.NextVal(ctx)
	classIDStr := uint64ToBase58Checked(classID)

	err := s.classInfoTable.Create(ctx, &ecocredit.ClassInfo{
		ClassId:  classIDStr,
		Designer: req.Designer,
		Issuers:  req.Issuers,
		Metadata: req.Metadata,
	})
	if err != nil {
		return nil, err
	}

	err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventCreateClass{
		ClassId:  classIDStr,
		Designer: req.Designer,
	})
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgCreateClassResponse{ClassId: classIDStr}, nil
}

func (s serverImpl) CreateBatch(ctx sdk.Context, req *ecocredit.MsgCreateBatchRequest) (*ecocredit.MsgCreateBatchResponse, error) {
	classID := req.ClassId
	if err := s.assertClassIssuer(ctx, classID, req.Issuer); err != nil {
		return nil, err
	}

	batchID := s.idSeq.NextVal(ctx)
	tradableSupply := apd.New(0, 0)
	retiredSupply := apd.New(0, 0)
	var maxDecimalPlaces uint32 = 0

	aclRule, err := types.NewAnyWithValue(&bank.ACLRule{AllowedAddresses: []string{s.moduleAddr.String()}})
	if err != nil {
		return nil, err
	}

	sendRule, err := types.NewAnyWithValue(&bank.BooleanRule{Enabled: true})
	if err != nil {
		return nil, err
	}

	res, err := s.bankMsgClient.CreateDenom(ctx, &bank.MsgCreateDenomRequest{
		NamespaceAdmin:   s.moduleAddr.String(),
		DenomNamespace:   "eco",
		DenomName:        fmt.Sprintf("%s/%s", classID, uint64ToBase58Checked(batchID)),
		DenomAdmin:       s.moduleAddr.String(),
		MintRule:         aclRule,
		SendRule:         sendRule,
		MoveRule:         aclRule,
		BurnRule:         aclRule,
		MaxDecimalPlaces: 0,
	})
	if err != nil {
		return nil, err
	}

	batchDenom := res.Denom

	store := ctx.KVStore(s.key)

	var mintIssuance []*bank.MsgMintRequest_Issuance

	for _, issuance := range req.Issuance {
		recipient := issuance.Recipient

		tradable, err := math.ParseNonNegativeDecimal(issuance.TradableUnits)
		if err != nil {
			return nil, err
		}

		if !tradable.IsZero() {
			err = math.Add(tradableSupply, tradableSupply, tradable)
			if err != nil {
				return nil, err
			}

			mintIssuance = append(mintIssuance, &bank.MsgMintRequest_Issuance{
				Recipient: recipient,
				Coins: []*bank.Coin{
					{
						Denom:  batchDenom,
						Amount: issuance.TradableUnits,
					},
				},
			})
		}
		retired, err := math.ParseNonNegativeDecimal(issuance.RetiredUnits)
		if err != nil {
			return nil, err
		}

		decPlaces := math.NumDecimalPlaces(retired)
		if decPlaces > maxDecimalPlaces {
			maxDecimalPlaces = decPlaces
		}

		if !retired.IsZero() {
			err = math.Add(retiredSupply, retiredSupply, retired)
			if err != nil {
				return nil, err
			}

			err = retire(ctx, store, recipient, batchDenomT(batchDenom), retired)
			if err != nil {
				return nil, err
			}
		}

		var sum apd.Decimal
		err = math.Add(&sum, tradable, retired)
		if err != nil {
			return nil, err
		}

		err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventReceive{
			Recipient:  recipient,
			BatchDenom: batchDenom,
			Units:      math.DecimalString(&sum),
		})
		if err != nil {
			return nil, err
		}
	}

	mintRes, err := s.bankMsgClient.Mint(ctx, &bank.MsgMintRequest{
		MinterAddress: s.moduleAddr.String(),
		Issuance:      mintIssuance,
	})
	if err != nil {
		return nil, err
	}

	if maxDecimalPlaces > mintRes.MaxDecimalPlaces {
		_, err := s.bankMsgClient.SetPrecision(ctx, &bank.MsgSetPrecisionRequest{
			DenomAdmin:       s.moduleAddr.String(),
			Denom:            batchDenom,
			MaxDecimalPlaces: maxDecimalPlaces,
		})
		if err != nil {
			return nil, err
		}
	}

	storehelpers.SetDecimal(store, RetiredSupplyKey(batchDenomT(batchDenom)), retiredSupply)

	var totalSupply apd.Decimal
	err = math.Add(&totalSupply, tradableSupply, retiredSupply)
	if err != nil {
		return nil, err
	}

	totalSupplyStr := math.DecimalString(&totalSupply)
	err = s.batchInfoTable.Create(ctx, &ecocredit.BatchInfo{
		ClassId:    classID,
		BatchDenom: batchDenom,
		Issuer:     req.Issuer,
		TotalUnits: totalSupplyStr,
		Metadata:   req.Metadata,
	})
	if err != nil {
		return nil, err
	}

	err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventCreateBatch{
		ClassId:    classID,
		BatchDenom: batchDenom,
		Issuer:     req.Issuer,
		TotalUnits: totalSupplyStr,
	})
	if err != nil {
		return nil, err
	}

	return &ecocredit.MsgCreateBatchResponse{BatchDenom: batchDenom}, nil
}

func (s serverImpl) Send(ctx sdk.Context, req *ecocredit.MsgSendRequest) (*ecocredit.MsgSendResponse, error) {
	store := ctx.KVStore(s.key)
	sender := req.Sender
	recipient := req.Recipient

	for _, credit := range req.Credits {
		denom := credit.BatchDenom

		tradable, err := math.ParseNonNegativeDecimal(credit.TradableUnits)
		if err != nil {
			return nil, err
		}

		retired, err := math.ParseNonNegativeDecimal(credit.RetiredUnits)
		if err != nil {
			return nil, err
		}

		var sum apd.Decimal
		err = math.Add(&sum, tradable, retired)
		if err != nil {
			return nil, err
		}

		_, err = s.bankMsgClient.Send(ctx, &bank.MsgSendRequest{
			FromAddress: sender,
			ToAddress:   recipient,
			Amount: []*bank.Coin{
				{
					Denom:  denom,
					Amount: credit.TradableUnits,
				},
			},
		})
		if err != nil {
			return nil, err
		}

		_, err = s.bankMsgClient.Burn(ctx, &bank.MsgBurnRequest{
			BurnerAddress: sender,
			Coins: []*bank.Coin{
				{
					Denom:  denom,
					Amount: credit.RetiredUnits,
				},
			},
		})

		// Add retired balance
		err = retire(ctx, store, recipient, batchDenomT(denom), retired)
		if err != nil {
			return nil, err
		}

		// Add retired supply
		err = storehelpers.GetAddAndSetDecimal(store, RetiredSupplyKey(batchDenomT(denom)), retired)
		if err != nil {
			return nil, err
		}

		err = ctx.EventManager().EmitTypedEvent(&ecocredit.EventReceive{
			Sender:     sender,
			Recipient:  recipient,
			BatchDenom: denom,
			Units:      math.DecimalString(&sum),
		})
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgSendResponse{}, nil
}

func (s serverImpl) Retire(ctx sdk.Context, req *ecocredit.MsgRetireRequest) (*ecocredit.MsgRetireResponse, error) {
	store := ctx.KVStore(s.key)
	holder := req.Holder
	for _, credit := range req.Credits {
		denom := credit.BatchDenom
		if !s.batchInfoTable.Has(ctx, orm.RowID(denom)) {
			return nil, sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, fmt.Sprintf("%s is not a valid credit denom", denom))
		}

		toRetire, err := math.ParsePositiveDecimal(credit.Units)
		if err != nil {
			return nil, err
		}

		_, err = s.bankMsgClient.Burn(ctx, &bank.MsgBurnRequest{
			BurnerAddress: holder,
			Coins: []*bank.Coin{
				{
					Denom:  denom,
					Amount: credit.Units,
				},
			},
		})
		if err != nil {
			return nil, err
		}

		//  Add retired balance
		err = retire(ctx, store, holder, batchDenomT(denom), toRetire)
		if err != nil {
			return nil, err
		}

		//  Add retired supply
		err = storehelpers.GetAddAndSetDecimal(store, RetiredSupplyKey(batchDenomT(denom)), toRetire)
		if err != nil {
			return nil, err
		}
	}

	return &ecocredit.MsgRetireResponse{}, nil
}

// assertClassIssuer makes sure that the issuer is part of issuers of given classID.
// Returns ErrUnauthorized otherwise.
func (s serverImpl) assertClassIssuer(ctx sdk.Context, classID, issuer string) error {
	classInfo, err := s.getClassInfo(ctx, classID)
	if err != nil {
		return err
	}
	for _, i := range classInfo.Issuers {
		if issuer == i {
			return nil
		}
	}
	return sdkerrors.ErrUnauthorized
}

func uint64ToBase58Checked(x uint64) string {
	buf := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(buf, x)
	return base58.CheckEncode(buf[:n], 0)
}

func retire(ctx sdk.Context, store sdk.KVStore, recipient string, batchDenom batchDenomT, retired *apd.Decimal) error {
	err := storehelpers.GetAddAndSetDecimal(store, RetiredBalanceKey(recipient, batchDenom), retired)
	if err != nil {
		return err
	}

	return ctx.EventManager().EmitTypedEvent(&ecocredit.EventRetire{
		Retirer:    recipient,
		BatchDenom: string(batchDenom),
		Units:      math.DecimalString(retired),
	})
}
