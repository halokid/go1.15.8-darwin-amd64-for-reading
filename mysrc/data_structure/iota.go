package main

import "log"

const (
  bit0, mask0 = 1 << iota, 1 << iota - 1
  bit1, mask1
  bit2, mask2
  bit3, mask3
)

func main() {
  log.Printf("bit0, mask0 -->>> %+v, %+v", bit0, mask0)
  log.Printf("bit1, mask1 -->>> %+v, %+v", bit1, mask1)
  log.Printf("bit2, mask2 -->>> %+v, %+v", bit2, mask2)
  log.Printf("bit3, mask3 -->>> %+v, %+v", bit3, mask3)
}

