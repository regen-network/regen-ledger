package table

import (
	"encoding/json"
	"io"

	"github.com/regen-network/regen-ledger/orm/v2/internal/store"
)

func (s Store) DefaultJSON() json.RawMessage {
	return json.RawMessage("[]")
}

func (s Store) ValidateJSON(reader io.Reader) error {
	panic("implement me")
}

func (s Store) ImportJSON(kvStore store.KVStore, reader io.Reader) error {
	panic("implement me")
}

func (s Store) ExportJSON(kvStore store.KVStore, writer io.Writer) error {
	panic("implement me")
}
