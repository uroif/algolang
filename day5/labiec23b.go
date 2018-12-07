package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	fmt.Printf("%d", findMax(n))
}

func findMax(n int) int {
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
