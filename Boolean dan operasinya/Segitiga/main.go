package main

import "fmt"

func main() {

	var (
		a, b, c int
		status  bool
	)

	fmt.Scan(&a, &b, &c)

	status = a+b > c && a+c > b && b+c > a

	fmt.Println(status)

}
