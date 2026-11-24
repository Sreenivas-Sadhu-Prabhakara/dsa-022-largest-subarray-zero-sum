package main

import "testing"

func TestLongestZeroSumSubarray(t *testing.T) {
	cases := []struct {
		in   []int
		want int
	}{
		{[]int{9, -3, 3, -1, 6, -5}, 5},
		{[]int{1, -1, 3, 2, -2, -8, 1, 7, 10, 23}, 5},
		{[]int{1, 2, 3}, 0},
		{[]int{0}, 1},
		{[]int{0, 0, 0}, 3},
		{[]int{}, 0},
		{[]int{1, -1}, 2},
	}
	for _, c := range cases {
		if got := LongestZeroSumSubarray(c.in); got != c.want {
			t.Errorf("LongestZeroSumSubarray(%v) = %d, want %d", c.in, got, c.want)
		}
	}
}
