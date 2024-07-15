package main

import "fmt"

func Fahrenheit(celcius float32) float32 {
	var suhu float32
	suhu = (9 * celcius / 5) + 32
	return suhu
}

func main() {

	var (
		i, n          int
		celcius, suhu float32
	)

	fmt.Scan(&n)

	i = 1
	for i <= n {
		fmt.Scan(&celcius)
		suhu = Fahrenheit(celcius)
		fmt.Printf("%.2f Celcius = %.2f Fahrenheit\n", celcius, suhu)

		i += 1
	}
}
