package signal

import (
	"os"
	sig "os/signal"
	"log"
	"fmt"
)

type SignalHandle func()

var (
	sigHandlesMgt = make(map[os.Signal][]SignalHandle)
	sigalCh       = make(chan os.Signal)
)

func RegistHandle(s os.Signal, handlers ...SignalHandle) {
	sigHandles, ok := sigHandlesMgt[s]
	if !ok {
		sig.Notify(sigalCh, s)
	}
	sigHandlesMgt[s] = append(sigHandles, handlers...)
}

func init() {
	go func() {
		for {
			sig := <-sigalCh
			fmt.Println("")
            log.Printf("recv signal:%v\n", sig)
			if handlers, ok := sigHandlesMgt[sig]; ok {
				for _, handler := range handlers {
					handler()
				}
			} else {
				fmt.Println("")
				log.Printf("unkown signal:%v\n", sig)
			}
		}
	}()
}
