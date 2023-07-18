package server

import (
	"Go-ProductMS/config"
	"Go-ProductMS/internal/interceptors"
	grpcproduct "Go-ProductMS/internal/product/delivery/grpc"
	"Go-ProductMS/internal/product/delivery/kafka"
	"Go-ProductMS/internal/product/repository"
	"Go-ProductMS/internal/product/usecase"
	"Go-ProductMS/pkg/logger"
	productService "Go-ProductMS/proto/product"
	"context"
	"crypto/tls"
	"fmt"
	"github.com/go-playground/validator/v10"
	grpcrecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"github.com/labstack/echo/v4"
	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"net"
	"os"
	"os/signal"
	"time"
)

const (
	certFile = "ssl/server.crt"
	keyFile  = "ssl/server.pem"
)

type server struct {
	log     logger.Logger
	cfg     *config.Config
	tracing opentracing.Tracer
	mongoDB *mongo.Client
}

func NewServer(
	log logger.Logger,
	cfg *config.Config,
	mongoDB *mongo.Client,
	tracer opentracing.Tracer,
) *server {
	return &server{
		log:     log,
		cfg:     cfg,
		mongoDB: mongoDB,
		tracing: tracer,
	}
}

func (s *server) Run() error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	validate := validator.New()

	productMgRepo := repository.NewProductMongoRepo(s.mongoDB)
	productUsecase := usecase.NewProductUsecase(productMgRepo, s.log, validate)

	im := interceptors.NewInterceptorManager(s.log, s.cfg)

	l, err := net.Listen("tcp", fmt.Sprintf("localhost:%s", s.cfg.Server.Port))
	if err != nil {
		return errors.Wrap(err, "failed to listen")
	}
	defer l.Close()

	cert, err := tls.LoadX509KeyPair(certFile, keyFile)
	if err != nil {
		s.log.Fatalf("failed to load key pair: %s", err)
	}

	s.log.Info("CERT loaded")

	grpcServer := grpc.NewServer(
		grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		grpc.KeepaliveParams(keepalive.ServerParameters{
			MaxConnectionIdle: s.cfg.Server.MaxConnectionIdle * time.Minute,
			MaxConnectionAge:  s.cfg.Server.MaxConnectionAge * time.Minute,
			Timeout:           s.cfg.Server.Timeout * time.Second,
			Time:              s.cfg.Server.Timeout * time.Minute,
		}),
		grpc.ChainUnaryInterceptor(
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_opentracing.UnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpcrecovery.UnaryServerInterceptor(),
			im.Logger,
		),
	)

	productSvc := grpcproduct.NewProductService(s.log, productUsecase, validate)
	productService.RegisterProductsServiceServer(grpcServer, productSvc)
	grpc_prometheus.Register(grpcServer)

	productsCG := kafka.NewProductsConsumerGroup(s.cfg.Kafka.Brokers, "products_group", s.log, s.cfg, productUsecase)
	productsCG.RunConsumers(ctx, cancel)

	go func() {
		s.log.Infof("starting grpc server at %s", s.cfg.Server.Port)
		s.log.Fatal(grpcServer.Serve(l))
	}()

	if s.cfg.Server.Development {
		reflection.Register(grpcServer)
	}

	metricsServer := echo.New()
	go func() {
		metricsServer.GET("/metrics", echo.WrapHandler(promhttp.Handler()))
		s.log.Infof("Metrics server is running on port: %s", s.cfg.Metrics.Port)
		if err := metricsServer.Start(s.cfg.Metrics.Port); err != nil {
			s.log.Error(err)
			cancel()
		}
	}()

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
