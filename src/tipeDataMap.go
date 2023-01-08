package main

import "fmt"

func main() {
	person := map[string]string{
		"name":    "Danial",
		"address": "Malang",
	}

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	book := make(map[string]string)
	book["id"] = "1"
	book["namaSekolah"] = "Kanesa"
	book["wrong"] = "ups"

	delete(book, "wrong")

	fmt.Println(book)
}
