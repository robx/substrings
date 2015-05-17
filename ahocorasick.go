package substring

import (
	"github.com/gansidui/ahocorasick"
)

type AhoCorasick struct {
	*ahocorasick.Matcher
	hasEmpty bool
}

func MakeAhoCorasick(patterns []string) Matcher {
	m := &AhoCorasick{
		Matcher: ahocorasick.NewMatcher(),
	}
	for _, p := range patterns {
		if p == "" {
			m.hasEmpty = true
		}
	}
	m.Build(patterns)
	return m
}

func (m AhoCorasick) Matches(s string) bool {
	// TODO: should short cut on first hit.
	return m.hasEmpty || len(m.Match(s)) > 0
}

func AnyContainsAnyAhoCorasick(ss []string, patterns []string) bool {
	m := MakeAhoCorasick(patterns)
	for _, s := range ss {
		if m.Matches(s) {
			return true
		}
	}
	return false
}
