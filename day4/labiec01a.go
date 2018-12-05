package main

import "fmt"

func main() {
	var a, b int 

	fmt.Scanf("%d %d\n", &a, &b)

	for i := 0; i <= (b - a); i++ {
		fmt.Printf("%d\n", a + i)
	}
}