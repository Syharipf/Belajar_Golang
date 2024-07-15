package main

import "fmt"

func simple_calc(c *int, b int) {
	var a int

	a = 10 + b - *c
	*c = *c + 4

	fmt.Println(a)
}

func main() {
	var a, b, c int

	fmt.Scan(&a, &b, &c)
	simple_calc(&c, b)
	fmt.Println(a, b, c)
	fmt.Scan(b)
	simple_calc(&b, c)
	fmt.Println(a, b, c)
}
