package main

import "fmt"

func main() {

	var v, s, t, total int

	fmt.Scan(&s, &v, &t)

	total = v*t + s

	fmt.Println(total)
}
