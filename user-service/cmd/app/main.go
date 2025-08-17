package main

import (
	"database/sql"
	"log"
	"net"
	"os"

	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/user"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/repository"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/service"
	grpcTransport "github.com/alibekkenny/simple-marketplace/user-service/internal/transport/grpc"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	dsn := os.Getenv("DATABASE_URL")
	jwtKey := os.Getenv("JWT_SECRET")
	addr := os.Getenv("SERVICE_ADDR")

	db, err := openDB(dsn)
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

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Printf("UserService running on %s", addr)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, err
}
