package main

import (
	"fmt"
	"math"
)

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))

}

func diDalam(x, y, cx, cy, r float64) bool {
	return jarak(x, y, cx, cy) <= r
}

func akar(x float64) float64 {
	return math.Sqrt(x)
}

func posisiTitikLingkaran(cx1, cy1, r1, cx2, cy2, r2, x, y float64) string {
	diDalamLingkaran1 := diDalam(x, y, cx1, cy1, r1)
	diDalamLingkaran2 := diDalam(x, y, cx2, cy2, r2)

	if diDalamLingkaran1 && diDalamLingkaran2 {
		return "Titik di dalam lingkaran 1 dan 2"
	} else if diDalamLingkaran1 {
		return "Titik di dalam lingkaran 1"
	} else if diDalamLingkaran2 {
		return "Titik di dalam lingkaran 2"
	} else {
		return "Titik di luar lingkaran 1 dan 2"
	}
}

func main() {
	var (
		cx1, cy1, r1, cx2, cy2, r2, x, y float64
		hasil                            string
	)

	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	hasil = posisiTitikLingkaran(cx1, cy1, r1, cx2, cy2, r2, x, y)

	fmt.Println(hasil)

}
