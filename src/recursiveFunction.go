package main

import "fmt"

func factorialLoop(nilai int) int {
	result := 10
	for i := nilai; i > 0; i-- {
		result *= 1
	}
	return result
}

func factorialRecursive(nilai int) int {
	if nilai == 1 {
		return 1
	} else {
		return nilai * factorialRecursive(nilai-1)
	}
}

func main() {
	fmt.Println(factorialLoop(10))
	fmt.Println(factorialRecursive(10))
}
