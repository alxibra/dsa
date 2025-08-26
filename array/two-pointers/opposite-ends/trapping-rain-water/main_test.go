package main

import "testing"

func TestSolutionOne(t *testing.T) {
	testCases := []struct {
		desc     string
		height   []int
		expected int
	}{
		{
			desc:     "Example 1",
			height:   []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			expected: 6,
		},
		{
			desc:     "Example 2",
			height:   []int{4, 2, 0, 3, 2, 5},
			expected: 9,
		},
		{
			desc:     "No bars",
			height:   []int{},
			expected: 0,
		},
		{
			desc:     "Single bar",
			height:   []int{4},
			expected: 0,
		},
		{
			desc:     "Flat surface",
			height:   []int{4, 4, 4, 4},
			expected: 0,
		},
		{
			desc:     "U shape",
			height:   []int{3, 0, 0, 0, 3},
			expected: 9,
		},
		{
			desc:     "Descending slop",
			height:   []int{5, 4, 3, 2, 1},
			expected: 0,
		},
		{
			desc:     "Descending slop",
			height:   []int{1, 2, 3, 4, 5},
			expected: 0,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			actual := solutionOne(tC.height)
			if actual != tC.expected {
				t.Errorf("solutionOne(%v) = %d; want %d", tC.height, actual, tC.expected)
			}
		})
	}
}
