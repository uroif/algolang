package main

import (
	"fmt"
)

func main() {
	var ntest int
	var s string

	fmt.Scanf("%d\n", &ntest)

	for ntest > 0 {
		ntest--

		fmt.Scan(&s)
		if isPalindrome(s) {
			fmt.Print("1 ")
		} else {
			fmt.Print("0 ")
		}
	}
}

func isPalindrome(s string) bool {
	length := len(s)
	for i, _ := range s {
		if s[i] != s[length - 1 - i] {
			return false
		}
	}
	return true
}