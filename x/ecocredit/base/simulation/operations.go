package simulation

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"

	regenerrors "github.com/regen-network/regen-ledger/types/v2/errors"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3"
	types "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	baskettypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/basket/types/v1"
	markettypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/marketplace/types/v1"
)

// Simulation operation weights constants
const ()

// ecocredit message types
var ()

// WeightedOperations returns all the operations from the module with their respective weights
func WeightedOperations(
	appParams simtypes.AppParams, cdc codec.JSONCodec,
	ak ecocredit.AccountKeeper, bk ecocredit.BankKeeper,
	govk ecocredit.GovKeeper,
	qryClient types.QueryServer, _ baskettypes.QueryServer,
	_ markettypes.QueryServer, authority sdk.AccAddress) simulation.WeightedOperations {

	var (
		weightMsgCreateClass              int
		weightMsgCreateBatch              int
		weightMsgSend                     int
		weightMsgRetire                   int
		weightMsgCancel                   int
		weightMsgUpdateClassAdmin         int
		weightMsgUpdateClassIssuers       int
		weightMsgUpdateClassMetadata      int
		weightMsgCreateProject            int
		weightMsgUpdateProjectAdmin       int
		weightMsgUpdateProjectMetadata    int
		weightMsgUpdateBatchMetadata      int
		weightMsgSealBatch                int
		weightMsgMintBatchCredits         int
		weightMsgBridge                   int
		weightMsgAddCreditType            int
		weightMsgAddClassCreator          int
		weightMsgRemoveClassCreator       int
		weightMsgSetClassCreatorAllowlist int
		weightMsgUpdateClassFee           int
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateClass, &weightMsgCreateClass, nil,
		func(_ *rand.Rand) {
			weightMsgCreateClass = WeightCreateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateProject, &weightMsgCreateProject, nil,
		func(_ *rand.Rand) {
			weightMsgCreateProject = WeightCreateProject
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCreateBatch, &weightMsgCreateBatch, nil,
		func(_ *rand.Rand) {
			weightMsgCreateBatch = WeightCreateBatch
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgSend, &weightMsgSend, nil,
		func(_ *rand.Rand) {
			weightMsgSend = WeightSend
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRetire, &weightMsgRetire, nil,
		func(_ *rand.Rand) {
			weightMsgRetire = WeightRetire
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgCancel, &weightMsgCancel, nil,
		func(_ *rand.Rand) {
			weightMsgCancel = WeightCancel
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassAdmin, &weightMsgUpdateClassAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassAdmin = WeightUpdateClass
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassIssuers, &weightMsgUpdateClassIssuers, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassIssuers = MsgUpdateClassIssuers
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassMetadata, &weightMsgUpdateClassMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassMetadata = WeightUpdateClassMetadata
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateProjectAdmin, &weightMsgUpdateProjectAdmin, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProjectAdmin = WeightUpdateProjectAdmin
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateProjectMetadata, &weightMsgUpdateProjectMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateProjectMetadata = WeightUpdateProjectMetadata
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateBatchMetadata, &weightMsgUpdateBatchMetadata, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateBatchMetadata = WeightUpdateBatchMetadata
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgMintBatchCredits, &weightMsgMintBatchCredits, nil,
		func(_ *rand.Rand) {
			weightMsgMintBatchCredits = WeightMintBatchCredits
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgSealBatch, &weightMsgSealBatch, nil,
		func(_ *rand.Rand) {
			weightMsgSealBatch = WeightSealBatch
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgBridge, &weightMsgBridge, nil,
		func(_ *rand.Rand) {
			weightMsgBridge = WeightBridge
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgAddCreditType, &weightMsgAddCreditType, nil,
		func(_ *rand.Rand) {
			weightMsgAddCreditType = WeightAddCreditType
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgAddClassCreator, &weightMsgAddClassCreator, nil,
		func(_ *rand.Rand) {
			weightMsgAddClassCreator = WeightAddClassCreator
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgRemoveClassCreator, &weightMsgRemoveClassCreator, nil,
		func(_ *rand.Rand) {
			weightMsgRemoveClassCreator = WeightRemoveClassCreator
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgSetClassCreatorAllowlist, &weightMsgSetClassCreatorAllowlist, nil,
		func(_ *rand.Rand) {
			weightMsgSetClassCreatorAllowlist = WeightSetClassCreatorAllowlist
		},
	)

	appParams.GetOrGenerate(cdc, OpWeightMsgUpdateClassFee, &weightMsgUpdateClassFee, nil,
		func(_ *rand.Rand) {
			weightMsgUpdateClassFee = WeightUpdateClassFee
		},
	)

	ops := simulation.WeightedOperations{
		simulation.NewWeightedOperation(
			weightMsgCreateClass,
			SimulateMsgCreateClass(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateProject,
			SimulateMsgCreateProject(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCreateBatch,
			SimulateMsgCreateBatch(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgSend,
			SimulateMsgSend(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgRetire,
			SimulateMsgRetire(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgCancel,
			SimulateMsgCancel(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassAdmin,
			SimulateMsgUpdateClassAdmin(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassIssuers,
			SimulateMsgUpdateClassIssuers(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassMetadata,
			SimulateMsgUpdateClassMetadata(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassAdmin,
			SimulateMsgUpdateProjectAdmin(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateProjectMetadata,
			SimulateMsgUpdateProjectMetadata(ak, bk, qryClient),
		),

		simulation.NewWeightedOperation(
			weightMsgMintBatchCredits,
			SimulateMsgMintBatchCredits(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgSealBatch,
			SimulateMsgSealBatch(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgBridge,
			SimulateMsgBridge(ak, bk, qryClient),
		),
		simulation.NewWeightedOperation(
			weightMsgAddCreditType,
			SimulateMsgAddCreditType(ak, bk, govk, qryClient, authority),
		),

		simulation.NewWeightedOperation(
			weightMsgAddClassCreator,
			SimulateMsgAddClassCreator(ak, bk, govk, qryClient, authority),
		),

		simulation.NewWeightedOperation(
			weightMsgRemoveClassCreator,
			SimulateMsgRemoveClassCreator(ak, bk, govk, qryClient, authority),
		),

		simulation.NewWeightedOperation(
			weightMsgSetClassCreatorAllowlist,
			SimulateMsgSetClassCreatorAllowlist(ak, bk, govk, qryClient, authority),
		),
		simulation.NewWeightedOperation(
			weightMsgUpdateClassFee,
			SimulateMsgUpdateClassFee(ak, bk, govk, qryClient, authority),
		),
	}

	return ops
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}

	return false
}

func getClassIssuers(ctx sdk.Context, qryClient types.QueryServer, className string, msgType string) ([]string, simtypes.OperationMsg, error) {
	classIssuers, err := qryClient.ClassIssuers(sdk.WrapSDKContext(ctx), &types.QueryClassIssuersRequest{ClassId: className})
	if err != nil {
		if ormerrors.IsNotFound(err) {
			return []string{}, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no credit classes"), nil
		}

		return []string{}, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	return classIssuers.Issuers, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func getRandomProjectFromClass(ctx context.Context, r *rand.Rand, qryClient types.QueryServer, msgType, classID string) (*types.ProjectInfo, simtypes.OperationMsg, error) {
	res, err := qryClient.ProjectsByClass(ctx, &types.QueryProjectsByClassRequest{
		ClassId: classID,
	})
	if err != nil {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	projects := res.Projects
	if len(projects) == 0 {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, "no projects"), nil
	}

	return projects[r.Intn(len(projects))], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func getRandomBatchFromProject(ctx context.Context, r *rand.Rand, qryClient types.QueryServer, msgType, projectID string) (*types.BatchInfo, simtypes.OperationMsg, error) {
	res, err := qryClient.BatchesByProject(ctx, &types.QueryBatchesByProjectRequest{
		ProjectId: projectID,
	})
	if err != nil {
		if regenerrors.ErrNotFound.Is(err) {
			return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, fmt.Sprintf("no credit batches for %s project", projectID)), nil
		}
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, err.Error()), err
	}

	batches := res.Batches
	if len(batches) == 0 {
		return nil, simtypes.NoOpMsg(ecocredit.ModuleName, msgType, fmt.Sprintf("no credit batches for %s project", projectID)), nil
	}
	return batches[r.Intn(len(batches))], simtypes.NoOpMsg(ecocredit.ModuleName, msgType, ""), nil
}

func randomIssuers(r *rand.Rand, accounts []simtypes.Account) []string {
	n := simtypes.RandIntBetween(r, 3, 10)

	var issuers []string
	issuersMap := make(map[string]bool)
	for i := 0; i < n; i++ {
		acc, _ := simtypes.RandomAcc(r, accounts)
		addr := acc.Address.String()
		if _, ok := issuersMap[addr]; ok {
			continue
		}
		issuersMap[acc.Address.String()] = true
		issuers = append(issuers, addr)
	}

	return issuers
}

func generateBatchIssuance(r *rand.Rand, accs []simtypes.Account) []*types.BatchIssuance {
	numIssuances := simtypes.RandIntBetween(r, 3, 10)
	res := make([]*types.BatchIssuance, numIssuances)

	for i := 0; i < numIssuances; i++ {
		recipient := accs[i]
		retiredAmount := simtypes.RandIntBetween(r, 0, 100)
		var retirementJurisdiction string
		if retiredAmount > 0 {
			retirementJurisdiction = "AD"
		}
		res[i] = &types.BatchIssuance{
			Recipient:              recipient.Address.String(),
			TradableAmount:         fmt.Sprintf("%d", simtypes.RandIntBetween(r, 10, 1000)),
			RetiredAmount:          fmt.Sprintf("%d", retiredAmount),
			RetirementJurisdiction: retirementJurisdiction,
		}
	}

	return res
}
