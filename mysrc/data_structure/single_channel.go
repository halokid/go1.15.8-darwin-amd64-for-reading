package main

import "log"

// single read
//func readChan(chanName chan int) {
func readChan(chanName <-chan int) {
  n := <- chanName
  log.Printf("n -->>> %v", n)
}

// single write
//func writeChan(chanName chan int) {
func writeChan(chanName chan<- int) {
  chanName <- 1
}

func main() {
  //var myChan = make(chan int, 10)
  var myChan = make(chan int, 1)

  writeChan(myChan)
  readChan(myChan)
}
