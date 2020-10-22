package binary

import (
	grpc_from0 "Proto/github.com/monkrus/grpc-from0"
	"fmt"
	"github.com/golang/protobuf/proto"
	"log"
)

func EncodeProto(RFW_ID string, LastBatchID int32, samplesWorK []*grpc_from0.Batch) ([]byte, error) {

	samples := &grpc_from0.RFD{
		RFWID:       RFW_ID,
		LastBatchId: LastBatchID,
		Batches:     samplesWorK,
	}

	data, err := proto.Marshal(samples)
	if err != nil {
		log.Fatal("marshaling error: ", err)
	}
	// printing out our raw protobuf object
	fmt.Println(data)
	return data, err
}

func DecodeProto(data []byte, file *grpc_from0.RFD) {

	err := proto.Unmarshal(data, file)
	if err != nil {
		log.Fatal("unmarshall error: ", err)
	}

}
