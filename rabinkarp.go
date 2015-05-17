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

func hashes(patterns []string) (int, uint32, map[uint32]*radix.Tree) {
	min := -1
	for _, p := range patterns {
		if l := len(p); min == -1 || l < min {
			min = l
		}
	}
	pow := hashPow(min)

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
	return min, pow, trees
}

func AnyContainsAnyKarpRabin(ss []string, patterns []string) bool {
	n, pow, trees := hashes(patterns)
	switch n {
	case -1:
		return false
	case 0:
		return len(ss) > 0
	}

	for _, s := range ss {
		if len(s) < n {
			continue
		}

		// Rabin-Karp search
		var h uint32
		for i := 0; i < n; i++ {
			h = h*primeRK + uint32(s[i])
		}
		if t := trees[h]; t != nil {
			_, val, _ := t.LongestPrefix(s)
			if val != nil {
				return true
			}
		}
		for i := n; i < len(s); {
			h *= primeRK
			h += uint32(s[i])
			h -= pow * uint32(s[i-n])
			i++
			if t := trees[h]; t != nil {
				_, val, _ := t.LongestPrefix(s[i-n:])
				if val != nil {
					return true
				}
			}
		}
	}
	return false
}
