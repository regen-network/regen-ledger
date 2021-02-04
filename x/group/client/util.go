package client

import (
	"encoding/json"
	"io/ioutil"

	"github.com/regen-network/regen-ledger/x/group"
	"github.com/spf13/pflag"
)

func parseMembersFlag(fs *pflag.FlagSet) ([]group.Member, error) {
	members := []group.Member{}
	membersFile, _ := fs.GetString(flagMembers)

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
