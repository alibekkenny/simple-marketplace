package grpc

import (
	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewOrderClient(addr string) (pb.OrderServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return pb.NewOrderServiceClient(conn), nil
}
