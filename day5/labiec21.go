package main

import (
	"fmt"
	"log"
)

func main() {
	var ntest, x, n, min int

	fmt.Scanf("%d\n", &ntest)

	for ntest > 0 {
		ntest--

		fmt.Scanf("%d\n", &n)
		fmt.Scan(&x)

		min = x
		k := 1
		for i := 2; i <= n; i++ {
			fmt.Scan(&x)
			if x < min {
				min = x
				k = 0
			}
			if x == min {
				k++
			}
		}
		fmt.Printf("%d %d\n", min, k)
	}
}
