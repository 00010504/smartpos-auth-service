package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"genproto/common"

	"github.com/Invan2/invan_auth_service/pkg/logger"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func (e *EventHandler) CreateCompany(ctx context.Context, event *kafka.Message) error {

	var req common.CompanyCreatedModel

	if err := json.Unmarshal(event.Value, &req); err != nil {
		return err
	}

	e.log.Info("company created", logger.Any("event", req))

	if err := e.strgPG.Company().Upsert(&req); err != nil {
		return err
	}

	e.log.Info("shop created", logger.Any("event", req))

	if req.Shop == nil {
		return errors.New("nimadir")
	}

	if err := e.strgPG.Shop().Upsert(req.Shop); err != nil {
		return err
	}

	return nil

}
