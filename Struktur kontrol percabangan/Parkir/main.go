package main

import "fmt"

func main() {

	var h1, m1, h2, m2, menitAw, menitAk, total, hh, mm int

	fmt.Scan(&h1, &m1, &h2, &m2)

	menitAw = h1*60 + m1
	menitAk = h2*60 + m2
	total = menitAk - menitAw

	if total < 0 {
		total += 12 * 60
	}

	hh = total / 60
	mm = total % 60

	fmt.Printf("%d jam %d menit", hh, mm)

}
