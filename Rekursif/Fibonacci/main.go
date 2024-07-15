package main

import "fmt"

func fibo(n int) int {
	if n <= 1 {
		return n
	} else {
		return fibo(n-1) + fibo(n-2)
	}
}

func main() {
	var n int
	fmt.Scan(&n)
	fmt.Println(fibo(n))
}
