package client

import (
	"fmt"
	"regexp"
	"strings"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/gogo/protobuf/proto"
	"github.com/spf13/cobra"

	"github.com/regen-network/regen-ledger/x/ecocredit"
)

// prints a query client response
func print(cctx sdkclient.Context, res proto.Message, err error) error {
	if err != nil {
		return err
	}
	return cctx.PrintProto(res)
}

func mkQueryClient(cmd *cobra.Command) (ecocredit.QueryClient, sdkclient.Context, error) {
	ctx, err := sdkclient.GetClientQueryContext(cmd)
	if err != nil {
		return nil, sdkclient.Context{}, err
	}
	return ecocredit.NewQueryClient(ctx), ctx, err
}

type credits struct {
	batchDenom string
	amount     string
}

var (
	reCreditAmt  = `[[:digit:]]+(?:\.[[:digit:]]+)?|\.[[:digit:]]+`
	reBatchDenom = `[a-zA-Z0-9]+\/[a-zA-Z0-9]+`
	reCredits    = regexp.MustCompile(fmt.Sprintf(`^(%s)\:(%s)$`, reCreditAmt, reBatchDenom))
)

func parseCancelCreditsList(creditsListStr string) ([]*ecocredit.MsgCancelRequest_CancelCredits, error) {
	creditsList, err := parseCreditsList(creditsListStr)
	if err != nil {
		return nil, err
	}

	cancelCreditsList := make([]*ecocredit.MsgCancelRequest_CancelCredits, len(creditsList))
	for i, credits := range creditsList {
		cancelCreditsList[i] = &ecocredit.MsgCancelRequest_CancelCredits {
			BatchDenom: credits.batchDenom,
			Amount:     credits.amount,
		}
	}

	return cancelCreditsList, nil
}

func parseCreditsList(creditsListStr string) ([]credits, error) {
	creditsListStr = strings.TrimSpace(creditsListStr)
	if len(creditsListStr) == 0 {
		return nil, nil
	}

	creditsStrs := strings.Split(creditsListStr, ",")
	creditsList := make([]credits, len(creditsStrs))
	for i, creditsStr := range creditsStrs {
		credits, err := parseCredits(creditsStr)
		if err != nil {
			return nil, err
		}

		creditsList[i] = credits
	}

	return creditsList, nil
}

func parseCredits(creditsStr string) (credits, error) {
	creditsStr = strings.TrimSpace(creditsStr)

	matches := reCredits.FindStringSubmatch(creditsStr)
	if matches == nil {
		return credits{}, fmt.Errorf("invalid credit expression: %s", creditsStr)
	}

	return credits {
		batchDenom: matches[2],
		amount:     matches[1],
	}, nil
}
