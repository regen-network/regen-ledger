package core

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type project struct {
	t       gocuke.TestingT
	project *Project
	err     error
}

func TestProject(t *testing.T) {
	gocuke.NewRunner(t, &project{}).Path("./features/state_project.feature").Run()
}

func (s *project) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *project) TheProject(a gocuke.DocString) {
	s.project = &Project{}
	err := jsonpb.UnmarshalString(a.Content, s.project)
	require.NoError(s.t, err)
}

func (s *project) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.project.Metadata = strings.Repeat("x", int(length))
}

func (s *project) TheProjectIsValidated() {
	s.err = s.project.Validate()
}

func (s *project) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *project) ExpectNoError() {
	require.NoError(s.t, s.err)
}
