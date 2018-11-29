package main

import (
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Scanf("%d %d %d", &a, &b, &c)

	if isTriangleSquare(a, b, c) {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
}

func isTriangleSquare(a, b, c int) bool {
	if a*a==b*b+c*c || a*a+b*b==c*c || b*b==c*c+a*a {
		return true
	} else {
		return false
	}
}