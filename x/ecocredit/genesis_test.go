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
			"expect error: no supply for tradable balance",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.Precisions = []*ecocredit.Precision{
					{
						BatchDenom:       "1/2",
						MaxDecimalPlaces: 3,
					},
				}
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:  "1",
						Designer: addr1.String(),
						Issuers:  []string{addr1.String(), addr2.String()},
						Metadata: []byte("meta-data"),
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ClassId:    "1",
						BatchDenom: "1/2",
						Issuer:     addr1.String(),
						TotalUnits: "1000",
						Metadata:   []byte("meta-data"),
					},
				}
				genesisState.TradableBalances = []*ecocredit.Balance{
					{
						Address:    addr2.String(),
						BatchDenom: "1/2",
						Balance:    "400.456",
					},
				}
				return genesisState
			},
			true,
			"tradable: supply is not found for 1/2 credit batch: not found",
		},
		{
			"expect error: no supply for retired balance",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.Precisions = []*ecocredit.Precision{
					{
						BatchDenom:       "1/2",
						MaxDecimalPlaces: 3,
					},
				}
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:  "1",
						Designer: addr1.String(),
						Issuers:  []string{addr1.String(), addr2.String()},
						Metadata: []byte("meta-data"),
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ClassId:    "1",
						BatchDenom: "1/2",
						Issuer:     addr1.String(),
						TotalUnits: "1000",
						Metadata:   []byte("meta-data"),
					},
				}
				genesisState.TradableBalances = []*ecocredit.Balance{
					{
						Address:    addr2.String(),
						BatchDenom: "1/2",
						Balance:    "400.456",
					},
					{
						Address:    addr1.String(),
						BatchDenom: "1/2",
						Balance:    "400.111",
					},
				}
				genesisState.TradableSupplies = []*ecocredit.Supply{
					{
						BatchDenom: "1/2",
						Supply:     "800.567",
					},
				}
				genesisState.RetiredBalances = []*ecocredit.Balance{
					{
						Address:    addr2.String(),
						BatchDenom: "1/2",
						Balance:    "100.123",
					},
				}
				return genesisState
			},
			true,
			"retired: supply is not found for 1/2 credit batch: not found",
		},
		{
			"expect error: invalid tradable supply",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:  "1",
						Designer: addr1.String(),
						Issuers:  []string{addr1.String(), addr2.String()},
						Metadata: []byte("meta-data"),
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ClassId:    "1",
						BatchDenom: "1/2",
						Issuer:     addr1.String(),
						TotalUnits: "1000",
						Metadata:   []byte("meta-data"),
					},
				}
				genesisState.TradableBalances = []*ecocredit.Balance{
					{
						Address:    addr2.String(),
						BatchDenom: "1/2",
						Balance:    "100",
					},
				}
				genesisState.TradableSupplies = []*ecocredit.Supply{
					{
						BatchDenom: "1/2",
						Supply:     "10",
					},
				}
				genesisState.RetiredBalances = []*ecocredit.Balance{
					{
						Address:    addr2.String(),
						BatchDenom: "1/2",
						Balance:    "100",
					},
				}
				return genesisState
			},
			true,
			"tradable: supply is incorrect for 1/2 credit batch, expected 10, got 100: invalid coins",
		},
		{
			"expect error: invalid retired supply",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:  "1",
						Designer: addr1.String(),
						Issuers:  []string{addr1.String(), addr2.String()},
						Metadata: []byte("meta-data"),
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ClassId:    "1",
						BatchDenom: "1/2",
						Issuer:     addr1.String(),
						TotalUnits: "1000",
						Metadata:   []byte("meta-data"),
					},
				}
				genesisState.TradableBalances = []*ecocredit.Balance{
					{
						Address:    addr2.String(),
						BatchDenom: "1/2",
						Balance:    "100",
					},
					{
						Address:    addr1.String(),
						BatchDenom: "1/2",
						Balance:    "100",
					},
				}
				genesisState.TradableSupplies = []*ecocredit.Supply{
					{
						BatchDenom: "1/2",
						Supply:     "200",
					},
				}
				genesisState.RetiredBalances = []*ecocredit.Balance{
					{
						Address:    addr2.String(),
						BatchDenom: "1/2",
						Balance:    "100",
					},
				}
				genesisState.RetiredSupplies = []*ecocredit.Supply{
					{
						BatchDenom: "1/2",
						Supply:     "200",
					},
				}
				return genesisState
			},
			true,
			"retired: supply is incorrect for 1/2 credit batch, expected 200, got 100: invalid coins",
		},
		{
			"valid test case",
			func() *ecocredit.GenesisState {
				genesisState := ecocredit.DefaultGenesisState()
				genesisState.Precisions = []*ecocredit.Precision{
					{
						BatchDenom:       "1/2",
						MaxDecimalPlaces: 3,
					},
				}
				genesisState.ClassInfo = []*ecocredit.ClassInfo{
					{
						ClassId:  "1",
						Designer: addr1.String(),
						Issuers:  []string{addr1.String(), addr2.String()},
						Metadata: []byte("meta-data"),
					},
				}
				genesisState.BatchInfo = []*ecocredit.BatchInfo{
					{
						ClassId:    "1",
						BatchDenom: "1/2",
						Issuer:     addr1.String(),
						TotalUnits: "1000",
						Metadata:   []byte("meta-data"),
					},
				}
				genesisState.TradableBalances = []*ecocredit.Balance{
					{
						Address:    addr2.String(),
						BatchDenom: "1/2",
						Balance:    "100.123",
					},
					{
						Address:    addr1.String(),
						BatchDenom: "1/2",
						Balance:    "100.123",
					},
				}
				genesisState.TradableSupplies = []*ecocredit.Supply{
					{
						BatchDenom: "1/2",
						Supply:     "200.246",
					},
				}
				genesisState.RetiredBalances = []*ecocredit.Balance{
					{
						Address:    addr2.String(),
						BatchDenom: "1/2",
						Balance:    "100.123",
					},
					{
						Address:    addr1.String(),
						BatchDenom: "1/2",
						Balance:    "100.123",
					},
				}
				genesisState.RetiredSupplies = []*ecocredit.Supply{
					{
						BatchDenom: "1/2",
						Supply:     "200.246",
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
