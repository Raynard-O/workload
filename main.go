package main

import (
	"Proto/controller"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()
	//e.GET("/", controller.Proto)
	e.POST("/client", controller.Options)

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	//e.GET("/j", controller.J)
	e.Logger.Fatal(e.Start(":8080"))
}
