package main

import (
  "log"
  "reflect"
)

// 反射第一定律:反射可以将interface类型变量转换成反 射对象

func main() {
  var x float64 = 3.4
  t := reflect.TypeOf(x)
  log.Printf("t -->>> %+v", t)

  v := reflect.ValueOf(x)
  log.Printf("v -->>> %+v", v)
}
