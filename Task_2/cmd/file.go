package main

import (
	"fmt"
	"sort"
)

// Example of tested file
func findKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	if k <= len(nums) && k > 0 {
		return nums[len(nums)-k]
	} else {
		fmt.Print("out of range")
	}
	return 0
}
