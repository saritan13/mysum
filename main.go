package main

import (
	"fmt"

	"github.com/calc/calc"
)

func main() {
	fmt.Println("This is sum")
	result := calc.Sum(10, 20)
	fmt.Println("Sum of two number", result)
	result = calc.Sub(10, 20)
	fmt.Println("Sum of two number", result)
	result = calc.Multi(10, 20)
	fmt.Println("Sum of two number", result)
	result1 := calc.Div(20.0, 10.0)
	fmt.Println("Sum of two number", result1)

}
