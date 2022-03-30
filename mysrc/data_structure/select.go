package main

import (
  "time"
  "log"
)

func addNumberToChan(chanName chan int) {
  for {
    chanName <- 1
    time.Sleep(1 * time.Second)
  }
}

func main() {
  var chan1 = make(chan int, 10)
  var chan2 = make(chan int, 10)

  go addNumberToChan(chan1)
  go addNumberToChan(chan2)

  for {
    select {
    case e := <- chan1:
      log.Printf("get element from chan1 -->>> %+v", e)
    case e := <- chan2 :
      log.Printf("get element from chan2 -->>> %+v", e)
    default:
     log.Printf("No element in chan1 and chan2")
     time.Sleep(1 * time.Second)
    }
  }
}