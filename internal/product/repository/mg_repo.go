package repository

import (
	"Go-ProductMS/internal/models"
	producterrors "Go-ProductMS/pkg/product_errors"
	"Go-ProductMS/pkg/util"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
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
	collection := p.mongoDB.Database(ProductDatabaseName).Collection(ProductCollectionName)

	upsert := true
	after := options.After
	opts := &options.FindOneAndUpdateOptions{
		ReturnDocument: &after,
		Upsert:         &upsert,
	}

	product.UpdatedAt = time.Now().UTC()
	var prod models.Product
	if err := collection.FindOneAndUpdate(ctx, bson.M{"_id": product.ProductID}, bson.M{"$set": product}, opts).Decode(&prod); err != nil {
		return nil, errors.Wrap(err, "Decode")
	}

	return &prod, nil
}

func (p *productMongoRepo) GetByID(ctx context.Context, productID primitive.ObjectID) (*models.Product, error) {
	collection := p.mongoDB.Database(ProductDatabaseName).Collection(ProductCollectionName)

	var prod models.Product
	if err := collection.FindOne(ctx, bson.M{"_id": productID}).Decode(&prod); err != nil {
		return nil, errors.Wrap(err, "Decode")
	}

	return &prod, nil
}

func (p *productMongoRepo) Search(ctx context.Context, search string, pagination *util.Pagination) (*models.ProductsList, error) {
	collection := p.mongoDB.Database(ProductDatabaseName).Collection(ProductCollectionName)

	filter := bson.D{
		{
			"$or",
			bson.A{
				bson.D{{"name", primitive.Regex{Pattern: search, Options: "gi"}}},
				bson.D{{"description", primitive.Regex{Pattern: search, Options: "gi"}}},
			},
		},
	}

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, errors.Wrap(err, "CountDocuments")
	}

	if count == 0 {
		return &models.ProductsList{
			TotalCount: 0,
			TotalPages: 0,
			Page:       0,
			Size:       0,
			Products:   []*models.Product{},
		}, nil
	}

	limit := int64(pagination.GetLimit())
	skip := int64(pagination.GetOffset())
	cursor, err := collection.Find(ctx, filter, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Find")
	}
	defer cursor.Close(ctx)

	products := make([]*models.Product, 0, pagination.GetSize())
	for cursor.Next(ctx) {
		var prod models.Product
		if err := cursor.Decode(&prod); err != nil {
			return nil, errors.Wrap(err, "Find")
		}
		products = append(products, &prod)
	}

	if err := cursor.Err(); err != nil {
		return nil, errors.Wrap(err, "cursor.Err")
	}

	return &models.ProductsList{
		TotalCount: count,
		TotalPages: int64(pagination.GetTotalPages(int(count))),
		Page:       int64(pagination.GetPage()),
		Size:       int64(pagination.GetSize()),
		HasMore:    pagination.GetHasMore(int(count)),
		Products:   products,
	}, nil
}
