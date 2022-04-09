package core_test

import (
	fmt "fmt"
	"testing"

	"github.com/cosmos/cosmos-sdk/orm/model/ormdb"
	"github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	"github.com/cosmos/cosmos-sdk/orm/testing/ormtest"
	"github.com/cosmos/cosmos-sdk/orm/types/ormjson"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/store"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
	"github.com/tendermint/tendermint/libs/log"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	dbm "github.com/tendermint/tm-db"
	"gotest.tools/assert"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
	"github.com/regen-network/regen-ledger/x/ecocredit/server"
)

func TestValidateGenesis(t *testing.T) {
	x := `{"regen.ecocredit.v1.BatchBalance":[{"address":"gydQIvR2RUi0N1RJnmgOLVSkcd4=","batch_id":"1","tradable":"90.003","retired":"9.997","escrowed":""}],"regen.ecocredit.v1.BatchInfo":[{"issuer":"WCBEyNFP/N5RoS4h43AqkjC6zA8=","project_id":"1","batch_denom":"BIO01-00000000-00000000-001","metadata":"batch metadata","start_date":null,"end_date":null,"issuance_date":"2022-04-08T10:40:10.774108141Z"},{"issuer":null,"project_id":"1","batch_denom":"BIO02-00000000-00000000-001","metadata":"batch metadata","start_date":null,"end_date":null,"issuance_date":"2022-04-08T10:40:10.774108556Z"}],"regen.ecocredit.v1.BatchSequence":[{"project_id":"P01","next_batch_id":"3"}],"regen.ecocredit.v1.BatchSupply":[{"batch_id":"1","tradable_amount":"90.003","retired_amount":"9.997","cancelled_amount":""}],"regen.ecocredit.v1.ClassInfo":[{"name":"BIO001","admin":"4A/V6LMEL2lZv9PZnkWSIDQzZM4=","metadata":"credit class metadata","credit_type":"BIO"},{"name":"BIO02","admin":"HK9YDsBMN1hU8tjfLTNy+qjbqLE=","metadata":"credit class metadata","credit_type":"BIO"}],"regen.ecocredit.v1.ClassIssuer":[{"class_id":"1","issuer":"1ygCfmJaPVMIvVEcpx6r+2gpurM="},{"class_id":"1","issuer":"KoXfzfqe+V/9x7C4XjnqDFB2Tl4="},{"class_id":"2","issuer":"KoXfzfqe+V/9x7C4XjnqDFB2Tl4="},{"class_id":"2","issuer":"lEjmu9Vooa24qp9vCMIlXGrMZoU="}],"regen.ecocredit.v1.ClassSequence":[{"credit_type":"BIO","next_class_id":"3"}],"regen.ecocredit.v1.Params":{"credit_class_fee":[{"denom":"stake","amount":"20000000"}],"basket_fee":[{"denom":"stake","amount":"20000000"}],"allowed_class_creators":[],"allowlist_enabled":false,"credit_types":[{"abbreviation":"C","name":"carbon","unit":"metric ton CO2 equivalent","precision":6}],"allowed_ask_denoms":[{"denom":"uregen","display_denom":"regen","exponent":18}]},"regen.ecocredit.v1.ProjectInfo":[{"name":"P01","admin":"gPFuHL7Hn+uVYD6XOR00du3C/Xg=","class_id":"1","project_location":"AQ","metadata":"project metadata"},{"name":"P02","admin":"CHkV2Tv6A7RXPJYTivVklbxXWP8=","class_id":"2","project_location":"AQ","metadata":"project metadata"}],"regen.ecocredit.v1.ProjectSequence":[{"class_id":"1","next_project_id":"3"}]}`

	jsonSource, err := ormjson.NewRawMessageSource([]byte(x))
	require.NoError(t, err)

	err = core.ValidateGenesis(jsonSource)
	require.NoError(t, err)
	require.True(t, false)
}

var (
	addr1 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	addr2 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
)

