package main

import "fmt"

func main()  {
	// Array
	var player [2]string
	// menambahkan nilai
	player[0] = "Nesta"
	player[1] = "Maldini"

	/* 
	
	mendeklarasikan dan menambahkan value dengan panjang array diketahui atau ditentukan.

	milanPlayer := [4]string{"Maldini","Cafu","Dida","Ancelloti"}
	
	*/
	
	//Array slice, array tanpa perlu mendeklarasikan panjang array tsb.

	milanPlayer := []string{"Maldini","Cafu","Dida","Ancelloti"}

	
	fmt.Println(player)
	
	fmt.Println(milanPlayer)

	// mencari value dengan index 0
	fmt.Println(milanPlayer[0]) 
	// mencari panjang array
	fmt.Println(len(milanPlayer)) 
	// mencari value dengan range index 1 sampai 2
	fmt.Println(milanPlayer[0:2]) 
}	