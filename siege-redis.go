package main

import (
    "time"
    "fmt"
    "github.com/satori/go.uuid"
    "flag"
    "os"
    "log"
    "sync/atomic"
)

func main(){
  defer func() {
    if r := recover(); r != nil {
      log.Printf("Runtime error caught: %v", r)
    }
  }()
  flag.Parse()
  if h {
    flag.Usage()
    return
  }
  CheckParam()
  log.Println("---start---------------")
  wg.Add(1)
  if wOrR{
     Write()
  }else{
     Read()
  }
  wg.Wait()
  log.Println("---test end---------------")
  fmt.Println("Result:")
  fmt.Println("write/read key number successfully:",*keyNum)
  fmt.Println("failed write/read key number:",*errNum)
}

func Read(){

}

func Write(){
  defer wg.Done()
  for maxThread > 0{
      //set
      maxThread--
      go WriteBase()
  }
  time.Sleep(time.Duration(uc) * time.Second)
}

func WriteBase(){
    for{
       u, _ := uuid.NewV4()
       //u1, _ := uuid.NewV4()
       //key := prefix+":"+u.String()
       //err := redisClient.SAdd(key,u1.String())
       key := "aa"
       err := redisClient.SAdd(key,u.String())
       if err.Err() != nil{
         fmt.Println("sadd: ",key,"err:",err.Err())
         go func(){
           atomic.AddInt64(errNum, 1)
         }()
       }else{
         go func(){
          atomic.AddInt64(keyNum, 1)
         }()
       }
    }
}

func usage() {
    fmt.Fprintf(os.Stderr, `Welcome to use siege-redis! 

For My Daring:@Qingxiaokui

Usage: siege-redis [-h help] [-t] [-p] [-pool] [-m] [-host] [-port] [-a] [-w] [-o]

Options:
`)
    flag.PrintDefaults()
}

func Stop() {
  log.Printf("Shutdown Script ...")
  defer wg.Done()
  if srv != nil {
    ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
    defer cancel()
    if err := srv.Shutdown(ctx); err != nil {
      log.Printf("Script Shutdown:", err)
    }
    log.Printf("Script exist")
  }
}

