package testsuite

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestGenesis(t *testing.T) {
	ff := setup(t)
	s := NewGenesisTestSuite(ff)
	suite.Run(t, s)
}
