package main

import "fmt"

func main() {

    names := []string {"muh", "syaugi", "gtbn"}
    for index, name := range names {
        fmt.Println("Index", index, "=", name)
    }
        fmt.Println("==========================================")

    for _, name := range names {
        fmt.Println("Nama", "=", name)
    }
	
	
}