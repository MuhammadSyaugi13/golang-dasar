package main

import "fmt"

type Man struct{
	Name string
}

func (man *Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("dalam method")
	fmt.Println(man)
}

func main() {
	ogi := Man{"Ogi"}
	ogi.Married()
	fmt.Println(ogi)
}
