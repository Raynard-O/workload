package router

import (
	"Proto/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New() *echo.Echo {
	e := echo.New()

	e.POST("/clients", controller.Options)

	e.POST("/client", controller.Client2)
	e.GET("/", func(context echo.Context) error {
		return context.JSON(200, "Home")
	})
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))


	return e
}
