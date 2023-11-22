package v1

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
)

type class struct {
	t     gocuke.TestingT
	class *Class
	err   error
}

func TestClass(t *testing.T) {
	gocuke.NewRunner(t, &class{}).Path("./features/state_class.feature").Run()
}

func (s *class) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *class) TheClass(a gocuke.DocString) {
	s.class = &Class{}
	err := jsonpb.UnmarshalString(a.Content, s.class)
	require.NoError(s.t, err)
}

func (s *class) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.class.Metadata = strings.Repeat("x", int(length))
}

func (s *class) TheClassIsValidated() {
	s.err = s.class.Validate()
}

func (s *class) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *class) ExpectNoError() {
	require.NoError(s.t, s.err)
}
