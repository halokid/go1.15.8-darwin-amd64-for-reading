package main

import (
  "log"
  "reflect"
)

// 反射第三定律:反射对象可修改，value值必须是可设置 的
func main() {
  var x float64 = 3.4
  //v := reflect.ValueOf(x)
  //v.SetFloat(7.1)
  // todo: 上例中，传入reflect.ValueOf()函数的其实是x的值，而非x本身。即通过v修改其值是无法影响x的，也即是无效 的修改，所以golang会报错。

  // todo: right style
  v := reflect.ValueOf(&x)
  v.Elem().SetFloat(7.1)
  log.Println("x Elem -->>>", v.Elem())
  log.Println("x -->>>", v.Elem().Interface())
}



