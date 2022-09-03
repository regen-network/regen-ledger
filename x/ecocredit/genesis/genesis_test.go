package genesis

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	gogoproto "github.com/gogo/protobuf/proto"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/types/known/timestamppb"

	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	sdk "github.com/cosmos/cosmos-sdk/types"

	basketapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	baseapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	basetypes "github.com/regen-network/regen-ledger/x/ecocredit/base/types/v1"
)

func TestValidateGenesis(t *testing.T) {
	t.Parallel()

	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	modDB, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)
	ss, err := baseapi.NewStateStore(modDB)
	require.NoError(t, err)

	require.NoError(t, ss.CreditTypeTable().Insert(ormCtx, &baseapi.CreditType{
		Abbreviation: "BIO",
		Name:         "biodiversity",
		Unit:         "acres",
		Precision:    6,
	}))

	require.NoError(t, ss.BatchBalanceTable().Insert(ormCtx,
		&baseapi.BatchBalance{
			BatchKey:       1,
			Address:        sdk.AccAddress("addr1"),
			TradableAmount: "90.003",
			RetiredAmount:  "9.997",
		}))

	batches := []*baseapi.Batch{
		{
			Issuer:       sdk.AccAddress("addr2"),
			ProjectKey:   1,
			Denom:        "BIO01-001-00000000-00000000-001",
			StartDate:    &timestamppb.Timestamp{Seconds: 100},
			EndDate:      &timestamppb.Timestamp{Seconds: 101},
			IssuanceDate: &timestamppb.Timestamp{Seconds: 102},
		},
		{
			Issuer:       sdk.AccAddress("addr3"),
			ProjectKey:   1,
			Denom:        "BIO02-001-00000000-00000000-001",
			StartDate:    &timestamppb.Timestamp{Seconds: 100},
			EndDate:      &timestamppb.Timestamp{Seconds: 101},
			IssuanceDate: &timestamppb.Timestamp{Seconds: 102},
		},
	}
	for _, b := range batches {
		require.NoError(t, ss.BatchTable().Insert(ormCtx, b))
	}

	require.NoError(t, ss.BatchSupplyTable().Insert(ormCtx,
		&baseapi.BatchSupply{
			BatchKey:       1,
			TradableAmount: "90.003",
			RetiredAmount:  "9.997",
		}))

	classes := []*baseapi.Class{
		{
			Id:               "BIO001",
			Admin:            sdk.AccAddress("addr4"),
			CreditTypeAbbrev: "BIO",
		},
		{
			Id:               "BIO002",
			Admin:            sdk.AccAddress("addr5"),
			CreditTypeAbbrev: "BIO",
		},
	}
	for _, c := range classes {
		require.NoError(t, ss.ClassTable().Insert(ormCtx, c))
	}

	projects := []*baseapi.Project{
		{
			Id:           "P01-001",
			Admin:        sdk.AccAddress("addr6"),
			ClassKey:     1,
			Jurisdiction: "AQ",
			Metadata:     "meta",
		},
		{
			Id:           "P02-001",
			Admin:        sdk.AccAddress("addr7"),
			ClassKey:     2,
			Jurisdiction: "AQ",
			Metadata:     "meta",
		},
	}
	for _, p := range projects {
		require.NoError(t, ss.ProjectTable().Insert(ormCtx, p))
	}

	target := ormjson.NewRawMessageTarget()
	require.NoError(t, modDB.ExportJSON(ormCtx, target))
	genesisJSON, err := target.JSON()
	require.NoError(t, err)

	params := basetypes.Params{AllowlistEnabled: true}
	err = ValidateGenesis(genesisJSON, params)
	require.NoError(t, err)
}

