package main

import "fmt"

func main() {
	nomor := 1
	nama := "dani"

	switch nomor {
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
	default:
		fmt.Println("Hanya 1 & 2")
		break
	}

	// dengan short statement
	switch length := len(nama); length > 5 {
	case true:
		fmt.Println("Terlalu panjang")
	case false:
		fmt.Println("Benar")
	}

	// tanpa kondisi
	panjang := len(nama)
	switch {
	case panjang > 10:
		fmt.Println("nama terlalu panjang")
	case panjang > 5:
		fmt.Println("nama lumayan panjang")
	default:
		fmt.Println("Benar")
	}
}
