package database

import (
	"context"
	"fmt"
	"time"

	"github.com/ehix/go-microservices/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

type DatabaseClient interface {
	Ready() bool

	GetAllCustomers(ctx context.Context, emailAddress string) ([]models.Customer, error)
	AddCustomer(ctx context.Context, customer *models.Customer) (*models.Customer, error)
	GetCustomerById(ctx context.Context, ID string) (*models.Customer, error)

	GetAllProducts(ctx context.Context, vendorId string) ([]models.Products, error)
	AddProduct(ctx context.Context, product *models.Products) (*models.Products, error)
	GetProductById(ctx context.Context, ID string) (*models.Products, error)

	GetAllServices(ctx context.Context) ([]models.Services, error)
	AddService(ctx context.Context, service *models.Services) (*models.Services, error)
	GetServiceById(ctx context.Context, ID string) (*models.Services, error)

	GetAllVendors(ctx context.Context) ([]models.Vendors, error)
	AddVendor(ctx context.Context, vendor *models.Vendors) (*models.Vendors, error)
	GetVendorById(ctx context.Context, ID string) (*models.Vendors, error)
}

type Client struct {
	DB *gorm.DB
}

func (c Client) Ready() bool {
	var ready string
	// Open a transaction to see the database
	tx := c.DB.Raw("SELECT 1 as ready").Scan(&ready)
	if tx.Error != nil {
		return false
	}
	if ready == "1" {
		return true
	}
	return false
}

// Create the client, would be a good idea to consume config here instead.
func NewDatabaseClient() (DatabaseClient, error) {
	// Used to connect to the DB
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=%s",
		"localhost", "postgres", "postgres", "postgres", 5432, "disable")

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "wisdom.",
		},
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		QueryFields: true,
	})

	if err != nil {
		return nil, err
	}

	client := Client{
		DB: db,
	}
	return client, nil
}
