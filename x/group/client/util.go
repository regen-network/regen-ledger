package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/regen-network/regen-ledger/x/group"
)

func parseMembers(membersFile string) ([]group.Member, error) {
	members := []group.Member{}

	if membersFile == "" {
		return members, nil
	}

	contents, err := ioutil.ReadFile(membersFile)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &members)
	if err != nil {
		return nil, err
	}

	return members, nil
}
