package main

import "fmt"

func main() {
	var input float64
	fmt.Print("Inserte la medida el lado del cuadrado: ")
	fmt.Scanf("%f", &input)
	output := input * input
	fmt.Println("Area del cuadrado: ", output)
}
