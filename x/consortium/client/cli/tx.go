package cli

import (
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
	var memo string

	cmd := proposalcli.GetCmdPropose(cdc, func(cmd *cobra.Command, args []string) (action proposal.ProposalAction, e error) {
		name := args[0]

		var t time.Time
		var err error
		if len(timeStr) != 0 {
			t, err = time.Parse(time.RFC3339, timeStr)
			if err != nil {
				panic(err)
			}
		}

		return consortium.ActionScheduleUpgrade{
			Plan: upgrade.UpgradePlan{
				Name:   name,
				Time:   t,
				Height: height,
				Memo:   memo,
			},
		}, nil
	})

	cmd.Args = cobra.ExactArgs(1)
	cmd.Use = "propose-upgrade <name> [--upgrade-time <time> | --upgrade-height <height>] [--upgrade-memo <memo>]"
	cmd.Short = "Propose an ESP version"
	cmd.Flags().StringVar(&timeStr, "upgrade-time", "", "The time after which the upgrade must happen in ISO8601/RFC3339 format (omit if using --upgrade-height)")
	cmd.Flags().Int64Var(&height, "upgrade-height", 0, "The height at which the upgrade must happen (omit if using --upgrade-time)")
	cmd.Flags().StringVar(&memo, "upgrade-memo", "", "Any memo to attach to the upgrade plan such as a git commit hash")
	return cmd
}
