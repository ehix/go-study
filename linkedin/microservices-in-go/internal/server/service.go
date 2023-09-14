package server

import (
	"net/http"

	"github.com/ehix/go-microservices/internal/database/dberrors"
	"github.com/ehix/go-microservices/internal/models"
	"github.com/labstack/echo/v4"
)

func (s *EchoServer) GetAllServices(ctx echo.Context) error {
	services, err := s.DB.GetAllServices(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err)
	}
	return ctx.JSON(http.StatusOK, services)
}

func (s *EchoServer) AddService(ctx echo.Context) error {
	service := new(models.Service) // get a pointer to service
	// built-in to bind a body to an object, checking if we can do it.
	if err := ctx.Bind(service); err != nil {
		return ctx.JSON(http.StatusUnsupportedMediaType, err)
	}
	service, err := s.DB.AddService(ctx.Request().Context(), service)
	if err != nil {
		switch err.(type) {
		case *dberrors.ConflictError:
			return ctx.JSON(http.StatusConflict, err)
		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusCreated, service)
}

func (s *EchoServer) GetServiceById(ctx echo.Context) error {
	// Pull the value for ID from the root URL
	ID := ctx.Param("id")
	service, err := s.DB.GetServiceById(ctx.Request().Context(), ID)
	if err != nil {
		switch err.(type) {
		case *dberrors.NotFoundError:
			return ctx.JSON(http.StatusNotFound, err)

		default:
			return ctx.JSON(http.StatusInternalServerError, err)
		}
	}
	return ctx.JSON(http.StatusOK, service)
}
