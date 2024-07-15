package main

import "fmt"

func main() {

	var (
		nilai        int
		buat, status bool
	)

	fmt.Scan(&nilai, &buat)

	status = (nilai > 55 && buat) || nilai > 90

	fmt.Println(status)
}
