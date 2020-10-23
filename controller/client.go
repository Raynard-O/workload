package controller

import (
	"Proto/library"
	"github.com/labstack/echo"
	"log"
)

func Options(ctx echo.Context) error {
	workload := new(library.DataParamsRequest)
	if err := ctx.Bind(workload); err != nil {
		log.Fatal(err)
	}
	if workload.BinarySerialization == "binary"{
		return P(ctx)
	}
	return J(ctx)
}
