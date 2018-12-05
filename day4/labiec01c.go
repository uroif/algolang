package main

import "fmt"

func main() {
	var a, b int 

	fmt.Scanf("%d %d", &a, &b)

	if a % 2 != 0 {
		a++
	}

	for i := a; i <= b; i=i+2 {
		fmt.Printf("%d\n", i)
	}
}