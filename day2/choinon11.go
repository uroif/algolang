package main

import "fmt"

func main() {
  var a, b int
  
  fmt.Scanf("%d %d", &a, &b)
  
  if a == 0 && b != 0 {
    fmt.Printf("Phuong trinh da cho vo nghiem")
  } else if a == 0 && b == 0 {
    fmt.Printf("Phuong trinh da cho vo so nghiem")
  } else {
    fmt.Printf("Phuong trinh da cho co 1 nghiem x = %f", -float64(b)/float64(a))
  }
}