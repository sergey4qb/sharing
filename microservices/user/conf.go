package main
//
//import (
//	"github.com/go-redis/redis"
//	"user/storage"
//)
//
//type App struct {
//	Storage *storage.Storage
//	//Order
//	//User
//}
//
//func NewApp() (*App, error) {
//	r := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379", //TODO Перенести в env
//		Password: "",
//		DB:       0,
//	})
//	_, err := r.Ping().Result()
//	if err != nil {
//		return &App{}, err
//	}
//	store := storage.NewStorage(r)
//	return &App{Storage: store}, nil
//}



