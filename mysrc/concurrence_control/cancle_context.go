package main

import (
  "context"
  "log"
  "time"
)

func HandleRequest(ctx context.Context) {
  go WriteRedis(ctx)
  go WriteDatabase(ctx)

  for {
    select {
    case <- ctx.Done():
      log.Println("HandleRequest Done.")
      return
    default:
      log.Println("HandleRequest running.")
      time.Sleep(2 * time.Second)
    }
  }
}

func WriteRedis(ctx context.Context) {
  for {
    select {
    case <- ctx.Done():
      log.Println("WriteRedis Done.")
      return
    default:
      log.Println("WriteRedis running.")
      time.Sleep(2 * time.Second)
    }
  }
}

func WriteDatabase(ctx context.Context) {
  for {
    select {
    case <- ctx.Done():
      log.Println("WriteDatabase Done.")
      return
    default:
      log.Println("WriteDatabase running.")
      time.Sleep(2 * time.Second)
    }
  }
}

func main() {
  ctx, cancel := context.WithCancel(context.Background())
  go HandleRequest(ctx)

  time.Sleep(5 * time.Second)
  log.Println("it's time to stop all sub goroutines!")
  cancel()

  // just for test whether sub goroutines exit or not
  time.Sleep(5 * time.Second)
}






