package controller

import (
	"Proto/github.com/binary"
	grpcfrom0 "Proto/github.com/monkrus/grpc-from0"
	"fmt"
	"github.com/labstack/echo"
	"log"
)


func Proto(c echo.Context, RFWID string, LASTBATCHID int32, batch []*grpcfrom0.Batch) error {

	data, err := binary.EncodeProto(RFWID, LASTBATCHID, batch)

	if err != nil {
		log.Fatal("marshaling error: ", err)
	}

	file := &grpcfrom0.RFD{}
	binary.DecodeProto(data, file)
	fmt.Print(file)


	return c.JSONBlob(200, data)
}
