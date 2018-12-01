package main

import (
	"fmt"
	"math"
)

func main() {
	var A, B, C, D int

	fmt.Scanf("%d %d %d %d", &A, &B, &C, &D)

	fmt.Printf("%d", secondBiggest(A, B, C, D))
}

func secondBiggest(a, b, c, d int) int {
	maxAB := math.Max(float64(a), float64(b))
	minAB := math.Min(float64(a), float64(b))

	maxCD := math.Max(float64(c), float64(d))
	minCD := math.Min(float64(c), float64(d))

	minX := math.Min(maxAB, maxCD)

	if minAB > maxCD {
		return int(minAB)
	} else if minCD > maxAB {
		return int(minCD)
	} else {
		return int(minX)
	}
}
