package main

import (
	"log"
	"net"

	"github.com/elliaaan/statistics-service/db"
	"github.com/elliaaan/statistics-service/internal/statistics"
	"github.com/elliaaan/statistics-service/models"

	statisticspb "github.com/elliaaan/proto-gen/pb/statistics"

	"github.com/nats-io/nats.go"
	"google.golang.org/grpc"
)

func main() {
	// 1. Подключаем базу данных
	database := db.InitDB()
	if err := database.AutoMigrate(&models.Event{}); err != nil {
		log.Fatalf("Ошибка миграции таблицы Event: %v", err)
	}

	// 2. Подключаемся к NATS
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatalf("Ошибка подключения к NATS: %v", err)
	}
	defer nc.Drain()

	// 3. Подписка на события от order/inventory сервисов
	statistics.SubscribeToOrderCreated(nc, database)
	// TODO: statistics.SubscribeToInventoryUpdated(nc, database) если будет нужно

	// 4. Создаём сервис и gRPC сервер
	service := &statistics.Service{DB: database}
	server := &statistics.GRPCServer{Service: service}

	grpcServer := grpc.NewServer()
	statisticspb.RegisterStatisticsServiceServer(grpcServer, server)

	listener, err := net.Listen("tcp", ":8082")
	if err != nil {
		log.Fatalf("Ошибка запуска gRPC-сервера: %v", err)
	}

	log.Println("📊 Statistics Service запущен на порту 8082")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Ошибка gRPC Serve: %v", err)
	}
}
