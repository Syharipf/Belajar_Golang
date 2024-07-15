package main

import "fmt"

func main() {

	var (
		kar    rune
		status bool
	)

	fmt.Scanf("%c", &kar)

	status = (kar >= 'A' && kar <= 'Z') || (kar >= 'a' && kar <= 'z') || (kar >= '0' && kar <= '9')

	fmt.Println(status)

}
