package main

import "fmt"

func main() {

	var (
		n, m, minimum, fpb int
	)

	fmt.Scan(&n, &m)

	minimum = m
	if n < m {
		minimum = n
	}

	for i := 1; i <= minimum; i++ {
		if n%i == 0 && m%i == 0 {
			fpb = i
		}
	}
	fmt.Println(fpb)
}
