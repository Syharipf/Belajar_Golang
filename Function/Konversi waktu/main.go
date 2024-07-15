package main

import "fmt"

func waktu(x, y, z int) int {
	var jam, menit, detik, total int
	jam = x * 3600
	menit = y * 60
	detik = z

	total = jam + menit + detik
	return total
}

func main() {

	var (
		x, y, z int
		time    int
	)

	fmt.Scan(&x, &y, &z)
	time = waktu(x, y, z)
	fmt.Printf("Hasil konversi = %d detik", time)
}
