package main

import "fmt"

func main() {
  var point float32
  
  fmt.Scanf("%f", &point)
  
  if point >= 9 {
    fmt.Printf("Xuat sac")
  } else if point >= 8 {
    fmt.Printf("Gioi")
  } else if point >= 7 {
    fmt.Printf("Kha")
  } else if point >= 6 {
    fmt.Printf("TBKha")
  } else if point >= 5 {
    fmt.Printf("Tbinh")
  } else {
    fmt.Printf("Yeu")
  }
}