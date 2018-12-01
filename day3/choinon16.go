package main

import (
	"math"
	"fmt"
)

func main() {
	var a, b, c int

	fmt.Scanf("%d %d %d", &a, &b, &c)

	fmt.Printf("%d %d", max(a, b, c), min(a, b, c))
}

func max(a, b, c int) int {
	max := math.Max(float64(a), float64(b))
	max = math.Max(max, float64(c))
	return int(max)
}

func min(a, b, c int) int {
	min := math.Min(float64(a), float64(b))
	min = math.Min(min, float64(c))
	return int(min)
}