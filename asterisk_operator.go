package main

import "fmt"

type Address struct{
	City, Province,  Country string
}

func main () {
	// tanpa asterisk
	fmt.Println("\n\n------------------Tanpa Asterisk------------------\n")

	var address1 Address = Address{"Subang", "jawa barat", "indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1
	address2.City = "Bandung"
	
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3, "\n")
	
	address3.City = "Bogor"
	address2 = &Address{"Jakarta", "DKI Jakarta", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	// dengan asterisk
	fmt.Println("\n\n------------------Dengan Asterisk------------------\n")

	*address3 = Address{"Nusa Dua", "Bali", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	
}