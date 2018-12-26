package main

import (
	"algolang/lib"
	"fmt"
	"sort"
	"strings"
)

func main() {
	var ntest, n int

	fmt.Scan(&ntest)

	for ntest > 0 {
		ntest--

		fmt.Scan(&n)
		slc := lib.InputSlice(n)
		sort.Ints(slc)

		fmt.Println(strings.Trim(strings.Replace(fmt.Sprint(slc), " ", " ", -1), "[]"))
	}
}