package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Scanf("%d %d %d", &a, &b, &c)

	if laTriangle(a, b, c) && isTriangleSquare(a, b, c) {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
}

func laTriangle(a, b, c int) bool {
	if a+b > c && b+c > a && a+c > b {
		return true
	}
	return false
}

func isTriangleSquare(a, b, c int) bool {
	if a*a == b*b+c*c || a*a+b*b == c*c || b*b == c*c+a*a {
		return true
	}
	return false
}
