package server

import (
	"log"
	"net/http"

	"github.com/ehix/go-microservices/internal/database"
	"github.com/ehix/go-microservices/internal/models"
	"github.com/labstack/echo/v4"
)

type Server interface {
	Start() error
	// Have these on k8s
	Readiness(ctx echo.Context) error
	Liveness(ctx echo.Context) error
	// Get all methods
	GetAllCustomers(ctx echo.Context) error
	GetAllProducts(ctx echo.Context) error
	GetAllServices(ctx echo.Context) error
	GetAllVendors(ctx echo.Context) error
}

type EchoServer struct {
	echo *echo.Echo
	DB   database.DatabaseClient
}

// This is set up for testing, being able to inject data into other elements, like a mock.
func NewEchoServer(db database.DatabaseClient) Server {
	server := &EchoServer{
		echo: echo.New(),
		DB:   db,
	}
	server.registerRoutes()
	return server
}

// Info passed usually implemented through config.
func (s *EchoServer) Start() error {
	if err := s.echo.Start(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatalf("server shutdown occurred %s", err)
		return err // despite fatal
	}
	return nil
}

// Groups for specific elements
func (s *EchoServer) registerRoutes() {
	// mimics k8s, bc we're building for k8s
	s.echo.GET("/readiness", s.Readiness)
	s.echo.GET("/liveness", s.Liveness)

	// create customer group
	cg := s.echo.Group("/customers")
	cg.GET("", s.GetAllCustomers)

	// create product group
	pg := s.echo.Group("/products")
	pg.GET("", s.GetAllProducts)

	// create services group
	sg := s.echo.Group("/services")
	sg.GET("", s.GetAllServices)

	// create vendors group
	vg := s.echo.Group("/vendors")
	vg.GET("", s.GetAllVendors)

}

func (s *EchoServer) Readiness(ctx echo.Context) error {
	ready := s.DB.Ready()
	if ready {
		return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
	}
	return ctx.JSON(http.StatusInternalServerError, models.Health{Status: "Failure"})
}

// Ping page, is everything up and running?
func (s *EchoServer) Liveness(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, models.Health{Status: "OK"})
}
