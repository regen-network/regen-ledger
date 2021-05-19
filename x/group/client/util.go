package client

import (
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/client"

	"github.com/regen-network/regen-ledger/x/group"
)

func parseMembers(clientCtx client.Context, membersFile string) ([]group.Member, error) {
	members := group.Members{}

	if membersFile == "" {
		return members.Members, nil
	}

	contents, err := ioutil.ReadFile(membersFile)
	if err != nil {
		return nil, err
	}

	err = clientCtx.JSONCodec.UnmarshalJSON(contents, &members)
	if err != nil {
		return nil, err
	}

	return members.Members, nil
}
