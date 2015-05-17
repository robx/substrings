package substring

import (
	radix "github.com/armon/go-radix"
)

func AnyContainsAnyRadix(ss []string, patterns []string) bool {
	r := radix.New()
	for _, p := range patterns {
		r.Insert(p, struct{}{})
	}

	for _, s := range ss {
		// i == len(s) handles ""
		for i := 0; i <= len(s); i++ {
			_, val, _ := r.LongestPrefix(s[i:])
			if val != nil {
				return true
			}
		}
	}
	return false
}
