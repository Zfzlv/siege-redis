package main

import (
    "time"
    "fmt"
    "github.com/satori/go.uuid"
    "flag"
    "os"
    "log"
    "sync/atomic"
    "strconv"
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
     go Write()
  }else{
     go Read()
  }
  //go Qps()
  go Progress()
  wg.Wait()
  fmt.Println("")
  log.Printf("---test end---------------\n")
  fmt.Println("Result:")
  if wOrR{
    fmt.Println("Number of keys written successfully:",*keyNum)
    fmt.Println("Number of write key failures:",*errNum)
  }else{
    fmt.Println("Number of keys successfully read:",*keyNum)
    fmt.Println("Number of failed to read key:",*errNum)
  }
}

func Read(){
  defer wg.Done()
  for maxThread > 0{
      maxThread--
      go ReadBase()
  }
  time.Sleep(time.Duration(uc) * time.Second)
}

func ReadBase(){
  for{
     key := prefix+":"+strconv.FormatInt(Incr(),10)
     value := redisClient.SMembers(key).Val()
     if len(value) <= 0{
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

func Write(){
  defer wg.Done()
  for maxThread > 0{
      maxThread--
      go WriteBase()
  }
  time.Sleep(time.Duration(uc) * time.Second)
}

func WriteBase(){
    for{
       key := prefix+":"
       if orderON{
          key += strconv.FormatInt(Incr(),10)
       }else{
          u := uuid.NewV4()
          key += u.String()
       }
       u1 := uuid.NewV4()
       err := redisClient.SAdd(key,u1.String())
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

For My Daring:%c[1;40;32m@echosimple%c[0m %c[1;40;33mhttps://www.zcool.com.cn/u/15027113%c[0m

Usage: siege-redis [-h help] [-t] [-p] [-pool] [-m] [-host] [-port] [-a] [-w] [-o]

Options:
`,0x1B,0x1B,0x1B,0x1B)
    flag.PrintDefaults()
}

func Stop() {
  log.Printf("premature closure...")
  defer wg.Done()
  log.Printf("Script exist")
}

func Incr() int64{
   return atomic.AddInt64(INCR, 1)
}