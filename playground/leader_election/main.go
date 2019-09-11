package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/shanexu/go-catman"
)

func main() {
	done := make(chan struct{})
	cm, err := catman.NewCatMan([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}
	cm.CMCreate("/leader_election", nil)
	cm.LeaderElector(context.Background(), func() {
		fmt.Println("takeLeaderShip")
		<-done
	}, "/testle", nil)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	close(done)
}
