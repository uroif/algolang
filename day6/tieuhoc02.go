package main

import (
	"fmt"
	"math"

	lib "../lib"
)

func main() {
	var ntest int

	fmt.Scan(&ntest)

	for ntest > 0 {
		ntest--

		var n int
		fmt.Scan(&n)

		slc := lib.InputSlice(n)
		// lib.PrintSlice(slc)
		fmt.Printf("%f %f", tinhTBC(slc), tinhTBN(slc))
		fmt.Println()
	}
}

func tinhTBC(slc []int) float64 {
	sum := 0
	for i := 0; i < len(slc); i++ {
		sum = sum + slc[i]
	}
	tbc := float64(sum) / float64(len(slc))
	return tbc
}

func tinhTBN(slc []int) float64 {
	product := 1
	for _, value := range slc {
		product = product * value
	}
	tbn := math.Pow(float64(product),1.0/float64(len(slc)))
	return tbn
}
