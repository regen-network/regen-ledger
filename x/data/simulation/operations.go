package simulation

import (
	"crypto"
	"fmt"
	"math/rand"
	"strings"

	"github.com/cosmos/cosmos-sdk/baseapp"
	simtestutil "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module/testutil"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	"github.com/regen-network/regen-ledger/x/data/v3"
)

// Simulation operation weights constants
const (
	OpWeightMsgAnchor           = "op_weight_msg_anchor"            //nolint: gosec
	OpWeightMsgAttest           = "op_weight_msg_attest"            //nolint: gosec
	OpWeightMsgDefineResolver   = "op_weight_msg_define_resolver"   //nolint: gosec
	OpWeightMsgRegisterResolver = "op_weight_msg_register_resolver" //nolint: gosec
)

var (
	TypeMsgAnchor           = sdk.MsgTypeURL(&data.MsgAnchor{})
	TypeMsgAttest           = sdk.MsgTypeURL(&data.MsgAttest{})
	TypeMsgDefineResolver   = sdk.MsgTypeURL(&data.MsgDefineResolver{})
	TypeMsgRegisterResolver = sdk.MsgTypeURL(&data.MsgRegisterResolver{})
)

const (
	WeightAnchor           = 100
	WeightAttest           = 100
	WeightRegisterResolver = 100
	WeightDefineResolver   = 100
)

// WeightedOperations returns all the operations from the data module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams,
	ak data.AccountKeeper, bk data.BankKeeper,
	qryClient data.QueryServer,
) simulation.WeightedOperations {
	var (
		weightMsgAnchor           int
		weightMsgAttest           int
		weightMsgDefineResolver   int
		weightMsgRegisterResolver int
	)

	appParams.GetOrGenerate(OpWeightMsgAnchor, &weightMsgAnchor, nil,
		func(_ *rand.Rand) {
			weightMsgAnchor = WeightAnchor
		},
	)

	appParams.GetOrGenerate(OpWeightMsgAttest, &weightMsgAttest, nil,
		func(_ *rand.Rand) {
			weightMsgAttest = WeightAttest
		},
	)

	appParams.GetOrGenerate(OpWeightMsgDefineResolver, &weightMsgDefineResolver, nil,
		func(_ *rand.Rand) {
			weightMsgDefineResolver = WeightDefineResolver
		},
	)

	appParams.GetOrGenerate(OpWeightMsgRegisterResolver, &weightMsgRegisterResolver, nil,
		func(_ *rand.Rand) {
			weightMsgRegisterResolver = WeightRegisterResolver
		},
	)

	ops := simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgAnchor,
			SimulateMsgAnchor(ak, bk),
		),
		simulation.NewWeightedOperation(
			weightMsgAnchor,
			SimulateMsgAttest(ak, bk),
		),
		simulation.NewWeightedOperation(
			weightMsgAnchor,
			SimulateMsgDefineResolver(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgAnchor,
			SimulateMsgRegisterResolver(ak, bk, qryClient),
		),
	}

	return ops
}

// SimulateMsgAnchor generates a MsgAnchor with random values.
func SimulateMsgAnchor(ak data.AccountKeeper, bk data.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		sender, _ := simtypes.RandomAcc(r, accs)

		contentHash, err := getContentHash(r)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgAnchor, err.Error()), nil, err
		}

		spendable := bk.SpendableCoins(sdkCtx, sender.Address)
		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgAnchor, "fee error"), nil, err
		}

		msg := &data.MsgAnchor{
			Sender:      sender.Address.String(),
			ContentHash: contentHash,
		}

		account := ak.GetAccount(sdkCtx, sender.Address)
		txGen := testutil.MakeTestEncodingConfig().TxConfig
		tx, err := simtestutil.GenSignedMockTx(
			r,
			txGen,
			[]sdk.Msg{msg},
			fees,
			simtestutil.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			sender.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgAnchor, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgAnchor, "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgAttest generates a MsgAttest with random values.
func SimulateMsgAttest(ak data.AccountKeeper, bk data.BankKeeper) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		attestor, _ := simtypes.RandomAcc(r, accs)

		spendable := bk.SpendableCoins(sdkCtx, attestor.Address)
		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgAttest, "fee error"), nil, err
		}

		content := []byte(simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 2, 100)))
		hash := crypto.BLAKE2b_256.New()
		_, err = hash.Write(content)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgAttest, err.Error()), nil, err
		}

		digest := hash.Sum(nil)
		msg := &data.MsgAttest{
			Attestor:      attestor.Address.String(),
			ContentHashes: []*data.ContentHash_Graph{getGraph(digest)},
		}

		txGen := testutil.MakeTestEncodingConfig().TxConfig

		account := ak.GetAccount(sdkCtx, attestor.Address)
		tx, err := simtestutil.GenSignedMockTx(
			r,
			txGen,
			[]sdk.Msg{msg},
			fees,
			simtestutil.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			attestor.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgAttest, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			if strings.Contains(err.Error(), "insufficient funds") {
				return simtypes.NoOpMsg(data.ModuleName, TypeMsgAttest, "not enough balance"), nil, nil
			}
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgAttest, "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

// SimulateMsgDefineResolver generates a MsgDefineResolver with random values.
func SimulateMsgDefineResolver(ak data.AccountKeeper, bk data.BankKeeper, qryClient data.QueryServer) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		manager, _ := simtypes.RandomAcc(r, accs)

		spendable := bk.SpendableCoins(sdkCtx, manager.Address)
		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgDefineResolver, "fee error"), nil, err
		}

		resolverURL := genResolverURL(r)
		ctx := sdkCtx
		result, err := qryClient.ResolversByURL(ctx, &data.QueryResolversByURLRequest{Url: resolverURL})
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgDefineResolver, err.Error()), nil, err
		}

		for _, resolver := range result.Resolvers {
			if resolver.Url == resolverURL && resolver.Manager == manager.Address.String() {
				return simtypes.NoOpMsg(data.ModuleName, TypeMsgDefineResolver, "resolver with the same URL and manager already exists"), nil, nil
			}
		}

		msg := &data.MsgDefineResolver{
			Definer:     manager.Address.String(),
			ResolverUrl: resolverURL,
		}

		account := ak.GetAccount(sdkCtx, manager.Address)
		txGen := testutil.MakeTestEncodingConfig().TxConfig
		tx, err := simtestutil.GenSignedMockTx(
			r,
			txGen,
			[]sdk.Msg{msg},
			fees,
			simtestutil.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			manager.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgDefineResolver, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			if strings.Contains(err.Error(), "resolver URL already exists") {
				return simtypes.NoOpMsg(data.ModuleName, TypeMsgDefineResolver, "resolver URL already exists"), nil, nil
			}
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgDefineResolver, "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
	}
}

