package v1

import (
	"testing"

	"github.com/regen-network/gocuke"
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
