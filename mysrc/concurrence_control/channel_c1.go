package main

import (
  "log"
  "time"
)

func Process(ch chan int) {
  // todo: do some work...
  time.Sleep(time.Second)

  // todo: write one element to channel, expression current goroutine already done
  ch <- 1
}

func main() {
  chanNum := 100
  channels := make([]chan int, chanNum)

  for i := 0; i < chanNum; i++ {
    channels[i] = make(chan int)
    go Process(channels[i])
  }

  for i, ch := range channels {
    <- ch
    log.Println("Routine ", i, " quit!")
  }
}

