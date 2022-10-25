package client

import (
	"fmt"
	"io/ioutil"

	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/x/intertx"
	v1 "github.com/regen-network/regen-ledger/x/intertx/types/v1"
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

			msg := v1.MsgRegisterAccount{
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
		Use:     "submit [path/to/sdk_msg.json]",
		Example: "regen tx intertx submit tx.json",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			cdc := codec.NewProtoCodec(clientCtx.InterfaceRegistry)

			var txMsg sdk.Msg
			if err := cdc.UnmarshalInterfaceJSON([]byte(args[0]), &txMsg); err != nil {

				// check for file path if JSON input is not provided
				contents, err := ioutil.ReadFile(args[0])
				if err != nil {
					return sdkerrors.ErrNotFound.Wrapf("could not real file at %s: %s", args[0], err.Error())
				}

				if err := cdc.UnmarshalInterfaceJSON(contents, &txMsg); err != nil {
					return sdkerrors.ErrJSONUnmarshal.Wrap(err.Error())
				}
			}

			anyBz, err := packAnyIntoMsg(txMsg)
			if err != nil {
				return err
			}

			msg := v1.MsgSubmitTx{
				Owner:        clientCtx.GetFromAddress().String(),
				ConnectionId: viper.GetString(FlagConnectionID),
				Msg:          anyBz,
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), &msg)
		},
	}

	cmd.Flags().String(FlagConnectionID, "", "the connection end identifier on the controller chain")
	_ = cmd.MarkFlagRequired(FlagConnectionID)

	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

// PackTxMsgAny marshals the sdk.Msg payload to a protobuf Any type
func packAnyIntoMsg(sdkMsg sdk.Msg) (*codectypes.Any, error) {
	msg, ok := sdkMsg.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("can't proto marshal %T", sdkMsg)
	}

	a, err := codectypes.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}

	return a, nil
}
