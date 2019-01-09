package cli

import (
	"encoding/hex"
	"github.com/spf13/cobra"
	agentcli "gitlab.com/regen-network/regen-ledger/x/agent/client/cli"
	"gitlab.com/regen-network/regen-ledger/x/esp"
	"gitlab.com/regen-network/regen-ledger/x/proposal"
	proposalcli "gitlab.com/regen-network/regen-ledger/x/proposal/client/cli"
	"strconv"

	//"github.com/twpayne/go-geom/encoding/ewkbhex"
	//"github.com/twpayne/go-geom/encoding/ewkb"

	"github.com/cosmos/cosmos-sdk/codec"
)

func GetCmdProposeVersion(cdc *codec.Codec) *cobra.Command {
	var verifiers []string

	cmd := proposalcli.GetCmdPropose(cdc, func(cmd *cobra.Command, args []string) (action proposal.ProposalAction, e error) {
		curatorStr := args[0]
		curator, err := strconv.ParseUint(curatorStr, 10, 64)
		if err != nil {
			return nil, err
		}

		name := args[1]

		version := args[2]

		verifierAgents := agentcli.AgentsFromArray(verifiers)

		return esp.ActionRegisterESPVersion{
			Curator: curator,
			Name:    name,
			Version: version,
			Spec:    esp.ESPVersionSpec{Verifiers: verifierAgents},
		}, nil
	})

	cmd.Args = cobra.ExactArgs(3)
	cmd.Use = "propose-version <curator> <name> <version> --verifiers <verifiers-list>"
	cmd.Short = "Propose an ESP version"
	cmd.Flags().StringArrayVar(&verifiers, "verifiers", []string{}, "ESP verifier agent ID's")
	return cmd
}

func GetCmdReportResult(cdc *codec.Codec) *cobra.Command {
	cmd := proposalcli.GetCmdPropose(cdc, func(cmd *cobra.Command, args []string) (action proposal.ProposalAction, e error) {
		curatorStr := args[0]
		curator, err := strconv.ParseUint(curatorStr, 10, 64)
		if err != nil {
			return nil, err
		}

		name := args[1]

		version := args[2]

		verifierStr := args[3]
		verifier, err := strconv.ParseUint(verifierStr, 10, 64)
		if err != nil {
			return nil, err
		}

		polygonHex := args[4]
		polygon, err := hex.DecodeString(polygonHex)
		if err != nil {
			return nil, err
		}

		data := args[5]

		return esp.ActionReportESPResult{
			Curator:  curator,
			Name:     name,
			Version:  version,
			Verifier: verifier,
			Result: esp.ESPResult{
				Data:        []byte(data),
				PolygonEWKB: polygon,
			},
		}, nil
	})

	cmd.Args = cobra.ExactArgs(6)
	cmd.Use = "propose-result <curator> <name> <version> <verifier> <polygon-ewkb-hex> <data>"
	cmd.Short = "Propose an ESP result"
	return cmd
}
