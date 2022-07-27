package grpc_controllers

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"log"
	ws "ws/gRPC/User"
	"ws/models"
)

const (
	USERADDR = "user:50051" // TODO Перенести в env
)

type Client struct {
	ws.UserClient
	ctx context.Context
}

func (c Client) Create(request models.User) error {
	userData := new(ws.UserData)
	userData.Id = new(ws.Id)
	fmt.Println(userData.Id)
	userData.Id.Id = request.ID
	userData.Name = request.Name
	userData.Address = request.Address

	fmt.Println("init create")
	_, err := c.CreateUser(c.ctx, userData)
	if err != nil {
		return errors.New(fmt.Sprintf("%v error in createUser", err))
	}
	return nil
}

func (c Client) Get(id uint32) (userResponse models.User, err error) {
	userId := new(ws.Id)
	userId.Id = uint32(id)
	user, err := c.GetUserByID(c.ctx, userId)
	if err != nil {
		return models.User{}, err
	}
	userResponse.ID = user.Id.Id
	userResponse.Name = user.Name
	userResponse.Address = user.Address
	return userResponse, nil
}

func (c Client) Update(id uint32, u models.User) error {
	var userDataUpdate ws.UserDataUpdate
	userDataUpdate.Id = new(ws.Id)
	userDataUpdate.Data = new(ws.UserData)
	userDataUpdate.Data.Id = new(ws.Id)
	userDataUpdate.Id.Id = id
	userDataUpdate.Data.Id.Id = u.ID
	userDataUpdate.Data.Name = u.Name
	userDataUpdate.Data.Address = u.Address

	_, err := c.UpdateUserByID(c.ctx, &userDataUpdate)
	if err != nil {
		return err
	}
	return nil
}

func (c Client) Delete(id uint32) error {
	userId := new(ws.Id)
	userId.Id = uint32(id)
	_, err := c.DeleteUserByID(c.ctx, userId)
	if err != nil {
		return err
	}
	return nil
}

func NewClient(ctx context.Context) Client {
	conn, err := grpc.Dial(USERADDR, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	c := ws.NewUserClient(conn)

	return Client{c, ctx}
}
