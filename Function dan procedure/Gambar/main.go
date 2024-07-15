package main

import "fmt"

func zoomIn(p, l, n int) {
	var pA, lA int
	pA = p * n
	lA = l * n

	fmt.Println(pA, lA)
}

func zoomOut(p, l, n int) {
	var pA, lA int
	pA = p / n
	lA = l / n

	fmt.Println(pA, lA)
}

func main() {
	var (
		p, l, n int
		s       string
	)

	fmt.Scan(&p, &l)
	fmt.Scan(&s, &n)

	if s == "+" {
		zoomIn(p, l, n)
	} else if s == "-" {
		zoomOut(p, l, n)
	}
}
