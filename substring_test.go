package substring

import (
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

func TestAnyContainsAny(t *testing.T) {
	testAnyContainsAny(t, AnyContainsAny)
}

func testAgainstNaive(t *testing.T, f func([]string, []string) bool) {
	if err := quick.CheckEqual(f, AnyContainsAnyNaive, nil); err != nil {
		t.Error(err)
	}
}

func TestAnyContainsAnyQuick(t *testing.T) {
	testAgainstNaive(t, AnyContainsAny)
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

func TestAnyContainsAnyKRQuick(t *testing.T) {
	testAgainstNaive(t, AnyContainsAnyKarpRabin)
}
