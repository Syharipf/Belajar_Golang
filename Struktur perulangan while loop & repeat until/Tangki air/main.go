package main

import "fmt"

func main() {

	var (
		x, T, isi int
		status    bool
	)

	isi = 0
	fmt.Scan(&T)
	for isi <= T {
		fmt.Scan(&x)
		isi += x
		status = isi == T
		fmt.Println(status)
	}
}
