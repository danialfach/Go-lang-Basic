package main

import "fmt"

func integer() (int, int) {
	return 16, 10
}

func main() {
	firstNumber, lastNumber := integer()
	fmt.Println(firstNumber, lastNumber)
}
