package main

import "fmt"

func main() {
	var n, sum int 

	fmt.Scanf("%d", &n)

	for i := 1; i <= n; i++ {
		sum = sum + i
	}
	fmt.Printf("%d", sum)
}