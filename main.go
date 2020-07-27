package main

import (
	"echo-framework-structure/api"
	"echo-framework-structure/config"
	"echo-framework-structure/models"
	"echo-framework-structure/services"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	c := config.GetInstance()

	m := models.NewModels(c)
	s := services.NewServices(m)

	api.Route(e, s)
	e.Logger.Fatal(e.Start(":3003"))
}
