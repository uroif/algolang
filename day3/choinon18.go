package main

import (
  "fmt"
)

func main() {
  var a int
  
  fmt.Scanf("%d\n", &a)
  
  if isDividedBy12(a) {
    fmt.Printf("Yes")
  } else {
    fmt.Printf("No")
  }
}

func isDividedBy12(n int) bool {
  if n % 12 == 0 {
    return true
  }
  return false
}