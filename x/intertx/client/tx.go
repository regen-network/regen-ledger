package client

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"

	"github.com/regen-network/regen-ledger/x/intertx"
	intertxv1 "github.com/regen-network/regen-ledger/x/intertx/types/v1"
)

const (
	// FlagConnectionID is the connection end identifier on the controller chain
	FlagConnectionID = "connection-id"
	// FlagVersion is the controller chain channel version
	FlagVersion = "version"
)

// GetTxCmd creates and returns the intertx tx command
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        intertx.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", intertx.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		getRegisterAccountCmd(),
		getSubmitTxCmd(),
	)

	return cmd
}

func getRegisterAccountCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "register --connection-id <connection-id> --version <version>",
		Short:   "register an interchain account",
		Long:    "register an interchain account for the chain corresponding to the connection-id.",
		Example: "regen tx intertx register --connection-id channel-10 --version v5",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			msg := intertxv1.MsgRegisterAccount{
				Owner:        clientCtx.GetFromAddress().String(),
				ConnectionId: viper.GetString(FlagConnectionID),
				Version:      viper.GetString(FlagVersion),
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String(FlagConnectionID, "", "the connection end identifier on the controller chain")
	cmd.Flags().String(FlagVersion, "", "")
	_ = cmd.MarkFlagRequired(FlagConnectionID)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func getSubmitTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "submit-tx [connection-id] [path/to/sdk_msg.json]",
		Example: "regen tx intertx submit-tx channel-5 tx.json",
		Args:    cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			connectionID := args[0]
			sdkMsgs := args[1]

			theTx, err := authclient.ReadTxFromFile(clientCtx, sdkMsgs)
			if err != nil {
				return err
			}

			innerMsgs := theTx.GetMsgs()
			if lenMsgs := len(innerMsgs); lenMsgs != 1 {
				return sdkerrors.ErrInvalidRequest.Wrapf("expected 1 msg, got %d", lenMsgs)
			}

			msg := intertxv1.NewMsgSubmitTx(clientCtx.GetFromAddress().String(), connectionID, innerMsgs[0])

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
