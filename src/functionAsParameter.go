package main

import "fmt"

type Filter func(string) string

func sayHelloWithFilter(name string, filter func(string) string) {
	fmt.Println("Hello", filter(name))
}

func spamFilter(name string) string {
	if name == "Dani" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Dani", spamFilter)

	filter := spamFilter("")
	fmt.Println("ganteng", filter)
}
