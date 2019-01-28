package cli

import (
	"encoding/binary"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"
	"github.com/spf13/cobra"
	"github.com/twpayne/go-geom"
	"github.com/twpayne/go-geom/encoding/ewkb"
	"github.com/twpayne/go-geom/encoding/geojson"
	utils2 "gitlab.com/regen-network/regen-ledger/utils"
	"gitlab.com/regen-network/regen-ledger/x/geo"
)

func GetCmdStoreGeometry(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		// TODO WKT, EWKT and EWKB little-endian
		Use:   "store <geometry> (expects geojson format by default)",
		Short: "store a geometry on the blockchain",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}

			account, err := cliCtx.GetFromAddress()
			if err != nil {
				return err
			}

			geojsonStr := args[0]

			var geometry geom.T

			err = geojson.Unmarshal([]byte(geojsonStr), &geometry)
			if err != nil {
				return err
			}

			bz, err := ewkb.Marshal(geometry, binary.LittleEndian)

			if err != nil {
				return err
			}

			msg := geo.MsgStoreGeometry{
				Data: geo.Geometry{
					EWKB: bz,
				},
				Signer: account,
			}
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true
			cliCtx.ResponsePrinter = utils2.PrintCLIResponse_StringData

			return utils.CompleteAndBroadcastTxCli(txBldr, cliCtx, []sdk.Msg{msg})
		},
	}
	return cmd
}
