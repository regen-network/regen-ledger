package client

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

func parseCancelCreditsList(creditsListStr string) ([]*ecocredit.MsgCancel_CancelCredits, error) {
	creditsList, err := parseCreditsList(creditsListStr)
	if err != nil {
		return nil, err
	}

	cancelCreditsList := make([]*ecocredit.MsgCancel_CancelCredits, len(creditsList))
	for i, credits := range creditsList {
		cancelCreditsList[i] = &ecocredit.MsgCancel_CancelCredits{
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

	return credits{
		batchDenom: matches[2],
		amount:     matches[1],
	}, nil
}

func parseDate(field string, date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, sdkerrors.ErrInvalidRequest.Wrapf("%s must have format yyyy-mm-dd, but received %v", field, date)
	}
	return t, nil
}
