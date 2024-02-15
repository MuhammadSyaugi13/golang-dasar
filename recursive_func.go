package main

import "fmt"

func faktorialRecursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * faktorialRecursive(value-1)
	}

}

func main() {
	fmt.Println(faktorialRecursive(4))
}
