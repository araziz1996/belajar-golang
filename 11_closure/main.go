package main

import (
	"fmt"
)

func penambah() func (int) int  {
	sum := 0

	return func (x int) int  {
		sum += x
		return sum
	}
}  

func main() {
	sum := penambah()

	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}