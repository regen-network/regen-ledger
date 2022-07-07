package core

import (
	"regexp"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/regen-network/regen-ledger/types/eth"
)

var reOriginTxId = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9 _\-]{0,127}$`)
var reOriginTxSource = regexp.MustCompile(`^[a-zA-Z0-9][a-zA-Z0-9 _\-]{0,31}$`)

func (o *OriginTx) Validate() error {
	if o.Id == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("origin_tx.id cannot be empty")
	}

	if !reOriginTxId.MatchString(o.Id) {
		return sdkerrors.ErrInvalidRequest.Wrap("origin_tx.id must be at most 128 characters long, valid characters: alpha-numberic, space, '-' or '_'")
	}

	if o.Source == "" {
		return sdkerrors.ErrInvalidRequest.Wrap("origin_tx.source cannot be empty")
	}

	if !reOriginTxSource.MatchString(o.Source) {
		return sdkerrors.ErrInvalidRequest.Wrap("origin_tx.source must be at most 32 characters long, valid characters: alpha-numberic, space, '-' or '_'")
	}

	if len(o.Contract) > 0 && !eth.IsValidAddress(o.Contract) {
		return sdkerrors.ErrInvalidAddress.Wrapf("origin_tx.contract must be a valid ethereum address")
	}

	if len(o.Note) > MaxNoteLength {
		return sdkerrors.ErrInvalidRequest.Wrapf("origin_tx.note must be at most %d characters long", MaxNoteLength)
	}

	return nil
}
