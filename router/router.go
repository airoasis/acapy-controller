package router

import (
	"github.com/airoasis/controller/handler"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
	"github.com/ziflex/lecho/v2"
)

func New() *echo.Echo {
	e := echo.New()
	e.Logger = lecho.From(log.Logger)
	e.Pre(middleware.RemoveTrailingSlash())
	//e.Use(middleware.Logger())
	e.Validator = NewValidator()

	e.POST("wallet", handler.CreateWallet)

	return e
}