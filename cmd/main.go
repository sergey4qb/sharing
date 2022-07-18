package main

import (
	"context"
	"log"
	"time"
	"tr_redis_ws/internal/server"
	"tr_redis_ws/internal/storage"
)

func main() {
	ctx:= context.Background()
	newStorage := storage.New().Redis
	storageRedis := storage.NewRedis(newStorage, time.Hour)
	serv := server.NewServer(ctx,storageRedis)
	if err := serv.Start(); err != nil {
		log.Fatal(err)
	}
}
