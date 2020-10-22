package main

import (
	"Proto/controller"
	"github.com/labstack/echo"
)

func main() {

	e := echo.New()
	//e.GET("/", controller.Proto)
	e.GET("/p", controller.P)
	e.Logger.Fatal(e.Start(":8080"))
}
