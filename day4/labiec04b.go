package main

import "fmt"

func main() {
	var x int

	for {
		fmt.Scanf("%d", &x)
		if x % 10 == 0 {
			break
		} else {
			fmt.Printf("%d ", x)
		}
	}
}