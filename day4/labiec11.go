package main

import "fmt"

func main() {
	var a, b int 

	fmt.Scanf("%d %d", &a, &b)

	for i := 1; i <= b; i++ {
		for j := 1; j <= a; j++ {
			fmt.Printf("*")
		}
		fmt.Printf("\n")
	}
}