package main

import (
	"fmt"
)

func main() {
	var ntest int

	fmt.Scanf("%d\n", &ntest)

	for ntest > 0 {
		ntest--

		var n int
		fmt.Scan(&n)

		slc := []int{1, 5, 10, 20, 50}
		rsl := doiTien(slc, n)

		for i, _ := range slc {
			fmt.Printf("(%d) %d ", slc[i], rsl[i])
		}
		fmt.Println()
	}
}

func doiTien(slc []int, n int) []int {
	rsl := []int{0, 0, 0, 0, 0}

	for i := len(slc) - 1; i >= 0; i-- {
		if slc[i] <= n {
			rsl[i] = n / slc[i]
			n = n % slc[i]
		}
	}
	return rsl
}
