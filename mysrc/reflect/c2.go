package main

import (
  "log"
  "reflect"
)

// 反射第二定律:反射可以将反射对象还原成interface对象
func main() {
  var x float64 = 3.4

  v := reflect.ValueOf(x)

  var y float64 = v.Interface().(float64)
  log.Printf("y -->>> %+v", y)
}
