package substring

import (
	radix "github.com/armon/go-radix"
)

// based on go stdlib strings/strings.go

const primeRK = 16777619

// hashStr returns the hash and the appropriate multiplicative
// factor for use in Rabin-Karp algorithm.
func hashStr(sep string) uint32 {
	hash := uint32(0)
	for i := 0; i < len(sep); i++ {
		hash = hash*primeRK + uint32(sep[i])
	}
	return hash
}

func hashPow(n int) uint32 {
	var pow, sq uint32 = 1, primeRK
	for i := n; i > 0; i >>= 1 {
		if i&1 != 0 {
			pow *= sq
		}
		sq *= sq
	}
	return pow
}

type RabinKarp struct {
	trees map[uint32]*radix.Tree
	n     int
	pow   uint32
}

var _ Matcher = RabinKarp{}

func MakeRabinKarp(patterns []string) RabinKarp {
	min := -1
	for _, p := range patterns {
		if l := len(p); min == -1 || l < min {
			min = l
		}
	}

	trees := make(map[uint32]*radix.Tree)
	for _, s := range patterns {
		h := hashStr(s[:min])
		t, ok := trees[h]
		if !ok {
			t = radix.New()
			trees[h] = t
		}
		t.Insert(s, struct{}{})
	}
	return RabinKarp{
		trees: trees,
		n:     min,
		pow:   hashPow(min),
	}
}

func (m RabinKarp) Matches(s string) bool {
	switch m.n {
	case -1:
		return false
	case 0:
		return true
	}
	if len(s) < m.n {
		return false
	}

	// Rabin-Karp search
	var h uint32
	for i := 0; i < m.n; i++ {
		h = h*primeRK + uint32(s[i])
	}
	if t := m.trees[h]; t != nil {
		_, val, _ := t.LongestPrefix(s)
		if val != nil {
			return true
		}
	}
	for i := m.n; i < len(s); {
		h *= primeRK
		h += uint32(s[i])
		h -= m.pow * uint32(s[i-m.n])
		i++
		if t := m.trees[h]; t != nil {
			_, val, _ := t.LongestPrefix(s[i-m.n:])
			if val != nil {
				return true
			}
		}
	}
	return false
}

func AnyContainsAnyKarpRabin(ss []string, patterns []string) bool {
	m := MakeRabinKarp(patterns)
	for _, s := range ss {
		if m.Matches(s) {
			return true
		}
	}
	return false
}
