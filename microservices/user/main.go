package main

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/gocql/gocql"
	"google.golang.org/grpc"
	"log"
	"net"
	"user/controllers"
	user "user/grpc/User"
	"user/storage"
	"user/storage/scylla/Connection"
)

//import (
//	"bytes"
//	"fmt"
//	"net/http"
//	"os"
//)
//const serverPort = 8000
//func main() {
//
//	requestURL := fmt.Sprintf("http://hello1:%d/hello", serverPort)
//	data := []byte(`{"id":1, "name": "name1"}`)
//	r := bytes.NewReader(data)
//	resp, err := http.Post(requestURL, "application/json", r)
//	if err != nil {
//		fmt.Printf("error making http request: %s\n", err)
//		os.Exit(1)
//	}
//	fmt.Println(resp.StatusCode)
//}

type App struct {
	//Storage *storage.Storage
	Controllers *controllers.UserControllers
	//Order
	//User
}

func NewApp() (*App, error) {
	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", //TODO Перенести в env
		Password: "",
		DB:       0,
	})
	_, err := redis.Ping().Result()
	if err != nil {
		return &App{}, err
	}
	conf := Connection.Config{
		Hosts:        []string{"localhost:9042"},
		Keyspace:     "testnew",
		Strategy:     "SimpleStrategy",
		RF:           1,
		ProtoVersion: 4,
		CL:           gocql.Quorum,
		NumConns:     10,
	}
	store, err := storage.NewStorage(redis, &conf)
	if err != nil {
		return &App{}, err
	}
	ctrls := controllers.NewUserControllers(store.Redis, store.Scylla)
	return &App{Controllers: ctrls}, nil
}

func main() {

	app, err := NewApp()
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051)) //TODO Перенести в env
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	user.RegisterUserServer(s, app.Controllers)
	//pb.RegisterUserServer(s, &training_grpc.Server{})
	log.Printf("Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
	//user := models.User{
	//	ID:        0,
	//	Name:      "name",
	//	Address:   "address",
	//	CreatedAt: time.Time{}.UTC(),
	//}
	//
}
