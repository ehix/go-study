package database

import (
	"context"

	"github.com/ehix/go-microservices/internal/models"
)

func (c Client) GetAllServices(ctx context.Context) ([]models.Services, error) {
	var services []models.Services
	result := c.DB.WithContext(ctx).Find(&services)
	// if you don't have an email, don't use it in the where clause, i.e. filter out.
	return services, result.Error
}
