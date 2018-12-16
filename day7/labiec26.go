package main

import (
	"fmt"
	lib "../lib"
)

func main() {
	var ntest int

	fmt.Scanf("%d\n", &ntest)

	for ntest > 0 {
		ntest--

		var n int
		fmt.Scan(&n)

		slc := lib.InputSlice(n)

		fmt.Println(isDayDanDau(slc))
	}
}

func isDayDanDau(slc []int) int {
	for i := 1; i < len(slc); i++ {
		if slc[i]*slc[i-1] >= 0 {
			return 0
		}
	}
	return 1
}
