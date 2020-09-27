package main

import "fmt"

func main() {
	var dia int
	var mes int

	fmt.Scan(&dia)

	fmt.Scan(&mes)

	switch mes {
	case 1:
		if dia < 21 {
			fmt.Print("Capricornio")
		} else {
			fmt.Print("Acuario")
		}
	case 2:
		if dia < 21 {
			fmt.Print("Acuario")
		} else {
			fmt.Print("Piscis")
		}
	case 3:
		if dia < 21 {
			fmt.Print("Piscis")
		} else {
			fmt.Print("Aries")
		}
	case 4:
		if dia < 21 {
			fmt.Print("Aries")
		} else {
			fmt.Print("Tauro")
		}
	case 5:
		if dia < 22 {
			fmt.Print("Tauro")
		} else {
			fmt.Print("Geminis")
		}
	case 6:
		if dia < 22 {
			fmt.Print("Geminis")
		} else {
			fmt.Print("Cancer")
		}
	case 7:
		if dia < 23 {
			fmt.Print("Cancer")
		} else {
			fmt.Print("Leo")
		}
	case 8:
		if dia < 23 {
			fmt.Print("Leo")
		} else {
			fmt.Print("Virgo")
		}
	case 9:
		if dia < 24 {
			fmt.Print("Virgo")
		} else {
			fmt.Print("Libra")
		}
	case 10:
		if dia < 25 {
			fmt.Print("Libra")
		} else {
			fmt.Print("Escorpio")
		}
	case 11:
		if dia < 23 {
			fmt.Print("Escorpio")
		} else {
			fmt.Print("Sagitario")
		}
	case 12:
		if dia < 22 {
			fmt.Print("Sagitario")
		} else {
			fmt.Print("Capricornio")
		}
	}
}
