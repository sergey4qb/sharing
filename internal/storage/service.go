package storage

import (
	"encoding/json"
	"errors"
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
	exist,err := r.Exists(string(user.ID)).Result()
	if err != nil {
		return err
	}
	if exist > 0{
		return errors.New("user already exist")
	}
	js, jserr:= json.Marshal(user)
	if jserr != nil {
		log.Fatal(jserr)
	}
	set := r.Set(string(user.ID), js, r.exp)
	if set.Err() != nil {
		return set.Err()
	}
	return nil
}

func (r *Redis) GetUser(id int) ([]byte, error) {
	fmt.Println("get result")
	result, err := r.Get(string(id)).Result()
	if err != nil {
		return []byte{}, err
	}

	return []byte(result), nil
}

func (r *Redis) Update(id int, data []byte) error {
	fmt.Println("update")
	exist,err := r.Exists(string(id)).Result()
	if err != nil {
		return err
	}
	if exist == 0{
		return errors.New("user doesn't exist")
	}
	getSet := r.GetSet(string(id), data)
	if getSet.Err() != nil {
		return getSet.Err()
	}
	return nil
}

func (r *Redis) Delete(id int) error {
	exist,err := r.Exists(string(id)).Result()
	if err != nil {
		return err
	}
	if exist == 0{
		return errors.New("user doesn't exist")
	}
	del := r.Del(string(id))
	if del.Err() != nil {
		return del.Err()
	}
	return nil
}

