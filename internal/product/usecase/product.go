package usecase

import (
	"Go-ProductMS/internal/models"
	"Go-ProductMS/internal/product/domain"
	"Go-ProductMS/pkg/logger"
	"Go-ProductMS/pkg/util"
	"context"
	"github.com/opentracing/opentracing-go"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProductUseCase interface {
	Create(ctx context.Context, product *models.Product) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) (*models.Product, error)
	GetByID(ctx context.Context, productID primitive.ObjectID) (*models.Product, error)
	Search(ctx context.Context, search string, pagination *util.Pagination) (*models.ProductsList, error)
}

type productUsecase struct {
	productRepo domain.MongoRepository
	log         logger.Logger
}

func NewProductUsecase(
	productRepo domain.MongoRepository,
	log logger.Logger,
) *productUsecase {
	return &productUsecase{
		productRepo: productRepo,
		log:         log,
	}
}

func (p *productUsecase) Create(ctx context.Context, product *models.Product) (*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productUsecase.Create")
	defer span.Finish()

	return p.productRepo.Create(ctx, product)
}

func (p *productUsecase) Update(ctx context.Context, product *models.Product) (*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productUsecase.Update")
	defer span.Finish()

	return p.productRepo.Update(ctx, product)
}

func (p *productUsecase) GetByID(ctx context.Context, productID primitive.ObjectID) (*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productUsecase.GetByID")
	defer span.Finish()

	return p.productRepo.GetByID(ctx, productID)
}

func (p *productUsecase) Search(ctx context.Context, search string, pagination *util.Pagination) (*models.ProductsList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productUsecase.Search")
	defer span.Finish()

	return p.productRepo.Search(ctx, search, pagination)
}
