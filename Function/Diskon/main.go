package main

import "fmt"

func discount(belanja int, member bool) int {

	var harga int
	if belanja > 100000 && member {
		harga = belanja - (belanja * 10 / 100)
	} else if belanja > 100000 && !member {
		harga = belanja - (belanja * 5 / 100)
	} else if belanja < 100000 {
		harga = belanja
	}
	return harga
}

func main() {

	var (
		belanja, total int
		member         bool
	)

	fmt.Scan(&belanja, &member)
	total = discount(belanja, member)
	fmt.Println(total)
}
