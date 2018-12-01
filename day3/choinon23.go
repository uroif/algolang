package main

import "fmt"

func main() {
	var d1, m1, y1, d2, m2, y2 int

	fmt.Scanf("%d %d %d\n", &d1, &m1, &y1)
	fmt.Scanf("%d %d %d\n", &d2, &m2, &y2)

	if birth(d1, m1, y1) < birth(d2, m2, y2) {
		fmt.Printf("1")
	} else if birth(d1, m1, y1) > birth(d2, m2, y2) {
		fmt.Printf("2")
	} else {
		fmt.Printf("0")
	}
}

func birth(d, m, y int) int {
	return y*10000 + m*100 + d
}
