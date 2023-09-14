package database

import (
	"context"
	"errors"

	"github.com/ehix/go-microservices/internal/database/dberrors"
	"github.com/ehix/go-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c Client) GetAllProducts(ctx context.Context, vendorId string) ([]models.Products, error) {
	var products []models.Products
	// Pass ctx bc allows data pass to write handlers in gorm to do specific things.
	result := c.DB.WithContext(ctx).Where(models.Products{VendorID: vendorId}).Find(&products)
	// if you don't have an email, don't use it in the where clause, i.e. filter out.
	return products, result.Error
}

func (c Client) AddProduct(ctx context.Context, product *models.Products) (*models.Products, error) {
	product.ProductID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&product)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return product, nil
}
