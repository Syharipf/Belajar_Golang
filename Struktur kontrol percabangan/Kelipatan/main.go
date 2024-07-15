package main

import "fmt"

func main() {

	var x int

	fmt.Scan(&x)

	if x%5 == 0 && x%3 == 0 {
		fmt.Println("Kelipatan 5")
		fmt.Println("Kelipatan 3")
	} else if x%5 == 0 {
		fmt.Println("Kelipatan 5")
	} else if x%3 == 0 {
		fmt.Println("Kelipatan 3")
	} else {
		fmt.Println("")
	}
}
