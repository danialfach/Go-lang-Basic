package main

import "fmt"

func main() {
	names := [...]string{"Danial", "Fachrudin", "Ganteng", "Banget", "jangan", "ganggu"}
	slice := names[0:6]

	fmt.Println(slice[0])
	fmt.Println(slice[1])

	days := [...]string{"Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu", "minggu"}
	daySlice1 := days[5:]
	daySlice1[0] = "Sabtu Baru"
	daySlice1[1] = "Minggu baru"
	fmt.Println(days) //[Senin,Selasa,Rabu,Kamis,Jumat,Sabtu,Minggu]

	daySlice2 := append(daySlice1, "Libur Baru")
	daySlice2[0] = "ups"
	fmt.Println(daySlice2)
	fmt.Println(days)

	newSlice := make([]string, 2, 5)
	newSlice[0] = "Dani"
	newSlice[1] = "Ganteng"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	fromSlice := days[:]
	toSlice := make([]string, len(fromSlice), cap(fromSlice))
	copy(toSlice, fromSlice)

	fmt.Println(toSlice)

	iniArray := [...]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
}
