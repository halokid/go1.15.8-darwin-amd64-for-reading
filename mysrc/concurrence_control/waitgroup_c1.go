package main

import (
  "log"
  "sync"
  "time"
)

func main() {
  var wg sync.WaitGroup

  wg.Add(2)   // set counter, the number is goroutines-number
  go func() {
    // todo: do some work
    time.Sleep(1 * time.Second)

    log.Println("Goroutine 1 finshed!")
    wg.Done()   // todo: the goroutine number will reduce 1
  } ()

  go func() {
    // todo: do some work
    time.Sleep(1 * time.Second)

    log.Println("Goroutine 2 finshed!")
    wg.Done()
  } ()

  wg.Wait()
  log.Println("All Goroutine finished!")
}




