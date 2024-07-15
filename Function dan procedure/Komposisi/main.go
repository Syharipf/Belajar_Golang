package main

import (
	"fmt"
)

func f(x int) int {
	return x * x
}

func g(x int) int {
	return x - 2
}

func h(x int) int {
	return x + 1
}

func fungsi(a, b, c int) (int, int, int) {

	var fogoh, gohof, hofog int

	fogoh = f(g(h(a)))
	gohof = g(h(f(b)))
	hofog = h(f(g(c)))

	return fogoh, gohof, hofog
}

func main() {
	var (
		a, b, c    int
		f1, f2, f3 int
	)

	fmt.Scan(&a, &b, &c)

	f1, f2, f3 = fungsi(a, b, c)

	fmt.Printf("(fogoh) (%d) = %d\n", a, f1)
	fmt.Printf("(gohof) (%d) = %d\n", b, f2)
	fmt.Printf("(hofog) (%d) = %d", c, f3)
}
