package main

import "log"

func addElement(slice []int, e int) []int {
  return append(slice, e)
}

func main() {
  //var array [10]int
  var array [20]int

  var slice = array[5:6]

  log.Printf("length of slice -->>> %+v", len(slice))
  log.Printf("capacity of slice -->>> %+v", cap(slice))
  log.Printf( "result of (&slice[0] == &array[5]) -->>> %+v", &slice[0] == &array[5])
  log.Printf( "result of (&slice[0] == &array[6]) -->>> %+v", &slice[0] == &array[6])

  // ----------------------------------------------------------
  var slicex []int
  slicex = append(slicex, 1, 2, 3)
  log.Println(len(slicex))
  log.Println(cap(slicex))

  newSlice := addElement(slicex, 4)
  log.Println(len(newSlice))
  log.Println(cap(newSlice))

  log.Printf("result of &slice[0] == &newSlice[0] -->>> %+v",
    &slicex[0] == &newSlice[0])

  log.Printf("result of &slice == &newSlice -->>> %+v",
    &slicex == &newSlice)

  // -------------------------------------------------------------
  orderLen := 5
  order := make([]uint16, 2 * orderLen)

  pollOrder := order[:orderLen:orderLen]
  //lockOrder := order[orderLen:][:orderLen:orderLen]
  lockOrder := order[orderLen:]

  log.Println("len(pollOrder) -->>>", len(pollOrder))
  log.Println("cap(pollOrder) -->>>", cap(pollOrder))

  log.Println("len(lockOrder) -->>>", len(lockOrder))
  log.Println("cap(lockOrder) -->>>", cap(lockOrder))

  s := make([]int, 2)
  log.Println("s -->>>", s)
  s = append(s, 1)
  s = append(s, 2)
  s = append(s, 3)
  s = append(s, 4)
  log.Println("s len -->>>", len(s))
  log.Println("s cap -->>>", cap(s))
  log.Printf("s address -->>> %p", &s)

  s = append(s, 5)
  s = append(s, 6)
  s = append(s, 7)
  s = append(s, 8)
  s = append(s, 9)
  log.Printf("after s address -->>> %p", &s)
  log.Println("s -->>>", s)
  log.Println("s len -->>>", len(s))
  log.Println("s cap -->>>", cap(s))

  a := make([]int, 10, 10)
  x := a[5:7]
  log.Println("x -->>>", x)
  log.Println("x len -->>>", len(x))
  log.Println("x cap -->>>", cap(x))

  aa := make([]int, 0)
  log.Println("aa -->>>", aa, cap(aa))
  log.Printf("aa point -->>> %p", &aa)

  aa = append(aa, 2)
  aa = append(aa, 3)
  log.Println("aa -->>>", aa, cap(aa))
  log.Printf("aa point -->>> %p", &aa)

  dd := make([]int, 0)
  cc := copy(aa, dd)
  log.Println("cc -->>>", cc)

  bb := make([]int, 0)
  log.Println("bb -->>>", bb)
}










