package main

import (
	"fmt"
	"time"

	"github.com/shanexu/go-catman"
)

func main() {
	cm, err := catman.NewCatMan([]string{"127.0.0.1:2181"}, time.Second)
	defer cm.Close()
	if err != nil {
		panic(err)
	}
	q := cm.NewDistributedQueue("/queue")

	_, err = q.Offer([]byte("hello"))
	if err != nil {
		panic(err)
	}

	data, err := q.Element()
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", data)

	data, err = q.Take()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", data)
}
