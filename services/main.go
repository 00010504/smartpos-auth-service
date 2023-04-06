package services

import (
	"genproto/notification_service"

	"github.com/Invan2/invan_auth_service/config"
	"github.com/Invan2/invan_auth_service/events"
	"github.com/Invan2/invan_auth_service/pkg/logger"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type gRPCServices struct {
	log   logger.Logger
	cfg   *config.Config
	kafka events.PubSubServer

	notificationService notification_service.SMSClient
}

type GRPCServices interface {
	NotificationService() notification_service.SMSClient
}

func NewGrpcServices(log logger.Logger, cfg *config.Config, kafka events.PubSubServer) (GRPCServices, error) {

	res := &gRPCServices{
		log:   log,
		cfg:   cfg,
		kafka: kafka,
	}

	connNotification, err := grpc.Dial(cfg.NotificationService, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, errors.Wrap(err, "error while creating notification service client")
	}

	res.notificationService = notification_service.NewSMSClient(connNotification)

	return res, nil
}

func (g *gRPCServices) NotificationService() notification_service.SMSClient {
	return g.notificationService
}
