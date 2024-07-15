package main

import "fmt"

func main() {

	var (
		jabatan                                       string
		masaKerja, jumlahAnak, gaji, tunjangan, total int
	)

	fmt.Scan(&jabatan, &masaKerja, &jumlahAnak)

	if jumlahAnak > 3 {
		jumlahAnak = 3
	}
	if jabatan == "staf" {
		if masaKerja < 5 {
			gaji = 4000
			tunjangan = jumlahAnak * 0
			total = gaji + tunjangan
			fmt.Printf("%d + %d = %d", gaji, tunjangan, total)
		}
		if masaKerja >= 5 && masaKerja <= 10 {
			gaji = 4000
			tunjangan = jumlahAnak * 100
			total = gaji + tunjangan
			fmt.Printf("%d + %d = %d", gaji, tunjangan, total)
		}
		if masaKerja > 10 {
			gaji = 5000
			tunjangan = jumlahAnak * 100
			total = gaji + tunjangan
			fmt.Printf("%d + %d = %d", gaji, tunjangan, total)
		}
	}
	if jabatan == "manager" {
		if masaKerja > 10 {
			gaji = 10000
			tunjangan = jumlahAnak * 300
			total = gaji + tunjangan
			fmt.Printf("%d + %d = %d", gaji, tunjangan, total)
		}
		if masaKerja <= 10 {
			gaji = 8500
			tunjangan = jumlahAnak * 300
			total = gaji + tunjangan
			fmt.Printf("%d + %d = %d", gaji, tunjangan, total)

		}
	}
	if jabatan == "direktur" {
		gaji = 20000
		tunjangan = jumlahAnak * 500
		total = gaji + tunjangan
		fmt.Printf("%d + %d = %d", gaji, tunjangan, total)
	}
}
