package main

import "fmt"

func main() {
	var n int 

	fmt.Scanf("%d", &n)

	for i := 1; i <= n/2; i++ {
		if n % i == 0 && i <= n/i {
			if i != n/i && i != 1 {
				fmt.Printf("%d %d ", i, n/i)
			} else {
				fmt.Printf("%d ", i)
			}
		}
	}
}
