package main

import "fmt"

type Address struct{
	City, Province,  Country string
}

func ChangeCountryToIndonesia(addressParam *Address){ //mengambil pointer di param 
	addressParam.Country = "Indonesia"
}

func main() {
	var address Address = Address{}
	ChangeCountryToIndonesia(&address)

	fmt.Println(address)
}