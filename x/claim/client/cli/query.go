package cli

import (
	"bytes"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/regen-network/regen-ledger/types"
	"github.com/regen-network/regen-ledger/x/claim"
	"github.com/spf13/cobra"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// GetSignaturesQueryCmd creates a query sub-command for the claim module using cmdName as the name of the sub-command.
func GetSignaturesQueryCmd(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "signatures <content-address>",
		Short: "get signatures for claim",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			content, err := types.DecodeBech32DataAddress(args[0])
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryStore(claim.KeySignatures(content), storeName)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return fmt.Errorf("no signatures for claim")
			}

			var sigs []sdk.AccAddress
			err = cdc.UnmarshalBinaryBare(res, &sigs)
			if err != nil {
				return err
			}

			var signature_str bytes.Buffer

			for _, sig := range sigs {
				signature_str.WriteString(sig) // need to convert type []AccAddress to string
			}
			return cliCtx.PrintOutput(signature_str.String())
		},
	}
}

// GetEvidenceQueryCmd creates a query sub-command for the claim module using cmdName as the name of the sub-command.
func GetEvidenceQueryCmd(storeName string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "evidence <evidence-address> <signer-address>",
		Short: "get evidence for claim",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			content, err := types.DecodeBech32DataAddress(args[0])
			if err != nil {
				return err
			}
			signer, err := sdk.AccAddressFromBech32(args[1])
			if err != nil {
				return err
			}

			res, err := cliCtx.QueryStore(claim.KeySignatureEvidence(content, signer), storeName)
			if err != nil {
				return err
			}

			if len(res) == 0 {
				return fmt.Errorf("no evidence for claim")
			}

			var evidence []types.DataAddress
			err = cdc.UnmarshalBinaryBare(res, &evidence)
			if err != nil {
				return err
			}
			return cliCtx.PrintOutput(evidence)
		},
	}
}

