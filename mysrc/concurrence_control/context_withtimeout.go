package main

import (
  "context"
  "log"
  "time"
)

func HandleRequestX(ctx context.Context) {
  go WriteRedisX(ctx)
  go WriteDatabaseX(ctx)

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

func WriteRedisX(ctx context.Context) {
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

func WriteDatabaseX(ctx context.Context) {
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
  ctx, _ := context.WithTimeout(context.Background(), 5 * time.Second)

  go HandleRequestX(ctx)

  time.Sleep(10 * time.Second)
}




