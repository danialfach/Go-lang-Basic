package main

import "fmt"

type Pelanggan struct {
	Name, Address string
	Age           int
}

func main() {
	var Dani Pelanggan
	Dani.Name = "Danial Fachrudin"
	Dani.Address = "Malang"
	Dani.Age = 17
	fmt.Println(Dani)

	Joko := Pelanggan{
		Name:    "Joko",
		Address: "Malang",
		Age:     17,
	}
	fmt.Println(Joko)

	dono := Pelanggan{"Dono", "Jakarta", 19}
	fmt.Println(dono)
}
