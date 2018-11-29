package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c int

	fmt.Scanf("%d %d %d", &a, &b, &c)

	delta := (b * b) - (4 * a * c)

	if delta < 0 {
		fmt.Printf("Phuong trinh da cho vo nghiem")
	} else if delta == 0 {
		x := -float64(b) / (2 * float64(a))
		fmt.Printf("Phuong trinh da cho co nghiem kep x1 = x2 = %f", x)
	} else {
		sqrtDelta := math.Sqrt(float64(delta))
		x1 := (-float64(b) + sqrtDelta) / (2 * float64(a))
		x2 := (-float64(b) - sqrtDelta) / (2 * float64(a))
		fmt.Printf("Phuong trinh da cho co 2 phan biet x1 = %f và x2 = %f", x1, x2)
	}
}