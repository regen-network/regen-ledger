package client

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strconv"
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

func parseMsgCreateBatch(clientCtx sdkclient.Context, batchFile string) (*ecocredit.MsgCreateBatch, error) {
	contents, err := ioutil.ReadFile(batchFile)
	if err != nil {
		return nil, err
	}

	var msg ecocredit.MsgCreateBatch
	err = clientCtx.Codec.UnmarshalJSON(contents, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil
}

type credits struct {
	batchDenom string
	amount     string
}

var (
	reCreditAmt = `[[:digit:]]+(?:\.[[:digit:]]+)?|\.[[:digit:]]+`
	reCredits   = regexp.MustCompile(fmt.Sprintf(`^(%s) (%s)$`, reCreditAmt, ecocredit.ReBatchDenom))
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
		return credits{}, ecocredit.ErrParseFailure.Wrapf("invalid credit expression: %s", creditsStr)
	}

	return credits{
		batchDenom: matches[2],
		amount:     matches[1],
	}, nil
}

// ParseDate parses a date using the format yyyy-mm-dd.
func ParseDate(field string, date string) (time.Time, error) {
	t, err := time.Parse("2006-01-02", date)
	if err != nil {
		return t, sdkerrors.ErrInvalidRequest.Wrapf("%s must have format yyyy-mm-dd, but received %v", field, date)
	}
	return t, nil
}

// parseAndSetDate is as helper function which sets the time do the provided argument if
// the ParseDate was successful.
func parseAndSetDate(dest **time.Time, field string, date string) error {
	t, err := ParseDate(field, date)
	if err == nil {
		*dest = &t
	}
	return err
}

// checkDuplicateKey checks duplicate keys in a JSON
func checkDuplicateKey(d *json.Decoder, path []string) error {
	// Get next token from JSON
	t, err := d.Token()
	if err != nil {
		return err
	}

	delim, ok := t.(json.Delim)

	// There's nothing to do for simple values (strings, numbers, bool, nil)
	if !ok {
		return nil
	}

	switch delim {
	case '{':
		keys := make(map[string]bool)
		for d.More() {
			// Get field key
			t, err := d.Token()
			if err != nil {
				return err
			}

			key := t.(string)
			// Check for duplicates
			if keys[key] {
				return fmt.Errorf("duplicate key %s", key)
			}
			keys[key] = true

			// Check value
			if err := checkDuplicateKey(d, append(path, key)); err != nil {
				return err
			}
		}
		// Consume trailing }
		if _, err := d.Token(); err != nil {
			return err
		}

	case '[':
		i := 0
		for d.More() {
			if err := checkDuplicateKey(d, append(path, strconv.Itoa(i))); err != nil {
				return err
			}
			i++
		}
		// Consume trailing ]
		if _, err := d.Token(); err != nil {
			return err
		}

	}

	return nil
}
