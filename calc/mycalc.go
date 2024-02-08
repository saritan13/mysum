package calc

import "log"

func Sum(a, b int) int {
	log.Print("Executing sum function")
	return a + b
}

func Sub(a, b int) int {
	log.Print("Executing sub function")
	return a - b
}

func Multi(a, b int) int {
	log.Print("Executing multi function")
	return a * b
}

func Div(a, b float32) float32 {
	log.Print("Executing div function")
	return a / b
}

func Mod(a, b int) int {
	log.Print("Executing mod function")
	return a % b
}

func Sqr(a int) int {
	log.Print("Executing Square function")
	return a / a
}

func Cube(a int) int {
	log.Print("Executing Cube function")
	return a * a * a
}
