package grpc

import (
	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/user"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewUserClient(addr string) (pb.UserServiceClient, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return pb.NewUserServiceClient(conn), nil
}
