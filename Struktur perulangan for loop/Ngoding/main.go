package main

import "fmt"

func main() {

	var (
		n, m, total int
		hasil       float32
	)

	fmt.Scan(&n)
	total = 0
	for i := 1; i <= n; i++ {
		fmt.Scan(&m)
		total += m
	}

	hasil = float32(total) / float32(n)
	fmt.Println(hasil)
}
