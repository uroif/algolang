package main

import (
	"fmt"
)

func main() {
	var n int

	for {
		status, _ := fmt.Scan(&n)
		if status == 0 {
			break
		}

		for i := 1; i < n; i++ {
			if isPerfectNumber(i) {
				fmt.Printf("%d ", i)
			}
		}
		fmt.Printf("\n")
	}

}

func isPerfectNumber(n int) bool {
	sum := 0

	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum = sum + i
		}
	}

	if sum == n {
		return true
	}
	return false
}
