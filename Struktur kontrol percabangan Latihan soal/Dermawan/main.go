package main

import "fmt"

func main() {

	var (
		n, uang, tabungan int
	)

	fmt.Scan(&n)

	tabungan = 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&uang)
		tabungan += uang
	}

	fmt.Println(tabungan)
}
