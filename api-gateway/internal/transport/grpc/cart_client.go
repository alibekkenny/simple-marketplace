package grpc

import (
	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCartClient(addr string) (pb.CartServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return pb.NewCartServiceClient(conn), nil
}
