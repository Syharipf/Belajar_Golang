package main

import "fmt"

func main() {

	var kar rune

	fmt.Scanf("%c", &kar)

	if kar >= 'A' && kar <= 'Z' || kar >= 'a' && kar <= 'z' {
		fmt.Println("Alphabet")
	} else {
		fmt.Println("Bukan Alphabet")
	}
}
