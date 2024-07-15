package main

import "fmt"

func main() {

	var x, y, div, mod int

	fmt.Scan(&x, &y)
	div = 0
	mod = x

	for mod >= y {
		mod = mod - y
		div = div + 1
	}
	fmt.Printf("%d mod %d = %d\n", x, y, mod)
	fmt.Printf("%d div %d = %d", x, y, div)
}
