package main

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"

	"genproto/auth_service"

	"github.com/Invan2/invan_auth_service/config"
	"github.com/Invan2/invan_auth_service/events"
	"github.com/Invan2/invan_auth_service/pkg/logger"
	"github.com/Invan2/invan_auth_service/services"
	"github.com/Invan2/invan_auth_service/services/listeners"
	"github.com/Invan2/invan_auth_service/storage"
	"github.com/confluentinc/confluent-kafka-go/kafka"

	"github.com/go-redis/redis/v8"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {

	ctx := context.Background()

	cfg := config.Load()
	log := logger.New(cfg.LogLevel, cfg.ServiceName)

	log.Info("config", logger.Any("config", cfg), logger.Any("env", os.Environ()))

	postgresURL := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresDatabase,
	)

	psqlConn, err := sqlx.Connect("postgres", postgresURL)
	if err != nil {
		log.Error("postgres", logger.Error(err))
		return
	}

	defer func() {
		if err := psqlConn.Close(); err != nil {
			log.Error("error while close postgres connectino", logger.Error(err))
			return
		}

		fmt.Println("postgres connection closed")
	}()

	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort),
		DB:       cfg.RedisDB,
		Password: cfg.RedisPassword,
	})

	defer func() {
		if err := redisClient.Close(); err != nil {
			log.Error("error while close redis connectino", logger.Error(err))
			return
		}
	}()

	statusCMd := redisClient.Ping(context.Background())
	if _, err := statusCMd.Result(); err != nil {
		log.Error("error while connecting redis", logger.Error(err))
		// return
	}

	log.Info("connected to redis")

	strgP := storage.NewStoragePg(log, psqlConn)
	strgR := storage.NewStorageR(log, redisClient)

	conf := kafka.ConfigMap{
		"bootstrap.servers":                     cfg.KafkaUrl,
		"group.id":                              config.ConsumerGroupID,
		"auto.offset.reset":                     "earliest",
		"go.events.channel.size":                1000000,
		"socket.keepalive.enable":               true,
		"metadata.max.age.ms":                   900000,
		"metadata.request.timeout.ms":           30000,
		"retries":                               1000000,
		"message.timeout.ms":                    300000,
		"socket.timeout.ms":                     30000,
		"max.in.flight.requests.per.connection": 5,
		"heartbeat.interval.ms":                 3000,
		"enable.idempotence":                    true,
	}

	log.Info("kafka config", logger.Any("config", conf))

	producer, err := kafka.NewProducer(&conf)
	if err != nil {
		log.Error("error while creating producer")
		return
	}

	defer func() {
		producer.Close()
	}()

	consumer, err := kafka.NewConsumer(&conf)
	if err != nil {
		log.Error("error while creating consumer", logger.Error(err))
		return
	}

	pubSub, err := events.NewPubSubServer(log, producer, consumer, strgP, strgR)
	if err != nil {
		log.Error("error while creating pub sub server", logger.Error(err))
		return
	}

	go func() {
		if err := pubSub.Run(ctx); err != nil {
			log.Error("error while run pubsub server", logger.Error(err))
			return
		}
	}()

	grpcServices, err := services.NewGrpcServices(log, &cfg, pubSub)
	if err != nil {
		log.Error("error while create grpc services", logger.Error(err))
		return
	}

	server := grpc.NewServer()
	auth_service.RegisterAuthServer(server, listeners.NewAuthService(log, cfg, grpcServices, strgP, strgR, pubSub))

	lis, err := net.Listen("tcp", fmt.Sprintf("%s%s", cfg.HttpHost, cfg.HttpPort))
	if err != nil {
		log.Error("error while create listener:", logger.Error(err))
		return
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		<-c
		fmt.Println("Gracefully shutting down...")
		server.GracefulStop()
		if err := pubSub.Shutdown(); err != nil {
			log.Error("error while shutdown pub sub server")
		}
	}()

	if err := server.Serve(lis); err != nil {
		log.Fatal("serve", logger.Error(err))
		return
	}
}
