package server

import (
	"context"
	"net/http"
	"tr_redis_ws/internal/server/controllers"
	"tr_redis_ws/internal/storage"
)


type Server struct {
	//Storage *storage.Redis
	Controllers *controllers.Controllers
}

func NewServer(ctx context.Context, redis *storage.Redis) *Server {
	//repo := storage.New()
	//newRedis := storage.NewRedis(repo.Client, time.Hour)
	return &Server{
		 //Storage: redis,
		 Controllers: controllers.NewControllers(redis),
	}
}

func (r *Server) Start() error{
	r.Routes()
	return http.ListenAndServe(":8080", nil)
}
func (r *Server) Routes() {
	http.HandleFunc("/", r.Controllers.WebSocket)
}