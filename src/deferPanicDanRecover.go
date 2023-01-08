package main

import "fmt"

func logging() {
	fmt.Println("Selesai memanggil function (defer)")
}

func runApplication() {
	defer logging()
	fmt.Println("Run Application (defer)")
}

func endApp() {
	fmt.Println("End App (panic)")
	message := recover() //recover
	fmt.Println("terjadi error", message)
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic("ini Erorr (panic)")
	}
}

func main() {
	runApplication()
	endApp()
}
