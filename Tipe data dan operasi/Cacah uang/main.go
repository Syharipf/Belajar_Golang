package main

import "fmt"

func main() {

	var x int

	fmt.Scan(&x)

	d1 := x / 10000
	x = x % 10000
	d2 := x / 5000
	x = x % 5000
	d3 := x / 1000

	fmt.Printf("%d lembar\n", d1)
	fmt.Printf("%d lembar \n", d2)
	fmt.Printf("%d lembar", d3)
}
