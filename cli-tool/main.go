package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
)

// the `flag` functions take the name of the flag, a default value, and a usage description

func calculate(num1 float64, num2 float64, operator string) (float64, error) {
	switch operator {
	case "add":
		return num1 + num2, nil
	case "sub":
		return num1 - num2, nil
	case "mul":
		return num1 * num2, nil
	case "div":
		if num2 == 0 {
			return 0, errors.New("cannot divide by zero")
		}
		return num1 / num2, nil
	default:
		return 0, fmt.Errorf("unknown operator %s", operator)
	}
}

func main() {
	num1 := flag.Float64("num1", 0, "Operand 1")
	num2 := flag.Float64("num2", 0, "Operand 2")
	operator := flag.String("operator", "","Operation to perform (add, subt, mul, div)")

	flag.Parse()

	result, err := calculate(*num1, *num2, *operator)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error calculating result: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("The result is: %f\n", result)
}