package storage

import (
	"encoding/json"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
	"tr_redis_ws/internal/models"
)


type Redis struct {
	*redis.Client
	exp     time.Duration
}

func NewRedis(store *redis.Client, duration time.Duration) *Redis {
	return &Redis{
		Client:store,
		exp:     duration,
	}
}

func (r *Redis) Create(user models.User) error {
	fmt.Println("createUserRedis")
	js, jserr:= json.Marshal(user)
	if jserr != nil {
		log.Fatal(jserr)
	}
	err := r.Set(string(user.ID), js, r.exp)
	if err != nil {
		return err.Err()
	}
	return nil
}

func (r *Redis) GetUser(id int) ([]byte, error) {
	result, err := r.Get(string(id)).Result()
	if err != nil {
		return []byte{}, err
	}
	fmt.Println(result, "get result")

	return []byte(result), nil
}

func (r *Redis) Update(data []byte) error {
	panic("implement me")
}

func (r *Redis) Delete(id int) (string, error) {
	panic("implement me")
}

