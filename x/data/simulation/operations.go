package simulation

import (
	"crypto"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/simapp/helpers"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/data"
)

// Simulation operation weights constants
const (
	OpWeightMsgAnchor           = "op_weight_msg_anchor"
	OpWeightMsgAttest           = "op_weight_msg_attest"
	OpWeightMsgDefineResolver   = "op_weight_msg_define_resolver"
	OpWeightMsgRegisterResolver = "op_weight_msg_register_resolver"

	ModuleName  = "data"
	resolverUrl = "https://foo.bar"
)

var (
	TypeMsgAnchor           = data.MsgAnchor{}.Type()
	TypeMsgAttest           = data.MsgAttest{}.Type()
	TypeMsgDefineResolver   = data.MsgDefineResolver{}.Type()
	TypeMsgRegisterResolver = data.MsgRegisterResolver{}.Type()
)

const (
	WeightAnchor           = 100
	WeightAttest           = 100
	WeightRegisterResolver = 100
	WeightDefineResolver   = 100
)

// WeightedOperations returns all the operations from the data module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak data.AccountKeeper, bk data.BankKeeper,
	qryClient data.QueryClient) simulation.WeightedOperations {

	var (
		weightMsgAnchor           int
		weightMsgAttest           int
		weightMsgDefineResolver   int
		weightMsgRegisterResolver int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgAnchor, &weightMsgAnchor, nil,
		func(_ *rand.Rand) {
			weightMsgAnchor = WeightAnchor
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgAttest, &weightMsgAttest, nil,
		func(_ *rand.Rand) {
			weightMsgAttest = WeightAttest
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgDefineResolver, &weightMsgDefineResolver, nil,
		func(_ *rand.Rand) {
			weightMsgDefineResolver = WeightDefineResolver
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRegisterResolver, &weightMsgRegisterResolver, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterResolver = WeightRegisterResolver
		},
	)

	ops := simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgAnchor,
			SimulateMsgAnchor(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgAnchor,
			SimulateMsgAnchor(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgAnchor,
			SimulateMsgAnchor(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgAnchor,
			SimulateMsgAnchor(ak, bk, qryClient),
		),
	}

	return ops
}

// SimulateMsgAnchor generates a MsgAnchor with random values.
func SimulateMsgAnchor(ak data.AccountKeeper, bk data.BankKeeper,
	qryClient data.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		sender, _ := simtypes.RandomAcc(r, accs)

		contentHash, err := getContentHash(r)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgAnchor, err.Error()), nil, err
		}

		spendable := bk.SpendableCoins(sdkCtx, sender.Address)
		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgAnchor, "fee error"), nil, err
		}

		msg := &data.MsgAnchor{
			Sender: sender.Address.String(),
			Hash:   contentHash,
		}

		account := ak.GetAccount(sdkCtx, sender.Address)
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			sender.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgAnchor, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgAnchor, "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

// SimulateMsgAttest generates a MsgAttest with random values.
func SimulateMsgAttest(ak data.AccountKeeper, bk data.BankKeeper,
	qryClient data.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		attestors := randomAttestors(r, accs)
		acc, _ := simtypes.RandomAcc(r, accs)

		spendable := bk.SpendableCoins(sdkCtx, acc.Address)
		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgAttest, "fee error"), nil, err
		}

		hash1, err := getContentHash(r)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgAttest, err.Error()), nil, err
		}

		msg := &data.MsgAttest{
			Attestors: attestors,
			Hash:      hash1.GetGraph(),
		}

		account := ak.GetAccount(sdkCtx, acc.Address)
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			acc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgAttest, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgAttest, "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

// SimulateMsgDefineResolver generates a MsgDefineResolver with random values.
func SimulateMsgDefineResolver(ak data.AccountKeeper, bk data.BankKeeper,
	qryClient data.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		manager, _ := simtypes.RandomAcc(r, accs)

		spendable := bk.SpendableCoins(sdkCtx, manager.Address)
		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgDefineResolver, "fee error"), nil, err
		}

		msg := &data.MsgDefineResolver{
			Manager:     manager.Address.String(),
			ResolverUrl: resolverUrl,
		}

		account := ak.GetAccount(sdkCtx, manager.Address)
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			manager.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgDefineResolver, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgDefineResolver, "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

// SimulateMsgRegisterResolver generates a MsgRegisterResolver with random values.
func SimulateMsgRegisterResolver(ak data.AccountKeeper, bk data.BankKeeper,
	qryClient data.QueryClient) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		ctx := sdk.WrapSDKContext(sdkCtx)
		res, err := qryClient.ResolverInfo(ctx, &data.QueryResolverInfoRequest{Url: resolverUrl})
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgRegisterResolver, err.Error()), nil, nil // not found
		}

		manager, err := sdk.AccAddressFromBech32(res.Manager)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgRegisterResolver, err.Error()), nil, err
		}

		managerAcc, found := simtypes.FindAccount(accs, manager)
		if !found {
			return simtypes.NoOpMsg(ModuleName, TypeMsgRegisterResolver, "not a sim account"), nil, nil
		}

		spendable := bk.SpendableCoins(sdkCtx, managerAcc.Address)
		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgRegisterResolver, "fee error"), nil, err
		}

		msg := &data.MsgRegisterResolver{
			Manager:    manager.String(),
			ResolverId: res.Id,
		}

		account := ak.GetAccount(sdkCtx, manager)
		txGen := simappparams.MakeTestEncodingConfig().TxConfig
		tx, err := helpers.GenTx(
			txGen,
			[]sdk.Msg{msg},
			fees,
			helpers.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			managerAcc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgRegisterResolver, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.Deliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(ModuleName, TypeMsgRegisterResolver, "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, "", nil), nil, nil
	}
}

func getContentHash(r *rand.Rand) (*data.ContentHash, error) {
	content := []byte(simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 2, 100)))
	hash := crypto.BLAKE2b_256.New()
	_, err := hash.Write(content)
	if err != nil {
		return nil, err
	}

	digest := hash.Sum(nil)
	if r.Float32() < 0.5 {
		graphHash := &data.ContentHash_Graph{
			Hash:                      digest,
			DigestAlgorithm:           data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			CanonicalizationAlgorithm: data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_URDNA2015,
		}
		return &data.ContentHash{Graph: graphHash}, nil
	} else {
		rawHash := &data.ContentHash_Raw{
			Hash:            digest,
			DigestAlgorithm: data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256,
			MediaType:       data.RawMediaType_RAW_MEDIA_TYPE_UNSPECIFIED,
		}
		return &data.ContentHash{Raw: rawHash}, nil
	}
}

func randomAttestors(r *rand.Rand, accounts []simtypes.Account) []string {
	n := simtypes.RandIntBetween(r, 3, 10)
	attestors := make([]string, n)
	for i := 0; i < n; i++ {
		acc, _ := simtypes.RandomAcc(r, accounts)
		attestors[i] = acc.Address.String()
	}

	return attestors
}
