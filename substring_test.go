package substring

import (
	"math/rand"
	"strings"
	"testing"
	"testing/quick"
)

func testAnyContainsAny(t *testing.T, f func([]string, []string) bool) {
	for _, td := range [][2][]string{
		{{""}, {""}},
		{{"hello"}, {""}},
		{{"b"}, {"b"}},
		{{"ab"}, {"b"}},
		{{"abc"}, {"b"}},
		{{"some", "more", "words"}, {"just", "one", "word"}},
		{{"some", "more", "words"}, {"really", "or", "fewer"}},
		{{"very very long and", "short"}, {"shorter", "longer", "y very l"}},
	} {
		if !f(td[0], td[1]) {
			t.Errorf("fail, %q contains %q", td[0], td[1])
		}
	}
	for _, td := range [][2][]string{
		{{}, {}},
		{{"some", "more", "words"}, {}},
		{{}, {"hello"}},
		{{}, {""}},
		{{""}, {"hello"}},
		{{"", "hello"}, {"HELLO"}},
		{{"some", "more", "words"}, {"no", "not", "really"}},
		{{"some", "short", "words"}, {"some longer", "patterns than that"}},
		{{"very very long and", "short"}, {"shorter", "longer", " very ly"}},
	} {
		if f(td[0], td[1]) {
			t.Errorf("fail, %q does not contain %q", td[0], td[1])
		}
	}
}

func TestAnyContainsAnyNaive(t *testing.T) {
	testAnyContainsAny(t, AnyContainsAnyNaive)
}

func testAgainstNaive(t *testing.T, f func([]string, []string) bool) {
	if err := quick.CheckEqual(f, AnyContainsAnyNaive, nil); err != nil {
		t.Error(err)
	}
}

func TestAnyContainsAnyRadix(t *testing.T) {
	testAnyContainsAny(t, AnyContainsAnyRadix)
}
func TestAnyContainsAnyRadixQuick(t *testing.T) {
	testAgainstNaive(t, AnyContainsAnyRadix)
}
func TestAnyContainsAnyKR(t *testing.T) {
	testAnyContainsAny(t, AnyContainsAnyKarpRabin)
}
func TestAnyContainsAnyAC(t *testing.T) {
	testAnyContainsAny(t, AnyContainsAnyAhoCorasick)
}
func TestAnyContainsAnyACB(t *testing.T) {
	testAnyContainsAny(t, AnyContainsAnyAhoCorasickB)
}
func TestAnyContainsAnyKRQuick(t *testing.T) {
	testAgainstNaive(t, AnyContainsAnyKarpRabin)
}
func TestAnyContainsAnyKRBrute(t *testing.T) {
	testAnyContainsAny(t, AnyContainsAnyKarpRabinBrute)
}
func TestAnyContainsAnyKRBruteQuick(t *testing.T) {
	testAgainstNaive(t, AnyContainsAnyKarpRabinBrute)
}
func TestAnyContainsAnyACQuick(t *testing.T) {
	testAgainstNaive(t, AnyContainsAnyAhoCorasick)
}
func TestAnyContainsAnyACBQuick(t *testing.T) {
	testAgainstNaive(t, AnyContainsAnyAhoCorasickB)
}

func makeBenchInputHard() string {
	tokens := [...]string{
		"<a>", "<p>", "<b>", "<strong>",
		"</a>", "</p>", "</b>", "</strong>",
		"hello", "world",
	}
	x := make([]byte, 0, 1<<20)
	for {
		i := rand.Intn(len(tokens))
		if len(x)+len(tokens[i]) >= 1<<20 {
			break
		}
		x = append(x, tokens[i]...)
	}
	return string(x)
}

var benchInputHard = makeBenchInputHard()

func benchmarkMatcherHard(b *testing.B, mm MatcherMaker) {
	m := mm([]string{"<>", "</pre>", "<b>hello world</b>"})
	for i := 0; i < b.N; i++ {
		m.Matches(benchInputHard)
	}
}

func benchmarkMatcherHard2(b *testing.B, mm MatcherMaker) {
	m := mm([]string{
		"<p>a couple more</p>",
		"and longer",
		"<p>some way longer</p>",
		"<pre>patterns than</pre>",
		"the other case",
		"no idea if this makes a real difference",
		"though, no really",
		"we'll just have to see",
		"hello world hello world",
	})
	for i := 0; i < b.N; i++ {
		m.Matches(benchInputHard)
	}
}

