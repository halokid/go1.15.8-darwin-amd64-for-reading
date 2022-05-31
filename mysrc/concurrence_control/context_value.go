package main

import (
  "context"
  "log"
  "time"
)

func HandleRequestY(ctx context.Context) {
  for {
    select {
    case <- ctx.Done():
      log.Println("HandleRequestY Done.")
      return
    default:
      log.Println("HandleRequestY running, parameter: ",
        ctx.Value("parameter"))
      time.Sleep(2 * time.Second)
    }
  }
}

func main() {
  ctx := context.WithValue(context.Background(), "parameter", "1")
  go HandleRequestY(ctx)

  //time.Sleep(10 * time.Second)
}
