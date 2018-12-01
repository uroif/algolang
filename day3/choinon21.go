package main

import (
	"fmt"
	"math"
)

func main() {
	var a int

	fmt.Scanf("%d", &a)

	if sqrNum(a) {
		fmt.Printf("Yes")
	} else {
		fmt.Printf("No")
	}
}

func sqrNum(a int) bool {
	x := math.Sqrt(float64(a))
	if int(x)*int(x) == a && a != 0 {
		return true
	}
	return false
}
