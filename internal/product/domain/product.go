package domain

import (
	"Go-ProductMS/internal/models"
	"Go-ProductMS/pkg/util"
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoRepository Product
type MongoRepository interface {
	Create(ctx context.Context, product *models.Product) (*models.Product, error)
	Update(ctx context.Context, product *models.Product) (*models.Product, error)
	GetByID(ctx context.Context, productID primitive.ObjectID) (*models.Product, error)
	Search(ctx context.Context, search string, pagination *util.Pagination) (*models.ProductsList, error)
}
