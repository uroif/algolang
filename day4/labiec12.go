package main

import "fmt"

func main() {
	var n, x int 

	fmt.Scanf("%d\n", &n)

	for i := 0; i < n; i++ {
		fmt.Scanf("%d\n", &x)

		if isPrime(x) {
			fmt.Printf("%d la so nguyen to\n", x)
		} else {
			fmt.Printf("%d khong phai la so nguyen to\n", x)
		}
	}
}

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	
	for i := 2; i <= n/2; i++ {
		if n % i == 0 {
			return false
		}
	}
	return true
}