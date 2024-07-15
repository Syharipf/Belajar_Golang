package main

import "fmt"

func main() {

	var (
		x   int
		kar string
	)

	fmt.Scan(&x)
	fmt.Scan(&kar)
	for i := 1; i <= x; i++ {
		fmt.Println(kar)
	}
}
