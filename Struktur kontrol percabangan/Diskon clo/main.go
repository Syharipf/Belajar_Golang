package main

import "fmt"

func main() {

	var (
		bAwal, bAkhir  int
		diskon, status bool
	)

	fmt.Scan(&bAwal, &status)
	diskon = true

	if diskon == status {
		bAkhir = bAwal - (bAwal * 35 / 100)
		fmt.Println(bAkhir)
	} else {
		fmt.Println(bAwal)
	}

}
