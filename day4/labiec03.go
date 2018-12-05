package main

import "fmt"

func main() {
	var n, x int 

	fmt.Scanf("%d\n", &n)

	sum := 0
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &x)
		sum = sum + x
	}
	fmt.Printf("%d", sum)
}