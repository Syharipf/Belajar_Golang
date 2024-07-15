package main

import "fmt"

func main() {

	var (
		username, password, x, y string
		total                    int
	)

	username = "admin"
	password = "admin"
	total = 0

	for username != x || password != y {
		total += 1
		fmt.Scan(&x, &y)
	}
	fmt.Println(total)
	fmt.Println("Login berhasil")
}
