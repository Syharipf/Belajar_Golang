package main

import "fmt"

func main() {

	var (
		n            int
		item, status bool
	)

	fmt.Scan(&n)
	status = true
	for i := 1; i <= n; i++ {
		for j := 0; j < 5; j++ {
			fmt.Scan(&item)
			if !item {
				status = false
			}
		}
	}
	fmt.Println(status)
}
