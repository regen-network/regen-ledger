package ecocredit_test

import (
	"fmt"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/ecocredit"
	"github.com/stretchr/testify/require"
	"github.com/tendermint/tendermint/crypto/ed25519"
)

var (
	addr1 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
	addr2 = sdk.AccAddress(ed25519.GenPrivKey().PubKey().Address())
)

func TestGenesisDefaultParams(t *testing.T) {
	genesis := ecocredit.DefaultGenesisState()
	params := ecocredit.DefaultParams()
	require.Equal(t, params.String(), genesis.Params.String())
}

func TestGenesisValidate(t *testing.T) {
	testCases := []struct {
		name        string
		gensisState func() *ecocredit.GenesisState
		expectErr   bool
		errorMsg    string
	}{
		{
			"empty genesis state",
			func() *ecocredit.GenesisState {
				return ecocredit.DefaultGenesisState()
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
			err := tc.gensisState().Validate()
			if tc.expectErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.errorMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

var defaultCreditTypes = ecocredit.DefaultGenesisState().Params.CreditTypes

func formatCreditTypeParamError(ct ecocredit.CreditType) error {
	return fmt.Errorf("credit type %+v does not match param type %+v: invalid type", ct, *defaultCreditTypes[0])
}
