package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan Ke ", counter)
		counter++
	}

	// for dengan statement
	for counter := 1; counter <= 10; counter++ {
		fmt.Println("ke ", counter)
	}

	// for dengan range
	names := []string{"Danial", "Fachrudin"}
	for index, name := range names {
		fmt.Println("index", index, "=", name)
	}
}
