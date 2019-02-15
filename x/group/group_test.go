package group

import (
	"flag"
	"github.com/DATA-DOG/godog"
	"github.com/DATA-DOG/godog/colors"
	"os"
	"testing"
)

var opt = godog.Options{
	Output: colors.Colored(os.Stdout),
	Format: "progress", // can define default values
}

func init() {
	godog.BindFlags("godog.", flag.CommandLine, &opt)
}

func TestMain(m *testing.M) {
	ci, found := os.LookupEnv("CI")
	if found && len(ci) != 0 {
		f, err := os.Create("test_output.xml")
		if err == nil {
			opt.Output = f
			opt.Format = "junit"
		}
	}

	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		FeatureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}

func aPublicKeyAddress() error {
	return godog.ErrPending
}

func aUserCreatesAGroupWithThatAddress() error {
	return godog.ErrPending
}

func aDecisionThresholdOf(arg1 int) error {
	return godog.ErrPending
}

func theyShouldGetANewGroupAddressBack() error {
	return godog.ErrPending
}

func aGroupID() error {
	return godog.ErrPending
}

func aUserGetsTheGroupDetailsOnTheCommandLine() error {
	return godog.ErrPending
}

func theyShouldGetBackTheGroupDetailsInJSONFormat() error {
	return godog.ErrPending
}



func FeatureContext(s *godog.Suite) {
	s.Step(`^a public key address$`, aPublicKeyAddress)
	s.Step(`^a user creates a group with that address$`, aUserCreatesAGroupWithThatAddress)
	s.Step(`^a decision threshold of (\d+)$`, aDecisionThresholdOf)
	s.Step(`^they should get a new group address back$`, theyShouldGetANewGroupAddressBack)
	s.Step(`^a group ID$`, aGroupID)
	s.Step(`^a user gets the group details on the command line$`, aUserGetsTheGroupDetailsOnTheCommandLine)
	s.Step(`^they should get back the group details in JSON format$`, theyShouldGetBackTheGroupDetailsInJSONFormat)
}

