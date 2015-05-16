package substring

import (
	"strings"
)

func AnyContainsAnyNaive(ss []string, substrs []string) bool {
	for _, s := range ss {
		for _, substr := range substrs {
			if strings.Contains(s, substr) {
				return true
			}
		}
	}
	return false
}

func AnyContainsAny(ss []string, substrs []string) bool {
	return AnyContainsAnyNaive(ss, substrs)
}
