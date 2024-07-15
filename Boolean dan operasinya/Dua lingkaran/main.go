package main

import "fmt"

func main() {

	var (
		A, B, tPusat int
		irisan       bool
	)

	fmt.Scan(&A, &B, &tPusat)

	irisan = (A + B) <= tPusat

	fmt.Println(irisan)
}
