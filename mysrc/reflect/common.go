package main

import (
  "io"
  "log"
  "os"
)

type Myint int

var i int
var j Myint

func main() {
  var r io.Reader
  tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
  if err != nil {
   return
  }
  r = tty
  log.Println(r)
}
