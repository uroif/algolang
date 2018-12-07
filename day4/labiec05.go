package main

import "fmt"

func main() {
	var x int 

	fmt.Scanf("%d", &x)

	for i := 0; i < x; i++ {
		for j := 1; j <= 10; j++ {
			fmt.Printf("%d ", i)
		}
		fmt.Printf("\n")
	}
}