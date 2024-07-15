package main

import "fmt"

func main() {

	var jam, menit, detik int

	fmt.Scan(&detik)

	jam = detik / 3600
	detik = detik % 3600
	menit = detik / 60
	detik = detik % 60

	fmt.Printf("%d jam %d menit %d detik", jam, menit, detik)
}
