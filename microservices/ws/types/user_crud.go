package types

import "ws/models"

type UserCrud interface {
	Create(request models.User) error
	Get(id uint32) (models.User, error)
	Update(id uint32, user models.User) error
	Delete(id uint32) error
}
