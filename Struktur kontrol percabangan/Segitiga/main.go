package main

import "fmt"

func main() {

	var a, b, c int

	fmt.Scan(&a, &b, &c)

	if a == 0 || b == 0 || c == 0 {
		fmt.Println("")
	} else if a == b && a == c && b == c {
		fmt.Println("Segitiga Sama Sisi")
	} else if a == b || a == c || b == c {
		fmt.Println("Segitiga Sama kaki")
	} else {
		fmt.Println("Segitiga Sembarang")
	}
}
