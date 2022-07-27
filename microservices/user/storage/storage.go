package storage

import (
	"github.com/go-redis/redis"
	"log"
	redisStorage "user/storage/redis"
	"user/storage/scylla"
	"user/storage/scylla/Connection"
	"user/storage/scylla/Tables"
)

type Storage struct {
	Redis *redisStorage.Redis
	Scylla  *scylla.Scylla //TODO Scylla make connection
}

func NewStorage(clientRedis *redis.Client, clientScylla *Connection.Config) (*Storage, error) {
	newRedis := redisStorage.NewRedis(clientRedis)
	TableNames := new(scylla.TableNames)
	TableNames.User = "users" //TODO TO env
	newScylla,err := scylla.NewScylla(clientScylla, TableNames)
	if err != nil {
		return &Storage{}, err
	}
	makeTable := Tables.NewTables(newScylla.Session)
	err = makeTable.MakeUserTables(newScylla.Keyspace, TableNames.User)
	if err != nil {
		log.Println(err)
	}
	return &Storage{Redis: newRedis, Scylla: newScylla}, nil
}