package main

import "fmt"

func main() {
	var n int

	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		for j := i + 1; j <= (2*i - 1); j++ {
			fmt.Print(2*i - j)
		}
		fmt.Println()
	}

	// 	for i := 1; i <= n; i++ {
	// 		for j := 1; j <= (2*i-1); j++ {
	// 			if j <= i {
	// 				fmt.Print(j)
	// 			} else {
	// 				fmt.Print(2*i-j)
	// 			}
	// 		}
	// 		fmt.Println()
	// 	}
}