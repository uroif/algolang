package main

import (
	"algolang/lib"
	"fmt"
)

func main() {
	var ntest, n, x int

	fmt.Scan(&ntest)

	for ntest > 0 {
		ntest--

		fmt.Scanf("%d %d\n", &n, &x)

		slc := lib.InputSlice(n)
		fmt.Println(search(slc, x))
	}
}

func search(slc []int, x int) []int] {
	var rsl []int
	for i, v := range slc {
		if v == x {
			rsl = append(rsl, i)
		}
	}
	return rsl
}