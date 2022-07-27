package redis

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis"
	"log"
	"time"
	"user/models"
)

type Redis struct{
	 *redis.Client
}

func NewRedis(conn *redis.Client) *Redis {
	return &Redis{conn}
}

func (redis *Redis) Create(user *models.User) error {
	fmt.Println("createUserRedis")
	exist,err := redis.Exists(string(user.ID)).Result()
	if err != nil {
		return err
	}
	if exist > 0{
		return errors.New("user already exist")
	}//TODO make status collection env file
	//user.CreatedAt = time.Now().Unix() // TODO Set location in env
	//user.UpdatedAt = time.Now().Unix() // TODO Set location in env
	js, jserr:= json.Marshal(user)
	if jserr != nil {
		log.Fatal(jserr)
	}
	set := redis.Set(string(user.ID), js, time.Hour)
	if set.Err() != nil {
		return set.Err()
	}
	return nil
}

func (redis *Redis) GetUser(id uint32) ([]byte, error) {
	fmt.Println("get result")
	result, err := redis.Get(string(id)).Result()
	if err != nil {
		return []byte{}, err
	}

	return []byte(result), nil
}

func (redis *Redis) Update(userID uint32, data []byte) error {
	fmt.Println("update")
	exist,err := redis.Exists(string(userID)).Result()
	if err != nil {
		return err
	}
	if exist == 0{
		return errors.New("user doesn't exist")//TODO make status collection env file
	}

	upd, err := redis.GetSet(string(userID), data).Result()
	if err != nil {
		return err
	}
	fmt.Println(upd)
	return nil
}

func (redis *Redis) Delete(id uint32)  error {
	exist,err := redis.Exists(string(id)).Result()
	if err != nil {
		return err
	}
	if exist == 0{
		return errors.New("user doesn't exist")//TODO make status collection env file
	}
	del := redis.Del(string(id))
	if del.Err() != nil {
		return del.Err()
	}
	return  nil // TODO make status collection env file
}
