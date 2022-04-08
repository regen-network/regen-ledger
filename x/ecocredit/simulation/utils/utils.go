package utils

import (
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	regentypes "github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

func GetAndShuffleClasses(sdkCtx sdk.Context, r *rand.Rand, qryClient core.QueryClient) ([]*core.ClassInfo, error) {
	ctx := regentypes.Context{Context: sdkCtx}
	res, err := qryClient.Classes(ctx, &core.QueryClassesRequest{})
	if err != nil {
		return nil, err
	}

	classes := res.Classes
	if len(classes) <= 1 {
		return classes, nil
	}

	r.Shuffle(len(classes), func(i, j int) { classes[i], classes[j] = classes[j], classes[i] })
	return classes, nil
}

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
func GenAndDeliverTxWithRandFees(txCtx simulation.OperationInput) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	account := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)
	spendable := txCtx.Bankkeeper.SpendableCoins(txCtx.Context, account.GetAddress())

	var fees sdk.Coins
	var err error

	coins, hasNeg := spendable.SafeSub(txCtx.CoinsSpentInMsg)
	if hasNeg {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "message doesn't leave room for fees"), nil, err
	}

	fees, err = simtypes.RandomFees(txCtx.R, txCtx.Context, coins)
	if err != nil {
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to generate fees"), nil, err
	}
	return GenAndDeliverTx(txCtx, fees)
}

// GenAndDeliverTx generates a transactions and delivers it.
func GenAndDeliverTx(txCtx simulation.OperationInput, fees sdk.Coins) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
	account := txCtx.AccountKeeper.GetAccount(txCtx.Context, txCtx.SimAccount.Address)
	tx, err := helpers.GenTx(
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

	_, _, err = txCtx.App.Deliver(txCtx.TxGen.TxEncoder(), tx)
	if err != nil {
		if strings.Contains(err.Error(), "insufficient funds") {
			return simtypes.NoOpMsg(ecocredit.ModuleName, txCtx.MsgType, "not enough balance"), nil, nil
		}
		return simtypes.NoOpMsg(txCtx.ModuleName, txCtx.MsgType, "unable to deliver tx"), nil, err
	}

	return simtypes.NewOperationMsg(txCtx.Msg, true, "", txCtx.Cdc), nil, nil
}

func GetRandomClass(ctx sdk.Context, r *rand.Rand, qryClient core.QueryClient, msgType string) (*core.ClassInfo, simtypes.OperationMsg, error) {
	classes, err := GetAndShuffleClasses(ctx, r, qryClient)
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no credit classes"), nil
		}
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	return classes[0], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
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
