package subjects

import (
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func Pub(nc *nats.Conn)  {
	for {
		msg, err:= nc.Request("help", []byte("ping1"), 5* time.Second)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(msg.Data))
		time.Sleep(3 * time.Second)
	}
}
func Sub(nc *nats.Conn,r chan *nats.Msg)  {
	_, err := nc.ChanSubscribe("help", r)
	if err != nil {
		log.Fatal(err)
	}

	for msg := range r {
		fmt.Println(string(msg.Data))
		nc.Publish(msg.Reply, []byte("pong1"))
	}
}
