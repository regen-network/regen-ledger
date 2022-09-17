package utils

import (
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/ecocredit"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func RandomExponent(r *rand.Rand, precision uint32) uint32 {
	exponents := []uint32{0, 1, 2, 3, 6, 9, 12, 15, 18, 21, 24}
	for {
		x := exponents[r.Intn(len(exponents))]
		if x > precision {
			return x
		}
	}
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// GenAndDeliverTxWithRandFees generates a transaction with a random fee and delivers it.
func GenAndDeliverTxWithRandFees(r *rand.Rand, txCtx simulation.OperationInput) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	account := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)
	spendable := txCtx.Bankkeeper.SpendableCoins(txCtx.Context, account.GetAddress())

	var fees sdk.Coins
	var err error

	coins, hasNeg := spendable.SafeSub(txCtx.CoinsSpentInMsg...)
	if hasNeg {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "message doesn't leave room for fees"), nil, err
	}

	fees, err = simtypes.RandomFees(txCtx.R, txCtx.Context, coins)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to generate fees"), nil, err
	}
	return GenAndDeliverTx(r, txCtx, fees)
}

// GenAndDeliverTx generates a transactions and delivers it.
func GenAndDeliverTx(r *rand.Rand, txCtx simulation.OperationInput, fees sdk.Coins) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	account := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)
	tx, err := helpers.GenSignedMockTx(
		r,
		txCtx.TxGen,
		[]sdk.Msg{txCtx.Msg},
		fees,
		6000000,
		txCtx.Context.ChainID(),
		[]uint64{account.GetAccountNumber()},
		[]uint64{account.GetSequence()},
		txCtx.SimAccount.PrivKey,
	)

	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to generate mock tx"), nil, err
	}

	_, _, err = txCtx.App.SimDeliver(txCtx.TxGen.TxEncoder(), tx)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return simtypes.NoOpMsg(ecocredit.ModuleName, txCtx.MsgType, "not enough balance"), nil, nil
		}
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to deliver tx"), nil, err
	}

	return simtypes.NewOperationMsg(txCtx.Msg, true, "", txCtx.Cdc), nil, nil
}

func GetClasses(sdkCtx sdk.Context, r *rand.Rand, qryClient basetypes.QueryServer, msgType string) ([]*basetypes.ClassInfo, simtypes.OperationMsg, error) {
	ctx := sdk.WrapSDKContext(sdkCtx)
	res, err := qryClient.Classes(ctx, &basetypes.QueryClassesRequest{})
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return []*basetypes.ClassInfo{}, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no classes"), nil
		}
		return []*basetypes.ClassInfo{}, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	return res.Classes, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func GetRandomClass(sdkCtx sdk.Context, r *rand.Rand, qryClient basetypes.QueryServer, msgType string) (*basetypes.ClassInfo, simtypes.OperationMsg, error) {
	classes, op, err := GetClasses(sdkCtx, r, qryClient, msgType)
	if len(classes) == 0 {
		return nil, op, err
	}

	return classes[r.Intn(len(classes))], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func GetAccountAndSpendableCoins(ctx sdk.Context, bk ecocredit.BankKeeper,
	accs []simtypes.Account, addr, msgType string) (sdk.Coins, *simtypes.Account, simtypes.OperationMsg, error) {
	accAddr, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return nil, nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	account, found := simtypes.FindAccount(accs, accAddr)
	if !found {
		return nil, &account, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "account not found"), nil
	}

	spendable := bk.SpendableCoins(ctx, accAddr)
	return spendable, &account, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

// RandomFee generate random credit class/basket creation fee
func RandomFee(r *rand.Rand) sdk.Coin {
	// 30% chance of fee using random denom
	if r.Int63n(101) <= 30 {
		return sdk.NewCoin(simtypes.RandStringOfLength(r, 4), simtypes.RandomAmount(r, sdk.NewInt(10000)))
	}
	return sdk.NewCoin(sdk.DefaultBondDenom, simtypes.RandomAmount(r, sdk.NewInt(10000)))
}

// RandomDeposit returns minimum deposit if account have enough balance
// else returns deposit amount between (1, balance)
func RandomDeposit(r *rand.Rand, ctx sdk.Context,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper, depositParams govtypes.DepositParams, addr sdk.AccAddress,
) (deposit sdk.Coins, skip bool, err error) {
	account := ak.GetAccount(ctx, addr)
	spendable := bk.SpendableCoins(ctx, account.GetAddress())

	if spendable.Empty() {
		return nil, true, nil // skip
	}

	minDeposit := depositParams.MinDeposit
	denomIndex := r.Intn(len(minDeposit))
	denom := minDeposit[denomIndex].Denom

	depositCoins := spendable.AmountOf(denom)
	if depositCoins.IsZero() {
		return nil, true, nil
	}

	amount := depositCoins
	if amount.GT(minDeposit[denomIndex].Amount) {
		amount = minDeposit[denomIndex].Amount
	} else {
		amount, err = simtypes.RandPositiveInt(r, depositCoins)
		if err != nil {
			return nil, false, err
		}
	}

	return sdk.Coins{sdk.NewCoin(denom, amount)}, false, nil
}
