package main

import (
	myCalc "GO/Task_1/pkg/myCalc"
	"fmt"
	"log"
)

func main() {
	var input1 int
	fmt.Print("Введите первое число: ")
	_, err := fmt.Scan(&input1)
	for err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение:")
		fmt.Print("Введите первое число: ")
		_, err = fmt.Scan(&input1)
	}

	var operator string
	fmt.Print("Выберите операцию (+, -, *, /): ")
	fmt.Scan(&operator)
	for !(operator == "+" || operator == "-" || operator == "*" || operator == "/") {
		fmt.Println("Некорректная операция. Пожалуйста, используйте символы +, -, * или /.")
		fmt.Print("Выберите операцию (+, -, *, /): ")
		fmt.Scan(&operator)
	}

	var input2 int
	fmt.Print("Введите второе число: ")
	_, err = fmt.Scan(&input2)
	for err != nil {
		fmt.Println("Некорректное число. Пожалуйста, введите числовое значение:")
		fmt.Print("Введите второе число: ")
		_, err = fmt.Scan(&input2)
	}

	result, err := myCalc.Calculator(input1, input2, operator)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Результат: %d %s %d = %d.", input1, operator, input2, result)
}
