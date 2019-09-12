package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shanexu/go-catman"
)

func main() {
	cm, err := catman.NewCatMan([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}
	defer cm.Close()
	l := cm.NewLeaderElectionSupport("myhost", "/election")
	l.AddListener(catman.LeaderElectionAwareFunc(func(event catman.ElectionEvent) {
		fmt.Println("ElectionEvent:", event)
	}))
	if err := l.Start(); err != nil {
		panic(err)
	}
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	if err := l.Stop(); err != nil {
		panic(err)
	}
}
