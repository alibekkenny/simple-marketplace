package app

import (
	"log"
	"net"

	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/product"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func NewGRPCServer(addr string, categoryHandler pb.CategoryServiceServer, productHandler pb.ProductServiceServer, productOfferHandler pb.ProductOfferServiceServer) error {
	grpcServer := grpc.NewServer()

	pb.RegisterCategoryServiceServer(grpcServer, categoryHandler)
	pb.RegisterProductServiceServer(grpcServer, productHandler)
	pb.RegisterProductOfferServiceServer(grpcServer, productOfferHandler)

	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}

	log.Printf("ProductService running on %v\n", addr)
	if err := grpcServer.Serve(lis); err != nil {
		return err
	}

	return nil
}
