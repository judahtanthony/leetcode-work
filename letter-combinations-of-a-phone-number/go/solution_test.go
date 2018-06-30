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

func TestLetterCombinations(t *testing.T) {
	for _, tt := range testcases {
		results := letterCombinations(tt.Input)
		sort.Strings(results)
		if !reflect.DeepEqual(results, tt.Output) {
			t.Fatalf("Actual %v; expecting %v\n", results, tt.Output)
		}
	}
}
