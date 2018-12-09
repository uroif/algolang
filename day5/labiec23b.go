package main

import "fmt"

func main() {
	var ntest, x int

	fmt.Scanf("%d\n", &ntest)

	for ntest > 0 {
		ntest--
		
		fmt.Scan(&x)
		fmt.Printf("%d\n", findMaxDigit(x))
	}
}

func findMaxDigit(n int) int {
	var x, max int

	for n != 0 {
		x = n % 10
		n = n / 10
		if x > max {
			max = x
		}
	}
	return max
}
