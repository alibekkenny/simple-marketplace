package app

import (
	"log"
	"net"

	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/order"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(addr string, cartHandler pb.CartServiceServer, orderHandler pb.OrderServiceServer) error {
	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, orderHandler)
	pb.RegisterCartServiceServer(grpcServer, cartHandler)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("ProductService running on %v\n", addr)
	grpcServer.Serve(lis)

	return nil
}
