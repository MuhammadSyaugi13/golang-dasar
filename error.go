package main

import (
	"errors"
	"fmt"
)

func pembagi(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("pembagi bernilai 0")
	} else {
		return nilai / pembagi, nil
	}
}

func main() {
	hasil, err := pembagi(100, 0)

	if err == nil {
		fmt.Println(hasil)
	} else {
		fmt.Println(err.Error())
	}

}
