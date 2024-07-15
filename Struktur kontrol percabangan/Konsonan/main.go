package main

import "fmt"

func main() {

	var kar rune

	fmt.Scanf("%c", &kar)

	if (kar >= '0' && kar <= '9') || ((kar >= 'a' && kar <= 'z') || (kar >= 'A' && kar <= 'Z')) && (kar == 'a' || kar == 'A' || kar == 'i' || kar == 'I' || kar == 'u' || kar == 'U' || kar == 'e' || kar == 'E' || kar == 'o' || kar == 'O') {
		fmt.Println("Bukan Konsonan")
	} else if (kar >= 'a' && kar <= 'z') || (kar >= 'A' && kar <= 'Z') {
		fmt.Println("Konsonan")
	}
}
