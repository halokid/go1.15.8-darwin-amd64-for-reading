package main

import (
  "log"
  "time"
)

func main() {
  chan1 := make(chan int)
  chan2 := make(chan int)

  go func() {
    chan1 <- 1
    time.Sleep(time.Second * 5)
  }()

  go func() {
    chan2 <- 1
    time.Sleep(time.Second * 5)
  }()

  select {
  case <-chan1:
    log.Println("chan1 ready")
  case <-chan2:
    log.Println("chan2 ready")
  default:
    log.Println("default")
  }

  log.Println("main exit")
}



