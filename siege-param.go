package main

import (
    "fmt"
    "runtime"
    "strings"
    "regexp"
    "flag"
    "strconv"
    "github.com/go-redis/redis"
    "syscall"
    "github.com/Zfzlv/siege-redis/signal"
    "sync"
)

var (
  redisClient *redis.Client
  keyNum,errNum,INCR *int64
  maxNum int64
  host,pwd,lon,prefix string
  port,maxThread,pool int
  h,wOrR,orderON bool
  uc int
  wg  sync.WaitGroup
)

func init(){
  runtime.GOMAXPROCS(runtime.NumCPU())
  keyAddress := int64(0)
  errAddress := int64(0)
  IncrAddress := int64(0)
  keyNum,errNum = &keyAddress,&errAddress
  INCR = &IncrAddress
  flag.BoolVar(&h, "h", false, "this help")
  flag.BoolVar(&wOrR, "w", false, "true means write,false means read")
  flag.BoolVar(&orderON, "o", false, "true means write by an ordered positive number, false means write by random string")
  flag.StringVar(&lon, "t", "1m", "test time;you can use these units:s,m,h")
  flag.StringVar(&host, "host", "10.0.18.183", "redis host")
  flag.StringVar(&pwd, "a", "TwcqX", "redis password")
  flag.StringVar(&prefix, "p", "test", "key prefix: like test + `:1`")
  flag.IntVar(&port, "port", 9221, "redis port")
  flag.IntVar(&pool, "pool", 10, "number of connection pool")
  flag.IntVar(&maxThread, "m", 100, "number of concurrent threads")
  flag.Int64Var(&maxNum, "maxNum", 10000, "the maximum ordinal number that can be read by the client")
  flag.Usage = usage
  signal.RegistHandle(syscall.SIGINT, Stop)
  signal.RegistHandle(syscall.SIGTERM, Stop)
  signal.RegistHandle(syscall.SIGKILL, Stop)
  signal.RegistHandle(syscall.SIGQUIT, Stop)
}

func CheckParam(){
  var err error
  redisClient = redis.NewClient(&redis.Options{
        Addr:     host+":"+strconv.Itoa(port),
        Password: pwd,
        DB:       0,
        PoolSize: pool,
  })
  re,_ := regexp.Compile(`^(\d+)(\S)[\s\S]*$`)
  u := re.ReplaceAllString(strings.TrimSpace(lon),"$1")
  uc,err = strconv.Atoi(u)
  if err != nil{
     panic("test time parse error:"+err.Error())
  }
  c := re.ReplaceAllString(lon,"$2")
  if c == "m"{
    uc = uc * 60
  }else if c == "h"{
    uc = uc * 3600
  }
  //list param
  fmt.Printf("host:%s\nport:%d\npassword:%s\npoolThread num:%d\nthread num:%d\n",host,port,pwd,pool,maxThread)
  fmt.Printf("siege time:%s\nwrite or read:%t\nwrite by an ordered positive number or not:%t\n",lon,wOrR,orderON)
  fmt.Printf("maximum ordinal number can be read :%d\nkey's prefix:%s\n",maxNum,prefix)
  v,err := redisClient.Ping().Result()
  if err != nil{
    panic("client ping err :"+err.Error())
  }
  fmt.Println("client ping result :",v)
}