package v1

import (
	"fmt"
	"strconv"
	"testing"

	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/reflect/protoreflect"

	api "github.com/regen-network/regen-ledger/api/v2/regen/ecocredit/v1"
)

type msgUpdateProjectEnrollment struct {
	gocuke.TestingT
	msg *MsgUpdateProjectEnrollment
	err error
}

func TestMsgUpdateProjectEnrollment(t *testing.T) {
	gocuke.NewRunner(t, &msgUpdateProjectEnrollment{}).Path("./features/msg_update_project_enrollment.feature").Run()
}

func (s *msgUpdateProjectEnrollment) Before() {
	s.msg = &MsgUpdateProjectEnrollment{}
}

func (s *msgUpdateProjectEnrollment) Issuer(a string) {
	s.msg.Issuer = a
}

func (s *msgUpdateProjectEnrollment) ProjectId(a string) {
	s.msg.ProjectId = a
}

func (s *msgUpdateProjectEnrollment) ClassId(a string) {
	s.msg.ClassId = a
}

func (s *msgUpdateProjectEnrollment) NewStatus(a string) {
	n, err := strconv.Atoi(a)
	if err == nil {
		s.msg.NewStatus = ProjectEnrollmentStatus(n)
	} else {
		var status api.ProjectEnrollmentStatus
		value := status.Descriptor().Values().ByName(protoreflect.Name(fmt.Sprintf("PROJECT_ENROLLMENT_STATUS_%s", a)))
		if value == nil {
			s.Fatalf("invalid status: %s", a)
		}
		s.msg.NewStatus = ProjectEnrollmentStatus(value.Number())
	}
}

func (s *msgUpdateProjectEnrollment) Metadata(a string) {
	s.msg.Metadata = a
}

func (s *msgUpdateProjectEnrollment) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgUpdateProjectEnrollment) ExpectErrorContains(a string) {
	if a == "" {
		require.NoError(s, s.err)
	} else {
		require.ErrorContains(s, s.err, a)
	}
}

func (s *msgUpdateProjectEnrollment) ExpectGetsignersReturns(a string) {
	require.Equal(s, a, s.msg.GetSigners()[0].String())
}
