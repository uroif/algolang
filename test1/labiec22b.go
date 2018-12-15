package main

import (
	"fmt"
)

func main() {
	var ntest, n int

	fmt.Scanf("%d\n", &ntest)

	for ntest > 0 {
		ntest--

		fmt.Scanf("%d\n", &n)
		fmt.Printf("%d\n", timSoLon2(n))
	}
}

func timSoLon2(n int) int {
	var x int
	max1 := 0
	max2 := 0

	if n <= 9 {
		return n
	}

	for n > 0 {
		x = n % 10
		n = n / 10
		if x > max1 {
			max2 = max1
			max1 = x
		} else if x < max1 && x > max2 || x == max1 && max2 == 0 {
			max2 = x
		}
	}
	max2 = findMax2(x, max1, max2)

	return max2
}

func find2Max(a, b, c int) (int, int) {
	if a <= b && a <= c && b != c || a == b {
		return b, c
	} else if b <= a && b <= c && a != c {
		return a, c
	}
	return a, b
}

func findMax2(a, b, c int) int {
	x, y := find2Max(a, b, c)

	if x >= y {
		return y
	}
	return x
}
