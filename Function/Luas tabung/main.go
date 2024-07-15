package main

import "fmt"

func luasTabung(r, t int) float32 {
	var hasil float32
	hasil = 2.0 * 3.14 * float32(r) * float32(r+t)
	return hasil
}

func main() {

	var (
		r, t int
		luas float32
	)

	fmt.Scan(&r, &t)
	luas = luasTabung(r, t)
	fmt.Println(luas)
}
