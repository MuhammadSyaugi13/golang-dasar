package main

import (
	"fmt"

	"golang.org/x/text/message"
)

func getPanic() {
	message := recover()
	fmt.Println("terjadi panic", message)
}

func startApp(value bool) {
	defer func() {
		messageRecover := recover()
		fmt.Println(messageRecover)
	}
	if value {
		panic("Ups error")
	}
}

func main() {
	startApp(true)
}
