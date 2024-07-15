package main

import "fmt"

func volumeT(r, t int) float32 {
	var hasil float32

	hasil = 3.14 * float32(r*r) * float32(t)
	return hasil
}

func main() {

	var (
		r, t   int
		volume float32
	)

	fmt.Scan(&r, &t)
	volume = volumeT(r, t)

	fmt.Println(volume)

}
