package main

import "fmt"

func main() {

	var nilai float32

	fmt.Scan(&nilai)

	if nilai < 2.0 {
		fmt.Println("Poor")
	} else if nilai >= 2.0 && nilai <= 2.75 {
		fmt.Println("Fair")
	} else if nilai >= 2.76 && nilai <= 3.00 {
		fmt.Println("Satisfactory")
	} else if nilai >= 3.01 && nilai <= 3.50 {
		fmt.Println("Very Good")
	} else if nilai > 3.50 {
		fmt.Println("Excellent")
	}
}
