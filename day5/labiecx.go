package main

import "fmt"

func main() {
	var ntest int 
	var x int

	fmt.Scanf("%d\n", &ntest)

	for ntest > 0 {
		ntest--

		fmt.Scanf("%d\n", &x)
		fmt.Printf("%d\n", cal(x))
	}
} 

func cal(n int) int {
	var i int

	if n <= 0 {
		return -1
	}

	for i = 0; n > 1; i++ {
		if n % 2 == 0 {
			n = n / 2
		} else {
			n = n * 3 + 1
		}
	}
	
	return i
}