package user

import (
	"context"
	"encoding/json"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"strconv"
	"time"
	"user/grpc/User"
	"user/models"
	"user/types"
)

type Crud struct {
	redis types.UserCrud
	scylla types.UserCrud
	user.UnimplementedUserServer
	user models.User
}

func NewCRUD(RedisCrud, ScyllaCrud types.UserCrud) *Crud {
	return &Crud{redis: RedisCrud, scylla: ScyllaCrud}
}

func (crud *Crud) CreateUser(ctx context.Context, data *user.UserData) (*emptypb.Empty, error) {
	var empty emptypb.Empty
	var userInternal models.User
	fmt.Println("CreateUserMethod") //TODO connect logger

	//userMock := models.User{
	//	ID:        0,
	//	Name:      "name",
	//	Address:   "address",
	//	CreatedAt: time.Time{}.UTC(),
	//}
	userInternal.ID = data.GetId().Id
	userInternal.Name = data.GetName()
	userInternal.Address = data.GetAddress()
	userInternal.CreatedAt = time.Now().Unix()
	userInternal.UpdatedAt = time.Now().Unix()
	fmt.Println(&userInternal)
	err := crud.scylla.Create(&userInternal)
	if err != nil {
		return &empty, err
	}
	err = crud.redis.Create(&userInternal)
	if err != nil {
		return &empty, err
	}

	return &empty, nil
}

func (crud *Crud) GetUserByID(ctx context.Context, id *user.Id) (userResponse *user.UserDataResponse, err error) {
	var response user.UserDataResponse
	var user models.User
	fmt.Println("GetUserCTRL")
	userRedis, err := crud.redis.GetUser(id.GetId())
	if err != nil {
		return nil, err
	}
	jsonErr := json.Unmarshal(userRedis, &user)
	if jsonErr != nil {
		log.Fatal(jsonErr)
		return userResponse, err
	}
	response.Id = id
	response.Name = user.Name
	response.Address = user.Address
	response.CreatedAt = strconv.FormatInt(user.CreatedAt, 10)
	response.UpdatedAt = strconv.FormatInt(user.UpdatedAt, 10)
	return &response, nil
}

func (crud *Crud) UpdateUserByID(ctx context.Context, update *user.UserDataUpdate) (*emptypb.Empty, error) {
	var empty emptypb.Empty
	userInternal := new(models.User)
	userInfo, err := crud.GetUserByID(ctx, update.GetId())
	if err != nil {
		return &empty, err
	}
	createdAttime, err := time.Parse(time.UnixDate, userInfo.GetCreatedAt())
	userInternal.ID = update.Data.GetId().Id
	userInternal.Name = update.Data.GetName()
	userInternal.Address = update.Data.GetAddress()
	userInternal.CreatedAt = createdAttime.Unix()
	userInternal.UpdatedAt = time.Now().Unix()
	marshalled, err := json.Marshal(userInternal)
	fmt.Println(marshalled)
	if err != nil {
		return &empty, err
	}
	err = crud.redis.Update(update.GetId().Id, marshalled)
	if err != nil {
		return &empty, err
	}
	return &empty, nil
}

func (crud *Crud) DeleteUserByID(ctx context.Context, id *user.Id) (*emptypb.Empty, error) {
	var empty emptypb.Empty
	err := crud.redis.Delete(id.Id)
	if err != nil {
		return &empty, err
	}
	return &empty, nil
}

func (crud *Crud) mustEmbedUnimplementedUserServer() {
	panic("implement me")
}
