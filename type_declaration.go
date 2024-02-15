package main

import "fmt"

func main() {
	type noKTP string

	var ktpOgi = "111111111"

	var contoh string = "2222222222"
	var contohKTP noKTP = noKTP(contoh)

	fmt.Println(ktpOgi)
	fmt.Println(contohKTP)

}