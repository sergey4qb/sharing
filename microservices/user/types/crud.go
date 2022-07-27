package types

import "user/models"

type UserCrud interface {
	Create(user *models.User) error
	GetUser(id uint32) ([]byte, error)
	Update(userID uint32, data []byte) error
	Delete(id uint32) error
}
