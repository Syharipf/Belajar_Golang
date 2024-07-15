package main

import "fmt"

func main() {

	var (
		jam, menit, jarak, tarif float64
	)

	fmt.Scan(&jam, &menit, &jarak)

	if (jam < 05 && jam > 22) || (jarak > 20) || (jam > 22) {
		fmt.Println("Maaf, kami tidak bisa melayani anda")
	} else {
		if (jam >= 5 && jam <= 9) || (jam >= 16 && jam <= 19) {
			if jarak > 0 && jarak <= 10 {
				tarif = jarak * 5000
				fmt.Println(int(tarif))
			} else if jarak > 10 && jarak <= 20 {
				tarif = jarak * 4500
				fmt.Println(int(tarif))
			}
		} else if (jam > 9 && jam < 16) || (jam > 19 && jam <= 22) {
			tarif = jarak * 4000
			fmt.Println(int(tarif))

		}
	}
}
