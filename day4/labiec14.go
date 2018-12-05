package main

import "fmt"

func main() {
	var n, x int 

	fmt.Scanf("%d\n", &n)

	fmt.Printf("%d\n", (n-1)/2)

	for i := 1; i <= (n-1)/2; i++ {
		x = n - i
		fmt.Printf("%d %d\n", x, i)
	}
}