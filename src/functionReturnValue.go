package main

import "fmt"

func getHello(nama string) string {
	return "Hello " + nama
}

func main() {
	memanggil := getHello("Dani")
	fmt.Println(memanggil)
}
