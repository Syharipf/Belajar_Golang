package main

import "fmt"

func main() {

	var (
		x, y, z int
		status  bool
	)

	fmt.Scan(&x, &y, &z)

	status = float32((x+y))/2 == float32(z) || float32((x+z))/2 == float32(y) || float32((y+z))/2 == float32(x)

	fmt.Println(status)
}
