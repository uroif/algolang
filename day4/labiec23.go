package main

import "fmt"

func main() {
	var n, x int 

	fmt.Scanf("%d\n", &n)

	for i := 1; i <= n; i++ {
		fmt.Scanf("%d\n", &x)
		fmt.Printf("%d\n", sumDigit(x))
	}
}

func sumDigit(n int) int {
	sum := 0
	for i := 10000; i >= 1; i = i/10 {
		sum = sum + n / i
		n = n % i
	}
	return sum
}