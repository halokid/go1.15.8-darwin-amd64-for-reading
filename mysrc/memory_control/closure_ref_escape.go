package main

import "fmt"

func Fibonacci() (func() int) {
  a, b := 0, 1
  fmt.Println("first a:", a, ", b:", b)
  return func() int {
    a, b = b, a + b
    fmt.Println("second a:", a, ", b:", b)
    fmt.Println("=============================")
    return a
  }
}

func main() {
  f := Fibonacci()

  for i := 0; i < 10; i++ {
    fmt.Printf("Fibonacci: %d\n", f())
  }
}
