package main

import "fmt"

func main() {
	var n int 

	fmt.Scanf("%d", &n)

	for i := 2; i < n/2; i++ {
		if n % i == 0 && i < n/i {
			fmt.Printf("%d %d ", i, n/i)
		}
	}
}