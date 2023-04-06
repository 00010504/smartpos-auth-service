package redis

import (
	"github.com/Invan2/invan_auth_service/pkg/logger"
	"github.com/go-redis/redis/v8"
)

type permissionRepo struct {
	db  *redis.Client
	log logger.Logger
}
