package main

import "fmt"

func NewMap (name string) map[string]string {
	if name == "" {
		return nil
	}else {
		return map[string]string {
			"name" : name,
		}
	}
}

func main () {
	data := NewMap("Ogi")

	if data == nil {
		fmt.Println("data kosong")
	}

	if data != nil {
		fmt.Println(data)
		fmt.Println(data["name"])
	}

}

