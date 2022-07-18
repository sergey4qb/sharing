package controllers

import "tr_redis_ws/internal/storage"

type Controllers struct {
	storage *storage.Redis
}

func NewControllers(redis *storage.Redis) *Controllers {
	return &Controllers{storage: redis}
}