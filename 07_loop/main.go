package main

import "fmt"

func main() {
	
	// Cara 1
	i := 1
	for i <= 10 {
		fmt.Println(i)
		// i = i + 1 atau
		i++
	}

	// Cara

	for i := 0; i <= 10; i++ {
		fmt.Printf("Number %d\n", i)
	}

	// FizzBuzz 
	count := 100
	for i := 0; i < count; i++ {
		if i % 15 == 0 {
			fmt.Println("FizzBuzz")
		} else if i % 3 == 0 {
			fmt.Println("Fizz")
		} else if i % 5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}