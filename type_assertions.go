package main

import "fmt"

func random() any {
	return 123
}

func main() {
	// var result any = random()
	// var resultString string = result.(string) //type assertion
	// fmt.Println(resultString)
	
	// var resultInt int = result.(int) //type assertion salah
	// fmt.Println(resultInt)

	
	/* type assertions dengan switch*/
	var result any = random()
	switch value := result.(type) {
		case string:
			fmt.Println("string", value)
		case int: 
			fmt.Println("integer", value)
		default:
			fmt.Println("unknown", value)
	}

	
	/* ./ type assertions dengan switch*/


}