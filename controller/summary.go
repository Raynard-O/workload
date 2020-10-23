package controller

import (
	"github.com/labstack/echo"
)

func Summary(c echo.Context) error {


	data, _ := DataSet2(c)

	//type test struct {
	//	Home string `json:"home"`
	//	School string `json:"school"`
	//}
	//
	//
	//testt := &test{
	//	Home:   "Montreal",
	//	School: "Concordia",
	//}
	return c.JSONPretty(201, &data, "")
}
