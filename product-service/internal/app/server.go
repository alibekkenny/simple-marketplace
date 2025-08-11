package app

import (
	"log"
	"net"

	pb "github.com/alibekkenny/simple-marketplace/product-service/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(addr string, categoryHandler pb.CategoryServiceServer) error {
	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryHandler)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("UserService running on %v\n", addr)
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
