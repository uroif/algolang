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
		fmt.Scanf("%d\n", &n)

		slc := lib.InputSlice(n)
		sum := 0

		for i := 0; i < len(slc); i++ {
			for j := 0; j < len(slc); j++ {
				if i < j && isPrime(slc[i] + slc[j]) {
					sum++
				}
			} 
		}

		fmt.Println(sum)
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}

	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
