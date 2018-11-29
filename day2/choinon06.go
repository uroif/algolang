package main

import "fmt"

func main() {
  var gio, luong int
  
  fmt.Scanf("%d %d", &gio, &luong)
  
  if gio <= 40 {
    fmt.Println(gio * luong)
  } else {
    fmt.Println(40*luong + (gio - 40)*2*luong)
  }
}