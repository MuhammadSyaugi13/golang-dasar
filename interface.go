package main

import "fmt"

type HasName interface {
	GetName() string
}

func SayHello(value HasName){
	fmt.Println("Hello", value.GetName())
}

type Person struct {
	name string
}

func (person Person) GetName() string {
	return person.name
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

	/**/
	animal :=  Animal{Name: "Kucing"}
	SayHello(animal)
}
