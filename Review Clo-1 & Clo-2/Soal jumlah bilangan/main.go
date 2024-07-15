package main

import "fmt"

func main() {

	var x, hasil, digit int

	fmt.Scan(&x)
	hasil = 0
	for x > 0 {
		digit = x % 10
		hasil += digit
		x /= 10
	}
	fmt.Println(hasil)
}
