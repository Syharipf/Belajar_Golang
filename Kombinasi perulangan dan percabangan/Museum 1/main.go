package main

import "fmt"

func main() {

	var (
		x, hari, totalPengunjungTiapHari, data, peningkatan int
		rataRata                                            float32
	)

	data = 0
	peningkatan = 0
	totalPengunjungTiapHari = 0
	hari = 0
	for {
		hari += 1
		totalPengunjungTiapHari += x
		fmt.Scan(&x)
		data = x
		if data < x {
			peningkatan += 1
		}
		if (x < 0) || (x > 200) {
			break
		}
	}
	rataRata = float32(totalPengunjungTiapHari) / float32(hari)
	fmt.Print(peningkatan, rataRata)
}
