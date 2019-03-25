package cli

import (
	"encoding/json"
	"fmt"
	"github.com/regen-network/regen-ledger/x/consortium"
	"github.com/regen-network/regen-ledger/x/proposal"
	proposalcli "github.com/regen-network/regen-ledger/x/proposal/client/cli"
	"github.com/regen-network/regen-ledger/x/upgrade"
	"github.com/spf13/cobra"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
)

func GetCmdProposeUpgrade(cdc *codec.Codec) *cobra.Command {
	var timeStr string
	var height int64
	var commit string

	cmd := proposalcli.GetCmdPropose(cdc, func(cmd *cobra.Command, args []string) (action proposal.ProposalAction, e error) {
		name := args[0]

		var t time.Time
		var err error
		if len(timeStr) != 0 {
			t, err = time.Parse(time.RFC3339, timeStr)
			if err != nil {
				return nil, fmt.Errorf("error parsing time: %+v", err)
			}

			if height != 0 {
				return nil, fmt.Errorf("only one of --time or --height should be specified")
			}
		}

		info := make(map[string]interface{})
		if len(commit) != 0 {
			info["commit"] = commit
		}

		jsonInfo, err := json.Marshal(info)

		return consortium.ActionScheduleUpgrade{
			Plan: upgrade.Plan{
				Name:   name,
				Time:   t,
				Height: height,
				Info:   string(jsonInfo),
			},
		}, nil
	})

	cmd.Args = cobra.ExactArgs(1)
	cmd.Use = "propose-upgrade <name> [--time <time> | --height <height>] [--commit <commit-hash>]"
	cmd.Short = "Propose an ESP version"
	cmd.Flags().StringVar(&timeStr, "time", "", "The time after which the upgrade must happen in ISO8601/RFC3339 format (omit if using --upgrade-height)")
	cmd.Flags().Int64Var(&height, "height", 0, "The height at which the upgrade must happen (omit if using --upgrade-time)")
	cmd.Flags().StringVar(&commit, "commit", "", "The git commit hash of the version of the software to upgrade to")
	return cmd
}
