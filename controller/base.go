package controller

import (
	"Proto/github.com/binary"
	grpc_from0 "Proto/github.com/monkrus/grpc-from0"
	"fmt"
	"github.com/labstack/echo"
	"log"
	"net/http"
)

// Response Struct
type Response struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}

// Data Struct
type Data struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// PaginatedData Struct
type PaginatedData struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
	Count   int64       `json:"count"`
	Pages   int64       `json:"pages"`
}

// Error Struct
type Error struct {
	Success bool        `json:"success"`
	Message interface{} `json:"message"`
}

func InternalError(c echo.Context, message string) error {
	return c.JSONPretty(http.StatusInternalServerError, Error{
		Message: message,
		Success: false,
	}, " ")
}


func BadRequestResponse(c echo.Context, message string) error {
	return c.JSONPretty(http.StatusBadRequest, Error{
		Message: message,
		Success: false,
	}, " ")
}

func DataResponse(c echo.Context, status int, data interface{}) error {
	return c.JSONPretty(status, Data{
		Data:    data,
		Success: true,
	}, " ")
}
func MessageResponse(c echo.Context, status int, message string) error {
	return c.JSONPretty(status, Error{
		Message: message,
		Success: status <= 201,
	}, " ")
}

func InvalidResponse(c echo.Context) error {
	return BadRequestResponse(c, "Invalid Request")
}



func EncodeJson(c echo.Context, RfwId string, LastBatchID int32, samplesWorK []*grpc_from0.Batch) error {
	RFW := &grpc_from0.RFD{
		RFWID:       RfwId,
		LastBatchId: LastBatchID,
		Batches:     samplesWorK,
	}

	return DataResponse(c, http.StatusAccepted, RFW)

}


func Proto(c echo.Context, RFWID string, LASTBATCHID int32, batch []*grpc_from0.Batch) error {

	data, err := binary.EncodeProto(RFWID, LASTBATCHID, batch)

	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	file := &grpc_from0.RFD{}
	binary.DecodeProto(data, file)
	fmt.Print(file)


	return c.JSONBlob(200, data)
}
