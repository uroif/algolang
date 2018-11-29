package main

import "fmt"

func main() {
  var a, b int
  
  fmt.Scanf("%d %d", &a, &b)
  
  if a > b {
    fmt.Printf("So lon nhat = %d", a)
  } else {
    fmt.Printf("So lon nhat = %d", b)
  }
}