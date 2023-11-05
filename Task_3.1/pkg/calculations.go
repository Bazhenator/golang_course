package calculations

import log "github.com/sirupsen/logrus"

//Claculate uses logs whether isLogged is true and returns the result of factorial(n)
func Calculate(n int64, isLogged bool) int64 {
	var result int64 = 1
	if *&isLogged {
		log.Println("Start calculations...")
		log.Printf("Calculate <%d>!", n)
	}
	for i := int64(1); i < n+1; i++ {
		result *= i
	}
	if *&isLogged {
		log.Println("Calculations complete!")
	}
	return result
}
