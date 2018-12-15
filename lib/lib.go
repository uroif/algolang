package lib

import "fmt"

func DoPrint(c string, number int) {
	for i := 0; i < number; i++ {
		fmt.Printf("%s", c)
	}
}

func InputSlice(n int) []int {
	var slice []int

	for i := 0; i < n; i++ {
		var value int
		fmt.Scan(&value)
		slice = append(slice, value)
	}
	return slice
}

func PrintSlice(slice []int) {
	for i := 0; i < len(slice); i++ {
		fmt.Printf("%d ", slice[i])
	}
	fmt.Printf("\n")
}
