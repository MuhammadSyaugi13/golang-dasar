package main

import "fmt"

type Address struct{
	City, Province,  Country string
}

func main () {
	// new alamat1 *Address= &Address{}
	alamat1 := new(Address)
	alamat2 := alamat1

	alamat2.Country = "Indonesia"

	fmt.Println(alamat1)
	fmt.Println(alamat2)

}