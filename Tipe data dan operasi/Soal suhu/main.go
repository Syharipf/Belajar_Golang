package main

import "fmt"

func main() {

	var c, r, f, k float32

	fmt.Scan(&c)

	r = c * 4.0 / 5.0
	f = (c * 9.0 / 5.0) + 32.0
	k = c + 273.15

	fmt.Println("Reamur     :", r)
	fmt.Println("Fahrenheit :", f)
	fmt.Println("Kelvin     :", k)

}
