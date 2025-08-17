package grpc

import (
	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewProductClient(addr string) (pb.ProductServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return pb.NewProductServiceClient(conn), nil
}
