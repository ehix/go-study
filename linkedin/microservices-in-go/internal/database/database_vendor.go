package database

import (
	"context"

	"github.com/ehix/go-microservices/internal/models"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendors, error) {
	var vendors []models.Vendors
	result := c.DB.WithContext(ctx).Find(&vendors)
	// if you don't have an email, don't use it in the where clause, i.e. filter out.
	return vendors, result.Error
}
