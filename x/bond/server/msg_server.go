package server

import (
	"context"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/google/uuid"
	"github.com/regen-network/regen-ledger/orm"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/types/math"
	"github.com/regen-network/regen-ledger/v2/x/bond"
	"time"
)

// TODO: Revisit this once we have proper gas fee framework.
// Tracking issues https://github.com/cosmos/cosmos-sdk/issues/9054, https://github.com/cosmos/cosmos-sdk/discussions/9072
const gasIssueBondFee = uint64(10)
const gasSellBondFee = uint64(10)

// IssueBond creates a new bond
//
// The admin is charged a fee for creating the bond. This is controlled by
// the global parameter IssueBondFee.
func (s serverImpl) IssueBond(goCtx context.Context, req *bond.MsgIssueBond) (*bond.MsgIssueBondResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)
	bondId := s.genBondID()

	ctx.GasMeter().ConsumeGas(gasIssueBondFee, "bond issuance")

	err := s.bondInfoTable.Create(ctx, &bond.BondInfo{
		Id:              bondId,
		EmissionDenom:   req.EmissionDenom,
		Name:            req.Name,
		Holder:          req.Holder,
		ParentBond:      "",
		FaceValue:       req.FaceValue,
		FaceCurrency:    req.FaceCurrency,
		CreationDate:    time.Now().Format("2006-01-02T15:04:05Z07:00"),
		IssuanceDate:    req.IssuanceDate.Format("2006-01-02"),
		MaturityDate:    req.MaturityDate.Format("2006-01-02"),
		CouponRate:      req.CouponRate,
		CouponFrequency: req.CouponFrequency,
		Status:          bond.BondInfo_ACTIVE,
		Project:         req.Project,
		Metadata:        req.Metadata,
	})

	if err != nil {
		return nil, err
	}

	return &bond.MsgIssueBondResponse{BondId: bondId}, nil
}

// SellBond sells a bond
//
// The admin is charged a fee for selling the bond. This is controlled by
// the global parameter SellBondFee.
func (s serverImpl) SellBond(goCtx context.Context, req *bond.MsgSellBond) (*bond.MsgIssueBondResponse, error) {
	ctx := types.UnwrapSDKContext(goCtx)

	var bondToSell bond.BondInfo
	err := s.bondInfoTable.GetOne(ctx, orm.RowID(req.BondId), &bondToSell)
	if err != nil || bondToSell.Status != bond.BondInfo_ACTIVE {
		return nil, sdkerrors.ErrInvalidRequest.Wrapf("%s is not a valid bond to sell", req.BondId)
	}

	bondToSell.Status = bond.BondInfo_CLOSED
	err = s.bondInfoTable.Update(ctx, &bondToSell)
	if err != nil {
		return nil, err
	}

	amountToSell, err := math.NewPositiveFixedDecFromString(req.Amount, bond.PRECISION)
	if err != nil {
		return nil, err
	}

	bondToSellFaceValue, err := math.NewPositiveFixedDecFromString(bondToSell.FaceValue, bond.PRECISION)
	if amountToSell.Cmp(bondToSellFaceValue) == 1 {
		return nil, bond.ErrInsufficientFunds
	}

	ctx.GasMeter().ConsumeGas(gasSellBondFee, "bond sell")

	// Create the sold bond with a new holder
	soldBond := bond.BondInfo{
		Id:              s.genBondID(),
		EmissionDenom:   bondToSell.EmissionDenom,
		Name:            bondToSell.Name,
		Holder:          req.Buyer,
		ParentBond:      bondToSell.Id,
		FaceValue:       amountToSell.String(),
		FaceCurrency:    bondToSell.FaceCurrency,
		CreationDate:    time.Now().Format("2006-01-02T15:04:05Z07:00"),
		IssuanceDate:    bondToSell.IssuanceDate,
		MaturityDate:    bondToSell.MaturityDate,
		CouponRate:      bondToSell.CouponRate,
		CouponFrequency: bondToSell.CouponFrequency,
		Status:          bond.BondInfo_ACTIVE,
		Project:         bondToSell.Project,
		Metadata:        bondToSell.Metadata,
	}

	err = s.bondInfoTable.Create(ctx, &soldBond)
	if err != nil {
		return nil, err
	}

	if amountToSell.Cmp(bondToSellFaceValue) == -1 {
		newFaceValue, err := bondToSellFaceValue.Sub(amountToSell)

		notSoldBond := bond.BondInfo{
			Id:              s.genBondID(),
			EmissionDenom:   bondToSell.EmissionDenom,
			Name:            bondToSell.Name,
			Holder:          bondToSell.Holder,
			ParentBond:      bondToSell.Id,
			FaceValue:       newFaceValue.String(),
			FaceCurrency:    bondToSell.FaceCurrency,
			CreationDate:    time.Now().Format("2006-01-02T15:04:05Z07:00"),
			IssuanceDate:    bondToSell.IssuanceDate,
			MaturityDate:    bondToSell.MaturityDate,
			CouponRate:      bondToSell.CouponRate,
			CouponFrequency: bondToSell.CouponFrequency,
			Status:          bond.BondInfo_ACTIVE,
			Project:         bondToSell.Project,
			Metadata:        bondToSell.Metadata,
		}

		err = s.bondInfoTable.Create(ctx, &notSoldBond)
		if err != nil {
			return nil, err
		}
	}

	return &bond.MsgIssueBondResponse{BondId: soldBond.Id}, nil
}

func (s serverImpl) genBondID() string {
	return uuid.NewString()
}
