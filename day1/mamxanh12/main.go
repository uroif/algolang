package main

import "fmt"

func main() {
	var A, B int
	
	fmt.Scanf("%d %d", &A, &B)

	result := A + B + A * B
	fmt.Printf("%d", result)
}
