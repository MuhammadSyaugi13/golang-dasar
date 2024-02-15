package main

import "fmt"

func getPanic() {
	message := recover()
	fmt.Println("terjadi panic", message)
}

func startApp(value bool) {
	defer getPanic()
	if value {
		panic("Ups error")
	}
}

func main() {
	startApp(true)
}
