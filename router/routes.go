package router

import (
	"Proto/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)
/* Routes*/
func New() *echo.Echo {
	e := echo.New()
	// post request for client data request ( for the hardcoded data set)
	e.POST("/clients", controller.Options)
	// post request for client data request
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
