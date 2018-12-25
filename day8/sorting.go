package main

import (
	"algolang/lib"
	"fmt"
)

func main() {
	var ntest, n int

	fmt.Scan(&ntest)

	for ntest > 0 {
		ntest--

		fmt.Scan(&n)
		slc := lib.InputSlice(n)
		// fmt.Println(slc)
		fmt.Println(sortIt(slc))
	}
}

func sortIt(slc []int) []int {
	length := len(slc)
	var newSlc []int
	var x int

	for i := 0; i < length; i++ {
		x, slc = findMin(slc)
		newSlc = append(newSlc, x)
		// log.Println(slc)
	}
	return newSlc
}

func findMin(slc []int) (int, []int) {
	var min int
	length := len(slc)

	for i := 1; i < length; i++ {
		if slc[i] >= slc[i-1] {
			min = slc[i-1]
		} else {
			min = slc[i]
		}
	}

	for i := 0; i < length; i++ {
		if min == slc[i] {
			slc[i] = slc[length-1]
			slc[length - 1] = 0
			slc = slc[:length-1]
		}
	}

	return min, slc
}
