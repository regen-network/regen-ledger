package cli

import (
	"encoding/hex"
	"github.com/spf13/cobra"
	"gitlab.com/regen-network/regen-ledger/x/agent"
	agentcli "gitlab.com/regen-network/regen-ledger/x/agent/client/cli"
	"gitlab.com/regen-network/regen-ledger/x/esp"
	"gitlab.com/regen-network/regen-ledger/x/proposal"
	proposalcli "gitlab.com/regen-network/regen-ledger/x/proposal/client/cli"

	//"github.com/twpayne/go-geom/encoding/ewkbhex"
	//"github.com/twpayne/go-geom/encoding/ewkb"

	"github.com/cosmos/cosmos-sdk/codec"
)

func GetCmdProposeVersion(cdc *codec.Codec) *cobra.Command {
	var verifiers []string

	cmd := proposalcli.GetCmdPropose(cdc, func(cmd *cobra.Command, args []string) (action proposal.ProposalAction, e error) {
		curator := agent.MustDecodeBech32AgentID(args[0])

		name := args[1]

		version := args[2]

		verifierAgents := agentcli.AgentsFromArray(verifiers)

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
	cmd.Flags().StringArrayVar(&verifiers, "verifiers", []string{}, "ESP verifier agent ID's")
	return cmd
}

func GetCmdReportResult(cdc *codec.Codec) *cobra.Command {
	cmd := proposalcli.GetCmdPropose(cdc, func(cmd *cobra.Command, args []string) (action proposal.ProposalAction, e error) {
		curator := agent.MustDecodeBech32AgentID(args[0])

		name := args[1]

		version := args[2]

		verifier := agent.MustDecodeBech32AgentID(args[3])

		polygonHex := args[4]
		polygon, err := hex.DecodeString(polygonHex)
		if err != nil {
			return nil, err
		}

		data := args[5]

		return esp.ActionReportESPResult{
			ESPResult: esp.ESPResult{
				Curator:     curator,
				Name:        name,
				Version:     version,
				Verifier:    verifier,
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
