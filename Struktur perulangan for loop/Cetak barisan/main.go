package main

import "fmt"

func main() {

	var x, y int

	fmt.Scan(&x, &y)

	for i := x; i <= y; i++ {
		fmt.Printf("%d ", i)
	}
}
