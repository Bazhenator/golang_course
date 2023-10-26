package main

import (
	"fmt"
	"sort"
)

// Example of tested file
func findKthLargest(nums []int, k int) int {
	//Array sorting with Ints func from sort package
	sort.Ints(nums)

	//Cheking if k is out of array's range
	if k <= len(nums) && k >= 0 {
		return nums[len(nums)-k]
	} else {
		fmt.Print("out of range")
	}
	return 0
}
