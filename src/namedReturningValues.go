package main

import "fmt"

func namaLengkap() (namaDepan, namaBelakang string) {
	namaDepan = "Danial"
	namaBelakang = "Fachrudin"

	return namaDepan, namaBelakang
}

func main() {
	depan, belakang := namaLengkap()
	fmt.Println(depan, belakang)
}
