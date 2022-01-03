package ecocredit

import (
	"testing"

	"github.com/stretchr/testify/require"
	"pgregory.net/rapid"
)

func TestProject(t *testing.T) {
	t.Run("TestValidateFormatProjectID", rapid.MakeCheck(testValidateFormatProjectID))
	t.Run("TestInvalidProjectsError", rapid.MakeCheck(testInvalidProjectsError))
	t.Run("TestValidateGenProjectID", rapid.MakeCheck(testValidateGenProjectID))
}

var genClassID = rapid.Custom(func(t *rapid.T) string {
	return rapid.StringMatching(`[A-Z]{1,3}`).Draw(t, "classID").(string)
})

// Property: ValidateProjectID(FormatProjectID(a)) == nil
func testValidateFormatProjectID(t *rapid.T) {
	classID := genClassID.Draw(t, "classID").(string)
	projectSeqNo := rapid.Uint64Range(1, 1000000).Draw(t, "projectSeqNo").(uint64)

	projectID := FormatProjectID(classID, projectSeqNo)
	err := ValidateProjectID(projectID)
	require.NoError(t, err)
}

var genProjectID = rapid.Custom(func(t *rapid.T) string {
	return rapid.StringMatching(`[A-Za-z0-9]{2,16}`).Draw(t, "projectID").(string)
})

// Property: ValidateProjectID(genProjectID(a)) == nil
func testValidateGenProjectID(t *rapid.T) {
	projectID := genProjectID.Draw(t, "projectID").(string)

	err := ValidateProjectID(projectID)
	require.NoError(t, err)
}

// genInvalidProjectID generates strings that don't conform to the ProjectID format
var genInvalidProjectID = rapid.OneOf(
	rapid.StringMatching(`^[a-zA-Z0-9]{1}$`),
	rapid.StringMatching(`^[a-zA-Z0-9]{17,}$`),
	rapid.StringMatching(`^[a-zA-Z0-9]*[!@#$&()\\-\x60.+,/\"]+[a-zA-Z0-9]*$`),
)

func testInvalidProjectsError(t *rapid.T) {
	projectID := genInvalidProjectID.Draw(t, "projectID").(string)
	require.Error(t, ValidateProjectID(projectID))
}

func TestProjectID(t *testing.T) {
	testCases := []struct {
		name      string
		projectID string
		expectErr bool
	}{
		{
			"valid project id",
			"A123",
			false,
		},
		{
			"invalid project id min length",
			"a",
			true,
		},
		{
			"invalid project id max length",
			"abcdef123456ghijklmnop789",
			true,
		},
		{
			"invalid project id special characters",
			"abcd@1",
			true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if err := ValidateProjectID(tc.projectID); tc.expectErr {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