func TestGenesisValidate(t *testing.T) {
	cdc := simapp.MakeTestEncodingConfig().Marshaler
	ecocreditKey := sdk.NewKVStoreKey("ecocredit")
	db := dbm.NewMemDB()
	cms := store.NewCommitMultiStore(db)
	cms.MountStoreWithDB(ecocreditKey, sdk.StoreTypeIAVL, db)
	assert.NilError(t, cms.LoadLatestVersion())
	ormCtx := ormtable.WrapContextDefault(ormtest.NewMemoryBackend())
	sdkCtx := sdk.NewContext(cms, tmproto.Header{}, false, log.NewNopLogger()).WithContext(ormCtx)

	ormdb, err := ormdb.NewModuleDB(&ecocredit.ModuleSchema, ormdb.ModuleDBOptions{})
	require.NoError(t, err)
	ss, err := api.NewStateStore(ormdb)
	require.Nil(t, err)

	jsonTarget := ormjson.NewRawMessageTarget()
	require.NoError(t, ormdb.DefaultJSON(jsonTarget))

	params := core.DefaultParams()
	require.NoError(t, server.MergeParamsIntoTarget(cdc, &params, jsonTarget))

	defaultGenesisBz, err := jsonTarget.JSON()
	require.NoError(t, err)

	testCases := []struct {
		name        string
		gensisState func() []byte
		expectErr   bool
		errorMsg    string
	}{
		{
			"empty genesis state",
			func() []byte {
				return defaultGenesisBz
			},
			false,
			"",
		},
		{
			"valid: no credit batches",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:    "1",
						Admin:      addr1.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
				}
				return genesisState
			},
			false,
			"",
		},
		{
			"invalid: credit type param",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:    "1",
						Admin:      addr1.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
				}
				genesisState.Params.CreditTypes = []*ecocredit.CreditType{{
					Name:         "carbon",
					Abbreviation: "C",
					Unit:         "metric ton CO2 equivalent",
					Precision:    7,
				}}
				return genesisState
			},
			true,
			"invalid precision 7: precision is currently locked to 6: invalid request",
		},
		{
			"invalid: duplicate credit type",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:    "1",
						Admin:      addr1.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
				}
				genesisState.Params.CreditTypes = []*ecocredit.CreditType{{
					Name:         "carbon",
					Abbreviation: "C",
					Unit:         "metric ton CO2 equivalent",
					Precision:    6,
				}, {
					Name:         "carbon",
					Abbreviation: "C",
					Unit:         "metric ton CO2 equivalent",
					Precision:    6,
				}}
				return genesisState
			},
			true,
			"duplicate credit type name in request: carbon: invalid request",
		},
		{
			"invalid: bad addresses in allowlist",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:    "1",
						Admin:      addr1.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
				}
				genesisState.Params.AllowlistEnabled = true
				genesisState.Params.AllowedClassCreators = []string{"-=!?#09)("}
				return genesisState
			},
			true,
			"invalid creator address: decoding bech32 failed",
		},
		{
			"invalid: type name does not match param name",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:  "1",
						Admin:    addr1.String(),
						Issuers:  []string{addr1.String(), addr2.String()},
						Metadata: []byte("meta-data"),
						CreditType: &ecocredit.CreditType{
							Name:         "badbadnotgood",
							Abbreviation: "C",
							Unit:         "metric ton CO2 equivalent",
							Precision:    6,
						},
					},
				}
				return genesisState
			},
			true,
			formatCreditTypeParamError(ecocredit.CreditType{"badbadnotgood", "C", "metric ton CO2 equivalent", 6}).Error(),
		},
		{
			"invalid: type unit does not match param unit",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:  "1",
						Admin:    addr1.String(),
						Issuers:  []string{addr1.String(), addr2.String()},
						Metadata: []byte("meta-data"),
						CreditType: &ecocredit.CreditType{
							Name:         "carbon",
							Abbreviation: "C",
							Unit:         "inches",
							Precision:    6,
						},
					},
				}
				return genesisState
			},
			true,
			formatCreditTypeParamError(ecocredit.CreditType{"carbon", "C", "inches", 6}).Error(),
		},
		{
			"invalid: non-existent abbreviation",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:  "1",
						Admin:    addr1.String(),
						Issuers:  []string{addr1.String(), addr2.String()},
						Metadata: []byte("meta-data"),
						CreditType: &ecocredit.CreditType{
							Name:         "carbon",
							Abbreviation: "F",
							Unit:         "metric ton CO2 equivalent",
							Precision:    6,
						},
					},
				}
				return genesisState
			},
			true,
			"unknown credit type abbreviation: F: not found",
		},
		{
			"expect error: supply is missing",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()

				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:    "1",
						Admin:      addr1.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
				}
				genesisState.ProjectInfo = []*ecocredit.ProjectInfo{
					{
						ProjectId:       "01",
						ClassId:         "1",
						Issuer:          addr1.String(),
						Metadata:        []byte("meta-data"),
						ProjectLocation: "AQ",
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ProjectId:   "01",
						BatchDenom:  "1/2",
						TotalAmount: "1000",
						Metadata:    []byte("meta-data"),
					},
				}
				genesisState.Balances = []*ecocredit.Balance{
					{
						Address:         addr2.String(),
						BatchDenom:      "1/2",
						TradableBalance: "400.456",
					},
				}
				return genesisState
			},
			true,
			"supply is not found for 1/2 credit batch: not found",
		},
		{
			"expect error: invalid supply",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:    "1",
						Admin:      addr1.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
				}
				genesisState.ProjectInfo = []*ecocredit.ProjectInfo{
					{
						ProjectId:       "01",
						ClassId:         "1",
						Issuer:          addr1.String(),
						Metadata:        []byte("meta-data"),
						ProjectLocation: "AQ",
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ProjectId:   "01",
						BatchDenom:  "1/2",
						TotalAmount: "1000",
						Metadata:    []byte("meta-data"),
					},
				}
				genesisState.Balances = []*ecocredit.Balance{
					{
						Address:         addr2.String(),
						BatchDenom:      "1/2",
						TradableBalance: "100",
						RetiredBalance:  "100",
					},
				}
				genesisState.Supplies = []*ecocredit.Supply{
					{
						BatchDenom:     "1/2",
						TradableSupply: "10",
					},
				}
				return genesisState
			},
			true,
			"supply is incorrect for 1/2 credit batch, expected 10, got 200: invalid coins",
		},
		{
			"valid test case",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:    "1",
						Admin:      addr1.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
				}
				genesisState.ProjectInfo = []*ecocredit.ProjectInfo{
					{
						ProjectId:       "01",
						ClassId:         "1",
						Issuer:          addr1.String(),
						Metadata:        []byte("meta-data"),
						ProjectLocation: "AQ",
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ProjectId:   "01",
						BatchDenom:  "1/2",
						TotalAmount: "1000",
						Metadata:    []byte("meta-data"),
					},
				}
				genesisState.Balances = []*ecocredit.Balance{
					{
						Address:         addr2.String(),
						BatchDenom:      "1/2",
						TradableBalance: "100.123",
						RetiredBalance:  "100.123",
					},
					{
						Address:         addr1.String(),
						BatchDenom:      "1/2",
						TradableBalance: "100.123",
						RetiredBalance:  "100.123",
					},
				}
				genesisState.Supplies = []*ecocredit.Supply{
					{
						BatchDenom:     "1/2",
						TradableSupply: "200.246",
						RetiredSupply:  "200.246",
					},
				}
				return genesisState
			},
			false,
			"",
		},
		{
			"valid test case, multiple classes",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:    "1",
						Admin:      addr1.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
					{
						ClassId:    "2",
						Admin:      addr2.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
				}
				genesisState.ProjectInfo = []*ecocredit.ProjectInfo{
					{
						ProjectId:       "01",
						ClassId:         "1",
						Issuer:          addr1.String(),
						Metadata:        []byte("meta-data"),
						ProjectLocation: "AQ",
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ProjectId:   "01",
						BatchDenom:  "1/2",
						TotalAmount: "1000",
						Metadata:    []byte("meta-data"),
					},
					{
						ProjectId:       "01",
						BatchDenom:      "2/2",
						AmountCancelled: "0",
						TotalAmount:     "1000",
						Metadata:        []byte("meta-data"),
					},
				}
				genesisState.Balances = []*ecocredit.Balance{
					{
						Address:         addr2.String(),
						BatchDenom:      "1/2",
						TradableBalance: "100.123",
						RetiredBalance:  "100.123",
					},
					{
						Address:         addr1.String(),
						BatchDenom:      "2/2",
						TradableBalance: "100.123",
						RetiredBalance:  "100.123",
					},
				}
				genesisState.Supplies = []*ecocredit.Supply{
					{
						BatchDenom:     "1/2",
						TradableSupply: "100.123",
						RetiredSupply:  "100.123",
					},
					{
						BatchDenom:     "2/2",
						TradableSupply: "100.123",
						RetiredSupply:  "100.123",
					},
				}
				return genesisState
			},
			false,
			"",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := core.ValidateGenesis(tc.gensisState())
			if tc.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errorMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

var defaultCreditTypes = core.DefaultParams().CreditTypes

func formatCreditTypeParamError(ct ecocredit.CreditType) error {
	return fmt.Errorf("credit type %+v does not match param type %+v: invalid type", ct, *defaultCreditTypes[0])
}
