package main

import "fmt"

func urut(a, b *int) {
	if *a > *b {
		*a, *b = *b, *a
	}
}

func main() {
	var a, b int
	for {
		fmt.Scan(&a, &b)

		if a == b {
			break
		}

		urut(&a, &b)
		fmt.Println(a, b)
	}
}
