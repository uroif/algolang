package main

import "fmt"

func main() {
	var year int

	fmt.Scanf("%d", &year)

	if isLeapYear(year) {
		fmt.Printf("1")
	} else {
		fmt.Printf("0")
	}
}

func isLeapYear(n int) bool {
	if n % 400 == 0 {
		return true
	} else if n % 100 == 0 {
		return false
	} else if n % 4 == 0 {
		return true
	}
	return false
}