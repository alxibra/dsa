package main

import (
	"reflect"
	"testing"

	"golang.org/x/exp/slices"
)

func TestSolutionOne(t *testing.T) {
	testCases := []struct {
		desc     string
		input    []int
		expected []int
	}{
		{
			desc:     "Mixed zeros and non-zeros",
			input:    []int{0, 1, 0, 3, 12},
			expected: []int{1, 3, 12, 0, 0},
		},
		{
			desc:     "All zeros",
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		{
			desc:     "No zeros",
			input:    []int{1, 2, 3},
			expected: []int{1, 2, 3},
		},
		{
			desc:     "Zeros at the beginning",
			input:    []int{0, 0, 1, 2},
			expected: []int{1, 2, 0, 0},
		},
		{
			desc:     "Single zero",
			input:    []int{0},
			expected: []int{0},
		},
		{
			desc:     "Single non-zero",
			input:    []int{5},
			expected: []int{5},
		},
		{
			desc:     "Alternating zero and non-zero",
			input:    []int{0, 1, 0, 2, 0, 3},
			expected: []int{1, 2, 3, 0, 0, 0},
		},
	}

	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			t.Parallel()
			actual := solutionOne(slices.Clone(tC.input))
			if !reflect.DeepEqual(actual, tC.expected) {
				t.Errorf("failed %s: got %v, want %v", tC.desc, actual, tC.expected)
			}
		})
	}
}
