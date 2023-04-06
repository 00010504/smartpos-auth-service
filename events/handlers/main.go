package handlers

import (
	"github.com/Invan2/invan_auth_service/pkg/logger"
	"github.com/Invan2/invan_auth_service/storage"
)

type EventHandler struct {
	log    logger.Logger
	strgPG storage.StoragePG
	strgR  storage.StorageR
}

func NewHandler(log logger.Logger, strgPG storage.StoragePG, strgR storage.StorageR) *EventHandler {
	return &EventHandler{
		log:    log,
		strgR:  strgR,
		strgPG: strgPG,
	}
}
