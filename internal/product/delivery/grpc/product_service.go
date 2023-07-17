package grpc

import (
	"Go-ProductMS/internal/models"
	"Go-ProductMS/internal/product/usecase"
	grpcErrors "Go-ProductMS/pkg/grpc_errors"
	"Go-ProductMS/pkg/logger"
	"Go-ProductMS/pkg/util"
	productSvc "Go-ProductMS/proto/product"
	"context"
	"github.com/go-playground/validator/v10"
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type productService struct {
	log            logger.Logger
	productUsecase usecase.ProductsUseCase
	validate       *validator.Validate
	productSvc.UnimplementedProductsServiceServer
}

func NewProductService(
	log logger.Logger,
	productUsecase usecase.ProductsUseCase,
	validate *validator.Validate,
) *productService {
	return &productService{
		log:            log,
		productUsecase: productUsecase,
		validate:       validate,
	}
}

func (p *productService) Create(ctx context.Context, req *productSvc.CreateReq) (*productSvc.CreateRes, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productService.Create")
	defer span.Finish()

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
	span, ctx := opentracing.StartSpanFromContext(ctx, "productService.Update")
	defer span.Finish()

	productID, err := primitive.ObjectIDFromHex(req.GetProductID())
	if err != nil {
		p.log.Errorf("primitive.ObjectIDFromHex: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	categoryID, err := primitive.ObjectIDFromHex(req.GetCategoryID())
	if err != nil {
		p.log.Errorf("primitive.ObjectIDFromHex: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	product := &models.Product{
		ProductID:   productID,
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

	update, err := p.productUsecase.Update(ctx, product)
	if err != nil {
		p.log.Errorf("productUC.Update: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	return &productSvc.UpdateRes{Product: update.ToProto()}, nil
}

func (p *productService) GetByID(ctx context.Context, req *productSvc.GetByIDReq) (*productSvc.GetByIDRes, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productService.GetByID")
	defer span.Finish()

	productID, err := primitive.ObjectIDFromHex(req.GetProductID())
	if err != nil {
		p.log.Errorf("primitive.ObjectIDFromHex: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	product, err := p.productUsecase.GetByID(ctx, productID)
	if err != nil {
		p.log.Errorf("productUC.GetByID: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	return &productSvc.GetByIDRes{Product: product.ToProto()}, nil
}

func (p *productService) Search(ctx context.Context, req *productSvc.SearchReq) (*productSvc.SearchRes, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productService.Search")
	defer span.Finish()

	pagination := util.NewPagination(int(req.GetPage()), int(req.GetSize()))

	products, err := p.productUsecase.Search(ctx, req.GetSearch(), pagination)
	if err != nil {
		p.log.Errorf("productUC.Search: %v", err)
		return nil, grpcErrors.ErrorResponse(err, err.Error())
	}

	p.log.Infof("PRODUCTS: %-v", products)

	return &productSvc.SearchRes{
		Products:   products.ToProtoList(),
		TotalCount: products.TotalCount,
		TotalPages: products.TotalPages,
		Page:       products.Page,
		Size:       products.Size,
		HasMore:    products.HasMore,
	}, nil
}
