package main

import "fmt"

func main() {

	var (
		member           string
		a, b, c, d, e    int
		cashback, diskon float64
		ganjil, genap    bool
	)

	fmt.Scan(&member, &a, &b, &c, &d, &e)

	genap = (a%2 == 0) && (b%2 == 0) && (c%2 == 0) && (d%2 == 0) && (e%2 == 0)
	ganjil = (a%2 != 0) && (b%2 != 0) && (c%2 != 0) && (d%2 != 0) && (e%2 != 0)

	if genap {
		diskon = float64(d+e) * 1.7
	} else if ganjil {
		cashback = float64(d+e) * 3.1
	} else {
		cashback = float64(a+b+c) * 3.1
		diskon = float64(d+e) * 1.7
	}

	if member == "yes" {
		cashback *= 1.15
		diskon *= 1.15
	}

	if cashback > 35 {
		cashback = 35
	}

	if diskon > 50 {
		diskon = 50
	}

	fmt.Printf("Cashback : %.2f, Diskon : %.2f", cashback, diskon)
}
