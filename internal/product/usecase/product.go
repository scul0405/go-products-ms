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

type ProductsUseCase interface {
	Create(ctx context.Context, product *models.Product) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) (*models.Product, error)
	GetByID(ctx context.Context, productID primitive.ObjectID) (*models.Product, error)
	Search(ctx context.Context, search string, pagination *util.Pagination) (*models.ProductsList, error)
}

type productsUsecase struct {
	productRepo domain.MongoRepository
	log         logger.Logger
}

func NewProductUsecase(
	productRepo domain.MongoRepository,
	log logger.Logger,
) *productsUsecase {
	return &productsUsecase{
		productRepo: productRepo,
		log:         log,
	}
}

func (p *productsUsecase) Create(ctx context.Context, product *models.Product) (*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsUsecase.Create")
	defer span.Finish()

	return p.productRepo.Create(ctx, product)
}

func (p *productsUsecase) Update(ctx context.Context, product *models.Product) (*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsUsecase.Update")
	defer span.Finish()

	return p.productRepo.Update(ctx, product)
}

func (p *productsUsecase) GetByID(ctx context.Context, productID primitive.ObjectID) (*models.Product, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsUsecase.GetByID")
	defer span.Finish()

	return p.productRepo.GetByID(ctx, productID)
}

func (p *productsUsecase) Search(ctx context.Context, search string, pagination *util.Pagination) (*models.ProductsList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "productsUsecase.Search")
	defer span.Finish()

	return p.productRepo.Search(ctx, search, pagination)
}
