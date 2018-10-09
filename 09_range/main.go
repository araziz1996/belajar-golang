package main

import "fmt"

func main()  {
	numbers := []int{32,23,24,65,78,54,12}

	//Melakukan perulangan menggunakan range

	for i, number := range numbers {
		fmt.Printf("ID - %d adalah %d\n", i, number)
	}
	// variabel i merupakan index dari range number

	// perulangan tanpa menggunakan index

	for _, number := range numbers{
		fmt.Printf("Number : %d\n", number)
	}

	// 

	sum := 0
	for _, number := range numbers {
		sum += number
	}

	fmt.Println(sum)

	// Perulangan dan menampilkan isi dari map

	departments := map[string]string{
		"Aziz":"IT Manager",
		"Balkin":"IT Support",
	}

	for key, value := range departments {
		fmt.Printf("%s : %s\n", key,value)
	}

	for key := range departments {
		fmt.Println("Name : "+ key)
	}

}