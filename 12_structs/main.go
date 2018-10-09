package main

import (
	"fmt"
)

// Mendefinisikan struct

type Person struct{
	namaDepan string
	namaBelakang string
	alamat string
	umur int
}

func main() {
	// orang1 := Person{namaDepan: "Aziz", namaBelakang: "Anis", alamat: "Pemalang", umur:22}


	orang1 := Person{"Aziz","Anis","Pemalang",22}

	fmt.Println(orang1)
	
}