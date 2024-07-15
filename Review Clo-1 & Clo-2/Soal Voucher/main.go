package main

import "fmt"

func main() {

	var (
		x, d1, d3, d4   int
		diskon, voucher bool
	)

	fmt.Scan(&x)
	d1 = x / 1000
	d3 = x / 10 % 10
	d4 = x % 10

	diskon = d3%2 == 0

	voucher = (d1+d3)%d4 == 0

	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", voucher)

}
