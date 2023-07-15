package grpc

import (
	"Go-ProductMS/internal/models"
	"Go-ProductMS/internal/product/usecase"
	grpcErrors "Go-ProductMS/pkg/grpc_errors"
	"Go-ProductMS/pkg/logger"
	productSvc "Go-ProductMS/proto/product"
	"context"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (p *productService) Create(ctx context.Context, req *productSvc.CreateReq) (*productSvc.CreateRes, error) {
	categoryID, err := primitive.ObjectIDFromHex(req.GetCategoryID())
	if err != nil {
		p.log.Errorf("primitive.ObjectIDFromHex: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	product := &models.Product{
		CategoryID:  categoryID,
		Name:        req.GetName(),
		Description: req.GetDescription(),
		Price:       req.GetPrice(),
		ImageURL:    &req.ImageURL,
		Photos:      req.GetPhotos(),
		Quantity:    req.GetQuantity(),
		Rating:      int(req.GetRating()),
	}

	if err := p.validate.StructCtx(ctx, product); err != nil {
		p.log.Errorf("validate.StructCtx: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	created, err := p.productUsecase.Create(ctx, product)
	if err != nil {
		p.log.Errorf("productUC.Create: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	return &productSvc.CreateRes{Product: created.ToProto()}, nil
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
