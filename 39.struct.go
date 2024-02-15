package main

import "fmt"

type Customer struct {
	Name, Address string
	Age           int
}

func main() {
	var ogi Customer
	fmt.Println(ogi)

	ogi.Name = "M Syaugi"
	ogi.Address = "Bogor"
	ogi.Age = 24

	fmt.Println(ogi)

	// struct literals
	joko := Customer{
		Name:    "Joko",
		Address: "Jawa",
		Age:     30,
	}

	fmt.Println(joko)

	budi := Customer{"budi", "Jakarta", 25}
	fmt.Println(budi)
}
