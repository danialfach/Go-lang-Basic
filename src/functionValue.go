package main

import "fmt"

func getGoodBye(nama string) string {
	return "Goodbye " + nama
}

func main() {
	goodBye := getGoodBye
	fmt.Println(goodBye("Dani"))
}
