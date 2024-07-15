package main

import "fmt"

func main() {

	var n, iterasi int

	fmt.Scan(&n)
	iterasi = 0
	for iterasi != n {
		iterasi += 1
		fmt.Print(iterasi, " ")
	}
}
