package main

import (
	"belajar-golang-dasar/database"
	// blank intialize seperti construct di php
	_ "belajar-golang-dasar/internal"
	"fmt"
)

func main() {
	var result = database.GetDatabase()
	fmt.Println(result)
}
