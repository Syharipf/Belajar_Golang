package main

import "fmt"

type balok struct {
	p, l, t, volume, luas int
}

func luasSelimut(x *balok) {
	x.luas = 2 * (x.p*x.l + x.l*x.t)
}

func volume(x *balok) {
	x.volume = x.p * x.l * x.t
}

func main() {
	var x balok

	fmt.Scan(&x.p, &x.l, &x.t)
	luasSelimut(&x)
	volume(&x)

	fmt.Println(x.luas, x.volume)

}
