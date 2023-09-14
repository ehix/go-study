package database

import (
	"context"

	"github.com/ehix/go-microservices/internal/models"
)

func (c Client) GetAllProducts(ctx context.Context, vendorId string) ([]models.Products, error) {
	var products []models.Products
	// Pass ctx bc allows data pass to write handlers in gorm to do specific things.
	result := c.DB.WithContext(ctx).Where(models.Products{VendorID: vendorId}).Find(&products)
	// if you don't have an email, don't use it in the where clause, i.e. filter out.
	return products, result.Error
}
