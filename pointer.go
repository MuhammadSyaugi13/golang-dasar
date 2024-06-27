package main

import "fmt"

type Address struct{
	City, Province,  Country string
}

func main () {
	
	// tanpa pointer
	fmt.Println("\n\n---------Tanpa Pointer------------------\n")
	address1 := Address{"Subang", "Jabar", "Indonesia"}
	address2 := address1

	fmt.Println(address1)
	fmt.Println(address2)
	
	address2.City = "Bandung"

	fmt.Println(address1)
	fmt.Println(address2)

	// dengan pointer
	fmt.Println("\n\n---------Dengan Pointer------------------\n")
	var addressp1 Address = Address{"Bogor", "Jabar", "Indonesia"}
	var addressp2 *Address = &addressp1

	fmt.Println(addressp1)
	fmt.Println(addressp2)
	
	addressp2.City = "Cianjur"

	fmt.Println(addressp1)
	fmt.Println(addressp2)

 }