package main

import (
	"fmt"
)

func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	pivot := arr[0]
	var left, right []int

	for _, v := range arr[1:] {
		if v <= pivot {
			left = append(left, v)
		} else {
			right = append(right, v)
		}
	}

	left = quickSort(left)
	right = quickSort(right)
	return append(append(left, pivot), right...)
}

// Example of tested file
func findKthLargest(nums []int, k int) int {
	//Checking Quick sort
	nums = quickSort(nums)
	//Cheking if k is out of array's range
	if k <= len(nums) && k >= 0 {
		return nums[len(nums)-k]
	} else {
		fmt.Print("out of range")
	}
	return 0
}
