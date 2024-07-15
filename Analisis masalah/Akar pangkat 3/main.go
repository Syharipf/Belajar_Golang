package main

import "fmt"

func main() {

	var (
		x, y   int
		status bool
	)

	fmt.Scan(&x, &y)

	status = x*x*x == y

	fmt.Println(status)
}
