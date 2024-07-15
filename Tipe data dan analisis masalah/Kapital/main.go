package main

import "fmt"

func main() {

	var (
		kar     rune
		kapital bool
	)

	fmt.Scanf("%c", &kar)

	kapital = kar >= 'A' && kar <= 'Z'

	fmt.Println(kapital)
}
