package main

import "fmt"

func main() {
	var n int
	var slc []int
	dem := make(map[int]int)

	for {
		stt, _ := fmt.Scan(&n)
		if stt == 0 {
			break
		}
		slc = append(slc, n)
	}

	for _, el := range slc {
		dem[el]++
	}

	for _, el := range slc {
		if dem[el] > 0 {
			fmt.Printf("%d %d\n", el, dem[el])
			dem[el] = 0
		}
	}
}
