package main

import "fmt"

func main() {

	var x, total int

	total = 0
	for {
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		total += x
	}
	fmt.Println(total)
}
