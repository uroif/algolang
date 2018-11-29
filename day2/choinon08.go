package main

import "fmt"

func main() {
  var a, b, c int

  fmt.Scanf("%d %d %d", &a, &b, &c)
  
  if a >= b && b >= c {
    fmt.Println(c, b, a)
  } else if a >= c && c >= b {
    fmt.Println(b, c, a)
  } else if b >= a && a >= c {
    fmt.Println(c, a, b)
  } else if b >= c && c >= a {
    fmt.Println(a, c, b)
  } else if c >= a && a >= b {
    fmt.Println(b, a, c)
  } else {
    fmt.Println(b, c, a)
  }
}