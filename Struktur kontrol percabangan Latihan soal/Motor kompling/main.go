package main

import "fmt"

func main() {

	var (
		gigi                     int
		statusKopling, statusGas bool
	)

	fmt.Scan(&gigi, &statusKopling, &statusGas)

	if statusKopling || gigi == 0 {
		fmt.Println("Mesin menyala dan motor tidak berjalan")
	} else if !statusGas {
		fmt.Println("Mesin mati")
	} else {
		fmt.Println("Motor berjalan")
	}
}
