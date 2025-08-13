package main

import (
	"reflect"
	"testing"
)

func TestSolutionOne(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "Basic case",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "Pair at the end",
			nums:   []int{1, 2, 3, 4, 6},
			target: 10,
			want:   []int{3, 4},
		},
		{
			name:   "No solution",
			nums:   []int{1, 2, 3},
			target: 100,
			want:   []int{},
		},
		{
			name:   "Empty array",
			nums:   []int{},
			target: 5,
			want:   []int{},
		},
		{
			name:   "Array with one element",
			nums:   []int{5},
			target: 5,
			want:   []int{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := solutionOne(tt.nums, tt.target)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solutionOne(%v, %d) = %v; want %v",
					tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
