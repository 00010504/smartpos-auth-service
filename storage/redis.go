package storage

import (
	"github.com/Invan2/invan_auth_service/pkg/logger"
	"github.com/go-redis/redis/v8"
)

type storageR struct {
}

type StorageR interface {
}

func NewStorageR(log logger.Logger, db *redis.Client) StorageR {
	return &storageR{}
}
