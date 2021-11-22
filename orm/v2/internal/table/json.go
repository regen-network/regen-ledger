package table

import (
	"encoding/json"
	"io"

	"github.com/regen-network/regen-ledger/orm/v2/backend/kv"
)

func (s TableModel) DefaultJSON() json.RawMessage {
	return json.RawMessage("[]")
}

func (s TableModel) ValidateJSON(reader io.Reader) error {
	panic("implement me")
}

func (s TableModel) ImportJSON(kvStore kv.KVStore, reader io.Reader) error {
	panic("implement me")
}

func (s TableModel) ExportJSON(kvStore kv.ReadKVStore, writer io.Writer) error {
	panic("implement me")
}
