package main

import "fmt"

func main() {
	var a = 6
	var b = 7
	var jumlah = a + b
	var total = a % b

	fmt.Println(jumlah)
	fmt.Println(total)

	var i = 10
	i += 5
	fmt.Println(i)

	var j = 5
	j++
	j++
	fmt.Println(j)
}
