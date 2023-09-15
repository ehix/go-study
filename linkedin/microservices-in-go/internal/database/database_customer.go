package database

import (
	"context"
	"errors"

	"github.com/ehix/go-microservices/internal/database/dberrors"
	"github.com/ehix/go-microservices/internal/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func (c Client) GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error) {
	var customers []models.Customer
	// Pass ctx bc allows data pass to write handlers in gorm to do specific things.
	result := c.DB.WithContext(ctx).Where(models.Customer{Email: emailAddress}).Find(&customers)
	// if you don't have an email, don't use it in the where clause, i.e. filter out.
	return customers, result.Error
}

func (c Client) AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	// create a new uuid
	customer.CustomerID = uuid.NewString()
	result := c.DB.WithContext(ctx).Create(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			// we could fix rather than return
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	return customer, nil
}

func (c Client) GetCustomerById(ctx context.Context, ID string) (*models.Customer, error) {
	customer := &models.Customer{}
	result := c.DB.WithContext(ctx).Where(&models.Customer{CustomerID: ID}).First(&customer)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, &dberrors.NotFoundError{Entity: "customer", ID: ID}
		}
		return nil, result.Error
	}
	return customer, nil
}

func (c Client) UpdateCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error) {
	// create a slice instead
	var customers []models.Customer
	// Clause: return obj with any update that we update with any changes triggered in DB.
	// Will guard against updates to customer ID in web-layer.
	// Think about what we'll allow to be updated, allow the entire model?
	result := c.DB.WithContext(ctx).
		Clauses(clause.Returning{}).
		Where(&models.Customer{CustomerID: customer.CustomerID}).
		Updates(models.Customer{
			FirstName: customer.FirstName,
			LastName:  customer.LastName,
			Email:     customer.Email,
			Phone:     customer.Phone,
			Address:   customer.Address,
		})
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, &dberrors.ConflictError{}
		}
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, &dberrors.NotFoundError{Entity: "customer", ID: customer.CustomerID}
	}
	return &customers[0], nil
}

func (c Client) DeleteCustomer(ctx context.Context, ID string) error {
	return c.DB.WithContext(ctx).Delete(&models.Customer{CustomerID: ID}).Error
}
