package gotest

import (
  "fmt"
  "os"
  "testing"
)

func TestMain(m *testing.M) {
  fmt.Println("TestMain setup...")

  // process test, inlude unit test, performance test, example test
  retCode := m.Run()

  fmt.Println("TestMain tear down...")
  os.Exit(retCode)
}