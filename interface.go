package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName) {
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	name string
}

func (person Person) GetName() string {
	return person.name
}

func (person Person) Jalan() string {
	return person.name + " sedang berjalan"
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	person := Person{name: "ogi"}
	SayHello(person)
	fmt.Println(person.Jalan())

	/**/
	animal := Animal{Name: "Kucing"}
	SayHello(animal)
}
