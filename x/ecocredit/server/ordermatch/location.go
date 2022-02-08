package ordermatch

import "strings"

func matchLocation(target string, filters []string) bool {
	if len(filters) == 0 {
		return true
	}

	for _, filter := range filters {
		// filters and project locations should have already been validated
		// when they were inserted so we can do a simple prefix check
		if strings.HasPrefix(target, filter) {
			return true
		}
	}

	return false
}
