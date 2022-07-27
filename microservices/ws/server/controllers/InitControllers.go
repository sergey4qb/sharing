package controllers

import (
	"context"
	"ws/grpc_controllers"
)

type Controllers struct {
	grpc grpc_controllers.Client
	ctx context.Context
}

func NewControllers(ctx context.Context,client grpc_controllers.Client) *Controllers {

	return &Controllers{
		grpc: 	client,
		ctx: ctx,
	}
}