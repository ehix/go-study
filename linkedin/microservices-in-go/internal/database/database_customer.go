package database

import (
	"context"

	"github.com/ehix/go-microservices/internal/models"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	// Pass ctx bc allows data pass to write handlers in gorm to do specific things.
	result := c.DB.WithContext(ctx).Where(models.Customer{Email: emailAddress}).Find(&customers)
	// if you don't have an email, don't use it in the where clause, i.e. filter out.
	return customers, result.Error
}
