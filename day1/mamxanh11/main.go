package main

import "fmt"

func main() {
	var A, B, C int
	
	fmt.Scanf("%d %d %d", &A, &B, &C)

	result := (A - B) * C
	fmt.Printf("%d", result)
}
