package main

import (
	"cal/calAdd"
	"cal/calDiv"
	"cal/calMul"
	"cal/calSub"
	"fmt"
)

func main() {
	var num1, num2 float64
	var operator string

	fmt.Print("Enter the first number: ")
	fmt.Scan(&num1)

	fmt.Print("Enter an operator (+, -, *, /): ")
	fmt.Scan(&operator)

	fmt.Print("Enter the second number: ")
	fmt.Scan(&num2)

	var result float64

	switch operator {
	case "+":
		result = calAdd.Add(num1, num2)
	case "-":
		result = calSub.Subtract(num1, num2)
	case "*":
		result = calMul.Mul(num1, num2)
	case "/":
		result = calDiv.Div(num1, num2)
	default:
		fmt.Println("Invalid operator.")
		return
	}

	fmt.Printf("Result: %.2f\n", result)
}
