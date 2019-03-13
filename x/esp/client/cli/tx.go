package cli

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/esp"
	"github.com/regen-network/regen-ledger/x/geo"
	"github.com/regen-network/regen-ledger/x/proposal"
	proposalcli "github.com/regen-network/regen-ledger/x/proposal/client/cli"
	"github.com/spf13/cobra"

	//"github.com/twpayne/go-geom/encoding/ewkbhex"
	//"github.com/twpayne/go-geom/encoding/ewkb"

	"github.com/cosmos/cosmos-sdk/codec"
)

func addrsFromBech32Array(arr []string) []sdk.AccAddress {
	n := len(arr)
	res := make([]sdk.AccAddress, n)
	for i := 0; i < n; i++ {
		str := arr[i]
		acc, err := sdk.AccAddressFromBech32(str)
		if err != nil {
			panic(err)
		}
		res[i] = acc
	}
	return res
}

func GetCmdProposeVersion(cdc *codec.Codec) *cobra.Command {
	var verifiers []string

	cmd := proposalcli.GetCmdPropose(cdc, func(cmd *cobra.Command, args []string) (action proposal.ProposalAction, e error) {
		curator, e := sdk.AccAddressFromBech32(args[0])
		if e != nil {
			return action, e
		}

		name := args[1]

		version := args[2]

		verifierAgents := addrsFromBech32Array(verifiers)

		return esp.ActionRegisterESPVersion{
			ESPVersionSpec: esp.ESPVersionSpec{
				Curator:   curator,
				Name:      name,
				Version:   version,
				Verifiers: verifierAgents,
			},
		}, nil
	})

	cmd.Args = cobra.ExactArgs(3)
	cmd.Use = "propose-version <curator> <name> <version> --verifiers <verifiers-list>"
	cmd.Short = "Propose an ESP version"
	cmd.Flags().StringArrayVar(&verifiers, "verifiers", []string{}, "ESP verifier group ID's")
	return cmd
}

func GetCmdReportResult(cdc *codec.Codec) *cobra.Command {
	cmd := proposalcli.GetCmdPropose(cdc, func(cmd *cobra.Command, args []string) (action proposal.ProposalAction, e error) {
		curator, e := sdk.AccAddressFromBech32(args[0])
		if e != nil {
			return action, e
		}

		name := args[1]

		version := args[2]

		verifier, e := sdk.AccAddressFromBech32(args[3])
		if e != nil {
			return action, e
		}

		geoId := geo.MustDecodeBech32GeoID(args[4])

		data := args[5]

		return esp.ActionReportESPResult{
			ESPResult: esp.ESPResult{
				Curator:  curator,
				Name:     name,
				Version:  version,
				Verifier: verifier,
				Data:     []byte(data),
				GeoID:    geoId,
			},
		}, nil
	})

	cmd.Args = cobra.ExactArgs(6)
	cmd.Use = "propose-result <curator> <name> <version> <verifier> <geo-id> <data>"
	cmd.Short = "Propose an ESP result"
	return cmd
}
