package main

import "fmt"

func main() {
	var second int
	
	fmt.Scanf("%d\n", &second)

	hour := second / 3600
	remain := second % 3600
	minute := remain / 60
	second = remain % 60
	fmt.Printf("%d : %d : %d", hour, minute, second)
}
