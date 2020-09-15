package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64

	fmt.Print("Inserte el radio: ")
	fmt.Scan(&r)

	output := math.Pi * math.Pow(r, 2)

	fmt.Println("Area del circulo: ", output)
}
