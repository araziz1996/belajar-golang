package main

import (
	"fmt"
)


func main() {

	a := 5
	b := &a

	fmt.Println(a,b)

	/*
		Output yang dihasilkan adalah
		value dari a = 5, b = 0xc00004c058 atau alamat memory dari value 5.
	
	*/

	fmt.Printf("%T\n ", b)
	// output tipe data dari b adalah *int

	// gunakan *nama-variabel untuk membaca value dari alamat memory
	fmt.Println(*b)
	
	// mengganti value dengan pointer

	*b = 10
	fmt.Println(a)
}