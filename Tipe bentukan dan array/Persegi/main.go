package main

import "fmt"

type geometry struct {
	area, perimeter int
}
type rectangle struct {
	length, width int
	color         string
	property      geometry
}

func isiData(persegi *rectangle) {
	fmt.Print("Masukkan panjang dan lebar: ")
	fmt.Scan(&persegi.length, &persegi.width)
	fmt.Print("Masukkan warna: ")
	fmt.Scan(&persegi.color)
}

func hitung(persegi *rectangle) {
	persegi.property.area = persegi.length * persegi.width
	persegi.property.perimeter = 2 * (persegi.length + persegi.width)
}

func main() {
	var persegi rectangle

	isiData(&persegi)
	hitung(&persegi)

	fmt.Printf("Luas persegi panjang: %d\n", persegi.property.area)
	fmt.Printf("Keliling persegi panjang: %d\n", persegi.property.perimeter)
}
