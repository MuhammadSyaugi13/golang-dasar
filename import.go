package main

import (
	"belajar-golang-dasar/helper"
	"fmt"
)

func main() {
	result := helper.SayHello("ogi")
	fmt.Println(result)

	fmt.Println(helper.Application)

	helper.ContohAccessModifier("ogi")
}
