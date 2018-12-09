package main

import (
	"fmt"
	// "../lib"
)

func main() {
	var n int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= (2*n -1); j++ {
			if j <= n-i || j >= n+i {
			fmt.Printf(" ")
			} else {
				fmt.Printf("*")
			}
		}
		fmt.Printf("\n")
	}

	// soDauCach := n - 1
	// soDauCham := 1

	// for i := 0; i < n; i++ {
	// 	lib.DoPrint(" ", soDauCach)
	// 	lib.DoPrint("*", soDauCham)
	// 	fmt.Printf("\n")

	// 	soDauCach--
	// 	soDauCham = soDauCham + 2
	// }
}
