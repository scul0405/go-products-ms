package server

import (
	"Go-ProductMS/config"
	grpcproduct "Go-ProductMS/internal/product/delivery/grpc"
	"Go-ProductMS/internal/product/repository"
	"Go-ProductMS/internal/product/usecase"
	"Go-ProductMS/pkg/logger"
	productService "Go-ProductMS/proto/product"
	"context"
	"github.com/go-playground/validator/v10"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/pkg/errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"time"
)

type server struct {
	log logger.Logger
	cfg *config.Config
}

func NewServer(
	log logger.Logger,
	cfg *config.Config,
) *server {
	return &server{
		log: log,
		cfg: cfg,
	}
}

func (s *server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	validate := validator.New()

	productMgRepo := repository.NewProductMongoRepo()
	productUsecase := usecase.NewProductUsecase(productMgRepo, s.log)

	l, err := net.Listen("tcp", s.cfg.Server.Port)
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	defer l.Close()

	grpcServer := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: s.cfg.Server.MaxConnectionIdle * time.Minute,
		MaxConnectionAge:  s.cfg.Server.MaxConnectionAge * time.Minute,
		Timeout:           s.cfg.Server.Timeout * time.Second,
		Time:              s.cfg.Server.Timeout * time.Minute,
	}),
		grpc.ChainUnaryInterceptor(
			grpcrecovery.UnaryServerInterceptor(),
		),
	)

	productSvc := grpcproduct.NewProductService(s.log, productUsecase, validate)
	productService.RegisterProductsServiceServer(grpcServer, productSvc)

	go func() {
		s.log.Infof("starting grpc server at %s", s.cfg.Server.Port)
		s.log.Fatal(grpcServer.Serve(l))
	}()

	if s.cfg.Server.Development {
		reflection.Register(grpcServer)
	}

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	select {
	case <-quit:
		s.log.Errorf("signal.Notify: %v", quit)
	case done := <-ctx.Done():
		s.log.Errorf("ctx.Done: %v", done)
	}

	grpcServer.GracefulStop()
	s.log.Info("Server Exited Properly")

	return nil
}