var words = []string{
	"let's", "see", "what", "happens", "with", "lots",
	"<em>of</em>", "words", "that", "aren't", "just",
	"one", "letter", "each", "but", "THAT", "aren't",
	"really", "particularly", "long", "either", "such",
	"<b>as</b>", "hello world", "none", "that", "match",
}

func benchmarkMatcherHard3(b *testing.B, mm MatcherMaker) {
	m := mm(words)
	for i := 0; i < b.N; i++ {
		m.Matches(benchInputHard)
	}
}

func benchmarkMatcherHard4(b *testing.B, mm MatcherMaker) {
	m := mm([]string{
		"xyz", "nothing", "matches", "xylophone", "man",
		"of", "in", "an", "any", "if", "it", "is", "don't", "really",
		"know", "what", "I'm", "doing", "here",
	})
	for i := 0; i < b.N; i++ {
		m.Matches(benchInputHard)
	}
}

func BenchmarkBruteHard(b *testing.B)          { benchmarkMatcherHard(b, MakeBrute) }
func BenchmarkRadixHard(b *testing.B)          { benchmarkMatcherHard(b, MakeRadix) }
func BenchmarkRabinKarpHard(b *testing.B)      { benchmarkMatcherHard(b, MakeRabinKarp) }
func BenchmarkRabinKarpBruteHard(b *testing.B) { benchmarkMatcherHard(b, MakeRabinKarpBrute) }
func BenchmarkAhoCorasickHard(b *testing.B)    { benchmarkMatcherHard(b, MakeAhoCorasick) }
func BenchmarkAhoCorasickBHard(b *testing.B)   { benchmarkMatcherHard(b, MakeAhoCorasickB) }

func BenchmarkBruteHard2(b *testing.B)          { benchmarkMatcherHard2(b, MakeBrute) }
func BenchmarkRadixHard2(b *testing.B)          { benchmarkMatcherHard2(b, MakeRadix) }
func BenchmarkRabinKarpHard2(b *testing.B)      { benchmarkMatcherHard2(b, MakeRabinKarp) }
func BenchmarkRabinKarpBruteHard2(b *testing.B) { benchmarkMatcherHard2(b, MakeRabinKarpBrute) }
func BenchmarkAhoCorasickHard2(b *testing.B)    { benchmarkMatcherHard2(b, MakeAhoCorasick) }
func BenchmarkAhoCorasickBHard2(b *testing.B)   { benchmarkMatcherHard2(b, MakeAhoCorasickB) }

func BenchmarkBruteHard3(b *testing.B)          { benchmarkMatcherHard3(b, MakeBrute) }
func BenchmarkRadixHard3(b *testing.B)          { benchmarkMatcherHard3(b, MakeRadix) }
func BenchmarkRabinKarpHard3(b *testing.B)      { benchmarkMatcherHard3(b, MakeRabinKarp) }
func BenchmarkRabinKarpBruteHard3(b *testing.B) { benchmarkMatcherHard3(b, MakeRabinKarpBrute) }
func BenchmarkAhoCorasickHard3(b *testing.B)    { benchmarkMatcherHard3(b, MakeAhoCorasick) }
func BenchmarkAhoCorasickBHard3(b *testing.B)   { benchmarkMatcherHard3(b, MakeAhoCorasickB) }

func BenchmarkBruteHard4(b *testing.B)          { benchmarkMatcherHard4(b, MakeBrute) }
func BenchmarkRadixHard4(b *testing.B)          { benchmarkMatcherHard4(b, MakeRadix) }
func BenchmarkRabinKarpHard4(b *testing.B)      { benchmarkMatcherHard4(b, MakeRabinKarp) }
func BenchmarkRabinKarpBruteHard4(b *testing.B) { benchmarkMatcherHard4(b, MakeRabinKarpBrute) }
func BenchmarkAhoCorasickHard4(b *testing.B)    { benchmarkMatcherHard4(b, MakeAhoCorasick) }
func BenchmarkAhoCorasickHardB4(b *testing.B)   { benchmarkMatcherHard4(b, MakeAhoCorasickB) }

