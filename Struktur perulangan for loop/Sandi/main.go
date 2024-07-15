package main

import "fmt"

func main() {

	var x, y, d1, d2, total int

	fmt.Scan(&x)
	total = 0
	for i := 1; i <= x; i++ {
		fmt.Scan(&y)
		d1 = y / 1000
		d2 = y % 10
		total += d1 + d2
	}
	fmt.Println(total)
}
