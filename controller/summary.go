package controller

import (
	"Proto/library"
	"github.com/labstack/echo"
)

func Summary(c echo.Context) error {

	workload := library.DataParamsRequest{

		BenchmarkType:  "NDBENCH",
		WorkloadMetric: "CPU",
		BatchUnit:      2,
		BatchID:         5,
		BatchSize:      0,
	}
	data, _ := DataSet(c)

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
