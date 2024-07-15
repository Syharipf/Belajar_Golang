package main

import "fmt"

func main() {

	var (
		berat                                                             int
		bumi, merkurius, venus, mars, yupiter, saturnus, uranus, neptunus float64
	)
	fmt.Scan(&berat)

	bumi = float64(berat) * 9.8
	merkurius = float64(berat) * 9.8 * 0.38
	venus = float64(berat) * 9.8 * 0.91
	mars = float64(berat) * 9.8 * 0.38
	yupiter = float64(berat) * 9.8 * 2.37
	saturnus = float64(berat) * 9.8 * 0.92
	uranus = float64(berat) * 9.8 * 0.89
	neptunus = float64(berat) * 9.8 * 1.13

	fmt.Println(int(merkurius), int(venus), int(bumi), int(mars), int(yupiter), int(saturnus), int(uranus), int(neptunus))
}
