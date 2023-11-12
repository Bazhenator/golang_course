package main

import (
	"fmt"
	"task_4/pkg/math"
)

var nums = []float64{
	1.0,
	0.0001,
	0.9333,
	0.0,
	0.2,
}

func main() {

	top := nums[0]
	bottom := nums[0]
	aN := nums[0]
	eps := nums[1]
	x := nums[2]
	copyX := nums[2]
	n := nums[3]
	m := nums[4]

	result := float64(aN)

	for aN > eps {
		n += 1.0
		x = math.MyPow(copyX, n)
		//fmt.Printf("Pow(x^%g) = %g \n", n, x)
		top = math.InfiniteMultiply(n, m)
		//fmt.Printf("InfiniteMultiply(%g) = %g \n", n, top)
		bottom = math.Factorial(n)
		//fmt.Printf("Factorial(%g) = %g \n", n, bottom)
		aN = (top * x) / bottom
		//fmt.Printf("a%g is %g \n \n", n, aN)
		result += aN
	}
	fmt.Printf("The result is: %g \n", result)
}
