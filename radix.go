package substring

import (
	radix "github.com/armon/go-radix"
)

type Radix struct {
	*radix.Tree
}

var _ Matcher = Radix{}

func MakeRadix(patterns []string) Radix {
	r := Radix{radix.New()}
	for _, p := range patterns {
		r.Insert(p, struct{}{})
	}
	return r
}

func (r Radix) Matches(s string) bool {
	// i == len(s) handles ""
	for i := 0; i <= len(s); i++ {
		_, val, _ := r.LongestPrefix(s[i:])
		if val != nil {
			return true
		}
	}
	return false
}

func AnyContainsAnyRadix(ss []string, patterns []string) bool {
	r := MakeRadix(patterns)
	for _, s := range ss {
		if r.Matches(s) {
			return true
		}
	}
	return false
}
