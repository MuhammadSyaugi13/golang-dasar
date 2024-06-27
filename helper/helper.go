package helper

import "fmt"

var version = "1.0.0"
var Application = "golang"

func sayGoodBye(name string) string {
	return "goodbye " + name
}

func SayHello(name string) string {
	return "hello " + name
}

// access modifier seperti public private pada php
func ContohAccessModifier(name string) {
	sayGoodBye(name)
	fmt.Println(version)
}
