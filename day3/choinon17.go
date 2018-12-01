package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)

	fmt.Printf("%d\n", isDividedBy10(a*b))
}

func isDividedBy10(n int) int {
	if n%10 == 0 {
		return 1
	}
	return 0
}
