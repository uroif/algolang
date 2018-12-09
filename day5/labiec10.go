package main

import "fmt"

func main() {
	var a, b int

	fmt.Scanf("%d %d", &a, &b)

	uCLN := timULN(a, b)
	bCNN := (a * b) / uCLN
	
	fmt.Printf("%d %d", uCLN, bCNN)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func timULN(a, b int) int {
	n := min(a/2, b/2)
	var uc int

	for i := n; i >= 1; i-- {
		if a%i == 0 && b%i == 0 {
			uc = i
			break
		}
	}
	return uc
}
