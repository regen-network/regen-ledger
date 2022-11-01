package v1

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	"gotest.tools/v3/assert"
)

type registerAccountSuite struct {
	t         gocuke.TestingT
	msg       *MsgRegisterAccount
	signBytes string
	err       error
}

func TestMsgRegisterAccount(t *testing.T) {
	gocuke.NewRunner(t, &registerAccountSuite{}).Path("./features/msg_register_account.feature").Run()
}

func (s *registerAccountSuite) Before(t gocuke.TestingT) {
	s.t = t
}

func (s *registerAccountSuite) TheMessage(a gocuke.DocString) {
	var msg MsgRegisterAccount
	err := json.Unmarshal([]byte(a.Content), &msg)
	assert.NilError(s.t, err)
	s.msg = &msg
}

func (s *registerAccountSuite) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *registerAccountSuite) ExpectNoError() {
	assert.NilError(s.t, s.err)
}

func (s *registerAccountSuite) ExpectTheError(a string) {
	assert.ErrorContains(s.t, s.err, a)
}

func (s *registerAccountSuite) MessageSignBytesQueried() {
	s.signBytes = string(s.msg.GetSignBytes())
}

func (s *registerAccountSuite) ExpectTheSignBytes(a gocuke.DocString) {
	buffer := new(bytes.Buffer)
	require.NoError(s.t, json.Compact(buffer, []byte(a.Content)))
	require.Equal(s.t, buffer.String(), s.signBytes)
}
