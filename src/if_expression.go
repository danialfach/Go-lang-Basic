package main

import "fmt"

func main() {
	name := "Dani"

	if name == "Dani" {
		fmt.Println("Iya")
	} else if name == "Danial" {
		fmt.Println("Oke")
	} else {
		fmt.Println("Tidak")
	}

	// if short statement
	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Benar")
	}
}
