package interceptors

import (
	"Go-ProductMS/config"
	"Go-ProductMS/pkg/logger"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"time"
)

type InterceptorManager struct {
	logger logger.Logger
	cfg    *config.Config
}

func NewInterceptorManager(logger logger.Logger, cfg *config.Config) *InterceptorManager {
	return &InterceptorManager{logger: logger, cfg: cfg}
}

// Logger Interceptor
func (im *InterceptorManager) Logger(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
	start := time.Now()
	md, _ := metadata.FromIncomingContext(ctx)
	reply, err := handler(ctx, req)
	im.logger.Infof("Method: %s, Time: %v, Metadata: %v, Err: %v", info.FullMethod, time.Since(start), md, err)

	return reply, err
}
