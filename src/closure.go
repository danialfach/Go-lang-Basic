package main

import "fmt"

func main() {
	counter := 0
	iniApaYa := func() {
		fmt.Println("ini namanya increment")
		counter++
	}

	iniApaYa()
	iniApaYa()
}
