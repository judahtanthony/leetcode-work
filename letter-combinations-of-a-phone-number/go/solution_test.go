package leetcode

import (
	"reflect"
	"sort"
	"testing"
)

type TestCases []struct {
	Input  string
	Output []string
}

// Input: "23"
// Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
var testcases = TestCases{
	{Input: "23", Output: []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}},
}

func TestLetterCombinationsRealloc(t *testing.T) {
	for _, tt := range testcases {
		results := letterCombinationsRealloc(tt.Input)
		sort.Strings(results)
		if !reflect.DeepEqual(results, tt.Output) {
			t.Fatalf("Actual %v; expecting %v\n", results, tt.Output)
		}
	}
}
func BenchmarkLetterCombinationsRealloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinationsRealloc("234")
	}
}

func TestLetterCombinationsReduced(t *testing.T) {
	for _, tt := range testcases {
		results := letterCombinationsReduced(tt.Input)
		sort.Strings(results)
		if !reflect.DeepEqual(results, tt.Output) {
			t.Fatalf("Actual %v; expecting %v\n", results, tt.Output)
		}
	}
}
func BenchmarkLetterCombinationsReduced(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinationsReduced("234")
	}
}

func TestLetterCombinations(t *testing.T) {
	for _, tt := range testcases {
		results := letterCombinations(tt.Input)
		sort.Strings(results)
		if !reflect.DeepEqual(results, tt.Output) {
			t.Fatalf("Actual %v; expecting %v\n", results, tt.Output)
		}
	}
}
func BenchmarkLetterCombinations(b *testing.B) {
	for i := 0; i < b.N; i++ {
		letterCombinations("234")
	}
}
