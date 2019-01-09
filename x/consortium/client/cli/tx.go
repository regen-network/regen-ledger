package cli

import (
	"github.com/spf13/cobra"
	"gitlab.com/regen-network/regen-ledger/x/consortium"
	"gitlab.com/regen-network/regen-ledger/x/proposal"
	proposalcli "gitlab.com/regen-network/regen-ledger/x/proposal/client/cli"
	"gitlab.com/regen-network/regen-ledger/x/upgrade"
	"strconv"

	"github.com/cosmos/cosmos-sdk/codec"
)

func GetCmdProposeUpgrade(cdc *codec.Codec) *cobra.Command {
	var memo string

	cmd := proposalcli.GetCmdPropose(cdc, func(cmd *cobra.Command, args []string) (action proposal.ProposalAction, e error) {
		blockHeightStr := args[0]
		blockHeight, err := strconv.ParseInt(blockHeightStr, 10, 64)
		if err != nil {
			return nil, err
		}

		return consortium.ActionScheduleUpgrade{
			UpgradeInfo: upgrade.UpgradeInfo{
				Height: blockHeight,
				Memo:   memo,
			},
		}, nil
	})

	cmd.Args = cobra.ExactArgs(1)
	cmd.Use = "propose-upgrade <block-height> [--upgrade-memo <memo>]"
	cmd.Short = "Propose an ESP version"
	cmd.Flags().StringVar(&memo, "upgrade-memo", "", "Any memo to attach to the upgrade plan such as a git commit hash")
	return cmd
}
