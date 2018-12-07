package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)

	fmt.Printf("%d", timULN(a, b))
}

func timULN(a, b int) int {
	var max int

	for i := 1; i <= a; i++ {
		if a % i == 0 && b % i == 0 {
			max = i
		}
	}
	return max
}

// func min(a, b int) int {
// 	if a < b {
// 		return a
// 	}
// 	return b
// }

// func timULN(a, b int) int {
// 	n := min(a/2, b/2)
// 	var uc int

// 	for i := 1; i <= n; i++ {
// 		if a % i == 0 && b % i == 0 {
// 			uc = i
// 		}
// 	}
// 	return uc
// }