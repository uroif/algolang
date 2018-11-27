package main

import "fmt"

func main() {
	var r int
	pi := 3.14
	
	fmt.Scanf("%d\n", &r)

	s := pi * float64(r * r)
	v := (4.0/3) * pi * float64(r * r * r)

	fmt.Printf("Dien tich hinh cau la: %f\n", s)
	fmt.Printf("The tich hinh cau la: %f", v)
}
