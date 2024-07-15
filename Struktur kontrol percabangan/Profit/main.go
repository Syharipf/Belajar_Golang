package main

import "fmt"

func main() {

	var x, y, hasil1, hasil2 float32

	fmt.Scan(&x, &y)

	if x > y {
		hasil1 = x - y
		fmt.Printf("Penurunan sebesar %d", int(hasil1))
	} else if x < y {
		hasil2 = y - x
		fmt.Printf("Peningkatan sebesar %d", int(hasil2))
	} else if x == y {
		fmt.Println("Tetap")
	}
}
