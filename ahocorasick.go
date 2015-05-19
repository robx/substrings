package substrings

import (
	ahocorasickb "github.com/cloudflare/ahocorasick"
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

type AhoCorasickB struct {
	*ahocorasickb.Matcher
	hasEmpty bool
}

func MakeAhoCorasickB(patterns []string) Matcher {
	hasEmpty := false
	for _, p := range patterns {
		if p == "" {
			hasEmpty = true
		}
	}
	return &AhoCorasickB{
		Matcher:  ahocorasickb.NewStringMatcher(patterns),
		hasEmpty: hasEmpty,
	}
}

func (m AhoCorasickB) Matches(s string) bool {
	// TODO: should short cut on first hit.
	return m.hasEmpty || len(m.Match([]byte(s))) > 0
}

func AnyContainsAnyAhoCorasickB(ss []string, patterns []string) bool {
	m := MakeAhoCorasickB(patterns)
	for _, s := range ss {
		if m.Matches(s) {
			return true
		}
	}
	return false
}
