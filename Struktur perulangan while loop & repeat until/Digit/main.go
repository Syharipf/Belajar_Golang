package main

import "fmt"

func main() {

	var x, total, digit int

	fmt.Scan(&x)
	total = 0
	for x > 0 {
		digit = x % 10
		fmt.Print(digit, " ")
		total += digit
		x /= 10
	}

	fmt.Println(total)
}
