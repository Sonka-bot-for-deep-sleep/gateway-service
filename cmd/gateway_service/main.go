package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Sonka-bot-for-deep-sleep/common/pkg/logger"
	"github.com/Sonka-bot-for-deep-sleep/gateway_service/application/config"
	pb "github.com/Sonka-bot-for-deep-sleep/proto_files/api"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	log, err := logger.New()
	if err != nil {
		fmt.Println(fmt.Errorf("Failed create logger instance: %w", err))
		return
	}

	ctx := context.Background()
	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	cfg, err := config.MustLoad()
	if err != nil {
		log.Error("Failed load eniviroment variables", zap.Error(err))
		return
	}
	if err = pb.RegisterUsersServiceHandlerFromEndpoint(ctx, mux, cfg.UsersUrlService, opts); err != nil {
		log.Error("Failed register gateway for users service", zap.Error(err))
		return
	}

	if err = pb.RegisterTimeServiceHandlerFromEndpoint(ctx, mux, cfg.TimeUrlService, opts); err != nil {
		log.Error("Failed register gateway for time service", zap.Error(err))
		return
	}

	log.Info("Server start work", zap.String("url", cfg.GatewayURL))
	if err := http.ListenAndServe(cfg.GatewayURL, mux); err != nil {
		log.Error("Failed create listener http server", zap.Error(err), zap.String("url", cfg.GatewayURL))
		return
	}
}
