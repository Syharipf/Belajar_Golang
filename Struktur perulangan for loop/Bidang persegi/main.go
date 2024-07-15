package main

import "fmt"

func main() {

	var n, s, keliling, luas int

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&s)
		keliling = s * s
		luas = 4 * s

		fmt.Println(keliling, luas)
	}
}
