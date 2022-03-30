package server

import (
	"github.com/regen-network/regen-ledger/x/data"
)

var _ data.MsgServer = serverImpl{}
