package main

import (
	"algolang/lib"
	"fmt"
	"strings"
)

func main() {
	var ntest, n, x int

	fmt.Scan(&ntest)

	for ntest > 0 {
		ntest--

		fmt.Scan(&n, &x)

		slc := lib.InputSlice(n)
		if len(search(slc, x)) == 0 {
			fmt.Println("-1")
		} else {
			slc = search(slc, x)
			fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(slc), " ", " ", -1), "[]"))
		}
	}
}

func search(slc []int, x int) []int {
	var rsl []int
	for i, v := range slc {
		if v == x {
			rsl = append(rsl, i+1)
		}
	}
	return rsl
}
