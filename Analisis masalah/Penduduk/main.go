package main

import "fmt"

func main() {

	var pAawal, jLahir, jImigrasi, jKematian, jEmigrasi, jAkhir int

	fmt.Scan(&pAawal, &jLahir, &jImigrasi, &jKematian, &jEmigrasi)

	jAkhir = pAawal + jLahir + jImigrasi - jKematian - jEmigrasi

	fmt.Println(jAkhir)
}
