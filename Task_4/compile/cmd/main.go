package main

import (
	"fmt"
)

func main() {

	top := float64(1.0)
	bottom := float64(1.0)
	aN := float64(1.0)
	eps := float64(0.0001)
	x := float64(0.9333)
	copyX := x
	n := float64(0.0)
	m := float64(0.2)

	result := float64(aN)

	for aN > eps {
		n += 1.0
		x = MyPow(copyX, n)
		//fmt.Printf("Pow(x^%g) = %g \n", n, x)
		top = InfiniteMultiply(n, m)
		//fmt.Printf("InfiniteMultiply(%g) = %g \n", n, top)
		bottom = Factorial(n)
		//fmt.Printf("Factorial(%g) = %g \n", n, bottom)
		aN = (top * x) / bottom
		//fmt.Printf("a%g is %g \n \n", n, aN)
		result += aN
	}
	fmt.Printf("The result is: %g \n", result)
}



