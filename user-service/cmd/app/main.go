package main

import (
	"database/sql"
	"net"
	"os"
	"time"

	pb "github.com/alibekkenny/simple-marketplace/shared/proto/genproto/user"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/app"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/config"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/repository"
	"github.com/alibekkenny/simple-marketplace/user-service/internal/service"
	grpcTransport "github.com/alibekkenny/simple-marketplace/user-service/internal/transport/grpc"
	"github.com/go-playground/validator/v10"
	_ "github.com/lib/pq"
	"github.com/rs/zerolog"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	zerolog.TimeFieldFormat = time.RFC3339
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	logger := zerolog.New(os.Stdout).With().Timestamp().Str("service", "user-service").Logger()

	cfg := config.Load()
	app := app.NewApplication(&logger, cfg)

	db, err := openDB(cfg.DSN)
	if err != nil {
		logger.Error().Err(err).Msg("failed to connect to db")
	}
	defer db.Close()

	repo := repository.NewPostgresRepository(db)
	validator := validator.New()
	service := service.NewUserService(repo, []byte(cfg.JWTKey), validator)

	handler := grpcTransport.NewUserHandler(service, app)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, handler)
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", cfg.Addr)
	if err != nil {
		logger.Error().Err(err).Msgf("failed to listen: %s", cfg.Addr)
	}

	logger.Info().Msgf("UserService running on %v", cfg.Addr)
	if err := grpcServer.Serve(lis); err != nil {
		logger.Error().Err(err).Msg("failed to serve")
	}
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	return db, err
}
