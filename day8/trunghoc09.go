package main

import (
	"algolang/lib"
	"fmt"
)

func main() {
	var ntest int

	fmt.Scan(&ntest)

	for ntest > 0 {
		ntest--

		var shop, month int
		check := false

		fmt.Scan(&shop, &month)

		for i := 1; i <= shop; i++ {
			var slc []int

			slc = lib.InputSlice(month)

			if incIsDown3M(slc) {
				fmt.Printf("%d ", i)
				check = true
			}
		}

		if !check {
			fmt.Print("0")
		}
		fmt.Println()
	}
}

func incIsDown3M(slc []int) bool {
	var a int
	b := 0
	for i := 1; i < len(slc); i++ {
		a = slc[i] - slc[i-1]
		if a < 0 {
			b++
		} else {
			b = 0
		}
		if b == 3 {
			return true
		}
	}
	return false
}
