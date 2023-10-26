package myCalc

import "testing"

func TestMyCalc(t *testing.T) {
	var x, y = 6, 3
	operators := []string{"+", "-", "*", "/"}
	expectedResults := []int{9, 3, 18, 2}
	for i, operator := range operators {
		realResult, _ := Calculator(x, y, operator)
		if expectedResults[i] != realResult {
			t.Errorf("expected result %d is not equal to real result %d!", expectedResults[i], realResult)
		}
	}
}
