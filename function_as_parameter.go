package main

import "fmt"

type FILTER func(string) string

func sayHelloWithFilter(name string, filter FILTER) {
	filteredName := filter(name)
	fmt.Println("Hello " + filteredName)
}

func spamFilter(name string) string {
	if name == "anjing" {
		return "..."
	} else {
		return name
	}
}

func main() {
	//function parameter
	sayHelloWithFilter("Aji", spamFilter)

	filter := spamFilter
	sayHelloWithFilter("anjing", filter) //function parameter
}
