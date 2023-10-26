package myCalc

import "fmt"

func Calculator(op1 int, op2 int, operator string) (int, error) {
	switch operator {
	case "+":
		return op1 + op2, nil
	case "-":
		return op1 - op2, nil
	case "*":
		return op1 * op2, nil
	case "/":
		if op2 == 0 {
			return -1, fmt.Errorf("division by zero was occured")
		}
		return op1 / op2, nil
	}
	return 0, nil
}
