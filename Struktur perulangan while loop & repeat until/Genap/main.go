package main

import "fmt"

func main() {

	var x, total int

	total = 0
	for x%2 == 0 {
		total += x
		fmt.Scan(&x)

	}
	fmt.Println(total)
}
