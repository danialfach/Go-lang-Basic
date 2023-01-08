package main

import "fmt"

func main() {
	type noKTP string

	var ktpDani noKTP = "anjai "
	fmt.Println(ktpDani)
	fmt.Println(noKTP("122424145"))
}
