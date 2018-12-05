package main

import "fmt"

func main() {
	var n, k int 

	fmt.Scanf("%d", &n)

	for i := 2; i < n; i++ {
		k = n / i
		if i * k == n {
			fmt.Printf("%d ", i)
		}
	}
}