func genResolverURL(r *rand.Rand) string {
	return fmt.Sprintf("https://%s.com", simtypes.RandStringOfLength(r, simtypes.RandIntBetween(r, 2, 3)))
}

// SimulateMsgRegisterResolver generates a MsgRegisterResolver with random values.
func SimulateMsgRegisterResolver(ak data.AccountKeeper, bk data.BankKeeper,
	qryClient data.QueryServer,
) simtypes.Operation {
	return func(
		r *rand.Rand, app *baseapp.BaseApp, sdkCtx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		ctx := sdkCtx
		resolverID := r.Uint64()
		res, err := qryClient.Resolver(ctx, &data.QueryResolverRequest{Id: resolverID})
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgRegisterResolver, err.Error()), nil, nil // not found
		}

		manager, err := sdk.AccAddressFromBech32(res.Resolver.Manager)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgRegisterResolver, err.Error()), nil, err
		}

		managerAcc, found := simtypes.FindAccount(accs, manager)
		if !found {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgRegisterResolver, "not a sim account"), nil, nil
		}

		spendable := bk.SpendableCoins(sdkCtx, managerAcc.Address)
		fees, err := simtypes.RandomFees(r, sdkCtx, spendable)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgRegisterResolver, "fee error"), nil, err
		}

		contentHash, err := getContentHash(r)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgRegisterResolver, err.Error()), nil, err
		}
		msg := &data.MsgRegisterResolver{
			Signer:        manager.String(),
			ResolverId:    resolverID,
			ContentHashes: []*data.ContentHash{contentHash},
		}

		account := ak.GetAccount(sdkCtx, manager)
		txGen := testutil.MakeTestEncodingConfig().TxConfig
		tx, err := simtestutil.GenSignedMockTx(
			r,
			txGen,
			[]sdk.Msg{msg},
			fees,
			simtestutil.DefaultGenTxGas,
			chainID,
			[]uint64{account.GetAccountNumber()},
			[]uint64{account.GetSequence()},
			managerAcc.PrivKey,
		)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgRegisterResolver, "unable to generate mock tx"), nil, err
		}

		_, _, err = app.SimDeliver(txGen.TxEncoder(), tx)
		if err != nil {
			return simtypes.NoOpMsg(data.ModuleName, TypeMsgRegisterResolver, "unable to deliver tx"), nil, err
		}

		return simtypes.NewOperationMsg(msg, true, ""), nil, nil
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
		return &data.ContentHash{Graph: getGraph(digest)}, nil
	}

	return &data.ContentHash{Raw: getRaw(digest)}, nil
}

func getGraph(digest []byte) *data.ContentHash_Graph {
	return &data.ContentHash_Graph{
		Hash:                      digest,
		DigestAlgorithm:           uint32(data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256),
		CanonicalizationAlgorithm: uint32(data.GraphCanonicalizationAlgorithm_GRAPH_CANONICALIZATION_ALGORITHM_RDFC_1_0),
	}
}

func getRaw(digest []byte) *data.ContentHash_Raw {
	return &data.ContentHash_Raw{
		Hash:            digest,
		DigestAlgorithm: uint32(data.DigestAlgorithm_DIGEST_ALGORITHM_BLAKE2B_256),
		FileExtension:   "bin",
	}
}
