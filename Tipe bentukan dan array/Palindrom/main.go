package main

import "fmt"

type Capacity [256]int

func isiData(a *Capacity, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan elemen ke-%d: ", i+1)
		fmt.Scan(&a[i])
	}
}

func reverseArray(a *Capacity, n int) {
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func isiPolindrom(a Capacity, n int) bool {
	for i := 0; i < n/2; i++ {
		if a[i] != a[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var (
		a Capacity
		n int
	)

	fmt.Print("Masukkan jumlah elemen (maksimal 256): ")
	fmt.Scan(&n)

	if n > 256 || n < 1 {
		fmt.Println("Jumlah elemen harus antara 1 dan 256")
		return
	}

	isiData(&a, n)

	fmt.Println("Array setelah diisi:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()

	reverseArray(&a, n)

	fmt.Println("Array setelah dibalik:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", a[i])
	}
	fmt.Println()

	if isiPolindrom(a, n) {
		fmt.Println("Array memiliki pola palindrom.")
	} else {
		fmt.Println("Array tidak memiliki pola palindrom.")
	}
}
