package main

import "fmt"

func main()  {
	// If Else Statement
	x := 12
	y := 4

	if  x < y {
		fmt.Printf("%d lebih besar dari %d\n", x , y)
	} else {
		fmt.Printf("%d lebih kecil dari %d\n", y,x)
	}

	// if else if
	musim := "Hujan"

	if musim == "Kemarau" {
		fmt.Println("Panas")
	} else if musim == "Hujan"{
		fmt.Println("Dingin")
	} else {
		fmt.Println("Anda tidak berada di Indonesia")
	}

	// switch case

	switch musim {
	case "Kemarau":
		fmt.Println("Panas")
	case "Hujan":
		fmt.Println("Dingin")
	default:
		fmt.Println("Anda tidak berada di Indonesia")
	}
}