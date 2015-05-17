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

func benchmarkMatcherHard3(b *testing.B, mm MatcherMaker) {
	m := mm([]string{
		"let's", "see", "what", "happens", "with", "lots",
		"<em>of</em>", "words", "that", "aren't", "just",
		"one", "letter", "each", "but", "THAT", "aren't",
		"really", "particularly", "long", "either", "such",
		"<b>as</b>", "hello world", "none", "that", "match",
	})
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

func BenchmarkBruteHard2(b *testing.B)          { benchmarkMatcherHard2(b, MakeBrute) }
func BenchmarkRadixHard2(b *testing.B)          { benchmarkMatcherHard2(b, MakeRadix) }
func BenchmarkRabinKarpHard2(b *testing.B)      { benchmarkMatcherHard2(b, MakeRabinKarp) }
func BenchmarkRabinKarpBruteHard2(b *testing.B) { benchmarkMatcherHard2(b, MakeRabinKarpBrute) }
func BenchmarkAhoCorasickHard2(b *testing.B)    { benchmarkMatcherHard2(b, MakeAhoCorasick) }

func BenchmarkBruteHard3(b *testing.B)          { benchmarkMatcherHard3(b, MakeBrute) }
func BenchmarkRadixHard3(b *testing.B)          { benchmarkMatcherHard3(b, MakeRadix) }
func BenchmarkRabinKarpHard3(b *testing.B)      { benchmarkMatcherHard3(b, MakeRabinKarp) }
func BenchmarkRabinKarpBruteHard3(b *testing.B) { benchmarkMatcherHard3(b, MakeRabinKarpBrute) }
func BenchmarkAhoCorasickHard3(b *testing.B)    { benchmarkMatcherHard3(b, MakeAhoCorasick) }

func BenchmarkBruteHard4(b *testing.B)          { benchmarkMatcherHard4(b, MakeBrute) }
func BenchmarkRadixHard4(b *testing.B)          { benchmarkMatcherHard4(b, MakeRadix) }
func BenchmarkRabinKarpHard4(b *testing.B)      { benchmarkMatcherHard4(b, MakeRabinKarp) }
func BenchmarkRabinKarpBruteHard4(b *testing.B) { benchmarkMatcherHard4(b, MakeRabinKarpBrute) }
func BenchmarkAhoCorasickHard4(b *testing.B)    { benchmarkMatcherHard4(b, MakeAhoCorasick) }

var benchInputTorture = strings.Repeat("ABC", 1<<10) + "123" + strings.Repeat("ABC", 1<<10)
var benchNeedleTorture = strings.Repeat("ABC", 1<<10+1)
