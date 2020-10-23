package controller

import "github.com/labstack/echo"

func Summary(c echo.Context) error {

	type test struct {
		Home string `json:"home"`
		School string `json:"school"`
	}


	testt := &test{
		Home:   "Montreal",
		School: "Concordia",
	}
	return c.JSONPretty(201, &testt, "")
}
