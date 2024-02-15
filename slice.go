package main

import "fmt"

func main() {
	names :=  [...]string{"ogi", "andi", "umar", "pipah", "budi", "farhan"}

	slice1 := names[4:6]

	fmt.Println(slice1)

	slice2 := names[:3]

	fmt.Println(slice2)

	days := [...]string{"senin", "selasa", "rabu", "kamis", "jumat", "sabtu", "minggu"}
	daySclice1 := days[5:]
	daySclice1[0] = "Sabtu Baru"
	daySclice1[1] = "Minggu Baru"

	fmt.Println(daySclice1)

	daySclice2 := append(daySclice1, "Libur Baru") //menabahkan nilai baru kedalam slice
	daySclice2[0] = "Ups"
	fmt.Println(daySclice2)
	fmt.Println(days)

	//membuat slice dari awal
	var newSlice []string = make([]string, 2, 5) //make (type, length, capacity)
	newSlice[0] = "Muh"
	newSlice[1] = "Ogi"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	newSlice2 := append(newSlice, "gtbn")
	fmt.Println(newSlice2)
	fmt.Println(len(newSlice2))
	fmt.Println(cap(newSlice2))

	fmt.Println("=============================")
	newSlice2[0] = "Moh"
	fmt.Println(newSlice2)
	fmt.Println(newSlice)

	fmt.Println("===============copy slice==============")
	fromSlice := days[:]
	toSlice := make([] string, len(fromSlice), cap(fromSlice))

	copy(toSlice, fromSlice)

	fmt.Println(fromSlice)
	fmt.Println(toSlice)


	fmt.Println("===============pebedaan array dan slice==============")
	iniArray := [...]int {1,2,3,4,5}
	iniSlice := []int {1,2,3,4,5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)


}