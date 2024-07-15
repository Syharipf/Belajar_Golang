package main

import "fmt"

func fibo(n int) int {
	var a, b, hasil, total int
	if n <= 1 {
		return 0
	}

	a, b = 0, 1
	total = a + b

	for i := 2; i < n; i++ {
		hasil = a + b
		total += hasil
		a = b
		b = hasil
	}

	return total
}

func main() {

	var n, hasil int

	fmt.Scan(&n)
	hasil = fibo(n)
	fmt.Println(hasil)
}
