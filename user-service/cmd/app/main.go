package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"

	pb "github.com/alibekkenny/simple-marketplace/user-service/genproto"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/repository"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/service"
	grpcTransport "github.com/alibekkenny/simple-marketplace/user-service/internal/transport/grpc"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("failed to load env")
	}

	dsn := os.Getenv("DSN")
	jwtKey := os.Getenv("JWT_SECRET")

	fmt.Println(dsn)

	db, err := provideDB(dsn)
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	defer db.Close()

	repo := repository.NewPostgresRepository(db)
	validator := validator.New()
	service := service.NewUserService(repo, []byte(jwtKey), validator)
	handler := grpcTransport.NewUserHandler(service)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("UserService running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func provideDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, err
}
