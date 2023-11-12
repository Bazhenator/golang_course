package main

func InfiniteMultiply(n float64, m float64) float64 {
	k := float64(1.0)
	result := m
	for k < n {
		result *= (m + k)
		k++
	}
	return result
}

func Factorial(n float64) float64 {
	result := float64(1.0)
	for i := float64(1.0); i < n+1; i++ {
		result *= i
	}
	return result
}

func MyPow(x float64, n float64) float64 {
	k := float64(0.0)
	result := float64(1.0)
	for k != n {
		result *= x
		k++
	}
	return result
}
