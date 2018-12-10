package main

import "fmt"

func main() {
	var ntest, n int

	fmt.Scanf("%d\n", &ntest)

	for ntest > 0 {
		ntest--

		fmt.Scanf("%d\n", &n)
		sumDigit := 0
		
		for n != 0 {
			sumDigit = sumDigit + n % 10
			n = n / 10
		}

		fmt.Printf("%d\n", sumDigit)
	}
}