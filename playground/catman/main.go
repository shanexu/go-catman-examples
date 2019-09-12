package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/samuel/go-zookeeper/zk"

	"github.com/shanexu/go-catman"
)

func main() {
	cm, err := catman.NewCatMan([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}
	defer cm.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	cs, err := cm.CMChildren("/children", catman.WatcherFunc(func(event zk.Event) {
		fmt.Println(event)
		wg.Done()
	}))
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", cs)
	wg.Wait()
}
