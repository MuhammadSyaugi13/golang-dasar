package main

import "fmt"

func main(){
	var name = "saya"
	var e uint8 = name[2] // bertipe data uint8
	fmt.Println(e)
	var eString = string(e)
	
	fmt.Println(eString)
}