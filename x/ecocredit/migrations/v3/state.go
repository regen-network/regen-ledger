package v3

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	basketapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/basket/v1"
	baseapi "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
	regentypes "github.com/regen-network/regen-ledger/types/v2"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/v3/base/types/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/v3/server/utils"
)

// MigrateState performs in-place store migrations from ConsensusVersion 2 to 3.
func MigrateState(sdkCtx sdk.Context, baseStore baseapi.StateStore, basketStore basketapi.StateStore, subspace paramtypes.Subspace) error {
	var params basetypes.Params
	subspace.GetParamSet(sdkCtx, &params)

	// validate credit class fee
	if err := params.CreditClassFee.Validate(); err != nil {
		return err
	}

	// migrate credit class fees
	classFees := regentypes.CoinsToProtoCoins(params.CreditClassFee)
	if err := baseStore.ClassFeeTable().Save(sdkCtx, &baseapi.ClassFee{
		Fee: classFees[0], // we assume there is one fee at the time of the upgrade
	}); err != nil {
		return err
	}

	// migrate credit class allow list
	if err := baseStore.ClassCreatorAllowlistTable().Save(sdkCtx, &baseapi.ClassCreatorAllowlist{
		Enabled: params.AllowlistEnabled,
	}); err != nil {
		return err
	}

	// migrate allowed class creators to orm table
	for _, creator := range params.AllowedClassCreators {
		addr, err := sdk.AccAddressFromBech32(creator)
		if err != nil {
			return err
		}

		if err := baseStore.AllowedClassCreatorTable().Save(sdkCtx, &baseapi.AllowedClassCreator{
			Address: addr,
		}); err != nil {
			return err
		}
	}

	// verify basket fee is valid
	if err := params.BasketFee.Validate(); err != nil {
		return err
	}

	// migrate basket params
	basketFees := regentypes.CoinsToProtoCoins(params.BasketFee)
	if err := basketStore.BasketFeeTable().Save(sdkCtx, &basketapi.BasketFee{
		Fee: basketFees[0], // we assume there is one fee at the time of the upgrade
	}); err != nil {
		return err
	}

	var balances []BalanceInfo
	if sdkCtx.ChainID() == "regen-1" {
		balances = getMainnetState()
	}

	if sdkCtx.ChainID() == "regen-redwood-1" {
		balances = getRedwoodState()
	}

	for _, balance := range balances {
		if err := migrateBatchBalance(sdkCtx, baseStore, balance); err != nil {
			return err
		}
	}

	return nil
}

const precision = 6

type Balance struct {
	BatchDenom string
	Amount     string
}

type BalanceInfo struct {
	Sender   string
	Balances []Balance
}

func getRedwoodState() []BalanceInfo {
	return []BalanceInfo{
		{
			Sender: "regen1df675r9vnf7pdedn4sf26svdsem3ugavgxmy46",
			Balances: []Balance{
				{
					BatchDenom: "C01-001-20170101-20180101-003",
					Amount:     "1",
				},
				{
					BatchDenom: "C01-001-20170101-20180101-005",
					Amount:     "2",
				},
			},
		},
		{
			Sender: "regen1m6d7al7yrgwv6j6sczt382x33yhxrtrxz2q09z",
			Balances: []Balance{
				{
					BatchDenom: "C01-001-20170606-20210601-007",
					Amount:     "0.01",
				},
			},
		},
		{
			Sender: "regen1m0qh5y4ejkz3l5n6jlrntxcqx9r0x9xjv4vpcp",
			Balances: []Balance{
				{
					BatchDenom: "C01-001-20170606-20210601-007",
					Amount:     "2.490054",
				},
			},
		},
	}

}

func getMainnetState() []BalanceInfo {
	return []BalanceInfo{
		{
			Sender: "regen1l8v5nzznewg9cnfn0peg22mpysdr3a8jcm4p8v",
			Balances: []Balance{
				{
					BatchDenom: "C02-001-20180101-20181231-001",
					Amount:     "0.05",
				},
			},
		},
	}
}

// migrateBatchBalance adds lost batch credits to sender account.
func migrateBatchBalance(ctx sdk.Context, baseStore baseapi.StateStore, balanceInfo BalanceInfo) error {
	senderAddr, err := sdk.AccAddressFromBech32(balanceInfo.Sender)
	if err != nil {
		return err
	}

	for _, balance := range balanceInfo.Balances {

		batchInfo, err := baseStore.BatchTable().GetByDenom(ctx, balance.BatchDenom)
		if err != nil {
			return err
		}

		userBalance, err := baseStore.BatchBalanceTable().Get(ctx, senderAddr, batchInfo.Key)
		if err != nil {
			return err
		}

		decs, err := utils.GetNonNegativeFixedDecs(precision, userBalance.TradableAmount, balance.Amount)
		if err != nil {
			return err
		}

		tradableAmount, creditsToAdd := decs[0], decs[1]

		result, err := tradableAmount.Add(creditsToAdd)
		if err != nil {
			return err
		}

		userBalance.TradableAmount = result.String()

		if err := baseStore.BatchBalanceTable().Update(ctx, userBalance); err != nil {
			return err
		}
	}
	return nil
}
