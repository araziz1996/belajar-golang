package main

import (
	"fmt"
	"math"
	// import package yang telah dibuat
	"belajar-golang/03_package/mypkg"
)

func main()  {
	fmt.Println(math.Floor(3.1)) // pembulatan kebawah
	fmt.Println(math.Ceil(3.1)) // pembulatan kebawah
	fmt.Println(math.Sqrt(16)) // squareroot
	
	fmt.Println(mypkg.Reverse("Aziz"))
}