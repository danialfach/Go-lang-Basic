package main

import "fmt"

func main() {
	const firstName = "Danial"
	const lastName = "Fachrudin"

	fmt.Println(firstName + lastName)

	const (
		age     = 17
		address = "Malang"
	)

	fmt.Println(age)
	fmt.Println(address)
}
