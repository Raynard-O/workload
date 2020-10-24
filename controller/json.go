package controller

import (
	grpc_from0 "Proto/github.com/monkrus/grpc-from0"
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func Options(c echo.Context) error {



	data, workload := DataSet(c)

	var size int = workload.BatchUnit
	var unit int = workload.BatchID
	bSize := workload.BatchSize
	total := size * unit
	totalT := total - 1
	result := data[bSize-total : bSize]

	var Bench int8
	switch workload.WorkloadMetric {
	case "CPU":
		Bench = 0
		//serverRFW.Benchmark_Type = "CPUUtilization_Average"
	case "NETIN":
		Bench = 1
		//serverRFW.Benchmark_Type = "NetworkIn_Average"
	case "NETOUT":
		Bench = 2
		//serverRFW.Benchmark_Type = "NetworkOut_Average"
	case "MEMUTI":
		Bench = 3
		//serverRFW.Benchmark_Type = "MemoryUtilization_Average"
	default:
		Bench = 4
		//serverRFW.Benchmark_Type = "Final_Target"
	}

	fmt.Printf("Workload Metrics : %d", Bench)

	var batch3 []*grpc_from0.Batch

	ID := 1
	for i := 0; i <= totalT; {

		z := total-i
		b := total-i-size
		test := result[b: z]
		var sam []*grpc_from0.Sample
		for k := size-1 ; k >= 0; k-- {
			if Bench == 0 {
				samp := &grpc_from0.Sample{
					CPUUtilization: test[k].CpuUtilizationAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 1 {
				samp := &grpc_from0.Sample{
					NetworkIN: test[k].NetworkInAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 2 {
				samp := &grpc_from0.Sample{
					NetworkOUT: test[k].NetworkOutAverage,
				}
				sam = append(sam, samp)
			}else if Bench == 3 {
				samp := &grpc_from0.Sample{
					MemoryUtilization: test[k].MemoryUtilizationAverage,
				}
				sam = append(sam, samp)
			}else {
				samp := &grpc_from0.Sample{
					MemoryUtilization: test[k].MemoryUtilizationAverage,
				}
				sam = append(sam, samp)
			}


		}
		batch := &grpc_from0.Batch{
			Batch_ID: int32(ID),
			Samples:  sam,
		}
		//log.Println(test)
		batch3 = append(batch3, batch)
		i += size
		ID++
	}
	if workload.BinarySerialization == "binary" {
		return Proto(c, workload.RFWID, int32(workload.BatchID), batch3)
	}

	return EncodeJson(c, workload.RFWID, int32(workload.BatchID), batch3)

}

//func J(c echo.Context) error {
//
//	data, workload := DataSet(c)
//	//size := strconv.Itoa(int(workload.Batch_Size))
//
//	var Bench int8
//	switch workload.WorkloadMetric {
//	case "CPU":
//		Bench = 0
//		//serverRFW.Benchmark_Type = "CPUUtilization_Average"
//	case "NETIN":
//		Bench = 1
//		//serverRFW.Benchmark_Type = "NetworkIn_Average"
//	case "NETOUT":
//		Bench = 2
//		//serverRFW.Benchmark_Type = "NetworkOut_Average"
//	case "MEMUTI":
//		Bench = 3
//		//serverRFW.Benchmark_Type = "MemoryUtilization_Average"
//	default:
//		Bench = 4
//		//serverRFW.Benchmark_Type = "Final_Target"
//	}
//
//	fmt.Println(Bench)
//
//	batchSize := (workload.BatchSize / workload.BatchUnit) - 1
//	var batch3 []*library.Batch
//
//	for i := batchSize; i > (batchSize)-workload.BatchID; i-- {
//		//fmt.Println(i)
//		f := batchSize - i
//
//		var sam []*library.Sample
//
//		c := i * workload.BatchUnit
//
//		for j := c; j < c+workload.BatchUnit; j++ {
//
//			switch Bench {
//			case 0:
//				samp := &library.Sample{
//					CpuUtilization: data[j].CpuUtilizationAverage,
//				}
//				sam = append(sam, samp)
//			case 1 :
//				samp := &library.Sample{
//					NetworkIN: data[j].NetworkInAverage,
//				}
//				sam = append(sam, samp)
//			case 2:
//				samp := &library.Sample{
//					NetworkOUT: data[j].NetworkOutAverage,
//				}
//				sam = append(sam, samp)
//			case 3:
//				samp := &library.Sample{
//					MemoryUtilization: data[j].MemoryUtilizationAverage,
//				}
//				sam = append(sam, samp)
//			default:
//				samp := &library.Sample{
//					FinalTarget: data[j].FinalTarget,
//				}
//				sam = append(sam, samp)
//			}
//
//		}
//
//		batch := &library.Batch{
//			Batch_ID: f + 1,
//			Samples:  sam,
//		}
//		batch3 = append(batch3, batch)
//	}
//	return  EncodeJson(c, workload.RFWID, workload.BatchID, batch3)
//
//}

func EncodeJson(c echo.Context, RfwId string, LastBatchID int32, samplesWorK []*grpc_from0.Batch) error {
	RFW := &grpc_from0.RFD{
		RFWID:       RfwId,
		LastBatchId: LastBatchID,
		Batches:     samplesWorK,
	}

	return c.JSONPretty(http.StatusOK, RFW, "")
}

