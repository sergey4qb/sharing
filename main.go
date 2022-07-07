package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"test/subjects"
	"time"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL)
	fmt.Println("Connected to NATS at:", nc.ConnectedUrl())
	ch := make(chan *nats.Msg, 64)
	go func() {
		subjects.Pub(nc)
	}()
	for{
		subjects.Sub(nc, ch)
		time.Sleep(3* time.Second)
	}

}

