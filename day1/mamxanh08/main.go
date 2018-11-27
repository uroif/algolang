package main

import "fmt"

func main() {
	var n int
	
	fmt.Scanf("%d\n", &n)

	binhPhuong := n * n
	lapPhuong := n * n * n

	fmt.Printf("Binh phuong cua %d la: %d\n", n, binhPhuong)
	fmt.Printf("Lap phuong cua %d la: %d", n, lapPhuong)
}
