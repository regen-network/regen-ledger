package util

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

func GodogMain(m *testing.M, suiteName string, featureContext func(suite *godog.Suite)) {
	/*
		ci, found := os.LookupEnv("CI")
		if found && len(ci) != 0 {
			testfileName := fmt.Sprintf("%s/TEST_%s.xml", os.Getenv("CI_PROJECT_DIR"), suiteName)
			f, err := os.Create(testfileName)
			if err != nil {
				fmt.Printf("Error creating test output file %s: %+v", testfileName, err)
			} else {
				opt.Output = f
				opt.Format = "junit"
			}
		}*/

	flag.Parse()
	opt.Paths = flag.Args()

	status := godog.RunWithOptions("godogs", func(s *godog.Suite) {
		featureContext(s)
	}, opt)

	if st := m.Run(); st > status {
		status = st
	}
	os.Exit(status)
}
