package main

import (
  "log"
  "time"
)

func main() {
  chan1 := make(chan int)
  chan2 := make(chan int)

  writeFlag := false

  go func() {
    for {
      if writeFlag {
        chan1 <- 1
      }
      time.Sleep(time.Second)
    }
  }()

  go func() {
    for {
      if writeFlag {
        chan2 <- 1
      }
      time.Sleep(time.Second)
    }
  }()

  select {
  case <-chan1:
    log.Println("chan1 ready")
  case <-chan2:
    log.Println("chan2 ready")

  }

  log.Println("main exit")
}



