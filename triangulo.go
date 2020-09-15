package main

import "fmt"

func main() {
	var b float64
	var h float64

	fmt.Print("Inserte la medida de la base: ")
	fmt.Scan(&b)

	fmt.Print("Inserte la medida de la altura: ")
	fmt.Scan(&h)

	output := (b * h) / 2

	fmt.Println("Area del triagulo: ", output)
}
