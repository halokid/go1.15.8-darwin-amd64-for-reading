package gotest

import (
  "testing"
)

func TestAdd(t *testing.T) {
  var a = 1
  var b = 2
  var expected = 3

  actual := Add(a, b)
  t.Logf("actual -->>> %+v, expected -->>> %+v",
    actual, expected)
}

