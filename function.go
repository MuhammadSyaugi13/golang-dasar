package main

import "fmt"

func main() {

    sayHello("ogi")
    fmt.Println(getHello("Aji"))

    firstname, lastname, _ := getFullName()
    fmt.Println(firstname)
    fmt.Println(lastname)

    a, b, c := getCompleteName()
    fmt.Println(a)
    fmt.Println(b)
    fmt.Println(c)
	
}

func sayHello(nama string) {
    fmt.Println("hallo");
}

func getHello(name string) string {
    hello := "Hello " + name
    return hello
}

// multiple return value
func getFullName() (string, string, int) {
    return "Moh", "Ogi", 25
}

//named return values
func getCompleteName() (firstname, middlename, lastname string){
    firstname = "m"
    middlename = "ogi"
    lastname = "gtbn"

    return firstname, middlename, lastname
}