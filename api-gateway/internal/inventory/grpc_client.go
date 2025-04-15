package inventory

import (
	"context"
	"log"
	"time"

	pb "github.com/elliaaan/proto-gen/pb/inventory/github.com/elliaaan/proto-gen/pb/inventory"
	"google.golang.org/grpc"
)

type Client struct {
	conn   *grpc.ClientConn
	Client pb.InventoryServiceClient
}

func NewInventoryClient(address string) *Client {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(3*time.Second))
	if err != nil {
		log.Fatalf("Failed to connect to inventory-service: %v", err)
	}

	client := pb.NewInventoryServiceClient(conn)
	return &Client{conn: conn, Client: client}
}

func (c *Client) ListProducts() ([]*pb.Product, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()

	res, err := c.Client.ListProducts(ctx, &pb.ListProductsRequest{})
	if err != nil {
		return nil, err
	}

	return res.Products, nil
}
