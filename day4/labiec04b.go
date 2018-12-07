package main

import (
	"fmt"
)

func main() {
	var x int
	
	for {
		status, _ := fmt.Scan(&x)
		if x%10 == 0 || status == 0 {
			break
		} else {
			fmt.Printf("%d ", x)
		}
	}
}
