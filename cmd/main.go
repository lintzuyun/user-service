package main

import (
	"fmt"
	"log"

	baseApp "github.com/WithWise/golang/app"
	baseLogger "github.com/WithWise/golang/logger"
	"go.uber.org/zap"

	"user-service/pkg/server"
	"user-service/pkg/user"
)

func main() {
	logger := createLogger()
	app := baseApp.NewApp(
		baseApp.WithLogger(logger),
	)

	userService := user.NewService(logger)

	grpcSrv := createGrpcServer(logger, userService)

	go func() {
		if err := grpcSrv.ListenAndServe(); err != nil {
			logger.Error("error running grpc server server", zap.Error(err))

			// a fatal log would call os.Exit(1), so cancel the context to trigger graceful shutdown
			app.Cancel()
		}
	}()

	serviceShutdownOrder := []baseApp.Stopper{grpcSrv}

	app.HandleGracefulShutdown(serviceShutdownOrder)
}

func createLogger() *zap.Logger {
	cfg, err := baseLogger.ReadConfig()
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to read logger config, error=%v", err))
	}

	logger, err := baseLogger.NewLogger(cfg)
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to create logger, error=%v", err))
	}

	return logger
}

func createGrpcServer(
	logger *zap.Logger,
	userSrv *user.UserService,
) *server.Service {
	cfg, err := server.ReadConfig()
	if err != nil {
		logger.Fatal(
			"failed to read config",
			zap.String("service", "grpcServer"),
			zap.Error(err),
		)
	}

	return server.NewServer(cfg, logger, userSrv)
}
