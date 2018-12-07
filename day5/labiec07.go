package main

import (
	"fmt"
	"../lib"
)

func main() {
	var n int

	fmt.Scan(&n)

	soDauCach := n - 1
	soDauCham := 1

	for i := 0; i < n; i++ {
		lib.DoPrint(" ", soDauCach)
		lib.DoPrint("*", soDauCham)
		fmt.Printf("\n")

		soDauCach--
		soDauCham = soDauCham + 2
	}
}
