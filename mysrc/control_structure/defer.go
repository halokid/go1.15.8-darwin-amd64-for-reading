package main

import "log"

func deferFuncParameter() {
  var aInt = 1

  defer log.Println("aInt -->>>", aInt)

  aInt = 2
  return
}

func printArray(arr *[3]int) {
  for i := range arr {
    log.Println(arr[i])
  }
}

func deferFuncPx() {
  var aArr = [3]int{1, 2, 3}

  defer printArray(&aArr)

  aArr[0] = 10
  return
}

func deferFuncPy() (result int) {
  log.Println("result 1 -->>>", result)
  i := 1
  log.Println("result 2 -->>>", result)

  defer func() {
    // todo: the func return result(i) to defer, so `result` be 1 here
    log.Println("result 4 -->>>", result)
    result++
    // todo: after result++, it change to 2
    log.Println("result 5 -->>>", result)
  }()

  log.Println("result 3 -->>>", result)
  return i    // todo: return i to `result`, result here is 1
}

func main() {
  deferFuncParameter()

  deferFuncPx()

  i := deferFuncPy()
  log.Println("i -->>>", i)
}


