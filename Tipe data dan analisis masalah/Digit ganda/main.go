package main

import "fmt"

func main() {

	var x, d1, d2 int

	fmt.Scan(&x)

	d1 = x / 10
	d2 = x % 10

	fmt.Printf("%d%d%d", d1, x, d2)
}