func benchmarkMake(b *testing.B, mm MatcherMaker) {
	for i := 0; i < b.N; i++ {
		mm(words)
	}
}

func BenchmarkMakeBrute(b *testing.B)          { benchmarkMake(b, MakeBrute) }
func BenchmarkMakeRadix(b *testing.B)          { benchmarkMake(b, MakeRadix) }
func BenchmarkMakeRabinKarp(b *testing.B)      { benchmarkMake(b, MakeRabinKarp) }
func BenchmarkMakeRabinKarpBrute(b *testing.B) { benchmarkMake(b, MakeRabinKarpBrute) }
func BenchmarkMakeAhoCorasick(b *testing.B)    { benchmarkMake(b, MakeAhoCorasick) }
func BenchmarkMakeAhoCorasickB(b *testing.B)   { benchmarkMake(b, MakeAhoCorasickB) }

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func makePatterns(min, max, n int) []string {
	var ps []string
	for i := 0; i < n; i++ {
		l := rand.Intn(max-min) + min
		p := make([]rune, l)
		for j := 0; j < l; j++ {
			p[j] = letters[rand.Intn(26)*2]
		}
		ps = append(ps, string(p))
	}
	return ps
}

func makeStrings(min, max, n, k int, patterns []string, match float64) []string {
	var ss []string
	for i := 0; i < n; i++ {
		var parts []string
		for j := 0; j < k; j++ {
			l := rand.Intn(max-min) + min
			p := make([]rune, l)
			for j := 0; j < l; j++ {
				p[j] = letters[rand.Intn(26)*2+1]
			}
			parts = append(parts, string(p))
		}
		if rand.Float64() < match {
			p := patterns[rand.Intn(len(patterns))]
			parts[rand.Intn(len(parts))] = p
		}
		ss = append(ss, strings.Join(parts, ""))
	}
	return ss
}

var patternsRand = makePatterns(5, 10, 50)
var stringsRand = makeStrings(5, 10, 50, 5, patternsRand, 0.1)

func benchmarkMatcherRand(b *testing.B, mm MatcherMaker) {
	m := mm(patternsRand)
	for i := 0; i < b.N; i++ {
		for _, s := range stringsRand {
			m.Matches(s)
		}
	}
}

func BenchmarkBruteRand(b *testing.B)          { benchmarkMatcherRand(b, MakeBrute) }
func BenchmarkRadixRand(b *testing.B)          { benchmarkMatcherRand(b, MakeRadix) }
func BenchmarkRabinKarpRand(b *testing.B)      { benchmarkMatcherRand(b, MakeRabinKarp) }
func BenchmarkRabinKarpBruteRand(b *testing.B) { benchmarkMatcherRand(b, MakeRabinKarpBrute) }
func BenchmarkAhoCorasickRand(b *testing.B)    { benchmarkMatcherRand(b, MakeAhoCorasick) }
func BenchmarkAhoCorasickBRand(b *testing.B)   { benchmarkMatcherRand(b, MakeAhoCorasickB) }

var stringsRandM = makeStrings(5, 10, 50, 5, patternsRand, 0.9)

func benchmarkMatcherRandM(b *testing.B, mm MatcherMaker) {
	m := mm(patternsRand)
	for i := 0; i < b.N; i++ {
		for _, s := range stringsRandM {
			m.Matches(s)
		}
	}
}

func BenchmarkBruteRandM(b *testing.B)          { benchmarkMatcherRandM(b, MakeBrute) }
func BenchmarkRadixRandM(b *testing.B)          { benchmarkMatcherRandM(b, MakeRadix) }
func BenchmarkRabinKarpRandM(b *testing.B)      { benchmarkMatcherRandM(b, MakeRabinKarp) }
func BenchmarkRabinKarpBruteRandM(b *testing.B) { benchmarkMatcherRandM(b, MakeRabinKarpBrute) }
func BenchmarkAhoCorasickRandM(b *testing.B)    { benchmarkMatcherRandM(b, MakeAhoCorasick) }
func BenchmarkAhoCorasickRandBM(b *testing.B)   { benchmarkMatcherRandM(b, MakeAhoCorasickB) }
