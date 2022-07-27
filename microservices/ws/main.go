package main

import (
	"context"
	"log"
	"ws/server"
)

func main() {
	ctx := context.Background()

	serv := server.NewServer(ctx)
	if err := serv.Start(); err != nil {
		log.Fatal(err)
	}
}
