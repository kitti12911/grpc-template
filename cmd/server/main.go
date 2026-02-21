package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/kitti12911/lib-util/logger"

	examplehand "grpc-template/internal/feature/example/handler"
	examplerepo "grpc-template/internal/feature/example/repository"
	exampleserv "grpc-template/internal/feature/example/service"
	"grpc-template/internal/server"
)

func main() {
	logger.New(
		logger.WithLevel(logger.LevelInfo),
		logger.WithServiceName("grpc-template"),
	)

	repo := examplerepo.NewExampleRepository()
	svc := exampleserv.NewExampleService(repo)
	handler := examplehand.NewExampleHandler(svc)

	srv, err := server.NewGRPCServer(50051, handler)

	if err != nil {
		slog.Error("failed to create gRPC server", "error", err)
		os.Exit(1)
	}

	go func() {
		if err := srv.Start(); err != nil {
			slog.Error("gRPC server error", "error", err)
			os.Exit(1)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	slog.Info("shutting down gRPC server...")
	srv.Stop()
	slog.Info("server stopped")
}
