package ecocredit

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestSuites(t *testing.T) {
	suite.Run(t, &ModelSuite{})
}
