package main

import "fmt"

func faktorial(n int) int {
	if n == 1 || n == 0 {
		return 1
	} else {
		return n * faktorial(n-1)
	}
}

func main() {
	var n int

	fmt.Scan(&n)
	fmt.Println(faktorial(n))

}
