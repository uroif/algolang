package main

import "fmt"

func main() {
  var a int
  
  fmt.Scanf("%d", &a)
  
  if a >= 0 {
    if a % 2 == 0 {
      fmt.Printf("%d la so chan", a)
    } else {
      fmt.Printf("%d la so le", a)
    }
  } else {
    fmt.Printf("%d la so am", a)
  }
}