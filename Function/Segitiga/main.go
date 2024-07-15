package main

import "fmt"

func segitiga(a, b, c int) string {

	if a+b > c && a+c > b && b+c > a {
		return "Segitiga"
	} else {
		return "Bukan segitiga"
	}
}

func main() {

	var (
		a, b, c  int
		triangle string
	)

	fmt.Scan(&a, &b, &c)
	triangle = segitiga(a, b, c)
	fmt.Println(triangle)
}
