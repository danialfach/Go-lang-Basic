package main

import "fmt"

func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	total := sumAll(20, 30)
	fmt.Println(total)

	// slice parameter
	numbers := []int{10, 10, 10, 10, 10, 2}
	total = sumAll(numbers...)
	fmt.Println(total)
}
