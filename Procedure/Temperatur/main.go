package main

import "fmt"

func temp(reamur, fahrenheit, kelvin *float32, celcius float32) {

	*reamur = (4.0 / 5.0) * celcius
	*fahrenheit = (9.0/5.0)*celcius + 32
	*kelvin = celcius + 273.15

}
func main() {
	var c, r, f, k float32

	fmt.Scan(&c)
	temp(&r, &f, &k, c)
	fmt.Printf("%.2fR %.2fF %.2fK", r, f, k)
}
