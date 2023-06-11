package hackerrank

import (
	"fmt"
	"testing"
)

type TC struct {
	Input    []string
	Expected []string
}

var testcases = []TC{
	TC{
		Input: []string{
			"a", "ab", "abc", "abcd",
		},
		Expected: []string{
			"a", "abc", "abcd", "ab",
		},
	},
	TC{
		Input: []string{
			"abc", "ab", "abcde", "a", "abcd",
		},
		Expected: []string{
			"a", "abc", "abcde", "abcd", "ab",
		},
	},
}

func TestCustomSort(t *testing.T) {
	for _, tc := range testcases {
		ans := customSorting(tc.Input)
		if len(ans) != len(tc.Expected) {
			t.Fatal("unexpected array size")
		}

		for i := 0; i < len(tc.Expected); i++ {
			if ans[i] != tc.Expected[i] {
				t.Fatal("incorrect answer")
			}
		}
	}
}

func TestCustomSortV2(t *testing.T) {
	for _, tc := range testcases {
		ans := customSortingV2(tc.Input)
		if len(ans) != len(tc.Expected) {
			t.Fatal("unexpected array size")
		}

		for i := 0; i < len(tc.Expected); i++ {
			if ans[i] != tc.Expected[i] {
				fmt.Println("actual: ", ans)
				fmt.Println("expected: ", tc.Expected)
				t.Fatal("incorrect answer")

			}
		}
	}
}
