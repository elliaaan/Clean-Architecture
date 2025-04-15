package main

import (
	"log"
	"net"
	"order-service/db"
	"order-service/internal/order"

	pb "github.com/elliaaan/proto-gen/pb/order/github.com/elliaaan/proto-gen/pb/order"
	"google.golang.org/grpc"
)

func main() {
	// 1. Подключаем базу данных
	database := db.InitDB()

	// 2. Инициализируем Repository → Service → gRPC Server
	repo := &order.Repository{DB: database}
	service := &order.Service{Repo: repo}
	server := &order.GRPCServer{Service: service}

	// 3. Настраиваем gRPC сервер
	grpcServer := grpc.NewServer()
	pb.RegisterOrderServiceServer(grpcServer, server)

	// 4. Слушаем порт 8081
	listener, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Order Service running on port 8081")
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
