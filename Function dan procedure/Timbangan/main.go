package main

import "fmt"

const pi = 3.14

func volume(r, t float64) float64 {
	return pi * r * r * t
}

func massa(r, t, m float64) float64 {
	return volume(r, t) * m
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else if m1 > m2 {
		fmt.Println(m1 - m2)
	} else {
		fmt.Println(m2 - m1)
	}
}

func main() {
	var (
		r, t1, m1, t2, m2 float64
		stor1, stor2      float64
	)

	fmt.Scan(&r)
	fmt.Scan(&t1, &m1)
	fmt.Scan(&t2, &m2)

	stor1 = massa(r, t1, m1)
	stor2 = massa(r, t2, m2)
	display(stor1, stor2)

}
