package main

import "fmt"

func faktorial(n int) int {
	var hasil int

	if n == 0 || n == 1 {
		return 1
	}

	hasil = 1
	for i := 2; i <= n; i++ {
		hasil *= i
	}
	return hasil
}

func permutasi(x, y int) int {
	return faktorial(x) / faktorial(x-y)
}

func kombinasi(x, y int) int {
	return faktorial(x) / (faktorial(x-y) * faktorial(y))
}

func main() {
	var a, b, c, d int

	fmt.Scan(&a, &b, &c, &d)

	if a >= c {
		fmt.Println(permutasi(a, c))
		fmt.Println(kombinasi(a, c))
	}

	if b >= d {
		fmt.Println(permutasi(b, d))
		fmt.Println(kombinasi(b, d))
	}
}
