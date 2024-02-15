package main

import "fmt"

func main() {

    var person map[string]string = map[string]string{}
    person["nama"] = "ogi"
    person["alamat"] = "bogor"

	fmt.Println(person)

    book := make(map[string]string) //membuat map
    book["judul"] = "buku golang" // menambahkan map
    book["penulis"] = "M. syaugi"
    book["penerbit"] = "Alam Jaya"

    fmt.Println(book)

    delete(book, "penerbit") //menghapus map

    fmt.Println(book)
	
	
}