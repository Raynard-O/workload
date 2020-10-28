package controller

import (
	grpc_from0 "Proto/github.com/monkrus/grpc-from0"

	"github.com/labstack/echo"
	"net/http"
)





func EncodeJson(c echo.Context, RfwId string, LastBatchID int32, samplesWorK []*grpc_from0.Batch) error {
	RFW := &grpc_from0.RFD{
		RFWID:       RfwId,
		LastBatchId: LastBatchID,
		Batches:     samplesWorK,
	}

	return c.JSONPretty(http.StatusOK, RFW, "")
}

