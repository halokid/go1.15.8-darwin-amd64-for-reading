package gotest

import (
  "testing"
  "time"
)

func parallelTest1(t *testing.T) {
  t.Parallel()
  time.Sleep(3 * time.Second)
}

func parallelTest2(t *testing.T) {
  t.Parallel()
  time.Sleep(2 * time.Second)
}

func parallelTest3(t *testing.T) {
  t.Parallel()
  time.Sleep(1 * time.Second)
}

// TestSubParallel 通过把多个子测试放到一个组中并发执行，同时多个子测试可以共享setup和tear-down
func TestSubParallel(t *testing.T) {
  // setup
  t.Logf("Setup")

  t.Run("group", func(t *testing.T) {
    t.Run("Test1", parallelTest1)
    t.Run("Test2", parallelTest2)
    t.Run("Test3", parallelTest3)
  })

  // tear down
  t.Logf("teardown")
}












