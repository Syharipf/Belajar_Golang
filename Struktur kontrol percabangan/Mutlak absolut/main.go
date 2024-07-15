package main

import "fmt"

func main() {

	var x, hasil int

	fmt.Scan(&x)

	if x < 0 {
		hasil = x * -1
		fmt.Println(hasil)
	} else {
		fmt.Println(x)
	}
}
