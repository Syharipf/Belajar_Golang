package main

import "fmt"

func main() {

	var (
		x, digitAk, digitAw int
		status              bool
	)

	fmt.Scan(&x)
	status = true
	digitAk = x % 10
	x /= 10

	for x > 0 && status {
		digitAw = x % 10
		if digitAk-digitAw != 1 && digitAk-digitAw != -1 {
			status = false
		}

		digitAk = digitAw
		x /= 10

	}
	fmt.Println(status)
}
