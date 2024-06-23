package main

import "fmt"

/*
* interface kosong (any) bisa mengembalikan semua macam jenis tipe data
*/
func Ups() any {
	// return "hello" //mengembalikan string
	// return 123 //mengembalikan integer
	return true //mengembalikan boolean
} 

func main() {
	interfaceKosong := Ups()
	fmt.Println(interfaceKosong)
}