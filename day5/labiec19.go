package main

import "fmt"

func main() {
	var ntest int

	var a, b int
	fmt.Scanf("%d\n", &ntest)

	for ntest > 0 {
		ntest--

		fmt.Scanf("%d %d", &a, &b)
		fmt.Println(findAndSum(a, b))
	}
}

func findAndSum(a, b int) (int, int) {
	n := 0
	sum := 0
	for i := a; i <= b; i++ {
		if isPrimeNumber(i) {
			sum = sum + i 
			n++
		}
	}
	return n, sum
}

func isPrimeNumber(a int) bool {
	if a < 2 {
		return false
	}
	for i := 2; i <= a/2; i++ {
		if a % i == 0 {
			return false
		}
	}
	return true
}