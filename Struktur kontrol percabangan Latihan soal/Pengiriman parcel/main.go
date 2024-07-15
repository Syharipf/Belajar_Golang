package main

import "fmt"

func main() {

	var berat, kg, gr, biaya, biayaTambahan1, biayaTambahan2 int

	fmt.Scan(&berat)

	kg = berat / 1000
	gr = berat % 1000
	biaya = kg * 10000
	biayaTambahan1 = gr * 5
	biayaTambahan2 = gr * 15
	if gr != 0 && gr < 500 {
		fmt.Printf("%d kg + %d gr\n", kg, gr)
		fmt.Printf("Rp %d + Rp %d = Rp %d ", biaya, biayaTambahan2, biaya+biayaTambahan2)
	} else if gr != 0 && (gr >= 500 && gr <= 999) {
		fmt.Printf("%d kg + %d gr\n", kg, gr)
		fmt.Printf("Rp %d + Rp %d = Rp %d ", biaya, biayaTambahan1, biaya+biayaTambahan1)
	} else {
		fmt.Printf("%d kg + %d gr\n", kg, gr)
		fmt.Printf("Rp %d + Rp 0 = Rp %d ", biaya, biaya)
	}
}
