package interfaces

import "tr_redis_ws/internal/models"

type UserStorage interface {
	Create(user models.User) error
	GetUser(id int) ([]byte, error)
	Update(data []byte) error
	Delete(id int) (string, error)
}
