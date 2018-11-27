package main

import "fmt"

func main() {
	var r int
	pi := 3.14
	
	fmt.Scanf("%d\n", &r)

	s := pi * float64(r * r)
	fmt.Printf("Dien tich hinh tron co ban kinh %d la: %f", r, s)
}
