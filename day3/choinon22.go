package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Scanf("%d %d %d", &a, &b, &c)

	if sub(a, b, c) {
		fmt.Printf("%d - %d = %d", a, b, c)
	} else if sub(a, c, b) {
		fmt.Printf("%d - %d = %d", a, c, b)
	} else if sub(b, c, a) {
		fmt.Printf("%d - %d = %d", b, c, a)
	} else if sub(b, a, c) {
		fmt.Printf("%d - %d = %d", b, a, c)
	} else if sub(c, b, a) {
		fmt.Printf("%d - %d = %d", c, b, a)
	} else if sub(c, a, b) {
		fmt.Printf("%d - %d = %d", c, a, b)
	} else {
		fmt.Printf("Không có đáp án.")
	}
}

func sub(a, b, c int) bool {
	if a - b == c {
		return true
	}
	return false
} 