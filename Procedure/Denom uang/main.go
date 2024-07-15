package main

import "fmt"

func denom(k10, k2, k1 *int, uang int) {
	*k10 = uang / 10000
	uang = uang % 10000
	*k2 = uang / 2000
	uang = uang % 2000
	*k1 = uang / 1000
}

func main() {
	var l1, l2, l3, uang int

	fmt.Scan(&uang)
	denom(&l1, &l2, &l3, uang)
	fmt.Printf("%d lembar 10000\n", l1)
	fmt.Printf("%d lembar 2000\n", l2)
	fmt.Printf("%d lembar 100", l3)
}
