package main

import "fmt"

func main() {
	var n int

	fmt.Scanf("%d\n", &n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= 2*n-1; j++ {
			if j == n-i+1 || j == n+i-1 || i == n {
				fmt.Printf("*")
			} else {
				fmt.Printf(" ")
			}
		}
		fmt.Printf("\n")
	}
}
