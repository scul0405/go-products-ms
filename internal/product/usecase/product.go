package usecase

import (
	"Go-ProductMS/internal/models"
	"Go-ProductMS/internal/product/domain"
	"Go-ProductMS/pkg/logger"
	"context"
)

type ProductUseCase interface {
	Create(ctx context.Context, product *models.Product) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) (*models.Product, error)
	GetByID(ctx context.Context, productID string) (*models.Product, error)
	Search(ctx context.Context, search string, page, size int64) ([]*models.Product, error)
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
	panic("implement me")
}

func (p *productUsecase) Update(ctx context.Context, product *models.Product) (*models.Product, error) {
	panic("implement me")
}

func (p *productUsecase) GetByID(ctx context.Context, productID string) (*models.Product, error) {
	panic("implement me")
}

func (p *productUsecase) Search(ctx context.Context, search string, page, size int64) ([]*models.Product, error) {
	panic("implement me")
}
