package main

import (
	"fmt"
)

// Example of tested file
func findKthLargest(nums []int, k int) int {
	//Checking Bubble sort
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(nums)-1; i++ {
			if nums[i] > nums[i+1] {
				c := nums[i]
				nums[i] = nums[i+1]
				nums[i+1] = c
				swapped = true
			}
		}
	}
	//Cheking if k is out of array's range
	if k <= len(nums) && k >= 0 {
		return nums[len(nums)-k]
	} else {
		fmt.Print("out of range")
	}
	return 0
}
