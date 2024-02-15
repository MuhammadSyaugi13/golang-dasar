package main
import "fmt"

func sumAll(numbers ...int) int {
    total := 0

    for _, number := range numbers {
        total += number
    }

    return total
}

func main(){
    fmt.Println(sumAll(10, 10, 10, 10, 10))

    //merubah slice menjadi vargs
    numbers := []int{10,10,10} //slice
    fmt.Println(sumAll(numbers...)) //numbers... adalah slice yg diubah menjadu vargs
}