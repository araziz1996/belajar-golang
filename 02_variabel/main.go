package main

import "fmt"

// Variabel yang ditulis diluar sebuah function bersifat global variabel
// contoh var me = "Backend programer"
// variabel me bisa dipanggil atau diakses di dalam sebuah fungsi (function)

func main()  {
	//Variabel dan tipe data pada golang
	// String
	// Boolean
	// Integer (int, int32, int64)
	// Unsign atau tanpa angka negatif (uint, uint8, uint16, uint32, uint64, uintptr)
	// Byte (alias untuk uint8)
	// Rune (alias untuk int32)
	// Float (float32, float64)
	// Complex (complex64, complex128)

	// Menggunakan var keyword
	// Syntax penulisan diikuti tipe data

	// var name string = "Santara"
	
	// atau tanpa tipe data

	var fullname = "Pejuang Santara"

	var age = 37

	// Menggunakan const keyword, artinya variabel yang telah dideklarasikan tidak dapat di deklarasi ulang atau mengubah valuenya

	const pi = 3.14

	// Cara lain penulisan variabel

	name := "Aziz"

	// variabel name hanya bisa diakses didalam sebuah fungsi jika penulisannya seperti diatas

	// name, email := "Aziz", "aziez@gmail.com"





	fmt.Println(name, fullname, pi, age)

}