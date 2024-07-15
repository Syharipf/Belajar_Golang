package main

import "fmt"

func main() {

	var (
		barang1, barang2, barang3 int
		harga1, harga2, harga3    float64
	)

	fmt.Scan(&barang1, &barang2, &barang3)

	harga1 = float64(barang1) + float64(barang1)*0.05
	harga2 = float64(barang2) + float64(barang2)*0.05
	harga3 = float64(barang3) + float64(barang3)*0.05

	fmt.Println(harga1, harga2, harga3)

}
