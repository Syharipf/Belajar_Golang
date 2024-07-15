package main

import "fmt"

func main() {

	var (
		n     int
		hasil bool
	)

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		hasil = n%i == 0
		fmt.Println(i, hasil)
	}
}
