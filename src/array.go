package main

import "fmt"

func main() {
	// 1
	var names [3]string
	names[0] = "Nama Saya "
	names[1] = "Danial "
	names[2] = "Fachrudin "
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	// 2
	var nilai = [3]int{
		90,
		20,
		65,
	}
	fmt.Println(nilai[1])

	// 3 function array
	var values = [...]int{
		60,
		50,
		30,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}
