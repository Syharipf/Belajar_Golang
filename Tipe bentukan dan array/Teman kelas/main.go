package main

import "fmt"

type teman [100]string

func tulis(a *teman, m int) {
	fmt.Println("Masukan nama teman kelas anda: ")
	for i := 0; i <= m-1; i++ {
		fmt.Scan(&a[i])
	}
}

func cetak(a teman, m int) {
	fmt.Println("Daftar teman kelas anda: ")
	for i := 0; i <= m-1; i++ {
		fmt.Println(a[i])
	}
}

func main() {
	var (
		t teman
		n int
	)
	fmt.Print("Masukan jumlah teman kelas anda: ")
	fmt.Scan(&n)
	tulis(&t, n)
	cetak(t, n)
}
