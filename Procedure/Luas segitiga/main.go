package main

import "fmt"

func hitungLuas(Luas *float32, alas, tinggi float32) {

	*Luas = alas * tinggi / 2
}

func main() {
	var (
		n     int
		a, t  float32
		hasil float32
	)

	fmt.Scan(&n)
	for i := 1; i <= n; i++ {
		fmt.Scan(&a, &t)
		hitungLuas(&hasil, a, t)
		fmt.Println(hasil)
	}
}
