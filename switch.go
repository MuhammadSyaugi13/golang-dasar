package main

import "fmt"

func main() {

    name := "ogisssss"
    length := len(name)
    switch{
        case length > 10:
            fmt.Println("nama terlalu panjang")
        case length > 5:
            fmt.Println("nama lumayan panjang")
        default:
            fmt.Println("nama sudah benar")

    }
	
	
}