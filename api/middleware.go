package api

import (
	"github.com/labstack/echo/v4"
)

func someMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return next(c)
	}
}
