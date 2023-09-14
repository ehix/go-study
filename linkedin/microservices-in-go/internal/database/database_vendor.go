package database

import (
	"context"
	"errors"

	"github.com/ehix/go-microservices/internal/database/dberrors"
	"github.com/ehix/go-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

func (c Client) GetAllVendors(ctx context.Context) ([]models.Vendors, error) {
	var vendors []models.Vendors
	result := c.DB.WithContext(ctx).Find(&vendors)
	// if you don't have an email, don't use it in the where clause, i.e. filter out.
	return vendors, result.Error
}

func (c Client) AddVendor(ctx context.Context, vendor *models.Vendors) (*models.Vendors, error) {
	vendor.VendorID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&vendor)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return vendor, nil
}

func (c Client) GetVendorById(ctx context.Context, ID string) (*models.Vendors, error) {
	vendor := &models.Vendors{}
	result := c.DB.WithContext(ctx).Where(&models.Vendors{VendorID: ID}).First((&vendor))
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "vendor", ID: ID}
		}
		return nil, result.Error
	}
	return vendor, nil
}
