package main

import "fmt"

// syntax penulisan
// func namafungsi (parameter tipe_data) tipe_data_hasil_return {}

func greeting(nama string) string  {
	return "Hello " + nama
}

func hasilPenjumlahan(nilai1, nilai2 int) int  {
	return nilai1 + nilai2
}

func main() {
	fmt.Println(greeting("Aziz"))
	fmt.Println(hasilPenjumlahan(12, 12))
}