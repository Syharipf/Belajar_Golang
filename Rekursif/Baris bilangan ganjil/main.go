package main

import "fmt"

func baris(n, current int) {
	if current > n {
		return
	}
	if current%2 != 0 {
		fmt.Printf("%d ", current)
	}
	baris(n, current+1)
}

func main() {
	var n int

	fmt.Scan(&n)
	fmt.Println(n)
	baris(n, 1)
}
