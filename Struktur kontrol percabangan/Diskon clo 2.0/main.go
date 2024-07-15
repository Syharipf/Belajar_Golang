package main

import "fmt"

func main() {

	var (
		belanja, potongan float64
		asesmen           bool
	)

	fmt.Scan(&belanja, &asesmen)
	if asesmen {
		potongan = belanja * 0.35
	} else {
		potongan = 0
	}

	fmt.Println(belanja - potongan)
}
