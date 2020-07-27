package services

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type (
	model interface {
		GetHelloWorldString() string
	}

	service struct {
		model model
	}
)

func NewServices(m model) *service {
	return &service{m}
}

func (s *service) ShowHelloWorld(c echo.Context) error {

	res := s.model.GetHelloWorldString()
	return c.JSON(http.StatusOK, res)
}
