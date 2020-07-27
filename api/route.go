package api

import "github.com/labstack/echo/v4"

type (
	services interface {
		ShowHelloWorld(c echo.Context) error
	}
)

func Route(e *echo.Echo, s services) {
	e.Use(someMiddleware)

	e.GET("/", s.ShowHelloWorld)

}
