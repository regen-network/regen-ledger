package ecocredit_test

import (
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
						ClassId:  "1",
						Designer: addr1.String(),
						Issuers:  []string{addr1.String(), addr2.String()},
						Metadata: []byte("meta-data"),
						CreditType: &ecocredit.CreditType{
							Name:         "carbon",
							Abbreviation: "C",
							Unit:         "ton",
							Precision:    6,
						},
					},
				}
				return genesisState
			},
			false,
			"",
		},
		{
			"expect error: supply is missing",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()

				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:  "1",
						Designer: addr1.String(),
						Issuers:  []string{addr1.String(), addr2.String()},
						Metadata: []byte("meta-data"),
						CreditType: &ecocredit.CreditType{
							Name:         "carbon",
							Abbreviation: "C",
							Unit:         "ton",
							Precision:    6,
						},
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ClassId:     "1",
						BatchDenom:  "1/2",
						Issuer:      addr1.String(),
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
						Designer:   addr1.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ClassId:     "1",
						BatchDenom:  "1/2",
						Issuer:      addr1.String(),
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
						Designer:   addr1.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ClassId:     "1",
						BatchDenom:  "1/2",
						Issuer:      addr1.String(),
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
						Designer:   addr1.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
					{
						ClassId:    "2",
						Designer:   addr2.String(),
						Issuers:    []string{addr1.String(), addr2.String()},
						Metadata:   []byte("meta-data"),
						CreditType: genesisState.Params.CreditTypes[0],
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ClassId:     "1",
						BatchDenom:  "1/2",
						Issuer:      addr1.String(),
						TotalAmount: "1000",
						Metadata:    []byte("meta-data"),
					},
					{
						ClassId:         "2",
						BatchDenom:      "2/2",
						Issuer:          addr2.String(),
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
				require.Equal(t, tc.errorMsg, err.Error())
			} else {
				require.NoError(t, err)
			}
		})
	}
}
