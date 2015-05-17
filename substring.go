package substring

import (
	"strings"
)

type Matcher interface {
	Matches(s string) bool
}

type MatcherMaker func(patterns []string) Matcher

type Brute []string

var _ Matcher = Brute([]string{})

func MakeBrute(patterns []string) Matcher {
	return Brute(patterns)
}

func (patterns Brute) Matches(s string) bool {
	for _, p := range patterns {
		if strings.Contains(s, p) {
			return true
		}
	}
	return false
}

func AnyContainsAnyNaive(ss []string, patterns []string) bool {
	b := Brute(patterns)
	for _, s := range ss {
		if b.Matches(s) {
			return true
		}
	}
	return false
}
