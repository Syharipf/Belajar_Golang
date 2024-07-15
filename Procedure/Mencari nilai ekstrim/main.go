package main

import "fmt"

func mencari(min, max *int, n1, n2 int) {

	if n1 > n2 {
		*min = n2
		*max = n1
	} else {
		*min = n1
		*max = n2
	}
}

func main() {

	var (
		a, b, c, d                             int
		min1, min2, max1, max2, min, max, temp int
	)

	fmt.Scan(&a, &b, &c, &d)
	mencari(&min1, &max1, a, b)
	mencari(&min2, &max2, c, d)
	mencari(&min, &temp, min1, min2)
	mencari(&temp, &max, max1, max2)
	fmt.Println(max, min)
}
