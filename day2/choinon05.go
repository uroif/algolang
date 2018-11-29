package main

import "fmt"

func main() {
  var a, b, c int
  
  fmt.Scanf("%d %d %d", &a, &b, &c)
  
  if a >= b && a >= c {
    fmt.Printf("%d la so lon nhat", a)
  } else if b >= a && b >= c {
    fmt.Printf("%d la so lon nhat", b)
  } else {
	fmt.Printf("%d la so lon nhat", c)
  }
}