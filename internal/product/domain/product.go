package domain

import (
	"Go-ProductMS/internal/models"
	"context"
)

// MongoRepository Product
type MongoRepository interface {
	Create(ctx context.Context, product *models.Product) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) (*models.Product, error)
	GetByID(ctx context.Context, productID string) (*models.Product, error)
	Search(ctx context.Context, search string, page, size int64) ([]*models.Product, error)
}
