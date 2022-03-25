package structvalid

import "fmt"

// StrMaxLen checks if given string s length is not bigger then max.
// Appends error message got the errs list otherwise.
func StrMaxLen(field, s string, max int, errs []string) []string {
	if len(s) > max {
		return append(errs, fmt.Sprintf("Max length of %q is %d.", field, len(s)))
	}
	return errs
}
