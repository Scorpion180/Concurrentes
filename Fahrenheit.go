package main

import "fmt"

func main() {
	var F float64

	fmt.Print("Inserte los grados Fahrenheit: ")
	fmt.Scan(&F)

	output := (F - 32) * 5 / 9

	fmt.Println("Celsius: ", output)
}
