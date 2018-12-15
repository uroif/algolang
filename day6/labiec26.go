package main

import "fmt"

func main() {
	var ntest int

	fmt.Scanf("%d\n", &ntest)

	for ntest > 0 {
		ntest--
		var n, x int
		fmt.Scanf("%d\n", &n)
		var slc []int

		for i := 0; i < n; i++ {
			fmt.Scan(&x)
			slc = append(slc, x)
		}

		fmt.Println(isDayDanDau(slc))
	}
}

func isDayDanDau(slc []int) int {
	for i := 0; i < len(slc)-1; i++ {
		if slc[i]*slc[i+1] >= 0 {
			return 0
		}
	}
	return 1
}
