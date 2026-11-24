package main

import "fmt"

// LongestZeroSumSubarray returns the length of the longest contiguous
// subarray whose elements sum to zero. A repeated prefix sum means the
// elements between the two occurrences cancel out.
func LongestZeroSumSubarray(nums []int) int {
	first := map[int]int{0: -1} // prefix sum -> earliest index it occurred
	sum, best := 0, 0
	for i, v := range nums {
		sum += v
		if j, ok := first[sum]; ok {
			if i-j > best {
				best = i - j
			}
		} else {
			first[sum] = i
		}
	}
	return best
}

func main() {
	fmt.Println(LongestZeroSumSubarray([]int{9, -3, 3, -1, 6, -5}))
}
