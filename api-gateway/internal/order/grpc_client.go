package order

import (
	"context"
	"log"
	"time"

	pb "github.com/elliaaan/proto-gen/pb/order/github.com/elliaaan/proto-gen/pb/order"
	"google.golang.org/grpc"
)

type OrderClient struct {
	client pb.OrderServiceClient
}

func NewOrderClient(address string) *OrderClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(time.Second*5))
	if err != nil {
		log.Fatalf("could not connect to order service: %v", err)
	}

	return &OrderClient{
		client: pb.NewOrderServiceClient(conn),
	}
}

func (o *OrderClient) CreateOrder(order *pb.Order) (*pb.OrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return o.client.CreateOrder(ctx, &pb.CreateOrderRequest{Order: order})
}

func (o *OrderClient) GetOrderByID(id uint64) (*pb.OrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return o.client.GetOrderByID(ctx, &pb.GetOrderRequest{Id: id})
}

func (o *OrderClient) UpdateOrder(order *pb.Order) (*pb.OrderResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return o.client.UpdateOrder(ctx, &pb.UpdateOrderRequest{Order: order})
}

func (o *OrderClient) DeleteOrder(id uint64) (*pb.Empty, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return o.client.DeleteOrder(ctx, &pb.DeleteOrderRequest{Id: id})
}

func (o *OrderClient) ListOrders() (*pb.ListOrdersResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	return o.client.ListOrders(ctx, &pb.ListOrdersRequest{})
}
