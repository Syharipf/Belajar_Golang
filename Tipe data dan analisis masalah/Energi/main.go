package main

import "fmt"

func main() {

	var e0, e1, c, cc int

	fmt.Scan(&e0, &c, &e1)

	cc = (e0 - e1) / c

	fmt.Println(cc)
}
