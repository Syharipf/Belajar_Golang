package main

import "fmt"

func main() {

	var (
		usia int
		KK   bool
	)

	fmt.Scan(&usia, &KK)

	if usia >= 17 && KK {
		fmt.Println("Bisa membuat KTP")
	} else {
		fmt.Println("Belum bisa membuat KTP")
	}
}
