package main

import logpkg "log"

func main() {
  var s string
  logpkg.Printf("s -->>> %+v", s)

  var bs []byte   // todo: not allocation memory
  logpkg.Printf("bs -->>> %+v", bs)
  //bs[0] = 'a'     // todo: no memory, will panic

  bs = make([]byte, 2)
  bs[0] = 'a'
  bs[1] = 'b'
  logpkg.Printf("bs1 -->>> %+v", bs)
  //bs[2] = 'j'     // todo: not enough memory, panic
  bs = append(bs, 'j')
  logpkg.Printf("bs2 -->>> %+v", bs)
}

