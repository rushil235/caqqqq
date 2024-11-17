package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scanln(&num1)

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Enter the second number: ")
	fmt.Scanln(&num2)

	result := 0.0

	switch operator {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		if num2 != 0 {
			result = num1 / num2
		} else {
			fmt.Println("Error: Division by zero.")
			return
		}
	default:
		fmt.Println("Invalid operator")
		return
	}

	fmt.Printf("Result: %f\n", result)
}
