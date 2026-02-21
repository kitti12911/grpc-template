package server

import (
	"fmt"
	"log/slog"
	"net"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/recovery"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"

	examplev1 "grpc-template/gen/grpc/example/v1"
	examplehand "grpc-template/internal/feature/example/handler"
	"grpc-template/internal/server/interceptor"
)

type GRPCServer struct {
	server   *grpc.Server
	listener net.Listener
	health   *health.Server
}

func NewGRPCServer(port int, exampleHandler *examplehand.ExampleHandler) (*GRPCServer, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return nil, fmt.Errorf("failed to listen on port %d: %w", port, err)
	}

	recoveryOpt := recovery.WithRecoveryHandler(func(p any) error {
		return status.Errorf(codes.Internal, "internal server error: %v", p)
	})

	srv := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			interceptor.ErrorHandler(),
			recovery.UnaryServerInterceptor(recoveryOpt),
		),
		grpc.ChainStreamInterceptor(
			recovery.StreamServerInterceptor(recoveryOpt),
		),
	)

	examplev1.RegisterExampleServiceServer(srv, exampleHandler)

	healthServer := health.NewServer()
	healthv1.RegisterHealthServer(srv, healthServer)
	healthServer.SetServingStatus("", healthv1.HealthCheckResponse_SERVING)
	healthServer.SetServingStatus(
		examplev1.ExampleService_ServiceDesc.ServiceName,
		healthv1.HealthCheckResponse_SERVING,
	)

	reflection.Register(srv)

	return &GRPCServer{
		server:   srv,
		listener: listener,
		health:   healthServer,
	}, nil
}

func (s *GRPCServer) Start() error {
	slog.Info("gRPC server listening", "addr", s.listener.Addr().String())
	return s.server.Serve(s.listener)
}

func (s *GRPCServer) Stop() {
	s.health.Shutdown()
	s.server.GracefulStop()
}
