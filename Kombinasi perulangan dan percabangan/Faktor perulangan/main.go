package main

import "fmt"

func main() {

	var (
		n, m, minimum int
	)

	fmt.Scan(&n, &m)

	minimum = n
	if n < m {
		minimum = m
	}

	for i := 1; i <= minimum; i++ {
		if n%i == 0 && m%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}
