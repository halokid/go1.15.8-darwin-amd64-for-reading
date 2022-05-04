package main

import "log"

func main() {
  chan1 := make(chan int)
  chan2 := make(chan int)

  go func() {
    close(chan1)
  }()

  go func() {
    close(chan2)
  }()

  select {
  case <-chan1:
    log.Println("chan1 ready")
  case <-chan2:
    log.Println("chan2 ready")

  }

  log.Println("main exit")
}
