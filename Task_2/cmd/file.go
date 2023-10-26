package main

import (
	"fmt"
)

func quickSort(arr []int) {
	if len(arr) <= 1 {
		return
	}

	pivot := arr[len(arr)/2]
	left, right := make([]int, 0), make([]int, 0)
	middle := make([]int, 0)

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v > pivot {
			right = append(right, v)
		} else {
			middle = append(middle, v)
		}
	}

	quickSort(left)
	quickSort(right)

	copy(arr, left)
	copy(arr[len(left):], middle)
	copy(arr[len(left)+len(middle):], right)
}

// Example of tested file
func findKthLargest(nums []int, k int) int {
	//Checking Quick sort
	quickSort(nums)
	//Cheking if k is out of array's range
	if k <= len(nums) && k >= 0 {
		return nums[len(nums)-k]
	} else {
		fmt.Print("out of range")
	}
	return 0
}
