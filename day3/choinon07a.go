package main

import (
  "fmt"
  "math"
)

func main() {
  var a, b, c int
  
  fmt.Scanf("%d %d %d", &a, &b, &c)
  
  if isTriangle(a, b, c) {
    fmt.Printf("Dien tich cua tam giac la %f\n", sTriangle(a, b, c)) 
  } else {
    fmt.Printf("Khong phai la tam giac")
  }
}

func isTriangle(a, b, c int) bool {
  if a + b > c && b + c > a && a + c > b {
    return true
  }
  return false
}

func sTriangle(a, b, c int) float64 {
  p := 0.5 * float64(a + b + c)
  s := math.Sqrt(p * (p - float64(a)) * (p - float64(b)) * (p - float64(c)))
  return s
}