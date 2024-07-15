package main

import "fmt"

func main() {

	var (
		belanjaAw, belanjaAk          int
		buat, kartu, diskon, cashback bool
	)

	fmt.Scan(&belanjaAw, buat)

	kartu = buat == true
	diskon = kartu && belanjaAw >= 100000
	cashback = kartu && belanjaAw >= 200000

	belanjaAk = 0
	if diskon && cashback {
		belanjaAk = belanjaAw - (belanjaAw * 10 / 100) - 75000
	} else if diskon {
		belanjaAk = belanjaAw - (belanjaAw * 10 / 100)
	}

	fmt.Println("Kartu?", kartu)
	fmt.Println("Diskon?", diskon)
	fmt.Println("Cashback?", cashback)
	fmt.Println(belanjaAk)
}
