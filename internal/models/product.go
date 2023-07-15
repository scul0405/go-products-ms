package models

import (
	productService "Go-ProductMS/proto/product"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

type Product struct {
	ProductID   primitive.ObjectID `json:"productId" bson:"_id,omitempty"`
	CategoryID  primitive.ObjectID `json:"categoryId" bson:"categoryId,omitempty"`
	Name        string             `json:"name" bson:"name,omitempty" validate:"required"`
	Description string             `json:"description" bson:"description,omitempty" validate:"required"`
	Price       float64            `json:"price" bson:"price,omitempty" validate:"required"`
	ImageURL    *string            `json:"imageUrl" bson:"imageUrl,omitempty"`
	Photos      []string           `json:"photos" bson:"photos,omitempty"`
	Quantity    int64              `json:"quantity" bson:"quantity" validate:"required"`
	Rating      int                `json:"rating" bson:"rating" validate:"required"`
	CreatedAt   time.Time          `json:"createdAt" bson:"createdAt,omitempty"`
	UpdatedAt   time.Time          `json:"updatedAt" bson:"updatedAt,omitempty"`
}

func (p *Product) GetImage() string {
	var img string
	if p.ImageURL != nil {
		img = *p.ImageURL
	}
	return img
}

func (p *Product) ToProto() *productService.Product {
	return &productService.Product{
		ID:          p.ProductID.String(),
		CategoryID:  p.CategoryID.String(),
		Name:        p.Name,
		Description: p.Description,
		Price:       p.Price,
		ImageURL:    p.GetImage(),
		Photos:      p.Photos,
		Quantity:    p.Quantity,
		Rating:      int32(p.Rating),
		Created_At:  timestamppb.New(p.CreatedAt),
		Updated_At:  timestamppb.New(p.UpdatedAt),
	}
}

// ProductsList All Products response with pagination
type ProductsList struct {
	TotalCount int64      `json:"totalCount"`
	TotalPages int64      `json:"totalPages"`
	Page       int64      `json:"page"`
	Size       int64      `json:"size"`
	HasMore    bool       `json:"hasMore"`
	Products   []*Product `json:"products"`
}

// ToProtoList convert products list to proto
func (p *ProductsList) ToProtoList() []*productService.Product {
	productsList := make([]*productService.Product, 0, len(p.Products))
	for _, product := range p.Products {
		productsList = append(productsList, product.ToProto())
	}
	return productsList
}
