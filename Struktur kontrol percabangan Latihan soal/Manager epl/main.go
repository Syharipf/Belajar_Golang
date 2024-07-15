package main

import "fmt"

func main() {

	var (
		pertandingan1, pertandingan2, pertandingan3, pertandingan4, pertandingan5 string
		jumlahKekalahan                                                           int
	)

	fmt.Scan(&pertandingan1, &pertandingan2, &pertandingan3, &pertandingan4, &pertandingan5)

	jumlahKekalahan = 0

	if pertandingan1 == "kalah" {
		jumlahKekalahan++
	} else {
		jumlahKekalahan = 0
	}
	if pertandingan2 == "kalah" {
		jumlahKekalahan++
	} else {
		jumlahKekalahan = 0
	}

	if pertandingan3 == "kalah" {
		jumlahKekalahan++
	} else {
		jumlahKekalahan = 0
	}

	if pertandingan4 == "kalah" {
		jumlahKekalahan++
	} else {
		jumlahKekalahan = 0
	}

	if pertandingan5 == "kalah" {
		jumlahKekalahan++
	}

	if jumlahKekalahan == 5 {
		fmt.Println("Dipecat")
	} else {
		fmt.Println("Tidak dipecat")
	}
}
