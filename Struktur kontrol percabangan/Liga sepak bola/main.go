package main

import "fmt"

func main() {

	var g1, g2, g3, g4, max, min int

	fmt.Scan(&g1, &g2, &g3, &g4)

	max = g1
	if g2 > max {
		max = g2
	} else if g3 > max {
		max = g3
	} else if g4 > max {
		max = g4
	}

	min = g1
	if g2 < min {
		min = g2
	} else if g3 < min {
		min = g3
	} else if g4 < min {
		min = g4
	}

	fmt.Println(max, min)
}
