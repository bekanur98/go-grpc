package main

import "fmt"

func main() {
  func(){
    fmt.Println("anonym func")
  } ()

  inc := increment()
  fmt.Println(inc())
  fmt.Println(inc())
  fmt.Println(inc())
  fmt.Println(inc())
}

func increment() func() int {
  count := 0
  return func() int {
    return count++
  }
}
