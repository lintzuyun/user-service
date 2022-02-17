package server

import (
	"fmt"
	"net"

	"user-service/pkg/user"

	pb "github.com/lintzuyun/protorepo-user-go-practice/v1"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// type Service struct {
// 	config Config
// 	server *grpc.Server

// 	pb.UnimplementedUserServiceServer

// 	serviceOptions
// }

// type serviceOptions struct {
// 	loggerOption
// }

type Service struct {
	config      Config
	server      *grpc.Server
	logger      *zap.Logger
	userService *user.UserService
	pb.UnimplementedUserServiceServer
}

func NewServer(cfg Config, logger *zap.Logger, userSrv *user.UserService) *Service {
	grpcSrv := grpc.NewServer()

	// register the gRPC server for reflection to expose available endpoints
	reflection.Register(grpcSrv)

	svc := &Service{
		config:      cfg,
		server:      grpcSrv,
		logger:      logger,
		userService: userSrv,
	}

	pb.RegisterUserServiceServer(grpcSrv, svc)

	return svc
}

// func NewServer(cfg Config, opts ...Option) *Service {
// 	srv := grpc.NewServer()

// 	// register the gRPC server for reflection to expose available endpoints
// 	reflection.Register(srv)

// 	svc := &Service{
// 		config: cfg,
// 		server: srv,
// 		serviceOptions: serviceOptions{
// 			loggerOption: loggerOption{
// 				logger: zap.NewNop(),
// 			},
// 		},
// 	}

// 	pb.RegisterUserServiceServer(srv, svc)

// 	return svc.withOptions(opts...)
// }

// func WithLogger(l *zap.Logger) Option {
// 	return loggerOption{logger: l}
// }

func (s *Service) Stop() {
	s.server.GracefulStop()
}

func (s *Service) ListenAndServe() error {
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%s", s.config.Host, s.config.Port))
	if err != nil {
		return errors.Errorf("failed to listen: %v", err)
	}

	s.logger.Info("starting grpc server", zap.String("address", lis.Addr().String()))
	if err2 := s.server.Serve(lis); err2 != nil {
		return errors.Errorf("failed to serve: %v", err2)
	}

	return nil
}
