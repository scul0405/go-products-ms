package repository

import (
	"Go-ProductMS/internal/models"
	producterrors "Go-ProductMS/pkg/product_errors"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	ProductDatabaseName   = "products"
	ProductCollectionName = "products"
)

// productMongoRepo
type productMongoRepo struct {
	mongoDB *mongo.Client
}

func NewProductMongoRepo(mongoDB *mongo.Client) *productMongoRepo {
	return &productMongoRepo{
		mongoDB: mongoDB,
	}
}

// Create creates a new product
func (p *productMongoRepo) Create(ctx context.Context, product *models.Product) (*models.Product, error) {
	collection := p.mongoDB.Database(ProductDatabaseName).Collection(ProductCollectionName)

	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	result, err := collection.InsertOne(ctx, product, &options.InsertOneOptions{})
	if err != nil {
		return nil, errors.Wrap(err, "error while inserting product")
	}

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, errors.Wrap(producterrors.ErrObjectIDTypeConversion, "error while converting object id")
	}

	product.ProductID = objectID

	return product, nil
}

func (p *productMongoRepo) Update(ctx context.Context, product *models.Product) (*models.Product, error) {
	panic("implement me")
}

func (p *productMongoRepo) GetByID(ctx context.Context, productID primitive.ObjectID) (*models.Product, error) {
	panic("implement me")
}

func (p *productMongoRepo) Search(ctx context.Context, search string, page, size int64) ([]*models.Product, error) {
	panic("implement me")
}
