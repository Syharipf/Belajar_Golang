package main

import (
	"fmt"
	"math"
)

type titik struct {
	x float64
	y float64
}

func jarak(p1, p2 titik) float64 {
	return akar(pangkat(p1.x-p2.x) + pangkat(p1.y-p2.y))
}

func akar(x float64) float64 {
	return math.Sqrt(x)
}

func pangkat(x float64) float64 {
	return math.Pow(x, 2)
}

func main() {
	var (
		p1, p2     titik
		a, b, c, d float64
		hasil      float64
	)
	fmt.Print("Masukan koordinat titik x1 & y1: ")
	fmt.Scan(&a, &b)
	fmt.Print("Masukan koordinat titik x2 & y2: ")
	fmt.Scan(&c, &d)

	p1 = titik{x: a, y: b}
	p2 = titik{x: c, y: d}

	hasil = jarak(p1, p2)
	fmt.Printf("Jarak antara titik (%.2f, %.2f) & (%.2f, %.2f) adalah %.2f", a, b, c, d, hasil)
}
