package controllers

import (
	"user/controllers/user"
	//user2 "user/grpc/User"
	"user/types"
)

type UserControllers struct {
	*user.Crud
}

func NewUserControllers(RedisCrud, ScyllaCrud types.UserCrud) *UserControllers {
	userCRUD := user.NewCRUD(RedisCrud, ScyllaCrud)
	return &UserControllers{userCRUD}
}
