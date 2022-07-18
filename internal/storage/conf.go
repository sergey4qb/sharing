package storage

import "github.com/go-redis/redis"

type Storage struct {
	Redis *redis.Client
}



func New() *Storage {
	return &Storage{redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379", //TODO Перенести в env
			Password: "",
			DB:       0,
		})}
}
