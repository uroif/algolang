package main

import (
	"fmt"
	"strings"
)

func main() {
	var ntest int
	var a, b string

	fmt.Scan(&ntest)

	for ntest > 0 {
		ntest--

		fmt.Scan(&a, &b)

		n := strings.Compare(a, b)
		if n == 0 {
			fmt.Println("1")
		} else {
			fmt.Println("0")
		}
	}
}
