package main

import (
	"fmt"
	"strings"
)

func main() {
	var ntest int
	var avenger, wanda string

	fmt.Scan(&ntest)

	for ntest > 0 {
		ntest--

		fmt.Scan(&avenger, &wanda)
		if strings.Contains(avenger, wanda) {
			fmt.Println(avenger + " chua xau " + wanda)
		} else {
			fmt.Println(wanda + " khong nam trong " + avenger)
		}
	}
}
