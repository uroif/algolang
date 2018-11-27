package main

import "fmt"

func main() {
	var hour, minute, second, X int
	
	fmt.Scanf("%d:%d:%d %d", &hour, &minute, &second, &X)

	time := (hour * 3600 + minute * 60 + second) + X
	second = time % (24 * 3600)

	hour = second / 3600
	remain := second % 3600
	minute = remain / 60
	second = remain % 60

	fmt.Printf("%d : %d : %d", hour, minute, second)
}
