package main

import "fmt"

func main()  {

	// Mendefinisakn map
	emails := make(map[string]string)

	// Mengisi Key - Value

	emails["Aziz"] = "aziz@email.com"
	emails["Danu"] = "danu@gmail.com"

	/*
		emails["Aziz"] merupakan Key
		"aziz@email" merupakan value dari key
	*/

	fmt.Println(emails)

	// Cara lain penulisan map
	departments := map[string]string{
		"Aziz":"IT Manager",
		"Balkin":"IT Support",
	}

	fmt.Println(departments)

	delete(departments,"Balkin")
	
	fmt.Println(departments)
}