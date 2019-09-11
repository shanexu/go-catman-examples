package main

import (
	"fmt"
	"time"

	"github.com/shanexu/go-catman"
)

type lockListener int

func (l lockListener) LockAcquired() {
	fmt.Printf("LockAcquired %d\n", l)
}

func (l lockListener) LockReleased() {
	fmt.Printf("LockReleased %d\n", l)
}

func main() {
	cm, err := catman.NewCatMan([]string{"127.0.0.1:2181"}, time.Second)
	if err != nil {
		panic(err)
	}
	go func() {
		l := cm.NewLock("/mylock", catman.OpenAclUnsafe, lockListener(0))
		l.Lock()
		fmt.Println("locked")
		time.Sleep(time.Second * 10)
		l.Unlock()
	}()

	l := cm.NewLock("/mylock", catman.OpenAclUnsafe, lockListener(1))
	l.Lock()
	fmt.Println("locked")
	time.Sleep(time.Second * 10)
	l.Unlock()
}
