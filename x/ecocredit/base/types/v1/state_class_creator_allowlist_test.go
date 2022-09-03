package v1

import (
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type classCreatorAllowlistSuite struct {
	t                     gocuke.TestingT
	classCreatorAllowlist *ClassCreatorAllowlist
	err                   error
}

func TestClassCreatorAllowlist(t *testing.T) {
	gocuke.NewRunner(t, &classCreatorAllowlistSuite{}).Path(
		"./features/state_class_creator_allowlist.feature",
	).Run()
}

func (s *classCreatorAllowlistSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *classCreatorAllowlistSuite) TheClassCreatorAllowlist(a gocuke.DocString) {
	s.classCreatorAllowlist = &ClassCreatorAllowlist{}
	err := jsonpb.UnmarshalString(a.Content, s.classCreatorAllowlist)
	require.NoError(s.t, err)
}

func (s *classCreatorAllowlistSuite) TheClassCreatorAllowlistIsValidated() {
	s.err = s.classCreatorAllowlist.Validate()
}

func (s *classCreatorAllowlistSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}
