package main

import "fmt"

type Customer struct {
}

func (customer Customer) sayHello() {
	fmt.Println("Halo, Nama Saya ", customer)
}

func main() {
	azzizah := Customer{Name: "Azizzah"}
	azzizah.sayHello()
}
