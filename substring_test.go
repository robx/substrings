package substring

import (
	"testing"
	"testing/quick"
)

func testAnyContainsAny(t *testing.T, f func([]string, []string) bool) {
	for _, td := range [][2][]string{
		{{""}, {""}},
		{{"hello"}, {""}},
		{{"some", "more", "words"}, {"just", "one", "word"}},
		{{"some", "more", "words"}, {"really", "or", "fewer"}},
	} {
		if !f(td[0], td[1]) {
			t.Errorf("fail, %v contains %v", td[0], td[1])
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
	} {
		if f(td[0], td[1]) {
			t.Errorf("fail, %v does not contain %v", td[0], td[1])
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
