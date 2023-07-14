package grpc

import (
	"Go-ProductMS/internal/product/usecase"
	"Go-ProductMS/pkg/logger"
	productSvc "Go-ProductMS/proto/product"
	"context"
	"github.com/go-playground/validator/v10"
)

type productService struct {
	log            logger.Logger
	productUsecase usecase.ProductUseCase
	validate       *validator.Validate
	productSvc.UnimplementedProductsServiceServer
}

func NewProductService(
	log logger.Logger,
	productUsecase usecase.ProductUseCase,
	validate *validator.Validate,
) *productService {
	return &productService{
		log:            log,
		productUsecase: productUsecase,
		validate:       validate,
	}
}

func (p *productService) Create(ctx context.Context, req *productSvc.UpdateReq) (*productSvc.UpdateRes, error) {
	panic("implement me")
}

func (p *productService) Update(ctx context.Context, req *productSvc.UpdateReq) (*productSvc.UpdateRes, error) {
	panic("implement me")
}

func (p *productService) GetByID(ctx context.Context, req *productSvc.GetByIDReq) (*productSvc.GetByIDRes, error) {
	panic("implement me")
}

func (p *productService) Search(ctx context.Context, req *productSvc.SearchReq) (*productSvc.SearchRes, error) {
	panic("implement me")
}
