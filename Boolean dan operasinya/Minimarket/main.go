package main

import "fmt"

func main() {

	var (
		belanja                       int
		buat, kartu, diskon, cashback bool
	)

	fmt.Scan(&belanja, &buat)

	kartu = buat
	diskon = belanja >= 100000 && kartu
	cashback = belanja >= 200000 && kartu

	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
}
