package server

import (
	"context"
	"net/http"
	"ws/grpc_controllers"
	"ws/server/controllers"
)

type Server struct {
	Controllers *controllers.Controllers
	Client *grpc_controllers.Client
}

func NewServer(ctx context.Context) *Server {
	//grpcCtrls := user.NewGrpcControllers()
	grpcClient := grpc_controllers.NewClient(ctx)
	return &Server{
		Controllers: controllers.NewControllers(ctx, grpcClient),
	}
}

func (r *Server) Start() error {
	r.Routes()
	return http.ListenAndServe(":8080", nil)
}
func (r *Server) Routes() {
	http.HandleFunc("/", r.Controllers.WebSocket)
}
