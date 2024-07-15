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

func main() {

	var x, y, factX, factY, perm int

	fmt.Scan(&x, &y)

	factX = faktorial(x)
	factY = faktorial(y)

	if x > y {
		perm = permutasi(x, y)
	} else {
		perm = permutasi(y, x)
	}

	fmt.Println(factX, factY, perm)
}