func TestGenesisValidate(t *testing.T) {
	t.Parallel()

	defaultParams := DefaultParams()
	addr1 := sdk.AccAddress("foobar")
	addr2 := sdk.AccAddress("fooBarBaz")
	testCases := []struct {
		id         string
		setupState func(ctx context.Context, ss baseapi.StateStore)
		params     basetypes.Params
		expectErr  bool
		errorMsg   string
	}{
		{
			"valid: no credit batches",
			func(ctx context.Context, ss baseapi.StateStore) {
				require.NoError(t, ss.CreditTypeTable().Insert(ctx, &baseapi.CreditType{
					Abbreviation: "C",
					Name:         "carbon",
					Unit:         "metric ton C02 equivalent",
					Precision:    6,
				}))
				require.NoError(t, ss.ClassTable().Insert(ctx, &baseapi.Class{
					Id:               "C01",
					Admin:            addr1,
					CreditTypeAbbrev: "C",
				}))
			},
			defaultParams,
			false,
			"",
		},
		{
			"invalid credit type abbreviation",
			func(ctx context.Context, ss baseapi.StateStore) {
				require.NoError(t, ss.CreditTypeTable().Insert(ctx, &baseapi.CreditType{
					Abbreviation: "1234",
					Name:         "carbon",
					Unit:         "kg",
					Precision:    6,
				}))
				require.NoError(t, ss.ClassTable().Insert(ctx, &baseapi.Class{
					Id:               "C01",
					Admin:            addr1,
					CreditTypeAbbrev: "C",
				}))
			},
			defaultParams,
			true,
			"must be 1-3 uppercase alphabetic characters: parse error",
		},
		{
			"invalid: credit type param",
			func(ctx context.Context, ss baseapi.StateStore) {
				require.NoError(t, ss.ClassTable().Insert(ctx, &baseapi.Class{
					Id:               "C01",
					Admin:            addr1,
					CreditTypeAbbrev: "C",
				}))
				require.NoError(t, ss.CreditTypeTable().Insert(ctx, &baseapi.CreditType{
					Abbreviation: "C",
					Name:         "carbon",
					Unit:         "kg",
					Precision:    7,
				}))
			},
			func() basetypes.Params {
				return defaultParams
			}(),
			true,
			"precision is currently locked to 6",
		},
		{
			"invalid: bad addresses in allowlist",
			func(ctx context.Context, ss baseapi.StateStore) {
			},
			func() basetypes.Params {
				p := DefaultParams()
				p.AllowlistEnabled = true
				p.AllowedClassCreators = []string{"-=!?#09)("}
				return p
			}(),
			true,
			"invalid creator address: decoding bech32 failed",
		},
		{
			"invalid: type id does not match param id",
			func(ctx context.Context, ss baseapi.StateStore) {
				require.NoError(t, ss.ClassTable().Insert(ctx, &baseapi.Class{
					Id:               "C01",
					Admin:            addr1,
					CreditTypeAbbrev: "F",
				}))
			},
			defaultParams,
			true,
			"credit type not exist",
		},
		{
			"expect error: balances are missing",
			func(ctx context.Context, ss baseapi.StateStore) {
				require.NoError(t, ss.CreditTypeTable().Insert(ctx, &baseapi.CreditType{
					Abbreviation: "C",
					Name:         "carbon",
					Unit:         "metric ton C02 equivalent",
					Precision:    6,
				}))
				denom := "C01-001-00000000-00000000-001"
				key, err := ss.ClassTable().InsertReturningID(ctx, &baseapi.Class{
					Id:               "C01",
					Admin:            addr1,
					CreditTypeAbbrev: "C",
				})
				require.NoError(t, err)

				pKey, err := ss.ProjectTable().InsertReturningID(ctx, &baseapi.Project{
					Id:           "P01-001",
					Admin:        addr1,
					ClassKey:     key,
					Jurisdiction: "AQ",
				})
				require.NoError(t, err)
				bKey, err := ss.BatchTable().InsertReturningID(ctx, &baseapi.Batch{
					Issuer:       addr1,
					ProjectKey:   pKey,
					Denom:        denom,
					StartDate:    &timestamppb.Timestamp{Seconds: 100},
					EndDate:      &timestamppb.Timestamp{Seconds: 101},
					IssuanceDate: &timestamppb.Timestamp{Seconds: 400},
				})
				require.NoError(t, err)
				require.NoError(t, ss.BatchSupplyTable().Insert(ctx, &baseapi.BatchSupply{
					BatchKey:       bKey,
					TradableAmount: "400.456",
				}))
			},
			defaultParams,
			true,
			"no balances were found",
		},
		{
			"expect error: invalid supply",
			func(ctx context.Context, ss baseapi.StateStore) {
				require.NoError(t, ss.CreditTypeTable().Insert(ctx, &baseapi.CreditType{
					Abbreviation: "C",
					Name:         "carbon",
					Unit:         "metric ton C02 equivalent",
					Precision:    6,
				}))
				denom := "C01-001-00000000-00000000-001"
				cKey, err := ss.ClassTable().InsertReturningID(ctx, &baseapi.Class{
					Id:               "C01",
					Admin:            addr1,
					CreditTypeAbbrev: "C",
				})
				require.NoError(t, err)

				pKey, err := ss.ProjectTable().InsertReturningID(ctx, &baseapi.Project{
					Id:           "P01-001",
					Admin:        addr1,
					ClassKey:     cKey,
					Jurisdiction: "AQ",
				})
				require.NoError(t, err)

				bKey, err := ss.BatchTable().InsertReturningID(ctx, &baseapi.Batch{
					Issuer:       addr1,
					ProjectKey:   pKey,
					Denom:        denom,
					StartDate:    &timestamppb.Timestamp{Seconds: 100},
					EndDate:      &timestamppb.Timestamp{Seconds: 101},
					IssuanceDate: &timestamppb.Timestamp{Seconds: 102},
				})
				require.NoError(t, err)
				require.NoError(t, ss.BatchBalanceTable().Insert(ctx, &baseapi.BatchBalance{
					BatchKey:       bKey,
					Address:        addr1,
					TradableAmount: "100",
					RetiredAmount:  "100",
				}))
				require.NoError(t, ss.BatchSupplyTable().Insert(ctx, &baseapi.BatchSupply{
					BatchKey:        bKey,
					TradableAmount:  "10",
					RetiredAmount:   "",
					CancelledAmount: "",
				}))
			},
			defaultParams,
			true,
			"supply is incorrect for 1 credit batch, expected 10, got 200: invalid coins",
		},
		{
			"valid test case",
			func(ctx context.Context, ss baseapi.StateStore) {
				require.NoError(t, ss.CreditTypeTable().Insert(ctx, &baseapi.CreditType{
					Abbreviation: "C",
					Name:         "carbon",
					Unit:         "metric ton C02 equivalent",
					Precision:    6,
				}))
				cKey, err := ss.ClassTable().InsertReturningID(ctx, &baseapi.Class{
					Id:               "C01",
					Admin:            addr1,
					CreditTypeAbbrev: "C",
				})
				require.NoError(t, err)
				pKey, err := ss.ProjectTable().InsertReturningID(ctx, &baseapi.Project{
					Id:           "P01-001",
					Admin:        addr1,
					ClassKey:     cKey,
					Jurisdiction: "AQ",
				})
				require.NoError(t, err)
				bKey, err := ss.BatchTable().InsertReturningID(ctx, &baseapi.Batch{
					Issuer:       addr1,
					ProjectKey:   pKey,
					Denom:        "C01-001-00000000-00000000-001",
					StartDate:    &timestamppb.Timestamp{Seconds: 100},
					EndDate:      &timestamppb.Timestamp{Seconds: 101},
					IssuanceDate: &timestamppb.Timestamp{Seconds: 102},
				})
				require.NoError(t, err)
				require.NoError(t, ss.BatchBalanceTable().Insert(ctx, &baseapi.BatchBalance{
					BatchKey:       bKey,
					Address:        addr1,
					TradableAmount: "100.123",
					RetiredAmount:  "100.123",
					EscrowedAmount: "10.000",
				}))
				require.NoError(t, ss.BatchBalanceTable().Insert(ctx, &baseapi.BatchBalance{
					BatchKey:       bKey,
					Address:        addr2,
					TradableAmount: "100.123",
					RetiredAmount:  "100.123",
				}))
				require.NoError(t, ss.BatchSupplyTable().Insert(ctx, &baseapi.BatchSupply{
					BatchKey:       bKey,
					TradableAmount: "210.246",
					RetiredAmount:  "200.246",
				}))
			},
			defaultParams,
			false,
			"",
		},
		{
			"valid test case, multiple classes and credit types",
			func(ctx context.Context, ss baseapi.StateStore) {
				require.NoError(t, ss.CreditTypeTable().Insert(ctx, &baseapi.CreditType{
					Abbreviation: "C",
					Name:         "carbon",
					Unit:         "metric ton C02 equivalent",
					Precision:    6,
				}))
				require.NoError(t, ss.CreditTypeTable().Insert(ctx, &baseapi.CreditType{
					Abbreviation: "BIO",
					Name:         "biodiversity",
					Unit:         "acres",
					Precision:    6,
				}))

				cKey, err := ss.ClassTable().InsertReturningID(ctx, &baseapi.Class{
					Id:               "C01",
					Admin:            addr1,
					CreditTypeAbbrev: "C",
				})
				require.NoError(t, err)
				cKeyBIO, err := ss.ClassTable().InsertReturningID(ctx, &baseapi.Class{
					Id:               "BIO01",
					Admin:            addr1,
					CreditTypeAbbrev: "BIO",
				})
				require.NoError(t, err)
				pKey, err := ss.ProjectTable().InsertReturningID(ctx, &baseapi.Project{
					Id:           "P01-001",
					Admin:        addr1,
					ClassKey:     cKey,
					Jurisdiction: "AQ",
				})
				require.NoError(t, err)
				pKeyBIO, err := ss.ProjectTable().InsertReturningID(ctx, &baseapi.Project{
					Id:           "P02-001",
					Admin:        addr1,
					ClassKey:     cKeyBIO,
					Jurisdiction: "AQ",
				})
				require.NoError(t, err)
				bKey, err := ss.BatchTable().InsertReturningID(ctx, &baseapi.Batch{
					Issuer:       addr1,
					ProjectKey:   pKey,
					Denom:        "C01-001-00000000-00000000-001",
					StartDate:    &timestamppb.Timestamp{Seconds: 100},
					EndDate:      &timestamppb.Timestamp{Seconds: 101},
					IssuanceDate: &timestamppb.Timestamp{Seconds: 102},
				})
				require.NoError(t, err)
				bKeyBIO, err := ss.BatchTable().InsertReturningID(ctx, &baseapi.Batch{
					Issuer:       addr1,
					ProjectKey:   pKeyBIO,
					Denom:        "BIO01-001-00000000-00000000-001",
					StartDate:    &timestamppb.Timestamp{Seconds: 100},
					EndDate:      &timestamppb.Timestamp{Seconds: 101},
					IssuanceDate: &timestamppb.Timestamp{Seconds: 102},
				})
				require.NoError(t, err)
				bKey2, err := ss.BatchTable().InsertReturningID(ctx, &baseapi.Batch{
					Issuer:       addr1,
					ProjectKey:   pKey,
					Denom:        "C01-001-00000000-00000000-002",
					StartDate:    &timestamppb.Timestamp{Seconds: 100},
					EndDate:      &timestamppb.Timestamp{Seconds: 101},
					IssuanceDate: &timestamppb.Timestamp{Seconds: 102},
				})
				require.NoError(t, err)
				require.NoError(t, ss.BatchBalanceTable().Insert(ctx, &baseapi.BatchBalance{
					BatchKey:       bKey,
					Address:        addr1,
					TradableAmount: "100.123",
					RetiredAmount:  "100.123",
				}))
				require.NoError(t, ss.BatchBalanceTable().Insert(ctx, &baseapi.BatchBalance{
					BatchKey:       bKey2,
					Address:        addr2,
					TradableAmount: "100.123",
					RetiredAmount:  "100.123",
				}))
				require.NoError(t, ss.BatchBalanceTable().Insert(ctx, &baseapi.BatchBalance{
					BatchKey:       bKeyBIO,
					Address:        addr1,
					TradableAmount: "105.2",
					EscrowedAmount: "102.2",
					RetiredAmount:  "207.1",
				}))
				require.NoError(t, ss.BatchBalanceTable().Insert(ctx, &baseapi.BatchBalance{
					BatchKey:       bKeyBIO,
					Address:        addr2,
					TradableAmount: "101.1",
					RetiredAmount:  "404.1",
				}))
				require.NoError(t, ss.BatchSupplyTable().Insert(ctx, &baseapi.BatchSupply{
					BatchKey:        bKey,
					TradableAmount:  "100.123",
					RetiredAmount:   "100.123",
					CancelledAmount: "",
				}))
				require.NoError(t, ss.BatchSupplyTable().Insert(ctx, &baseapi.BatchSupply{
					BatchKey:        bKey2,
					TradableAmount:  "100.123",
					RetiredAmount:   "100.123",
					CancelledAmount: "",
				}))

				require.NoError(t, ss.BatchSupplyTable().Insert(ctx, &baseapi.BatchSupply{
					BatchKey:        bKeyBIO,
					TradableAmount:  "308.5",
					RetiredAmount:   "611.2",
					CancelledAmount: "",
				}))
			},
			defaultParams,
			false,
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.id, func(t *testing.T) {
			jsn := setupStateAndExportJSON(t, tc.setupState)
			err := ValidateGenesis(jsn, tc.params)
			if tc.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errorMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestValidateGenesisWithBasketBalance(t *testing.T) {
	t.Parallel()

	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	modDB, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)
	ss, err := baseapi.NewStateStore(modDB)
	require.NoError(t, err)

	bsktStore, err := basketapi.NewStateStore(modDB)
	require.NoError(t, err)

	require.NoError(t, ss.CreditTypeTable().Insert(ormCtx, &baseapi.CreditType{
		Abbreviation: "BIO",
		Name:         "biodiversity",
		Unit:         "acres",
		Precision:    6,
	}))

	require.NoError(t, ss.BatchBalanceTable().Insert(ormCtx,
		&baseapi.BatchBalance{
			BatchKey:       1,
			Address:        sdk.AccAddress("addr1"),
			TradableAmount: "90.003",
			RetiredAmount:  "9.997",
		}))

	require.NoError(t, ss.BatchBalanceTable().Insert(ormCtx,
		&baseapi.BatchBalance{
			BatchKey:       2,
			Address:        sdk.AccAddress("addr1"),
			TradableAmount: "1.234",
			EscrowedAmount: "1.234",
			RetiredAmount:  "0",
		}))

	batches := []*baseapi.Batch{
		{
			Issuer:       sdk.AccAddress("addr2"),
			ProjectKey:   1,
			Denom:        "C01-001-20200101-20210101-001",
			StartDate:    &timestamppb.Timestamp{Seconds: 100},
			EndDate:      &timestamppb.Timestamp{Seconds: 101},
			IssuanceDate: &timestamppb.Timestamp{Seconds: 102},
		},
		{
			Issuer:       sdk.AccAddress("addr3"),
			ProjectKey:   1,
			Denom:        "BIO02-001-20200101-20210101-001",
			StartDate:    &timestamppb.Timestamp{Seconds: 100},
			EndDate:      &timestamppb.Timestamp{Seconds: 101},
			IssuanceDate: &timestamppb.Timestamp{Seconds: 102},
		},
	}
	for _, b := range batches {
		require.NoError(t, ss.BatchTable().Insert(ormCtx, b))
	}

	require.NoError(t, ss.BatchSupplyTable().Insert(ormCtx,
		&baseapi.BatchSupply{
			BatchKey:       1,
			TradableAmount: "190.003",
			RetiredAmount:  "9.997",
		}),
	)

	require.NoError(t, ss.BatchSupplyTable().Insert(ormCtx,
		&baseapi.BatchSupply{
			BatchKey:       2,
			TradableAmount: "12.468",
			RetiredAmount:  "0",
		}),
	)

	class := baseapi.Class{
		Id:               "BIO001",
		Admin:            sdk.AccAddress("addr4"),
		CreditTypeAbbrev: "BIO",
	}
	require.NoError(t, ss.ClassTable().Insert(ormCtx, &class))

	project := baseapi.Project{
		Id:           "P01-001",
		Admin:        sdk.AccAddress("addr6"),
		ClassKey:     1,
		Jurisdiction: "AQ",
		Metadata:     "meta",
	}
	require.NoError(t, ss.ProjectTable().Insert(ormCtx, &project))

	startDate1, err := types.ParseDate("start date", "2020-01-01")
	require.NoError(t, err)
	startDate2, err := types.ParseDate("start date", "2020-01-01")
	require.NoError(t, err)

	basketBalances := []*basketapi.BasketBalance{
		{
			BasketId:       1,
			BatchDenom:     "C01-001-20200101-20210101-001",
			Balance:        "100",
			BatchStartDate: timestamppb.New(startDate1),
		},
		{
			BasketId:       2,
			BatchDenom:     "BIO02-001-20200101-20210101-001",
			Balance:        "10.000",
			BatchStartDate: timestamppb.New(startDate2),
		},
	}
	for _, b := range basketBalances {
		require.NoError(t, bsktStore.BasketBalanceTable().Insert(ormCtx, b))
	}

	target := ormjson.NewRawMessageTarget()
	require.NoError(t, modDB.ExportJSON(ormCtx, target))
	genesisJSON, err := target.JSON()
	require.NoError(t, err)

	params := basetypes.Params{AllowlistEnabled: true}
	err = ValidateGenesis(genesisJSON, params)
	require.NoError(t, err)
}

// setupStateAndExportJSON sets up state as defined in the setupFunc function and then exports the ORM data as JSON.
func setupStateAndExportJSON(t *testing.T, setupFunc func(ctx context.Context, ss baseapi.StateStore)) json.RawMessage {
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	modDB, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)
	ss, err := baseapi.NewStateStore(modDB)
	require.NoError(t, err)
	setupFunc(ormCtx, ss)
	target := ormjson.NewRawMessageTarget()
	require.NoError(t, modDB.ExportJSON(ormCtx, target))
	jsn, err := target.JSON()
	require.NoError(t, err)
	return jsn
}

func TestMergeClassFeeIntoTarget(t *testing.T) {
	cdc := codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
	db, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)

	jsonTarget := ormjson.NewRawMessageTarget()
	err = db.DefaultJSON(jsonTarget)
	require.NoError(t, err)

	classFee := DefaultClassFee()
	err = MergeClassFeeIntoTarget(cdc, classFee, jsonTarget)
	require.NoError(t, err)

	raw, err := jsonTarget.JSON()
	require.NoError(t, err)

	jsonSource, err := ormjson.NewRawMessageSource(raw)
	require.NoError(t, err)

	r, err := jsonSource.OpenReader(protoreflect.FullName(gogoproto.MessageName(&classFee)))
	require.NoError(t, err)

	var expected baseapi.ClassFee
	err = (&jsonpb.Unmarshaler{AllowUnknownFields: true}).Unmarshal(r, &expected)
	require.NoError(t, err)

	require.NotEmpty(t, classFee.Fee)
	require.Equal(t, expected.Fee.Amount, classFee.Fee.Amount.String())
	require.Equal(t, expected.Fee.Denom, classFee.Fee.Denom)
}

func TestMergeBasketFeeIntoTarget(t *testing.T) {
	cdc := codec.NewProtoCodec(codectypes.NewInterfaceRegistry())
	db, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)

	jsonTarget := ormjson.NewRawMessageTarget()
	err = db.DefaultJSON(jsonTarget)
	require.NoError(t, err)

	basketFee := DefaultBasketFee()
	err = MergeBasketFeeIntoTarget(cdc, basketFee, jsonTarget)
	require.NoError(t, err)

	raw, err := jsonTarget.JSON()
	require.NoError(t, err)

	jsonSource, err := ormjson.NewRawMessageSource(raw)
	require.NoError(t, err)

	r, err := jsonSource.OpenReader(protoreflect.FullName(gogoproto.MessageName(&basketFee)))
	require.NoError(t, err)

	var expected basketapi.BasketFee
	err = (&jsonpb.Unmarshaler{AllowUnknownFields: true}).Unmarshal(r, &expected)
	require.NoError(t, err)

	require.NotEmpty(t, basketFee.Fee)
	require.Equal(t, basketFee.Fee.Amount.String(), expected.Fee.Amount)
	require.Equal(t, basketFee.Fee.Denom, expected.Fee.Denom)
}
